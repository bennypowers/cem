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
	ts "github.com/tree-sitter/go-tree-sitter"
)

// analyzeSlotDiagnostics finds invalid slot attribute values and suggests corrections
func analyzeSlotDiagnostics(ctx types.ServerContext, doc types.Document) []protocol.Diagnostic {
	return AnalyzeSlotDiagnosticsForTest(ctx, doc)
}

// AnalyzeSlotDiagnosticsForTest is the exported version for testing
func AnalyzeSlotDiagnosticsForTest(ctx types.ServerContext, doc types.Document) []protocol.Diagnostic {
	var diagnostics []protocol.Diagnostic

	content, err := doc.Content()
	if err != nil {
		return diagnostics
	}

	helpers.SafeDebugLog("[DIAGNOSTICS] Analyzing slot attributes in document (length=%d)", len(content))

	slotMatches := findSlotAttributes(content)
	helpers.SafeDebugLog("[DIAGNOSTICS] Found %d slot attributes", len(slotMatches))

	for _, match := range slotMatches {
		parentElement := findParentElement(doc, match.Range.Start)
		if parentElement == "" {
			helpers.SafeDebugLog("[DIAGNOSTICS] No parent element found for slot at line %d", match.Range.Start.Line)
			continue
		}

		helpers.SafeDebugLog("[DIAGNOSTICS] Found slot '%s' with parent element '%s'", match.Value, parentElement)

		availableSlots, exists := ctx.Slots(parentElement)
		if !exists || len(availableSlots) == 0 {
			helpers.SafeDebugLog("[DIAGNOSTICS] No slots available for element '%s'", parentElement)
			continue
		}

		slotNames := make([]string, len(availableSlots))
		for i, slot := range availableSlots {
			slotNames[i] = slot.Name
		}

		helpers.SafeDebugLog("[DIAGNOSTICS] Available slots for '%s': %v", parentElement, slotNames)

		slotValue := match.Value
		if slices.Contains(slotNames, slotValue) {
			continue
		}

		helpers.SafeDebugLog("[DIAGNOSTICS] Invalid slot '%s' for element '%s'", slotValue, parentElement)

		closestMatch, distance := findClosestMatch(slotValue, slotNames, 3)

		severity := protocol.DiagnosticSeverityError
		source := "cem-lsp"
		var diagnostic protocol.Diagnostic
		diagnostic.Range = match.Range
		diagnostic.Severity = &severity
		diagnostic.Source = &source

		if closestMatch != "" && distance <= 2 {
			diagnostic.Message = fmt.Sprintf("Unknown slot '%s' for element '%s'. Did you mean '%s'?", slotValue, parentElement, closestMatch)

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

	return diagnostics
}

// SlotMatch represents a found slot attribute in the document
type SlotMatch struct {
	Value string
	Range protocol.Range
}

// findSlotAttributes finds all slot="value" attributes in the document content
func findSlotAttributes(content string) []SlotMatch {
	var matches []SlotMatch
	lines := strings.Split(content, "\n")

	for lineIdx, line := range lines {
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
					break
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

			slotValue := line[quoteStart+1 : quoteEnd]
			if slotValue != "" {
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
				}
				matches = append(matches, match)
			}

			idx = quoteEnd + 1
		}
	}

	return matches
}

// findParentElement finds the parent custom element for a slotted element
// using tree-sitter AST traversal, with a string-scanning fallback.
func findParentElement(doc types.Document, position protocol.Position) string {
	if tree := doc.Tree(); tree != nil {
		if tag := findParentElementTreeSitter(doc, tree, position); tag != "" {
			return tag
		}
	}
	return findParentElementFallback(doc, position)
}

