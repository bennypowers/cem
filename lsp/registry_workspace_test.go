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
package lsp_test

import (
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestWorkspaceManifestLoading_npm tests that npm workspace packages are loaded into the registry
func TestWorkspaceManifestLoading_npm(t *testing.T) {
	fixturePath := filepath.Join("test", "fixtures", "workspace-npm")

	// Create workspace context
	wsCtx := workspace.NewFileSystemWorkspaceContext(fixturePath)

	// Create registry and load from workspace
	registry, err := lsp.NewRegistryWithDefaults()
	require.NoError(t, err, "Failed to create registry")
	err = registry.LoadFromWorkspace(wsCtx)
	require.NoError(t, err, "Failed to load from workspace")

	// Test that workspace package manifests are loaded
	t.Run("Component A manifest loaded", func(t *testing.T) {
		element, exists := registry.Element("my-element-a")
		assert.True(t, exists, "Expected my-element-a to be loaded from workspace package")
		if exists {
			assert.Equal(t, "my-element-a", element.TagName)
			assert.Len(t, element.Attributes, 1)
			assert.Equal(t, "label", element.Attributes[0].Name)
		}
	})

	t.Run("Component B manifest loaded", func(t *testing.T) {
		element, exists := registry.Element("my-element-b")
		assert.True(t, exists, "Expected my-element-b to be loaded from workspace package")
		if exists {
			assert.Equal(t, "my-element-b", element.TagName)
			assert.Len(t, element.Attributes, 1)
			assert.Equal(t, "title", element.Attributes[0].Name)
		}
	})

	t.Run("Both elements from different workspace packages", func(t *testing.T) {
		allTagNames := registry.AllTagNames()
		assert.Contains(t, allTagNames, "my-element-a", "Expected my-element-a in all tag names")
		assert.Contains(t, allTagNames, "my-element-b", "Expected my-element-b in all tag names")
		assert.GreaterOrEqual(t, len(allTagNames), 2, "Expected at least 2 elements loaded")
	})
}

// TestWorkspaceManifestLoading_yarn tests that yarn workspace packages are loaded into the registry
func TestWorkspaceManifestLoading_yarn(t *testing.T) {
	fixturePath := filepath.Join("test", "fixtures", "workspace-yarn")

	// Create workspace context
	wsCtx := workspace.NewFileSystemWorkspaceContext(fixturePath)

	// Create registry and load from workspace
	registry, err := lsp.NewRegistryWithDefaults()
	require.NoError(t, err, "Failed to create registry")
	err = registry.LoadFromWorkspace(wsCtx)
	require.NoError(t, err, "Failed to load from workspace")

	// Test that workspace package manifests are loaded
	t.Run("Component A manifest loaded", func(t *testing.T) {
		element, exists := registry.Element("my-element-a")
		assert.True(t, exists, "Expected my-element-a to be loaded from yarn workspace package")
		if exists {
			assert.Equal(t, "my-element-a", element.TagName)
		}
	})

	t.Run("Component B manifest loaded", func(t *testing.T) {
		element, exists := registry.Element("my-element-b")
		assert.True(t, exists, "Expected my-element-b to be loaded from yarn workspace package")
		if exists {
			assert.Equal(t, "my-element-b", element.TagName)
		}
	})

	t.Run("Yarn workspace detection", func(t *testing.T) {
		// Verify that yarn.lock is detected and workspace packages are loaded
		allTagNames := registry.AllTagNames()
		assert.Contains(t, allTagNames, "my-element-a")
		assert.Contains(t, allTagNames, "my-element-b")
	})
}

// TestWorkspaceManifestLoading_pnpm tests that pnpm workspace packages are loaded into the registry
func TestWorkspaceManifestLoading_pnpm(t *testing.T) {
	fixturePath := filepath.Join("test", "fixtures", "workspace-pnpm")

	// Create workspace context
	wsCtx := workspace.NewFileSystemWorkspaceContext(fixturePath)

	// Create registry and load from workspace
	registry, err := lsp.NewRegistryWithDefaults()
	require.NoError(t, err, "Failed to create registry")
	err = registry.LoadFromWorkspace(wsCtx)
	require.NoError(t, err, "Failed to load from workspace")

	// Test that workspace package manifests are loaded
	t.Run("Component A manifest loaded", func(t *testing.T) {
		element, exists := registry.Element("my-element-a")
		assert.True(t, exists, "Expected my-element-a to be loaded from pnpm workspace package")
		if exists {
			assert.Equal(t, "my-element-a", element.TagName)
		}
	})

	t.Run("Component B manifest loaded", func(t *testing.T) {
		element, exists := registry.Element("my-element-b")
		assert.True(t, exists, "Expected my-element-b to be loaded from pnpm workspace package")
		if exists {
			assert.Equal(t, "my-element-b", element.TagName)
		}
	})

	t.Run("pnpm-workspace.yaml detection", func(t *testing.T) {
		// Verify that pnpm-workspace.yaml is detected and workspace packages are loaded
		allTagNames := registry.AllTagNames()
		assert.Contains(t, allTagNames, "my-element-a")
		assert.Contains(t, allTagNames, "my-element-b")
	})
}
