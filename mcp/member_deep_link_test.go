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

// Inline: protocol-level integration test; the SDK's anchored URI-template
// matching is the behavior under test, so the reads must go through a real
// client session rather than calling resource handlers directly

// TestMemberDeepLinkRouting drives cem://element/{tagName}/methods/{memberName}
// through the SDK transport. The go-sdk matches resource URIs against
// registered templates with an anchored regexp where each {var} matches a
// single path segment, so the parent methods resource cannot route
// four-segment URIs; the element-methods-member template exists for them.
// Regression test for the deep links documented in element-methods.md.
func TestMemberDeepLinkRouting(t *testing.T) {
	workspace := testworkspace.NewMapWorkspaceContext(t, "./testdata/fixtures/element-methods-integration")
	err := workspace.Init()
	require.NoError(t, err)

	cemServer, err := NewServer(workspace)
	require.NoError(t, err)

	err = cemServer.registry.LoadManifests()
	require.NoError(t, err)

	ctx := context.Background()

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

	t.Run("member deep link routes and returns JSON", func(t *testing.T) {
		result, err := clientSession.ReadResource(ctx, &mcp.ReadResourceParams{
			URI: "cem://element/counter-element/methods/increment",
		})
		require.NoError(t, err, "member deep link should route through the SDK")
		require.NotNil(t, result)
		require.Len(t, result.Contents, 1)
		assert.Equal(t, "application/json", result.Contents[0].MIMEType)
		assert.Contains(t, result.Contents[0].Text, `"name": "increment"`)
	})

	t.Run("field deep link routes", func(t *testing.T) {
		result, err := clientSession.ReadResource(ctx, &mcp.ReadResourceParams{
			URI: "cem://element/counter-element/methods/count",
		})
		require.NoError(t, err, "field deep link should route through the SDK")
		require.Len(t, result.Contents, 1)
		assert.Contains(t, result.Contents[0].Text, `"name": "count"`)
	})

	t.Run("parent methods resource still routes", func(t *testing.T) {
		result, err := clientSession.ReadResource(ctx, &mcp.ReadResourceParams{
			URI: "cem://element/counter-element/methods",
		})
		require.NoError(t, err, "three-segment URI should route to the markdown resource")
		require.Len(t, result.Contents, 1)
		assert.Equal(t, "text/markdown", result.Contents[0].MIMEType)
	})

	t.Run("unknown member errors at protocol level", func(t *testing.T) {
		_, err := clientSession.ReadResource(ctx, &mcp.ReadResourceParams{
			URI: "cem://element/counter-element/methods/nonexistent",
		})
		assert.Error(t, err, "unknown member should return an error")
	})
}
