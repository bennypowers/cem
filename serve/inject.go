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
)

// webSocketClientScript is the client-side JavaScript for live reload
const webSocketClientScript = `<script>
(function() {
  const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
  const ws = new WebSocket(protocol + '//' + window.location.host + '/__cem-reload');

  // Expose socket on window for debugging
  window.__cemReloadSocket = ws;

  ws.onopen = function() {
    console.log('[cem-serve] WebSocket connected');
  };

  ws.onmessage = function(event) {
    const data = JSON.parse(event.data);
    console.log('[cem-serve] Received reload payload:', data);
    if (data.type === 'reload') {
      console.log('[cem-serve] Reloading page:', data.reason, data.files);
      window.location.reload();
    }
  };

  ws.onclose = function() {
    console.log('[cem-serve] Connection closed, retrying in 1s...');
    setTimeout(function() {
      window.location.reload();
    }, 1000);
  };

  ws.onerror = function(error) {
    console.error('[cem-serve] WebSocket error:', error);
  };
})();
</script>`

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
			w.Write(bodyBytes)
			return
		}

		// Inject script before </head> or at start of <body>
		html := string(bodyBytes)
		injected := injectScript(html, webSocketClientScript)

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
		w.Write([]byte(injected))
	})
}

// injectScript injects a script into HTML
func injectScript(html, script string) string {
	// Try to inject before </head>
	if strings.Contains(html, "</head>") {
		return strings.Replace(html, "</head>", script+"\n</head>", 1)
	}

	// Try to inject at start of <body>
	if strings.Contains(html, "<body>") {
		return strings.Replace(html, "<body>", "<body>\n"+script, 1)
	}

	// Fallback: inject at end of HTML
	if strings.Contains(html, "</html>") {
		return strings.Replace(html, "</html>", script+"\n</html>", 1)
	}

	// Last resort: append to end
	return html + "\n" + script
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
