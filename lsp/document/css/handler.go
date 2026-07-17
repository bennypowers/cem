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
package css

import (
	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	Q "bennypowers.dev/cem/internal/treesitter"
	"go.lsp.dev/protocol"
)

// Handler implements language-specific operations for CSS documents.
type Handler struct{}

// NewHandler creates a new CSS language handler
func NewHandler(_ *Q.QueryManager) (*Handler, error) {
	return &Handler{}, nil
}

func (h *Handler) Language() string { return "css" }

// CreateDocument creates a new CSS document
func (h *Handler) CreateDocument(uri, content string, version int32) types.Document {
	doc := NewCSSDocument(uri, content, version)
	if err := doc.Parse(content); err != nil {
		helpers.SafeDebugLog("[CSS] Failed to parse document %s: %v", uri, err)
	}
	return doc
}

func (h *Handler) FindCustomElements(_ types.Document) ([]types.CustomElementMatch, error) {
	return nil, nil
}

func (h *Handler) AnalyzeCompletionContext(_ types.Document, _ protocol.Position) *types.CompletionAnalysis {
	return &types.CompletionAnalysis{}
}

func (h *Handler) FindElementAtPosition(_ types.Document, _ protocol.Position) *types.CustomElementMatch {
	return nil
}

func (h *Handler) FindAttributeAtPosition(_ types.Document, _ protocol.Position) (*types.AttributeMatch, string) {
	return nil, ""
}

func (h *Handler) Close() {}
