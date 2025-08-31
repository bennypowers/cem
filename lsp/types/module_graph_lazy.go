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
package types

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/queries"
)

// Lazy Module Graph Building Methods
// These methods provide on-demand module graph construction for specific import paths
// rather than analyzing the entire workspace upfront.

// SetWorkspaceRoot stores the workspace root for lazy file parsing
func (mg *ModuleGraph) SetWorkspaceRoot(workspaceRoot string) {
	mg.workspaceRoot = workspaceRoot
}

// BuildForImportPath builds module graph entries for a specific import path
func (mg *ModuleGraph) BuildForImportPath(importPath string) error {
	if mg.workspaceRoot == "" {
		return fmt.Errorf("workspace root not set for lazy module graph building")
	}

	start := time.Now()
	defer func() {
		mg.metrics.RecordDuration("lazy_build_time", time.Since(start))
	}()

	mg.metrics.IncrementCounter("lazy_build_calls")

	// First, check if this import path corresponds to any known manifest modules
	// The key insight: instead of resolving to files and parsing,
	// we should discover dependencies between manifest modules
	manifestModule := mg.findManifestModuleForImportPath(importPath)
	if manifestModule != "" {
		helpers.SafeDebugLog("[MODULE_GRAPH] Found manifest module '%s' for import '%s'", manifestModule, importPath)
		// Parse the source file for this manifest module to discover its dependencies
		if err := mg.buildDependenciesForManifestModule(manifestModule); err != nil {
			helpers.SafeDebugLog("[MODULE_GRAPH] Warning: Failed to build dependencies for manifest module %s: %v", manifestModule, err)
		}
	}

	// Fallback: resolve to file paths and parse (for non-manifest imports)
	filePaths := mg.resolveImportPathToFiles(importPath)
	if len(filePaths) > 0 {
		helpers.SafeDebugLog("[MODULE_GRAPH] Building graph for import %s -> %d files", importPath, len(filePaths))

		// Parse only the specific files related to this import
		// Use a set to track already processed files to avoid infinite recursion
		processedFiles := make(map[string]bool)

		for _, filePath := range filePaths {
			if err := mg.processFileWithDependencies(filePath, processedFiles); err != nil {
				mg.metrics.IncrementCounter("lazy_parse_errors")
				helpers.SafeDebugLog("[MODULE_GRAPH] Warning: Failed to process file %s: %v", filePath, err)
				// Continue with other files
			}
		}
	}

	// CRITICAL: Resolve re-export chains after parsing files
	// This is essential for making re-exported elements available
	// (e.g., my-tab available through my-tabs.js re-export)
	mg.resolveReExportChains()
	helpers.SafeDebugLog("[MODULE_GRAPH] Resolved re-export chains for lazy build of %s", importPath)

	// Cache invalidation is handled automatically by AddModuleDependency() calls
	// during the recursive dependency building process, so no additional cache clearing needed

	mg.metrics.IncrementCounter("lazy_imports_built")
	return nil
}

// resolveImportPathToFiles converts an import path to actual file paths that need to be parsed
// This function resolves relative imports relative to the workspace root (legacy behavior)
func (mg *ModuleGraph) resolveImportPathToFiles(importPath string) []string {
	return mg.resolveImportPathToFilesRelativeTo(importPath, "")
}

