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
	"context"
	"encoding/json"
	"fmt"
	"time"

	G "bennypowers.dev/cem/generate"
	C "bennypowers.dev/cem/internal/config"
	W "bennypowers.dev/cem/internal/workspace"
	"bennypowers.dev/cem/serve/middleware"
	importmappkg "bennypowers.dev/cem/serve/middleware/importmap"
	"bennypowers.dev/cem/serve/middleware/routes"
)

// contextWithShutdown creates a context that is cancelled when either:
// - The provided timeout is reached, or
// - The server begins shutting down
// This prevents resource leaks during long-running operations.
func (s *Server) contextWithShutdown(timeout time.Duration) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	// Spawn goroutine to cancel context if server shuts down
	go func() {
		select {
		case <-s.Done():
			cancel()
		case <-ctx.Done():
			// Context already cancelled/timed out
		}
	}()

	return ctx, cancel
}

// IsWorkspace returns true if server is running in workspace mode
func (s *Server) IsWorkspace() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.isWorkspace
}

// WorkspacePackages returns discovered workspace packages (workspace mode only)
func (s *Server) WorkspacePackages() []middleware.WorkspacePackage {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.workspacePackages
}

// DemoRoutes returns the pre-computed demo routing table (both workspace and single-package mode)
func (s *Server) DemoRoutes() map[string]*middleware.DemoRouteEntry {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.demoRoutes
}

// SourceControlRootURL returns the source control root URL for demo routing
func (s *Server) SourceControlRootURL() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.sourceControlRootURL
}

// DemoURLPrefix returns the URL path prefix to strip from demo URLs for local routing.
func (s *Server) DemoURLPrefix() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.demoURLPrefix
}

// InitializeWorkspaceMode detects and initializes workspace mode if applicable
func (s *Server) InitializeWorkspaceMode() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.watchDir == "" {
		return fmt.Errorf("no watch directory set")
	}

	// Check if this is a workspace
	if !W.IsWorkspaceMode(s.watchDir, s.fs) {
		s.isWorkspace = false
		return nil
	}

	// Generate fresh manifests for all workspace packages
	// This ensures we use current demo URLs instead of stale data from disk
	s.mu.Unlock() // Unlock before calling regenerateWorkspaceManifests
	packages, err := s.generateInitialWorkspaceManifests(s.watchDir)
	s.mu.Lock() // Re-lock for remaining initialization

	if err != nil {
		return fmt.Errorf("generating workspace manifests: %w", err)
	}

	// If no packages have customElements field, fall back to single-package mode
	// This handles workspaces where packages haven't been configured yet
	if len(packages) == 0 {
		s.logger.Debug("Workspace detected but no packages have customElements field - using single-package mode")
		s.isWorkspace = false
		return nil
	}

	s.logger.Debug("Detected workspace mode - found %d packages with manifests", len(packages))
	s.isWorkspace = true
	s.workspaceRoot = s.watchDir
	s.workspacePackages = packages

	// Build routing table from workspace packages
	pkgContexts := make([]routes.PackageContext, len(packages))
	for i, pkg := range packages {
		pkgContexts[i] = routes.PackageContext{
			Name:     pkg.Name,
			Path:     pkg.Path,
			Manifest: pkg.Manifest,
		}
	}
	workspaceRoutingTable, err := routes.BuildWorkspaceRoutingTable(pkgContexts)
	if err != nil {
		return fmt.Errorf("building workspace routing table: %w", err)
	}
	s.demoRoutes = workspaceRoutingTable
	s.logger.Debug("Built routing table with %d demo routes", len(workspaceRoutingTable))

	// Generate workspace import map using middleware package
	if s.config.ImportMap.Generate {
		importMap, err := importmappkg.Generate(s.workspaceRoot, &importmappkg.Config{
			InputMapPath:      s.config.ImportMap.OverrideFile,
			ConfigOverride:    s.buildConfigOverride(),
			WorkspacePackages: packages,
			Logger:            s.logger,
			FS:                s.fs,
		})
		if err != nil {
			s.logger.Warning("Failed to generate workspace import map: %v", err)
			s.importMap = nil
		} else {
			s.importMap = importMap
			s.logger.Info("Generated workspace import map")
		}
	} else {
		s.logger.Debug("Import map generation disabled")
	}

	return nil
}

// loadWorkspaceRootConfig loads and caches the root workspace config.
// Returns nil if watchDir is empty or config cannot be loaded.
func (s *Server) loadWorkspaceRootConfig() *C.CemConfig {
	if s.watchDir == "" {
		return nil
	}
	rootCtx := W.NewFileSystemWorkspaceContext(s.watchDir)
	if err := rootCtx.Init(); err != nil {
		s.logger.Debug("Root context init failed: %v", err)
		return nil
	}
	cfg, err := rootCtx.Config()
	if err != nil {
		s.logger.Debug("Root config load failed: %v", err)
		return nil
	}
	return cfg
}

