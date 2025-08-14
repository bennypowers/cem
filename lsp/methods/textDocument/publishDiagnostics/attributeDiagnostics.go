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
package publishDiagnostics

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

//go:embed data/global_attributes.json
var globalAttributesJSON []byte

// MDNCompatData represents the structure of MDN browser compatibility data
type MDNCompatData struct {
	HTML struct {
		GlobalAttributes map[string]interface{} `json:"global_attributes"`
	} `json:"html"`
}

var globalAttributes map[string]bool

// Initialize global attributes from embedded MDN data
func init() {
	globalAttributes = make(map[string]bool)

	var mdnData MDNCompatData
	if err := json.Unmarshal(globalAttributesJSON, &mdnData); err == nil {
		for attrName := range mdnData.HTML.GlobalAttributes {
			// Convert data_attributes to data-* pattern
			if attrName == "data_attributes" {
				// We'll handle data-* pattern separately
				continue
			}
			globalAttributes[attrName] = true
		}

		helpers.SafeDebugLog("[DIAGNOSTICS] Loaded %d global HTML attributes from MDN data", len(globalAttributes))
	} else {
		helpers.SafeDebugLog("[DIAGNOSTICS] Failed to load MDN global attributes data: %v", err)
	}
}

// AttributeMatch represents a found attribute in the document
type AttributeMatch struct {
	Name     string
	TagName  string // The tag this attribute belongs to
	Line     uint32
	StartCol uint32
	EndCol   uint32
}

// analyzeAttributeDiagnostics finds unknown attributes and suggests corrections
func analyzeAttributeDiagnostics(ctx DiagnosticsContext, doc types.Document) []protocol.Diagnostic {
	return AnalyzeAttributeDiagnosticsForTest(ctx, doc)
}

// AnalyzeAttributeDiagnosticsForTest is the exported version for testing
func AnalyzeAttributeDiagnosticsForTest(ctx DiagnosticsContext, doc types.Document) []protocol.Diagnostic {
	var diagnostics []protocol.Diagnostic

	// Get document content to search for attributes
	content := ""
	if docWithContent, ok := doc.(interface{ Content() string }); ok {
		content = docWithContent.Content()
	} else {
		return diagnostics
	}

	helpers.SafeDebugLog("[DIAGNOSTICS] Analyzing attributes in document (length=%d)", len(content))

	// Find all attributes in the document
	attributeMatches := findAttributes(content)
	helpers.SafeDebugLog("[DIAGNOSTICS] Found %d attributes", len(attributeMatches))

	for _, match := range attributeMatches {
		// Skip if it's a global HTML attribute
		if isGlobalAttribute(match.Name) {
			continue
		}

		// For custom elements, check if the attribute is defined in the manifest
		if isCustomElement(match.TagName) {
			// Get attributes for this custom element
			if attrs := getCustomElementAttributes(ctx, match.TagName); attrs != nil {
				// Check if this attribute is defined for this custom element
				found := false
				for _, attr := range attrs {
					if attr.Name == match.Name {
						found = true
						break
					}
				}
				if found {
					continue // Valid custom element attribute
				}

				// Suggest corrections for custom element attributes
				suggestion := findClosestAttribute(match.Name, attrs)
				if suggestion != "" {
					diagnostics = append(diagnostics, createAttributeDiagnostic(match, suggestion))
				} else {
					diagnostics = append(diagnostics, createUnknownAttributeDiagnostic(match))
				}
			} else {
				// No manifest found for this custom element, warn about unknown attribute
				diagnostics = append(diagnostics, createUnknownAttributeDiagnostic(match))
			}
		} else {
			// For standard HTML elements, only suggest if it might be a typo of a global attribute
			suggestion := findClosestGlobalAttribute(match.Name)
			if suggestion != "" {
				diagnostics = append(diagnostics, createAttributeDiagnostic(match, suggestion))
			}
			// Don't warn about unknown attributes on standard elements -
			// there are too many element-specific attributes to track
		}
	}

	helpers.SafeDebugLog("[DIAGNOSTICS] Generated %d attribute diagnostics", len(diagnostics))
	return diagnostics
}

