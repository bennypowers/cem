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
	source         []byte
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
