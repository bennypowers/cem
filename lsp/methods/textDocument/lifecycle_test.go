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
package textDocument_test

import (
	"testing"

	_ "bennypowers.dev/cem/internal/languages/registry"
	"bennypowers.dev/cem/lsp/document"
	textDocument "bennypowers.dev/cem/lsp/methods/textDocument"
	"bennypowers.dev/cem/lsp/testhelpers"
	protocol "github.com/bennypowers/glsp/protocol_3_17"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDocumentChangeHandling(t *testing.T) {
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	// Test document URI and initial content
	uri := "file:///test.html"
	initialContent := `<div><my-element></my-element></div>`
	updatedContent := `<div><my-element attr="new"></my-element></div>`

	// Open document directly
	doc := dm.OpenDocument(uri, initialContent, 1)
	if doc == nil {
		t.Fatal("Failed to open document")
	}

	// Verify initial document
	docContent, err := doc.Content()
	if err != nil {
		t.Fatalf("Failed to get document content: %v", err)
	}
	if docContent != initialContent {
		t.Errorf("Initial content mismatch. Expected: %s, Got: %s", initialContent, docContent)
	}

	// Update document directly - simulating content change
	updatedDoc := dm.UpdateDocument(uri, updatedContent, 2)
	if updatedDoc == nil {
		t.Fatal("Failed to update document")
	}

	// Verify document was updated
	updatedDocContent, err := updatedDoc.Content()
	if err != nil {
		t.Fatalf("Failed to get updated document content: %v", err)
	}
	if updatedDocContent != updatedContent {
		t.Errorf("Updated content mismatch. Expected: %s, Got: %s", updatedContent, updatedDocContent)
	}
	if updatedDoc.Version() != 2 {
		t.Errorf("Version mismatch. Expected: 2, Got: %d", updatedDoc.Version())
	}

	// Verify that tree parsing still works after update
	elements, err := updatedDoc.FindCustomElements(dm)
	if err != nil {
		t.Fatalf("Failed to find custom elements after change: %v", err)
	}

	if len(elements) != 1 {
		t.Errorf("Expected 1 custom element, got %d", len(elements))
		return
	}

	element := elements[0]
	if element.TagName != "my-element" {
		t.Errorf("Expected tag name 'my-element', got '%s'", element.TagName)
	}

	// Verify the new attribute is detected
	if _, hasAttr := element.Attributes["attr"]; !hasAttr {
		t.Error("Expected 'attr' attribute to be detected after document change")
	}

	// Test that we can find the element at a position
	// Position should be somewhere in the tag name
	pos := protocol.Position{Line: 0, Character: 6} // Somewhere in "<my-element"
	foundElement := updatedDoc.FindElementAtPosition(pos, dm)
	if foundElement == nil {
		t.Error("Should find element at position after document change")
	} else if foundElement.TagName != "my-element" {
		t.Errorf("Found wrong element. Expected: my-element, Got: %s", foundElement.TagName)
	}
}

// TestNilDocumentManagerHandling tests the regression fix for nil pointer dereference
// when AnalyzeCompletionContextTS is called with nil document manager (e.g., from diagnostics)
func TestNilDocumentManagerHandling(t *testing.T) {
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	// Test HTML content that could trigger completion context analysis
	uri := "file:///test.html"
	htmlContent := `<my-element><div slot="header">Content</div></my-element>`

	// Open document
	doc := dm.OpenDocument(uri, htmlContent, 1)
	if doc == nil {
		t.Fatal("Failed to open document")
	}

	// Test that AnalyzeCompletionContextTS doesn't crash when called with nil document manager
	// This simulates what happens when diagnostics or other non-completion contexts call the function
	pos := protocol.Position{Line: 0, Character: 20} // Position in the slot attribute

	// This should not panic - previously would cause: panic: runtime error: invalid memory address or nil pointer dereference
	result := doc.AnalyzeCompletionContextTS(pos, nil)

	// The result might be nil or have minimal analysis, but it shouldn't crash
	// The important thing is that it doesn't panic
	if result != nil {
		t.Logf("Analysis result with nil dm: Type=%v, TagName=%s", result.Type, result.TagName)
	} else {
		t.Log("Analysis returned nil with nil dm (acceptable)")
	}

	// Verify the same call works fine with a proper document manager
	resultWithDM := doc.AnalyzeCompletionContextTS(pos, dm)
	if resultWithDM != nil {
		t.Logf("Analysis result with proper dm: Type=%v, TagName=%s", resultWithDM.Type, resultWithDM.TagName)
	}

	// The key assertion: no panic occurred during the nil dm call
	// If we reach this point, the test passes
}

func TestDidOpen(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	dm, err := document.NewDocumentManager()
	require.NoError(t, err)
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	params := &protocol.DidOpenTextDocumentParams{
		TextDocument: protocol.TextDocumentItem{
			URI:     "file:///test.html",
			Text:    "<my-element></my-element>",
			Version: 1,
		},
	}

	err = textDocument.DidOpen(ctx, nil, params)
	require.NoError(t, err)

	doc := dm.Document("file:///test.html")
	require.NotNil(t, doc)
	content, err := doc.Content()
	require.NoError(t, err)
	assert.Equal(t, "<my-element></my-element>", content)
}

func TestDidChange_FullReplacement(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	dm, err := document.NewDocumentManager()
	require.NoError(t, err)
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	dm.OpenDocument("file:///test.html", "<div>old</div>", 1)

	params := &protocol.DidChangeTextDocumentParams{
		TextDocument: protocol.VersionedTextDocumentIdentifier{
			TextDocumentIdentifier: protocol.TextDocumentIdentifier{URI: "file:///test.html"},
			Version:                2,
		},
		ContentChanges: []any{
			protocol.TextDocumentContentChangeEvent{
				Text: "<div>new</div>",
			},
		},
	}

	err = textDocument.DidChange(ctx, nil, params)
	require.NoError(t, err)

	doc := dm.Document("file:///test.html")
	require.NotNil(t, doc)
	content, err := doc.Content()
	require.NoError(t, err)
	assert.Equal(t, "<div>new</div>", content)
}

func TestDidChange_IncrementalEdit(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	dm, err := document.NewDocumentManager()
	require.NoError(t, err)
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	dm.OpenDocument("file:///test.html", "<div>hello</div>", 1)

	params := &protocol.DidChangeTextDocumentParams{
		TextDocument: protocol.VersionedTextDocumentIdentifier{
			TextDocumentIdentifier: protocol.TextDocumentIdentifier{URI: "file:///test.html"},
			Version:                2,
		},
		ContentChanges: []any{
			protocol.TextDocumentContentChangeEvent{
				Range: &protocol.Range{
					Start: protocol.Position{Line: 0, Character: 5},
					End:   protocol.Position{Line: 0, Character: 10},
				},
				Text: "world",
			},
		},
	}

	err = textDocument.DidChange(ctx, nil, params)
	require.NoError(t, err)

	doc := dm.Document("file:///test.html")
	require.NotNil(t, doc)
	content, err := doc.Content()
	require.NoError(t, err)
	assert.Equal(t, "<div>world</div>", content)
}

func TestDidChange_DocumentNotFound(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	dm, err := document.NewDocumentManager()
	require.NoError(t, err)
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	params := &protocol.DidChangeTextDocumentParams{
		TextDocument: protocol.VersionedTextDocumentIdentifier{
			TextDocumentIdentifier: protocol.TextDocumentIdentifier{URI: "file:///missing.html"},
			Version:                1,
		},
		ContentChanges: []any{},
	}

	err = textDocument.DidChange(ctx, nil, params)
	assert.NoError(t, err)
}

func TestDidClose(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	dm, err := document.NewDocumentManager()
	require.NoError(t, err)
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	dm.OpenDocument("file:///test.html", "<div>hello</div>", 1)
	assert.NotNil(t, dm.Document("file:///test.html"))

	params := &protocol.DidCloseTextDocumentParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: "file:///test.html"},
	}

	err = textDocument.DidClose(ctx, nil, params)
	require.NoError(t, err)
	assert.Nil(t, dm.Document("file:///test.html"))
}

func TestAnalyzeCompletionContext_NilDoc(t *testing.T) {
	_, err := textDocument.AnalyzeCompletionContext(nil, protocol.Position{}, "")
	assert.Error(t, err)
}

func TestAnalyzeCompletionContext_OutOfBounds(t *testing.T) {
	dm, err := document.NewDocumentManager()
	require.NoError(t, err)
	defer dm.Close()

	doc := dm.OpenDocument("file:///test.html", "single line", 1)
	_, err = textDocument.AnalyzeCompletionContext(doc, protocol.Position{Line: 99, Character: 0}, "")
	assert.Error(t, err)
}
