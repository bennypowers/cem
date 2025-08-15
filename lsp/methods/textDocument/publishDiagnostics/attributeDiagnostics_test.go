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

	"bennypowers.dev/cem/lsp/methods/textDocument/publishDiagnostics"
	"bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type mockAttributeDiagnosticsContext struct {
	content    string
	attributes map[string]map[string]*M.Attribute
}

func (m *mockAttributeDiagnosticsContext) Document(uri string) types.Document {
	return &mockAttributeDocument{content: m.content}
}

func (m *mockAttributeDiagnosticsContext) Slots(tagName string) ([]M.Slot, bool) {
	return nil, false
}

func (m *mockAttributeDiagnosticsContext) Attributes(tagName string) (map[string]*M.Attribute, bool) {
	attrs, exists := m.attributes[tagName]
	return attrs, exists
}

func (m *mockAttributeDiagnosticsContext) AllTagNames() []string {
	tags := make([]string, 0, len(m.attributes))
	for tag := range m.attributes {
		tags = append(tags, tag)
	}
	return tags
}

func (m *mockAttributeDiagnosticsContext) ElementDefinition(tagName string) (types.ElementDefinition, bool) {
	return nil, false
}

func (m *mockAttributeDiagnosticsContext) ElementSource(tagName string) (string, bool) {
	return "", false
}

type mockAttributeDocument struct {
	content string
}

func (m *mockAttributeDocument) Content() string {
	return m.content
}

func (m *mockAttributeDocument) Version() int32 {
	return 0
}

func (m *mockAttributeDocument) URI() string {
	return "test://test.html"
}

func (m *mockAttributeDocument) FindElementAtPosition(position protocol.Position, dm any) *types.CustomElementMatch {
	return nil
}

func (m *mockAttributeDocument) FindAttributeAtPosition(position protocol.Position, dm any) (*types.AttributeMatch, string) {
	return nil, ""
}

func (m *mockAttributeDocument) FindCustomElements(dm any) ([]types.CustomElementMatch, error) {
	return nil, nil
}

func (m *mockAttributeDocument) AnalyzeCompletionContextTS(position protocol.Position, dm any) *types.CompletionAnalysis {
	return nil
}

func (m *mockAttributeDocument) GetTemplateContext(position protocol.Position) string {
	return ""
}

func (m *mockAttributeDocument) GetScriptTags() []types.ScriptTag {
	return nil
}

func (m *mockAttributeDocument) FindModuleScript() (protocol.Position, bool) {
	return protocol.Position{}, false
}

func TestAttributeDiagnostics_GlobalAttributes(t *testing.T) {
	ctx := &mockAttributeDiagnosticsContext{
		content:    `<div class="test" id="main" data-value="42">Hello</div>`,
		attributes: map[string]map[string]*M.Attribute{},
	}

	diagnostics := publishDiagnostics.AnalyzeAttributeDiagnosticsForTest(ctx, ctx.Document("test.html"))

	// Should have no diagnostics for global attributes
	if len(diagnostics) != 0 {
		t.Errorf("Expected 0 diagnostics for global attributes, got %d", len(diagnostics))
	}
}

func TestAttributeDiagnostics_CustomElementValidAttribute(t *testing.T) {
	ctx := &mockAttributeDiagnosticsContext{
		content: `<my-element size="large" color="red">Content</my-element>`,
		attributes: map[string]map[string]*M.Attribute{
			"my-element": {
				"size":  &M.Attribute{FullyQualified: M.FullyQualified{Name: "size"}},
				"color": &M.Attribute{FullyQualified: M.FullyQualified{Name: "color"}},
			},
		},
	}

	diagnostics := publishDiagnostics.AnalyzeAttributeDiagnosticsForTest(ctx, ctx.Document("test.html"))

	// Should have no diagnostics for valid custom element attributes
	if len(diagnostics) != 0 {
		t.Errorf("Expected 0 diagnostics for valid custom element attributes, got %d", len(diagnostics))
	}
}

