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

package importmap

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"bennypowers.dev/cem/serve/middleware"
)

// MiddlewareConfig holds configuration for the import map injection middleware
type MiddlewareConfig struct {
	// Context provides access to dev server state
	Context middleware.DevServerContext
}

// New creates a middleware that injects import maps into HTML responses
func New(config MiddlewareConfig) middleware.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Skip internal routes that need direct ResponseWriter access (WebSocket, etc.)
			// or are known non-HTML responses
			if strings.HasPrefix(r.URL.Path, "/__cem") ||
				r.URL.Path == "/custom-elements.json" {
				next.ServeHTTP(w, r)
				return
			}

			// Capture the response
			rec := middleware.NewResponseRecorder()
			next.ServeHTTP(rec, r)

			// Get body content
			bodyBytes := rec.Body()

			// Check if this is an HTML response
			contentType := rec.Header().Get("Content-Type")
			isHTML := middleware.IsHTMLResponse(contentType)
			if !isHTML && len(bodyBytes) > 0 {
				// Content sniffing fallback
				sample := bodyBytes
				if len(bodyBytes) > 100 {
					sample = bodyBytes[:100]
				}
				isHTML = strings.Contains(string(sample), "<html")
			}

			if !isHTML {
				// Not HTML, write original response
				writeResponse(w, rec)
				return
			}

			// Get import map from context (pre-computed during server initialization)
			importMapIface := config.Context.ImportMap()

			// Skip if no import map
			if importMapIface == nil {
				writeResponse(w, rec)
				return
			}

			// Convert to concrete type to marshal
			var importMapObj any = importMapIface

			// Convert import map to JSON
			importMapJSON, err := json.MarshalIndent(importMapObj, "  ", "  ")
			if err != nil {
				config.Context.Logger().Error("Failed to marshal import map: %v", err)
				// Continue without import map
				writeResponse(w, rec)
				return
			}

			// Inject import map into HTML
			htmlStr := string(bodyBytes)
			injected := injectImportMap(htmlStr, string(importMapJSON))

			// Copy headers (except Content-Length, it will be wrong)
			middleware.CopyHeaders(w.Header(), rec.Header(), "Content-Length")

			// Set content type if not set
			if contentType == "" {
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
			}

			// Write status code
			w.WriteHeader(rec.StatusCode())

			// Write modified body
			if _, err := w.Write([]byte(injected)); err != nil {
				config.Context.Logger().Error("Failed to write import map injection response: %v", err)
			}
		})
	}
}

// injectImportMap injects or replaces an import map script in HTML
// If an import map already exists, it replaces it; otherwise it injects a new one
func injectImportMap(htmlStr string, importMapJSON string) string {
	// Regex pattern to match <script type="importmap">...</script>
	// Handles both single and double quotes, and content with newlines
	importMapPattern := regexp.MustCompile(`(?s)<script\s+type=["']importmap["'][^>]*>.*?</script>`)

	// Check if an import map already exists
	if importMapPattern.MatchString(htmlStr) {
		// Replace existing import map with new content
		replacement := `<script type="importmap">
` + importMapJSON + `
  </script>`
		return importMapPattern.ReplaceAllString(htmlStr, replacement)
	}

	// No existing import map found, inject a new one before </head>
	return injectImportMapFallback(htmlStr, importMapJSON)
}

// injectImportMapFallback uses string replacement to inject import map
// Used when HTML parsing fails or when injecting (not replacing)
func injectImportMapFallback(html string, importMapJSON string) string {
	importMapScript := "<script type=\"importmap\">\n  " + importMapJSON + "\n  </script>"

	// Try to inject before </head>
	if idx := strings.Index(html, "</head>"); idx != -1 {
		return html[:idx] + importMapScript + "\n" + html[idx:]
	}

	// Fall back to start of <body>
	if idx := strings.Index(html, "<body"); idx != -1 {
		// Find the end of the <body> tag
		closeIdx := strings.Index(html[idx:], ">")
		if closeIdx != -1 {
			insertPos := idx + closeIdx + 1
			return html[:insertPos] + "\n" + importMapScript + html[insertPos:]
		}
	}

	// Last resort: inject at start of HTML
	if idx := strings.Index(html, "<html"); idx != -1 {
		closeIdx := strings.Index(html[idx:], ">")
		if closeIdx != -1 {
			insertPos := idx + closeIdx + 1
			return html[:insertPos] + "\n" + importMapScript + html[insertPos:]
		}
	}

	// Can't find a good place, return unchanged
	return html
}

// writeResponse writes the recorded response to the ResponseWriter
func writeResponse(w http.ResponseWriter, rec *middleware.ResponseRecorder) {
	// Copy all headers
	for key, values := range rec.Header() {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// Write status code
	w.WriteHeader(rec.StatusCode())

	// Write body
	_, _ = w.Write(rec.Body())
}
