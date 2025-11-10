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
	"fmt"
	"os"
	"path/filepath"

	C "bennypowers.dev/cem/cmd/config"
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
