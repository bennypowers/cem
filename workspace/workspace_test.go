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

package workspace_test

import (
	"context"
	"io"
	"io/fs"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/mcp"
	cemworkspace "bennypowers.dev/cem/workspace"
	sdk "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The exported constructor is filesystem-backed, so these tests exercise a
// real fixture directory rather than a MapFS. The fixture is shared with the
// mcp package's Register tests to prove that external embedders can mount
// Register with the public constructor alone. Assertions are inline rather
// than golden files because the contract is a small, stable set of values from
// one shared fixture; a golden would add indirection without extra coverage.
const fixtureDir = "../mcp/testdata/fixtures/basic-integration"

func TestNewFileSystemWorkspaceContext(t *testing.T) {
	ctx, err := cemworkspace.NewFileSystemWorkspaceContext(fixtureDir)
	require.NoError(t, err)
	require.NotNil(t, ctx)
	t.Cleanup(func() { require.NoError(t, ctx.Cleanup()) })

	assert.Equal(t, fixtureDir, ctx.Root())
	assert.Equal(t, "", ctx.ConfigFile())
	assert.Equal(t, filepath.Join(fixtureDir, "custom-elements.json"), ctx.CustomElementsManifestPath())

	t.Run("Config", func(t *testing.T) {
		cfg, err := ctx.Config()
		require.NoError(t, err)
		require.NotNil(t, cfg)
		assert.Equal(t, fixtureDir, cfg.ProjectDir)
	})

	t.Run("PackageJSON", func(t *testing.T) {
		pkg, err := ctx.PackageJSON()
		require.NoError(t, err)
		require.NotNil(t, pkg)
		assert.Equal(t, "test-package", pkg.Name)
		assert.Equal(t, "1.0.0", pkg.Version)
		assert.Equal(t, "custom-elements.json", pkg.CustomElements)
	})

	t.Run("Manifest", func(t *testing.T) {
		manifest, err := ctx.Manifest()
		require.NoError(t, err)
		require.NotNil(t, manifest)
		assert.Equal(t, "2.1.1", manifest.SchemaVersion)
		require.NotEmpty(t, manifest.Modules)
	})

	t.Run("ReadFile", func(t *testing.T) {
		rc, err := ctx.ReadFile(filepath.Join(fixtureDir, "package.json"))
		require.NoError(t, err)
		defer func() { require.NoError(t, rc.Close()) }()

		content, err := io.ReadAll(rc)
		require.NoError(t, err)
		assert.Contains(t, string(content), `"test-package"`)
	})

	t.Run("Stat", func(t *testing.T) {
		info, err := ctx.Stat(filepath.Join(fixtureDir, "custom-elements.json"))
		require.NoError(t, err)
		assert.False(t, info.IsDir())
		assert.Positive(t, info.Size())
	})

	t.Run("ReadDir", func(t *testing.T) {
		entries, err := ctx.ReadDir(fixtureDir)
		require.NoError(t, err)

		names := make([]string, 0, len(entries))
		for _, entry := range entries {
			names = append(names, entry.Name())
		}
		assert.Contains(t, names, "package.json")
		assert.Contains(t, names, "custom-elements.json")
	})
}

func TestNewFileSystemWorkspaceContextMissingDirectory(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "does-not-exist")

	ctx, err := cemworkspace.NewFileSystemWorkspaceContext(dir)
	require.Error(t, err)
	assert.Nil(t, ctx)
	assert.ErrorIs(t, err, fs.ErrNotExist)
}

func TestNewFileSystemWorkspaceContextRejectsNonDirectory(t *testing.T) {
	file := filepath.Join(fixtureDir, "package.json")

	ctx, err := cemworkspace.NewFileSystemWorkspaceContext(file)
	require.Error(t, err)
	assert.Nil(t, ctx)
	assert.ErrorIs(t, err, fs.ErrInvalid)
	assert.ErrorContains(t, err, "not a directory")
}

func TestNewFileSystemWorkspaceContextInvalidConfig(t *testing.T) {
	dir := "testdata/invalid-config"

	ctx, err := cemworkspace.NewFileSystemWorkspaceContext(dir)
	require.Error(t, err)
	assert.Nil(t, ctx)
	assert.ErrorContains(t, err, "initializing")
	assert.ErrorContains(t, err, dir)
}

