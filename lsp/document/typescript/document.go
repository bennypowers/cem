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
	"reflect"
	"strings"
	"sync"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	Q "bennypowers.dev/cem/queries"
	protocol "github.com/tliron/glsp/protocol_3_16"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// TypeScriptDocument represents a TypeScript document with tree-sitter parsing
type TypeScriptDocument struct {
	uri        string
	content    string
	version    int32
	language   string
	tree       *ts.Tree
	parser     *ts.Parser
	scriptTags []types.ScriptTag
	mu         sync.RWMutex
}

// Base document interface methods

// URI returns the document URI
func (d *TypeScriptDocument) URI() string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.uri
}

// Content returns the document content
func (d *TypeScriptDocument) Content() (string, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.content, nil
}

// Version returns the document version
func (d *TypeScriptDocument) Version() int32 {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.version
}

// Language returns the document language
func (d *TypeScriptDocument) Language() string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.language
}

// Tree returns the document's syntax tree
func (d *TypeScriptDocument) Tree() *ts.Tree {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.tree
}

// ScriptTags returns the parsed script tags
func (d *TypeScriptDocument) ScriptTags() []types.ScriptTag {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.scriptTags
}

// UpdateContent updates the document content
func (d *TypeScriptDocument) UpdateContent(content string, version int32) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.content = content
	d.version = version
}

// SetTree sets the document's syntax tree
func (d *TypeScriptDocument) SetTree(tree *ts.Tree) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.tree != nil {
		d.tree.Close()
	}
	d.tree = tree
}

// Parser returns the document's parser
func (d *TypeScriptDocument) Parser() *ts.Parser {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.parser
}

// SetParser sets the document's parser
func (d *TypeScriptDocument) SetParser(parser *ts.Parser) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.parser = parser
}

// SetScriptTags sets the script tags
func (d *TypeScriptDocument) SetScriptTags(scriptTags []types.ScriptTag) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.scriptTags = scriptTags
}

// Close cleans up document resources
func (d *TypeScriptDocument) Close() {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.tree != nil {
		d.tree.Close()
		d.tree = nil
	}

	if d.parser != nil {
		d.parser = nil
	}
}

// ByteRangeToProtocolRange converts byte range to protocol range
func (d *TypeScriptDocument) ByteRangeToProtocolRange(content string, startByte, endByte uint) protocol.Range {
	return protocol.Range{
		Start: d.byteOffsetToPosition(startByte),
		End:   d.byteOffsetToPosition(endByte),
	}
}

// Helper method for byte offset to position conversion
func (d *TypeScriptDocument) byteOffsetToPosition(offset uint) protocol.Position {
	line := uint32(0)
	char := uint32(0)

	for i, r := range d.content {
		if uint(i) >= offset {
			break
		}

		if r == '\n' {
			line++
			char = 0
		} else {
			char++
		}
	}

	return protocol.Position{
		Line:      line,
		Character: char,
	}
}

// Parse parses the TypeScript content using tree-sitter
func (d *TypeScriptDocument) Parse(content string) error {
	d.UpdateContent(content, d.version)

	parser := Q.RetrieveTypeScriptParser()
	if parser == nil {
		return fmt.Errorf("failed to get TypeScript parser")
	}
	d.SetParser(parser)

	tree := parser.Parse([]byte(content), nil)
	d.SetTree(tree)

	return nil
}

// findCustomElements finds custom elements in TypeScript template literals
func (d *TypeScriptDocument) findCustomElements(handler *Handler) ([]types.CustomElementMatch, error) {
	tree := d.Tree()
	if tree == nil {
		return nil, fmt.Errorf("no tree available for document")
	}

	_, err := d.Content()
	if err != nil {
		return nil, err
	}

	var elements []types.CustomElementMatch

	// First, find HTML template literals
	templates, err := d.findHTMLTemplates(handler)
	if err != nil {
		return nil, fmt.Errorf("failed to find HTML templates: %w", err)
	}

	// Parse each template as HTML to find custom elements
	for _, template := range templates {
		templateElements, err := d.parseHTMLInTemplate(template, handler)
		if err != nil {
			helpers.SafeDebugLog("[TypeScript] Failed to parse template: %v", err)
			continue
		}
		elements = append(elements, templateElements...)
	}

	return elements, nil
}

