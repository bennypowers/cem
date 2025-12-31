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
	"fmt"
	"sync"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	protocol "github.com/tliron/glsp/protocol_3_16"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// BaseDocument provides common document functionality that language-specific documents can embed
type BaseDocument struct {
	uri        string
	content    string
	version    int32
	language   string
	tree       *ts.Tree
	parser     *ts.Parser
	scriptTags []types.ScriptTag
	mu         sync.RWMutex
}

// NewBaseDocument creates a new base document
func NewBaseDocument(uri, content string, version int32, language string) *BaseDocument {
	return &BaseDocument{
		uri:      uri,
		content:  content,
		version:  version,
		language: language,
	}
}

// URI returns the document URI
func (d *BaseDocument) URI() string {
	if d == nil {
		return ""
	}

	defer func() {
		if r := recover(); r != nil {
			helpers.SafeDebugLog("[DOCUMENT] PANIC in URI(): %v", r)
		}
	}()

	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.uri
}

// Content returns the document content
func (d *BaseDocument) Content() (string, error) {
	// Defensive programming: check for nil pointer
	if d == nil {
		return "", fmt.Errorf("document is nil")
	}

	// Additional safety check - ensure the mutex is properly initialized
	// This is a workaround for potential concurrent access issues
	defer func() {
		if r := recover(); r != nil {
			helpers.SafeDebugLog("[DOCUMENT] PANIC in Content(): %v", r)
			// Log document state for debugging
			helpers.SafeDebugLog("[DOCUMENT] Document state: uri=%s, content_len=%d", d.uri, len(d.content))
		}
	}()

	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.content, nil
}

// Version returns the document version
func (d *BaseDocument) Version() int32 {
	if d == nil {
		return 0
	}

	defer func() {
		if r := recover(); r != nil {
			helpers.SafeDebugLog("[DOCUMENT] PANIC in Version(): %v", r)
		}
	}()

	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.version
}

// Language returns the document language
func (d *BaseDocument) Language() string {
	if d == nil {
		return ""
	}

	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.language
}

// Tree returns the document's syntax tree
func (d *BaseDocument) Tree() *ts.Tree {
	if d == nil {
		return nil
	}

	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.tree
}

// SetTree sets the document's syntax tree (protected method for subclasses)
func (d *BaseDocument) SetTree(tree *ts.Tree) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.tree != nil {
		d.tree.Close()
	}
	d.tree = tree
}

// SetParser sets the document's parser (protected method for subclasses)
func (d *BaseDocument) SetParser(parser *ts.Parser) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.parser = parser
}

// GetParser returns the document's parser (protected method for subclasses)
func (d *BaseDocument) GetParser() *ts.Parser {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.parser
}

// UpdateContent updates the document content (protected method for subclasses)
func (d *BaseDocument) UpdateContent(content string, version int32) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.content = content
	d.version = version
}

// ScriptTags returns the parsed script tags for HTML documents
func (d *BaseDocument) ScriptTags() []types.ScriptTag {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.scriptTags
}

// ImportMap returns nil by default (HTML documents override this to provide import maps)
func (d *BaseDocument) ImportMap() map[string]string {
	return nil
}

// SetScriptTags sets the script tags (protected method for subclasses)
func (d *BaseDocument) SetScriptTags(scriptTags []types.ScriptTag) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.scriptTags = scriptTags
}

// Close cleans up document resources
func (d *BaseDocument) Close() {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.tree != nil {
		d.tree.Close()
		d.tree = nil
	}

	if d.parser != nil {
		// Return parser to pool based on language - this will be handled by language handlers
		d.parser = nil
	}
}

// Default implementations that can be overridden by language-specific documents

// Parse performs initial parsing of the document (default implementation)
func (d *BaseDocument) Parse(content string) error {
	d.UpdateContent(content, d.version)
	// Language-specific parsing will be implemented by subclasses
	return nil
}

// CompletionPrefix extracts the prefix being typed for filtering completions (default implementation)
func (d *BaseDocument) CompletionPrefix(analysis *types.CompletionAnalysis) string {
	// Default implementation - can be overridden by language-specific documents
	return ""
}

// FindModuleScript finds the first module script tag and returns insertion position (default implementation)
func (d *BaseDocument) FindModuleScript() (protocol.Position, bool) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	for _, script := range d.scriptTags {
		if script.IsModule {
			// Return position at the end of the script content for insertion
			return protocol.Position{
				Line:      script.ContentRange.End.Line,
				Character: 0, // Start of line for clean insertion
			}, true
		}
	}

	return protocol.Position{}, false
}

// FindInlineModuleScript finds the first inline module script (no src) and returns insertion position (default implementation)
func (d *BaseDocument) FindInlineModuleScript() (protocol.Position, bool) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	for _, script := range d.scriptTags {
		if script.IsModule && script.Src == "" {
			// Return position at the end of the script content for insertion
			return protocol.Position{
				Line:      script.ContentRange.End.Line,
				Character: 0, // Start of line for clean insertion
			}, true
		}
	}

	return protocol.Position{}, false
}

// ByteRangeToProtocolRange converts byte range to protocol range (default implementation)
func (d *BaseDocument) ByteRangeToProtocolRange(content string, startByte, endByte uint) protocol.Range {
	// Basic implementation - can be enhanced by language-specific documents
	return protocol.Range{
		Start: d.byteOffsetToPosition(startByte),
		End:   d.byteOffsetToPosition(endByte),
	}
}

// Helper method for byte offset to position conversion
func (d *BaseDocument) byteOffsetToPosition(offset uint) protocol.Position {
	line := uint32(0)
	char := uint32(0)

	for i, r := range d.content {
		if uint(i) >= offset {
			break
		}

		if r == '\n' {
			line++
			char = 0
		} else {
			char++
		}
	}

	return protocol.Position{
		Line:      line,
		Character: char,
	}
}
