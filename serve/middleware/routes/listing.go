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
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/url"
	"path/filepath"
	"sort"
	"strings"

	M "bennypowers.dev/cem/manifest"
)

// ElementListing represents a custom element with its demos
type ElementListing struct {
	TagName     string
	Summary     template.HTML
	Description template.HTML
	Demos       []DemoListing
}

// DemoListing represents a single demo
type DemoListing struct {
	Name        string
	URL         string
	Slug        string
	PackageName string // Package name (for workspace mode, empty for single-package)
}

// RenderElementListing renders the root listing page with all elements
func RenderElementListing(manifestBytes []byte, importMap string, packageName string) (string, error) {
	if len(manifestBytes) == 0 {
		// Empty manifest - show helpful message
		title := packageName
		if title == "" {
			title = "Component Browser"
		}
		return renderDemoChrome(ChromeData{
			TagName:   "", // Empty for index page
			DemoTitle: title,
			DemoHTML: template.HTML(`
				<div class="empty-state">
					<p>No custom elements found in manifest.</p>
					<p class="help-text">Add some components to your project and they'll appear here.</p>
				</div>
			`),
			ImportMap:   template.HTML(importMap),
			PackageName: title,
		})
	}

	// Parse manifest
	var pkg M.Package
	if err := json.Unmarshal(manifestBytes, &pkg); err != nil {
		return "", fmt.Errorf("parsing manifest: %w", err)
	}

	// Build navigation (this also extracts elements and sorts them)
	navigationHTML, title, err := BuildSinglePackageNavigation(manifestBytes, packageName)
	if err != nil {
		return "", fmt.Errorf("building navigation: %w", err)
	}

	// Extract elements for the main listing content
	elements, err := extractElementListings(&pkg, packageName)
	if err != nil {
		return "", fmt.Errorf("extracting element listings: %w", err)
	}

	if len(elements) == 0 {
		return renderDemoChrome(ChromeData{
			TagName:   "cem-serve",
			DemoTitle: title,
			DemoHTML: template.HTML(`
				<div class="empty-state">
					<p>No demos found.</p>
					<p class="help-text">Add demo files to your components and they'll appear here.</p>
				</div>
			`),
			ImportMap:   template.HTML(importMap),
			PackageName: title,
		})
	}

	// Sort elements alphabetically
	sort.Slice(elements, func(i, j int) bool {
		return elements[i].TagName < elements[j].TagName
	})

	// Wrap in a single package for consistent template structure
	packages := []PackageNavigation{
		{
			Name:     title,
			Elements: elements,
		},
	}

	// Render with template
	var buf bytes.Buffer
	err = WorkspaceListingTemplate.Execute(&buf, map[string]interface{}{
		"Packages": packages,
	})
	if err != nil {
		return "", fmt.Errorf("executing element listing template: %w", err)
	}

	return renderDemoChrome(ChromeData{
		TagName:        "", // Empty for index page
		DemoTitle:      title,
		DemoHTML:       template.HTML(buf.String()),
		ImportMap:      template.HTML(importMap),
		PackageName:    title,
		NavigationHTML: navigationHTML,
	})
}

// prettifyRoute converts a route path to a human-readable name
// e.g., /elements/accordion/demo/ → "Demo"
// e.g., /elements/accordion/demo/accents/ → "Accents"
// e.g., /demo/basic.html → "Basic"
func prettifyRoute(route string) string {
	// Remove trailing slash
	route = strings.TrimSuffix(route, "/")

	// Get the last path segment
	parts := strings.Split(route, "/")
	if len(parts) == 0 {
		return "Demo"
	}

	lastPart := parts[len(parts)-1]

	// Remove file extension if present
	if ext := filepath.Ext(lastPart); ext != "" {
		lastPart = strings.TrimSuffix(lastPart, ext)
	}

	// If it's just "demo" or empty, use "Demo"
	if lastPart == "" || lastPart == "demo" {
		return "Demo"
	}

	// Replace hyphens/underscores with spaces and title case
	lastPart = strings.ReplaceAll(lastPart, "-", " ")
	lastPart = strings.ReplaceAll(lastPart, "_", " ")

	// Simple title case: capitalize first letter
	if len(lastPart) > 0 {
		lastPart = strings.ToUpper(lastPart[:1]) + lastPart[1:]
	}

	return lastPart
}

// extractLocalRoute extracts the local route from a demo URL (same logic as routing table)
func extractLocalRoute(demoURL string) string {
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

	return localRoute
}