// findHTMLTemplates finds HTML template literals in the TypeScript document
func (d *TypeScriptDocument) findHTMLTemplates(handler *Handler) ([]TemplateContext, error) {
	tree := d.Tree()
	if tree == nil {
		return nil, fmt.Errorf("no tree available")
	}

	content, err := d.Content()
	if err != nil {
		return nil, err
	}

	var templates []TemplateContext

	// Create a fresh HTML templates query matcher for thread safety
	templatesMatcher, err := Q.GetCachedQueryMatcher(handler.queryManager, "typescript", "htmlTemplates")
	if err != nil {
		return nil, fmt.Errorf("failed to create HTML templates matcher: %w", err)
	}
	defer templatesMatcher.Close()

	// Find HTML templates with different parent capture names
	parentCaptureNames := []string{"html.template", "html.generic.template", "html.options.template", "innerHTML.assignment"}

	for _, parentName := range parentCaptureNames {
		for captureMap := range templatesMatcher.ParentCaptures(tree.RootNode(), []byte(content), parentName) {
			// Process different template literal capture names
			templateCaptureNames := []string{"template.literal", "generic.template.literal", "options.template.literal", "innerHTML.template"}

			for _, captureName := range templateCaptureNames {
				if templateCaptures, exists := captureMap[captureName]; exists {
					for _, capture := range templateCaptures {
						template := TemplateContext{
							Range:   d.ByteRangeToProtocolRange(content, capture.StartByte, capture.EndByte),
							content: capture.Text,
							Type:    "html",
						}
						templates = append(templates, template)
					}
				}
			}
		}
	}

	return templates, nil
}

// parseHTMLInTemplate parses HTML content within a template literal
func (d *TypeScriptDocument) parseHTMLInTemplate(template TemplateContext, handler *Handler) ([]types.CustomElementMatch, error) {
	templateContent, err := template.Content()
	if err != nil {
		return nil, err
	}

	// Create a temporary HTML parser to parse the template content
	htmlParser := Q.GetHTMLParser()
	if htmlParser == nil {
		return nil, fmt.Errorf("failed to get HTML parser")
	}
	defer Q.PutHTMLParser(htmlParser)

	htmlTree := htmlParser.Parse([]byte(templateContent), nil)
	if htmlTree == nil {
		return nil, fmt.Errorf("failed to parse HTML template")
	}
	defer htmlTree.Close()

	var elements []types.CustomElementMatch

	// Get HTML custom elements query from the query manager
	htmlCustomElements, err := Q.GetCachedQueryMatcher(handler.queryManager, "html", "customElements")
	if err != nil {
		return nil, fmt.Errorf("failed to get HTML custom elements query: %w", err)
	}

	// Find custom elements in the HTML template
	for captureMap := range htmlCustomElements.ParentCaptures(htmlTree.RootNode(), []byte(templateContent), "element") {
		if tagNames, exists := captureMap["tag.name"]; exists {
			for _, capture := range tagNames {
				if strings.Contains(capture.Text, "-") {
					// Adjust positions relative to the template start
					element := types.CustomElementMatch{
						TagName:    capture.Text,
						Range:      d.adjustRangeToTemplate(capture, template),
						Attributes: make(map[string]types.AttributeMatch),
					}
					elements = append(elements, element)
				}
			}
		}
	}

	return elements, nil
}

// adjustRangeToTemplate adjusts a range from template content to document coordinates
func (d *TypeScriptDocument) adjustRangeToTemplate(
	capture Q.CaptureInfo,
	template TemplateContext,
) protocol.Range {
	// This is a simplified version - a full implementation would need to:
	// 1. Calculate the byte offset of the template start in the document
	// 2. Add the capture's byte offsets to get document positions
	// 3. Convert byte positions back to line/character positions

	// For now, return the template range as a placeholder
	return template.Range
}

