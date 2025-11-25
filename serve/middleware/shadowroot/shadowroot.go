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
			// Check if this is an element template request with attrs query params
			isElementWithAttrs := strings.HasPrefix(r.URL.Path, "/__cem/elements/") &&
				strings.HasSuffix(r.URL.Path, ".html") &&
				hasAttrsQueryParams(r.URL.Query())

			// Skip internal routes that need direct ResponseWriter access (WebSocket, etc.)
			// But allow element template requests with attrs to be processed
			if strings.HasPrefix(r.URL.Path, "/__cem") && !isElementWithAttrs {
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
				_, _ = w.Write(rec.Body())
				return
			}

			var htmlToProcess []byte

			// If this is an element template request with attrs, construct element HTML
			if isElementWithAttrs {
				elementHTML, err := m.constructElementHTML(r)
				if err != nil {
					logger.Error("Failed to construct element HTML", "error", err)
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				htmlToProcess = elementHTML
			} else {
				htmlToProcess = rec.Body()
			}

			// Process HTML and inject shadow roots
			processed, err := m.injectShadowRoots(htmlToProcess)
			if err != nil {
				logger.Error("Failed to inject shadow roots", "error", err)
				// Fall back to original HTML on error
				for k, v := range rec.Header() {
					w.Header()[k] = v
				}
				w.WriteHeader(rec.StatusCode())
				_, _ = w.Write(rec.Body())
				return
			}

			// If content=shadow, extract just the shadow root content
			if r.URL.Query().Get("content") == "shadow" {
				shadowContent, err := extractShadowContent(processed)
				if err != nil {
					logger.Error("Failed to extract shadow content", "error", err)
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				processed = shadowContent
			}

			// Write processed response
			for k, v := range rec.Header() {
				w.Header()[k] = v
			}
			w.WriteHeader(rec.StatusCode())
			_, _ = w.Write(processed)
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
			if err := writeToken(&out, token); err != nil {
				return nil, err
			}

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
						if _, err := out.WriteString(shadowRootHTML); err != nil {
							return nil, err
						}
					} else {
						if _, err := out.Write(processed); err != nil {
							return nil, err
						}
					}
				}
			}

		default:
			// Write all other tokens as-is
			if _, err := out.Write(z.Raw()); err != nil {
				return nil, err
			}
		}
	}
}

// shouldInjectShadowRoot checks if element should get shadow root injection
func (m *Middleware) shouldInjectShadowRoot(tagName string) bool {
	// Must be custom element (contains hyphen) AND have template
	return strings.Contains(tagName, "-") && m.knownElements.Has(tagName)
}

// writeToken writes an HTML token to the output buffer
func writeToken(w io.Writer, token html.Token) error {
	if _, err := w.Write([]byte("<")); err != nil {
		return err
	}
	if token.Type == html.EndTagToken {
		if _, err := w.Write([]byte("/")); err != nil {
			return err
		}
	}
	if _, err := w.Write([]byte(token.Data)); err != nil {
		return err
	}

	for _, attr := range token.Attr {
		if _, err := w.Write([]byte(" ")); err != nil {
			return err
		}
		if _, err := w.Write([]byte(attr.Key)); err != nil {
			return err
		}
		if attr.Val != "" {
			if _, err := w.Write([]byte(`="`)); err != nil {
				return err
			}
			if _, err := w.Write([]byte(html.EscapeString(attr.Val))); err != nil {
				return err
			}
			if _, err := w.Write([]byte(`"`)); err != nil {
				return err
			}
		}
	}

	if token.Type == html.SelfClosingTagToken {
		if _, err := w.Write([]byte(" /")); err != nil {
			return err
		}
	}
	if _, err := w.Write([]byte(">")); err != nil {
		return err
	}
	return nil
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

// hasAttrsQueryParams checks if the query has attrs[*] parameters
func hasAttrsQueryParams(query map[string][]string) bool {
	for key := range query {
		if strings.HasPrefix(key, "attrs[") {
			return true
		}
	}
	return false
}

// constructElementHTML builds element HTML from request path and query params
func (m *Middleware) constructElementHTML(r *http.Request) ([]byte, error) {
	// Extract element name from path: /__cem/elements/pf-v6-button/pf-v6-button.html
	path := strings.TrimPrefix(r.URL.Path, "/__cem/elements/")
	parts := strings.Split(path, "/")
	if len(parts) == 0 {
		return nil, io.ErrUnexpectedEOF
	}
	elementName := parts[0]

	// Parse attrs from query
	query := r.URL.Query()
	var attrPairs []string
	for key, values := range query {
		if strings.HasPrefix(key, "attrs[") && strings.HasSuffix(key, "]") {
			attrName := key[6 : len(key)-1] // Extract "key" from "attrs[key]"
			if len(values) > 0 {
				value := values[0]
				if value == "" {
					attrPairs = append(attrPairs, attrName)
				} else {
					// HTML escape the value
					escaped := strings.ReplaceAll(value, `"`, `&quot;`)
					attrPairs = append(attrPairs, attrName+`="`+escaped+`"`)
				}
			}
		}
	}

	attrString := strings.Join(attrPairs, " ")
	var html string
	if attrString != "" {
		html = "<" + elementName + " " + attrString + "></" + elementName + ">"
	} else {
		html = "<" + elementName + "></" + elementName + ">"
	}

	return []byte(html), nil
}

// extractShadowContent extracts just the shadow root content from DSD HTML
func extractShadowContent(dsdHTML []byte) ([]byte, error) {
	doc, err := html.Parse(bytes.NewReader(dsdHTML))
	if err != nil {
		return nil, err
	}

	// Find <template shadowrootmode="open">
	var templateNode *html.Node
	var findTemplate func(*html.Node)
	findTemplate = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "template" {
			for _, attr := range n.Attr {
				if attr.Key == "shadowrootmode" && attr.Val == "open" {
					templateNode = n
					return
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findTemplate(c)
			if templateNode != nil {
				return
			}
		}
	}
	findTemplate(doc)

	if templateNode == nil {
		return nil, io.ErrUnexpectedEOF
	}

	// Render the children of the template node
	var buf bytes.Buffer
	for c := templateNode.FirstChild; c != nil; c = c.NextSibling {
		err := html.Render(&buf, c)
		if err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}
