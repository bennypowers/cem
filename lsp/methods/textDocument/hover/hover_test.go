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

// Cursor positions for each test fixture
var cursorPositions = map[string]protocol.Position{
	"element-hover-html":              {Line: 6, Character: 5},
	"element-hover-typescript":        {Line: 5, Character: 10},
	"hover-attribute-html":            {Line: 6, Character: 20},
	"multiline-attributes-hover-html": {Line: 6, Character: 13},
}

func TestHover_Fixtures(t *testing.T) {
	testutil.RunLSPFixtures(t, "testdata", func(t *testing.T, fixture *testutil.LSPFixture) {
		// Get cursor position from map
		cursor, ok := cursorPositions[fixture.Name]
		if !ok {
			t.Fatalf("No cursor position defined for fixture %s", fixture.Name)
		}

		// Parse manifest if present
		ctx := testhelpers.NewMockServerContext()
		if fixture.Manifest != nil {
			var pkg M.Package
			err := json.Unmarshal(fixture.Manifest, &pkg)
			if err != nil {
				t.Fatalf("Failed to parse manifest: %v", err)
			}
			ctx.AddManifest(&pkg)
		}

		// Create document manager
		dm, err := document.NewDocumentManager()
		if err != nil {
			t.Fatalf("Failed to create DocumentManager: %v", err)
		}
		defer dm.Close()
		ctx.SetDocumentManager(dm)

		// Determine URI based on input type
		var uri string
		if fixture.InputType == "ts" {
			uri = "file:///test.ts"
		} else {
			uri = "file:///test.html"
		}

		// Open document
		doc := dm.OpenDocument(uri, fixture.InputContent, 1)
		ctx.AddDocument(uri, doc)

		// Load expected Hover response
		var expected protocol.Hover
		err = fixture.GetExpected("expected", &expected)
		if err != nil {
			t.Fatalf("Failed to get expected hover: %v", err)
		}

		// Create hover params
		params := &protocol.HoverParams{
			TextDocumentPositionParams: protocol.TextDocumentPositionParams{
				TextDocument: protocol.TextDocumentIdentifier{URI: uri},
				Position:     cursor,
			},
		}

		// Call hover
		result, err := hover.Hover(ctx, nil, params)
		if err != nil {
			t.Fatalf("Hover failed: %v", err)
		}

		if result == nil {
			t.Fatal("Expected hover result, got nil")
		}

		// Verify hover matches expected
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

		// Verify range if provided in expected
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
