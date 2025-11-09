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
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// injectWebSocketClient injects the WebSocket client script into HTML responses
func injectWebSocketClient(next http.Handler, enabled bool) http.Handler {
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
		rec := &responseRecorder{
			header:     make(http.Header),
			body:       &bytes.Buffer{},
			statusCode: http.StatusOK,
		}

		next.ServeHTTP(rec, r)

		// Get body content
		bodyBytes := rec.body.Bytes()

		// Check if this is an HTML response
		contentType := rec.header.Get("Content-Type")
		isHTML := strings.Contains(contentType, "text/html") ||
			(contentType == "" && len(bodyBytes) > 0 && strings.Contains(string(bodyBytes[:min(100, len(bodyBytes))]), "<html"))

		if !isHTML {
			// Not HTML, write original response
			for k, v := range rec.header {
				for _, vv := range v {
					w.Header().Add(k, vv)
				}
			}
			w.WriteHeader(rec.statusCode)
			if _, err := w.Write(bodyBytes); err != nil {
				// Client disconnected or write error - can't respond
				return
			}
			return
		}

		// Inject script before </head> or at start of <body>
		html := string(bodyBytes)
		script := "<script type=\"module\" src=\"/__cem/websocket-client.js\"></script>"
		injected := injectScript(html, script)

		// Copy headers
		for k, v := range rec.header {
			if k != "Content-Length" { // Don't copy Content-Length, it will be wrong
				for _, vv := range v {
					w.Header().Add(k, vv)
				}
			}
		}

		// Set content type if not set
		if contentType == "" {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
		}

		w.WriteHeader(rec.statusCode)
		if _, err := w.Write([]byte(injected)); err != nil {
			// Client disconnected or write error - can't respond
			return
		}
	})
}

// injectScript injects a script into HTML by parsing and manipulating the DOM
func injectScript(htmlStr, script string) string {
	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		// Fallback to string replacement if parsing fails
		if strings.Contains(htmlStr, "</head>") {
			return strings.Replace(htmlStr, "</head>", script+"\n</head>", 1)
		}
		return htmlStr + "\n" + script
	}

	// Parse the script string into nodes
	scriptNodes, err := html.ParseFragment(strings.NewReader(script), &html.Node{
		Type:     html.ElementNode,
		Data:     "body",
		DataAtom: 0,
	})
	if err != nil || len(scriptNodes) == 0 {
		// Fallback if script parsing fails
		if strings.Contains(htmlStr, "</head>") {
			return strings.Replace(htmlStr, "</head>", script+"\n</head>", 1)
		}
		return htmlStr + "\n" + script
	}

	// Find the <head> element and inject before its closing tag
	var findHead func(*html.Node) *html.Node
	findHead = func(n *html.Node) *html.Node {
		if n.Type == html.ElementNode && n.Data == "head" {
			return n
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if result := findHead(c); result != nil {
				return result
			}
		}
		return nil
	}

	head := findHead(doc)
	if head != nil {
		// Append script nodes to head
		for _, scriptNode := range scriptNodes {
			head.AppendChild(scriptNode)
		}
	} else {
		// No head found, try to find body and prepend
		var findBody func(*html.Node) *html.Node
		findBody = func(n *html.Node) *html.Node {
			if n.Type == html.ElementNode && n.Data == "body" {
				return n
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if result := findBody(c); result != nil {
					return result
				}
			}
			return nil
		}

		body := findBody(doc)
		if body != nil {
			// Prepend script nodes to body
			for i := len(scriptNodes) - 1; i >= 0; i-- {
				if body.FirstChild != nil {
					body.InsertBefore(scriptNodes[i], body.FirstChild)
				} else {
					body.AppendChild(scriptNodes[i])
				}
			}
		}
	}

	// Render the modified DOM back to HTML
	var buf bytes.Buffer
	if err := html.Render(&buf, doc); err != nil {
		// Fallback on render error
		if strings.Contains(htmlStr, "</head>") {
			return strings.Replace(htmlStr, "</head>", script+"\n</head>", 1)
		}
		return htmlStr + "\n" + script
	}

	return buf.String()
}

// responseRecorder captures HTTP response for modification
type responseRecorder struct {
	header        http.Header
	body          *bytes.Buffer
	statusCode    int
	headerWritten bool
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	if !r.headerWritten {
		r.WriteHeader(http.StatusOK)
	}
	return r.body.Write(b)
}

func (r *responseRecorder) WriteHeader(statusCode int) {
	if !r.headerWritten {
		r.statusCode = statusCode
		r.headerWritten = true
	}
}

func (r *responseRecorder) Header() http.Header {
	return r.header
}

// Ensure responseRecorder implements http.ResponseWriter
var _ http.ResponseWriter = (*responseRecorder)(nil)
