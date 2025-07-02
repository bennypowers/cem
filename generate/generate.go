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
	"path/filepath"
	"runtime"
	"slices"
	"sync"

	"github.com/pterm/pterm"

	C "bennypowers.dev/cem/cmd/config"
	DT "bennypowers.dev/cem/designtokens"
	DD "bennypowers.dev/cem/generate/demodiscovery"
	Q "bennypowers.dev/cem/generate/queries"
	M "bennypowers.dev/cem/manifest"

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

func preprocess(cfg *C.CemConfig) (r preprocessResult, errs error) {
	r.excludePatterns = make([]string, 0, len(cfg.Generate.Exclude)+len(defaultExcludePatterns))
	r.includedFiles = make([]string, 0)
	r.excludePatterns = append(r.excludePatterns, cfg.Generate.Exclude...)
	if !cfg.Generate.NoDefaultExcludes {
		r.excludePatterns = append(r.excludePatterns, defaultExcludePatterns...)
	}
	if cfg.Generate.DesignTokens.Spec != "" {
		tokens, err := DT.LoadDesignTokens(cfg)
		if err != nil {
			errs = errors.Join(errs, err)
		}
		r.designTokens = tokens
	}
	if cfg.Generate.DemoDiscovery.FileGlob != "" {
		demoFiles, err := DS.Glob(cfg.Generate.DemoDiscovery.FileGlob)
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

func process(
	cfg *C.CemConfig,
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
	jobsChan := make(chan string, len(result.includedFiles))

	// Fill jobs channel with files to process
	for _, file := range result.includedFiles {
		jobsChan <- file
	}
	close(jobsChan)


	wg.Add(numWorkers)
	for range numWorkers {
		go func() {
			defer wg.Done()
			parser := Q.GetTypeScriptParser()
			defer Q.PutTypeScriptParser(parser)
			for file := range jobsChan {
				module, tagAliases, logger, err := processModule(file, cfg, qm, parser)

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
				if err != nil {
					errsMu.Lock()
					errsList = append(errsList, err)
					errsMu.Unlock()
				}
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
	file string,
	cfg *C.CemConfig,
	qm *Q.QueryManager,
	parser *ts.Parser,
) (module *M.Module, tagAliases map[string]string, logCtx *LogCtx, errs error) {
	mp := NewModuleProcessor(file, parser, cfg, qm)
	if cfg.Verbose {
		mp.logger.Section.Printf("Module: %s", mp.logger.File)
	}
	module, tagAliases, err := mp.Collect()
	pterm.Print(mp.logger.Buffer.String())
	parser.Reset()
	mp.Close()
	return module, tagAliases, mp.logger, err
}

func postprocess(
	cfg *C.CemConfig,
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

	packageJsonDir := cfg.ProjectDir
	if packageJsonDir == "" {
		packageJsonDir = filepath.Join(cfg.ConfigFile, "../")
	}
	packageJsonPath := filepath.Join(packageJsonDir, "package.json")

	// Because demo discovery and design tokens may mutate modules, we need to coordinate by pointer
	for i := range modules {
		wg.Add(1)
		go func(module *M.Module) {
			defer wg.Done()

			resolvedPath, err := M.ResolveExportPath(packageJsonPath, module.Path)
			if err != nil {
				pterm.Error.Println(err)
			} else {
				module.Path = resolvedPath
			}
			if result.designTokens != nil {
				DT.MergeDesignTokensToModule(module, *result.designTokens)
			}
			// Discover demos and attach to manifest
			if len(demoMap) > 0 {
				err := DD.DiscoverDemos(cfg, allTagAliases, module, qm, demoMap)
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
func Generate(cfg *C.CemConfig) (manifest *string, errs error) {
	qm, err := Q.NewQueryManager()
	if err != nil {
		pterm.Fatal.Printfln("Could not create QueryManager: %v", err)
	}
	defer qm.Close()
	result, err := preprocess(cfg)
	if err != nil {
		errs = errors.Join(errs, err)
	}
	modules, logs, aliases, err := process(cfg, result, qm)
	if err != nil {
		errs = errors.Join(errs, err)
	}
	pkg, err := postprocess(cfg, result, aliases, qm, modules)
	if err != nil {
		errs = errors.Join(errs, err)
	}
	if cfg.Verbose {
		RenderBarChart(logs)
	}
	manifestStr, err := M.SerializeToString(&pkg)
	if err != nil {
		return nil, errors.Join(errs, err)
	}
	return &manifestStr, errs
}
