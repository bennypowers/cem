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
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"bennypowers.dev/cem/serve/middleware"
	"golang.org/x/net/html"
)

// ImportMap returns the cached import map (may be nil)
func (s *Server) ImportMap() middleware.ImportMap {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.importMap
}

// resolveSourceFile checks if a .js file has a corresponding .ts source
func (s *Server) resolveSourceFile(path string) string {
	// If it's a .js file, check if .ts exists
	if filepath.Ext(path) == ".js" {
		tsPath := path[:len(path)-3] + ".ts"
		watchDir := s.WatchDir()
		if watchDir != "" {
			fullTsPath := filepath.Join(watchDir, tsPath)
			if fs := s.FileSystem(); fs != nil {
				if _, err := fs.Stat(fullTsPath); err == nil {
					return tsPath
				}
			} else if _, err := os.Stat(fullTsPath); err == nil {
				return tsPath
			}
		}
	}
	return path
}

// extractModuleImports parses an HTML file and extracts ES module import specifiers
// Returns both import specifiers from inline scripts and src URLs from script tags
func (s *Server) extractModuleImports(htmlPath string) ([]string, error) {
	s.mu.RLock()
	fs := s.fs
	s.mu.RUnlock()

	content, err := fs.ReadFile(htmlPath)
	if err != nil {
		return nil, err
	}

	imports := make(map[string]bool) // Use map to deduplicate
	doc, err := html.Parse(strings.NewReader(string(content)))
	if err != nil {
		return nil, err
	}

	// Import regex patterns
	importFromRe := regexp.MustCompile(`import\s+(?:[^'"]*?)\s*from\s*['"]([^'"]+)['"]`)
	importRe := regexp.MustCompile(`import\s*['"]([^'"]+)['"]`)

	var visit func(*html.Node)
	visit = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "script" {
			// Check if it's a module script
			isModule := false
			var srcAttr string
			for _, attr := range n.Attr {
				if attr.Key == "type" && attr.Val == "module" {
					isModule = true
				}
				if attr.Key == "src" {
					srcAttr = attr.Val
				}
			}

			if isModule {
				// External module script
				if srcAttr != "" {
					imports[srcAttr] = true
				}

				// Inline module script - extract import statements
				if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
					scriptContent := n.FirstChild.Data

					// Extract "import ... from 'specifier'" statements
					matches := importFromRe.FindAllStringSubmatch(scriptContent, -1)
					for _, match := range matches {
						if len(match) > 1 {
							imports[match[1]] = true
						}
					}

					// Extract "import 'specifier'" statements
					matches = importRe.FindAllStringSubmatch(scriptContent, -1)
					for _, match := range matches {
						if len(match) > 1 && !strings.Contains(match[0], " from ") {
							imports[match[1]] = true
						}
					}
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visit(c)
		}
	}

	visit(doc)

	// Convert map to slice
	result := make([]string, 0, len(imports))
	for imp := range imports {
		result = append(result, imp)
	}

	return result, nil
}

// resolveImportToPath converts an import specifier to a file path
// Handles bare specifiers (via import map), relative paths, and absolute paths
// contextDir is the directory of the HTML file making the import (for resolving relative imports)
func (s *Server) resolveImportToPath(importSpec string, contextDir string) []string {
	paths := make([]string, 0, 2)

	// If it's a relative path, resolve it relative to the context directory
	if strings.HasPrefix(importSpec, "./") || strings.HasPrefix(importSpec, "../") {
		var resolved string
		if contextDir != "" {
			// Resolve relative to the HTML file's directory
			resolved = filepath.Join(contextDir, importSpec)
		} else {
			// No context available, use import spec as-is
			resolved = importSpec
		}
		normalized := filepath.Clean(resolved)
		paths = append(paths, normalized)
		s.logger.Debug("Resolved relative import %s (context: %s) -> %s", importSpec, contextDir, normalized)
		return paths
	}

	// If it's an absolute path, use it directly
	if strings.HasPrefix(importSpec, "/") {
		normalized := filepath.Clean(importSpec)
		paths = append(paths, normalized)
		s.logger.Debug("Resolved absolute import %s -> %s", importSpec, normalized)
		return paths
	}

	// Try to resolve via import map
	s.mu.RLock()
	importMap := s.importMap
	s.mu.RUnlock()

	if importMap != nil {
		// First try exact match
		if resolved, ok := importMap.Imports[importSpec]; ok {
			s.logger.Debug("Import map entry: %s -> %s", importSpec, resolved)
			if parsedURL, err := url.Parse(resolved); err == nil && parsedURL.Path != "" {
				paths = append(paths, parsedURL.Path)
				s.logger.Debug("Resolved bare import %s -> %s (via import map exact match)", importSpec, parsedURL.Path)
			} else {
				s.logger.Debug("Failed to parse URL or extract path from: %s", resolved)
			}
		} else {
			// Try prefix matching for entries ending with "/"
			// Find the longest matching prefix
			var longestPrefix string
			var longestPrefixValue string
			for key, value := range importMap.Imports {
				if strings.HasSuffix(key, "/") && strings.HasPrefix(importSpec, key) {
					if len(key) > len(longestPrefix) {
						longestPrefix = key
						longestPrefixValue = value
					}
				}
			}

			if longestPrefix != "" {
				// Replace the prefix
				suffix := strings.TrimPrefix(importSpec, longestPrefix)
				resolved := longestPrefixValue + suffix
				s.logger.Debug("Import map prefix match: %s (prefix: %s -> %s)", importSpec, longestPrefix, longestPrefixValue)
				if parsedURL, err := url.Parse(resolved); err == nil && parsedURL.Path != "" {
					paths = append(paths, parsedURL.Path)
					s.logger.Debug("Resolved bare import %s -> %s (via import map prefix match)", importSpec, parsedURL.Path)
				} else {
					s.logger.Debug("Failed to parse URL or extract path from: %s", resolved)
				}
			} else {
				s.logger.Debug("Import %s not found in import map (have %d entries)", importSpec, len(importMap.Imports))
			}
		}

		// TODO(enhancement): Add support for scoped imports for path-specific resolution.
		// The import map spec supports a 'scopes' section for context-dependent module resolution.
		// This would allow different import resolutions based on the importing file's path.
		// Currently, we only resolve from the global 'imports' section, which handles most cases.
		// Ref: https://github.com/WICG/import-maps#scoping-examples
	} else {
		s.logger.Debug("No import map available to resolve %s", importSpec)
	}

	return paths
}
