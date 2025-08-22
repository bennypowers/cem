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
	"regexp"
	"strconv"
	"strings"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// attributeRegex parses attribute name-value pairs
// Matches: name="value", name='value', name=value, or just name
var attributeRegex = regexp.MustCompile(`(\w+)(?:\s*=\s*(?:"([^"]*)"|'([^']*)'|([^\s>]+)))?`)

// analyzeAttributeValueDiagnostics validates attribute values against their type definitions
func analyzeAttributeValueDiagnostics(ctx types.ServerContext, doc types.Document) []protocol.Diagnostic {
	return AnalyzeAttributeValueDiagnosticsForTest(ctx, doc)
}

// AnalyzeAttributeValueDiagnosticsForTest is the exported version for testing
func AnalyzeAttributeValueDiagnosticsForTest(ctx types.ServerContext, doc types.Document) []protocol.Diagnostic {
	var diagnostics []protocol.Diagnostic

	// Get document content to search for attributes
	content, err := doc.Content()
	if err != nil {
		return diagnostics
	}

	helpers.SafeDebugLog("[VALUE_DIAGNOSTICS] Analyzing attribute values in document (length=%d)", len(content))

	// Find all attributes in the document with their values
	attributeMatches := findAttributesWithValues(content)
	helpers.SafeDebugLog("[VALUE_DIAGNOSTICS] Found %d attributes with values", len(attributeMatches))

	for _, match := range attributeMatches {
		// Only validate custom elements
		if !isCustomElement(match.TagName) {
			continue
		}

		// Get attribute definition for this custom element
		if attrs, exists := ctx.Attributes(match.TagName); exists {
			if attr, attrExists := attrs[match.Name]; attrExists && attr != nil {
				// Validate the attribute value against its type
				valueDiagnostics := validateAttributeValue(attr, match)
				diagnostics = append(diagnostics, valueDiagnostics...)
			}
		}
	}

	helpers.SafeDebugLog("[VALUE_DIAGNOSTICS] Generated %d attribute value diagnostics", len(diagnostics))
	return diagnostics
}

// validateAttributeValue validates an attribute value against its type definition
func validateAttributeValue(attr *M.Attribute, match AttributeMatch) []protocol.Diagnostic {
	var diagnostics []protocol.Diagnostic

	// Skip validation if no type information
	if attr.Type == nil || attr.Type.Text == "" {
		return diagnostics
	}

	typeText := strings.TrimSpace(attr.Type.Text)
	lowerTypeText := strings.ToLower(typeText)

	helpers.SafeDebugLog("[VALUE_DIAGNOSTICS] Validating attribute %s.%s with type '%s', value='%s', hasValue=%t",
		match.TagName, match.Name, typeText, match.Value, match.HasValue)

	switch {
	case lowerTypeText == "boolean":
		diagnostics = append(diagnostics, validateBooleanAttribute(match)...)
	case lowerTypeText == "number":
		diagnostics = append(diagnostics, validateNumberAttribute(match)...)
	case lowerTypeText == "string":
		// String attributes are very permissive - almost any value is valid
		// Generic string validation is not generalizable due to wide variety of use cases:
		// - URLs, file paths, CSS selectors, arbitrary text, IDs, JSON strings, etc.
		// - Each string attribute type would need domain-specific validation rules
		// - Better to provide no validation than incorrect validation
	case strings.Contains(typeText, "|"):
		diagnostics = append(diagnostics, validateUnionType(typeText, match)...)
	case isLiteralType(typeText):
		diagnostics = append(diagnostics, validateLiteralType(typeText, match)...)
	case isArrayType(lowerTypeText):
		diagnostics = append(diagnostics, createArrayTypeInfo(match)...)
	default:
		helpers.SafeDebugLog("[VALUE_DIAGNOSTICS] Skipping validation for unknown type: %s", typeText)
	}

	return diagnostics
}

