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

	"bennypowers.dev/cem/internal/textutil"
	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	"go.lsp.dev/protocol"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// changeRange returns the range of a partial change, or nil for a whole-document change.
func changeRange(change protocol.TextDocumentContentChangeEvent) *protocol.Range {
	if partial, ok := change.(*protocol.TextDocumentContentChangePartial); ok {
		r := partial.Range
		return &r
	}
	return nil
}

// changeText returns the replacement text from a content change event.
func changeText(change protocol.TextDocumentContentChangeEvent) string {
	switch c := change.(type) {
	case *protocol.TextDocumentContentChangePartial:
		return c.Text
	case *protocol.TextDocumentContentChangeWholeDocument:
		return c.Text
	default:
		return ""
	}
}

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
		if changeRange(change) == nil {
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

		if r := changeRange(change); r != nil {
			// Incremental change
			oldLength = ip.calculateOldLength(r, oldContent)
			newLength = uint(len(changeText(change)))
		} else {
			// Full document replacement
			hasFullDocumentChange = true
			oldLength = uint(len(oldContent))
			newLength = uint(len(changeText(change)))
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

	if doc.Parser() == nil {
		helpers.SafeDebugLog("[INCREMENTAL] No parser available, falling back to full recreation")
		return types.ParseResult{Success: false}
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
		if r := changeRange(change); r != nil {
			edit := ip.convertToTreeSitterEdit(r, changeText(change), oldContent)
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

	if doc.Parser() == nil {
		return types.ParseResult{Success: false}
	}

	oldTree := doc.Tree()
	newTree := doc.Parser().Parse([]byte(newContent), nil)

	return types.ParseResult{
		Success:         true,
		UsedIncremental: false,
		NewTree:         newTree,
		OldTree:         oldTree,
	}
}

// convertToTreeSitterEdit converts LSP change range and text to tree-sitter edit
func (ip *incrementalParser) convertToTreeSitterEdit(r *protocol.Range, text string, oldContent string) ts.InputEdit {
	if r == nil {
		return ts.InputEdit{}
	}

	startByte := ip.positionToByteOffset(r.Start, oldContent)
	oldEndByte := ip.positionToByteOffset(r.End, oldContent)
	newEndByte := startByte + uint(len(text))

	oldLines := strings.Split(oldContent, "\n")

	return ts.InputEdit{
		StartByte:      startByte,
		OldEndByte:     oldEndByte,
		NewEndByte:     newEndByte,
		StartPosition:  ip.lspPositionToTreeSitterPoint(r.Start, oldLines),
		OldEndPosition: ip.lspPositionToTreeSitterPoint(r.End, oldLines),
		NewEndPosition: ip.calculateNewEndPoint(r.Start, text, oldLines),
	}
}

// positionToByteOffset converts LSP position (UTF-16 code units) to byte offset (UTF-8)
func (ip *incrementalParser) positionToByteOffset(pos protocol.Position, content string) uint {
	lines := strings.Split(content, "\n")
	var offset uint = 0

	for i := uint32(0); i < pos.Line && i < uint32(len(lines)); i++ {
		offset += uint(len(lines[i])) + 1
	}

	if pos.Line < uint32(len(lines)) {
		line := lines[pos.Line]
		offset += textutil.UTF16ToByteOffset(line, pos.Character)
	}

	return offset
}

// lspPositionToTreeSitterPoint converts LSP position (UTF-16) to tree-sitter Point (UTF-8 byte column)
func (ip *incrementalParser) lspPositionToTreeSitterPoint(pos protocol.Position, lines []string) ts.Point {
	if pos.Line >= uint32(len(lines)) {
		return ts.Point{Row: uint(pos.Line), Column: 0}
	}
	line := lines[pos.Line]
	return ts.Point{
		Row:    uint(pos.Line),
		Column: textutil.UTF16ToByteOffset(line, pos.Character),
	}
}

// calculateNewEndPoint calculates the end point after applying an edit.
// Positions use UTF-8 byte offsets for tree-sitter Points.
func (ip *incrementalParser) calculateNewEndPoint(startPos protocol.Position, newText string, oldLines []string) ts.Point {
	newLines := strings.Split(newText, "\n")
	if len(newLines) == 0 {
		return ip.lspPositionToTreeSitterPoint(startPos, oldLines)
	}

	if len(newLines) == 1 {
		startByteColumn := uint(0)
		if startPos.Line < uint32(len(oldLines)) {
			startByteColumn = textutil.UTF16ToByteOffset(oldLines[startPos.Line], startPos.Character)
		}
		return ts.Point{
			Row:    uint(startPos.Line),
			Column: startByteColumn + uint(len(newText)),
		}
	}

	lastLine := newLines[len(newLines)-1]
	return ts.Point{
		Row:    uint(startPos.Line) + uint(len(newLines)-1),
		Column: uint(len(lastLine)),
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
