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

package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"bennypowers.dev/cem/serve/middleware"
)

// TestChain_SingleMiddleware tests Chain with a single middleware
func TestChain_SingleMiddleware(t *testing.T) {
	called := false
	mw := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			called = true
			w.Header().Set("X-Test", "middleware")
			next.ServeHTTP(w, r)
		})
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	chained := middleware.Chain(handler, mw)

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()
	chained.ServeHTTP(rec, req)

	if !called {
		t.Error("Expected middleware to be called")
	}
	if rec.Header().Get("X-Test") != "middleware" {
		t.Error("Expected middleware to set header")
	}
}

// TestChain_MultipleMiddlewares tests Chain with multiple middlewares
func TestChain_MultipleMiddlewares(t *testing.T) {
	var order []string

	mw1 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			order = append(order, "mw1-before")
			next.ServeHTTP(w, r)
			order = append(order, "mw1-after")
		})
	}

	mw2 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			order = append(order, "mw2-before")
			next.ServeHTTP(w, r)
			order = append(order, "mw2-after")
		})
	}

	mw3 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			order = append(order, "mw3-before")
			next.ServeHTTP(w, r)
			order = append(order, "mw3-after")
		})
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		order = append(order, "handler")
		w.WriteHeader(http.StatusOK)
	})

	// Chain applies in reverse, so mw3 is outermost, mw1 is innermost
	chained := middleware.Chain(handler, mw3, mw2, mw1)

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()
	chained.ServeHTTP(rec, req)

	// Chain applies middlewares in reverse order, so:
	// mw3 wraps mw2 wraps mw1 wraps handler
	// Therefore: mw3-before, mw2-before, mw1-before, handler, mw1-after, mw2-after, mw3-after
	expected := []string{
		"mw3-before", "mw2-before", "mw1-before",
		"handler",
		"mw1-after", "mw2-after", "mw3-after",
	}

	if len(order) != len(expected) {
		t.Fatalf("Expected %d calls, got %d", len(expected), len(order))
	}

	for i, exp := range expected {
		if order[i] != exp {
			t.Errorf("At position %d: expected %s, got %s", i, exp, order[i])
		}
	}
}

// TestChain_NoMiddlewares tests Chain with no middlewares
func TestChain_NoMiddlewares(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	chained := middleware.Chain(handler)

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()
	chained.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}
}

// TestResponseRecorder_Write tests basic write operations
func TestResponseRecorder_Write(t *testing.T) {
	rec := middleware.NewResponseRecorder()

	n, err := rec.Write([]byte("hello"))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if n != 5 {
		t.Errorf("Expected 5 bytes written, got %d", n)
	}

	body := rec.Body()
	if string(body) != "hello" {
		t.Errorf("Expected body 'hello', got '%s'", string(body))
	}
}

// TestResponseRecorder_WriteHeader tests status code capture
func TestResponseRecorder_WriteHeader(t *testing.T) {
	rec := middleware.NewResponseRecorder()

	rec.WriteHeader(http.StatusNotFound)

	if rec.StatusCode() != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", rec.StatusCode())
	}
}

// TestResponseRecorder_DefaultStatus tests default status code
func TestResponseRecorder_DefaultStatus(t *testing.T) {
	rec := middleware.NewResponseRecorder()

	// Write without calling WriteHeader
	_, _ = rec.Write([]byte("test"))

	if rec.StatusCode() != http.StatusOK {
		t.Errorf("Expected default status 200, got %d", rec.StatusCode())
	}
}

// TestResponseRecorder_Header tests header manipulation
func TestResponseRecorder_Header(t *testing.T) {
	rec := middleware.NewResponseRecorder()

	rec.Header().Set("Content-Type", "application/json")
	rec.Header().Add("X-Custom", "value1")
	rec.Header().Add("X-Custom", "value2")

	if rec.Header().Get("Content-Type") != "application/json" {
		t.Error("Expected Content-Type header to be set")
	}

	values := rec.Header().Values("X-Custom")
	if len(values) != 2 {
		t.Errorf("Expected 2 X-Custom headers, got %d", len(values))
	}
}

