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
package references_test

import (
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/methods/textDocument/references"
	"bennypowers.dev/cem/lsp/testhelpers"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// Cursor positions for each test fixture (in Go code, not JSON)
var cursorPositions = map[string]protocol.Position{
	"basic-references":    {Line: 0, Character: 5},
	"workspace-search":    {Line: 0, Character: 5},
	"gitignore-filtering": {Line: 0, Character: 5},
	"no-element":          {Line: 0, Character: 5},
	"missing-document":    {Line: 0, Character: 0},
}

func TestReferences_Fixtures(t *testing.T) {
	testutil.RunLSPFixtures(t, "testdata", func(t *testing.T, fixture *testutil.LSPFixture) {
		// Get cursor position from map
		cursor, ok := cursorPositions[fixture.Name]
		if !ok {
			t.Fatalf("No cursor position defined for fixture %s", fixture.Name)
		}

		// Create MockServerContext with proper DocumentManager
		ctx := testhelpers.NewMockServerContext()

		// Create DocumentManager and add documents
		dm, err := document.NewDocumentManager()
		if err != nil {
			t.Fatalf("Failed to create DocumentManager: %v", err)
		}
		defer dm.Close()
		ctx.SetDocumentManager(dm)

		// Determine URI based on input type and scenario
		var uri string
		if fixture.InputType == "ts" {
			uri = "file:///test.ts"
		} else {
			// Use scenario-specific filename for HTML
			switch fixture.Name {
			case "workspace-search", "gitignore-filtering":
				uri = "file:///index.html"
			case "missing-document":
				uri = "file:///missing.html"
			default:
				uri = "file:///test1.html"
			}
		}

		// Open the main document (except for missing-document fixture)
		if fixture.Name != "missing-document" {
			doc := dm.OpenDocument(uri, fixture.InputContent, 1)
			ctx.AddDocument(uri, doc)
		}

		// Create in-memory filesystem with workspace files
		// Check if workspace directory exists first
		workspaceDir := filepath.Join("testdata", fixture.Name, "workspace")
		var mapFS platform.FileSystem
		if _, err := os.Stat(workspaceDir); os.IsNotExist(err) {
			// Workspace directory doesn't exist (e.g., no-element test) - use empty filesystem
			mapFS = platform.NewMapFS(nil)
		} else if err != nil {
			// Other stat error
			t.Logf("Warning: error checking workspace directory: %v", err)
			mapFS = platform.NewMapFS(nil)
		} else {
			// Workspace exists - load fixtures using testutil helper
			mapFS = testutil.NewFixtureFS(t, filepath.Join(fixture.Name, "workspace"), ".")
		}

		ctx.SetFileSystem(mapFS)
		ctx.SetWorkspaceRoot(".")

		// Load expected LSP response (Location[])
		var expectedLocations []protocol.Location
		err = fixture.GetExpected("expected", &expectedLocations)
		if err != nil {
			t.Fatalf("Failed to get expected locations: %v", err)
		}

		// Create reference params
		params := &protocol.ReferenceParams{
			TextDocumentPositionParams: protocol.TextDocumentPositionParams{
				TextDocument: protocol.TextDocumentIdentifier{
					URI: uri,
				},
				Position: cursor,
			},
			Context: protocol.ReferenceContext{
				IncludeDeclaration: true,
			},
		}

		// Call References function
		locations, err := references.References(ctx, &glsp.Context{}, params)

		// Verify results
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Verify count matches
		if len(locations) != len(expectedLocations) {
			t.Errorf("Expected %d locations, got %d", len(expectedLocations), len(locations))
			t.Logf("Actual locations: %+v", locations)
		}

		// Verify each expected location is present
		for _, expected := range expectedLocations {
			found := false
			for _, actual := range locations {
				if actual.URI == expected.URI &&
					actual.Range.Start.Line == expected.Range.Start.Line &&
					actual.Range.Start.Character == expected.Range.Start.Character &&
					actual.Range.End.Line == expected.Range.End.Line &&
					actual.Range.End.Character == expected.Range.End.Character {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Expected location not found: %s at %d:%d-%d:%d",
					expected.URI, expected.Range.Start.Line, expected.Range.Start.Character,
					expected.Range.End.Line, expected.Range.End.Character)
			}
		}
	})
}
