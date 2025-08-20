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
	"fmt"
	"strings"

	"bennypowers.dev/cem/lsp/types"
	protocol "github.com/tliron/glsp/protocol_3_16"
)


// AnalyzeCompletionContext determines what completion should be provided using tree-sitter
func AnalyzeCompletionContext(doc types.Document, position protocol.Position, triggerChar string) (*types.CompletionAnalysis, error) {
	return AnalyzeCompletionContextWithDM(doc, position, triggerChar, nil)
}

// AnalyzeCompletionContextWithDM determines what completion should be provided using tree-sitter with DocumentManager
func AnalyzeCompletionContextWithDM(doc types.Document, position protocol.Position, triggerChar string, dm any) (*types.CompletionAnalysis, error) {
	if doc == nil {
		return nil, fmt.Errorf("document is nil")
	}

	// Get line content for the analysis
	content, err := doc.Content()
	if err != nil {
		return nil, fmt.Errorf("failed to get document content: %w", err)
	}

	lines := splitLines(content)
	if int(position.Line) >= len(lines) {
		return nil, fmt.Errorf("position line %d is out of bounds (document has %d lines)", position.Line, len(lines))
	}

	lineContent := lines[position.Line]
	if int(position.Character) > len(lineContent) {
		return nil, fmt.Errorf("position character %d is out of bounds for line %d (line has %d characters)", position.Character, position.Line, len(lineContent))
	}

	// Check if document supports tree-sitter analysis
	docWithTreeSitter, ok := doc.(interface {
		AnalyzeCompletionContextTS(position protocol.Position, dm any) *types.CompletionAnalysis
	})
	if !ok {
		return nil, fmt.Errorf("document does not support tree-sitter completion context analysis")
	}

	// Use tree-sitter analysis - this is the only supported method
	tsAnalysis := docWithTreeSitter.AnalyzeCompletionContextTS(position, dm)
	if tsAnalysis == nil {
		return nil, fmt.Errorf("tree-sitter completion context analysis returned nil")
	}

	// Update trigger char and line content
	tsAnalysis.TriggerChar = triggerChar
	tsAnalysis.LineContent = lineContent

	return tsAnalysis, nil
}

// splitLines splits content into lines, handling different line ending types
func splitLines(content string) []string {
	// Handle Windows (\r\n), Unix (\n), and old Mac (\r) line endings
	content = strings.ReplaceAll(content, "\r\n", "\n")
	content = strings.ReplaceAll(content, "\r", "\n")
	return strings.Split(content, "\n")
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
