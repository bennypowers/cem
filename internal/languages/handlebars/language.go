// Package handlebars provides the tree-sitter Handlebars language integration.
package handlebars

import (
	"sync"

	"bennypowers.dev/cem/internal/languages"
	tsHandlebars "bennypowers.dev/tree-sitter-handlebars/bindings/go"
	ts "github.com/tree-sitter/go-tree-sitter"
)

var lang *language

func init() {
	tsLang := ts.NewLanguage(tsHandlebars.Language())
	lang = &language{tsLang: tsLang, pool: languages.NewParserPool(tsLang)}
	languages.Register(lang)
}

type language struct {
	tsLang *ts.Language
	pool   *sync.Pool
}

func (l *language) Name() string            { return "handlebars" }
func (l *language) TSLanguage() *ts.Language { return l.tsLang }
func (l *language) QueryDir() string        { return "handlebars" }
func (l *language) QueryNames(_ languages.Scope) []string { return nil }

func GetParser() *ts.Parser              { return lang.pool.Get().(*ts.Parser) }
func PutParser(parser *ts.Parser)        { parser.Reset(); lang.pool.Put(parser) }
func (l *language) GetParser() *ts.Parser      { return GetParser() }
func (l *language) PutParser(parser *ts.Parser) { PutParser(parser) }
func TSLanguage() *ts.Language           { return lang.tsLang }
