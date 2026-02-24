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
	"encoding/json"
	"strings"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/methods/textDocument/definition"
	"bennypowers.dev/cem/lsp/testhelpers"
	M "bennypowers.dev/cem/manifest"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestDefinition(t *testing.T) {
	// Load test manifest with source information
	fs := testutil.NewFixtureFS(t, "legacy/slot-completions-test", "/test")
	manifestBytes, err := fs.ReadFile("/test/manifest.json")
	if err != nil {
		t.Fatalf("Failed to read test manifest: %v", err)
	}

	var pkg M.Package
	err = json.Unmarshal(manifestBytes, &pkg)
	if err != nil {
		t.Fatalf("Failed to parse manifest: %v", err)
	}

	// Create a mock document manager
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	// Test context that provides the registry data
	ctx := testhelpers.NewMockServerContext()
	ctx.SetDocumentManager(dm)
	ctx.AddManifest(&pkg)

	tests := []struct {
		name           string
		html           string
		position       protocol.Position
		expectLocation bool
		expectedPath   string
		description    string
	}{
		{
			name:           "Definition for custom element tag name",
			html:           `<card-element></card-element>`,
			position:       protocol.Position{Line: 0, Character: 5}, // Inside "card-element"
			expectLocation: true,
			expectedPath:   "card-element.js", // Should prefer .ts if available
			description:    "Should provide definition location for custom element",
		},
		{
			name:           "No definition for non-custom element",
			html:           `<div></div>`,
			position:       protocol.Position{Line: 0, Character: 2}, // Inside "div"
			expectLocation: false,
			description:    "Should not provide definition for standard HTML elements",
		},
		{
			name:           "No definition outside element",
			html:           `<card-element></card-element>`,
			position:       protocol.Position{Line: 0, Character: 20}, // After closing tag
			expectLocation: false,
			description:    "Should not provide definition when cursor is not on an element",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock document for this test case
			uri := "test://test.html"
			doc := dm.OpenDocument(uri, tt.html, 1)
			if doc == nil {
				t.Fatal("Failed to open test document")
			}

			// Also add document to the mock context for Document() lookup
			ctx.AddDocument(uri, doc)

			// Create definition params
			params := &protocol.DefinitionParams{
				TextDocumentPositionParams: protocol.TextDocumentPositionParams{
					TextDocument: protocol.TextDocumentIdentifier{URI: uri},
					Position:     tt.position,
				},
			}

			// Call the definition function
			result, err := definition.Definition(ctx, nil, params)
			if err != nil {
				t.Fatalf("Definition failed: %v", err)
			}

			if tt.expectLocation {
				if result == nil {
					t.Error("Expected definition location, but got nil")
					return
				}

				location, ok := result.(protocol.Location)
				if !ok {
					t.Errorf("Expected protocol.Location, got %T", result)
					return
				}

				// Check that URI contains the expected path
				if tt.expectedPath != "" && !strings.Contains(location.URI, tt.expectedPath) {
					t.Errorf("Expected URI to contain '%s', got '%s'", tt.expectedPath, location.URI)
				}

				t.Logf("Definition URI: %s", location.URI)
			} else {
				if result != nil {
					t.Errorf("Expected no definition location, but got: %v", result)
				}
			}

		})
	}
}

func TestDefinition_AttributeOnHTMLElement(t *testing.T) {
	// Regression test: go-to-definition on an attribute in HTML should work.
	// Previously, FindElementAtPosition used the tag name range, so clicking
	// on an attribute returned nil (position outside the tag name bytes).

	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	ctx := testhelpers.NewMockServerContext()
	ctx.SetDocumentManager(dm)

	// Register element definition
	elementDef := &testhelpers.MockElementDefinition{
		ModulePathStr:  "components/button-element.js",
		PackageNameStr: "demo-project",
	}
	ctx.ElementDefsMap["button-element"] = elementDef
	ctx.SetWorkspaceRoot("testdata/integration/definition-test-fixtures")

	htmlContent := `<button-element variant="primary"></button-element>`

	uri := "file:///test-attr.html"
	doc := dm.OpenDocument(uri, htmlContent, 1)
	ctx.AddDocument(uri, doc)

	tests := []struct {
		name     string
		position protocol.Position
		wantNil  bool
		desc     string
	}{
		{
			name:     "cursor on tag name",
			position: protocol.Position{Line: 0, Character: 5},
			wantNil:  false,
			desc:     "definition on tag name should work",
		},
		{
			name:     "cursor on attribute name",
			position: protocol.Position{Line: 0, Character: 18},
			wantNil:  false,
			desc:     "definition on attribute name should work",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := &protocol.DefinitionParams{
				TextDocumentPositionParams: protocol.TextDocumentPositionParams{
					TextDocument: protocol.TextDocumentIdentifier{URI: uri},
					Position:     tt.position,
				},
			}

			result, err := definition.Definition(ctx, nil, params)
			if err != nil {
				t.Fatalf("Definition failed: %v", err)
			}

			if tt.wantNil {
				if result != nil {
					t.Errorf("expected nil, got %v", result)
				}
			} else {
				if result == nil {
					t.Fatalf("expected definition location, got nil: %s", tt.desc)
				}
				location, ok := result.(protocol.Location)
				if !ok {
					t.Fatalf("expected protocol.Location, got %T", result)
				}
				t.Logf("definition URI: %s at %d:%d", location.URI, location.Range.Start.Line, location.Range.Start.Character)
			}
		})
	}
}

// TestResolveSourcePath tests the path resolution logic specifically
func TestResolveSourcePath(t *testing.T) {
	tests := []struct {
		name          string
		modulePath    string
		workspaceRoot string
		expected      string
	}{
		{
			name:          "Relative path with ./ prefix",
			modulePath:    "./rh-card/rh-card.js",
			workspaceRoot: "/workspace/project",
			expected:      "file:///workspace/project/rh-card/rh-card.js",
		},
		{
			name:          "Relative path without ./ prefix",
			modulePath:    "components/my-button.js",
			workspaceRoot: "/workspace/project",
			expected:      "file:///workspace/project/components/my-button.js",
		},
		{
			name:          "Absolute path",
			modulePath:    "/absolute/path/to/element.js",
			workspaceRoot: "/workspace/project",
			expected:      "file:///absolute/path/to/element.js",
		},
		{
			name:          "Already file:// URI",
			modulePath:    "file:///already/a/uri.js",
			workspaceRoot: "/workspace/project",
			expected:      "file:///already/a/uri.js",
		},
		{
			name:          "Empty module path",
			modulePath:    "",
			workspaceRoot: "/workspace/project",
			expected:      "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test-specific implementation
			testDef := &testElementDefinition{
				modulePath: tt.modulePath,
			}

			result := definition.ResolveSourcePathForTesting(testDef, tt.workspaceRoot)

			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

// testElementDefinition is a test-only implementation of types.ElementDefinition
type testElementDefinition struct {
	modulePath string
	sourceHref string
}

func (t *testElementDefinition) ModulePath() string {
	return t.modulePath
}

func (t *testElementDefinition) SourceHref() string {
	return t.sourceHref
}

func (t *testElementDefinition) PackageName() string {
	return ""
}

func (t *testElementDefinition) Element() *M.CustomElement {
	return nil // For testing path resolution, we don't need the actual element
}
