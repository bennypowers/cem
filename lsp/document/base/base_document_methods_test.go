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
package base_test

import (
	"testing"

	"bennypowers.dev/cem/lsp/document/base"
	"bennypowers.dev/cem/lsp/types"
	protocol "github.com/bennypowers/glsp/protocol_3_17"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Inline: pure function, scalar assertions
// Getters, setters, and nil-safety checks on BaseDocument public API.

func newTestDoc(content string) *base.BaseDocument {
	return base.NewBaseDocument("file:///test.html", content, 1, "html", nil, nil)
}

func TestBaseDocument_Getters(t *testing.T) {
	doc := newTestDoc("<div>hello</div>")
	assert.Equal(t, "file:///test.html", doc.URI())
	content, err := doc.Content()
	require.NoError(t, err)
	assert.Equal(t, "<div>hello</div>", content)
	assert.Equal(t, int32(1), doc.Version())
	assert.Equal(t, "html", doc.Language())
}

func TestBaseDocument_Tree_NilByDefault(t *testing.T) {
	doc := newTestDoc("<div>hello</div>")
	assert.Nil(t, doc.Tree())
}

func TestBaseDocument_UpdateContent(t *testing.T) {
	doc := newTestDoc("<div>old</div>")
	doc.UpdateContent("<div>new</div>", 2)
	content, err := doc.Content()
	require.NoError(t, err)
	assert.Equal(t, "<div>new</div>", content)
	assert.Equal(t, int32(2), doc.Version())
}

func TestBaseDocument_ScriptTags(t *testing.T) {
	doc := newTestDoc("<div>hello</div>")
	assert.Empty(t, doc.ScriptTags())

	tags := []types.ScriptTag{{Type: "module", IsModule: true}}
	doc.SetScriptTags(tags)
	assert.Len(t, doc.ScriptTags(), 1)
}

func TestBaseDocument_ImportMap(t *testing.T) {
	doc := newTestDoc("<div>hello</div>")
	assert.Nil(t, doc.ImportMap())
}

func TestBaseDocument_CompletionPrefix(t *testing.T) {
	doc := newTestDoc("<div>hello</div>")
	assert.Empty(t, doc.CompletionPrefix(nil))
}

func TestBaseDocument_FindModuleScript(t *testing.T) {
	doc := newTestDoc("<div>hello</div>")
	_, found := doc.FindModuleScript()
	assert.False(t, found)

	doc.SetScriptTags([]types.ScriptTag{
		{Type: "text/javascript"},
		{Type: "module", IsModule: true},
	})
	_, found = doc.FindModuleScript()
	assert.True(t, found)
}

func TestBaseDocument_FindInlineModuleScript(t *testing.T) {
	doc := newTestDoc("<div>hello</div>")
	_, found := doc.FindInlineModuleScript()
	assert.False(t, found)

	doc.SetScriptTags([]types.ScriptTag{
		{Type: "module", IsModule: true, Src: "external.js"},
		{Type: "module", IsModule: true, Src: ""},
	})
	_, found = doc.FindInlineModuleScript()
	assert.True(t, found)
}

func TestBaseDocument_TreeAndContent(t *testing.T) {
	doc := newTestDoc("<div>hello</div>")
	tree, content, release := doc.TreeAndContent()
	defer release()
	assert.Nil(t, tree)
	assert.Equal(t, "<div>hello</div>", content)
}

func TestBaseDocument_ByteRangeToProtocolRange(t *testing.T) {
	doc := newTestDoc("line1\nline2\nline3")
	content := "line1\nline2\nline3"
	r := doc.ByteRangeToProtocolRange(content, 6, 11)
	assert.Equal(t, uint32(1), r.Start.Line)
	assert.Equal(t, uint32(0), r.Start.Character)
	assert.Equal(t, uint32(1), r.End.Line)
	assert.Equal(t, uint32(5), r.End.Character)
}

func TestBaseDocument_FindElementAtPosition_NilHandler(t *testing.T) {
	doc := newTestDoc("<div>hello</div>")
	assert.Nil(t, doc.FindElementAtPosition(protocol.Position{}, nil))
}

func TestBaseDocument_FindCustomElements_NilHandler(t *testing.T) {
	doc := newTestDoc("<div>hello</div>")
	elements, err := doc.FindCustomElements(nil)
	assert.NoError(t, err)
	assert.Empty(t, elements)
}

func TestBaseDocument_AnalyzeCompletionContextTS_NilHandler(t *testing.T) {
	doc := newTestDoc("<div>hello</div>")
	result := doc.AnalyzeCompletionContextTS(protocol.Position{}, nil)
	assert.NotNil(t, result)
	assert.Equal(t, types.CompletionUnknown, result.Type)
}
