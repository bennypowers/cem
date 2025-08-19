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
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/methods/textDocument/completion"
	"bennypowers.dev/cem/lsp/testhelpers"
	"bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// TestEndToEndStartTagCompletion tests start tag completion through the main Completion function
// This simulates what happens when a user types "<" in their editor
func TestEndToEndStartTagCompletion(t *testing.T) {
	// Load test manifest with custom elements
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

	// Create a completion context that returns documents
	ctx := &endToEndCompletionContext{
		registry: registry,
		documents: make(map[string]types.Document),
	}

	// Test scenarios that reproduce the user's issue
	tests := []struct {
		name                string
		documentContent     string
		position            protocol.Position
		triggerChar         *string
		expectedMinResults  int
		expectedTags        []string
		description         string
	}{
		{
			name:            "Start tag completion after typing <",
			documentContent: "<",
			position:        protocol.Position{Line: 0, Character: 1},
			triggerChar:     stringPtr("<"),
			expectedMinResults: 3,
			expectedTags:    []string{"card-element", "dialog-element", "my-custom-element"},
			description:     "Should provide start tag completions when user types <",
		},
		{
			name:            "Start tag completion in existing HTML",
			documentContent: "<div>\n  <",
			position:        protocol.Position{Line: 1, Character: 3},
			triggerChar:     stringPtr("<"),
			expectedMinResults: 3,
			expectedTags:    []string{"card-element", "dialog-element", "my-custom-element"},
			description:     "Should provide start tag completions in existing HTML context",
		},
		{
			name:            "Start tag completion with partial content",
			documentContent: "<card",
			position:        protocol.Position{Line: 0, Character: 5},
			triggerChar:     nil,
			expectedMinResults: 1,
			expectedTags:    []string{"card-element"},
			description:     "Should provide filtered start tag completions with partial typing",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up the document for this test
			uri := "test://test.html"
			doc := testhelpers.NewMockDocument(tt.documentContent)
			ctx.documents[uri] = doc

			// Create completion parameters
			params := &protocol.CompletionParams{
				TextDocumentPositionParams: protocol.TextDocumentPositionParams{
					TextDocument: protocol.TextDocumentIdentifier{URI: uri},
					Position:     tt.position,
				},
			}

			if tt.triggerChar != nil {
				params.Context = &protocol.CompletionContext{
					TriggerKind:      protocol.CompletionTriggerKindTriggerCharacter,
					TriggerCharacter: tt.triggerChar,
				}
			}

			// Call the main completion function
			result, err := completion.Completion(ctx, nil, params)
			if err != nil {
				t.Fatalf("Completion failed: %v", err)
			}

			// Convert result to completion items
			completionItems, ok := result.([]protocol.CompletionItem)
			if !ok {
				t.Fatalf("Expected []protocol.CompletionItem, got %T", result)
			}

			// Check minimum number of results
			if len(completionItems) < tt.expectedMinResults {
				t.Errorf("Expected at least %d completions, got %d: %v",
					tt.expectedMinResults, len(completionItems), extractLabels(completionItems))
			}

			// Check that expected tags are present
			foundTags := make(map[string]bool)
			for _, item := range completionItems {
				foundTags[item.Label] = true
			}

			for _, expectedTag := range tt.expectedTags {
				if !foundTags[expectedTag] {
					t.Errorf("Expected tag '%s' not found in completions: %v",
						expectedTag, extractLabels(completionItems))
				}
			}

			t.Logf("✓ %s", tt.description)
			t.Logf("  Found %d completions: %v", len(completionItems), extractLabels(completionItems))
		})
	}
}

// endToEndCompletionContext implements CompletionContext for end-to-end testing
type endToEndCompletionContext struct {
	registry  *lsp.Registry
	documents map[string]types.Document
}

func (ctx *endToEndCompletionContext) Document(uri string) types.Document {
	if doc, exists := ctx.documents[uri]; exists {
		return doc
	}
	return nil
}

func (ctx *endToEndCompletionContext) AllTagNames() []string {
	return ctx.registry.AllTagNames()
}

func (ctx *endToEndCompletionContext) Element(tagName string) (*M.CustomElement, bool) {
	return ctx.registry.Element(tagName)
}

func (ctx *endToEndCompletionContext) Attributes(tagName string) (map[string]*M.Attribute, bool) {
	return ctx.registry.Attributes(tagName)
}

func (ctx *endToEndCompletionContext) Slots(tagName string) ([]M.Slot, bool) {
	return ctx.registry.Slots(tagName)
}

// Helper functions
func stringPtr(s string) *string {
	return &s
}

func extractLabels(items []protocol.CompletionItem) []string {
	labels := make([]string, len(items))
	for i, item := range items {
		labels[i] = item.Label
	}
	return labels
}