// resolveImportPathToFilesRelativeTo converts an import path to actual file paths,
// resolving relative imports relative to the given file path
func (mg *ModuleGraph) resolveImportPathToFilesRelativeTo(importPath, currentFilePath string) []string {
	var filePaths []string

	helpers.SafeDebugLog("[MODULE_GRAPH] 🔍 Resolving import path '%s' to files in workspace '%s' (relative to '%s')", importPath, mg.workspaceRoot, currentFilePath)

	// Handle different import path formats
	if strings.HasPrefix(importPath, "./") || strings.HasPrefix(importPath, "../") {
		// Relative import - resolve relative to current file's directory
		var baseDir string
		if currentFilePath != "" {
			// Get the directory of the current file
			baseDir = filepath.Dir(currentFilePath)
		} else {
			// Fallback to workspace root for backward compatibility
			baseDir = mg.workspaceRoot
		}

		fullPath := filepath.Join(baseDir, importPath)
		helpers.SafeDebugLog("[MODULE_GRAPH] 📁 Trying relative path: %s (base dir: %s)", fullPath, baseDir)
		if mg.fileExists(fullPath) {
			filePaths = append(filePaths, fullPath)
			helpers.SafeDebugLog("[MODULE_GRAPH] ✅ Found relative file: %s", fullPath)
		}

		// For imports like './rh-tab.js', also try corresponding TypeScript source
		// This handles TypeScript projects where imports reference .js but sources are .ts
		if strings.HasSuffix(importPath, ".js") {
			tsPath := strings.TrimSuffix(fullPath, ".js") + ".ts"
			if mg.fileExists(tsPath) {
				filePaths = append(filePaths, tsPath)
				helpers.SafeDebugLog("[MODULE_GRAPH] ✅ Found TypeScript source: %s", tsPath)
			}
		}

		// Also try with common extensions
		for _, ext := range []string{".js", ".ts", ".mjs"} {
			extPath := fullPath + ext
			if mg.fileExists(extPath) {
				// Avoid duplicates
				isDuplicate := false
				for _, existing := range filePaths {
					if existing == extPath {
						isDuplicate = true
						break
					}
				}
				if !isDuplicate {
					filePaths = append(filePaths, extPath)
					helpers.SafeDebugLog("[MODULE_GRAPH] ✅ Found file with extension: %s", extPath)
				}
			}
		}
	} else if strings.Contains(importPath, "/") {
		// Package import like '@rhds/elements/rh-tabs/rh-tabs.js' - look in workspace for matching files
		// This handles cases where packages are developed locally

		// Extract the meaningful part after the package name
		// For '@rhds/elements/rh-tabs/rh-tabs.js' -> 'rh-tabs/rh-tabs.js'
		var relativePath string
		if strings.HasPrefix(importPath, "@") {
			// Handle scoped packages like @rhds/elements/path/file.js
			parts := strings.SplitN(importPath, "/", 3)
			helpers.SafeDebugLog("[MODULE_GRAPH] 📦 Scoped package parts: %v", parts)
			if len(parts) >= 3 {
				relativePath = parts[2] // Get 'rh-tabs/rh-tabs.js'
				helpers.SafeDebugLog("[MODULE_GRAPH] 📁 Extracted relative path: %s", relativePath)
			} else {
				relativePath = filepath.Base(importPath)
				helpers.SafeDebugLog("[MODULE_GRAPH] 📁 Fallback to base path: %s", relativePath)
			}
		} else {
			// Handle regular packages like some-package/path/file.js
			parts := strings.SplitN(importPath, "/", 2)
			helpers.SafeDebugLog("[MODULE_GRAPH] 📦 Regular package parts: %v", parts)
			if len(parts) >= 2 {
				relativePath = parts[1]
				helpers.SafeDebugLog("[MODULE_GRAPH] 📁 Extracted relative path: %s", relativePath)
			} else {
				relativePath = filepath.Base(importPath)
				helpers.SafeDebugLog("[MODULE_GRAPH] 📁 Fallback to base path: %s", relativePath)
			}
		}

		possiblePaths := []string{
			filepath.Join(mg.workspaceRoot, importPath),               // Try full import path
			filepath.Join(mg.workspaceRoot, relativePath),             // Try relative path
			filepath.Join(mg.workspaceRoot, "elements", relativePath), // Try in elements/ dir
			filepath.Join(mg.workspaceRoot, "src", relativePath),      // Try in src/ dir
			filepath.Join(mg.workspaceRoot, "lib", relativePath),      // Try in lib/ dir
		}

		helpers.SafeDebugLog("[MODULE_GRAPH] 🔍 Checking %d possible paths for import '%s'", len(possiblePaths), importPath)
		for i, path := range possiblePaths {
			helpers.SafeDebugLog("[MODULE_GRAPH] 📁 Path %d: %s", i+1, path)
			if mg.fileExists(path) {
				filePaths = append(filePaths, path)
				helpers.SafeDebugLog("[MODULE_GRAPH] ✅ FOUND: %s", path)
			} else {
				helpers.SafeDebugLog("[MODULE_GRAPH] ❌ Not found: %s", path)
			}
		}
	}

	return filePaths
}

