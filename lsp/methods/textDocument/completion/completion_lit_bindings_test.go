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
package completion_test

import (
	"encoding/json"
	"testing"

	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/methods/textDocument/completion"
	"bennypowers.dev/cem/lsp/testhelpers"
	M "bennypowers.dev/cem/manifest"
	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Inline: testing presence/absence of specific labels in completion results;
// golden would be non-deterministic (map iteration order) and less readable
// than direct label assertions.
func TestLitPropertyCompletionUsesFieldName(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()

	dm, err := document.NewDocumentManager()
	require.NoError(t, err)
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	manifestJSON := `{
		"schemaVersion": "1.0.0",
		"modules": [{
			"kind": "javascript-module",
			"path": "src/test.ts",
			"declarations": [{
				"kind": "class",
				"name": "TestElement",
				"customElement": true,
				"tagName": "test-element",
				"members": [
					{
						"kind": "field",
						"name": "myAttr",
						"type": {"text": "string"},
						"attribute": "my-attr"
					},
					{
						"kind": "field",
						"name": "disabled",
						"type": {"text": "boolean"},
						"attribute": "disabled"
					}
				],
				"attributes": [
					{
						"name": "my-attr",
						"type": {"text": "string"},
						"fieldName": "myAttr"
					},
					{
						"name": "disabled",
						"type": {"text": "boolean"},
						"fieldName": "disabled"
					}
				]
			}],
			"exports": [{
				"kind": "custom-element-definition",
				"name": "test-element",
				"declaration": {"name": "TestElement"}
			}]
		}]
	}`
	var pkg M.Package
	require.NoError(t, json.Unmarshal([]byte(manifestJSON), &pkg))
	ctx.AddManifest(&pkg)

	content := "const tpl = html`<test-element .`;"
	docURI := "file:///test.ts"
	doc := dm.OpenDocument(docURI, content, 1)
	ctx.AddDocument(docURI, doc)

	result, err := completion.Completion(ctx, &protocol.CompletionParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{URI: uri.URI(docURI)},
			Position:     protocol.Position{Line: 0, Character: 32},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, result)

	items := result
	require.NotEmpty(t, items, "should have property completions")

	byLabel := make(map[string]protocol.CompletionItem)
	for _, item := range items {
		byLabel[item.Label] = item
	}

	assert.Contains(t, byLabel, ".myAttr", "should offer .myAttr (fieldName), not .my-attr")
	assert.NotContains(t, byLabel, ".my-attr", "should NOT offer .my-attr (attribute name)")
	assert.Contains(t, byLabel, ".disabled", "should offer .disabled")

	if item, ok := byLabel[".myAttr"]; ok {
		require.False(t, item.InsertText.IsZero(), "InsertText should be set")
		v, _ := item.InsertText.Get()
		assert.Equal(t, "myAttr", v, "InsertText should be the field name")
	}
}

// Inline: testing label filtering (boolean-only) -- see TestLitPropertyCompletionUsesFieldName.
func TestLitBooleanCompletionFiltersToBoolean(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()

	dm, err := document.NewDocumentManager()
	require.NoError(t, err)
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	manifestJSON := `{
		"schemaVersion": "1.0.0",
		"modules": [{
			"kind": "javascript-module",
			"path": "src/test.ts",
			"declarations": [{
				"kind": "class",
				"name": "TestElement",
				"customElement": true,
				"tagName": "test-element",
				"attributes": [
					{
						"name": "variant",
						"type": {"text": "string"}
					},
					{
						"name": "disabled",
						"type": {"text": "boolean"}
					}
				]
			}],
			"exports": [{
				"kind": "custom-element-definition",
				"name": "test-element",
				"declaration": {"name": "TestElement"}
			}]
		}]
	}`
	var pkg M.Package
	require.NoError(t, json.Unmarshal([]byte(manifestJSON), &pkg))
	ctx.AddManifest(&pkg)

	content := "const tpl = html`<test-element ?`;"
	docURI := "file:///test.ts"
	doc := dm.OpenDocument(docURI, content, 1)
	ctx.AddDocument(docURI, doc)

	result, err := completion.Completion(ctx, &protocol.CompletionParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{URI: uri.URI(docURI)},
			Position:     protocol.Position{Line: 0, Character: 32},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, result)

	items := result

	for _, item := range items {
		assert.NotEqual(t, "?variant", item.Label, "non-boolean 'variant' should not appear in ? completions")
	}

	labels := make(map[string]bool)
	for _, item := range items {
		labels[item.Label] = true
	}
	assert.True(t, labels["?disabled"], "boolean 'disabled' should appear in ? completions")
}

// Inline: testing label filtering (events-only) -- see TestLitPropertyCompletionUsesFieldName.
func TestLitEventCompletionFiltersToEvents(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()

	dm, err := document.NewDocumentManager()
	require.NoError(t, err)
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	manifestJSON := `{
		"schemaVersion": "1.0.0",
		"modules": [{
			"kind": "javascript-module",
			"path": "src/test.ts",
			"declarations": [{
				"kind": "class",
				"name": "TestElement",
				"customElement": true,
				"tagName": "test-element",
				"attributes": [
					{"name": "disabled", "type": {"text": "boolean"}}
				],
				"events": [
					{"name": "change", "type": {"text": "CustomEvent"}, "description": "Fired on change"},
					{"name": "input", "type": {"text": "Event"}}
				]
			}],
			"exports": [{
				"kind": "custom-element-definition",
				"name": "test-element",
				"declaration": {"name": "TestElement"}
			}]
		}]
	}`
	var pkg M.Package
	require.NoError(t, json.Unmarshal([]byte(manifestJSON), &pkg))
	ctx.AddManifest(&pkg)

	content := "const tpl = html`<test-element @`;"
	docURI := "file:///test.ts"
	doc := dm.OpenDocument(docURI, content, 1)
	ctx.AddDocument(docURI, doc)

	result, err := completion.Completion(ctx, &protocol.CompletionParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{URI: uri.URI(docURI)},
			Position:     protocol.Position{Line: 0, Character: 32},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, result)

	items := result

	labels := make(map[string]bool)
	for _, item := range items {
		labels[item.Label] = true
	}

	assert.True(t, labels["@change"], "should offer @change event")
	assert.True(t, labels["@input"], "should offer @input event")
	assert.False(t, labels["disabled"], "should NOT offer attributes in @ completions")
	assert.False(t, labels["@disabled"], "should NOT offer @disabled (not an event)")
}