func TestAttributeDiagnostics_CustomElementInvalidAttribute(t *testing.T) {
	ctx := &mockAttributeDiagnosticsContext{
		content: `<my-element siz="large" colour="red">Content</my-element>`,
		attributes: map[string]map[string]*M.Attribute{
			"my-element": {
				"size":  &M.Attribute{FullyQualified: M.FullyQualified{Name: "size"}},
				"color": &M.Attribute{FullyQualified: M.FullyQualified{Name: "color"}},
			},
		},
	}

	diagnostics := publishDiagnostics.AnalyzeAttributeDiagnosticsForTest(ctx, ctx.Document("test.html"))

	// Should have 2 diagnostics for invalid attributes
	if len(diagnostics) != 2 {
		t.Errorf("Expected 2 diagnostics for invalid attributes, got %d", len(diagnostics))
		return
	}

	// Check first diagnostic (siz -> size)
	if diagnostics[0].Message != "Unknown attribute 'siz'. Did you mean 'size'?" {
		t.Errorf("Unexpected diagnostic message: %s", diagnostics[0].Message)
	}

	// Check second diagnostic (colour -> color)
	if diagnostics[1].Message != "Unknown attribute 'colour'. Did you mean 'color'?" {
		t.Errorf("Unexpected diagnostic message: %s", diagnostics[1].Message)
	}
}

func TestAttributeDiagnostics_ScriptTypeModule(t *testing.T) {
	ctx := &mockAttributeDiagnosticsContext{
		content:    `<script type="module">import './my-element.js';</script>`,
		attributes: map[string]map[string]*M.Attribute{},
	}

	diagnostics := publishDiagnostics.AnalyzeAttributeDiagnosticsForTest(ctx, ctx.Document("test.html"))

	// Should have 0 diagnostics - script[type] is outside CEM scope
	if len(diagnostics) != 0 {
		t.Errorf("Expected 0 diagnostics for script[type] (outside CEM scope), got %d", len(diagnostics))
		for _, diag := range diagnostics {
			t.Errorf("Unexpected diagnostic: %s", diag.Message)
		}
	}
}

func TestAttributeDiagnostics_GlobalAttributeTypoOnCustomElement(t *testing.T) {
	ctx := &mockAttributeDiagnosticsContext{
		content: `<my-custom-element clas="test" titl="tooltip">Hello</my-custom-element>`,
		attributes: map[string]map[string]*M.Attribute{
			"my-custom-element": {}, // Custom element with no defined attributes
		},
	}

	diagnostics := publishDiagnostics.AnalyzeAttributeDiagnosticsForTest(ctx, ctx.Document("test.html"))

	// Should have 2 diagnostics for global attribute typos on custom elements
	if len(diagnostics) != 2 {
		t.Errorf("Expected 2 diagnostics for global attribute typos on custom element, got %d", len(diagnostics))
		for i, diag := range diagnostics {
			t.Errorf("Diagnostic %d: %s", i, diag.Message)
		}
		return
	}

	// Check diagnostic messages
	found := make(map[string]bool)
	for _, diag := range diagnostics {
		found[diag.Message] = true
	}

	expected := []string{
		"Unknown attribute 'clas'. Did you mean 'class'?",
		"Unknown attribute 'titl'. Did you mean 'title'?",
	}

	for _, exp := range expected {
		if !found[exp] {
			t.Errorf("Expected diagnostic message not found: %s", exp)
		}
	}
}

func TestAttributeDiagnostics_StandardElementIgnored(t *testing.T) {
	ctx := &mockAttributeDiagnosticsContext{
		content:    `<div clas="test" titl="tooltip">Hello</div>`,
		attributes: map[string]map[string]*M.Attribute{},
	}

	diagnostics := publishDiagnostics.AnalyzeAttributeDiagnosticsForTest(ctx, ctx.Document("test.html"))

	// Should have 0 diagnostics - standard HTML elements are outside CEM scope (except slot)
	if len(diagnostics) != 0 {
		t.Errorf("Expected 0 diagnostics for standard HTML elements (outside CEM scope), got %d", len(diagnostics))
		for _, diag := range diagnostics {
			t.Errorf("Unexpected diagnostic: %s", diag.Message)
		}
	}
}
