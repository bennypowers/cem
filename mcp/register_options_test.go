/*
Copyright © 2026 Benny Powers <web@bennypowers.com>

SPDX-License-Identifier: GPL-3.0-or-later

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
	"fmt"
	"io/fs"
	"path/filepath"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"bennypowers.dev/cem/internal/platform"
	testworkspace "bennypowers.dev/cem/internal/platform/testutil/workspace"
	cemworkspace "bennypowers.dev/cem/internal/workspace"
	M "bennypowers.dev/cem/manifest"
	cemtypes "bennypowers.dev/cem/types"
	fileworkspace "bennypowers.dev/cem/workspace"
	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Inline: these tests assert exact protocol surfaces and handler outcomes that
// Register itself produces; goldens would add indirection without improving
// diagnosis of a failure.

func newRegisterOptionsTestServer(t *testing.T) *mcp.Server {
	t.Helper()
	return mcp.NewServer(&mcp.Implementation{
		Name:    "host-server",
		Version: "test",
	}, nil)
}

func newRegisterOptionsTestWorkspace(t *testing.T) *testworkspace.MapWorkspaceContext {
	t.Helper()
	workspace := testworkspace.NewMapWorkspaceContext(t, "./testdata/fixtures/basic-integration")
	require.NoError(t, workspace.Init())
	return workspace
}

func addHostTool(server *mcp.Server, name, text string) {
	server.AddTool(&mcp.Tool{
		Name:        name,
		Description: "host tool",
		InputSchema: &jsonschema.Schema{Type: "object"},
	}, func(context.Context, *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: text}},
		}, nil
	})
}

func addHostResource(server *mcp.Server, uri, text string) {
	server.AddResource(&mcp.Resource{
		URI:  uri,
		Name: "host-resource",
	}, func(context.Context, *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
		return &mcp.ReadResourceResult{
			Contents: []*mcp.ResourceContents{{URI: uri, Text: text}},
		}, nil
	})
}

func addHostResourceTemplate(server *mcp.Server, uriTemplate, text string) {
	server.AddResourceTemplate(&mcp.ResourceTemplate{
		URITemplate: uriTemplate,
		Name:        "host-template",
	}, func(_ context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
		return &mcp.ReadResourceResult{
			Contents: []*mcp.ResourceContents{{URI: req.Params.URI, Text: text}},
		}, nil
	})
}

func callRegisterOptionsTestTool(t *testing.T, client *mcp.ClientSession, name string, args map[string]any) *mcp.CallToolResult {
	t.Helper()
	result, err := client.CallTool(context.Background(), &mcp.CallToolParams{
		Name:      name,
		Arguments: args,
	})
	require.NoError(t, err)
	require.NotNil(t, result)
	return result
}

func readRegisterOptionsTestResource(t *testing.T, client *mcp.ClientSession, uri string) string {
	t.Helper()
	result, err := client.ReadResource(context.Background(), &mcp.ReadResourceParams{URI: uri})
	require.NoError(t, err)
	require.NotEmpty(t, result.Contents)
	return result.Contents[0].Text
}

func TestRegisterToolFilters(t *testing.T) {
	workspace := newRegisterOptionsTestWorkspace(t)
	config := ServerConfig{MaxDescriptionLength: 2000}

	t.Run("include selects exact names", func(t *testing.T) {
		server := newRegisterOptionsTestServer(t)
		handle, err := Register(server, workspace, config, WithTools(Selection{
			Include: []string{"generate_config", "validate_html"},
		}))
		require.NoError(t, err)
		require.NotNil(t, handle)

		client := connectRegisterTestClient(t, server)
		assert.Equal(t, []string{"generate_config", "validate_html"}, listRegisterTestToolNames(t, client))
	})

	t.Run("exclude wins over include", func(t *testing.T) {
		server := newRegisterOptionsTestServer(t)
		_, err := Register(server, workspace, config, WithTools(Selection{
			Include: []string{"generate_config", "validate_html"},
			Exclude: []string{"validate_html"},
		}))
		require.NoError(t, err)

		client := connectRegisterTestClient(t, server)
		assert.Equal(t, []string{"generate_config"}, listRegisterTestToolNames(t, client))
	})

	t.Run("exclude everything from a category", func(t *testing.T) {
		server := newRegisterOptionsTestServer(t)
		_, err := Register(server, workspace, config, WithTools(Selection{
			Exclude: []string{"generate_config", "generate_html", "validate_config", "validate_html"},
		}))
		require.NoError(t, err)

		client := connectRegisterTestClient(t, server)
		assert.Empty(t, listRegisterTestToolNames(t, client))
	})

	t.Run("filtered tool is absent from list and call fails", func(t *testing.T) {
		server := newRegisterOptionsTestServer(t)
		_, err := Register(server, workspace, config, WithTools(Selection{
			Include: []string{"generate_config"},
		}))
		require.NoError(t, err)

		client := connectRegisterTestClient(t, server)
		assert.Equal(t, []string{"generate_config"}, listRegisterTestToolNames(t, client))

		_, err = client.CallTool(context.Background(), &mcp.CallToolParams{
			Name: "validate_html",
		})
		require.Error(t, err, "a filtered tool must not be callable")

		result := callRegisterOptionsTestTool(t, client, "generate_config", map[string]any{"focus": "generate"})
		assert.False(t, result.IsError)
		require.NotEmpty(t, result.Content)
		text, ok := result.Content[0].(*mcp.TextContent)
		require.True(t, ok)
		assert.Contains(t, text.Text, "Current config")
	})
}

func TestRegisterResourceAndTemplateFilters(t *testing.T) {
	workspace := newRegisterOptionsTestWorkspace(t)
	config := ServerConfig{MaxDescriptionLength: 2000}

	t.Run("resource include and exclude", func(t *testing.T) {
		server := newRegisterOptionsTestServer(t)
		_, err := Register(server, workspace, config, WithResources(Selection{
			Include: []string{"cem://elements", "cem://packages"},
			Exclude: []string{"cem://elements"},
		}))
		require.NoError(t, err)

		client := connectRegisterTestClient(t, server)
		assert.Equal(t, []string{"cem://packages"}, listRegisterTestResourceURIs(t, client))
		assert.NotEmpty(t, listRegisterTestResourceTemplateURIs(t, client),
			"resource filters must not affect templates")
	})

	t.Run("template include selects exact URIs", func(t *testing.T) {
		server := newRegisterOptionsTestServer(t)
		_, err := Register(server, workspace, config, WithTemplates(Selection{
			Include: []string{"cem://element/{tagName}"},
		}))
		require.NoError(t, err)

		client := connectRegisterTestClient(t, server)
		assert.Equal(t, []string{"cem://element/{tagName}"}, listRegisterTestResourceTemplateURIs(t, client))
		assert.NotEmpty(t, listRegisterTestResourceURIs(t, client),
			"template filters must not affect static resources")
	})

	t.Run("template exclude wins over include", func(t *testing.T) {
		server := newRegisterOptionsTestServer(t)
		_, err := Register(server, workspace, config, WithTemplates(Selection{
			Include: []string{"cem://element/{tagName}", "cem://element/{tagName}/attributes"},
			Exclude: []string{"cem://element/{tagName}"},
		}))
		require.NoError(t, err)

		client := connectRegisterTestClient(t, server)
		assert.Equal(t,
			[]string{"cem://element/{tagName}/attributes"},
			listRegisterTestResourceTemplateURIs(t, client),
		)
	})

	t.Run("exclude everything from a category", func(t *testing.T) {
		server := newRegisterOptionsTestServer(t)
		_, err := Register(server, workspace, config,
			WithResources(Selection{Exclude: []string{"cem://elements"}}),
			WithTemplates(Selection{Exclude: []string{"cem://element/{tagName}"}}),
		)
		require.NoError(t, err)

		client := connectRegisterTestClient(t, server)
		resources := listRegisterTestResourceURIs(t, client)
		assert.NotContains(t, resources, "cem://elements")
		assert.Contains(t, resources, "cem://packages")
		templates := listRegisterTestResourceTemplateURIs(t, client)
		assert.NotContains(t, templates, "cem://element/{tagName}")
		assert.Contains(t, templates, "cem://element/{tagName}/attributes")
	})
}

func TestRegisterFilterValidation(t *testing.T) {
	workspace := newRegisterOptionsTestWorkspace(t)
	config := ServerConfig{MaxDescriptionLength: 2000}

	t.Run("unknown include entry names the entry", func(t *testing.T) {
		tests := []struct {
			name string
			opt  RegisterOption
			want string
		}{
			{"tool", WithTools(Selection{Include: []string{"missing_tool"}}), "missing_tool"},
			{"resource", WithResources(Selection{Include: []string{"cem://missing"}}), "cem://missing"},
			{"template", WithTemplates(Selection{Include: []string{"cem://missing/{x}"}}), "cem://missing/{x}"},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				server := newRegisterOptionsTestServer(t)
				_, err := Register(server, workspace, config, test.opt)
				require.Error(t, err)
				assert.Contains(t, err.Error(), test.want)

				client := connectRegisterTestClient(t, server)
				assert.Empty(t, listRegisterTestToolNames(t, client))
				assert.Empty(t, listRegisterTestResourceURIs(t, client))
				assert.Empty(t, listRegisterTestResourceTemplateURIs(t, client))
			})
		}
	})

	t.Run("matching is case-sensitive", func(t *testing.T) {
		server := newRegisterOptionsTestServer(t)
		_, err := Register(server, workspace, config, WithTools(Selection{
			Include: []string{"Generate_Config"},
		}))
		require.ErrorContains(t, err, "Generate_Config")
	})

	t.Run("unknown exclude entries are accepted", func(t *testing.T) {
		server := newRegisterOptionsTestServer(t)
		_, err := Register(server, workspace, config,
			WithTools(Selection{Exclude: []string{"missing_tool"}}),
			WithResources(Selection{Exclude: []string{"cem://missing"}}),
			WithTemplates(Selection{Exclude: []string{"cem://missing/{x}"}}),
		)
		require.NoError(t, err)

		client := connectRegisterTestClient(t, server)
		assert.Len(t, listRegisterTestToolNames(t, client), 4)
		assert.NotEmpty(t, listRegisterTestResourceURIs(t, client))
		assert.NotEmpty(t, listRegisterTestResourceTemplateURIs(t, client))
	})

	t.Run("empty strings are rejected", func(t *testing.T) {
		tests := []struct {
			name string
			opt  RegisterOption
		}{
			{"tool include", WithTools(Selection{Include: []string{""}})},
			{"tool exclude", WithTools(Selection{Exclude: []string{""}})},
			{"resource include", WithResources(Selection{Include: []string{""}})},
			{"resource exclude", WithResources(Selection{Exclude: []string{""}})},
			{"template include", WithTemplates(Selection{Include: []string{""}})},
			{"template exclude", WithTemplates(Selection{Exclude: []string{""}})},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				server := newRegisterOptionsTestServer(t)
				_, err := Register(server, workspace, config, test.opt)
				require.ErrorContains(t, err, "must not be empty")
			})
		}
	})

	t.Run("nil option is rejected without panic", func(t *testing.T) {
		server := newRegisterOptionsTestServer(t)
		require.NotPanics(t, func() {
			_, err := Register(server, workspace, config, nil)
			require.ErrorContains(t, err, "register option must not be nil")
		})
	})
}

func TestRegisterRepeatedFilterOptionsAccumulate(t *testing.T) {
	workspace := newRegisterOptionsTestWorkspace(t)
	config := ServerConfig{MaxDescriptionLength: 2000}

	t.Run("includes union across calls", func(t *testing.T) {
		server := newRegisterOptionsTestServer(t)
		_, err := Register(server, workspace, config,
			WithTools(Selection{Include: []string{"generate_config"}}),
			WithTools(Selection{Include: []string{"validate_config"}}),
		)
		require.NoError(t, err)

		client := connectRegisterTestClient(t, server)
		assert.Equal(t, []string{"generate_config", "validate_config"}, listRegisterTestToolNames(t, client))
	})

	t.Run("excludes union across calls and beat includes", func(t *testing.T) {
		server := newRegisterOptionsTestServer(t)
		_, err := Register(server, workspace, config,
			WithTools(Selection{Include: []string{"generate_config", "validate_config"}}),
			WithTools(Selection{Exclude: []string{"generate_config"}}),
			WithTools(Selection{Exclude: []string{"validate_config"}}),
		)
		require.NoError(t, err)

		client := connectRegisterTestClient(t, server)
		assert.Empty(t, listRegisterTestToolNames(t, client))
	})

	t.Run("resource includes union across calls", func(t *testing.T) {
		server := newRegisterOptionsTestServer(t)
		_, err := Register(server, workspace, config,
			WithResources(Selection{Include: []string{"cem://elements"}}),
			WithResources(Selection{Include: []string{"cem://packages"}}),
		)
		require.NoError(t, err)

		client := connectRegisterTestClient(t, server)
		assert.Equal(t, []string{"cem://elements", "cem://packages"}, listRegisterTestResourceURIs(t, client))
	})

	t.Run("template includes union across calls", func(t *testing.T) {
		server := newRegisterOptionsTestServer(t)
		_, err := Register(server, workspace, config,
			WithTemplates(Selection{Include: []string{"cem://element/{tagName}"}}),
			WithTemplates(Selection{Include: []string{"cem://element/{tagName}/attributes"}}),
		)
		require.NoError(t, err)

		client := connectRegisterTestClient(t, server)
		assert.Equal(t,
			[]string{"cem://element/{tagName}", "cem://element/{tagName}/attributes"},
			listRegisterTestResourceTemplateURIs(t, client),
		)
	})
}

func TestRegisterConflictPolicies(t *testing.T) {
	workspace := newRegisterOptionsTestWorkspace(t)
	config := ServerConfig{MaxDescriptionLength: 2000}

	t.Run("ConflictError is atomic and keeps the host handler", func(t *testing.T) {
		server := newRegisterOptionsTestServer(t)
		addHostTool(server, "generate_config", "host handler")

		_, err := Register(server, workspace, config, WithConflictPolicy(ConflictError))
		require.ErrorContains(t, err, "generate_config")

		client := connectRegisterTestClient(t, server)
		assert.Equal(t, []string{"generate_config"}, listRegisterTestToolNames(t, client))
		assert.Empty(t, listRegisterTestResourceURIs(t, client),
			"a failed registration must not add cem resources")
		assert.Empty(t, listRegisterTestResourceTemplateURIs(t, client),
			"a failed registration must not add cem resource templates")

		result := callRegisterOptionsTestTool(t, client, "generate_config", nil)
		require.NotEmpty(t, result.Content)
		text, ok := result.Content[0].(*mcp.TextContent)
		require.True(t, ok)
		assert.Equal(t, "host handler", text.Text)
	})

	t.Run("ConflictError detects resource and template conflicts", func(t *testing.T) {
		server := newRegisterOptionsTestServer(t)
		addHostResource(server, "cem://elements", "host resource")
		addHostResourceTemplate(server, "cem://element/{tagName}", "host template")

		_, err := Register(server, workspace, config, WithConflictPolicy(ConflictError))
		require.ErrorContains(t, err, `resource "cem://elements"`)
		require.ErrorContains(t, err, `resource template "cem://element/{tagName}"`)

		client := connectRegisterTestClient(t, server)
		assert.Empty(t, listRegisterTestToolNames(t, client),
			"a failed registration must not add cem tools")
		assert.Equal(t, []string{"cem://elements"}, listRegisterTestResourceURIs(t, client))
		assert.Equal(t, []string{"cem://element/{tagName}"}, listRegisterTestResourceTemplateURIs(t, client))
		assert.Equal(t, "host resource", readRegisterOptionsTestResource(t, client, "cem://elements"))
		assert.Equal(t, "host template", readRegisterOptionsTestResource(t, client, "cem://element/test-element"))
	})

	t.Run("ConflictError exposes conflicts programmatically", func(t *testing.T) {
		server := newRegisterOptionsTestServer(t)
		addHostTool(server, "generate_config", "host handler")
		addHostResource(server, "cem://elements", "host resource")
		addHostResourceTemplate(server, "cem://element/{tagName}", "host template")

		_, err := Register(server, workspace, config, WithConflictPolicy(ConflictError))

		var conflict *RegistrationConflictError
		require.ErrorAs(t, err, &conflict)
		assert.Equal(t, []string{"generate_config"}, conflict.Tools)
		assert.Equal(t, []string{"cem://elements"}, conflict.Resources)
		assert.Equal(t, []string{"cem://element/{tagName}"}, conflict.ResourceTemplates)
	})

	t.Run("ConflictError succeeds on an empty surface", func(t *testing.T) {
		server := newRegisterOptionsTestServer(t)
		_, err := Register(server, workspace, config, WithConflictPolicy(ConflictError))
		require.NoError(t, err)

		client := connectRegisterTestClient(t, server)
		assert.Len(t, listRegisterTestToolNames(t, client), 4)
		assert.NotEmpty(t, listRegisterTestResourceURIs(t, client))
		assert.NotEmpty(t, listRegisterTestResourceTemplateURIs(t, client))
	})

	t.Run("ConflictSkip keeps the host handler and registers the rest", func(t *testing.T) {
		server := newRegisterOptionsTestServer(t)
		addHostTool(server, "generate_config", "host handler")

		_, err := Register(server, workspace, config, WithConflictPolicy(ConflictSkip))
		require.NoError(t, err)

		client := connectRegisterTestClient(t, server)
		assert.Equal(t,
			[]string{"generate_config", "generate_html", "validate_config", "validate_html"},
			listRegisterTestToolNames(t, client),
		)

		result := callRegisterOptionsTestTool(t, client, "generate_config", nil)
		require.NotEmpty(t, result.Content)
		text, ok := result.Content[0].(*mcp.TextContent)
		require.True(t, ok)
		assert.Equal(t, "host handler", text.Text)

		result = callRegisterOptionsTestTool(t, client, "validate_config", nil)
		assert.False(t, result.IsError, "non-conflicting cem tools should work")
	})

	t.Run("ConflictSkip keeps existing resources and templates", func(t *testing.T) {
		server := newRegisterOptionsTestServer(t)
		addHostResource(server, "cem://elements", "host resource")
		addHostResourceTemplate(server, "cem://element/{tagName}", "host template")

		_, err := Register(server, workspace, config, WithConflictPolicy(ConflictSkip))
		require.NoError(t, err)

		client := connectRegisterTestClient(t, server)
		assert.Len(t, listRegisterTestToolNames(t, client), 4,
			"tool conflicts must not be affected by resource conflicts")

		resources := listRegisterTestResourceURIs(t, client)
		assert.Contains(t, resources, "cem://elements")
		assert.Contains(t, resources, "cem://packages")
		templates := listRegisterTestResourceTemplateURIs(t, client)
		assert.Contains(t, templates, "cem://element/{tagName}")
		assert.Contains(t, templates, "cem://element/{tagName}/attributes")

		assert.Equal(t, "host resource", readRegisterOptionsTestResource(t, client, "cem://elements"))
		assert.Equal(t, "host template", readRegisterOptionsTestResource(t, client, "cem://element/test-element"))
	})

	t.Run("ConflictReplace replaces the host handler", func(t *testing.T) {
		server := newRegisterOptionsTestServer(t)
		addHostTool(server, "generate_config", "host handler")

		_, err := Register(server, workspace, config, WithConflictPolicy(ConflictReplace))
		require.NoError(t, err)

		client := connectRegisterTestClient(t, server)
		result := callRegisterOptionsTestTool(t, client, "generate_config", map[string]any{"focus": "generate"})
		require.NotEmpty(t, result.Content)
		text, ok := result.Content[0].(*mcp.TextContent)
		require.True(t, ok)
		assert.Contains(t, text.Text, "Current config")
		assert.NotEqual(t, "host handler", text.Text)
	})

	t.Run("filtered-out conflicts do not fail ConflictError", func(t *testing.T) {
		server := newRegisterOptionsTestServer(t)
		addHostTool(server, "generate_config", "host handler")

		_, err := Register(server, workspace, config,
			WithTools(Selection{Include: []string{"validate_html"}}),
			WithConflictPolicy(ConflictError),
		)
		require.NoError(t, err)

		client := connectRegisterTestClient(t, server)
		assert.Equal(t, []string{"generate_config", "validate_html"}, listRegisterTestToolNames(t, client))
	})

	t.Run("default options replace on conflict", func(t *testing.T) {
		server := newRegisterOptionsTestServer(t)
		addHostTool(server, "generate_config", "host handler")

		_, err := Register(server, workspace, config)
		require.NoError(t, err)

		client := connectRegisterTestClient(t, server)
		result := callRegisterOptionsTestTool(t, client, "generate_config", map[string]any{"focus": "generate"})
		require.NotEmpty(t, result.Content)
		text, ok := result.Content[0].(*mcp.TextContent)
		require.True(t, ok)
		assert.Contains(t, text.Text, "Current config")
	})

	t.Run("last conflict policy wins", func(t *testing.T) {
		server := newRegisterOptionsTestServer(t)
		addHostTool(server, "generate_config", "host handler")

		_, err := Register(server, workspace, config,
			WithConflictPolicy(ConflictError),
			WithConflictPolicy(ConflictReplace),
		)
		require.NoError(t, err)

		client := connectRegisterTestClient(t, server)
		result := callRegisterOptionsTestTool(t, client, "generate_config", map[string]any{"focus": "generate"})
		require.NotEmpty(t, result.Content)
		text, ok := result.Content[0].(*mcp.TextContent)
		require.True(t, ok)
		assert.Contains(t, text.Text, "Current config")
	})

	t.Run("out-of-range policy is rejected before registration", func(t *testing.T) {
		server := newRegisterOptionsTestServer(t)
		_, err := Register(server, workspace, config, WithConflictPolicy(ConflictPolicy(99)))
		require.ErrorContains(t, err, "unknown conflict policy")

		client := connectRegisterTestClient(t, server)
		assert.Empty(t, listRegisterTestToolNames(t, client))
		assert.Empty(t, listRegisterTestResourceURIs(t, client))
		assert.Empty(t, listRegisterTestResourceTemplateURIs(t, client))
	})
}

func TestRegisterDoesNotAllocateFileWatcher(t *testing.T) {
	workspace := newRegisterOptionsTestWorkspace(t)
	config := ServerConfig{MaxDescriptionLength: 2000}

	server := newRegisterOptionsTestServer(t)
	handle, err := Register(server, workspace, config)
	require.NoError(t, err)
	require.NotNil(t, handle)
	assert.Nil(t, handle.registry.current.Load().LSPRegistry().GetFileWatcher(),
		"Register must not allocate a file watcher")

	require.NoError(t, handle.Reload(context.Background()))
	assert.Nil(t, handle.registry.current.Load().LSPRegistry().GetFileWatcher(),
		"Reload must not allocate a file watcher")

	cemServer, err := NewServerWithConfig(workspace, config)
	require.NoError(t, err)
	t.Cleanup(func() { _ = cemServer.Registry().LSPRegistry().StopFileWatching() })
	assert.NotNil(t, cemServer.Registry().LSPRegistry().GetFileWatcher(),
		"NewServerWithConfig must keep its watch lifecycle")
}

func TestServerReloadFromWatcherRefreshesManifest(t *testing.T) {
	workspace, manifestPath := newReloadTestWorkspace(t, reloadTestManifest("original description"))

	server, err := NewServerWithConfig(workspace, ServerConfig{})
	require.NoError(t, err)
	t.Cleanup(func() { _ = server.Registry().LSPRegistry().StopFileWatching() })

	require.NoError(t, server.Registry().LoadManifests())

	client := connectRegisterTestClient(t, server.server)
	assert.Contains(t, readRegisterOptionsTestResource(t, client, "cem://element/reload-element"), "original description")

	writeReloadTestFile(t, manifestPath, reloadTestManifest("updated description"))
	server.reloadFromWatcher()

	assert.Contains(t, readRegisterOptionsTestResource(t, client, "cem://element/reload-element"), "updated description",
		"the watcher path must invalidate the manifest cache before reloading")
}

func TestRegistrationReloadReadsUpdatedManifest(t *testing.T) {
	workspace, manifestPath := newReloadTestWorkspace(t, reloadTestManifest("original description"))

	server := newRegisterOptionsTestServer(t)
	handle, err := Register(server, workspace, ServerConfig{})
	require.NoError(t, err)

	client := connectRegisterTestClient(t, server)
	assert.Contains(t, readRegisterOptionsTestResource(t, client, "cem://element/reload-element"), "original description")

	writeReloadTestFile(t, manifestPath, reloadTestManifest("updated description"))
	require.NoError(t, handle.Reload(context.Background()))

	assert.Contains(t, readRegisterOptionsTestResource(t, client, "cem://element/reload-element"), "updated description")
	assert.NotContains(t, readRegisterOptionsTestResource(t, client, "cem://element/reload-element"), "original description")
}

func TestRegistrationReloadMalformedManifestKeepsPreviousState(t *testing.T) {
	workspace, manifestPath := newReloadTestWorkspace(t, reloadTestManifest("original description"))

	server := newRegisterOptionsTestServer(t)
	handle, err := Register(server, workspace, ServerConfig{})
	require.NoError(t, err)

	client := connectRegisterTestClient(t, server)
	assert.Contains(t, readRegisterOptionsTestResource(t, client, "cem://element/reload-element"), "original description")

	writeReloadTestFile(t, manifestPath, `{"schemaVersion": "2.1.1", "modules": [`)

	err = handle.Reload(context.Background())
	require.Error(t, err, "a malformed manifest must fail reload")

	content := readRegisterOptionsTestResource(t, client, "cem://element/reload-element")
	assert.Contains(t, content, "original description",
		"the previous registry state must stay readable after a failed reload")

	writeReloadTestFile(t, manifestPath, reloadTestManifest("recovered description"))
	require.NoError(t, handle.Reload(context.Background()))
	assert.Contains(t, readRegisterOptionsTestResource(t, client, "cem://element/reload-element"), "recovered description",
		"a failed reload must not poison later reloads")
}

func TestRegistrationReloadConcurrentWithResourceReads(t *testing.T) {
	workspace, manifestPath := newReloadTestWorkspace(t, reloadTestManifest("version-0"))

	server := newRegisterOptionsTestServer(t)
	handle, err := Register(server, workspace, ServerConfig{})
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	const readers = 4
	readersDone := make(chan struct{})
	startReaders := make(chan struct{})
	errCh := make(chan error, readers+1)

	clients := make([]*mcp.ClientSession, readers)
	for i := range clients {
		clients[i] = connectRegisterTestClient(t, server)
	}

	var readerWG sync.WaitGroup
	for i := 0; i < readers; i++ {
		readerWG.Add(1)
		go func(client *mcp.ClientSession) {
			defer readerWG.Done()

			<-startReaders
			for {
				select {
				case <-readersDone:
					return
				default:
					result, err := client.ReadResource(ctx, &mcp.ReadResourceParams{
						URI: "cem://element/reload-element",
					})
					if err != nil {
						errCh <- fmt.Errorf("concurrent read: %w", err)
						return
					}
					if len(result.Contents) == 0 || result.Contents[0].Text == "" {
						errCh <- fmt.Errorf("concurrent read returned empty contents")
						return
					}
				}
			}
		}(clients[i])
	}

	close(startReaders)

	for i := 0; i < 10; i++ {
		writeReloadTestFile(t, manifestPath, reloadTestManifest(fmt.Sprintf("version-%d", i+1)))
		if err := handle.Reload(ctx); err != nil {
			errCh <- fmt.Errorf("concurrent reload: %w", err)
			break
		}
	}

	close(readersDone)
	readerWG.Wait()
	close(errCh)

	for err := range errCh {
		require.NoError(t, err)
	}

	assert.Contains(t, readRegisterOptionsTestResource(t, connectRegisterTestClient(t, server), "cem://element/reload-element"),
		"version-10")
}

func TestRegistrationReloadZeroValueReturnsError(t *testing.T) {
	var handle *Registration
	require.NotPanics(t, func() {
		require.Error(t, handle.Reload(context.Background()))
	})

	handle = &Registration{}
	require.NotPanics(t, func() {
		require.ErrorContains(t, handle.Reload(context.Background()), "reloadable")
	})
}

func TestRegisterConflictPreflightFiresInitializedHandler(t *testing.T) {
	workspace := newRegisterOptionsTestWorkspace(t)
	config := ServerConfig{MaxDescriptionLength: 2000}

	var defaultCalls int
	defaultServer := mcp.NewServer(&mcp.Implementation{
		Name:    "host-server",
		Version: "test",
	}, &mcp.ServerOptions{
		InitializedHandler: func(context.Context, *mcp.InitializedRequest) {
			defaultCalls++
		},
	})
	_, err := Register(defaultServer, workspace, config)
	require.NoError(t, err)
	assert.Zero(t, defaultCalls, "ConflictReplace must not open a preflight session")

	var conflictCalls int
	conflictServer := mcp.NewServer(&mcp.Implementation{
		Name:    "host-server",
		Version: "test",
	}, &mcp.ServerOptions{
		InitializedHandler: func(context.Context, *mcp.InitializedRequest) {
			conflictCalls++
		},
	})
	_, err = Register(conflictServer, workspace, config, WithConflictPolicy(ConflictError))
	require.NoError(t, err)
	assert.Equal(t, 1, conflictCalls, "preflight opens and initializes one in-memory session")
}

func TestRegisterSecondCallIsAdditive(t *testing.T) {
	workspace := newRegisterOptionsTestWorkspace(t)
	config := ServerConfig{MaxDescriptionLength: 2000}

	server := newRegisterOptionsTestServer(t)
	first, err := Register(server, workspace, config, WithTools(Selection{
		Include: []string{"generate_config"},
	}))
	require.NoError(t, err)
	second, err := Register(server, workspace, config, WithTools(Selection{
		Include: []string{"validate_config"},
	}))
	require.NoError(t, err)
	require.NotSame(t, first, second)

	client := connectRegisterTestClient(t, server)
	assert.Equal(t, []string{"generate_config", "validate_config"}, listRegisterTestToolNames(t, client))

	result := callRegisterOptionsTestTool(t, client, "generate_config", map[string]any{"focus": "generate"})
	assert.False(t, result.IsError)
	result = callRegisterOptionsTestTool(t, client, "validate_config", nil)
	assert.False(t, result.IsError)
}

func TestRegisterValidatesOptionsBeforeWorkspaceUse(t *testing.T) {
	workspace := rootPanicWorkspace{}
	server := newRegisterOptionsTestServer(t)

	require.NotPanics(t, func() {
		_, err := Register(server, workspace, ServerConfig{}, WithTools(Selection{Include: []string{""}}))
		require.ErrorContains(t, err, "must not be empty")
	})

	require.NotPanics(t, func() {
		_, err := Register(server, workspace, ServerConfig{}, WithConflictPolicy(ConflictPolicy(99)))
		require.ErrorContains(t, err, "unknown conflict policy")
	})
}

func TestRegistrationReloadContextErrors(t *testing.T) {
	workspace, _ := newReloadTestWorkspace(t, reloadTestManifest("original description"))

	server := newRegisterOptionsTestServer(t)
	handle, err := Register(server, workspace, ServerConfig{})
	require.NoError(t, err)

	require.NotPanics(t, func() {
		//nolint:staticcheck // deliberate nil context guard test
		//lint:ignore SA1012 the nil context guard is the behavior under test
		require.ErrorContains(t, handle.Reload(nil), "context")
	})

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	require.ErrorIs(t, handle.Reload(ctx), context.Canceled)

	client := connectRegisterTestClient(t, server)
	assert.Contains(t, readRegisterOptionsTestResource(t, client, "cem://element/reload-element"), "original description",
		"rejected reloads must not change the served state")
}

func TestRegistrationReloadCanceledAfterLoadDoesNotSwap(t *testing.T) {
	base, _ := newReloadTestWorkspace(t, reloadTestManifest("original description"))

	ctx, cancel := context.WithCancel(context.Background())
	workspace := &cancelOnManifestWorkspace{
		WorkspaceContext: base,
		cancel:           cancel,
	}

	server := newRegisterOptionsTestServer(t)
	handle, err := Register(server, workspace, ServerConfig{})
	require.NoError(t, err)

	before := handle.registry.current.Load()
	workspace.armed.Store(true)

	err = handle.Reload(ctx)
	require.ErrorIs(t, err, context.Canceled)
	assert.Same(t, before, handle.registry.current.Load(),
		"a canceled reload must not swap the registry after loading")

	client := connectRegisterTestClient(t, server)
	assert.Contains(t, readRegisterOptionsTestResource(t, client, "cem://element/reload-element"), "original description")
}

func TestValidateReloadManifestStatErrors(t *testing.T) {
	workspace, _ := newReloadTestWorkspace(t, reloadTestManifest("original description"))

	require.NoError(t, validateReloadManifest(statErrorWorkspace{
		WorkspaceContext: workspace,
		err:              fs.ErrNotExist,
	}), "a missing manifest is handled by the in-memory generation fallback")

	require.ErrorIs(t, validateReloadManifest(statErrorWorkspace{
		WorkspaceContext: workspace,
		err:              fs.ErrPermission,
	}), fs.ErrPermission)

	require.NoError(t, validateReloadManifest(statErrorWorkspace{
		WorkspaceContext: workspace,
		err:              cemworkspace.ErrRemoteUnsupported,
	}), "unsupported local inspection is handled by the in-memory generation fallback")

	require.NoError(t, validateReloadManifest(emptyManifestPathWorkspace{WorkspaceContext: workspace}),
		"a workspace without a manifest path is not a reload failure")
}

func TestRegistrationReloadWithoutCacheInvalidatorUsesCachedWorkspaceData(t *testing.T) {
	base, manifestPath := newReloadTestWorkspace(t, reloadTestManifest("original description"))
	workspace := nonInvalidatingWorkspace{WorkspaceContext: base}

	server := newRegisterOptionsTestServer(t)
	handle, err := Register(server, workspace, ServerConfig{})
	require.NoError(t, err)

	client := connectRegisterTestClient(t, server)
	assert.Contains(t, readRegisterOptionsTestResource(t, client, "cem://element/reload-element"), "original description")
	before := handle.registry.current.Load()

	writeReloadTestFile(t, manifestPath, reloadTestManifest("updated description"))
	require.NoError(t, handle.Reload(context.Background()))

	after := handle.registry.current.Load()
	require.NotSame(t, before, after, "Reload must rebuild the registry even without types.CacheInvalidator")
	assert.NotEmpty(t, after.GetAllElements())

	assert.Contains(t, readRegisterOptionsTestResource(t, client, "cem://element/reload-element"), "original description",
		"without types.CacheInvalidator, Reload rebuilds from the workspace's cached data")
	assert.NotContains(t, readRegisterOptionsTestResource(t, client, "cem://element/reload-element"), "updated description")
}

func TestRegistrationReloadFailsWhenAdditionalPackageFails(t *testing.T) {
	workspace, _ := newReloadTestWorkspace(t, reloadTestManifest("original description"))

	server := newRegisterOptionsTestServer(t)
	handle, err := Register(server, workspace, ServerConfig{
		AdditionalPackages: []string{"./testdata/fixtures/does-not-exist"},
	})
	require.NoError(t, err, "Register logs additional package failures")

	client := connectRegisterTestClient(t, server)
	assert.Contains(t, readRegisterOptionsTestResource(t, client, "cem://element/reload-element"), "original description")

	err = handle.Reload(context.Background())
	require.ErrorContains(t, err, "additional package")
	require.ErrorContains(t, err, "does-not-exist")

	assert.Contains(t, readRegisterOptionsTestResource(t, client, "cem://element/reload-element"), "original description",
		"a failed reload must leave the previous state active")
}

func TestRegistrationReloadKeepsConfigSnapshotOnFailure(t *testing.T) {
	workspace, _, configPath := newReloadTestWorkspaceWithConfig(t,
		reloadTestManifest("original description"),
		"serve:\n  port: 5000\n",
	)

	server := newRegisterOptionsTestServer(t)
	handle, err := Register(server, workspace, ServerConfig{})
	require.NoError(t, err)

	active := handle.registry.current.Load()
	cfg, err := active.Config()
	require.NoError(t, err)
	require.Equal(t, 5000, cfg.Serve.Port)

	writeReloadTestFile(t, configPath, "serve: [\n")

	require.Error(t, handle.Reload(context.Background()))

	assert.Same(t, active, handle.registry.current.Load(),
		"a failed reload must keep the previous generation")
	cfg, err = active.Config()
	require.NoError(t, err)
	assert.Equal(t, 5000, cfg.Serve.Port,
		"the previous generation must keep its configuration snapshot")

	writeReloadTestFile(t, configPath, "serve:\n  port: 6000\n")
	require.NoError(t, handle.Reload(context.Background()))
	next := handle.registry.current.Load()
	require.NotSame(t, active, next)
	cfg, err = next.Config()
	require.NoError(t, err)
	assert.Equal(t, 6000, cfg.Serve.Port,
		"a successful reload must serve the updated configuration")
}

func TestRegistrationReloadReusesDocumentManager(t *testing.T) {
	workspace, _ := newReloadTestWorkspace(t, reloadTestManifest("original description"))

	server := newRegisterOptionsTestServer(t)
	handle, err := Register(server, workspace, ServerConfig{})
	require.NoError(t, err)

	before := handle.registry.current.Load().DocumentManager()
	for range 5 {
		require.NoError(t, handle.Reload(context.Background()))
		after := handle.registry.current.Load().DocumentManager()
		assert.Same(t, before, after, "every reload generation must reuse the tree-sitter document manager")
	}
}

func TestRegistrationConcurrentReloads(t *testing.T) {
	workspace, _ := newReloadTestWorkspace(t, reloadTestManifest("stable description"))

	server := newRegisterOptionsTestServer(t)
	handle, err := Register(server, workspace, ServerConfig{})
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	const reloaders = 4
	const reloadsPerGoroutine = 5
	errCh := make(chan error, reloaders)

	var wg sync.WaitGroup
	for range reloaders {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range reloadsPerGoroutine {
				if err := handle.Reload(ctx); err != nil {
					errCh <- err
					return
				}
			}
		}()
	}

	wg.Wait()
	close(errCh)
	for err := range errCh {
		require.NoError(t, err)
	}

	client := connectRegisterTestClient(t, server)
	assert.Contains(t, readRegisterOptionsTestResource(t, client, "cem://element/reload-element"), "stable description")
}

func newReloadTestWorkspace(t *testing.T, manifest string) (cemtypes.WorkspaceContext, string) {
	t.Helper()

	workspace, manifestPath, _ := newReloadTestWorkspaceWithConfig(t, manifest, "")
	return workspace, manifestPath
}

func newReloadTestWorkspaceWithConfig(t *testing.T, manifest, config string) (cemtypes.WorkspaceContext, string, string) {
	t.Helper()

	dir := t.TempDir()
	writeReloadTestFile(t, filepath.Join(dir, "package.json"),
		`{"name": "reload-package", "version": "1.0.0", "customElements": "custom-elements.json"}`)

	manifestPath := filepath.Join(dir, "custom-elements.json")
	writeReloadTestFile(t, manifestPath, manifest)

	configPath := filepath.Join(dir, ".config", "cem.yaml")
	if config != "" {
		require.NoError(t, platform.NewOSFileSystem().MkdirAll(filepath.Dir(configPath), 0o755))
		writeReloadTestFile(t, configPath, config)
	}

	workspace, err := fileworkspace.NewFileSystemWorkspaceContext(dir)
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, workspace.Cleanup()) })
	return workspace, manifestPath, configPath
}

// statErrorWorkspace overrides Stat so tests can exercise validation errors.
type statErrorWorkspace struct {
	cemtypes.WorkspaceContext
	err error
}

func (w statErrorWorkspace) Stat(string) (fs.FileInfo, error) {
	return nil, w.err
}

// nonInvalidatingWorkspace hides the base workspace's InvalidateConfig and
// InvalidateManifest methods so it models an embedder that caches data but
// does not implement types.CacheInvalidator.
type nonInvalidatingWorkspace struct {
	cemtypes.WorkspaceContext
}

// cancelOnManifestWorkspace cancels a reload context when Manifest is read
// after the test arms it, so the post-load cancellation check is exercised.
type cancelOnManifestWorkspace struct {
	cemtypes.WorkspaceContext
	cancel context.CancelFunc
	armed  atomic.Bool
}

func (w *cancelOnManifestWorkspace) Manifest() (*M.Package, error) {
	if w.armed.Load() {
		w.cancel()
	}
	return w.WorkspaceContext.Manifest()
}

// emptyManifestPathWorkspace models a workspace that does not expose a
// manifest path, for example one that relies entirely on generation.
type emptyManifestPathWorkspace struct {
	cemtypes.WorkspaceContext
}

func (w emptyManifestPathWorkspace) CustomElementsManifestPath() string {
	return ""
}

// rootPanicWorkspace proves option validation does not touch the workspace.
type rootPanicWorkspace struct {
	cemtypes.WorkspaceContext
}

func (w rootPanicWorkspace) Root() string {
	panic("workspace Root must not be called during option validation")
}

func writeReloadTestFile(t *testing.T, path, content string) {
	t.Helper()
	require.NoError(t, platform.NewOSFileSystem().WriteFile(path, []byte(content), 0o644))
}

func reloadTestManifest(description string) string {
	return fmt.Sprintf(`{
  "schemaVersion": "2.1.1",
  "modules": [
    {
      "kind": "javascript-module",
      "path": "reload-element.js",
      "declarations": [
        {
          "kind": "class",
          "name": "ReloadElement",
          "tagName": "reload-element",
          "customElement": true,
          "description": %q,
          "attributes": [
            {
              "name": "label",
              "type": {"text": "string"},
              "description": "The label"
            }
          ]
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "reload-element",
          "declaration": {
            "name": "ReloadElement",
            "module": "reload-element.js"
          }
        }
      ]
    }
  ]
}`, description)
}
