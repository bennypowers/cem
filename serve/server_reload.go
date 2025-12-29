/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

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

package serve

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"time"

	importmappkg "bennypowers.dev/cem/serve/middleware/importmap"
	"bennypowers.dev/cem/serve/middleware/types"
)

// CreateReloadMessage creates a reload message JSON
func (s *Server) CreateReloadMessage(files []string, reason string) ([]byte, error) {
	msg := ReloadMessage{
		Type:   "reload",
		Reason: reason,
		Files:  files,
	}
	return json.Marshal(msg)
}

// BroadcastReload broadcasts a reload event to all WebSocket clients
func (s *Server) BroadcastReload(files []string, reason string) error {
	if s.wsManager == nil {
		return nil // Reload disabled
	}

	msgBytes, err := s.CreateReloadMessage(files, reason)
	if err != nil {
		return err
	}

	s.logger.Debug("Broadcasting reload: reason=%s, files=%v, clients=%d",
		reason, files, s.wsManager.ConnectionCount())

	return s.wsManager.Broadcast(msgBytes)
}

// BroadcastError broadcasts an error notification to all WebSocket clients
func (s *Server) BroadcastError(title, message, file string) error {
	if s.wsManager == nil {
		return nil // WebSocket disabled
	}

	msg := ErrorMessage{
		Type:    "error",
		Title:   title,
		Message: message,
		File:    file,
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	s.logger.Debug("Broadcasting error: title=%s, file=%s, clients=%d",
		title, file, s.wsManager.ConnectionCount())

	return s.wsManager.Broadcast(msgBytes)
}

// logCacheStats periodically logs cache statistics
func (s *Server) logCacheStats(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if s.transformCache != nil {
				stats := s.transformCache.Stats()
				s.logger.Debug("Transform cache stats: %d entries, %.1f%% hit rate, %d MB / %d MB",
					stats.Entries,
					stats.HitRate,
					stats.SizeBytes/(1024*1024),
					stats.MaxSize/(1024*1024),
				)
			}
		case <-s.shutdown:
			return
		}
	}
}

// shouldRebuildPathResolver checks if any changed files require path resolver rebuild
func (s *Server) shouldRebuildPathResolver(filesToProcess []string) bool {
	s.mu.RLock()
	pathResolverSourceFiles := s.pathResolverSourceFiles
	s.mu.RUnlock()

	for _, changedFile := range filesToProcess {
		// Normalize changed file to absolute path for comparison
		// pathResolverSourceFiles contains absolute paths from ParseTsConfig
		absChangedFile, err := filepath.Abs(changedFile)
		if err != nil {
			absChangedFile = changedFile // Fall back to original if Abs fails
		}

		for _, sourceFile := range pathResolverSourceFiles {
			if absChangedFile == sourceFile {
				s.logger.Debug("PathResolver source file changed: %s", sourceFile)
				return true
			}
		}
	}
	return false
}

// filterRelevantFiles filters source files and separates TS/JS files for manifest regeneration
func (s *Server) filterRelevantFiles(filesToProcess []string) (relevantFiles, tsJsFiles []string) {
	relevantFiles = make([]string, 0, len(filesToProcess))
	tsJsFiles = make([]string, 0, len(filesToProcess))

	for _, filePath := range filesToProcess {
		ext := filepath.Ext(filePath)
		isRelevant := ext == ".ts" || ext == ".js" || ext == ".css" || ext == ".html"

		if !isRelevant {
			var relPath string
			if s.watchDir != "" {
				if rel, err := filepath.Rel(s.watchDir, filePath); err == nil {
					relPath = rel
				} else {
					relPath = filePath
				}
			} else {
				relPath = filePath
			}
			s.logger.Debug("Ignoring non-source file: %s", relPath)
			continue
		}

		relevantFiles = append(relevantFiles, filePath)

		// Collect TS/JS files for manifest regeneration
		if ext == ".ts" || ext == ".js" {
			tsJsFiles = append(tsJsFiles, filePath)
		}
	}

	return relevantFiles, tsJsFiles
}

