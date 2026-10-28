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
	lineStart  int
	dynamic    bool
	attributes []parsedAttribute
	children   []*elementNode
}

type parsedAttribute struct {
	name        string
	value       string
	line        int
	column      int
	lineStart   int
	dynamic     bool
	nameDynamic bool
}

// parseContext carries the template ranges the parser needs to classify
// tokens. dynamic covers all template syntax; expressions covers only output
// expressions, which can compose a tag or attribute name.
type parseContext struct {
	dynamic     []ts.Range
	expressions []ts.Range
}

func formatFamily(format Format) (formatFamilyValue, error) {
	switch strings.ToLower(strings.TrimSpace(string(format))) {
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

	var (
		ctx    parseContext
		masked = []byte(source)
	)
	switch family {
	case formatJinja, formatLiquid:
		textRanges, dynamic, err := analyzeJinjaRanges(source, family == formatLiquid)
		if err != nil {
			return nil, err
		}
		// Mask template syntax instead of restricting the parser to the text
		// ranges. Included ranges concatenate text, which turns an unquoted
		// template value such as variant={{ v }} into a dangling variant= and
		// can drop the whole tag. Masking preserves byte length and line
		// breaks, so node positions still index the original source.
		ctx.dynamic = dynamic
		ctx.expressions = jinjaExpressionRanges(source)
		var comments []ts.Range
		if family == formatLiquid {
			comments = liquidCommentSpans(source)
		}
		masked = maskTemplate(source, textRanges, comments)
	case formatHandlebars:
		masks, expressions := handlebarsRanges(source)
		ctx.dynamic = masks
		ctx.expressions = expressions
		if len(masks) > 0 {
			masked = maskSpans(source, masks)
		}
	}

	tree := parser.Parse(masked, nil)
	if tree == nil {
		return nil, fmt.Errorf("markup: failed to parse markup")
	}
	defer tree.Close()

	// Node byte ranges are identical in the masked and original source, so
	// slice names and values from the original text.
	elements, _ := elementsFromNode(tree.RootNode(), []byte(source), ctx, false)
	return elements, nil
}

// elementsFromNode walks structural nodes. When skipTags is set, the node is
// an element's own tag container, so its start/end tags are not emitted as
// separate elements. ERROR nodes that wrap a start_tag are still traversed.
// The second result reports that an unclosed raw-text element consumed the
// rest of the document, so ancestor traversals must stop too.
func elementsFromNode(node *ts.Node, source []byte, ctx parseContext, skipTags bool) ([]*elementNode, bool) {
	var elements []*elementNode
	cursor := node.Walk()
	defer cursor.Close()

	for ok := cursor.GotoFirstChild(); ok; ok = cursor.GotoNextSibling() {
		child := cursor.Node()
		switch child.Kind() {
		case "element":
			if el := parseTag(child, source, ctx); el != nil {
				if !isTextContentElement(el.tagName) {
					children, stop := elementsFromNode(child, source, ctx, true)
					el.children = children
					elements = append(elements, el)
					if stop {
						return elements, true
					}
				} else {
					elements = append(elements, el)
				}
				if el.tagName == "plaintext" {
					// Everything after <plaintext> is text.
					return elements, true
				}
				if isTextContentElement(el.tagName) && !hasEndTag(child) {
					// A self-closing or unclosed raw-text element makes the
					// rest of the container text, so later markup is inert.
					return elements, true
				}
			} else {
				children, stop := elementsFromNode(child, source, ctx, false)
				elements = append(elements, children...)
				if stop {
					return elements, true
				}
			}
		case "self_closing_tag", "start_tag":
			if skipTags {
				continue
			}
			if el := parseTag(child, source, ctx); el != nil {
				elements = append(elements, el)
				if isTextContentElement(el.tagName) {
					// An unclosed or self-closing raw-text element consumes
					// the rest of its container as text.
					return elements, true
				}
			}
		case "end_tag":
			// End tags never yield elements, in either traversal mode.
		default:
			if selfClosingRawTextError(child, source) {
				// tree-sitter wraps <script/> and <style/> in an ERROR node.
				return elements, true
			}
			children, stop := elementsFromNode(child, source, ctx, false)
			elements = append(elements, children...)
			if stop {
				return elements, true
			}
		}
	}
	return elements, false
}

// selfClosingRawTextError reports whether an ERROR node is a self-closing
// script or style tag, which tree-sitter does not represent as an element.
func selfClosingRawTextError(node *ts.Node, source []byte) bool {
	if node.Kind() != "ERROR" {
		return false
	}
	tagName := findChildKind(node, "tag_name")
	if tagName == nil {
		return false
	}
	name := asciiLower(string(source[tagName.StartByte():tagName.EndByte()]))
	if !isTextContentElement(name) {
		return false
	}
	return strings.Contains(string(source[node.StartByte():node.EndByte()]), "/>")
}

// hasEndTag reports whether an element node contains its own end tag.
func hasEndTag(node *ts.Node) bool {
	cursor := node.Walk()
	defer cursor.Close()

	for ok := cursor.GotoFirstChild(); ok; ok = cursor.GotoNextSibling() {
		if cursor.Node().Kind() == "end_tag" {
			return true
		}
	}
	return false
}

func parseTag(node *ts.Node, source []byte, ctx parseContext) *elementNode {
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
		dynamic: overlapsAny(ctx.dynamic, tagName.StartByte(), tagName.EndByte()) ||
			adjacentToAny(ctx.expressions, tagName.StartByte(), tagName.EndByte()),
	}
	el.line, el.column, el.lineStart = positionFor(tagName)

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
		line, column, lineStart := positionFor(nameNode)
		valueStart, valueEnd, hasValue := attributeValueSpan(child)
		attr := parsedAttribute{
			name:      name,
			value:     attributeValue(child, source),
			line:      line,
			column:    column,
			lineStart: lineStart,
			dynamic:   hasValue && overlapsAny(ctx.dynamic, valueStart, valueEnd),
			nameDynamic: overlapsAny(ctx.dynamic, nameNode.StartByte(), nameNode.EndByte()) ||
				adjacentToAny(ctx.expressions, nameNode.StartByte(), nameNode.EndByte()),
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

// positionFor returns a one-based line, a zero-based UTF-8 byte column, and
// the byte offset where that line starts. Tree-sitter reports columns as byte
// offsets within the line; consumers convert the column to UTF-16 at the
// exported API boundary.
func positionFor(node *ts.Node) (line, column, lineStart int) {
	point := node.StartPosition()
	column = int(point.Column)
	lineStart = int(node.StartByte()) - column
	if lineStart < 0 {
		lineStart = 0
	}
	return int(point.Row) + 1, column, lineStart
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

// attributeValueSpan returns the value node's byte range. Dynamics are
// measured on the value alone, so template syntax between a name and '=' does
// not make a static value dynamic.
func attributeValueSpan(attribute *ts.Node) (start, end uint, ok bool) {
	cursor := attribute.Walk()
	defer cursor.Close()

	for ok := cursor.GotoFirstChild(); ok; ok = cursor.GotoNextSibling() {
		node := cursor.Node()
		switch node.Kind() {
		case "quoted_attribute_value", "attribute_value":
			return node.StartByte(), node.EndByte(), true
		}
	}
	return 0, 0, false
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
	case "iframe", "noembed", "noframes", "plaintext",
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
	dynamic = retainTerminatedConstructs(source, dynamic)
	sort.Slice(dynamic, func(i, j int) bool {
		return dynamic[i].EndByte < dynamic[j].EndByte
	})
	if liquid {
		text = excludeRanges(source, text, liquidCommentSpans(source))
	}
	return text, dynamic, nil
}

// retainTerminatedConstructs drops syntax leaves that belong to an
// unterminated template construct. The Jinja grammar wraps everything after a
// stray opener in one expression node, so keeping those leaves would mark all
// following markup as dynamic and silently skip validation.
func retainTerminatedConstructs(source string, ranges []ts.Range) []ts.Range {
	if len(ranges) == 0 {
		return nil
	}
	constructs := templateConstructs(source, 0, len(source))
	if len(constructs) == 0 {
		return nil
	}

	kept := ranges[:0]
	for _, r := range ranges {
		i := sort.Search(len(constructs), func(i int) bool {
			return constructs[i].EndByte > r.StartByte
		})
		if i < len(constructs) && constructs[i].StartByte <= r.StartByte && constructs[i].EndByte >= r.EndByte {
			kept = append(kept, r)
		}
	}
	return kept
}

// handlebarsRanges returns the ranges occupied by Handlebars expressions and
// comments, plus the subset that are output expressions. Masking them before
// the HTML parse keeps block helpers and helper string arguments from becoming
// attributes or elements. Unterminated expressions are left as text, and
// scanning resumes so later expressions are still masked.
func handlebarsRanges(source string) (masks, expressions []ts.Range) {
	tracker := &pointTracker{source: []byte(source)}

	add := func(start, end int, expression bool) {
		r := tracker.between(start, end)
		masks = append(masks, r)
		if expression {
			expressions = append(expressions, r)
		}
	}

	for i := 0; i+1 < len(source); {
		rel := strings.Index(source[i:], "{{")
		if rel < 0 {
			break
		}
		start := i + rel
		i = start + 2

		// An escaped mustache is literal text, not an expression.
		if isEscaped(source, start) {
			continue
		}

		// Handlebars comments contain arbitrary text, so they do not follow
		// expression quoting rules.
		if strings.HasPrefix(source[start:], "{{!--") {
			relEnd := strings.Index(source[start+5:], "--}}")
			if relEnd < 0 {
				continue
			}
			end := start + 5 + relEnd + 4
			add(start, end, false)
			i = end
			continue
		}
		if strings.HasPrefix(source[start:], "{{!") {
			relEnd := strings.Index(source[start+3:], "}}")
			if relEnd < 0 {
				continue
			}
			end := start + 3 + relEnd + 2
			add(start, end, false)
			i = end
			continue
		}

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
			continue
		}
		add(start, end, !handlebarsIsBlock(source, start, width))
		i = end
	}
	return masks, expressions
}

// handlebarsIsBlock reports whether a mustache is a block helper, partial, or
// else branch. These separate tokens instead of composing a name.
func handlebarsIsBlock(source string, start, width int) bool {
	j := start + width
	for j < len(source) && (source[j] == '~' || isHTMLSpace(source[j])) {
		j++
	}
	if j >= len(source) {
		return false
	}
	switch source[j] {
	case '#', '/', '^', '>', '!':
		return true
	}
	if strings.HasPrefix(source[j:], "else") {
		after := j + 4
		return after >= len(source) || isHTMLSpace(source[after]) || source[after] == '}' || source[after] == '~'
	}
	return false
}

// isEscaped reports whether the byte at start follows an odd number of
// backslashes, which makes a mustache literal in Handlebars.
func isEscaped(source string, start int) bool {
	backslashes := 0
	for i := start - 1; i >= 0 && source[i] == '\\'; i-- {
		backslashes++
	}
	return backslashes%2 == 1
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
		if name == "raw" {
			// Raw bodies are literal text, so a comment marker inside one
			// must not pair with a later endcomment.
			closed := false
			for j := end; j < len(source); {
				_, end2, name2, ok2 := nextLiquidTag(source, j)
				if !ok2 {
					break
				}
				if name2 == "endraw" {
					i = end2
					closed = true
					break
				}
				j = end2
			}
			if !closed {
				return spans
			}
			continue
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

// nextLiquidTag finds the next Liquid tag at or after from. Output
// expressions are skipped so a tag-looking string inside {{ ... }} does not
// start a comment block.
func nextLiquidTag(source string, from int) (start, end int, name string, ok bool) {
	for from < len(source) {
		tagRel := strings.Index(source[from:], "{%")
		if tagRel < 0 {
			return 0, 0, "", false
		}
		tagStart := from + tagRel

		// Only output expressions between from and the tag can precede it, so
		// search that gap instead of rescanning the rest of the file.
		if outRel := strings.Index(source[from:tagStart], "{{"); outRel >= 0 {
			outStart := from + outRel
			end := templateExpressionEnd(source, outStart)
			if end < 0 {
				// A stray opener is literal text; keep looking for tags.
				from = outStart + 2
				continue
			}
			from = end
			continue
		}

		start = tagStart
		closeRel := templateDelimitedEnd(source, start+2, len(source), "%}")
		if closeRel < 0 {
			// An unterminated tag is literal text; resume after its opener so
			// later tags are still recognized.
			from = start + 2
			continue
		}
		end = closeRel

		inner := strings.TrimSpace(source[start+2 : closeRel-2])
		inner = strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(inner, "-"), "-"))
		if fields := strings.Fields(inner); len(fields) > 0 {
			name = fields[0]
		}
		return start, end, name, true
	}
	return 0, 0, "", false
}

// maskTemplate masks each template construct in the gaps between keep ranges,
// leaving whitespace between constructs intact. Comment blocks are masked
// whole because their bodies must not be validated.
func maskTemplate(source string, keep, commentBlocks []ts.Range) []byte {
	var spans []ts.Range

	prev := 0
	for _, r := range keep {
		spans = append(spans, templateMasks(source, prev, int(r.StartByte), commentBlocks)...)
		if int(r.EndByte) > prev {
			prev = int(r.EndByte)
		}
	}
	spans = append(spans, templateMasks(source, prev, len(source), commentBlocks)...)
	return maskSpans(source, spans)
}

func templateMasks(source string, start, end int, commentBlocks []ts.Range) []ts.Range {
	var spans []ts.Range

	cursor := start
	for _, block := range commentBlocks {
		blockStart, blockEnd := int(block.StartByte), int(block.EndByte)
		if blockEnd <= start || blockStart >= end {
			continue
		}
		if blockStart > cursor {
			spans = append(spans, templateConstructs(source, cursor, min(blockStart, end))...)
		}
		if blockStart < start {
			blockStart = start
		}
		if blockEnd > end {
			blockEnd = end
		}
		spans = append(spans, ts.Range{StartByte: uint(blockStart), EndByte: uint(blockEnd)})
		if blockEnd > cursor {
			cursor = blockEnd
		}
	}
	if cursor < end {
		spans = append(spans, templateConstructs(source, cursor, end)...)
	}
	return spans
}

// templateConstructs returns the {{ ... }}, {% ... %}, and {# ... #} ranges in
// source[start:end]. Unterminated constructs are left unmasked.
func templateConstructs(source string, start, end int) []ts.Range {
	var spans []ts.Range

	for i := start; i < end; {
		next, opener := nextTemplateOpener(source, i, end)
		if next < 0 {
			break
		}

		var closeEnd int
		switch opener {
		case '{':
			closeEnd = templateExpressionEnd(source, next)
		case '%':
			closeEnd = templateDelimitedEnd(source, next+2, end, "%}")
		case '#':
			closeEnd = templateCommentEnd(source, next+2, end)
		}
		if closeEnd <= next {
			// Leave an unterminated opener as literal text, but keep scanning
			// so later constructs are still masked.
			i = next + 2
			continue
		}
		if closeEnd > end {
			// The construct crosses a text range boundary; its own gap masks it.
			break
		}
		spans = append(spans, ts.Range{StartByte: uint(next), EndByte: uint(closeEnd)})
		i = closeEnd
	}
	return spans
}

func nextTemplateOpener(source string, from, end int) (int, byte) {
	for i := from; i+1 < end; i++ {
		if source[i] != '{' {
			continue
		}
		switch source[i+1] {
		case '{':
			return i, '{'
		case '%':
			return i, '%'
		case '#':
			return i, '#'
		}
	}
	return -1, 0
}

// templateCommentEnd returns the end offset of a Jinja comment. Comment bodies
// are opaque text, so quotes and apostrophes have no special meaning.
func templateCommentEnd(source string, from, end int) int {
	rel := strings.Index(source[from:end], "#}")
	if rel < 0 {
		return -1
	}
	return from + rel + 2
}

func templateDelimitedEnd(source string, from, end int, closer string) int {
	var quote byte
	for i := from; i < end; {
		c := source[i]
		if quote != 0 {
			if c == '\\' {
				i += 2
				continue
			}
			if c == quote {
				quote = 0
			}
			i++
			continue
		}
		if c == '\'' || c == '"' {
			quote = c
			i++
			continue
		}
		if strings.HasPrefix(source[i:end], closer) {
			return i + len(closer)
		}
		i++
	}
	return -1
}

// maskSpans replaces the given spans with a context-aware placeholder in one
// forward pass. A space is inert between attributes and inside quoted values;
// an unquoted value or attribute-name position needs value characters so the
// mask does not leave a dangling equals sign or split the tag. Line breaks are
// kept so positions stay aligned with the original.
func maskSpans(source string, spans []ts.Range) []byte {
	masked := []byte(source)
	var state maskState
	prev := 0

	for _, span := range spans {
		start, end := int(span.StartByte), int(span.EndByte)
		state.advance(source, prev, start)
		placeholder := state.placeholder(source, end, spanIsExpression(source, start, end))
		for i := start; i < end && i < len(source); i++ {
			switch source[i] {
			case '\n', '\r':
				// Preserve line breaks.
			default:
				masked[i] = placeholder
			}
		}
		state.applyPlaceholder(placeholder)
		prev = end
	}
	state.advance(source, prev, len(source))
	return masked
}

// maskState tracks HTML tag, comment, and attribute-value context in one pass.
type maskState struct {
	inTag       bool
	inComment   bool
	quote       byte
	afterEquals bool
	unquoted    bool
}

// advance steps over source[start:end], treating HTML comment bodies as
// opaque so quotes and angle brackets inside them do not change tag state.
func (s *maskState) advance(source string, start, end int) {
	for i := start; i < end && i < len(source); {
		if s.inComment {
			if strings.HasPrefix(source[i:], "-->") {
				s.inComment = false
				i += 3
				continue
			}
			i++
			continue
		}
		if !s.inTag && strings.HasPrefix(source[i:], "<!--") {
			s.inComment = true
			i += 4
			continue
		}
		s.step(source[i])
		i++
	}
}

func (s *maskState) step(c byte) {
	if !s.inTag {
		if c == '<' {
			s.inTag = true
			s.quote = 0
			s.afterEquals = false
			s.unquoted = false
		}
		return
	}

	if s.quote != 0 {
		if c == s.quote {
			s.quote = 0
			s.afterEquals = false
			s.unquoted = false
		}
		return
	}

	switch c {
	case '>':
		s.inTag = false
		s.afterEquals = false
		s.unquoted = false
	case '=':
		s.afterEquals = true
		s.unquoted = false
	case '"', '\'':
		s.quote = c
		s.afterEquals = false
		s.unquoted = false
	case ' ', '\t', '\n', '\r', '\f':
		s.unquoted = false
	default:
		if s.afterEquals {
			s.unquoted = true
			s.afterEquals = false
		}
	}
}

// applyPlaceholder advances the state past a masked span as if its placeholder
// had been source text.
func (s *maskState) applyPlaceholder(placeholder byte) {
	if s.quote != 0 || s.inComment {
		return
	}
	if s.afterEquals || s.unquoted {
		// A statement between value branches does not end the value, so keep
		// the following static text in the same unquoted value.
		s.unquoted = true
		s.afterEquals = false
		return
	}
	s.afterEquals = false
	s.unquoted = placeholder == 'x'
}

func (s *maskState) placeholder(source string, end int, expression bool) byte {
	if s.inComment {
		return ' '
	}
	if !s.inTag || s.quote != 0 {
		return ' '
	}
	if s.unquoted {
		return 'x'
	}
	if s.afterEquals {
		if expression {
			return 'x'
		}
		// A comment or statement after '=' leaves the value position for the
		// following static value.
		return ' '
	}

	// Template syntax used as an attribute name is followed by '='; an
	// identifier keeps the rest of the tag parseable.
	if !expression {
		return ' '
	}
	j := end
	for j < len(source) && isHTMLSpace(source[j]) {
		j++
	}
	if j < len(source) && source[j] == '=' {
		return 'x'
	}
	return ' '
}

// spanIsExpression reports whether a masked span is an output expression
// rather than a statement, comment, or block helper.
func spanIsExpression(source string, start, end int) bool {
	i := start
	for i < end && isHTMLSpace(source[i]) {
		i++
	}
	if i+1 >= end || source[i] != '{' {
		return false
	}
	switch source[i+1] {
	case '{':
		j := i + 2
		for j < end && (source[j] == '~' || isHTMLSpace(source[j])) {
			j++
		}
		if j >= end {
			return true
		}
		switch source[j] {
		case '!', '#', '/', '^', '>':
			return false
		}
		if strings.HasPrefix(source[i:], "{{else") {
			after := i + 6
			if after >= end || isHTMLSpace(source[after]) || source[after] == '}' || source[after] == '~' {
				return false
			}
		}
		return true
	case '%', '#':
		return false
	}
	return false
}

func isHTMLSpace(c byte) bool {
	switch c {
	case ' ', '\t', '\n', '\r', '\f':
		return true
	}
	return false
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

// adjacentToAny reports whether [start, end) touches any range. Only output
// expressions qualify, because they can compose a tag or attribute name;
// statements and comments separate tokens rather than joining them.
func adjacentToAny(ranges []ts.Range, start, end uint) bool {
	if len(ranges) == 0 {
		return false
	}
	i := sort.Search(len(ranges), func(i int) bool {
		return ranges[i].EndByte >= start
	})
	if i < len(ranges) && ranges[i].EndByte == start {
		return true
	}
	i = sort.Search(len(ranges), func(i int) bool {
		return ranges[i].StartByte >= end
	})
	return i < len(ranges) && ranges[i].StartByte == end
}

// templateExpressionEnd returns the end offset of a {{ ... }} expression,
// honoring quoted strings. It returns -1 for an unterminated expression.
func templateExpressionEnd(source string, start int) int {
	j := start + 2
	var quote byte
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
		if strings.HasPrefix(source[j:], "}}") {
			return j + 2
		}
		j++
	}
	return -1
}

// jinjaExpressionRanges returns the ranges of {{ ... }} output expressions.
// Statements and comments are deliberately excluded.
func jinjaExpressionRanges(source string) []ts.Range {
	tracker := &pointTracker{source: []byte(source)}
	var ranges []ts.Range

	for i := 0; i+1 < len(source); {
		rel := strings.Index(source[i:], "{{")
		if rel < 0 {
			break
		}
		start := i + rel
		end := templateExpressionEnd(source, start)
		if end < 0 {
			break
		}
		ranges = append(ranges, tracker.between(start, end))
		i = end
	}
	return ranges
}
