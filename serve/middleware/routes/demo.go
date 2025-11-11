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

package routes

import (
	"encoding/json"
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	M "bennypowers.dev/cem/manifest"
)

// DemoRouteEntry maps a local route to demo metadata
type DemoRouteEntry struct {
	LocalRoute  string    // e.g., "/elements/accordion/demo/"
	TagName     string    // e.g., "rh-accordion"
	Demo        *M.Demo   // Demo metadata from manifest
	FilePath    string    // Relative file path from watch dir
	PackageName string    // Package name (for workspace mode)
	PackagePath string    // Absolute path to package directory (for workspace mode)
}

// BuildDemoRoutingTable creates a routing table from manifest
func BuildDemoRoutingTable(manifestBytes []byte) (map[string]*DemoRouteEntry, error) {
	if len(manifestBytes) == 0 {
		return nil, fmt.Errorf("no manifest available")
	}

	var pkg M.Package
	if err := json.Unmarshal(manifestBytes, &pkg); err != nil {
		return nil, fmt.Errorf("parsing manifest: %w", err)
	}

	// TODO: Get source control root URL from config
	// For now, we'll extract it from the demo source.href if needed
	var sourceControlRootURL string

	routes := make(map[string]*DemoRouteEntry)

	for _, renderableDemo := range pkg.RenderableDemos() {
		demoURL := renderableDemo.Demo.URL

		// Extract local route from canonical URL (strip origin)
		localRoute := demoURL
		if parsed, err := url.Parse(demoURL); err == nil && parsed.Path != "" {
			localRoute = parsed.Path
		}

		// Normalize relative paths (./foo -> /foo)
		if strings.HasPrefix(localRoute, "./") {
			localRoute = localRoute[1:] // Remove leading dot
		}
		if !strings.HasPrefix(localRoute, "/") {
			localRoute = "/" + localRoute
		}

		// Normalize: ensure trailing slash for directory-style URLs
		if localRoute != "/" && localRoute[len(localRoute)-1] != '/' && filepath.Ext(localRoute) == "" {
			localRoute += "/"
		}

		// Determine file path from Source.Href
		filePath := demoURL
		if renderableDemo.Demo.Source != nil && renderableDemo.Demo.Source.Href != "" {
			sourceHref := renderableDemo.Demo.Source.Href

			// If we have a sourceControlRootURL and the source href starts with it,
			// everything after the root URL is the file path
			if sourceControlRootURL != "" && strings.HasPrefix(sourceHref, sourceControlRootURL) {
				filePath = strings.TrimPrefix(sourceHref, sourceControlRootURL)
			} else if parsed, err := url.Parse(sourceHref); err == nil && parsed.Path != "" {
				// Try to extract file path from GitHub/GitLab URLs
				// URLs like: https://github.com/org/repo/tree/main/path/to/file.html
				// or:        https://github.com/org/repo/blob/main/path/to/file.html
				extractedPath := ""
				path := parsed.Path
				for _, prefix := range []string{"/tree/", "/blob/"} {
					if idx := strings.Index(path, prefix); idx >= 0 {
						// Find next slash after the branch name
						branchStart := idx + len(prefix)
						if slashIdx := strings.Index(path[branchStart:], "/"); slashIdx >= 0 {
							extractedPath = path[branchStart+slashIdx+1:]
							break
						}
					}
				}

				if extractedPath != "" {
					filePath = extractedPath
				} else {
					// Fallback: extract the filename and combine with module directory
					moduleDir := filepath.Dir(renderableDemo.Module.Path)
					filename := filepath.Base(parsed.Path)
					filePath = filepath.Join(moduleDir, "demo", filename)
				}
			}
		} else if strings.HasPrefix(demoURL, "./") || strings.HasPrefix(demoURL, "/") {
			// Relative path - use as-is (strip leading ./ if present)
			filePath = strings.TrimPrefix(demoURL, "./")
			filePath = strings.TrimPrefix(filePath, "/")
		}

		// Normalize filePath to ensure it's always relative (security: prevent directory traversal)
		// Strip leading "./" and "/" repeatedly to ensure relative path
		for strings.HasPrefix(filePath, "./") || strings.HasPrefix(filePath, "/") {
			filePath = strings.TrimPrefix(filePath, "./")
			filePath = strings.TrimPrefix(filePath, "/")
		}

		// Check for duplicate routes before assignment
		if existing, exists := routes[localRoute]; exists {
			return nil, fmt.Errorf("duplicate demo route %q: %s (tagName: %s) conflicts with existing %s (tagName: %s)",
				localRoute,
				filePath,
				renderableDemo.CustomElementDeclaration.TagName,
				existing.FilePath,
				existing.TagName)
		}

		routes[localRoute] = &DemoRouteEntry{
			LocalRoute: localRoute,
			TagName:    renderableDemo.CustomElementDeclaration.TagName,
			Demo:       renderableDemo.Demo,
			FilePath:   filePath,
		}
	}

	return routes, nil
}


