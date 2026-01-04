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

	G "bennypowers.dev/cem/generate"
	"bennypowers.dev/cem/serve/middleware"
	importmappkg "bennypowers.dev/cem/serve/middleware/importmap"
	"bennypowers.dev/cem/serve/middleware/routes"
	W "bennypowers.dev/cem/workspace"
)

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

// discoverWorkspacePackages finds all packages in workspace with manifests
func discoverWorkspacePackages(rootDir string) ([]middleware.WorkspacePackage, error) {
	packages, err := W.LoadWorkspaceManifests(rootDir)
	if err != nil {
		return nil, err
	}

	// Convert workspace.PackageWithManifest to middleware.WorkspacePackage
	contexts := make([]middleware.WorkspacePackage, len(packages))
	for i, pkg := range packages {
		contexts[i] = middleware.WorkspacePackage{
			Name:     pkg.Name,
			Path:     pkg.Path,
			Manifest: pkg.Manifest,
		}
	}

	return contexts, nil
}

// InitializeWorkspaceMode detects and initializes workspace mode if applicable
func (s *Server) InitializeWorkspaceMode() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.watchDir == "" {
		return fmt.Errorf("no watch directory set")
	}

	// Check if this is a workspace
	if !W.IsWorkspaceMode(s.watchDir) {
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

// generateInitialWorkspaceManifests generates fresh manifests for all workspace packages
// Used during server initialization. Returns packages with freshly generated manifests.
func (s *Server) generateInitialWorkspaceManifests(watchDir string) ([]middleware.WorkspacePackage, error) {
	// Find all workspace packages
	packageDirs, err := W.FindPackagesWithManifests(watchDir)
	if err != nil {
		return nil, fmt.Errorf("finding workspace packages: %w", err)
	}

	if len(packageDirs) == 0 {
		return nil, nil // Return empty list, not error
	}

	// Generate fresh manifests for each package
	packages := make([]middleware.WorkspacePackage, 0, len(packageDirs))

	for _, pkgInfo := range packageDirs {
		// Create workspace context for this package
		workspace := W.NewFileSystemWorkspaceContext(pkgInfo.Path)
		if err := workspace.Init(); err != nil {
			s.logger.Warning("Failed to initialize workspace for package %s: %v", pkgInfo.Name, err)
			continue
		}

		// Create generate session
		session, err := G.NewGenerateSession(workspace)
		if err != nil {
			s.logger.Warning("Failed to create session for package %s: %v", pkgInfo.Name, err)
			continue
		}

		// Generate manifest
		ctx := context.Background()
		pkg, err := session.GenerateFullManifest(ctx)
		if err != nil {
			s.logger.Warning("Failed to generate manifest for package %s: %v", pkgInfo.Name, err)
			session.Close()
			continue
		}

		// Marshal to JSON
		manifestBytes, err := json.MarshalIndent(pkg, "", "  ")
		session.Close()
		if err != nil {
			s.logger.Warning("Failed to marshal manifest for package %s: %v", pkgInfo.Name, err)
			continue
		}

		packages = append(packages, middleware.WorkspacePackage{
			Name:     pkgInfo.Name,
			Path:     pkgInfo.Path,
			Manifest: manifestBytes,
		})
	}

	return packages, nil
}

// regenerateWorkspaceManifests regenerates manifests for all workspace packages
// and rebuilds the routing table. Used during file change events in workspace mode.
func (s *Server) regenerateWorkspaceManifests() (int, error) {
	s.mu.RLock()
	watchDir := s.watchDir
	s.mu.RUnlock()

	if watchDir == "" {
		return 0, fmt.Errorf("no watch directory set")
	}

	// Find all workspace packages
	packageDirs, err := W.FindPackagesWithManifests(watchDir)
	if err != nil {
		return 0, fmt.Errorf("finding workspace packages: %w", err)
	}

	if len(packageDirs) == 0 {
		s.logger.Warning("No packages found in workspace during regeneration")
		return 0, nil
	}

	// Generate fresh manifests for each package
	packages := make([]middleware.WorkspacePackage, 0, len(packageDirs))
	totalSize := 0

	for _, pkgInfo := range packageDirs {
		// Create workspace context for this package
		workspace := W.NewFileSystemWorkspaceContext(pkgInfo.Path)
		if err := workspace.Init(); err != nil {
			s.logger.Warning("Failed to initialize workspace for package %s: %v", pkgInfo.Name, err)
			continue
		}

		// Create generate session
		session, err := G.NewGenerateSession(workspace)
		if err != nil {
			s.logger.Warning("Failed to create session for package %s: %v", pkgInfo.Name, err)
			continue
		}

		// Generate manifest
		ctx := context.Background()
		pkg, err := session.GenerateFullManifest(ctx)
		if err != nil {
			s.logger.Warning("Failed to generate manifest for package %s: %v", pkgInfo.Name, err)
			session.Close()
			continue
		}

		// Marshal to JSON
		manifestBytes, err := json.MarshalIndent(pkg, "", "  ")
		session.Close()
		if err != nil {
			s.logger.Warning("Failed to marshal manifest for package %s: %v", pkgInfo.Name, err)
			continue
		}

		packages = append(packages, middleware.WorkspacePackage{
			Name:     pkgInfo.Name,
			Path:     pkgInfo.Path,
			Manifest: manifestBytes,
		})

		totalSize += len(manifestBytes)
	}

	if len(packages) == 0 {
		return 0, fmt.Errorf("failed to regenerate any package manifests")
	}

	// Build routing table from regenerated packages
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
		s.logger.Warning("Failed to build workspace routing table: %w", err)
		// Continue anyway - still update packages even if routing fails
	}

	// Update server state under write lock
	s.mu.Lock()
	defer s.mu.Unlock()

	s.workspacePackages = packages
	s.demoRoutes = workspaceRoutingTable
	s.logger.Debug("Regenerated %d workspace package manifests, built routing table with %d demo routes",
		len(packages), len(workspaceRoutingTable))

	return totalSize, nil
}
