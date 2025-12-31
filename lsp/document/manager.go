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
	"strings"
	"sync"

	"bennypowers.dev/cem/lsp/document/html"
	"bennypowers.dev/cem/lsp/document/tsx"
	"bennypowers.dev/cem/lsp/document/typescript"
	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	Q "bennypowers.dev/cem/queries"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// documentManager implements the types.Manager interface
type documentManager struct {
	documents         map[string]types.Document
	queryManager      *Q.QueryManager
	languageHandlers  map[string]types.LanguageHandler
	incrementalParser types.IncrementalParser
	mu                sync.RWMutex
	// Per-URI locks to serialize tree-sitter operations on the same document
	// Tree-sitter C objects are NOT thread-safe - concurrent access causes segfaults
	uriLocks   map[string]*sync.Mutex
	uriLocksMu sync.Mutex
}

// NewDocumentManager creates a new document manager with language handlers
func NewDocumentManager() (types.Manager, error) {
	// Use LSP-specific query selector for performance
	queryManager, err := Q.NewQueryManager(Q.LSPQueries())
	if err != nil {
		return nil, err
	}

	dm := &documentManager{
		documents:         make(map[string]types.Document),
		queryManager:      queryManager,
		languageHandlers:  make(map[string]types.LanguageHandler),
		incrementalParser: newIncrementalParser(types.ParseStrategyAuto),
		uriLocks:          make(map[string]*sync.Mutex),
	}

	// Initialize language handlers
	if err := dm.initializeLanguageHandlers(); err != nil {
		return nil, fmt.Errorf("failed to initialize language handlers: %w", err)
	}

	return dm, nil
}

// getURILock returns or creates a mutex for the given URI
// This ensures tree-sitter operations on the same document are serialized
func (dm *documentManager) getURILock(uri string) *sync.Mutex {
	dm.uriLocksMu.Lock()
	defer dm.uriLocksMu.Unlock()

	if lock, exists := dm.uriLocks[uri]; exists {
		return lock
	}

	lock := &sync.Mutex{}
	dm.uriLocks[uri] = lock
	return lock
}

// cleanupURILock removes the mutex for a URI when the document is closed
func (dm *documentManager) cleanupURILock(uri string) {
	dm.uriLocksMu.Lock()
	defer dm.uriLocksMu.Unlock()
	delete(dm.uriLocks, uri)
}

// OpenDocument tracks a new document
func (dm *documentManager) OpenDocument(uri, content string, version int32) types.Document {
	// Serialize all tree-sitter operations for this URI to prevent concurrent access
	uriLock := dm.getURILock(uri)
	uriLock.Lock()
	defer uriLock.Unlock()

	dm.mu.Lock()
	defer dm.mu.Unlock()

	// Close existing document if it exists
	if existing, exists := dm.documents[uri]; exists {
		existing.Close()
	}

	language := getLanguageFromURI(uri)

	// Get the appropriate language handler
	handler, exists := dm.languageHandlers[language]
	if !exists {
		helpers.SafeDebugLog("[DOCUMENT] No handler for language %s, using HTML as fallback", language)
		// Fallback to HTML handler for unknown file types
		handler = dm.languageHandlers["html"]
		if handler == nil {
			helpers.SafeDebugLog("[DOCUMENT] No HTML handler available, creating basic document")
			return nil
		}
	}

	doc := handler.CreateDocument(uri, content, version)
	dm.documents[uri] = doc
	return doc
}

// UpdateDocument updates an existing document with incremental parsing
func (dm *documentManager) UpdateDocument(uri, content string, version int32) types.Document {
	// This method maintains compatibility with simple content updates
	// For full incremental parsing, use UpdateDocumentWithChanges
	return dm.UpdateDocumentWithChanges(uri, content, version, nil)
}

