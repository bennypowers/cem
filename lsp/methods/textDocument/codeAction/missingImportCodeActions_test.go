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
package codeAction

import (
	"testing"

	"bennypowers.dev/cem/lsp/types"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// mockCodeActionContext implements CodeActionContext for testing
type mockCodeActionContext struct{}

func (m *mockCodeActionContext) Document(uri string) types.Document {
	return nil // Not needed for this test
}

func (m *mockCodeActionContext) RawDocumentManager() any {
	return nil // Mock implementation
}

func TestCreateMissingImportAction(t *testing.T) {
	tests := []struct {
		name           string
		data           map[string]any
		documentURI    string
		expectAction   bool
		expectedTitle  string
		expectedImport string
	}{
		{
			name: "Valid missing import data for TypeScript",
			data: map[string]any{
				"type":       string(types.DiagnosticTypeMissingImport),
				"original":   "my-element",
				"suggestion": "import './my-element.js'",
				"importPath": "./my-element.js",
				"tagName":    "my-element",
				"range": map[string]any{
					"start": map[string]any{"line": float64(5), "character": float64(10)},
					"end":   map[string]any{"line": float64(5), "character": float64(20)},
				},
			},
			documentURI:    "file:///test.ts",
			expectAction:   true,
			expectedTitle:  "Add import for 'my-element'",
			expectedImport: `import "./my-element.js";`, // Local module import
		},
		{
			name: "Valid missing import data for HTML",
			data: map[string]any{
				"type":       string(types.DiagnosticTypeMissingImport),
				"original":   "my-element",
				"suggestion": "import './my-element.js'",
				"importPath": "./my-element.js",
				"tagName":    "my-element",
				"range": map[string]any{
					"start": map[string]any{"line": float64(5), "character": float64(10)},
					"end":   map[string]any{"line": float64(5), "character": float64(20)},
				},
			},
			documentURI:    "file:///test.html",
			expectAction:   true,
			expectedTitle:  "Add import for 'my-element'",
			expectedImport: `<script type="module">
	import "./my-element.js";
</script>`, // Local module import as inline script
		},
		{
			name: "Package import for TypeScript",
			data: map[string]any{
				"type":       string(types.DiagnosticTypeMissingImport),
				"original":   "ui-button",
				"suggestion": "import '@my-org/components'",
				"importPath": "@my-org/components",
				"tagName":    "ui-button",
				"range": map[string]any{
					"start": map[string]any{"line": float64(2), "character": float64(5)},
					"end":   map[string]any{"line": float64(2), "character": float64(14)},
				},
			},
			documentURI:    "file:///test.ts",
			expectAction:   true,
			expectedTitle:  "Add import for 'ui-button'",
			expectedImport: `import "@my-org/components";`, // Package import
		},
		{
			name: "Wrong diagnostic type",
			data: map[string]any{
				"type":       string(types.DiagnosticTypeTagSuggestion),
				"original":   "my-element",
				"suggestion": "my-other-element",
			},
			documentURI:  "file:///test.ts",
			expectAction: false,
		},
		{
			name: "Invalid data structure",
			data: map[string]any{
				"type": "invalid-type",
			},
			documentURI:  "file:///test.ts",
			expectAction: false,
		},
		{
			name:         "Empty data",
			data:         map[string]any{},
			documentURI:  "file:///test.ts",
			expectAction: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diagnostic := &protocol.Diagnostic{
				Range: protocol.Range{
					Start: protocol.Position{Line: 5, Character: 10},
					End:   protocol.Position{Line: 5, Character: 20},
				},
				Message: "Test diagnostic",
			}

			// Create a mock context
			mockCtx := &mockCodeActionContext{}

			action := createMissingImportAction(mockCtx, diagnostic, tt.data, tt.documentURI)

			if tt.expectAction {
				if action == nil {
					t.Fatal("Expected action to be created, but got nil")
				}

				if action.Title != tt.expectedTitle {
					t.Errorf("Expected title '%s', got '%s'", tt.expectedTitle, action.Title)
				}

				if action.Kind == nil || *action.Kind != protocol.CodeActionKindQuickFix {
					t.Errorf("Expected kind 'quickfix', got %v", action.Kind)
				}

				if action.Edit == nil || action.Edit.Changes == nil {
					t.Fatal("Expected edit with changes, got nil")
				}

				changes, exists := action.Edit.Changes[tt.documentURI]
				if !exists {
					t.Fatal("Expected changes for document URI")
				}

				if len(changes) != 1 {
					t.Fatalf("Expected 1 text edit, got %d", len(changes))
				}

				if changes[0].NewText != tt.expectedImport+"\n" {
					t.Errorf("Expected import text '%s\\n', got '%s'", tt.expectedImport, changes[0].NewText)
				}

				// Verify diagnostic is attached
				if len(action.Diagnostics) != 1 {
					t.Errorf("Expected 1 diagnostic, got %d", len(action.Diagnostics))
				}
			} else {
				if action != nil {
					t.Errorf("Expected no action, but got: %+v", action)
				}
			}
		})
	}
}
