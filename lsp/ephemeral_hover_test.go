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
package lsp_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/methods/textDocument/hover"
	W "bennypowers.dev/cem/workspace"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// Cursor positions for each ephemeral hover test fixture
var ephemeralHoverPositions = map[string]protocol.Position{
	"element-hover":     {Line: 6, Character: 5},
	"attribute-hover":   {Line: 6, Character: 19},
	"ts-template-hover": {Line: 14, Character: 10},
}

func TestEphemeralHover(t *testing.T) {
	testutil.RunLSPFixtures(t, "testdata/ephemeral-hover", func(t *testing.T, fixture *testutil.LSPFixture) {
		cursor, ok := ephemeralHoverPositions[fixture.Name]
		if !ok {
			t.Fatalf("No cursor position defined for fixture %s", fixture.Name)
		}

		workspace := W.NewFileSystemWorkspaceContext("/test")
		if err := workspace.Init(); err != nil {
			t.Fatalf("Failed to init workspace: %v", err)
		}

		server, err := lsp.NewServer(workspace, lsp.TransportStdio)
		if err != nil {
			t.Fatalf("Failed to create server: %v", err)
		}
		defer func() { _ = server.Close() }()

		dm, err := server.DocumentManager()
		if err != nil {
			t.Fatalf("Failed to get document manager: %v", err)
		}

		// If there's an additional element.ts file, open and synthesize it first
		if elementTS, ok := fixture.AdditionalFiles["element.ts"]; ok {
			tsURI := "file:///element.ts"
			dm.OpenDocument(tsURI, elementTS, 1)
			server.SynthesizeEphemeralElements(tsURI)

		}

		// Open the primary input as a document
		var uri string
		if fixture.InputType == "ts" {
			uri = "file:///test.ts"

			// For TypeScript inputs, the input IS the element definition —
			// synthesize it for self-referencing hover
			dm.OpenDocument(uri, fixture.InputContent, 1)
			server.SynthesizeEphemeralElements(uri)

		} else {
			uri = "file:///test.html"
			dm.OpenDocument(uri, fixture.InputContent, 1)
		}

		params := &protocol.HoverParams{
			TextDocumentPositionParams: protocol.TextDocumentPositionParams{
				TextDocument: protocol.TextDocumentIdentifier{URI: uri},
				Position:     cursor,
			},
		}

		result, err := hover.Hover(server, nil, params)
		if err != nil {
			t.Fatalf("Hover failed: %v", err)
		}

		if result == nil {
			t.Fatal("Expected hover result, got nil")
		}

		// Compare against golden expected.json
		goldenPath := filepath.Join("testdata/ephemeral-hover", fixture.Name, "expected.json")

		actualJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			t.Fatalf("Failed to marshal hover result: %v", err)
		}

		if *testutil.Update {
			if err := os.WriteFile(goldenPath, actualJSON, 0644); err != nil {
				t.Fatalf("Failed to write golden file: %v", err)
			}
			t.Logf("Updated golden file: %s", goldenPath)
			return
		}

		var expected protocol.Hover
		if err := fixture.GetExpected("expected", &expected); err != nil {
			t.Fatalf("Failed to get expected hover: %v", err)
		}

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
