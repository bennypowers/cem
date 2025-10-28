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
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/methods/textDocument/references"
	"bennypowers.dev/cem/lsp/testhelpers"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestReferences(t *testing.T) {
	// Create MockServerContext with proper DocumentManager
	ctx := testhelpers.NewMockServerContext()

	// Create DocumentManager and add documents
	dm, err := document.NewDocumentManager()
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

func TestReferences_WorkspaceSearch(t *testing.T) {
	// Create MockServerContext
	ctx := testhelpers.NewMockServerContext()

	// Create DocumentManager
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	// Set workspace root to test fixtures directory
	ctx.SetWorkspaceRoot("test-fixtures/workspace-search")

	// Open one document (the others will be searched from disk)
	doc := dm.OpenDocument("file:///test-fixtures/workspace-search/index.html", `<rh-card>test</rh-card>`, 1)
	ctx.AddDocument("file:///test-fixtures/workspace-search/index.html", doc)

	// Create request for references to rh-card
	params := &protocol.ReferenceParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{
				URI: "file:///test-fixtures/workspace-search/index.html",
			},
			Position: protocol.Position{Line: 0, Character: 5}, // Inside rh-card
		},
		Context: protocol.ReferenceContext{
			IncludeDeclaration: true,
		},
	}

	// Call References function - should search both open doc and workspace files
	locations, err := references.References(ctx, &glsp.Context{}, params)

	// Verify results
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Should find references in:
	// 1. Open document (index.html)
	// 2. component.ts (found via workspace search)
	// Note: We expect at least 2 references, but might find more depending on file contents
	if len(locations) < 2 {
		t.Errorf("Expected at least 2 locations (open doc + workspace), got %d", len(locations))
	}

	// Verify we found references in both HTML and TypeScript files
	// Note: URIs will be absolute paths since workspace root gets converted to absolute
	foundHTML := false
	foundTS := false
	for _, location := range locations {
		// Check by filename since URIs will be absolute paths
		if filepath.Base(string(location.URI)) == "index.html" {
			foundHTML = true
		}
		if filepath.Base(string(location.URI)) == "component.ts" {
			foundTS = true
		}
	}

	if !foundHTML {
		t.Error("Expected to find references in index.html")
	}
	if !foundTS {
		t.Error("Expected to find references in component.ts via workspace search")
	}
}

func TestReferences_GitignoreFiltering(t *testing.T) {
	// Create MockServerContext
	ctx := testhelpers.NewMockServerContext()

	// Create DocumentManager
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	// Set workspace root to test fixtures directory (has .gitignore)
	ctx.SetWorkspaceRoot("test-fixtures/workspace-search")

	// Open a document to trigger search
	doc := dm.OpenDocument("file:///test-fixtures/workspace-search/index.html", `<rh-card>test</rh-card>`, 1)
	ctx.AddDocument("file:///test-fixtures/workspace-search/index.html", doc)

	// Create request
	params := &protocol.ReferenceParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{
				URI: "file:///test-fixtures/workspace-search/index.html",
			},
			Position: protocol.Position{Line: 0, Character: 5},
		},
		Context: protocol.ReferenceContext{
			IncludeDeclaration: true,
		},
	}

	// Call References function
	locations, err := references.References(ctx, &glsp.Context{}, params)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify that ignored/should-be-skipped.html is NOT in results
	for _, location := range locations {
		if location.URI == "file:///test-fixtures/workspace-search/ignored/should-be-skipped.html" {
			t.Error("Found reference in gitignored file - should have been filtered out")
		}
	}
}

func TestReferences_NoElement(t *testing.T) {
	// Create MockServerContext
	ctx := testhelpers.NewMockServerContext()

	// Create DocumentManager
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	// Open document with no custom elements at cursor
	doc := dm.OpenDocument("file:///test.html", `<div>regular html</div>`, 1)
	ctx.AddDocument("file:///test.html", doc)

	// Create request at position with no element
	params := &protocol.ReferenceParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{
				URI: "file:///test.html",
			},
			Position: protocol.Position{Line: 0, Character: 2}, // Inside <div>
		},
		Context: protocol.ReferenceContext{
			IncludeDeclaration: true,
		},
	}

	// Call References function
	locations, err := references.References(ctx, &glsp.Context{}, params)

	// Should return empty array, not error
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(locations) != 0 {
		t.Errorf("Expected 0 locations for non-custom-element, got %d", len(locations))
	}
}

func TestReferences_DocumentNotFound(t *testing.T) {
	// Create MockServerContext without adding document
	ctx := testhelpers.NewMockServerContext()

	// Create DocumentManager
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	// Create request for non-existent document
	params := &protocol.ReferenceParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{
				URI: "file:///non-existent.html",
			},
			Position: protocol.Position{Line: 0, Character: 5},
		},
		Context: protocol.ReferenceContext{
			IncludeDeclaration: true,
		},
	}

	// Call References function
	locations, err := references.References(ctx, &glsp.Context{}, params)

	// Should return empty array, not error
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(locations) != 0 {
		t.Errorf("Expected 0 locations for non-existent document, got %d", len(locations))
	}
}
