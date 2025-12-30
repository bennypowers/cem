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
package lsp_test

import (
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/lsp/document"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestIncrementalParsingIntegration(t *testing.T) {
	fixtureDir := filepath.Join("testdata", "integration", "incremental-parsing-test")

	// Setup document manager
	dm, err := document.NewDocumentManager()
	require.NoError(t, err, "Should create document manager")
	defer dm.Close()

	tests := []struct {
		name           string
		initialFile    string
		updatedFile    string
		uri            string
		changePosition struct {
			line      uint32
			startChar uint32
			endChar   uint32
		}
		expectedLanguage string
	}{
		{
			name:        "HTML Document Incremental Parsing",
			initialFile: "html-initial.html",
			updatedFile: "html-updated.html",
			uri:         "file:///test.html",
			changePosition: struct {
				line      uint32
				startChar uint32
				endChar   uint32
			}{
				line:      4,
				startChar: 20, // Start of "primary"
				endChar:   27, // End of "primary"
			},
			expectedLanguage: "html",
		},
		{
			name:        "TypeScript Document Incremental Parsing",
			initialFile: "typescript-initial.ts",
			updatedFile: "typescript-updated.ts",
			uri:         "file:///test.ts",
			changePosition: struct {
				line      uint32
				startChar uint32
				endChar   uint32
			}{
				line:      4,
				startChar: 24, // Start of "primary" in template literal
				endChar:   31, // End of "primary"
			},
			expectedLanguage: "typescript",
		},
		{
			name:        "TSX Document Incremental Parsing",
			initialFile: "tsx-initial.tsx",
			updatedFile: "tsx-updated.tsx",
			uri:         "file:///test.tsx",
			changePosition: struct {
				line      uint32
				startChar uint32
				endChar   uint32
			}{
				line:      4,
				startChar: 20, // Start of "primary" in JSX
				endChar:   27, // End of "primary"
			},
			expectedLanguage: "tsx",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read initial content from fixture
			initialPath := filepath.Join(fixtureDir, tt.initialFile)
			initialContent, err := os.ReadFile(initialPath)
			require.NoError(t, err, "Should read initial fixture file")

			// Read updated content from fixture
			updatedPath := filepath.Join(fixtureDir, tt.updatedFile)
			updatedContentBytes, err := os.ReadFile(updatedPath)
			require.NoError(t, err, "Should read updated fixture file")
			updatedContentStr := string(updatedContentBytes)

			// Create document with initial content
			doc := dm.OpenDocument(tt.uri, string(initialContent), 1)
			require.NotNil(t, doc, "Should create document")
			assert.Equal(t, tt.expectedLanguage, doc.Language(), "Should detect correct language")

			// Define incremental change (primary -> secondary)
			changes := []protocol.TextDocumentContentChangeEvent{
				{
					Range: &protocol.Range{
						Start: protocol.Position{
							Line:      tt.changePosition.line,
							Character: tt.changePosition.startChar,
						},
						End: protocol.Position{
							Line:      tt.changePosition.line,
							Character: tt.changePosition.endChar,
						},
					},
					Text: "secondary",
				},
			}

			// Test document manager incremental update
			updatedDoc := dm.UpdateDocumentWithChanges(tt.uri, updatedContentStr, 2, changes)
			require.NotNil(t, updatedDoc, "Should get updated document")

			// Verify content was updated
			newContent, err := updatedDoc.Content()
			require.NoError(t, err, "Should get updated content")
			assert.Equal(t, updatedContentStr, newContent, "Content should be updated")

			// Verify version was updated
			assert.Equal(t, int32(2), updatedDoc.Version(), "Version should be updated")

			// Verify tree is valid
			if updatedDoc.Tree() != nil {
				root := updatedDoc.Tree().RootNode()
				assert.NotNil(t, root, "Should have valid tree root")
				// Tree should span the entire content
				contentLength := uint(len(newContent))
				assert.GreaterOrEqual(t, root.EndByte(), contentLength-10, "Tree should cover most of content")
			}
		})
	}
}

