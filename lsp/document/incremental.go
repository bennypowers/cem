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
package document

import (
	"strings"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	protocol "github.com/tliron/glsp/protocol_3_16"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// incrementalParser implements the types.IncrementalParser interface
type incrementalParser struct {
	strategy types.ParseStrategy
}

// newIncrementalParser creates a new incremental parser
func newIncrementalParser(strategy types.ParseStrategy) types.IncrementalParser {
	return &incrementalParser{
		strategy: strategy,
	}
}

// ParseWithStrategy parses a document using the configured strategy
func (ip *incrementalParser) ParseWithStrategy(doc types.Document, newContent string, changes []protocol.TextDocumentContentChangeEvent) types.ParseResult {
	if doc.Tree() == nil {
		return ip.performFullParse(doc, newContent)
	}

	strategy := ip.determineStrategy(doc, newContent, changes)

	switch strategy {
	case types.ParseStrategyIncremental:
		return ip.attemptIncrementalParse(doc, newContent, changes)
	case types.ParseStrategyFull:
		return ip.performFullParse(doc, newContent)
	case types.ParseStrategyAuto:
		oldContent, err := doc.Content()
		if err != nil {
			return types.ParseResult{Success: false, Error: err}
		}
		analysis := ip.analyzeChanges(changes, oldContent)
		if analysis.ShouldUseIncremental {
			return ip.attemptIncrementalParse(doc, newContent, changes)
		}
		return ip.performFullParse(doc, newContent)
	default:
		return ip.performFullParse(doc, newContent)
	}
}

// determineStrategy determines which parsing strategy to use
func (ip *incrementalParser) determineStrategy(doc types.Document, newContent string, changes []protocol.TextDocumentContentChangeEvent) types.ParseStrategy {
	if ip.strategy != types.ParseStrategyAuto {
		return ip.strategy
	}

	// Auto strategy logic
	if len(changes) == 0 {
		return types.ParseStrategyFull
	}

	// Check if we have a full document replacement
	for _, change := range changes {
		if change.Range == nil {
			return types.ParseStrategyFull
		}
	}

	return types.ParseStrategyIncremental
}

// changeAnalysis analyzes the characteristics of document changes
type changeAnalysis struct {
	TotalChanges         int  // Total number of changes
	LargestChange        uint // Size of largest single change
	IsSmallChange        bool // Whether changes are considered "small"
	ShouldUseIncremental bool // Recommendation for parsing strategy
}

// analyzeChanges analyzes LSP content changes to determine parsing strategy
func (ip *incrementalParser) analyzeChanges(changes []protocol.TextDocumentContentChangeEvent, oldContent string) *changeAnalysis {
	analysis := &changeAnalysis{
		TotalChanges: len(changes),
	}

	var totalChangeSize uint = 0
	hasFullDocumentChange := false

	for _, change := range changes {
		var oldLength, newLength uint

		if change.Range != nil {
			// Incremental change
			oldLength = ip.calculateOldLength(change.Range, oldContent)
			newLength = uint(len(change.Text))
		} else {
			// Full document replacement
			hasFullDocumentChange = true
			oldLength = uint(len(oldContent))
			newLength = uint(len(change.Text))
		}

		totalChangeSize += max(oldLength, newLength)

		if oldLength > analysis.LargestChange {
			analysis.LargestChange = oldLength
		}
		if newLength > analysis.LargestChange {
			analysis.LargestChange = newLength
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
func (ip *incrementalParser) calculateOldLength(changeRange *protocol.Range, content string) uint {
	lines := strings.Split(content, "\n")
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

// attemptIncrementalParse attempts incremental parsing with fallback to full parsing
func (ip *incrementalParser) attemptIncrementalParse(doc types.Document, newContent string, changes []protocol.TextDocumentContentChangeEvent) types.ParseResult {
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
		return types.ParseResult{
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
func (ip *incrementalParser) performFullParse(doc types.Document, newContent string) types.ParseResult {
	helpers.SafeDebugLog("[INCREMENTAL] Performing full parse for %s", doc.URI())

	oldTree := doc.Tree()
	newTree := doc.Parser().Parse([]byte(newContent), nil)

	return types.ParseResult{
		Success:         true,
		UsedIncremental: false,
		NewTree:         newTree,
		OldTree:         oldTree,
	}
}

// convertToTreeSitterEdit converts LSP change to tree-sitter edit
func (ip *incrementalParser) convertToTreeSitterEdit(change protocol.TextDocumentContentChangeEvent, oldContent string) ts.InputEdit {
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
func (ip *incrementalParser) positionToByteOffset(pos protocol.Position, content string) uint {
	lines := strings.Split(content, "\n")
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
func (ip *incrementalParser) calculateNewEndPoint(startPos protocol.Position, newText string) ts.Point {
	lines := strings.Split(newText, "\n")
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
func (ip *incrementalParser) validateIncrementalParse(tree *ts.Tree, content string) bool {
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
func (ip *incrementalParser) hasErrorNodes(node *ts.Node) bool {
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

// max returns the maximum of two uint values
func max(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}