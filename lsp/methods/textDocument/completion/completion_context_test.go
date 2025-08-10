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
package completion_test

import (
	"testing"

	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/methods/textDocument"
	"bennypowers.dev/cem/lsp/methods/textDocument/completion"
	"bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type MockCompletionContext struct {
	registry  *MockRegistry
	documents map[string]types.Document
}

func (m *MockCompletionContext) GetDocument(uri string) types.Document {
	return m.documents[uri]
}

func (m *MockCompletionContext) GetAllTagNames() []string {
	return m.registry.GetAllTagNames()
}

func (m *MockCompletionContext) GetElement(tagName string) (*M.CustomElement, bool) {
	return m.registry.GetElement(tagName)
}

func (m *MockCompletionContext) GetAttributes(tagName string) (map[string]*M.Attribute, bool) {
	return m.registry.GetAttributes(tagName)
}

func (m *MockCompletionContext) GetSlots(tagName string) ([]M.Slot, bool) {
	return m.registry.GetSlots(tagName)
}

// TestCompletionContextAnalysis tests cursor position analysis for different completion scenarios
func TestCompletionContextAnalysis(t *testing.T) {
	tests := []struct {
		name         string
		content      string
		line         uint32
		character    uint32
		triggerChar  string
		expectedType types.CompletionContextType
		expectedTag  string
		expectedAttr string
	}{
		{
			name:         "Tag name completion after <",
			content:      "<my-",
			line:         0,
			character:    4,
			triggerChar:  "",
			expectedType: types.CompletionTagName,
		},
		{
			name:         "Tag name completion with trigger char",
			content:      "<",
			line:         0,
			character:    1,
			triggerChar:  "<",
			expectedType: types.CompletionTagName,
		},
		{
			name:         "Attribute name completion",
			content:      "<my-element ",
			line:         0,
			character:    12,
			triggerChar:  " ",
			expectedType: types.CompletionAttributeName,
			expectedTag:  "my-element",
		},
		{
			name:         "Attribute value completion",
			content:      "<my-element disabled=\"",
			line:         0,
			character:    21,
			triggerChar:  "\"",
			expectedType: types.CompletionAttributeValue,
			expectedTag:  "my-element",
			expectedAttr: "disabled",
		},
		{
			name:         "Template literal tag completion",
			content:      "const html = html`<my-",
			line:         0,
			character:    22,
			triggerChar:  "",
			expectedType: types.CompletionTagName,
		},
		{
			name:         "Template literal attribute completion",
			content:      "element.innerHTML = `<my-element ",
			line:         0,
			character:    33,
			triggerChar:  " ",
			expectedType: types.CompletionAttributeName,
			expectedTag:  "my-element",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock document
			doc := &MockDocument{content: tt.content}

			position := protocol.Position{
				Line:      tt.line,
				Character: tt.character,
			}

			analysis := textDocument.AnalyzeCompletionContext(doc, position, tt.triggerChar)

			if analysis.Type != tt.expectedType {
				t.Errorf("Expected completion type %d, got %d", tt.expectedType, analysis.Type)
			}

			if tt.expectedTag != "" && analysis.TagName != tt.expectedTag {
				t.Errorf("Expected tag name %q, got %q", tt.expectedTag, analysis.TagName)
			}

			if tt.expectedAttr != "" && analysis.AttributeName != tt.expectedAttr {
				t.Errorf("Expected attribute name %q, got %q", tt.expectedAttr, analysis.AttributeName)
			}
		})
	}
}

// TestCompletionPrefixExtraction tests prefix extraction for filtering completions
func TestCompletionPrefixExtraction(t *testing.T) {
	tests := []struct {
		name           string
		beforeCursor   string
		completionType types.CompletionContextType
		tagName        string
		expectedPrefix string
	}{
		{
			name:           "Tag name prefix",
			beforeCursor:   "<my-",
			completionType: types.CompletionTagName,
			expectedPrefix: "my-",
		},
		{
			name:           "Attribute name prefix",
			beforeCursor:   "<my-element dis",
			completionType: types.CompletionAttributeName,
			expectedPrefix: "dis",
		},
		{
			name:           "No prefix",
			beforeCursor:   "<",
			completionType: types.CompletionTagName,
			expectedPrefix: "",
		},
		{
			name:           "Attribute value prefix",
			beforeCursor:   "<my-element disabled=\"tr",
			completionType: types.CompletionAttributeValue,
			expectedPrefix: "tr",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analysis := &types.CompletionAnalysis{
				Type:    tt.completionType,
				TagName: tt.tagName,
			}

			prefix := textDocument.GetCompletionPrefix(tt.beforeCursor, analysis)

			if prefix != tt.expectedPrefix {
				t.Errorf("Expected prefix %q, got %q", tt.expectedPrefix, prefix)
			}
		})
	}
}

