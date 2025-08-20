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
package symbol_test

import (
	"testing"

	"bennypowers.dev/cem/lsp/methods/workspace/symbol"
	"bennypowers.dev/cem/lsp/testhelpers"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestWorkspaceSymbol(t *testing.T) {
	// Set up mock context
	ctx := testhelpers.NewMockServerContext()
	ctx.TagNames = []string{
		"my-button",
		"my-card",
		"my-dialog",
		"ui-icon",
		"ui-table",
	}
	ctx.SetWorkspaceRoot("/workspace")

	// Set up element sources
	sources := map[string]string{
		"my-button": "src/components/my-button.ts",
		"my-card":   "src/components/my-card.ts",
		"my-dialog": "src/dialogs/my-dialog.ts",
		"ui-icon":   "/absolute/path/to/ui-icon.ts",
		"ui-table":  "components/ui-table.js",
	}
	for tagName, source := range sources {
		ctx.AddElementDefinition(tagName, &testhelpers.MockElementDefinition{
			ModulePathStr: source,
		})
	}

	// Set up descriptions
	descriptions := map[string]string{
		"my-button": "A reusable button component",
		"my-card":   "Card layout component",
		"ui-icon":   "Icon display component",
	}
	for tagName, desc := range descriptions {
		ctx.AddElementDescription(tagName, desc)
	}

	mockGlspContext := &glsp.Context{}

	tests := []struct {
		name          string
		query         string
		expectedCount int
		expectedNames []string
	}{
		{
			name:          "Empty query returns all symbols",
			query:         "",
			expectedCount: 5,
			expectedNames: []string{"my-button - A reusable button component", "my-card - Card layout component", "my-dialog", "ui-icon - Icon display component", "ui-table"},
		},
		{
			name:          "Query 'my' filters correctly",
			query:         "my",
			expectedCount: 3,
			expectedNames: []string{"my-button - A reusable button component", "my-card - Card layout component", "my-dialog"},
		},
		{
			name:          "Query 'ui' filters correctly",
			query:         "ui",
			expectedCount: 2,
			expectedNames: []string{"ui-icon - Icon display component", "ui-table"},
		},
		{
			name:          "Case insensitive search",
			query:         "MY",
			expectedCount: 3,
			expectedNames: []string{"my-button - A reusable button component", "my-card - Card layout component", "my-dialog"},
		},
		{
			name:          "Specific element name",
			query:         "button",
			expectedCount: 1,
			expectedNames: []string{"my-button - A reusable button component"},
		},
		{
			name:          "No matches",
			query:         "nonexistent",
			expectedCount: 0,
			expectedNames: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Execute workspace symbol request
			params := &protocol.WorkspaceSymbolParams{
				Query: tt.query,
			}

			symbols, err := symbol.Symbol(ctx, mockGlspContext, params)
			if err != nil {
				t.Fatalf("Symbol() error = %v", err)
			}

			// Check count
			if len(symbols) != tt.expectedCount {
				t.Errorf("Expected %d symbols, got %d", tt.expectedCount, len(symbols))
			}

			// Check names
			actualNames := make([]string, len(symbols))
			for i, symbol := range symbols {
				actualNames[i] = symbol.Name
			}

			for _, expectedName := range tt.expectedNames {
				found := false
				for _, actualName := range actualNames {
					if actualName == expectedName {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected symbol '%s' not found in results: %v", expectedName, actualNames)
				}
			}
		})
	}
}

func TestWorkspaceSymbolProperties(t *testing.T) {
	// Test symbol properties like kind, location, etc.
	ctx := testhelpers.NewMockServerContext()
	ctx.TagNames = []string{"test-element"}
	ctx.SetWorkspaceRoot("/workspace")
	ctx.AddElementDefinition("test-element", &testhelpers.MockElementDefinition{
		ModulePathStr: "src/test-element.ts",
		SourceHrefStr: "src/test-element.ts",
	})
	ctx.AddElementDescription("test-element", "Test element for verification")

	mockGlspContext := &glsp.Context{}
	params := &protocol.WorkspaceSymbolParams{Query: "test"}

	symbols, err := symbol.Symbol(ctx, mockGlspContext, params)
	if err != nil {
		t.Fatalf("Symbol() error = %v", err)
	}

	if len(symbols) != 1 {
		t.Fatalf("Expected 1 symbol, got %d", len(symbols))
	}

	symbol := symbols[0]

	// Check symbol kind
	if symbol.Kind != protocol.SymbolKindClass {
		t.Errorf("Expected SymbolKindClass, got %v", symbol.Kind)
	}

	// Check name includes description
	expectedName := "test-element - Test element for verification"
	if symbol.Name != expectedName {
		t.Errorf("Expected name '%s', got '%s'", expectedName, symbol.Name)
	}

	// Check location URI
	expectedURI := "file:///workspace/src/test-element.ts"
	if symbol.Location.URI != expectedURI {
		t.Errorf("Expected URI '%s', got '%s'", expectedURI, symbol.Location.URI)
	}

	// Check container name
	expectedContainer := "src/test-element.ts"
	if symbol.ContainerName == nil || *symbol.ContainerName != expectedContainer {
		var actualContainer string
		if symbol.ContainerName != nil {
			actualContainer = *symbol.ContainerName
		}
		t.Errorf("Expected container '%s', got '%s'", expectedContainer, actualContainer)
	}
}

func TestWorkspaceSymbolAbsolutePaths(t *testing.T) {
	// Test handling of absolute paths
	ctx := testhelpers.NewMockServerContext()
	ctx.TagNames = []string{"abs-element"}
	ctx.SetWorkspaceRoot("/workspace")
	ctx.AddElementDefinition("abs-element", &testhelpers.MockElementDefinition{
		ModulePathStr: "/absolute/path/to/element.ts",
	})

	mockGlspContext := &glsp.Context{}
	params := &protocol.WorkspaceSymbolParams{Query: "abs"}

	symbols, err := symbol.Symbol(ctx, mockGlspContext, params)
	if err != nil {
		t.Fatalf("Symbol() error = %v", err)
	}

	if len(symbols) != 1 {
		t.Fatalf("Expected 1 symbol, got %d", len(symbols))
	}

	symbol := symbols[0]

	// Check absolute path is used directly
	expectedURI := "file:///absolute/path/to/element.ts"
	if symbol.Location.URI != expectedURI {
		t.Errorf("Expected URI '%s', got '%s'", expectedURI, symbol.Location.URI)
	}
}

func TestWorkspaceSymbolNoSource(t *testing.T) {
	// Test elements without source information
	ctx := testhelpers.NewMockServerContext()
	ctx.TagNames = []string{"no-source-element"}
	ctx.SetWorkspaceRoot("/workspace")
	// Intentionally not adding element definition - no source for this element

	mockGlspContext := &glsp.Context{}
	params := &protocol.WorkspaceSymbolParams{Query: "no-source"}

	symbols, err := symbol.Symbol(ctx, mockGlspContext, params)
	if err != nil {
		t.Fatalf("Symbol() error = %v", err)
	}

	if len(symbols) != 1 {
		t.Fatalf("Expected 1 symbol, got %d", len(symbols))
	}

	symbol := symbols[0]

	// Check element is still included but with empty URI
	if symbol.Location.URI != "" {
		t.Errorf("Expected empty URI for element without source, got '%s'", symbol.Location.URI)
	}

	// Name should just be tag name without description
	if symbol.Name != "no-source-element" {
		t.Errorf("Expected name 'no-source-element', got '%s'", symbol.Name)
	}
}
