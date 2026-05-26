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
package workspace_test

import (
	"io"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMapWorkspaceContext_Init(t *testing.T) {
	ctx := workspace.NewMapWorkspaceContext(t, "testdata/workspace-context")
	err := ctx.Init()
	require.NoError(t, err)
}

func TestMapWorkspaceContext_Root(t *testing.T) {
	ctx := workspace.NewMapWorkspaceContext(t, "testdata/workspace-context")
	err := ctx.Init()
	require.NoError(t, err)

	assert.Equal(t, "/", ctx.Root())
}

func TestMapWorkspaceContext_PackageJSON(t *testing.T) {
	ctx := workspace.NewMapWorkspaceContext(t, "testdata/workspace-context")
	err := ctx.Init()
	require.NoError(t, err)

	pkg, err := ctx.PackageJSON()
	require.NoError(t, err)
	require.NotNil(t, pkg)
	assert.Equal(t, "test-workspace-context", pkg.Name)
	assert.Equal(t, "custom-elements.json", pkg.CustomElements)
}

func TestMapWorkspaceContext_Manifest(t *testing.T) {
	ctx := workspace.NewMapWorkspaceContext(t, "testdata/workspace-context")
	err := ctx.Init()
	require.NoError(t, err)

	manifest, err := ctx.Manifest()
	require.NoError(t, err)
	require.NotNil(t, manifest)
	assert.Equal(t, "2.1.0", manifest.SchemaVersion)
	assert.Len(t, manifest.Modules, 1)
}

func TestMapWorkspaceContext_CustomElementsManifestPath(t *testing.T) {
	ctx := workspace.NewMapWorkspaceContext(t, "testdata/workspace-context")
	err := ctx.Init()
	require.NoError(t, err)

	assert.Equal(t, "/custom-elements.json", ctx.CustomElementsManifestPath())
}

func TestMapWorkspaceContext_ReadFile(t *testing.T) {
	ctx := workspace.NewMapWorkspaceContext(t, "testdata/workspace-context")
	err := ctx.Init()
	require.NoError(t, err)

	rc, err := ctx.ReadFile("/src/my-element.js")
	require.NoError(t, err)
	defer func() { _ = rc.Close() }()

	data, err := io.ReadAll(rc)
	require.NoError(t, err)
	assert.Contains(t, string(data), "MyElement")
}

func TestMapWorkspaceContext_ReadFile_NotFound(t *testing.T) {
	ctx := workspace.NewMapWorkspaceContext(t, "testdata/workspace-context")
	err := ctx.Init()
	require.NoError(t, err)

	_, err = ctx.ReadFile("/nonexistent.js")
	assert.Error(t, err)
}

func TestMapWorkspaceContext_Glob(t *testing.T) {
	ctx := workspace.NewMapWorkspaceContext(t, "testdata/workspace-context")
	err := ctx.Init()
	require.NoError(t, err)

	matches, err := ctx.Glob("*.json")
	require.NoError(t, err)
	assert.Contains(t, matches, "package.json")
	assert.Contains(t, matches, "custom-elements.json")
}

func TestMapWorkspaceContext_PathResolution(t *testing.T) {
	ctx := workspace.NewMapWorkspaceContext(t, "testdata/workspace-context")
	err := ctx.Init()
	require.NoError(t, err)

	t.Run("ModulePathToFS", func(t *testing.T) {
		result := ctx.ModulePathToFS("src/my-element.js")
		assert.Equal(t, "/src/my-element.js", result)
	})

	t.Run("FSPathToModule", func(t *testing.T) {
		mod, err := ctx.FSPathToModule("/src/my-element.js")
		require.NoError(t, err)
		assert.Equal(t, "src/my-element.js", mod)
	})

	t.Run("ResolveModuleDependency_relative", func(t *testing.T) {
		resolved, err := ctx.ResolveModuleDependency("src/my-element.js", "./helper.js")
		require.NoError(t, err)
		assert.Equal(t, "src/helper.js", resolved)
	})
}

func TestMapWorkspaceContext_Config(t *testing.T) {
	ctx := workspace.NewMapWorkspaceContext(t, "testdata/workspace-context")
	err := ctx.Init()
	require.NoError(t, err)

	cfg, err := ctx.Config()
	require.NoError(t, err)
	assert.NotNil(t, cfg)
}

func TestMapWorkspaceContext_DesignTokensCache(t *testing.T) {
	ctx := workspace.NewMapWorkspaceContext(t, "testdata/workspace-context")
	err := ctx.Init()
	require.NoError(t, err)

	cache := ctx.DesignTokensCache()
	assert.NotNil(t, cache)
}

func TestMapWorkspaceContext_Cleanup(t *testing.T) {
	ctx := workspace.NewMapWorkspaceContext(t, "testdata/workspace-context")
	err := ctx.Init()
	require.NoError(t, err)
	assert.NoError(t, ctx.Cleanup())
}
