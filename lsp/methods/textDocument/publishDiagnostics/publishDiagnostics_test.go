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
	"bennypowers.dev/cem/lsp/types"
	"bennypowers.dev/cem/manifest"
)

func TestSlotDiagnosticsBasic(t *testing.T) {
	// This is a placeholder test to demonstrate the slot diagnostics structure
	// Full integration testing should be done at the LSP server level

	// Test that our mock implementations work
	ctx := &mockDiagnosticsContext{
		documents: map[string]types.Document{
			"test://test.html": testhelpers.NewMockDocument(`<my-element><div slot="heade">Content</div></my-element>`),
		},
		slots: map[string][]manifest.Slot{
			"my-element": {
				{FullyQualified: manifest.FullyQualified{Name: "header"}},
				{FullyQualified: manifest.FullyQualified{Name: "footer"}},
			},
		},
	}

	// Verify mock context works
	doc := ctx.GetDocument("test://test.html")
	if doc == nil {
		t.Fatal("Expected document to be found")
	}

	content, err := doc.Content()
	if err != nil {
		t.Fatalf("Failed to get document content: %v", err)
	}
	if content != `<my-element><div slot="heade">Content</div></my-element>` {
		t.Errorf("Expected content to match, got: %s", content)
	}

	slots, exists := ctx.GetSlots("my-element")
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

// Mock implementations for testing

type mockDiagnosticsContext struct {
	documents map[string]types.Document
	slots     map[string][]manifest.Slot
	tagNames  []string
}

func (m *mockDiagnosticsContext) GetDocument(uri string) types.Document {
	return m.documents[uri]
}

func (m *mockDiagnosticsContext) GetSlots(tagName string) ([]manifest.Slot, bool) {
	slots, exists := m.slots[tagName]
	return slots, exists
}

func (m *mockDiagnosticsContext) AllTagNames() []string {
	return m.tagNames
}
