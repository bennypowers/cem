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
package tsx

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"bennypowers.dev/cem/lsp/types"
	Q "bennypowers.dev/cem/queries"
	protocol "github.com/tliron/glsp/protocol_3_16"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// Completion scoring constants for prioritizing tree-sitter captures
const (
	// Base scores for different completion types (lower scores are more specific)
	SCORE_TAG_NAME_BASE       = 10 // Most specific: tag name completion
	SCORE_ATTR_NAME_BASE      = 20 // Medium specificity: attribute name completion
	SCORE_ATTR_VALUE_BASE     = 30 // Least specific: attribute value completion

	// Bonus/penalty scoring factors
	SCORE_CURSOR_BONUS        = 50 // Large bonus for captures containing cursor
	SCORE_BROAD_CAPTURE_PENALTY = 5  // Penalty multiplier for broad captures (ERROR nodes)

	// Thresholds for completion analysis
	MAX_REASONABLE_TAG_LENGTH = 10 // Threshold for reasonable tag name length
)

// TSXDocument represents a TSX document with tree-sitter parsing
type TSXDocument struct {
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
func (d *TSXDocument) URI() string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.uri
}

// Content returns the document content
func (d *TSXDocument) Content() (string, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.content, nil
}

// Version returns the document version
func (d *TSXDocument) Version() int32 {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.version
}

// Language returns the document language
func (d *TSXDocument) Language() string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.language
}

// Tree returns the document's syntax tree
func (d *TSXDocument) Tree() *ts.Tree {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.tree
}

// ScriptTags returns the parsed script tags
func (d *TSXDocument) ScriptTags() []types.ScriptTag {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.scriptTags
}

// UpdateContent updates the document content
func (d *TSXDocument) UpdateContent(content string, version int32) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.content = content
	d.version = version
}

// SetTree sets the document's syntax tree
func (d *TSXDocument) SetTree(tree *ts.Tree) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.tree != nil {
		d.tree.Close()
	}
	d.tree = tree
}

// Parser returns the document's parser
func (d *TSXDocument) Parser() *ts.Parser {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.parser
}

// SetParser sets the document's parser
func (d *TSXDocument) SetParser(parser *ts.Parser) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.parser = parser
}

// SetScriptTags sets the script tags
func (d *TSXDocument) SetScriptTags(scriptTags []types.ScriptTag) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.scriptTags = scriptTags
}

