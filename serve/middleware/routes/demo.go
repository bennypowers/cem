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
	"bennypowers.dev/cem/serve/middleware/types"
)

// DemoRouteEntry is a type alias for types.DemoRouteEntry
type DemoRouteEntry = types.DemoRouteEntry

// BuildDemoRoutingTable creates a routing table from manifest
func BuildDemoRoutingTable(manifestBytes []byte, sourceControlRootURL string) (map[string]*DemoRouteEntry, error) {
	if len(manifestBytes) == 0 {
		return nil, fmt.Errorf("no manifest available")
	}

	var pkg M.Package
	if err := json.Unmarshal(manifestBytes, &pkg); err != nil {
		return nil, fmt.Errorf("parsing manifest: %w", err)
	}

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
		if localRoute != "/" && !strings.HasSuffix(localRoute, "/") && filepath.Ext(localRoute) == "" {
			localRoute += "/"
		}

		// Determine file path from Source.Href
		filePath := demoURL
		if renderableDemo.Demo.Source != nil && renderableDemo.Demo.Source.Href != "" {
			moduleDir := filepath.Dir(renderableDemo.Module.Path)
			filePath = resolveSourceHrefToFilePath(renderableDemo.Demo.Source.Href, moduleDir, sourceControlRootURL)
		} else if strings.HasPrefix(demoURL, "./") || strings.HasPrefix(demoURL, "/") {
			// Relative path - use as-is (strip leading ./ if present)
			filePath = strings.TrimPrefix(demoURL, "./")
			filePath = strings.TrimPrefix(filePath, "/")
		}

		// Normalize and validate filePath (security: prevent directory traversal)
		var err error
		filePath, err = normalizeAndValidateDemoPath(filePath, renderableDemo.CustomElementDeclaration.TagName, demoURL)
		if err != nil {
			return nil, err
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
			LocalRoute:  localRoute,
			TagName:     renderableDemo.CustomElementDeclaration.TagName,
			Demo:        renderableDemo.Demo,
			Declaration: renderableDemo.CustomElementDeclaration,
			FilePath:    filePath,
		}
	}

	return routes, nil
}

// resolveSourceHrefToFilePath extracts a file path from a source href.
// Handles URLs (GitHub/GitLab) and relative file paths.
// If sourceControlRootURL is provided and the href starts with it, returns the path relative to that root.
// For GitHub/GitLab URLs, extracts the file path after /tree/branch/ or /blob/branch/.
// For non-URL hrefs, treats them as file paths.
// Falls back to moduleDir/demo/filename for URLs without recognized patterns.
func resolveSourceHrefToFilePath(sourceHref, moduleDir, sourceControlRootURL string) string {
	// If we have a sourceControlRootURL and the source href starts with it,
	// everything after the root URL is the file path
	if sourceControlRootURL != "" && strings.HasPrefix(sourceHref, sourceControlRootURL) {
		return strings.TrimPrefix(sourceHref, sourceControlRootURL)
	}

	// Try to parse as URL
	parsed, err := url.Parse(sourceHref)
	if err == nil && parsed.Scheme != "" && parsed.Path != "" {
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
			return extractedPath
		}

		// Fallback: extract the filename and combine with module directory
		filename := filepath.Base(parsed.Path)
		return filepath.Join(moduleDir, "demo", filename)
	}

	// Not a URL, treat as file path
	return strings.TrimPrefix(sourceHref, "/")
}

// normalizeAndValidateDemoPath normalizes and validates a demo file path to prevent directory traversal.
// It cleans the path, ensures it's relative, and rejects paths that attempt to traverse above the root.
// Returns an error if the path is invalid or attempts directory traversal.
func normalizeAndValidateDemoPath(filePath, tagName, demoURL string) (string, error) {
	// Clean the path to resolve any ".." or "." segments
	filePath = filepath.Clean(filePath)

	// Strip leading slash if present (we want relative paths)
	filePath = strings.TrimPrefix(filePath, "/")

	// Reject paths that attempt to traverse above the root
	if strings.HasPrefix(filePath, "..") || filepath.IsAbs(filePath) {
		return "", fmt.Errorf("security: demo file path %q attempts directory traversal (tagName: %s, demo: %s)",
			filePath,
			tagName,
			demoURL)
	}

	return filePath, nil
}

