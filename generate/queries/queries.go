package queries

import (
	"embed"
	"errors"
	"fmt"
	"iter"
	"log"
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
}{
	ts.NewLanguage(tsTypescript.LanguageTypescript()),
	ts.NewLanguage(tsJsdoc.Language()),
	ts.NewLanguage(tsCss.Language()),
	ts.NewLanguage(tsHtml.Language()),
}

// ---- Parser Pooling Section ----

// HTML parser pool
var htmlParserPool = sync.Pool{
	New: func() any {
		parser := ts.NewParser()
		parser.SetLanguage(languages.html)
		return parser
	},
}

// CSS parser pool
var cssParserPool = sync.Pool{
	New: func() any {
		parser := ts.NewParser()
		parser.SetLanguage(languages.css)
		return parser
	},
}

// JSDoc parser pool
var jsdocParserPool = sync.Pool{
	New: func() any {
		parser := ts.NewParser()
		parser.SetLanguage(languages.jsdoc)
		return parser
	},
}

// TypeScript parser pool
var typescriptParserPool = sync.Pool{
	New: func() any {
		parser := ts.NewParser()
		parser.SetLanguage(languages.typescript)
		return parser
	},
}

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

// ---- End Parser Pooling Section ----

type QueryManagerI interface {
	Close()
	getQuery(name string) (*ts.Query, error)
}

type QueryManager struct {
	typescript map[string]*ts.Query
	jsdoc      map[string]*ts.Query
	css        map[string]*ts.Query
	html       map[string]*ts.Query
}

func NewQueryManager() (*QueryManager, error) {
	start := time.Now()
	qm := &QueryManager{
		typescript: make(map[string]*ts.Query),
		jsdoc:      make(map[string]*ts.Query),
		css:        make(map[string]*ts.Query),
		html:       make(map[string]*ts.Query),
	}

	data, err := queries.ReadDir(".")
	if err != nil {
		return nil, err
	}

	for _, direntry := range data {
		if direntry.IsDir() {
			language := direntry.Name()
			// embeds must always use /, never \
			// so we need to use path.Join here, never filepath.Join
			data, err := queries.ReadDir(path.Join(".", language))
			if err != nil {
				return nil, err
			}
			for _, entry := range data {
				if !entry.IsDir() {
					file := entry.Name()
					queryName := strings.Split(path.Base(file), ".")[0]
					// embeds must always use /, never \
					// so we need to use path.Join here, never filepath.Join
					data, err := queries.ReadFile(path.Join(".", language, file))
					if err != nil {
						return nil, err
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
					}
					if tsLang == nil {
						log.Fatalf("unknown language %s", language)
					}
					query, qerr := ts.NewQuery(tsLang, queryText)
					if qerr != nil {
						log.Fatal(qerr)
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
					}
				}
			}
		}
	}
	pterm.Debug.Println("Constructing queries took", time.Since(start))
	return qm, nil
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
	qm.cursor.Close()
}

func (qm QueryMatcher) GetCaptureNameByIndex(index uint32) string {
	return qm.query.CaptureNames()[index]
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
