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

	"bennypowers.dev/cem/lsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestDocumentChangeHandling(t *testing.T) {
	dm, err := lsp.NewDocumentManager()
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
	if doc.GetContent() != initialContent {
		t.Errorf("Initial content mismatch. Expected: %s, Got: %s", initialContent, doc.GetContent())
	}

	// Update document directly - simulating content change
	updatedDoc := dm.UpdateDocument(uri, updatedContent, 2)
	if updatedDoc == nil {
		t.Fatal("Failed to update document")
	}

	// Verify document was updated
	if updatedDoc.GetContent() != updatedContent {
		t.Errorf("Updated content mismatch. Expected: %s, Got: %s", updatedContent, updatedDoc.GetContent())
	}
	if updatedDoc.GetVersion() != 2 {
		t.Errorf("Version mismatch. Expected: 2, Got: %d", updatedDoc.GetVersion())
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
