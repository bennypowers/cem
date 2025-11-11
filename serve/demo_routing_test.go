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

// TestDemoRouting_BasicRoute verifies demo URLs route correctly
func TestDemoRouting_BasicRoute(t *testing.T) {
	// Load fixtures from testdata into MapFileSystem
	mfs := testutil.NewFixtureFS(t, "demo-routing", "/test")
	manifestBytes := testutil.LoadFixtureFile(t, "demo-routing/manifest.json")

	// Create server with MapFileSystem
	server, err := serve.NewServerWithConfig(serve.Config{
		Port:   0,
		Reload: true,
		FS:     mfs,
	})
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer func() {
		if err := server.Close(); err != nil {
			t.Errorf("Failed to close server: %v", err)
		}
	}()

	err = server.SetWatchDir("/test")
	if err != nil {
		t.Fatalf("Failed to set watch dir: %v", err)
	}

	err = server.SetManifest(manifestBytes)
	if err != nil {
		t.Fatalf("Failed to set manifest: %v", err)
	}

	// Request demo URL (matching the URL from manifest: ./demo/basic.html)
	req := httptest.NewRequest(http.MethodGet, "/demo/basic.html", nil)
	w := httptest.NewRecorder()

	server.Handler().ServeHTTP(w, req)

	// Should return 200 with demo content
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	body := w.Body.String()

	// Should contain chrome wrapper
	if !strings.Contains(body, "<cem-serve-chrome") {
		t.Error("Expected cem-serve-chrome element")
	}

	// Should contain demo content
	if !strings.Contains(body, "my-element") {
		t.Error("Expected demo element in response")
	}

	if !strings.Contains(body, "Hello World") {
		t.Error("Expected demo content in response")
	}
}
