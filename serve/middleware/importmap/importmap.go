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
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	"bennypowers.dev/cem/serve/middleware"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
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

// injectImportMap injects or replaces an import map in HTML using DOM parsing.
// The import map is placed before the first <script> in <head>, or as the last
// child of <head> if no scripts exist. Any existing import map is removed first.
func injectImportMap(htmlStr string, importMapJSON string) string {
	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return injectImportMapFallback(htmlStr, importMapJSON)
	}

	head := findElement(doc, "head")
	if head == nil {
		return injectImportMapFallback(htmlStr, importMapJSON)
	}

	removeExistingImportMaps(doc)

	scriptTag := `<script type="importmap">` + "\n" + importMapJSON + "\n</script>"
	newNodes, err := html.ParseFragment(strings.NewReader(scriptTag), &html.Node{
		Type:     html.ElementNode,
		Data:     "head",
		DataAtom: atom.Head,
	})
	if err != nil || len(newNodes) == 0 {
		return injectImportMapFallback(htmlStr, importMapJSON)
	}

	firstScript := findFirstScript(head)
	for _, node := range newNodes {
		if firstScript != nil {
			head.InsertBefore(node, firstScript)
		} else {
			head.AppendChild(node)
		}
	}

	var buf bytes.Buffer
	if err := html.Render(&buf, doc); err != nil {
		return injectImportMapFallback(htmlStr, importMapJSON)
	}
	return buf.String()
}

func findElement(n *html.Node, tag string) *html.Node {
	if n.Type == html.ElementNode && n.Data == tag {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if result := findElement(c, tag); result != nil {
			return result
		}
	}
	return nil
}

func isImportMapScript(n *html.Node) bool {
	if n.Type != html.ElementNode || n.Data != "script" {
		return false
	}
	for _, attr := range n.Attr {
		if attr.Key == "type" && attr.Val == "importmap" {
			return true
		}
	}
	return false
}

func removeExistingImportMaps(root *html.Node) {
	var toRemove []*html.Node
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if isImportMapScript(c) {
				toRemove = append(toRemove, c)
			} else {
				walk(c)
			}
		}
	}
	walk(root)
	for _, n := range toRemove {
		n.Parent.RemoveChild(n)
	}
}

func findFirstScript(head *html.Node) *html.Node {
	for c := head.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "script" {
			return c
		}
	}
	return nil
}

// injectImportMapFallback uses string replacement when DOM parsing fails.
// Inserts before the first <script> in the document, or before </head>.
func injectImportMapFallback(htmlStr string, importMapJSON string) string {
	importMapScript := "<script type=\"importmap\">\n" + importMapJSON + "\n</script>"

	if idx := strings.Index(htmlStr, "<script"); idx != -1 {
		return htmlStr[:idx] + importMapScript + "\n" + htmlStr[idx:]
	}

	if idx := strings.Index(htmlStr, "</head>"); idx != -1 {
		return htmlStr[:idx] + importMapScript + "\n" + htmlStr[idx:]
	}

	if idx := strings.Index(htmlStr, "<body"); idx != -1 {
		closeIdx := strings.Index(htmlStr[idx:], ">")
		if closeIdx != -1 {
			insertPos := idx + closeIdx + 1
			return htmlStr[:insertPos] + "\n" + importMapScript + htmlStr[insertPos:]
		}
	}

	return htmlStr
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