// analyzeCompletionContext analyzes completion context for TypeScript documents
func (d *TypeScriptDocument) analyzeCompletionContext(
	position protocol.Position,
	handler *Handler,
) *types.CompletionAnalysis {
	analysis := &types.CompletionAnalysis{
		Type: types.CompletionUnknown,
	}

	tree := d.Tree()
	if tree == nil {
		return analysis
	}

	content, err := d.Content()
	if err != nil {
		return analysis
	}

	// Convert position to byte offset
	byteOffset := d.positionToByteOffset(position, content)

	// Create a fresh completion context query matcher for thread safety
	completionMatcher, err := Q.GetCachedQueryMatcher(handler.queryManager, "typescript", "completionContext")
	if err != nil {
		helpers.SafeDebugLog("[TypeScript] Failed to create completion context matcher: %v", err)
		return analysis
	}
	defer completionMatcher.Close()

	// Use tree-sitter to analyze context
	helpers.SafeDebugLog("[TypeScript] Analyzing completion context at offset %d", byteOffset)

	captureCount := 0
	for captureMap := range completionMatcher.ParentCaptures(tree.RootNode(), []byte(content), "context") {
		captureCount++
		helpers.SafeDebugLog("[TypeScript] Found capture map %d with keys: %v", captureCount, getCaptureMapKeys(captureMap))

		// Analyze captures to determine completion context
		for captureName, captures := range captureMap {
			helpers.SafeDebugLog("[TypeScript] Checking capture %s with %d captures", captureName, len(captures))
			for _, capture := range captures {
				helpers.SafeDebugLog("[TypeScript] Capture %s: text='%s' start=%d end=%d (cursor at %d)",
					captureName, capture.Text, capture.StartByte, capture.EndByte, byteOffset)
				if capture.StartByte <= byteOffset && byteOffset <= capture.EndByte {
					helpers.SafeDebugLog("[TypeScript] Match found for %s", captureName)
					switch captureName {
					case "template.content":
						// Inside an HTML template literal
						helpers.SafeDebugLog("[TypeScript] Setting IsLitTemplate = true")
						analysis.IsLitTemplate = true
						// Analyze the HTML content within the template
						analysis = d.analyzeTemplateCompletionContext(byteOffset, analysis, handler)
					case "interpolation":
						// Inside TypeScript interpolation - no HTML completions
						helpers.SafeDebugLog("[TypeScript] Found interpolation, returning CompletionUnknown")
						analysis.Type = types.CompletionUnknown
						return analysis
					}
				}
			}
		}
	}

	if captureCount == 0 {
		helpers.SafeDebugLog("[TypeScript] No captures found for completion context analysis")
	}

	return analysis
}

// getCaptureMapKeys returns the keys from a capture map for debugging
func getCaptureMapKeys(captureMap Q.CaptureMap) []string {
	keys := make([]string, 0, len(captureMap))
	for k := range captureMap {
		keys = append(keys, k)
	}
	return keys
}

// analyzeTemplateCompletionContext analyzes completion context within template literals
func (d *TypeScriptDocument) analyzeTemplateCompletionContext(
	byteOffset uint,
	analysis *types.CompletionAnalysis,
	handler *Handler,
) *types.CompletionAnalysis {
	helpers.SafeDebugLog("[TypeScript] analyzeTemplateCompletionContext called with offset %d", byteOffset)
	// Find which template the cursor is in
	templates, err := d.findHTMLTemplates(handler)
	if err != nil {
		helpers.SafeDebugLog("[TypeScript] Failed to find HTML templates: %v", err)
		return analysis
	}
	helpers.SafeDebugLog("[TypeScript] Found %d templates", len(templates))

	content, err := d.Content()
	if err != nil {
		return analysis
	}

	for _, template := range templates {
		templateStart := d.positionToByteOffset(template.Range.Start, content)
		templateEnd := d.positionToByteOffset(template.Range.End, content)

		if templateStart <= byteOffset && byteOffset <= templateEnd {
			// Cursor is in this template - analyze the HTML content
			relativeOffset := byteOffset - templateStart
			templateContent, _ := template.Content()

			// Use HTML completion context analysis on the template content
			analysis = d.analyzeTemplateContentAsHTML(templateContent, relativeOffset, analysis, handler)
			break
		}
	}

	return analysis
}

