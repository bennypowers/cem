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
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
)

// TestTransformCSS_Basic tests basic CSS to constructable stylesheet transformation
func TestTransformCSS_Basic(t *testing.T) {
	cssContent := `:host {
  display: block;
  color: var(--color, red);
}`

	result := TransformCSS([]byte(cssContent), "test.css")

	// Should wrap in CSSStyleSheet
	if !strings.Contains(result, "new CSSStyleSheet()") {
		t.Error("Output missing CSSStyleSheet constructor")
	}

	// Should use replaceSync
	if !strings.Contains(result, "replaceSync") {
		t.Error("Output missing replaceSync call")
	}

	// Should export default
	if !strings.Contains(result, "export default") {
		t.Error("Output missing default export")
	}

	// Should contain the original CSS
	if !strings.Contains(result, ":host") {
		t.Error("Output missing original CSS content")
	}

	// Should escape backticks in CSS
	cssWithBacktick := "content: `test`;"
	escaped := TransformCSS([]byte(cssWithBacktick), "test.css")
	if strings.Contains(escaped, "content: `test`;") {
		t.Error("Backticks in CSS not escaped")
	}

	// Should only escape ${ not all $ (matches Lit behavior)
	cssWithDollar := "width: $var; height: ${expr};"
	dollarEscaped := TransformCSS([]byte(cssWithDollar), "test.css")
	// $var should NOT be escaped
	if !strings.Contains(dollarEscaped, "$var") {
		t.Error("Single $ incorrectly escaped (should only escape ${)")
	}
	// ${ should be escaped
	if strings.Contains(dollarEscaped, "${expr}") && !strings.Contains(dollarEscaped, "\\${") {
		t.Error("${ not properly escaped")
	}

	// Should escape </ to prevent script tag injection
	cssWithScriptTag := "content: '</script>';"
	scriptEscaped := TransformCSS([]byte(cssWithScriptTag), "test.css")
	if strings.Contains(scriptEscaped, "</script>") && !strings.Contains(scriptEscaped, "<\\/script>") {
		t.Error("</ not properly escaped")
	}
}

// TestServCSS_TransformsToModule tests HTTP serving of CSS as JavaScript module
func TestServeCSS_TransformsToModule(t *testing.T) {
	// Load CSS fixture into in-memory filesystem
	mfs := testutil.NewFixtureFS(t, "transforms/http-css", "/test")

	server, err := NewServerWithConfig(Config{
		Port:   0,
		Reload: true,
		FS:     mfs,
	})
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer func() { _ = server.Close() }()

	if err := server.SetWatchDir("/test"); err != nil {
		t.Fatalf("Failed to set watch dir: %v", err)
	}

	// Request the CSS file (server should serve transformed CSS from fixture)
	req := httptest.NewRequest("GET", "/simple-host.css", nil)
	w := httptest.NewRecorder()

	server.Handler().ServeHTTP(w, req)

	// Should return 200
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// Should have JavaScript content type (not text/css)
	contentType := w.Header().Get("Content-Type")
	if !strings.Contains(contentType, "application/javascript") {
		t.Errorf("Expected JavaScript content type, got %s", contentType)
	}

	// Should be transformed to constructable stylesheet
	body := w.Body.String()
	if !strings.Contains(body, "CSSStyleSheet") {
		t.Error("CSS not transformed to CSSStyleSheet module")
	}

	// Should contain original CSS
	if !strings.Contains(body, ":host") {
		t.Error("Transformed output missing original CSS")
	}
}

// TestServeCSS_WarnsWithoutImportAttribute tests warning for CSS imports without 'with { type: "css" }'
func TestServeCSS_WarnsWithoutImportAttribute(t *testing.T) {
	t.Skip("Warning implementation TBD - need to detect import context")
	// This test verifies that when CSS is imported without `with { type: 'css' }`,
	// the server logs a warning. Implementation requires examining the referrer
	// to detect how the CSS is being imported.
}

// TestServeCSS_RejectsPathTraversal tests that path traversal attacks are blocked
func TestServeCSS_RejectsPathTraversal(t *testing.T) {
	// Load path-traversal fixtures:
	// Fixture structure creates:
	//   /parent/legitimate.css (inside watch dir)
	//   /parent/secret.css (outside watch dir, at parent level)
	mfs := testutil.NewFixtureFS(t, "transforms/http-css/path-traversal", "/parent")

	server, err := NewServerWithConfig(Config{
		Port:   0,
		Reload: true,
		FS:     mfs,
	})
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer func() { _ = server.Close() }()

	// Set watch dir to /parent (both files are in this dir, so we test same-level access control)
	// Note: Unlike TypeScript test, CSS fixtures are in same directory
	if err := server.SetWatchDir("/parent"); err != nil {
		t.Fatalf("Failed to set watch dir: %v", err)
	}

	// Path traversal to access file in parent directory
	attacks := []string{
		"/../secret.css",                    // Direct parent access
		"/subdir/../../secret.css",          // Via non-existent subdir
		"/./../secret.css",                  // Mixed with current dir
		"/./../../parent/../secret.css",     // Complex traversal
	}

	for _, attack := range attacks {
		t.Run(attack, func(t *testing.T) {
			req := httptest.NewRequest("GET", attack, nil)
			w := httptest.NewRecorder()

			server.Handler().ServeHTTP(w, req)

			// Should return 404, not expose the secret file
			if w.Code != http.StatusNotFound {
				t.Errorf("Expected status 404 for path traversal, got %d", w.Code)
			}

			// Should NOT contain secret content
			body := w.Body.String()
			if strings.Contains(body, "SECRET") {
				t.Errorf("Path traversal succeeded! Secret content leaked: %s", body)
			}
		})
	}

	// Verify legitimate access still works
	t.Run("legitimate access", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/legitimate.css", nil)
		w := httptest.NewRecorder()

		server.Handler().ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status 200 for legitimate request, got %d", w.Code)
		}

		body := w.Body.String()
		if !strings.Contains(body, "color: blue") {
			t.Error("Legitimate CSS file not served correctly")
		}
	})
}