// BuildWorkspaceRoutingTable builds a combined routing table from all packages
// Returns error if route conflicts are detected or if package routing errors occurred
func BuildWorkspaceRoutingTable(packages []PackageContext) (map[string]*DemoRouteEntry, error) {
	routes := make(map[string]*DemoRouteEntry)
	conflicts := make(map[string][]routeConflict)
	var packageErrors []packageRoutingError

	for _, pkg := range packages {
		// Build routing table for this package
		pkgRoutes, err := buildPackageRoutingTable(pkg)
		if err != nil {
			// Collect package routing errors for reporting
			packageErrors = append(packageErrors, packageRoutingError{
				PackageName: pkg.Name,
				PackagePath: pkg.Path,
				Error:       err,
			})
			// Skip packages with routing errors but continue processing others
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

	// Report package routing errors (non-fatal warnings)
	if len(packageErrors) > 0 {
		// If there are also conflicts, combine both errors
		if len(conflicts) > 0 {
			pkgErr := formatPackageRoutingErrors(packageErrors)
			conflictErr := formatRouteConflictsError(conflicts)
			return nil, fmt.Errorf("%w\n\n%w", pkgErr, conflictErr)
		}
		// Only package errors - return as error so callers are aware
		return routes, formatPackageRoutingErrors(packageErrors)
	}

	// If conflicts detected, return detailed error
	if len(conflicts) > 0 {
		return nil, formatRouteConflictsError(conflicts)
	}

	return routes, nil
}

// packageRoutingError represents a package that failed to build routing table
type packageRoutingError struct {
	PackageName string
	PackagePath string
	Error       error
}

// routeConflict represents a routing conflict between packages
type routeConflict struct {
	Route       string
	PackageName string
	PackagePath string
	FilePath    string
}

// formatPackageRoutingErrors creates a detailed error message for package routing failures
func formatPackageRoutingErrors(errors []packageRoutingError) error {
	var msg strings.Builder
	msg.WriteString("Package routing errors detected (packages skipped):\n\n")

	for _, pkgErr := range errors {
		msg.WriteString(fmt.Sprintf("Package '%s' (%s):\n", pkgErr.PackageName, pkgErr.PackagePath))
		msg.WriteString(fmt.Sprintf("  Error: %v\n\n", pkgErr.Error))
	}

	return fmt.Errorf("%s", msg.String())
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
		if localRoute != "/" && !strings.HasSuffix(localRoute, "/") && filepath.Ext(localRoute) == "" {
			localRoute += "/"
		}

		// Resolve file path
		moduleDir := filepath.Dir(renderableDemo.Module.Path)
		var filePath string
		if renderableDemo.Demo.Source != nil && renderableDemo.Demo.Source.Href != "" {
			filePath = resolveSourceHrefToFilePath(renderableDemo.Demo.Source.Href, moduleDir, "")
		} else {
			// Fallback: try to find demo file in module directory
			filePath = filepath.Join(moduleDir, "demo", "index.html")
		}

		// Normalize and validate filePath (security: prevent directory traversal)
		var err error
		filePath, err = normalizeAndValidateDemoPath(filePath, renderableDemo.CustomElementDeclaration.TagName, demoURL)
		if err != nil {
			return nil, err
		}

		// Make filePath relative to package path for workspace mode
		// This prevents path duplication when joining with PackagePath later
		// Note: filePath is guaranteed to be relative after normalizeAndValidateDemoPath
		relFilePath := filePath

		// Strip workspace-relative prefix from filePath if present
		// Example: pkg.Path="/home/user/cem/examples/intermediate" (absolute)
		//          filePath="examples/intermediate/elements/ui-button/demo/sizes.html" (relative)
		//          We need to strip "examples/intermediate/" to get "elements/ui-button/demo/sizes.html"
		// We can't compare pkg.Path directly with filePath since one is absolute and one is relative,
		// so we try increasingly longer suffixes of pkg.Path to find the workspace-relative portion.
		normalizedFilePath := filepath.ToSlash(filePath)
		pathParts := strings.Split(filepath.ToSlash(pkg.Path), "/")
		for i := len(pathParts) - 1; i >= 0; i-- {
			possiblePrefix := strings.Join(pathParts[i:], "/")
			if strings.HasPrefix(normalizedFilePath, possiblePrefix+"/") {
				relFilePath = strings.TrimPrefix(normalizedFilePath, possiblePrefix+"/")
				break
			}
		}

		// Check for duplicate routes before assignment
		if existing, exists := routes[localRoute]; exists {
			return nil, fmt.Errorf("duplicate demo route %q in package %s (%s): %s (tagName: %s) conflicts with existing %s (tagName: %s)",
				localRoute,
				pkg.Name,
				pkg.Path,
				relFilePath,
				renderableDemo.CustomElementDeclaration.TagName,
				existing.FilePath,
				existing.TagName)
		}

		entry := &DemoRouteEntry{
			LocalRoute:  localRoute,
			TagName:     renderableDemo.CustomElementDeclaration.TagName,
			Demo:        renderableDemo.Demo,
			Declaration: renderableDemo.CustomElementDeclaration,
			FilePath:    relFilePath,
			PackageName: pkg.Name,
			PackagePath: pkg.Path,
		}

		routes[localRoute] = entry
	}

	return routes, nil
}