// TestResponseRecorder_MultipleWrites tests multiple write calls
func TestResponseRecorder_MultipleWrites(t *testing.T) {
	rec := middleware.NewResponseRecorder()

	_, _ = rec.Write([]byte("hello"))
	_, _ = rec.Write([]byte(" "))
	_, _ = rec.Write([]byte("world"))

	body := rec.Body()
	if string(body) != "hello world" {
		t.Errorf("Expected 'hello world', got '%s'", string(body))
	}
}

// TestResponseRecorder_WriteHeaderOnce tests that WriteHeader is only called once
func TestResponseRecorder_WriteHeaderOnce(t *testing.T) {
	rec := middleware.NewResponseRecorder()

	rec.WriteHeader(http.StatusNotFound)
	rec.WriteHeader(http.StatusOK) // Should be ignored

	if rec.StatusCode() != http.StatusNotFound {
		t.Errorf("Expected status 404 (first call), got %d", rec.StatusCode())
	}
}

// TestCopyHeaders_BasicCopy tests basic header copying
func TestCopyHeaders_BasicCopy(t *testing.T) {
	src := make(http.Header)
	src.Set("Content-Type", "text/html")
	src.Set("X-Custom", "value")

	dst := make(http.Header)

	middleware.CopyHeaders(dst, src)

	if dst.Get("Content-Type") != "text/html" {
		t.Error("Expected Content-Type to be copied")
	}
	if dst.Get("X-Custom") != "value" {
		t.Error("Expected X-Custom to be copied")
	}
}

// TestCopyHeaders_WithExclusions tests header copying with exclusions
func TestCopyHeaders_WithExclusions(t *testing.T) {
	src := make(http.Header)
	src.Set("Content-Type", "text/html")
	src.Set("Content-Length", "100")
	src.Set("X-Custom", "value")

	dst := make(http.Header)

	middleware.CopyHeaders(dst, src, "Content-Length")

	if dst.Get("Content-Type") != "text/html" {
		t.Error("Expected Content-Type to be copied")
	}
	if dst.Get("X-Custom") != "value" {
		t.Error("Expected X-Custom to be copied")
	}
	if dst.Get("Content-Length") != "" {
		t.Error("Expected Content-Length to be excluded")
	}
}

// TestCopyHeaders_MultipleValues tests copying headers with multiple values
func TestCopyHeaders_MultipleValues(t *testing.T) {
	src := make(http.Header)
	src.Add("X-Custom", "value1")
	src.Add("X-Custom", "value2")
	src.Add("X-Custom", "value3")

	dst := make(http.Header)

	middleware.CopyHeaders(dst, src)

	values := dst.Values("X-Custom")
	if len(values) != 3 {
		t.Fatalf("Expected 3 values, got %d", len(values))
	}
	if values[0] != "value1" || values[1] != "value2" || values[2] != "value3" {
		t.Error("Expected all values to be copied in order")
	}
}

// TestIsHTMLResponse_HTMLContentType tests HTML detection by content type
func TestIsHTMLResponse_HTMLContentType(t *testing.T) {
	tests := []struct {
		name        string
		contentType string
		want        bool
	}{
		{"HTML with charset", "text/html; charset=utf-8", true},
		{"Plain HTML", "text/html", true},
		{"JSON", "application/json", false},
		{"JavaScript", "application/javascript", false},
		{"Plain text", "text/plain", false},
		{"Empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := middleware.IsHTMLResponse(tt.contentType)
			if got != tt.want {
				t.Errorf("IsHTMLResponse(%q) = %v, want %v", tt.contentType, got, tt.want)
			}
		})
	}
}

// TestResponseRecorder_ImplementsResponseWriter tests interface compliance
func TestResponseRecorder_ImplementsResponseWriter(t *testing.T) {
	var _ http.ResponseWriter = middleware.NewResponseRecorder()
}
