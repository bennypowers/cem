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

// TestWorkspacePatterns_NestedDoublestar tests that "**" doublestar patterns
// find deeply nested workspace packages
func TestWorkspacePatterns_NestedDoublestar(t *testing.T) {
	fixturePath := filepath.Join("test", "fixtures", "workspace-nested")

	// Create workspace context
	wsCtx := workspace.NewFileSystemWorkspaceContext(fixturePath)

	// Create registry and load from workspace
	registry, err := lsp.NewRegistryWithDefaults()
	require.NoError(t, err, "Failed to create registry")
	err = registry.LoadFromWorkspace(wsCtx)
	require.NoError(t, err, "Failed to load from workspace")

	// Test that all deeply nested packages are found
	// Pattern: "components/**" should find:
	// - components/ui/button
	// - components/ui/input
	// - components/layout/grid

	t.Run("ui-button loaded from nested path", func(t *testing.T) {
		element, exists := registry.Element("ui-button")
		assert.True(t, exists, "Expected ui-button to be loaded from components/ui/button")
		if exists {
			assert.Equal(t, "ui-button", element.TagName)
		}
	})

	t.Run("ui-input loaded from nested path", func(t *testing.T) {
		element, exists := registry.Element("ui-input")
		assert.True(t, exists, "Expected ui-input to be loaded from components/ui/input")
		if exists {
			assert.Equal(t, "ui-input", element.TagName)
		}
	})

	t.Run("layout-grid loaded from nested path", func(t *testing.T) {
		element, exists := registry.Element("layout-grid")
		assert.True(t, exists, "Expected layout-grid to be loaded from components/layout/grid")
		if exists {
			assert.Equal(t, "layout-grid", element.TagName)
		}
	})

	t.Run("all nested elements found", func(t *testing.T) {
		allTagNames := registry.AllTagNames()
		assert.Contains(t, allTagNames, "ui-button", "Expected ui-button in all tag names")
		assert.Contains(t, allTagNames, "ui-input", "Expected ui-input in all tag names")
		assert.Contains(t, allTagNames, "layout-grid", "Expected layout-grid in all tag names")
		assert.GreaterOrEqual(t, len(allTagNames), 3, "Expected at least 3 elements from nested packages")
	})
}

// TestWorkspacePatterns_Negation tests that "!" negation patterns
// exclude matching packages from being loaded
func TestWorkspacePatterns_Negation(t *testing.T) {
	fixturePath := filepath.Join("test", "fixtures", "workspace-negation")

	// Create workspace context
	wsCtx := workspace.NewFileSystemWorkspaceContext(fixturePath)

	// Create registry and load from workspace
	registry, err := lsp.NewRegistryWithDefaults()
	require.NoError(t, err, "Failed to create registry")
	err = registry.LoadFromWorkspace(wsCtx)
	require.NoError(t, err, "Failed to load from workspace")

	// Pattern: ["packages/*", "!packages/legacy"]
	// Should load: component-a, component-b
	// Should NOT load: legacy

	t.Run("component-a loaded (not negated)", func(t *testing.T) {
		element, exists := registry.Element("comp-a")
		assert.True(t, exists, "Expected comp-a to be loaded")
		if exists {
			assert.Equal(t, "comp-a", element.TagName)
		}
	})

	t.Run("component-b loaded (not negated)", func(t *testing.T) {
		element, exists := registry.Element("comp-b")
		assert.True(t, exists, "Expected comp-b to be loaded")
		if exists {
			assert.Equal(t, "comp-b", element.TagName)
		}
	})

	t.Run("legacy NOT loaded (negated)", func(t *testing.T) {
		_, exists := registry.Element("legacy-comp")
		assert.False(t, exists, "Expected legacy-comp to be EXCLUDED by negation pattern")
	})

	t.Run("negation pattern applied correctly", func(t *testing.T) {
		allTagNames := registry.AllTagNames()
		assert.Contains(t, allTagNames, "comp-a", "Expected comp-a to be loaded")
		assert.Contains(t, allTagNames, "comp-b", "Expected comp-b to be loaded")
		assert.NotContains(t, allTagNames, "legacy-comp", "Expected legacy-comp to be EXCLUDED")
	})
}

// TestWorkspacePatterns_EdgeCases tests edge cases like empty patterns, invalid patterns, etc.
func TestWorkspacePatterns_EdgeCases(t *testing.T) {
	t.Run("empty workspace patterns", func(t *testing.T) {
		// Create a temporary fixture without workspaces
		fixturePath := filepath.Join("test", "fixtures", "workspace-npm")

		wsCtx := workspace.NewFileSystemWorkspaceContext(fixturePath)
		registry, err := lsp.NewRegistryWithDefaults()
		require.NoError(t, err, "Failed to create registry")

		// Should not crash, just return no workspace packages
		err = registry.LoadFromWorkspace(wsCtx)
		assert.NoError(t, err, "Should handle empty workspace gracefully")
	})
}
