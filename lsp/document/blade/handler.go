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
package blade

import (
	"bennypowers.dev/cem/lsp/document/html"
	"bennypowers.dev/cem/lsp/types"
	Q "bennypowers.dev/cem/queries"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// Handler implements language-specific operations for Blade template documents
// (.blade.php). Blade's tree-sitter grammar extends tree-sitter-html, producing
// identical node names (element, start_tag, tag_name, attribute). The handler
// parses with the Blade grammar and runs HTML query patterns compiled against
// the Blade language via a shared html.Handler configured for "blade" queries.
type Handler struct {
	htmlHandler *html.Handler
}

// NewHandler creates a new Blade language handler. It configures an html.Handler
// with "blade" query language so the same .scm patterns run against the Blade
// grammar's node types.
func NewHandler(queryManager *Q.QueryManager) (*Handler, error) {
	htmlHandler, err := html.NewHandlerWithLanguage(queryManager, "blade")
	if err != nil {
		return nil, err
	}

	return &Handler{
		htmlHandler: htmlHandler,
	}, nil
}

// Language returns the language identifier.
func (h *Handler) Language() string {
	return "blade"
}

// CreateDocument parses the content with the Blade grammar and creates an
// HTMLDocument with the resulting tree. The document's queries use "blade"
// language, so captures match the Blade grammar's node type IDs.
func (h *Handler) CreateDocument(uri, content string, version int32) types.Document {
	parser := Q.GetBladeParser()
	defer Q.PutBladeParser(parser)
	tree := parser.Parse([]byte(content), nil)

	return h.htmlHandler.CreateDocumentWithTree(uri, content, version, tree)
}

// FindCustomElements delegates to the blade-configured HTML handler.
func (h *Handler) FindCustomElements(doc types.Document) ([]types.CustomElementMatch, error) {
	return h.htmlHandler.FindCustomElements(doc)
}

// AnalyzeCompletionContext delegates to the blade-configured HTML handler.
func (h *Handler) AnalyzeCompletionContext(doc types.Document, position protocol.Position) *types.CompletionAnalysis {
	return h.htmlHandler.AnalyzeCompletionContext(doc, position)
}

// FindElementAtPosition delegates to the blade-configured HTML handler.
func (h *Handler) FindElementAtPosition(doc types.Document, position protocol.Position) *types.CustomElementMatch {
	return h.htmlHandler.FindElementAtPosition(doc, position)
}

// FindAttributeAtPosition delegates to the blade-configured HTML handler.
func (h *Handler) FindAttributeAtPosition(doc types.Document, position protocol.Position) (*types.AttributeMatch, string) {
	return h.htmlHandler.FindAttributeAtPosition(doc, position)
}

// FindHeadInsertionPoint delegates to the blade-configured HTML handler.
func (h *Handler) FindHeadInsertionPoint(doc types.Document) (protocol.Position, bool) {
	return h.htmlHandler.FindHeadInsertionPoint(doc)
}

// Close cleans up the underlying HTML handler.
func (h *Handler) Close() {
	if h.htmlHandler != nil {
		h.htmlHandler.Close()
	}
}
