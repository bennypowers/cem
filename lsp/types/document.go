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
package types

import (
	M "bennypowers.dev/cem/manifest"
	Q "bennypowers.dev/cem/queries"
	protocol "github.com/tliron/glsp/protocol_3_16"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// Document interface for LSP document operations
type Document interface {
	FindElementAtPosition(position protocol.Position, dm any) *CustomElementMatch
	FindAttributeAtPosition(position protocol.Position, dm any) (*AttributeMatch, string)
	Content() (string, error) // Returns content with proper error handling
	Version() int32
	URI() string
	Language() string // Returns the document language
	FindCustomElements(dm any) ([]CustomElementMatch, error)
	AnalyzeCompletionContextTS(position protocol.Position, dm any) *CompletionAnalysis
	CompletionPrefix(analysis *CompletionAnalysis) string                            // Get completion prefix using tree-sitter
	ScriptTags() []ScriptTag                                                         // Get parsed script tags
	FindModuleScript() (protocol.Position, bool)                                     // Find insertion point in module script
	FindInlineModuleScript() (protocol.Position, bool)                               // Find insertion point in inline module script (no src)
	FindHeadInsertionPoint(dm any) (protocol.Position, bool)                         // Find insertion point in <head> section
	ByteRangeToProtocolRange(content string, startByte, endByte uint) protocol.Range // Convert byte range to protocol range
	Close()                                                                          // Clean up document resources

	// Tree-sitter methods for incremental parsing support
	Tree() *ts.Tree                              // Get the current syntax tree
	SetTree(tree *ts.Tree)                       // Set the syntax tree
	Parser() *ts.Parser                          // Get the tree-sitter parser
	SetParser(parser *ts.Parser)                 // Set the tree-sitter parser
	UpdateContent(content string, version int32) // Update document content
}

// LanguageHandler defines the interface for language-specific operations
type LanguageHandler interface {
	Language() string
	CreateDocument(uri, content string, version int32) Document
	FindCustomElements(doc Document) ([]CustomElementMatch, error)
	AnalyzeCompletionContext(doc Document, position protocol.Position) *CompletionAnalysis
	FindElementAtPosition(doc Document, position protocol.Position) *CustomElementMatch
	FindAttributeAtPosition(doc Document, position protocol.Position) (*AttributeMatch, string)
	Close()
}

// Manager interface for document lifecycle management
type Manager interface {
	// Document lifecycle
	OpenDocument(uri, content string, version int32) Document
	UpdateDocument(uri, content string, version int32) Document
	UpdateDocumentWithChanges(uri, content string, version int32, changes []protocol.TextDocumentContentChangeEvent) Document
	CloseDocument(uri string)
	Document(uri string) Document
	AllDocuments() []Document

	// Core query management (language-agnostic)
	QueryManager() *Q.QueryManager

	// Cleanup
	Close()
}

// CustomElementMatch represents a found custom element
type CustomElementMatch struct {
	TagName    string
	Range      protocol.Range
	Attributes map[string]AttributeMatch
}

// AttributeMatch represents a found attribute
type AttributeMatch struct {
	Name  string
	Value string
	Range protocol.Range
}

// ElementDefinition represents a custom element with its source information
type ElementDefinition interface {
	ModulePath() string
	PackageName() string
	SourceHref() string
	Element() *M.CustomElement
}

// ScriptTag represents a tracked script tag in HTML documents
type ScriptTag struct {
	Range        protocol.Range    // Range of the entire script tag
	ContentRange protocol.Range    // Range of the script content (inside the tag)
	Type         string            // "module", "classic", or empty
	Src          string            // src attribute value (if any)
	IsModule     bool              // true if type="module"
	Imports      []ImportStatement // Parsed import statements
}

// ImportStatement represents an import statement within a script tag
type ImportStatement struct {
	Range      protocol.Range // Range of the import statement
	ImportPath string         // The imported path/module
	Type       string         // "static" or "dynamic"
}

// ParseStrategy represents different parsing strategies for document updates
type ParseStrategy int

const (
	// ParseStrategyFull always performs a full reparse
	ParseStrategyFull ParseStrategy = iota
	// ParseStrategyIncremental attempts incremental parsing with fallback to full
	ParseStrategyIncremental
	// ParseStrategyAuto automatically chooses based on change characteristics
	ParseStrategyAuto
)

// ParseResult represents the result of a parsing operation
type ParseResult struct {
	Success         bool
	UsedIncremental bool
	Error           error
	NewTree         *ts.Tree
	OldTree         *ts.Tree // For cleanup purposes
}

// IncrementalParser interface for incremental document parsing
type IncrementalParser interface {
	// ParseWithStrategy parses a document using the configured strategy
	ParseWithStrategy(doc Document, newContent string, changes []protocol.TextDocumentContentChangeEvent) ParseResult
}
