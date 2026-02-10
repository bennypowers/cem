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
	"strings"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/methods/textDocument/publishDiagnostics"
	"bennypowers.dev/cem/lsp/testhelpers"
	M "bennypowers.dev/cem/manifest"
)

// TestDiagnostics_NestedElementRanges verifies that tag diagnostics in
// TypeScript template literals produce ranges covering the tag name, not
// the entire template. This is a regression test for the bug where
// adjustRangeToTemplate returned the whole template range.
func TestDiagnostics_NestedElementRanges(t *testing.T) {
	testutil.RunLSPFixtures(t, "testdata-regression", func(t *testing.T, fixture *testutil.LSPFixture) {
		if fixture.Name != "nested-element-diagnostics-typescript" {
			t.Skip("not applicable to this test")
		}

		ctx := testhelpers.NewMockServerContext()
		if fixture.Manifest != nil {
			var pkg M.Package
			if err := json.Unmarshal(fixture.Manifest, &pkg); err != nil {
				t.Fatalf("Failed to parse manifest: %v", err)
			}
			ctx.AddManifest(&pkg)
		}

		dm, err := document.NewDocumentManager()
		if err != nil {
			t.Fatalf("Failed to create DocumentManager: %v", err)
		}
		defer dm.Close()
		ctx.SetDocumentManager(dm)

		uri := "file:///test.ts"
		doc := dm.OpenDocument(uri, fixture.InputContent, 1)
		ctx.AddDocument(uri, doc)

		diagnostics := publishDiagnostics.AnalyzeTagNameDiagnosticsForTest(ctx, doc)

		// We expect diagnostics for both elements:
		// - outer-surface: known but not imported
		// - inner-icon: unknown element
		// The critical assertion: NONE should span the entire template range
		// (lines 4-10). Each should cover only the tag name on its own line.

		// input.ts template layout:
		// Line 5:  "      <outer-surface id="container" role="alert">"
		//           outer-surface at chars 7-20
		// Line 7:  "          <inner-icon id="icon" set="ui" icon="check"></inner-icon>"
		//           inner-icon at chars 11-21

		expectedRanges := map[string]struct {
			line      uint32
			startChar uint32
			endChar   uint32
		}{
			"outer-surface": {line: 5, startChar: 7, endChar: 20},
			"inner-icon":    {line: 7, startChar: 11, endChar: 21},
		}

		for tagName, expected := range expectedRanges {
			t.Run(tagName, func(t *testing.T) {
				// Find the diagnostic whose range matches the expected line
				found := false
				for _, diag := range diagnostics {
					if diag.Range.Start.Line != expected.line {
						continue
					}
					found = true
					if diag.Range.End.Line != expected.line {
						t.Errorf("Expected diagnostic to end on line %d, got line %d",
							expected.line, diag.Range.End.Line)
					}
					if diag.Range.Start.Character != expected.startChar {
						t.Errorf("Expected start character %d, got %d",
							expected.startChar, diag.Range.Start.Character)
					}
					if diag.Range.End.Character != expected.endChar {
						t.Errorf("Expected end character %d, got %d",
							expected.endChar, diag.Range.End.Character)
					}
					if !strings.Contains(diag.Message, "'"+tagName+"'") {
						t.Errorf("Expected diagnostic message to mention '%s', got: %s",
							tagName, diag.Message)
					}
					break
				}
				if !found {
					t.Errorf("No diagnostic found on line %d for %s", expected.line, tagName)
				}
			})
		}
	})
}
