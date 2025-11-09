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
	TagName string
	Demos   []DemoListing
}

// DemoListing represents a single demo
type DemoListing struct {
	Name string
	URL  string
	Slug string
}

// renderElementListing renders the root listing page with all elements
func renderElementListing(manifestBytes []byte, importMap string) (string, error) {
	if len(manifestBytes) == 0 {
		// Empty manifest - show helpful message
		return renderDemoChrome(ChromeData{
			TagName:   "cem-serve",
			DemoTitle: "Component Browser",
			DemoHTML: template.HTML(`
				<div style="text-align: center; padding: var(--__cem-spacing-xl); color: var(--__cem-text-secondary);">
					<p>No custom elements found in manifest.</p>
					<p style="font-size: var(--__cem-font-size-sm);">Add some components to your project and they'll appear here.</p>
				</div>
			`),
			ImportMap: template.HTML(importMap),
		})
	}

	// Parse manifest
	var pkg M.Package
	if err := json.Unmarshal(manifestBytes, &pkg); err != nil {
		return "", fmt.Errorf("parsing manifest: %w", err)
	}

	// Extract all elements with demos
	elements := extractElementListings(&pkg)

	if len(elements) == 0 {
		return renderDemoChrome(ChromeData{
			TagName:   "cem-serve",
			DemoTitle: "Component Browser",
			DemoHTML: template.HTML(`
				<div style="text-align: center; padding: var(--__cem-spacing-xl); color: var(--__cem-text-secondary);">
					<p>No demos found.</p>
					<p style="font-size: var(--__cem-font-size-sm);">Add demo files to your components and they'll appear here.</p>
				</div>
			`),
			ImportMap: template.HTML(importMap),
		})
	}

	// Sort elements alphabetically
	sort.Slice(elements, func(i, j int) bool {
		return elements[i].TagName < elements[j].TagName
	})

	// Build listing HTML
	listingHTML := buildListingHTML(elements)

	return renderDemoChrome(ChromeData{
		TagName:   "cem-serve",
		DemoTitle: "Component Browser",
		DemoHTML:  template.HTML(listingHTML),
		ImportMap: template.HTML(importMap),
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
func extractElementListings(pkg *M.Package) []ElementListing {
	elements := []ElementListing{}
	elementMap := make(map[string]*ElementListing)

	// Use RenderableDemos to get all demos with proper context
	for _, renderableDemo := range pkg.RenderableDemos() {
		tagName := renderableDemo.CustomElementDeclaration.TagName
		demo := renderableDemo.Demo

		// Get or create element listing
		listing, exists := elementMap[tagName]
		if !exists {
			listing = &ElementListing{
				TagName: tagName,
				Demos:   make([]DemoListing, 0),
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
			Name: name,
			URL:  localRoute,
			Slug: slugify(name),
		})
	}

	// Convert map to slice
	for _, listing := range elementMap {
		elements = append(elements, *listing)
	}

	return elements
}

// buildListingHTML builds the HTML for the element listing
func buildListingHTML(elements []ElementListing) string {
	var buf bytes.Buffer

	buf.WriteString(`<cem-serve-listing>`)
	buf.WriteString(`<div style="display: grid; gap: var(--__cem-spacing-lg);">`)

	for _, element := range elements {
		buf.WriteString(`<section style="background: var(--__cem-bg-secondary); border: var(--__cem-border-width) solid var(--__cem-border-color); border-radius: var(--__cem-border-radius); padding: var(--__cem-spacing-lg);">`)

		// Element name
		buf.WriteString(fmt.Sprintf(`<h2 style="margin: 0 0 var(--__cem-spacing-md) 0; color: var(--__cem-accent-color); font-size: var(--__cem-font-size-lg);"><code>&lt;%s&gt;</code></h2>`, element.TagName))

		// Demo list
		buf.WriteString(`<ul style="list-style: none; padding: 0; margin: 0; display: grid; gap: var(--__cem-spacing-sm);">`)
		for _, demo := range element.Demos {
			buf.WriteString(fmt.Sprintf(`<li><a href="%s" style="color: var(--__cem-text-primary); text-decoration: none; padding: var(--__cem-spacing-sm) var(--__cem-spacing-md); background: var(--__cem-bg-tertiary); border-radius: var(--__cem-border-radius); display: block; transition: background 0.2s;" onmouseover="this.style.background='var(--__cem-accent-color)'; this.style.color='white';" onmouseout="this.style.background='var(--__cem-bg-tertiary)'; this.style.color='var(--__cem-text-primary)';">%s</a></li>`,
				demo.URL, demo.Name))
		}
		buf.WriteString(`</ul>`)

		buf.WriteString(`</section>`)
	}

	buf.WriteString(`</div>`)
	buf.WriteString(`</cem-serve-listing>`)

	return buf.String()
}

// slugify converts a string to a URL-safe slug
func slugify(s string) string {
	// Simple slugification - lowercase and replace spaces with hyphens
	// TODO: More robust slugification
	s = template.HTMLEscapeString(s)
	result := ""
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			result += string(r)
		} else if r == ' ' || r == '-' {
			result += "-"
		}
	}
	// Convert to lowercase
	result = template.HTMLEscapeString(result)
	return result
}
