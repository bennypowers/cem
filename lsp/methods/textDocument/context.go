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
package textDocument

import (
	"regexp"
	"strings"

	"bennypowers.dev/cem/lsp/types"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// Pre-compiled regex patterns for performance
var (
	quotePattern = regexp.MustCompile(`(?s)[\w-]+\s*=\s*["'][^"']*$|[\w-]+\s*=\s*$`)
	attrPattern  = regexp.MustCompile(`([\w-]+)\s*=\s*["'][^"']*$|([\w-]+)\s*=\s*$`)
	valuePattern = regexp.MustCompile(`["']([^"']*)$`)
)

// CompletionContext represents what the user is trying to complete
type CompletionContextType int

const (
	CompletionUnknown             CompletionContextType = iota
	CompletionTagName                                   // After < or inside tag name
	CompletionAttributeName                             // Inside a tag, completing attribute
	CompletionAttributeValue                            // Inside attribute value quotes
	CompletionLitEventBinding                           // @event-name (lit event binding)
	CompletionLitPropertyBinding                        // .property (lit property binding)
	CompletionLitBooleanAttribute                       // ?attribute (lit boolean attribute)
)

// CompletionAnalysis holds the analysis of cursor position for completion
type CompletionAnalysis struct {
	Type          types.CompletionContextType
	TagName       string // For attribute completion, which tag we're in
	AttributeName string // For value completion, which attribute
	TriggerChar   string // What character triggered completion
	LineContent   string // Content of the current line
	IsLitTemplate bool   // True if we're in a tagged template literal (not innerHTML)
	LitSyntax     string // The Lit syntax prefix: "@", ".", or "?"
}

// AnalyzeCompletionContext determines what completion should be provided using tree-sitter
func AnalyzeCompletionContext(doc types.Document, position protocol.Position, triggerChar string) *types.CompletionAnalysis {
	return AnalyzeCompletionContextWithDM(doc, position, triggerChar, nil)
}

// AnalyzeCompletionContextWithDM determines what completion should be provided using tree-sitter with DocumentManager
func AnalyzeCompletionContextWithDM(doc types.Document, position protocol.Position, triggerChar string, dm any) *types.CompletionAnalysis {
	// Get the line content for fallback analysis
	content, err := doc.Content()
	if err != nil {
		content = ""
	}

	lines := strings.Split(content, "\n")
	if int(position.Line) >= len(lines) {
		return &types.CompletionAnalysis{Type: types.CompletionUnknown}
	}

	lineContent := lines[position.Line]

	analysis := &types.CompletionAnalysis{
		TriggerChar: triggerChar,
		LineContent: lineContent,
	}

	// Try to use tree-sitter based analysis if document supports it
	if docWithTreeSitter, ok := doc.(interface {
		AnalyzeCompletionContextTS(position protocol.Position, dm any) *types.CompletionAnalysis
	}); ok {
		if tsAnalysis := docWithTreeSitter.AnalyzeCompletionContextTS(position, dm); tsAnalysis != nil {
			// Convert from types.CompletionAnalysis to types.CompletionAnalysis
			return &types.CompletionAnalysis{
				Type:          types.CompletionContextType(tsAnalysis.Type),
				TagName:       tsAnalysis.TagName,
				AttributeName: tsAnalysis.AttributeName,
				TriggerChar:   triggerChar,
				LineContent:   lineContent,
				IsLitTemplate: tsAnalysis.IsLitTemplate,
				LitSyntax:     tsAnalysis.LitSyntax,
			}
		}
	}

	// Fallback to text-based analysis
	charPos := int(position.Character)
	beforeCursor := ""
	if charPos <= len(lineContent) {
		beforeCursor = lineContent[:charPos]
	}

	// Check for template literal context first using document's tree-sitter analysis
	if doc != nil {
		if templateType := getTemplateContext(doc, position); templateType != "" {
			analysis.IsLitTemplate = (templateType == "html") // True for html`` but not innerHTML
			return analyzeTemplateContext(beforeCursor, analysis)
		}
	}

	// Otherwise analyze as HTML
	return analyzeHTMLContext(beforeCursor, analysis)
}

// getTemplateContext checks if cursor is inside a template literal and returns the type
// Returns: "html" for html“ templates, "innerHTML" for innerHTML assignments, "outerHTML" for outerHTML, "" if not in template
func getTemplateContext(doc types.Document, position protocol.Position) string {
	// Try to use the document's built-in template detection if available
	if docWithTemplates, ok := doc.(interface {
		GetTemplateContext(position protocol.Position) string
	}); ok {
		return docWithTemplates.GetTemplateContext(position)
	}

	// Fallback: Use simple heuristic based on content analysis
	content, err := doc.Content()
	if err != nil {
		content = ""
	}

	// If this looks like a TypeScript file, use template literal detection
	lines := strings.Split(content, "\n")
	if int(position.Line) >= len(lines) {
		return ""
	}

	// Look backwards from cursor position to find template literal start
	for i := int(position.Line); i >= 0; i-- {
		line := lines[i]
		if i == int(position.Line) {
			// For current line, only check up to cursor position
			if int(position.Character) < len(line) {
				line = line[:position.Character]
			}
		}

		// Look for template literal patterns (simplified heuristic)
		if strings.Contains(line, "html`") {
			// Check if we're likely inside a template literal
			if strings.Count(line, "`")%2 == 1 {
				return "html"
			}
		}
		if strings.Contains(line, "innerHTML") && strings.Contains(line, "`") {
			if strings.Count(line, "`")%2 == 1 {
				return "innerHTML"
			}
		}
		if strings.Contains(line, "outerHTML") && strings.Contains(line, "`") {
			if strings.Count(line, "`")%2 == 1 {
				return "outerHTML"
			}
		}

		// If we find a complete template literal or statement end, stop looking
		if strings.Contains(line, ";") || strings.Contains(line, "}") {
			break
		}
	}

	return ""
}

// analyzeLitSyntax detects Lit-specific attribute syntax (@, ., ?)
func analyzeLitSyntax(tagContent string) *types.CompletionAnalysis {
	// Look for Lit syntax patterns: @event, .property, ?boolean
	// Pattern: <tag-name @|.|?partial-name

	tagName := extractTagName(tagContent)
	if tagName == "" {
		return nil
	}

	// Remove tag name and < to focus on attributes
	afterTag := strings.TrimPrefix(tagContent, "<"+tagName)

	// Look for the last occurrence of Lit syntax
	litPrefixes := []string{"@", ".", "?"}
	var lastPrefix string
	var lastIndex int = -1

	for _, prefix := range litPrefixes {
		if idx := strings.LastIndex(afterTag, prefix); idx > lastIndex {
			lastIndex = idx
			lastPrefix = prefix
		}
	}

	if lastIndex == -1 {
		return nil
	}

	// Check if we're immediately after the prefix or completing the attribute name
	afterPrefix := afterTag[lastIndex+1:]

	// If there's a space or = after the prefix content, we're not completing this attribute
	if strings.Contains(afterPrefix, " ") || strings.Contains(afterPrefix, "=") {
		return nil
	}

	analysis := &types.CompletionAnalysis{
		TagName:       tagName,
		LitSyntax:     lastPrefix,
		IsLitTemplate: true,
	}

	switch lastPrefix {
	case "@":
		analysis.Type = types.CompletionLitEventBinding
	case ".":
		analysis.Type = types.CompletionLitPropertyBinding
	case "?":
		analysis.Type = types.CompletionLitBooleanAttribute
	}

	return analysis
}

// analyzeTemplateContext analyzes completion context within template literals
func analyzeTemplateContext(beforeCursor string, analysis *types.CompletionAnalysis) *types.CompletionAnalysis {
	// Same logic as HTML since template literals contain HTML
	return analyzeHTMLContext(beforeCursor, analysis)
}

// analyzeHTMLContext analyzes completion context in HTML content
func analyzeHTMLContext(beforeCursor string, analysis *types.CompletionAnalysis) *types.CompletionAnalysis {

	// Remove any content before the last < to focus on current tag
	lastLT := strings.LastIndex(beforeCursor, "<")
	if lastLT == -1 {
		analysis.Type = types.CompletionUnknown
		return analysis
	}

	// Get content since the last <
	tagContent := beforeCursor[lastLT:]

	// Check for Lit-specific syntax if we're in a tagged template literal
	if analysis.IsLitTemplate {
		if litContext := analyzeLitSyntax(tagContent); litContext != nil {
			return litContext
		}
	}

	// Check if we're in an attribute value (inside quotes)
	if isInAttributeValue(tagContent) {
		analysis.Type = types.CompletionAttributeValue
		analysis.TagName, analysis.AttributeName = extractTagAndAttribute(tagContent)
		return analysis
	}

	// Check if we're completing an attribute name
	if isCompletingAttributeName(tagContent) {
		analysis.Type = types.CompletionAttributeName
		analysis.TagName = extractTagName(tagContent)
		return analysis
	}

	// Check if we're completing a tag name
	if isCompletingTagName(tagContent) {
		analysis.Type = types.CompletionTagName
		return analysis
	}

	analysis.Type = types.CompletionUnknown
	return analysis
}

// isInAttributeValue checks if cursor is inside attribute value quotes
func isInAttributeValue(tagContent string) bool {
	// Look for pattern: attr="value| or attr='value| or attr=|
	// Updated pattern to handle tag names with hyphens and attributes with hyphens
	// This matches when we have an attribute followed by = (cursor right after = or inside quotes)
	// Updated to handle multi-line cases with (?s) flag to make . match newlines
	return quotePattern.MatchString(tagContent)
}

// isCompletingAttributeName checks if cursor is positioned to complete attribute name
func isCompletingAttributeName(tagContent string) bool {
	// Must have a tag name and be after whitespace
	// Pattern: <tag-name space|
	// Or: <tag-name attr="value" space|

	// Extract tag name first
	tagName := extractTagName(tagContent)
	if tagName == "" {
		return false
	}

	// Remove the tag name and < from consideration
	afterTag := strings.TrimPrefix(tagContent, "<"+tagName)

	// If there's any content after tag name, we could be completing attributes
	// Check that we're not inside quotes and there's whitespace before cursor
	if len(strings.TrimSpace(afterTag)) > 0 || strings.HasSuffix(tagContent, " ") {
		return !isInAttributeValue(tagContent)
	}

	return false
}

// isCompletingTagName checks if cursor is positioned to complete tag name
func isCompletingTagName(tagContent string) bool {
	// Pattern: <| or <some-partial-name|
	// Must not contain spaces (which would indicate we're past tag name)
	return !strings.Contains(tagContent[1:], " ") && !strings.Contains(tagContent, "=")
}

// extractTagAndAttribute extracts both tag name and attribute name from tag content
func extractTagAndAttribute(tagContent string) (string, string) {
	tagName := extractTagName(tagContent)

	// Find the last attribute name before the cursor
	// Pattern: attr="value or attr='value or attr=
	// Updated pattern to handle attributes with hyphens and missing quotes
	matches := attrPattern.FindStringSubmatch(tagContent)
	if len(matches) >= 2 {
		var attrName string
		// First capture group (with quotes)
		if matches[1] != "" {
			attrName = matches[1]
		} else if matches[2] != "" {
			// Second capture group (without quotes)
			attrName = matches[2]
		}

		// Special case for slot attribute: allow any tag name, not just custom elements
		if attrName == "slot" {
			// Extract any tag name (not just custom elements) for slot attributes
			tagName = extractAnyTagName(tagContent)
		}

		return tagName, attrName
	}

	return tagName, ""
}

// extractAnyTagName extracts any tag name, including standard HTML elements
func extractAnyTagName(tagContent string) string {
	// Remove < and extract first word
	content := strings.TrimPrefix(tagContent, "<")
	parts := strings.Fields(content)
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}

// extractTagName extracts the tag name from tag content
func extractTagName(tagContent string) string {
	// Remove < and extract first word
	content := strings.TrimPrefix(tagContent, "<")
	parts := strings.Fields(content)
	if len(parts) > 0 {
		// Only return if it looks like a custom element (contains hyphen)
		tagName := parts[0]
		if strings.Contains(tagName, "-") {
			return tagName
		}
	}
	return ""
}

// IsCustomElementTag checks if a tag name is a custom element
func IsCustomElementTag(tagName string) bool {
	// Custom elements must contain a hyphen and start with lowercase letter
	if len(tagName) == 0 || !strings.Contains(tagName, "-") {
		return false
	}

	// Must start with lowercase letter
	firstChar := tagName[0]
	return firstChar >= 'a' && firstChar <= 'z'
}

// GetCompletionPrefix extracts the prefix being typed for filtering
func GetCompletionPrefix(beforeCursor string, analysis *types.CompletionAnalysis) string {
	switch analysis.Type {
	case types.CompletionTagName:
		// Extract partial tag name after <
		if idx := strings.LastIndex(beforeCursor, "<"); idx != -1 {
			tagPart := beforeCursor[idx+1:]
			// Remove any attributes that might have been parsed
			if spaceIdx := strings.Index(tagPart, " "); spaceIdx != -1 {
				tagPart = tagPart[:spaceIdx]
			}
			return tagPart
		}

	case types.CompletionAttributeName:
		// Extract partial attribute name
		// Look for the last word that's not followed by =
		parts := strings.Fields(beforeCursor)
		if len(parts) > 0 {
			lastPart := parts[len(parts)-1]
			// If it doesn't contain =, it's likely a partial attribute name
			if !strings.Contains(lastPart, "=") && !strings.HasPrefix(lastPart, "<") {
				return lastPart
			}
		}

	case types.CompletionAttributeValue:
		// Extract partial value inside quotes
		matches := valuePattern.FindStringSubmatch(beforeCursor)
		if len(matches) >= 2 {
			return matches[1]
		}
	}

	return ""
}
