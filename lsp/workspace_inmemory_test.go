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
	"os"
	"path/filepath"
	"strings"
	"testing"

	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/methods/textDocument/publishDiagnostics"
	"bennypowers.dev/cem/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestWorkspaceInMemoryGeneration_MissingManifest tests that when a workspace package
// declares customElements but the manifest file doesn't exist, the LSP generates it in-memory
func TestWorkspaceInMemoryGeneration_MissingManifest(t *testing.T) {
	// This test uses the PatternFly Elements-style structure:
	// - Workspace package declares "customElements": "custom-elements.json"
	// - But the file doesn't exist yet (hasn't been built)
	// - LSP should generate the manifest in-memory and use it for diagnostics
	fixturePath, err := filepath.Abs(filepath.Join("test", "fixtures", "workspace-missing-manifest"))
	require.NoError(t, err, "Failed to get absolute path")

	wsCtx := workspace.NewFileSystemWorkspaceContext(fixturePath)
	registry, err := lsp.NewRegistryWithDefaults()
	require.NoError(t, err, "Failed to create registry")

	err = registry.LoadFromWorkspace(wsCtx)
	require.NoError(t, err, "Failed to load from workspace")

	t.Run("Component element loaded from generated manifest", func(t *testing.T) {
		// The fixture has a workspace package "my-components" with:
		// - package.json: "customElements": "custom-elements.json"
		// - my-button.ts: defines <my-button>
		// - No custom-elements.json file (should be generated in-memory)

		element, exists := registry.Element("my-button")
		assert.True(t, exists, "Expected my-button to be loaded from in-memory generated manifest")
		if exists {
			assert.Equal(t, "my-button", element.TagName)
			// Note: Attributes may not be fully parsed without proper imports in the test fixture
			// The important thing is that the element itself is recognized
		}
	})

	t.Run("No false positive diagnostics for generated elements", func(t *testing.T) {
		// Create an LSP server and test that using <my-button> doesn't trigger diagnostics
		server, err := lsp.NewServer(wsCtx, lsp.TransportStdio)
		require.NoError(t, err, "Failed to create server")
		defer func() {
			_ = server.Close()
		}()

		err = server.InitializeForTesting()
		require.NoError(t, err, "Failed to initialize server")

		dm, err := document.NewDocumentManager()
		require.NoError(t, err, "Failed to create DocumentManager")
		defer dm.Close()

		// Test HTML file that uses <my-button>
		htmlPath := filepath.Join(fixturePath, "test.html")
		content, err := os.ReadFile(htmlPath)
		require.NoError(t, err, "Failed to read test HTML file")

		doc := dm.OpenDocument("file://"+htmlPath, string(content), 1)
		require.NotNil(t, doc, "Failed to open document")

		diagnostics := publishDiagnostics.AnalyzeTagNameDiagnosticsForTest(server, doc)

		// CRITICAL: Should have NO "unknown element" diagnostic for <my-button>
		for _, diag := range diagnostics {
			if strings.Contains(diag.Message, "my-button") && strings.Contains(diag.Message, "Unknown") {
				t.Errorf("FALSE POSITIVE: Got 'unknown element' diagnostic for in-memory generated element: %s", diag.Message)
			}
		}
	})
}

// TestWorkspaceInMemoryGeneration_WithExistingManifest tests that when a manifest file exists,
// it's used instead of generating in-memory (existing behavior preserved)
func TestWorkspaceInMemoryGeneration_WithExistingManifest(t *testing.T) {
	// This test verifies that existing manifests still work as before
	fixturePath, err := filepath.Abs(filepath.Join("test", "fixtures", "workspace-npm"))
	require.NoError(t, err, "Failed to get absolute path")

	wsCtx := workspace.NewFileSystemWorkspaceContext(fixturePath)
	registry, err := lsp.NewRegistryWithDefaults()
	require.NoError(t, err, "Failed to create registry")

	err = registry.LoadFromWorkspace(wsCtx)
	require.NoError(t, err, "Failed to load from workspace")

	t.Run("Existing manifest still loaded correctly", func(t *testing.T) {
		element, exists := registry.Element("my-element-a")
		assert.True(t, exists, "Expected my-element-a to be loaded from existing manifest")
		if exists {
			assert.Equal(t, "my-element-a", element.TagName)
		}
	})
}

// TestWorkspaceRootInMemoryGeneration tests that when a workspace root (not a workspace package)
// has no manifest file, it generates one in-memory. This is the RHDS case where:
// - package.json exists but has no "customElements" field OR
// - package.json has "customElements" but the file doesn't exist
func TestWorkspaceRootInMemoryGeneration(t *testing.T) {
	// Use the existing workspace-missing-manifest fixture which simulates a workspace root
	fixturePath, err := filepath.Abs(filepath.Join("test", "fixtures", "workspace-missing-manifest"))
	require.NoError(t, err, "Failed to get absolute path")

	wsCtx := workspace.NewFileSystemWorkspaceContext(fixturePath)
	registry, err := lsp.NewRegistryWithDefaults()
	require.NoError(t, err, "Failed to create registry")

	err = registry.LoadFromWorkspace(wsCtx)
	require.NoError(t, err, "Failed to load from workspace")

	t.Run("Workspace root manifest generated in-memory", func(t *testing.T) {
		// The fixture workspace itself should have elements loaded from in-memory generation
		allTags := registry.AllTagNames()
		assert.NotEmpty(t, allTags, "Expected workspace to have elements from in-memory generation")

		// Check that my-button is available (from workspace packages)
		element, exists := registry.Element("my-button")
		assert.True(t, exists, "Expected my-button to be loaded")
		if exists {
			assert.Equal(t, "my-button", element.TagName)
		}
	})

	t.Run("Package name preserved from package.json", func(t *testing.T) {
		// Verify that elements loaded from workspace root have correct package name
		if def, exists := registry.ElementDefinition("my-button"); exists {
			packageName := def.PackageName()
			// The fixture workspace package has name "my-components"
			assert.Equal(t, "my-components", packageName, "Expected package name to be preserved from package.json")
		}
	})
}
