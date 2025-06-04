package generate

import (
	"log"

	ts "github.com/tree-sitter/go-tree-sitter"
	tsts "github.com/tree-sitter/tree-sitter-typescript/bindings/go"
)

type QueryThing struct {
	query *ts.Query
	cursor *ts.QueryCursor
}

var Treesitter = ts.NewLanguage(tsts.LanguageTypescript())

func NewQueryThing(queryName string, language *ts.Language) (*QueryThing, func()) {
	queryText, err := loadQueryFile(queryName)
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

func getCapturesFromMatch(match *ts.QueryMatch, query *ts.Query) map[string][]ts.Node {
	names := query.CaptureNames()
	captures := make(map[string][]ts.Node);
	for _, capture := range match.Captures {
		name := names[capture.Index]
		_, ok := captures[name]
		if !ok {
			captures[name] = make([]ts.Node, 0)
		}
		captures[name] = append(captures[name], capture.Node)
	}
	return captures
}

