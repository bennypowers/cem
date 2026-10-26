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
		return formatHTML, nil
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
		ranges := jinjaHTMLRanges(source)
		if len(ranges) > 0 {
			if err := parser.SetIncludedRanges(ranges); err != nil {
				return nil, fmt.Errorf("markup: failed to set template ranges: %w", err)
			}
			defer func() { _ = parser.SetIncludedRanges(nil) }()
		}
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
	for i := uint(0); i < node.ChildCount(); i++ {
		child := node.Child(i)
		if child == nil {
			continue
		}

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
		tagName: string(source[tagName.StartByte():tagName.EndByte()]),
		line:    int(tagName.StartPosition().Row) + 1,
		column:  int(tagName.StartPosition().Column) + 1,
	}

	for i := uint(0); i < startTag.ChildCount(); i++ {
		child := startTag.Child(i)
		if child == nil || child.Kind() != "attribute" {
			continue
		}

		nameNode := findChildKind(child, "attribute_name")
		if nameNode == nil {
			continue
		}

		el.attributes = append(el.attributes, parsedAttribute{
			name:   string(source[nameNode.StartByte():nameNode.EndByte()]),
			value:  attributeValue(child, source),
			line:   int(nameNode.StartPosition().Row) + 1,
			column: int(nameNode.StartPosition().Column) + 1,
		})
	}

	return el
}

func attributeValue(attribute *ts.Node, source []byte) string {
	for i := uint(0); i < attribute.ChildCount(); i++ {
		child := attribute.Child(i)
		if child == nil {
			continue
		}
		switch child.Kind() {
		case "quoted_attribute_value":
			if value := findChildKind(child, "attribute_value"); value != nil {
				return string(source[value.StartByte():value.EndByte()])
			}
			// Interpolated values have no attribute_value node; the value is
			// dynamic, so leave it empty and skip value validation.
			return ""
		case "attribute_value":
			return string(source[child.StartByte():child.EndByte()])
		}
	}
	return ""
}

func findChildKind(node *ts.Node, kind string) *ts.Node {
	for i := uint(0); i < node.ChildCount(); i++ {
		child := node.Child(i)
		if child != nil && child.Kind() == kind {
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
	jinjaTextQueryErr  error
)

// jinjaHTMLRanges returns the source byte ranges that the Jinja grammar marks
// as plain text (that is, HTML regions). The HTML parser is restricted to
// those ranges so template statements and expressions are not parsed as
// markup. Unbalanced control flow still yields usable text ranges.
func jinjaHTMLRanges(source string) []ts.Range {
	jinjaTextQueryOnce.Do(func() {
		jinjaTextQuery, jinjaTextQueryErr = ts.NewQuery(jinjalang.TSLanguage(), `(text) @html`)
	})
	if jinjaTextQueryErr != nil {
		return nil
	}

	parser := jinjalang.BorrowParser()
	if parser == nil {
		return nil
	}
	defer jinjalang.ReturnParser(parser)

	sourceBytes := []byte(source)
	tree := parser.Parse(sourceBytes, nil)
	if tree == nil {
		return nil
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
	return ranges
}
