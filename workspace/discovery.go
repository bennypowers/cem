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

package workspace

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	C "bennypowers.dev/cem/cmd/config"
	"gopkg.in/yaml.v3"
)

// PackageInfo represents a discovered workspace package
type PackageInfo struct {
	Name              string // Package name from package.json
	Path              string // Absolute path to package directory
	CustomElementsRef string // Value of customElements field (path to manifest)
}

// packageJSON represents the structure we need from package.json
type packageJSON struct {
	Name           string      `json:"name"`
	Workspaces     interface{} `json:"workspaces"` // Can be []string or object with "packages" field
	CustomElements string      `json:"customElements"`
}

// DiscoverWorkspacePackages discovers all workspace packages from workspace patterns
// Returns map of package name -> absolute path to package directory
// Supports negated patterns (prefixed with !) to exclude packages
func DiscoverWorkspacePackages(rootDir string, workspacesField interface{}) (map[string]string, error) {
	result := make(map[string]string)

	var patterns []string

	// Handle different workspace field formats
	switch v := workspacesField.(type) {
	case []interface{}:
		for _, item := range v {
			if str, ok := item.(string); ok {
				patterns = append(patterns, str)
			}
		}
	case map[string]interface{}:
		if packages, ok := v["packages"].([]interface{}); ok {
			for _, item := range packages {
				if str, ok := item.(string); ok {
					patterns = append(patterns, str)
				}
			}
		}
	}

	// Split patterns into includes and excludes
	var includePatterns []string
	var excludePatterns []string

	for _, pattern := range patterns {
		if len(pattern) > 0 && pattern[0] == '!' {
			// Negated pattern - strip the ! and add to excludes
			excludePatterns = append(excludePatterns, pattern[1:])
		} else {
			includePatterns = append(includePatterns, pattern)
		}
	}

	// Process include patterns to find all matching directories
	for _, pattern := range includePatterns {
		matches, err := filepath.Glob(filepath.Join(rootDir, pattern))
		if err != nil {
			continue
		}

		for _, match := range matches {
			info, err := os.Stat(match)
			if err != nil || !info.IsDir() {
				continue
			}

			// Read package.json in this workspace
			pkgPath := filepath.Join(match, "package.json")
			pkg, err := readPackageJSON(pkgPath)
			if err != nil {
				continue
			}

			if pkg.Name != "" {
				absPath, err := filepath.Abs(match)
				if err != nil {
					absPath = match
				}
				result[pkg.Name] = absPath
			}
		}
	}

	// Remove packages that match exclude patterns
	for _, excludePattern := range excludePatterns {
		matches, err := filepath.Glob(filepath.Join(rootDir, excludePattern))
		if err != nil {
			continue
		}

		for _, match := range matches {
			// Read package.json to get the package name
			pkgPath := filepath.Join(match, "package.json")
			pkg, err := readPackageJSON(pkgPath)
			if err != nil {
				continue
			}

			// Remove from result if it was found
			delete(result, pkg.Name)
		}
	}

	return result, nil
}

// FindPackagesWithManifests finds all workspace packages that have a customElements field
// This is what the serve command uses to auto-discover packages to serve
func FindPackagesWithManifests(rootDir string) ([]PackageInfo, error) {
	// Read root package.json to get workspaces
	rootPkgPath := filepath.Join(rootDir, "package.json")
	rootPkg, err := readPackageJSON(rootPkgPath)
	if err != nil {
		return nil, err
	}

	// If no workspaces, not a monorepo
	if rootPkg.Workspaces == nil {
		return nil, nil
	}

	// Discover all workspace packages
	packages, err := DiscoverWorkspacePackages(rootDir, rootPkg.Workspaces)
	if err != nil {
		return nil, err
	}

	// Filter to only packages with customElements field
	var result []PackageInfo
	for name, path := range packages {
		pkgPath := filepath.Join(path, "package.json")
		pkg, err := readPackageJSON(pkgPath)
		if err != nil {
			continue
		}

		if pkg.CustomElements != "" {
			result = append(result, PackageInfo{
				Name:              name,
				Path:              path,
				CustomElementsRef: pkg.CustomElements,
			})
		}
	}

	return result, nil
}

// readPackageJSON reads and parses a package.json file
func readPackageJSON(path string) (*packageJSON, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var pkg packageJSON
	if err := json.Unmarshal(data, &pkg); err != nil {
		return nil, err
	}

	return &pkg, nil
}

// LoadWorkspaceConfig loads config from workspace root
// Returns nil if no workspace root found or no config file exists
func LoadWorkspaceConfig(packageDir string) (*C.CemConfig, error) {
	// Find workspace root by looking for package.json with workspaces field
	workspaceRoot := findWorkspaceRootWithWorkspaces(packageDir)
	if workspaceRoot == "" {
		return nil, nil // Not in a workspace
	}

	configPath := filepath.Join(workspaceRoot, ".config", "cem.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, nil // No workspace config
	}

	// Read and parse config
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config C.CemConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	config.ProjectDir = workspaceRoot
	config.ConfigFile = configPath

	return &config, nil
}