// analyzeTemplateContentAsHTML analyzes HTML content within a template
func (d *TypeScriptDocument) analyzeTemplateContentAsHTML(
	templateContent string,
	relativeOffset uint,
	analysis *types.CompletionAnalysis,
	handler *Handler,
) *types.CompletionAnalysis {
	// Create temporary HTML parser
	htmlParser := Q.GetHTMLParser()
	if htmlParser == nil {
		return analysis
	}
	defer Q.PutHTMLParser(htmlParser)

	htmlTree := htmlParser.Parse([]byte(templateContent), nil)
	if htmlTree == nil {
		return analysis
	}
	defer htmlTree.Close()

	// Get HTML completion context query
	htmlCompletionContext, err := Q.GetCachedQueryMatcher(handler.queryManager, "html", "completionContext")
	if err != nil {
		return analysis
	}

	// First pass: collect all tag names from all capture maps
	var allTagNames []Q.CaptureInfo
	for captureMap := range htmlCompletionContext.ParentCaptures(htmlTree.RootNode(), []byte(templateContent), "context") {
		if tagNames, exists := captureMap["tag.name"]; exists {
			allTagNames = append(allTagNames, tagNames...)
		}
	}

	// Second pass: analyze completion context at the relative offset
	for captureMap := range htmlCompletionContext.ParentCaptures(htmlTree.RootNode(), []byte(templateContent), "context") {
		// Find completion context at the relative offset
		for captureName, captures := range captureMap {
			for _, capture := range captures {
				if capture.StartByte <= relativeOffset && relativeOffset <= capture.EndByte {
					switch captureName {
					case "tag.name":
						analysis.Type = types.CompletionTagName
						analysis.TagName = capture.Text
						analysis.IsLitTemplate = true
					case "attr.name":
						attrText := capture.Text
						// Check for Lit syntax patterns in attribute names
						if strings.HasPrefix(attrText, "@") {
							analysis.Type = types.CompletionLitEventBinding
							analysis.LitSyntax = "@"
							analysis.IsLitTemplate = true
							// Remove @ prefix for the actual attribute name
							if len(attrText) > 1 {
								analysis.AttributeName = attrText[1:]
							}
						} else if strings.HasPrefix(attrText, ".") {
							analysis.Type = types.CompletionLitPropertyBinding
							analysis.LitSyntax = "."
							analysis.IsLitTemplate = true
							// Remove . prefix for the actual property name
							if len(attrText) > 1 {
								analysis.AttributeName = attrText[1:]
							}
						} else if strings.HasPrefix(attrText, "?") {
							analysis.Type = types.CompletionLitBooleanAttribute
							analysis.LitSyntax = "?"
							analysis.IsLitTemplate = true
							// Remove ? prefix for the actual attribute name
							if len(attrText) > 1 {
								analysis.AttributeName = attrText[1:]
							}
						} else {
							analysis.Type = types.CompletionAttributeName
							analysis.AttributeName = attrText
							analysis.IsLitTemplate = true
						}

						// Find the associated tag name for attribute completions from all collected tag names
						for _, tagCapture := range allTagNames {
							if tagCapture.EndByte <= relativeOffset {
								analysis.TagName = tagCapture.Text
							}
						}
					case "attr.value":
						analysis.Type = types.CompletionAttributeValue
						analysis.IsLitTemplate = true
						// For attribute value completion, find the attribute name
						if attrNames, hasAttr := captureMap["attr.name"]; hasAttr {
							for _, attrCapture := range attrNames {
								if attrCapture.EndByte <= relativeOffset {
									analysis.AttributeName = attrCapture.Text
								}
							}
						}
						// Also find the tag name from all collected tag names
						for _, tagCapture := range allTagNames {
							if tagCapture.EndByte <= relativeOffset {
								analysis.TagName = tagCapture.Text
							}
						}
					}
				}
			}
		}
	}

	return analysis
}

// positionToByteOffset converts LSP position to byte offset
func (d *TypeScriptDocument) positionToByteOffset(pos protocol.Position, content string) uint {
	var offset uint = 0
	lines := strings.Split(content, "\n")

	// Add bytes for complete lines before the target line
	for i := uint32(0); i < pos.Line && i < uint32(len(lines)); i++ {
		offset += uint(len(lines[i])) + 1 // +1 for newline
	}

	// Add bytes for characters in the target line
	if pos.Line < uint32(len(lines)) {
		line := lines[pos.Line]
		if pos.Character < uint32(len(line)) {
			offset += uint(pos.Character)
		} else {
			offset += uint(len(line))
		}
	}

	return offset
}

// CompletionPrefix extracts the prefix being typed for filtering completions
func (d *TypeScriptDocument) CompletionPrefix(analysis *types.CompletionAnalysis) string {
	if analysis == nil {
		return ""
	}

	switch analysis.Type {
	case types.CompletionTagName:
		return analysis.TagName
	case types.CompletionAttributeName:
		return analysis.AttributeName
	case types.CompletionAttributeValue:
		return analysis.AttributeName // For attribute values, we filter on the attribute name
	default:
		return ""
	}
}