// TestCustomElementTagValidation tests custom element tag name validation
func TestCustomElementTagValidation(t *testing.T) {
	tests := []struct {
		name     string
		tagName  string
		expected bool
	}{
		{"Valid custom element", "my-element", true},
		{"Valid with numbers", "my-element2", true},
		{"Valid with multiple hyphens", "my-custom-element", true},
		{"Invalid no hyphen", "div", false},
		{"Invalid starts with uppercase", "My-element", false},
		{"Invalid empty", "", false},
		{"Invalid only hyphen", "-", false},
		{"Invalid starts with hyphen", "-element", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := textDocument.IsCustomElementTag(tt.tagName)
			if result != tt.expected {
				t.Errorf("Expected %v for tag %q, got %v", tt.expected, tt.tagName, result)
			}
		})
	}
}

// TestCompletionIntegration tests the full completion flow
func TestCompletionIntegration(t *testing.T) {
	// Create test registry with custom elements
	registry := &MockRegistry{
		elements: map[string]*M.CustomElement{
			"my-button": {
				TagName: "my-button",
				Attributes: []M.Attribute{
					{FullyQualified: M.FullyQualified{Name: "disabled"}, Type: &M.Type{Text: "boolean"}},
					{FullyQualified: M.FullyQualified{Name: "variant"}, Type: &M.Type{Text: "string"}},
				},
			},
			"my-input": {
				TagName: "my-input",
				Attributes: []M.Attribute{
					{FullyQualified: M.FullyQualified{Name: "value"}, Type: &M.Type{Text: "string"}},
					{FullyQualified: M.FullyQualified{Name: "required"}, Type: &M.Type{Text: "boolean"}},
				},
			},
		},
	}

	// Create completion context
	ctx := &MockCompletionContext{
		registry: registry,
		documents: map[string]types.Document{
			"test://test.html": &MockDocument{content: "<my-"},
		},
	}

	// Test tag name completion
	t.Run("Tag name completion", func(t *testing.T) {
		params := &protocol.CompletionParams{
			TextDocumentPositionParams: protocol.TextDocumentPositionParams{
				TextDocument: protocol.TextDocumentIdentifier{URI: "test://test.html"},
				Position:     protocol.Position{Line: 0, Character: 4},
			},
		}

		result, err := completion.Completion(ctx, nil, params)
		if err != nil {
			t.Fatalf("Completion failed: %v", err)
		}

		items, ok := result.([]protocol.CompletionItem)
		if !ok {
			t.Fatalf("Expected []CompletionItem, got %T", result)
		}

		// Should contain both custom elements
		if len(items) != 2 {
			t.Errorf("Expected 2 completion items, got %d", len(items))
		}

		// Check that we have the expected elements
		labels := make([]string, len(items))
		for i, item := range items {
			labels[i] = item.Label
		}

		expectedLabels := []string{"my-button", "my-input"}
		for _, expected := range expectedLabels {
			found := false
			for _, label := range labels {
				if label == expected {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Expected to find %q in completion items %v", expected, labels)
			}
		}
	})

	// Test attribute completion
	t.Run("Attribute completion", func(t *testing.T) {
		ctx.documents["test://test.html"] = &MockDocument{content: "<my-button "}

		params := &protocol.CompletionParams{
			TextDocumentPositionParams: protocol.TextDocumentPositionParams{
				TextDocument: protocol.TextDocumentIdentifier{URI: "test://test.html"},
				Position:     protocol.Position{Line: 0, Character: 11},
			},
		}

		result, err := completion.Completion(ctx, nil, params)
		if err != nil {
			t.Fatalf("Completion failed: %v", err)
		}

		items, ok := result.([]protocol.CompletionItem)
		if !ok {
			t.Fatalf("Expected []CompletionItem, got %T", result)
		}

		// Should contain attributes for my-button
		if len(items) != 2 {
			t.Errorf("Expected 2 attribute completion items, got %d", len(items))
		}

		labels := make([]string, len(items))
		for i, item := range items {
			labels[i] = item.Label
		}

		expectedLabels := []string{"disabled", "variant"}
		for _, expected := range expectedLabels {
			found := false
			for _, label := range labels {
				if label == expected {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Expected to find %q in attribute completion items %v", expected, labels)
			}
		}
	})

	// Test boolean attribute value completion
	t.Run("Boolean attribute value completion", func(t *testing.T) {
		ctx.documents["test://test.html"] = &MockDocument{content: "<my-button disabled=\""}

		params := &protocol.CompletionParams{
			TextDocumentPositionParams: protocol.TextDocumentPositionParams{
				TextDocument: protocol.TextDocumentIdentifier{URI: "test://test.html"},
				Position:     protocol.Position{Line: 0, Character: 20},
			},
		}

		result, err := completion.Completion(ctx, nil, params)
		if err != nil {
			t.Fatalf("Completion failed: %v", err)
		}

		items, ok := result.([]protocol.CompletionItem)
		if !ok {
			t.Fatalf("Expected []CompletionItem, got %T", result)
		}

		// Boolean attributes should not provide value completions
		// (presence = true, absence = false)
		if len(items) != 0 {
			t.Errorf("Expected 0 boolean value completion items (presence=true), got %d", len(items))
		}
	})
}

// Mock implementations for testing

type MockDocument struct {
	content string
}

func (m *MockDocument) GetContent() string {
	return m.content
}

func (m *MockDocument) FindElementAtPosition(position protocol.Position, dm any) *types.CustomElementMatch {
	// Simple mock implementation - not used in these tests
	return nil
}

func (m *MockDocument) FindAttributeAtPosition(position protocol.Position, dm any) (*types.AttributeMatch, string) {
	// Simple mock implementation - not used in these tests
	return nil, ""
}

// GetTemplateContext returns empty string for HTML documents (not a template)
func (m *MockDocument) GetTemplateContext(position protocol.Position) string {
	return ""
}

// GetVersion returns the document version
func (m *MockDocument) GetVersion() int32 {
	return 1
}

// GetURI returns the document URI
func (m *MockDocument) GetURI() string {
	return "test://mock"
}

// FindCustomElements returns empty list for mock
func (m *MockDocument) FindCustomElements(dm any) ([]types.CustomElementMatch, error) {
	return nil, nil
}

// AnalyzeCompletionContextTS returns nil for mock
func (m *MockDocument) AnalyzeCompletionContextTS(position protocol.Position, dm any) *types.CompletionAnalysis {
	return nil
}

type MockRegistry struct {
	elements map[string]*M.CustomElement
}

func (m *MockRegistry) GetElement(tagName string) (*M.CustomElement, bool) {
	elem, exists := m.elements[tagName]
	return elem, exists
}

func (m *MockRegistry) GetAttributes(tagName string) (map[string]*M.Attribute, bool) {
	elem, exists := m.elements[tagName]
	if !exists {
		return nil, false
	}

	attrs := make(map[string]*M.Attribute)
	for i := range elem.Attributes {
		attr := &elem.Attributes[i]
		attrs[attr.Name] = attr
	}
	return attrs, true
}

func (m *MockRegistry) GetAllTagNames() []string {
	names := make([]string, 0, len(m.elements))
	for name := range m.elements {
		names = append(names, name)
	}
	return names
}

func (m *MockRegistry) GetSlots(tagName string) ([]M.Slot, bool) {
	elem, exists := m.elements[tagName]
	if !exists {
		return nil, false
	}
	return elem.Slots, true
}

// TestCompletionContextBehavior tests the true desired behavior for completion context analysis
// This represents what the system should do, regardless of implementation (regex vs tree-sitter)
func TestCompletionContextBehavior(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		position protocol.Position
		want     types.CompletionContextType
		wantTag  string
		wantAttr string
	}{
		// HTML Tag Name Completion
		{
			name:     "Tag name after opening bracket",
			content:  "<",
			position: protocol.Position{Line: 0, Character: 1},
			want:     types.CompletionTagName,
		},
		{
			name:     "Partial tag name",
			content:  "<my-el",
			position: protocol.Position{Line: 0, Character: 6},
			want:     types.CompletionTagName,
		},
		{
			name:     "Tag name with hyphen",
			content:  "<custom-element",
			position: protocol.Position{Line: 0, Character: 15},
			want:     types.CompletionTagName,
		},

		// HTML Attribute Name Completion
		{
			name:     "Attribute name after space",
			content:  "<my-element ",
			position: protocol.Position{Line: 0, Character: 12},
			want:     types.CompletionAttributeName,
			wantTag:  "my-element",
		},
		{
			name:     "Partial attribute name",
			content:  "<my-element dis",
			position: protocol.Position{Line: 0, Character: 15},
			want:     types.CompletionAttributeName,
			wantTag:  "my-element",
		},
		{
			name:     "Attribute name after other attributes",
			content:  "<my-element class=\"foo\" ",
			position: protocol.Position{Line: 0, Character: 24},
			want:     types.CompletionAttributeName,
			wantTag:  "my-element",
		},
		{
			name:     "Attribute name with hyphens",
			content:  "<my-element data-test",
			position: protocol.Position{Line: 0, Character: 21},
			want:     types.CompletionAttributeName,
			wantTag:  "my-element",
		},

		// HTML Attribute Value Completion
		{
			name:     "Attribute value in double quotes",
			content:  "<my-element disabled=\"",
			position: protocol.Position{Line: 0, Character: 21},
			want:     types.CompletionAttributeValue,
			wantTag:  "my-element",
			wantAttr: "disabled",
		},
		{
			name:     "Attribute value in single quotes",
			content:  "<my-element type='",
			position: protocol.Position{Line: 0, Character: 18},
			want:     types.CompletionAttributeValue,
			wantTag:  "my-element",
			wantAttr: "type",
		},
		{
			name:     "Partial attribute value",
			content:  "<my-element variant=\"prim",
			position: protocol.Position{Line: 0, Character: 25},
			want:     types.CompletionAttributeValue,
			wantTag:  "my-element",
			wantAttr: "variant",
		},
		{
			name:     "Attribute value with hyphens",
			content:  "<my-element data-test=\"some-val",
			position: protocol.Position{Line: 0, Character: 31},
			want:     types.CompletionAttributeValue,
			wantTag:  "my-element",
			wantAttr: "data-test",
		},

		// Invalid HTML attributes (Lit syntax is only valid in template literals)
		{
			name:     "Invalid attribute with @ (should be treated as attribute name)",
			content:  "<my-element @",
			position: protocol.Position{Line: 0, Character: 13},
			want:     types.CompletionAttributeName,
			wantTag:  "my-element",
		},
		{
			name:     "Invalid attribute with @ partial (should be treated as attribute name)",
			content:  "<my-element @cli",
			position: protocol.Position{Line: 0, Character: 16},
			want:     types.CompletionAttributeName,
			wantTag:  "my-element",
		},
		{
			name:     "Invalid attribute with dot (should be treated as attribute name)",
			content:  "<my-element .",
			position: protocol.Position{Line: 0, Character: 13},
			want:     types.CompletionAttributeName,
			wantTag:  "my-element",
		},
		{
			name:     "Invalid attribute with dot partial (should be treated as attribute name)",
			content:  "<my-element .val",
			position: protocol.Position{Line: 0, Character: 16},
			want:     types.CompletionAttributeName,
			wantTag:  "my-element",
		},
		{
			name:     "Invalid attribute with question mark (should be treated as attribute name)",
			content:  "<my-element ?",
			position: protocol.Position{Line: 0, Character: 13},
			want:     types.CompletionAttributeName,
			wantTag:  "my-element",
		},
		{
			name:     "Invalid attribute with question mark partial (should be treated as attribute name)",
			content:  "<my-element ?disab",
			position: protocol.Position{Line: 0, Character: 18},
			want:     types.CompletionAttributeName,
			wantTag:  "my-element",
		},

		// Multi-line scenarios
		{
			name: "Tag name on new line",
			content: `<div>
  <my-el`,
			position: protocol.Position{Line: 1, Character: 8},
			want:     types.CompletionTagName,
		},
		{
			name: "Attribute on new line",
			content: `<my-element
  dis`,
			position: protocol.Position{Line: 1, Character: 5},
			want:     types.CompletionAttributeName,
			wantTag:  "my-element",
		},
		{
			name: "Attribute value on new line",
			content: `<my-element
  variant="`,
			position: protocol.Position{Line: 1, Character: 11},
			want:     types.CompletionAttributeValue,
			wantTag:  "my-element",
			wantAttr: "variant",
		},

		// Edge cases and negative scenarios
		{
			name:     "Inside tag content (no completion)",
			content:  "<my-element>content",
			position: protocol.Position{Line: 0, Character: 15},
			want:     types.CompletionUnknown,
		},
		{
			name:     "After closing tag",
			content:  "<my-element></my-element>",
			position: protocol.Position{Line: 0, Character: 25},
			want:     types.CompletionUnknown,
		},
		{
			name:     "Not in a tag context",
			content:  "Some text content",
			position: protocol.Position{Line: 0, Character: 10},
			want:     types.CompletionUnknown,
		},
		{
			name:     "Standard HTML element (no completion)",
			content:  "<div ",
			position: protocol.Position{Line: 0, Character: 5},
			want:     types.CompletionUnknown, // Only custom elements
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create document manager for tree-sitter analysis
			dm, err := lsp.NewDocumentManager()
			if err != nil {
				t.Fatalf("Failed to create document manager: %v", err)
			}
			defer dm.Close()

			// Create real document with tree-sitter parsing
			doc := dm.OpenDocument("test://test.html", tt.content, 1)

			// Analyze context using tree-sitter implementation
			analysis := doc.AnalyzeCompletionContextTS(tt.position, dm)
			if analysis == nil {
				analysis = &types.CompletionAnalysis{Type: types.CompletionUnknown}
			}

			// Verify completion type
			if analysis.Type != tt.want {
				t.Errorf("Expected completion type %d (%s), got %d",
					tt.want, completionTypeString(tt.want), analysis.Type)
			}

			// Verify tag name if expected
			if tt.wantTag != "" && analysis.TagName != tt.wantTag {
				t.Errorf("Expected tag name %q, got %q", tt.wantTag, analysis.TagName)
			}

			// Verify attribute name if expected
			if tt.wantAttr != "" && analysis.AttributeName != tt.wantAttr {
				t.Errorf("Expected attribute name %q, got %q", tt.wantAttr, analysis.AttributeName)
			}
		})
	}
}

