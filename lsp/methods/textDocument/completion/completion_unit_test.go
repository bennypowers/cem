package completion_test

import (
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
	W "bennypowers.dev/cem/workspace"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// testCompletionAdapter implements the context interfaces needed for completion testing
type testCompletionAdapter struct {
	registry *lsp.Registry
	docMgr   *lsp.DocumentManager
}

func (t *testCompletionAdapter) Document(uri string) types.Document {
	return t.docMgr.Document(uri)
}

func (t *testCompletionAdapter) AllTagNames() []string {
	return t.registry.AllTagNames()
}

func (t *testCompletionAdapter) ElementDefinition(tagName string) (types.ElementDefinition, bool) {
	return t.registry.ElementDefinition(tagName)
}

func (t *testCompletionAdapter) Attributes(tagName string) (map[string]*M.Attribute, bool) {
	return t.registry.Attributes(tagName)
}

func (t *testCompletionAdapter) Slots(tagName string) ([]M.Slot, bool) {
	return t.registry.Slots(tagName)
}

func (t *testCompletionAdapter) Element(tagName string) (*M.CustomElement, bool) {
	return t.registry.Element(tagName)
}

func (t *testCompletionAdapter) GetDocumentManager() *lsp.DocumentManager {
	return t.docMgr
}

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
	server, err := lsp.NewServer(workspace)
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

	// Create adapter
	adapter := &testCompletionAdapter{
		registry: server.Registry(),
		docMgr:   dm,
	}

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
