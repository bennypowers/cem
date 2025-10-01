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
	"fmt"
	"strings"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
	Q "bennypowers.dev/cem/queries"
	"bennypowers.dev/cem/validations"
	"github.com/agext/levenshtein"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// AttributeMatch represents a found attribute in the document
type AttributeMatch struct {
	Name     string
	TagName  string // The tag this attribute belongs to
	Value    string // The attribute value (empty if no value)
	HasValue bool   // Whether attribute has an explicit value
	Line     uint32
	StartCol uint32
	EndCol   uint32
}

// analyzeAttributeDiagnostics finds unknown attributes and suggests corrections
func analyzeAttributeDiagnostics(ctx types.ServerContext, doc types.Document) []protocol.Diagnostic {
	diagnostics, err := AnalyzeAttributeDiagnosticsForTest(ctx, doc)
	if err != nil {
		helpers.SafeDebugLog("[DIAGNOSTICS] Attribute analysis failed: %v", err)
		return []protocol.Diagnostic{}
	}
	return diagnostics
}

// AnalyzeAttributeDiagnosticsForTest is the exported version for testing
func AnalyzeAttributeDiagnosticsForTest(ctx types.ServerContext, doc types.Document) ([]protocol.Diagnostic, error) {
	var diagnostics []protocol.Diagnostic

	// Get document content to search for attributes
	content, err := doc.Content()
	if err != nil {
		return nil, fmt.Errorf("failed to get document content: %w", err)
	}

	helpers.SafeDebugLog("[DIAGNOSTICS] Analyzing attributes in document (length=%d)", len(content))

	// Find all attributes in the document using tree-sitter
	attributeMatches, err := findAttributes(ctx, doc, content)
	if err != nil {
		return nil, fmt.Errorf("failed to find attributes: %w", err)
	}
	helpers.SafeDebugLog("[DIAGNOSTICS] Found %d attributes", len(attributeMatches))

	for _, match := range attributeMatches {
		// For custom elements, validate all their attributes
		if isCustomElement(match.TagName) {
			// Skip if it's a global HTML attribute (valid on all elements)
			if validations.IsGlobalAttribute(match.Name) {
				continue
			}

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
					// Check for global attribute typos as fallback
					globalSuggestion := findClosestGlobalAttribute(match.Name)
					if globalSuggestion != "" {
						diagnostics = append(diagnostics, createAttributeDiagnostic(match, globalSuggestion))
					} else {
						diagnostics = append(diagnostics, createUnknownAttributeDiagnostic(match))
					}
				}
			} else {
				// No manifest found for this custom element, warn about unknown attribute
				diagnostics = append(diagnostics, createUnknownAttributeDiagnostic(match))
			}
		} else {
			// For standard HTML elements, CEM-LSP ONLY validates slot attributes
			// All other attributes (type, class, id, etc.) are outside CEM's scope
			if match.Name == "slot" {
				// Slot attribute validation is handled by slotDiagnostics.go
				// We don't process slot attributes here to avoid duplication
				continue
			}
			// Ignore ALL other attributes on standard HTML elements
		}
	}

	helpers.SafeDebugLog("[DIAGNOSTICS] Generated %d attribute diagnostics", len(diagnostics))
	return diagnostics, nil
}

// findAttributes finds all attributes in the document using tree-sitter
func findAttributes(ctx types.ServerContext, doc types.Document, content string) ([]AttributeMatch, error) {
	var matches []AttributeMatch

	// Get the tree-sitter tree
	tree := doc.Tree()
	if tree == nil {
		return nil, fmt.Errorf("no tree-sitter tree available")
	}

	// Get query manager from context
	qm, err := ctx.QueryManager()
	if err != nil {
		return nil, fmt.Errorf("failed to get QueryManager: %w", err)
	}

	// Create query matcher for attributes
	matcher, err := Q.NewQueryMatcher(qm, "html", "attributes")
	if err != nil {
		return nil, fmt.Errorf("failed to create attributes query matcher: %w", err)
	}
	defer matcher.Close()

	root := tree.RootNode()
	text := []byte(content)

	// Process start tags
	for captureMap := range matcher.ParentCaptures(root, text, "start.tag") {
		processTagCaptures(captureMap, content, &matches)
	}

	// Process self-closing tags
	for captureMap := range matcher.ParentCaptures(root, text, "self.closing.tag") {
		processTagCaptures(captureMap, content, &matches)
	}

	helpers.SafeDebugLog("[ATTR_DIAG] Tree-sitter found %d attributes", len(matches))
	return matches, nil
}

// processTagCaptures processes a single tag's captures and extracts attributes
func processTagCaptures(captureMap Q.CaptureMap, content string, matches *[]AttributeMatch) {
	var tagName string
	var attrNames []Q.CaptureInfo
	var quotedValues []Q.CaptureInfo
	var unquotedValues []Q.CaptureInfo

	// Extract tag name
	if tagNodes, ok := captureMap["tag.name"]; ok && len(tagNodes) > 0 {
		tagName = tagNodes[0].Text
	}

	if tagName == "" {
		return
	}

	// Extract attribute names
	if names, ok := captureMap["attr.name"]; ok {
		attrNames = names
	}

	// Extract quoted values
	if quoted, ok := captureMap["attr.value.quoted"]; ok {
		quotedValues = quoted
	}

	// Extract unquoted values
	if unquoted, ok := captureMap["attr.value.unquoted"]; ok {
		unquotedValues = unquoted
	}

	// Process each attribute
	for i, attrNameInfo := range attrNames {
		attrName := attrNameInfo.Text
		attrValue := ""
		hasValue := false

		// Check for quoted value
		if i < len(quotedValues) {
			rawValue := quotedValues[i].Text
			// Strip quotes
			if len(rawValue) >= 2 {
				attrValue = rawValue[1 : len(rawValue)-1]
			}
			hasValue = true
		} else if i < len(unquotedValues) {
			attrValue = unquotedValues[i].Text
			hasValue = true
		}

		// Convert byte positions to line/column
		lines := strings.Split(content[:attrNameInfo.StartByte], "\n")
		line := uint32(len(lines) - 1)
		startCol := uint32(len(lines[len(lines)-1]))

		endLines := strings.Split(content[:attrNameInfo.EndByte], "\n")
		endCol := uint32(len(endLines[len(endLines)-1]))

		match := AttributeMatch{
			Name:     attrName,
			TagName:  tagName,
			Value:    attrValue,
			HasValue: hasValue,
			Line:     line,
			StartCol: startCol,
			EndCol:   endCol,
		}

		*matches = append(*matches, match)
	}
}

// isCustomElement checks if a tag name looks like a custom element (contains hyphen)
func isCustomElement(tagName string) bool {
	return strings.Contains(tagName, "-") && !strings.HasPrefix(tagName, "-") && !strings.HasSuffix(tagName, "-")
}

// getCustomElementAttributes gets the attributes for a custom element from the manifest
func getCustomElementAttributes(ctx types.ServerContext, tagName string) []M.Attribute {
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
		distance := levenshtein.Distance(strings.ToLower(target), strings.ToLower(attr.Name), nil)
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

	globalAttributes := validations.GetGlobalAttributes()
	for attrName := range globalAttributes {
		distance := levenshtein.Distance(targetLower, attrName, nil)
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