// collectAffectedFiles collects files from transform cache and module graph
func (s *Server) collectAffectedFiles(changedPath string) []string {
	affectedFiles := make(map[string]bool)

	// Add transform cache invalidations
	if s.transformCache != nil {
		invalidatedFiles := s.transformCache.Invalidate(changedPath)
		if len(invalidatedFiles) > 0 {
			s.logger.Debug("Transform cache invalidated %d files: %v", len(invalidatedFiles), invalidatedFiles)
			for _, file := range invalidatedFiles {
				affectedFiles[file] = true
			}
		}
	}

	// Add module graph affected files (includes non-transformed .js files)
	moduleGraphFiles := s.getModuleGraphAffectedFiles(changedPath)
	if len(moduleGraphFiles) > 0 {
		for _, file := range moduleGraphFiles {
			affectedFiles[file] = true
		}
	}

	// Convert map to slice for downstream use
	var result []string
	for file := range affectedFiles {
		result = append(result, file)
	}

	s.logger.Debug("Total affected files: %d (transform cache + module graph)", len(result))
	return result
}

// regenerateManifestIfNeeded regenerates manifest when TS/JS files change
func (s *Server) regenerateManifestIfNeeded(tsJsFiles []string) {
	if len(tsJsFiles) == 0 {
		return
	}

	manifestStart := time.Now()
	s.logger.Debug("Regenerating manifest incrementally for %d file(s)...", len(tsJsFiles))
	manifestSize, err := s.RegenerateManifestIncremental(tsJsFiles)
	manifestDuration := time.Since(manifestStart)
	if err != nil {
		s.logger.Error("Failed to regenerate manifest incrementally: %v", err)
		// Continue anyway - we still want to reload the page
	} else {
		s.logger.Info("Manifest regenerated incrementally (%d bytes) in %v", manifestSize, manifestDuration)
	}
}

// regenerateImportMapIfNeeded regenerates import map when package.json or file structure changes
func (s *Server) regenerateImportMapIfNeeded(event FileEvent) {
	// Only regenerate when necessary and enabled:
	// - package.json was modified (exports may have changed)
	// - files were created (new modules may need to be mapped)
	// - files were deleted (old modules may need to be unmapped)
	if !s.config.ImportMap.Generate || (!event.HasPackageJSON && !event.HasCreates && !event.HasDeletes) {
		return
	}

	importMapStart := time.Now()

	// Extract state under read lock to avoid blocking during I/O
	s.mu.RLock()
	isWorkspace := s.isWorkspace
	workspaceRoot := s.workspaceRoot
	workspacePackages := s.workspacePackages
	watchDir := s.watchDir
	configOverride := s.buildConfigOverride()
	s.mu.RUnlock()

	// Perform I/O without lock to avoid blocking other operations
	var importMap *importmappkg.ImportMap
	var err error
	if isWorkspace {
		// Workspace mode: regenerate workspace import map
		importMap, err = importmappkg.Generate(workspaceRoot, &importmappkg.Config{
			InputMapPath:      s.config.ImportMap.OverrideFile,
			ConfigOverride:    configOverride,
			WorkspacePackages: workspacePackages,
			Logger:            s.logger,
			FS:                s.fs,
		})
	} else {
		// Single-package mode: regenerate import map
		importMap, err = importmappkg.Generate(watchDir, &importmappkg.Config{
			InputMapPath:   s.config.ImportMap.OverrideFile,
			ConfigOverride: configOverride,
			Logger:         s.logger,
			FS:             s.fs,
		})
	}

	// Update state under write lock
	importMapDuration := time.Since(importMapStart)
	s.mu.Lock()
	if err != nil {
		s.logger.Warning("Failed to regenerate import map: %v", err)
	} else {
		s.importMap = importMap
		s.logger.Info("Regenerated import map in %v", importMapDuration)
	}
	s.mu.Unlock()
}

// broadcastSmartReload sends reload messages to affected pages or all clients
func (s *Server) broadcastSmartReload(changedPath, relPath string, invalidatedFiles []string) {
	// Smart reload: only reload pages that import the changed file or its dependents
	affectedPageURLs := s.getAffectedPageURLs(changedPath, invalidatedFiles)

	if len(affectedPageURLs) == 0 {
		// If smart reload found no affected pages, check if we're in a "no routes" state
		// (e.g. no manifest yet). In this case, fallback to broadcasting to all clients.
		s.mu.RLock()
		noDemoRoutes := len(s.demoRoutes) == 0
		s.mu.RUnlock()

		if noDemoRoutes {
			s.logger.Debug("No demo routes found, falling back to broadcast all for %s", relPath)
			files := []string{relPath}
			if err := s.BroadcastReload(files, "file-change-fallback"); err != nil {
				s.logger.Error("Failed to broadcast reload: %v", err)
			}
			return
		}

		s.logger.Debug("No pages affected by changes to %s", relPath)
		return
	}

	// Broadcast reload only to affected pages
	files := []string{relPath}
	msgBytes, err := s.CreateReloadMessage(files, "file-change")
	if err != nil {
		s.logger.Error("Failed to create reload message: %v", err)
		return
	}

	if s.wsManager != nil {
		err = s.wsManager.BroadcastToPages(msgBytes, affectedPageURLs)
		if err != nil {
			s.logger.Error("Failed to broadcast reload: %v", err)
		}
	}
}

