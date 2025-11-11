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
	"path/filepath"
	"slices"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
)

// TestGetAffectedPageURLs_ImportedFileChange tests that when an imported file changes,
// the demo page that imports it is marked as affected
func TestGetAffectedPageURLs_ImportedFileChange(t *testing.T) {
	// Load fixtures into in-memory filesystem
	mfs := testutil.NewFixtureFS(t, "smart-reload", "/test")
	manifestBytes := testutil.LoadFixtureFile(t, "smart-reload/manifest.json")

	// Create server with fixtures
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

	if err := server.SetManifest(manifestBytes); err != nil {
		t.Fatalf("Failed to set manifest: %v", err)
	}

	// Simulate a change to the imported JS file
	changedFile := filepath.Join("/test", "demo", "my-element.js")
	affectedPages := server.getAffectedPageURLs(changedFile, nil)

	// The demo page should be affected because it imports my-element.js
	if len(affectedPages) == 0 {
		t.Fatal("Expected demo page to be affected when imported file changes")
	}

	expectedRoute := "/demo/basic.html"
	found := slices.Contains(affectedPages, expectedRoute)

	if !found {
		t.Errorf("Expected affected pages to include %s, got: %v", expectedRoute, affectedPages)
	}
}

// TestGetAffectedPageURLs_HTMLFileChange tests that when a demo HTML file itself changes,
// that demo page is marked as affected
func TestGetAffectedPageURLs_HTMLFileChange(t *testing.T) {
	// Load fixtures into in-memory filesystem
	mfs := testutil.NewFixtureFS(t, "smart-reload", "/test")
	manifestBytes := testutil.LoadFixtureFile(t, "smart-reload/manifest.json")

	// Create server with fixtures
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

	if err := server.SetManifest(manifestBytes); err != nil {
		t.Fatalf("Failed to set manifest: %v", err)
	}

	// Simulate a change to the HTML file itself
	changedFile := filepath.Join("/test", "demo", "basic.html")
	affectedPages := server.getAffectedPageURLs(changedFile, nil)

	// The demo page should be affected because it IS the changed file
	if len(affectedPages) == 0 {
		t.Fatal("Expected demo page to be affected when HTML file changes")
	}

	expectedRoute := "/demo/basic.html"
	found := slices.Contains(affectedPages, expectedRoute)

	if !found {
		t.Errorf("Expected affected pages to include %s, got: %v", expectedRoute, affectedPages)
	}
}
