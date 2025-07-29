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
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"maps"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	DT "bennypowers.dev/cem/designtokens"
	DD "bennypowers.dev/cem/generate/demodiscovery"
	Q "bennypowers.dev/cem/generate/queries"
	M "bennypowers.dev/cem/manifest"
	W "bennypowers.dev/cem/workspace"
	DS "github.com/bmatcuk/doublestar"
	"github.com/fsnotify/fsnotify"
	"github.com/pterm/pterm"
)

// GenerateSession holds reusable state for efficient generation cycles.
// The QueryManager is expensive to initialize (loads all tree-sitter queries),
// so we reuse it across multiple generation runs in watch mode.
type GenerateSession struct {
	ctx              W.WorkspaceContext
	queryManager     *Q.QueryManager
	inMemoryManifest *M.Package
	depTracker       *FileDependencyTracker
	mu               sync.RWMutex // protects inMemoryManifest
}

// Type aliases for dependency tracking maps
type FileHashMap map[string][32]byte
type ModuleDependencyMap map[string]*ModuleDependencies
type CssReverseDepMap map[string][]string

// FileDependencyTracker tracks file dependencies and content hashes for incremental updates
type FileDependencyTracker struct {
	mu            sync.RWMutex
	fileHashes    FileHashMap         // FS path -> hash
	moduleDeps    ModuleDependencyMap // Module path -> dependencies
	cssDepReverse CssReverseDepMap    // FS path -> module paths that depend on it
	lastScanTime  time.Time
	ctx           W.WorkspaceContext
}

// ModuleDependencies tracks dependencies for a specific module
type ModuleDependencies struct {
	ModulePath    string
	StyleImports  []string
	ImportedFiles []string
	LastModified  time.Time
}

// NewFileDependencyTracker creates a new dependency tracker
func NewFileDependencyTracker(ctx W.WorkspaceContext) *FileDependencyTracker {
	return &FileDependencyTracker{
		fileHashes:    make(FileHashMap),
		moduleDeps:    make(ModuleDependencyMap),
		cssDepReverse: make(CssReverseDepMap),
		ctx:           ctx,
	}
}

// UpdateFileHash computes and stores the hash for a file (expects FS path)
func (fdt *FileDependencyTracker) UpdateFileHash(fsPath string) ([32]byte, error) {
	fdt.mu.Lock()
	defer fdt.mu.Unlock()

	file, err := os.Open(fsPath)
	if err != nil {
		return [32]byte{}, err
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return [32]byte{}, err
	}

	hash := [32]byte(hasher.Sum(nil))
	fdt.fileHashes[fsPath] = hash
	return hash, nil
}

// HasFileChanged checks if a file has changed since last scan (expects FS path)
func (fdt *FileDependencyTracker) HasFileChanged(fsPath string) (bool, error) {
	fdt.mu.RLock()
	lastHash, exists := fdt.fileHashes[fsPath]
	fdt.mu.RUnlock()

	if !exists {
		return true, nil // New file
	}

	currentHash, err := fdt.UpdateFileHash(fsPath)
	if err != nil {
		return true, err // Assume changed if we can't read it
	}

	return currentHash != lastHash, nil
}

// GetModulesAffectedByFiles returns modules that need rebuilding due to file changes
// Expects module paths as input, returns module paths
func (fdt *FileDependencyTracker) GetModulesAffectedByFiles(changedModulePaths []string) []string {
	fdt.mu.RLock()
	defer fdt.mu.RUnlock()

	affectedModules := make(map[string]bool)

	for _, modulePath := range changedModulePaths {
		// Direct module file changes
		if deps, exists := fdt.moduleDeps[modulePath]; exists {
			affectedModules[deps.ModulePath] = true
		}

		// CSS dependency changes - convert module path to FS path for CSS reverse lookup
		fsPath := fdt.ctx.ModulePathToFS(modulePath)
		if modules, exists := fdt.cssDepReverse[fsPath]; exists {
			for _, module := range modules {
				affectedModules[module] = true
			}
		}
	}

	result := make([]string, 0, len(affectedModules))
	for module := range affectedModules {
		result = append(result, module)
	}
	return result
}

