/*
Copyright © 2025 Benny Powers

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package generate

import (
	"cmp"
	"errors"
	"fmt"
	"runtime"
	"slices"
	"sync"

	"github.com/pterm/pterm"

	DT "bennypowers.dev/cem/designtokens"
	DD "bennypowers.dev/cem/generate/demodiscovery"
	Q "bennypowers.dev/cem/generate/queries"
	M "bennypowers.dev/cem/manifest"
	W "bennypowers.dev/cem/workspace"

	"maps"

	DS "github.com/bmatcuk/doublestar"
	ts "github.com/tree-sitter/go-tree-sitter"
)

var defaultExcludePatterns = []string{
	"**/*.d.ts",
}

// Checks whether a file matches any of the given patterns (with doublestar)
func matchesAnyPattern(file string, patterns []string) bool {
	for _, pattern := range patterns {
		match, err := DS.PathMatch(pattern, file)
		if err == nil && match {
			return true
		}
	}
	return false
}

type preprocessResult struct {
	demoFiles       []string
	designTokens    *DT.DesignTokens
	excludePatterns []string
	includedFiles   []string
}

// preprocess handles config merging for generate command
func preprocess(ctx W.WorkspaceContext) (r preprocessResult, errs error) {
	cfg, err := ctx.Config()
	if err != nil {
		return r, err
	}
	cfg.Generate.Exclude = append(cfg.Generate.Exclude, []string{}...)
	r.excludePatterns = make([]string, 0, len(cfg.Generate.Exclude)+len(defaultExcludePatterns))

	r.includedFiles = make([]string, 0)
	r.excludePatterns = append(r.excludePatterns, cfg.Generate.Exclude...)

	if !cfg.Generate.NoDefaultExcludes {
		r.excludePatterns = append(r.excludePatterns, defaultExcludePatterns...)
	}

	if cfg.Generate.DesignTokens.Spec != "" {
		tokens, err := DT.LoadDesignTokens(ctx)
		if err != nil {
			errs = errors.Join(errs, err)
		}
		r.designTokens = tokens
	}
	if cfg.Generate.DemoDiscovery.FileGlob != "" {
		demoFiles, err := ctx.Glob(cfg.Generate.DemoDiscovery.FileGlob)
		r.demoFiles = demoFiles
		if err != nil {
			errs = errors.Join(errs, err)
		}
	}
	for _, file := range cfg.Generate.Files {
		if !matchesAnyPattern(file, r.excludePatterns) {
			r.includedFiles = append(r.includedFiles, file)
		}
	}
	return r, errs
}

type processJob struct {
	file string
	ctx  W.WorkspaceContext
}

// process actually processes a module file
func process(
	ctx W.WorkspaceContext,
	result preprocessResult,
	qm *Q.QueryManager,
) (modules []M.Module, logs []*LogCtx, aliases map[string]string, errs error) {
	numWorkers := runtime.NumCPU()
	pterm.Info.Printf("Starting Generation with %d workers\n", numWorkers)

	var wg sync.WaitGroup
	var aliasesMu sync.Mutex
	var modulesMu sync.Mutex
	var errsMu sync.Mutex
	var logsMu sync.Mutex
	errsList := make([]error, 0)
	logs = make([]*LogCtx, 0, len(result.includedFiles))

	aliases = make(map[string]string) // Ensure aliases is initialized
	jobsChan := make(chan processJob, len(result.includedFiles))

	// Fill jobs channel with files to process
	for _, file := range result.includedFiles {
		jobsChan <- processJob{file: file, ctx: ctx}
	}
	close(jobsChan)

	wg.Add(numWorkers)
	for range numWorkers {
		go func() {
			defer wg.Done()
			parser := Q.GetTypeScriptParser()
			defer Q.PutTypeScriptParser(parser)
			for job := range jobsChan {
				module, tagAliases, logger, err := processModule(job, qm, parser)
				if err != nil {
					errsMu.Lock()
					errsList = append(errsList, err)
					errsMu.Unlock()
				}

				// Save log for later bar chart (always save duration for bar chart)
				logsMu.Lock()
				logs = append(logs, logger)
				logsMu.Unlock()

				// Write to aliases in a threadsafe manner
				aliasesMu.Lock()
				maps.Copy(aliases, tagAliases)
				aliasesMu.Unlock()
				modulesMu.Lock()
				if module != nil {
					modules = append(modules, *module)
				}
				modulesMu.Unlock()
			}
		}()
	}

	wg.Wait()

	if len(errsList) > 0 {
		errs = errors.Join(errsList...)
	}

	return modules, logs, aliases, errs
}

func processModule(
	job processJob,
	qm *Q.QueryManager,
	parser *ts.Parser,
) (module *M.Module, tagAliases map[string]string, logCtx *LogCtx, errs error) {
	defer parser.Reset()
	mp, err := NewModuleProcessor(job.ctx, job.file, parser, qm)
	if err != nil {
		return nil, nil, nil, err
	}
	defer mp.Close()
	cfg, err := job.ctx.Config()
	if err != nil {
		return nil, nil, nil, err
	}
	if cfg.Verbose {
		mp.logger.Section.Printf("Module: %s", mp.logger.File)
	}
	module, tagAliases, err = mp.Collect()
	return module, tagAliases, mp.logger, err
}

func postprocess(
	ctx W.WorkspaceContext,
	result preprocessResult,
	allTagAliases map[string]string,
	qm *Q.QueryManager,
	modules []M.Module,
) (pkg M.Package, errs error) {
	var wg sync.WaitGroup
	var errsMu sync.Mutex
	errsList := make([]error, 0)

	// Build the demo map once
	demoMap, err := DD.NewDemoMap(result.demoFiles)
	if err != nil {
		errsList = append(errsList, err)
	}

	// Because demo discovery and design tokens may mutate modules, we need to coordinate by pointer
	for i := range modules {
		wg.Add(1)
		go func(module *M.Module) {
			defer wg.Done()
			if result.designTokens != nil {
				DT.MergeDesignTokensToModule(module, *result.designTokens)
			}
			// Discover demos and attach to manifest
			if len(demoMap) > 0 {
				err := DD.DiscoverDemos(ctx, allTagAliases, module, qm, demoMap)
				if err != nil {
					errsMu.Lock()
					errsList = append(errsList, err)
					errsMu.Unlock()
				}
			}
		}(&modules[i])
	}
	wg.Wait()
	// No need to close a channel here, just return
	pkg = M.NewPackage(modules)
	slices.SortStableFunc(pkg.Modules, func(a, b M.Module) int {
		return cmp.Compare(a.Path, b.Path)
	})
	if len(errsList) > 0 {
		errs = errors.Join(errsList...)
	}
	return pkg, errs
}

// Generates a custom-elements manifest from a list of typescript files
func Generate(ctx W.WorkspaceContext) (manifest *string, errs error) {
	cfg, err := ctx.Config()
	if err != nil {
		return nil, err
	}
	qm, err := Q.NewQueryManager()
	if err != nil {
		return nil, err
	}
	defer qm.Close()
	result, err := preprocess(ctx)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("module preprocess failed:\n%w", err))
	}
	modules, logs, aliases, err := process(ctx, result, qm)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("module process failed:\n%w", err))
	}
	pkg, err := postprocess(ctx, result, aliases, qm, modules)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("module postprocess failed:\n%w", err))
	}
	if cfg.Verbose {
		RenderBarChart(logs)
	}
	manifestStr, err := M.SerializeToString(&pkg)
	if err != nil {
		return nil, errors.Join(errs, fmt.Errorf("module serialize failed:\n%w", err))
	}
	return &manifestStr, errs
}
