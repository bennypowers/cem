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
	"github.com/agext/levenshtein"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// analyzeSlotDiagnostics finds invalid slot attribute values and suggests corrections
func analyzeSlotDiagnostics(ctx types.ServerContext, doc types.Document) []protocol.Diagnostic {
	var diagnostics []protocol.Diagnostic

	// Get document content to search for slot attributes
	content, err := doc.Content()
	if err != nil {
		return diagnostics
	}

	helpers.SafeDebugLog("[DIAGNOSTICS] Analyzing slot attributes in document (length=%d)", len(content))

	// Find all slot attributes in the document
	slotMatches := findSlotAttributes(content)
	helpers.SafeDebugLog("[DIAGNOSTICS] Found %d slot attributes", len(slotMatches))

	for _, match := range slotMatches {
		// Find the parent element to get available slots
		parentElement := findParentElement(content, match.Position)
		if parentElement == "" {
			helpers.SafeDebugLog("[DIAGNOSTICS] No parent element found for slot at position %d", match.Position)
			continue
		}

		helpers.SafeDebugLog("[DIAGNOSTICS] Found slot '%s' with parent element '%s'", match.Value, parentElement)

		// Get available slots for the parent element
		availableSlots, exists := ctx.Slots(parentElement)
		if !exists || len(availableSlots) == 0 {
			helpers.SafeDebugLog("[DIAGNOSTICS] No slots available for element '%s'", parentElement)
			continue
		}

		// Extract slot names from manifest
		slotNames := make([]string, len(availableSlots))
		for i, slot := range availableSlots {
			slotNames[i] = slot.Name
		}

		helpers.SafeDebugLog("[DIAGNOSTICS] Available slots for '%s': %v", parentElement, slotNames)

		// Check if the slot value is valid
		slotValue := match.Value
		isValid := slices.Contains(slotNames, slotValue)

		if !isValid {
			helpers.SafeDebugLog("[DIAGNOSTICS] Invalid slot '%s' for element '%s'", slotValue, parentElement)

			// Find closest match using Levenshtein distance
			closestMatch, distance := findClosestMatch(slotValue, slotNames, 3)

			severity := protocol.DiagnosticSeverityError
			source := "cem-lsp"
			var diagnostic protocol.Diagnostic
			diagnostic.Range = match.Range
			diagnostic.Severity = &severity
			diagnostic.Source = &source

			if closestMatch != "" && distance <= 2 {
				diagnostic.Message = fmt.Sprintf("Unknown slot '%s' for element '%s'. Did you mean '%s'?", slotValue, parentElement, closestMatch)

				// Add code action for quick fix
				autofixData := &types.AutofixData{
					Type:       types.DiagnosticTypeSlotSuggestion,
					Original:   slotValue,
					Suggestion: closestMatch,
					Range:      match.Range,
				}
				diagnostic.Data = autofixData.ToMap()
			} else {
				availableList := strings.Join(slotNames, "', '")
				diagnostic.Message = fmt.Sprintf("Unknown slot '%s' for element '%s'. Available slots: '%s'", slotValue, parentElement, availableList)
			}

			diagnostics = append(diagnostics, diagnostic)
			helpers.SafeDebugLog("[DIAGNOSTICS] Added diagnostic for invalid slot '%s'", slotValue)
		}
	}

	return diagnostics
}

// SlotMatch represents a found slot attribute in the document
type SlotMatch struct {
	Value    string
	Range    protocol.Range
	Position int
}

// findSlotAttributes finds all slot="value" attributes in the document content
func findSlotAttributes(content string) []SlotMatch {
	var matches []SlotMatch
	lines := strings.Split(content, "\n")

	for lineIdx, line := range lines {
		// Look for slot="value" patterns
		idx := 0
		for {
			slotIdx := strings.Index(line[idx:], "slot=")
			if slotIdx == -1 {
				break
			}
			slotIdx += idx

			// Find the quote after slot=
			quoteStart := -1
			for i := slotIdx + 5; i < len(line); i++ {
				if line[i] == '"' || line[i] == '\'' {
					quoteStart = i
					break
				}
				if line[i] != ' ' && line[i] != '\t' {
					break // No quote found immediately
				}
			}

			if quoteStart == -1 {
				idx = slotIdx + 5
				continue
			}

			// Find the closing quote
			quote := line[quoteStart]
			quoteEnd := -1
			for i := quoteStart + 1; i < len(line); i++ {
				if line[i] == quote {
					quoteEnd = i
					break
				}
			}

			if quoteEnd == -1 {
				idx = quoteStart + 1
				continue
			}

			// Extract the slot value
			slotValue := line[quoteStart+1 : quoteEnd]
			if slotValue != "" { // Only process non-empty slot values
				match := SlotMatch{
					Value: slotValue,
					Range: protocol.Range{
						Start: protocol.Position{
							Line:      uint32(lineIdx),
							Character: uint32(quoteStart + 1),
						},
						End: protocol.Position{
							Line:      uint32(lineIdx),
							Character: uint32(quoteEnd),
						},
					},
					Position: slotIdx,
				}
				matches = append(matches, match)
			}

			idx = quoteEnd + 1
		}
	}

	return matches
}

// findParentElement finds the parent custom element tag name for a slot attribute
func findParentElement(content string, slotPosition int) string {
	// Work backwards from the slot position to find the opening tag
	beforeSlot := content[:slotPosition]

	// Find the last opening tag before this position
	lastTagStart := -1
	for i := len(beforeSlot) - 1; i >= 0; i-- {
		if beforeSlot[i] == '<' && (i == 0 || beforeSlot[i-1] != '/') {
			lastTagStart = i
			break
		}
	}

	if lastTagStart == -1 {
		return ""
	}

	// Extract the tag name
	tagContent := content[lastTagStart:]
	tagEnd := strings.IndexAny(tagContent, " \t\n\r>")
	if tagEnd == -1 {
		return ""
	}

	tagName := tagContent[1:tagEnd] // Skip the '<'

	// Validate that this looks like a custom element (has a dash)
	if strings.Contains(tagName, "-") {
		return tagName
	}

	return ""
}

// findClosestMatch finds the closest matching string from a list using Levenshtein distance
// Returns the best match and its distance, or empty string if no good match found
func findClosestMatch(target string, candidates []string, maxDistance int) (string, int) {
	bestMatch := ""
	bestDistance := maxDistance + 1

	targetLower := strings.ToLower(target)

	for _, candidate := range candidates {
		candidateLower := strings.ToLower(candidate)
		distance := levenshtein.Distance(targetLower, candidateLower, nil)

		if distance < bestDistance && distance <= maxDistance {
			bestMatch = candidate
			bestDistance = distance
		}
	}

	if bestMatch == "" {
		return "", bestDistance
	}

	return bestMatch, bestDistance
}