// TestTemplateContextBehavior tests completion context in TypeScript template literals
func TestTemplateContextBehavior(t *testing.T) {
	tests := []struct {
		name            string
		content         string
		position        protocol.Position
		want            types.CompletionContextType
		wantTag         string
		wantAttr        string
		wantLitTemplate bool
		wantLitSyntax   string
	}{
		// Template literal contexts
		{
			name: "Tag completion in html`` template",
			content: `const template = html` + "`" + `
  <my-el
` + "`" + `;`,
			position:        protocol.Position{Line: 1, Character: 8},
			want:            types.CompletionTagName,
			wantLitTemplate: true,
		},
		{
			name: "Attribute completion in html`` template",
			content: `const template = html` + "`" + `
  <my-element dis
` + "`" + `;`,
			position:        protocol.Position{Line: 1, Character: 15},
			want:            types.CompletionAttributeName,
			wantTag:         "my-element",
			wantLitTemplate: true,
		},
		{
			name: "Lit event binding in template",
			content: `const template = html` + "`" + `
  <my-element @cli
` + "`" + `;`,
			position:        protocol.Position{Line: 1, Character: 16},
			want:            types.CompletionLitEventBinding,
			wantTag:         "my-element",
			wantLitTemplate: true,
			wantLitSyntax:   "@",
		},
		{
			name: "Lit property binding in template",
			content: `const template = html` + "`" + `
  <my-element .prop
` + "`" + `;`,
			position:        protocol.Position{Line: 1, Character: 17},
			want:            types.CompletionLitPropertyBinding,
			wantTag:         "my-element",
			wantLitTemplate: true,
			wantLitSyntax:   ".",
		},
		{
			name: "Lit boolean attribute in template",
			content: `const template = html` + "`" + `
  <my-element ?disab
` + "`" + `;`,
			position:        protocol.Position{Line: 1, Character: 18},
			want:            types.CompletionLitBooleanAttribute,
			wantTag:         "my-element",
			wantLitTemplate: true,
			wantLitSyntax:   "?",
		},

		// innerHTML context (not Lit template, no special syntax)
		{
			name: "Tag completion in innerHTML",
			content: `element.innerHTML = ` + "`" + `
  <my-el
` + "`" + `;`,
			position:        protocol.Position{Line: 1, Character: 8},
			want:            types.CompletionTagName,
			wantLitTemplate: false,
		},
		{
			name: "Regular attribute in innerHTML (no Lit syntax)",
			content: `element.innerHTML = ` + "`" + `
  <my-element @cli
` + "`" + `;`,
			position:        protocol.Position{Line: 1, Character: 16},
			want:            types.CompletionAttributeName,
			wantTag:         "my-element",
			wantLitTemplate: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock document that reports template context
			doc := &MockTemplateDocument{
				content:       tt.content,
				isLitTemplate: tt.wantLitTemplate,
			}

			// Analyze context
			analysis := textDocument.AnalyzeCompletionContext(doc, tt.position, "")

			// Verify completion type
			if analysis.Type != tt.want {
				t.Errorf("Expected completion type %d (%s), got %d",
					tt.want, completionTypeString(tt.want), analysis.Type)
			}

			// Verify template context
			if analysis.IsLitTemplate != tt.wantLitTemplate {
				t.Errorf("Expected IsLitTemplate %t, got %t", tt.wantLitTemplate, analysis.IsLitTemplate)
			}

			// Verify Lit syntax if expected
			if tt.wantLitSyntax != "" && analysis.LitSyntax != tt.wantLitSyntax {
				t.Errorf("Expected LitSyntax %q, got %q", tt.wantLitSyntax, analysis.LitSyntax)
			}

			// Verify tag name if expected
			if tt.wantTag != "" && analysis.TagName != tt.wantTag {
				t.Errorf("Expected tag name %q, got %q", tt.wantTag, analysis.TagName)
			}

			// Verify attribute name if expected
			if tt.wantAttr != "" && analysis.AttributeName != tt.wantAttr {
				t.Errorf("Expected attribute name %q, got %q", tt.wantAttr, analysis.AttributeName)
			}
		})
	}
}

