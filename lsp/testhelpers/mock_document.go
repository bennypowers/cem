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
package testhelpers

import (
	"regexp"
	"strings"

	"bennypowers.dev/cem/lsp/types"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// MockDocument provides a unified mock implementation of types.Document for testing
type MockDocument struct {
	ContentStr      string
	URIStr          string
	VersionNum      int32
	CustomElements  []types.CustomElementMatch
	ScriptTagsList  []types.ScriptTag
	TemplateContext string
}

func (m *MockDocument) Content() (string, error) {
	return m.ContentStr, nil
}

func (m *MockDocument) URI() string {
	if m.URIStr == "" {
		return "test://test.html"
	}
	return m.URIStr
}

func (m *MockDocument) Version() int32 {
	return m.VersionNum
}

func (m *MockDocument) FindElementAtPosition(position protocol.Position, dm any) *types.CustomElementMatch {
	// Simple implementation that finds elements based on position within their range
	for _, element := range m.CustomElements {
		if isPositionInRange(position, element.Range) {
			return &element
		}
	}
	return nil
}

func (m *MockDocument) FindAttributeAtPosition(position protocol.Position, dm any) (*types.AttributeMatch, string) {
	// Find element and then check its attributes
	for _, element := range m.CustomElements {
		for _, attr := range element.Attributes {
			if isPositionInRange(position, attr.Range) {
				return &attr, element.TagName
			}
		}
	}
	return nil, ""
}

func (m *MockDocument) FindCustomElements(dm any) ([]types.CustomElementMatch, error) {
	return m.CustomElements, nil
}

func (m *MockDocument) AnalyzeCompletionContextTS(position protocol.Position, dm any) *types.CompletionAnalysis {
	// For testing template behavior, we need to analyze the content more carefully
	if m.TemplateContext == "" {
		return nil // Fall back to regex analysis
	}

	// Get the line content at position
	lines := strings.Split(m.ContentStr, "\n")
	if int(position.Line) >= len(lines) {
		return nil
	}
	lineContent := lines[position.Line]
	beforeCursor := ""
	if int(position.Character) <= len(lineContent) {
		beforeCursor = lineContent[:position.Character]
	}

	// Basic analysis for template context
	analysis := &types.CompletionAnalysis{
		IsLitTemplate: true,
		LineContent:   lineContent,
	}

	// Detect Lit syntax patterns - be more specific about patterns
	if regexp.MustCompile(`<[a-z][a-z0-9-]*\s+@[a-z]*$`).MatchString(beforeCursor) {
		// Event binding: <my-element @cli|
		analysis.Type = types.CompletionLitEventBinding
		analysis.LitSyntax = "@"
		if tagMatch := regexp.MustCompile(`<([a-z][a-z0-9-]*)`).FindStringSubmatch(beforeCursor); len(tagMatch) > 1 {
			analysis.TagName = tagMatch[1]
		}
	} else if regexp.MustCompile(`<[a-z][a-z0-9-]*\s+\.[a-z]*$`).MatchString(beforeCursor) {
		// Property binding: <my-element .prop|
		analysis.Type = types.CompletionLitPropertyBinding
		analysis.LitSyntax = "."
		if tagMatch := regexp.MustCompile(`<([a-z][a-z0-9-]*)`).FindStringSubmatch(beforeCursor); len(tagMatch) > 1 {
			analysis.TagName = tagMatch[1]
		}
	} else if regexp.MustCompile(`<[a-z][a-z0-9-]*\s+\?[a-z]*$`).MatchString(beforeCursor) {
		// Boolean attribute: <my-element ?disab|
		analysis.Type = types.CompletionLitBooleanAttribute
		analysis.LitSyntax = "?"
		if tagMatch := regexp.MustCompile(`<([a-z][a-z0-9-]*)`).FindStringSubmatch(beforeCursor); len(tagMatch) > 1 {
			analysis.TagName = tagMatch[1]
		}
	} else if regexp.MustCompile(`<[a-z][a-z0-9-]*\s+[a-z-]*$`).MatchString(beforeCursor) {
		// Attribute name completion: <my-element dis|
		analysis.Type = types.CompletionAttributeName
		if tagMatch := regexp.MustCompile(`<([a-z][a-z0-9-]*)`).FindStringSubmatch(beforeCursor); len(tagMatch) > 1 {
			analysis.TagName = tagMatch[1]
		}
	} else {
		// Tag name completion: <my-el|
		analysis.Type = types.CompletionTagName
		analysis.TagName = m.TemplateContext
	}

	return analysis
}

func (m *MockDocument) GetScriptTags() []types.ScriptTag {
	return m.ScriptTagsList
}

func (m *MockDocument) FindModuleScript() (protocol.Position, bool) {
	for _, script := range m.ScriptTagsList {
		if script.IsModule {
			return script.ContentRange.End, true
		}
	}
	return protocol.Position{}, false
}

func (m *MockDocument) FindInlineModuleScript() (protocol.Position, bool) {
	for _, script := range m.ScriptTagsList {
		if script.IsModule && script.Src == "" {
			return script.ContentRange.End, true
		}
	}
	return protocol.Position{}, false
}

func (m *MockDocument) FindHeadInsertionPoint(dm any) (protocol.Position, bool) {
	// Simple mock: if content contains <head>, return position after it
	if strings.Contains(m.ContentStr, "<head>") {
		lines := strings.Split(m.ContentStr, "\n")
		for i, line := range lines {
			if strings.Contains(line, "<head>") {
				return protocol.Position{Line: uint32(i + 1), Character: 0}, true
			}
		}
	}
	return protocol.Position{}, false
}

func (m *MockDocument) ByteRangeToProtocolRange(content string, startByte, endByte uint) protocol.Range {
	return protocol.Range{
		Start: protocol.Position{Line: 0, Character: uint32(startByte)},
		End:   protocol.Position{Line: 0, Character: uint32(endByte)},
	}
}

// Helper function to check if position is within range
func isPositionInRange(pos protocol.Position, r protocol.Range) bool {
	if pos.Line < r.Start.Line || pos.Line > r.End.Line {
		return false
	}
	if pos.Line == r.Start.Line && pos.Character < r.Start.Character {
		return false
	}
	if pos.Line == r.End.Line && pos.Character > r.End.Character {
		return false
	}
	return true
}

// NewMockDocument creates a new MockDocument with sensible defaults
func NewMockDocument(content string) *MockDocument {
	return &MockDocument{
		ContentStr: content,
		URIStr:     "test://test.html",
		VersionNum: 1,
	}
}

// NewMockDocumentWithElements creates a MockDocument with predefined elements
func NewMockDocumentWithElements(content string, elements []types.CustomElementMatch) *MockDocument {
	return &MockDocument{
		ContentStr:     content,
		URIStr:         "test://test.html",
		VersionNum:     1,
		CustomElements: elements,
	}
}
