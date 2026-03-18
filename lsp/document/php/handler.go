/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package php

import (
	"fmt"
	"sync"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	Q "bennypowers.dev/cem/queries"
	protocol "github.com/tliron/glsp/protocol_3_16"
	sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_php "github.com/tree-sitter/tree-sitter-php/bindings/go"
)

// Handler implements language-specific operations for PHP documents.
// It uses a two-stage pipeline:
//  1. tree-sitter-php extracts HTML text nodes from PHP source
//  2. The HTML handler parses the reconstructed HTML for custom elements
//
// PHP blocks are replaced with whitespace to preserve line/column positions.
type Handler struct {
	htmlHandler  types.LanguageHandler
	queryManager *Q.QueryManager
	phpParser    *sitter.Parser
	textQuery    *sitter.Query
	mu           sync.RWMutex
}

var phpLang = sitter.NewLanguage(tree_sitter_php.LanguagePHP())

// NewHandler creates a new PHP language handler
func NewHandler(queryManager *Q.QueryManager, htmlHandler types.LanguageHandler) (*Handler, error) {
	parser := sitter.NewParser()
	if err := parser.SetLanguage(phpLang); err != nil {
		return nil, fmt.Errorf("failed to set PHP language: %w", err)
	}

	textQuery, qerr := sitter.NewQuery(phpLang, `(text) @html`)
	if qerr != nil {
		parser.Close()
		return nil, fmt.Errorf("failed to compile text query: %w", qerr)
	}

	return &Handler{
		htmlHandler:  htmlHandler,
		queryManager: queryManager,
		phpParser:    parser,
		textQuery:    textQuery,
	}, nil
}

// Language returns the language identifier
func (h *Handler) Language() string {
	return "php"
}

// extractHTML uses tree-sitter-php to find HTML text nodes, then replaces
// PHP blocks with whitespace to produce valid HTML with preserved positions.
func (h *Handler) extractHTML(source []byte) []byte {
	h.mu.Lock()
	h.phpParser.Reset()
	tree := h.phpParser.Parse(source, nil)
	h.mu.Unlock()

	if tree == nil {
		return nil
	}
	defer tree.Close()

	// Start with a buffer where PHP blocks are whitespace
	buf := make([]byte, len(source))
	for i, b := range source {
		if b == '\n' {
			buf[i] = '\n'
		} else {
			buf[i] = ' '
		}
	}

	// Copy HTML text nodes back into position
	cursor := sitter.NewQueryCursor()
	defer cursor.Close()

	matches := cursor.Matches(h.textQuery, tree.RootNode(), source)
	for m := matches.Next(); m != nil; m = matches.Next() {
		for _, c := range m.Captures {
			n := c.Node
			copy(buf[n.StartByte():n.EndByte()], source[n.StartByte():n.EndByte()])
		}
	}

	return buf
}

// CreateDocument creates a new PHP document by extracting HTML and delegating
// to the HTML handler
func (h *Handler) CreateDocument(uri, content string, version int32) types.Document {
	htmlContent := h.extractHTML([]byte(content))
	if htmlContent == nil {
		helpers.SafeDebugLog("[PHP] Failed to extract HTML from %s, falling back to raw content", uri)
		return h.htmlHandler.CreateDocument(uri, content, version)
	}

	return h.htmlHandler.CreateDocument(uri, string(htmlContent), version)
}

// FindCustomElements delegates to the HTML handler
func (h *Handler) FindCustomElements(doc types.Document) ([]types.CustomElementMatch, error) {
	return h.htmlHandler.FindCustomElements(doc)
}

// AnalyzeCompletionContext delegates to the HTML handler
func (h *Handler) AnalyzeCompletionContext(doc types.Document, position protocol.Position) *types.CompletionAnalysis {
	return h.htmlHandler.AnalyzeCompletionContext(doc, position)
}

// FindElementAtPosition delegates to the HTML handler
func (h *Handler) FindElementAtPosition(doc types.Document, position protocol.Position) *types.CustomElementMatch {
	return h.htmlHandler.FindElementAtPosition(doc, position)
}

// FindAttributeAtPosition delegates to the HTML handler
func (h *Handler) FindAttributeAtPosition(doc types.Document, position protocol.Position) (*types.AttributeMatch, string) {
	return h.htmlHandler.FindAttributeAtPosition(doc, position)
}

// Close cleans up resources
func (h *Handler) Close() {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.phpParser != nil {
		h.phpParser.Close()
		h.phpParser = nil
	}
	if h.textQuery != nil {
		h.textQuery.Close()
		h.textQuery = nil
	}
}