// UpdateDocumentWithChanges updates an existing document using incremental parsing when possible
func (dm *documentManager) UpdateDocumentWithChanges(uri, content string, version int32, changes []protocol.TextDocumentContentChangeEvent) types.Document {
	helpers.SafeDebugLog("[DOCUMENT] UpdateDocumentWithChanges: URI=%s, Version=%d, ContentLength=%d, Changes=%d",
		uri, version, len(content), len(changes))

	// Serialize all tree-sitter operations for this URI to prevent concurrent access
	uriLock := dm.getURILock(uri)
	uriLock.Lock()
	defer uriLock.Unlock()

	dm.mu.Lock()
	defer dm.mu.Unlock()

	doc, exists := dm.documents[uri]
	if !exists {
		helpers.SafeDebugLog("[DOCUMENT] Document not found, creating new one: %s", uri)
		// Note: We already hold the URI lock, OpenDocument will try to acquire it again
		// Need to unlock before calling OpenDocument to avoid deadlock
		dm.mu.Unlock()
		uriLock.Unlock()
		return dm.OpenDocument(uri, content, version)
	}

	// Get the language handler for this document
	language := getLanguageFromURI(uri)
	handler, handlerExists := dm.languageHandlers[language]
	if !handlerExists {
		handler = dm.languageHandlers["html"] // Fallback to HTML
	}

	// Try incremental parsing if we have changes and an existing tree
	if len(changes) > 0 && doc.Tree() != nil {
		helpers.SafeDebugLog("[DOCUMENT] Attempting incremental parsing for %s", uri)

		// Use incremental parser to analyze and potentially parse incrementally
		result := dm.incrementalParser.ParseWithStrategy(doc, content, changes)

		if result.Success {
			helpers.SafeDebugLog("[DOCUMENT] Incremental parsing %s: UsedIncremental=%t",
				uri, result.UsedIncremental)

			// Update document content and version
			doc.UpdateContent(content, version)

			// Update the tree with the new tree from parsing
			if result.NewTree != nil {
				doc.SetTree(result.NewTree)
			}

			// Note: OldTree cleanup is handled by SetTree method

			// IMPORTANT: Re-parse script tags and importmap ONLY if script content changed
			// This ensures diagnostics reflect the current document state while being performant
			if dm.scriptContentChanged(doc, changes) {
				helpers.SafeDebugLog("[DOCUMENT] Script content changed, re-parsing metadata for %s", uri)
				dm.reparseDocumentMetadata(doc, handler)
			} else {
				helpers.SafeDebugLog("[DOCUMENT] Script content unchanged, skipping metadata re-parse for %s", uri)
			}

			return doc
		} else {
			helpers.SafeDebugLog("[DOCUMENT] Incremental parsing failed for %s: %v", uri, result.Error)
		}
	}

	// Fall back to full document recreation
	helpers.SafeDebugLog("[DOCUMENT] Using full document recreation for %s", uri)

	if handler != nil {
		// Close the old document
		doc.Close()
		// Create a new document with updated content
		newDoc := handler.CreateDocument(uri, content, version)
		dm.documents[uri] = newDoc
		return newDoc
	}

	return doc
}

// scriptContentChanged checks if any of the document changes affected script tags or importmaps
func (dm *documentManager) scriptContentChanged(doc types.Document, changes []protocol.TextDocumentContentChangeEvent) bool {
	if doc == nil || len(changes) == 0 {
		return false
	}

	// Get existing script tags
	scriptTags := doc.ScriptTags()
	if len(scriptTags) == 0 {
		// No existing script tags - check if changes might add one
		// Look for "<script" in the change text
		for _, change := range changes {
			if strings.Contains(change.Text, "<script") {
				return true
			}
		}
		return false
	}

	// Check if any change overlaps with existing script tag ranges
	for _, change := range changes {
		if change.Range == nil {
			// Full document change - script content definitely changed
			return true
		}

		changeStart := change.Range.Start
		changeEnd := change.Range.End

		for _, scriptTag := range scriptTags {
			// Check if change overlaps with script tag range
			if rangesOverlap(changeStart, changeEnd, scriptTag.Range.Start, scriptTag.Range.End) {
				return true
			}
		}

		// Also check if the change text contains script-related content
		if strings.Contains(change.Text, "<script") || strings.Contains(change.Text, "</script>") {
			return true
		}
	}

	return false
}

// rangesOverlap checks if two ranges overlap
func rangesOverlap(start1, end1, start2, end2 protocol.Position) bool {
	// Range 1 starts before range 2 ends, AND range 1 ends after range 2 starts
	return !positionAfter(start1, end2) && !positionBefore(end1, start2)
}

// positionBefore checks if pos1 is strictly before pos2
func positionBefore(pos1, pos2 protocol.Position) bool {
	if pos1.Line < pos2.Line {
		return true
	}
	if pos1.Line == pos2.Line && pos1.Character < pos2.Character {
		return true
	}
	return false
}

// positionAfter checks if pos1 is strictly after pos2
func positionAfter(pos1, pos2 protocol.Position) bool {
	if pos1.Line > pos2.Line {
		return true
	}
	if pos1.Line == pos2.Line && pos1.Character > pos2.Character {
		return true
	}
	return false
}

