package generate

import (
	"log"
	"embed"
	"fmt"
	"iter"

	ts "github.com/tree-sitter/go-tree-sitter"
	tsts "github.com/tree-sitter/tree-sitter-typescript/bindings/go"
	tsjsdoc "github.com/tree-sitter/tree-sitter-jsdoc/bindings/go"
)


var Languages = struct{Typescript *ts.Language; Jsdoc *ts.Language}{
	Typescript: ts.NewLanguage(tsts.LanguageTypescript()),
	Jsdoc: ts.NewLanguage(tsjsdoc.Language()),
}

//go:embed queries/*.scm
var queries embed.FS

type NodeInfo struct {
	Node *ts.Node;
	Text string;
}

type CaptureMap = map[string][]NodeInfo

type QueryMatcher struct {
	query *ts.Query
	cursor *ts.QueryCursor
}

func LoadQueryFile(queryName string) (string, error) {
	path := fmt.Sprintf("queries/%s.scm", queryName)
	data, err := queries.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func NewQueryMatcher(queryName string, language *ts.Language) (*QueryMatcher, func()) {
	queryText, err := LoadQueryFile(queryName)
	if err != nil {
		log.Fatal(err) // it's ok to die here because these queries are compiled in via embed
	}
	query, qerr := ts.NewQuery(language, queryText)
	if qerr != nil {
		log.Fatal(qerr) // it's ok to die here because these queries are compiled in via embed
	}
	cursor := ts.NewQueryCursor()
	thing := QueryMatcher{ query, cursor }
	return &thing, func() {
		thing.query.Close()
		thing.cursor.Close()
	}
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
	captures := make(CaptureMap);
	for _, capture := range match.Captures {
		name := names[capture.Index]
		_, ok := captures[name]
		if !ok {
			captures[name] = make([]NodeInfo, 0)
		}
		captures[name] = append(captures[name], NodeInfo{
			Node: &capture.Node,
			Text: capture.Node.Utf8Text(code),
		})
	}
	return captures
}