// isGlobalAttribute checks if an attribute is a valid global HTML attribute
func isGlobalAttribute(name string) bool {
	nameLower := strings.ToLower(name)

	// Check if it's in our global attributes list
	if globalAttributes[nameLower] {
		return true
	}

	// Check for data-* attributes (always valid)
	if strings.HasPrefix(nameLower, "data-") {
		return true
	}

	// Check for aria-* attributes (always valid)
	if strings.HasPrefix(nameLower, "aria-") {
		return true
	}

	// Check for event handler attributes starting with "on"
	if strings.HasPrefix(nameLower, "on") && len(nameLower) > 2 {
		return true
	}

	return false
}

// findAttributes finds all attributes in the document content
func findAttributes(content string) []AttributeMatch {
	var matches []AttributeMatch
	lines := strings.Split(content, "\n")

	for lineIdx, line := range lines {
		// Look for attributes in opening tags
		idx := 0
		for {
			tagStart := strings.Index(line[idx:], "<")
			if tagStart == -1 {
				break
			}
			tagStart += idx

			// Skip closing tags and comments
			if tagStart+1 < len(line) && (line[tagStart+1] == '/' || line[tagStart+1] == '!') {
				idx = tagStart + 1
				continue
			}

			// Find the tag name first
			tagNameEnd := strings.IndexAny(line[tagStart+1:], " \t\n\r/>")
			if tagNameEnd == -1 {
				idx = tagStart + 1
				continue
			}
			tagName := line[tagStart+1 : tagStart+1+tagNameEnd]

			// Find the end of the tag
			tagEnd := strings.Index(line[tagStart:], ">")
			if tagEnd == -1 {
				idx = tagStart + 1
				continue
			}
			tagEnd += tagStart

			// Extract the attribute section
			attrSection := line[tagStart+1+tagNameEnd : tagEnd]

			// Find attributes in this section
			attrMatches := findAttributesInSection(attrSection, tagName, uint32(lineIdx), uint32(tagStart+1+tagNameEnd))
			matches = append(matches, attrMatches...)

			idx = tagEnd + 1
		}
	}

	return matches
}

// findAttributesInSection finds attributes within a tag's attribute section
func findAttributesInSection(section string, tagName string, line uint32, startOffset uint32) []AttributeMatch {
	var matches []AttributeMatch

	// Simple regex-like parsing for attribute names
	i := 0
	for i < len(section) {
		// Skip whitespace
		for i < len(section) && (section[i] == ' ' || section[i] == '\t' || section[i] == '\n' || section[i] == '\r') {
			i++
		}
		if i >= len(section) {
			break
		}

		// Find attribute name (ends with =, whitespace, or end of section)
		nameStart := i
		for i < len(section) && section[i] != '=' && section[i] != ' ' && section[i] != '\t' && section[i] != '\n' && section[i] != '\r' {
			i++
		}

		if i > nameStart {
			attrName := section[nameStart:i]

			// Skip if this looks like a value rather than an attribute name
			if strings.Contains(attrName, `"`) || strings.Contains(attrName, `'`) {
				continue
			}

			matches = append(matches, AttributeMatch{
				Name:     attrName,
				TagName:  tagName,
				Line:     line,
				StartCol: startOffset + uint32(nameStart),
				EndCol:   startOffset + uint32(i),
			})
		}

		// Skip the value if present
		if i < len(section) && section[i] == '=' {
			i++ // skip =
			// Skip whitespace
			for i < len(section) && (section[i] == ' ' || section[i] == '\t') {
				i++
			}
			// Skip quoted value
			if i < len(section) && (section[i] == '"' || section[i] == '\'') {
				quote := section[i]
				i++ // skip opening quote
				for i < len(section) && section[i] != quote {
					i++
				}
				if i < len(section) {
					i++ // skip closing quote
				}
			} else {
				// Skip unquoted value
				for i < len(section) && section[i] != ' ' && section[i] != '\t' && section[i] != '\n' && section[i] != '\r' {
					i++
				}
			}
		}
	}

	return matches
}

// isCustomElement checks if a tag name looks like a custom element (contains hyphen)
func isCustomElement(tagName string) bool {
	return strings.Contains(tagName, "-") && !strings.HasPrefix(tagName, "-") && !strings.HasSuffix(tagName, "-")
}

