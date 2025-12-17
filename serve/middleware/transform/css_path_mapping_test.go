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

	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/serve/middleware/transform"
)

// TestCSSMiddleware_SrcDistPattern tests CSS file resolution with src/dist separation
func TestCSSMiddleware_SrcDistPattern(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/src-dist", "/test")

	// Add a CSS file in src/
	fs.AddFile("/test/src/components/button.css", `
.button {
	background: blue;
	color: white;
}
`, 0644)

	// Create middleware with import attributes (always transforms)
	middleware := transform.NewCSS(transform.CSSConfig{
		WatchDirFunc: func() string { return "/test" },
		Logger:       nil,
		Enabled:      false, // Disabled, but import attributes should still work
		FS:           fs,
		PathMappings: map[string]string{
			"/dist/": "/src/",
		},
	})

	// Request /dist/components/button.css with import attributes
	// Should resolve to /src/components/button.css
	req := httptest.NewRequest("GET", "/dist/components/button.css?__cem-import-attrs[type]=css", nil)
	rec := httptest.NewRecorder()

	terminalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Terminal handler called - CSS middleware should have handled request")
	})

	handler := middleware(terminalHandler)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	contentType := rec.Header().Get("Content-Type")
	if !strings.Contains(contentType, "application/javascript") {
		t.Errorf("Expected Content-Type to contain 'application/javascript', got %s", contentType)
	}

	body := rec.Body.String()
	if len(body) == 0 {
		t.Fatal("Response body is empty")
	}

	// Should be a JavaScript module with CSS string
	if !strings.Contains(body, "export default") {
		t.Error("Response should be a JavaScript module with 'export default'")
	}
}

// TestCSSMiddleware_WorkspaceMode tests CSS resolution with package prefix
func TestCSSMiddleware_WorkspaceMode(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/workspace-mode", "/test")

	// Add a CSS file in workspace package src/
	fs.AddFile("/test/packages/webawesome/src/components/alert/alert.css", `
.alert {
	padding: 1rem;
	border-radius: 4px;
}
`, 0644)

	middleware := transform.NewCSS(transform.CSSConfig{
		WatchDirFunc: func() string { return "/test" },
		Logger:       nil,
		Enabled:      false,
		FS:           fs,
		PathMappings: map[string]string{
			"/packages/webawesome/dist/": "/packages/webawesome/src/",
		},
	})

	// Request with workspace path: /packages/webawesome/dist/components/alert/alert.css
	// Should resolve to: /packages/webawesome/src/components/alert/alert.css
	req := httptest.NewRequest("GET", "/packages/webawesome/dist/components/alert/alert.css?__cem-import-attrs[type]=css", nil)
	rec := httptest.NewRecorder()

	terminalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Terminal handler called - should transform")
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

	if !strings.Contains(body, ".alert") {
		t.Error("Transformed CSS should contain .alert class")
	}
}

// TestCSSMiddleware_ExplicitPathMappings tests user-configured CSS path mappings
func TestCSSMiddleware_ExplicitPathMappings(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/explicit-mappings", "/test")

	// Add CSS file in sources/
	fs.AddFile("/test/sources/styles/theme.css", `
:root {
	--primary-color: blue;
}
`, 0644)

	middleware := transform.NewCSS(transform.CSSConfig{
		WatchDirFunc: func() string { return "/test" },
		Logger:       nil,
		Enabled:      false,
		FS:           fs,
		PathMappings: map[string]string{
			"/lib/": "/sources/",
		},
	})

	// Request using custom mapping: /lib/styles/theme.css -> /sources/styles/theme.css
	req := httptest.NewRequest("GET", "/lib/styles/theme.css?__cem-import-attrs[type]=css", nil)
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
	if !strings.Contains(body, "--primary-color") {
		t.Error("Transformed CSS should contain CSS variable")
	}
}

// TestCSSMiddleware_CoLocated tests that co-located CSS files still work
func TestCSSMiddleware_CoLocated(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/in-place", "/test")

	// Add CSS file in same directory as components
	fs.AddFile("/test/components/button.css", `
.button { display: block; }
`, 0644)

	middleware := transform.NewCSS(transform.CSSConfig{
		WatchDirFunc: func() string { return "/test" },
		Logger:       nil,
		Enabled:      false,
		FS:           fs,
		PathMappings: nil,
	})

	req := httptest.NewRequest("GET", "/components/button.css?__cem-import-attrs[type]=css", nil)
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
	if !strings.Contains(body, "display: block") {
		t.Error("Should contain CSS content")
	}
}

// TestCSSMiddleware_NotFound tests 404 when CSS file doesn't exist
func TestCSSMiddleware_NotFound(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/src-dist", "/test")

	middleware := transform.NewCSS(transform.CSSConfig{
		WatchDirFunc: func() string { return "/test" },
		Logger:       nil,
		Enabled:      false,
		FS:           fs,
		PathMappings: nil,
	})

	// Request non-existent file
	req := httptest.NewRequest("GET", "/dist/components/nonexistent.css?__cem-import-attrs[type]=css", nil)
	rec := httptest.NewRecorder()

	terminalCalled := false
	terminalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		terminalCalled = true
		w.WriteHeader(http.StatusNotFound)
	})

	handler := middleware(terminalHandler)
	handler.ServeHTTP(rec, req)

	// Should pass through to terminal handler (which returns 404)
	if !terminalCalled {
		t.Error("Terminal handler should be called for non-existent files")
	}
}

// TestCSSMiddleware_WithoutImportAttributes tests that CSS without import attrs is not transformed when disabled
func TestCSSMiddleware_WithoutImportAttributes(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/src-dist", "/test")
	fs.AddFile("/test/src/components/button.css", `.button { color: red; }`, 0644)

	middleware := transform.NewCSS(transform.CSSConfig{
		WatchDirFunc: func() string { return "/test" },
		Logger:       nil,
		Enabled:      false, // Disabled and no import attributes
		FS:           fs,
		PathMappings: nil,
	})

	// Request without import attributes - should pass through
	req := httptest.NewRequest("GET", "/dist/components/button.css", nil)
	rec := httptest.NewRecorder()

	terminalCalled := false
	terminalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		terminalCalled = true
		w.WriteHeader(http.StatusOK)
	})

	handler := middleware(terminalHandler)
	handler.ServeHTTP(rec, req)

	if !terminalCalled {
		t.Error("Should pass through when disabled and no import attributes")
	}
}

// TestCSSMiddleware_EnabledWithIncludePattern tests CSS transforms with include patterns
func TestCSSMiddleware_EnabledWithIncludePattern(t *testing.T) {
	fs := testutil.NewFixtureFS(t, "path-mappings/src-dist", "/test")
	fs.AddFile("/test/src/styles/global.css", `.global { margin: 0; }`, 0644)

	middleware := transform.NewCSS(transform.CSSConfig{
		WatchDirFunc: func() string { return "/test" },
		Logger:       nil,
		Enabled:      true,
		Include:      []string{"**/*.css"},
		FS:           fs,
		PathMappings: map[string]string{
			"/dist/": "/src/",
		},
	})

	// Request without import attributes, but should transform due to enabled + include pattern
	req := httptest.NewRequest("GET", "/dist/styles/global.css", nil)
	rec := httptest.NewRecorder()

	terminalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Should transform when enabled with matching include pattern")
	})

	handler := middleware(terminalHandler)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rec.Code)
	}

	body := rec.Body.String()
	if !strings.Contains(body, "export default") {
		t.Error("Should be transformed to JavaScript module")
	}
}