// generateManifestForPackage generates a manifest for a single workspace package.
// rootCfg is the workspace root config (may be nil if unavailable).
func (s *Server) generateManifestForPackage(pkgInfo W.PackageInfo, rootCfg *C.CemConfig) (*middleware.WorkspacePackage, error) {
	s.logger.Debug("Generating manifest for workspace package %s", pkgInfo.Name)
	workspace := W.NewFileSystemWorkspaceContext(pkgInfo.Path, W.WithFileSystem(s.fs))
	if err := workspace.Init(); err != nil {
		return nil, fmt.Errorf("initializing workspace for %s: %w", pkgInfo.Name, err)
	}

	if rootCfg != nil {
		pkgCfg, err := workspace.Config()
		if err != nil {
			return nil, fmt.Errorf("loading config for %s: %w", pkgInfo.Name, err)
		}
		if len(rootCfg.Generate.Files) > 0 {
			resolved, err := W.ResolveWorkspaceFiles(s.watchDir, rootCfg.Generate.Files, pkgInfo.Path, s.fs)
			if err != nil {
				s.logger.Debug("ResolveWorkspaceFiles failed for %s: %v", pkgInfo.Name, err)
			} else {
				s.logger.Debug("Resolved %d files for %s", len(resolved), pkgInfo.Name)
				pkgCfg.Generate.Files = append(pkgCfg.Generate.Files, resolved...)
			}
		}
		if len(rootCfg.Generate.Exclude) > 0 {
			resolvedExclude, err := W.ResolveWorkspaceFiles(s.watchDir, rootCfg.Generate.Exclude, pkgInfo.Path, s.fs)
			if err == nil {
				pkgCfg.Generate.Exclude = append(pkgCfg.Generate.Exclude, resolvedExclude...)
			}
		}
		if rootCfg.Generate.DemoDiscovery.FileGlob != "" && pkgCfg.Generate.DemoDiscovery.FileGlob == "" {
			demoFiles, err := W.ResolveWorkspaceFiles(s.watchDir, []string{rootCfg.Generate.DemoDiscovery.FileGlob}, pkgInfo.Path, s.fs)
			if err == nil && len(demoFiles) > 0 {
				pkgCfg.Generate.DemoDiscovery = rootCfg.Generate.DemoDiscovery
				pkgCfg.Generate.DemoDiscovery.FileGlob = W.DerivePackageGlob(demoFiles)
			}
		}
	}

	// Create generate session
	session, err := G.NewGenerateSession(workspace, s.fs)
	if err != nil {
		return nil, fmt.Errorf("creating session for %s: %w", pkgInfo.Name, err)
	}
	defer session.Close()

	// Generate manifest with cancellable context (respects shutdown signal)
	genCtx, cancel := s.contextWithShutdown(30 * time.Second)
	defer cancel()
	pkg, err := session.GenerateFullManifest(genCtx)
	if err != nil {
		return nil, fmt.Errorf("generating manifest for %s: %w", pkgInfo.Name, err)
	}
	s.logger.Debug("Package %s: generated %d modules", pkgInfo.Name, len(pkg.Modules))

	// Marshal to JSON
	manifestBytes, err := json.MarshalIndent(pkg, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("marshaling manifest for %s: %w", pkgInfo.Name, err)
	}

	return &middleware.WorkspacePackage{
		Name:     pkgInfo.Name,
		Path:     pkgInfo.Path,
		Manifest: manifestBytes,
	}, nil
}

// generateInitialWorkspaceManifests generates fresh manifests for all workspace packages
// Used during server initialization. Returns packages with freshly generated manifests.
func (s *Server) generateInitialWorkspaceManifests(watchDir string) ([]middleware.WorkspacePackage, error) {
	// Find all workspace packages
	packageDirs, err := W.FindPackagesWithManifests(watchDir, s.fs)
	if err != nil {
		return nil, fmt.Errorf("finding workspace packages: %w", err)
	}

	if len(packageDirs) == 0 {
		return nil, nil // Return empty list, not error
	}

	rootCfg := s.loadWorkspaceRootConfig()

	packages := make([]middleware.WorkspacePackage, 0, len(packageDirs))

	for _, pkgInfo := range packageDirs {
		pkg, err := s.generateManifestForPackage(pkgInfo, rootCfg)
		if err != nil {
			s.logger.Warning("Failed to generate manifest for package %s: %v", pkgInfo.Name, err)
			continue
		}
		packages = append(packages, *pkg)
	}

	return packages, nil
}

