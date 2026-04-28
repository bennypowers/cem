// Package blade provides the tree-sitter Blade language integration.
// Blade extends tree-sitter-html, so it reuses html/*.scm query files
// compiled against the Blade grammar.
package blade

import (
	"embed"
	"sync"

	"bennypowers.dev/cem/internal/languages"
	tsBlade "github.com/EmranMR/tree-sitter-blade/bindings/go"
	ts "github.com/tree-sitter/go-tree-sitter"
)

//go:embed queries/*.scm
var queryFiles embed.FS

var lang *language

func init() {
	tsLang := ts.NewLanguage(tsBlade.Language())
	lang = &language{tsLang: tsLang, pool: languages.NewParserPool(tsLang)}
	languages.Register(lang)
}

type language struct {
	tsLang *ts.Language
	pool   *sync.Pool
}

func (l *language) Name() string             { return "blade" }
func (l *language) TSLanguage() *ts.Language { return l.tsLang }
func (l *language) QueryFS() embed.FS        { return queryFiles }

func (l *language) QueryNames(scope languages.Scope) []string {
	switch scope {
	case languages.ScopeLSP, languages.ScopeAll:
		return []string{"customElements", "completionContext", "scriptTags", "headElements", "attributes"}
	}
	return nil
}

// BorrowParser borrows a Blade parser from the pool.
func BorrowParser() *ts.Parser { return lang.pool.Get().(*ts.Parser) }

// ReturnParser returns a Blade parser to the pool.
func ReturnParser(parser *ts.Parser) { parser.Reset(); lang.pool.Put(parser) }

func (l *language) BorrowParser() *ts.Parser       { return BorrowParser() }
func (l *language) ReturnParser(parser *ts.Parser) { ReturnParser(parser) }

// TSLanguage returns the tree-sitter Blade language grammar.
func TSLanguage() *ts.Language { return lang.tsLang }
