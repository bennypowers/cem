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

package importmap

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/serve/middleware"
	"bennypowers.dev/cem/serve/middleware/types"
	W "bennypowers.dev/cem/workspace"
)

// packageJSON represents the structure we need from package.json
type packageJSON struct {
	Name         string            `json:"name"`
	Dependencies map[string]string `json:"dependencies"`
	Workspaces   any               `json:"workspaces"` // Can be []string or object with "packages" field
	Exports      any               `json:"exports"`
	Main         string            `json:"main"`
}

// findWorkspaceRoot walks up the directory tree to find the workspace root
// Returns the directory containing node_modules or workspace configuration
// Stops at git repository boundaries to avoid breaking out of submodules
func findWorkspaceRoot(startDir string, fs platform.FileSystem) string {
	dir := startDir
	for {
		// Check if node_modules exists in this directory
		nodeModulesPath := filepath.Join(dir, "node_modules")
		if stat, err := fs.Stat(nodeModulesPath); err == nil && stat.IsDir() {
			return dir
		}

		// Check if there's a package.json with workspaces field
		pkgPath := filepath.Join(dir, "package.json")
		if pkg, err := readPackageJSON(pkgPath, fs); err == nil && pkg.Workspaces != nil {
			return dir
		}

		// Stop if we've reached a git repository root (don't go higher)
		gitDir := filepath.Join(dir, ".git")
		if stat, err := fs.Stat(gitDir); err == nil && stat.IsDir() {
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

// Generate generates an import map from package.json and configuration
func Generate(rootDir string, config *Config) (*ImportMap, error) {
	// Default config to empty if not provided
	if config == nil {
		config = &Config{}
	}

	// Default to OS filesystem if not provided
	fs := config.FS
	if fs == nil {
		fs = platform.NewOSFileSystem()
	}

	// Normalize rootDir for workspace discovery and relative path calculations
	if absRoot, err := filepath.Abs(rootDir); err == nil {
		rootDir = absRoot
	}

	// Workspace mode: generate flattened import map from packages
	if len(config.WorkspacePackages) > 0 {
		return generateWorkspaceImportMap(rootDir, config.WorkspacePackages, config.Logger, fs)
	}

	// Single-package mode: generate import map with scopes
	result := &ImportMap{
		Imports: make(map[string]string),
		Scopes:  make(map[string]map[string]string),
	}

	// Alias config for easier access (logger may be nil, checked before use)
	cfg := config

	// 1. Parse root package.json
	rootPkgPath := filepath.Join(rootDir, "package.json")
	rootPkg, err := readPackageJSON(rootPkgPath, fs)
	if err != nil {
		// Empty package.json is OK - apply user/CLI overrides and return
		if errors.Is(err, os.ErrNotExist) {
			if err := applyOverrides(result, cfg, fs); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("reading package.json: %w", err)
	}

	// 1.5. Find workspace root if we're in a workspace package
	// This handles the case where serve is run from a subdirectory in a monorepo
	workspaceRoot := findWorkspaceRoot(rootDir, fs)
	if workspaceRoot != rootDir {
		if cfg.Logger != nil {
			cfg.Logger.Debug("Detected workspace subdirectory, using workspace root: %s", workspaceRoot)
		}
	}

	// 2. Discover workspace packages (monorepo support)
	// If we found a workspace root, load its package.json for workspace config
	workspacePackages := make(map[string]string) // name -> path
	workspacePkg := rootPkg
	if workspaceRoot != rootDir {
		// Load workspace root's package.json
		workspaceRootPkgPath := filepath.Join(workspaceRoot, "package.json")
		if pkg, err := readPackageJSON(workspaceRootPkgPath, fs); err == nil {
			workspacePkg = pkg
		}
	}
	if workspacePkg.Workspaces != nil {
		workspaces, err := W.DiscoverWorkspacePackages(workspaceRoot, workspacePkg.Workspaces)
		if err != nil {
			if cfg.Logger != nil {
				cfg.Logger.Warning("Failed to discover workspaces: %v", err)
			}
		} else {
			workspacePackages = workspaces
		}
	}

	// 2.5. If not a monorepo (no workspaces), process root package's own exports
	// Only add if the package has exports or main field (skip packages that are just dependency containers)
	if rootPkg.Workspaces == nil && rootPkg.Name != "" && (rootPkg.Exports != nil || rootPkg.Main != "") {
		if err := addPackageExportsToImportMap(result, rootPkg.Name, rootDir, rootDir, cfg.Logger, fs); err != nil {
			if cfg.Logger != nil {
				cfg.Logger.Warning("Failed to add root package %s exports: %v", rootPkg.Name, err)
			}
		}
	}

	// 3. Resolve workspace packages first (they have priority over node_modules)
	for name, pkgPath := range workspacePackages {
		if err := addPackageExportsToImportMap(result, name, pkgPath, rootDir, cfg.Logger, fs); err != nil {
			if cfg.Logger != nil {
				cfg.Logger.Warning("Failed to add workspace package %s exports: %v", name, err)
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
			if !fs.Exists(depPath) {
				if cfg.Logger != nil {
					cfg.Logger.Warning("Dependency %s listed in package.json but not found in node_modules", depName)
				}
				continue
			}

			if err := addPackageExportsToImportMap(result, depName, depPath, rootDir, cfg.Logger, fs); err != nil {
				if cfg.Logger != nil {
					cfg.Logger.Warning("Failed to add package %s exports: %v", depName, err)
				}
			}
		}
	}

	// 4.5. Add scopes for transitive dependencies in node_modules
	// For each package in the dependency tree, add scopes for their dependencies
	// Use workspaceRoot for node_modules location (handles monorepo case)
	if err := addTransitiveDependenciesToScopes(result, workspaceRoot, rootPkg, cfg.Logger, fs); err != nil {
		if cfg.Logger != nil {
			cfg.Logger.Warning("Failed to add transitive dependencies: %v", err)
		}
	}

	// 4.6. Add scopes for workspace packages and root package
	// Workspace packages and root package also need to resolve bare specifiers
	// Use workspaceRoot for node_modules location (handles monorepo case)
	if err := addWorkspaceScopesToImportMap(result, workspaceRoot, rootDir, rootPkg, workspacePackages, cfg.Logger, fs); err != nil {
		if cfg.Logger != nil {
			cfg.Logger.Warning("Failed to add workspace scopes: %v", err)
		}
	}

	// 5. Apply user overrides and CLI overrides
	if err := applyOverrides(result, cfg, fs); err != nil {
		return nil, err
	}

	return result, nil
}

// isValidImportKey returns true if the key is valid for import maps.
// Keys must not contain control characters that could break import map parsing.
func isValidImportKey(key string) bool {
	return !strings.ContainsAny(key, "\x00\n\r")
}

// mergeImportsIntoResult merges imports from source into result, validating keys.
// Invalid keys are skipped and logged if logger is provided.
func mergeImportsIntoResult(result *ImportMap, source map[string]string, sourceName string, logger types.Logger) {
	for key, value := range source {
		if !isValidImportKey(key) {
			if logger != nil {
				logger.Warning("Skipping invalid import map key from %s (contains control characters): %q", sourceName, key)
			}
			continue
		}
		result.Imports[key] = value
	}
}

// mergeScopesIntoResult merges scopes from source into result, validating all keys.
// Invalid scope keys or import keys within scopes are skipped and logged if logger is provided.
func mergeScopesIntoResult(result *ImportMap, source map[string]map[string]string, sourceName string, logger types.Logger) {
	if len(source) == 0 {
		return
	}

	if result.Scopes == nil {
		result.Scopes = make(map[string]map[string]string)
	}

	for scopeKey, sourceScopeMap := range source {
		// Validate scope key
		if !isValidImportKey(scopeKey) {
			if logger != nil {
				logger.Warning("Skipping invalid scope key from %s (contains control characters): %q", sourceName, scopeKey)
			}
			continue
		}

		if result.Scopes[scopeKey] == nil {
			// No existing scope for this key, validate and copy the source scope map
			validatedScopeMap := make(map[string]string)
			for importKey, importValue := range sourceScopeMap {
				if !isValidImportKey(importKey) {
					if logger != nil {
						logger.Warning("Skipping invalid import key in scope %q from %s: %q", scopeKey, sourceName, importKey)
					}
					continue
				}
				validatedScopeMap[importKey] = importValue
			}
			result.Scopes[scopeKey] = validatedScopeMap
		} else {
			// Merge individual import entries, source entries override existing
			for importKey, importValue := range sourceScopeMap {
				if !isValidImportKey(importKey) {
					if logger != nil {
						logger.Warning("Skipping invalid import key in scope %q from %s: %q", scopeKey, sourceName, importKey)
					}
					continue
				}
				result.Scopes[scopeKey][importKey] = importValue
			}
		}
	}
}

// applyOverrides merges user override file and config overrides into the result
// This is called both in the normal flow and when package.json is missing
// Priority (highest wins): Config override > Override file > Auto-generated
func applyOverrides(result *ImportMap, cfg *Config, fs platform.FileSystem) error {
	// Merge with user override file (if provided)
	if cfg.InputMapPath != "" {
		// Validate override file path to prevent common mistakes and improve error messages
		// Note: Path traversal is allowed here since this is a config file path specified by
		// the server owner/developer, not user input from web clients.
		cleanPath := filepath.Clean(cfg.InputMapPath)

		// Check if file exists before attempting to read
		if _, err := fs.Stat(cleanPath); err != nil {
			if os.IsNotExist(err) {
				return fmt.Errorf("import map override file not found: %s", cleanPath)
			}
			return fmt.Errorf("cannot access import map override file %s: %w", cleanPath, err)
		}

		userMap, err := readImportMapFile(cleanPath, fs)
		if err != nil {
			return fmt.Errorf("reading user override file %s: %w", cleanPath, err)
		}
		// User overrides win - deep merge
		mergeImportsIntoResult(result, userMap.Imports, "override file", cfg.Logger)
		mergeScopesIntoResult(result, userMap.Scopes, "override file", cfg.Logger)
	}

	// Apply config overrides (highest priority)
	if cfg.ConfigOverride != nil {
		mergeImportsIntoResult(result, cfg.ConfigOverride.Imports, "config override", cfg.Logger)
		mergeScopesIntoResult(result, cfg.ConfigOverride.Scopes, "config override", cfg.Logger)
	}

	return nil
}

// readPackageJSON reads and parses a package.json file
func readPackageJSON(path string, fs platform.FileSystem) (*packageJSON, error) {
	data, err := fs.ReadFile(path)
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
func resolvePackageEntryPoint(pkgPath string, fs platform.FileSystem) (string, error) {
	pkgJSONPath := filepath.Join(pkgPath, "package.json")
	data, err := fs.ReadFile(pkgJSONPath)
	if err != nil {
		return "", err
	}

	var pkg struct {
		Exports any    `json:"exports"`
		Main    string `json:"main"`
	}
	if err := json.Unmarshal(data, &pkg); err != nil {
		return "", err
	}

	// Try exports field first
	if pkg.Exports != nil {
		// Try to resolve the "." export
		if exportsMap, ok := pkg.Exports.(map[string]any); ok {
			if rootExport, ok := exportsMap["."].(string); ok {
				return strings.TrimPrefix(rootExport, "./"), nil
			}
			// Try nested conditions
			if rootValue, ok := exportsMap["."]; ok {
				if resolved := resolveSimpleExportValue(rootValue, ""); resolved != "" {
					return strings.TrimPrefix(resolved, "./"), nil
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
func readImportMapFile(path string, fs platform.FileSystem) (*ImportMap, error) {
	data, err := fs.ReadFile(path)
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
func addPackageExportsToImportMap(importMap *ImportMap, pkgName, pkgPath, rootDir string, logger types.Logger, fs platform.FileSystem) error {
	pkgJSONPath := filepath.Join(pkgPath, "package.json")
	pkg, err := readPackageJSON(pkgJSONPath, fs)
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

		case map[string]any:
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

	// No exports or main field - add fallback trailing slash mapping
	// This allows packages like @patternfly/icons to work with subpath imports
	// e.g., @patternfly/icons/fas/copy.js -> /node_modules/@patternfly/icons/fas/copy.js
	importMap.Imports[pkgName+"/"] = pkgWebPath + "/"
	if logger != nil {
		logger.Warning("Package %s has no exports or main field, added trailing slash mapping: %s/ -> %s/", pkgName, pkgName, pkgWebPath)
	}
	return nil
}

// resolveExportValue resolves an export value to a web path
func resolveExportValue(exportValue any, pkgPath, rootDir string) (string, error) {
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

	case map[string]any:
		// Conditional exports - prioritize "import" condition
		if importValue, ok := v["import"]; ok {
			// Recursively resolve if it's a nested condition
			if importPath, ok := importValue.(string); ok {
				return pkgWebPath + "/" + strings.TrimPrefix(importPath, "./"), nil
			}
			// Recursively resolve nested conditions
			if importMap, ok := importValue.(map[string]any); ok {
				return resolveExportValue(importMap, pkgPath, rootDir)
			}
		}
		// Try default condition
		if defaultValue, ok := v["default"]; ok {
			if defaultPath, ok := defaultValue.(string); ok {
				return pkgWebPath + "/" + strings.TrimPrefix(defaultPath, "./"), nil
			}
			// Recursively resolve nested default
			if defaultMap, ok := defaultValue.(map[string]any); ok {
				return resolveExportValue(defaultMap, pkgPath, rootDir)
			}
		}
	}

	return "", fmt.Errorf("could not resolve export value")
}

// addTransitiveDependenciesToScopes builds scopes for packages in the dependency tree
// by recursively traversing dependencies from the root package.
//
// Note: This function only considers dependencies installed in the top-level node_modules
// directory (rootDir/node_modules). Dependencies in nested node_modules directories
// (e.g., node_modules/foo/node_modules/bar) are not traversed unless they're hoisted to
// the root. This is generally acceptable for modern package managers (npm, pnpm, yarn)
// which hoist dependencies by default, but may result in incomplete scopes for
// non-hoisted layouts.
func addTransitiveDependenciesToScopes(importMap *ImportMap, rootDir string, rootPkg *packageJSON, logger types.Logger, fs platform.FileSystem) error {
	if rootPkg == nil {
		return nil
	}

	// Build dependency tree starting from root dependencies
	visited := make(map[string]bool)
	nodeModulesPath := filepath.Join(rootDir, "node_modules")

	// Process each root dependency recursively
	if rootPkg.Dependencies != nil {
		for depName := range rootPkg.Dependencies {
			if err := buildScopesForPackage(importMap, depName, nodeModulesPath, visited, fs); err != nil {
				if logger != nil {
					logger.Warning("Failed to build scopes for %s: %v", depName, err)
				}
			}
		}
	}

	return nil
}

// buildScopesForPackage recursively builds scopes for a package and its dependencies
func buildScopesForPackage(importMap *ImportMap, pkgName string, nodeModulesPath string, visited map[string]bool, fs platform.FileSystem) error {
	// Skip if already processed
	if visited[pkgName] {
		return nil
	}
	visited[pkgName] = true

	// Read package.json
	pkgPath := filepath.Join(nodeModulesPath, pkgName)
	pkgJSONPath := filepath.Join(pkgPath, "package.json")
	pkg, err := readPackageJSON(pkgJSONPath, fs)
	if err != nil {
		return err
	}

	// If this package has no dependencies, nothing to scope
	if len(pkg.Dependencies) == 0 {
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
		if !fs.Exists(depPath) {
			continue // Dependency not installed, skip
		}

		// Read dependency's package.json
		depPkgPath := filepath.Join(depPath, "package.json")
		depPkg, err := readPackageJSON(depPkgPath, fs)
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
		if err := buildScopesForPackage(importMap, depName, nodeModulesPath, visited, fs); err != nil {
			// Continue on error, just skip this dependency's tree
			continue
		}
	}

	return nil
}

// addDependencyExportsToScope adds dependency exports to a scope map
func addDependencyExportsToScope(
	scope map[string]string,
	depName, depWebPath string,
	exports any,
) {
	switch exp := exports.(type) {
	case string:
		// Simple string export
		scope[depName] = depWebPath + "/" + strings.TrimPrefix(exp, "./")

	case map[string]any:
		// Check if this is condition-only (no subpaths)
		hasSubpaths := false
		for exportPath := range exp {
			if strings.HasPrefix(exportPath, ".") {
				hasSubpaths = true
				break
			}
		}

		if !hasSubpaths {
			// Condition-only export, resolve it (handles nested conditions recursively)
			if importValue, ok := exp["import"]; ok {
				if resolved := resolveSimpleExportValue(importValue, depWebPath); resolved != "" {
					scope[depName] = resolved
				}
			} else if defaultValue, ok := exp["default"]; ok {
				if resolved := resolveSimpleExportValue(defaultValue, depWebPath); resolved != "" {
					scope[depName] = resolved
				}
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
func resolveSimpleExportValue(exportValue any, depWebPath string) string {
	switch v := exportValue.(type) {
	case string:
		trimmed := strings.TrimPrefix(v, "./")
		if depWebPath == "" {
			return "./" + trimmed
		}
		return depWebPath + "/" + trimmed
	case map[string]any:
		// Try import condition first
		if importValue, ok := v["import"]; ok {
			if importPath, ok := importValue.(string); ok {
				trimmed := strings.TrimPrefix(importPath, "./")
				if depWebPath == "" {
					return "./" + trimmed
				}
				return depWebPath + "/" + trimmed
			}
			// Recursively resolve nested import conditions
			if importMap, ok := importValue.(map[string]any); ok {
				return resolveSimpleExportValue(importMap, depWebPath)
			}
		}
		// Try default condition
		if defaultValue, ok := v["default"]; ok {
			if defaultPath, ok := defaultValue.(string); ok {
				trimmed := strings.TrimPrefix(defaultPath, "./")
				if depWebPath == "" {
					return "./" + trimmed
				}
				return depWebPath + "/" + trimmed
			}
			// Recursively resolve nested default
			if defaultMap, ok := defaultValue.(map[string]any); ok {
				return resolveSimpleExportValue(defaultMap, depWebPath)
			}
		}
	}
	return ""
}

// addWorkspaceScopesToImportMap adds scopes for workspace packages and root package
// so they can resolve bare specifiers to dependencies.
//
// workspaceRoot is where node_modules is located, rootDir is for calculating relative paths.
//
// Security: Scope prefixes are computed from filepath.Rel(rootDir, pkgPath). When rootDir is a
// subdirectory of workspaceRoot, workspace packages outside rootDir will yield prefixes
// containing ".." (e.g., "/../packages/components/"). This is safe because the HTTP handler
// (serve/server.go:1370-1372) normalizes all request paths and rejects path traversal attempts
// that would escape the workspace root. The ".." segments in import map scopes are merely
// browser-relative URLs that get validated server-side before serving any files.
func addWorkspaceScopesToImportMap(importMap *ImportMap, workspaceRoot, rootDir string, rootPkg *packageJSON, workspacePackages map[string]string, logger types.Logger, fs platform.FileSystem) error {
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
			if !fs.Exists(depPath) {
				if logger != nil {
					logger.Debug("Dependency %s not found in node_modules, skipping", depName)
				}
				return // Dependency not installed
			}

			depPkgPath := filepath.Join(depPath, "package.json")
			depPkg, err := readPackageJSON(depPkgPath, fs)
			if err != nil {
				if logger != nil {
					logger.Warning("Failed to read package.json for %s: %v", depName, err)
				}
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
	for pkgName, pkgPath := range workspacePackages {
		pkgJSONPath := filepath.Join(pkgPath, "package.json")
		pkg, err := readPackageJSON(pkgJSONPath, fs)
		if err != nil {
			if logger != nil {
				logger.Warning("Failed to read package.json for workspace package %s: %v", pkgName, err)
			}
			continue
		}
		if pkg.Dependencies == nil {
			continue
		}

		// Calculate web path for this workspace package
		pkgRelPath, err := filepath.Rel(rootDir, pkgPath)
		if err != nil {
			if logger != nil {
				logger.Warning("Failed to compute relative path for workspace package %s: %v", pkgName, err)
			}
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

// generateWorkspaceImportMap generates an import map for workspace mode (private implementation)
// Only includes dependencies from packages with customElements fields
// Global imports: workspace packages + their direct dependencies
// Scopes: transitive dependencies
func generateWorkspaceImportMap(workspaceRoot string, packages []middleware.WorkspacePackage, logger types.Logger, fs platform.FileSystem) (*ImportMap, error) {
	result := &ImportMap{
		Imports: make(map[string]string),
		Scopes:  make(map[string]map[string]string),
	}

	// Collect all package names that have customElements (to exclude from node_modules)
	workspacePackageNames := make(map[string]bool)
	for _, pkg := range packages {
		workspacePackageNames[pkg.Name] = true
	}

	// Add all workspace packages to global imports with their exports
	for _, pkg := range packages {
		if logger != nil {
			logger.Debug("Processing workspace package: %s at %s", pkg.Name, pkg.Path)
		}
		// Add package exports to import map (this handles subpaths like @patternfly/elements/pf-avatar/pf-avatar.js)
		if err := addPackageExportsToImportMap(result, pkg.Name, pkg.Path, workspaceRoot, logger, fs); err != nil {
			if logger != nil {
				logger.Warning("Failed to add exports for %s: %v - using fallback", pkg.Name, err)
			}
			// If exports resolution fails, add basic mapping as fallback
			pkgRelPath, err := filepath.Rel(workspaceRoot, pkg.Path)
			if err != nil {
				if logger != nil {
					logger.Warning("Failed to compute relative path for %s: %v", pkg.Name, err)
				}
				continue
			}
			pkgWebPath := "/" + filepath.ToSlash(pkgRelPath)

			if entryPoint, err := resolvePackageEntryPoint(pkg.Path, fs); err == nil && entryPoint != "" {
				result.Imports[pkg.Name] = pkgWebPath + "/" + entryPoint
				if logger != nil {
					logger.Debug("Added fallback mapping: %s -> %s", pkg.Name, pkgWebPath+"/"+entryPoint)
				}
			} else {
				result.Imports[pkg.Name+"/"] = pkgWebPath + "/"
				if logger != nil {
					logger.Debug("Added fallback trailing slash mapping: %s/ -> %s/", pkg.Name, pkgWebPath)
				}
			}
		} else {
			if logger != nil {
				logger.Debug("Successfully added exports for %s (%d imports)", pkg.Name, len(result.Imports))
			}
		}
	}

	// Collect dependencies from all packages (from package.json, not all workspace packages)
	allDeps := make(map[string]bool)
	for _, pkg := range packages {
		pkgJSONPath := filepath.Join(pkg.Path, "package.json")
		data, err := fs.ReadFile(pkgJSONPath)
		if err != nil {
			continue
		}

		var pkgJSON struct {
			Dependencies map[string]string `json:"dependencies"`
		}
		if err := json.Unmarshal(data, &pkgJSON); err != nil {
			continue
		}

		// Collect dependency names (excluding workspace packages)
		for depName := range pkgJSON.Dependencies {
			if !workspacePackageNames[depName] {
				allDeps[depName] = true
			}
		}
	}

	// Resolve all dependencies from workspace root's node_modules
	for depName := range allDeps {
		depPath := filepath.Join(workspaceRoot, "node_modules", depName)
		if !fs.Exists(depPath) {
			if logger != nil {
				logger.Debug("Dependency %s not found in workspace node_modules", depName)
			}
			continue
		}

		// Add to global imports with proper exports resolution
		if err := addPackageExportsToImportMap(result, depName, depPath, workspaceRoot, logger, fs); err != nil {
			// Skip packages that fail to resolve
			if logger != nil {
				logger.Warning("Failed to resolve exports for dependency %s in workspace mode: %v", depName, err)
			}
			continue
		}
	}

	// Add scopes for transitive dependencies
	// Convert allDeps map[string]bool to map[string]string for scope building
	depsMap := make(map[string]string)
	for depName := range allDeps {
		depsMap[depName] = "*" // Version doesn't matter for scope building
	}

	// Use the existing scope-building logic from single-package mode
	if err := addTransitiveDependenciesToScopes(result, workspaceRoot, &packageJSON{Dependencies: depsMap}, logger, fs); err != nil {
		// Log but don't fail - workspace mode is best-effort
		if logger != nil {
			logger.Warning("Failed to add transitive dependencies: %v", err)
		}
	}

	return result, nil
}
