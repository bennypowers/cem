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

	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/methods/textDocument"
	"bennypowers.dev/cem/lsp/methods/textDocument/completion"
	"bennypowers.dev/cem/lsp/testhelpers"
	"bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// TestStartTagCompletion tests start tag completion scenarios
// Consolidated from TestEndToEndStartTagCompletion and TestStartTagCompletionRegression
func TestStartTagCompletion(t *testing.T) {
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

	// Create context and add the test manifest
	ctx := testhelpers.NewMockServerContext()
	ctx.AddManifest(&pkg)

	// Debug: Check what tags were loaded
	allTags := ctx.AllTagNames()
	t.Logf("Loaded %d tags: %v", len(allTags), allTags)

	// Create and set a real DocumentManager
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	t.Run("Context Analysis", func(t *testing.T) {
		// Tests that verify context analysis correctly detects tag name completion
		tests := []struct {
			name         string
			html         string
			position     protocol.Position
			expectedType types.CompletionContextType
			description  string
		}{
			{
				name:         "After <",
				html:         "<",
				position:     protocol.Position{Line: 0, Character: 1},
				expectedType: types.CompletionTagName,
				description:  "Should detect tag name completion after typing <",
			},
			{
				name:         "Partial tag name",
				html:         "<my-c",
				position:     protocol.Position{Line: 0, Character: 5},
				expectedType: types.CompletionTagName,
				description:  "Should detect tag name completion for partial input",
			},
			{
				name:         "After < with newlines",
				html:         "\n<\n",
				position:     protocol.Position{Line: 1, Character: 1},
				expectedType: types.CompletionTagName,
				description:  "Should detect tag name completion across newlines",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				doc := dm.OpenDocument("test://test.html", tt.html, 1)
				analysis, err := textDocument.AnalyzeCompletionContextWithDM(doc, tt.position, "", dm)
				if err != nil {
					t.Fatalf("Failed to analyze context: %v", err)
				}

				if analysis.Type != tt.expectedType {
					t.Errorf("Expected context type %v, got %v", tt.expectedType, analysis.Type)
				}

				t.Logf("✓ %s", tt.description)
			})
		}
	})

	t.Run("Completion Results", func(t *testing.T) {
		// Tests that verify actual completion results
		tests := []struct {
			name               string
			documentContent    string
			position           protocol.Position
			triggerChar        *string
			expectedMinResults int
			expectedTags       []string
			description        string
		}{
			{
				name:               "Basic start tag",
				documentContent:    "<",
				position:           protocol.Position{Line: 0, Character: 1},
				triggerChar:        stringPtr("<"),
				expectedMinResults: 3,
				expectedTags:       []string{"card-element", "dialog-element", "my-custom-element"},
				description:        "Should suggest all custom elements after typing <",
			},
			{
				name:               "Partial tag name filtering",
				documentContent:    "<my-c",
				position:           protocol.Position{Line: 0, Character: 5},
				triggerChar:        nil,
				expectedMinResults: 1,
				expectedTags:       []string{"my-custom-element"},
				description:        "Should filter completions based on prefix",
			},
			{
				name:               "Start tag in context",
				documentContent:    "<div>\n  <",
				position:           protocol.Position{Line: 1, Character: 3},
				triggerChar:        stringPtr("<"),
				expectedMinResults: 3,
				expectedTags:       []string{"card-element", "dialog-element", "my-custom-element"},
				description:        "Should suggest custom elements within existing HTML",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				doc := dm.OpenDocument("test://test.html", tt.documentContent, 1)
				ctx.AddDocument("test://test.html", doc)

				params := &protocol.CompletionParams{
					TextDocumentPositionParams: protocol.TextDocumentPositionParams{
						TextDocument: protocol.TextDocumentIdentifier{URI: "test://test.html"},
						Position:     tt.position,
					},
					Context: &protocol.CompletionContext{
						TriggerCharacter: tt.triggerChar,
					},
				}

				result, err := completion.Completion(ctx, &glsp.Context{}, params)
				if err != nil {
					t.Fatalf("Completion failed: %v", err)
				}

				items, ok := result.([]protocol.CompletionItem)
				if !ok {
					t.Fatalf("Expected []CompletionItem, got %T", result)
				}

				if len(items) < tt.expectedMinResults {
					t.Errorf("Expected at least %d completions, got %d", tt.expectedMinResults, len(items))
				}

				// Verify expected tags are present
				foundTags := make(map[string]bool)
				for _, item := range items {
					foundTags[item.Label] = true
				}

				for _, expectedTag := range tt.expectedTags {
					if !foundTags[expectedTag] {
						t.Errorf("Expected tag '%s' not found in completions", expectedTag)
						t.Logf("Got tags: %v", getCompletionLabels(items))
					}
				}

				t.Logf("✓ %s (%d completions)", tt.description, len(items))
			})
		}
	})
}

// Helper functions
func stringPtr(s string) *string {
	return &s
}

func getCompletionLabels(items []protocol.CompletionItem) []string {
	labels := make([]string, len(items))
	for i, item := range items {
		labels[i] = item.Label
	}
	return labels
}
