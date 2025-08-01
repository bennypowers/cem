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
	"context"
	"errors"
	"fmt"
	"slices"
	"sync"

	DT "bennypowers.dev/cem/designtokens"
	DD "bennypowers.dev/cem/generate/demodiscovery"
	Q "bennypowers.dev/cem/generate/queries"
	M "bennypowers.dev/cem/manifest"
	W "bennypowers.dev/cem/workspace"

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
	// Create jobs for all included files
	jobs := make([]processJob, 0, len(result.includedFiles))
	for _, file := range result.includedFiles {
		jobs = append(jobs, processJob{file: file, ctx: ctx})
	}

	// Use parallel processor
	processor := NewModuleBatchProcessor(qm, nil, cssParseCache) // No dependency tracker for simple processing, use global cache
	processingResult := processor.ProcessModulesSimple(context.Background(), jobs, ModuleProcessorSimpleFunc(processModule))

	return processingResult.Modules, processingResult.Logs, processingResult.Aliases, processingResult.Errors
}

func processModule(
	job processJob,
	qm *Q.QueryManager,
	parser *ts.Parser,
	cssCache CssCache,
) (module *M.Module, tagAliases map[string]string, logCtx *LogCtx, errs error) {
	defer parser.Reset()
	mp, err := NewModuleProcessor(job.ctx, job.file, parser, qm, cssCache)
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

func processModuleWithDeps(
	job processJob,
	qm *Q.QueryManager,
	parser *ts.Parser,
	depTracker *FileDependencyTracker,
	cssCache CssCache,
) (module *M.Module, tagAliases map[string]string, logCtx *LogCtx, errs error) {
	defer parser.Reset()
	mp, err := NewModuleProcessor(job.ctx, job.file, parser, qm, cssCache)
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

	// Record dependencies for incremental rebuilds
	if depTracker != nil && module != nil {
		styleImports := make([]string, 0, len(mp.styleImportsBindingToSpecMap))
		for _, spec := range mp.styleImportsBindingToSpecMap {
			styleImports = append(styleImports, spec)
		}

		importedFiles := make([]string, 0, len(mp.importBindingToSpecMap))
		for _, imp := range mp.importBindingToSpecMap {
			importedFiles = append(importedFiles, imp.spec)
		}

		if depErr := depTracker.RecordModuleDependencies(job.file, styleImports, importedFiles); depErr != nil {
			// Join dependency resolution errors with any existing processing errors
			err = errors.Join(err, depErr)
		}
	}

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
	session, err := NewGenerateSession(ctx)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	pkg, err := session.GenerateFullManifest(context.Background())
	if err != nil {
		return nil, err
	}

	manifestStr, err := M.SerializeToString(pkg)
	if err != nil {
		return nil, fmt.Errorf("module serialize failed: %w", err)
	}
	return &manifestStr, nil
}
