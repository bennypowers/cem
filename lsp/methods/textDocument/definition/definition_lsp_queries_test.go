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
package definition_test

import (
	"testing"

	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/methods/textDocument/definition"
	"bennypowers.dev/cem/lsp/testhelpers"
	"bennypowers.dev/cem/queries"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestDefinition_LSPQueriesIncludeClasses(t *testing.T) {
	// Regression test: Ensure LSPQueries includes the classes query needed for go-to-definition
	// This prevents the issue where go-to-definition went to top of file instead of actual definition

	// Verify classes query is included in LSP queries
	lspQueries := queries.LSPQueries()
	hasClassesQuery := false
	for _, query := range lspQueries.TypeScript {
		if query == "classes" {
			hasClassesQuery = true
			break
		}
	}

	if !hasClassesQuery {
		t.Fatal("LSPQueries must include 'classes' query for go-to-definition to work correctly")
	}

	t.Log("✅ LSPQueries correctly includes 'classes' query")
}

func TestDefinition_GoesToActualDefinitionNotTopOfFile(t *testing.T) {
	// Regression test for the specific issue: go-to-definition should go to the @customElement
	// decorator location, not the top of the file (line 0, character 0)

	ctx := testhelpers.NewMockServerContext()
	ctx.SetWorkspaceRoot("definition-test-fixtures")

	// Create a real DocumentManager with LSP queries (includes classes)
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	// Add HTML document with custom element
	htmlContent := `<!DOCTYPE html>
<html>
<body>
  <card-element></card-element>
</body>
</html>`

	uri := "file:///test.html"
	htmlDoc := dm.OpenDocument(uri, htmlContent, 1)
	ctx.Documents[uri] = htmlDoc

	// Add element definition pointing to TypeScript file
	elementDef := &testhelpers.MockElementDefinition{
		ModulePathStr:  "components/card-element.js", // Manifest points to .js
		PackageNameStr: "demo-project",
		SourceHrefStr:  "",
		ElementPtr:     nil,
	}
	ctx.ElementDefsMap["card-element"] = elementDef

	// Create definition request inside the custom element
	params := &protocol.DefinitionParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{URI: uri},
			Position:     protocol.Position{Line: 3, Character: 5}, // Inside <card-element>
		},
	}

	result, err := definition.Definition(ctx, nil, params)

	if err != nil {
		t.Fatalf("Definition request failed: %v", err)
	}

	if result == nil {
		t.Fatal("Expected definition result, got nil")
	}

	location, ok := result.(protocol.Location)
	if !ok {
		t.Fatalf("Expected protocol.Location, got %T", result)
	}

	// Should resolve to TypeScript file (path resolution working)
	expectedURI := "file://definition-test-fixtures/components/card-element.ts"
	if location.URI != expectedURI {
		t.Errorf("Expected URI %s, got %s", expectedURI, location.URI)
	}

	// CRITICAL: Should NOT go to top of file (the bug we fixed)
	if location.Range.Start.Line == 0 && location.Range.Start.Character == 0 {
		t.Error("REGRESSION: Definition went to top of file (0,0) instead of actual definition location")
		t.Error("This indicates the classes query is not available or tree-sitter queries are failing")
	}

	// Should go to actual @customElement decorator location (around line 10)
	expectedLine := uint32(10) // @customElement('card-element') is on line 10
	if location.Range.Start.Line != expectedLine {
		t.Errorf("Expected definition at line %d (@customElement decorator), got line %d",
			expectedLine, location.Range.Start.Line)
	}

	// Should go to the tag name string within the decorator
	expectedCharacter := uint32(16) // 'card-element' starts at character 16
	if location.Range.Start.Character != expectedCharacter {
		t.Errorf("Expected definition at character %d (tag name in decorator), got character %d",
			expectedCharacter, location.Range.Start.Character)
	}

	t.Logf("✅ Definition correctly resolved to: %s at Line:%d Character:%d",
		location.URI, location.Range.Start.Line, location.Range.Start.Character)
}

func TestDefinition_TreeSitterClassesQueryWorks(t *testing.T) {
	// Test that tree-sitter classes query can actually find definitions in TypeScript content
	// This ensures the query manager setup is correct for go-to-definition

	// Real TypeScript content with @customElement decorator
	tsContent := `import { LitElement, html, css } from 'lit';
import { customElement, property } from 'lit/decorators.js';

/**
 * A test element for verification
 */
@customElement('test-element')
export class TestElement extends LitElement {
  render() {
    return html` + "`<div>test</div>`" + `;
  }
}`

	// Create query manager with LSP queries
	qm, err := queries.NewQueryManager(queries.LSPQueries())
	if err != nil {
		t.Fatalf("Failed to create query manager: %v", err)
	}
	defer qm.Close()

	// Test FindTagNameDefinitionInSource
	tagRange, err := queries.FindTagNameDefinitionInSource([]byte(tsContent), "test-element", qm)
	if err != nil {
		t.Fatalf("FindTagNameDefinitionInSource failed: %v", err)
	}

	if tagRange == nil {
		t.Error("FindTagNameDefinitionInSource should find the @customElement decorator")
		t.Error("This indicates the classes query is not working properly")
	} else {
		// Should find the 'test-element' string in the decorator
		if tagRange.Start.Line != 6 { // @customElement is on line 6 (0-indexed)
			t.Errorf("Expected tag name at line 6, got line %d", tagRange.Start.Line)
		}
		t.Logf("✅ Tag name definition found at Line:%d Character:%d",
			tagRange.Start.Line, tagRange.Start.Character)
	}

	// Test FindClassDeclarationInSource
	classRange, err := queries.FindClassDeclarationInSource([]byte(tsContent), "TestElement", qm)
	if err != nil {
		t.Fatalf("FindClassDeclarationInSource failed: %v", err)
	}

	if classRange == nil {
		t.Error("FindClassDeclarationInSource should find the class declaration")
		t.Error("This indicates the classes query is not working properly")
	} else {
		// Should find the class name
		if classRange.Start.Line != 7 { // class declaration is on line 7 (0-indexed)
			t.Errorf("Expected class declaration at line 7, got line %d", classRange.Start.Line)
		}
		t.Logf("✅ Class declaration found at Line:%d Character:%d",
			classRange.Start.Line, classRange.Start.Character)
	}
}