// validateBooleanAttribute validates boolean attributes according to HTML semantics
func validateBooleanAttribute(match AttributeMatch) []protocol.Diagnostic {
	var diagnostics []protocol.Diagnostic

	// Boolean attributes: presence = true, absence = false
	if !match.HasValue {
		// No value is perfectly valid for boolean attributes
		return diagnostics
	}

	// Boolean attribute has a value - this might be problematic
	if match.Value == "false" {
		// This is a common mistake - boolean attribute with "false" is still true
		diagnostics = append(diagnostics, protocol.Diagnostic{
			Range: protocol.Range{
				Start: protocol.Position{Line: match.Line, Character: match.StartCol},
				End:   protocol.Position{Line: match.Line, Character: match.EndCol},
			},
			Message:  fmt.Sprintf("Boolean attribute '%s' with value 'false' is still true. Remove the attribute entirely to make it false.", match.Name),
			Severity: &[]protocol.DiagnosticSeverity{protocol.DiagnosticSeverityWarning}[0],
			Source:   &[]string{"cem-lsp"}[0],
		})
	} else if match.Value == "true" {
		// This is redundant but not wrong
		diagnostics = append(diagnostics, protocol.Diagnostic{
			Range: protocol.Range{
				Start: protocol.Position{Line: match.Line, Character: match.StartCol},
				End:   protocol.Position{Line: match.Line, Character: match.EndCol},
			},
			Message:  fmt.Sprintf("Boolean attribute '%s' with value 'true' is redundant. Use <%s %s> instead.", match.Name, match.TagName, match.Name),
			Severity: &[]protocol.DiagnosticSeverity{protocol.DiagnosticSeverityInformation}[0],
			Source:   &[]string{"cem-lsp"}[0],
		})
	} else if match.Value != "" && match.Value != match.Name {
		// Non-standard boolean value
		diagnostics = append(diagnostics, protocol.Diagnostic{
			Range: protocol.Range{
				Start: protocol.Position{Line: match.Line, Character: match.StartCol},
				End:   protocol.Position{Line: match.Line, Character: match.EndCol},
			},
			Message:  fmt.Sprintf("Boolean attribute '%s' should not have value '%s'. Use <%s %s> instead.", match.Name, match.Value, match.TagName, match.Name),
			Severity: &[]protocol.DiagnosticSeverity{protocol.DiagnosticSeverityInformation}[0],
			Source:   &[]string{"cem-lsp"}[0],
		})
	}

	return diagnostics
}

// validateNumberAttribute validates numeric attribute values
func validateNumberAttribute(match AttributeMatch) []protocol.Diagnostic {
	var diagnostics []protocol.Diagnostic

	if !match.HasValue || match.Value == "" {
		// Empty value for number attribute is usually invalid
		diagnostics = append(diagnostics, protocol.Diagnostic{
			Range: protocol.Range{
				Start: protocol.Position{Line: match.Line, Character: match.StartCol},
				End:   protocol.Position{Line: match.Line, Character: match.EndCol},
			},
			Message:  fmt.Sprintf("Number attribute '%s' requires a numeric value", match.Name),
			Severity: &[]protocol.DiagnosticSeverity{protocol.DiagnosticSeverityError}[0],
			Source:   &[]string{"cem-lsp"}[0],
		})
		return diagnostics
	}

	// Try to parse as a number
	if _, err := strconv.ParseFloat(match.Value, 64); err != nil {
		diagnostics = append(diagnostics, protocol.Diagnostic{
			Range: protocol.Range{
				Start: protocol.Position{Line: match.Line, Character: match.StartCol},
				End:   protocol.Position{Line: match.Line, Character: match.EndCol},
			},
			Message:  fmt.Sprintf("Expected number for attribute '%s', got '%s'", match.Name, match.Value),
			Severity: &[]protocol.DiagnosticSeverity{protocol.DiagnosticSeverityError}[0],
			Source:   &[]string{"cem-lsp"}[0],
		})
	}

	return diagnostics
}