// getCustomElementAttributes gets the attributes for a custom element from the manifest
func getCustomElementAttributes(ctx DiagnosticsContext, tagName string) []M.Attribute {
	// Use the Attributes method from DiagnosticsContext
	if attrMap, exists := ctx.Attributes(tagName); exists {
		// Convert map to slice
		attrs := make([]M.Attribute, 0, len(attrMap))
		for _, attr := range attrMap {
			if attr != nil {
				attrs = append(attrs, *attr)
			}
		}
		return attrs
	}
	return nil
}

// findClosestAttribute finds the closest matching attribute from a list of manifest attributes
func findClosestAttribute(target string, attributes []M.Attribute) string {
	if len(attributes) == 0 {
		return ""
	}

	bestMatch := ""
	bestDistance := 999

	for _, attr := range attributes {
		distance := levenshteinDistance(strings.ToLower(target), strings.ToLower(attr.Name))
		if distance < bestDistance && distance <= 3 { // Only suggest if reasonably close
			bestDistance = distance
			bestMatch = attr.Name
		}
	}

	return bestMatch
}

// findClosestGlobalAttribute finds the closest matching global HTML attribute
func findClosestGlobalAttribute(target string) string {
	bestMatch := ""
	bestDistance := 999

	targetLower := strings.ToLower(target)

	for attrName := range globalAttributes {
		distance := levenshteinDistance(targetLower, attrName)
		if distance < bestDistance && distance <= 2 { // Only suggest if close
			bestDistance = distance
			bestMatch = attrName
		}
	}

	return bestMatch
}

// createAttributeDiagnostic creates a diagnostic for an attribute with a suggested correction
func createAttributeDiagnostic(match AttributeMatch, suggestion string) protocol.Diagnostic {
	message := fmt.Sprintf("Unknown attribute '%s'", match.Name)
	if suggestion != "" {
		message = fmt.Sprintf("Unknown attribute '%s'. Did you mean '%s'?", match.Name, suggestion)
	}

	diagnostic := protocol.Diagnostic{
		Range: protocol.Range{
			Start: protocol.Position{Line: match.Line, Character: match.StartCol},
			End:   protocol.Position{Line: match.Line, Character: match.EndCol},
		},
		Message:  message,
		Severity: &[]protocol.DiagnosticSeverity{protocol.DiagnosticSeverityWarning}[0],
		Source:   &[]string{"cem-lsp"}[0],
	}

	// Add autofix data if we have a suggestion
	if suggestion != "" {
		autofixData := &types.AutofixData{
			Type:       types.DiagnosticTypeAttributeSuggestion,
			Original:   match.Name,
			Suggestion: suggestion,
			Range:      diagnostic.Range,
		}
		diagnostic.Data = autofixData.ToMap()
	}

	return diagnostic
}

// createUnknownAttributeDiagnostic creates a diagnostic for a completely unknown attribute
func createUnknownAttributeDiagnostic(match AttributeMatch) protocol.Diagnostic {
	var message string
	if isCustomElement(match.TagName) {
		message = fmt.Sprintf("Unknown attribute '%s' for custom element '%s'", match.Name, match.TagName)
	} else {
		message = fmt.Sprintf("Unknown attribute '%s' for element '%s'", match.Name, match.TagName)
	}

	return protocol.Diagnostic{
		Range: protocol.Range{
			Start: protocol.Position{Line: match.Line, Character: match.StartCol},
			End:   protocol.Position{Line: match.Line, Character: match.EndCol},
		},
		Message:  message,
		Severity: &[]protocol.DiagnosticSeverity{protocol.DiagnosticSeverityWarning}[0],
		Source:   &[]string{"cem-lsp"}[0],
	}
}

// levenshteinDistance calculates the edit distance between two strings
func levenshteinDistance(a, b string) int {
	if len(a) == 0 {
		return len(b)
	}
	if len(b) == 0 {
		return len(a)
	}

	matrix := make([][]int, len(a)+1)
	for i := range matrix {
		matrix[i] = make([]int, len(b)+1)
		matrix[i][0] = i
	}
	for j := 0; j <= len(b); j++ {
		matrix[0][j] = j
	}

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			cost := 0
			if a[i-1] != b[j-1] {
				cost = 1
			}
			matrix[i][j] = minAttribute(
				matrix[i-1][j]+1,      // deletion
				matrix[i][j-1]+1,      // insertion
				matrix[i-1][j-1]+cost, // substitution
			)
		}
	}

	return matrix[len(a)][len(b)]
}

func minAttribute(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= c {
		return b
	}
	return c
}
