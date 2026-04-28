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
package queries

import (
	"embed"
	"fmt"
	"strings"

	ts "github.com/tree-sitter/go-tree-sitter"
)

//go:embed */*.scm
var queries embed.FS

type NoCaptureError struct {
	Capture string
	Query   string
}

func (e *NoCaptureError) Error() string {
	return fmt.Sprintf("No nodes for capture %s in query %s", e.Capture, e.Query)
}

// QuerySelector defines which queries to load for performance
type QuerySelector struct {
	HTML       []string // HTML query names to load
	TypeScript []string // TypeScript query names to load
	CSS        []string // CSS query names to load
	JSDoc      []string // JSDoc query names to load
	TSX        []string // TSX query names to load
	Blade      []string // Blade query names to load (reuses html/*.scm)
}

// AllQueries returns a selector that loads all available queries
func AllQueries() QuerySelector {
	return QuerySelector{
		HTML:       []string{"slotsAndParts", "customElements", "completionContext"},
		TypeScript: []string{"classMemberDeclaration", "classes", "declarations", "imports", "htmlTemplates", "completionContext"},
		CSS:        []string{"cssCustomProperties"},
		JSDoc:      []string{"jsdoc"},
	}
}

// LSPQueries returns a selector that loads only queries needed for LSP
func LSPQueries() QuerySelector {
	return QuerySelector{
		HTML:       []string{"customElements", "completionContext", "scriptTags", "headElements", "attributes"},
		TypeScript: []string{"htmlTemplates", "completionContext", "imports", "classes", "classMemberDeclaration", "exports", "importAttributes", "definedElements"},
		CSS:        []string{},
		JSDoc:      []string{},
		TSX:        []string{"customElements", "completionContext"},
		Blade:      []string{"customElements", "completionContext", "scriptTags", "headElements", "attributes"},
	}
}

// GenerateQueries returns a selector that loads only queries needed for generate
func GenerateQueries() QuerySelector {
	return QuerySelector{
		HTML:       []string{"slotsAndParts"},
		TypeScript: []string{"classMemberDeclaration", "classes", "declarations", "imports", "typeAliases", "exports"},
		CSS:        []string{"cssCustomProperties"},
		JSDoc:      []string{"jsdoc"},
	}
}

type ParentNodeCaptures struct {
	NodeId   uintptr
	Captures CaptureMap
}

type CaptureInfo struct {
	NodeId    int
	Text      string
	StartByte uint
	EndByte   uint
}

type CaptureMap = map[string][]CaptureInfo

func GetDescendantById(root *ts.Node, id int) *ts.Node {
	c := root.Walk()
	defer c.Close()
	var find func(node *ts.Node) *ts.Node
	find = func(node *ts.Node) *ts.Node {
		if int(node.Id()) == id {
			return node
		}
		for i := range int(node.ChildCount()) {
			child := node.Child(uint(i))
			if child == nil {
				continue
			}
			if res := find(child); res != nil {
				return res
			}
		}
		return nil
	}

	return find(root)
}

// Position represents a line/character position
type Position struct {
	Line      uint32
	Character uint32
}

// Range represents a start/end range
type Range struct {
	Start Position
	End   Position
}

// NodeToRange converts a tree-sitter node to a Range using byte-to-position conversion
func NodeToRange(node *ts.Node, content []byte) Range {
	start := byteOffsetToPosition(content, node.StartByte())
	end := byteOffsetToPosition(content, node.EndByte())
	return Range{
		Start: start,
		End:   end,
	}
}

// Helper functions

// byteOffsetToPosition converts a byte offset to line/character position
func byteOffsetToPosition(content []byte, offset uint) Position {
	line := uint32(0)
	char := uint32(0)

	for i, b := range content {
		if uint(i) >= offset {
			break
		}

		if b == '\n' {
			line++
			char = 0
		} else {
			char++
		}
	}

	return Position{
		Line:      line,
		Character: char,
	}
}

func matchesAttribute(memberName, decoratorAttrName, targetAttr string) bool {
	// If decorator explicitly sets attribute name, use that
	if decoratorAttrName != "" {
		return decoratorAttrName == targetAttr
	}

	// Otherwise, check if member name matches (camelCase property → kebab-case attribute)
	if memberName == targetAttr {
		return true
	}

	// Convert camelCase to kebab-case for comparison
	kebabCase := camelToKebab(memberName)
	return kebabCase == targetAttr
}

func camelToKebab(camelCase string) string {
	var result strings.Builder
	for i, r := range camelCase {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteByte('-')
		}
		result.WriteRune(r - ('A' - 'a'))
	}
	return result.String()
}

func findSlotInTemplate(htmlContent string, slotName string, queryManager *QueryManager) *Range {
	// Parse the HTML template content
	parser := GetHTMLParser()
	defer PutHTMLParser(parser)

	tree := parser.Parse([]byte(htmlContent), nil)
	if tree == nil {
		return nil
	}
	defer tree.Close()

	// Get HTML slot queries
	templateQueries, err := NewQueryMatcher(queryManager, "html", "slotsAndParts")
	if err != nil {
		return nil
	}
	defer templateQueries.Close()

	htmlBytes := []byte(htmlContent)
	// Find slot with matching name
	for captureMap := range templateQueries.ParentCaptures(tree.RootNode(), htmlBytes, "slot") {
		if attrValues, ok := captureMap["attr.value"]; ok {
			for _, attrValue := range attrValues {
				if attrValue.Text == slotName {
					node := GetDescendantById(tree.RootNode(), attrValue.NodeId)
					if node != nil {
						r := NodeToRange(node, htmlBytes)
						return &r
					}
				}
			}
		}
	}

	return nil
}

func adjustTemplateRange(templateRange *Range, templateNode *ts.Node, content []byte) *Range {
	// Convert the template node's position to line/character
	templateStart := byteOffsetToPosition(content, templateNode.StartByte())

	return &Range{
		Start: Position{
			Line:      templateStart.Line + templateRange.Start.Line,
			Character: templateStart.Character + templateRange.Start.Character,
		},
		End: Position{
			Line:      templateStart.Line + templateRange.End.Line,
			Character: templateStart.Character + templateRange.End.Character,
		},
	}
}
