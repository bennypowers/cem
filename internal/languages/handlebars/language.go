// Package handlebars provides the tree-sitter Handlebars language integration.
package handlebars

import (
	"embed"
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

func (l *language) Name() string                          { return "handlebars" }
func (l *language) TSLanguage() *ts.Language              { return l.tsLang }
func (l *language) QueryFS() embed.FS                     { return embed.FS{} }
func (l *language) QueryNames(_ languages.Scope) []string { return nil }

func BorrowParser() *ts.Parser                        { return lang.pool.Get().(*ts.Parser) }
func ReturnParser(parser *ts.Parser)                  { parser.Reset(); lang.pool.Put(parser) }
func (l *language) BorrowParser() *ts.Parser       { return BorrowParser() }
func (l *language) ReturnParser(parser *ts.Parser) { ReturnParser(parser) }
func TSLanguage() *ts.Language                     { return lang.tsLang }
