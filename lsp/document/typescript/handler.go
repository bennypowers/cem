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
package typescript

import (
	"fmt"
	"sync"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	Q "bennypowers.dev/cem/queries"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// Handler implements language-specific operations for TypeScript documents
type Handler struct {
	queryManager        *Q.QueryManager
	tsHtmlTemplates     *Q.QueryMatcher
	tsCompletionContext *Q.QueryMatcher
	mu                  sync.RWMutex
}

// NewHandler creates a new TypeScript language handler
func NewHandler(queryManager *Q.QueryManager) (*Handler, error) {
	h := &Handler{
		queryManager: queryManager,
	}

	var err error
	if h.tsHtmlTemplates, err = Q.GetCachedQueryMatcher(h.queryManager, "typescript", "htmlTemplates"); err != nil {
		return nil, fmt.Errorf("failed to load TypeScript HTML templates query: %w", err)
	}

	if h.tsCompletionContext, err = Q.GetCachedQueryMatcher(h.queryManager, "typescript", "completionContext"); err != nil {
		return nil, fmt.Errorf("failed to load TypeScript completion context query: %w", err)
	}

	return h, nil
}

// Language returns the language identifier
func (h *Handler) Language() string {
	return "typescript"
}

// CreateDocument creates a new TypeScript document
func (h *Handler) CreateDocument(uri, content string, version int32) types.Document {
	doc := &TypeScriptDocument{
		uri:      uri,
		content:  content,
		version:  version,
		language: "typescript",
	}

	// Parse the document
	if err := doc.Parse(content); err != nil {
		helpers.SafeDebugLog("[TypeScript] Failed to parse document %s: %v", uri, err)
	}

	return doc
}

// FindCustomElements finds custom elements in TypeScript template literals
func (h *Handler) FindCustomElements(doc types.Document) ([]types.CustomElementMatch, error) {
	tsDoc, ok := doc.(*TypeScriptDocument)
	if !ok {
		return nil, fmt.Errorf("document is not a TypeScript document")
	}

	return tsDoc.findCustomElements(h)
}

// AnalyzeCompletionContext analyzes completion context for TypeScript documents
func (h *Handler) AnalyzeCompletionContext(doc types.Document, position protocol.Position) *types.CompletionAnalysis {
	tsDoc, ok := doc.(*TypeScriptDocument)
	if !ok {
		return &types.CompletionAnalysis{}
	}

	return tsDoc.analyzeCompletionContext(position, h)
}

// FindElementAtPosition finds a custom element at the given position
func (h *Handler) FindElementAtPosition(doc types.Document, position protocol.Position) *types.CustomElementMatch {
	elements, err := h.FindCustomElements(doc)
	if err != nil {
		helpers.SafeDebugLog("[TypeScript] Failed to find elements: %v", err)
		return nil
	}

	for _, element := range elements {
		if isPositionInRange(position, element.Range) {
			return &types.CustomElementMatch{
				TagName:    element.TagName,
				Range:      element.Range,
				Attributes: element.Attributes,
			}
		}
	}

	return nil
}

// FindAttributeAtPosition finds an attribute at the given position
func (h *Handler) FindAttributeAtPosition(doc types.Document, position protocol.Position) (*types.AttributeMatch, string) {
	elements, err := h.FindCustomElements(doc)
	if err != nil {
		helpers.SafeDebugLog("[TypeScript] Failed to find elements: %v", err)
		return nil, ""
	}

	for _, element := range elements {
		for _, attr := range element.Attributes {
			if isPositionInRange(position, attr.Range) {
				return &types.AttributeMatch{
					Name:  attr.Name,
					Value: attr.Value,
					Range: attr.Range,
				}, element.TagName
			}
		}
	}

	return nil, ""
}

// Close cleans up the handler resources
func (h *Handler) Close() {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.tsHtmlTemplates != nil {
		h.tsHtmlTemplates.Close()
		h.tsHtmlTemplates = nil
	}
	if h.tsCompletionContext != nil {
		h.tsCompletionContext.Close()
		h.tsCompletionContext = nil
	}
}

// Helper function to check if position is within range
func isPositionInRange(pos protocol.Position, r protocol.Range) bool {
	if pos.Line < r.Start.Line || pos.Line > r.End.Line {
		return false
	}

	if pos.Line == r.Start.Line && pos.Character < r.Start.Character {
		return false
	}

	if pos.Line == r.End.Line && pos.Character > r.End.Character {
		return false
	}

	return true
}