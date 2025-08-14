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
package publishDiagnostics

import (
	"testing"

	"bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// Mock implementations for tag diagnostics testing
type mockDocument struct {
	content string
}

func (m *mockDocument) Content() string {
	return m.content
}

func (m *mockDocument) Version() int32 {
	return 1
}

func (m *mockDocument) URI() string {
	return "file:///test.html"
}

func (m *mockDocument) FindElementAtPosition(position protocol.Position, dm any) *types.CustomElementMatch {
	return nil
}

func (m *mockDocument) FindAttributeAtPosition(position protocol.Position, dm any) (*types.AttributeMatch, string) {
	return nil, ""
}

func (m *mockDocument) FindCustomElements(dm any) ([]types.CustomElementMatch, error) {
	return nil, nil
}

func (m *mockDocument) AnalyzeCompletionContextTS(position protocol.Position, dm any) *types.CompletionAnalysis {
	return nil
}

type mockTagDiagnosticsContext struct {
	availableElements map[string]bool
	elementSources    map[string]string
}

func (m *mockTagDiagnosticsContext) Document(uri string) types.Document {
	return &mockDocument{content: "<my-button>Test</my-button><my-buttom>Test</my-buttom><missing-import-element>Test</missing-import-element><completely-different-element>Test</completely-different-element>"}
}

func (m *mockTagDiagnosticsContext) Slots(tagName string) ([]M.Slot, bool) {
	return nil, false
}

func (m *mockTagDiagnosticsContext) Attributes(tagName string) (map[string]*M.Attribute, bool) {
	return nil, false
}

func (m *mockTagDiagnosticsContext) AllTagNames() []string {
	var result []string
	for tag := range m.availableElements {
		if m.availableElements[tag] {
			result = append(result, tag)
		}
	}
	return result
}

func (m *mockTagDiagnosticsContext) ElementDefinition(tagName string) (types.ElementDefinition, bool) {
	// Elements that exist in registry but are not in the available list (not imported)
	if tagName == "missing-import-element" {
		// Only return definition if element is NOT in available elements (simulating not imported)
		if !m.availableElements[tagName] {
			return &mockElementDefinition{modulePath: "./missing-import-element.js"}, true
		}
	}
	return nil, false
}

func (m *mockTagDiagnosticsContext) ElementSource(tagName string) (string, bool) {
	// Only return source if element is NOT in available elements (simulating not imported)
	if !m.availableElements[tagName] {
		source, exists := m.elementSources[tagName]
		return source, exists
	}
	return "", false
}

type mockElementDefinition struct {
	modulePath string
}

func (m *mockElementDefinition) ModulePath() string {
	return m.modulePath
}

func (m *mockElementDefinition) SourceHref() string {
	return ""
}

func TestTagNameDiagnostics(t *testing.T) {
	tests := []struct {
		name                     string
		availableElements        map[string]bool
		elementSources           map[string]string
		expectedDiagnosticsCount int
		expectedMessages         []string
		expectedTypes            []string
	}{
		{
			name: "Unknown element with suggestion",
			availableElements: map[string]bool{
				"my-button": true,
			},
			elementSources: map[string]string{
				"missing-import-element": "./missing-import-element.js", // Element exists but not imported
			},
			expectedDiagnosticsCount: 3, // my-buttom (typo), missing-import-element, completely-different-element
			expectedMessages: []string{
				"Unknown custom element 'my-buttom'. Did you mean 'my-button'?",
				"Custom element 'missing-import-element' is not imported. Add import from './missing-import-element.js'",
				"Unknown custom element 'completely-different-element'. Available elements: 'my-button'",
			},
			expectedTypes: []string{
				string(types.DiagnosticTypeTagSuggestion),
				string(types.DiagnosticTypeMissingImport),
				"", // Large distance doesn't get suggestion
			},
		},
		{
			name: "Missing import element",
			availableElements: map[string]bool{
				"my-button": true,
			},
			elementSources: map[string]string{
				"missing-import-element": "./missing-import-element.js",
			},
			expectedDiagnosticsCount: 3,
			expectedMessages: []string{
				"Unknown custom element 'my-buttom'. Did you mean 'my-button'?",
				"Custom element 'missing-import-element' is not imported. Add import from './missing-import-element.js'",
				"Unknown custom element 'completely-different-element'. Available elements: 'my-button'",
			},
			expectedTypes: []string{
				string(types.DiagnosticTypeTagSuggestion),
				string(types.DiagnosticTypeMissingImport),
				"", // Large distance doesn't get suggestion
			},
		},
		{
			name: "Large distance without suggestion",
			availableElements: map[string]bool{
				"my-button":      true,
				"other-element":  true,
				"third-element":  true,
				"fourth-element": true,
				"fifth-element":  true,
			},
			elementSources:           map[string]string{},
			expectedDiagnosticsCount: 3, // my-buttom, missing-import-element, completely-different-element
			expectedMessages: []string{
				"Unknown custom element 'my-buttom'. Did you mean 'my-button'?",                                                                                               // Close distance
				"Unknown custom element 'missing-import-element'. Available elements: 'fifth-element', 'fourth-element', 'my-button', 'other-element', 'third-element'",       // 5 elements, sorted
				"Unknown custom element 'completely-different-element'. Available elements: 'fifth-element', 'fourth-element', 'my-button', 'other-element', 'third-element'", // 5 elements, sorted
			},
			expectedTypes: []string{
				string(types.DiagnosticTypeTagSuggestion), // Close distance gets suggestion
				"", // Large distance doesn't get a suggestion type in diagnostic data
				"", // Large distance doesn't get a suggestion type in diagnostic data
			},
		},
		{
			name: "Large distance with many elements",
			availableElements: map[string]bool{
				"my-button":      true,
				"other-element":  true,
				"third-element":  true,
				"fourth-element": true,
				"fifth-element":  true,
				"sixth-element":  true, // More than 5, should show documentation message
			},
			elementSources:           map[string]string{},
			expectedDiagnosticsCount: 3,
			expectedMessages: []string{
				"Unknown custom element 'my-buttom'. Did you mean 'my-button'?",                                                                // Close distance still gets suggestion
				"Unknown custom element 'missing-import-element'. Check available elements in your project's manifest or documentation.",       // Many elements, show documentation message
				"Unknown custom element 'completely-different-element'. Check available elements in your project's manifest or documentation.", // Many elements, show documentation message
			},
			expectedTypes: []string{
				string(types.DiagnosticTypeTagSuggestion),
				"",
				"",
			},
		},
		{
			name: "All elements available",
			availableElements: map[string]bool{
				"my-button":                    true,
				"my-buttom":                    true, // Fixed the typo, so now available
				"missing-import-element":       true,
				"completely-different-element": true,
			},
			elementSources:           map[string]string{},
			expectedDiagnosticsCount: 0,
			expectedMessages:         []string{},
			expectedTypes:            []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &mockTagDiagnosticsContext{
				availableElements: tt.availableElements,
				elementSources:    tt.elementSources,
			}

			diagnostics := analyzeTagNameDiagnostics(ctx, ctx.Document(""))

			if len(diagnostics) != tt.expectedDiagnosticsCount {
				t.Errorf("Expected %d diagnostics, got %d", tt.expectedDiagnosticsCount, len(diagnostics))
			}

			for i, diagnostic := range diagnostics {
				if i < len(tt.expectedMessages) {
					if diagnostic.Message != tt.expectedMessages[i] {
						t.Errorf("Expected message '%s', got '%s'", tt.expectedMessages[i], diagnostic.Message)
					}
				}

				if i < len(tt.expectedTypes) {
					if diagnostic.Data != nil {
						if dataMap, ok := diagnostic.Data.(map[string]any); ok {
							if diagType, exists := dataMap["type"]; exists {
								if diagType != tt.expectedTypes[i] {
									t.Errorf("Expected diagnostic type '%s', got '%s'", tt.expectedTypes[i], diagType)
								}
							}
						}
					}
				}

				// Verify all diagnostics are from cem-lsp
				if diagnostic.Source == nil || *diagnostic.Source != "cem-lsp" {
					t.Errorf("Expected source 'cem-lsp', got %v", diagnostic.Source)
				}

				// Verify severity is Error
				if diagnostic.Severity == nil || *diagnostic.Severity != protocol.DiagnosticSeverityError {
					t.Errorf("Expected severity Error, got %v", diagnostic.Severity)
				}
			}
		})
	}
}

