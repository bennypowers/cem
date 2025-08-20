package completion_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/methods/textDocument/completion"
	"bennypowers.dev/cem/lsp/testhelpers"
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

	// Create context and add the test manifest
	ctx := testhelpers.NewMockServerContext()
	ctx.AddManifest(&pkg)

	// Create a real document manager
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

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
			doc := dm.OpenDocument("test://test.html", tt.html, 1)

			// Call the slot completion function directly
			completions := completion.GetAttributeValueCompletionsWithContext(ctx, doc, tt.position, "", "slot")

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

	ctx := testhelpers.NewMockServerContext()
	ctx.AddManifest(&pkg)

	// Create and set a real DocumentManager
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	// Test specific completion details
	doc := dm.OpenDocument("test://test.html", `<card-element><button slot="`, 1)
	position := protocol.Position{Line: 0, Character: 28}

	completions := completion.GetAttributeValueCompletionsWithContext(ctx, doc, position, "", "slot")

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


// TestSlotAttributeNameSuggestion tests basic slot attribute name suggestions functionality
// For comprehensive regression testing, see TestSlotAttributeNameSuggestionRegression
func TestSlotAttributeNameSuggestion(t *testing.T) {
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

	// Create context and add the test manifest
	ctx := testhelpers.NewMockServerContext()
	ctx.AddManifest(&pkg)

	// Create and set a real DocumentManager
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	tests := []struct {
		name                string
		html                string
		position            protocol.Position
		tagName             string
		shouldSuggestSlot   bool
		description         string
	}{
		{
			name:              "Slot attribute suggested for button inside card-element",
			html:              `<card-element><button `,
			position:          protocol.Position{Line: 0, Character: 22},
			tagName:           "button",
			shouldSuggestSlot: true,
			description:       "Should suggest slot attribute for button inside card-element with named slots",
		},
		{
			name:              "Slot attribute suggested for span inside dialog-element",
			html:              `<dialog-element><span `,
			position:          protocol.Position{Line: 0, Character: 22},
			tagName:           "span",
			shouldSuggestSlot: true,
			description:       "Should suggest slot attribute for span inside dialog-element with named slots",
		},
		{
			name:              "Slot attribute suggested for custom element inside card-element",
			html:              `<card-element><my-custom-element `,
			position:          protocol.Position{Line: 0, Character: 33},
			tagName:           "my-custom-element",
			shouldSuggestSlot: true,
			description:       "Should suggest slot attribute for custom element inside card-element",
		},
		{
			name:              "No slot attribute for element not inside custom element",
			html:              `<div><button `,
			position:          protocol.Position{Line: 0, Character: 13},
			tagName:           "button",
			shouldSuggestSlot: false,
			description:       "Should not suggest slot attribute when not inside custom element",
		},
		{
			name:              "No slot attribute for top-level element",
			html:              `<button `,
			position:          protocol.Position{Line: 0, Character: 8},
			tagName:           "button",
			shouldSuggestSlot: false,
			description:       "Should not suggest slot attribute for top-level element",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock document for this test case
			doc := dm.OpenDocument("test://test.html", tt.html, 1)

			
			// Call the attribute completion function with context
			completions := completion.GetAttributeCompletionsWithContext(ctx, doc, tt.position, tt.tagName)

			// Check if slot attribute is suggested
			foundSlot := false
			for _, completion := range completions {
				if completion.Label == "slot" {
					foundSlot = true
					
					// Verify the completion has proper structure
					if completion.InsertText == nil || *completion.InsertText != `slot="$0"` {
						t.Errorf("Expected slot completion to have snippet insert text 'slot=\"$0\"', got: %v", completion.InsertText)
					}
					
					if completion.Kind == nil || *completion.Kind != protocol.CompletionItemKindProperty {
						t.Errorf("Expected slot completion to have Property kind")
					}
					
					break
				}
			}

			if tt.shouldSuggestSlot && !foundSlot {
				t.Errorf("Expected slot attribute to be suggested but it wasn't. Got completions: %v", 
					getSlotCompletionLabels(completions))
			}

			if !tt.shouldSuggestSlot && foundSlot {
				t.Errorf("Expected slot attribute NOT to be suggested but it was. Got completions: %v", 
					getSlotCompletionLabels(completions))
			}

			t.Logf("Test: %s", tt.description)
			t.Logf("Found %d completions: %v", len(completions), getSlotCompletionLabels(completions))
		})
	}
}
