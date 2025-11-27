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

package transform

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
)

func TestTypeScript_DisabledByConfig(t *testing.T) {
	// Load fixtures into in-memory filesystem
	mfs := testutil.NewFixtureFS(t, "transforms/config-test", "/test")

	// Create middleware with Enabled=false
	middleware := NewTypeScript(TypeScriptConfig{
		WatchDirFunc: func() string { return "/test" },
		Enabled:      false,
		FS:           mfs,
	})

	// Create test handler that should be called (middleware should skip)
	called := false
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	})

	// Request .js URL (which would normally transform .ts file)
	req := httptest.NewRequest("GET", "/test.js", nil)
	rec := httptest.NewRecorder()

	handler := middleware(next)
	handler.ServeHTTP(rec, req)

	// When disabled, middleware should pass through to next handler
	if !called {
		t.Fatal("Expected next handler to be called when transform is disabled")
	}
}

func TestTypeScript_EnabledByConfig(t *testing.T) {
	// Load fixtures into in-memory filesystem
	mfs := testutil.NewFixtureFS(t, "transforms/config-test", "/test")

	// Create middleware with Enabled=true and cache
	cache := NewCache(1024 * 1024)
	middleware := NewTypeScript(TypeScriptConfig{
		WatchDirFunc:     func() string { return "/test" },
		TsconfigRawFunc:  func() string { return "" },
		Cache:            cache,
		Logger:           &mockLogger{},
		Enabled:          true,
		Target:           "ES2022",
		ErrorBroadcaster: &mockErrorBroadcaster{},
		FS:               mfs,
	})

	// Create test handler
	called := false
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	})

	// Request .js URL (should transform .ts file)
	req := httptest.NewRequest("GET", "/test.js", nil)
	rec := httptest.NewRecorder()

	handler := middleware(next)
	handler.ServeHTTP(rec, req)

	// When enabled, middleware should handle transform (not call next)
	if called {
		t.Fatal("Expected transform middleware to handle request when enabled")
	}

	// Should return JavaScript content
	result := rec.Result()
	if result.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", result.StatusCode)
	}

	body, _ := io.ReadAll(result.Body)
	if len(body) == 0 {
		t.Fatal("Expected transformed JavaScript content")
	}
}

func TestCSS_DisabledByConfig(t *testing.T) {
	// Load fixtures into in-memory filesystem
	mfs := testutil.NewFixtureFS(t, "transforms/config-test", "/test")

	// Create middleware with Enabled=false
	middleware := NewCSS(CSSConfig{
		WatchDirFunc: func() string { return "/test" },
		Enabled:      false,
		FS:           mfs,
	})

	// Create test handler that should be called
	called := false
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	})

	// Request .css URL (normally would transform)
	req := httptest.NewRequest("GET", "/test.css", nil)
	rec := httptest.NewRecorder()

	handler := middleware(next)
	handler.ServeHTTP(rec, req)

	// When disabled, should pass through
	if !called {
		t.Fatal("Expected next handler to be called when CSS transform is disabled")
	}
}

func TestCSS_EnabledByConfig(t *testing.T) {
	// Load fixtures into in-memory filesystem
	mfs := testutil.NewFixtureFS(t, "transforms/config-test", "/test")

	// Create middleware with Enabled=true and include pattern
	middleware := NewCSS(CSSConfig{
		WatchDirFunc: func() string { return "/test" },
		Enabled:      true,
		Include:      []string{"**/*.css"}, // Opt-in to transform CSS
		Logger:       &mockLogger{},
		FS:           mfs,
	})

	// Create test handler
	called := false
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	})

	// Request .css URL
	req := httptest.NewRequest("GET", "/test.css", nil)
	rec := httptest.NewRecorder()

	handler := middleware(next)
	handler.ServeHTTP(rec, req)

	// When enabled, should transform (not call next)
	if called {
		t.Fatal("Expected CSS transform middleware to handle request when enabled")
	}

	// Should return JavaScript module
	result := rec.Result()
	if result.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", result.StatusCode)
	}

	contentType := result.Header.Get("Content-Type")
	if contentType != "application/javascript; charset=utf-8" {
		t.Errorf("Expected JavaScript content-type, got %s", contentType)
	}
}

// Mock error broadcaster for testing
type mockErrorBroadcaster struct{}

func (m *mockErrorBroadcaster) BroadcastError(title, message, filename string) {
	// No-op for testing
}
