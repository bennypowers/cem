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
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestRootListing_ShowsAllElements verifies GET / shows element listing
func TestRootListing_ShowsAllElements(t *testing.T) {
	tmpDir := t.TempDir()

	// Create manifest with multiple elements and demos
	manifest := map[string]interface{}{
		"schemaVersion": "1.0.0",
		"modules": []interface{}{
			map[string]interface{}{
				"path": "src/button/button.ts",
				"declarations": []interface{}{
					map[string]interface{}{
						"name":    "MyButton",
						"tagName": "my-button",
						"demos": []interface{}{
							map[string]interface{}{
								"name": "Basic",
								"url":  "./demo/basic.html",
							},
							map[string]interface{}{
								"name": "Variants",
								"url":  "./demo/variants.html",
							},
						},
					},
				},
			},
			map[string]interface{}{
				"path": "src/card/card.ts",
				"declarations": []interface{}{
					map[string]interface{}{
						"name":    "MyCard",
						"tagName": "my-card",
						"demos": []interface{}{
							map[string]interface{}{
								"name": "Simple",
								"url":  "./demo/simple.html",
							},
						},
					},
				},
			},
		},
	}

	manifestBytes, _ := json.Marshal(manifest)
	manifestPath := filepath.Join(tmpDir, "custom-elements.json")
	err := os.WriteFile(manifestPath, manifestBytes, 0644)
	if err != nil {
		t.Fatalf("Failed to write manifest: %v", err)
	}

	server, err := NewServer(0)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	err = server.SetWatchDir(tmpDir)
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

	// Should link to demo URLs
	if !strings.Contains(body, "/components/my-button/demo/") {
		t.Error("Expected demo URL for my-button")
	}

	if !strings.Contains(body, "/components/my-card/demo/") {
		t.Error("Expected demo URL for my-card")
	}
}

// TestRootListing_EmptyManifest verifies listing with no elements
func TestRootListing_EmptyManifest(t *testing.T) {
	tmpDir := t.TempDir()

	manifest := map[string]interface{}{
		"schemaVersion": "1.0.0",
		"modules":       []interface{}{},
	}

	manifestBytes, _ := json.Marshal(manifest)

	server, err := NewServer(0)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	err = server.SetWatchDir(tmpDir)
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
	if !strings.Contains(body, "No elements found") && !strings.Contains(body, "no demos") {
		t.Error("Expected message about no elements found")
	}
}

// TestRootListing_SortedByElement verifies alphabetical sorting
func TestRootListing_SortedByElement(t *testing.T) {
	tmpDir := t.TempDir()

	manifest := map[string]interface{}{
		"schemaVersion": "1.0.0",
		"modules": []interface{}{
			map[string]interface{}{
				"path": "src/zebra.ts",
				"declarations": []interface{}{
					map[string]interface{}{
						"tagName": "my-zebra",
						"demos": []interface{}{
							map[string]interface{}{"name": "Demo", "url": "./demo.html"},
						},
					},
				},
			},
			map[string]interface{}{
				"path": "src/apple.ts",
				"declarations": []interface{}{
					map[string]interface{}{
						"tagName": "my-apple",
						"demos": []interface{}{
							map[string]interface{}{"name": "Demo", "url": "./demo.html"},
						},
					},
				},
			},
			map[string]interface{}{
				"path": "src/mango.ts",
				"declarations": []interface{}{
					map[string]interface{}{
						"tagName": "my-mango",
						"demos": []interface{}{
							map[string]interface{}{"name": "Demo", "url": "./demo.html"},
						},
					},
				},
			},
		},
	}

	manifestBytes, _ := json.Marshal(manifest)

	server, err := NewServer(0)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	err = server.SetWatchDir(tmpDir)
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
	if !(applePos < mangoPos && mangoPos < zebraPos) {
		t.Errorf("Expected alphabetical order: apple(%d) < mango(%d) < zebra(%d)",
			applePos, mangoPos, zebraPos)
	}
}