// validateUnionType validates attribute values against union type definitions
func validateUnionType(typeText string, match AttributeMatch) []protocol.Diagnostic {
	var diagnostics []protocol.Diagnostic

	if !match.HasValue || match.Value == "" {
		// Empty value for union type might be invalid
		diagnostics = append(diagnostics, protocol.Diagnostic{
			Range: protocol.Range{
				Start: protocol.Position{Line: match.Line, Character: match.StartCol},
				End:   protocol.Position{Line: match.Line, Character: match.EndCol},
			},
			Message:  fmt.Sprintf("Union type attribute '%s' requires a value", match.Name),
			Severity: &[]protocol.DiagnosticSeverity{protocol.DiagnosticSeverityError}[0],
			Source:   &[]string{"cem-lsp"}[0],
		})
		return diagnostics
	}

	// Parse union options from type text (e.g., "red" | "green" | "blue")
	options := parseUnionOptions(typeText)
	if len(options) == 0 {
		return diagnostics // Skip validation if we can't parse the union
	}

	// Check if the value matches any of the union options
	for _, option := range options {
		if match.Value == option {
			return diagnostics // Valid value found
		}
	}

	// Value doesn't match any option - suggest the closest one
	suggestion := findClosestUnionOption(match.Value, options)
	if suggestion != "" {
		diagnostic := protocol.Diagnostic{
			Range: protocol.Range{
				Start: protocol.Position{Line: match.Line, Character: match.StartCol},
				End:   protocol.Position{Line: match.Line, Character: match.EndCol},
			},
			Message:  fmt.Sprintf("Expected one of: %s for attribute '%s', got '%s'. Did you mean '%s'?", formatUnionOptions(options), match.Name, match.Value, suggestion),
			Severity: &[]protocol.DiagnosticSeverity{protocol.DiagnosticSeverityError}[0],
			Source:   &[]string{"cem-lsp"}[0],
		}

		// Add autofix data
		autofixData := &types.AutofixData{
			Type:       types.DiagnosticTypeAttributeValueSuggestion,
			Original:   match.Value,
			Suggestion: suggestion,
			Range:      diagnostic.Range,
		}
		diagnostic.Data = autofixData.ToMap()

		diagnostics = append(diagnostics, diagnostic)
	} else {
		diagnostics = append(diagnostics, protocol.Diagnostic{
			Range: protocol.Range{
				Start: protocol.Position{Line: match.Line, Character: match.StartCol},
				End:   protocol.Position{Line: match.Line, Character: match.EndCol},
			},
			Message:  fmt.Sprintf("Expected one of: %s for attribute '%s', got '%s'", formatUnionOptions(options), match.Name, match.Value),
			Severity: &[]protocol.DiagnosticSeverity{protocol.DiagnosticSeverityError}[0],
			Source:   &[]string{"cem-lsp"}[0],
		})
	}

	return diagnostics
}

// validateLiteralType validates attribute values against literal type definitions
func validateLiteralType(typeText string, match AttributeMatch) []protocol.Diagnostic {
	var diagnostics []protocol.Diagnostic

	// Extract the literal value (remove quotes)
	expectedValue := strings.Trim(typeText, `"'`)

	if !match.HasValue || match.Value == "" {
		diagnostics = append(diagnostics, protocol.Diagnostic{
			Range: protocol.Range{
				Start: protocol.Position{Line: match.Line, Character: match.StartCol},
				End:   protocol.Position{Line: match.Line, Character: match.EndCol},
			},
			Message:  fmt.Sprintf("Expected literal value '%s' for attribute '%s'", expectedValue, match.Name),
			Severity: &[]protocol.DiagnosticSeverity{protocol.DiagnosticSeverityError}[0],
			Source:   &[]string{"cem-lsp"}[0],
		})
		return diagnostics
	}

	if match.Value != expectedValue {
		// Check for case mismatch
		if strings.EqualFold(match.Value, expectedValue) {
			diagnostic := protocol.Diagnostic{
				Range: protocol.Range{
					Start: protocol.Position{Line: match.Line, Character: match.StartCol},
					End:   protocol.Position{Line: match.Line, Character: match.EndCol},
				},
				Message:  fmt.Sprintf("Expected literal value '%s', got '%s' (case mismatch)", expectedValue, match.Value),
				Severity: &[]protocol.DiagnosticSeverity{protocol.DiagnosticSeverityError}[0],
				Source:   &[]string{"cem-lsp"}[0],
			}

			// Add autofix data for case correction
			autofixData := &types.AutofixData{
				Type:       types.DiagnosticTypeAttributeValueSuggestion,
				Original:   match.Value,
				Suggestion: expectedValue,
				Range:      diagnostic.Range,
			}
			diagnostic.Data = autofixData.ToMap()

			diagnostics = append(diagnostics, diagnostic)
		} else {
			diagnostics = append(diagnostics, protocol.Diagnostic{
				Range: protocol.Range{
					Start: protocol.Position{Line: match.Line, Character: match.StartCol},
					End:   protocol.Position{Line: match.Line, Character: match.EndCol},
				},
				Message:  fmt.Sprintf("Expected literal value '%s' for attribute '%s', got '%s'", expectedValue, match.Name, match.Value),
				Severity: &[]protocol.DiagnosticSeverity{protocol.DiagnosticSeverityError}[0],
				Source:   &[]string{"cem-lsp"}[0],
			})
		}
	}

	return diagnostics
}

