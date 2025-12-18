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
	"errors"
	"fmt"
	"iter"
	"path"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/pterm/pterm"

	ts "github.com/tree-sitter/go-tree-sitter"
	tsCss "github.com/tree-sitter/tree-sitter-css/bindings/go"
	tsHtml "github.com/tree-sitter/tree-sitter-html/bindings/go"
	tsJsdoc "github.com/tree-sitter/tree-sitter-jsdoc/bindings/go"
	tsTypescript "github.com/tree-sitter/tree-sitter-typescript/bindings/go"
)

//go:embed */*.scm
var queries embed.FS

var ErrNoQueryManager = errors.New("QueryManager is nil")

type NoCaptureError struct {
	Capture string
	Query   string
}

func (e *NoCaptureError) Error() string {
	return fmt.Sprintf("No nodes for capture %s in query %s", e.Capture, e.Query)
}

// ---- Languages struct holds pre-initialized language grammars ----
var languages = struct {
	typescript *ts.Language
	jsdoc      *ts.Language
	css        *ts.Language
	html       *ts.Language
	tsx        *ts.Language
}{
	ts.NewLanguage(tsTypescript.LanguageTypescript()),
	ts.NewLanguage(tsJsdoc.Language()),
	ts.NewLanguage(tsCss.Language()),
	ts.NewLanguage(tsHtml.Language()),
	ts.NewLanguage(tsTypescript.LanguageTSX()), // TSX uses TypeScript's TSX dialect
}

// ---- Parser Pooling Section ----

// HTML parser pool
var htmlParserPool = sync.Pool{
	New: func() any {
		parser := ts.NewParser()
		if err := parser.SetLanguage(languages.html); err != nil {
			panic(fmt.Sprintf("failed to set HTML language: %v", err))
		}
		return parser
	},
}

// CSS parser pool
var cssParserPool = sync.Pool{
	New: func() any {
		parser := ts.NewParser()
		if err := parser.SetLanguage(languages.css); err != nil {
			panic(fmt.Sprintf("failed to set CSS language: %v", err))
		}
		return parser
	},
}

// JSDoc parser pool
var jsdocParserPool = sync.Pool{
	New: func() any {
		parser := ts.NewParser()
		if err := parser.SetLanguage(languages.jsdoc); err != nil {
			panic(fmt.Sprintf("failed to set JSDoc language: %v", err))
		}
		return parser
	},
}

// TypeScript parser pool
var typescriptParserPool = sync.Pool{
	New: func() any {
		parser := ts.NewParser()
		if err := parser.SetLanguage(languages.typescript); err != nil {
			panic(fmt.Sprintf("failed to set TypeScript language: %v", err))
		}
		return parser
	},
}

// TSX parser pool
var tsxParserPool = sync.Pool{
	New: func() any {
		parser := ts.NewParser()
		if err := parser.SetLanguage(languages.tsx); err != nil {
			panic(fmt.Sprintf("failed to set TSX language: %v", err))
		}
		return parser
	},
}

// Note: Removed cursor pooling due to stateful nature of QueryCursor objects.
// Cursors maintain internal state that can affect subsequent queries,
// leading to unpredictable behavior. Always create fresh cursors.

// GetHTMLParser returns a pooled HTML parser.
// Always call PutHTMLParser when done.
func GetHTMLParser() *ts.Parser {
	return htmlParserPool.Get().(*ts.Parser)
}

// PutHTMLParser returns a parser to the HTML pool.
func PutHTMLParser(parser *ts.Parser) {
	parser.Reset()
	htmlParserPool.Put(parser)
}

// GetCSSParser returns a pooled CSS parser.
// Always call PutCSSParser when done.
func GetCSSParser() *ts.Parser {
	return cssParserPool.Get().(*ts.Parser)
}

// PutCSSParser returns a parser to the CSS pool.
func PutCSSParser(parser *ts.Parser) {
	parser.Reset()
	cssParserPool.Put(parser)
}

// GetJSDocParser returns a pooled JSDoc parser.
// Always call PutJSDocParser when done.
func GetJSDocParser() *ts.Parser {
	return jsdocParserPool.Get().(*ts.Parser)
}

// PutJSDocParser returns a parser to the JSDoc pool.
func PutJSDocParser(parser *ts.Parser) {
	parser.Reset()
	jsdocParserPool.Put(parser)
}

