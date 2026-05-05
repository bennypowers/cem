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
package inlayHint_test

import (
	"encoding/json"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/methods/textDocument/inlayHint"
	"bennypowers.dev/cem/lsp/testhelpers"
	"bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
	protocol "github.com/bennypowers/glsp/protocol_3_17"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type manifestFixture struct {
	Attributes map[string]map[string]M.Attribute `json:"attributes"`
	Slots      map[string][]M.Slot               `json:"slots"`
}

func fullRange() protocol.Range {
	return protocol.Range{
		Start: protocol.Position{Line: 0, Character: 0},
		End:   protocol.Position{Line: 1000, Character: 0},
	}
}

func TestInlayHint_Fixtures(t *testing.T) {
	testutil.RunLSPFixtures(t, "testdata", func(t *testing.T, fixture *testutil.LSPFixture) {
		ctx := testhelpers.NewMockServerContext()

		dm, err := document.NewDocumentManager()
		require.NoError(t, err)
		defer dm.Close()
		ctx.SetDocumentManager(dm)

		doc := dm.OpenDocument("test.html", fixture.InputHTML, 1)
		ctx.AddDocument("test.html", doc)

		if len(fixture.Manifest) > 0 {
			var mf manifestFixture
			require.NoError(t, json.Unmarshal(fixture.Manifest, &mf))

			for tagName, attrs := range mf.Attributes {
				attrMap := make(map[string]*M.Attribute)
				for attrName, attr := range attrs {
					attr.Name = attrName
					a := attr
					attrMap[attrName] = &a
				}
				ctx.AddAttributes(tagName, attrMap)
			}

			for tagName, slots := range mf.Slots {
				ctx.AddSlots(tagName, slots)
			}
		}

		if settingsRaw, ok := fixture.ExpectedMap["settings"]; ok {
			if settingsMap, ok := settingsRaw.(map[string]any); ok {
				if inlayHints, ok := settingsMap["inlayHints"].(bool); ok {
					cfg := ctx.Config()
					cfg.InlayHints = &inlayHints
					ctx.SetConfig(cfg)
				}
			}
		}

		result, err := inlayHint.InlayHint(ctx, nil, &protocol.InlayHintParams{
			TextDocument: protocol.TextDocumentIdentifier{URI: "test.html"},
			Range:        fullRange(),
		})
		require.NoError(t, err)

		var expected []protocol.InlayHint
		require.NoError(t, fixture.GetExpected("expected", &expected))
		require.Len(t, result, len(expected), "hint count mismatch")

		expectedLabels := make(map[string]bool, len(expected))
		for _, e := range expected {
			if label, ok := e.Label.(string); ok {
				expectedLabels[label] = true
			}
		}
		for _, hint := range result {
			label, ok := hint.Label.(string)
			require.True(t, ok, "hint label should be string")
			assert.True(t, expectedLabels[label], "unexpected hint label: %s", label)
		}
	})
}

func TestInlayHint_NilDocument(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	result, err := inlayHint.InlayHint(ctx, nil, &protocol.InlayHintParams{
		TextDocument: protocol.TextDocumentIdentifier{URI: "nonexistent.html"},
		Range:        fullRange(),
	})
	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestInlayHint_EnabledByDefault(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	assert.True(t, ctx.InlayHintsEnabled())
}

func TestConfig_DefaultEnabled(t *testing.T) {
	cfg := types.DefaultConfig()
	require.NotNil(t, cfg.InlayHints)
	assert.True(t, *cfg.InlayHints)
}

func TestConfig_NilMeansEnabled(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	ctx.SetConfig(types.ServerConfig{InlayHints: nil})
	assert.True(t, ctx.InlayHintsEnabled())
}
