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
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/testhelpers"
	W "bennypowers.dev/cem/workspace"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestBasicAttributeValueCompletions(t *testing.T) {
	// Setup fixture workspace
	fixtureDir, err := filepath.Abs("test-fixtures/completion-basic")
	if err != nil {
		t.Fatalf("Failed to get fixture path: %v", err)
	}

	// The generate watcher will create the manifest from the TypeScript source file

	// Create workspace context after manifest exists
	workspace := W.NewFileSystemWorkspaceContext(fixtureDir)
	err = workspace.Init()
	if err != nil {
		t.Fatalf("Failed to initialize workspace: %v", err)
	}

	// Create server (this loads manifests automatically)
	server, err := lsp.NewServer(workspace, lsp.TransportStdio)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	// Initialize the server
	err = server.InitializeForTesting()
	if err != nil {
		t.Fatalf("Failed to initialize server: %v", err)
	}

	// Create document manager
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	// Create adapter using centralized MockServerContext
	adapter := testhelpers.NewMockServerContext()
	// Copy registry data to mock context for interface compatibility
	registry := server.Registry()
	for _, tagName := range registry.AllTagNames() {
		if element, exists := registry.Element(tagName); exists {
			adapter.AddElement(tagName, element)
		}
		if attrs, exists := registry.Attributes(tagName); exists {
			adapter.AddAttributes(tagName, attrs)
		}
		if slots, exists := registry.Slots(tagName); exists {
			adapter.AddSlots(tagName, slots)
		}
	}
	adapter.SetDocumentManager(dm)

	// Debug: Check what elements are loaded
	allTags := adapter.AllTagNames()
	t.Logf("Loaded tags: %v", allTags)
	if attrs, exists := adapter.Attributes("test-alert"); exists {
		t.Logf("test-alert attributes: %d found", len(attrs))
		for name, attr := range attrs {
			t.Logf("  - %s: %s", name, attr.Type.Text)
		}
	} else {
		t.Logf("test-alert element not found in registry")
	}

	// Test basic attribute value completions
	t.Run("state attribute completions", func(t *testing.T) {
		items, err := TestHelper.GetAttributeValueCompletionsUsingMainEntry(adapter, "file:///test.html", "test-alert", "state")
		if err != nil {
			t.Fatalf("Failed to get completions: %v", err)
		}

		// Verify we get the expected completions (including error from union type)
		expectedValues := []string{"info", "success", "warning", "error"}
		found := make(map[string]bool)

		for _, item := range items {
			found[item.Label] = true
		}

		for _, expected := range expectedValues {
			if !found[expected] {
				t.Errorf("Expected completion '%s' not found. Got: %v", expected, getFixtureCompletionLabels(items))
			}
		}

		// Should have at least 4 completions (may include default value)
		if len(items) < 4 {
			t.Errorf("Expected at least 4 completions, got %d: %v", len(items), getFixtureCompletionLabels(items))
		}
	})

	t.Run("variant attribute completions", func(t *testing.T) {
		items, err := TestHelper.GetAttributeValueCompletionsUsingMainEntry(adapter, "file:///test.html", "test-alert", "variant")
		if err != nil {
			t.Fatalf("Failed to get completions: %v", err)
		}

		// Verify we get the expected completions (including ghost from union type)
		expectedValues := []string{"primary", "secondary", "ghost"}
		found := make(map[string]bool)

		for _, item := range items {
			found[item.Label] = true
		}

		for _, expected := range expectedValues {
			if !found[expected] {
				t.Errorf("Expected completion '%s' not found. Got: %v", expected, getFixtureCompletionLabels(items))
			}
		}

		// Should have at least 3 completions (may include default value and duplicates)
		if len(items) < 3 {
			t.Errorf("Expected at least 3 completions, got %d: %v", len(items), getFixtureCompletionLabels(items))
		}
	})

	t.Run("boolean attribute completions", func(t *testing.T) {
		items, err := TestHelper.GetAttributeValueCompletionsUsingMainEntry(adapter, "file:///test.html", "test-alert", "disabled")
		if err != nil {
			t.Fatalf("Failed to get completions: %v", err)
		}

		// Boolean attributes should have no value completions (presence=true, absence=false)
		if len(items) != 0 {
			t.Errorf("Expected 0 completions for boolean attribute, got %d: %v", len(items), getFixtureCompletionLabels(items))
		}
	})
}

// Helper function to extract completion labels for debugging
func getFixtureCompletionLabels(items []protocol.CompletionItem) []string {
	labels := make([]string, len(items))
	for i, item := range items {
		labels[i] = item.Label
	}
	return labels
}
