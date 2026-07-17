/*
Copyright © 2026 Benny Powers

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
package document_test

import (
	"testing"

	_ "bennypowers.dev/cem/internal/languages/registry"
	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/types"
	"go.lsp.dev/protocol"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Inline: lifecycle state transitions
// DocumentManager open/close/update lifecycle and incremental editing.

func getContent(t *testing.T, doc types.Document) string {
	t.Helper()
	content, err := doc.Content()
	require.NoError(t, err)
	return content
}

func TestNewDocumentManager(t *testing.T) {
	dm, err := document.NewDocumentManager()
	require.NoError(t, err)
	defer dm.Close()
	assert.NotNil(t, dm.QueryManager())
}

func TestDocumentManager_OpenAndRetrieve(t *testing.T) {
	dm, err := document.NewDocumentManager()
	require.NoError(t, err)
	defer dm.Close()

	doc := dm.OpenDocument("file:///test.html", "<div>hello</div>", 1)
	require.NotNil(t, doc)
	assert.Equal(t, "file:///test.html", doc.URI())
	assert.Equal(t, "<div>hello</div>", getContent(t, doc))
	assert.Equal(t, int32(1), doc.Version())

	retrieved := dm.Document("file:///test.html")
	assert.NotNil(t, retrieved)
	assert.Equal(t, doc.URI(), retrieved.URI())
}

func TestDocumentManager_OpenTypeScript(t *testing.T) {
	dm, err := document.NewDocumentManager()
	require.NoError(t, err)
	defer dm.Close()

	doc := dm.OpenDocument("file:///test.ts", "const x = 1;", 1)
	require.NotNil(t, doc)
	assert.Equal(t, "typescript", doc.Language())
}

func TestDocumentManager_OpenTSX(t *testing.T) {
	dm, err := document.NewDocumentManager()
	require.NoError(t, err)
	defer dm.Close()

	doc := dm.OpenDocument("file:///test.tsx", "const el = <div/>;", 1)
	require.NotNil(t, doc)
	assert.Equal(t, "tsx", doc.Language())
}

func TestDocumentManager_DocumentNotFound(t *testing.T) {
	dm, err := document.NewDocumentManager()
	require.NoError(t, err)
	defer dm.Close()

	assert.Nil(t, dm.Document("file:///nonexistent.html"))
}

func TestDocumentManager_AllDocuments(t *testing.T) {
	dm, err := document.NewDocumentManager()
	require.NoError(t, err)
	defer dm.Close()

	assert.Empty(t, dm.AllDocuments())

	dm.OpenDocument("file:///a.html", "<a>", 1)
	dm.OpenDocument("file:///b.html", "<b>", 1)

	docs := dm.AllDocuments()
	assert.Len(t, docs, 2)
}

func TestDocumentManager_CloseDocument(t *testing.T) {
	dm, err := document.NewDocumentManager()
	require.NoError(t, err)
	defer dm.Close()

	dm.OpenDocument("file:///test.html", "<div>hello</div>", 1)
	assert.NotNil(t, dm.Document("file:///test.html"))

	dm.CloseDocument("file:///test.html")
	assert.Nil(t, dm.Document("file:///test.html"))
}

func TestDocumentManager_UpdateDocument(t *testing.T) {
	dm, err := document.NewDocumentManager()
	require.NoError(t, err)
	defer dm.Close()

	dm.OpenDocument("file:///test.html", "<div>old</div>", 1)

	updated := dm.UpdateDocument("file:///test.html", "<div>new</div>", 2)
	require.NotNil(t, updated)
	assert.Equal(t, "<div>new</div>", getContent(t, updated))
	assert.Equal(t, int32(2), updated.Version())
}

func TestDocumentManager_UpdateCreatesIfMissing(t *testing.T) {
	dm, err := document.NewDocumentManager()
	require.NoError(t, err)
	defer dm.Close()

	doc := dm.UpdateDocument("file:///new.html", "<p>created</p>", 1)
	require.NotNil(t, doc, "UpdateDocument on missing document should create it")
	assert.Equal(t, "<p>created</p>", getContent(t, doc))
	assert.Equal(t, int32(1), doc.Version())

	retrieved := dm.Document("file:///new.html")
	require.NotNil(t, retrieved, "created document should be retrievable")
	assert.Equal(t, doc.URI(), retrieved.URI())
}

func TestDocumentManager_ReopenReplacesExisting(t *testing.T) {
	dm, err := document.NewDocumentManager()
	require.NoError(t, err)
	defer dm.Close()

	dm.OpenDocument("file:///test.html", "<div>first</div>", 1)
	doc2 := dm.OpenDocument("file:///test.html", "<div>second</div>", 2)
	assert.Equal(t, "<div>second</div>", getContent(t, doc2))
	assert.Len(t, dm.AllDocuments(), 1)
}

func TestDocumentManager_UpdateWithChanges_FullReplace(t *testing.T) {
	dm, err := document.NewDocumentManager()
	require.NoError(t, err)
	defer dm.Close()

	dm.OpenDocument("file:///test.html", "<div>old</div>", 1)

	changes := []protocol.TextDocumentContentChangeEvent{
		&protocol.TextDocumentContentChangeWholeDocument{Text: "<div>new</div>"},
	}
	doc := dm.UpdateDocumentWithChanges("file:///test.html", "<div>new</div>", 2, changes)
	require.NotNil(t, doc)
	assert.Equal(t, "<div>new</div>", getContent(t, doc))
}

func TestDocumentManager_UpdateWithChanges_NilChanges(t *testing.T) {
	dm, err := document.NewDocumentManager()
	require.NoError(t, err)
	defer dm.Close()

	dm.OpenDocument("file:///test.html", "<div>hello</div>", 1)

	doc := dm.UpdateDocumentWithChanges("file:///test.html", "<div>world</div>", 2, nil)
	require.NotNil(t, doc)
	assert.Equal(t, "<div>world</div>", getContent(t, doc))
}

func TestDocumentManager_UpdateWithChanges_IncrementalEdit(t *testing.T) {
	dm, err := document.NewDocumentManager()
	require.NoError(t, err)
	defer dm.Close()

	dm.OpenDocument("file:///test.html", "<div>hello</div>", 1)

	changes := []protocol.TextDocumentContentChangeEvent{
		&protocol.TextDocumentContentChangePartial{
			Range: protocol.Range{
				Start: protocol.Position{Line: 0, Character: 5},
				End:   protocol.Position{Line: 0, Character: 10},
			},
			Text: "world",
		},
	}
	doc := dm.UpdateDocumentWithChanges("file:///test.html", "<div>world</div>", 2, changes)
	require.NotNil(t, doc)
	assert.Equal(t, int32(2), doc.Version())
	assert.Equal(t, "<div>world</div>", getContent(t, doc))
}

func TestDocumentManager_QueryManager(t *testing.T) {
	dm, err := document.NewDocumentManager()
	require.NoError(t, err)
	defer dm.Close()

	qm := dm.QueryManager()
	assert.NotNil(t, qm)
}

func TestDocumentManager_CloseNonexistent(t *testing.T) {
	dm, err := document.NewDocumentManager()
	require.NoError(t, err)
	defer dm.Close()

	dm.CloseDocument("file:///nonexistent.html")
}
