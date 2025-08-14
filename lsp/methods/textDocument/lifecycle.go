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
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// LifecycleContext provides the dependencies needed for document lifecycle
type LifecycleContext interface {
	TextDocumentManager() DocumentManager
	publishDiagnostics.DiagnosticsContext
}

// DocumentManager interface for document lifecycle operations
type DocumentManager interface {
	OpenDocument(uri, content string, version int32) types.Document
	UpdateDocument(uri, content string, version int32) types.Document
	CloseDocument(uri string)
	Document(uri string) types.Document
}

// DidOpen handles textDocument/didOpen notifications
func DidOpen(ctx LifecycleContext, context *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	helpers.SafeDebugLog("[LIFECYCLE] DidOpen: URI=%s, Version=%d, ContentLength=%d",
		params.TextDocument.URI, params.TextDocument.Version, len(params.TextDocument.Text))

	dm := ctx.TextDocumentManager()
	doc := dm.OpenDocument(
		params.TextDocument.URI,
		params.TextDocument.Text,
		params.TextDocument.Version,
	)

	if doc != nil {
		helpers.SafeDebugLog("[LIFECYCLE] Successfully opened document: %s", params.TextDocument.URI)

		// Trigger diagnostics for the newly opened document
		if err := publishDiagnostics.PublishDiagnostics(ctx, context, params.TextDocument.URI); err != nil {
			helpers.SafeDebugLog("[LIFECYCLE] Failed to publish diagnostics for %s: %v", params.TextDocument.URI, err)
		}
	} else {
		helpers.SafeDebugLog("[LIFECYCLE] Failed to open document: %s", params.TextDocument.URI)
	}

	return nil
}

// DidChange handles textDocument/didChange notifications
func DidChange(ctx LifecycleContext, context *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	helpers.SafeDebugLog("[LIFECYCLE] DidChange: URI=%s, Version=%d, Changes=%d",
		params.TextDocument.URI, params.TextDocument.Version, len(params.ContentChanges))

	dm := ctx.TextDocumentManager()

	// Handle incremental changes
	doc := dm.Document(params.TextDocument.URI)
	if doc == nil {
		helpers.SafeDebugLog("[LIFECYCLE] No existing document found for URI: %s", params.TextDocument.URI)
		return nil
	}
	helpers.SafeDebugLog("[LIFECYCLE] Found existing document for URI: %s", params.TextDocument.URI)

	// Get current document content
	var currentContent string
	if docWithContent, ok := doc.(interface{ Content() string }); ok {
		currentContent = docWithContent.Content()
		helpers.SafeDebugLog("[LIFECYCLE] Current document content length: %d", len(currentContent))
	} else {
		helpers.SafeDebugLog("[LIFECYCLE] Document doesn't implement Content interface")
		return nil
	}

	// Apply content changes (can be incremental or full document)
	newContent := currentContent
	for i, change := range params.ContentChanges {
		if changeEvent, ok := change.(protocol.TextDocumentContentChangeEvent); ok {
			helpers.SafeDebugLog("[LIFECYCLE] Processing change %d of %d", i+1, len(params.ContentChanges))

			// Check if this is a full document change (no range specified)
			if changeEvent.Range == nil {
				// Full document replacement
				newContent = changeEvent.Text
				helpers.SafeDebugLog("[LIFECYCLE] Full document replacement (length=%d)", len(newContent))
			} else {
				// Incremental change - apply to current content
				newContent = applyIncrementalChange(newContent, changeEvent)
				helpers.SafeDebugLog("[LIFECYCLE] Applied incremental change, new length=%d", len(newContent))
			}
		} else {
			helpers.SafeDebugLog("[LIFECYCLE] Failed to cast content change %d to TextDocumentContentChangeEvent", i+1)
		}
	}

	helpers.SafeDebugLog("[LIFECYCLE] Updating document with content (length=%d)", len(newContent))
	updatedDoc := dm.UpdateDocument(
		params.TextDocument.URI,
		newContent,
		params.TextDocument.Version,
	)

	if updatedDoc != nil {
		helpers.SafeDebugLog("[LIFECYCLE] Successfully updated document: %s", params.TextDocument.URI)

		// Trigger diagnostics for the updated document
		if err := publishDiagnostics.PublishDiagnostics(ctx, context, params.TextDocument.URI); err != nil {
			helpers.SafeDebugLog("[LIFECYCLE] Failed to publish diagnostics for %s: %v", params.TextDocument.URI, err)
		}
	} else {
		helpers.SafeDebugLog("[LIFECYCLE] Failed to update document: %s", params.TextDocument.URI)
	}

	return nil
}

