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

// TestSimpleRepoInMemoryGeneration tests the RHDS case:
// - Simple repo (no workspaces)
// - package.json exists with name but NO customElements field
// - TypeScript source files with custom elements exist
// - LSP should generate manifest in-memory for the workspace root
func TestSimpleRepoInMemoryGeneration(t *testing.T) {
	fixturePath, err := filepath.Abs(filepath.Join("test", "fixtures", "simple-repo-no-manifest"))
	require.NoError(t, err, "Failed to get absolute path")

	// Verify this is a simple repo (no workspaces in package.json)
	packageJSONPath := filepath.Join(fixturePath, "package.json")
	data, err := os.ReadFile(packageJSONPath)
	require.NoError(t, err, "Failed to read package.json")
	require.NotContains(t, string(data), "workspaces", "Fixture should NOT have workspaces field")

	wsCtx := workspace.NewFileSystemWorkspaceContext(fixturePath)
	registry, err := lsp.NewRegistryWithDefaults()
	require.NoError(t, err, "Failed to create registry")

	err = registry.LoadFromWorkspace(wsCtx)
	require.NoError(t, err, "Failed to load from workspace")

	t.Run("Elements loaded from in-memory generation", func(t *testing.T) {
		// The fixture has my-card.ts and my-button.ts
		// These should be loaded via in-memory manifest generation

		cardElement, cardExists := registry.Element("my-card")
		assert.True(t, cardExists, "Expected my-card to be loaded from in-memory generated manifest")
		if cardExists {
			assert.Equal(t, "my-card", cardElement.TagName)
			t.Logf("✓ my-card loaded with %d attributes", len(cardElement.Attributes))
		}

		buttonElement, buttonExists := registry.Element("my-button")
		assert.True(t, buttonExists, "Expected my-button to be loaded from in-memory generated manifest")
		if buttonExists {
			assert.Equal(t, "my-button", buttonElement.TagName)
			t.Logf("✓ my-button loaded with %d attributes", len(buttonElement.Attributes))
		}

		// Verify we loaded at least these 2 elements
		allTags := registry.AllTagNames()
		assert.GreaterOrEqual(t, len(allTags), 2, "Expected at least 2 elements from in-memory generation")
		t.Logf("Total elements loaded: %d", len(allTags))
	})

	t.Run("Package name preserved from package.json", func(t *testing.T) {
		// Verify elements have correct package name from package.json
		if def, exists := registry.ElementDefinition("my-card"); exists {
			packageName := def.PackageName()
			assert.Equal(t, "@test/simple-components", packageName,
				"Expected package name to be preserved from package.json")
			t.Logf("✓ Package name: %s", packageName)
		}
	})

	t.Run("No false positive diagnostics", func(t *testing.T) {
		// Create LSP server and verify no "unknown element" diagnostics
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

		// Open test.html which uses <my-card> and <my-button>
		htmlPath := filepath.Join(fixturePath, "test.html")
		content, err := os.ReadFile(htmlPath)
		require.NoError(t, err, "Failed to read test.html")

		doc := dm.OpenDocument("file://"+htmlPath, string(content), 1)
		require.NotNil(t, doc, "Failed to open document")

		diagnostics := publishDiagnostics.AnalyzeTagNameDiagnosticsForTest(server, doc)

		// Should have NO "unknown element" diagnostics for our elements
		for _, diag := range diagnostics {
			message := diag.Message
			if (strings.Contains(message, "my-card") || strings.Contains(message, "my-button")) &&
				strings.Contains(message, "Unknown") {
				t.Errorf("FALSE POSITIVE: Got 'unknown element' diagnostic for in-memory generated element: %s", message)
			}
		}

		t.Logf("✓ No false positive diagnostics (%d total diagnostics)", len(diagnostics))
	})

	t.Run("Element definitions available", func(t *testing.T) {
		// Verify that element definitions from in-memory generated manifest are available
		// Note: Attributes may not be fully parsed in this fixture since it lacks proper imports
		// The important thing is that the elements themselves are recognized

		cardDef, exists := registry.ElementDefinition("my-card")
		assert.True(t, exists, "Expected my-card element definition to exist")
		if exists {
			assert.Equal(t, "@test/simple-components", cardDef.PackageName())
			t.Logf("✓ my-card element definition available (package: %s)", cardDef.PackageName())
		}

		buttonDef, exists := registry.ElementDefinition("my-button")
		assert.True(t, exists, "Expected my-button element definition to exist")
		if exists {
			assert.Equal(t, "@test/simple-components", buttonDef.PackageName())
			t.Logf("✓ my-button element definition available (package: %s)", buttonDef.PackageName())
		}
	})
}

// TestSimpleRepoInMemoryGeneration_ManifestCount verifies manifest statistics
func TestSimpleRepoInMemoryGeneration_ManifestCount(t *testing.T) {
	fixturePath, err := filepath.Abs(filepath.Join("test", "fixtures", "simple-repo-no-manifest"))
	require.NoError(t, err, "Failed to get absolute path")

	wsCtx := workspace.NewFileSystemWorkspaceContext(fixturePath)
	registry, err := lsp.NewRegistryWithDefaults()
	require.NoError(t, err, "Failed to create registry")

	err = registry.LoadFromWorkspace(wsCtx)
	require.NoError(t, err, "Failed to load from workspace")

	// Should have exactly 1 manifest (the in-memory generated workspace manifest)
	manifestCount := len(registry.Manifests)
	assert.Equal(t, 1, manifestCount, "Expected exactly 1 manifest (workspace in-memory)")

	// Should have elements from that manifest
	elementCount := len(registry.AllTagNames())
	assert.GreaterOrEqual(t, elementCount, 2, "Expected at least 2 elements from workspace")

	t.Logf("✓ Loaded %d elements from %d manifest(s)", elementCount, manifestCount)
}
