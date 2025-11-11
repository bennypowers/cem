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

	"bennypowers.dev/cem/internal/platform"
)

// newListingTestFS creates an in-memory filesystem for listing tests
func newListingTestFS() platform.FileSystem {
	mfs := platform.NewMapFileSystem(nil)
	// Add minimal directory structure
	mfs.AddFile("/test/package.json", `{"name":"test"}`, 0644)
	return mfs
}

// TestRootListing_ShowsAllElements verifies GET / shows element listing
func TestRootListing_ShowsAllElements(t *testing.T) {
	// Load manifest fixture
	manifestBytes, err := os.ReadFile(filepath.Join("testdata", "listing", "multiple-elements.json"))
	if err != nil {
		t.Fatalf("Failed to read manifest fixture: %v", err)
	}

	mfs := newListingTestFS()
	server, err := NewServerWithConfig(Config{
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

	// Request root
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	server.Handler().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	body := w.Body.String()

	// Should contain listing chrome
	if !strings.Contains(body, "<cem-serve-listing") {
		t.Error("Expected cem-serve-listing element")
	}

	// Should be grouped by element (sidebar nav)
	if !strings.Contains(body, "my-button") {
		t.Error("Expected my-button in listing")
	}

	if !strings.Contains(body, "my-card") {
		t.Error("Expected my-card in listing")
	}

	// Should show demo links
	if !strings.Contains(body, "Basic") {
		t.Error("Expected 'Basic' demo link")
	}

	if !strings.Contains(body, "Variants") {
		t.Error("Expected 'Variants' demo link")
	}

	if !strings.Contains(body, "Simple") {
		t.Error("Expected 'Simple' demo link")
	}

	// Should link to demo URLs (from manifest: ./demo/basic.html, etc)
	if !strings.Contains(body, "/demo/basic.html") {
		t.Error("Expected demo URL /demo/basic.html")
	}

	if !strings.Contains(body, "/demo/simple.html") {
		t.Error("Expected demo URL /demo/simple.html")
	}
}

// TestRootListing_EmptyManifest verifies listing with no elements
func TestRootListing_EmptyManifest(t *testing.T) {
	// Load empty manifest fixture
	manifestBytes, err := os.ReadFile(filepath.Join("testdata", "listing", "empty-manifest.json"))
	if err != nil {
		t.Fatalf("Failed to read manifest fixture: %v", err)
	}

	mfs := newListingTestFS()
	server, err := NewServerWithConfig(Config{
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

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	server.Handler().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	body := w.Body.String()

	// Should show helpful message
	if !strings.Contains(body, "No demos found") {
		t.Error("Expected message about no demos found")
	}
}

// TestRootListing_SortedByElement verifies alphabetical sorting
func TestRootListing_SortedByElement(t *testing.T) {
	// Load sorting test fixture
	manifestBytes, err := os.ReadFile(filepath.Join("testdata", "listing", "sorting-test.json"))
	if err != nil {
		t.Fatalf("Failed to read manifest fixture: %v", err)
	}

	mfs := newListingTestFS()
	server, err := NewServerWithConfig(Config{
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

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	server.Handler().ServeHTTP(w, req)

	body := w.Body.String()

	// Find positions of elements
	applePos := strings.Index(body, "my-apple")
	mangoPos := strings.Index(body, "my-mango")
	zebraPos := strings.Index(body, "my-zebra")

	if applePos == -1 || mangoPos == -1 || zebraPos == -1 {
		t.Fatal("Expected all elements in output")
	}

	// Should be in alphabetical order
	if applePos >= mangoPos || mangoPos >= zebraPos {
		t.Errorf("Expected alphabetical order: apple(%d) < mango(%d) < zebra(%d)",
			applePos, mangoPos, zebraPos)
	}
}