// RetrieveTypeScriptParser returns a singleton TypeScript parser.
// Do not Close() this parser.
func RetrieveTypeScriptParser() *ts.Parser {
	return typescriptParserPool.Get().(*ts.Parser)
}

// PutTypeScriptParser returns a parser to the TypeScript pool.
func PutTypeScriptParser(parser *ts.Parser) {
	parser.Reset()
	typescriptParserPool.Put(parser)
}

// GetTSXParser returns a pooled TSX parser.
// Always call PutTSXParser when done.
func GetTSXParser() *ts.Parser {
	return tsxParserPool.Get().(*ts.Parser)
}

// PutTSXParser returns a parser to the TSX pool.
func PutTSXParser(parser *ts.Parser) {
	parser.Reset()
	tsxParserPool.Put(parser)
}

// ---- End Parser Pooling Section ----

// QuerySelector defines which queries to load for performance
type QuerySelector struct {
	HTML       []string // HTML query names to load
	TypeScript []string // TypeScript query names to load
	CSS        []string // CSS query names to load
	JSDoc      []string // JSDoc query names to load
	TSX        []string // TSX query names to load
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
		TypeScript: []string{"htmlTemplates", "completionContext", "imports", "classes", "exports", "importAttributes"},
		CSS:        []string{},
		JSDoc:      []string{},
		TSX:        []string{"customElements", "completionContext"}, // TSX queries for custom element detection
	}
}

// GenerateQueries returns a selector that loads only queries needed for generate
func GenerateQueries() QuerySelector {
	return QuerySelector{
		HTML:       []string{"slotsAndParts"},
		TypeScript: []string{"classMemberDeclaration", "classes", "declarations", "imports", "typeAliases"},
		CSS:        []string{"cssCustomProperties"},
		JSDoc:      []string{"jsdoc"},
	}
}

type QueryManagerI interface {
	Close()
	getQuery(name string) (*ts.Query, error)
}

type QueryManager struct {
	typescript map[string]*ts.Query
	jsdoc      map[string]*ts.Query
	css        map[string]*ts.Query
	html       map[string]*ts.Query
	tsx        map[string]*ts.Query
}

func NewQueryManager(selector QuerySelector) (*QueryManager, error) {
	start := time.Now()
	qm := &QueryManager{
		typescript: make(map[string]*ts.Query),
		jsdoc:      make(map[string]*ts.Query),
		css:        make(map[string]*ts.Query),
		html:       make(map[string]*ts.Query),
		tsx:        make(map[string]*ts.Query),
	}

	// Load only the requested queries
	for _, queryName := range selector.HTML {
		if err := qm.loadQuery("html", queryName); err != nil {
			qm.Close()
			return nil, fmt.Errorf("failed to load HTML query %s: %w", queryName, err)
		}
	}

	for _, queryName := range selector.TypeScript {
		if err := qm.loadQuery("typescript", queryName); err != nil {
			qm.Close()
			return nil, fmt.Errorf("failed to load TypeScript query %s: %w", queryName, err)
		}
	}

	for _, queryName := range selector.CSS {
		if err := qm.loadQuery("css", queryName); err != nil {
			qm.Close()
			return nil, fmt.Errorf("failed to load CSS query %s: %w", queryName, err)
		}
	}

	for _, queryName := range selector.JSDoc {
		if err := qm.loadQuery("jsdoc", queryName); err != nil {
			qm.Close()
			return nil, fmt.Errorf("failed to load JSDoc query %s: %w", queryName, err)
		}
	}

	for _, queryName := range selector.TSX {
		if err := qm.loadQuery("tsx", queryName); err != nil {
			qm.Close()
			return nil, fmt.Errorf("failed to load TSX query %s: %w", queryName, err)
		}
	}

	pterm.Debug.Println("Constructing selected queries took", time.Since(start))
	return qm, nil
}

