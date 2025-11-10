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

	W "bennypowers.dev/cem/workspace"
)

// ImportMap represents an ES module import map
type ImportMap struct {
	Imports map[string]string            `json:"imports"`
	Scopes  map[string]map[string]string `json:"scopes,omitempty"`
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

// findWorkspaceRoot walks up the directory tree to find the workspace root
// Returns the directory containing node_modules or workspace configuration
// Stops at git repository boundaries to avoid breaking out of submodules
func findWorkspaceRoot(startDir string) string {
	dir := startDir
	for {
		// Check if node_modules exists in this directory
		nodeModulesPath := filepath.Join(dir, "node_modules")
		if stat, err := os.Stat(nodeModulesPath); err == nil && stat.IsDir() {
			return dir
		}

		// Check if there's a package.json with workspaces field
		pkgPath := filepath.Join(dir, "package.json")
		if pkg, err := readPackageJSON(pkgPath); err == nil && pkg.Workspaces != nil {
			return dir
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

// GenerateImportMap generates an import map from package.json and configuration
func GenerateImportMap(rootDir string, config *ImportMapConfig) (*ImportMap, error) {
	result := &ImportMap{
		Imports: make(map[string]string),
		Scopes:  make(map[string]map[string]string),
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

	// 1.5. Find workspace root if we're in a workspace package
	// This handles the case where serve is run from a subdirectory in a monorepo
	workspaceRoot := findWorkspaceRoot(rootDir)
	if workspaceRoot != rootDir {
		if logger.Logger != nil {
			logger.Logger.Debug("Detected workspace subdirectory, using workspace root: %s", workspaceRoot)
		}
	}

	// 2. Discover workspace packages (monorepo support)
	// If we found a workspace root, load its package.json for workspace config
	workspacePackages := make(map[string]string) // name -> path
	workspacePkg := rootPkg
	if workspaceRoot != rootDir {
		// Load workspace root's package.json
		workspaceRootPkgPath := filepath.Join(workspaceRoot, "package.json")
		if pkg, err := readPackageJSON(workspaceRootPkgPath); err == nil {
			workspacePkg = pkg
		}
	}
	if workspacePkg.Workspaces != nil {
		workspaces, err := W.DiscoverWorkspacePackages(workspaceRoot, workspacePkg.Workspaces)
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
	// Use workspaceRoot for node_modules location (handles monorepo case)
	if rootPkg.Dependencies != nil {
		for depName := range rootPkg.Dependencies {
			// Skip if already resolved as workspace package
			if _, isWorkspace := workspacePackages[depName]; isWorkspace {
				continue
			}

			// Look for package in workspace root's node_modules
			depPath := filepath.Join(workspaceRoot, "node_modules", depName)

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

	// 4.5. Add scopes for transitive dependencies in node_modules
	// For each package in the dependency tree, add scopes for their dependencies
	// Use workspaceRoot for node_modules location (handles monorepo case)
	if err := addTransitiveDependenciesToScopes(result, workspaceRoot, rootPkg, logger.Logger); err != nil {
		if logger.Logger != nil {
			logger.Logger.Warning("Failed to add transitive dependencies: %v", err)
		}
	}

	// 4.6. Add scopes for workspace packages and root package
	// Workspace packages and root package also need to resolve bare specifiers
	// Use workspaceRoot for node_modules location (handles monorepo case)
	if err := addWorkspaceScopesToImportMap(result, workspaceRoot, rootDir, rootPkg, workspacePackages, logger.Logger); err != nil {
		if logger.Logger != nil {
			logger.Logger.Warning("Failed to add workspace scopes: %v", err)
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

// resolvePackageEntryPoint resolves a package's main entry point from package.json
// Returns the relative path to the entry point (e.g., "index.js" or "dist/index.mjs")
// Returns empty string if no entry point can be determined
func resolvePackageEntryPoint(pkgPath string) (string, error) {
	pkgJSONPath := filepath.Join(pkgPath, "package.json")
	data, err := os.ReadFile(pkgJSONPath)
	if err != nil {
		return "", err
	}

	var pkg struct {
		Exports interface{} `json:"exports"`
		Main    string      `json:"main"`
	}
	if err := json.Unmarshal(data, &pkg); err != nil {
		return "", err
	}

	// Try exports field first
	if pkg.Exports != nil {
		// Try to resolve the "." export
		if exportsMap, ok := pkg.Exports.(map[string]interface{}); ok {
			if rootExport, ok := exportsMap["."].(string); ok {
				return strings.TrimPrefix(rootExport, "./"), nil
			}
			// Try nested conditions
			if rootValue, ok := exportsMap["."]; ok {
				if resolved := resolveSimpleExportValue(rootValue, ""); resolved != "" {
					return strings.TrimPrefix(resolved, "/"), nil
				}
			}
		}
		// Try string exports
		if exportStr, ok := pkg.Exports.(string); ok {
			return strings.TrimPrefix(exportStr, "./"), nil
		}
	}

	// Fallback to main field
	if pkg.Main != "" {
		return strings.TrimPrefix(pkg.Main, "./"), nil
	}

	return "", nil
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
		if importValue, ok := v["import"]; ok {
			// Recursively resolve if it's a nested condition
			if importPath, ok := importValue.(string); ok {
				return pkgWebPath + "/" + strings.TrimPrefix(importPath, "./"), nil
			}
			// Recursively resolve nested conditions
			if importMap, ok := importValue.(map[string]interface{}); ok {
				return resolveExportValue(importMap, pkgPath, rootDir)
			}
		}
		// Try default condition
		if defaultValue, ok := v["default"]; ok {
			if defaultPath, ok := defaultValue.(string); ok {
				return pkgWebPath + "/" + strings.TrimPrefix(defaultPath, "./"), nil
			}
			// Recursively resolve nested default
			if defaultMap, ok := defaultValue.(map[string]interface{}); ok {
				return resolveExportValue(defaultMap, pkgPath, rootDir)
			}
		}
	}

	return "", fmt.Errorf("could not resolve export value")
}

// addTransitiveDependenciesToScopes builds scopes for packages in the dependency tree
// by recursively traversing dependencies from the root package
func addTransitiveDependenciesToScopes(importMap *ImportMap, rootDir string, rootPkg *packageJSON, logger Logger) error {
	if rootPkg == nil {
		return nil
	}

	// Build dependency tree starting from root dependencies
	visited := make(map[string]bool)
	nodeModulesPath := filepath.Join(rootDir, "node_modules")

	// Process each root dependency recursively
	if rootPkg.Dependencies != nil {
		for depName := range rootPkg.Dependencies {
			if err := buildScopesForPackage(importMap, depName, nodeModulesPath, visited); err != nil {
				if logger != nil {
					logger.Warning("Failed to build scopes for %s: %v", depName, err)
				}
			}
		}
	}

	return nil
}

// buildScopesForPackage recursively builds scopes for a package and its dependencies
func buildScopesForPackage(importMap *ImportMap, pkgName string, nodeModulesPath string, visited map[string]bool) error {
	// Skip if already processed
	if visited[pkgName] {
		return nil
	}
	visited[pkgName] = true

	// Read package.json
	pkgPath := filepath.Join(nodeModulesPath, pkgName)
	pkgJSONPath := filepath.Join(pkgPath, "package.json")
	pkg, err := readPackageJSON(pkgJSONPath)
	if err != nil {
		return err
	}

	// If this package has no dependencies, nothing to scope
	if pkg.Dependencies == nil || len(pkg.Dependencies) == 0 {
		return nil
	}

	// Create scope key for this package
	scopeKey := "/node_modules/" + pkgName + "/"
	if importMap.Scopes[scopeKey] == nil {
		importMap.Scopes[scopeKey] = make(map[string]string)
	}

	// Add each dependency to this package's scope
	for depName := range pkg.Dependencies {
		// Check if dependency exists
		depPath := filepath.Join(nodeModulesPath, depName)
		if _, err := os.Stat(depPath); os.IsNotExist(err) {
			continue // Dependency not installed, skip
		}

		// Read dependency's package.json
		depPkgPath := filepath.Join(depPath, "package.json")
		depPkg, err := readPackageJSON(depPkgPath)
		if err != nil {
			continue
		}

		// Add dependency to this package's scope
		depWebPath := "/node_modules/" + depName
		if depPkg.Exports != nil {
			addDependencyExportsToScope(importMap.Scopes[scopeKey], depName, depWebPath, depPkg.Exports)
		} else if depPkg.Main != "" {
			mainPath := depWebPath + "/" + strings.TrimPrefix(depPkg.Main, "./")
			importMap.Scopes[scopeKey][depName] = mainPath
			importMap.Scopes[scopeKey][depName+"/"] = depWebPath + "/"
		} else {
			importMap.Scopes[scopeKey][depName+"/"] = depWebPath + "/"
		}

		// Recursively process this dependency's dependencies
		if err := buildScopesForPackage(importMap, depName, nodeModulesPath, visited); err != nil {
			// Continue on error, just skip this dependency's tree
			continue
		}
	}

	return nil
}

// addDependencyExportsToScope adds dependency exports to a scope map
func addDependencyExportsToScope(scope map[string]string, depName, depWebPath string, exports interface{}) {
	switch exp := exports.(type) {
	case string:
		// Simple string export
		scope[depName] = depWebPath + "/" + strings.TrimPrefix(exp, "./")

	case map[string]interface{}:
		// Check if this is condition-only (no subpaths)
		hasSubpaths := false
		for exportPath := range exp {
			if strings.HasPrefix(exportPath, ".") {
				hasSubpaths = true
				break
			}
		}

		if !hasSubpaths {
			// Condition-only export, resolve it
			if importPath, ok := exp["import"].(string); ok {
				scope[depName] = depWebPath + "/" + strings.TrimPrefix(importPath, "./")
			} else if defaultPath, ok := exp["default"].(string); ok {
				scope[depName] = depWebPath + "/" + strings.TrimPrefix(defaultPath, "./")
			}
		} else {
			// Has subpaths - process each export path
			for exportPath, exportValue := range exp {
				var importKey string
				if exportPath == "." {
					importKey = depName
				} else {
					cleanPath := strings.TrimPrefix(exportPath, "./")
					importKey = depName + "/" + cleanPath
				}

				// Handle wildcard patterns like "./decorators/*" -> "./decorators/*.js"
				if strings.Contains(exportPath, "*") {
					if targetStr, ok := exportValue.(string); ok && strings.Contains(targetStr, "*") {
						// Extract prefixes before the wildcard
						// "./decorators/*" becomes "decorators/"
						// "./decorators/*.js" becomes "decorators/"
						srcParts := strings.Split(strings.TrimPrefix(exportPath, "./"), "*")
						dstParts := strings.Split(strings.TrimPrefix(targetStr, "./"), "*")
						srcPrefix := srcParts[0]
						dstPrefix := dstParts[0]
						// Add trailing slash mapping for wildcards
						scope[depName+"/"+srcPrefix] = depWebPath + "/" + dstPrefix
					}
					continue
				}

				// Resolve export value
				if resolvedPath := resolveSimpleExportValue(exportValue, depWebPath); resolvedPath != "" {
					scope[importKey] = resolvedPath
				}
			}
		}
	}
}

// resolveSimpleExportValue resolves an export value to a path (simplified version for scopes)
func resolveSimpleExportValue(exportValue interface{}, depWebPath string) string {
	switch v := exportValue.(type) {
	case string:
		return depWebPath + "/" + strings.TrimPrefix(v, "./")
	case map[string]interface{}:
		// Try import condition first
		if importValue, ok := v["import"]; ok {
			if importPath, ok := importValue.(string); ok {
				return depWebPath + "/" + strings.TrimPrefix(importPath, "./")
			}
			// Recursively resolve nested import conditions
			if importMap, ok := importValue.(map[string]interface{}); ok {
				return resolveSimpleExportValue(importMap, depWebPath)
			}
		}
		// Try default condition
		if defaultValue, ok := v["default"]; ok {
			if defaultPath, ok := defaultValue.(string); ok {
				return depWebPath + "/" + strings.TrimPrefix(defaultPath, "./")
			}
			// Recursively resolve nested default
			if defaultMap, ok := defaultValue.(map[string]interface{}); ok {
				return resolveSimpleExportValue(defaultMap, depWebPath)
			}
		}
	}
	return ""
}

// addWorkspaceScopesToImportMap adds scopes for workspace packages and root package
// so they can resolve bare specifiers to dependencies
// workspaceRoot is where node_modules is located, rootDir is for calculating relative paths
func addWorkspaceScopesToImportMap(importMap *ImportMap, workspaceRoot, rootDir string, rootPkg *packageJSON, workspacePackages map[string]string, logger Logger) error {
	nodeModulesPath := filepath.Join(workspaceRoot, "node_modules")

	// Helper function to add dependencies to a scope (including transitive dependencies)
	addDependenciesToScope := func(scopePrefix string, dependencies map[string]string) {
		if len(dependencies) == 0 {
			return
		}

		if importMap.Scopes[scopePrefix] == nil {
			importMap.Scopes[scopePrefix] = make(map[string]string)
		}

		// Track visited packages to avoid infinite loops
		visited := make(map[string]bool)

		// Recursively add a package and its transitive dependencies to the scope
		var addPackageAndDeps func(string)
		addPackageAndDeps = func(depName string) {
			if visited[depName] {
				return
			}
			visited[depName] = true

			depPath := filepath.Join(nodeModulesPath, depName)
			if _, err := os.Stat(depPath); os.IsNotExist(err) {
				return // Dependency not installed
			}

			depPkgPath := filepath.Join(depPath, "package.json")
			depPkg, err := readPackageJSON(depPkgPath)
			if err != nil {
				return
			}

			// Add this package to the scope
			depWebPath := "/node_modules/" + depName
			if depPkg.Exports != nil {
				addDependencyExportsToScope(importMap.Scopes[scopePrefix], depName, depWebPath, depPkg.Exports)
			} else if depPkg.Main != "" {
				mainPath := depWebPath + "/" + strings.TrimPrefix(depPkg.Main, "./")
				importMap.Scopes[scopePrefix][depName] = mainPath
				importMap.Scopes[scopePrefix][depName+"/"] = depWebPath + "/"
			} else {
				importMap.Scopes[scopePrefix][depName+"/"] = depWebPath + "/"
			}

			// Recursively add its dependencies
			if depPkg.Dependencies != nil {
				for transDepName := range depPkg.Dependencies {
					addPackageAndDeps(transDepName)
				}
			}
		}

		// Add each direct dependency (which will recursively add transitives)
		for depName := range dependencies {
			addPackageAndDeps(depName)
		}
	}

	// Add scopes for root package (if not a monorepo, or if it has dependencies)
	if rootPkg != nil && rootPkg.Dependencies != nil {
		// Root package files are served from "/"
		addDependenciesToScope("/", rootPkg.Dependencies)
	}

	// Add scopes for each workspace package
	for _, pkgPath := range workspacePackages {
		pkgJSONPath := filepath.Join(pkgPath, "package.json")
		pkg, err := readPackageJSON(pkgJSONPath)
		if err != nil || pkg.Dependencies == nil {
			continue
		}

		// Calculate web path for this workspace package
		pkgRelPath, err := filepath.Rel(rootDir, pkgPath)
		if err != nil {
			continue
		}

		var scopePrefix string
		if pkgRelPath == "." {
			scopePrefix = "/"
		} else {
			scopePrefix = "/" + filepath.ToSlash(pkgRelPath) + "/"
		}

		addDependenciesToScope(scopePrefix, pkg.Dependencies)
	}

	return nil
}
