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

	_ "bennypowers.dev/cem/internal/languages/registry"
	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/methods/textDocument/codeAction"
	"bennypowers.dev/cem/lsp/testhelpers"
	protocol "github.com/bennypowers/glsp/protocol_3_17"
)

func TestCodeActionCSSAmbiguousComment(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()

	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	cssContent := ".foo {\n  /** Blue colors */\n  background: var(--blue) var(--dark-blue);\n}\n"
	uri := "file:///test.css"
	doc := dm.OpenDocument(uri, cssContent, 1)
	ctx.AddDocument(uri, doc)

	severity := protocol.DiagnosticSeverityWarning
	source := "cem-lsp"
	diagnostic := protocol.Diagnostic{
		Range: protocol.Range{
			Start: protocol.Position{Line: 1, Character: 2},
			End:   protocol.Position{Line: 1, Character: 20},
		},
		Severity: &severity,
		Source:   &source,
		Message:  "Ambiguous comment ignored: more than one var() call in declaration.",
		Data: map[string]any{
			"type":        "css-ambiguous-comment",
			"commentText": "/** Blue colors */",
			"deleteRange": map[string]any{
				"start": map[string]any{"line": float64(1), "character": float64(0)},
				"end":   map[string]any{"line": float64(2), "character": float64(0)},
			},
			"properties": []any{
				map[string]any{
					"name":           "--blue",
					"insertPosition": map[string]any{"line": float64(2), "character": float64(14)},
				},
				map[string]any{
					"name":           "--dark-blue",
					"insertPosition": map[string]any{"line": float64(2), "character": float64(26)},
				},
			},
		},
	}

	params := &protocol.CodeActionParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: uri},
		Context:      protocol.CodeActionContext{Diagnostics: []protocol.Diagnostic{diagnostic}},
	}

	result, err := codeAction.CodeAction(ctx, nil, params)
	if err != nil {
		t.Fatalf("CodeAction failed: %v", err)
	}

	actions, ok := result.([]protocol.CodeAction)
	if !ok {
		t.Fatalf("Expected []protocol.CodeAction, got %T", result)
	}

	if len(actions) != 2 {
		t.Fatalf("Expected 2 code actions, got %d", len(actions))
	}

	// Action 1: associate with --blue
	if actions[0].Title != "Associate comment with `--blue`" {
		t.Errorf("Action 0: expected title with --blue, got %q", actions[0].Title)
	}
	if actions[0].Kind == nil || *actions[0].Kind != protocol.CodeActionKindQuickFix {
		t.Errorf("Action 0: expected QuickFix kind")
	}
	edits0 := actions[0].Edit.Changes[uri]
	if len(edits0) != 2 {
		t.Fatalf("Action 0: expected 2 edits, got %d", len(edits0))
	}
	if edits0[0].NewText != "" {
		t.Errorf("Action 0 edit 0: expected empty NewText (delete), got %q", edits0[0].NewText)
	}
	if edits0[0].Range.Start.Line != 1 || edits0[0].Range.End.Line != 2 {
		t.Errorf("Action 0 edit 0: delete range wrong: %v", edits0[0].Range)
	}
	if edits0[1].NewText != "/** Blue colors */ " {
		t.Errorf("Action 0 edit 1: expected comment text + space, got %q", edits0[1].NewText)
	}
	if edits0[1].Range.Start.Line != 2 || edits0[1].Range.Start.Character != 14 {
		t.Errorf("Action 0 edit 1: insert position wrong: %v", edits0[1].Range.Start)
	}

	// Action 2: associate with --dark-blue
	if actions[1].Title != "Associate comment with `--dark-blue`" {
		t.Errorf("Action 1: expected title with --dark-blue, got %q", actions[1].Title)
	}
	edits1 := actions[1].Edit.Changes[uri]
	if len(edits1) != 2 {
		t.Fatalf("Action 1: expected 2 edits, got %d", len(edits1))
	}
	if edits1[1].Range.Start.Character != 26 {
		t.Errorf("Action 1 edit 1: insert position wrong: char %d, want 26", edits1[1].Range.Start.Character)
	}
}
