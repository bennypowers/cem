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
	if strings.HasPrefix(importPath, "./") {
		// Convert './my-icon.js' to 'my-icon.js'
		return strings.TrimPrefix(importPath, "./")
	}

	// Handle parent directory imports like '../shared/utils.js'
	if strings.HasPrefix(importPath, "../") {
		// For now, just remove the prefix - more sophisticated resolution
		// could be added if needed
		return strings.TrimPrefix(importPath, "../")
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

	// Check file name matching (last resort)
	importFile := strings.TrimSuffix(strings.Split(normalizedImport, "/")[len(strings.Split(normalizedImport, "/"))-1], ".js")
	sourceFile := strings.TrimSuffix(strings.Split(normalizedSource, "/")[len(strings.Split(normalizedSource, "/"))-1], ".js")
	return importFile == sourceFile
}