// LoadPackageConfigWithWorkspaceDefaults loads package config and merges with workspace defaults
// Package settings override workspace settings
func LoadPackageConfigWithWorkspaceDefaults(packageDir string) (*C.CemConfig, error) {
	// Load workspace config (if exists)
	workspaceConfig, err := LoadWorkspaceConfig(packageDir)
	if err != nil {
		return nil, err
	}

	// Load package config
	packageConfigPath := filepath.Join(packageDir, ".config", "cem.yaml")
	packageConfig := &C.CemConfig{
		ProjectDir: packageDir,
		ConfigFile: packageConfigPath,
	}

	if _, err := os.Stat(packageConfigPath); err == nil {
		data, err := os.ReadFile(packageConfigPath)
		if err != nil {
			return nil, err
		}

		if err := yaml.Unmarshal(data, packageConfig); err != nil {
			return nil, err
		}
	}

	// If no workspace config, just return package config
	if workspaceConfig == nil {
		return packageConfig, nil
	}

	// Merge: package config overrides workspace defaults
	// For serve config: use workspace defaults but allow package to override
	if packageConfig.Serve.Port == 0 && workspaceConfig.Serve.Port != 0 {
		packageConfig.Serve.Port = workspaceConfig.Serve.Port
	}
	if !packageConfig.Serve.OpenBrowser && workspaceConfig.Serve.OpenBrowser {
		packageConfig.Serve.OpenBrowser = workspaceConfig.Serve.OpenBrowser
	}

	return packageConfig, nil
}

// findWorkspaceRootWithWorkspaces finds workspace root by looking for package.json with workspaces field
// Returns empty string if not in a workspace
func findWorkspaceRootWithWorkspaces(startDir string) string {
	dir := startDir
	for {
		// Check if there's a package.json with workspaces field
		pkgPath := filepath.Join(dir, "package.json")
		if pkg, err := readPackageJSON(pkgPath); err == nil && pkg.Workspaces != nil {
			return dir
		}

		// Stop if we've reached a git repository root (don't go higher)
		gitDir := filepath.Join(dir, ".git")
		if stat, err := os.Stat(gitDir); err == nil && stat.IsDir() {
			return "" // Hit git boundary without finding workspace
		}

		// Move up one directory
		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached filesystem root
			return ""
		}
		dir = parent
	}
}

// FindPackagesForFiles determines which workspace packages are affected by a set of changed files.
// Returns a deduplicated list of PackageInfo for all affected packages.
func FindPackagesForFiles(rootDir string, filePaths []string) ([]PackageInfo, error) {
	// Get all packages with manifests
	packages, err := FindPackagesWithManifests(rootDir)
	if err != nil {
		return nil, err
	}

	// Build a set of affected package paths to deduplicate
	affectedSet := make(map[string]PackageInfo)

	for _, filePath := range filePaths {
		// Make file path absolute for comparison
		absFilePath, err := filepath.Abs(filePath)
		if err != nil {
			absFilePath = filePath
		}

		for _, pkg := range packages {
			// Check if the file is under this package's directory
			if strings.HasPrefix(absFilePath, pkg.Path+string(filepath.Separator)) {
				affectedSet[pkg.Path] = pkg
			}
		}
	}

	// Convert map to slice
	result := make([]PackageInfo, 0, len(affectedSet))
	for _, pkg := range affectedSet {
		result = append(result, pkg)
	}

	return result, nil
}

// IsWorkspaceMode determines if a directory is a monorepo workspace
// Returns true if the directory has a workspaces field in package.json
func IsWorkspaceMode(dir string) bool {
	// Read root package.json to check for workspaces field
	rootPkgPath := filepath.Join(dir, "package.json")
	rootPkg, err := readPackageJSON(rootPkgPath)
	if err != nil {
		return false
	}

	// Workspace mode is determined by presence of workspaces field
	return rootPkg.Workspaces != nil
}

// PackageWithManifest represents a workspace package with its loaded manifest
type PackageWithManifest struct {
	Name     string // Package name from package.json
	Path     string // Absolute path to package directory
	Manifest []byte // Loaded custom elements manifest JSON
}

// LoadWorkspaceManifests loads manifest data for all workspace packages
// Returns packages with their manifests loaded, skipping any that fail to load
func LoadWorkspaceManifests(rootDir string) ([]PackageWithManifest, error) {
	packages, err := FindPackagesWithManifests(rootDir)
	if err != nil {
		return nil, err
	}

	var result []PackageWithManifest
	for _, pkg := range packages {
		// Load package's manifest
		manifestPath := filepath.Join(pkg.Path, pkg.CustomElementsRef)
		manifestData, err := os.ReadFile(manifestPath)
		if err != nil {
			// Skip packages with missing manifests
			continue
		}

		result = append(result, PackageWithManifest{
			Name:     pkg.Name,
			Path:     pkg.Path,
			Manifest: manifestData,
		})
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no packages with manifests found in workspace")
	}

	return result, nil
}