// RecordModuleDependencies stores dependencies for a module
func (fdt *FileDependencyTracker) RecordModuleDependencies(modulePath string, styleImports, importedFiles []string) {
	fdt.mu.Lock()
	defer fdt.mu.Unlock()

	// Resolve all dependency paths using workspace context
	resolvedStyleImports := make([]string, 0, len(styleImports))
	resolvedImportedFiles := make([]string, 0, len(importedFiles))

	for _, styleImport := range styleImports {
		resolved := fdt.ctx.ResolveModuleDependency(modulePath, styleImport)
		resolvedStyleImports = append(resolvedStyleImports, resolved)
	}

	for _, importFile := range importedFiles {
		resolved := fdt.ctx.ResolveModuleDependency(modulePath, importFile)
		resolvedImportedFiles = append(resolvedImportedFiles, resolved)
	}

	deps := &ModuleDependencies{
		ModulePath:    modulePath,
		StyleImports:  resolvedStyleImports,
		ImportedFiles: resolvedImportedFiles,
		LastModified:  time.Now(),
	}

	fdt.moduleDeps[modulePath] = deps

	// Update reverse CSS dependencies using FS paths for CSS files
	for _, cssModulePath := range resolvedStyleImports {
		cssFS := fdt.ctx.ModulePathToFS(cssModulePath)
		if fdt.cssDepReverse[cssFS] == nil {
			fdt.cssDepReverse[cssFS] = make([]string, 0)
		}
		// Avoid duplicates
		found := false
		for _, existing := range fdt.cssDepReverse[cssFS] {
			if existing == modulePath {
				found = true
				break
			}
		}
		if !found {
			fdt.cssDepReverse[cssFS] = append(fdt.cssDepReverse[cssFS], modulePath)
		}
	}
}

// NewGenerateSession creates a new session with initialized QueryManager.
// This is expensive (tree-sitter query loading) and should be done once
// per watch session or single generation run.
func NewGenerateSession(ctx W.WorkspaceContext) (*GenerateSession, error) {
	qm, err := Q.NewQueryManager()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize QueryManager: %w", err)
	}

	return &GenerateSession{
		ctx:          ctx,
		queryManager: qm,
		depTracker:   NewFileDependencyTracker(ctx),
	}, nil
}

// Close releases resources held by the session
func (gs *GenerateSession) Close() {
	if gs.queryManager != nil {
		gs.queryManager.Close()
	}
}

