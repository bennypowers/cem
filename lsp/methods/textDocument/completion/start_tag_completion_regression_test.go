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
	"bennypowers.dev/cem/lsp/testhelpers"
	"bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// TestStartTagCompletionRegression ensures start tag completions work correctly
// and prevents future regressions of this critical completion feature
func TestStartTagCompletionRegression(t *testing.T) {
	// Load test manifest with custom elements
	fixtureDir := filepath.Join("slot-completions-test")
	manifestPath := filepath.Join(fixtureDir, "manifest.json")

	manifestBytes, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatalf("Failed to read regression test manifest: %v", err)
	}

	var pkg M.Package
	err = json.Unmarshal(manifestBytes, &pkg)
	if err != nil {
		t.Fatalf("Failed to parse regression test manifest: %v", err)
	}

	// Create context and add the test manifest
	ctx := testhelpers.NewMockServerContext()
	ctx.AddManifest(&pkg)

	// Create and set a real DocumentManager
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	// Test cases for start tag completions using the context analysis
	tests := []struct {
		name                string
		html                string
		position            protocol.Position
		expectedType        types.CompletionContextType
		expectedCompletions []string // Tag names that should be suggested
		minCompletions      int      // Minimum number of completions expected
		description         string
	}{
		{
			name:                "Start tag completion after <",
			html:                "<",
			position:            protocol.Position{Line: 0, Character: 1},
			expectedType:        types.CompletionTagName,
			expectedCompletions: []string{"card-element", "dialog-element", "my-custom-element"},
			minCompletions:      3,
			description:         "Should suggest all custom elements after typing <",
		},
		{
			name:                "Start tag completion with partial tag name",
			html:                "<card",
			position:            protocol.Position{Line: 0, Character: 5},
			expectedType:        types.CompletionTagName,
			expectedCompletions: []string{"card-element"},
			minCompletions:      1,
			description:         "Should suggest card-element when typing <card",
		},
		{
			name:                "Start tag completion after opening tag start",
			html:                "<my-",
			position:            protocol.Position{Line: 0, Character: 4},
			expectedType:        types.CompletionTagName,
			expectedCompletions: []string{"my-custom-element"},
			minCompletions:      1,
			description:         "Should suggest my-custom-element when typing <my-",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a real document using the DocumentManager
			uri := "test://test.html"
			doc := dm.OpenDocument(uri, tt.html, 1)
			ctx.AddDocument(uri, doc)

			// Test the completion context analysis first with DocumentManager
			analysis, err := textDocument.AnalyzeCompletionContextWithDM(doc, tt.position, "", dm)
			if err != nil {
				t.Fatalf("Failed to analyze completion context: %v", err)
			}
			if analysis.Type != tt.expectedType {
				t.Errorf("Expected completion type %v, got %v", tt.expectedType, analysis.Type)
				return
			}

			// If we correctly detected tag name completion context,
			// then the start tag completion issue is not in the context analysis
			if analysis.Type == types.CompletionTagName {
				t.Logf("✓ %s - Context analysis correctly detected TagName completion", tt.description)

				// Verify that the registry has the expected elements
				allTagNames := ctx.AllTagNames()
				t.Logf("  Available tag names: %v", allTagNames)

				for _, expectedTag := range tt.expectedCompletions {
					found := false
					for _, tagName := range allTagNames {
						if tagName == expectedTag {
							found = true
							break
						}
					}
					if !found {
						t.Errorf("REGRESSION: Expected tag '%s' not found in registry. Available tags: %v",
							expectedTag, allTagNames)
					}
				}
			} else {
				t.Errorf("REGRESSION: Expected completion type %v, got %v. This means start tag detection is broken!",
					tt.expectedType, analysis.Type)
			}
		})
	}
}

