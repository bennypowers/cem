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

package cors_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"bennypowers.dev/cem/serve/middleware/cors"
)

// TestCORS_SetsCORSHeaders tests that CORS headers are set
func TestCORS_SetsCORSHeaders(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mw := cors.New()
	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	if rec.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Errorf("Expected CORS header to be '*', got '%s'", rec.Header().Get("Access-Control-Allow-Origin"))
	}
}

// TestCORS_SetsSecurityHeaders tests that security headers are set
func TestCORS_SetsSecurityHeaders(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mw := cors.New()
	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	if rec.Header().Get("X-Content-Type-Options") != "nosniff" {
		t.Errorf("Expected X-Content-Type-Options header to be 'nosniff', got '%s'",
			rec.Header().Get("X-Content-Type-Options"))
	}
}

// TestCORS_CallsNextHandler tests that the next handler is called
func TestCORS_CallsNextHandler(t *testing.T) {
	nextCalled := false
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCalled = true
		w.WriteHeader(http.StatusOK)
	})

	mw := cors.New()
	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	if !nextCalled {
		t.Error("Expected next handler to be called")
	}
}

// TestCORS_PreservesStatusCode tests that status code from handler is preserved
func TestCORS_PreservesStatusCode(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	mw := cors.New()
	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", rec.Code)
	}
}

// TestCORS_PreservesBody tests that response body is preserved
func TestCORS_PreservesBody(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("test body"))
	})

	mw := cors.New()
	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	if rec.Body.String() != "test body" {
		t.Errorf("Expected body 'test body', got '%s'", rec.Body.String())
	}
}

// TestCORS_PreservesHandlerHeaders tests that headers set by handler are preserved
func TestCORS_PreservesHandlerHeaders(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Custom", "value")
		w.WriteHeader(http.StatusOK)
	})

	mw := cors.New()
	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	if rec.Header().Get("Content-Type") != "application/json" {
		t.Error("Expected Content-Type header to be preserved")
	}
	if rec.Header().Get("X-Custom") != "value" {
		t.Error("Expected X-Custom header to be preserved")
	}
	// CORS headers should also be present
	if rec.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Error("Expected CORS headers to be set")
	}
}

// TestCORS_MultipleRequests tests that middleware works for multiple requests
func TestCORS_MultipleRequests(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mw := cors.New()
	wrapped := mw(handler)

	for i := 0; i < 10; i++ {
		req := httptest.NewRequest("GET", "/test", nil)
		rec := httptest.NewRecorder()
		wrapped.ServeHTTP(rec, req)

		if rec.Header().Get("Access-Control-Allow-Origin") != "*" {
			t.Errorf("Request %d: Expected CORS header to be set", i)
		}
	}
}

// TestCORS_DifferentMethods tests that CORS works with different HTTP methods
func TestCORS_DifferentMethods(t *testing.T) {
	methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH"}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mw := cors.New()
	wrapped := mw(handler)

	for _, method := range methods {
		req := httptest.NewRequest(method, "/test", nil)
		rec := httptest.NewRecorder()
		wrapped.ServeHTTP(rec, req)

		if rec.Header().Get("Access-Control-Allow-Origin") != "*" {
			t.Errorf("Method %s: Expected CORS header to be set", method)
		}
	}
}