// GenerateFullManifest performs a complete generation using the existing logic.
// This is used for the initial generation in watch mode and for regular generate command.
func (gs *GenerateSession) GenerateFullManifest(ctx context.Context) (*M.Package, error) {
	// Check for cancellation
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	// Use existing generation logic with our reusable QueryManager
	result, err := gs.preprocessWithContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("preprocess failed: %w", err)
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	modules, logs, aliases, err := gs.processWithContext(ctx, result)
	if err != nil {
		return nil, fmt.Errorf("process failed: %w", err)
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	pkg, err := gs.postprocessWithContext(ctx, result, aliases, modules)
	if err != nil {
		return nil, fmt.Errorf("postprocess failed: %w", err)
	}

	// Update in-memory manifest
	gs.mu.Lock()
	gs.inMemoryManifest = &pkg
	gs.mu.Unlock()

	cfg, _ := gs.ctx.Config()
	if cfg != nil && cfg.Verbose {
		RenderBarChart(logs)
	}

	return &pkg, nil
}

// GetInMemoryManifest returns a copy of the current in-memory manifest.
// This is safe for concurrent access and intended for LSP integration.
func (gs *GenerateSession) GetInMemoryManifest() *M.Package {
	gs.mu.RLock()
	defer gs.mu.RUnlock()

	if gs.inMemoryManifest == nil {
		return nil
	}

	// Return a copy to prevent concurrent modification issues
	// TODO: Consider if we need deep copying for LSP use case
	manifest := *gs.inMemoryManifest
	return &manifest
}

// preprocessWithContext is the existing preprocess logic with cancellation support
func (gs *GenerateSession) preprocessWithContext(ctx context.Context) (preprocessResult, error) {
	select {
	case <-ctx.Done():
		return preprocessResult{}, ctx.Err()
	default:
	}

	return preprocess(gs.ctx)
}

// processWithContext is the existing process logic with cancellation support
func (gs *GenerateSession) processWithContext(ctx context.Context, result preprocessResult) ([]M.Module, []*LogCtx, map[string]string, error) {
	select {
	case <-ctx.Done():
		return nil, nil, nil, ctx.Err()
	default:
	}

	// Use the dependency-tracking version to collect file dependencies
	return gs.processWithDeps(ctx, result)
}

// processWithDeps processes files while tracking dependencies for incremental rebuilds
func (gs *GenerateSession) processWithDeps(ctx context.Context, result preprocessResult) ([]M.Module, []*LogCtx, map[string]string, error) {
	select {
	case <-ctx.Done():
		return nil, nil, nil, ctx.Err()
	default:
	}

	numWorkers := runtime.NumCPU()
	pterm.Info.Printf("Starting Generation with %d workers\n", numWorkers)

	var wg sync.WaitGroup
	var aliasesMu sync.Mutex
	var modulesMu sync.Mutex
	var errsMu sync.Mutex
	var logsMu sync.Mutex
	errsList := make([]error, 0)
	logs := make([]*LogCtx, 0, len(result.includedFiles))
	aliases := make(map[string]string)
	var modules []M.Module

	jobsChan := make(chan processJob, len(result.includedFiles))

	// Fill jobs channel with files to process
	for _, file := range result.includedFiles {
		jobsChan <- processJob{file: file, ctx: gs.ctx}
	}
	close(jobsChan)

	wg.Add(numWorkers)
	for range numWorkers {
		go func() {
			defer wg.Done()
			parser := Q.GetTypeScriptParser()
			defer Q.PutTypeScriptParser(parser)
			for job := range jobsChan {
				module, tagAliases, logger, err := processModuleWithDeps(job, gs.queryManager, parser, gs.depTracker)
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

	var errs error
	if len(errsList) > 0 {
		errs = errors.Join(errsList...)
	}

	return modules, logs, aliases, errs
}

// processSpecificModules processes only the specified module files
func (gs *GenerateSession) processSpecificModules(ctx context.Context, result preprocessResult, modulePaths []string) ([]M.Module, []*LogCtx, map[string]string, error) {
	select {
	case <-ctx.Done():
		return nil, nil, nil, ctx.Err()
	default:
	}

	numWorkers := runtime.NumCPU()
	if len(modulePaths) < numWorkers {
		numWorkers = len(modulePaths) // Don't over-parallelize for few modules
	}
	pterm.Debug.Printf("Starting incremental processing with %d workers for %d modules\n", numWorkers, len(modulePaths))

	var wg sync.WaitGroup
	var aliasesMu sync.Mutex
	var modulesMu sync.Mutex
	var errsMu sync.Mutex
	var logsMu sync.Mutex
	errsList := make([]error, 0)
	logs := make([]*LogCtx, 0, len(modulePaths))
	aliases := make(map[string]string)
	var modules []M.Module

	jobsChan := make(chan processJob, len(modulePaths))

	// Filter to only process modules that are in the included files list
	validModules := make([]string, 0, len(modulePaths))
	includedSet := make(map[string]bool)
	for _, file := range result.includedFiles {
		includedSet[file] = true
	}

	for _, modulePath := range modulePaths {
		if includedSet[modulePath] {
			validModules = append(validModules, modulePath)
		} else {
			pterm.Debug.Printf("Skipping module not in included files: %s\n", modulePath)
		}
	}

	// Fill jobs channel with valid modules to process
	for _, file := range validModules {
		jobsChan <- processJob{file: file, ctx: gs.ctx}
	}
	close(jobsChan)

	if len(validModules) == 0 {
		return modules, logs, aliases, nil
	}

	wg.Add(numWorkers)
	for range numWorkers {
		go func() {
			defer wg.Done()
			parser := Q.GetTypeScriptParser()
			defer Q.PutTypeScriptParser(parser)
			for job := range jobsChan {
				module, tagAliases, logger, err := processModuleWithDeps(job, gs.queryManager, parser, gs.depTracker)
				if err != nil {
					errsMu.Lock()
					errsList = append(errsList, err)
					errsMu.Unlock()
				}

				// Save log for performance tracking
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

	var errs error
	if len(errsList) > 0 {
		errs = errors.Join(errsList...)
	}

	return modules, logs, aliases, errs
}

// postprocessWithContext is the existing postprocess logic with cancellation support
func (gs *GenerateSession) postprocessWithContext(ctx context.Context, result preprocessResult, aliases map[string]string, modules []M.Module) (M.Package, error) {
	select {
	case <-ctx.Done():
		return M.Package{}, ctx.Err()
	default:
	}

	// TODO: Add cancellation points within the postprocess function
	// For now, we'll use the existing postprocess function
	return postprocess(gs.ctx, result, aliases, gs.queryManager, modules)
}

// WatchSession manages the long-lived watch mode state
type WatchSession struct {
	ctx             W.WorkspaceContext
	globs           []string
	generateSession *GenerateSession
	debounceTimer   *time.Timer
	mu              sync.Mutex
	cancelCurrent   context.CancelFunc
	lastWrittenHash map[string][32]byte // file path -> SHA256 hash of content we wrote
	pendingChanges  map[string]bool     // files that have changed and are pending processing
}

// NewWatchSession creates a new watch session with the given workspace context and globs
func NewWatchSession(ctx W.WorkspaceContext, globs []string) (*WatchSession, error) {
	generateSession, err := NewGenerateSession(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create generate session: %w", err)
	}

	return &WatchSession{
		ctx:             ctx,
		globs:           globs,
		generateSession: generateSession,
		lastWrittenHash: make(map[string][32]byte),
		pendingChanges:  make(map[string]bool),
	}, nil
}

// Close releases resources held by the watch session
func (ws *WatchSession) Close() {
	if ws.generateSession != nil {
		ws.generateSession.Close()
	}
}

// RunWatch starts the file watching mode with initial generation
func (ws *WatchSession) RunWatch() error {
	pterm.Info.Println("Starting watch mode...")

	// Do initial generation
	pterm.Info.Println("Starting Generation with 4 workers")
	start := time.Now()

	if err := ws.generateOnce(context.Background()); err != nil {
		pterm.Error.Printf("Initial generation failed: %v\n", err)
		return err
	}
	duration := time.Since(start)
	pterm.Success.Printf("Generated in %s\n", ColorizeDuration(duration).Sprint(duration))

	// Set up file watcher
	watcher, err := ws.setupWatcher()
	if err != nil {
		return fmt.Errorf("failed to setup file watcher: %w", err)
	}
	defer watcher.Close()

	pterm.Info.Println("Watching for file changes... (Ctrl+C to stop)")

	// Watch for changes
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return nil
			}
			ws.handleFileChange(event)
		case err, ok := <-watcher.Errors:
			if !ok {
				return nil
			}
			pterm.Warning.Printf("File watcher error: %v\n", err)
		}
	}
}

// setupWatcher creates and configures the file system watcher
func (ws *WatchSession) setupWatcher() (*fsnotify.Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	// Get all directories that might contain files matching our globs
	dirs := make(map[string]bool)
	dirs[ws.ctx.Root()] = true // Always watch the root

	// Expand globs to get current files and their directories
	for _, glob := range ws.globs {
		files, err := ws.ctx.Glob(glob)
		if err != nil {
			continue
		}
		for _, file := range files {
			dir := filepath.Dir(file)
			if !filepath.IsAbs(dir) {
				dir = filepath.Join(ws.ctx.Root(), dir)
			}
			dirs[dir] = true
		}
	}

	// Add demo discovery directories if configured
	cfg, err := ws.ctx.Config()
	if err == nil && cfg.Generate.DemoDiscovery.FileGlob != "" {
		demoFiles, err := ws.ctx.Glob(cfg.Generate.DemoDiscovery.FileGlob)
		if err == nil {
			for _, file := range demoFiles {
				dir := filepath.Dir(file)
				if !filepath.IsAbs(dir) {
					dir = filepath.Join(ws.ctx.Root(), dir)
				}
				dirs[dir] = true
			}
		}
	}

	// Watch all identified directories
	for dir := range dirs {
		if err := watcher.Add(dir); err != nil {
			pterm.Warning.Printf("Failed to watch directory %s: %v\n", dir, err)
		}
	}

	return watcher, nil
}

// handleFileChange processes a file system event with debouncing
func (ws *WatchSession) handleFileChange(event fsnotify.Event) {
	// Only care about write and create events
	if event.Op&fsnotify.Write == 0 && event.Op&fsnotify.Create == 0 {
		return
	}

	// Ignore changes that we just made to prevent infinite loops
	if ws.isOurWrite(event.Name) {
		pterm.Debug.Printf("Ignoring our own write to: %s\n", event.Name)
		return
	}

	// Only process files that match our input globs
	if !ws.matchesInputGlobs(event.Name) {
		pterm.Debug.Printf("Ignoring file not matching input globs: %s\n", event.Name)
		return
	}

	pterm.Debug.Printf("File change detected: %s (op: %s)\n", event.Name, event.Op)

	ws.mu.Lock()
	defer ws.mu.Unlock()

	// Track this file as changed - convert FS path to module path
	modulePath, err := ws.ctx.FSPathToModule(event.Name)
	if err != nil {
		// If conversion fails, use the FS path as-is
		modulePath = event.Name
	}
	ws.pendingChanges[modulePath] = true

	// Cancel any existing timer
	if ws.debounceTimer != nil {
		ws.debounceTimer.Stop()
	}

	// Start a new debounce timer (100ms delay)
	ws.debounceTimer = time.AfterFunc(100*time.Millisecond, func() {
		ws.processChanges()
	})
}

// isOurWrite checks if the file change was caused by our own write operation
func (ws *WatchSession) isOurWrite(filePath string) bool {
	// Read the current file content and compute its hash
	file, err := os.Open(filePath)
	if err != nil {
		return false // If we can't read it, assume it's not our write
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return false
	}

	currentHash := [32]byte(hasher.Sum(nil))

	// Check if this hash matches what we last wrote to this file
	ws.mu.Lock()
	lastHash, exists := ws.lastWrittenHash[filepath.Clean(filePath)]
	ws.mu.Unlock()

	return exists && currentHash == lastHash
}

// matchesInputGlobs checks if a file path matches any of our input globs
func (ws *WatchSession) matchesInputGlobs(filePath string) bool {
	// Convert absolute path to relative path from workspace root
	relPath, err := filepath.Rel(ws.ctx.Root(), filePath)
	if err != nil {
		// If we can't get relative path, try with the original path
		relPath = filePath
	}

	// Check against each glob pattern
	for _, glob := range ws.globs {
		matched, err := filepath.Match(glob, relPath)
		if err == nil && matched {
			return true
		}

		// Also try with doublestar for more complex patterns like **/*.ts
		matched, err = DS.PathMatch(glob, relPath)
		if err == nil && matched {
			return true
		}
	}

	return false
}

// processChanges handles the actual regeneration after debouncing
func (ws *WatchSession) processChanges() {
	ws.mu.Lock()

	// Cancel any in-progress generation
	if ws.cancelCurrent != nil {
		ws.cancelCurrent()
	}

	// Create new cancellable context for this generation cycle
	ctx, cancel := context.WithCancel(context.Background())
	ws.cancelCurrent = cancel

	// Capture the list of changed files and reset pending changes
	changedFiles := make([]string, 0, len(ws.pendingChanges))
	for file := range ws.pendingChanges {
		changedFiles = append(changedFiles, file)
	}
	ws.pendingChanges = make(map[string]bool)

	ws.mu.Unlock()

	pterm.Info.Println("File changes detected, regenerating...")
	start := time.Now()

	// Try incremental processing first
	if len(changedFiles) > 0 {
		pterm.Debug.Printf("Changed files: %v\n", changedFiles)
		pkg, err := ws.generateSession.ProcessChangedFiles(ctx, changedFiles)
		if err != nil {
			if ctx.Err() == context.Canceled {
				pterm.Warning.Println("Generation cancelled due to new file changes")
			} else {
				pterm.Error.Printf("Generation failed: %v\n", err)
			}
			return
		}

		// Write the manifest
		manifestStr, err := M.SerializeToString(pkg)
		if err != nil {
			pterm.Error.Printf("Failed to serialize manifest: %v\n", err)
			return
		}

		if err := ws.writeManifest(manifestStr); err != nil {
			pterm.Error.Printf("Failed to write manifest: %v\n", err)
			return
		}
	} else {
		// Fallback to full rebuild if no specific files tracked
		if err := ws.generateOnce(ctx); err != nil {
			if ctx.Err() == context.Canceled {
				pterm.Warning.Println("Generation cancelled due to new file changes")
			} else {
				pterm.Error.Printf("Generation failed: %v\n", err)
			}
			return
		}
	}

	duration := time.Since(start)
	pterm.Success.Printf("Regenerated in %s\n", ColorizeDuration(duration).Sprint(duration))
}

// generateOnce performs a single generation cycle, respecting cancellation
func (ws *WatchSession) generateOnce(ctx context.Context) error {
	// Use the reusable generate session for efficient regeneration
	pkg, err := ws.generateSession.GenerateFullManifest(ctx)
	if err != nil {
		return err
	}

	// Serialize to string
	manifestStr, err := M.SerializeToString(pkg)
	if err != nil {
		return fmt.Errorf("failed to serialize manifest: %w", err)
	}

	// Write the manifest
	return ws.writeManifest(manifestStr)
}

// writeManifest writes the manifest to the configured output path
func (ws *WatchSession) writeManifest(manifestStr string) error {
	cfg, err := ws.ctx.Config()
	if err != nil {
		return err
	}

	// Determine output path (copied from main generate logic)
	outputPath := cfg.Generate.Output
	if outputPath == "" {
		pkgjson, err := ws.ctx.PackageJSON()
		if err != nil {
			return err
		}
		if pkgjson != nil && pkgjson.CustomElements != "" {
			outputPath = filepath.Join(ws.ctx.Root(), pkgjson.CustomElements)
		}
	}

	if outputPath == "" {
		// If no output path, just skip writing (or could print to stdout)
		return nil
	}

	// Prepare the content we're about to write
	content := manifestStr + "\n"
	contentBytes := []byte(content)

	writer, err := ws.ctx.OutputWriter(outputPath)
	if err != nil {
		return err
	}

	_, err = writer.Write(contentBytes)
	if err != nil {
		writer.Close()
		return err
	}

	// Close the writer first to ensure all data is flushed
	err = writer.Close()
	if err != nil {
		return err
	}

	// Record the hash of what we just wrote to prevent infinite loops
	// We do this AFTER closing to ensure the file is fully written
	hash := sha256.Sum256(contentBytes)
	ws.mu.Lock()
	ws.lastWrittenHash[filepath.Clean(outputPath)] = hash
	ws.mu.Unlock()

	return nil
}

// ProcessChangedFiles performs incremental processing for a set of changed files
func (gs *GenerateSession) ProcessChangedFiles(ctx context.Context, changedFiles []string) (*M.Package, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	// Invalidate CSS cache for any changed CSS files (convert module paths to FS paths)
	cssFiles := make([]string, 0)
	for _, modulePath := range changedFiles {
		if filepath.Ext(modulePath) == ".css" {
			fsPath := gs.ctx.ModulePathToFS(modulePath)
			cssFiles = append(cssFiles, fsPath)
		}
	}
	if len(cssFiles) > 0 {
		cssParseCache.Invalidate(cssFiles)
		pterm.Debug.Printf("Invalidated CSS cache for files: %v\n", cssFiles)
	}

	// Determine which modules are affected by the changes
	affectedModules := gs.depTracker.GetModulesAffectedByFiles(changedFiles)

	if len(affectedModules) == 0 {
		// No modules affected, return current manifest
		pterm.Debug.Printf("No modules affected by changes: %v\n", changedFiles)
		return gs.GetInMemoryManifest(), nil
	}

	// Check if we have a base manifest to work with
	currentManifest := gs.GetInMemoryManifest()
	if currentManifest == nil || len(affectedModules) > 3 {
		// No base manifest or too many affected modules - do full rebuild
		pterm.Debug.Printf("Files changed: %v, affected modules: %v - performing full rebuild (no base or too many changes)\n", changedFiles, affectedModules)
		return gs.GenerateFullManifest(ctx)
	}

	// Try incremental processing
	pterm.Debug.Printf("Files changed: %v, affected modules: %v - attempting incremental rebuild\n", changedFiles, affectedModules)
	return gs.ProcessModulesIncremental(ctx, affectedModules)
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

// MergeModulesIntoManifest merges updated modules into the existing in-memory manifest
func (gs *GenerateSession) MergeModulesIntoManifest(updatedModules []M.Module) {
	gs.mu.Lock()
	defer gs.mu.Unlock()

	if gs.inMemoryManifest == nil {
		// No existing manifest, create a new one
		pkg := M.NewPackage(updatedModules)
		gs.inMemoryManifest = &pkg
		return
	}

	// Create a map of existing modules by path for efficient lookup
	existingModules := make(map[string]*M.Module)
	for i := range gs.inMemoryManifest.Modules {
		module := &gs.inMemoryManifest.Modules[i]
		existingModules[module.Path] = module
	}

	// Merge or replace modules
	for _, updatedModule := range updatedModules {
		if existingModule, exists := existingModules[updatedModule.Path]; exists {
			// Replace existing module
			*existingModule = updatedModule
		} else {
			// Add new module
			gs.inMemoryManifest.Modules = append(gs.inMemoryManifest.Modules, updatedModule)
		}
	}
}

// ProcessModulesIncremental processes only the specified modules incrementally
func (gs *GenerateSession) ProcessModulesIncremental(ctx context.Context, modulePaths []string) (*M.Package, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	pterm.Debug.Printf("Processing %d modules incrementally: %v\n", len(modulePaths), modulePaths)

	// Get the preprocessing result (we need some global context)
	result, err := gs.preprocessWithContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("preprocess failed during incremental build: %w", err)
	}

	// Process only the affected modules
	updatedModules, logs, aliases, err := gs.processSpecificModules(ctx, result, modulePaths)
	if err != nil {
		return nil, fmt.Errorf("incremental module processing failed: %w", err)
	}

	// Merge the updated modules into the existing manifest
	gs.MergeModulesIntoManifest(updatedModules)

	// Get the updated manifest
	updatedManifest := gs.GetInMemoryManifest()
	if updatedManifest == nil {
		return nil, fmt.Errorf("failed to get updated manifest after incremental processing")
	}

	// Apply demo discovery and design tokens to affected modules only
	if err := gs.applyPostProcessingToModules(ctx, result, aliases, updatedModules); err != nil {
		pterm.Warning.Printf("Incremental post-processing failed: %v\n", err)
		// Don't fail the entire build for post-processing issues
	}

	// Log performance info
	cfg, _ := gs.ctx.Config()
	if cfg != nil && cfg.Verbose {
		pterm.Debug.Printf("Processed %d modules incrementally\n", len(logs))
	}

	return updatedManifest, nil
}

// ProcessChangedFilesIncremental performs true incremental processing (work in progress)
func (gs *GenerateSession) ProcessChangedFilesIncremental(ctx context.Context, changedFiles []string) (*M.Package, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	// Determine which modules are affected by the changes
	affectedModules := gs.depTracker.GetModulesAffectedByFiles(changedFiles)

	if len(affectedModules) == 0 {
		// No modules affected, return current manifest
		return gs.GetInMemoryManifest(), nil
	}

	// TODO: Implement true incremental processing
	// This would involve:
	// 1. Process only the affected module files
	// 2. Handle CSS dependency invalidation (clear CSS cache for changed CSS files)
	// 3. Update only affected modules in the manifest
	// 4. Re-run demo discovery and design tokens only for affected modules

	return gs.ProcessModulesIncremental(ctx, affectedModules)
}

// applyPostProcessingToModules applies demo discovery and design tokens to specific modules only
func (gs *GenerateSession) applyPostProcessingToModules(ctx context.Context, result preprocessResult, allTagAliases map[string]string, modules []M.Module) error {
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

	// Build the demo map once if needed
	var demoMap map[string][]string
	if len(result.demoFiles) > 0 {
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
				err := DD.DiscoverDemos(gs.ctx, allTagAliases, module, gs.queryManager, demoMap)
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
