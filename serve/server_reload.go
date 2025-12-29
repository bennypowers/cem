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

		// Check if any changed files are pathResolver source files (tsconfig.json, config files)
		// If so, rebuild the pathResolver with updated URL rewrites
		pathResolverNeedsRebuild := false
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
					pathResolverNeedsRebuild = true
					s.logger.Debug("PathResolver source file changed: %s", sourceFile)
					break
				}
			}
			if pathResolverNeedsRebuild {
				break
			}
		}

		if pathResolverNeedsRebuild {
			s.logger.Info("Rebuilding path resolver due to source file change")
			if err := s.rebuildPathResolver(); err != nil {
				s.logger.Error("Failed to rebuild path resolver: %v", err)
				// Continue processing - don't block other hot-reload logic
			}
			// Note: rebuildPathResolver() handles broadcasting reload message
		}

		// Filter to only relevant source files
		relevantFiles := make([]string, 0, len(filesToProcess))
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
		}

		// Skip if no relevant files
		if len(relevantFiles) == 0 {
			continue
		}

		// Process the first relevant file (for manifest regeneration, etc.)
		// In the future, we could optimize to only regenerate once for multiple files
		changedPath := relevantFiles[0]
		relPath := changedPath
		if s.watchDir != "" {
			if rel, err := filepath.Rel(s.watchDir, changedPath); err == nil {
				relPath = rel
			}
		}

		// Resolve .js to .ts if source exists
		displayPath := s.resolveSourceFile(relPath)
		s.logger.Info("File changed: %s", displayPath)

		ext := filepath.Ext(changedPath)

		// Collect all affected files from both transform cache and module graph
		// This gives us complete dependency tracking:
		// - Transform cache: tracks dependencies from esbuild transforms
		// - Module graph: tracks ALL module imports from manifest generation
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
		var invalidatedFiles []string
		for file := range affectedFiles {
			invalidatedFiles = append(invalidatedFiles, file)
		}

		s.logger.Debug("Total affected files: %d (transform cache + module graph)", len(invalidatedFiles))

		// Regenerate manifest if a source file changed
		if ext == ".ts" || ext == ".js" {
			manifestStart := time.Now()
			s.logger.Debug("Regenerating manifest incrementally for %s file change...", ext)
			manifestSize, err := s.RegenerateManifestIncremental([]string{changedPath})
			manifestDuration := time.Since(manifestStart)
			if err != nil {
				s.logger.Error("Failed to regenerate manifest incrementally: %v", err)
				// Continue anyway - we still want to reload the page
			} else {
				s.logger.Info("Manifest regenerated incrementally (%d bytes) in %v", manifestSize, manifestDuration)
			}

			// Note: Import map regeneration is skipped for .ts/.js changes
			// Import maps are built from package.json exports, not source files
			// They only need to be regenerated when package.json changes
		}

		// Regenerate import map only when necessary and enabled:
		// - package.json was modified (exports may have changed)
		// - files were created (new modules may need to be mapped)
		// - files were deleted (old modules may need to be unmapped)
		if s.config.ImportMap.Generate && (event.HasPackageJSON || event.HasCreates || event.HasDeletes) {
			importMapStart := time.Now()
			s.mu.Lock()
			if s.isWorkspace {
				// Workspace mode: regenerate workspace import map
				importMap, err := importmappkg.Generate(s.workspaceRoot, &importmappkg.Config{
					InputMapPath:      s.config.ImportMap.OverrideFile,
					ConfigOverride:    s.buildConfigOverride(),
					WorkspacePackages: s.workspacePackages,
					Logger:            s.logger,
					FS:                s.fs,
				})
				if err != nil {
					s.logger.Warning("Failed to regenerate workspace import map: %v", err)
				} else {
					s.importMap = importMap
					importMapDuration := time.Since(importMapStart)
					s.logger.Info("Regenerated workspace import map in %v", importMapDuration)
				}
			} else {
				// Single-package mode: regenerate import map
				importMap, err := importmappkg.Generate(s.watchDir, &importmappkg.Config{
					InputMapPath:   s.config.ImportMap.OverrideFile,
					ConfigOverride: s.buildConfigOverride(),
					Logger:         s.logger,
					FS:             s.fs,
				})
				if err != nil {
					s.logger.Warning("Failed to regenerate import map: %v", err)
				} else {
					s.importMap = importMap
					importMapDuration := time.Since(importMapStart)
					s.logger.Info("Regenerated import map in %v", importMapDuration)
				}
			}
			s.mu.Unlock()
		}

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
				// Create a "broadcast all" message (using empty affected list effectively broadcasts to all if we change logic,
				// but here we just pass the file and rely on client side or simply assume all pages need reload)
				// Actually BroadcastToPages with nil/empty list skips.
				// We should use Broadcast() instead.
				files := []string{relPath}
				if err := s.BroadcastReload(files, "file-change-fallback"); err != nil {
					s.logger.Error("Failed to broadcast reload: %v", err)
				}
				continue
			}

			s.logger.Debug("No pages affected by changes to %s", relPath)
			continue
		}

		// Broadcast reload only to affected pages
		files := []string{relPath}
		msgBytes, err := s.CreateReloadMessage(files, "file-change")
		if err != nil {
			s.logger.Error("Failed to create reload message: %v", err)
			continue
		}

		if s.wsManager != nil {
			err = s.wsManager.BroadcastToPages(msgBytes, affectedPageURLs)
			if err != nil {
				s.logger.Error("Failed to broadcast reload: %v", err)
			}
		}
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