// fileExists checks if a file exists
func (mg *ModuleGraph) fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// processFile processes a single file and adds its exports to the module graph
func (mg *ModuleGraph) processFile(filePath string) error {
	helpers.SafeDebugLog("[MODULE_GRAPH] Processing file: %s", filePath)

	// Convert to relative path for the module graph
	relPath, err := filepath.Rel(mg.workspaceRoot, filePath)
	if err != nil {
		relPath = filePath // Fallback to absolute path
	}

	// Parse exports from this file
	parseStart := time.Now()
	parseErr := mg.parseFileExports(filePath, mg.workspaceRoot)
	mg.metrics.RecordDuration("lazy_file_parse_time", time.Since(parseStart))

	if parseErr != nil {
		mg.metrics.IncrementCounter("lazy_parse_errors")
		return parseErr
	}

	mg.metrics.IncrementCounter("lazy_files_processed")
	helpers.SafeDebugLog("[MODULE_GRAPH] Successfully processed file: %s", relPath)
	return nil
}

// processFileWithDependencies processes a file and recursively processes its dependencies
func (mg *ModuleGraph) processFileWithDependencies(filePath string, processedFiles map[string]bool) error {
	// Check if already processed to avoid infinite recursion
	if processedFiles[filePath] {
		helpers.SafeDebugLog("[MODULE_GRAPH] Skipping already processed file: %s", filePath)
		return nil
	}

	// Mark as processed
	processedFiles[filePath] = true
	helpers.SafeDebugLog("[MODULE_GRAPH] Processing file with dependencies: %s", filePath)

	// Process the file itself
	if err := mg.processFile(filePath); err != nil {
		return err
	}

	// Get dependencies (imports) from this file and process them recursively
	dependencies, err := mg.extractFileDependencies(filePath)
	if err != nil {
		helpers.SafeDebugLog("[MODULE_GRAPH] Warning: Failed to extract dependencies from %s: %v", filePath, err)
		// Continue without dependencies rather than failing completely
		return nil
	}

	helpers.SafeDebugLog("[MODULE_GRAPH] Found %d dependencies for %s: %v", len(dependencies), filePath, dependencies)

	// Recursively process each dependency
	for _, depPath := range dependencies {
		if err := mg.processFileWithDependencies(depPath, processedFiles); err != nil {
			helpers.SafeDebugLog("[MODULE_GRAPH] Warning: Failed to process dependency %s: %v", depPath, err)
			// Continue with other dependencies
		}
	}

	return nil
}

