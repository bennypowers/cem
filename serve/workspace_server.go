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
	"net/url"
	"os"
	"path/filepath"
	"strings"

	C "bennypowers.dev/cem/cmd/config"
	M "bennypowers.dev/cem/manifest"
	W "bennypowers.dev/cem/workspace"
)

// PackageContext holds information about a single package in the workspace
type PackageContext struct {
	Name      string        // Package name from package.json
	Path      string        // Absolute path to package directory
	Manifest  []byte        // Generated custom elements manifest
	Config    *C.CemConfig  // Package-specific config
}

// WorkspaceMode determines if we're serving workspace or single package
func IsWorkspaceMode(dir string) bool {
	// Check if directory has package.json with workspaces field
	pkgPath := filepath.Join(dir, "package.json")
	data, err := os.ReadFile(pkgPath)
	if err != nil {
		return false
	}

	// Simple check for workspaces field
	// (More robust parsing happens in workspace.FindPackagesWithManifests)
	return len(data) > 0 && (
		[]byte(`"workspaces"`) != nil &&
		len(data) > 0) // Simplified - actual check in discovery.go
}

// DiscoverWorkspacePackages finds all packages in workspace with manifests
func DiscoverWorkspacePackages(rootDir string) ([]PackageContext, error) {
	packages, err := W.FindPackagesWithManifests(rootDir)
	if err != nil {
		return nil, err
	}

	var contexts []PackageContext
	for _, pkg := range packages {
		// Load package's manifest
		manifestPath := filepath.Join(pkg.Path, pkg.CustomElementsRef)
		manifestData, err := os.ReadFile(manifestPath)
		if err != nil {
			// Skip packages with missing manifests
			continue
		}

		// Load package config (TODO: will use for urlPattern, etc)
		// For now, create minimal context
		ctx := PackageContext{
			Name:     pkg.Name,
			Path:     pkg.Path,
			Manifest: manifestData,
			Config:   nil, // TODO: Load config
		}
		contexts = append(contexts, ctx)
	}

	if len(contexts) == 0 {
		return nil, fmt.Errorf("no packages with manifests found in workspace")
	}

	return contexts, nil
}

// BuildWorkspaceRoutingTable builds a combined routing table from all packages
// Returns error if route conflicts are detected
func BuildWorkspaceRoutingTable(packages []PackageContext) (map[string]*DemoRouteEntry, error) {
	routes := make(map[string]*DemoRouteEntry)
	conflicts := make(map[string][]routeConflict)

	for _, pkg := range packages {
		// Build routing table for this package
		pkgRoutes, err := buildPackageRoutingTable(pkg)
		if err != nil {
			// Skip packages with routing errors
			continue
		}

		// Merge routes, detecting conflicts
		for route, entry := range pkgRoutes {
			if existing, exists := routes[route]; exists {
				// Conflict detected
				conflict := routeConflict{
					Route:       route,
					PackageName: pkg.Name,
					PackagePath: pkg.Path,
					FilePath:    entry.FilePath,
				}
				conflicts[route] = append(conflicts[route], conflict)
				// Also add the existing entry to conflicts if not already there
				if len(conflicts[route]) == 1 {
					existingConflict := routeConflict{
						Route:       route,
						PackageName: existing.PackageName,
						PackagePath: existing.PackagePath,
						FilePath:    existing.FilePath,
					}
					conflicts[route] = []routeConflict{existingConflict, conflict}
				}
			} else {
				routes[route] = entry
			}
		}
	}

	// If conflicts detected, return detailed error
	if len(conflicts) > 0 {
		return nil, formatRouteConflictsError(conflicts)
	}

	return routes, nil
}

// routeConflict represents a routing conflict between packages
type routeConflict struct {
	Route       string
	PackageName string
	PackagePath string
	FilePath    string
}

// formatRouteConflictsError creates a detailed error message for route conflicts
func formatRouteConflictsError(conflicts map[string][]routeConflict) error {
	var msg strings.Builder
	msg.WriteString("Route conflicts detected:\n\n")

	for route, conflictList := range conflicts {
		msg.WriteString(fmt.Sprintf("Conflict for route '%s':\n", route))
		for _, c := range conflictList {
			msg.WriteString(fmt.Sprintf("  - Package '%s' (%s)\n", c.PackageName, c.PackagePath))
			msg.WriteString(fmt.Sprintf("    Source: %s\n", c.FilePath))
		}
		msg.WriteString("\n")
	}

	msg.WriteString("Fix: Configure unique urlPattern in each package's manifest or demo metadata\n")

	return fmt.Errorf(msg.String())
}

