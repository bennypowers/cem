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
	"testing"
)

// TestDemoRouting_BasicRoute verifies demo URLs route correctly
func TestDemoRouting_BasicRoute(t *testing.T) {
	// Create test directory with demo file
	tmpDir := t.TempDir()

	// Create demo HTML with metadata
	demoDir := filepath.Join(tmpDir, "demo")
	err := os.MkdirAll(demoDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create demo directory: %v", err)
	}

	demoHTML := `<!DOCTYPE html>
<html>
<head>
  <meta itemprop="name" content="Basic Demo">
</head>
<body>
  <my-element></my-element>
</body>
</html>`

	demoPath := filepath.Join(demoDir, "basic.html")
	err = os.WriteFile(demoPath, []byte(demoHTML), 0644)
	if err != nil {
		t.Fatalf("Failed to write demo file: %v", err)
	}

	// Create server with demo routing
	server, err := NewServer(0)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	err = server.SetWatchDir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to set watch dir: %v", err)
	}

	// Request demo URL
	req := httptest.NewRequest(http.MethodGet, "/components/my-element/demo/basic-demo/", nil)
	w := httptest.NewRecorder()

	server.Handler().ServeHTTP(w, req)

	// Should return 200 with demo content
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// Should contain demo element
	body := w.Body.String()
	if body == "" {
		t.Error("Expected demo content, got empty response")
	}

	// TODO: More assertions once we implement chrome
}
