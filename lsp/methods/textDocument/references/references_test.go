package references_test

import (
	"testing"

	"bennypowers.dev/cem/lsp/methods/textDocument/references"
	"bennypowers.dev/cem/lsp/testhelpers"
	"bennypowers.dev/cem/lsp/types"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// MockReferencesContext implements types.ReferencesContext for testing
type MockReferencesContext struct {
	documents []types.Document
}

func (m *MockReferencesContext) Document(uri string) types.Document {
	for _, doc := range m.documents {
		if doc.URI() == uri {
			return doc
		}
	}
	return nil
}

func (m *MockReferencesContext) AllDocuments() []types.Document {
	return m.documents
}

func (m *MockReferencesContext) ElementDefinition(tagName string) (types.ElementDefinition, bool) {
	return nil, false
}

func (m *MockReferencesContext) WorkspaceRoot() string {
	return "/test/workspace"
}


func TestReferences(t *testing.T) {
	// Create mock documents with rh-card elements
	doc1 := testhelpers.NewMockDocumentWithElements(
		`<rh-card>content</rh-card>`,
		[]types.CustomElementMatch{
			{
				TagName: "rh-card",
				Range: protocol.Range{
					Start: protocol.Position{Line: 0, Character: 1},
					End:   protocol.Position{Line: 0, Character: 8},
				},
			},
		},
	)
	doc1.URIStr = "file:///test1.html"

	doc2 := testhelpers.NewMockDocumentWithElements(
		"html`<rh-card variant=\"primary\"></rh-card>`",
		[]types.CustomElementMatch{
			{
				TagName: "rh-card",
				Range: protocol.Range{
					Start: protocol.Position{Line: 0, Character: 6},
					End:   protocol.Position{Line: 0, Character: 13},
				},
			},
		},
	)
	doc2.URIStr = "file:///test2.ts"

	ctx := &MockReferencesContext{
		documents: []types.Document{doc1, doc2},
	}

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

	// Check first location (from test1.html)
	if locations[0].URI != "file:///test1.html" {
		t.Errorf("Expected URI file:///test1.html, got %s", locations[0].URI)
	}

	// Check second location (from test2.ts)
	if locations[1].URI != "file:///test2.ts" {
		t.Errorf("Expected URI file:///test2.ts, got %s", locations[1].URI)
	}
}
