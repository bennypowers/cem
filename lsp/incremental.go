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

// calculateOldLength calculates the length of text being replaced
func (ip *IncrementalParser) calculateOldLength(changeRange *protocol.Range, content string) uint {
	lines := splitLines(content)
	if changeRange.Start.Line >= uint32(len(lines)) {
		return 0
	}

	startLine := changeRange.Start.Line
	endLine := changeRange.End.Line

	if endLine >= uint32(len(lines)) {
		endLine = uint32(len(lines)) - 1
	}

	if startLine == endLine {
		// Single line change
		line := lines[startLine]
		if changeRange.End.Character <= uint32(len(line)) {
			return uint(changeRange.End.Character - changeRange.Start.Character)
		}
		return uint(len(line)) - uint(changeRange.Start.Character)
	}

	// Multi-line change
	var totalLength uint = 0

	// First line: from start character to end of line
	if changeRange.Start.Character < uint32(len(lines[startLine])) {
		totalLength += uint(len(lines[startLine])) - uint(changeRange.Start.Character)
	}
	totalLength += 1 // newline

	// Middle lines: full lines
	for i := startLine + 1; i < endLine; i++ {
		totalLength += uint(len(lines[i])) + 1 // +1 for newline
	}

	// Last line: from start to end character
	if endLine > startLine && changeRange.End.Character <= uint32(len(lines[endLine])) {
		totalLength += uint(changeRange.End.Character)
	}

	return totalLength
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
func (ip *IncrementalParser) convertToTreeSitterEdit(change protocol.TextDocumentContentChangeEvent, oldContent string) ts.InputEdit {
	if change.Range == nil {
		return ts.InputEdit{}
	}

	// Convert positions to byte offsets
	startByte := ip.positionToByteOffset(change.Range.Start, oldContent)
	oldEndByte := ip.positionToByteOffset(change.Range.End, oldContent)
	newEndByte := startByte + uint(len(change.Text))

	return ts.InputEdit{
		StartByte:  startByte,
		OldEndByte: oldEndByte,
		NewEndByte: newEndByte,
		StartPosition: ts.Point{
			Row:    uint(change.Range.Start.Line),
			Column: uint(change.Range.Start.Character),
		},
		OldEndPosition: ts.Point{
			Row:    uint(change.Range.End.Line),
			Column: uint(change.Range.End.Character),
		},
		NewEndPosition: ip.calculateNewEndPoint(change.Range.Start, change.Text),
	}
}

// positionToByteOffset converts LSP position to byte offset
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
		if pos.Character < uint32(len(line)) {
			offset += uint(pos.Character)
		} else {
			offset += uint(len(line))
		}
	}

	return offset
}

// calculateNewEndPoint calculates the end point after applying an edit
func (ip *IncrementalParser) calculateNewEndPoint(startPos protocol.Position, newText string) ts.Point {
	lines := splitLines(newText)
	if len(lines) == 0 {
		return ts.Point{Row: uint(startPos.Line), Column: uint(startPos.Character)}
	}

	if len(lines) == 1 {
		// Single line
		return ts.Point{
			Row:    uint(startPos.Line),
			Column: uint(startPos.Character) + uint(len(newText)),
		}
	}

	// Multi-line
	return ts.Point{
		Row:    uint(startPos.Line) + uint(len(lines)-1),
		Column: uint(len(lines[len(lines)-1])),
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

