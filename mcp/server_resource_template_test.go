/*
Copyright © 2026 Benny Powers <web@bennypowers.com>

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
package mcp

import (
	"context"
	"testing"

	testworkspace "bennypowers.dev/cem/internal/platform/testutil/workspace"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Inline: protocol-level integration test, scalar assertions on resource lookups

func TestResourceTemplatesRegisteredCorrectly(t *testing.T) {
	workspace := testworkspace.NewMapWorkspaceContext(t, "./testdata/fixtures/multiple-elements-integration")
	err := workspace.Init()
	require.NoError(t, err)

	cemServer, err := NewServer(workspace)
	require.NoError(t, err)

	err = cemServer.registry.LoadManifests()
	require.NoError(t, err)

	ctx := context.Background()

	// Connect client to server via in-memory transport
	serverTransport, clientTransport := mcp.NewInMemoryTransports()

	serverSession, err := cemServer.server.Connect(ctx, serverTransport, nil)
	require.NoError(t, err)
	defer func() { _ = serverSession.Close() }()

	client := mcp.NewClient(&mcp.Implementation{
		Name:    "test-client",
		Version: "1.0.0",
	}, nil)

	clientSession, err := client.Connect(ctx, clientTransport, nil)
	require.NoError(t, err)
	defer func() { _ = clientSession.Close() }()

	t.Run("static resource cem://elements is readable", func(t *testing.T) {
		result, err := clientSession.ReadResource(ctx, &mcp.ReadResourceParams{
			URI: "cem://elements",
		})
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.NotEmpty(t, result.Contents)
	})

	t.Run("templated resource cem://element/{tagName} is readable", func(t *testing.T) {
		result, err := clientSession.ReadResource(ctx, &mcp.ReadResourceParams{
			URI: "cem://element/button-element",
		})
		require.NoError(t, err, "reading cem://element/button-element should not return an error")
		require.NotNil(t, result)
		assert.NotEmpty(t, result.Contents)
	})

	t.Run("templated resource with sub-path is readable", func(t *testing.T) {
		result, err := clientSession.ReadResource(ctx, &mcp.ReadResourceParams{
			URI: "cem://element/button-element/attributes",
		})
		require.NoError(t, err, "reading cem://element/button-element/attributes should not return an error")
		require.NotNil(t, result)
		assert.NotEmpty(t, result.Contents)
	})

	t.Run("resource templates appear in ListResourceTemplates", func(t *testing.T) {
		result, err := clientSession.ListResourceTemplates(ctx, &mcp.ListResourceTemplatesParams{})
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.NotEmpty(t, result.ResourceTemplates, "templated resources should appear in ListResourceTemplates")

		templateURIs := make(map[string]bool)
		for _, tmpl := range result.ResourceTemplates {
			templateURIs[tmpl.URITemplate] = true
		}
		assert.True(t, templateURIs["cem://element/{tagName}"], "cem://element/{tagName} should be a registered template")
	})

	t.Run("templated resource with non-existent parameter returns error", func(t *testing.T) {
		_, err := clientSession.ReadResource(ctx, &mcp.ReadResourceParams{
			URI: "cem://element/does-not-exist",
		})
		assert.Error(t, err, "requesting a non-existent element should return an error")
	})

	t.Run("static resources do not appear in ListResourceTemplates", func(t *testing.T) {
		result, err := clientSession.ListResourceTemplates(ctx, &mcp.ListResourceTemplatesParams{})
		require.NoError(t, err)
		require.NotNil(t, result)

		for _, tmpl := range result.ResourceTemplates {
			assert.NotEqual(t, "cem://elements", tmpl.URITemplate, "static resources should not be in template list")
		}
	})
}