// handleFileChanges listens for file change events and triggers reload
func (s *Server) handleFileChanges() {
	if s.watcher == nil {
		return
	}

	for event := range s.watcher.Events() {
		// Process all files in the batched event
		filesToProcess := event.Paths
		if len(filesToProcess) == 0 {
			filesToProcess = []string{event.Path}
		}

		// Check if path resolver needs rebuild (tsconfig.json, config files changed)
		if s.shouldRebuildPathResolver(filesToProcess) {
			s.logger.Info("Rebuilding path resolver due to source file change")
			if err := s.rebuildPathResolver(); err != nil {
				s.logger.Error("Failed to rebuild path resolver: %v", err)
				// Continue processing - don't block other hot-reload logic
			}
			// Note: rebuildPathResolver() handles broadcasting reload message
		}

		// Filter to only relevant source files and collect TS/JS files
		relevantFiles, tsJsFiles := s.filterRelevantFiles(filesToProcess)
		if len(relevantFiles) == 0 {
			continue
		}

		// Use first file for display/logging purposes
		changedPath := relevantFiles[0]
		relPath := changedPath
		if s.watchDir != "" {
			if rel, err := filepath.Rel(s.watchDir, changedPath); err == nil {
				relPath = rel
			}
		}

		// Resolve .js to .ts if source exists and log
		displayPath := s.resolveSourceFile(relPath)
		s.logger.Info("File changed: %s", displayPath)

		// Collect affected files from transform cache and module graph
		invalidatedFiles := s.collectAffectedFiles(changedPath)

		// Regenerate manifest if TS/JS files changed
		s.regenerateManifestIfNeeded(tsJsFiles)

		// Regenerate import map if package.json or file structure changed
		s.regenerateImportMapIfNeeded(event)

		// Broadcast smart reload to affected pages
		s.broadcastSmartReload(changedPath, relPath, invalidatedFiles)
	}
}

// getModuleGraphAffectedFiles queries the generate session's dependency tracker
// to find all modules transitively affected by a file change.
// Returns filesystem paths of affected modules.
func (s *Server) getModuleGraphAffectedFiles(changedPath string) []string {
	s.mu.RLock()
	session := s.generateSession
	watchDir := s.watchDir
	s.mu.RUnlock()

	if session == nil || watchDir == "" {
		return nil
	}

	// Get the workspace context from the session
	ctx := session.WorkspaceContext()
	if ctx == nil {
		return nil
	}

	// Convert filesystem path to module path
	modulePath, err := ctx.FSPathToModule(changedPath)
	if err != nil {
		s.logger.Debug("Failed to convert FS path to module path for %s: %v", changedPath, err)
		return nil
	}

	// Query the dependency tracker
	depTracker := session.DependencyTracker()
	if depTracker == nil {
		return nil
	}

	affectedModulePaths := depTracker.GetModulesAffectedByFiles([]string{modulePath})
	if len(affectedModulePaths) == 0 {
		return nil
	}

	// Convert module paths back to filesystem paths
	affectedFSPaths := make([]string, 0, len(affectedModulePaths))
	for _, modPath := range affectedModulePaths {
		fsPath := ctx.ModulePathToFS(modPath)
		affectedFSPaths = append(affectedFSPaths, fsPath)
	}

	s.logger.Debug("Module graph analysis: %d modules affected by %s", len(affectedFSPaths), modulePath)
	return affectedFSPaths
}

