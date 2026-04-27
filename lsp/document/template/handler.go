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
	tree_sitter_embedded_template "github.com/tree-sitter/tree-sitter-embedded-template/bindings/go"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

// Handler implements language-specific operations for template documents
// (Nunjucks, Jinja2, Twig, Liquid, Handlebars, ERB, EJS).
// It uses a two-stage pipeline:
//  1. tree-sitter grammar extracts text/content nodes (HTML) from template source
//  2. The HTML handler parses the reconstructed HTML for custom elements
//
// Template blocks are replaced with whitespace to preserve line/column positions.
type Handler struct {
	htmlHandler types.LanguageHandler
	jinjaParser *sitter.Parser
	jinjaQuery  *sitter.Query
	hbsParser   *sitter.Parser
	hbsQuery    *sitter.Query
	etParser    *sitter.Parser
	etQuery     *sitter.Query
	mu          sync.Mutex
}

var (
	jinjaLang = sitter.NewLanguage(tree_sitter_jinja.Language())
	hbsLang   = sitter.NewLanguage(tree_sitter_handlebars.Language())
	etLang    = sitter.NewLanguage(tree_sitter_embedded_template.Language())
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

	etParser := sitter.NewParser()
	if err := etParser.SetLanguage(etLang); err != nil {
		jinjaParser.Close()
		jinjaQuery.Close()
		hbsParser.Close()
		hbsQuery.Close()
		return nil, fmt.Errorf("failed to set embedded-template language: %w", err)
	}

	etQuery, qerr := sitter.NewQuery(etLang, `(content) @html`)
	if qerr != nil {
		jinjaParser.Close()
		jinjaQuery.Close()
		hbsParser.Close()
		hbsQuery.Close()
		etParser.Close()
		return nil, fmt.Errorf("failed to compile embedded-template content query: %w", qerr)
	}

	return &Handler{
		htmlHandler: htmlHandler,
		jinjaParser: jinjaParser,
		jinjaQuery:  jinjaQuery,
		hbsParser:   hbsParser,
		hbsQuery:    hbsQuery,
		etParser:    etParser,
		etQuery:     etQuery,
	}, nil
}

func (h *Handler) Language() string {
	return "template"
}

// templateFamily determines which tree-sitter grammar to use based on file
// extension. Liquid files use the Jinja grammar since Liquid shares the same
// delimiter syntax ({%...%}, {{...}}). Liquid-specific constructs like
// {% raw %}...{% endraw %} are handled as custom statements by the Jinja
// grammar; HTML inside raw blocks is extracted on a best-effort basis
// depending on how the grammar parses the block content.
func templateFamily(uri string) string {
	lower := strings.ToLower(uri)
	switch {
	case strings.HasSuffix(lower, ".hbs"):
		return "handlebars"
	case strings.HasSuffix(lower, ".erb"), strings.HasSuffix(lower, ".ejs"):
		return "embedded-template"
	default:
		return "jinja"
	}
}

// extractHTML uses tree-sitter to find text/content nodes (HTML content), then
// replaces template blocks with whitespace to produce valid HTML with preserved
// positions. The mutex is held for the entire operation to prevent Close from
// freeing the parser or query while they are in use.
func (h *Handler) extractHTML(source []byte, uri string) []byte {
	family := templateFamily(uri)

	h.mu.Lock()
	defer h.mu.Unlock()

	var parser *sitter.Parser
	var query *sitter.Query

	switch family {
	case "handlebars":
		parser = h.hbsParser
		query = h.hbsQuery
	case "embedded-template":
		parser = h.etParser
		query = h.etQuery
	default:
		parser = h.jinjaParser
		query = h.jinjaQuery
	}

	if parser == nil || query == nil {
		return nil
	}
	parser.Reset()
	tree := parser.Parse(source, nil)

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
	if h.etParser != nil {
		h.etParser.Close()
		h.etParser = nil
	}
	if h.etQuery != nil {
		h.etQuery.Close()
		h.etQuery = nil
	}
}