// BuildWorkspaceRoutingTable builds a combined routing table from all packages
// Returns error if route conflicts are detected
func BuildWorkspaceRoutingTable(packages []PackageContext) (map[string]*DemoRouteEntry, error) {
	routes := make(map[string]*DemoRouteEntry)
	conflicts := make(map[string][]routeConflict)

	for _, pkg := range packages {
		// Build routing table for this package
		pkgRoutes, err := buildPackageRoutingTable(pkg)
		if err != nil {
			// Skip packages with routing errors
			continue
		}

		// Merge routes, detecting conflicts
		for route, entry := range pkgRoutes {
			if existing, exists := routes[route]; exists {
				// Conflict detected
				conflict := routeConflict{
					Route:       route,
					PackageName: pkg.Name,
					PackagePath: pkg.Path,
					FilePath:    entry.FilePath,
				}
				conflicts[route] = append(conflicts[route], conflict)
				// Also add the existing entry to conflicts if not already there
				if len(conflicts[route]) == 1 {
					existingConflict := routeConflict{
						Route:       route,
						PackageName: existing.PackageName,
						PackagePath: existing.PackagePath,
						FilePath:    existing.FilePath,
					}
					conflicts[route] = []routeConflict{existingConflict, conflict}
				}
			} else {
				routes[route] = entry
			}
		}
	}

	// If conflicts detected, return detailed error
	if len(conflicts) > 0 {
		return nil, formatRouteConflictsError(conflicts)
	}

	return routes, nil
}

// routeConflict represents a routing conflict between packages
type routeConflict struct {
	Route       string
	PackageName string
	PackagePath string
	FilePath    string
}

// formatRouteConflictsError creates a detailed error message for route conflicts
func formatRouteConflictsError(conflicts map[string][]routeConflict) error {
	var msg strings.Builder
	msg.WriteString("Route conflicts detected:\n\n")

	for route, conflictList := range conflicts {
		msg.WriteString(fmt.Sprintf("Conflict for route '%s':\n", route))
		for _, c := range conflictList {
			msg.WriteString(fmt.Sprintf("  - Package '%s' (%s)\n", c.PackageName, c.PackagePath))
			msg.WriteString(fmt.Sprintf("    Source: %s\n", c.FilePath))
		}
		msg.WriteString("\n")
	}

	msg.WriteString("Fix: Configure unique urlPattern in each package's manifest or demo metadata\n")

	return fmt.Errorf("%s", msg.String())
}

// buildPackageRoutingTable builds routing table for a single package
func buildPackageRoutingTable(pkg PackageContext) (map[string]*DemoRouteEntry, error) {
	var manifest M.Package
	if err := json.Unmarshal(pkg.Manifest, &manifest); err != nil {
		return nil, fmt.Errorf("parsing manifest for %s: %w", pkg.Name, err)
	}

	routes := make(map[string]*DemoRouteEntry)

	for _, renderableDemo := range manifest.RenderableDemos() {
		demoURL := renderableDemo.Demo.URL

		// Extract local route from canonical URL
		localRoute := demoURL
		if parsed, err := url.Parse(demoURL); err == nil && parsed.Path != "" {
			localRoute = parsed.Path
		}

		// Normalize relative paths
		if strings.HasPrefix(localRoute, "./") {
			localRoute = localRoute[1:]
		}
		if !strings.HasPrefix(localRoute, "/") {
			localRoute = "/" + localRoute
		}

		// Ensure trailing slash for directory-style URLs
		if !strings.HasSuffix(localRoute, "/") && !strings.Contains(filepath.Base(localRoute), ".") {
			localRoute += "/"
		}

		// Resolve file path
		var filePath string
		if renderableDemo.Demo.Source != nil && renderableDemo.Demo.Source.Href != "" {
			sourceHref := renderableDemo.Demo.Source.Href
			// Try to extract file path from GitHub/GitLab URLs
			// URLs like: https://github.com/org/repo/tree/main/path/to/file.html
			if parsed, err := url.Parse(sourceHref); err == nil && parsed.Path != "" {
				extractedPath := ""
				path := parsed.Path
				for _, prefix := range []string{"/tree/", "/blob/"} {
					if idx := strings.Index(path, prefix); idx >= 0 {
						// Find next slash after the branch name
						branchStart := idx + len(prefix)
						if slashIdx := strings.Index(path[branchStart:], "/"); slashIdx >= 0 {
							extractedPath = path[branchStart+slashIdx+1:]
							break
						}
					}
				}

				if extractedPath != "" {
					filePath = extractedPath
				} else {
					// Fallback: extract filename and combine with module directory
					moduleDir := filepath.Dir(renderableDemo.Module.Path)
					filename := filepath.Base(parsed.Path)
					filePath = filepath.Join(moduleDir, "demo", filename)
				}
			} else {
				filePath = sourceHref
			}
		} else {
			// Fallback: try to find demo file in module directory
			moduleDir := filepath.Dir(renderableDemo.Module.Path)
			filePath = filepath.Join(moduleDir, "demo", "index.html")
		}

		entry := &DemoRouteEntry{
			LocalRoute:  localRoute,
			TagName:     renderableDemo.CustomElementDeclaration.TagName,
			Demo:        renderableDemo.Demo,
			FilePath:    filePath,
			PackageName: pkg.Name,
			PackagePath: pkg.Path,
		}

		routes[localRoute] = entry
	}

	return routes, nil
}