// regenerateAffectedWorkspacePackages regenerates manifests only for packages
// affected by the changed files. Preserves cached manifests for unchanged packages.
// If changedFiles is nil, regenerates all packages (used for full regeneration).
// Returns the total manifest size in bytes and any error.
func (s *Server) regenerateAffectedWorkspacePackages(changedFiles []string) (int, error) {
	s.mu.RLock()
	watchDir := s.watchDir
	existingPackages := s.workspacePackages
	s.mu.RUnlock()

	if watchDir == "" {
		return 0, fmt.Errorf("no watch directory set")
	}

	// Determine which packages to regenerate
	var affectedPkgInfos []W.PackageInfo
	var err error

	if changedFiles == nil {
		// Full regeneration: get all packages
		affectedPkgInfos, err = W.FindPackagesWithManifests(watchDir, s.fs)
		if err != nil {
			return 0, fmt.Errorf("finding workspace packages: %w", err)
		}
		if len(affectedPkgInfos) == 0 {
			s.logger.Warning("No packages found in workspace during regeneration")
			return 0, nil
		}
	} else {
		// Incremental: only affected packages
		affectedPkgInfos, err = W.FindPackagesForFiles(watchDir, changedFiles, s.fs)
		if err != nil {
			return 0, fmt.Errorf("finding affected packages: %w", err)
		}
		if len(affectedPkgInfos) == 0 {
			s.logger.Debug("No packages affected by changed files, skipping regeneration")
			totalSize := 0
			for _, pkg := range existingPackages {
				totalSize += len(pkg.Manifest)
			}
			return totalSize, nil
		}
		s.logger.Debug("Regenerating %d affected package(s) out of %d total",
			len(affectedPkgInfos), len(existingPackages))
	}

	// Build a set of affected package paths for quick lookup
	affectedPaths := make(map[string]bool)
	for _, pkg := range affectedPkgInfos {
		affectedPaths[pkg.Path] = true
	}

	// Create new packages slice, preserving unaffected packages and regenerating affected ones
	packages := make([]middleware.WorkspacePackage, 0, max(len(existingPackages), len(affectedPkgInfos)))
	totalSize := 0

	// Copy over unaffected packages from the cache
	for _, existingPkg := range existingPackages {
		if !affectedPaths[existingPkg.Path] {
			packages = append(packages, existingPkg)
			totalSize += len(existingPkg.Manifest)
		}
	}

	rootCfg := s.loadWorkspaceRootConfig()

	for _, pkgInfo := range affectedPkgInfos {
		pkg, err := s.generateManifestForPackage(pkgInfo, rootCfg)
		if err != nil {
			s.logger.Warning("Failed to regenerate manifest for package %s: %v", pkgInfo.Name, err)
			// Try to preserve the old manifest if regeneration fails
			for _, existingPkg := range existingPackages {
				if existingPkg.Path == pkgInfo.Path {
					packages = append(packages, existingPkg)
					totalSize += len(existingPkg.Manifest)
					break
				}
			}
			continue
		}
		packages = append(packages, *pkg)
		totalSize += len(pkg.Manifest)
	}

	if len(packages) == 0 {
		return 0, fmt.Errorf("failed to regenerate any package manifests")
	}

	// Build routing table from updated packages
	pkgContexts := make([]routes.PackageContext, len(packages))
	for i, pkg := range packages {
		pkgContexts[i] = routes.PackageContext{
			Name:     pkg.Name,
			Path:     pkg.Path,
			Manifest: pkg.Manifest,
		}
	}

	workspaceRoutingTable, err := routes.BuildWorkspaceRoutingTable(pkgContexts)
	if err != nil {
		s.logger.Warning("Failed to build workspace routing table: %v", err)
	}

	// Update server state under write lock
	s.mu.Lock()
	defer s.mu.Unlock()

	s.workspacePackages = packages
	s.demoRoutes = workspaceRoutingTable

	preserved := len(packages) - len(affectedPkgInfos)
	if preserved > 0 {
		s.logger.Debug("Regenerated %d package(s), preserved %d cached, routing table has %d routes",
			len(affectedPkgInfos), preserved, len(workspaceRoutingTable))
	} else {
		s.logger.Debug("Regenerated %d package(s), routing table has %d routes",
			len(affectedPkgInfos), len(workspaceRoutingTable))
	}

	return totalSize, nil
}

// regenerateWorkspaceManifests regenerates manifests for all workspace packages.
func (s *Server) regenerateWorkspaceManifests() (int, error) {
	return s.regenerateAffectedWorkspacePackages(nil)
}
