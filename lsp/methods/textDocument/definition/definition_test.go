package definition_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/methods/textDocument"
	"bennypowers.dev/cem/lsp/methods/textDocument/definition"
	"bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestDefinition(t *testing.T) {
	// Load test manifest with source information
	fixtureDir := filepath.Join("slot-completions-test")
	manifestPath := filepath.Join(fixtureDir, "manifest.json")

	manifestBytes, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatalf("Failed to read test manifest: %v", err)
	}

	var pkg M.Package
	err = json.Unmarshal(manifestBytes, &pkg)
	if err != nil {
		t.Fatalf("Failed to parse manifest: %v", err)
	}

	// Create registry and add the test manifest
	registry := lsp.NewRegistry()
	registry.AddManifest(&pkg)

	// Create a mock document manager
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	// Test context that provides the registry data
	ctx := &testDefinitionContext{
		registry: registry,
		dm:       dm,
	}

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
				if tt.expectedPath != "" && !contains(location.URI, tt.expectedPath) {
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

// contains is a helper function to check if a string contains a substring
func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// testDefinitionContext implements DefinitionContext for testing
type testDefinitionContext struct {
	registry *lsp.Registry
	dm       *lsp.DocumentManager
}

func (ctx *testDefinitionContext) GetDocument(uri string) types.Document {
	return ctx.dm.GetDocument(uri)
}

func (ctx *testDefinitionContext) GetElementDefinition(tagName string) (types.ElementDefinition, bool) {
	return ctx.registry.GetElementDefinition(tagName)
}

func (ctx *testDefinitionContext) GetWorkspaceRoot() string {
	// Return current directory for tests
	if wd, err := os.Getwd(); err == nil {
		return wd
	}
	return ""
}

func (ctx *testDefinitionContext) GetRawDocumentManager() interface{} {
	return ctx.dm
}

// MockElementDefinition implements textDocument.ElementDefinition
type MockElementDefinition struct {
	def *lsp.ElementDefinition
}

func (m *MockElementDefinition) GetModulePath() string {
	return m.def.ModulePath
}

func (m *MockElementDefinition) GetSourceHref() string {
	if m.def.Source != nil {
		return m.def.Source.Href
	}
	return ""
}

// MockDocumentAdapter adapts lsp.Document to textDocument.Document
type MockDocumentAdapter struct {
	doc *lsp.Document
	dm  *lsp.DocumentManager
}

func (m *MockDocumentAdapter) FindElementAtPosition(position protocol.Position, dm interface{}) *textDocument.CustomElementMatch {
	element := m.doc.FindElementAtPosition(position, m.dm)
	if element == nil {
		return nil
	}

	// Convert attributes
	attrs := make(map[string]textDocument.AttributeMatch)
	for k, v := range element.Attributes {
		attrs[k] = textDocument.AttributeMatch{
			Name:  v.Name,
			Value: v.Value,
			Range: v.Range,
		}
	}

	return &textDocument.CustomElementMatch{
		TagName:    element.TagName,
		Range:      element.Range,
		Attributes: attrs,
	}
}

func (m *MockDocumentAdapter) FindAttributeAtPosition(position protocol.Position, dm interface{}) (*types.AttributeMatch, string) {
	attr, tagName := m.doc.FindAttributeAtPosition(position, m.dm)
	if attr == nil {
		return nil, ""
	}

	return &types.AttributeMatch{
		Name:  attr.Name,
		Value: attr.Value,
		Range: attr.Range,
	}, tagName
}

func (m *MockDocumentAdapter) GetContent() string {
	return m.doc.Content
}

func (m *MockDocumentAdapter) AnalyzeCompletionContextTS(position protocol.Position, dm any) *types.CompletionAnalysis {
	return m.doc.AnalyzeCompletionContextTS(position, m.dm)
}

func (m *MockDocumentAdapter) GetTemplateContext(position protocol.Position) string {
	return ""
}

// GetVersion returns the document version
func (m *MockDocumentAdapter) GetVersion() int32 {
	return m.doc.GetVersion()
}

// GetURI returns the document URI
func (m *MockDocumentAdapter) GetURI() string {
	return m.doc.GetURI()
}

// FindCustomElements returns custom elements
func (m *MockDocumentAdapter) FindCustomElements(dm any) ([]types.CustomElementMatch, error) {
	return m.doc.FindCustomElements(dm)
}
