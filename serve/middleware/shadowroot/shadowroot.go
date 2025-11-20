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

package shadowroot

import (
	"bytes"
	"embed"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"

	"bennypowers.dev/cem/serve/middleware"
	"bennypowers.dev/cem/serve/middleware/types"
	"bennypowers.dev/cem/set"
)

// Middleware injects declarative shadow DOM for custom elements
type Middleware struct {
	logger         types.Logger
	broadcaster    types.ErrorBroadcaster
	knownElements  set.Set[string]
	renderShadow   ShadowRootRenderer
}

// ShadowRootRenderer renders shadow DOM template for an element
type ShadowRootRenderer func(elementName string, data interface{}) (string, error)

// New creates a new shadow root injection middleware
func New(logger types.Logger, broadcaster types.ErrorBroadcaster, templatesFS embed.FS, renderer ShadowRootRenderer) middleware.Middleware {
	knownElements := buildKnownElements(templatesFS)
	return newWithKnownElements(logger, broadcaster, knownElements, renderer)
}

// NewWithPrefix creates middleware with a custom templates directory prefix
func NewWithPrefix(logger types.Logger, broadcaster types.ErrorBroadcaster, templatesFS embed.FS, prefix string, renderer ShadowRootRenderer) middleware.Middleware {
	knownElements := buildKnownElementsWithPrefix(templatesFS, prefix)
	return newWithKnownElements(logger, broadcaster, knownElements, renderer)
}

// newWithKnownElements creates middleware with a pre-built set of known elements
func newWithKnownElements(logger types.Logger, broadcaster types.ErrorBroadcaster, knownElements set.Set[string], renderer ShadowRootRenderer) middleware.Middleware {
	m := &Middleware{
		logger:        logger,
		broadcaster:   broadcaster,
		knownElements: knownElements,
		renderShadow:  renderer,
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Skip internal routes that need direct ResponseWriter access (WebSocket, etc.)
			if strings.HasPrefix(r.URL.Path, "/__cem") {
				next.ServeHTTP(w, r)
				return
			}

			// Capture response
			rec := middleware.NewResponseRecorder()
			next.ServeHTTP(rec, r)

			// Only process HTML responses
			contentType := rec.Header().Get("Content-Type")
			if !middleware.IsHTMLResponse(contentType) {
				// Pass through non-HTML responses
				for k, v := range rec.Header() {
					w.Header()[k] = v
				}
				w.WriteHeader(rec.StatusCode())
				w.Write(rec.Body())
				return
			}

			// Process HTML and inject shadow roots
			processed, err := m.injectShadowRoots(rec.Body())
			if err != nil {
				logger.Error("Failed to inject shadow roots", "error", err)
				// Fall back to original HTML on error
				for k, v := range rec.Header() {
					w.Header()[k] = v
				}
				w.WriteHeader(rec.StatusCode())
				w.Write(rec.Body())
				return
			}

			// Write processed response
			for k, v := range rec.Header() {
				w.Header()[k] = v
			}
			w.WriteHeader(rec.StatusCode())
			w.Write(processed)
		})
	}
}

// injectShadowRoots processes HTML and injects shadow roots for custom elements
func (m *Middleware) injectShadowRoots(htmlBytes []byte) ([]byte, error) {
	return m.processHTML(bytes.NewReader(htmlBytes))
}

// processHTML recursively processes HTML tokens and injects shadow roots
func (m *Middleware) processHTML(reader io.Reader) ([]byte, error) {
	var out bytes.Buffer
	z := html.NewTokenizer(reader)

	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			if err := z.Err(); err != io.EOF {
				return nil, err
			}
			return out.Bytes(), nil

		case html.StartTagToken, html.SelfClosingTagToken:
			token := z.Token()

			// Write opening tag
			writeToken(&out, token)

			// Check if this is a known custom element
			if m.shouldInjectShadowRoot(token.Data) {
				// Render shadow root with attributes as template data
				attrData := attributesToTemplateData(token.Attr)
				shadowRootHTML, err := m.renderShadow(token.Data, attrData)

				if err != nil {
					m.logAndBroadcastError(token.Data, err)
				} else {
					// Recursively process shadow root HTML for nested custom elements
					processed, err := m.processHTML(strings.NewReader(shadowRootHTML))
					if err != nil {
						m.logger.Warning("Failed to process nested shadow root",
							"element", token.Data,
							"error", err)
						// Write unprocessed shadow root on error
						out.WriteString(shadowRootHTML)
					} else {
						out.Write(processed)
					}
				}
			}

		default:
			// Write all other tokens as-is
			out.Write(z.Raw())
		}
	}
}

// shouldInjectShadowRoot checks if element should get shadow root injection
func (m *Middleware) shouldInjectShadowRoot(tagName string) bool {
	// Must be custom element (contains hyphen) AND have template
	return strings.Contains(tagName, "-") && m.knownElements.Has(tagName)
}

// writeToken writes an HTML token to the output buffer
func writeToken(w io.Writer, token html.Token) {
	w.Write([]byte("<"))
	if token.Type == html.EndTagToken {
		w.Write([]byte("/"))
	}
	w.Write([]byte(token.Data))

	for _, attr := range token.Attr {
		w.Write([]byte(" "))
		w.Write([]byte(attr.Key))
		if attr.Val != "" {
			w.Write([]byte(`="`))
			w.Write([]byte(html.EscapeString(attr.Val)))
			w.Write([]byte(`"`))
		}
	}

	if token.Type == html.SelfClosingTagToken {
		w.Write([]byte(" /"))
	}
	w.Write([]byte(">"))
}

// logAndBroadcastError logs and broadcasts missing template errors
func (m *Middleware) logAndBroadcastError(elementName string, err error) {
	m.logger.Warning("Shadow root template not found or failed to render",
		"element", elementName,
		"error", err)

	if m.broadcaster != nil {
		m.broadcaster.BroadcastError(
			"Shadow Root Template Error",
			err.Error(),
			elementName,
		)
	}
}
