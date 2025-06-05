package generate

import (
	"log"
	"embed"
	"fmt"
	"iter"

	ts "github.com/tree-sitter/go-tree-sitter"
	tsts "github.com/tree-sitter/tree-sitter-typescript/bindings/go"
)

var Typescript = ts.NewLanguage(tsts.LanguageTypescript())

//go:embed queries/*.scm
var queries embed.FS

type NodeInfo struct {
	Node *ts.Node;
	Text string;
}

type CaptureMap = map[string][]NodeInfo

func LoadQueryFile(queryName string) (string, error) {
	path := fmt.Sprintf("queries/%s.scm", queryName)
	data, err := queries.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func AllQueryMatches(qc *ts.QueryCursor, q *ts.Query, node *ts.Node, text []byte) iter.Seq[*ts.QueryMatch] {
	qm := qc.Matches(q, node, text)
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

func GetCapturesFromMatch(match *ts.QueryMatch, query *ts.Query, code []byte) CaptureMap {
	names := query.CaptureNames()
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

type QueryThing struct {
	query *ts.Query
	cursor *ts.QueryCursor
}

func NewQueryThing(queryName string, language *ts.Language) (*QueryThing, func()) {
	queryText, err := LoadQueryFile(queryName)
	if err != nil {
		log.Fatal(err)
	}
	query, qerr := ts.NewQuery(language, queryText)
	defer query.Close()
	if qerr != nil {
		log.Fatal(qerr)
	}
	cursor := ts.NewQueryCursor()
	defer cursor.Close()
	thing := QueryThing{ query, cursor }
	return &thing, func() {
		thing.query.Close()
		thing.cursor.Close()
	}
}

