// Package html provides the tree-sitter HTML language integration.
package html

import (
	"embed"
	"sync"

	"bennypowers.dev/cem/internal/languages"
	ts "github.com/tree-sitter/go-tree-sitter"
	tsHtml "github.com/tree-sitter/tree-sitter-html/bindings/go"
)

//go:embed queries/*.scm
var queryFiles embed.FS

var lang *language

func init() {
	tsLang := ts.NewLanguage(tsHtml.Language())
	lang = &language{
		tsLang: tsLang,
		pool:   languages.NewParserPool(tsLang),
	}
	languages.Register(lang)
}

type language struct {
	tsLang *ts.Language
	pool   *sync.Pool
}

func (l *language) Name() string            { return "html" }
func (l *language) TSLanguage() *ts.Language { return l.tsLang }
func (l *language) QueryFS() embed.FS      { return queryFiles }

func (l *language) QueryNames(scope languages.Scope) []string {
	switch scope {
	case languages.ScopeLSP:
		return []string{"customElements", "completionContext", "scriptTags", "headElements", "attributes"}
	case languages.ScopeGenerate:
		return []string{"slotsAndParts"}
	case languages.ScopeAll:
		return []string{"slotsAndParts", "customElements", "completionContext", "scriptTags", "headElements", "attributes"}
	}
	return nil
}

// GetParser borrows an HTML parser from the pool.
func GetParser() *ts.Parser { return lang.pool.Get().(*ts.Parser) }

// PutParser returns an HTML parser to the pool.
func PutParser(parser *ts.Parser) { parser.Reset(); lang.pool.Put(parser) }

func (l *language) GetParser() *ts.Parser      { return GetParser() }
func (l *language) PutParser(parser *ts.Parser) { PutParser(parser) }

// TSLanguage returns the tree-sitter HTML language grammar.
func TSLanguage() *ts.Language { return lang.tsLang }