// createArrayTypeInfo provides informational message about array types
func createArrayTypeInfo(match AttributeMatch) []protocol.Diagnostic {
	return []protocol.Diagnostic{
		{
			Range: protocol.Range{
				Start: protocol.Position{Line: match.Line, Character: match.StartCol},
				End:   protocol.Position{Line: match.Line, Character: match.EndCol},
			},
			Message:  "Array attributes support multiple formats (JSON, comma-separated, space-separated). Refer to component documentation.",
			Severity: &[]protocol.DiagnosticSeverity{protocol.DiagnosticSeverityInformation}[0],
			Source:   &[]string{"cem-lsp"}[0],
		},
	}
}

// Helper functions for type checking

// isLiteralType checks if a type definition is a literal type (e.g., "primary" or 'primary')
func isLiteralType(typeText string) bool {
	typeText = strings.TrimSpace(typeText)
	return (strings.HasPrefix(typeText, "'") && strings.HasSuffix(typeText, "'")) ||
		(strings.HasPrefix(typeText, "\"") && strings.HasSuffix(typeText, "\""))
}

// isArrayType checks if a type definition is an array type
func isArrayType(typeText string) bool {
	return strings.Contains(typeText, "[]") || strings.Contains(typeText, "array")
}

// parseUnionOptions extracts options from a union type string
func parseUnionOptions(typeText string) []string {
	var options []string

	// Split by | and clean up quotes and whitespace
	parts := strings.Split(typeText, "|")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		part = strings.Trim(part, `"'`) // Remove quotes
		if part != "" {
			options = append(options, part)
		}
	}

	return options
}

// findClosestUnionOption finds the closest matching option using Levenshtein distance
func findClosestUnionOption(value string, options []string) string {
	if len(options) == 0 {
		return ""
	}

	bestMatch := ""
	bestDistance := 999

	for _, option := range options {
		distance := levenshteinDistance(strings.ToLower(value), strings.ToLower(option))
		if distance < bestDistance && distance <= 2 { // Only suggest if reasonably close
			bestDistance = distance
			bestMatch = option
		}
	}

	return bestMatch
}

// formatUnionOptions formats union options for display in error messages
func formatUnionOptions(options []string) string {
	if len(options) == 0 {
		return ""
	}

	quoted := make([]string, len(options))
	for i, option := range options {
		quoted[i] = fmt.Sprintf("'%s'", option)
	}

	if len(quoted) == 1 {
		return quoted[0]
	}

	if len(quoted) == 2 {
		return quoted[0] + " or " + quoted[1]
	}

	return strings.Join(quoted[:len(quoted)-1], ", ") + " or " + quoted[len(quoted)-1]
}

// findAttributesWithValues finds all attributes in the document content with their values
func findAttributesWithValues(content string) []AttributeMatch {
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

			// Find attributes in this section with their values
			attrMatches := findAttributesWithValuesInSection(attrSection, tagName, uint32(lineIdx), uint32(tagStart+1+tagNameEnd))
			matches = append(matches, attrMatches...)

			idx = tagEnd + 1
		}
	}

	return matches
}

// findAttributesWithValuesInSection finds attributes within a tag's attribute section with their values
func findAttributesWithValuesInSection(section string, tagName string, line uint32, startOffset uint32) []AttributeMatch {
	var matches []AttributeMatch

	// Use pre-compiled regex to parse attribute name-value pairs
	attrMatches := attributeRegex.FindAllStringSubmatch(section, -1)

	for _, attrMatch := range attrMatches {
		if len(attrMatch) < 2 {
			continue
		}

		attrName := attrMatch[1]
		var attrValue string
		hasValue := false

		// Check which capture group has the value
		if attrMatch[2] != "" { // Double-quoted value
			attrValue = attrMatch[2]
			hasValue = true
		} else if attrMatch[3] != "" { // Single-quoted value
			attrValue = attrMatch[3]
			hasValue = true
		} else if attrMatch[4] != "" { // Unquoted value
			attrValue = attrMatch[4]
			hasValue = true
		} else {
			// No value (boolean attribute)
			attrValue = ""
			hasValue = false
		}

		// Find the position of this attribute in the section
		attrPos := strings.Index(section, attrMatch[0])
		if attrPos == -1 {
			continue
		}

		matches = append(matches, AttributeMatch{
			Name:     attrName,
			TagName:  tagName,
			Value:    attrValue,
			HasValue: hasValue,
			Line:     line,
			StartCol: startOffset + uint32(attrPos),
			EndCol:   startOffset + uint32(attrPos+len(attrMatch[0])),
		})
	}

	return matches
}