/*
Copyright © 2025 Benny Powers

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
package manifest

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

var ErrNotExported = errors.New("not exported by package.json")

// PackageJSON represents the subset of package.json we care about.
type PackageJSON struct {
	Name           string `json:"name"`
	Version        string `json:"version"`
	Exports        any    `json:"exports,omitempty"`
	CustomElements string `json:"customElements"`
	Types          string `json:"types,omitempty"`
	Typings        string `json:"typings,omitempty"`
	Main           string `json:"main,omitempty"`
}

// ResolveImportSubpath resolves an import subpath to a file path within the package.
// subpath is like "./lib/types" or "." for root.
// Returns the resolved file path relative to the package root.
// This is the inverse of ResolveExportPath: given a consumer-facing subpath,
// it returns the actual file path within the package.
func ResolveImportSubpath(pkg *PackageJSON, subpath string) (string, error) {
	if pkg == nil {
		return "", fmt.Errorf("package.json is nil")
	}

	// Normalize subpath
	if subpath == "" {
		subpath = "."
	}

	// Try exports map first
	if pkg.Exports != nil {
		resolved, err := resolveSubpathFromExports(pkg.Exports, subpath)
		if err == nil {
			return resolved, nil
		}
	}

	// Fallback for root subpath: use types → typings → main fields
	if subpath == "." {
		if pkg.Types != "" {
			return strings.TrimPrefix(pkg.Types, "./"), nil
		}
		if pkg.Typings != "" {
			return strings.TrimPrefix(pkg.Typings, "./"), nil
		}
		if pkg.Main != "" {
			return strings.TrimPrefix(pkg.Main, "./"), nil
		}
	}

	return "", fmt.Errorf("cannot resolve subpath %q: %w", subpath, ErrNotExported)
}

// resolveSubpathFromExports resolves a subpath against the exports field.
func resolveSubpathFromExports(exports any, subpath string) (string, error) {
	// Handle string exports (single export)
	if expStr, ok := exports.(string); ok {
		if subpath == "." {
			return strings.TrimPrefix(expStr, "./"), nil
		}
		return "", fmt.Errorf("subpath %q not found in string export", subpath)
	}

	exportsMap, ok := exports.(map[string]any)
	if !ok {
		return "", fmt.Errorf("exports field is not a map")
	}

	// Direct match for the subpath key
	if expVal, found := exportsMap[subpath]; found {
		return resolveExportValue(expVal)
	}

	// Wildcard pattern matching
	for key, val := range exportsMap {
		if !strings.Contains(key, "*") {
			continue
		}
		keyPrefix := strings.Split(key, "*")[0]
		keySuffix := strings.Split(key, "*")[1]
		if strings.HasPrefix(subpath, keyPrefix) && strings.HasSuffix(subpath, keySuffix) {
			starValue := subpath[len(keyPrefix):]
			if keySuffix != "" {
				starValue = starValue[:len(starValue)-len(keySuffix)]
			}
			resolved, err := resolveExportValue(val)
			if err != nil {
				continue
			}
			return strings.Replace(resolved, "*", starValue, 1), nil
		}
	}

	return "", fmt.Errorf("subpath %q not found in exports", subpath)
}

// resolveExportValue extracts a file path from an export value,
// preferring the "types" condition, then "import", then "default".
func resolveExportValue(val any) (string, error) {
	switch v := val.(type) {
	case string:
		return strings.TrimPrefix(v, "./"), nil
	case map[string]any:
		// Prefer types condition for type resolution
		for _, condition := range []string{"types", "import", "default"} {
			if cond, ok := v[condition]; ok {
				if s, ok := cond.(string); ok {
					return strings.TrimPrefix(s, "./"), nil
				}
			}
		}
	}
	return "", fmt.Errorf("cannot resolve export value")
}

// ResolveExportPath resolves a file path relative to the package root
// through the `exports` block in package.json, returning the corresponding
// package path as it would be used in an import.
// The returned path is always without a leading './'.
func ResolveExportPath(packageJson *PackageJSON, relFilePath string) (string, error) {
	if packageJson == nil {
		return relFilePath, nil
	}
	cleanRel := filepath.ToSlash(relFilePath)
	cleanRel = strings.TrimPrefix(cleanRel, "./")

	// Handle string exports (single export).
	if expStr, ok := packageJson.Exports.(string); ok {
		exportFile := strings.TrimPrefix(expStr, "./")
		if cleanRel == exportFile {
			return "", nil // returns ""
		}
		return "", fmt.Errorf("%s: %w", relFilePath, ErrNotExported)
	}

	exportsMap, ok := packageJson.Exports.(map[string]any)
	if !ok {
		// gracefully skip, treat as "not exported"
		return relFilePath, nil
	}

	for expKey, expVal := range exportsMap {
		if strings.HasPrefix(expKey, "./") {
			expTarget := ""
			switch v := expVal.(type) {
			case string:
				expTarget = strings.TrimPrefix(v, "./")
			case map[string]any:
				// For condition objects, try "default" or first string.
				if def, ok := v["default"]; ok {
					if s, ok := def.(string); ok {
						expTarget = strings.TrimPrefix(s, "./")
					}
				} else {
					for _, v2 := range v {
						if s, ok := v2.(string); ok {
							expTarget = strings.TrimPrefix(s, "./")
							break
						}
					}
				}
			}

			if expTarget == cleanRel {
				return strings.TrimPrefix(expKey, "./"), nil
			}

			// Fixed: subpath pattern matching, match cleanRel against expTarget with *
			if strings.Contains(expTarget, "*") {
				targetPrefix := strings.Split(expTarget, "*")[0]
				targetSuffix := strings.Split(expTarget, "*")[1]
				if strings.HasPrefix(cleanRel, targetPrefix) && strings.HasSuffix(cleanRel, targetSuffix) {
					starValue := cleanRel[len(targetPrefix):]
					if targetSuffix != "" {
						starValue = starValue[:len(starValue)-len(targetSuffix)]
					}
					// Form the package path by substituting in the starValue to the export key
					return strings.TrimPrefix(strings.Replace(expKey, "*", starValue, 1), "./"), nil
				}
			}
		}

		// Special case for "." export
		if expKey == "." && expVal != nil {
			target := ""
			switch v := expVal.(type) {
			case string:
				target = strings.TrimPrefix(v, "./")
			case map[string]any:
				if def, ok := v["default"]; ok {
					if s, ok := def.(string); ok {
						target = strings.TrimPrefix(s, "./")
					}
				}
			}
			if target == cleanRel {
				return "", nil // root export, no path
			}
		}
	}

	return "", fmt.Errorf("%s: %w", relFilePath, ErrNotExported)
}
