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

	DT "bennypowers.dev/cem/internal/designtokens"
	DD "bennypowers.dev/cem/generate/demodiscovery"
	"bennypowers.dev/cem/internal/logging"
	"bennypowers.dev/cem/internal/platform"
	M "bennypowers.dev/cem/manifest"
	Q "bennypowers.dev/cem/internal/treesitter"
	"bennypowers.dev/cem/types"

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
	designTokens    types.DesignTokens
	excludePatterns []string
	includedFiles   []string
}

// preprocess handles config merging for generate command
func preprocess(ctx types.WorkspaceContext) (r preprocessResult, errs error) {
	cfg, err := ctx.Config()
	if err != nil {
		return r, err
	}
	cfg.Generate.Exclude = append(cfg.Generate.Exclude, []string{}...)
	r.excludePatterns = make([]string, 0, len(cfg.Generate.Exclude)+len(defaultExcludePatterns))

	r.includedFiles = make([]string, 0)
	r.excludePatterns = append(r.excludePatterns, cfg.Generate.Exclude...)

	if cfg.Generate.NoDefaultExcludes == nil || !*cfg.Generate.NoDefaultExcludes {
		r.excludePatterns = append(r.excludePatterns, defaultExcludePatterns...)
	}

	if cfg.Generate.DesignTokens.Spec != "" {
		designTokens, err := validateAndLoadDesignTokens(ctx)
		if err != nil {
			errs = errors.Join(errs, err)
		} else {
			r.designTokens = designTokens
		}
	}
	if cfg.Generate.DemoDiscovery.FileGlob != "" {
		demoFiles, err := ctx.Glob(cfg.Generate.DemoDiscovery.FileGlob)
		r.demoFiles = demoFiles
		if err != nil {
			logging.Debug("demo discovery glob: %v", err)
		}
	}
	for _, filePattern := range cfg.Generate.Files {
		expandedFiles, err := ctx.Glob(filePattern)
		for _, file := range expandedFiles {
			if !matchesAnyPattern(file, r.excludePatterns) {
				r.includedFiles = append(r.includedFiles, file)
			}
		}
		if err != nil {
			errs = errors.Join(errs, fmt.Errorf("expanding glob %s: %w", filePattern, err))
		}
	}
	return r, errs
}

type processJob struct {
	file string
	ctx  types.WorkspaceContext
}

func processModule(
	job processJob,
	qm *Q.QueryManager,
	parser *ts.Parser,
	depTracker *FileDependencyTracker,
	cssCache CssCache,
	fsys platform.FileSystem,
) (module *M.Module, tagAliases map[string]string, typeAliases map[string]string, imports map[string]importInfo, logCtx *LogCtx, errs error) {
	defer parser.Reset()
	mp, err := NewModuleProcessor(job.ctx, job.file, parser, qm, cssCache, fsys)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	defer mp.Close()
	if mp.logger.Verbose {
		fmt.Fprintf(mp.logger.Buffer, "\n== Module: %s ==\n\n", mp.logger.File)
	}
	module, tagAliases, typeAliases, imports, err = mp.Collect()

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

	return module, tagAliases, typeAliases, imports, mp.logger, err
}

func postprocess(
	ctx types.WorkspaceContext,
	result preprocessResult,
	allTagAliases map[string]string,
	typeAliases moduleTypeAliasesMap,
	imports moduleImportsMap,
	qm *Q.QueryManager,
	modules []M.Module,
	fsys platform.FileSystem,
) (pkg M.Package, errs error) {
	var wg sync.WaitGroup
	var errsMu sync.Mutex
	errsList := make([]error, 0)

	// Auto-derive aliases for elements without explicit @alias
	var allTagNames []string
	for i := range modules {
		for _, decl := range modules[i].Declarations {
			if ced, ok := decl.(*M.CustomElementDeclaration); ok && ced.TagName != "" {
				allTagNames = append(allTagNames, ced.TagName)
			}
		}
	}
	allTagAliases = DD.AutoDeriveAliases(allTagNames, allTagAliases)

	// Build the demo map once
	cfg, cfgErr := ctx.Config()
	if cfgErr != nil {
		errsList = append(errsList, cfgErr)
	}
	var urlPattern string
	if cfgErr == nil {
		urlPattern = cfg.Generate.DemoDiscovery.URLPattern
	}
	demoMap, err := DD.NewDemoMapWithPattern(ctx, result.demoFiles, urlPattern, allTagAliases, fsys)
	if err != nil {
		errsList = append(errsList, err)
	}

	// Because demo discovery and design tokens may mutate modules, we need to coordinate by pointer
	for i := range modules {
		wg.Add(1)
		go func(module *M.Module) {
			defer wg.Done()
			if result.designTokens != nil {
				DT.MergeDesignTokensToModule(module, result.designTokens)
			}
			// Discover demos and attach to manifest
			if len(demoMap) > 0 {
				err := DD.DiscoverDemos(ctx, allTagAliases, module, qm, demoMap, fsys)
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
	// Resolve re-exported custom element definitions across modules
	resolveReExportedCEDefinitions(&pkg)

	// Resolve type aliases (including cross-package external types)
	externalResolver := NewExternalTypeResolver(ctx, qm, fsys)
	if err := ResolveTypeAliases(&pkg, typeAliases, imports, externalResolver); err != nil {
		errsList = append(errsList, fmt.Errorf("type resolution failed: %w", err))
	}
	if len(errsList) > 0 {
		errs = errors.Join(errsList...)
	}
	return pkg, errs
}

// Generates a custom-elements manifest from a list of typescript files
func Generate(ctx types.WorkspaceContext, fsys platform.FileSystem) (manifest *string, errs error) {
	session, err := NewGenerateSession(ctx, fsys)
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

// validateAndLoadDesignTokens loads design tokens from cache.
// Returns a proper error if the cached object cannot be loaded.
func validateAndLoadDesignTokens(ctx types.WorkspaceContext) (types.DesignTokens, error) {
	tokens, err := ctx.DesignTokensCache().LoadOrReuse(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load design tokens: %w", err)
	}

	if tokens == nil {
		return nil, fmt.Errorf("design tokens cache returned nil - check design tokens spec configuration")
	}

	return tokens, nil
}
