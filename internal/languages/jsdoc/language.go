// Package jsdoc provides the tree-sitter JSDoc language integration.
package jsdoc

import (
	"embed"
	"sync"

	"bennypowers.dev/cem/internal/languages"
	ts "github.com/tree-sitter/go-tree-sitter"
	tsJsdoc "github.com/tree-sitter/tree-sitter-jsdoc/bindings/go"
)

//go:embed queries/*.scm
var queryFiles embed.FS

var lang *language

func init() {
	tsLang := ts.NewLanguage(tsJsdoc.Language())
	lang = &language{tsLang: tsLang, pool: languages.NewParserPool(tsLang)}
	languages.Register(lang)
}

type language struct {
	tsLang *ts.Language
	pool   *sync.Pool
}

func (l *language) Name() string             { return "jsdoc" }
func (l *language) TSLanguage() *ts.Language { return l.tsLang }
func (l *language) QueryFS() embed.FS        { return queryFiles }

func (l *language) QueryNames(scope languages.Scope) []string {
	switch scope {
	case languages.ScopeGenerate, languages.ScopeAll:
		return []string{"jsdoc"}
	}
	return nil
}

func BorrowParser() *ts.Parser                        { return lang.pool.Get().(*ts.Parser) }
func ReturnParser(parser *ts.Parser)                  { parser.Reset(); lang.pool.Put(parser) }
func (l *language) BorrowParser() *ts.Parser       { return BorrowParser() }
func (l *language) ReturnParser(parser *ts.Parser) { ReturnParser(parser) }
func TSLanguage() *ts.Language                     { return lang.tsLang }