// extractElementListings extracts element listings from manifest using RenderableDemos
func extractElementListings(pkg *M.Package, packageName string) ([]ElementListing, error) {
	elements := []ElementListing{}
	elementMap := make(map[string]*ElementListing)

	// Use RenderableDemos to get all demos with proper context
	for _, renderableDemo := range pkg.RenderableDemos() {
		tagName := renderableDemo.CustomElementDeclaration.TagName
		demo := renderableDemo.Demo

		// Get or create element listing
		listing, exists := elementMap[tagName]
		if !exists {
			summaryHTML, err := markdownToHTML(renderableDemo.CustomElementDeclaration.Summary)
			if err != nil {
				return nil, fmt.Errorf("failed to convert summary to HTML for %s: %w", tagName, err)
			}
			descriptionHTML, err := markdownToHTML(renderableDemo.CustomElementDeclaration.Description)
			if err != nil {
				return nil, fmt.Errorf("failed to convert description to HTML for %s: %w", tagName, err)
			}
			listing = &ElementListing{
				TagName:     tagName,
				Summary:     template.HTML(summaryHTML),
				Description: template.HTML(descriptionHTML),
				Demos:       make([]DemoListing, 0),
			}
			elementMap[tagName] = listing
		}

		// Extract local route from demo URL
		localRoute := extractLocalRoute(demo.URL)

		// Use description as name if available, otherwise prettify the route
		name := demo.Description
		if name == "" {
			// Use the local route but make it prettier
			// e.g., /elements/accordion/demo/ → "Demo"
			// e.g., /elements/accordion/demo/accents/ → "Accents"
			name = prettifyRoute(localRoute)
		}

		listing.Demos = append(listing.Demos, DemoListing{
			Name:        name,
			URL:         localRoute,
			Slug:        slugify(name),
			PackageName: packageName,
		})
	}

	// Convert map to slice
	for _, listing := range elementMap {
		elements = append(elements, *listing)
	}

	return elements, nil
}

// slugify converts a string to a URL-safe slug
func slugify(s string) string {
	// Simple slugification - lowercase and replace spaces with hyphens
	// TODO: More robust slugification
	result := ""
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			result += string(r)
		} else if r == ' ' || r == '-' {
			result += "-"
		}
	}
	// Convert to lowercase
	return strings.ToLower(result)
}

// PackageNavigation represents a package with its elements for navigation
type PackageNavigation struct {
	Name     string
	Elements []ElementListing
}

// BuildSinglePackageNavigation builds navigation HTML for single-package mode
func BuildSinglePackageNavigation(manifestBytes []byte, packageName string) (template.HTML, string, error) {
	if len(manifestBytes) == 0 {
		return "", packageName, nil
	}

	var pkg M.Package
	if err := json.Unmarshal(manifestBytes, &pkg); err != nil {
		return "", packageName, fmt.Errorf("parsing manifest: %w", err)
	}

	elements, err := extractElementListings(&pkg, packageName)
	if err != nil {
		return "", packageName, fmt.Errorf("extracting element listings: %w", err)
	}
	if len(elements) == 0 {
		return "", packageName, nil
	}

	// Sort elements alphabetically
	sort.Slice(elements, func(i, j int) bool {
		return elements[i].TagName < elements[j].TagName
	})

	// Use package name if provided, otherwise default
	title := packageName
	if title == "" {
		title = "Component Browser"
	}

	// Create single package navigation
	packages := []PackageNavigation{
		{
			Name:     title,
			Elements: elements,
		},
	}

	return renderNavigationHTML(packages), title, nil
}

