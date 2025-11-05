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
	"testing"

	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/methods/textDocument/publishDiagnostics"
	"bennypowers.dev/cem/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestWorkspaceDiagnostics_NoFalsePositives_npm tests that workspace sibling elements
// do NOT produce "unknown element" diagnostics when properly imported
func TestWorkspaceDiagnostics_NoFalsePositives_npm(t *testing.T) {
	fixturePath, err := filepath.Abs(filepath.Join("test", "fixtures", "workspace-npm"))
	require.NoError(t, err, "Failed to get absolute path")

	// Create workspace context
	wsCtx := workspace.NewFileSystemWorkspaceContext(fixturePath)

	// Create LSP server
	server, err := lsp.NewServer(wsCtx, lsp.TransportStdio)
	require.NoError(t, err, "Failed to create server")
	defer server.Close()

	// Initialize server (this loads manifests)
	err = server.InitializeForTesting()
	require.NoError(t, err, "Failed to initialize server")

	// Create DocumentManager
	dm, err := document.NewDocumentManager()
	require.NoError(t, err, "Failed to create DocumentManager")
	defer dm.Close()

	// Read the test HTML file
	htmlPath := filepath.Join(fixturePath, "packages", "component-b", "test.html")
	content, err := os.ReadFile(htmlPath)
	require.NoError(t, err, "Failed to read test HTML file")

	// Open document
	doc := dm.OpenDocument("file://"+htmlPath, string(content), 1)
	require.NotNil(t, doc, "Failed to open document")

	// Analyze tag diagnostics
	diagnostics := publishDiagnostics.AnalyzeTagNameDiagnosticsForTest(server, doc)

	// CRITICAL: Should have NO "unknown element" diagnostic for <my-element-a>
	// because it's a workspace sibling that IS imported
	for _, diag := range diagnostics {
		if contains(diag.Message, "my-element-a") && contains(diag.Message, "Unknown") {
			t.Errorf("FALSE POSITIVE: Got 'unknown element' diagnostic for workspace sibling: %s", diag.Message)
		}
		if contains(diag.Message, "my-element-a") && contains(diag.Message, "not imported") {
			t.Errorf("FALSE POSITIVE: Got 'missing import' diagnostic for workspace sibling: %s", diag.Message)
		}
	}

	// Verify the element is actually in the registry
	element, exists := server.Element("my-element-a")
	assert.True(t, exists, "my-element-a should be loaded from workspace sibling")
	if exists {
		assert.Equal(t, "my-element-a", element.TagName)
	}
}

// TestWorkspaceDiagnostics_NoFalsePositives_yarn tests yarn workspace sibling elements
func TestWorkspaceDiagnostics_NoFalsePositives_yarn(t *testing.T) {
	fixturePath, err := filepath.Abs(filepath.Join("test", "fixtures", "workspace-yarn"))
	require.NoError(t, err, "Failed to get absolute path")

	// Create workspace context
	wsCtx := workspace.NewFileSystemWorkspaceContext(fixturePath)

	// Create LSP server
	server, err := lsp.NewServer(wsCtx, lsp.TransportStdio)
	require.NoError(t, err, "Failed to create server")
	defer server.Close()

	// Initialize server
	err = server.InitializeForTesting()
	require.NoError(t, err, "Failed to initialize server")

	// Create DocumentManager
	dm, err := document.NewDocumentManager()
	require.NoError(t, err, "Failed to create DocumentManager")
	defer dm.Close()

	// Read the test HTML file
	htmlPath := filepath.Join(fixturePath, "packages", "component-b", "test.html")
	content, err := os.ReadFile(htmlPath)
	require.NoError(t, err, "Failed to read test HTML file")

	// Open document
	doc := dm.OpenDocument("file://"+htmlPath, string(content), 1)
	require.NotNil(t, doc, "Failed to open document")

	// Analyze tag diagnostics
	diagnostics := publishDiagnostics.AnalyzeTagNameDiagnosticsForTest(server, doc)

	// Should have NO false positive diagnostics
	for _, diag := range diagnostics {
		if contains(diag.Message, "my-element-a") && contains(diag.Message, "Unknown") {
			t.Errorf("FALSE POSITIVE: Got 'unknown element' diagnostic for yarn workspace sibling: %s", diag.Message)
		}
		if contains(diag.Message, "my-element-a") && contains(diag.Message, "not imported") {
			t.Errorf("FALSE POSITIVE: Got 'missing import' diagnostic for yarn workspace sibling: %s", diag.Message)
		}
	}

	// Verify the element is in the registry
	element, exists := server.Element("my-element-a")
	assert.True(t, exists, "my-element-a should be loaded from yarn workspace sibling")
	if exists {
		assert.Equal(t, "my-element-a", element.TagName)
	}
}

// TestWorkspaceDiagnostics_NoFalsePositives_pnpm tests pnpm workspace sibling elements
func TestWorkspaceDiagnostics_NoFalsePositives_pnpm(t *testing.T) {
	fixturePath, err := filepath.Abs(filepath.Join("test", "fixtures", "workspace-pnpm"))
	require.NoError(t, err, "Failed to get absolute path")

	// Create workspace context
	wsCtx := workspace.NewFileSystemWorkspaceContext(fixturePath)

	// Create LSP server
	server, err := lsp.NewServer(wsCtx, lsp.TransportStdio)
	require.NoError(t, err, "Failed to create server")
	defer server.Close()

	// Initialize server
	err = server.InitializeForTesting()
	require.NoError(t, err, "Failed to initialize server")

	// Create DocumentManager
	dm, err := document.NewDocumentManager()
	require.NoError(t, err, "Failed to create DocumentManager")
	defer dm.Close()

	// Read the test HTML file
	htmlPath := filepath.Join(fixturePath, "packages", "component-b", "test.html")
	content, err := os.ReadFile(htmlPath)
	require.NoError(t, err, "Failed to read test HTML file")

	// Open document
	doc := dm.OpenDocument("file://"+htmlPath, string(content), 1)
	require.NotNil(t, doc, "Failed to open document")

	// Analyze tag diagnostics
	diagnostics := publishDiagnostics.AnalyzeTagNameDiagnosticsForTest(server, doc)

	// Should have NO false positive diagnostics
	for _, diag := range diagnostics {
		if contains(diag.Message, "my-element-a") && contains(diag.Message, "Unknown") {
			t.Errorf("FALSE POSITIVE: Got 'unknown element' diagnostic for pnpm workspace sibling: %s", diag.Message)
		}
		if contains(diag.Message, "my-element-a") && contains(diag.Message, "not imported") {
			t.Errorf("FALSE POSITIVE: Got 'missing import' diagnostic for pnpm workspace sibling: %s", diag.Message)
		}
	}

	// Verify the element is in the registry
	element, exists := server.Element("my-element-a")
	assert.True(t, exists, "my-element-a should be loaded from pnpm workspace sibling")
	if exists {
		assert.Equal(t, "my-element-a", element.TagName)
	}
}

// contains is a helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) &&
		(s[:len(substr)] == substr || s[len(s)-len(substr):] == substr ||
		containsMiddle(s, substr)))
}

func containsMiddle(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
