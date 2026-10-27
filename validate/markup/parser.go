/*
Copyright © 2026 Benny Powers <web@bennypowers.com>

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
package markup

import (
	"fmt"
	"html"
	"strings"
	"sync"

	htmllang "bennypowers.dev/cem/internal/languages/html"
	jinjalang "bennypowers.dev/cem/internal/languages/jinja"
	"bennypowers.dev/cem/internal/validations"
	ts "github.com/tree-sitter/go-tree-sitter"
)

type formatFamilyValue int

const (
	formatHTML formatFamilyValue = iota
	formatJinja
	// formatHandlebars parses with the HTML grammar but keeps template value
	// semantics, so "{{ ... }}" attribute values are not validated literally.
	formatHandlebars
)

type elementNode struct {
	tagName    string
	line       int
	column     int
	attributes []parsedAttribute
	children   []*elementNode
}

type parsedAttribute struct {
	name   string
	value  string
	line   int
	column int
}

func formatFamily(format string) (formatFamilyValue, error) {
	switch strings.ToLower(strings.TrimSpace(format)) {
	case "", "html", "htm":
		return formatHTML, nil
	case "twig", "njk", "jinja", "jinja2", "j2", "liquid":
		return formatJinja, nil
	case "hbs", "handlebars":
		// Parsing Handlebars as HTML is best effort; linking the Handlebars
		// grammar would pull an unnecessary dependency into this package.
		// Template values still use template semantics.
		return formatHandlebars, nil
	default:
		return 0, fmt.Errorf("markup: unsupported format %q", format)
	}
}

func parseElements(source string, family formatFamilyValue) ([]*elementNode, error) {
	parser := htmllang.BorrowParser()
	if parser == nil {
		return nil, fmt.Errorf("markup: failed to borrow HTML parser")
	}
	defer htmllang.ReturnParser(parser)

	if family == formatJinja {
		ranges, err := jinjaHTMLRanges(source)
		if err != nil {
			return nil, err
		}
		// With no text ranges there is no HTML to validate. Parsing the raw
		// template would treat template syntax as markup, so stop here.
		if len(ranges) == 0 {
			return nil, nil
		}
		if err := parser.SetIncludedRanges(ranges); err != nil {
			return nil, fmt.Errorf("markup: failed to set template ranges: %w", err)
		}
		defer func() { _ = parser.SetIncludedRanges(nil) }()
	}

	sourceBytes := []byte(source)
	tree := parser.Parse(sourceBytes, nil)
	if tree == nil {
		return nil, fmt.Errorf("markup: failed to parse markup")
	}
	defer tree.Close()

	return elementsFromNode(tree.RootNode(), sourceBytes), nil
}

func elementsFromNode(node *ts.Node, source []byte) []*elementNode {
	var elements []*elementNode
	cursor := node.Walk()
	defer cursor.Close()

	for ok := cursor.GotoFirstChild(); ok; ok = cursor.GotoNextSibling() {
		child := cursor.Node()
		switch child.Kind() {
		case "element":
			if el := parseTag(child, source); el != nil {
				el.children = elementsFromNode(child, source)
				elements = append(elements, el)
			}
		case "self_closing_tag":
			if el := parseTag(child, source); el != nil {
				elements = append(elements, el)
			}
		default:
			elements = append(elements, elementsFromNode(child, source)...)
		}
	}
	return elements
}

func parseTag(node *ts.Node, source []byte) *elementNode {
	startTag := findChildKind(node, "start_tag")
	if startTag == nil {
		startTag = node
	}

	tagName := findChildKind(startTag, "tag_name")
	if tagName == nil {
		return nil
	}

	el := &elementNode{
		tagName: asciiLower(string(source[tagName.StartByte():tagName.EndByte()])),
	}
	el.line, el.column = positionFor(tagName)

	// Duplicate attributes keep their first value and position, matching HTML
	// parsing and browser behavior. Later duplicates are ignored.
	index := make(map[string]int, startTag.ChildCount())
	cursor := startTag.Walk()
	defer cursor.Close()

	for ok := cursor.GotoFirstChild(); ok; ok = cursor.GotoNextSibling() {
		child := cursor.Node()
		if child.Kind() != "attribute" {
			continue
		}

		nameNode := findChildKind(child, "attribute_name")
		if nameNode == nil {
			continue
		}

		name := asciiLower(string(source[nameNode.StartByte():nameNode.EndByte()]))
		line, column := positionFor(nameNode)
		attr := parsedAttribute{
			name:   name,
			value:  attributeValue(child, source),
			line:   line,
			column: column,
		}

		if _, ok := index[name]; ok {
			continue
		}
		index[name] = len(el.attributes)
		el.attributes = append(el.attributes, attr)
	}

	return el
}

func asciiLower(value string) string {
	for i := 0; i < len(value); i++ {
		if value[i] < 'A' || value[i] > 'Z' {
			continue
		}
		lower := []byte(value)
		for j := i; j < len(lower); j++ {
			if lower[j] >= 'A' && lower[j] <= 'Z' {
				lower[j] += 'a' - 'A'
			}
		}
		return string(lower)
	}
	return value
}

// unescapeStaticValue decodes HTML character references in a static attribute
// value. Interpolated template values have no attribute_value node and never
// reach this path.
func unescapeStaticValue(raw []byte) string {
	return html.UnescapeString(string(raw))
}

// positionFor returns a one-based line and a one-based UTF-8 byte column.
// Tree-sitter reports columns as byte offsets within the line, and using them
// directly keeps position lookup O(1) for every node.
func positionFor(node *ts.Node) (line, column int) {
	point := node.StartPosition()
	return int(point.Row) + 1, int(point.Column) + 1
}

func attributeValue(attribute *ts.Node, source []byte) string {
	cursor := attribute.Walk()
	defer cursor.Close()

	for ok := cursor.GotoFirstChild(); ok; ok = cursor.GotoNextSibling() {
		child := cursor.Node()
		switch child.Kind() {
		case "quoted_attribute_value":
			if value := findChildKind(child, "attribute_value"); value != nil {
				return unescapeStaticValue(source[value.StartByte():value.EndByte()])
			}
			// Interpolated values have no attribute_value node; the value is
			// dynamic, so leave it empty and skip value validation.
			return ""
		case "attribute_value":
			return unescapeStaticValue(source[child.StartByte():child.EndByte()])
		}
	}
	return ""
}

func findChildKind(node *ts.Node, kind string) *ts.Node {
	cursor := node.Walk()
	defer cursor.Close()

	for ok := cursor.GotoFirstChild(); ok; ok = cursor.GotoNextSibling() {
		child := cursor.Node()
		if child.Kind() == kind {
			return child
		}
	}
	return nil
}

// isCustomElementTag matches the same shape the LSP and MCP tree-sitter
// queries use: ^[a-z][a-z0-9]*-[a-z0-9-]*$.
func isCustomElementTag(tagName string) bool {
	if len(tagName) < 2 {
		return false
	}
	if tagName[0] < 'a' || tagName[0] > 'z' {
		return false
	}

	hyphen := -1
	for i := 1; i < len(tagName); i++ {
		c := tagName[i]
		if c == '-' {
			hyphen = i
			break
		}
		if !isLowerDigit(c) {
			return false
		}
	}
	if hyphen == -1 {
		return false
	}

	for i := hyphen + 1; i < len(tagName); i++ {
		c := tagName[i]
		if !isLowerDigit(c) && c != '-' {
			return false
		}
	}
	return true
}

func isLowerDigit(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

func isGlobalAttribute(name string) bool {
	return validations.IsGlobalAttribute(name)
}

var (
	jinjaTextQueryOnce sync.Once
	jinjaTextQuery     *ts.Query
	// ts.NewQuery returns *ts.QueryError, not error. Keeping the concrete
	// pointer avoids storing a typed nil in a non-nil error interface.
	jinjaTextQueryErr *ts.QueryError
)

// jinjaHTMLRanges returns the source byte ranges that the Jinja grammar marks
// as plain text (that is, HTML regions). The HTML parser is restricted to
// those ranges so template statements and expressions are not parsed as
// markup. Unbalanced control flow still yields usable text ranges.
func jinjaHTMLRanges(source string) ([]ts.Range, error) {
	jinjaTextQueryOnce.Do(func() {
		jinjaTextQuery, jinjaTextQueryErr = ts.NewQuery(jinjalang.TSLanguage(), `(text) @html`)
	})
	if jinjaTextQueryErr != nil {
		return nil, fmt.Errorf("markup: failed to compile Jinja text query: %w", jinjaTextQueryErr)
	}

	parser := jinjalang.BorrowParser()
	if parser == nil {
		return nil, fmt.Errorf("markup: failed to borrow Jinja parser")
	}
	defer jinjalang.ReturnParser(parser)

	sourceBytes := []byte(source)
	tree := parser.Parse(sourceBytes, nil)
	if tree == nil {
		return nil, fmt.Errorf("markup: failed to parse template")
	}
	defer tree.Close()

	cursor := ts.NewQueryCursor()
	defer cursor.Close()

	var ranges []ts.Range
	matches := cursor.Matches(jinjaTextQuery, tree.RootNode(), sourceBytes)
	for match := matches.Next(); match != nil; match = matches.Next() {
		for _, capture := range match.Captures {
			ranges = append(ranges, capture.Node.Range())
		}
	}
	return ranges, nil
}
