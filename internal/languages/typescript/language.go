// Package typescript provides the tree-sitter TypeScript language integration.
package typescript

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
	tsLang := ts.NewLanguage(tsTypescript.LanguageTypescript())
	lang = &language{tsLang: tsLang, pool: languages.NewParserPool(tsLang)}
	languages.Register(lang)
}

type language struct {
	tsLang *ts.Language
	pool   *sync.Pool
}

func (l *language) Name() string            { return "typescript" }
func (l *language) TSLanguage() *ts.Language { return l.tsLang }
func (l *language) QueryFS() embed.FS      { return queryFiles }

func (l *language) QueryNames(scope languages.Scope) []string {
	switch scope {
	case languages.ScopeLSP:
		return []string{"htmlTemplates", "completionContext", "imports", "classes", "classMemberDeclaration", "exports", "importAttributes", "definedElements"}
	case languages.ScopeGenerate:
		return []string{"classMemberDeclaration", "classes", "declarations", "imports", "typeAliases", "exports"}
	case languages.ScopeAll:
		return []string{"classMemberDeclaration", "classes", "declarations", "imports", "htmlTemplates", "completionContext", "typeAliases", "exports", "importAttributes", "definedElements"}
	}
	return nil
}

// GetParser borrows a TypeScript parser from the pool.
func GetParser() *ts.Parser { return lang.pool.Get().(*ts.Parser) }

// PutParser returns a TypeScript parser to the pool.
func PutParser(parser *ts.Parser) { parser.Reset(); lang.pool.Put(parser) }

func (l *language) GetParser() *ts.Parser      { return GetParser() }
func (l *language) PutParser(parser *ts.Parser) { PutParser(parser) }

// TSLanguage returns the tree-sitter TypeScript language grammar.
func TSLanguage() *ts.Language { return lang.tsLang }
