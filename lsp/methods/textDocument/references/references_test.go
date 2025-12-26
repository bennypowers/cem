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
	"io/fs"
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
	"basic-references":      {Line: 0, Character: 5},
	"workspace-search":      {Line: 0, Character: 5},
	"gitignore-filtering":   {Line: 0, Character: 5},
	"no-element":            {Line: 0, Character: 5},
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
			if fixture.Name == "workspace-search" || fixture.Name == "gitignore-filtering" {
				uri = "file:///index.html"
			} else {
				uri = "file:///test1.html"
			}
		}

		// Open the main document
		doc := dm.OpenDocument(uri, fixture.InputContent, 1)
		ctx.AddDocument(uri, doc)

		// Create in-memory filesystem with predictable paths
		workspaceFiles := make(map[string]string)
		workspaceDir := filepath.Join("testdata", fixture.Name, "workspace")

		// Load all workspace files into the in-memory filesystem
		err = filepath.WalkDir(workspaceDir, func(path string, d fs.DirEntry, err error) error {
			if err != nil || d.IsDir() {
				return err
			}

			// Read file content from disk
			contentBytes, readErr := os.ReadFile(path)
			if readErr != nil {
				return readErr
			}

			// Get relative path from workspace directory
			relPath, err := filepath.Rel(workspaceDir, path)
			if err != nil {
				return err
			}

			// Store in map with clean path (e.g., "component.ts", "ignored/should-be-skipped.html")
			workspaceFiles[relPath] = string(contentBytes)
			return nil
		})
		if err != nil && err != filepath.SkipDir {
			// Workspace directory might not exist for some tests (like no-element)
			// That's okay
		}

		// Create MapFS and set it on the context
		mapFS := platform.NewMapFS(workspaceFiles)
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
					actual.Range.Start.Character == expected.Range.Start.Character {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Expected location not found: %s at %d:%d",
					expected.URI, expected.Range.Start.Line, expected.Range.Start.Character)
			}
		}
	})
}
