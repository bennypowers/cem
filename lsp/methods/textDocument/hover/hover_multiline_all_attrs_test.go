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
package hover_test

import (
	"encoding/json"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/methods/textDocument/hover"
	"bennypowers.dev/cem/lsp/testhelpers"
	M "bennypowers.dev/cem/manifest"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// TestHover_MultilineAllAttributes tests hover on all attributes that failed in the benchmark.
// This is a regression test for issues #176 and #179, ensuring our attribute value matching fix
// works for ALL attributes on a multi-line element.
//
// Benchmark failures (on docs/lsp/benchmarks branch) showed these 5 attributes all failed:
// - variant, size, disabled, loading, icon
//
// Root cause: Benchmark test positions were outdated, but this test validates the fix works.
func TestHover_MultilineAllAttributes(t *testing.T) {
	testutil.RunLSPFixtures(t, "testdata-regression", func(t *testing.T, fixture *testutil.LSPFixture) {
		if fixture.Name != "multiline-all-attributes" {
			t.Skip("not applicable to this test")
		}

		// Setup context
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

		uri := "file:///test.html"
		doc := dm.OpenDocument(uri, fixture.InputContent, 1)
		ctx.AddDocument(uri, doc)

		// Test each attribute that failed in the benchmark
		testCases := []struct {
			name     string
			position protocol.Position
		}{
			{"variant", protocol.Position{Line: 6, Character: 13}},
			{"size", protocol.Position{Line: 7, Character: 13}},
			{"disabled", protocol.Position{Line: 8, Character: 13}},
			{"loading", protocol.Position{Line: 9, Character: 13}},
			{"icon", protocol.Position{Line: 10, Character: 13}},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				// Load expected result from fixture
				var expected protocol.Hover
				if err := fixture.GetExpected(tc.name, &expected); err != nil {
					t.Fatalf("Failed to get expected hover: %v", err)
				}

				// Call hover
				params := &protocol.HoverParams{
					TextDocumentPositionParams: protocol.TextDocumentPositionParams{
						TextDocument: protocol.TextDocumentIdentifier{URI: uri},
						Position:     tc.position,
					},
				}

				result, err := hover.Hover(ctx, nil, params)
				if err != nil {
					t.Fatalf("Hover failed: %v", err)
				}

				// Structured comparison
				actualContents, ok := result.Contents.(protocol.MarkupContent)
				if !ok {
					t.Fatalf("Expected Contents to be MarkupContent, got %T", result.Contents)
				}

				expectedContents, ok := expected.Contents.(protocol.MarkupContent)
				if !ok {
					t.Fatalf("Expected Contents in expected.json to be MarkupContent, got %T", expected.Contents)
				}

				if actualContents.Kind != expectedContents.Kind {
					t.Errorf("Expected kind %s, got %s", expectedContents.Kind, actualContents.Kind)
				}

				if actualContents.Value != expectedContents.Value {
					t.Errorf("Expected value:\n%s\n\nGot:\n%s", expectedContents.Value, actualContents.Value)
				}

				// Verify range
				if expected.Range != nil && result.Range != nil {
					if result.Range.Start.Line != expected.Range.Start.Line ||
						result.Range.Start.Character != expected.Range.Start.Character ||
						result.Range.End.Line != expected.Range.End.Line ||
						result.Range.End.Character != expected.Range.End.Character {
						t.Errorf("Range mismatch.\nExpected: %+v\nGot: %+v", expected.Range, result.Range)
					}
				}
			})
		}
	})
}
