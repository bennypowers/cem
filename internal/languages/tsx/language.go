// Package tsx provides the tree-sitter TSX language integration.
package tsx

import (
	"embed"
	"sync"

	"bennypowers.dev/cem/internal/languages"
	ts "github.com/tree-sitter/go-tree-sitter"
	tsTypescript "github.com/tree-sitter/tree-sitter-typescript/bindings/go"
)

//go:embed queries/*.scm
var queryFiles embed.FS

var lang *language

func init() {
	tsLang := ts.NewLanguage(tsTypescript.LanguageTSX())
	lang = &language{tsLang: tsLang, pool: languages.NewParserPool(tsLang)}
	languages.Register(lang)
}

type language struct {
	tsLang *ts.Language
	pool   *sync.Pool
}

func (l *language) Name() string             { return "tsx" }
func (l *language) TSLanguage() *ts.Language { return l.tsLang }
func (l *language) QueryFS() embed.FS        { return queryFiles }

func (l *language) QueryNames(scope languages.Scope) []string {
	switch scope {
	case languages.ScopeLSP, languages.ScopeAll:
		return []string{"customElements", "completionContext"}
	}
	return nil
}

func BorrowParser() *ts.Parser                        { return lang.pool.Get().(*ts.Parser) }
func ReturnParser(parser *ts.Parser)                  { parser.Reset(); lang.pool.Put(parser) }
func (l *language) BorrowParser() *ts.Parser       { return BorrowParser() }
func (l *language) ReturnParser(parser *ts.Parser) { ReturnParser(parser) }
func TSLanguage() *ts.Language                     { return lang.tsLang }