// findParentElementTreeSitter walks up the tree-sitter AST from the slot
// attribute position, skipping the element that owns the slot attribute,
// and returns the first ancestor custom element tag name.
func findParentElementTreeSitter(doc types.Document, tree *ts.Tree, position protocol.Position) string {
	content, err := doc.Content()
	if err != nil {
		return ""
	}
	contentBytes := []byte(content)

	node := tree.RootNode().NamedDescendantForPointRange(
		ts.Point{Row: uint(position.Line), Column: uint(position.Character)},
		ts.Point{Row: uint(position.Line), Column: uint(position.Character)},
	)

	// Walk up to find the element that contains this slot attribute,
	// then continue to find its parent custom element.
	// Only match "element" nodes (not start_tag) to avoid double-counting
	// the owner element, since start_tag is a child of element.
	skippedOwner := false
	for node != nil {
		if node.Kind() == "element" {
			if !skippedOwner {
				skippedOwner = true
				node = node.Parent()
				continue
			}

			if tagName := elementTagName(node, contentBytes); helpers.IsCustomElementTag(tagName) {
				return tagName
			}
		}

		node = node.Parent()
	}

	return ""
}

// elementTagName extracts the tag name string from an element node
// by finding the tag_name node inside its start_tag or self_closing_tag child.
func elementTagName(element *ts.Node, content []byte) string {
	for i := range element.ChildCount() {
		child := element.Child(i)
		if child == nil {
			continue
		}
		kind := child.Kind()
		if kind != "start_tag" && kind != "self_closing_tag" {
			continue
		}
		// Find the tag_name child within the start_tag
		for j := range child.ChildCount() {
			gc := child.Child(j)
			if gc != nil && gc.Kind() == "tag_name" {
				start, end := gc.ByteRange()
				if start < end && end <= uint(len(content)) {
					return string(content[start:end])
				}
			}
		}
		break
	}
	return ""
}

// findParentElementFallback uses string scanning when tree-sitter is unavailable.
// It correctly computes a content-relative byte offset from the LSP position,
// then walks backward through the content tracking nesting depth to find the
// direct parent custom element.
func findParentElementFallback(doc types.Document, position protocol.Position) string {
	content, err := doc.Content()
	if err != nil {
		return ""
	}

	lines := strings.Split(content, "\n")
	if int(position.Line) >= len(lines) {
		return ""
	}

	// Convert LSP position to byte offset
	offset := 0
	for i := 0; i < int(position.Line); i++ {
		offset += len(lines[i]) + 1 // +1 for newline
	}
	offset += int(position.Character)
	if offset > len(content) {
		offset = len(content)
	}

	// Walk backward tracking nesting depth to find the parent custom element.
	// depth tracks how many levels of closing tags we've seen without matching
	// openers; when we find an opening tag at depth 0, it's our direct parent.
	depth := 0
	skippedOwner := false
	inComment := false

	for i := offset - 1; i >= 0; i-- {
		// Track HTML comments when walking backward: we see --> first,
		// then <!-- later. Skip everything between them.
		if !inComment && i >= 2 && content[i] == '>' && content[i-2:i+1] == "-->" {
			inComment = true
			continue
		}
		if inComment {
			if content[i] == '<' && i+4 <= len(content) && content[i:i+4] == "<!--" {
				inComment = false
			}
			continue
		}

		if content[i] != '<' {
			continue
		}

		// Extract tag name
		endIdx := i + 1
		isClosing := endIdx < len(content) && content[endIdx] == '/'
		if isClosing {
			endIdx++
		}
		for endIdx < len(content) && content[endIdx] != ' ' && content[endIdx] != '>' && content[endIdx] != '\n' && content[endIdx] != '\t' && content[endIdx] != '/' {
			endIdx++
		}
		if endIdx <= i+1+boolToInt(isClosing) {
			continue
		}

		start := i + 1
		if isClosing {
			start = i + 2
		}
		tagName := content[start:endIdx]

		if isClosing {
			if strings.Contains(tagName, "-") {
				depth++
			}
			continue
		}

		// Check for self-closing
		isSelfClosing := false
		for j := endIdx; j < len(content) && content[j] != '>'; j++ {
			if content[j] == '/' && j+1 < len(content) && content[j+1] == '>' {
				isSelfClosing = true
				break
			}
		}
		if isSelfClosing {
			continue
		}

		if !strings.Contains(tagName, "-") {
			continue
		}

		if depth > 0 {
			depth--
			continue
		}

		// depth == 0: this is the direct ancestor custom element
		if !skippedOwner {
			// Skip the element that owns the slot attribute
			skippedOwner = true
			continue
		}

		return tagName
	}

	return ""
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// findClosestMatch finds the closest matching string from a list using Levenshtein distance
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
