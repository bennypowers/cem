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
package publishDiagnostics_test

import (
	"encoding/json"
	"testing"

	_ "bennypowers.dev/cem/internal/languages/registry"
	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/methods/textDocument/codeAction"
	"bennypowers.dev/cem/lsp/methods/textDocument/publishDiagnostics"
	"bennypowers.dev/cem/lsp/testhelpers"
	"go.lsp.dev/protocol"
	urilib "go.lsp.dev/uri"
)

func TestCssDiagnostics_Fixtures(t *testing.T) {
	testutil.RunLSPFixtures(t, "testdata/css-diagnostics", func(t *testing.T, fixture *testutil.LSPFixture) {
		ctx := testhelpers.NewMockServerContext()

		dm, err := document.NewDocumentManager()
		if err != nil {
			t.Fatalf("Failed to create DocumentManager: %v", err)
		}
		defer dm.Close()
		ctx.SetDocumentManager(dm)

		uri := "file:///test.css"
		doc := dm.OpenDocument(uri, fixture.InputContent, 1)
		ctx.AddDocument(uri, doc)

		actual := publishDiagnostics.AnalyzeCssDiagnosticsForTest(ctx, doc)

		var expected []protocol.Diagnostic
		if err := fixture.GetExpected("expected", &expected); err != nil {
			t.Fatalf("Failed to load expected diagnostics: %v", err)
		}

		if len(actual) != len(expected) {
			t.Errorf("Expected %d diagnostics, got %d", len(expected), len(actual))
			for i, diag := range actual {
				t.Errorf("  Diagnostic %d: %s (range: %v)", i, diag.Message, diag.Range)
			}
			return
		}

		for i, exp := range expected {
			act := actual[i]

			if act.Range.Start.Line != exp.Range.Start.Line ||
				act.Range.Start.Character != exp.Range.Start.Character ||
				act.Range.End.Line != exp.Range.End.Line ||
				act.Range.End.Character != exp.Range.End.Character {
				t.Errorf("Diagnostic %d: expected range %v, got %v", i, exp.Range, act.Range)
			}

			if act.Message != exp.Message {
				t.Errorf("Diagnostic %d: expected message %q, got %q", i, exp.Message, act.Message)
			}

			if exp.Severity != 0 && act.Severity != exp.Severity {
				t.Errorf("Diagnostic %d: expected severity %v, got %v", i, exp.Severity, act.Severity)
			}
		}
	})
}

func TestCssDiagnostics_DataField(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	uri := "file:///test.css"
	css := ".foo {\n  /** spacing values */\n  padding: var(--pad-top) var(--pad-right) var(--pad-bottom);\n}\n"
	doc := dm.OpenDocument(uri, css, 1)
	ctx.AddDocument(uri, doc)

	diags := publishDiagnostics.AnalyzeCssDiagnosticsForTest(ctx, doc)
	if len(diags) != 1 {
		t.Fatalf("Expected 1 diagnostic, got %d", len(diags))
	}

	var data map[string]any
	if err := json.Unmarshal(diags[0].Data, &data); err != nil {
		t.Fatalf("Expected Data to be map[string]any, got error: %v", err)
	}

	if data["type"] != "css-ambiguous-comment" {
		t.Errorf("Expected type css-ambiguous-comment, got %v", data["type"])
	}

	if data["commentText"] != "/** spacing values */" {
		t.Errorf("Expected commentText, got %v", data["commentText"])
	}

	propsRaw, ok := data["properties"].([]any)
	if !ok {
		t.Fatalf("Expected properties to be []any, got %T", data["properties"])
	}
	if len(propsRaw) != 3 {
		t.Fatalf("Expected 3 properties, got %d", len(propsRaw))
	}
	props := make([]map[string]any, len(propsRaw))
	for i, p := range propsRaw {
		m, ok := p.(map[string]any)
		if !ok {
			t.Fatalf("Property %d: expected map[string]any, got %T", i, p)
		}
		props[i] = m
	}

	expectedNames := []string{"--pad-top", "--pad-right", "--pad-bottom"}
	for i, name := range expectedNames {
		if props[i]["name"] != name {
			t.Errorf("Property %d: expected name %q, got %v", i, name, props[i]["name"])
		}
		pos, ok := props[i]["insertPosition"].(map[string]any)
		if !ok {
			t.Errorf("Property %d: missing insertPosition", i)
			continue
		}
		if pos["line"] != float64(2) {
			t.Errorf("Property %d: expected insertPosition line 2, got %v", i, pos["line"])
		}
	}

	deleteRange, ok := data["deleteRange"].(map[string]any)
	if !ok {
		t.Fatalf("Expected deleteRange map, got %T", data["deleteRange"])
	}
	start, ok := deleteRange["start"].(map[string]any)
	if !ok {
		t.Fatalf("Expected deleteRange.start map, got %T", deleteRange["start"])
	}
	end, ok := deleteRange["end"].(map[string]any)
	if !ok {
		t.Fatalf("Expected deleteRange.end map, got %T", deleteRange["end"])
	}
	if start["line"] != float64(1) || start["character"] != float64(0) {
		t.Errorf("deleteRange.start wrong: %v", start)
	}
	if end["line"] != float64(2) || end["character"] != float64(0) {
		t.Errorf("deleteRange.end wrong: %v", end)
	}
}

func TestCssDiagnostics_RoundTripCodeAction(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	uri := "file:///test.css"
	css := ".foo {\n  /** Blue colors */\n  background: var(--blue) var(--dark-blue);\n}\n"
	doc := dm.OpenDocument(uri, css, 1)
	ctx.AddDocument(uri, doc)

	diags := publishDiagnostics.AnalyzeCssDiagnosticsForTest(ctx, doc)
	if len(diags) != 1 {
		t.Fatalf("Expected 1 diagnostic, got %d", len(diags))
	}

	serialized, err := protocol.Marshal(diags[0])
	if err != nil {
		t.Fatalf("Failed to marshal diagnostic: %v", err)
	}
	var roundTripped protocol.Diagnostic
	if err := protocol.Unmarshal(serialized, &roundTripped); err != nil {
		t.Fatalf("Failed to unmarshal diagnostic: %v", err)
	}

	params := &protocol.CodeActionParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: urilib.URI(uri)},
		Context:      protocol.CodeActionContext{Diagnostics: []protocol.Diagnostic{roundTripped}},
	}

	result, err := codeAction.CodeAction(ctx, params)
	if err != nil {
		t.Fatalf("CodeAction failed: %v", err)
	}

	actions := result

	if len(actions) != 2 {
		t.Fatalf("Expected 2 code actions after round-trip, got %d", len(actions))
	}

	if actions[0].Title != "Associate comment with `--blue`" {
		t.Errorf("Action 0 title: got %q", actions[0].Title)
	}
	if actions[1].Title != "Associate comment with `--dark-blue`" {
		t.Errorf("Action 1 title: got %q", actions[1].Title)
	}

	// Verify edits exist and are well-formed
	for i, action := range actions {
		edits := action.Edit.Changes[urilib.URI(uri)]
		if len(edits) != 2 {
			t.Errorf("Action %d: expected 2 edits, got %d", i, len(edits))
		}
	}
}
