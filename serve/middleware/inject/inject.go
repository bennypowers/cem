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

package inject

import (
	"net/http"
	"strings"

	"bennypowers.dev/cem/serve/middleware"
)

// New creates a middleware that injects a WebSocket client script into HTML responses.
// The script is injected into the <head> element if present, otherwise at the start of <body>.
func New(enabled bool, scriptPath string) middleware.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !enabled {
				next.ServeHTTP(w, r)
				return
			}

			// Only inject into HTML responses
			if !strings.HasSuffix(r.URL.Path, ".html") && r.URL.Path != "/" {
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

			// Inject script before </head> or at start of <body>
			htmlStr := string(bodyBytes)
			script := "<script type=\"module\" src=\"" + scriptPath + "\"></script>"
			injected := InjectScript(htmlStr, script)

			// Copy headers (except Content-Length, it will be wrong)
			middleware.CopyHeaders(w.Header(), rec.Header(), "Content-Length")

			// Set content type if not set
			if contentType == "" {
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
			}

			w.WriteHeader(rec.StatusCode())
			if _, err := w.Write([]byte(injected)); err != nil {
				// Client disconnected or write error - can't respond
				return
			}
		})
	}
}

// writeResponse writes the captured response to the http.ResponseWriter
func writeResponse(w http.ResponseWriter, rec *middleware.ResponseRecorder) {
	middleware.CopyHeaders(w.Header(), rec.Header())
	w.WriteHeader(rec.StatusCode())
	if _, err := w.Write(rec.Body()); err != nil {
		// Client disconnected or write error - can't respond
		return
	}
}