func (qm *QueryManager) loadQuery(language, queryName string) error {
	// Read the query file
	// Use path.Join (not filepath.Join) - embed.FS requires POSIX / separators
	queryPath := path.Join(language, queryName+".scm")
	data, err := queries.ReadFile(queryPath)
	if err != nil {
		return fmt.Errorf("failed to read query file %s: %w", queryPath, err)
	}

	queryText := string(data)
	var tsLang *ts.Language
	switch language {
	case "typescript":
		tsLang = languages.typescript
	case "jsdoc":
		tsLang = languages.jsdoc
	case "css":
		tsLang = languages.css
	case "html":
		tsLang = languages.html
	case "tsx":
		tsLang = languages.tsx
	default:
		return fmt.Errorf("unknown language %s", language)
	}

	query, qerr := ts.NewQuery(tsLang, queryText)
	if qerr != nil {
		return fmt.Errorf("failed to parse query %s: %w", queryName, qerr)
	}

	switch language {
	case "typescript":
		qm.typescript[queryName] = query
	case "jsdoc":
		qm.jsdoc[queryName] = query
	case "css":
		qm.css[queryName] = query
	case "html":
		qm.html[queryName] = query
	case "tsx":
		qm.tsx[queryName] = query
	}

	return nil
}

func (qm *QueryManager) Close() {
	for _, query := range qm.typescript {
		query.Close()
	}
	for _, query := range qm.jsdoc {
		query.Close()
	}
	for _, query := range qm.css {
		query.Close()
	}
	for _, query := range qm.html {
		query.Close()
	}
	for _, query := range qm.tsx {
		query.Close()
	}
}

