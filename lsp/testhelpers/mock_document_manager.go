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
package testhelpers

import (
	"bennypowers.dev/cem/lsp/types"
	"bennypowers.dev/cem/queries"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// MockDocumentManager provides a unified mock implementation of DocumentManager for all tests
type MockDocumentManager struct {
	Documents map[string]types.Document
	QueryMgr  *queries.QueryManager
}

// Verify MockDocumentManager implements DocumentManager
var _ types.DocumentManager = (*MockDocumentManager)(nil)

// NewMockDocumentManager creates a new mock document manager
func NewMockDocumentManager() *MockDocumentManager {
	return &MockDocumentManager{
		Documents: make(map[string]types.Document),
	}
}

// OpenDocument creates a new document
func (dm *MockDocumentManager) OpenDocument(uri, content string, version int32) types.Document {
	doc := NewMockDocument(content)
	doc.URIStr = uri
	doc.VersionNum = version
	dm.Documents[uri] = doc
	return doc
}

// UpdateDocument updates an existing document
func (dm *MockDocumentManager) UpdateDocument(uri, content string, version int32) types.Document {
	return dm.UpdateDocumentWithChanges(uri, content, version, nil)
}

// UpdateDocumentWithChanges updates an existing document with incremental changes
func (dm *MockDocumentManager) UpdateDocumentWithChanges(uri, content string, version int32, changes []protocol.TextDocumentContentChangeEvent) types.Document {
	if doc, exists := dm.Documents[uri]; exists {
		// Update existing document
		if mockDoc, ok := doc.(*MockDocument); ok {
			mockDoc.ContentStr = content
			mockDoc.VersionNum = version
			return mockDoc
		}
	}
	// Create new document if it doesn't exist
	return dm.OpenDocument(uri, content, version)
}

// CloseDocument removes a document from tracking
func (dm *MockDocumentManager) CloseDocument(uri string) {
	delete(dm.Documents, uri)
}

// Document retrieves a tracked document
func (dm *MockDocumentManager) Document(uri string) types.Document {
	return dm.Documents[uri]
}

// AllDocuments returns all tracked documents
func (dm *MockDocumentManager) AllDocuments() []types.Document {
	docs := make([]types.Document, 0, len(dm.Documents))
	for _, doc := range dm.Documents {
		docs = append(docs, doc)
	}
	return docs
}

// Close cleans up the document manager
func (dm *MockDocumentManager) Close() {
	// Clean up all documents
	dm.Documents = make(map[string]types.Document)
}

// QueryManager returns the query manager for tree-sitter queries
func (dm *MockDocumentManager) QueryManager() *queries.QueryManager {
	return dm.QueryMgr
}

// SetQueryManager sets the query manager for tests
func (dm *MockDocumentManager) SetQueryManager(qm *queries.QueryManager) {
	dm.QueryMgr = qm
}

// AddDocument adds a document to the mock document manager
func (dm *MockDocumentManager) AddDocument(uri string, doc types.Document) {
	dm.Documents[uri] = doc
}