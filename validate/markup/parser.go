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
	"sort"
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
	// formatLiquid uses the Jinja grammar but excludes Liquid comment bodies.
	formatLiquid
	// formatHandlebars parses with the HTML grammar but masks Handlebars
	// expressions, so they are never treated as attributes or markup.
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
	name    string
	value   string
	line    int
	column  int
	dynamic bool
}

func formatFamily(format string) (formatFamilyValue, error) {
	switch strings.ToLower(strings.TrimSpace(format)) {
	case "", "html", "htm":
		return formatHTML, nil
	case "twig", "njk", "jinja", "jinja2", "j2":
		return formatJinja, nil
	case "liquid":
		return formatLiquid, nil
	case "hbs", "handlebars":
		// Parsing Handlebars as HTML is best effort; linking the Handlebars
		// grammar would pull an unnecessary dependency into this package.
		// Template values still use template semantics.
		return formatHandlebars, nil
	default:
		return 0, fmt.Errorf("%w %q", ErrUnsupportedFormat, format)
	}
}

func parseElements(source string, family formatFamilyValue) ([]*elementNode, error) {
	parser := htmllang.BorrowParser()
	if parser == nil {
		return nil, fmt.Errorf("markup: failed to borrow HTML parser")
	}
	defer htmllang.ReturnParser(parser)

	var dynamicRanges []ts.Range
	switch family {
	case formatJinja, formatLiquid:
		textRanges, dynamic, err := analyzeJinjaRanges(source, family == formatLiquid)
		if err != nil {
			return nil, err
		}
		// With no text ranges there is no HTML to validate. Parsing the raw
		// template would treat template syntax as markup, so stop here.
		if len(textRanges) == 0 {
			return nil, nil
		}
		if err := parser.SetIncludedRanges(textRanges); err != nil {
			return nil, fmt.Errorf("markup: failed to set template ranges: %w", err)
		}
		defer func() { _ = parser.SetIncludedRanges(nil) }()
		dynamicRanges = dynamic
	case formatHandlebars:
		dynamicRanges = handlebarsRanges(source)
		textRanges := complementRanges(source, dynamicRanges)
		if len(textRanges) == 0 {
			return nil, nil
		}
		if err := parser.SetIncludedRanges(textRanges); err != nil {
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

	return elementsFromNode(tree.RootNode(), sourceBytes, dynamicRanges, false), nil
}

// elementsFromNode walks structural nodes. When skipTags is set, the node is
// an element's own tag container, so its start/end tags are not emitted as
// separate elements. ERROR nodes that wrap a start_tag are still traversed.
func elementsFromNode(node *ts.Node, source []byte, dynamicRanges []ts.Range, skipTags bool) []*elementNode {
	var elements []*elementNode
	cursor := node.Walk()
	defer cursor.Close()

	for ok := cursor.GotoFirstChild(); ok; ok = cursor.GotoNextSibling() {
		child := cursor.Node()
		switch child.Kind() {
		case "element":
			if el := parseTag(child, source, dynamicRanges); el != nil {
				if !isTextContentElement(el.tagName) {
					el.children = elementsFromNode(child, source, dynamicRanges, true)
				}
				elements = append(elements, el)
			} else {
				elements = append(elements, elementsFromNode(child, source, dynamicRanges, false)...)
			}
		case "self_closing_tag", "start_tag":
			if skipTags {
				continue
			}
			if el := parseTag(child, source, dynamicRanges); el != nil {
				elements = append(elements, el)
			}
		case "end_tag":
			// End tags never yield elements, in either traversal mode.
		default:
			elements = append(elements, elementsFromNode(child, source, dynamicRanges, false)...)
		}
	}
	return elements
}

func parseTag(node *ts.Node, source []byte, dynamicRanges []ts.Range) *elementNode {
	startTag := findChildKind(node, "start_tag")
	if startTag == nil {
		// Self-closing tags are nested as element -> self_closing_tag.
		startTag = findChildKind(node, "self_closing_tag")
	}
	if startTag == nil {
		// The node itself may be a start_tag or self_closing_tag, for example
		// when it is reached through an ERROR node for an unclosed tag.
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
			name:    name,
			value:   attributeValue(child, source),
			line:    line,
			column:  column,
			dynamic: overlapsAny(dynamicRanges, child.StartByte(), child.EndByte()),
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

// isTextContentElement reports whether HTML parses the element's content as
// text (raw text or RCDATA), so markup-looking text inside it is not an
// element. This mirrors the HTML parsing rules browsers apply.
func isTextContentElement(tagName string) bool {
	switch tagName {
	case "iframe", "noembed", "noframes", "noscript", "plaintext",
		"script", "style", "textarea", "title", "xmp":
		return true
	}
	return false
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

// analyzeJinjaRanges parses a template once and returns both the HTML text
// ranges and the ranges occupied by template syntax. Dynamic ranges let the
// HTML pass mark attributes whose values mix static text with expressions.
// Liquid comments are removed from the text ranges because they never render.
func analyzeJinjaRanges(source string, liquid bool) (text, dynamic []ts.Range, err error) {
	jinjaTextQueryOnce.Do(func() {
		jinjaTextQuery, jinjaTextQueryErr = ts.NewQuery(jinjalang.TSLanguage(), `(text) @html`)
	})
	if jinjaTextQueryErr != nil {
		return nil, nil, fmt.Errorf("markup: failed to compile Jinja text query: %w", jinjaTextQueryErr)
	}

	parser := jinjalang.BorrowParser()
	if parser == nil {
		return nil, nil, fmt.Errorf("markup: failed to borrow Jinja parser")
	}
	defer jinjalang.ReturnParser(parser)

	sourceBytes := []byte(source)
	tree := parser.Parse(sourceBytes, nil)
	if tree == nil {
		return nil, nil, fmt.Errorf("markup: failed to parse template")
	}
	defer tree.Close()

	cursor := ts.NewQueryCursor()
	defer cursor.Close()

	matches := cursor.Matches(jinjaTextQuery, tree.RootNode(), sourceBytes)
	for match := matches.Next(); match != nil; match = matches.Next() {
		for _, capture := range match.Captures {
			text = append(text, capture.Node.Range())
		}
	}

	dynamic = jinjaDynamicRanges(tree.RootNode())
	sort.Slice(dynamic, func(i, j int) bool {
		return dynamic[i].EndByte < dynamic[j].EndByte
	})
	if liquid {
		text = excludeRanges(source, text, liquidCommentSpans(source))
	}
	return text, dynamic, nil
}

// handlebarsRanges returns the byte ranges occupied by Handlebars expressions
// and comments. Masking them before the HTML parse keeps block helpers such as
// {{#if}} and helper string arguments from becoming attributes or elements.
// Unterminated expressions are left as HTML so later markup is still checked.
func handlebarsRanges(source string) []ts.Range {
	tracker := &pointTracker{source: []byte(source)}
	var ranges []ts.Range

	for i := 0; i+1 < len(source); {
		rel := strings.Index(source[i:], "{{")
		if rel < 0 {
			break
		}
		start := i + rel
		width := 2
		if start+2 < len(source) && source[start+2] == '{' {
			width = 3
		}
		closing := strings.Repeat("}", width)

		j := start + width
		var quote byte
		end := -1
		for j < len(source) {
			c := source[j]
			if quote != 0 {
				if c == '\\' {
					j += 2
					continue
				}
				if c == quote {
					quote = 0
				}
				j++
				continue
			}
			if c == '\'' || c == '"' {
				quote = c
				j++
				continue
			}
			if strings.HasPrefix(source[j:], closing) {
				end = j + width
				break
			}
			j++
		}
		if end < 0 {
			break
		}
		ranges = append(ranges, tracker.between(start, end))
		i = end
	}
	return ranges
}

// liquidCommentSpans returns the ranges of Liquid comment blocks. Comments do
// not nest, so the first endcomment closes the block.
func liquidCommentSpans(source string) []ts.Range {
	tracker := &pointTracker{source: []byte(source)}
	var spans []ts.Range

	for i := 0; i < len(source); {
		start, end, name, ok := nextLiquidTag(source, i)
		if !ok {
			break
		}
		if name != "comment" {
			i = end
			continue
		}

		closed := false
		for j := end; j < len(source); {
			_, end2, name2, ok2 := nextLiquidTag(source, j)
			if !ok2 {
				break
			}
			if name2 == "endcomment" {
				spans = append(spans, tracker.between(start, end2))
				i = end2
				closed = true
				break
			}
			j = end2
		}
		if !closed {
			return spans
		}
	}
	return spans
}

// nextLiquidTag finds the next Liquid tag at or after from.
func nextLiquidTag(source string, from int) (start, end int, name string, ok bool) {
	rel := strings.Index(source[from:], "{%")
	if rel < 0 {
		return 0, 0, "", false
	}
	start = from + rel
	closeRel := strings.Index(source[start+2:], "%}")
	if closeRel < 0 {
		return 0, 0, "", false
	}
	end = start + 2 + closeRel + 2

	inner := strings.TrimSpace(source[start+2 : start+2+closeRel])
	inner = strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(inner, "-"), "-"))
	if fields := strings.Fields(inner); len(fields) > 0 {
		name = fields[0]
	}
	return start, end, name, true
}

// complementRanges returns the source ranges not covered by masks. Masks must
// be sorted and disjoint, as handlebarsRanges returns them.
func complementRanges(source string, masks []ts.Range) []ts.Range {
	tracker := &pointTracker{source: []byte(source)}
	var out []ts.Range

	prev := 0
	for _, mask := range masks {
		start, end := int(mask.StartByte), int(mask.EndByte)
		if start > prev {
			out = append(out, tracker.between(prev, start))
		}
		if end > prev {
			prev = end
		}
	}
	if prev < len(source) {
		out = append(out, tracker.between(prev, len(source)))
	}
	return out
}

// excludeRanges subtracts masks from ranges. Both inputs must be sorted, and
// masks must be disjoint.
func excludeRanges(source string, ranges, masks []ts.Range) []ts.Range {
	if len(masks) == 0 {
		return ranges
	}

	tracker := &pointTracker{source: []byte(source)}
	var out []ts.Range
	for _, r := range ranges {
		start, end := int(r.StartByte), int(r.EndByte)
		for _, mask := range masks {
			maskStart, maskEnd := int(mask.StartByte), int(mask.EndByte)
			if maskEnd <= start || maskStart >= end {
				continue
			}
			if maskStart > start {
				out = append(out, tracker.between(start, maskStart))
			}
			if maskEnd > start {
				start = maskEnd
			}
			if start >= end {
				break
			}
		}
		if start < end {
			out = append(out, tracker.between(start, end))
		}
	}
	return out
}

// pointTracker converts byte offsets to tree-sitter points in one forward
// pass. Offsets must be queried in non-decreasing order.
type pointTracker struct {
	source []byte
	offset int
	point  ts.Point
}

func (t *pointTracker) at(offset int) ts.Point {
	if offset < t.offset {
		t.offset, t.point = 0, ts.Point{}
	}
	for t.offset < offset && t.offset < len(t.source) {
		if t.source[t.offset] == '\n' {
			t.point.Row++
			t.point.Column = 0
		} else {
			t.point.Column++
		}
		t.offset++
	}
	return t.point
}

func (t *pointTracker) between(start, end int) ts.Range {
	return ts.Range{
		StartByte:  uint(start),
		EndByte:    uint(end),
		StartPoint: t.at(start),
		EndPoint:   t.at(end),
	}
}

// jinjaDynamicRanges collects leaf nodes that represent template syntax.
// Leaf ranges avoid the enclosing statement span, which also covers the HTML
// between {% ... %} delimiters.
func jinjaDynamicRanges(root *ts.Node) []ts.Range {
	var ranges []ts.Range
	cursor := root.Walk()
	defer cursor.Close()

	if !cursor.GotoFirstChild() {
		return nil
	}
	for {
		node := cursor.Node()
		if node.ChildCount() == 0 {
			switch node.Kind() {
			case "text", "raw_text":
			default:
				ranges = append(ranges, node.Range())
			}
		}

		if cursor.GotoFirstChild() {
			continue
		}
		for !cursor.GotoNextSibling() {
			if !cursor.GotoParent() {
				return ranges
			}
		}
	}
}

// overlapsAny reports whether [start, end) intersects any range. Ranges must
// be sorted by EndByte and disjoint, which is true for template syntax leaves.
func overlapsAny(ranges []ts.Range, start, end uint) bool {
	if len(ranges) == 0 || start >= end {
		return false
	}
	i := sort.Search(len(ranges), func(i int) bool {
		return ranges[i].EndByte > start
	})
	return i < len(ranges) && ranges[i].StartByte < end
}
