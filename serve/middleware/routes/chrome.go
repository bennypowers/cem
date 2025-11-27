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
)

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
}

// renderDemoChrome renders the demo chrome template with given data.
// If template execution fails, it broadcasts the error and returns a minimal page with error overlay.
func renderDemoChrome(data ChromeData) (string, error) {
	// Render markdown description if present
	if data.Description != "" {
		var buf bytes.Buffer
		err := md.Convert([]byte(data.Description), &buf)
		if err != nil {
			// Markdown conversion error - broadcast and continue with plain text
			if errorBroadcaster != nil {
				_ = errorBroadcaster.BroadcastError(
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
	err := DemoChromeTemplate.Execute(&buf, data)
	if err != nil {
		// Template execution failed - broadcast error and return minimal error page
		if errorBroadcaster != nil {
			_ = errorBroadcaster.BroadcastError(
				"Template Execution Error",
				"Failed to render page template: "+err.Error(),
				data.TagName,
			)
		}

		// Return minimal HTML with error message instead of failing completely
		// This allows the error overlay to display the error
		errorHTML := `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>Template Error</title>
  <link rel="stylesheet" href="/__cem/pf-tokens.css">
  <link rel="stylesheet" href="/__cem/demo-chrome.css">
  <script type="module">
    import '/__cem/elements/cem-transform-error-overlay/cem-transform-error-overlay.js';
  </script>
</head>
<body>
  <div style="padding: 2rem; max-width: 60rem; margin: 0 auto;">
    <h1>Template Rendering Error</h1>
    <p>The page template failed to render. Check the error overlay and server logs for details.</p>
    <pre style="background: #f5f5f5; padding: 1rem; border-radius: 4px; overflow-x: auto;">` +
			template.HTMLEscapeString(err.Error()) +
			`</pre>
  </div>
  <cem-transform-error-overlay id="error-overlay"></cem-transform-error-overlay>
</body>
</html>`
		return errorHTML, nil
	}

	return buf.String(), nil
}
