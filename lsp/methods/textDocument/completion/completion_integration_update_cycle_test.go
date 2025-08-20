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
package completion_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	G "bennypowers.dev/cem/generate"
	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/methods/textDocument/completion"
	"bennypowers.dev/cem/lsp/testhelpers"
	M "bennypowers.dev/cem/manifest"
	W "bennypowers.dev/cem/workspace"
)

// TestCompletionUpdateCycle tests the complete end-to-end cycle:
// 1. Initial TypeScript source file with union type
// 2. Generate manifest with initial completions
// 3. Edit source file to add new union value
// 4. Verify generate watcher detects changes
// 5. Verify completions update with new value
func TestCompletionUpdateCycle(t *testing.T) {
	// Create a temporary workspace directory
	tempDir := t.TempDir()
	fixtureDir := filepath.Join("..", "..", "..", "test", "fixtures", "completion-update-cycle")

	// Copy fixture files to temp directory
	CopyFixtureFile(t, fixtureDir, "package.json", tempDir, "package.json")

	// Create .config directory and copy cem.yaml
	configDir := filepath.Join(tempDir, ".config")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		t.Fatalf("Failed to create config directory: %v", err)
	}
	CopyFixtureFile(t, fixtureDir, "cem.yaml", configDir, "cem.yaml")

	// Create elements directory structure
	elementsDir := filepath.Join(tempDir, "elements", "test-element")
	if err := os.MkdirAll(elementsDir, 0755); err != nil {
		t.Fatalf("Failed to create elements directory: %v", err)
	}

	// Copy initial TypeScript element file
	CopyFixtureFile(t, fixtureDir, "initial-element.ts", elementsDir, "test-element.ts")

	// Create dist directory
	distDir := filepath.Join(tempDir, "dist")
	if err := os.MkdirAll(distDir, 0755); err != nil {
		t.Fatalf("Failed to create dist directory: %v", err)
	}

	// Create workspace and generate initial manifest
	workspace := W.NewFileSystemWorkspaceContext(tempDir)
	if err := workspace.Init(); err != nil {
		t.Fatalf("Failed to initialize workspace: %v", err)
	}

	// Run initial generation to create the manifest
	generateSession, err := G.NewGenerateSession(workspace)
	if err != nil {
		t.Fatalf("Failed to create generate session: %v", err)
	}
	defer generateSession.Close()

	pkg, err := generateSession.GenerateFullManifest(context.Background())
	if err != nil {
		t.Fatalf("Failed to generate initial manifest: %v", err)
	}

	// Write initial manifest to file
	manifestStr, err := M.SerializeToString(pkg)
	if err != nil {
		t.Fatalf("Failed to serialize manifest: %v", err)
	}

	manifestPath := filepath.Join(distDir, "custom-elements.json")
	if err := os.WriteFile(manifestPath, []byte(manifestStr), 0644); err != nil {
		t.Fatalf("Failed to write manifest: %v", err)
	}

	// Create a full LSP server instance like TestServerLevelIntegration
	server, err := lsp.NewServer(workspace)
	if err != nil {
		t.Fatalf("Failed to create LSP server: %v", err)
	}
	defer server.Close()

	// Initialize the server (this loads manifests and starts file watching AND generate watcher)
	err = server.InitializeForTesting()
	if err != nil {
		t.Fatalf("Failed to initialize server: %v", err)
	}

	registry := server.Registry()

	// Create document manager
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	// Set up completion context using centralized MockServerContext
	ctx := testhelpers.NewMockServerContext()
	// Copy registry data to mock context for interface compatibility
	for _, tagName := range registry.AllTagNames() {
		if element, exists := registry.Element(tagName); exists {
			ctx.AddElement(tagName, element)
		}
		if attrs, exists := registry.Attributes(tagName); exists {
			ctx.AddAttributes(tagName, attrs)
		}
		if slots, exists := registry.Slots(tagName); exists {
			ctx.AddSlots(tagName, slots)
		}
	}
	ctx.SetDocumentManager(dm)

	// Create HTML demo file
	demoDir := filepath.Join(tempDir, "demo")
	if err := os.MkdirAll(demoDir, 0755); err != nil {
		t.Fatalf("Failed to create demo directory: %v", err)
	}

	htmlFilePath := filepath.Join(demoDir, "index.html")
	CopyFixtureFile(t, fixtureDir, "demo.html", demoDir, "index.html")

	// Load HTML content and open document
	htmlContent, err := os.ReadFile(htmlFilePath)
	if err != nil {
		t.Fatalf("Failed to read HTML content: %v", err)
	}

	htmlURI := "file://" + htmlFilePath
	doc := dm.OpenDocument(htmlURI, string(htmlContent), 1)
	if doc == nil {
		t.Fatalf("Failed to open HTML document")
	}

	// Test initial completions - should have one, two, three
	initialItems := completion.GetAttributeValueCompletions(ctx, "test-element", "prop")
	t.Logf("Initial completions: %v", testhelpers.GetCompletionLabels(initialItems))

	// Verify initial completions contain expected values
	expectedInitial := []string{"one", "two", "three"}
	for _, expected := range expectedInitial {
		found := false
		for _, item := range initialItems {
			if item.Label == expected {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("Initial completions missing '%s'. Got: %v", expected, testhelpers.GetCompletionLabels(initialItems))
		}
	}

	// The server has already started file watching and generate watcher via InitializeForTesting()

	// Now simulate editing the source file to add 'four'
	t.Logf("=== User edits test-element.ts to add 'four' ===")

	CopyFixtureFile(t, fixtureDir, "updated-element.ts", elementsDir, "test-element.ts")

	// Ensure the file is synced to disk
	if err := syncFile(filepath.Join(elementsDir, "test-element.ts")); err != nil {
		t.Logf("Warning: Could not sync file: %v", err)
	}

	// Verify the file was actually updated
	updatedContent, err := os.ReadFile(filepath.Join(elementsDir, "test-element.ts"))
	if err != nil {
		t.Fatalf("Failed to read updated file: %v", err)
	}
	t.Logf("Updated file content: %s", string(updatedContent))

	t.Logf("File updated, waiting for generate watcher to process changes...")

	// Wait for the generate watcher to detect changes and regenerate
	time.Sleep(5 * time.Second)

	// NOTE: Don't call ReloadManifestsDirectly() here as it overwrites the in-memory
	// updates from the generate watcher with old file content

	// Give a small buffer for any remaining async operations
	time.Sleep(100 * time.Millisecond)

	// Debug: Check what the generate command actually wrote to the manifest
	manifestContent, err := os.ReadFile(manifestPath)
	if err == nil {
		t.Logf("Manifest file after generation: %s", string(manifestContent))
	}

	// Test updated completions - should now include 'four'
	updatedItems := completion.GetAttributeValueCompletions(ctx, "test-element", "prop")
	t.Logf("Updated completions: %v", testhelpers.GetCompletionLabels(updatedItems))

	// Verify updated completions contain all original values plus 'four'
	expectedUpdated := []string{"one", "two", "three", "four"}
	for _, expected := range expectedUpdated {
		found := false
		for _, item := range updatedItems {
			if item.Label == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Updated completions missing '%s'. Got: %v", expected, testhelpers.GetCompletionLabels(updatedItems))

			// Debug: Check what's in the registry
			if attrs, exists := ctx.Attributes("test-element"); exists {
				if propAttr, hasProp := attrs["prop"]; hasProp {
					t.Logf("Registry shows prop attribute type: %s", propAttr.Type.Text)
				}
			}

			// Debug: Check the manifest file content
			if manifestContent, err := os.ReadFile(manifestPath); err == nil {
				t.Logf("Current manifest content: %s", string(manifestContent))
			}
		}
	}

	// Verify 'four' specifically was added
	hasUpdatedValue := false
	for _, item := range updatedItems {
		if item.Label == "four" {
			hasUpdatedValue = true
			break
		}
	}

	if hasUpdatedValue {
		t.Logf("✅ SUCCESS: 'four' found in completions!")
	} else {
		t.Errorf("❌ FAILURE: 'four' not found in updated completions")
	}
}

// copyFixtureFile copies a file from the fixture directory to the target location

// syncFile ensures a file is synced to disk
func syncFile(path string) error {
	file, err := os.OpenFile(path, os.O_RDWR, 0)
	if err != nil {
		return err
	}
	defer file.Close()
	return file.Sync()
}

