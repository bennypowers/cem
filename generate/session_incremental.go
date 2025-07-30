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
	"context"
	"errors"
	"path/filepath"
	"sync"

	DT "bennypowers.dev/cem/designtokens"
	DD "bennypowers.dev/cem/generate/demodiscovery"
	M "bennypowers.dev/cem/manifest"
	"github.com/pterm/pterm"
)

// ProcessChangedFiles performs incremental processing for a set of changed files
func (gs *GenerateSession) ProcessChangedFiles(ctx context.Context, changedFiles []string) (*M.Package, error) {
	return gs.ProcessChangedFilesWithSkip(ctx, changedFiles, false)
}

// ProcessChangedFilesWithSkip performs incremental processing for a set of changed files with optional demo discovery skipping
func (gs *GenerateSession) ProcessChangedFilesWithSkip(ctx context.Context, changedFiles []string, skipDemoDiscovery bool) (*M.Package, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	// Check verbose flag once
	cfg, _ := gs.setupCtx.Config()
	verbose := cfg != nil && cfg.Verbose

	// Invalidate CSS cache for any changed CSS files (convert module paths to FS paths)
	cssFiles := make([]string, 0)
	for _, modulePath := range changedFiles {
		if filepath.Ext(modulePath) == ".css" {
			fsPath := gs.setupCtx.ModulePathToFS(modulePath)
			cssFiles = append(cssFiles, fsPath)
		}
	}
	if len(cssFiles) > 0 {
		gs.setupCtx.CssCache().Invalidate(cssFiles)
		if verbose {
			pterm.Debug.Printf("Invalidated CSS cache for files: %v\n", cssFiles)
		}
	}

	// Determine which modules are affected by the changes
	affectedModules := gs.setupCtx.DependencyTracker().GetModulesAffectedByFiles(changedFiles)

	if len(affectedModules) == 0 {
		// No modules affected, return current manifest
		if verbose {
			pterm.Debug.Printf("No modules affected by changes: %v\n", changedFiles)
		}
		return gs.InMemoryManifest(), nil
	}

	// later on we can make this configurable, but for simplicity's sake, it's ok to hard code now
	maxAffectedModulesBeforeFullRebuild := 3
	// Check if we have a base manifest to work with
	currentManifest := gs.InMemoryManifest()
	if currentManifest == nil || len(affectedModules) > maxAffectedModulesBeforeFullRebuild {
		// No base manifest or too many affected modules - do full rebuild
		if verbose {
			pterm.Debug.Printf("Files changed: %v, affected modules: %v - performing full rebuild (no base or too many changes)\n", changedFiles, affectedModules)
		}
		return gs.GenerateFullManifest(ctx)
	}

	// Try incremental processing
	if verbose {
		pterm.Debug.Printf("Files changed: %v, affected modules: %v - attempting incremental rebuild\n", changedFiles, affectedModules)
	}
	return gs.ProcessModulesIncrementalWithSkip(ctx, affectedModules, skipDemoDiscovery)
}

// ProcessModulesIncremental processes only the specified modules incrementally
func (gs *GenerateSession) ProcessModulesIncremental(ctx context.Context, modulePaths []string) (*M.Package, error) {
	return gs.ProcessModulesIncrementalWithSkip(ctx, modulePaths, false)
}

// ProcessModulesIncrementalWithSkip processes only the specified modules incrementally with optional demo discovery skipping
func (gs *GenerateSession) ProcessModulesIncrementalWithSkip(ctx context.Context, modulePaths []string, skipDemoDiscovery bool) (*M.Package, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	// Check verbose flag once
	cfg, _ := gs.setupCtx.Config()
	verbose := cfg != nil && cfg.Verbose

	if verbose {
		pterm.Debug.Printf("Processing %d modules incrementally: %v\n", len(modulePaths), modulePaths)
	}

	// Get the preprocessing result (we need some global context)
	result, err := gs.preprocessWithContext(ctx)
	if err != nil {
		return nil, WrapIncrementalError("preprocess", err)
	}

	// Process only the affected modules
	updatedModules, logs, aliases, err := gs.processSpecificModules(ctx, result, modulePaths)
	if err != nil {
		return nil, WrapIncrementalError("module processing", err)
	}

	// Merge the updated modules into the existing manifest
	gs.MergeModulesIntoManifest(updatedModules)

	// Get the updated manifest
	updatedManifest := gs.InMemoryManifest()
	if updatedManifest == nil {
		return nil, NewError("updated manifest is nil after incremental processing")
	}

	// Apply demo discovery and design tokens to affected modules only
	if err := gs.applyPostProcessingToModules(ctx, result, aliases, updatedModules, skipDemoDiscovery); err != nil {
		pterm.Warning.Printf("Incremental post-processing failed: %v\n", err)
		// Don't fail the entire build for post-processing issues
	}

	// Log performance info
	if verbose {
		pterm.Debug.Printf("Processed %d modules incrementally\n", len(logs))
	}

	return updatedManifest, nil
}

