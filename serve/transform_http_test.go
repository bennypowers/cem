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

package serve_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/serve"
)

// TestServeTypeScript_TransformsOnRequest tests that requesting a .js file
// serves the corresponding .ts file transformed to JavaScript
func TestServeTypeScript_TransformsOnRequest(t *testing.T) {
	// Load TypeScript fixture into in-memory filesystem
	mfs := testutil.NewFixtureFS(t, "transforms/http-typescript", "/test")

	// Create server with MapFileSystem
	server, err := serve.NewServerWithConfig(serve.Config{
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

	// Request the .js file (server should serve transformed .ts from fixture)
	req := httptest.NewRequest("GET", "/simple-greeting.js", nil)
	w := httptest.NewRecorder()

	server.Handler().ServeHTTP(w, req)

	// Should return 200
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// Should have JavaScript content type
	contentType := w.Header().Get("Content-Type")
	if !strings.Contains(contentType, "application/javascript") {
		t.Errorf("Expected JavaScript content type, got %s", contentType)
	}

	// Should have transformed output (no TypeScript type annotations)
	body := w.Body.String()
	if strings.Contains(body, ": string") {
		t.Error("Output still contains TypeScript type annotation ': string'")
	}

	// Should contain the actual code
	if !strings.Contains(body, "greeting") {
		t.Error("Output missing variable 'greeting'")
	}

	// Should have inline source map
	if !strings.Contains(body, "//# sourceMappingURL=") {
		t.Error("Output missing inline source map")
	}
}

// TestServeTypeScript_OnlyWhenTsExists tests that .js files are served normally
// when no corresponding .ts file exists
func TestServeTypeScript_OnlyWhenTsExists(t *testing.T) {
	// Load plain JavaScript fixture (no .ts counterpart)
	mfs := testutil.NewFixtureFS(t, "transforms/http-typescript", "/test")

	server, err := serve.NewServerWithConfig(serve.Config{
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

	// Request the .js file from fixture
	req := httptest.NewRequest("GET", "/plain-javascript.js", nil)
	w := httptest.NewRecorder()

	server.Handler().ServeHTTP(w, req)

	// Should return 200
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// Should serve the original JS file unchanged
	body := w.Body.String()
	expectedContent := "export const greeting = \"Hello, JavaScript!\";\n"
	if body != expectedContent {
		t.Errorf("Expected original JS content, got: %s", body)
	}
}

// TestServeTypeScript_RejectsPathTraversal tests that path traversal attacks are blocked
func TestServeTypeScript_RejectsPathTraversal(t *testing.T) {
	// Load path-traversal fixtures:
	// Fixture structure creates:
	//   /parent/project/legitimate.ts (inside watch dir)
	//   /parent/secret.ts (outside watch dir)
	mfs := testutil.NewFixtureFS(t, "transforms/http-typescript/path-traversal", "/parent")

	server, err := serve.NewServerWithConfig(serve.Config{
		Port:   0,
		Reload: true,
		FS:     mfs,
	})
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer func() { _ = server.Close() }()

	// Set watch dir to /parent/project (only legitimate.ts should be accessible)
	if err := server.SetWatchDir("/parent/project"); err != nil {
		t.Fatalf("Failed to set watch dir: %v", err)
	}

	// Path traversal to access file in parent directory
	// Request .js but the server should check for .ts
	attacks := []string{
		"/../secret.js",                     // Direct parent access (.js request)
		"/../secret.ts",                     // Direct parent access (.ts request)
		"/subdir/../../secret.js",           // Via non-existent subdir
		"/./../secret.ts",                   // Mixed with current dir
		"/./../../project/../secret.js",     // Complex traversal
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
			if strings.Contains(body, "SECRET123") || strings.Contains(body, "API_KEY") {
				t.Errorf("Path traversal succeeded! Secret content leaked: %s", body)
			}
		})
	}

	// Verify legitimate access still works
	t.Run("legitimate access", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/legitimate.js", nil)
		w := httptest.NewRecorder()

		server.Handler().ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status 200 for legitimate request, got %d", w.Code)
		}

		body := w.Body.String()
		if !strings.Contains(body, "42") {
			t.Error("Legitimate TypeScript file not served correctly")
		}
	})
}