// AnalyzeCompletionContextTS analyzes completion context using tree-sitter queries (interface method)
func (d *TypeScriptDocument) AnalyzeCompletionContextTS(position protocol.Position, dm any) *types.CompletionAnalysis {
	// Use reflection to call GetLanguageHandler to avoid circular imports
	dmValue := reflect.ValueOf(dm)
	if dmValue.Kind() == reflect.Pointer && !dmValue.IsNil() {
		method := dmValue.MethodByName("GetLanguageHandler")
		if method.IsValid() {
			results := method.Call([]reflect.Value{reflect.ValueOf("typescript")})
			if len(results) > 0 && !results[0].IsNil() {
				handler := results[0].Interface()
				if h, ok := handler.(interface {
					AnalyzeCompletionContext(types.Document, protocol.Position) *types.CompletionAnalysis
				}); ok {
					result := h.AnalyzeCompletionContext(d, position)
					return result
				}
			}
		}
	}
	// Fallback: return safe default
	return &types.CompletionAnalysis{
		Type: types.CompletionUnknown,
	}
}

// FindElementAtPosition finds a custom element at the given position (interface method)
func (d *TypeScriptDocument) FindElementAtPosition(
	position protocol.Position,
	dm any,
) *types.CustomElementMatch {
	// Use reflection to call GetLanguageHandler to avoid circular imports
	dmValue := reflect.ValueOf(dm)
	if dmValue.Kind() == reflect.Pointer && !dmValue.IsNil() {
		method := dmValue.MethodByName("GetLanguageHandler")
		if method.IsValid() {
			results := method.Call([]reflect.Value{reflect.ValueOf("typescript")})
			if len(results) > 0 && !results[0].IsNil() {
				handler := results[0].Interface()
				if h, ok := handler.(interface {
					FindElementAtPosition(types.Document, protocol.Position) *types.CustomElementMatch
				}); ok {
					return h.FindElementAtPosition(d, position)
				}
			}
		}
	}
	// Fallback: return nil if handler not available
	return nil
}

// FindAttributeAtPosition finds an attribute at the given position (interface method)
func (d *TypeScriptDocument) FindAttributeAtPosition(
	position protocol.Position,
	dm any,
) (*types.AttributeMatch, string) {
	// This method needs to be implemented for the types.Document interface
	// For now, return nil as this functionality is handled by the handler
	return nil, ""
}

// FindCustomElements finds custom elements in the document (interface method)
func (d *TypeScriptDocument) FindCustomElements(
	dm any,
) ([]types.CustomElementMatch, error) {
	// Use reflection to call GetLanguageHandler to avoid circular imports
	dmValue := reflect.ValueOf(dm)
	if dmValue.Kind() == reflect.Ptr && !dmValue.IsNil() {
		method := dmValue.MethodByName("GetLanguageHandler")
		if method.IsValid() {
			results := method.Call([]reflect.Value{reflect.ValueOf("typescript")})
			if len(results) > 0 && !results[0].IsNil() {
				handler := results[0].Interface()
				if h, ok := handler.(interface{ FindCustomElements(types.Document) ([]types.CustomElementMatch, error) }); ok {
					return h.FindCustomElements(d)
				}
			}
		}
	}
	// Fallback: return empty slice if handler not available
	return []types.CustomElementMatch{}, nil
}

// FindModuleScript finds insertion point in module script
func (d *TypeScriptDocument) FindModuleScript() (protocol.Position, bool) {
	// TypeScript files are modules by default, return end of file
	content, err := d.Content()
	if err != nil {
		return protocol.Position{}, false
	}
	lines := strings.Split(content, "\n")
	return protocol.Position{Line: uint32(len(lines)), Character: 0}, true
}

// FindInlineModuleScript finds insertion point in inline module script
func (d *TypeScriptDocument) FindInlineModuleScript() (protocol.Position, bool) {
	return d.FindModuleScript()
}

// FindHeadInsertionPoint finds insertion point in <head> section
func (d *TypeScriptDocument) FindHeadInsertionPoint(dm any) (protocol.Position, bool) {
	// TypeScript files don't have <head> sections
	return protocol.Position{}, false
}

