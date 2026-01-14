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
	"fmt"
	"maps"
	"os"
	"path/filepath"

	"bennypowers.dev/cem/cmd/config"
	"bennypowers.dev/cem/serve/middleware"
	importmappkg "bennypowers.dev/cem/serve/middleware/importmap"
	"bennypowers.dev/cem/serve/middleware/transform"
)

// packageJSON represents the structure we need from package.json for workspace detection
type packageJSON struct {
	Workspaces any `json:"workspaces"` // Can be []string or object with "packages" field
}

// findWorkspaceRootForServe walks up the directory tree to find the workspace root
// Returns the directory containing node_modules or workspace configuration
// Stops at git repository boundaries to avoid breaking out of submodules
//
// Note: This is a simplified version of workspace.FindWorkspaceRoot that only looks for
// node_modules or workspace config without the VCS preference logic. It shares logic with
// serve/middleware/importmap/generate.go:findWorkspaceRoot (which accepts FileSystem for testing).
// Consider consolidating these implementations in the future.
func findWorkspaceRootForServe(startDir string) string {
	dir := startDir
	for {
		// Check if node_modules exists in this directory
		nodeModulesPath := filepath.Join(dir, "node_modules")
		if stat, err := os.Stat(nodeModulesPath); err == nil && stat.IsDir() {
			return dir
		}

		// Check if there's a package.json with workspaces field
		pkgPath := filepath.Join(dir, "package.json")
		if data, err := os.ReadFile(pkgPath); err == nil {
			var pkg packageJSON
			if json.Unmarshal(data, &pkg) == nil && pkg.Workspaces != nil {
				return dir
			}
		}

		// Stop if we've reached a git repository root (don't go higher)
		gitDir := filepath.Join(dir, ".git")
		if stat, err := os.Stat(gitDir); err == nil && stat.IsDir() {
			// Hit git boundary without finding workspace root, return start dir
			return startDir
		}

		// Move up one directory
		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached filesystem root, return original directory
			return startDir
		}
		dir = parent
	}
}

// URLRewrites returns the configured URL rewrites for request path resolution
func (s *Server) URLRewrites() []config.URLRewrite {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.urlRewrites
}

// PathResolver returns the cached path resolver for efficient URL rewriting
func (s *Server) PathResolver() middleware.PathResolver {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.pathResolver
}

// buildConfigOverride converts serve config override to import map config override
// Creates defensive copies to prevent shared mutable state.
func (s *Server) buildConfigOverride() *importmappkg.ImportMap {
	if len(s.config.ImportMap.Override.Imports) == 0 && len(s.config.ImportMap.Override.Scopes) == 0 {
		return nil
	}

	result := &importmappkg.ImportMap{
		Imports: make(map[string]string, len(s.config.ImportMap.Override.Imports)),
		Scopes:  make(map[string]map[string]string, len(s.config.ImportMap.Override.Scopes)),
	}

	// Copy imports
	maps.Copy(result.Imports, s.config.ImportMap.Override.Imports)

	// Deep copy scopes
	for scopeKey, scopeMap := range s.config.ImportMap.Override.Scopes {
		result.Scopes[scopeKey] = make(map[string]string, len(scopeMap))
		maps.Copy(result.Scopes[scopeKey], scopeMap)
	}

	return result
}

// parseConfigFileRewrites parses URL rewrites from the config file.
// Returns rewrites and source file paths for hot-reload tracking.
// Handles validation and logs warnings for invalid rewrites.
func (s *Server) parseConfigFileRewrites() ([]config.URLRewrite, []string) {
	if s.config.ConfigFile == "" {
		return nil, nil
	}

	rewrites, sourceFiles, err := transform.ParseConfigFileURLRewrites(s.config.ConfigFile, s.fs)
	if err != nil {
		s.logger.Debug("No config file URL rewrites: %v", err)
		return nil, nil
	}

	if len(rewrites) == 0 {
		return nil, sourceFiles
	}

	// Validate URL rewrites to prevent runtime panics
	if err := transform.ValidateURLRewrites(rewrites); err != nil {
		s.logger.Warning("Invalid URL rewrites in config file: %v", err)
		return nil, sourceFiles
	}

	s.logger.Debug("Parsed %d URL rewrites from %s", len(rewrites), s.config.ConfigFile)
	for _, r := range rewrites {
		s.logger.Debug("  %s -> %s", r.URLPattern, r.URLTemplate)
	}

	return rewrites, sourceFiles
}

