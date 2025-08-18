package references_test

import (
	"testing"

	"bennypowers.dev/cem/lsp/methods/textDocument/references"
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

// MockDocument implements types.Document for testing
type MockDocument struct {
	uri      string
	content  string
	elements []types.CustomElementMatch
}

func (m *MockDocument) FindElementAtPosition(position protocol.Position, dm any) *types.CustomElementMatch {
	return nil
}

func (m *MockDocument) FindAttributeAtPosition(position protocol.Position, dm any) (*types.AttributeMatch, string) {
	return nil, ""
}

func (m *MockDocument) Content() string {
	return m.content
}

func (m *MockDocument) Version() int32 {
	return 1
}

func (m *MockDocument) URI() string {
	return m.uri
}

func (m *MockDocument) FindCustomElements(dm any) ([]types.CustomElementMatch, error) {
	return m.elements, nil
}

func (m *MockDocument) AnalyzeCompletionContextTS(position protocol.Position, dm any) *types.CompletionAnalysis {
	// Mock implementation that recognizes when position is inside rh-card
	if position.Line == 0 && position.Character >= 1 && position.Character <= 8 {
		return &types.CompletionAnalysis{
			Type:    types.CompletionTagName,
			TagName: "rh-card",
		}
	}
	return nil
}

func (m *MockDocument) GetScriptTags() []types.ScriptTag {
	return nil
}

func (m *MockDocument) FindModuleScript() (protocol.Position, bool) {
	return protocol.Position{}, false
}

func (m *MockDocument) ByteRangeToProtocolRange(content string, startByte, endByte uint) protocol.Range {
	return protocol.Range{
		Start: protocol.Position{Line: 0, Character: 0},
		End:   protocol.Position{Line: 0, Character: 10},
	}
}

func TestReferences(t *testing.T) {
	// Create mock documents with rh-card elements
	doc1 := &MockDocument{
		uri:     "file:///test1.html",
		content: `<rh-card>content</rh-card>`,
		elements: []types.CustomElementMatch{
			{
				TagName: "rh-card",
				Range: protocol.Range{
					Start: protocol.Position{Line: 0, Character: 1},
					End:   protocol.Position{Line: 0, Character: 8},
				},
			},
		},
	}

	doc2 := &MockDocument{
		uri:     "file:///test2.ts",
		content: "html`<rh-card variant=\"primary\"></rh-card>`",
		elements: []types.CustomElementMatch{
			{
				TagName: "rh-card",
				Range: protocol.Range{
					Start: protocol.Position{Line: 0, Character: 6},
					End:   protocol.Position{Line: 0, Character: 13},
				},
			},
		},
	}

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
