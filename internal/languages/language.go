// Package languages defines the interface for tree-sitter language integrations
// and provides a registry for language packages to register themselves.
// Each language lives in its own subpackage (e.g. internal/languages/html) and
// registers via init(). This enables per-language review isolation: changes to
// one language don't touch other language packages.
package languages

import (
	"fmt"
	"sync"

	ts "github.com/tree-sitter/go-tree-sitter"
)

// Scope controls which queries are loaded for a language.
type Scope int

const (
	// ScopeLSP loads queries needed for LSP operations.
	ScopeLSP Scope = iota
	// ScopeGenerate loads queries needed for manifest generation.
	ScopeGenerate
	// ScopeAll loads all available queries.
	ScopeAll
)

// Language defines the interface for a tree-sitter language integration.
type Language interface {
	// Name returns the language identifier (e.g. "html", "blade").
	Name() string
	// TSLanguage returns the compiled tree-sitter grammar.
	TSLanguage() *ts.Language
	// QueryDir returns the embed.FS subdirectory containing .scm query files.
	// Languages that reuse another language's queries return that language's
	// directory (e.g. Blade returns "html").
	QueryDir() string
	// QueryNames returns which queries to load for the given scope.
	QueryNames(scope Scope) []string
	// GetParser borrows a parser from the pool. Always defer PutParser.
	GetParser() *ts.Parser
	// PutParser returns a parser to the pool.
	PutParser(parser *ts.Parser)
}

// NewParserPool creates a sync.Pool that produces parsers configured for the
// given tree-sitter language. Used by per-language packages in their init.
func NewParserPool(lang *ts.Language) *sync.Pool {
	return &sync.Pool{
		New: func() any {
			parser := ts.NewParser()
			if err := parser.SetLanguage(lang); err != nil {
				panic(fmt.Sprintf("failed to set language: %v", err))
			}
			return parser
		},
	}
}
