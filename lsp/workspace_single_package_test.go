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

// TestWorkspaceSinglePackage tests a workspace with a single package directly (like PatternFly Elements)
// Pattern: workspaces: ["./elements"] where ./elements is the package itself
func TestWorkspaceSinglePackage(t *testing.T) {
	fixturePath, err := filepath.Abs(filepath.Join("testdata", "integration", "workspace-elements-single"))
	require.NoError(t, err, "Failed to get absolute path")

	wsCtx := workspace.NewFileSystemWorkspaceContext(fixturePath)
	registry, err := lsp.NewRegistryWithDefaults()
	require.NoError(t, err, "Failed to create registry")

	err = registry.LoadFromWorkspace(wsCtx)
	require.NoError(t, err, "Failed to load from workspace")

	t.Run("Element loaded from generated manifest", func(t *testing.T) {
		element, exists := registry.Element("my-element")
		assert.True(t, exists, "Expected my-element to be loaded from in-memory generated manifest")
		if exists {
			assert.Equal(t, "my-element", element.TagName)
		}
	})

	t.Run("No false positive diagnostics", func(t *testing.T) {
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

		htmlPath := filepath.Join(fixturePath, "test.html")
		content, err := os.ReadFile(htmlPath)
		require.NoError(t, err, "Failed to read test HTML file")

		doc := dm.OpenDocument("file://"+htmlPath, string(content), 1)
		require.NotNil(t, doc, "Failed to open document")

		diagnostics := publishDiagnostics.AnalyzeTagNameDiagnosticsForTest(server, doc)

		// Should have NO "unknown element" diagnostic for <my-element>
		for _, diag := range diagnostics {
			if strings.Contains(diag.Message, "my-element") && strings.Contains(diag.Message, "Unknown") {
				t.Errorf("FALSE POSITIVE: Got 'unknown element' diagnostic for workspace element: %s", diag.Message)
			}
		}
	})
}