// mergeURLRewrites merges config file rewrites with tsconfig rewrites.
// Config file rewrites take precedence (prepended first for first-match-wins).
func (s *Server) mergeURLRewrites(configRewrites, tsconfigRewrites []config.URLRewrite) []config.URLRewrite {
	if len(configRewrites) == 0 {
		return tsconfigRewrites
	}

	s.logger.Debug("Adding %d URL rewrites from config file (will override tsconfig)", len(configRewrites))
	// Prepend config rewrites so they're tried first (first-match-wins in resolver)
	return append(configRewrites, tsconfigRewrites...)
}

// RebuildPathResolverForTest exposes rebuildPathResolver for testing.
// This is a test-only wrapper that allows tests to trigger path resolver rebuilds.
func (s *Server) RebuildPathResolverForTest() error {
	return s.rebuildPathResolver()
}

// rebuildPathResolver rebuilds the path resolver when tsconfig or config files change.
// This method is called during hot-reload to update URL rewrites without restarting the server.
func (s *Server) rebuildPathResolver() error {
	s.mu.Lock()

	if s.watchDir == "" {
		s.mu.Unlock()
		return fmt.Errorf("no watch directory set")
	}

	// Re-parse config file URL rewrites from disk (hot-reload needs fresh file contents)
	configURLRewrites, configSourceFiles := s.parseConfigFileRewrites()

	// Re-parse tsconfig files from disk (hot-reload needs fresh file contents)
	// Similar to SetWatchDir, but we re-read files instead of using cached config
	tsconfigPaths := []string{
		filepath.Join(s.watchDir, "tsconfig.settings.json"),
		filepath.Join(s.watchDir, "tsconfig.json"),
	}

	var urlRewrites []config.URLRewrite
	var tsconfigSourceFiles []string

	for _, tsconfigPath := range tsconfigPaths {
		rewrites, sourceFiles, err := transform.ParseTsConfig(tsconfigPath, s.fs)
		if err == nil {
			urlRewrites = rewrites
			tsconfigSourceFiles = sourceFiles
			if len(rewrites) > 0 {
				s.logger.Debug("Rebuilt URL rewrites from %s", tsconfigPath)
				for _, r := range rewrites {
					s.logger.Debug("  %s -> %s", r.URLPattern, r.URLTemplate)
				}
			}
			break
		}
	}

	// Merge URL rewrites: config file takes precedence over tsconfig
	urlRewrites = s.mergeURLRewrites(configURLRewrites, urlRewrites)

	s.logger.Info("Rebuilding path resolver with %d URL rewrites", len(urlRewrites))
	s.urlRewrites = urlRewrites

	// Track both config file and tsconfig source files for hot-reload
	s.pathResolverSourceFiles = append(configSourceFiles, tsconfigSourceFiles...)

	// Rebuild path resolver
	s.pathResolver = transform.NewPathResolver(s.watchDir, s.urlRewrites, s.fs, s.logger)

	// Invalidate transform cache (path resolution changed, so cached transforms may be stale)
	if s.transformCache != nil {
		s.transformCache.Clear()
		s.logger.Debug("Cleared transform cache after path resolver rebuild")
	}

	// Unlock before calling setupMiddleware to avoid deadlock
	// setupMiddleware calls accessor methods that acquire s.mu.RLock()
	s.mu.Unlock()

	// Rebuild middleware pipeline with new pathResolver (no lock needed)
	s.setupMiddleware()

	// Broadcast reload to connected clients (URL rewrites changed)
	// This happens after successful rebuild to notify browsers to refresh
	if err := s.BroadcastReload([]string{"config"}, "url-rewrites-changed"); err != nil {
		s.logger.Error("Failed to broadcast reload after path resolver rebuild: %v", err)
		// Don't return error - rebuild succeeded, broadcast failure is not critical
	}

	return nil
}

