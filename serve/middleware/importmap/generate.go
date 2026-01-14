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
	"bennypowers.dev/mappa/packagejson"
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
// This is an internal version that accepts a FileSystem for testing.
// Production code should use workspace.FindWorkspaceRoot instead.
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

// Generate generates an import map from package.json and configuration.
//
// In workspace mode (config.WorkspacePackages is set), it generates an import map
// that only includes packages with customElements fields - this is cem-specific behavior.
//
// In single-package mode, it uses mappa for import map generation with automatic
// workspace discovery and scope handling.
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

	// Workspace mode: generate flattened import map from packages with customElements.
	// This is cem-specific behavior - only packages with customElements fields are included.
	if len(config.WorkspacePackages) > 0 {
		result, err := generateWorkspaceImportMap(rootDir, config.WorkspacePackages, config.Logger, fs)
		if err != nil {
			return nil, err
		}

		// Apply user overrides and CLI overrides
		if err := applyOverrides(result, config, fs); err != nil {
			return nil, err
		}

		return result, nil
	}

	// Single-package mode: use mappa for import map generation.
	// Mappa handles workspace auto-discovery from the workspaces field.

	// Check if package.json exists
	rootPkgPath := filepath.Join(rootDir, "package.json")
	rootPkg, err := readPackageJSON(rootPkgPath, fs)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// Empty package.json is OK - apply user/CLI overrides and return
			result := &ImportMap{
				Imports: make(map[string]string),
				Scopes:  make(map[string]map[string]string),
			}
			if err := applyOverrides(result, config, fs); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("reading package.json: %w", err)
	}

	// Find workspace root using cem's logic (respects git submodule boundaries)
	workspaceRoot := findWorkspaceRoot(rootDir, fs)
	if workspaceRoot != rootDir && config.Logger != nil {
		config.Logger.Debug("Detected workspace subdirectory, using workspace root: %s", workspaceRoot)
	}

	// Determine if we should include root exports
	// Only include if not a monorepo and has exports or main
	includeRootExports := rootPkg.Workspaces == nil && rootPkg.Name != "" &&
		(rootPkg.Exports != nil || rootPkg.Main != "")

	// Use mappa for import map resolution (auto-discovers workspace packages)
	result, err := resolveWithMappa(
		workspaceRoot, // Use workspace root for resolution
		fs,
		config.Logger,
		nil, // No input map - we'll apply overrides separately
		nil, // Let mappa auto-discover workspace packages
		includeRootExports,
	)
	if err != nil {
		return nil, fmt.Errorf("generating import map: %w", err)
	}

	// Apply user overrides and CLI overrides
	if err := applyOverrides(result, config, fs); err != nil {
		return nil, err
	}

	return result, nil
}

// GenerateWithGraph generates an import map and returns a dependency graph for incremental updates.
// Use this for initial resolution in scenarios where you'll need incremental updates (e.g., hot-reload).
//
// The returned IncrementalResult contains both the import map and a dependency graph that tracks
// which packages depend on which. This graph can be passed to GenerateIncremental for efficient
// updates when packages change.
func GenerateWithGraph(rootDir string, config *Config) (*IncrementalResult, error) {
	// Default config to empty if not provided
	if config == nil {
		config = &Config{}
	}

	// Default to OS filesystem if not provided
	fs := config.FS
	if fs == nil {
		fs = platform.NewOSFileSystem()
	}

	// Normalize rootDir
	if absRoot, err := filepath.Abs(rootDir); err == nil {
		rootDir = absRoot
	}

	// Workspace mode with explicit packages is not supported for incremental updates
	// (the packages list can change between calls, making incremental updates unreliable)
	if len(config.WorkspacePackages) > 0 {
		// Fall back to full generation
		importMap, err := generateWorkspaceImportMap(rootDir, config.WorkspacePackages, config.Logger, fs)
		if err != nil {
			return nil, err
		}
		if err := applyOverrides(importMap, config, fs); err != nil {
			return nil, err
		}
		return &IncrementalResult{ImportMap: importMap, Graph: nil}, nil
	}

	// Check if package.json exists
	rootPkgPath := filepath.Join(rootDir, "package.json")
	rootPkg, err := readPackageJSON(rootPkgPath, fs)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			result := &ImportMap{
				Imports: make(map[string]string),
				Scopes:  make(map[string]map[string]string),
			}
			if err := applyOverrides(result, config, fs); err != nil {
				return nil, err
			}
			return &IncrementalResult{ImportMap: result, Graph: nil}, nil
		}
		return nil, fmt.Errorf("reading package.json: %w", err)
	}

	// Find workspace root
	workspaceRoot := findWorkspaceRoot(rootDir, fs)

	// Determine if we should include root exports
	includeRootExports := rootPkg.Workspaces == nil && rootPkg.Name != "" &&
		(rootPkg.Exports != nil || rootPkg.Main != "")

	// Create a persistent cache for incremental updates
	cache := packagejson.NewMemoryCache()

	// Use mappa for resolution with graph
	result, err := resolveWithMappaGraph(
		workspaceRoot,
		fs,
		config.Logger,
		nil,
		nil, // Let mappa auto-discover workspace packages
		includeRootExports,
		cache,
	)
	if err != nil {
		return nil, fmt.Errorf("generating import map: %w", err)
	}

	// Apply user overrides
	if err := applyOverrides(result.ImportMap, config, fs); err != nil {
		return nil, err
	}

	return result, nil
}

