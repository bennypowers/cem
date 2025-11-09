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
	"os"
	"path/filepath"
	"strings"
)

// ImportMap represents an ES module import map
type ImportMap struct {
	Imports map[string]string `json:"imports"`
}

// ImportMapConfig configures import map generation
type ImportMapConfig struct {
	InputMapPath string            // Path to user override file
	CLIOverrides map[string]string // CLI flag overrides (highest priority)
	Logger       Logger            // Logger for warnings
}

// packageJSON represents the structure we need from package.json
type packageJSON struct {
	Name         string                 `json:"name"`
	Dependencies map[string]string      `json:"dependencies"`
	Workspaces   interface{}            `json:"workspaces"` // Can be []string or object with "packages" field
	Exports      interface{}            `json:"exports"`
	Main         string                 `json:"main"`
}

// GenerateImportMap generates an import map from package.json and configuration
func GenerateImportMap(rootDir string, config *ImportMapConfig) (*ImportMap, error) {
	result := &ImportMap{
		Imports: make(map[string]string),
	}

	// Default logger to no-op if not provided
	logger := config
	if logger == nil {
		logger = &ImportMapConfig{}
	}

	// 1. Parse root package.json
	rootPkgPath := filepath.Join(rootDir, "package.json")
	rootPkg, err := readPackageJSON(rootPkgPath)
	if err != nil {
		// Empty package.json is OK - just return empty import map
		if os.IsNotExist(err) {
			return result, nil
		}
		return nil, fmt.Errorf("reading package.json: %w", err)
	}

	// 2. Discover workspace packages (monorepo support)
	workspacePackages := make(map[string]string) // name -> path
	if rootPkg.Workspaces != nil {
		workspaces, err := discoverWorkspaces(rootDir, rootPkg.Workspaces)
		if err != nil {
			if logger.Logger != nil {
				logger.Logger.Warning("Failed to discover workspaces: %v", err)
			}
		} else {
			workspacePackages = workspaces
		}
	}

	// 2.5. If not a monorepo (no workspaces), process root package's own exports
	if rootPkg.Workspaces == nil && rootPkg.Name != "" {
		if err := addPackageExportsToImportMap(result, rootPkg.Name, rootDir, rootDir, logger.Logger); err != nil {
			if logger.Logger != nil {
				logger.Logger.Warning("Failed to add root package %s exports: %v", rootPkg.Name, err)
			}
		}
	}

	// 3. Resolve workspace packages first (they have priority over node_modules)
	for name, pkgPath := range workspacePackages {
		if err := addPackageExportsToImportMap(result, name, pkgPath, rootDir, logger.Logger); err != nil {
			if logger.Logger != nil {
				logger.Logger.Warning("Failed to add workspace package %s exports: %v", name, err)
			}
		}
	}

	// 4. Resolve dependencies from node_modules (skip if already in workspaces)
	if rootPkg.Dependencies != nil {
		for depName := range rootPkg.Dependencies {
			// Skip if already resolved as workspace package
			if _, isWorkspace := workspacePackages[depName]; isWorkspace {
				continue
			}

			// Check for scoped package
			depPath := filepath.Join(rootDir, "node_modules", depName)

			// Check if package exists
			if _, err := os.Stat(depPath); os.IsNotExist(err) {
				if logger.Logger != nil {
					logger.Logger.Warning("Dependency %s listed in package.json but not found in node_modules", depName)
				}
				continue
			}

			if err := addPackageExportsToImportMap(result, depName, depPath, rootDir, logger.Logger); err != nil {
				if logger.Logger != nil {
					logger.Logger.Warning("Failed to add package %s exports: %v", depName, err)
				}
			}
		}
	}

	// 5. Merge with user override file (if provided)
	if logger.InputMapPath != "" {
		userMap, err := readImportMapFile(logger.InputMapPath)
		if err != nil {
			return nil, fmt.Errorf("reading user override file: %w", err)
		}
		// User overrides win - deep merge
		for key, value := range userMap.Imports {
			result.Imports[key] = value
		}
	}

	// 6. Apply CLI overrides (highest priority)
	if logger.CLIOverrides != nil {
		for key, value := range logger.CLIOverrides {
			result.Imports[key] = value
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

// discoverWorkspaces finds all workspace packages
func discoverWorkspaces(rootDir string, workspacesField interface{}) (map[string]string, error) {
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

	// For each pattern, find matching directories
	for _, pattern := range patterns {
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
				result[pkg.Name] = match
			}
		}
	}

	return result, nil
}

// readImportMapFile reads a user-provided import map file
func readImportMapFile(path string) (*ImportMap, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var importMap ImportMap
	if err := json.Unmarshal(data, &importMap); err != nil {
		return nil, fmt.Errorf("invalid JSON in import map file: %w", err)
	}

	return &importMap, nil
}

// addPackageExportsToImportMap adds all exports from a package to the import map
func addPackageExportsToImportMap(importMap *ImportMap, pkgName, pkgPath, rootDir string, logger Logger) error {
	pkgJSONPath := filepath.Join(pkgPath, "package.json")
	pkg, err := readPackageJSON(pkgJSONPath)
	if err != nil {
		return err
	}

	pkgRelPath, err := filepath.Rel(rootDir, pkgPath)
	if err != nil {
		return err
	}
	// Handle case where pkgPath == rootDir (relative path is ".")
	var pkgWebPath string
	if pkgRelPath == "." {
		pkgWebPath = ""
	} else {
		pkgWebPath = "/" + filepath.ToSlash(pkgRelPath)
	}

	// Process exports field if present
	if pkg.Exports != nil {
		switch exports := pkg.Exports.(type) {
		case string:
			// Simple string export: "exports": "./index.js"
			importMap.Imports[pkgName] = pkgWebPath + "/" + strings.TrimPrefix(exports, "./")

		case map[string]interface{}:
			// Check if this is a condition-only export (no subpaths)
			// e.g., { "import": "./dist/index.mjs", "require": "./dist/index.cjs" }
			hasSubpaths := false
			for exportPath := range exports {
				if strings.HasPrefix(exportPath, ".") {
					hasSubpaths = true
					break
				}
			}
			if !hasSubpaths {
				// This is a condition-only export map
				resolved, err := resolveExportValue(exports, pkgPath, rootDir)
				if err != nil {
					return err
				}
				importMap.Imports[pkgName] = resolved
				break
			}

			// Process each export path
			for exportPath, exportValue := range exports {
				// Get the import map key
				var importKey string
				if exportPath == "." {
					importKey = pkgName
				} else {
					// Remove leading "./" from export path
					cleanPath := strings.TrimPrefix(exportPath, "./")
					importKey = pkgName + "/" + cleanPath
				}

				// Handle wildcard patterns like "./*" or "./utils/*"
				if strings.Contains(exportPath, "*") {
					if targetStr, ok := exportValue.(string); ok && strings.Contains(targetStr, "*") {
						// Extract source and destination prefixes
						// e.g., "./utils/*" -> "./src/utils/*" becomes "utils/" -> "/src/utils/"
						srcPrefix := strings.TrimSuffix(strings.TrimPrefix(exportPath, "./"), "*")
						dstPrefix := strings.TrimSuffix(strings.TrimPrefix(targetStr, "./"), "*")
						importMap.Imports[pkgName+"/"+srcPrefix] = pkgWebPath + "/" + dstPrefix
					}
					continue
				}

				// Resolve the export value for non-wildcard exports
				resolved, err := resolveExportValue(exportValue, pkgPath, rootDir)
				if err != nil {
					// Skip exports we can't resolve
					continue
				}

				importMap.Imports[importKey] = resolved
			}
		}

		// Do NOT add default trailing-slash mapping for packages with exports
		// Only root wildcard exports (./*) should add trailing slash (handled above)
		// Adding a broad trailing slash would incorrectly expose all package internals
		return nil
	}

	// Fallback to main field if no exports
	if pkg.Main != "" {
		mainPath := pkgWebPath + "/" + strings.TrimPrefix(pkg.Main, "./")
		importMap.Imports[pkgName] = mainPath
		importMap.Imports[pkgName+"/"] = pkgWebPath + "/"
		return nil
	}

	return fmt.Errorf("no exports or main field")
}

// resolveExportValue resolves an export value to a web path
func resolveExportValue(exportValue interface{}, pkgPath, rootDir string) (string, error) {
	// Calculate relative path and handle root package case
	pkgRelPath, err := filepath.Rel(rootDir, pkgPath)
	if err != nil {
		return "", err
	}
	var pkgWebPath string
	if pkgRelPath == "." {
		pkgWebPath = ""
	} else {
		pkgWebPath = "/" + filepath.ToSlash(pkgRelPath)
	}

	switch v := exportValue.(type) {
	case string:
		// Direct string path
		return pkgWebPath + "/" + strings.TrimPrefix(v, "./"), nil

	case map[string]interface{}:
		// Conditional exports - prioritize "import" condition
		if importPath, ok := v["import"].(string); ok {
			return pkgWebPath + "/" + strings.TrimPrefix(importPath, "./"), nil
		}
		// Try default condition
		if defaultPath, ok := v["default"].(string); ok {
			return pkgWebPath + "/" + strings.TrimPrefix(defaultPath, "./"), nil
		}
	}

	return "", fmt.Errorf("could not resolve export value")
}