func TestTagNameValidation(t *testing.T) {
	tests := []struct {
		name     string
		tagName  string
		expected bool
	}{
		{"Valid custom element", "my-element", true},
		{"Valid with numbers", "my-element-2", true},
		{"Valid with multiple hyphens", "my-complex-element-name", true},
		{"Invalid no hyphen", "myelement", false},
		{"Invalid uppercase", "My-Element", false},
		{"Invalid xml prefix", "xml-element", false},
		{"Invalid xmlns prefix", "xmlns-element", false},
		{"Invalid empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidCustomElementName(tt.tagName)
			if result != tt.expected {
				t.Errorf("isValidCustomElementName('%s') = %v, expected %v", tt.tagName, result, tt.expected)
			}
		})
	}
}

func TestFindCustomElementTags(t *testing.T) {
	content := `<div class="container">
		<my-element attr="value">
			<nested-element>Content</nested-element>
		</my-element>
		<another-element/>
		<div class="regular-div">Not custom</div>
	</div>`

	matches := findCustomElementTags(content)

	expectedTags := []string{"my-element", "nested-element", "another-element"}
	if len(matches) != len(expectedTags) {
		t.Errorf("Expected %d matches, got %d", len(expectedTags), len(matches))
		for i, match := range matches {
			t.Logf("Match %d: '%s'", i, match.Value)
		}
	}

	for i, match := range matches {
		if i < len(expectedTags) {
			if match.Value != expectedTags[i] {
				t.Errorf("Expected tag '%s', got '%s'", expectedTags[i], match.Value)
			}
		}
	}
}
