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
package configuration_test

import (
	"testing"

	"bennypowers.dev/cem/lsp/methods/workspace/configuration"
	"bennypowers.dev/cem/lsp/testhelpers"
	protocol "github.com/bennypowers/glsp/protocol_3_17"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDidChangeConfiguration_NilSettings(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	err := configuration.DidChangeConfiguration(ctx, nil, &protocol.DidChangeConfigurationParams{
		Settings: nil,
	})
	require.NoError(t, err)
	assert.True(t, ctx.InlayHintsEnabled())
}

func TestDidChangeConfiguration_DisableInlayHints(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	err := configuration.DidChangeConfiguration(ctx, nil, &protocol.DidChangeConfigurationParams{
		Settings: map[string]any{
			"cem": map[string]any{
				"inlayHints": false,
			},
		},
	})
	require.NoError(t, err)
	assert.False(t, ctx.InlayHintsEnabled())
}

func TestDidChangeConfiguration_EnableInlayHints(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	disabled := false
	cfg := ctx.Config()
	cfg.InlayHints = &disabled
	ctx.SetConfig(cfg)
	assert.False(t, ctx.InlayHintsEnabled())

	err := configuration.DidChangeConfiguration(ctx, nil, &protocol.DidChangeConfigurationParams{
		Settings: map[string]any{
			"cem": map[string]any{
				"inlayHints": true,
			},
		},
	})
	require.NoError(t, err)
	assert.True(t, ctx.InlayHintsEnabled())
}

func TestDidChangeConfiguration_UnknownNamespace(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	err := configuration.DidChangeConfiguration(ctx, nil, &protocol.DidChangeConfigurationParams{
		Settings: map[string]any{
			"other-tool": map[string]any{
				"inlayHints": false,
			},
		},
	})
	require.NoError(t, err)
	assert.True(t, ctx.InlayHintsEnabled())
}

func TestDidChangeConfiguration_UnknownNamespacePreservesPriorConfig(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	disabled := false
	cfg := ctx.Config()
	cfg.InlayHints = &disabled
	ctx.SetConfig(cfg)
	assert.False(t, ctx.InlayHintsEnabled())

	err := configuration.DidChangeConfiguration(ctx, nil, &protocol.DidChangeConfigurationParams{
		Settings: map[string]any{
			"python": map[string]any{
				"formatting": true,
			},
		},
	})
	require.NoError(t, err)
	assert.False(t, ctx.InlayHintsEnabled(), "non-cem config change must not reset prior settings")
}

func TestDidChangeConfiguration_NonMapSettings(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()
	err := configuration.DidChangeConfiguration(ctx, nil, &protocol.DidChangeConfigurationParams{
		Settings: "invalid",
	})
	require.NoError(t, err)
	assert.True(t, ctx.InlayHintsEnabled())
}