// getAffectedPageURLs returns page URLs that import the changed file or its dependents
func (s *Server) getAffectedPageURLs(changedPath string, invalidatedFiles []string) []string {
	s.logger.Debug("getAffectedPageURLs called with changedPath=%s, invalidatedFiles=%d", changedPath, len(invalidatedFiles))
	s.mu.RLock()
	watchDir := s.watchDir
	s.mu.RUnlock()

	// Get all files affected by this change (from cache invalidation)
	// Convert absolute paths to relative paths for comparison
	affectedFiles := make(map[string]bool)

	// Convert changed path to relative
	if rel, err := filepath.Rel(watchDir, changedPath); err == nil {
		affectedFiles[rel] = true
		affectedFiles["/"+rel] = true // Also store with leading slash
		s.logger.Debug("Changed file (relative): %s", rel)
		s.logger.Debug("Affected files map keys: %v", affectedFiles)
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

	s.logger.Debug("Found %d affected file paths to check", len(affectedFiles))

	// Get demo routes to find HTML file paths
	s.mu.RLock()
	demoRoutes := s.demoRoutes
	s.mu.RUnlock()

	if demoRoutes == nil {
		s.logger.Debug("No demo routes available for smart reload")
		return nil
	}

	s.logger.Debug("Checking %d demo routes for affected imports", len(demoRoutes))
	affectedPages := make([]string, 0)

	// For each demo route, check if it imports any affected files
	for routePath, routeEntry := range demoRoutes {
		// First, check if the changed file IS this demo HTML file
		if affectedFiles[routeEntry.FilePath] ||
			affectedFiles["/"+routeEntry.FilePath] ||
			affectedFiles[strings.TrimPrefix(routeEntry.FilePath, "/")] {
			affectedPages = append(affectedPages, routePath)
			s.logger.Debug("Page %s is the changed file itself", routePath)
			continue // Already added, skip import checking for this route
		}

		htmlPath := filepath.Join(watchDir, routeEntry.FilePath)

		// Extract imports from HTML
		imports, err := s.extractModuleImports(htmlPath)
		if err != nil {
			s.logger.Debug("Failed to parse %s: %v", routeEntry.FilePath, err)
			continue
		}

		if len(imports) == 0 {
			continue // Skip routes with no imports
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
					affectedPages = append(affectedPages, routePath)
					s.logger.Debug("Page %s imports affected file %s (via %s)", routePath, normalizedResolved, importSpec)
					goto nextRoute // Found a match, move to next route
				}
			}
		}
	nextRoute:
	}

	if len(affectedPages) == 0 {
		s.logger.Debug("No pages import any of the affected files")
	}

	return affectedPages
}
