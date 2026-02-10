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

// TestHover_NestedElements tests that hovering on a custom element nested
// inside another custom element in a TypeScript template literal returns
// hover docs for the inner element, not the outer one.
// This is a regression test for the bug where hovering on <rh-icon> inside
// <rh-surface> in rh-alert.ts showed docs for rh-surface.
func TestHover_NestedElements(t *testing.T) {
	testutil.RunLSPFixtures(t, "testdata-regression", func(t *testing.T, fixture *testutil.LSPFixture) {
		if fixture.Name != "nested-element-hover-typescript" {
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

		// Line 7 of input.ts:
		//           <inner-icon id="icon" set="ui" icon="check"></inner-icon>
		// char:      0123456789...
		// inner-icon tag name is at characters 11-20
		testCases := []struct {
			name     string
			position protocol.Position
		}{
			// Hover on the inner-icon tag name (character 15 is within "inner-icon")
			{"element", protocol.Position{Line: 7, Character: 15}},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
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

				if result == nil {
					t.Fatal("Expected hover result, got nil")
				}

				actualContents, ok := result.Contents.(protocol.MarkupContent)
				if !ok {
					t.Fatalf("Expected Contents to be MarkupContent, got %T", result.Contents)
				}

				// Unmarshal expected golden into a typed struct so Contents
				// deserializes as MarkupContent rather than map[string]interface{}.
				var expected struct {
					Contents protocol.MarkupContent `json:"contents"`
					Range    *protocol.Range        `json:"range"`
				}
				if err := fixture.GetExpected(tc.name, &expected); err != nil {
					t.Fatalf("Failed to get expected hover: %v", err)
				}

				if actualContents.Value != expected.Contents.Value {
					t.Errorf("Expected value:\n%s\n\nGot:\n%s", expected.Contents.Value, actualContents.Value)
				}
			})
		}
	})
}
