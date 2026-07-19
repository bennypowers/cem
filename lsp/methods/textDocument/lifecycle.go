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
	"strings"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/methods/textDocument/publishDiagnostics"
	"bennypowers.dev/cem/lsp/types"
	"go.lsp.dev/protocol"
)

// DocumentManager interface for document lifecycle operations
type DocumentManager interface {
	OpenDocument(uri, content string, version int32) types.Document
	UpdateDocument(uri, content string, version int32) types.Document
	UpdateDocumentWithChanges(uri, content string, version int32, changes []protocol.TextDocumentContentChangeEvent) types.Document
	CloseDocument(uri string)
	Document(uri string) types.Document
}

// DidOpen handles textDocument/didOpen notifications
func DidOpen(ctx types.ServerContext, params *protocol.DidOpenTextDocumentParams) error {
	uri := string(params.TextDocument.URI)
	helpers.SafeDebugLog("[LIFECYCLE] DidOpen: URI=%s, Version=%d, ContentLength=%d",
		uri, params.TextDocument.Version, len(params.TextDocument.Text))

	dm, err := ctx.DocumentManager()
	if err != nil {
		return err
	}

	doc := dm.OpenDocument(
		uri,
		params.TextDocument.Text,
		params.TextDocument.Version,
	)

	if doc != nil {
		helpers.SafeDebugLog("[LIFECYCLE] Successfully opened document: %s", uri)

		// Synthesize ephemeral element declarations before diagnostics,
		// so locally-defined elements are available for diagnostic checks
		ctx.SynthesizeEphemeralElements(uri)

		if !ctx.UsePullDiagnostics() {
			if err := publishDiagnostics.PublishDiagnostics(ctx, uri); err != nil {
				helpers.SafeDebugLog("[LIFECYCLE] Failed to publish diagnostics for %s: %v", uri, err)
			}
		}
	} else {
		helpers.SafeDebugLog("[LIFECYCLE] Failed to open document: %s", uri)
	}

	return nil
}

// DidChange handles textDocument/didChange notifications
func DidChange(ctx types.ServerContext, params *protocol.DidChangeTextDocumentParams) error {
	uri := string(params.TextDocument.URI)
	helpers.SafeDebugLog("[LIFECYCLE] DidChange: URI=%s, Version=%d, Changes=%d",
		uri, params.TextDocument.Version, len(params.ContentChanges))

	dm, err := ctx.DocumentManager()
	if err != nil {
		return err
	}

	// Handle incremental changes
	doc := dm.Document(uri)
	if doc == nil {
		helpers.SafeDebugLog("[LIFECYCLE] No existing document found for URI: %s", uri)
		return nil
	}
	helpers.SafeDebugLog("[LIFECYCLE] Found existing document for URI: %s", uri)

	// Get current document content
	currentContent, err := doc.Content()
	if err != nil {
		helpers.SafeDebugLog("[LIFECYCLE] Error getting document content: %v", err)
		return nil
	}
	helpers.SafeDebugLog("[LIFECYCLE] Current document content length: %d", len(currentContent))

	// Apply content changes (can be incremental or full document)
	newContent := currentContent
	for i, change := range params.ContentChanges {
		helpers.SafeDebugLog("[LIFECYCLE] Processing change %d of %d", i+1, len(params.ContentChanges))

		switch c := change.(type) {
		case *protocol.TextDocumentContentChangePartial:
			newContent = applyIncrementalChange(newContent, c)
			helpers.SafeDebugLog("[LIFECYCLE] Applied incremental change, new length=%d", len(newContent))
		case *protocol.TextDocumentContentChangeWholeDocument:
			newContent = c.Text
			helpers.SafeDebugLog("[LIFECYCLE] Full document replacement (length=%d)", len(newContent))
		}
	}

	helpers.SafeDebugLog("[LIFECYCLE] Updating document with content (length=%d)", len(newContent))

	updatedDoc := dm.UpdateDocumentWithChanges(
		uri,
		newContent,
		params.TextDocument.Version,
		params.ContentChanges,
	)

	if updatedDoc != nil {
		helpers.SafeDebugLog("[LIFECYCLE] Successfully updated document: %s", uri)

		// Synthesize ephemeral element declarations before diagnostics
		ctx.SynthesizeEphemeralElements(uri)

		if !ctx.UsePullDiagnostics() {
			if err := publishDiagnostics.PublishDiagnostics(ctx, uri); err != nil {
				helpers.SafeDebugLog("[LIFECYCLE] Failed to publish diagnostics for %s: %v", uri, err)
			}
		}
	} else {
		helpers.SafeDebugLog("[LIFECYCLE] Failed to update document: %s", uri)
	}

	return nil
}

// DidClose handles textDocument/didClose notifications
func DidClose(ctx types.ServerContext, params *protocol.DidCloseTextDocumentParams) error {
	uri := string(params.TextDocument.URI)
	dm, err := ctx.DocumentManager()
	if err != nil {
		return err
	}
	dm.CloseDocument(uri)

	// Clean up ephemeral data for the closed document.
	// SynthesizeEphemeralElements will see doc == nil (already closed)
	// and remove any stale ephemeral data for this URI.
	ctx.SynthesizeEphemeralElements(uri)

	return nil
}