func TestNewFileSystemWorkspaceContextMissingManifest(t *testing.T) {
	ctx, err := cemworkspace.NewFileSystemWorkspaceContext("testdata/missing-manifest")
	require.NoError(t, err)
	require.NotNil(t, ctx)
	t.Cleanup(func() { require.NoError(t, ctx.Cleanup()) })

	manifest, err := ctx.Manifest()
	require.Error(t, err)
	assert.Nil(t, manifest)
	assert.ErrorIs(t, err, fs.ErrNotExist)
}

func TestNewFileSystemWorkspaceContextDiscoversConfig(t *testing.T) {
	dir := "testdata/config"

	ctx, err := cemworkspace.NewFileSystemWorkspaceContext(dir)
	require.NoError(t, err)
	require.NotNil(t, ctx)
	t.Cleanup(func() { require.NoError(t, ctx.Cleanup()) })

	assert.Equal(t, filepath.Join(dir, ".config", "cem.yaml"), ctx.ConfigFile())

	cfg, err := ctx.Config()
	require.NoError(t, err)
	require.NotNil(t, cfg)
	assert.Equal(t, filepath.Join(dir, "dist", "custom-elements.json"), cfg.Generate.Output)
	assert.Equal(t, cfg.Generate.Output, ctx.CustomElementsManifestPath())
}

// TestNewFileSystemWorkspaceContextSupportsRegister proves that the public
// constructor is sufficient for an external embedder: no internal package and
// no internal test helper is referenced here. The constructor initializes the
// workspace, so Register can be called immediately.
func TestNewFileSystemWorkspaceContextSupportsRegister(t *testing.T) {
	ctx, err := cemworkspace.NewFileSystemWorkspaceContext(fixtureDir)
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, ctx.Cleanup()) })

	hostServer := sdk.NewServer(&sdk.Implementation{
		Name:    "external-embedder",
		Version: "test",
	}, nil)
	require.NoError(t, mcp.Register(hostServer, ctx, mcp.ServerConfig{MaxDescriptionLength: 2000}))

	client := connectInMemoryClient(t, hostServer)

	t.Run("tools are listed", func(t *testing.T) {
		result, err := client.ListTools(context.Background(), &sdk.ListToolsParams{})
		require.NoError(t, err)
		require.NotEmpty(t, result.Tools)

		names := make([]string, 0, len(result.Tools))
		for _, tool := range result.Tools {
			names = append(names, tool.Name)
		}
		assert.Contains(t, names, "generate_config")
	})

	t.Run("resources are listed", func(t *testing.T) {
		result, err := client.ListResources(context.Background(), &sdk.ListResourcesParams{})
		require.NoError(t, err)
		require.NotEmpty(t, result.Resources)

		uris := make([]string, 0, len(result.Resources))
		for _, resource := range result.Resources {
			uris = append(uris, resource.URI)
		}
		assert.Contains(t, uris, "cem://elements")
	})

	t.Run("resource templates are listed", func(t *testing.T) {
		result, err := client.ListResourceTemplates(context.Background(), &sdk.ListResourceTemplatesParams{})
		require.NoError(t, err)
		require.NotEmpty(t, result.ResourceTemplates)

		uris := make([]string, 0, len(result.ResourceTemplates))
		for _, template := range result.ResourceTemplates {
			uris = append(uris, template.URITemplate)
		}
		assert.Contains(t, uris, "cem://element/{tagName}")
	})

	t.Run("element resource is readable", func(t *testing.T) {
		result, err := client.ReadResource(context.Background(), &sdk.ReadResourceParams{
			URI: "cem://element/test-element",
		})
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotEmpty(t, result.Contents)
		assert.Equal(t, "cem://element/test-element", result.Contents[0].URI)
		assert.Contains(t, result.Contents[0].Text, "test-element")
	})
}

func connectInMemoryClient(t *testing.T, server *sdk.Server) *sdk.ClientSession {
	t.Helper()

	ctx := context.Background()
	serverTransport, clientTransport := sdk.NewInMemoryTransports()

	serverSession, err := server.Connect(ctx, serverTransport, nil)
	require.NoError(t, err)
	t.Cleanup(func() { _ = serverSession.Close() })

	client := sdk.NewClient(&sdk.Implementation{
		Name:    "external-embedder-client",
		Version: "1.0.0",
	}, nil)

	clientSession, err := client.Connect(ctx, clientTransport, nil)
	require.NoError(t, err)
	t.Cleanup(func() { _ = clientSession.Close() })

	return clientSession
}
