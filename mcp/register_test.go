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
	"encoding/json"
	"sort"
	"testing"

	testworkspace "bennypowers.dev/cem/internal/platform/testutil/workspace"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Inline: protocol-level integration tests. The behavior under test is the
// surface registered by Register compared with NewServerWithConfig, which is
// discovered through a real client session rather than a static golden file.

func connectRegisterTestClient(t *testing.T, server *mcp.Server) *mcp.ClientSession {
	t.Helper()

	ctx := context.Background()
	serverTransport, clientTransport := mcp.NewInMemoryTransports()

	serverSession, err := server.Connect(ctx, serverTransport, nil)
	require.NoError(t, err)
	t.Cleanup(func() { _ = serverSession.Close() })

	client := mcp.NewClient(&mcp.Implementation{
		Name:    "register-test-client",
		Version: "1.0.0",
	}, nil)

	clientSession, err := client.Connect(ctx, clientTransport, nil)
	require.NoError(t, err)
	t.Cleanup(func() { _ = clientSession.Close() })

	return clientSession
}

func listRegisterTestToolNames(t *testing.T, session *mcp.ClientSession) []string {
	t.Helper()

	result, err := session.ListTools(context.Background(), &mcp.ListToolsParams{})
	require.NoError(t, err)

	names := make([]string, 0, len(result.Tools))
	for _, tool := range result.Tools {
		names = append(names, tool.Name)
	}
	sort.Strings(names)
	return names
}

func listRegisterTestToolSchemas(t *testing.T, session *mcp.ClientSession) map[string]string {
	t.Helper()

	result, err := session.ListTools(context.Background(), &mcp.ListToolsParams{})
	require.NoError(t, err)

	schemas := make(map[string]string, len(result.Tools))
	for _, tool := range result.Tools {
		schema, err := json.Marshal(tool.InputSchema)
		require.NoError(t, err)
		schemas[tool.Name] = string(schema)
	}
	return schemas
}

func listRegisterTestResourceURIs(t *testing.T, session *mcp.ClientSession) []string {
	t.Helper()

	result, err := session.ListResources(context.Background(), &mcp.ListResourcesParams{})
	require.NoError(t, err)

	uris := make([]string, 0, len(result.Resources))
	for _, resource := range result.Resources {
		uris = append(uris, resource.URI)
	}
	sort.Strings(uris)
	return uris
}

func listRegisterTestResourceTemplateURIs(t *testing.T, session *mcp.ClientSession) []string {
	t.Helper()

	result, err := session.ListResourceTemplates(context.Background(), &mcp.ListResourceTemplatesParams{})
	require.NoError(t, err)

	uris := make([]string, 0, len(result.ResourceTemplates))
	for _, template := range result.ResourceTemplates {
		uris = append(uris, template.URITemplate)
	}
	sort.Strings(uris)
	return uris
}

// TestRegisterMatchesNewServerSurface asserts that mounting cem on a
// caller-owned server exposes exactly the tools and resources that
// NewServerWithConfig exposes.
func TestRegisterMatchesNewServerSurface(t *testing.T) {
	workspace := testworkspace.NewMapWorkspaceContext(t, "./testdata/fixtures/basic-integration")
	require.NoError(t, workspace.Init())

	config := ServerConfig{MaxDescriptionLength: 2000}

	cemServer, err := NewServerWithConfig(workspace, config)
	require.NoError(t, err)
	t.Cleanup(func() {
		_ = cemServer.Registry().LSPRegistry().StopFileWatching()
	})

	hostServer := mcp.NewServer(&mcp.Implementation{
		Name:    "host-server",
		Version: "test",
	}, nil)
	require.NoError(t, Register(hostServer, workspace, config))

	cemClient := connectRegisterTestClient(t, cemServer.server)
	hostClient := connectRegisterTestClient(t, hostServer)

	t.Run("tool names match", func(t *testing.T) {
		cemTools := listRegisterTestToolNames(t, cemClient)
		hostTools := listRegisterTestToolNames(t, hostClient)

		assert.NotEmpty(t, cemTools, "cem server should register tools")
		assert.Equal(t, cemTools, hostTools)
	})

	t.Run("tool input schemas match", func(t *testing.T) {
		assert.Equal(t,
			listRegisterTestToolSchemas(t, cemClient),
			listRegisterTestToolSchemas(t, hostClient),
		)
	})

	t.Run("resource URIs match", func(t *testing.T) {
		cemResources := listRegisterTestResourceURIs(t, cemClient)
		hostResources := listRegisterTestResourceURIs(t, hostClient)

		assert.NotEmpty(t, cemResources, "cem server should register resources")
		assert.Equal(t, cemResources, hostResources)
	})

	t.Run("resource template URIs match", func(t *testing.T) {
		cemTemplates := listRegisterTestResourceTemplateURIs(t, cemClient)
		hostTemplates := listRegisterTestResourceTemplateURIs(t, hostClient)

		assert.NotEmpty(t, cemTemplates, "cem server should register resource templates")
		assert.Equal(t, cemTemplates, hostTemplates)
	})
}