// GenerateIncremental performs an incremental import map update.
// Only the changed packages and their transitive dependents are re-resolved.
//
// This is much faster than full regeneration for hot-reload scenarios where only
// a single package.json file changed.
//
// If the update cannot be performed incrementally (e.g., missing previous state),
// it falls back to full resolution.
func GenerateIncremental(rootDir string, config *Config, update IncrementalUpdate) (*IncrementalResult, error) {
	// If no previous state, fall back to full generation
	if update.PreviousMap == nil || update.PreviousGraph == nil || len(update.ChangedPackages) == 0 {
		return GenerateWithGraph(rootDir, config)
	}

	// Default config to empty if not provided
	if config == nil {
		config = &Config{}
	}

	// Default to OS filesystem if not provided
	fs := config.FS
	if fs == nil {
		fs = platform.NewOSFileSystem()
	}

	// Normalize rootDir
	if absRoot, err := filepath.Abs(rootDir); err == nil {
		rootDir = absRoot
	}

	// Find workspace root
	workspaceRoot := findWorkspaceRoot(rootDir, fs)

	// Create cache for this resolution
	cache := packagejson.NewMemoryCache()

	// Perform incremental resolution
	result, err := resolveIncrementalWithMappa(
		workspaceRoot,
		fs,
		config.Logger,
		update,
		cache,
	)
	if err != nil {
		return nil, fmt.Errorf("incremental import map update: %w", err)
	}

	// Apply user overrides
	if err := applyOverrides(result.ImportMap, config, fs); err != nil {
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

// generateWorkspaceImportMap generates an import map for workspace mode (private implementation).
// Only includes dependencies from packages with customElements fields.
// This is cem-specific behavior - mappa handles the actual import map generation,
// but we pre-filter workspace packages to only those with customElements.
func generateWorkspaceImportMap(workspaceRoot string, packages []middleware.WorkspacePackage, logger types.Logger, fs platform.FileSystem) (*ImportMap, error) {
	if logger != nil {
		logger.Debug("Generating workspace import map for %d packages with customElements", len(packages))
	}

	// Convert middleware.WorkspacePackage to workspacePackageInfo for mappa adapter
	workspacePackages := make([]workspacePackageInfo, len(packages))
	for i, pkg := range packages {
		workspacePackages[i] = workspacePackageInfo{
			Name: pkg.Name,
			Path: pkg.Path,
		}
		if logger != nil {
			logger.Debug("Processing workspace package: %s at %s", pkg.Name, pkg.Path)
		}
	}

	// Use mappa for import map resolution with the pre-filtered workspace packages
	result, err := resolveWithMappa(
		workspaceRoot,
		fs,
		logger,
		nil,               // No input map - overrides applied separately
		workspacePackages, // Pre-filtered packages with customElements
		false,             // Don't include root exports in workspace mode
	)
	if err != nil {
		return nil, fmt.Errorf("generating workspace import map: %w", err)
	}

	if logger != nil {
		logger.Debug("Generated workspace import map with %d imports and %d scopes",
			len(result.Imports), len(result.Scopes))
	}

	return result, nil
}
