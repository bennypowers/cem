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

	phplang "bennypowers.dev/cem/internal/languages/php"
	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	protocol "github.com/tliron/glsp/protocol_3_16"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

// Handler implements language-specific operations for PHP documents.
// It uses tree-sitter language injection:
//  1. tree-sitter-php extracts HTML region byte ranges from PHP source
//  2. The HTML parser is restricted to those ranges via SetIncludedRanges
//
// Byte positions in the resulting tree map to the original source automatically.
// Parsers are obtained from pools for concurrency; the query is compiled once
// at init time and shared immutably.
type Handler struct {
	htmlHandler types.LanguageHandler
	textQuery   *sitter.Query
}

// NewHandler creates a new PHP language handler
func NewHandler(htmlHandler types.LanguageHandler) (*Handler, error) {
	textQuery, qerr := sitter.NewQuery(phplang.TSLanguage(), `(text) @html`)
	if qerr != nil {
		return nil, fmt.Errorf("failed to compile text query: %w", qerr)
	}

	return &Handler{
		htmlHandler: htmlHandler,
		textQuery:   textQuery,
	}, nil
}

// Language returns the language identifier
func (h *Handler) Language() string {
	return "php"
}

// htmlRanges uses tree-sitter-php to find HTML text nodes and returns their
// byte ranges for use with tree-sitter language injection.
func (h *Handler) htmlRanges(source []byte) []sitter.Range {
	parser := phplang.GetParser()
	defer phplang.PutParser(parser)

	tree := parser.Parse(source, nil)
	if tree == nil {
		return nil
	}
	defer tree.Close()

	cursor := sitter.NewQueryCursor()
	defer cursor.Close()

	var ranges []sitter.Range
	matches := cursor.Matches(h.textQuery, tree.RootNode(), source)
	for m := matches.Next(); m != nil; m = matches.Next() {
		for _, c := range m.Captures {
			n := c.Node
			ranges = append(ranges, n.Range())
		}
	}

	return ranges
}

// CreateDocument creates a new PHP document by extracting HTML ranges and
// delegating to the HTML handler with tree-sitter language injection.
func (h *Handler) CreateDocument(uri, content string, version int32) types.Document {
	ranges := h.htmlRanges([]byte(content))
	if rp, ok := h.htmlHandler.(types.RangeParser); ok {
		return rp.CreateDocumentWithRanges(uri, content, version, ranges)
	}
	helpers.SafeDebugLog("[PHP] htmlHandler does not implement RangeParser for %s", uri)
	return h.htmlHandler.CreateDocument(uri, content, version)
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

// Close cleans up query resources
func (h *Handler) Close() {
	if h.textQuery != nil {
		h.textQuery.Close()
		h.textQuery = nil
	}
}
