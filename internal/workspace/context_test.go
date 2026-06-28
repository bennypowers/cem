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
package workspace_test

import (
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/internal/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Inline: integration test, scalar assertions
func TestFileSystemWorkspaceContext_Manifest(t *testing.T) {
	root := absFixture(t, "project-with-manifest")
	ctx := workspace.NewFileSystemWorkspaceContext(root)
	require.NoError(t, ctx.Init())

	t.Run("CustomElementsManifestPath", func(t *testing.T) {
		assert.Equal(t, filepath.Join(root, "custom-elements.json"), ctx.CustomElementsManifestPath())
	})

	t.Run("Manifest loads and caches", func(t *testing.T) {
		m, err := ctx.Manifest()
		require.NoError(t, err)
		require.NotNil(t, m)
		assert.Len(t, m.Modules, 1)
		assert.Equal(t, "src/button.js", m.Modules[0].Path)

		m2, err := ctx.Manifest()
		require.NoError(t, err)
		assert.Same(t, m, m2, "second call returns cached manifest")
	})
}

func TestFileSystemWorkspaceContext_ModulePathToFS(t *testing.T) {
	root := absFixture(t, "project-with-manifest")
	ctx := workspace.NewFileSystemWorkspaceContext(root)

	assert.Equal(t, filepath.Join(root, "src/button.js"), ctx.ModulePathToFS("src/button.js"))
	absPath := filepath.FromSlash("/abs/path.js")
	assert.Equal(t, absPath, ctx.ModulePathToFS(absPath))
}

func TestFileSystemWorkspaceContext_Cleanup(t *testing.T) {
	root := absFixture(t, "project-with-manifest")
	ctx := workspace.NewFileSystemWorkspaceContext(root)
	assert.NoError(t, ctx.Cleanup())
}

func TestFileSystemWorkspaceContext_DesignTokensCache(t *testing.T) {
	root := absFixture(t, "project-with-manifest")
	ctx := workspace.NewFileSystemWorkspaceContext(root)
	assert.NotNil(t, ctx.DesignTokensCache())
}

func TestFileSystemWorkspaceContext_FSPathToModule(t *testing.T) {
	root := absFixture(t, "project-with-manifest")
	ctx := workspace.NewFileSystemWorkspaceContext(root)

	rel, err := ctx.FSPathToModule(filepath.Join(root, "src/button.js"))
	require.NoError(t, err)
	assert.Equal(t, "src/button.js", rel)
}

func TestLoadWorkspaceManifests(t *testing.T) {
	root := absFixture(t, "workspace-mode-single-package")

	pkgs, err := workspace.LoadWorkspaceManifests(root, platform.NewOSFileSystem())
	require.NoError(t, err)
	require.Len(t, pkgs, 1)
	assert.Equal(t, "@test/elements", pkgs[0].Name)
	assert.NotEmpty(t, pkgs[0].Manifest)
}

func TestLoadWorkspaceManifests_NonWorkspace(t *testing.T) {
	root := absFixture(t, "non-workspace")
	_, err := workspace.LoadWorkspaceManifests(root, platform.NewOSFileSystem())
	assert.Error(t, err)
}

func TestGetContextForSpec(t *testing.T) {
	t.Run("local path", func(t *testing.T) {
		root := absFixture(t, "project-with-manifest")
		ctx, err := workspace.GetContextForSpec(root)
		require.NoError(t, err)
		assert.NotNil(t, ctx)
	})

	t.Run("URL spec", func(t *testing.T) {
		ctx, err := workspace.GetContextForSpec("https://example.com/manifest.json")
		require.NoError(t, err)
		assert.NotNil(t, ctx)
	})

	t.Run("npm spec without local node_modules", func(t *testing.T) {
		ctx, err := workspace.GetContextForSpec("npm:@nonexistent/pkg@1.0.0")
		require.NoError(t, err)
		assert.NotNil(t, ctx)
	})
}
