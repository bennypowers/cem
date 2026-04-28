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
package base

import (
	"strings"
	"sync"

	"bennypowers.dev/cem/internal/textutil"
	"bennypowers.dev/cem/lsp/types"
	protocol "github.com/tliron/glsp/protocol_3_16"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// BaseDocument provides common document functionality that language-specific documents can embed.
// Always use NewBaseDocument to construct; a zero-value BaseDocument will panic.
type BaseDocument struct {
	uri          string
	content      string
	version      int32
	language     string
	tree         *ts.Tree
	parser       *ts.Parser
	scriptTags   []types.ScriptTag
	returnParser func(*ts.Parser)
	self         types.Document
	mu           sync.RWMutex
}

// NewBaseDocument creates a new base document with a callback for returning parsers to the correct pool.
// The self parameter should be set to the embedding struct (the concrete document type) so that
// reflection-based handler delegation passes the correct type for handler type assertions.
func NewBaseDocument(uri, content string, version int32, language string, returnParser func(*ts.Parser), self types.Document) *BaseDocument {
	return &BaseDocument{
		uri:          uri,
		content:      content,
		version:      version,
		language:     language,
		returnParser: returnParser,
		self:         self,
	}
}

// URI returns the document URI
func (d *BaseDocument) URI() string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.uri
}

// Content returns the document content
func (d *BaseDocument) Content() (string, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.content, nil
}

// Version returns the document version
func (d *BaseDocument) Version() int32 {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.version
}

// Language returns the document language
func (d *BaseDocument) Language() string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.language
}

// Tree returns the document's syntax tree
func (d *BaseDocument) Tree() *ts.Tree {
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

// SetParser sets the document's parser, returning any previous parser to the pool.
func (d *BaseDocument) SetParser(parser *ts.Parser) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.parser != nil && d.parser != parser {
		if d.returnParser != nil {
			d.returnParser(d.parser)
		}
	}
	d.parser = parser
}

// Parser returns the document's parser (protected method for subclasses)
func (d *BaseDocument) Parser() *ts.Parser {
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

// Close cleans up document resources, returning the parser to its pool.
func (d *BaseDocument) Close() {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.tree != nil {
		d.tree.Close()
		d.tree = nil
	}

	if d.parser != nil {
		if d.returnParser != nil {
			d.returnParser(d.parser)
		}
		d.parser = nil
	}
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

// FindInlineModuleScript finds the first inline module script (no src attribute).
func (d *BaseDocument) FindInlineModuleScript() (protocol.Position, bool) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	for _, script := range d.scriptTags {
		if script.IsModule && script.Src == "" {
			return protocol.Position{
				Line:      script.ContentRange.End.Line,
				Character: 0,
			}, true
		}
	}

	return protocol.Position{}, false
}

// TreeAndContent returns the tree and content atomically under one lock.
// Use this from sub-package methods that need both values consistently.
//
// The returned tree is safe to use after the lock releases because
// DocumentManager serializes all operations on the same URI via per-URI
// locks, preventing concurrent Close() from freeing the tree mid-use.
func (d *BaseDocument) TreeAndContent() (*ts.Tree, string) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.tree, d.content
}

// ByteRangeToProtocolRange converts byte range to protocol range.
func (d *BaseDocument) ByteRangeToProtocolRange(content string, startByte, endByte uint) protocol.Range {
	return protocol.Range{
		Start: d.ByteOffsetToPosition(startByte, content),
		End:   d.ByteOffsetToPosition(endByte, content),
	}
}

// ByteOffsetToPosition converts a byte offset to a protocol position within the given content.
// Character is in UTF-16 code units per the LSP specification.
func (d *BaseDocument) ByteOffsetToPosition(offset uint, content string) protocol.Position {
	line := uint32(0)
	lineStart := uint(0)

	for i := uint(0); i < offset && i < uint(len(content)); i++ {
		if content[i] == '\n' {
			line++
			lineStart = i + 1
		}
	}

	return protocol.Position{
		Line:      line,
		Character: textutil.ByteOffsetToUTF16(content[lineStart:], offset-lineStart),
	}
}

// PositionToByteOffset converts a protocol position to a byte offset within the given content.
// Character is interpreted as UTF-16 code units per the LSP specification.
func (d *BaseDocument) PositionToByteOffset(pos protocol.Position, content string) uint {
	var offset uint
	lines := strings.Split(content, "\n")

	for i := uint32(0); i < pos.Line && i < uint32(len(lines)); i++ {
		offset += uint(len(lines[i])) + 1
	}

	if pos.Line < uint32(len(lines)) {
		line := lines[pos.Line]
		offset += textutil.UTF16ToByteOffset(line, pos.Character)
	}

	return offset
}

// outer returns the concrete document type for handler delegation.
// When embedded, self points to the outer struct; standalone, it falls back to d.
func (d *BaseDocument) outer() types.Document {
	if d.self != nil {
		return d.self
	}
	return d
}

// FindElementAtPosition finds a custom element at the given position.
func (d *BaseDocument) FindElementAtPosition(position protocol.Position, hp types.HandlerProvider) *types.CustomElementMatch {
	if hp == nil {
		return nil
	}
	handler := hp.GetLanguageHandler(d.language)
	if handler == nil {
		return nil
	}
	return handler.FindElementAtPosition(d.outer(), position)
}

// FindAttributeAtPosition finds an attribute at the given position.
func (d *BaseDocument) FindAttributeAtPosition(position protocol.Position, hp types.HandlerProvider) (*types.AttributeMatch, string) {
	if hp == nil {
		return nil, ""
	}
	handler := hp.GetLanguageHandler(d.language)
	if handler == nil {
		return nil, ""
	}
	return handler.FindAttributeAtPosition(d.outer(), position)
}

// FindCustomElements finds custom elements in the document.
func (d *BaseDocument) FindCustomElements(hp types.HandlerProvider) ([]types.CustomElementMatch, error) {
	if hp == nil {
		return []types.CustomElementMatch{}, nil
	}
	handler := hp.GetLanguageHandler(d.language)
	if handler == nil {
		return []types.CustomElementMatch{}, nil
	}
	return handler.FindCustomElements(d.outer())
}

// AnalyzeCompletionContextTS analyzes completion context using tree-sitter queries.
func (d *BaseDocument) AnalyzeCompletionContextTS(position protocol.Position, hp types.HandlerProvider) *types.CompletionAnalysis {
	if hp == nil {
		return &types.CompletionAnalysis{Type: types.CompletionUnknown}
	}
	handler := hp.GetLanguageHandler(d.language)
	if handler == nil {
		return &types.CompletionAnalysis{Type: types.CompletionUnknown}
	}
	return handler.AnalyzeCompletionContext(d.outer(), position)
}

// FindHeadInsertionPoint finds insertion point in <head> section.
func (d *BaseDocument) FindHeadInsertionPoint(hp types.HandlerProvider) (protocol.Position, bool) {
	if hp == nil {
		return protocol.Position{}, false
	}
	handler := hp.GetLanguageHandler(d.language)
	if handler == nil {
		return protocol.Position{}, false
	}
	if h, ok := handler.(interface {
		FindHeadInsertionPoint(types.Document) (protocol.Position, bool)
	}); ok {
		return h.FindHeadInsertionPoint(d.outer())
	}
	return protocol.Position{}, false
}
