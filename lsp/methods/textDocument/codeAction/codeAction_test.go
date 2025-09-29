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
	protocol "github.com/tliron/glsp/protocol_3_16"
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
	severity := protocol.DiagnosticSeverityError
	source := "cem-lsp"
	diagnostic := protocol.Diagnostic{
		Range: protocol.Range{
			Start: protocol.Position{Line: 0, Character: 25},
			End:   protocol.Position{Line: 0, Character: 30},
		},
		Severity: &severity,
		Source:   &source,
		Message:  "Unknown slot 'heade' for element 'my-element'. Did you mean 'header'?",
		Data: map[string]any{
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
		},
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
	result, err := codeAction.CodeAction(ctx, nil, params)
	if err != nil {
		t.Fatalf("CodeAction failed: %v", err)
	}

	// Verify result
	actions, ok := result.([]protocol.CodeAction)
	if !ok {
		t.Fatalf("Expected []protocol.CodeAction, got %T", result)
	}

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
	result, err := codeAction.CodeAction(ctx, nil, params)
	if err != nil {
		t.Fatalf("CodeAction failed: %v", err)
	}

	// Verify result
	actions, ok := result.([]protocol.CodeAction)
	if !ok {
		t.Fatalf("Expected []protocol.CodeAction, got %T", result)
	}

	if len(actions) != 0 {
		t.Errorf("Expected 0 code actions for valid content, got %d", len(actions))
	}
}

// Mock implementations for testing
