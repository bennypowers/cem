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

package transform_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"bennypowers.dev/cem/cmd/config"
	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/serve/middleware/transform"
)

// TestTypeScriptMiddleware_SrcDistPattern tests full request cycle with src/dist separation
func TestTypeScriptMiddleware_SrcDistPattern(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/src-dist", "/test")
	cache := transform.NewCache(1024 * 1024) // 1MB cache

	// Create TypeScript middleware with path mapping support
	middleware := transform.NewTypeScript(transform.TypeScriptConfig{
		WatchDirFunc: func() string { return "/test" },
		TsconfigRawFunc: func() string {
			// Return minimal tsconfig
			return `{"compilerOptions": {"target": "ES2020"}}`
		},
		Cache:   cache,
		Pool:    nil, // Skip pool for test simplicity
		Logger:  nil,
		Enabled: true,
		FS:      fs,
		PathResolver: transform.NewPathResolver("/test", []config.URLRewrite{
			{URLPattern: "/dist/:path*", URLTemplate: "/src/{{.path}}"},
		}, fs, nil),
	})

	// Create test request for /dist/components/widget.js
	req := httptest.NewRequest("GET", "/dist/components/widget.js", nil)
	rec := httptest.NewRecorder()

	// Terminal handler (should not be called if transform succeeds)
	terminalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Terminal handler called - transform middleware should have handled request")
	})

	// Apply middleware
	handler := middleware(terminalHandler)
	handler.ServeHTTP(rec, req)

	// Verify response
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	// Verify Content-Type
	contentType := rec.Header().Get("Content-Type")
	if !strings.Contains(contentType, "application/javascript") {
		t.Errorf("Expected Content-Type to contain 'application/javascript', got %s", contentType)
	}

	// Verify TypeScript was transformed
	body := rec.Body.String()
	if len(body) == 0 {
		t.Fatal("Response body is empty")
	}

	// Verify TypeScript-specific syntax is removed
	if strings.Contains(body, ": string") || strings.Contains(body, "interface WidgetOptions") {
		t.Error("Response still contains TypeScript syntax")
	}

	// Verify source map references correct file
	if !strings.Contains(body, "sourceMappingURL=") {
		t.Error("Response missing source map")
	}
}

// TestTypeScriptMiddleware_WorkspaceMode tests request with package prefix
func TestTypeScriptMiddleware_WorkspaceMode(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/workspace-mode", "/test")
	cache := transform.NewCache(1024 * 1024)

	middleware := transform.NewTypeScript(transform.TypeScriptConfig{
		WatchDirFunc:    func() string { return "/test" },
		TsconfigRawFunc: func() string { return `{}` },
		Cache:           cache,
		Pool:            nil,
		Logger:          nil,
		Enabled:         true,
		FS:              fs,
		PathResolver: transform.NewPathResolver("/test", []config.URLRewrite{
			{URLPattern: "/packages/webawesome/dist/:path*", URLTemplate: "/packages/webawesome/src/{{.path}}"},
		}, fs, nil),
	})

	// Request with package prefix: /packages/webawesome/dist/components/alert.js
	req := httptest.NewRequest("GET", "/packages/webawesome/dist/components/alert.js", nil)
	rec := httptest.NewRecorder()

	terminalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Terminal handler called - transform should have handled request")
	})

	handler := middleware(terminalHandler)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	body := rec.Body.String()
	if len(body) == 0 {
		t.Fatal("Response body is empty")
	}

	// Verify TypeScript enums and types are removed
	if strings.Contains(body, "type AlertVariant") || strings.Contains(body, "AlertVariant = ") {
		t.Error("Response still contains TypeScript type definitions")
	}
}

// TestTypeScriptMiddleware_ExplicitPathMappings tests user-configured mappings
func TestTypeScriptMiddleware_ExplicitPathMappings(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/explicit-mappings", "/test")
	cache := transform.NewCache(1024 * 1024)

	// Configure explicit path mapping
	middleware := transform.NewTypeScript(transform.TypeScriptConfig{
		WatchDirFunc:    func() string { return "/test" },
		TsconfigRawFunc: func() string { return `{}` },
		Cache:           cache,
		Pool:            nil,
		Logger:          nil,
		Enabled:         true,
		FS:              fs,
		PathResolver: transform.NewPathResolver("/test", []config.URLRewrite{
			{URLPattern: "/lib/:path*", URLTemplate: "/sources/{{.path}}"}, // Custom mapping
		}, fs, nil),
	})

	// Request using custom mapping
	req := httptest.NewRequest("GET", "/lib/helpers/formatter.js", nil)
	rec := httptest.NewRecorder()

	terminalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Terminal handler called")
	})

	handler := middleware(terminalHandler)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	body := rec.Body.String()
	if len(body) == 0 {
		t.Fatal("Response body is empty")
	}

	// Verify class was transformed
	if !strings.Contains(body, "Formatter") {
		t.Error("Response missing Formatter class")
	}
}

