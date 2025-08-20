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
	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/types"
)

// CreateTestDocument creates a real document using the provided DocumentManager
// Tests should call dm.OpenDocument() directly, but this helper is for compatibility
func CreateTestDocument(dm types.DocumentManager, content string) types.Document {
	return dm.OpenDocument("test://document.html", content, 1)
}

// CreateTestDocumentWithURI creates a real document with a specific URI
func CreateTestDocumentWithURI(dm types.DocumentManager, uri, content string) types.Document {
	return dm.OpenDocument(uri, content, 1)
}

// Legacy compatibility functions - these will be removed in favor of direct dm.OpenDocument() calls

// NewMockDocument creates a real document using a new DocumentManager
// DEPRECATED: Tests should create their own DocumentManager and call dm.OpenDocument()
func NewMockDocument(content string) types.Document {
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		panic("Failed to create DocumentManager for test: " + err.Error())
	}
	defer dm.Close()
	
	// Call OpenDocument on the DocumentManager as requested
	return dm.OpenDocument("test://mock.html", content, 1)
}

// NewMockDocumentWithElements creates a real document using a new DocumentManager
// DEPRECATED: Tests should create their own DocumentManager and call dm.OpenDocument()
func NewMockDocumentWithElements(content string, elements []types.CustomElementMatch) types.Document {
	// Ignore the elements parameter - real Documents use tree-sitter to find elements
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		panic("Failed to create DocumentManager for test: " + err.Error())
	}
	defer dm.Close()
	
	// Call OpenDocument on the DocumentManager as requested
	return dm.OpenDocument("test://mock.html", content, 1)
}



// MockDocument represents a document type for testing
// DEPRECATED: Tests should use real Documents via dm.OpenDocument()
type MockDocument = types.Document