// extractFileDependencies extracts the import dependencies from a file
func (mg *ModuleGraph) extractFileDependencies(filePath string) ([]string, error) {
	content, err := mg.fileParser.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	// Get TypeScript parser from pool
	parser := queries.RetrieveTypeScriptParser()
	defer queries.PutTypeScriptParser(parser)

	// Parse the content
	tree := parser.Parse(content, nil)
	if tree == nil {
		helpers.SafeDebugLog("[MODULE_GRAPH] DEBUG: Tree parsing failed for %s", filePath)
		return nil, nil // Skip files that fail to parse
	}
	defer tree.Close()
	helpers.SafeDebugLog("[MODULE_GRAPH] DEBUG: Successfully parsed tree for %s", filePath)

	// Get cached export matcher (which includes import queries)
	if mg.queryManager == nil {
		return nil, fmt.Errorf("query manager not available for extracting dependencies from file %s", filePath)
	}

	exportMatcher, err := queries.GetCachedQueryMatcher(mg.queryManager, "typescript", "exports")
	if err != nil {
		return nil, fmt.Errorf("failed to get cached export matcher for file %s: %w", filePath, err)
	}

	var dependencies []string
	dependencySet := make(map[string]bool) // Use set to avoid duplicates

	// Process matches to extract import sources
	matchCount := 0
	for match := range exportMatcher.AllQueryMatches(tree.RootNode(), content) {
		matchCount++
		for _, capture := range match.Captures {
			captureName := exportMatcher.GetCaptureNameByIndex(capture.Index)
			captureText := strings.TrimSpace(capture.Node.Utf8Text(content))

			// Look for import sources
			if captureName == "import.source" || captureName == "import.dynamic.source" {
				if captureText != "" && !dependencySet[captureText] {
					helpers.SafeDebugLog("[MODULE_GRAPH] Processing import source: %s", captureText)
					// Resolve import path to file paths relative to current file
					resolvedPaths := mg.resolveImportPathToFilesRelativeTo(captureText, filePath)
					helpers.SafeDebugLog("[MODULE_GRAPH] Resolved import '%s' to %d paths: %v", captureText, len(resolvedPaths), resolvedPaths)

					// Add all resolved file paths as dependencies
					for _, resolvedPath := range resolvedPaths {
						if !dependencySet[resolvedPath] {
							dependencies = append(dependencies, resolvedPath)
							dependencySet[resolvedPath] = true
						}
					}
				}
			}
		}
	}

	helpers.SafeDebugLog("[MODULE_GRAPH] Extracted %d dependencies from %s", len(dependencies), filePath)
	return dependencies, nil
}

// findManifestModuleForImportPath finds a manifest module that matches an import path
// Returns the manifest module path if found, empty string otherwise
func (mg *ModuleGraph) findManifestModuleForImportPath(importPath string) string {
	if mg.manifestResolver == nil {
		return ""
	}

	// Use the ManifestResolver to find matching manifest modules
	manifestModules := mg.manifestResolver.FindManifestModulesForImportPath(importPath)
	if len(manifestModules) > 0 {
		// Return the first matching module - in practice, there should usually be only one
		return manifestModules[0]
	}

	return ""
}

// buildDependenciesForManifestModule builds the complete dependency tree for a manifest module
// by recursively tracking all imports and mapping them to manifest modules.
// This ensures transitive dependency resolution works correctly (e.g., importing my-tabs.js
// provides access to transitively imported elements like my-icon from my-tab.js).
func (mg *ModuleGraph) buildDependenciesForManifestModule(manifestModule string) error {
	visited := make(map[string]bool)
	return mg.buildDependenciesRecursive(manifestModule, visited)
}

