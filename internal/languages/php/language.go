// Package php provides the tree-sitter PHP language integration.
package php

import (
	"embed"
	"sync"

	"bennypowers.dev/cem/internal/languages"
	ts "github.com/tree-sitter/go-tree-sitter"
	tsPHP "github.com/tree-sitter/tree-sitter-php/bindings/go"
)

var lang *language

func init() {
	tsLang := ts.NewLanguage(tsPHP.LanguagePHP())
	lang = &language{tsLang: tsLang, pool: languages.NewParserPool(tsLang)}
	languages.Register(lang)
}

type language struct {
	tsLang *ts.Language
	pool   *sync.Pool
}

func (l *language) Name() string            { return "php" }
func (l *language) TSLanguage() *ts.Language { return l.tsLang }
func (l *language) QueryFS() embed.FS      { return embed.FS{} }
func (l *language) QueryNames(_ languages.Scope) []string { return nil }

func GetParser() *ts.Parser              { return lang.pool.Get().(*ts.Parser) }
func PutParser(parser *ts.Parser)        { parser.Reset(); lang.pool.Put(parser) }
func (l *language) GetParser() *ts.Parser      { return GetParser() }
func (l *language) PutParser(parser *ts.Parser) { PutParser(parser) }
func TSLanguage() *ts.Language           { return lang.tsLang }
