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
package html

import (
	"fmt"
	"strings"
	"sync"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	Q "bennypowers.dev/cem/queries"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// Handler implements language-specific operations for HTML documents
type Handler struct {
	queryManager *Q.QueryManager
	mu           sync.RWMutex
}

// NewHandler creates a new HTML language handler
func NewHandler(queryManager *Q.QueryManager) (*Handler, error) {
	h := &Handler{
		queryManager: queryManager,
	}

	return h, nil
}

// Language returns the language identifier
func (h *Handler) Language() string {
	return "html"
}

// CreateDocument creates a new HTML document
func (h *Handler) CreateDocument(uri, content string, version int32) types.Document {
	doc := &HTMLDocument{
		uri:      uri,
		content:  content,
		version:  version,
		language: "html",
	}

	// Parse the document
	if err := doc.Parse(content); err != nil {
		helpers.SafeDebugLog("[HTML] Failed to parse document %s: %v", uri, err)
	}

	// Parse script tags using the handler's tree-sitter implementation
	if scriptTags, err := h.ParseScriptTags(doc); err != nil {
		helpers.SafeDebugLog("[HTML] Failed to parse script tags with handler: %v", err)
	} else {
		doc.SetScriptTags(scriptTags)
	}

	return doc
}

// FindCustomElements finds custom elements in the document
func (h *Handler) FindCustomElements(doc types.Document) ([]types.CustomElementMatch, error) {
	htmlDoc, ok := doc.(*HTMLDocument)
	if !ok {
		return nil, fmt.Errorf("document is not an HTML document")
	}

	return htmlDoc.findCustomElements(h)
}

// AnalyzeCompletionContext analyzes completion context for HTML documents
func (h *Handler) AnalyzeCompletionContext(doc types.Document, position protocol.Position) *types.CompletionAnalysis {
	htmlDoc, ok := doc.(*HTMLDocument)
	if !ok {
		return &types.CompletionAnalysis{}
	}

	return htmlDoc.analyzeCompletionContext(position, h)
}

// FindElementAtPosition finds a custom element at the given position
func (h *Handler) FindElementAtPosition(doc types.Document, position protocol.Position) *types.CustomElementMatch {
	elements, err := h.FindCustomElements(doc)
	if err != nil {
		helpers.SafeDebugLog("[HTML] Failed to find elements: %v", err)
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
		helpers.SafeDebugLog("[HTML] Failed to find elements: %v", err)
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

// ParseScriptTags parses script tags in the HTML document using tree-sitter
func (h *Handler) ParseScriptTags(doc types.Document) ([]types.ScriptTag, error) {
	htmlDoc, ok := doc.(*HTMLDocument)
	if !ok {
		return nil, fmt.Errorf("document is not an HTML document")
	}

	tree := htmlDoc.Tree()
	if tree == nil {
		return nil, fmt.Errorf("no tree available")
	}

	content, err := htmlDoc.Content()
	if err != nil {
		return nil, err
	}

	var scriptTags []types.ScriptTag

	// Create a fresh script tags query matcher for thread safety
	scriptTagsMatcher, err := Q.GetCachedQueryMatcher(h.queryManager, "html", "scriptTags")
	if err != nil {
		return nil, fmt.Errorf("failed to create script tags matcher: %w", err)
	}
	defer scriptTagsMatcher.Close()

	for captureMap := range scriptTagsMatcher.ParentCaptures(tree.RootNode(), []byte(content), "script") {
		scriptTag := types.ScriptTag{}

		// Extract the script element range
		if scriptCaptures, ok := captureMap["script"]; ok && len(scriptCaptures) > 0 {
			capture := scriptCaptures[0]
			scriptTag.Range = htmlDoc.ByteRangeToProtocolRange(content, capture.StartByte, capture.EndByte)
		}

		// Extract script content
		if contentCaptures, ok := captureMap["content"]; ok && len(contentCaptures) > 0 {
			capture := contentCaptures[0]
			scriptTag.ContentRange = htmlDoc.ByteRangeToProtocolRange(content, capture.StartByte, capture.EndByte)

			// Parse imports from the script content
			scriptTag.Imports = h.parseImportStatements(capture.Text)
		}

		// Extract attributes (type, src, etc.)
		if attrNames, ok := captureMap["attr.name"]; ok {
			if attrValues, ok := captureMap["attr.value"]; ok {
				// Build map of attribute name to value
				for i, nameCapture := range attrNames {
					if i < len(attrValues) {
						attrName := strings.ToLower(nameCapture.Text)
						attrValue := attrValues[i].Text

						// Remove quotes from attribute values
						if len(attrValue) >= 2 && (attrValue[0] == '"' || attrValue[0] == '\'') {
							attrValue = attrValue[1 : len(attrValue)-1]
						}

						switch attrName {
						case "type":
							scriptTag.Type = attrValue
							scriptTag.IsModule = (attrValue == "module")
						case "src":
							scriptTag.Src = attrValue
						}
					}
				}
			}
		}

		scriptTags = append(scriptTags, scriptTag)
	}

	return scriptTags, nil
}

// parseImportStatements parses import statements from script content
func (h *Handler) parseImportStatements(content string) []types.ImportStatement {
	var imports []types.ImportStatement

	// Simple regex-based parsing for now - this should be enhanced with proper TypeScript parsing
	lines := strings.SplitSeq(content, "\n")
	for line := range lines {
		line = strings.TrimSpace(line)

		// Static imports: import 'module' or import { x } from 'module'
		if strings.HasPrefix(line, "import ") {
			// Extract module path from import statement
			if path := extractImportPath(line); path != "" {
				imports = append(imports, types.ImportStatement{
					ImportPath: path,
					Type:       "static",
				})
			}
		}

		// Dynamic imports: import('module')
		if strings.Contains(line, "import(") {
			if path := extractDynamicImportPath(line); path != "" {
				imports = append(imports, types.ImportStatement{
					ImportPath: path,
					Type:       "dynamic",
				})
			}
		}
	}

	return imports
}

// extractImportPath extracts the module path from a static import statement
func extractImportPath(line string) string {
	// Look for 'module' or "module" patterns
	if idx := strings.Index(line, "'"); idx != -1 {
		rest := line[idx+1:]
		if endIdx := strings.Index(rest, "'"); endIdx != -1 {
			return rest[:endIdx]
		}
	}
	if idx := strings.Index(line, "\""); idx != -1 {
		rest := line[idx+1:]
		if endIdx := strings.Index(rest, "\""); endIdx != -1 {
			return rest[:endIdx]
		}
	}
	return ""
}

// extractDynamicImportPath extracts the module path from a dynamic import statement
func extractDynamicImportPath(line string) string {
	// Look for import('module') or import("module") patterns
	if idx := strings.Index(line, "import("); idx != -1 {
		rest := line[idx+7:]
		if strings.HasPrefix(rest, "'") {
			if endIdx := strings.Index(rest[1:], "'"); endIdx != -1 {
				return rest[1 : endIdx+1]
			}
		}
		if strings.HasPrefix(rest, "\"") {
			if endIdx := strings.Index(rest[1:], "\""); endIdx != -1 {
				return rest[1 : endIdx+1]
			}
		}
	}
	return ""
}

// Close cleans up the handler resources
func (h *Handler) Close() {
	h.mu.Lock()
	defer h.mu.Unlock()

	// No persistent QueryMatcher instances to clean up
	// QueryMatchers are now created fresh per operation and closed immediately
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

