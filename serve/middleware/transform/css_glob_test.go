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
	"testing"

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/serve/middleware/transform"
)

// mockErrorBroadcaster implements ErrorBroadcaster interface for testing
type mockErrorBroadcaster struct{}

func (m *mockErrorBroadcaster) BroadcastError(title, message, filename string) {
	// No-op for testing
}

// TestCSSGlobFiltering tests that CSS include/exclude patterns work correctly
func TestCSSGlobFiltering(t *testing.T) {
	tests := []struct {
		name               string
		requestPath        string
		include            []string
		exclude            []string
		expectTransformed  bool
		description        string
	}{
		{
			name:              "no patterns - transforms all CSS",
			requestPath:       "/styles/main.css",
			include:           nil,
			exclude:           nil,
			expectTransformed: true,
			description:       "When no patterns specified, all CSS files are transformed",
		},
		{
			name:              "include pattern matches",
			requestPath:       "/elements/button.css",
			include:           []string{"elements/**/*.css"},
			exclude:           nil,
			expectTransformed: true,
			description:       "File matches include pattern, should be transformed",
		},
		{
			name:              "include pattern doesn't match",
			requestPath:       "/docs/styles.css",
			include:           []string{"elements/**/*.css"},
			exclude:           nil,
			expectTransformed: false,
			description:       "File doesn't match include pattern, should not be transformed",
		},
		{
			name:              "exclude pattern matches",
			requestPath:       "/docs/styles.css",
			include:           nil,
			exclude:           []string{"docs/**/*.css"},
			expectTransformed: false,
			description:       "File matches exclude pattern, should not be transformed",
		},
		{
			name:              "exclude pattern doesn't match",
			requestPath:       "/elements/button.css",
			include:           nil,
			exclude:           []string{"docs/**/*.css"},
			expectTransformed: true,
			description:       "File doesn't match exclude pattern, should be transformed",
		},
		{
			name:              "include matches, exclude doesn't - transforms",
			requestPath:       "/elements/button.css",
			include:           []string{"elements/**/*.css"},
			exclude:           []string{"docs/**/*.css"},
			expectTransformed: true,
			description:       "File matches include and not exclude, should be transformed",
		},
		{
			name:              "both include and exclude match - exclude wins",
			requestPath:       "/elements/internal.css",
			include:           []string{"elements/**/*.css"},
			exclude:           []string{"**/*internal.css"},
			expectTransformed: false,
			description:       "File matches both include and exclude, exclude takes precedence",
		},
		{
			name:              "multiple include patterns - any match",
			requestPath:       "/components/card.css",
			include:           []string{"elements/**/*.css", "components/**/*.css"},
			exclude:           nil,
			expectTransformed: true,
			description:       "File matches one of multiple include patterns",
		},
		{
			name:              "simple glob pattern",
			requestPath:       "/button.css",
			include:           []string{"*.css"},
			exclude:           nil,
			expectTransformed: true,
			description:       "Simple glob pattern matches file in root",
		},
		{
			name:              "nested path with doublestar",
			requestPath:       "/deep/nested/path/styles.css",
			include:           []string{"**/styles.css"},
			exclude:           nil,
			expectTransformed: true,
			description:       "Doublestar glob matches deeply nested file",
		},
		{
			name:              "empty string in include patterns",
			requestPath:       "/elements/button.css",
			include:           []string{"", "elements/**/*.css"},
			exclude:           nil,
			expectTransformed: true,
			description:       "Empty pattern in array is ignored, other patterns apply",
		},
		{
			name:              "empty string in exclude patterns",
			requestPath:       "/elements/button.css",
			include:           nil,
			exclude:           []string{"", "docs/**/*.css"},
			expectTransformed: true,
			description:       "Empty pattern in exclude array is ignored, other patterns apply",
		},
		{
			name:              "very deeply nested path",
			requestPath:       "/a/b/c/d/e/f/g/h/i/j/k/l/m/n/o/p/q/r/s/t/u/v/w/x/y/z/deep.css",
			include:           []string{"**/deep.css"},
			exclude:           nil,
			expectTransformed: true,
			description:       "Doublestar pattern matches very deeply nested file (26 levels)",
		},
		{
			name:              "root-level CSS with include pattern",
			requestPath:       "/styles.css",
			include:           []string{"*.css"},
			exclude:           nil,
			expectTransformed: true,
			description:       "Simple glob pattern matches root-level CSS file",
		},
		{
			name:              "root-level CSS excluded",
			requestPath:       "/styles.css",
			include:           nil,
			exclude:           []string{"*.css"},
			expectTransformed: false,
			description:       "Root-level CSS file can be excluded",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create in-memory filesystem
			mfs := platform.NewMapFileSystem(nil)
			mfs.AddFile("/test-root"+tt.requestPath, ":host { color: red; }", 0644)

			// Create middleware with test config
			middleware := transform.NewCSS(transform.CSSConfig{
				WatchDirFunc:     func() string { return "/test-root" },
				Logger:           &mockLogger{},
				ErrorBroadcaster: &mockErrorBroadcaster{},
				ConfigFile:       "cem.config.yaml",
				Enabled:          true,
				Include:          tt.include,
				Exclude:          tt.exclude,
				FS:               mfs,
			})

			// Create test handler that the middleware wraps
			nextCalled := false
			next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				nextCalled = true
				// Serve original CSS (not transformed)
				w.Header().Set("Content-Type", "text/css; charset=utf-8")
				_, _ = w.Write([]byte(":host { color: red; }"))
			})

			// Wrap with middleware
			handler := middleware(next)

			// Make request
			req := httptest.NewRequest("GET", tt.requestPath, nil)
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)

			// Check if transformation happened
			isJavaScript := rec.Header().Get("Content-Type") == "application/javascript; charset=utf-8"

			if tt.expectTransformed {
				if nextCalled {
					t.Errorf("Expected CSS to be transformed, but next handler was called")
					t.Logf("Description: %s", tt.description)
				}
				if !isJavaScript {
					t.Errorf("Expected JavaScript content-type, got %s", rec.Header().Get("Content-Type"))
					t.Logf("Description: %s", tt.description)
				}
			} else {
				if !nextCalled {
					t.Errorf("Expected CSS to be passed through, but middleware transformed it")
					t.Logf("Description: %s", tt.description)
				}
				if isJavaScript {
					t.Errorf("Expected original CSS content-type, got JavaScript")
					t.Logf("Description: %s", tt.description)
				}
			}
		})
	}
}

