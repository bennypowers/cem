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
package helpers

import (
	"strings"
)

// NormalizeImportPath converts a relative import path to an absolute module path
// This handles simple relative path prefixes for module graph building
func NormalizeImportPath(importPath string) string {
	// Handle relative imports like './my-icon.js'
	if after, ok := strings.CutPrefix(importPath, "./"); ok {
		// Convert './my-icon.js' to 'my-icon.js'
		return after
	}

	// Handle parent directory imports like '../shared/utils.js'
	if after, ok := strings.CutPrefix(importPath, "../"); ok {
		// For now, just remove the prefix - more sophisticated resolution
		// could be added if needed
		return after
	}

	// Handle absolute imports (npm packages, etc.) - return as-is
	return importPath
}

// NormalizePathForMatching normalizes a file path for comparison in diagnostics
// This handles more sophisticated path normalization for import path matching
func NormalizePathForMatching(path string) string {
	// Remove common prefixes/suffixes
	path = strings.TrimPrefix(path, "./")
	path = strings.TrimPrefix(path, "../")
	path = strings.TrimPrefix(path, "/")

	// Handle npm package paths like @rhds/elements/rh-card/rh-card.js
	// vs manifest paths like ./dist/rh-card.js
	if strings.Contains(path, "/") {
		// Keep the last two segments for better matching
		parts := strings.Split(path, "/")
		if len(parts) >= 2 {
			return strings.Join(parts[len(parts)-2:], "/")
		}
	}

	return path
}

// PathsMatch checks if an import path matches an element source path using sophisticated matching logic
func PathsMatch(importPath, elementSource string) bool {
	// IMPORTANT: Check for package-level bare specifier imports FIRST (before normalization)
	// When importing "package-name" (without subpath), it should match elements from that package
	// like "package-name/index.js", "package-name/button.js", etc.
	// This handles self-imports where a package imports itself by name.
	//
	// Examples:
	// - "large-component-library" matches "large-component-library/index.js"
	// - "@rhds/elements" matches "@rhds/elements/rh-card/rh-card.js"
	// - "my-package" matches "my-package/src/component.js"
	//
	// Use the original (non-normalized) paths for this check to preserve package names
	if strings.HasPrefix(elementSource, importPath+"/") {
		return true
	}

	// Normalize both paths for comparison
	normalizedImport := NormalizePathForMatching(importPath)
	normalizedSource := NormalizePathForMatching(elementSource)

	// Direct match on normalized paths
	if normalizedImport == normalizedSource {
		return true
	}

	// Check if import path ends with the element source (relative imports)
	if strings.HasSuffix(normalizedImport, normalizedSource) {
		return true
	}

	// Check if element source ends with import path (package imports)
	if strings.HasSuffix(normalizedSource, normalizedImport) {
		return true
	}

	// Also check normalized paths for package prefix matching
	// This handles cases where elementSource has been normalized but still contains package structure
	if strings.HasPrefix(normalizedSource, normalizedImport+"/") {
		return true
	}

	// Check file name matching (last resort)
	importParts := strings.Split(normalizedImport, "/")
	sourceParts := strings.Split(normalizedSource, "/")
	importFile := strings.TrimSuffix(importParts[len(importParts)-1], ".js")
	sourceFile := strings.TrimSuffix(sourceParts[len(sourceParts)-1], ".js")
	return importFile == sourceFile
}