// Close cleans up document resources
func (d *TSXDocument) Close() {
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
func (d *TSXDocument) ByteRangeToProtocolRange(content string, startByte, endByte uint) protocol.Range {
	return protocol.Range{
		Start: d.byteOffsetToPosition(startByte),
		End:   d.byteOffsetToPosition(endByte),
	}
}

// Helper method for byte offset to position conversion
func (d *TSXDocument) byteOffsetToPosition(offset uint) protocol.Position {
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

// Parse parses the TSX content using tree-sitter
func (d *TSXDocument) Parse(content string) error {
	d.UpdateContent(content, d.version)

	parser := Q.GetTSXParser()
	if parser == nil {
		return fmt.Errorf("failed to get TSX parser")
	}
	d.SetParser(parser)

	tree := parser.Parse([]byte(content), nil)
	d.SetTree(tree)

	return nil
}

// findCustomElements finds custom elements in TSX documents
func (d *TSXDocument) findCustomElements(handler *Handler) ([]types.CustomElementMatch, error) {
	tree := d.Tree()
	if tree == nil {
		return nil, fmt.Errorf("no tree available for document")
	}

	content, err := d.Content()
	if err != nil {
		return nil, err
	}

	var elements []types.CustomElementMatch
	// Use a map to deduplicate elements based on position
	elementMap := make(map[string]types.CustomElementMatch)

	// Create a fresh query matcher for thread safety
	customElementsMatcher, err := Q.GetCachedQueryMatcher(handler.queryManager, "tsx", "customElements")
	if err != nil {
		return nil, fmt.Errorf("failed to create custom elements matcher: %w", err)
	}
	defer customElementsMatcher.Close()

	// Handle multiple parent capture types: element, self.closing.tag, start.tag, incomplete.element
	parentCaptureNames := []string{"element", "self.closing.tag", "start.tag", "incomplete.element"}

	for _, parentName := range parentCaptureNames {
		for captureMap := range customElementsMatcher.ParentCaptures(tree.RootNode(), []byte(content), parentName) {
			if tagNames, exists := captureMap["tag.name"]; exists {
				for _, tagCapture := range tagNames {
					if strings.Contains(tagCapture.Text, "-") {
						// Create a unique key based on tag name and position
						elementKey := fmt.Sprintf("%s:%d:%d", tagCapture.Text, tagCapture.StartByte, tagCapture.EndByte)

						// Skip if we've already processed this element
						if _, exists := elementMap[elementKey]; exists {
							continue
						}

						element := types.CustomElementMatch{
							TagName:    tagCapture.Text,
							Range:      d.ByteRangeToProtocolRange(content, tagCapture.StartByte, tagCapture.EndByte),
							Attributes: make(map[string]types.AttributeMatch),
						}

						// Look for attributes in the same capture map
						if attrNames, hasAttrs := captureMap["attr.name"]; hasAttrs {
							for _, attrCapture := range attrNames {
								attr := types.AttributeMatch{
									Name:  attrCapture.Text,
									Range: d.ByteRangeToProtocolRange(content, attrCapture.StartByte, attrCapture.EndByte),
								}

								// Look for attribute value
								if attrValues, hasValues := captureMap["attr.value"]; hasValues {
									for _, valueCapture := range attrValues {
										// Simple heuristic: if value is near the attribute name, associate it
										if valueCapture.StartByte > attrCapture.EndByte &&
											valueCapture.StartByte < attrCapture.EndByte+100 { // Within 100 chars
											attr.Value = valueCapture.Text
											break
										}
									}
								}

								element.Attributes[attr.Name] = attr
							}
						}

						elementMap[elementKey] = element
					}
				}
			}
		}
	}

	// Convert map back to slice
	for _, element := range elementMap {
		elements = append(elements, element)
	}

	return elements, nil
}

// analyzeCompletionContext analyzes completion context for TSX documents
func (d *TSXDocument) analyzeCompletionContext(position protocol.Position, handler *Handler) *types.CompletionAnalysis {
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
	completionMatcher, err := Q.GetCachedQueryMatcher(handler.queryManager, "tsx", "completionContext")
	if err != nil {
		return analysis
	}
	defer completionMatcher.Close()

	// Use tree-sitter to analyze context
	// Collect all matching captures first, then prioritize
	var possibleMatches []struct {
		captureType string
		capture     Q.CaptureInfo
		captureMap  Q.CaptureMap
	}

	for captureMap := range completionMatcher.ParentCaptures(tree.RootNode(), []byte(content), "context") {
		// Find which contexts the cursor is in by checking byte offsets
		for captureName, captures := range captureMap {
			// Skip generic "context" captures, only score specific completion types
			if captureName == "context" {
				continue
			}
			for _, capture := range captures {
				// Include captures that contain the cursor
				if capture.StartByte <= byteOffset && byteOffset <= capture.EndByte {
					possibleMatches = append(possibleMatches, struct {
						captureType string
						capture     Q.CaptureInfo
						captureMap  Q.CaptureMap
					}{captureName, capture, captureMap})
				} else if captureName == "attr.name.completion" && capture.EndByte <= byteOffset && byteOffset-capture.EndByte <= 10 {
					// Include attribute name captures that end just before the cursor (for attribute value completion)
					possibleMatches = append(possibleMatches, struct {
						captureType string
						capture     Q.CaptureInfo
						captureMap  Q.CaptureMap
					}{captureName, capture, captureMap})
				}
			}
		}
	}

	if len(possibleMatches) == 0 {
		// Check for the special case where cursor is right after a custom element tag
		// Look for captures that end just before the cursor position
		for captureMap := range completionMatcher.ParentCaptures(tree.RootNode(), []byte(content), "context") {
			for captureName, captures := range captureMap {
				for _, capture := range captures {
					// If cursor is 1-2 positions after a tag name completion capture
					if captureName == "tag.name.completion" &&
						capture.EndByte < byteOffset &&
						byteOffset-capture.EndByte <= 2 {
						// Check if the capture text contains a custom element
						if strings.Contains(capture.Text, "-") {
							// This is likely attribute name completion context
							analysis.Type = types.CompletionAttributeName
							analysis.TagName = extractCustomElementFromText(capture.Text)
							return analysis
						}
					}
				}
			}
		}
		return analysis
	}

	// Prioritize the most specific match that has enough context
	// Priority order: attr.value.completion > attr.name.completion > tag.name.completion
	// But also prefer matches that include tag/attribute context
	var bestMatch *struct {
		captureType string
		capture     Q.CaptureInfo
		captureMap  Q.CaptureMap
	}

	for i, match := range possibleMatches {
		if bestMatch == nil {
			bestMatch = &possibleMatches[i]
			continue
		}

		// Calculate a score for this match based on specificity and context
		currentScore := scoreMatch(match.captureType, match.captureMap, match.capture, byteOffset)
		bestScore := scoreMatch(bestMatch.captureType, bestMatch.captureMap, bestMatch.capture, byteOffset)

		if currentScore < bestScore { // Lower scores are better (more specific)
			bestMatch = &possibleMatches[i]
		}
	}

	if bestMatch != nil {
		switch bestMatch.captureType {
		case "tag.name.completion":
			analysis.Type = types.CompletionTagName
			analysis.TagName = bestMatch.capture.Text
		case "attr.name.completion":
			analysis.Type = types.CompletionAttributeName
			analysis.AttributeName = bestMatch.capture.Text
			// Try to find the tag name for this attribute
			analysis.TagName = d.findTagNameForAttribute(byteOffset, bestMatch.captureMap)
		case "attr.value.completion":
			analysis.Type = types.CompletionAttributeValue
			// Look across all captures to find the best tag and attribute names
			tagName := d.findBestTagName(byteOffset, possibleMatches)
			attrName := d.findBestAttributeName(byteOffset, possibleMatches)
			analysis.TagName = tagName
			analysis.AttributeName = attrName
		case "interpolation":
			// Inside TypeScript interpolation - no JSX completions
			analysis.Type = types.CompletionUnknown
			return analysis
		}
	}

	return analysis
}

// findTagNameForAttribute finds the tag name associated with an attribute at the given position
func (d *TSXDocument) findTagNameForAttribute(byteOffset uint, captureMap Q.CaptureMap) string {
	// Look for tag name captures that are before the current position
	if tagNames, exists := captureMap["tag.name.completion"]; exists {
		var closestTag *Q.CaptureInfo
		var closestDistance = ^uint(0) // Max uint value

		for _, tagCapture := range tagNames {
			// For attribute completions, we want tag captures that contain the cursor OR end just before it
			if (tagCapture.StartByte <= byteOffset && byteOffset <= tagCapture.EndByte) ||
				(tagCapture.EndByte <= byteOffset) {

				var distance uint
				if byteOffset <= tagCapture.EndByte {
					// Cursor is inside the tag capture - prioritize this
					distance = 0
				} else {
					distance = byteOffset - tagCapture.EndByte
				}

				if distance < closestDistance {
					closestDistance = distance
					closestTag = &tagCapture
				}
			}
		}

		if closestTag != nil {
			// Extract the actual tag name from the capture text
			tagName := extractCustomElementFromText(closestTag.Text)
			if tagName == "" {
				// Fallback: if extraction fails, try to use the text directly if it looks like a tag name
				if strings.Contains(closestTag.Text, "-") && !strings.Contains(closestTag.Text, " ") {
					tagName = closestTag.Text
				}
			}
			return tagName
		}
	}

	return ""
}

// findBestTagName finds the best tag name from all possible matches
func (d *TSXDocument) findBestTagName(byteOffset uint, possibleMatches []struct {
	captureType string
	capture     Q.CaptureInfo
	captureMap  Q.CaptureMap
}) string {
	var bestTagName string
	var bestDistance = ^uint(0) // Max uint value

	// Look through all captures for tag names
	for _, match := range possibleMatches {
		if tagNames, exists := match.captureMap["tag.name.completion"]; exists {
			for _, tagCapture := range tagNames {
				// Prefer captures that contain or are close to the cursor
				var distance uint
				if tagCapture.StartByte <= byteOffset && byteOffset <= tagCapture.EndByte {
					distance = 0 // Inside the capture - perfect
				} else if tagCapture.EndByte <= byteOffset {
					distance = byteOffset - tagCapture.EndByte
				} else {
					continue // Skip captures that start after the cursor
				}

				if distance < bestDistance {
					tagName := extractCustomElementFromText(tagCapture.Text)
					if tagName != "" {
						bestTagName = tagName
						bestDistance = distance
					}
				}
			}
		}
	}

	return bestTagName
}

// findBestAttributeName finds the best attribute name from all possible matches
func (d *TSXDocument) findBestAttributeName(byteOffset uint, possibleMatches []struct {
	captureType string
	capture     Q.CaptureInfo
	captureMap  Q.CaptureMap
}) string {
	var bestAttrName string
	var bestDistance = ^uint(0) // Max uint value

	// Look through all captures for attribute names across ALL capture maps
	for _, match := range possibleMatches {
		// Check for attr.name.completion captures
		if attrNames, exists := match.captureMap["attr.name.completion"]; exists {
			for _, attrCapture := range attrNames {
				// Prefer captures that are close to but before the cursor
				if attrCapture.EndByte <= byteOffset {
					distance := byteOffset - attrCapture.EndByte
					if distance < bestDistance {
						attrName := extractAttributeNameFromText(attrCapture.Text)
						if attrName != "" {
							bestAttrName = attrName
							bestDistance = distance
						}
					}
				}
			}
		}

		// Also check for any capture that might be an attribute name completion context
		if match.captureType == "attr.name.completion" {
			if match.capture.EndByte <= byteOffset {
				distance := byteOffset - match.capture.EndByte
				if distance < bestDistance {
					attrName := extractAttributeNameFromText(match.capture.Text)
					if attrName != "" {
						bestAttrName = attrName
						bestDistance = distance
					}
				}
			}
		}
	}

	return bestAttrName
}

// positionToByteOffset converts LSP position to byte offset
func (d *TSXDocument) positionToByteOffset(pos protocol.Position, content string) uint {
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

// extractCustomElementFromText extracts the custom element name from capture text
func extractCustomElementFromText(text string) string {
	// Look for the last occurrence of a custom element pattern
	parts := strings.Split(text, "<")
	if len(parts) < 2 {
		return ""
	}

	// Get the part after the last "<"
	lastPart := parts[len(parts)-1]

	// Extract the tag name (up to first space or >)
	tagName := lastPart
	if spaceIdx := strings.Index(tagName, " "); spaceIdx != -1 {
		tagName = tagName[:spaceIdx]
	}
	if gtIdx := strings.Index(tagName, ">"); gtIdx != -1 {
		tagName = tagName[:gtIdx]
	}

	// Only return if it contains a hyphen (custom element)
	if strings.Contains(tagName, "-") {
		return tagName
	}

	return ""
}

// extractAttributeNameFromText extracts the attribute name from capture text
func extractAttributeNameFromText(text string) string {
	// If it's already a simple attribute name, return it
	if !strings.Contains(text, " ") && !strings.Contains(text, "<") && !strings.Contains(text, "=") {
		return text
	}

	// Look for the last word that looks like an attribute name
	words := strings.Fields(text)
	for i := len(words) - 1; i >= 0; i-- {
		word := words[i]
		// Remove trailing = if present
		word = strings.TrimSuffix(word, "=")
		// Attribute names should be simple identifiers
		if len(word) > 0 && !strings.Contains(word, "<") && !strings.Contains(word, ">") {
			return word
		}
	}

	return ""
}

// scoreMatch scores a completion match based on type specificity and capture precision.
//
// Scoring Strategy:
// - Lower scores are better (more specific matches win)
// - Tag name completion gets the lowest base score (most specific)
// - Attribute value completion gets the highest base score (least specific)
// - Large bonus for captures that contain the cursor position
// - Penalties for overly broad captures (ERROR nodes) and long captures
func scoreMatch(captureType string, captureMap Q.CaptureMap, capture Q.CaptureInfo, byteOffset uint) int {
	score := 0

	// Base score for match type (lower scores win for more specific captures)
	switch captureType {
	case "tag.name.completion":
		score += SCORE_TAG_NAME_BASE // Most specific completion type
	case "attr.name.completion":
		score += SCORE_ATTR_NAME_BASE // Medium specificity
	case "attr.value.completion":
		score += SCORE_ATTR_VALUE_BASE // Least specific (most complex context)
	}

	// Large bonus for captures that contain the cursor (lower score = better)
	// This ensures that captures containing the cursor are strongly preferred
	if capture.StartByte <= byteOffset && byteOffset <= capture.EndByte {
		score -= SCORE_CURSOR_BONUS // Significant bonus for containing cursor
	}

	// Penalty for overly broad capture maps (ERROR nodes often have many captures)
	// More captures usually means less precision in tree-sitter parsing
	if len(captureMap) > 2 {
		score += len(captureMap) * SCORE_BROAD_CAPTURE_PENALTY // Penalize broad captures
	}

	// Major penalty for overly long captures (likely ERROR node captures)
	// Reasonable tag/attribute names are usually short
	captureLength := capture.EndByte - capture.StartByte
	if captureLength > MAX_REASONABLE_TAG_LENGTH {
		score += int(captureLength) // Penalize based on length
	}

	return score
}

// AnalyzeCompletionContextTS analyzes completion context using tree-sitter queries (interface method)
func (d *TSXDocument) AnalyzeCompletionContextTS(position protocol.Position, dm any) *types.CompletionAnalysis {
	// Use reflection to call GetLanguageHandler to avoid circular imports
	dmValue := reflect.ValueOf(dm)
	if dmValue.Kind() == reflect.Pointer && !dmValue.IsNil() {
		method := dmValue.MethodByName("GetLanguageHandler")
		if method.IsValid() {
			results := method.Call([]reflect.Value{reflect.ValueOf("tsx")})
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

// CompletionPrefix extracts the prefix being typed for filtering completions
func (d *TSXDocument) CompletionPrefix(analysis *types.CompletionAnalysis) string {
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

// FindElementAtPosition finds a custom element at the given position (interface method)
func (d *TSXDocument) FindElementAtPosition(position protocol.Position, dm any) *types.CustomElementMatch {
	// Use reflection to call GetLanguageHandler to avoid circular imports
	dmValue := reflect.ValueOf(dm)
	if dmValue.Kind() == reflect.Pointer && !dmValue.IsNil() {
		method := dmValue.MethodByName("GetLanguageHandler")
		if method.IsValid() {
			results := method.Call([]reflect.Value{reflect.ValueOf("tsx")})
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
func (d *TSXDocument) FindAttributeAtPosition(position protocol.Position, dm any) (*types.AttributeMatch, string) {
	// This method needs to be implemented for the types.Document interface
	// For now, return nil as this functionality is handled by the handler
	return nil, ""
}

// FindCustomElements finds custom elements in the document (interface method)
func (d *TSXDocument) FindCustomElements(dm any) ([]types.CustomElementMatch, error) {
	// Use reflection to call GetLanguageHandler to avoid circular imports
	dmValue := reflect.ValueOf(dm)
	if dmValue.Kind() == reflect.Pointer && !dmValue.IsNil() {
		method := dmValue.MethodByName("GetLanguageHandler")
		if method.IsValid() {
			results := method.Call([]reflect.Value{reflect.ValueOf("tsx")})
			if len(results) > 0 && !results[0].IsNil() {
				handler := results[0].Interface()
				if h, ok := handler.(interface {
					FindCustomElements(types.Document) ([]types.CustomElementMatch, error)
				}); ok {
					return h.FindCustomElements(d)
				}
			}
		}
	}
	// Fallback: return empty slice if handler not available
	return []types.CustomElementMatch{}, nil
}

// FindModuleScript finds insertion point in module script
func (d *TSXDocument) FindModuleScript() (protocol.Position, bool) {
	// For TSX files, find the end of the file content
	content, err := d.Content()
	if err != nil {
		return protocol.Position{}, false
	}

	lines := strings.Split(content, "\n")
	return protocol.Position{Line: uint32(len(lines)), Character: 0}, true
}

// FindInlineModuleScript finds insertion point in inline module script
func (d *TSXDocument) FindInlineModuleScript() (protocol.Position, bool) {
	return d.FindModuleScript()
}

// FindHeadInsertionPoint finds insertion point in <head> section
func (d *TSXDocument) FindHeadInsertionPoint(dm any) (protocol.Position, bool) {
	// TSX files don't have <head> sections
	return protocol.Position{}, false
}
