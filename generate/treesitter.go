package generate

import (
	"embed"
	"errors"
	"fmt"
	"iter"
	"log"
	"path"
	"path/filepath"
	"slices"
	"strings"

	ts "github.com/tree-sitter/go-tree-sitter"
	tsCss "github.com/tree-sitter/tree-sitter-css/bindings/go"
	tsHtml "github.com/tree-sitter/tree-sitter-html/bindings/go"
	tsJsdoc "github.com/tree-sitter/tree-sitter-jsdoc/bindings/go"
	tsTypescript "github.com/tree-sitter/tree-sitter-typescript/bindings/go"
)

//go:embed queries/*/*.scm
var queries embed.FS

type NoCaptureError struct {
	CaptureName string
	QueryName string
}

func (e *NoCaptureError) Error() string {
  return fmt.Sprintf("No nodes for capture %s in query %s", e.CaptureName, e.QueryName)
}

var Languages = struct{
	typescript *ts.Language
	jsdoc *ts.Language
	css *ts.Language
	html *ts.Language
}{
	ts.NewLanguage(tsTypescript.LanguageTypescript()),
	ts.NewLanguage(tsJsdoc.Language()),
	ts.NewLanguage(tsCss.Language()),
	ts.NewLanguage(tsHtml.Language()),
}

type QueryManagerI interface {
	Close()
	GetQuery(name string) (*ts.Query, error)
}

type QueryManager struct {
	typescript map[string]*ts.Query
	jsdoc map[string]*ts.Query
	css map[string]*ts.Query
	html map[string]*ts.Query
}

func NewQueryManager() (*QueryManager, error) {
	qm := &QueryManager{
		typescript: make(map[string]*ts.Query),
		jsdoc: make(map[string]*ts.Query),
		css: make(map[string]*ts.Query),
		html: make(map[string]*ts.Query),
	}

	data, err := queries.ReadDir("queries")
	if err != nil {
		log.Fatal(err) // it's ok to die here because these queries are compiled in via embed
	}

	for _, direntry := range data {
		if direntry.IsDir() {
			language := direntry.Name()
			data, err := queries.ReadDir(path.Join("queries", language))
			if err != nil {
				log.Fatal(err) // it's ok to die here because these queries are compiled in via embed
			}
			for _, entry := range data {
				// todo: jsdoc/*.scm, typescript/*.scm, etc
				if !entry.IsDir() {
					file := entry.Name()
					queryName := strings.Split(filepath.Base(file), ".")[0]
					data, err := queries.ReadFile(filepath.Join("queries", language, file))
					if err != nil {
						log.Fatal(err) // it's ok to die here because these queries are compiled in via embed
					}
					queryText := string(data)
					var tsLang *ts.Language
					switch language {
					case "typescript":
						tsLang = Languages.typescript
					case "jsdoc":
						tsLang = Languages.jsdoc
					case "css":
						tsLang = Languages.css
					case "html":
						tsLang = Languages.html
					}
					if tsLang == nil {
						// it's ok to die here because these queries are compiled in via embed
						log.Fatalf("unknown language %s", language)
					}
					query, qerr := ts.NewQuery(tsLang, queryText)
					if qerr != nil {
						log.Fatal(qerr) // it's ok to die here because these queries are compiled in via embed
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

func (qm *QueryManager) GetQuery(queryName string, language string) (*ts.Query, error) {
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
		return nil, errors.New(fmt.Sprintf("unknown query %s", queryName))
	}
	return q, nil
}

type ParentNodeCaptures struct {
	NodeId uintptr;
	Captures CaptureMap
}

type CaptureInfo struct {
	NodeId int;
	Text string;
	StartByte uint;
	EndByte uint;
}

type CaptureMap = map[string][]CaptureInfo

type QueryMatcher struct {
	query *ts.Query
	cursor *ts.QueryCursor
}

func (qm QueryMatcher) Close() {
	qm.cursor.Close()
}

func NewQueryMatcher(
	manager *QueryManager,
	language string,
	queryName string,
) (*QueryMatcher, error) {
	query, error := manager.GetQuery(queryName, language)
	if error != nil {
		return nil, error
	}
	cursor := ts.NewQueryCursor()
	qm := QueryMatcher{ query, cursor }
	return &qm, nil
}

func (q QueryMatcher) AllQueryMatches(node *ts.Node, text []byte) iter.Seq[*ts.QueryMatch] {
	qm := q.cursor.Matches(q.query, node, text)
	return func(yield func (qm *ts.QueryMatch) bool) {
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
//   for captures := range matcher.ParentCaptures(root, code, "field") {
//     // captures represents the captured nodes for a single field
//     addFieldToDeclaration(captures, declaration)
//   }
func (q *QueryMatcher) ParentCaptures(root *ts.Node, code []byte, parentCaptureName string) iter.Seq[CaptureMap] {
	names := q.query.CaptureNames()

	type pgroup struct { capMap CaptureMap; startByte uint }

	// Group matches by parent node id
	parentGroups := make(map[int]pgroup)

	// Collect all matches
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
			continue // skip matches without parent
		}
		pid := int(parentNode.Id())
		startByte := parentNode.StartByte()
		_, ok := parentGroups[pid]
		if !ok {
			capmap := make(CaptureMap)
			parentGroups[pid] = pgroup{capmap, startByte}
		}
		// Collect all captures for this match
		for _, cap := range match.Captures {
			name := names[cap.Index]
			text := cap.Node.Utf8Text(code)
			ci := CaptureInfo{
				NodeId: int(cap.Node.Id()),
				Text:   text,
				StartByte: cap.Node.StartByte(),
				EndByte: cap.Node.EndByte(),
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

	// Iterator over aggregated groups
	return func(yield func(CaptureMap) bool) {
		for _, group := range sorted {
			if !yield(group.capMap) {
				break
			}
		}
	}
}

func GetDescendantById(root *ts.Node, id int) (*ts.Node) {
	c := root.Walk()
	defer c.Close()
	// get the descendant with the id
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
