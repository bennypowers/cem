package generate

import (
	"embed"
	"errors"
	"fmt"
	"path/filepath"
	"iter"
	"log"
	"slices"
	"strings"

	ts "github.com/tree-sitter/go-tree-sitter"
	tsjsdoc "github.com/tree-sitter/tree-sitter-jsdoc/bindings/go"
	tsts "github.com/tree-sitter/tree-sitter-typescript/bindings/go"
)

type NoCaptureError struct {
	CaptureName string
	QueryName string
}

func (e *NoCaptureError) Error() string {
  return fmt.Sprintf("No nodes for capture %s in query %s", e.CaptureName, e.QueryName)
}

var Languages = struct{Typescript *ts.Language; Jsdoc *ts.Language}{
	Typescript: ts.NewLanguage(tsts.LanguageTypescript()),
	Jsdoc: ts.NewLanguage(tsjsdoc.Language()),
}

//go:embed queries/*.scm
var queries embed.FS

type ParentNodeCaptures struct {
	NodeId uintptr;
	Captures CaptureMap
}

type CaptureInfo struct {
	NodeId int;
	Text string;
}

type CaptureMap = map[string][]CaptureInfo

type QueryMatcher struct {
	query *ts.Query
	cursor *ts.QueryCursor
}

func (qm QueryMatcher) Close() {
	qm.query.Close()
	qm.cursor.Close()
}

var queryMap map[string]string

func preloadQueries() {
	if len(queryMap) > 0 {
		return
	}
  queryMap = make(map[string]string)
	data, err := queries.ReadDir("queries")
	if err != nil {
		log.Fatal(err) // it's ok to die here because these queries are compiled in via embed
	}
	for _, entry := range data {
		if !entry.IsDir() {
			file := entry.Name()
			queryName := strings.Split(filepath.Base(file), ".")[0]
			data, err := queries.ReadFile(filepath.Join("queries", file))
			if err != nil {
				log.Fatal(err) // it's ok to die here because these queries are compiled in via embed
			}
			contents := string(data)
			queryMap[queryName] = contents
		}
	}
}

func NewQueryMatcher(queryName string, language *ts.Language) (*QueryMatcher, error) {
	preloadQueries()
	queryText, ok := queryMap[queryName]
	if !ok {
		return nil, errors.New(fmt.Sprintf("unknown query %s", queryName))
	}
	query, qerr := ts.NewQuery(language, queryText)
	if qerr != nil {
		log.Fatal(qerr) // it's ok to die here because these queries are compiled in via embed
	}
	cursor := ts.NewQueryCursor()
	thing := QueryMatcher{ query, cursor }
	return &thing, nil
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

func (q QueryMatcher) GetCapturesFromMatch(match *ts.QueryMatch, code []byte) CaptureMap {
	names := q.query.CaptureNames()
	captures := make(CaptureMap)
	for _, capture := range match.Captures {
		name := names[capture.Index]
		_, ok := captures[name]
		if !ok {
			captures[name] = make([]CaptureInfo, 0)
		}
		text := capture.Node.Utf8Text(code)
		captures[name] = append(captures[name], CaptureInfo{
			NodeId: int(capture.Node.Id()),
			Text: text,
		})
	}
	return captures
}

// ParentIterator returns an iterator over unique parent node captures as identified by the given parent capture name.
// For each unique parent node (e.g., a field or method), it aggregates all captures from all query matches sharing 
// that parent node into a single CaptureMap. This allows you to collect all related captures (such as attributes, 
// decorators, etc.) for each parent node in the source AST.
//
// Example usage:
//
//   for captures := range matcher.ParentIterator(root, code, "field") {
//     // captures represents the captured nodes for a single field
//     addFieldToDeclaration(captures, declaration)
//   }
func (q *QueryMatcher) ParentCaptures(root *ts.Node, code []byte, parentCaptureName string) iter.Seq[CaptureMap] {
	names := q.query.CaptureNames()

	type pgroup struct { capMap CaptureMap; startByte int }

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
		startByte := int(parentNode.StartByte())
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
		return a.startByte - b.startByte
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