// DidClose handles textDocument/didClose notifications
func DidClose(ctx LifecycleContext, context *glsp.Context, params *protocol.DidCloseTextDocumentParams) error {
	dm := ctx.TextDocumentManager()
	dm.CloseDocument(params.TextDocument.URI)
	return nil
}

// applyIncrementalChange applies an incremental text change to existing content
func applyIncrementalChange(content string, change protocol.TextDocumentContentChangeEvent) string {
	if change.Range == nil {
		// This should not happen in incremental changes, but handle it safely
		return change.Text
	}

	lines := strings.Split(content, "\n")

	startLine := int(change.Range.Start.Line)
	startChar := int(change.Range.Start.Character)
	endLine := int(change.Range.End.Line)
	endChar := int(change.Range.End.Character)

	helpers.SafeDebugLog("[LIFECYCLE] Applying incremental change: range=start(%d,%d) end(%d,%d), text='%s'",
		startLine, startChar, endLine, endChar, change.Text)

	// Handle out-of-bounds cases safely
	if startLine >= len(lines) {
		helpers.SafeDebugLog("[LIFECYCLE] Start line %d beyond document end (%d lines), appending", startLine, len(lines))
		// Pad with empty lines if needed
		for len(lines) <= startLine {
			lines = append(lines, "")
		}
		lines[startLine] = change.Text
		return strings.Join(lines, "\n")
	}

	if endLine >= len(lines) {
		helpers.SafeDebugLog("[LIFECYCLE] End line %d beyond document end (%d lines), adjusting", endLine, len(lines))
		endLine = len(lines) - 1
		endChar = len(lines[endLine])
	}

	// Handle single line change
	if startLine == endLine {
		line := lines[startLine]
		if startChar > len(line) {
			startChar = len(line)
		}
		if endChar > len(line) {
			endChar = len(line)
		}

		// Replace the range in the line
		newLine := line[:startChar] + change.Text + line[endChar:]
		lines[startLine] = newLine
		helpers.SafeDebugLog("[LIFECYCLE] Single line change: '%s' -> '%s'", line, newLine)
	} else {
		// Multi-line change
		var newLines []string

		// Add lines before the change
		newLines = append(newLines, lines[:startLine]...)

		// Add the first line with the change
		startLineContent := lines[startLine]
		if startChar > len(startLineContent) {
			startChar = len(startLineContent)
		}

		// Add the last line after the change
		endLineContent := ""
		if endLine < len(lines) {
			endLineContent = lines[endLine]
			if endChar > len(endLineContent) {
				endChar = len(endLineContent)
			} else {
				endLineContent = endLineContent[endChar:]
			}
		}

		// Combine the change text with surrounding content
		changeLines := strings.Split(change.Text, "\n")
		if len(changeLines) == 1 {
			// Single line replacement
			newLine := startLineContent[:startChar] + change.Text + endLineContent
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

		// Add remaining lines after the change
		if endLine+1 < len(lines) {
			newLines = append(newLines, lines[endLine+1:]...)
		}

		lines = newLines
		helpers.SafeDebugLog("[LIFECYCLE] Multi-line change: %d lines -> %d lines", len(strings.Split(content, "\n")), len(lines))
	}

	return strings.Join(lines, "\n")
}