// reparseDocumentMetadata re-parses script tags and importmap for HTML documents
// This is called after incremental updates to ensure metadata is up-to-date
func (dm *documentManager) reparseDocumentMetadata(doc types.Document, handler types.LanguageHandler) {
	if doc == nil || handler == nil {
		return
	}

	// Only HTML documents have script tags and importmaps
	if doc.Language() != "html" {
		return
	}

	// Use type assertion to access HTML-specific handler methods
	type scriptTagParser interface {
		ParseScriptTags(doc types.Document) ([]types.ScriptTag, error)
		ParseImportMap(doc types.Document) (map[string]string, error)
	}

	if htmlHandler, ok := handler.(scriptTagParser); ok {
		// Re-parse script tags
		if scriptTags, err := htmlHandler.ParseScriptTags(doc); err != nil {
			helpers.SafeDebugLog("[DOCUMENT] Failed to re-parse script tags: %v", err)
		} else {
			// Use type assertion to set script tags
			type scriptTagSetter interface {
				SetScriptTags([]types.ScriptTag)
			}
			if setter, ok := doc.(scriptTagSetter); ok {
				setter.SetScriptTags(scriptTags)
				helpers.SafeDebugLog("[DOCUMENT] Re-parsed %d script tags", len(scriptTags))
			}
		}

		// Re-parse import map
		if importMap, err := htmlHandler.ParseImportMap(doc); err != nil {
			helpers.SafeDebugLog("[DOCUMENT] Failed to re-parse import map: %v", err)
		} else if importMap != nil {
			// Use type assertion to set import map
			type importMapSetter interface {
				SetImportMap(map[string]string)
			}
			if setter, ok := doc.(importMapSetter); ok {
				setter.SetImportMap(importMap)
				helpers.SafeDebugLog("[DOCUMENT] Re-parsed import map with %d entries", len(importMap))
			}
		}
	}
}

// Document retrieves a tracked document
func (dm *documentManager) Document(uri string) types.Document {
	dm.mu.RLock()
	defer dm.mu.RUnlock()

	return dm.documents[uri]
}

// AllDocuments returns all tracked documents
func (dm *documentManager) AllDocuments() []types.Document {
	dm.mu.RLock()
	defer dm.mu.RUnlock()

	documents := make([]types.Document, 0, len(dm.documents))
	for _, doc := range dm.documents {
		documents = append(documents, doc)
	}
	return documents
}

// CloseDocument removes a document from tracking
func (dm *documentManager) CloseDocument(uri string) {
	// Serialize all tree-sitter operations for this URI to prevent concurrent access
	uriLock := dm.getURILock(uri)
	uriLock.Lock()
	defer uriLock.Unlock()

	dm.mu.Lock()
	defer dm.mu.Unlock()

	if doc, exists := dm.documents[uri]; exists {
		doc.Close()
		delete(dm.documents, uri)
	}

	// Clean up the URI lock after closing the document
	dm.cleanupURILock(uri)
}

// QueryManager returns the query manager for tree-sitter queries
func (dm *documentManager) QueryManager() *Q.QueryManager {
	return dm.queryManager
}

// Close cleans up the document manager
func (dm *documentManager) Close() {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	for _, doc := range dm.documents {
		doc.Close()
	}

	for _, handler := range dm.languageHandlers {
		handler.Close()
	}

	if dm.queryManager != nil {
		dm.queryManager.Close()
	}
}

// initializeLanguageHandlers creates and registers all language handlers
func (dm *documentManager) initializeLanguageHandlers() error {
	// Import the language handlers
	htmlHandler, err := createHTMLHandler(dm.queryManager)
	if err != nil {
		return fmt.Errorf("failed to create HTML handler: %w", err)
	}
	dm.addLanguageHandler(htmlHandler)

	tsHandler, err := createTypeScriptHandler(dm.queryManager)
	if err != nil {
		return fmt.Errorf("failed to create TypeScript handler: %w", err)
	}
	dm.addLanguageHandler(tsHandler)

	tsxHandler, err := createTSXHandler(dm.queryManager)
	if err != nil {
		return fmt.Errorf("failed to create TSX handler: %w", err)
	}
	dm.addLanguageHandler(tsxHandler)

	return nil
}

// addLanguageHandler adds a language handler (internal method)
func (dm *documentManager) addLanguageHandler(handler types.LanguageHandler) {
	dm.languageHandlers[handler.Language()] = handler
}

// GetLanguageHandler gets a language handler by name (for internal use by documents)
func (dm *documentManager) GetLanguageHandler(language string) types.LanguageHandler {
	dm.mu.RLock()
	defer dm.mu.RUnlock()
	return dm.languageHandlers[language]
}

// Factory functions for language handlers

// createHTMLHandler creates a new HTML language handler
func createHTMLHandler(queryManager *Q.QueryManager) (types.LanguageHandler, error) {
	return html.NewHandler(queryManager)
}

// createTypeScriptHandler creates a new TypeScript language handler
func createTypeScriptHandler(queryManager *Q.QueryManager) (types.LanguageHandler, error) {
	return typescript.NewHandler(queryManager)
}

// createTSXHandler creates a new TSX language handler
func createTSXHandler(queryManager *Q.QueryManager) (types.LanguageHandler, error) {
	return tsx.NewHandler(queryManager)
}
