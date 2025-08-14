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
	"slices"
	"strings"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// analyzeTagNameDiagnostics finds invalid custom element tag names and suggests corrections
func analyzeTagNameDiagnostics(ctx DiagnosticsContext, doc types.Document) []protocol.Diagnostic {
	var diagnostics []protocol.Diagnostic

	// Get document content to search for custom element tags
	content := ""
	if docWithContent, ok := doc.(interface{ Content() string }); ok {
		content = docWithContent.Content()
	} else {
		return diagnostics
	}

	helpers.SafeDebugLog("[DIAGNOSTICS] Analyzing custom element tag names in document (length=%d)", len(content))

	// Find all custom element tag names in the document
	tagMatches := findCustomElementTags(content)
	helpers.SafeDebugLog("[DIAGNOSTICS] Found %d custom element tags", len(tagMatches))

	// Get all available tag names from manifest
	availableTagNames := ctx.AllTagNames()
	helpers.SafeDebugLog("[DIAGNOSTICS] Available tag names: %v", availableTagNames)

	for _, match := range tagMatches {
		tagName := match.Value
		helpers.SafeDebugLog("[DIAGNOSTICS] Checking tag '%s'", tagName)

		// Check if the tag exists in the manifest
		isValid := slices.Contains(availableTagNames, tagName)

		if !isValid {
			// Check if this element exists but might need an import
			if _, exists := ctx.ElementDefinition(tagName); exists {
				if importPath, hasSource := ctx.ElementSource(tagName); hasSource {
					// Element exists but may need import - create missing import diagnostic
					// Note: importPath currently contains module path, needs package.json resolution
					helpers.SafeDebugLog("[DIAGNOSTICS] Element '%s' exists but may need import from '%s'", tagName, importPath)

					severity := protocol.DiagnosticSeverityError
					source := "cem-lsp"
					var diagnostic protocol.Diagnostic
					diagnostic.Range = match.Range
					diagnostic.Severity = &severity
					diagnostic.Source = &source
					diagnostic.Message = fmt.Sprintf("Custom element '%s' is not imported. Add import from '%s'", tagName, importPath)

					// Add missing import autofix data
					autofixData := &types.AutofixData{
						Type:       types.DiagnosticTypeMissingImport,
						Original:   tagName,
						Suggestion: fmt.Sprintf("import '%s'", importPath),
						Range:      match.Range,
						ImportPath: importPath,
						TagName:    tagName,
					}
					diagnostic.Data = autofixData.ToMap()

					diagnostics = append(diagnostics, diagnostic)
					helpers.SafeDebugLog("[DIAGNOSTICS] Added missing import diagnostic for tag '%s'", tagName)
					continue
				}
			}
			helpers.SafeDebugLog("[DIAGNOSTICS] Invalid custom element tag '%s'", tagName)

			// Find closest match using Levenshtein distance
			closestMatch, distance := findClosestMatch(tagName, availableTagNames, 3)

			severity := protocol.DiagnosticSeverityError
			source := "cem-lsp"
			var diagnostic protocol.Diagnostic
			diagnostic.Range = match.Range
			diagnostic.Severity = &severity
			diagnostic.Source = &source

			if closestMatch != "" && distance <= 2 {
				diagnostic.Message = fmt.Sprintf("Unknown custom element '%s'. Did you mean '%s'?", tagName, closestMatch)

				// Add code action data for quick fix
				autofixData := &types.AutofixData{
					Type:       types.DiagnosticTypeTagSuggestion,
					Original:   tagName,
					Suggestion: closestMatch,
					Range:      match.Range,
				}
				diagnostic.Data = autofixData.ToMap()
			} else {
				if len(availableTagNames) > 0 {
					// For large distances, avoid showing random elements which might not be helpful
					if len(availableTagNames) <= 5 {
						// Only show elements if there are very few available
						// Sort for consistent output in tests and user experience
						sortedNames := make([]string, len(availableTagNames))
						copy(sortedNames, availableTagNames)
						slices.Sort(sortedNames)
						availableList := strings.Join(sortedNames, "', '")
						diagnostic.Message = fmt.Sprintf("Unknown custom element '%s'. Available elements: '%s'", tagName, availableList)
					} else {
						// For projects with many elements, suggest checking documentation instead
						diagnostic.Message = fmt.Sprintf("Unknown custom element '%s'. Check available elements in your project's manifest or documentation.", tagName)
					}
				} else {
					diagnostic.Message = fmt.Sprintf("Unknown custom element '%s'. No custom elements found in manifest.", tagName)
				}
			}

			diagnostics = append(diagnostics, diagnostic)
			helpers.SafeDebugLog("[DIAGNOSTICS] Added diagnostic for invalid tag '%s'", tagName)
		}
	}

	return diagnostics
}

// TagMatch represents a found custom element tag in the document
type TagMatch struct {
	Value string
	Range protocol.Range
}

// findCustomElementTags finds all custom element tag names in the document content
func findCustomElementTags(content string) []TagMatch {
	var matches []TagMatch
	lines := strings.Split(content, "\n")

	for lineIdx, line := range lines {
		// Look for opening tags with hyphens (custom elements)
		idx := 0
		for {
			tagStart := strings.Index(line[idx:], "<")
			if tagStart == -1 {
				break
			}
			tagStart += idx

			// Skip closing tags
			if tagStart+1 < len(line) && line[tagStart+1] == '/' {
				idx = tagStart + 1
				continue
			}

			// Find the end of the tag name (stop at space, /, >, or other delimiters)
			tagEnd := -1
			delimiters := " \t\n\r/>="
			spaceIdx := strings.IndexAny(line[tagStart+1:], delimiters)
			if spaceIdx != -1 {
				tagEnd = tagStart + 1 + spaceIdx
			}

			if tagEnd == -1 {
				idx = tagStart + 1
				continue
			}

			// Extract the tag name
			tagName := line[tagStart+1 : tagEnd]

			// Validate that this looks like a custom element (has a dash and is valid)
			if strings.Contains(tagName, "-") && isValidCustomElementName(tagName) {
				match := TagMatch{
					Value: tagName,
					Range: protocol.Range{
						Start: protocol.Position{
							Line:      uint32(lineIdx),
							Character: uint32(tagStart + 1),
						},
						End: protocol.Position{
							Line:      uint32(lineIdx),
							Character: uint32(tagEnd),
						},
					},
				}
				matches = append(matches, match)
			}

			idx = tagEnd
		}
	}

	return matches
}

// isValidCustomElementName checks if a tag name follows custom element naming rules
func isValidCustomElementName(tagName string) bool {
	// Basic validation - must contain hyphen and not start with invalid characters
	if !strings.Contains(tagName, "-") {
		return false
	}

	// Cannot start with certain reserved patterns
	if strings.HasPrefix(tagName, "xml") || strings.HasPrefix(tagName, "xmlns") {
		return false
	}

	// Must not contain uppercase (CE spec allows case-insensitive but lowercase is convention)
	if strings.ToLower(tagName) != tagName {
		return false
	}

	return true
}

// Helper function for minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
