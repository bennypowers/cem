package completion_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/methods/textDocument/completion"
	"bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestSlotAttributeCompletions(t *testing.T) {
	// Load test manifest with slot definitions
	fixtureDir := filepath.Join("slot-completions-test")
	manifestPath := filepath.Join(fixtureDir, "manifest.json")

	manifestBytes, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatalf("Failed to read test manifest: %v", err)
	}

	var pkg M.Package
	err = json.Unmarshal(manifestBytes, &pkg)
	if err != nil {
		t.Fatalf("Failed to parse manifest: %v", err)
	}

	// Create registry and add the test manifest
	registry := lsp.NewTestRegistry()
	registry.AddManifest(&pkg)

	// Create a mock document manager
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	// Test context that provides the registry data
	ctx := &testSlotCompletionContext{
		registry: registry,
	}

	tests := []struct {
		name          string
		html          string
		position      protocol.Position
		expectedSlots []string
		description   string
	}{
		{
			name:          "Slot completion for direct child of card-element",
			html:          `<card-element><button slot="`,
			position:      protocol.Position{Line: 0, Character: 28},
			expectedSlots: []string{"header", "footer", "actions"},
			description:   "Should provide slot completions for button inside card-element (excluding default slot)",
		},
		{
			name:          "Slot completion for direct child of dialog-element",
			html:          `<dialog-element><span slot="`,
			position:      protocol.Position{Line: 0, Character: 27},
			expectedSlots: []string{"title", "content"},
			description:   "Should provide slot completions for span inside dialog-element",
		},
		{
			name:          "No slot completion for non-custom element parent",
			html:          `<div><button slot="`,
			position:      protocol.Position{Line: 0, Character: 19},
			expectedSlots: []string{},
			description:   "Should not provide slot completions when parent is not a custom element",
		},
		{
			name:          "No slot completion when no parent found",
			html:          `<button slot="`,
			position:      protocol.Position{Line: 0, Character: 14},
			expectedSlots: []string{},
			description:   "Should not provide slot completions when no parent element is found",
		},
		{
			name:          "Slot completion with nested elements",
			html:          `<card-element><div><button slot="`,
			position:      protocol.Position{Line: 0, Character: 33},
			expectedSlots: []string{"header", "footer", "actions"},
			description:   "Should provide slot completions for nested child elements (excluding default slot)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock document for this test case
			mockDoc := NewMockDocument(tt.html)

			// Call the slot completion function directly
			completions := completion.GetAttributeValueCompletionsWithContext(ctx, mockDoc, tt.position, "", "slot")

			// Check if we got the expected completions
			if len(tt.expectedSlots) == 0 {
				if len(completions) != 0 {
					t.Errorf("Expected no completions, but got %d", len(completions))
				}
				return
			}

			// Create a map of found labels for easy checking
			foundLabels := make(map[string]bool)
			for _, completion := range completions {
				foundLabels[completion.Label] = true
			}

			// Check that all expected slot names are present
			for _, expectedSlot := range tt.expectedSlots {
				if !foundLabels[expectedSlot] {
					t.Errorf("Expected slot '%s' not found. Got labels: %v",
						expectedSlot, getSlotCompletionLabels(completions))
				}
			}

			// Log completion details for debugging
			t.Logf("Test: %s", tt.description)
			t.Logf("Found %d completions: %v", len(completions), getSlotCompletionLabels(completions))
		})
	}
}

func TestSlotCompletionDetails(t *testing.T) {
	// Load test manifest
	fixtureDir := filepath.Join("slot-completions-test")
	manifestPath := filepath.Join(fixtureDir, "manifest.json")

	manifestBytes, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatalf("Failed to read test manifest: %v", err)
	}

	var pkg M.Package
	err = json.Unmarshal(manifestBytes, &pkg)
	if err != nil {
		t.Fatalf("Failed to parse manifest: %v", err)
	}

	registry := lsp.NewTestRegistry()
	registry.AddManifest(&pkg)

	ctx := &testSlotCompletionContext{
		registry: registry,
	}

	// Test specific completion details
	mockDoc := NewMockDocument(`<card-element><button slot="`)
	position := protocol.Position{Line: 0, Character: 28}

	completions := completion.GetAttributeValueCompletionsWithContext(ctx, mockDoc, position, "", "slot")

	// Should have 3 completions: header, footer, actions (no default slot)
	if len(completions) != 3 {
		t.Errorf("Expected 3 completions, got %d", len(completions))
	}

	// Check that completions have proper details
	for _, completion := range completions {
		if completion.Kind == nil || *completion.Kind != protocol.CompletionItemKindValue {
			t.Errorf("Completion '%s' should have Value kind", completion.Label)
		}

		if completion.Detail == nil {
			t.Errorf("Completion '%s' should have detail", completion.Label)
		}

		if completion.Documentation == nil {
			t.Errorf("Completion '%s' should have documentation", completion.Label)
		}
	}
}

// Helper function to extract labels from completion items
func getSlotCompletionLabels(completions []protocol.CompletionItem) []string {
	labels := make([]string, len(completions))
	for i, completion := range completions {
		labels[i] = completion.Label
	}
	return labels
}

// testSlotCompletionContext implements CompletionContext for testing slot completions
type testSlotCompletionContext struct {
	registry *lsp.Registry
}

func (ctx *testSlotCompletionContext) Document(uri string) types.Document {
	return nil // Not needed for these tests
}

func (ctx *testSlotCompletionContext) AllTagNames() []string {
	return ctx.registry.AllTagNames()
}

func (ctx *testSlotCompletionContext) Element(tagName string) (*M.CustomElement, bool) {
	return ctx.registry.Element(tagName)
}

func (ctx *testSlotCompletionContext) Attributes(tagName string) (map[string]*M.Attribute, bool) {
	return ctx.registry.Attributes(tagName)
}

func (ctx *testSlotCompletionContext) Slots(tagName string) ([]M.Slot, bool) {
	return ctx.registry.Slots(tagName)
}
