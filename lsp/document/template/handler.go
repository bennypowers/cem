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
package template

import (
	"fmt"
	"strings"
	"sync"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	tree_sitter_handlebars "bennypowers.dev/tree-sitter-handlebars/bindings/go"
	tree_sitter_jinja "bennypowers.dev/tree-sitter-jinja-dialects/bindings/go"
	protocol "github.com/tliron/glsp/protocol_3_16"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

// Handler implements language-specific operations for template documents
// (Nunjucks, Jinja2, Twig, Liquid, Handlebars).
// It uses a two-stage pipeline:
//  1. tree-sitter grammar extracts text nodes (HTML content) from template source
//  2. The HTML handler parses the reconstructed HTML for custom elements
//
// Template blocks are replaced with whitespace to preserve line/column positions.
type Handler struct {
	htmlHandler types.LanguageHandler
	jinjaParser *sitter.Parser
	jinjaQuery  *sitter.Query
	hbsParser   *sitter.Parser
	hbsQuery    *sitter.Query
	mu          sync.RWMutex
}

var (
	jinjaLang = sitter.NewLanguage(tree_sitter_jinja.Language())
	hbsLang   = sitter.NewLanguage(tree_sitter_handlebars.Language())
)

func NewHandler(htmlHandler types.LanguageHandler) (*Handler, error) {
	jinjaParser := sitter.NewParser()
	if err := jinjaParser.SetLanguage(jinjaLang); err != nil {
		return nil, fmt.Errorf("failed to set jinja language: %w", err)
	}

	jinjaQuery, qerr := sitter.NewQuery(jinjaLang, `(text) @html`)
	if qerr != nil {
		jinjaParser.Close()
		return nil, fmt.Errorf("failed to compile jinja text query: %w", qerr)
	}

	hbsParser := sitter.NewParser()
	if err := hbsParser.SetLanguage(hbsLang); err != nil {
		jinjaParser.Close()
		jinjaQuery.Close()
		return nil, fmt.Errorf("failed to set handlebars language: %w", err)
	}

	hbsQuery, qerr := sitter.NewQuery(hbsLang, `(text) @html`)
	if qerr != nil {
		jinjaParser.Close()
		jinjaQuery.Close()
		hbsParser.Close()
		return nil, fmt.Errorf("failed to compile handlebars text query: %w", qerr)
	}

	return &Handler{
		htmlHandler: htmlHandler,
		jinjaParser: jinjaParser,
		jinjaQuery:  jinjaQuery,
		hbsParser:   hbsParser,
		hbsQuery:    hbsQuery,
	}, nil
}

func (h *Handler) Language() string {
	return "template"
}

// templateFamily determines which tree-sitter grammar to use based on file extension.
func templateFamily(uri string) string {
	if strings.HasSuffix(strings.ToLower(uri), ".hbs") {
		return "handlebars"
	}
	return "jinja"
}

// extractHTML uses tree-sitter to find text nodes (HTML content), then replaces
// template blocks with whitespace to produce valid HTML with preserved positions.
func (h *Handler) extractHTML(source []byte, uri string) []byte {
	family := templateFamily(uri)

	h.mu.Lock()
	var parser *sitter.Parser
	var query *sitter.Query

	switch family {
	case "handlebars":
		parser = h.hbsParser
		query = h.hbsQuery
	default:
		parser = h.jinjaParser
		query = h.jinjaQuery
	}

	if parser == nil || query == nil {
		h.mu.Unlock()
		return nil
	}
	parser.Reset()
	tree := parser.Parse(source, nil)
	h.mu.Unlock()

	if tree == nil {
		return nil
	}
	defer tree.Close()

	buf := make([]byte, len(source))
	for i, b := range source {
		if b == '\n' {
			buf[i] = '\n'
		} else {
			buf[i] = ' '
		}
	}

	cursor := sitter.NewQueryCursor()
	defer cursor.Close()

	matches := cursor.Matches(query, tree.RootNode(), source)
	for m := matches.Next(); m != nil; m = matches.Next() {
		for _, c := range m.Captures {
			n := c.Node
			copy(buf[n.StartByte():n.EndByte()], source[n.StartByte():n.EndByte()])
		}
	}

	return buf
}

func (h *Handler) CreateDocument(uri, content string, version int32) types.Document {
	htmlContent := h.extractHTML([]byte(content), uri)
	if htmlContent == nil {
		helpers.SafeDebugLog("[TEMPLATE] failed to extract HTML from %s, falling back to raw content", uri)
		return h.htmlHandler.CreateDocument(uri, content, version)
	}

	return h.htmlHandler.CreateDocument(uri, string(htmlContent), version)
}

func (h *Handler) FindCustomElements(doc types.Document) ([]types.CustomElementMatch, error) {
	return h.htmlHandler.FindCustomElements(doc)
}

func (h *Handler) AnalyzeCompletionContext(doc types.Document, position protocol.Position) *types.CompletionAnalysis {
	return h.htmlHandler.AnalyzeCompletionContext(doc, position)
}

func (h *Handler) FindElementAtPosition(doc types.Document, position protocol.Position) *types.CustomElementMatch {
	return h.htmlHandler.FindElementAtPosition(doc, position)
}

func (h *Handler) FindAttributeAtPosition(doc types.Document, position protocol.Position) (*types.AttributeMatch, string) {
	return h.htmlHandler.FindAttributeAtPosition(doc, position)
}

func (h *Handler) Close() {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.jinjaParser != nil {
		h.jinjaParser.Close()
		h.jinjaParser = nil
	}
	if h.jinjaQuery != nil {
		h.jinjaQuery.Close()
		h.jinjaQuery = nil
	}
	if h.hbsParser != nil {
		h.hbsParser.Close()
		h.hbsParser = nil
	}
	if h.hbsQuery != nil {
		h.hbsQuery.Close()
		h.hbsQuery = nil
	}
}