// applyIncrementalChange applies an incremental text change to existing content
func applyIncrementalChange(content string, change *protocol.TextDocumentContentChangePartial) string {
	lines := strings.Split(content, "\n")
	changeParams := extractChangeParameters(change, lines)

	helpers.SafeDebugLog("[LIFECYCLE] Applying incremental change: range=start(%d,%d) end(%d,%d), text='%s'",
		changeParams.startLine, changeParams.startChar, changeParams.endLine, changeParams.endChar, change.Text)

	// Handle out-of-bounds start line
	if changeParams.startLine >= len(lines) {
		return handleOutOfBoundsChange(lines, changeParams.startLine, change.Text)
	}

	// Adjust end line if out of bounds
	changeParams = adjustEndLineIfNeeded(changeParams, lines)

	// Apply the change based on whether it's single or multi-line
	if changeParams.startLine == changeParams.endLine {
		return applySingleLineChange(lines, changeParams, change.Text)
	}

	return applyMultiLineChange(lines, changeParams, change.Text)
}

// changeParameters holds the extracted and validated change parameters
type changeParameters struct {
	startLine int
	startChar int
	endLine   int
	endChar   int
}

// extractChangeParameters extracts and validates change parameters from the LSP change event
func extractChangeParameters(change *protocol.TextDocumentContentChangePartial, lines []string) changeParameters {
	return changeParameters{
		startLine: int(change.Range.Start.Line),
		startChar: int(change.Range.Start.Character),
		endLine:   int(change.Range.End.Line),
		endChar:   int(change.Range.End.Character),
	}
}

// handleOutOfBoundsChange handles changes where the start line is beyond the document end
func handleOutOfBoundsChange(lines []string, startLine int, text string) string {
	helpers.SafeDebugLog("[LIFECYCLE] Start line %d beyond document end (%d lines), appending", startLine, len(lines))
	// Pad with empty lines if needed
	for len(lines) <= startLine {
		lines = append(lines, "")
	}
	lines[startLine] = text
	return strings.Join(lines, "\n")
}

// adjustEndLineIfNeeded adjusts the end line and character if they're beyond the document bounds
func adjustEndLineIfNeeded(params changeParameters, lines []string) changeParameters {
	if params.endLine >= len(lines) {
		helpers.SafeDebugLog("[LIFECYCLE] End line %d beyond document end (%d lines), adjusting", params.endLine, len(lines))
		params.endLine = len(lines) - 1
		params.endChar = len(lines[params.endLine])
	}
	return params
}

// applySingleLineChange applies a change that occurs within a single line
func applySingleLineChange(lines []string, params changeParameters, text string) string {
	line := lines[params.startLine]

	// Clamp character positions to line bounds
	if params.startChar > len(line) {
		params.startChar = len(line)
	}
	if params.endChar > len(line) {
		params.endChar = len(line)
	}

	// Replace the range in the line
	newLine := line[:params.startChar] + text + line[params.endChar:]
	lines[params.startLine] = newLine
	helpers.SafeDebugLog("[LIFECYCLE] Single line change: '%s' -> '%s'", line, newLine)

	return strings.Join(lines, "\n")
}

// applyMultiLineChange applies a change that spans multiple lines
func applyMultiLineChange(lines []string, params changeParameters, text string) string {
	var newLines []string

	// Add lines before the change
	newLines = append(newLines, lines[:params.startLine]...)

	// Get start and end line content
	startLineContent := lines[params.startLine]
	if params.startChar > len(startLineContent) {
		params.startChar = len(startLineContent)
	}

	endLineContent := ""
	if params.endLine < len(lines) {
		endLineContent = lines[params.endLine]
		if params.endChar <= len(endLineContent) {
			endLineContent = endLineContent[params.endChar:]
		}
	}

	// Apply the change text
	newLines = appendChangeText(newLines, startLineContent, endLineContent, params.startChar, text)

	// Add remaining lines after the change
	if params.endLine+1 < len(lines) {
		newLines = append(newLines, lines[params.endLine+1:]...)
	}

	helpers.SafeDebugLog("[LIFECYCLE] Multi-line change: %d lines -> %d lines", len(lines), len(newLines))
	return strings.Join(newLines, "\n")
}

// appendChangeText adds the change text to the new lines, handling single and multi-line replacements
func appendChangeText(newLines []string, startLineContent, endLineContent string, startChar int, text string) []string {
	changeLines := strings.Split(text, "\n")

	if len(changeLines) == 1 {
		// Single line replacement
		newLine := startLineContent[:startChar] + text + endLineContent
		newLines = append(newLines, newLine)
	} else {
		// Multi-line replacement
		// First line of change
		newLines = append(newLines, startLineContent[:startChar]+changeLines[0])
		// Middle lines of change
		newLines = append(newLines, changeLines[1:len(changeLines)-1]...)
		// Last line of change
		newLines = append(newLines, changeLines[len(changeLines)-1]+endLineContent)
	}

	return newLines
}
