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
package codeAction_test

import (
	"testing"

	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/methods/textDocument/codeAction"
	"bennypowers.dev/cem/lsp/testhelpers"
	"go.lsp.dev/protocol"
)

func TestCodeActionSlotSuggestion(t *testing.T) {
	// Create mock context
	ctx := testhelpers.NewMockServerContext()

	// Create DocumentManager and add document
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	doc := dm.OpenDocument("test://test.html", `<my-element><div slot="heade">Content</div></my-element>`, 1)
	ctx.AddDocument("test://test.html", doc)

	// Create mock diagnostic with slot suggestion data
	dataMap := map[string]any{
		"type":       "slot-suggestion",
		"original":   "heade",
		"suggestion": "header",
		"range": map[string]any{
			"start": map[string]any{
				"line":      float64(0),
				"character": float64(25),
			},
			"end": map[string]any{
				"line":      float64(0),
				"character": float64(30),
			},
		},
	}
	data, _ := protocol.Marshal(dataMap)
	diagnostic := protocol.Diagnostic{
		Range: protocol.Range{
			Start: protocol.Position{Line: 0, Character: 25},
			End:   protocol.Position{Line: 0, Character: 30},
		},
		Severity: protocol.DiagnosticSeverityError,
		Source:   protocol.NewOptional("cem-lsp"),
		Message:  protocol.String("Unknown slot 'heade' for element 'my-element'. Did you mean 'header'?"),
		Data:     data,
	}

	// Create code action request
	params := &protocol.CodeActionParams{
		TextDocument: protocol.TextDocumentIdentifier{
			URI: "test://test.html",
		},
		Context: protocol.CodeActionContext{
			Diagnostics: []protocol.Diagnostic{diagnostic},
		},
	}

	// Call CodeAction
	result, err := codeAction.CodeAction(ctx, params)
	if err != nil {
		t.Fatalf("CodeAction failed: %v", err)
	}

	// Verify result
	actions := result

	if len(actions) != 1 {
		t.Fatalf("Expected 1 code action, got %d", len(actions))
	}

	action := actions[0]

	// Verify action properties
	expectedTitle := "Change 'heade' to 'header'"
	if action.Title != expectedTitle {
		t.Errorf("Expected title '%s', got '%s'", expectedTitle, action.Title)
	}

	if action.Kind == nil || *action.Kind != protocol.CodeActionKindQuickFix {
		t.Errorf("Expected kind QuickFix, got %v", action.Kind)
	}

	// Verify edit
	if action.Edit == nil {
		t.Fatal("Expected edit to be present")
	}

	edits, exists := action.Edit.Changes["test://test.html"]
	if !exists {
		t.Fatal("Expected edit for test://test.html")
	}

	if len(edits) != 1 {
		t.Fatalf("Expected 1 edit, got %d", len(edits))
	}

	edit := edits[0]
	if edit.NewText != "header" {
		t.Errorf("Expected new text 'header', got '%s'", edit.NewText)
	}

	// Verify range
	expectedRange := protocol.Range{
		Start: protocol.Position{Line: 0, Character: 25},
		End:   protocol.Position{Line: 0, Character: 30},
	}
	if edit.Range != expectedRange {
		t.Errorf("Expected range %+v, got %+v", expectedRange, edit.Range)
	}

	// Verify diagnostics
	if len(action.Diagnostics) != 1 {
		t.Errorf("Expected 1 diagnostic, got %d", len(action.Diagnostics))
	}
}

func TestCodeActionNoDiagnostics(t *testing.T) {
	// Create mock context
	ctx := testhelpers.NewMockServerContext()

	// Create DocumentManager and add document
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	doc := dm.OpenDocument("test://test.html", `<my-element><div slot="header">Content</div></my-element>`, 1)
	ctx.AddDocument("test://test.html", doc)

	// Create code action request with no diagnostics
	params := &protocol.CodeActionParams{
		TextDocument: protocol.TextDocumentIdentifier{
			URI: "test://test.html",
		},
		Context: protocol.CodeActionContext{
			Diagnostics: []protocol.Diagnostic{},
		},
	}

	// Call CodeAction
	result, err := codeAction.CodeAction(ctx, params)
	if err != nil {
		t.Fatalf("CodeAction failed: %v", err)
	}

	// Verify result
	actions := result

	if len(actions) != 0 {
		t.Errorf("Expected 0 code actions for valid content, got %d", len(actions))
	}
}

func TestCodeActionIsPreferred(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	doc := dm.OpenDocument("test://test.html", `<my-element><div slot="heade">Content</div></my-element>`, 1)
	ctx.AddDocument("test://test.html", doc)

	tests := []struct {
		name string
		data map[string]any
	}{
		{
			name: "slot suggestion",
			data: map[string]any{
				"type": "slot-suggestion", "original": "heade", "suggestion": "header",
				"range": map[string]any{
					"start": map[string]any{"line": float64(0), "character": float64(25)},
					"end":   map[string]any{"line": float64(0), "character": float64(30)},
				},
			},
		},
		{
			name: "tag suggestion",
			data: map[string]any{
				"type": "tag-suggestion", "original": "my-elem", "suggestion": "my-element",
				"range": map[string]any{
					"start": map[string]any{"line": float64(0), "character": float64(1)},
					"end":   map[string]any{"line": float64(0), "character": float64(8)},
				},
			},
		},
		{
			name: "attribute suggestion",
			data: map[string]any{
				"type": "attribute-suggestion", "original": "colr", "suggestion": "color",
				"range": map[string]any{
					"start": map[string]any{"line": float64(0), "character": float64(12)},
					"end":   map[string]any{"line": float64(0), "character": float64(16)},
				},
			},
		},
		{
			name: "attribute value suggestion",
			data: map[string]any{
				"type": "attribute-value-suggestion", "original": "rde", "suggestion": "red",
				"range": map[string]any{
					"start": map[string]any{"line": float64(0), "character": float64(18)},
					"end":   map[string]any{"line": float64(0), "character": float64(21)},
				},
			},
		},
		{
			name: "missing import",
			data: map[string]any{
				"type": "missing-import", "original": "", "suggestion": "",
				"tagName": "my-element", "importPath": "@scope/my-element",
				"range": map[string]any{
					"start": map[string]any{"line": float64(0), "character": float64(0)},
					"end":   map[string]any{"line": float64(0), "character": float64(10)},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ttData, _ := protocol.Marshal(tt.data)
			params := &protocol.CodeActionParams{
				TextDocument: protocol.TextDocumentIdentifier{URI: "test://test.html"},
				Context: protocol.CodeActionContext{
					Diagnostics: []protocol.Diagnostic{{
						Range:  protocol.Range{Start: protocol.Position{Line: 0, Character: 0}, End: protocol.Position{Line: 0, Character: 5}},
						Source: protocol.NewOptional("cem-lsp"),
						Data:   ttData,
					}},
				},
			}

			result, err := codeAction.CodeAction(ctx, params)
			if err != nil {
				t.Fatalf("CodeAction failed: %v", err)
			}

			actions := result
			if len(actions) != 1 {
				t.Fatalf("Expected 1 action, got %d", len(actions))
			}

			if actions[0].IsPreferred == nil || !*actions[0].IsPreferred {
				t.Error("Expected IsPreferred to be true")
			}
		})
	}
}

// Mock implementations for testing
