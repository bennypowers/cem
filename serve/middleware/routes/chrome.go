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
	"html/template"

	M "bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/serve/middleware"
)

// PackageWithManifest represents a package with its modules for workspace tree organization
type PackageWithManifest struct {
	Name    string
	Modules []M.Module
}

// ChromeData represents template data for demo chrome
type ChromeData struct {
	TagName        string
	DemoTitle      string
	DemoHTML       template.HTML
	EnabledKnobs   string
	KnobsHTML      template.HTML // Rendered knobs controls HTML
	ImportMap      template.HTML // Use HTML instead of JS for importmap script content
	Description    template.HTML
	ShadowMode     bool
	SourceURL      string        // Source file URL (e.g., GitHub blob URL)
	CanonicalURL   string        // Canonical demo URL
	PackageName    string        // Package name for title (listing pages)
	NavigationHTML template.HTML // Navigation drawer HTML (listing pages)
	ManifestJSON   template.JS   // Full manifest JSON for client-side tools
	Manifest       *M.Package    // Parsed manifest for server-side tree rendering
	Packages       []PackageWithManifest // Workspace packages with modules (for package-level tree)
	State          CemServeState // Persisted UI state for SSR (color scheme, drawer, tree)
}

// TemplateErrorData represents template data for the error page
type TemplateErrorData struct {
	ErrorMessage string
	Title        string
	Message      string
	File         string
}

// renderDemoChrome renders the demo chrome template with given data.
// If template execution fails, it broadcasts the error and returns a minimal page with error overlay.
func renderDemoChrome(templates *TemplateRegistry, ctx middleware.DevServerContext, data ChromeData) (string, error) {
	// Render markdown description if present
	if data.Description != "" {
		var buf bytes.Buffer
		err := md.Convert([]byte(data.Description), &buf)
		if err != nil {
			// Markdown conversion error - broadcast and continue with plain text
			if ctx != nil {
				_ = ctx.BroadcastError(
					"Markdown Conversion Error",
					"Failed to convert description markdown: "+err.Error(),
					data.TagName,
				)
			}
			// Keep original text as-is
		} else {
			data.Description = template.HTML(buf.String())
		}
	}

	var buf bytes.Buffer
	err := templates.DemoChromeTemplate.Execute(&buf, data)
	if err != nil {
		// Template execution failed - broadcast error and return minimal error page
		title := "Template Execution Error"
		message := "Failed to render page template: " + err.Error()
		if ctx != nil {
			_ = ctx.BroadcastError(
				title,
				message,
				data.TagName,
			)
		}

		// Render error page using template with SSR-rendered error overlay
		errorData := TemplateErrorData{
			ErrorMessage: err.Error(),
			Title:        title,
			Message:      message,
			File:         data.TagName,
		}

		var errorBuf bytes.Buffer
		if templateErr := templates.TemplateErrorTemplate.Execute(&errorBuf, errorData); templateErr != nil {
			// If even the error template fails, fall back to absolute minimal HTML
			return `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>Critical Template Error</title>
</head>
<body>
  <h1>Critical Template Error</h1>
  <p>Both the page template and error template failed to render.</p>
  <pre>Original error: ` + template.HTMLEscapeString(err.Error()) + `</pre>
  <pre>Error template error: ` + template.HTMLEscapeString(templateErr.Error()) + `</pre>
</body>
</html>`, nil
		}

		return errorBuf.String(), nil
	}

	return buf.String(), nil
}