// TestRegisterRejectsNilArguments covers the exported API's argument
// validation so callers get an error rather than a panic.
func TestRegisterRejectsNilArguments(t *testing.T) {
	workspace := testworkspace.NewMapWorkspaceContext(t, "./testdata/fixtures/basic-integration")
	require.NoError(t, workspace.Init())

	t.Run("nil server", func(t *testing.T) {
		err := Register(nil, workspace, ServerConfig{})
		require.ErrorContains(t, err, "server must not be nil")
	})

	t.Run("nil workspace", func(t *testing.T) {
		server := mcp.NewServer(&mcp.Implementation{
			Name:    "host-server",
			Version: "test",
		}, nil)
		err := Register(server, nil, ServerConfig{})
		require.ErrorContains(t, err, "workspace must not be nil")
	})

	t.Run("typed-nil workspace", func(t *testing.T) {
		var typedNil *testworkspace.MapWorkspaceContext
		server := mcp.NewServer(&mcp.Implementation{
			Name:    "host-server",
			Version: "test",
		}, nil)
		err := Register(server, typedNil, ServerConfig{})
		require.ErrorContains(t, err, "workspace must not be nil")
	})
}

// TestRegisterAdditionalPackages exercises the AdditionalPackages branch
// through a local fixture path. The fixture is read from disk because
// additional packages resolve through the package workspace loader rather
// than the test MapFS.
func TestRegisterAdditionalPackages(t *testing.T) {
	workspace := testworkspace.NewMapWorkspaceContext(t, "./testdata/fixtures/basic-integration")
	require.NoError(t, workspace.Init())

	hostServer := mcp.NewServer(&mcp.Implementation{
		Name:    "host-server",
		Version: "test",
	}, nil)
	require.NoError(t, Register(hostServer, workspace, ServerConfig{
		MaxDescriptionLength: 2000,
		AdditionalPackages:   []string{"./testdata/fixtures/multiple-elements-integration"},
	}))

	client := connectRegisterTestClient(t, hostServer)

	result, err := client.ReadResource(context.Background(), &mcp.ReadResourceParams{
		URI: "cem://element/button-element",
	})
	require.NoError(t, err, "an element from an additional package should be readable")
	require.NotEmpty(t, result.Contents)
	assert.NotEmpty(t, result.Contents[0].Text)
}

// TestRegisterToolCall works through a caller-owned server and verifies that
// registered handlers are wired to a usable MCP context.
func TestRegisterToolCall(t *testing.T) {
	workspace := testworkspace.NewMapWorkspaceContext(t, "./testdata/fixtures/basic-integration")
	require.NoError(t, workspace.Init())

	hostServer := mcp.NewServer(&mcp.Implementation{
		Name:    "host-server",
		Version: "test",
	}, nil)
	require.NoError(t, Register(hostServer, workspace, ServerConfig{MaxDescriptionLength: 2000}))

	client := connectRegisterTestClient(t, hostServer)

	result, err := client.CallTool(context.Background(), &mcp.CallToolParams{
		Name:      "generate_config",
		Arguments: map[string]any{"focus": "generate"},
	})
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.False(t, result.IsError, "generate_config should succeed on the fixture workspace")
	require.NotEmpty(t, result.Content)

	text, ok := result.Content[0].(*mcp.TextContent)
	require.True(t, ok, "generate_config should return text content")
	assert.Contains(t, text.Text, "Current config")
	assert.Contains(t, text.Text, "Schema: generate")
}

// TestRegisterResourceRead works through a caller-owned server and verifies
// that registered resource handlers can read from the loaded manifest.
func TestRegisterResourceRead(t *testing.T) {
	workspace := testworkspace.NewMapWorkspaceContext(t, "./testdata/fixtures/basic-integration")
	require.NoError(t, workspace.Init())

	hostServer := mcp.NewServer(&mcp.Implementation{
		Name:    "host-server",
		Version: "test",
	}, nil)
	require.NoError(t, Register(hostServer, workspace, ServerConfig{MaxDescriptionLength: 2000}))

	client := connectRegisterTestClient(t, hostServer)

	result, err := client.ReadResource(context.Background(), &mcp.ReadResourceParams{
		URI: "cem://element/test-element",
	})
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotEmpty(t, result.Contents)
	assert.NotEmpty(t, result.Contents[0].Text)
}
