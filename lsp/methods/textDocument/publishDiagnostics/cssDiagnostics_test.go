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
	"testing"

	_ "bennypowers.dev/cem/internal/languages/registry"
	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/methods/textDocument/publishDiagnostics"
	"bennypowers.dev/cem/lsp/testhelpers"
	protocol "github.com/bennypowers/glsp/protocol_3_17"
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

			if exp.Severity != nil && (act.Severity == nil || *act.Severity != *exp.Severity) {
				t.Errorf("Diagnostic %d: expected severity %v, got %v", i, *exp.Severity, act.Severity)
			}
		}
	})
}