// processSpecificModules processes only the specified module files
func (gs *GenerateSession) processSpecificModules(ctx context.Context, result preprocessResult, modulePaths []string) ([]M.Module, []*LogCtx, map[string]string, error) {
	// Check verbose flag once
	cfg, _ := gs.setupCtx.Config()
	verbose := cfg != nil && cfg.Verbose

	// Filter to only process modules that are in the included files list
	includedSet := make(map[string]bool)
	for _, file := range result.includedFiles {
		includedSet[file] = true
	}

	// Create jobs for valid modules
	validJobs := make([]processJob, 0, len(modulePaths))
	for _, modulePath := range modulePaths {
		if includedSet[modulePath] {
			validJobs = append(validJobs, processJob{file: modulePath, ctx: gs.setupCtx.WorkspaceContext})
		} else {
			if verbose {
				pterm.Debug.Printf("Skipping module not in included files: %s\n", modulePath)
			}
		}
	}

	if len(validJobs) == 0 {
		return make([]M.Module, 0), make([]*LogCtx, 0), make(map[string]string), nil
	}

	// Use parallel processor with dependency tracking and optimized worker count
	processor := NewModuleBatchProcessor(gs.setupCtx.QueryManager(), gs.setupCtx.DependencyTracker(), gs.setupCtx.CssCache())
	processor.SetWorkerCount(len(validJobs)) // Optimize for small incremental builds

	if verbose {
		pterm.Debug.Printf("Starting incremental processing with optimized workers for %d modules\n", len(validJobs))
	}
	processingResult := processor.ProcessModules(ctx, validJobs, ModuleProcessorFunc(processModuleWithDeps))

	return processingResult.Modules, processingResult.Logs, processingResult.Aliases, processingResult.Errors
}

// MergeModulesIntoManifest merges updated modules into the existing in-memory manifest
// using O(1) lookups via the persistent module index for better performance.
func (gs *GenerateSession) MergeModulesIntoManifest(updatedModules []M.Module) {
	gs.mu.Lock()
	defer gs.mu.Unlock()

	if gs.inMemoryManifest == nil {
		// No existing manifest, create a new one
		pkg := M.NewPackage(updatedModules)
		gs.inMemoryManifest = &pkg
		gs.rebuildModuleIndex()
		return
	}

	// Merge or replace modules using the persistent index (O(1) lookups)
	for _, updatedModule := range updatedModules {
		if existingModule, exists := gs.moduleIndex[updatedModule.Path]; exists {
			// Replace existing module content
			*existingModule = updatedModule
			// Index still points to the same memory location, so no update needed
		} else {
			// Add new module
			gs.inMemoryManifest.Modules = append(gs.inMemoryManifest.Modules, updatedModule)
			// Update index with pointer to the new module
			newModule := &gs.inMemoryManifest.Modules[len(gs.inMemoryManifest.Modules)-1]
			gs.moduleIndex[updatedModule.Path] = newModule
		}
	}
}

// applyPostProcessingToModules applies demo discovery and design tokens to specific modules only, with optional demo discovery skipping
func (gs *GenerateSession) applyPostProcessingToModules(ctx context.Context, result preprocessResult, allTagAliases map[string]string, modules []M.Module, skipDemoDiscovery bool) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if len(modules) == 0 {
		return nil
	}

	var wg sync.WaitGroup
	var errsMu sync.Mutex
	errsList := make([]error, 0)

	// Build the demo map once if needed and not skipped
	var demoMap map[string][]string
	if !skipDemoDiscovery && len(result.demoFiles) > 0 {
		var err error
		demoMap, err = DD.NewDemoMap(result.demoFiles)
		if err != nil {
			errsList = append(errsList, err)
		}
	}

	// Process each updated module
	for i := range modules {
		wg.Add(1)
		go func(module *M.Module) {
			defer wg.Done()

			// Apply design tokens if available
			if result.designTokens != nil {
				DT.MergeDesignTokensToModule(module, *result.designTokens)
			}

			// Discover demos and attach to module if available
			if len(demoMap) > 0 {
				err := DD.DiscoverDemos(gs.setupCtx.WorkspaceContext, allTagAliases, module, gs.setupCtx.QueryManager(), demoMap)
				if err != nil {
					errsMu.Lock()
					errsList = append(errsList, err)
					errsMu.Unlock()
				}
			}
		}(&modules[i])
	}

	wg.Wait()

	if len(errsList) > 0 {
		return errors.Join(errsList...)
	}

	return nil
}

// UpdateDemoDiscovery re-runs demo discovery when new demo files appear
func (gs *GenerateSession) UpdateDemoDiscovery(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	// TODO: Implement incremental demo discovery
	// For now, this would require a full rebuild to re-associate demos
	return nil
}

// UpdateDesignTokens reloads and re-applies design tokens when token files change
func (gs *GenerateSession) UpdateDesignTokens(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	// TODO: Implement incremental design token updates
	// For now, this would require a full rebuild to re-apply tokens
	return nil
}

// ProcessChangedFilesIncremental performs true incremental processing (work in progress)
func (gs *GenerateSession) ProcessChangedFilesIncremental(ctx context.Context, changedFiles []string) (*M.Package, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	// Determine which modules are affected by the changes
	affectedModules := gs.setupCtx.DependencyTracker().GetModulesAffectedByFiles(changedFiles)

	if len(affectedModules) == 0 {
		// No modules affected, return current manifest
		return gs.InMemoryManifest(), nil
	}

	return gs.ProcessModulesIncremental(ctx, affectedModules)
}