// TestCSSGlobFiltering_NonexistentFile tests that non-existent files are passed to next handler
func TestCSSGlobFiltering_NonexistentFile(t *testing.T) {
	// Create in-memory filesystem WITHOUT the requested file
	mfs := platform.NewMapFileSystem(nil)
	// Intentionally not adding /test-root/missing.css

	middleware := transform.NewCSS(transform.CSSConfig{
		WatchDirFunc:     func() string { return "/test-root" },
		Logger:           &mockLogger{},
		ErrorBroadcaster: &mockErrorBroadcaster{},
		ConfigFile:       "cem.config.yaml",
		Enabled:          true,
		Include:          []string{"**/*.css"}, // Pattern would match if file existed
		Exclude:          nil,
		FS:               mfs,
	})

	nextCalled := false
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCalled = true
		http.NotFound(w, r)
	})

	handler := middleware(next)

	req := httptest.NewRequest("GET", "/missing.css", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if !nextCalled {
		t.Error("Expected next handler to be called for non-existent file")
	}

	if rec.Code != http.StatusNotFound {
		t.Errorf("Expected 404 status for non-existent file, got %d", rec.Code)
	}
}

// TestCSSGlobFiltering_E2E tests glob filtering with full server integration
func TestCSSGlobFiltering_E2E(t *testing.T) {
	// Create in-memory filesystem with multiple CSS files
	mfs := platform.NewMapFileSystem(nil)
	mfs.AddFile("/test/elements/button.css", ":host { color: blue; }", 0644)
	mfs.AddFile("/test/docs/styles.css", "body { margin: 0; }", 0644)
	mfs.AddFile("/test/internal.css", ".private { display: none; }", 0644)

	// Create middleware that only transforms elements/ CSS
	middleware := transform.NewCSS(transform.CSSConfig{
		WatchDirFunc:     func() string { return "/test" },
		Logger:           &mockLogger{},
		ErrorBroadcaster: &mockErrorBroadcaster{},
		ConfigFile:       "cem.config.yaml",
		Enabled:          true,
		Include:          []string{"elements/**/*.css"},
		Exclude:          nil,
		FS:               mfs,
	})

	// Next handler (fallback for non-transformed files)
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("/* not transformed */"))
	})

	handler := middleware(next)

	tests := []struct {
		path                string
		expectTransformed   bool
	}{
		{"/elements/button.css", true},
		{"/docs/styles.css", false},
		{"/internal.css", false},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.path, nil)
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)

			isJavaScript := rec.Header().Get("Content-Type") == "application/javascript; charset=utf-8"

			if tt.expectTransformed != isJavaScript {
				t.Errorf("Path %s: expected transformed=%v, got JavaScript=%v",
					tt.path, tt.expectTransformed, isJavaScript)
			}
		})
	}
}