func (qm *QueryManager) getQuery(queryName string, language string) (*ts.Query, error) {
	var q *ts.Query
	var ok bool
	switch language {
	case "typescript":
		q, ok = qm.typescript[queryName]
	case "jsdoc":
		q, ok = qm.jsdoc[queryName]
	case "css":
		q, ok = qm.css[queryName]
	case "html":
		q, ok = qm.html[queryName]
	case "tsx":
		q, ok = qm.tsx[queryName]
	}
	if !ok {
		return nil, fmt.Errorf("unknown query %s", queryName)
	}
	return q, nil
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

type QueryMatcher struct {
	query  *ts.Query
	cursor *ts.QueryCursor
}

func (qm QueryMatcher) Close() {
	// NOTE: we don't close queries here, only at the end of execution in QueryManager.Close
	// Close the cursor since we're not pooling
	qm.cursor.Close()
}

func (qm QueryMatcher) GetCaptureNameByIndex(index uint32) string {
	return qm.query.CaptureNames()[index]
}

func (qm QueryMatcher) CaptureCount() int {
	return len(qm.query.CaptureNames())
}

func (qm QueryMatcher) GetCaptureIndexForName(name string) (uint, bool) {
	return qm.query.CaptureIndexForName(name)
}

func (qm QueryMatcher) SetByteRange(start uint, end uint) {
	qm.cursor.SetByteRange(start, end)
}

func NewQueryMatcher(
	manager *QueryManager,
	language string,
	queryName string,
) (*QueryMatcher, error) {
	if manager == nil {
		return nil, ErrNoQueryManager
	}
	query, err := manager.getQuery(queryName, language)
	if err != nil {
		return nil, err
	}
	cursor := ts.NewQueryCursor()
	qm := QueryMatcher{query, cursor}
	return &qm, nil
}

func (q QueryMatcher) AllQueryMatches(node *ts.Node, text []byte) iter.Seq[*ts.QueryMatch] {
	qm := q.cursor.Matches(q.query, node, text)
	return func(yield func(qm *ts.QueryMatch) bool) {
		for {
			m := qm.Next()
			if m == nil {
				break
			}
			if !yield(m) {
				return
			}
		}
	}
}

// ParentCaptures returns an iterator over unique parent node captures as identified by the given parent capture name.
// For each unique parent node (e.g., a field or method), it aggregates all captures from all query matches sharing
// that parent node into a single CaptureMap. This allows you to collect all related captures (such as attributes,
// decorators, etc.) for each parent node in the source AST.
//
// Example usage:
//
//	for captures := range matcher.ParentCaptures(root, code, "field") {
//	  // captures represents the captured nodes for a single field
//	  addFieldToDeclaration(captures, declaration)
//	}
func (q *QueryMatcher) ParentCaptures(root *ts.Node, code []byte, parentCaptureName string) iter.Seq[CaptureMap] {
	names := q.query.CaptureNames()

	type pgroup struct {
		capMap    CaptureMap
		startByte uint
	}

	parentGroups := make(map[int]pgroup)

	for match := range q.AllQueryMatches(root, code) {
		var parentNode *ts.Node
		for _, cap := range match.Captures {
			name := names[cap.Index]
			if name == parentCaptureName {
				parentNode = &cap.Node
				break
			}
		}
		if parentNode == nil {
			continue
		}
		pid := int(parentNode.Id())
		startByte := parentNode.StartByte()
		_, ok := parentGroups[pid]
		if !ok {
			capmap := make(CaptureMap)
			parentGroups[pid] = pgroup{capmap, startByte}
		}
		for _, cap := range match.Captures {
			name := names[cap.Index]
			text := cap.Node.Utf8Text(code)
			ci := CaptureInfo{
				NodeId:    int(cap.Node.Id()),
				Text:      text,
				StartByte: cap.Node.StartByte(),
				EndByte:   cap.Node.EndByte(),
			}
			if _, hasMap := parentGroups[pid].capMap[name]; !hasMap {
				parentGroups[pid].capMap[name] = make([]CaptureInfo, 0)
			}
			if !slices.ContainsFunc(parentGroups[pid].capMap[name], func(m CaptureInfo) bool {
				return m.NodeId == ci.NodeId
			}) {
				parentGroups[pid].capMap[name] = append(parentGroups[pid].capMap[name], ci)
			}
		}
	}

	sorted := make([]pgroup, 0)
	for _, group := range parentGroups {
		sorted = append(sorted, group)
	}

	slices.SortStableFunc(sorted, func(a pgroup, b pgroup) int {
		return int(a.startByte) - int(b.startByte)
	})

	return func(yield func(CaptureMap) bool) {
		for _, group := range sorted {
			if !yield(group.capMap) {
				break
			}
		}
	}
}

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

// NodeToRange converts a tree-sitter node to a Range using byte-to-position conversion
func NodeToRange(node *ts.Node, content []byte) Range {
	start := byteOffsetToPosition(content, node.StartByte())
	end := byteOffsetToPosition(content, node.EndByte())
	return Range{
		Start: start,
		End:   end,
	}
}

// FindClassDeclarationInSource finds the class declaration position in TypeScript source content
func FindClassDeclarationInSource(content []byte, className string, queryManager *QueryManager) (*Range, error) {
	// Get TypeScript parser and parse
	parser := RetrieveTypeScriptParser()
	defer PutTypeScriptParser(parser)

	tree := parser.Parse(content, nil)
	if tree == nil {
		return nil, nil
	}
	defer tree.Close()

	// Get class queries
	classQueries, err := NewQueryMatcher(queryManager, "typescript", "classes")
	if err != nil {
		return nil, err
	}
	defer classQueries.Close()

	// Find class declaration using ParentCaptures pattern
	for captureMap := range classQueries.ParentCaptures(tree.RootNode(), content, "class") {
		if classNames, ok := captureMap["class.name"]; ok {
			for _, classNameCapture := range classNames {
				if classNameCapture.Text == className {
					node := GetDescendantById(tree.RootNode(), classNameCapture.NodeId)
					if node != nil {
						r := NodeToRange(node, content)
						return &r, nil
					}
				}
			}
		}
	}

	return nil, nil
}

// FindTagNameDefinitionInSource finds where the tag name is defined (e.g., @customElement decorator)
func FindTagNameDefinitionInSource(content []byte, tagName string, queryManager *QueryManager) (*Range, error) {
	// Get TypeScript parser and parse
	parser := RetrieveTypeScriptParser()
	defer PutTypeScriptParser(parser)

	tree := parser.Parse(content, nil)
	if tree == nil {
		return nil, nil
	}
	defer tree.Close()

	// Get class queries
	classQueries, err := NewQueryMatcher(queryManager, "typescript", "classes")
	if err != nil {
		return nil, err
	}
	defer classQueries.Close()

	// Find tag name in decorator using ParentCaptures pattern
	for captureMap := range classQueries.ParentCaptures(tree.RootNode(), content, "class") {
		if tagNames, ok := captureMap["tag-name"]; ok {
			for _, tagNameCapture := range tagNames {
				if tagNameCapture.Text == tagName {
					node := GetDescendantById(tree.RootNode(), tagNameCapture.NodeId)
					if node != nil {
						r := NodeToRange(node, content)
						return &r, nil
					}
				}
			}
		}
	}

	return nil, nil
}

// FindAttributeDeclarationInSource finds the attribute/property declaration in source content
func FindAttributeDeclarationInSource(content []byte, attributeName string, queryManager *QueryManager) (*Range, error) {
	// Get TypeScript parser and parse
	parser := RetrieveTypeScriptParser()
	defer PutTypeScriptParser(parser)

	tree := parser.Parse(content, nil)
	if tree == nil {
		return nil, nil
	}
	defer tree.Close()

	// Get member queries
	memberQueries, err := NewQueryMatcher(queryManager, "typescript", "classMemberDeclaration")
	if err != nil {
		return nil, err
	}
	defer memberQueries.Close()

	// Find attribute declaration using ParentCaptures pattern
	for captureMap := range memberQueries.ParentCaptures(tree.RootNode(), content, "field") {
		var memberName string
		var attributeNameInDecorator string
		var memberNodeId int

		// Collect information from this field
		if memberNames, ok := captureMap["member.name"]; ok && len(memberNames) > 0 {
			memberName = memberNames[0].Text
			memberNodeId = memberNames[0].NodeId
		}
		if attrNames, ok := captureMap["field.attr.name"]; ok && len(attrNames) > 0 {
			attributeNameInDecorator = attrNames[0].Text
		}

		// Check if this member matches our target attribute
		if matchesAttribute(memberName, attributeNameInDecorator, attributeName) {
			node := GetDescendantById(tree.RootNode(), memberNodeId)
			if node != nil {
				r := NodeToRange(node, content)
				return &r, nil
			}
		}
	}

	return nil, nil
}

// FindSlotDefinitionInSource finds the slot definition in template within source content
func FindSlotDefinitionInSource(content []byte, slotName string, queryManager *QueryManager) (*Range, error) {
	// Get TypeScript parser and parse
	parser := RetrieveTypeScriptParser()
	defer PutTypeScriptParser(parser)

	tree := parser.Parse(content, nil)
	if tree == nil {
		return nil, nil
	}
	defer tree.Close()

	// Get class queries to find render templates
	classQueries, err := NewQueryMatcher(queryManager, "typescript", "classes")
	if err != nil {
		return nil, err
	}
	defer classQueries.Close()

	// Find HTML template strings in render methods
	for captureMap := range classQueries.ParentCaptures(tree.RootNode(), content, "class") {
		if templates, ok := captureMap["render.template"]; ok {
			for _, template := range templates {
				templateContent := template.Text
				// Remove template literal backticks
				if len(templateContent) >= 2 && templateContent[0] == '`' && templateContent[len(templateContent)-1] == '`' {
					templateContent = templateContent[1 : len(templateContent)-1]
				}

				slotRange := findSlotInTemplate(templateContent, slotName, queryManager)
				if slotRange != nil {
					// Adjust range to account for position within the TypeScript file
					templateNode := GetDescendantById(tree.RootNode(), template.NodeId)
					if templateNode != nil {
						adjustedRange := adjustTemplateRange(slotRange, templateNode, content)
						return adjustedRange, nil
					}
				}
			}
		}
	}

	return nil, nil
}

// Helper functions

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

// Thread-safe singleton QueryManager (there can be only one!)
var (
	globalQueryManager *QueryManager
	globalQueryOnce    sync.Once
	globalQueryError   error
)

// GetGlobalQueryManager returns the singleton QueryManager instance
func GetGlobalQueryManager() (*QueryManager, error) {
	globalQueryOnce.Do(func() {
		manager, err := NewQueryManager(LSPQueries())
		if err != nil {
			globalQueryError = err
			return
		}
		globalQueryManager = manager
	})

	if globalQueryError != nil {
		return nil, globalQueryError
	}

	if globalQueryManager == nil {
		return nil, fmt.Errorf("failed to initialize global query manager")
	}

	return globalQueryManager, nil
}

// GetCachedQueryMatcher returns a query matcher with cached query but fresh cursor
// This prevents concurrent access to the same cursor which causes segmentation faults
func GetCachedQueryMatcher(manager *QueryManager, language, queryName string) (*QueryMatcher, error) {
	if manager == nil {
		return nil, ErrNoQueryManager
	}

	// Get the cached query (thread-safe to share)
	query, err := manager.getQuery(queryName, language)
	if err != nil {
		return nil, err
	}

	// Always create a fresh cursor (NOT thread-safe to share)
	cursor := ts.NewQueryCursor()
	matcher := QueryMatcher{query, cursor}
	return &matcher, nil
}
