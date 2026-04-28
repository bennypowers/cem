// Package css provides the tree-sitter CSS language integration.
package css

import (
	"sync"

	"bennypowers.dev/cem/internal/languages"
	ts "github.com/tree-sitter/go-tree-sitter"
	tsCss "github.com/tree-sitter/tree-sitter-css/bindings/go"
)

var lang *language

func init() {
	tsLang := ts.NewLanguage(tsCss.Language())
	lang = &language{tsLang: tsLang, pool: languages.NewParserPool(tsLang)}
	languages.Register(lang)
}

type language struct {
	tsLang *ts.Language
	pool   *sync.Pool
}

func (l *language) Name() string            { return "css" }
func (l *language) TSLanguage() *ts.Language { return l.tsLang }
func (l *language) QueryDir() string        { return "css" }

func (l *language) QueryNames(scope languages.Scope) []string {
	switch scope {
	case languages.ScopeGenerate, languages.ScopeAll:
		return []string{"cssCustomProperties"}
	}
	return nil
}

func GetParser() *ts.Parser              { return lang.pool.Get().(*ts.Parser) }
func PutParser(parser *ts.Parser)        { parser.Reset(); lang.pool.Put(parser) }
func (l *language) GetParser() *ts.Parser      { return GetParser() }
func (l *language) PutParser(parser *ts.Parser) { PutParser(parser) }
func TSLanguage() *ts.Language           { return lang.tsLang }