// SetWatchDir sets the directory to watch for file changes
func (s *Server) SetWatchDir(dir string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Convert to absolute path for consistent behavior
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return fmt.Errorf("converting to absolute path: %w", err)
	}

	s.watchDir = absDir

	// Detect workspace root for monorepo support (only in single-package mode)
	// In workspace mode, server already runs from workspace root
	// In single-package mode, store workspace root separately so /node_modules/
	// can be served from workspace root while keeping package isolation
	if !s.isWorkspace {
		workspaceRoot := findWorkspaceRootForServe(absDir)
		if workspaceRoot != absDir {
			s.logger.Debug("Detected workspace root at %s", workspaceRoot)
			s.logger.Debug("Will serve /node_modules/ from workspace root while keeping package isolation")
			s.workspaceRoot = workspaceRoot
		} else {
			s.workspaceRoot = ""
		}
	}

	// Parse config file for URL rewrites (if config file path was set)
	configURLRewrites, configSourceFiles := s.parseConfigFileRewrites()

	// Load tsconfig - try tsconfig.settings.json first (common in monorepos),
	// then fall back to tsconfig.json
	tsconfigPaths := []string{
		filepath.Join(absDir, "tsconfig.settings.json"),
		filepath.Join(absDir, "tsconfig.json"),
	}

	loaded := false
	for _, tsconfigPath := range tsconfigPaths {
		if data, err := s.fs.ReadFile(tsconfigPath); err == nil {
			s.tsconfigRaw = string(data)
			s.logger.Debug("Loaded TypeScript config from %s", tsconfigPath)
			loaded = true
			break
		}
	}

	if !loaded {
		s.tsconfigRaw = ""
		s.logger.Debug("No tsconfig found, using default transform settings")
	}

	// Parse tsconfig for URL rewrites (supports src/dist separation)
	var urlRewrites []config.URLRewrite
	var tsconfigSourceFiles []string // Track all files used
	for _, tsconfigPath := range tsconfigPaths {
		rewrites, sourceFiles, err := transform.ParseTsConfig(tsconfigPath, s.fs)
		if err == nil {
			urlRewrites = rewrites
			tsconfigSourceFiles = sourceFiles // Store the files
			if len(rewrites) > 0 {
				s.logger.Debug("Extracted URL rewrites from %s", tsconfigPath)
				for _, r := range rewrites {
					s.logger.Debug("  %s -> %s", r.URLPattern, r.URLTemplate)
				}
			}
			break
		}
	}

	// Merge URL rewrites: config file takes precedence over tsconfig
	s.urlRewrites = s.mergeURLRewrites(configURLRewrites, urlRewrites)

	// Track both config file and tsconfig source files for hot-reload
	s.pathResolverSourceFiles = append(configSourceFiles, tsconfigSourceFiles...)

	// Initialize path resolver once (cached for all requests)
	s.pathResolver = transform.NewPathResolver(s.watchDir, s.urlRewrites, s.fs, s.logger)

	// Generate import map for single-package mode
	// (Workspace mode generates in InitializeWorkspaceMode instead)
	// Use GenerateWithGraph to enable incremental updates on hot-reload
	if !s.isWorkspace && s.config.ImportMap.Generate {
		result, err := importmappkg.GenerateWithGraph(dir, &importmappkg.Config{
			InputMapPath:   s.config.ImportMap.OverrideFile,
			ConfigOverride: s.buildConfigOverride(),
			Logger:         s.logger,
			FS:             s.fs,
		})
		if err != nil {
			s.logger.Warning("Failed to generate import map: %v", err)
			s.importMap = nil
			s.importMapGraph = nil
		} else {
			s.importMap = result.ImportMap
			s.importMapGraph = result.Graph
			s.logger.Debug("Generated import map for single-package mode")
		}
	} else if !s.config.ImportMap.Generate {
		s.logger.Debug("Import map generation disabled")
	}

	// Rebuild middleware pipeline with new pathResolver
	s.setupMiddleware()

	return nil
}
