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
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestServeTypeScript_TransformsOnRequest tests that requesting a .js file
// serves the corresponding .ts file transformed to JavaScript
func TestServeTypeScript_TransformsOnRequest(t *testing.T) {
	// Create a temporary directory with a TypeScript file
	tmpDir := t.TempDir()

	// Write a TypeScript file
	tsContent := `export const greeting: string = "Hello, TypeScript!";`
	tsPath := filepath.Join(tmpDir, "test.ts")
	if err := os.WriteFile(tsPath, []byte(tsContent), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	// Create server
	server, err := NewServer(0)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer func() { _ = server.Close() }()

	if err := server.SetWatchDir(tmpDir); err != nil {
		t.Fatalf("Failed to set watch dir: %v", err)
	}

	// Request the .js file (server should serve transformed .ts)
	req := httptest.NewRequest("GET", "/test.js", nil)
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
	tmpDir := t.TempDir()

	// Write a plain .js file (no .ts counterpart)
	jsContent := `export const greeting = "Hello, JavaScript!";`
	jsPath := filepath.Join(tmpDir, "plain.js")
	if err := os.WriteFile(jsPath, []byte(jsContent), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	server, err := NewServer(0)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer func() { _ = server.Close() }()

	if err := server.SetWatchDir(tmpDir); err != nil {
		t.Fatalf("Failed to set watch dir: %v", err)
	}

	// Request the .js file
	req := httptest.NewRequest("GET", "/plain.js", nil)
	w := httptest.NewRecorder()

	server.Handler().ServeHTTP(w, req)

	// Should return 200
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// Should serve the original JS file unchanged
	body := w.Body.String()
	if body != jsContent {
		t.Errorf("Expected original JS content, got: %s", body)
	}
}

// TestServeTypeScript_RejectsPathTraversal tests that path traversal attacks are blocked
func TestServeTypeScript_RejectsPathTraversal(t *testing.T) {
	// Create a parent directory with watch dir as subdirectory
	parentDir := t.TempDir()
	watchDir := filepath.Join(parentDir, "project")
	if err := os.Mkdir(watchDir, 0755); err != nil {
		t.Fatalf("Failed to create watch dir: %v", err)
	}

	// Create legitimate TypeScript in watch directory
	legitimateTs := filepath.Join(watchDir, "legitimate.ts")
	if err := os.WriteFile(legitimateTs, []byte("export const x: number = 42;"), 0644); err != nil {
		t.Fatalf("Failed to write legitimate TS: %v", err)
	}

	// Create a "sensitive" TypeScript file OUTSIDE watch directory but in parent
	secretTs := filepath.Join(parentDir, "secret.ts")
	secretContent := "export const API_KEY: string = 'SECRET123';"
	if err := os.WriteFile(secretTs, []byte(secretContent), 0644); err != nil {
		t.Fatalf("Failed to write secret TS: %v", err)
	}

	server, err := NewServer(0)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer func() { _ = server.Close() }()

	if err := server.SetWatchDir(watchDir); err != nil {
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
