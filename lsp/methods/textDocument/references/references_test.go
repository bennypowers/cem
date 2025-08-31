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
	"testing"

	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/methods/textDocument/references"
	"bennypowers.dev/cem/lsp/testhelpers"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestReferences(t *testing.T) {
	// Create MockServerContext with proper DocumentManager
	ctx := testhelpers.NewMockServerContext()

	// Create DocumentManager and add documents
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	// Add documents using real DocumentManager
	doc1 := dm.OpenDocument("file:///test1.html", `<rh-card>content</rh-card>`, 1)
	doc2 := dm.OpenDocument("file:///test2.ts", "html`<rh-card variant=\"primary\"></rh-card>`", 1)

	ctx.AddDocument("file:///test1.html", doc1)
	ctx.AddDocument("file:///test2.ts", doc2)
	ctx.SetWorkspaceRoot("/test/workspace")

	// Create request for references to rh-card
	params := &protocol.ReferenceParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{
				URI: "file:///test1.html",
			},
			Position: protocol.Position{Line: 0, Character: 5}, // Inside rh-card
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

	if len(locations) != 2 {
		t.Fatalf("Expected 2 locations, got %d", len(locations))
	}

	// Check that both expected URIs are present (order doesn't matter)
	expectedURIs := map[string]bool{
		"file:///test1.html": false,
		"file:///test2.ts":   false,
	}

	for _, location := range locations {
		if _, exists := expectedURIs[location.URI]; exists {
			expectedURIs[location.URI] = true
		} else {
			t.Errorf("Unexpected URI: %s", location.URI)
		}
	}

	// Verify all expected URIs were found
	for uri, found := range expectedURIs {
		if !found {
			t.Errorf("Expected URI %s not found in results", uri)
		}
	}
}