// TestTypeScriptMiddleware_Caching tests that transformed files are cached
func TestTypeScriptMiddleware_Caching(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/src-dist", "/test")
	cache := transform.NewCache(1024 * 1024)

	middleware := transform.NewTypeScript(transform.TypeScriptConfig{
		WatchDirFunc:    func() string { return "/test" },
		TsconfigRawFunc: func() string { return `{}` },
		Cache:           cache,
		Pool:            nil,
		Logger:          nil,
		Enabled:         true,
		FS:              fs,
		PathResolver: transform.NewPathResolver("/test", []config.URLRewrite{
			{URLPattern: "/dist/:path*", URLTemplate: "/src/{{.path}}"},
		}, fs, nil),
	})

	terminalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	// First request - cache miss
	req1 := httptest.NewRequest("GET", "/dist/components/widget.js", nil)
	rec1 := httptest.NewRecorder()
	handler := middleware(terminalHandler)
	handler.ServeHTTP(rec1, req1)

	firstResponse := rec1.Body.String()

	// Second request - should hit cache
	req2 := httptest.NewRequest("GET", "/dist/components/widget.js", nil)
	rec2 := httptest.NewRecorder()
	handler.ServeHTTP(rec2, req2)

	secondResponse := rec2.Body.String()

	// Both responses should be identical
	if firstResponse != secondResponse {
		t.Error("Cached response differs from original")
	}

	// Verify cache stats
	stats := cache.Stats()
	if stats.Hits == 0 {
		t.Error("Expected cache hit, got 0 hits")
	}
}

// TestTypeScriptMiddleware_Disabled tests that middleware passes through when disabled
func TestTypeScriptMiddleware_Disabled(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/src-dist", "/test")

	middleware := transform.NewTypeScript(transform.TypeScriptConfig{
		WatchDirFunc:    func() string { return "/test" },
		TsconfigRawFunc: func() string { return `{}` },
		Cache:           nil,
		Pool:            nil,
		Logger:          nil,
		Enabled:         false, // Disabled
		FS:              fs,
		PathResolver:    nil,
	})

	terminalCalled := false
	terminalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		terminalCalled = true
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "/dist/components/widget.js", nil)
	rec := httptest.NewRecorder()

	handler := middleware(terminalHandler)
	handler.ServeHTTP(rec, req)

	if !terminalCalled {
		t.Error("Terminal handler not called when middleware disabled")
	}
}

// TestTypeScriptMiddleware_PassthroughNonTypeScript tests that non-TS files pass through
func TestTypeScriptMiddleware_PassthroughNonTypeScript(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/in-place", "/test")
	// Add a plain .js file (no corresponding .ts)
	fs.AddFile("/test/components/plain.js", "export const value = 42;", 0644)

	middleware := transform.NewTypeScript(transform.TypeScriptConfig{
		WatchDirFunc:    func() string { return "/test" },
		TsconfigRawFunc: func() string { return `{}` },
		Cache:           transform.NewCache(1024 * 1024),
		Pool:            nil,
		Logger:          nil,
		Enabled:         true,
		FS:              fs,
		PathResolver:    nil,
	})

	terminalCalled := false
	terminalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		terminalCalled = true
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "/components/plain.js", nil)
	rec := httptest.NewRecorder()

	handler := middleware(terminalHandler)
	handler.ServeHTTP(rec, req)

	// Should pass through to terminal handler since no .ts source exists
	if !terminalCalled {
		t.Error("Expected passthrough for plain .js file")
	}
}

// TestTypeScriptMiddleware_BackwardCompatibility tests in-place compilation still works
func TestTypeScriptMiddleware_BackwardCompatibility(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/in-place", "/test")
	cache := transform.NewCache(1024 * 1024)

	middleware := transform.NewTypeScript(transform.TypeScriptConfig{
		WatchDirFunc:    func() string { return "/test" },
		TsconfigRawFunc: func() string { return `{}` },
		Cache:           cache,
		Pool:            nil,
		Logger:          nil,
		Enabled:         true,
		FS:              fs,
		PathResolver:    transform.NewPathResolver("/test", nil, fs, nil), // No explicit mappings, but still handles co-located
	})

	// Request co-located file
	req := httptest.NewRequest("GET", "/components/button.js", nil)
	rec := httptest.NewRecorder()

	terminalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Terminal handler called - should transform co-located .ts file")
	})

	handler := middleware(terminalHandler)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	body := rec.Body.String()
	if len(body) == 0 {
		t.Fatal("Response body is empty")
	}

	// Verify transformation occurred
	if strings.Contains(body, ": string") {
		t.Error("Response still contains TypeScript syntax")
	}
}
