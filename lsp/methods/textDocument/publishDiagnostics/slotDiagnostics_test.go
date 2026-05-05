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

	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/methods/textDocument/publishDiagnostics"
	"bennypowers.dev/cem/lsp/testhelpers"
	"bennypowers.dev/cem/manifest"
	protocol "github.com/bennypowers/glsp/protocol_3_17"
)

// SlotManifestFixture represents the manifest format for slot diagnostics tests
type SlotManifestFixture struct {
	Slots map[string][]manifest.Slot `json:"slots"`
}

// TestSlotDiagnostics_Fixtures runs slot diagnostics tests from fixtures.
// Regression test for issue #249: slot diagnostics validated against wrong
// parent custom element in nested structures.
func TestSlotDiagnostics_Fixtures(t *testing.T) {
	testutil.RunLSPFixtures(t, "testdata/slot-diagnostics", func(t *testing.T, fixture *testutil.LSPFixture) {
		ctx := testhelpers.NewMockServerContext()

		dm, err := document.NewDocumentManager()
		if err != nil {
			t.Fatalf("Failed to create DocumentManager: %v", err)
		}
		defer dm.Close()
		ctx.SetDocumentManager(dm)

		doc := dm.OpenDocument("test.html", fixture.InputHTML, 1)
		ctx.AddDocument("test.html", doc)

		if len(fixture.Manifest) > 0 {
			var manifestData SlotManifestFixture
			if err := json.Unmarshal(fixture.Manifest, &manifestData); err != nil {
				t.Fatalf("Failed to parse manifest.json: %v", err)
			}
			for tagName, slots := range manifestData.Slots {
				ctx.AddSlots(tagName, slots)
			}
		}

		diagnostics := publishDiagnostics.AnalyzeSlotDiagnosticsForTest(ctx, doc)

		var expected []protocol.Diagnostic
		if err := fixture.GetExpected("expected", &expected); err != nil {
			t.Fatalf("Failed to load expected diagnostics: %v", err)
		}

		if len(diagnostics) != len(expected) {
			t.Errorf("Expected %d diagnostics, got %d", len(expected), len(diagnostics))
			for i, diag := range diagnostics {
				t.Errorf("  Diagnostic %d: %s (line %d)", i, diag.Message, diag.Range.Start.Line)
			}
			return
		}

		for i, exp := range expected {
			act := diagnostics[i]

			if act.Range.Start.Line != exp.Range.Start.Line ||
				act.Range.Start.Character != exp.Range.Start.Character ||
				act.Range.End.Line != exp.Range.End.Line ||
				act.Range.End.Character != exp.Range.End.Character {
				t.Errorf("Diagnostic %d: expected range %v, got %v", i, exp.Range, act.Range)
			}

			if act.Message != exp.Message {
				t.Errorf("Diagnostic %d: expected message %q, got %q", i, exp.Message, act.Message)
			}
		}
	})
}