func TestIncrementalParsingFallback(t *testing.T) {
	fixtureDir := filepath.Join("testdata", "integration", "incremental-parsing-test")

	dm, err := document.NewDocumentManager()
	require.NoError(t, err, "Should create document manager")
	defer dm.Close()

	t.Run("Large Change Fallback to Full Parse", func(t *testing.T) {
		// Read fixture files
		initialPath := filepath.Join(fixtureDir, "large-change-initial.html")
		initialContent, err := os.ReadFile(initialPath)
		require.NoError(t, err, "Should read initial fixture file")

		updatedPath := filepath.Join(fixtureDir, "large-change-updated.html")
		updatedContent, err := os.ReadFile(updatedPath)
		require.NoError(t, err, "Should read updated fixture file")

		doc := dm.OpenDocument("file:///test-large.html", string(initialContent), 1)
		require.NotNil(t, doc, "Should create document")

		// Large change that should trigger full parse
		changes := []protocol.TextDocumentContentChangeEvent{
			{
				Range: &protocol.Range{
					Start: protocol.Position{Line: 0, Character: 0},
					End:   protocol.Position{Line: 0, Character: uint32(len(initialContent))},
				},
				Text: string(updatedContent),
			},
		}

		// Test document manager handles large changes
		updatedDoc := dm.UpdateDocumentWithChanges("file:///test-large.html", string(updatedContent), 2, changes)
		require.NotNil(t, updatedDoc, "Should get updated document")

		// Verify content was updated
		newContent, err := updatedDoc.Content()
		require.NoError(t, err, "Should get updated content")
		assert.Equal(t, string(updatedContent), newContent, "Content should be updated")

		// Verify tree is valid
		if updatedDoc.Tree() != nil {
			root := updatedDoc.Tree().RootNode()
			assert.NotNil(t, root, "Should have valid tree root")
		}
	})

	t.Run("Full Document Change Fallback", func(t *testing.T) {
		initialContent := `<my-card variant="primary">Original</my-card>`

		doc := dm.OpenDocument("file:///test-full.html", initialContent, 1)
		require.NotNil(t, doc, "Should create document")

		// Full document replacement (Range is nil)
		newContent := `<ui-button size="large">Completely Different</ui-button>`
		changes := []protocol.TextDocumentContentChangeEvent{
			{
				Range: nil, // Full document change
				Text:  newContent,
			},
		}

		// Test document manager handles full document changes
		updatedDoc := dm.UpdateDocumentWithChanges("file:///test-full.html", newContent, 2, changes)
		require.NotNil(t, updatedDoc, "Should get updated document")

		// Verify content was updated
		actualContent, err := updatedDoc.Content()
		require.NoError(t, err, "Should get updated content")
		assert.Equal(t, newContent, actualContent, "Content should be updated")

		// Verify tree is valid
		if updatedDoc.Tree() != nil {
			root := updatedDoc.Tree().RootNode()
			assert.NotNil(t, root, "Should have valid tree root")
		}
	})
}

func TestIncrementalParsingValidation(t *testing.T) {
	fixtureDir := filepath.Join("testdata", "integration", "incremental-parsing-test")

	dm, err := document.NewDocumentManager()
	require.NoError(t, err, "Should create document manager")
	defer dm.Close()

	t.Run("Parse Result Validation", func(t *testing.T) {
		// Use HTML fixture for validation test
		initialPath := filepath.Join(fixtureDir, "html-initial.html")
		initialContent, err := os.ReadFile(initialPath)
		require.NoError(t, err, "Should read initial fixture file")

		updatedPath := filepath.Join(fixtureDir, "html-updated.html")
		updatedContent, err := os.ReadFile(updatedPath)
		require.NoError(t, err, "Should read updated fixture file")

		doc := dm.OpenDocument("file:///test-validation.html", string(initialContent), 1)
		require.NotNil(t, doc, "Should create document")

		// Valid small change (primary -> secondary)
		changes := []protocol.TextDocumentContentChangeEvent{
			{
				Range: &protocol.Range{
					Start: protocol.Position{Line: 4, Character: 20}, // Start of "primary"
					End:   protocol.Position{Line: 4, Character: 27}, // End of "primary"
				},
				Text: "secondary",
			},
		}

		// Test incremental update through document manager
		updatedDoc := dm.UpdateDocumentWithChanges("file:///test-validation.html", string(updatedContent), 2, changes)
		require.NotNil(t, updatedDoc, "Should get updated document")

		// Validate tree and content
		if updatedDoc.Tree() != nil {
			root := updatedDoc.Tree().RootNode()
			assert.NotNil(t, root, "Should have root node")

			// Tree should span the entire content
			newContent, _ := updatedDoc.Content()
			contentLength := uint(len(newContent))
			assert.GreaterOrEqual(t, root.EndByte(), contentLength-10, "Tree should cover most of content")
		}
	})

	t.Run("Empty Document Handling", func(t *testing.T) {
		// Test with document that starts empty
		doc := dm.OpenDocument("file:///empty.html", "", 1)
		require.NotNil(t, doc, "Should create document")

		// Add content to empty document
		newContent := "<my-card>New content</my-card>"
		changes := []protocol.TextDocumentContentChangeEvent{
			{
				Range: &protocol.Range{
					Start: protocol.Position{Line: 0, Character: 0},
					End:   protocol.Position{Line: 0, Character: 0},
				},
				Text: newContent,
			},
		}

		// Should handle empty to content transition
		updatedDoc := dm.UpdateDocumentWithChanges("file:///empty.html", newContent, 2, changes)
		require.NotNil(t, updatedDoc, "Should get updated document")

		// Verify content
		content, err := updatedDoc.Content()
		require.NoError(t, err, "Should get content")
		assert.Equal(t, newContent, content, "Content should be updated")

		// Verify tree exists
		assert.NotNil(t, updatedDoc.Tree(), "Should have tree after content added")
	})
}