// buildAffectedFilesMap creates a map of all files affected by a change
// Includes the changed file and all transitively invalidated files
// Stores both relative and absolute forms for matching flexibility
func (s *Server) buildAffectedFilesMap(changedPath string, invalidatedFiles []string, watchDir string) map[string]bool {
	affectedFiles := make(map[string]bool)

	// Convert changed path to relative
	if rel, err := filepath.Rel(watchDir, changedPath); err == nil {
		affectedFiles[rel] = true
		affectedFiles["/"+rel] = true // Also store with leading slash
		s.logger.Debug("Changed file (relative): %s", rel)
	} else {
		affectedFiles[changedPath] = true
		s.logger.Debug("Could not make relative, using absolute: %s", changedPath)
	}

	// Add all transitively invalidated files (also convert to relative)
	for _, path := range invalidatedFiles {
		if rel, err := filepath.Rel(watchDir, path); err == nil {
			affectedFiles[rel] = true
			affectedFiles["/"+rel] = true
		} else {
			affectedFiles[path] = true
		}
	}

	s.logger.Debug("Affected files map keys: %v", affectedFiles)
	s.logger.Debug("Found %d affected file paths to check", len(affectedFiles))

	return affectedFiles
}

// isRouteAffectedByFile checks if a demo route is affected by any of the changed files
// Returns true if the route's HTML file is affected or if it imports any affected files
func (s *Server) isRouteAffectedByFile(routePath string, routeEntry *types.DemoRouteEntry, affectedFiles map[string]bool, watchDir string) bool {
	// First, check if the changed file IS this demo HTML file
	if affectedFiles[routeEntry.FilePath] ||
		affectedFiles["/"+routeEntry.FilePath] ||
		affectedFiles[strings.TrimPrefix(routeEntry.FilePath, "/")] {
		s.logger.Debug("Page %s is the changed file itself", routePath)
		return true
	}

	htmlPath := filepath.Join(watchDir, routeEntry.FilePath)

	// Extract imports from HTML
	imports, err := s.extractModuleImports(htmlPath)
	if err != nil {
		s.logger.Debug("Failed to parse %s: %v", routeEntry.FilePath, err)
		return false
	}

	if len(imports) == 0 {
		return false // No imports to check
	}

	// Get the directory of the HTML file for resolving relative imports
	demoDir := filepath.Dir(routeEntry.FilePath)

	// Resolve imports to file paths and check if any match affected files
	for _, importSpec := range imports {
		resolvedPaths := s.resolveImportToPath(importSpec, demoDir)
		if len(resolvedPaths) == 0 {
			continue
		}

		for _, resolvedPath := range resolvedPaths {
			// Normalize the resolved path
			normalizedResolved := filepath.Clean(resolvedPath)

			// Also try with .ts extension instead of .js
			normalizedResolvedTS := normalizedResolved
			if strings.HasSuffix(normalizedResolved, ".js") {
				normalizedResolvedTS = normalizedResolved[:len(normalizedResolved)-3] + ".ts"
			}

			// Check if this resolved path matches any affected file
			if affectedFiles[normalizedResolved] ||
				affectedFiles[normalizedResolvedTS] ||
				affectedFiles[strings.TrimPrefix(normalizedResolved, "/")] ||
				affectedFiles[strings.TrimPrefix(normalizedResolvedTS, "/")] {
				s.logger.Debug("Page %s imports affected file %s (via %s)", routePath, normalizedResolved, importSpec)
				return true
			}
		}
	}

	return false
}

// getAffectedPageURLs returns page URLs that import the changed file or its dependents
func (s *Server) getAffectedPageURLs(changedPath string, invalidatedFiles []string) []string {
	s.logger.Debug("getAffectedPageURLs called with changedPath=%s, invalidatedFiles=%d", changedPath, len(invalidatedFiles))

	s.mu.RLock()
	watchDir := s.watchDir
	demoRoutes := s.demoRoutes
	s.mu.RUnlock()

	if demoRoutes == nil {
		s.logger.Debug("No demo routes available for smart reload")
		return nil
	}

	// Build map of affected files with all path variants
	affectedFiles := s.buildAffectedFilesMap(changedPath, invalidatedFiles, watchDir)

	s.logger.Debug("Checking %d demo routes for affected imports", len(demoRoutes))
	affectedPages := make([]string, 0)

	// Check each route to see if it's affected
	for routePath, routeEntry := range demoRoutes {
		if s.isRouteAffectedByFile(routePath, routeEntry, affectedFiles, watchDir) {
			affectedPages = append(affectedPages, routePath)
		}
	}

	if len(affectedPages) == 0 {
		s.logger.Debug("No pages import any of the affected files")
	}

	return affectedPages
}