// MockTemplateDocument extends MockDocument with template context support
type MockTemplateDocument struct {
	content       string
	isLitTemplate bool
}

func (m *MockTemplateDocument) GetContent() string {
	return m.content
}

func (m *MockTemplateDocument) FindElementAtPosition(position protocol.Position, dm any) *types.CustomElementMatch {
	return nil
}

func (m *MockTemplateDocument) FindAttributeAtPosition(position protocol.Position, dm any) (*types.AttributeMatch, string) {
	return nil, ""
}

func (m *MockTemplateDocument) GetTemplateContext(position protocol.Position) string {
	if m.isLitTemplate {
		return "html"
	}
	return "innerHTML"
}

// GetVersion returns the document version
func (m *MockTemplateDocument) GetVersion() int32 {
	return 1
}

// GetURI returns the document URI
func (m *MockTemplateDocument) GetURI() string {
	return "test://mock-template"
}

// FindCustomElements returns empty list for mock
func (m *MockTemplateDocument) FindCustomElements(dm any) ([]types.CustomElementMatch, error) {
	return nil, nil
}

// AnalyzeCompletionContextTS returns nil for mock
func (m *MockTemplateDocument) AnalyzeCompletionContextTS(position protocol.Position, dm any) *types.CompletionAnalysis {
	return nil
}

// Helper function to convert completion type to string for better error messages
func completionTypeString(t types.CompletionContextType) string {
	switch t {
	case types.CompletionUnknown:
		return "Unknown"
	case types.CompletionTagName:
		return "TagName"
	case types.CompletionAttributeName:
		return "AttributeName"
	case types.CompletionAttributeValue:
		return "AttributeValue"
	case types.CompletionLitEventBinding:
		return "LitEventBinding"
	case types.CompletionLitPropertyBinding:
		return "LitPropertyBinding"
	case types.CompletionLitBooleanAttribute:
		return "LitBooleanAttribute"
	default:
		return "Invalid"
	}
}
