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
package publishDiagnostics_test

import (
	"testing"

	"bennypowers.dev/cem/lsp/testhelpers"
	"bennypowers.dev/cem/manifest"
)

func TestSlotDiagnosticsBasic(t *testing.T) {
	// This is a placeholder test to demonstrate the slot diagnostics structure
	// Full integration testing should be done at the LSP server level

	// Create test context using MockServerContext
	ctx := testhelpers.NewMockServerContext()
	
	// Add test document
	doc := testhelpers.NewMockDocument(`<my-element><div slot="heade">Content</div></my-element>`)
	ctx.AddDocument("test://test.html", doc)
	
	// Create registry with test slots
	registry := testhelpers.NewMockRegistry()
	registry.AddSlots("my-element", []manifest.Slot{
		{FullyQualified: manifest.FullyQualified{Name: "header"}},
		{FullyQualified: manifest.FullyQualified{Name: "footer"}},
	})
	ctx.SetRegistry(registry)

	// Verify mock context works
	retrievedDoc := ctx.Document("test://test.html")
	if retrievedDoc == nil {
		t.Fatal("Expected document to be found")
	}

	content, err := retrievedDoc.Content()
	if err != nil {
		t.Fatalf("Failed to get document content: %v", err)
	}
	if content != `<my-element><div slot="heade">Content</div></my-element>` {
		t.Errorf("Expected content to match, got: %s", content)
	}

	slots, exists := ctx.Slots("my-element")
	if !exists {
		t.Fatal("Expected slots to exist for my-element")
	}

	if len(slots) != 2 {
		t.Errorf("Expected 2 slots, got %d", len(slots))
	}

	if slots[0].Name != "header" {
		t.Errorf("Expected first slot to be 'header', got: %s", slots[0].Name)
	}
}