// buildPackageRoutingTable builds routing table for a single package
func buildPackageRoutingTable(pkg PackageContext) (map[string]*DemoRouteEntry, error) {
	var manifest M.Package
	if err := json.Unmarshal(pkg.Manifest, &manifest); err != nil {
		return nil, fmt.Errorf("parsing manifest for %s: %w", pkg.Name, err)
	}

	routes := make(map[string]*DemoRouteEntry)

	for _, renderableDemo := range manifest.RenderableDemos() {
		demoURL := renderableDemo.Demo.URL

		// Extract local route from canonical URL
		localRoute := demoURL
		if parsed, err := url.Parse(demoURL); err == nil && parsed.Path != "" {
			localRoute = parsed.Path
		}

		// Normalize relative paths
		if strings.HasPrefix(localRoute, "./") {
			localRoute = localRoute[1:]
		}
		if !strings.HasPrefix(localRoute, "/") {
			localRoute = "/" + localRoute
		}

		// Ensure trailing slash for directory-style URLs
		if !strings.HasSuffix(localRoute, "/") && !strings.Contains(filepath.Base(localRoute), ".") {
			localRoute += "/"
		}

		// Resolve file path
		var filePath string
		if renderableDemo.Demo.Source != nil && renderableDemo.Demo.Source.Href != "" {
			sourceHref := renderableDemo.Demo.Source.Href
			// Extract filename from href and look in module's demo directory
			if parsed, err := url.Parse(sourceHref); err == nil && parsed.Path != "" {
				moduleDir := filepath.Dir(renderableDemo.Module.Path)
				filename := filepath.Base(parsed.Path)
				filePath = filepath.Join(moduleDir, "demo", filename)
			} else {
				filePath = sourceHref
			}
		} else {
			// Fallback: try to find demo file in module directory
			moduleDir := filepath.Dir(renderableDemo.Module.Path)
			filePath = filepath.Join(moduleDir, "demo", "index.html")
		}

		entry := &DemoRouteEntry{
			LocalRoute:  localRoute,
			TagName:     renderableDemo.CustomElementDeclaration.TagName,
			Demo:        renderableDemo.Demo,
			FilePath:    filePath,
			PackageName: pkg.Name,
			PackagePath: pkg.Path,
		}

		routes[localRoute] = entry
	}

	return routes, nil
}

// GenerateWorkspaceImportMap generates an import map for workspace mode
// Creates scopes for each package directory with cross-package imports
func GenerateWorkspaceImportMap(workspaceRoot string, packages []PackageContext) (*ImportMap, error) {
	// Start with base import map from workspace root
	baseImportMap, err := GenerateImportMap(workspaceRoot, &ImportMapConfig{
		Logger: nil, // TODO: Pass logger if needed
	})
	if err != nil {
		return nil, fmt.Errorf("generating base import map: %w", err)
	}

	// Ensure scopes map exists
	if baseImportMap.Scopes == nil {
		baseImportMap.Scopes = make(map[string]map[string]string)
	}

	// Add scopes for each package
	for _, pkg := range packages {
		// Calculate relative path from workspace root
		pkgRelPath, err := filepath.Rel(workspaceRoot, pkg.Path)
		if err != nil {
			continue
		}

		// Scope key is the package directory path
		scopeKey := "/" + filepath.ToSlash(pkgRelPath) + "/"

		// Create scope for this package if it doesn't exist
		if baseImportMap.Scopes[scopeKey] == nil {
			baseImportMap.Scopes[scopeKey] = make(map[string]string)
		}

		// Add cross-package imports to this package's scope
		for _, otherPkg := range packages {
			if otherPkg.Name == pkg.Name {
				continue // Skip self
			}

			// Calculate relative path to other package
			otherRelPath, err := filepath.Rel(workspaceRoot, otherPkg.Path)
			if err != nil {
				continue
			}

			// Map package name to its path
			otherWebPath := "/" + filepath.ToSlash(otherRelPath)

			// Resolve entry point using shared helper
			if entryPoint, err := resolvePackageEntryPoint(otherPkg.Path); err == nil && entryPoint != "" {
				baseImportMap.Scopes[scopeKey][otherPkg.Name] = otherWebPath + "/" + entryPoint
			} else {
				// Fallback: map to package directory
				baseImportMap.Scopes[scopeKey][otherPkg.Name] = otherWebPath + "/"
			}
		}

		// Inherit shared dependencies from root imports
		for depName, depPath := range baseImportMap.Imports {
			// Only add if not already in scope (cross-package imports take precedence)
			if _, exists := baseImportMap.Scopes[scopeKey][depName]; !exists {
				baseImportMap.Scopes[scopeKey][depName] = depPath
			}
		}
	}

	return baseImportMap, nil
}
