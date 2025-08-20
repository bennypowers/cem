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
package lsp

import (
	"regexp"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

// ParsedElement represents a custom element found in a document
type ParsedElement struct {
	TagName    string
	Attributes map[string]string
	Range      protocol.Range
}

// ParsedAttribute represents an attribute found in a document
type ParsedAttribute struct {
	Name    string
	Value   string
	Range   protocol.Range
	TagName string // The element this attribute belongs to
}

// DocumentParser provides parsing capabilities for HTML and TypeScript files
type DocumentParser struct {
	content string
	uri     string
	lines   []string
}

// NewDocumentParser creates a new parser for the given document content
func NewDocumentParser(uri, content string) *DocumentParser {
	return &DocumentParser{
		content: content,
		uri:     uri,
		lines:   strings.Split(content, "\n"),
	}
}

// ParseCustomElements finds all custom elements in the document
func (p *DocumentParser) ParseCustomElements() []ParsedElement {
	if strings.HasSuffix(p.uri, ".html") {
		return p.parseHTMLElements()
	} else if strings.HasSuffix(p.uri, ".ts") || strings.HasSuffix(p.uri, ".js") {
		return p.parseTypeScriptElements()
	}
	return []ParsedElement{}
}

// parseHTMLElements finds custom elements in HTML content
func (p *DocumentParser) parseHTMLElements() []ParsedElement {
	var elements []ParsedElement

	// Regex to match custom element tags (must contain a hyphen)
	// Matches opening tags like <my-element>, <my-element attr="value">
	tagRegex := regexp.MustCompile(`<([a-z][a-z0-9]*-[a-z0-9-]*)\b[^>]*>`)

	matches := tagRegex.FindAllStringSubmatch(p.content, -1)
	for _, match := range matches {
		if len(match) >= 2 {
			tagName := match[1]

			// Find the position of this match in the document
			matchStart := strings.Index(p.content, match[0])
			if matchStart == -1 {
				continue
			}

			// Convert byte offset to line/character position
			startPos := p.offsetToPosition(matchStart)
			endPos := p.offsetToPosition(matchStart + len(tagName) + 1) // +1 for '<'

			elements = append(elements, ParsedElement{
				TagName: tagName,
				Range: protocol.Range{
					Start: startPos,
					End:   endPos,
				},
			})
		}
	}

	return elements
}

// parseTypeScriptElements finds custom elements in TypeScript template literals
func (p *DocumentParser) parseTypeScriptElements() []ParsedElement {
	var elements []ParsedElement

	// Look for template literals that contain HTML
	// This is a simplified approach - in practice you'd want to use tree-sitter
	templateRegex := regexp.MustCompile("(?s)`[^`]*<[^>]*>[^`]*`|html`[^`]*`")
	templateMatches := templateRegex.FindAllString(p.content, -1)

	for _, template := range templateMatches {
		// Find custom elements within this template literal
		tagRegex := regexp.MustCompile(`<([a-z][a-z0-9]*-[a-z0-9-]*)\b[^>]*>`)
		matches := tagRegex.FindAllStringSubmatch(template, -1)

		for _, match := range matches {
			if len(match) >= 2 {
				tagName := match[1]

				// Find position within the full document
				templateStart := strings.Index(p.content, template)
				if templateStart == -1 {
					continue
				}

				tagStart := strings.Index(template, match[0])
				if tagStart == -1 {
					continue
				}

				absoluteStart := templateStart + tagStart
				startPos := p.offsetToPosition(absoluteStart)
				endPos := p.offsetToPosition(absoluteStart + len(tagName) + 1)

				elements = append(elements, ParsedElement{
					TagName: tagName,
					Range: protocol.Range{
						Start: startPos,
						End:   endPos,
					},
				})
			}
		}
	}

	return elements
}

// FindElementAtPosition finds the custom element at the given position
func (p *DocumentParser) FindElementAtPosition(position protocol.Position) *ParsedElement {
	elements := p.ParseCustomElements()

	for _, element := range elements {
		if p.isPositionInRange(position, element.Range) {
			return &element
		}
	}

	return nil
}

// FindAttributeAtPosition finds the attribute at the given position
func (p *DocumentParser) FindAttributeAtPosition(position protocol.Position) *ParsedAttribute {
	// This is a simplified implementation
	// In a real implementation, you'd parse the full element structure
	lineContent := ""
	if int(position.Line) < len(p.lines) {
		lineContent = p.lines[position.Line]
	}

	// Look for attribute patterns around the cursor position
	// This is a basic implementation that could be improved
	attrRegex := regexp.MustCompile(`(\w+(?:-\w+)*)\s*=\s*["'][^"']*["']`)
	matches := attrRegex.FindAllStringSubmatch(lineContent, -1)

	for _, match := range matches {
		if len(match) >= 2 {
			attrName := match[1]
			attrStart := strings.Index(lineContent, match[0])
			if attrStart == -1 {
				continue
			}

			// Check if the position is within this attribute
			if position.Character >= uint32(attrStart) &&
				position.Character <= uint32(attrStart+len(attrName)) {
				return &ParsedAttribute{
					Name: attrName,
					Range: protocol.Range{
						Start: protocol.Position{Line: position.Line, Character: uint32(attrStart)},
						End:   protocol.Position{Line: position.Line, Character: uint32(attrStart + len(attrName))},
					},
				}
			}
		}
	}

	return nil
}

// offsetToPosition converts a byte offset to line/character position
func (p *DocumentParser) offsetToPosition(offset int) protocol.Position {
	line := 0
	char := 0

	for i, r := range p.content {
		if i >= offset {
			break
		}

		if r == '\n' {
			line++
			char = 0
		} else {
			char++
		}
	}

	return protocol.Position{
		Line:      uint32(line),
		Character: uint32(char),
	}
}

// isPositionInRange checks if a position is within a range
func (p *DocumentParser) isPositionInRange(pos protocol.Position, r protocol.Range) bool {
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
