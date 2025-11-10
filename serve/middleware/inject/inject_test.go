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

package inject_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"bennypowers.dev/cem/serve/middleware/inject"
)

// TestInject_DisabledNoInjection tests that no injection happens when disabled
func TestInject_DisabledNoInjection(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte("<html><head></head><body>test</body></html>"))
	})

	mw := inject.New(false, "/script.js")
	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/test.html", nil)
	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	body := rec.Body.String()
	if strings.Contains(body, "/script.js") {
		t.Error("Expected no script injection when disabled")
	}
}

// TestInject_InjectsIntoHTML tests that script is injected into HTML
func TestInject_InjectsIntoHTML(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte("<html><head></head><body>test</body></html>"))
	})

	mw := inject.New(true, "/__cem/websocket-client.js")
	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/test.html", nil)
	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	body := rec.Body.String()
	if !strings.Contains(body, `<script type="module" src="/__cem/websocket-client.js"></script>`) {
		t.Error("Expected script to be injected into HTML")
	}
}

// TestInject_OnlyInjectsIntoHTMLFiles tests that injection only happens for .html files
func TestInject_OnlyInjectsIntoHTMLFiles(t *testing.T) {
	tests := []struct {
		name   string
		path   string
		inject bool
	}{
		{"HTML file", "/test.html", true},
		{"Root path", "/", true},
		{"JS file", "/script.js", false},
		{"CSS file", "/style.css", false},
		{"JSON file", "/data.json", false},
		{"Image", "/image.png", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/html")
				_, _ = w.Write([]byte("<html><head></head><body>test</body></html>"))
			})

			mw := inject.New(true, "/script.js")
			wrapped := mw(handler)

			req := httptest.NewRequest("GET", tt.path, nil)
			rec := httptest.NewRecorder()
			wrapped.ServeHTTP(rec, req)

			body := rec.Body.String()
			hasScript := strings.Contains(body, "/script.js")
			if hasScript != tt.inject {
				t.Errorf("Path %s: expected inject=%v, got inject=%v", tt.path, tt.inject, hasScript)
			}
		})
	}
}

// TestInject_OnlyInjectsIntoHTMLContentType tests content type detection
func TestInject_OnlyInjectsIntoHTMLContentType(t *testing.T) {
	tests := []struct {
		name        string
		path        string
		contentType string
		body        string
		inject      bool
	}{
		{"HTML content type", "/test.html", "text/html", "<html><head></head><body>test</body></html>", true},
		{"HTML with charset", "/test.html", "text/html; charset=utf-8", "<html><head></head><body>test</body></html>", true},
		{"JSON content type with .html path", "/test.html", "application/json", `{"message":"hello"}`, false},
		{"JavaScript with .html path", "/test.html", "application/javascript", "console.log('test');", false},
		{"Plain text with .html path", "/test.html", "text/plain", "plain text content", false},
		{"No content type with HTML", "/test.html", "", "<html><head></head><body>test</body></html>", true},
		{"No content type without HTML", "/test.html", "", "plain text", false},
		{"Root path with HTML", "/", "text/html", "<html><head></head><body>test</body></html>", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if tt.contentType != "" {
					w.Header().Set("Content-Type", tt.contentType)
				}
				_, _ = w.Write([]byte(tt.body))
			})

			mw := inject.New(true, "/script.js")
			wrapped := mw(handler)

			req := httptest.NewRequest("GET", tt.path, nil)
			rec := httptest.NewRecorder()
			wrapped.ServeHTTP(rec, req)

			body := rec.Body.String()
			hasScript := strings.Contains(body, "/script.js")
			if hasScript != tt.inject {
				t.Errorf("Path %s, Content-Type %s: expected inject=%v, got inject=%v", tt.path, tt.contentType, tt.inject, hasScript)
			}
		})
	}
}

// TestInject_InjectsIntoHead tests that script is injected into <head>
func TestInject_InjectsIntoHead(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte("<html><head><title>Test</title></head><body>content</body></html>"))
	})

	mw := inject.New(true, "/script.js")
	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/test.html", nil)
	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	body := rec.Body.String()
	// Script should be in head
	if !strings.Contains(body, "<head>") {
		t.Fatal("Expected HTML to have <head>")
	}

	headEnd := strings.Index(body, "</head>")
	scriptPos := strings.Index(body, "/script.js")

	if scriptPos > headEnd {
		t.Error("Expected script to be injected before </head>")
	}
}