// buildPackageListingsFromRoutes groups routes by package and element, extracting
// summaries and descriptions to build structured package navigation data.
// Returns a sorted slice of PackageNavigation with alphabetically ordered packages and elements.
func buildPackageListingsFromRoutes(routes map[string]*DemoRouteEntry) ([]PackageNavigation, error) {
	// Group routes by package, then by element
	packageElements := make(map[string]map[string][]*DemoRouteEntry)
	for _, route := range routes {
		if packageElements[route.PackageName] == nil {
			packageElements[route.PackageName] = make(map[string][]*DemoRouteEntry)
		}
		packageElements[route.PackageName][route.TagName] = append(packageElements[route.PackageName][route.TagName], route)
	}

	// Build package navigation list
	var packageNav []PackageNavigation
	for pkgName, elements := range packageElements {
		var elementListings []ElementListing
		for tagName, tagRoutes := range elements {
			// Sort demos by URL
			sort.Slice(tagRoutes, func(i, j int) bool {
				return tagRoutes[i].LocalRoute < tagRoutes[j].LocalRoute
			})

			var demoListings []DemoListing
			var summary, description template.HTML
			for _, route := range tagRoutes {
				// Get summary and description from first route's declaration
				if route.Declaration != nil && summary == "" {
					summaryHTML, err := markdownToHTML(route.Declaration.Summary)
					if err != nil {
						return nil, fmt.Errorf("failed to convert summary to HTML for %s: %w", tagName, err)
					}
					descriptionHTML, err := markdownToHTML(route.Declaration.Description)
					if err != nil {
						return nil, fmt.Errorf("failed to convert description to HTML for %s: %w", tagName, err)
					}
					summary = template.HTML(summaryHTML)
					description = template.HTML(descriptionHTML)
				}

				demoName := route.Demo.Description
				if demoName == "" {
					demoName = prettifyRoute(route.LocalRoute)
				}
				demoListings = append(demoListings, DemoListing{
					Name:        demoName,
					URL:         route.LocalRoute,
					Slug:        slugify(demoName),
					PackageName: route.PackageName,
				})
			}

			elementListings = append(elementListings, ElementListing{
				TagName:     tagName,
				Summary:     summary,
				Description: description,
				Demos:       demoListings,
			})
		}

		// Sort elements alphabetically
		sort.Slice(elementListings, func(i, j int) bool {
			return elementListings[i].TagName < elementListings[j].TagName
		})

		packageNav = append(packageNav, PackageNavigation{
			Name:     pkgName,
			Elements: elementListings,
		})
	}

	// Sort packages alphabetically
	sort.Slice(packageNav, func(i, j int) bool {
		return packageNav[i].Name < packageNav[j].Name
	})

	return packageNav, nil
}

// BuildWorkspaceNavigation builds navigation HTML for workspace mode
func BuildWorkspaceNavigation(packages []PackageContext) (template.HTML, error) {
	if len(packages) == 0 {
		return "", nil
	}

	// Build routing table
	routes, err := BuildWorkspaceRoutingTable(packages)
	if err != nil {
		return "", err
	}

	// Build package listings from routes
	packageNav, err := buildPackageListingsFromRoutes(routes)
	if err != nil {
		return "", err
	}

	return renderNavigationHTML(packageNav), nil
}

// renderNavigationHTML generates navigation drawer HTML from package navigation data
func renderNavigationHTML(packages []PackageNavigation) template.HTML {
	// Return empty if no packages
	if len(packages) == 0 {
		return template.HTML("")
	}

	var buf bytes.Buffer

	data := map[string]interface{}{
		"Packages":      packages,
		"SinglePackage": len(packages) == 1,
	}

	if err := NavigationTemplate.Execute(&buf, data); err != nil {
		// Fail gracefully on template error
		return template.HTML("")
	}

	return template.HTML(buf.String())
}

// RenderWorkspaceListing renders the workspace index page with all packages
func RenderWorkspaceListing(packages []PackageContext, importMap string) (string, error) {
	if len(packages) == 0 {
		return renderDemoChrome(ChromeData{
			TagName:     "cem-serve",
			DemoTitle:   "Workspace Browser",
			PackageName: "Workspace",
			DemoHTML: template.HTML(`
				<div class="empty-state">
					<p>No packages found in workspace.</p>
				</div>
			`),
			ImportMap: template.HTML(importMap),
		})
	}

	// Build navigation (this also builds element listings)
	navigationHTML, err := BuildWorkspaceNavigation(packages)
	if err != nil {
		return "", fmt.Errorf("building workspace navigation: %w", err)
	}

	// Build routing table for element listings in main content
	routes, err := BuildWorkspaceRoutingTable(packages)
	if err != nil {
		return "", fmt.Errorf("building workspace routing table: %w", err)
	}

	// Build package listings from routes using shared helper
	packageListings, err := buildPackageListingsFromRoutes(routes)
	if err != nil {
		return "", err
	}

	// Render with template
	var buf bytes.Buffer
	err = WorkspaceListingTemplate.Execute(&buf, map[string]interface{}{
		"Packages": packageListings,
	})
	if err != nil {
		return "", fmt.Errorf("executing workspace listing template: %w", err)
	}

	return renderDemoChrome(ChromeData{
		TagName:        "cem-serve",
		DemoTitle:      "Workspace Browser",
		DemoHTML:       template.HTML(buf.String()),
		ImportMap:      template.HTML(importMap),
		PackageName:    "Workspace",
		NavigationHTML: navigationHTML,
	})
}
