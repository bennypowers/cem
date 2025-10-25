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
package lsp

import (
	"unicode/utf16"
	"unicode/utf8"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	protocol "github.com/tliron/glsp/protocol_3_16"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// Use types from the types package
type ParseStrategy = types.ParseStrategy

const (
	ParseStrategyFull        = types.ParseStrategyFull
	ParseStrategyIncremental = types.ParseStrategyIncremental
	ParseStrategyAuto        = types.ParseStrategyAuto
)

// UTF-16 Conversion Helpers
// LSP uses UTF-16 code units for character positions, while tree-sitter uses UTF-8 byte offsets.
// We need to convert between these two encodings using stdlib unicode/utf16 package.

// UTF16ToByteOffset converts a UTF-16 code unit offset to a UTF-8 byte offset
func UTF16ToByteOffset(text string, utf16Offset uint32) uint {
	var byteOffset uint = 0
	var utf16Count uint32 = 0

	for byteOffset < uint(len(text)) && utf16Count < utf16Offset {
		r, size := utf8.DecodeRuneInString(text[byteOffset:])
		// Check for invalid UTF-8: RuneError with size == 1
		// Note: Valid U+FFFD has size == 3, so we need to check both
		if r == utf8.RuneError && size == 1 {
			// Invalid UTF-8, treat as single byte
			byteOffset++
			utf16Count++
			continue
		}

		byteOffset += uint(size)
		// Use stdlib to count UTF-16 code units for this rune
		utf16Count += uint32(utf16.RuneLen(r))
	}

	return byteOffset
}

// ByteOffsetToUTF16 converts a UTF-8 byte offset to a UTF-16 code unit offset
func ByteOffsetToUTF16(text string, byteOffset uint) uint32 {
	var currentByte uint = 0
	var utf16Count uint32 = 0

	for currentByte < byteOffset && currentByte < uint(len(text)) {
		r, size := utf8.DecodeRuneInString(text[currentByte:])
		// Check for invalid UTF-8: RuneError with size == 1
		// Note: Valid U+FFFD has size == 3, so we need to check both
		if r == utf8.RuneError && size == 1 {
			// Invalid UTF-8, treat as single byte
			currentByte++
			utf16Count++
			continue
		}

		currentByte += uint(size)
		// Use stdlib to count UTF-16 code units
		utf16Count += uint32(utf16.RuneLen(r))
	}

	return utf16Count
}

// DocumentEdit represents a change to a document
type DocumentEdit struct {
	StartPosition protocol.Position // Start position of the edit
	EndPosition   protocol.Position // End position of the edit (before edit)
	OldLength     uint              // Length of text being replaced
	NewLength     uint              // Length of new text
	Text          string            // New text content
}

// ChangeAnalysis analyzes the characteristics of document changes
type ChangeAnalysis struct {
	TotalChanges         int            // Total number of changes
	LargestChange        uint           // Size of largest single change
	IsSmallChange        bool           // Whether changes are considered "small"
	ShouldUseIncremental bool           // Recommendation for parsing strategy
	Edits                []DocumentEdit // Individual edits extracted from LSP changes
}

// IncrementalParser manages incremental parsing for documents
type IncrementalParser struct {
	strategy ParseStrategy
}

// NewIncrementalParser creates a new incremental parser with the specified strategy
func NewIncrementalParser(strategy ParseStrategy) *IncrementalParser {
	return &IncrementalParser{
		strategy: strategy,
	}
}

// AnalyzeChanges analyzes LSP content changes to determine parsing strategy
func (ip *IncrementalParser) AnalyzeChanges(changes []protocol.TextDocumentContentChangeEvent, oldContent string) *ChangeAnalysis {
	analysis := &ChangeAnalysis{
		TotalChanges: len(changes),
		Edits:        make([]DocumentEdit, 0, len(changes)),
	}

	var totalChangeSize uint = 0

	hasFullDocumentChange := false

	for _, change := range changes {
		var edit DocumentEdit

		if change.Range != nil {
			// Incremental change
			edit = DocumentEdit{
				StartPosition: change.Range.Start,
				EndPosition:   change.Range.End,
				Text:          change.Text,
				NewLength:     uint(len(change.Text)),
			}

			// Calculate old length from range and content
			edit.OldLength = ip.calculateOldLength(change.Range, oldContent)
		} else {
			// Full document replacement
			hasFullDocumentChange = true
			edit = DocumentEdit{
				StartPosition: protocol.Position{Line: 0, Character: 0},
				EndPosition:   protocol.Position{Line: 999999, Character: 999999}, // End of document
				Text:          change.Text,
				NewLength:     uint(len(change.Text)),
				OldLength:     uint(len(oldContent)),
			}
		}

		analysis.Edits = append(analysis.Edits, edit)
		totalChangeSize += max(edit.OldLength, edit.NewLength)

		if edit.OldLength > analysis.LargestChange {
			analysis.LargestChange = edit.OldLength
		}
		if edit.NewLength > analysis.LargestChange {
			analysis.LargestChange = edit.NewLength
		}
	}

	// Heuristics for determining if incremental parsing should be used
	const smallChangeThreshold = 1000  // Characters
	const maxChangesForIncremental = 5 // Number of changes

	analysis.IsSmallChange = analysis.LargestChange < smallChangeThreshold
	analysis.ShouldUseIncremental = analysis.IsSmallChange &&
		analysis.TotalChanges <= maxChangesForIncremental &&
		!hasFullDocumentChange // Never use incremental for full document changes

	helpers.SafeDebugLog("[INCREMENTAL] Change analysis: changes=%d, largest=%d, small=%t, useIncremental=%t",
		analysis.TotalChanges, analysis.LargestChange, analysis.IsSmallChange, analysis.ShouldUseIncremental)

	return analysis
}

// calculateOldLength calculates the byte length of text being replaced
// Uses UTF-16 to UTF-8 conversion for accurate byte-based calculations
func (ip *IncrementalParser) calculateOldLength(changeRange *protocol.Range, content string) uint {
	if changeRange == nil {
		return uint(len(content))
	}

	// Use positionToByteOffset for accurate UTF-16 to UTF-8 conversion
	startByte := ip.positionToByteOffset(changeRange.Start, content)
	endByte := ip.positionToByteOffset(changeRange.End, content)

	if endByte < startByte {
		return 0
	}

	return endByte - startByte
}

// ParseWithStrategy parses a document using the configured strategy
func (ip *IncrementalParser) ParseWithStrategy(doc types.Document, newContent string, changes []protocol.TextDocumentContentChangeEvent) ParseResult {
	if doc.Tree() == nil {
		return ip.performFullParse(doc, newContent)
	}

	strategy := ip.determineStrategy(doc, newContent, changes)

	switch strategy {
	case ParseStrategyIncremental:
		return ip.attemptIncrementalParse(doc, newContent, changes)
	case ParseStrategyFull:
		return ip.performFullParse(doc, newContent)
	case ParseStrategyAuto:
		oldContent, err := doc.Content()
		if err != nil {
			return ParseResult{Success: false, Error: err}
		}
		analysis := ip.AnalyzeChanges(changes, oldContent)
		if analysis.ShouldUseIncremental {
			return ip.attemptIncrementalParse(doc, newContent, changes)
		}
		return ip.performFullParse(doc, newContent)
	default:
		return ip.performFullParse(doc, newContent)
	}
}

// Use ParseResult from types package
type ParseResult = types.ParseResult

// determineStrategy determines which parsing strategy to use
func (ip *IncrementalParser) determineStrategy(doc types.Document, newContent string, changes []protocol.TextDocumentContentChangeEvent) ParseStrategy {
	if ip.strategy != ParseStrategyAuto {
		return ip.strategy
	}

	// Auto strategy logic
	if len(changes) == 0 {
		return ParseStrategyFull
	}

	// Check if we have a full document replacement
	for _, change := range changes {
		if change.Range == nil {
			return ParseStrategyFull
		}
	}

	return ParseStrategyIncremental
}

// attemptIncrementalParse attempts incremental parsing with fallback to full parsing
func (ip *IncrementalParser) attemptIncrementalParse(doc types.Document, newContent string, changes []protocol.TextDocumentContentChangeEvent) ParseResult {
	helpers.SafeDebugLog("[INCREMENTAL] Attempting incremental parse for %s", doc.URI())

	if doc.Tree() == nil {
		helpers.SafeDebugLog("[INCREMENTAL] No existing tree, falling back to full parse")
		return ip.performFullParse(doc, newContent)
	}

	// Keep reference to old tree
	oldTree := doc.Tree()

	// Get old content for edit conversion
	oldContent, err := doc.Content()
	if err != nil {
		helpers.SafeDebugLog("[INCREMENTAL] Failed to get document content: %v", err)
		return ip.performFullParse(doc, newContent)
	}

	// Apply edits to the tree for incremental parsing
	for _, change := range changes {
		if change.Range != nil {
			edit := ip.convertToTreeSitterEdit(change, oldContent)
			oldTree.Edit(&edit)
		}
	}

	// Attempt incremental parsing
	newTree := doc.Parser().Parse([]byte(newContent), oldTree)

	// Validate the incremental parse result
	if ip.validateIncrementalParse(newTree, newContent) {
		helpers.SafeDebugLog("[INCREMENTAL] Incremental parse successful for %s", doc.URI())
		return ParseResult{
			Success:         true,
			UsedIncremental: true,
			NewTree:         newTree,
			OldTree:         oldTree,
		}
	}

	// Incremental parsing failed, fall back to full parsing
	helpers.SafeDebugLog("[INCREMENTAL] Incremental parse failed, falling back to full parse for %s", doc.URI())
	newTree.Close() // Clean up failed incremental tree

	return ip.performFullParse(doc, newContent)
}

// performFullParse performs a full document parse
func (ip *IncrementalParser) performFullParse(doc types.Document, newContent string) ParseResult {
	helpers.SafeDebugLog("[INCREMENTAL] Performing full parse for %s", doc.URI())

	oldTree := doc.Tree()
	newTree := doc.Parser().Parse([]byte(newContent), nil)

	return ParseResult{
		Success:         true,
		UsedIncremental: false,
		NewTree:         newTree,
		OldTree:         oldTree,
	}
}

// convertToTreeSitterEdit converts LSP change to tree-sitter edit
// LSP uses UTF-16 for positions, tree-sitter uses UTF-8 byte offsets
func (ip *IncrementalParser) convertToTreeSitterEdit(change protocol.TextDocumentContentChangeEvent, oldContent string) ts.InputEdit {
	if change.Range == nil {
		return ts.InputEdit{}
	}

	// Convert positions to byte offsets
	startByte := ip.positionToByteOffset(change.Range.Start, oldContent)
	oldEndByte := ip.positionToByteOffset(change.Range.End, oldContent)
	newEndByte := startByte + uint(len(change.Text))

	// Convert LSP positions (UTF-16) to tree-sitter Points (UTF-8 byte columns)
	oldLines := splitLines(oldContent)
	startPoint := ip.lspPositionToTreeSitterPoint(change.Range.Start, oldLines)
	oldEndPoint := ip.lspPositionToTreeSitterPoint(change.Range.End, oldLines)
	newEndPoint := ip.calculateNewEndPoint(change.Range.Start, change.Text, oldLines)

	return ts.InputEdit{
		StartByte:      startByte,
		OldEndByte:     oldEndByte,
		NewEndByte:     newEndByte,
		StartPosition:  startPoint,
		OldEndPosition: oldEndPoint,
		NewEndPosition: newEndPoint,
	}
}

// lspPositionToTreeSitterPoint converts LSP position (UTF-16) to tree-sitter Point (UTF-8 byte column)
func (ip *IncrementalParser) lspPositionToTreeSitterPoint(pos protocol.Position, lines []string) ts.Point {
	if pos.Line >= uint32(len(lines)) {
		return ts.Point{Row: uint(pos.Line), Column: 0}
	}

	line := lines[pos.Line]
	// Convert UTF-16 character offset to UTF-8 byte offset within the line
	byteColumn := UTF16ToByteOffset(line, pos.Character)

	return ts.Point{
		Row:    uint(pos.Line),
		Column: byteColumn,
	}
}

// positionToByteOffset converts LSP position (UTF-16) to byte offset (UTF-8)
func (ip *IncrementalParser) positionToByteOffset(pos protocol.Position, content string) uint {
	lines := splitLines(content)
	var offset uint = 0

	// Add bytes for complete lines before the target line
	for i := uint32(0); i < pos.Line && i < uint32(len(lines)); i++ {
		offset += uint(len(lines[i])) + 1 // +1 for newline
	}

	// Add bytes for characters in the target line
	if pos.Line < uint32(len(lines)) {
		line := lines[pos.Line]
		// Convert UTF-16 character offset to UTF-8 byte offset within the line
		byteOffsetInLine := UTF16ToByteOffset(line, pos.Character)
		offset += byteOffsetInLine
	}

	return offset
}

// PositionToByteOffset is the exported version for testing
func (ip *IncrementalParser) PositionToByteOffset(pos protocol.Position, content string) uint {
	return ip.positionToByteOffset(pos, content)
}

// calculateNewEndPoint calculates the end point after applying an edit
// LSP startPos is in UTF-16, tree-sitter Point.Column needs to be in UTF-8 bytes
func (ip *IncrementalParser) calculateNewEndPoint(startPos protocol.Position, newText string, oldLines []string) ts.Point {
	newLines := splitLines(newText)
	if len(newLines) == 0 {
		// No text inserted, return start position converted to byte column
		return ip.lspPositionToTreeSitterPoint(startPos, oldLines)
	}

	if len(newLines) == 1 {
		// Single line insertion - stay on same row
		// New column = byte offset of start position + byte length of new text
		startByteColumn := uint(0)
		if startPos.Line < uint32(len(oldLines)) {
			line := oldLines[startPos.Line]
			startByteColumn = UTF16ToByteOffset(line, startPos.Character)
		}

		return ts.Point{
			Row:    uint(startPos.Line),
			Column: startByteColumn + uint(len(newText)), // Byte offset
		}
	}

	// Multi-line insertion - new row is startLine + number of new lines - 1
	// New column is byte length of last new line (since it's a complete new line)
	lastLine := newLines[len(newLines)-1]
	return ts.Point{
		Row:    uint(startPos.Line) + uint(len(newLines)-1),
		Column: uint(len(lastLine)), // Byte length of last line
	}
}

// validateIncrementalParse validates that incremental parsing succeeded
func (ip *IncrementalParser) validateIncrementalParse(tree *ts.Tree, content string) bool {
	if tree == nil {
		return false
	}

	root := tree.RootNode()
	if root == nil {
		return false
	}

	// Check if tree spans the entire content
	contentBytes := uint(len(content))
	if root.EndByte() < contentBytes {
		// Tree doesn't cover all content, likely failed
		helpers.SafeDebugLog("[INCREMENTAL] Validation failed: tree end (%d) < content length (%d)",
			root.EndByte(), contentBytes)
		return false
	}

	// Check for error nodes (indicates parsing issues)
	if ip.hasErrorNodes(root) {
		helpers.SafeDebugLog("[INCREMENTAL] Validation failed: tree contains error nodes")
		return false
	}

	return true
}

// hasErrorNodes checks if the tree contains any error nodes
func (ip *IncrementalParser) hasErrorNodes(node *ts.Node) bool {
	if node.IsError() {
		return true
	}

	// Check children for errors
	for i := uint(0); i < node.ChildCount(); i++ {
		child := node.Child(i)
		if child != nil && ip.hasErrorNodes(child) {
			return true
		}
	}

	return false
}

// splitLines splits content into lines
func splitLines(content string) []string {
	if content == "" {
		return []string{}
	}

	lines := make([]string, 0)
	start := 0

	for i, r := range content {
		if r == '\n' {
			lines = append(lines, content[start:i])
			start = i + 1
		}
	}

	// Add the last line if it doesn't end with newline
	if start < len(content) {
		lines = append(lines, content[start:])
	}

	return lines
}

