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
	"reflect"
	"strings"
	"sync"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	Q "bennypowers.dev/cem/queries"
	protocol "github.com/tliron/glsp/protocol_3_16"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// HTMLDocument represents an HTML document with tree-sitter parsing
type HTMLDocument struct {
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
func (d *HTMLDocument) URI() string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.uri
}

// Content returns the document content
func (d *HTMLDocument) Content() (string, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.content, nil
}

// Version returns the document version
func (d *HTMLDocument) Version() int32 {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.version
}

// Language returns the document language
func (d *HTMLDocument) Language() string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.language
}

// Tree returns the document's syntax tree
func (d *HTMLDocument) Tree() *ts.Tree {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.tree
}

// ScriptTags returns the parsed script tags
func (d *HTMLDocument) ScriptTags() []types.ScriptTag {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.scriptTags
}

// UpdateContent updates the document content
func (d *HTMLDocument) UpdateContent(content string, version int32) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.content = content
	d.version = version
}

// SetTree sets the document's syntax tree
func (d *HTMLDocument) SetTree(tree *ts.Tree) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.tree != nil {
		d.tree.Close()
	}
	d.tree = tree
}

// Parser returns the document's parser
func (d *HTMLDocument) Parser() *ts.Parser {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.parser
}

// SetParser sets the document's parser
func (d *HTMLDocument) SetParser(parser *ts.Parser) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.parser = parser
}

// SetScriptTags sets the script tags
func (d *HTMLDocument) SetScriptTags(scriptTags []types.ScriptTag) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.scriptTags = scriptTags
}