// buildDependenciesRecursive recursively builds dependencies for a manifest module and all
// its transitive dependencies. The visited map prevents infinite recursion in circular
// dependency scenarios while ensuring complete dependency tree construction.
func (mg *ModuleGraph) buildDependenciesRecursive(manifestModule string, visited map[string]bool) error {
	if mg.workspaceRoot == "" {
		helpers.SafeDebugLog("[MODULE_GRAPH] No workspace root set, cannot build dependencies for manifest module: %s", manifestModule)
		return nil
	}

	// Check if we've already processed this module to prevent infinite recursion
	if visited[manifestModule] {
		helpers.SafeDebugLog("[MODULE_GRAPH] Skipping already processed manifest module: %s", manifestModule)
		return nil
	}
	visited[manifestModule] = true

	helpers.SafeDebugLog("[MODULE_GRAPH] Building dependencies for manifest module: %s", manifestModule)

	// Find the source file for this manifest module to parse its imports
	sourceFileCandidates := mg.findSourceFileForManifestModule(manifestModule)
	if len(sourceFileCandidates) == 0 {
		helpers.SafeDebugLog("[MODULE_GRAPH] No source file candidates found for manifest module: %s", manifestModule)
		return nil
	}

	// Try each source file candidate until we find one that exists and can be parsed
	for _, sourceFile := range sourceFileCandidates {
		if !mg.fileExists(sourceFile) {
			continue
		}

		helpers.SafeDebugLog("[MODULE_GRAPH] Found source file for manifest module '%s': %s", manifestModule, sourceFile)

		// Parse the source file to extract its import dependencies
		dependencies, err := mg.extractFileDependencies(sourceFile)
		if err != nil {
			helpers.SafeDebugLog("[MODULE_GRAPH] Warning: Failed to extract dependencies from %s: %v", sourceFile, err)
			continue // Try next candidate
		}

		// Process each dependency and recursively build its dependency tree
		mg.processDependencies(manifestModule, dependencies, visited)

		helpers.SafeDebugLog("[MODULE_GRAPH] Successfully built dependencies for manifest module '%s' with %d dependencies", manifestModule, len(dependencies))
		return nil
	}

	helpers.SafeDebugLog("[MODULE_GRAPH] Could not find or parse source file for manifest module: %s", manifestModule)
	return nil
}

// processDependencies processes all file dependencies for a manifest module, mapping them
// to manifest modules and recursively building their dependency trees.
func (mg *ModuleGraph) processDependencies(manifestModule string, dependencies []string, visited map[string]bool) {
	for _, depFilePath := range dependencies {
		// Convert file path to manifest module path using ManifestResolver
		depManifestModule := mg.manifestResolver.GetManifestModulePath(depFilePath)
		if depManifestModule == "" {
			helpers.SafeDebugLog("[MODULE_GRAPH] No manifest module found for dependency file: %s", depFilePath)
			continue
		}

		// Track the module dependency relationship at manifest level
		mg.dependencyTracker.AddModuleDependency(manifestModule, depManifestModule)
		helpers.SafeDebugLog("[MODULE_GRAPH] Manifest module dependency: %s -> %s (via file %s)", manifestModule, depManifestModule, depFilePath)

		// Recursively build dependencies for the dependency module
		if err := mg.buildDependenciesRecursive(depManifestModule, visited); err != nil {
			helpers.SafeDebugLog("[MODULE_GRAPH] Warning: Failed to recursively build dependencies for %s: %v", depManifestModule, err)
			// Continue with other dependencies rather than failing completely
		}
	}
}

// findSourceFileForManifestModule attempts to find the source file for a given manifest module
func (mg *ModuleGraph) findSourceFileForManifestModule(manifestModule string) []string {
	var candidates []string

	// Common patterns for finding source files from manifest modules:
	// 1. "rh-tabs/rh-tabs.js" -> "elements/rh-tabs/rh-tabs.ts" (Red Hat Design System)
	// 2. "rh-tabs/rh-tabs.js" -> "rh-tabs/rh-tabs.ts" (direct mapping)
	// 3. "rh-tabs/rh-tabs.js" -> "src/rh-tabs/rh-tabs.ts" (src directory)
	// 4. "rh-tabs/rh-tabs.js" -> "lib/rh-tabs/rh-tabs.ts" (lib directory)

	// Extract path without extension
	pathWithoutExt := strings.TrimSuffix(manifestModule, filepath.Ext(manifestModule))

	// Try different directory structures and file extensions
	prefixes := []string{"", "elements/", "src/", "lib/", "components/"}
	extensions := []string{".ts", ".js", ".mjs"}

	for _, prefix := range prefixes {
		for _, ext := range extensions {
			candidate := filepath.Join(mg.workspaceRoot, prefix+pathWithoutExt+ext)
			candidates = append(candidates, candidate)
		}
	}

	helpers.SafeDebugLog("[MODULE_GRAPH] Generated %d source file candidates for manifest module '%s': %v", len(candidates), manifestModule, candidates)
	return candidates
}