// TestInject_PreservesStatusCode tests that status code is preserved
func TestInject_PreservesStatusCode(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("<html><head></head><body>404</body></html>"))
	})

	mw := inject.New(true, "/script.js")
	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/test.html", nil)
	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", rec.Code)
	}
}

// TestInject_PreservesHeaders tests that headers are preserved
func TestInject_PreservesHeaders(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("X-Custom", "value")
		w.Header().Set("Cache-Control", "no-cache")
		_, _ = w.Write([]byte("<html><head></head><body>test</body></html>"))
	})

	mw := inject.New(true, "/script.js")
	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/test.html", nil)
	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	if rec.Header().Get("Content-Type") != "text/html; charset=utf-8" {
		t.Error("Expected Content-Type header to be preserved")
	}
	if rec.Header().Get("X-Custom") != "value" {
		t.Error("Expected X-Custom header to be preserved")
	}
	if rec.Header().Get("Cache-Control") != "no-cache" {
		t.Error("Expected Cache-Control header to be preserved")
	}
}

// TestInject_DoesNotSetContentLengthHeader tests that Content-Length is not set
func TestInject_DoesNotSetContentLengthHeader(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Content-Length", "100") // Original length
		_, _ = w.Write([]byte("<html><head></head><body>test</body></html>"))
	})

	mw := inject.New(true, "/script.js")
	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/test.html", nil)
	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	// Content-Length should be removed since body was modified
	if rec.Header().Get("Content-Length") == "100" {
		t.Error("Expected Content-Length to be updated or removed after injection")
	}
}

// TestInject_SetsContentTypeIfMissing tests that Content-Type is set if missing
func TestInject_SetsContentTypeIfMissing(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Don't set Content-Type
		_, _ = w.Write([]byte("<html><head></head><body>test</body></html>"))
	})

	mw := inject.New(true, "/script.js")
	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/test.html", nil)
	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	contentType := rec.Header().Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		t.Errorf("Expected Content-Type to be set to text/html, got '%s'", contentType)
	}
}

// TestInject_HandlesEmptyHTML tests injection into minimal HTML
func TestInject_HandlesEmptyHTML(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte("<html></html>"))
	})

	mw := inject.New(true, "/script.js")
	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/test.html", nil)
	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	body := rec.Body.String()
	if !strings.Contains(body, "/script.js") {
		t.Error("Expected script to be injected even in minimal HTML")
	}
}

// TestInject_HandlesHTMLWithoutHead tests injection when no <head> tag
func TestInject_HandlesHTMLWithoutHead(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte("<html><body>test</body></html>"))
	})

	mw := inject.New(true, "/script.js")
	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/test.html", nil)
	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	body := rec.Body.String()
	if !strings.Contains(body, "/script.js") {
		t.Error("Expected script to be injected even without <head>")
	}
}

// TestInject_MultipleRequests tests that middleware works for multiple requests
func TestInject_MultipleRequests(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte("<html><head></head><body>test</body></html>"))
	})

	mw := inject.New(true, "/script.js")
	wrapped := mw(handler)

	for i := 0; i < 10; i++ {
		req := httptest.NewRequest("GET", "/test.html", nil)
		rec := httptest.NewRecorder()
		wrapped.ServeHTTP(rec, req)

		body := rec.Body.String()
		if !strings.Contains(body, "/script.js") {
			t.Errorf("Request %d: Expected script to be injected", i)
		}
	}
}

// TestInject_CustomScriptPath tests injection with custom script path
func TestInject_CustomScriptPath(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte("<html><head></head><body>test</body></html>"))
	})

	customPath := "/custom/path/to/script.js"
	mw := inject.New(true, customPath)
	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/test.html", nil)
	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	body := rec.Body.String()
	if !strings.Contains(body, customPath) {
		t.Errorf("Expected custom script path '%s' to be injected", customPath)
	}
}

// TestInject_NonHTMLPassthrough tests that non-HTML responses pass through unchanged
func TestInject_NonHTMLPassthrough(t *testing.T) {
	originalBody := `{"message": "hello"}`
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(originalBody))
	})

	mw := inject.New(true, "/script.js")
	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/test.html", nil)
	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	body := rec.Body.String()
	if body != originalBody {
		t.Errorf("Expected non-HTML body to be unchanged, got '%s'", body)
	}
	if strings.Contains(body, "/script.js") {
		t.Error("Expected no script injection into JSON response")
	}
}