// Close cleans up document resources
func (d *HTMLDocument) Close() {
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
func (d *HTMLDocument) ByteRangeToProtocolRange(content string, startByte, endByte uint) protocol.Range {
	return protocol.Range{
		Start: d.byteOffsetToPosition(startByte),
		End:   d.byteOffsetToPosition(endByte),
	}
}

// Helper method for byte offset to position conversion
func (d *HTMLDocument) byteOffsetToPosition(offset uint) protocol.Position {
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

// Parse parses the HTML content using tree-sitter
func (d *HTMLDocument) Parse(content string) error {
	d.UpdateContent(content, d.version)

	parser := Q.GetHTMLParser()
	if parser == nil {
		return fmt.Errorf("failed to get HTML parser")
	}
	d.SetParser(parser)

	tree := parser.Parse([]byte(content), nil)
	d.SetTree(tree)

	// Script tags are now parsed by the handler after document creation
	// This avoids circular import issues while using tree-sitter

	return nil
}

// findCustomElements finds custom elements in the HTML document
func (d *HTMLDocument) findCustomElements(handler *Handler) ([]types.CustomElementMatch, error) {
	var elements []types.CustomElementMatch

	// Safety checks
	if handler == nil {
		helpers.SafeDebugLog("[HTML] Handler is nil in findCustomElements")
		return elements, nil
	}

	// Hold read lock for the ENTIRE duration of tree usage to prevent Close() from freeing it
	d.mu.RLock()
	defer d.mu.RUnlock()

	tree := d.tree  // Access tree directly while holding lock
	if tree == nil {
		return elements, nil
	}

	content := d.content  // Access content directly while holding lock

	// Create a fresh query matcher for thread safety
	matcher, err := Q.GetCachedQueryMatcher(handler.queryManager, "html", "customElements")
	if err != nil {
		helpers.SafeDebugLog("[HTML] Failed to create custom elements matcher: %v", err)
		return elements, nil
	}
	defer matcher.Close()

	root := tree.RootNode()
	contentBytes := []byte(content)

	// Ensure tree size doesn't exceed content length
	if root.EndByte() > uint(len(contentBytes)) {
		// Tree is out of sync with content, return empty results
		return elements, nil
	}

	for captureMap := range matcher.ParentCaptures(root, contentBytes, "element") {
		if tagNames, ok := captureMap["tag.name"]; ok && len(tagNames) > 0 {
			tagName := tagNames[0].Text
			tagRange := d.ByteRangeToProtocolRange(content, tagNames[0].StartByte, tagNames[0].EndByte)

			// Collect attributes
			attributes := make(map[string]types.AttributeMatch)
			if attrNames, ok := captureMap["attr.name"]; ok {
				// Build a map of attribute values by their byte position for proper matching
				valuesByPosition := make(map[uint]string)

				// Collect quoted attribute values
				if attrValues, ok := captureMap["attr.value"]; ok {
					for _, attrValue := range attrValues {
						// Strip quotes from quoted attribute values
						value := attrValue.Text
						if len(value) >= 2 && (value[0] == '"' || value[0] == '\'') {
							value = value[1 : len(value)-1] // Remove first and last character (quotes)
						}
						valuesByPosition[attrValue.StartByte] = value
					}
				}

				// Collect unquoted attribute values
				if unquotedValues, ok := captureMap["attr.unquoted.value"]; ok {
					for _, unquotedValue := range unquotedValues {
						valuesByPosition[unquotedValue.StartByte] = unquotedValue.Text
					}
				}

				for _, attrName := range attrNames {
					attrMatch := types.AttributeMatch{
						Name:  attrName.Text,
						Range: d.ByteRangeToProtocolRange(content, attrName.StartByte, attrName.EndByte),
					}

					// Find the value that comes immediately after this attribute name
					// Look for the closest value that starts after the attribute name ends
					var closestValue string
					var closestDistance = ^uint(0) // Max uint value
					for valuePos, value := range valuesByPosition {
						if valuePos > attrName.EndByte {
							distance := valuePos - attrName.EndByte
							if distance < closestDistance {
								closestDistance = distance
								closestValue = value
							}
						}
					}
					if closestValue != "" {
						attrMatch.Value = closestValue
					}

					attributes[attrName.Text] = attrMatch
				}
			}

			elements = append(elements, types.CustomElementMatch{
				TagName:    tagName,
				Range:      tagRange,
				Attributes: attributes,
			})
		}
	}

	return elements, nil
}

// analyzeCompletionContext analyzes completion context for HTML documents
func (d *HTMLDocument) analyzeCompletionContext(position protocol.Position, handler *Handler) *types.CompletionAnalysis {
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

	// Debug logging for the failing test case
	helpers.SafeDebugLog("[HTML] analyzeCompletionContext: content=%q, position=%+v, byteOffset=%d", content, position, byteOffset)

	// Create a fresh completion context query matcher for thread safety
	completionMatcher, err := Q.GetCachedQueryMatcher(handler.queryManager, "html", "completionContext")
	if err != nil {
		helpers.SafeDebugLog("[HTML] Failed to create completion context matcher: %v", err)
		return analysis
	}
	defer completionMatcher.Close()

	// Use tree-sitter to analyze context
	// First, collect all captures across all capture maps
	allTagNames := []Q.CaptureInfo{}
	allAttrNames := []Q.CaptureInfo{}
	allAttrValues := []Q.CaptureInfo{}

	captureCount := 0
	for captureMap := range completionMatcher.ParentCaptures(tree.RootNode(), []byte(content), "context") {
		captureCount++
		// Collect all captures of each type
		if tagNames, exists := captureMap["tag.name"]; exists {
			allTagNames = append(allTagNames, tagNames...)
		}
		if attrNames, exists := captureMap["attr.name"]; exists {
			allAttrNames = append(allAttrNames, attrNames...)
		}
		if attrValues, exists := captureMap["attr.value"]; exists {
			allAttrValues = append(allAttrValues, attrValues...)
		}
	}

	// Process captures in priority order: tag.name > attr.value > attr.name
	// But prefer more specific (smaller) captures over broader ones

	// Check tag names first, but only if they're legitimate tag name captures
	for _, capture := range allTagNames {
		if capture.StartByte <= byteOffset && byteOffset <= capture.EndByte {
			// Only consider tag name completion if the capture text looks like a tag name
			// and doesn't contain attribute syntax or other HTML syntax
			captureText := capture.Text
			if !strings.Contains(captureText, " ") &&
				!strings.Contains(captureText, "=") &&
				!strings.Contains(captureText, ">") &&
				!strings.Contains(captureText, "\n") {
				// This looks like a pure tag name without attributes or other syntax
				analysis.Type = types.CompletionTagName
				analysis.TagName = captureText
				return analysis
			}
		}
	}

	// Check attribute values second
	for _, capture := range allAttrValues {
		if capture.StartByte <= byteOffset && byteOffset <= capture.EndByte {
			// For attribute values, we're more lenient since they can contain various content
			// Only reject captures that are clearly overly broad (contain multiple HTML elements)
			captureText := capture.Text
			// Count occurrences of < and newlines to detect overly broad captures
			openBrackets := strings.Count(captureText, "<")
			// If it contains multiple opening brackets, it's likely overly broad
			// For attribute values, one opening bracket is normal (the start of the element)
			// Reject only if there are multiple opening brackets or if we detect other problematic patterns
			newlineCount := strings.Count(captureText, "\n")

			if openBrackets <= 1 && newlineCount <= 1 {
				analysis.Type = types.CompletionAttributeValue
				// Find the closest preceding attribute name and tag name
				for _, attrCapture := range allAttrNames {
					if attrCapture.EndByte <= byteOffset {
						analysis.AttributeName = attrCapture.Text
					}
				}
				for _, tagCapture := range allTagNames {
					if tagCapture.EndByte <= byteOffset {
						analysis.TagName = tagCapture.Text
					}
				}
				return analysis
			}
		}
	}

	// Check attribute names last
	for _, capture := range allAttrNames {
		if capture.StartByte <= byteOffset && byteOffset <= capture.EndByte {
			// Only consider attribute name completion if the capture text looks like a real attribute name
			// and doesn't contain HTML syntax that suggests it's an overly broad capture
			captureText := capture.Text
			if !strings.Contains(captureText, "<") &&
				!strings.Contains(captureText, ">") &&
				!strings.Contains(captureText, "\n") &&
				!strings.Contains(captureText, "=") {
				analysis.Type = types.CompletionAttributeName
				analysis.AttributeName = captureText
				// Find the closest preceding tag name
				for _, tagCapture := range allTagNames {
					if tagCapture.EndByte <= byteOffset {
						analysis.TagName = tagCapture.Text
					}
				}
				return analysis
			}
		}
	}

	// Special case: check if cursor is positioned just after a custom element tag name or attribute (attribute name completion)
	if analysis.Type == types.CompletionUnknown && captureCount > 0 {
		for captureMap := range completionMatcher.ParentCaptures(tree.RootNode(), []byte(content), "context") {
			for captureName, captures := range captureMap {
				if captureName == "tag.name" {
					for _, capture := range captures {
						// Check if cursor is positioned 1-3 characters after the tag name (whitespace)
						if strings.Contains(capture.Text, "-") && // Custom element
							capture.EndByte < byteOffset &&
							byteOffset <= capture.EndByte+3 { // Within 3 chars after tag name

							// Check if the content after tag name is just whitespace
							afterTag := content[capture.EndByte:byteOffset]
							if strings.TrimSpace(afterTag) == "" {
								analysis.Type = types.CompletionAttributeName
								analysis.TagName = capture.Text
								return analysis
							}
						}
					}
				}
			}
		}

		// Additional special case: check if cursor is positioned after a complete attribute
		// This handles cases like <test-component color="red" |> where cursor is after the space
		for captureMap := range completionMatcher.ParentCaptures(tree.RootNode(), []byte(content), "context") {
			var associatedTagName string
			var latestAttrEnd uint

			// Find the associated tag name and latest attribute end position
			if tagNames, exists := captureMap["tag.name"]; exists && len(tagNames) > 0 {
				for _, tagCapture := range tagNames {
					if strings.Contains(tagCapture.Text, "-") && tagCapture.StartByte < byteOffset {
						associatedTagName = tagCapture.Text
					}
				}
			}

			// Find the latest attribute value that ends before our cursor
			if attrValues, exists := captureMap["attr.value"]; exists {
				for _, attrValue := range attrValues {
					if attrValue.EndByte < byteOffset && attrValue.EndByte > latestAttrEnd {
						latestAttrEnd = attrValue.EndByte
					}
				}
			}

			// Check unquoted attribute values too
			if unquotedValues, exists := captureMap["attr.unquoted.value"]; exists {
				for _, unquotedValue := range unquotedValues {
					if unquotedValue.EndByte < byteOffset && unquotedValue.EndByte > latestAttrEnd {
						latestAttrEnd = unquotedValue.EndByte
					}
				}
			}

			// If we found a tag name and an attribute ending before our cursor, check the gap
			if associatedTagName != "" && latestAttrEnd > 0 {
				// Check if the content between the attribute end and cursor is just whitespace
				gapContent := content[latestAttrEnd:byteOffset]
				// For quoted attributes, skip the closing quote
				if len(gapContent) > 0 && (gapContent[0] == '"' || gapContent[0] == '\'') {
					gapContent = gapContent[1:]
				}

				if strings.TrimSpace(gapContent) == "" && len(gapContent) <= 5 { // Allow reasonable whitespace
					analysis.Type = types.CompletionAttributeName
					analysis.TagName = associatedTagName
					return analysis
				}
			}
		}
	}

	// Final fallback: if no captures found, check if cursor is right after a '<' for tag name completion
	if analysis.Type == types.CompletionUnknown {
		// Check if the character immediately before the cursor is '<'
		if byteOffset > 0 && byteOffset <= uint(len(content)) {
			beforeCursor := content[:byteOffset]
			// Look for the last '<' character
			lastOpenBracket := strings.LastIndex(beforeCursor, "<")
			if lastOpenBracket != -1 {
				// Check if everything after the last '<' until cursor is just a valid tag name start (no spaces, attributes, etc.)
				afterBracket := beforeCursor[lastOpenBracket+1:]

				// If cursor is immediately after '<' or after a valid partial tag name
				if afterBracket == "" || (strings.TrimSpace(afterBracket) == afterBracket &&
					!strings.Contains(afterBracket, " ") &&
					!strings.Contains(afterBracket, "=") &&
					!strings.Contains(afterBracket, ">")) {
					analysis.Type = types.CompletionTagName
					// Don't set TagName - let the prefix extraction handle it
					return analysis
				}
			}
		}
	}

	if captureCount == 0 {
		helpers.SafeDebugLog("[HTML] No captures found for completion context analysis")
	}

	// Debug logging for the result
	helpers.SafeDebugLog("[HTML] analyzeCompletionContext result: Type=%d, TagName=%q", analysis.Type, analysis.TagName)

	return analysis
}

// positionToByteOffset converts LSP position to byte offset
func (d *HTMLDocument) positionToByteOffset(pos protocol.Position, content string) uint {
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
func (d *HTMLDocument) CompletionPrefix(analysis *types.CompletionAnalysis) string {
	if analysis == nil {
		return ""
	}

	content, err := d.Content()
	if err != nil {
		return ""
	}

	switch analysis.Type {
	case types.CompletionTagName:
		// Extract the tag name being typed (after the last "<")
		if lastOpenBracket := strings.LastIndex(content, "<"); lastOpenBracket != -1 {
			afterBracket := content[lastOpenBracket+1:]
			// Extract up to first space, >, or end of content
			for i, r := range afterBracket {
				if r == ' ' || r == '>' || r == '\n' || r == '\t' {
					return afterBracket[:i]
				}
			}
			return afterBracket
		}
		return ""

	case types.CompletionAttributeName:
		// Extract the attribute name being typed (after the last space in the tag)
		// Find the last occurrence of "<tagname" to locate the tag start
		words := strings.Fields(content)
		if len(words) > 0 {
			lastWord := words[len(words)-1]
			// If the last word doesn't contain "<" or "=", it's likely a partial attribute name
			if !strings.Contains(lastWord, "<") && !strings.Contains(lastWord, "=") {
				return lastWord
			}
		}
		return ""

	case types.CompletionAttributeValue:
		// Extract the value being typed (after the last quote)
		// Look for the last occurrence of " or '
		lastDoubleQuote := strings.LastIndex(content, "\"")
		lastSingleQuote := strings.LastIndex(content, "'")
		lastQuote := max(lastSingleQuote, lastDoubleQuote)
		if lastQuote != -1 {
			afterQuote := content[lastQuote+1:]
			// Return everything after the quote (the partial value being typed)
			return afterQuote
		}
		return ""

	default:
		return ""
	}
}

// AnalyzeCompletionContextTS analyzes completion context using tree-sitter queries (interface method)
func (d *HTMLDocument) AnalyzeCompletionContextTS(position protocol.Position, dm any) *types.CompletionAnalysis {
	helpers.SafeDebugLog("[HTML] AnalyzeCompletionContextTS called with position %+v", position)
	// Use reflection to call GetLanguageHandler to avoid circular imports
	dmValue := reflect.ValueOf(dm)
	if dmValue.Kind() == reflect.Pointer && !dmValue.IsNil() {
		method := dmValue.MethodByName("GetLanguageHandler")
		if method.IsValid() {
			results := method.Call([]reflect.Value{reflect.ValueOf("html")})
			if len(results) > 0 && !results[0].IsNil() {
				handler := results[0].Interface()
				if h, ok := handler.(interface {
					AnalyzeCompletionContext(types.Document, protocol.Position) *types.CompletionAnalysis
				}); ok {
					result := h.AnalyzeCompletionContext(d, position)
					helpers.SafeDebugLog("[HTML] AnalyzeCompletionContext returned: Type=%d, TagName=%s, AttributeName=%s",
						result.Type, result.TagName, result.AttributeName)
					return result
				} else {
					helpers.SafeDebugLog("[HTML] Handler type assertion failed for AnalyzeCompletionContext")
				}
			} else {
				helpers.SafeDebugLog("[HTML] GetLanguageHandler returned nil for html")
			}
		} else {
			helpers.SafeDebugLog("[HTML] GetLanguageHandler method not found")
		}
	} else {
		helpers.SafeDebugLog("[HTML] Invalid dm value")
	}
	// Fallback: return safe default
	helpers.SafeDebugLog("[HTML] Returning fallback CompletionUnknown")
	return &types.CompletionAnalysis{
		Type: types.CompletionUnknown,
	}
}

// FindElementAtPosition finds a custom element at the given position (interface method)
func (d *HTMLDocument) FindElementAtPosition(position protocol.Position, dm any) *types.CustomElementMatch {
	// Use reflection to call GetLanguageHandler to avoid circular imports
	dmValue := reflect.ValueOf(dm)
	if dmValue.Kind() == reflect.Pointer && !dmValue.IsNil() {
		method := dmValue.MethodByName("GetLanguageHandler")
		if method.IsValid() {
			results := method.Call([]reflect.Value{reflect.ValueOf("html")})
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
func (d *HTMLDocument) FindAttributeAtPosition(position protocol.Position, dm any) (*types.AttributeMatch, string) {
	// Use reflection to call GetLanguageHandler to avoid circular imports
	dmValue := reflect.ValueOf(dm)
	if dmValue.Kind() == reflect.Pointer && !dmValue.IsNil() {
		method := dmValue.MethodByName("GetLanguageHandler")
		if method.IsValid() {
			results := method.Call([]reflect.Value{reflect.ValueOf("html")})
			if len(results) > 0 && !results[0].IsNil() {
				handler := results[0].Interface()
				if h, ok := handler.(interface {
					FindAttributeAtPosition(types.Document, protocol.Position) (*types.AttributeMatch, string)
				}); ok {
					return h.FindAttributeAtPosition(d, position)
				}
			}
		}
	}
	// Fallback: return nil if handler not available
	return nil, ""
}

// FindCustomElements finds custom elements in the document (interface method)
func (d *HTMLDocument) FindCustomElements(dm any) ([]types.CustomElementMatch, error) {
	// Use reflection to call GetLanguageHandler to avoid circular imports
	dmValue := reflect.ValueOf(dm)
	if dmValue.Kind() == reflect.Ptr && !dmValue.IsNil() {
		method := dmValue.MethodByName("GetLanguageHandler")
		if method.IsValid() {
			results := method.Call([]reflect.Value{reflect.ValueOf("html")})
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
func (d *HTMLDocument) FindModuleScript() (protocol.Position, bool) {
	// For HTML documents, find the first module script tag
	scriptTags := d.ScriptTags()
	for _, script := range scriptTags {
		if script.IsModule && script.Src == "" {
			// Found inline module script, return end of content
			return script.ContentRange.End, true
		}
	}
	return protocol.Position{}, false
}

// FindInlineModuleScript finds insertion point in inline module script
func (d *HTMLDocument) FindInlineModuleScript() (protocol.Position, bool) {
	return d.FindModuleScript()
}

// FindHeadInsertionPoint finds insertion point in <head> section
func (d *HTMLDocument) FindHeadInsertionPoint(dm any) (protocol.Position, bool) {
	// Simple implementation - find </head> tag and insert before it
	content, err := d.Content()
	if err != nil {
		return protocol.Position{}, false
	}

	lines := strings.Split(content, "\n")
	for i, line := range lines {
		if strings.Contains(line, "</head>") {
			return protocol.Position{Line: uint32(i), Character: 0}, true
		}
	}
	return protocol.Position{}, false
}

