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
	"fmt"
	"strings"

	importmappkg "bennypowers.dev/cem/serve/middleware/importmap"
)

// checkBareSpecifiers extracts import specifiers from transformed code and warns
// about bare specifiers that resolve to transitive dependencies (not in the import map).
// This helps users understand why the browser shows "bare specifier not remapped" errors.
func (s *Server) checkBareSpecifiers(requestPath string, code []byte) {
	specifiers, err := extractImportsFromScript(string(code))
	if err != nil {
		s.logger.Debug("Failed to extract imports from %s: %v", requestPath, err)
		return
	}

	if len(specifiers) == 0 {
		return
	}

	s.mu.RLock()
	im := s.importMap
	graph := s.importMapGraph
	s.mu.RUnlock()

	for _, specifier := range specifiers {
		if !isBareSpecifier(specifier) {
			continue
		}

		if isResolvedInImportMap(specifier, im) {
			continue
		}

		pkgName := extractPackageName(specifier)
		if pkgName == "" {
			continue
		}

		// Deduplicate: only warn once per specifier
		if _, loaded := s.warnedSpecifiers.LoadOrStore(specifier, true); loaded {
			continue
		}

		// Check if this is a known transitive dependency
		var owners []string
		if graph != nil {
			owners = graph.Dependents(pkgName)
		}

		msg := formatTransitiveWarning(specifier, pkgName, owners)
		s.logger.Warning("%s", msg)
		_ = s.BroadcastError("Unresolved Import", msg, requestPath)
	}
}

// extractPackageName extracts the npm package name from an import specifier.
// "@scope/pkg/path" → "@scope/pkg", "pkg/path" → "pkg", "pkg" → "pkg"
func extractPackageName(specifier string) string {
	if strings.HasPrefix(specifier, "@") {
		// Scoped package: @scope/name or @scope/name/path
		parts := strings.SplitN(specifier, "/", 3)
		if len(parts) < 2 {
			return ""
		}
		return parts[0] + "/" + parts[1]
	}
	// Regular package: name or name/path
	parts := strings.SplitN(specifier, "/", 2)
	return parts[0]
}

// isBareSpecifier returns true if the specifier is a bare module specifier
// (not relative, absolute, or a URL).
func isBareSpecifier(specifier string) bool {
	return specifier != "" &&
		!strings.HasPrefix(specifier, ".") &&
		!strings.HasPrefix(specifier, "/") &&
		!strings.HasPrefix(specifier, "http://") &&
		!strings.HasPrefix(specifier, "https://")
}

// isResolvedInImportMap checks whether a bare specifier is resolved by the import map,
// either via exact match or longest prefix match.
func isResolvedInImportMap(specifier string, im *importmappkg.ImportMap) bool {
	if im == nil || len(im.Imports) == 0 {
		return false
	}

	// Exact match
	if _, ok := im.Imports[specifier]; ok {
		return true
	}

	// Prefix match: find entries ending with "/" that match the specifier prefix
	for key := range im.Imports {
		if strings.HasSuffix(key, "/") && strings.HasPrefix(specifier, key) {
			return true
		}
	}

	return false
}

// formatTransitiveWarning formats a user-facing warning message for an unresolved
// bare specifier that is a transitive dependency.
func formatTransitiveWarning(specifier, pkgName string, owners []string) string {
	var b strings.Builder
	fmt.Fprintf(&b, "Cannot resolve \"%s\"\n\n", specifier)

	if len(owners) > 0 {
		ownerList := strings.Join(owners, "\", \"")
		fmt.Fprintf(&b, "\"%s\" is not a direct dependency — it's a transitive dependency of \"%s\".\n\n", pkgName, ownerList)
		if len(owners) == 1 {
			fmt.Fprintf(&b, "Add \"%s\" to your package.json dependencies,\nor import from \"%s\" instead.", pkgName, owners[0])
		} else {
			fmt.Fprintf(&b, "Add \"%s\" to your package.json dependencies,\nor import from one of its dependents instead.", pkgName)
		}
	} else {
		fmt.Fprintf(&b, "\"%s\" is not in the import map and is not a known dependency.\n\n", pkgName)
		fmt.Fprintf(&b, "Add \"%s\" to your package.json dependencies.", pkgName)
	}

	return b.String()
}
