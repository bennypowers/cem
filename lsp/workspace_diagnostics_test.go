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
	"io"
	"strings"
	"testing"

	testworkspace "bennypowers.dev/cem/internal/platform/testutil/workspace"
	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/methods/textDocument/publishDiagnostics"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.lsp.dev/protocol"
)

// Inline: integration test, scalar assertions
// Verifies no false-positive "unknown element" diagnostics for workspace sibling
// elements across npm, yarn, and pnpm workspace managers.

// TestWorkspaceDiagnostics_NoFalsePositives_npm tests that workspace sibling elements
// do NOT produce "unknown element" diagnostics when properly imported
func TestWorkspaceDiagnostics_NoFalsePositives_npm(t *testing.T) {
	wsCtx := testworkspace.NewMapWorkspaceContext(t, "testdata/integration/workspace-npm")

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

	htmlRC, err := wsCtx.ReadFile("/packages/component-b/test.html")
	require.NoError(t, err, "Failed to read test HTML file")
	content, err := io.ReadAll(htmlRC)
	_ = htmlRC.Close()
	require.NoError(t, err)
	htmlPath := "/packages/component-b/test.html"

	// Open document
	doc := dm.OpenDocument("file://"+htmlPath, string(content), 1)
	require.NotNil(t, doc, "Failed to open document")

	// Analyze tag diagnostics
	diagnostics := publishDiagnostics.AnalyzeTagNameDiagnosticsForTest(server, doc)

	// CRITICAL: Should have NO "unknown element" diagnostic for <my-element-a>
	// because it's a workspace sibling that IS imported
	for _, diag := range diagnostics {
		msg := string(diag.Message.(protocol.String))
		if strings.Contains(msg, "my-element-a") && strings.Contains(msg, "Unknown") {
			t.Errorf("FALSE POSITIVE: Got 'unknown element' diagnostic for workspace sibling: %s", msg)
		}
		if strings.Contains(msg, "my-element-a") && strings.Contains(msg, "not imported") {
			t.Errorf("FALSE POSITIVE: Got 'missing import' diagnostic for workspace sibling: %s", msg)
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
	wsCtx := testworkspace.NewMapWorkspaceContext(t, "testdata/integration/workspace-yarn")

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

	htmlRC, err := wsCtx.ReadFile("/packages/component-b/test.html")
	require.NoError(t, err, "Failed to read test HTML file")
	content, err := io.ReadAll(htmlRC)
	_ = htmlRC.Close()
	require.NoError(t, err)
	htmlPath := "/packages/component-b/test.html"

	// Open document
	doc := dm.OpenDocument("file://"+htmlPath, string(content), 1)
	require.NotNil(t, doc, "Failed to open document")

	// Analyze tag diagnostics
	diagnostics := publishDiagnostics.AnalyzeTagNameDiagnosticsForTest(server, doc)

	// Should have NO false positive diagnostics
	for _, diag := range diagnostics {
		msg := string(diag.Message.(protocol.String))
		if strings.Contains(msg, "my-element-a") && strings.Contains(msg, "Unknown") {
			t.Errorf("FALSE POSITIVE: Got 'unknown element' diagnostic for yarn workspace sibling: %s", msg)
		}
		if strings.Contains(msg, "my-element-a") && strings.Contains(msg, "not imported") {
			t.Errorf("FALSE POSITIVE: Got 'missing import' diagnostic for yarn workspace sibling: %s", msg)
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
	wsCtx := testworkspace.NewMapWorkspaceContext(t, "testdata/integration/workspace-pnpm")

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

	htmlRC, err := wsCtx.ReadFile("/packages/component-b/test.html")
	require.NoError(t, err, "Failed to read test HTML file")
	content, err := io.ReadAll(htmlRC)
	_ = htmlRC.Close()
	require.NoError(t, err)
	htmlPath := "/packages/component-b/test.html"

	// Open document
	doc := dm.OpenDocument("file://"+htmlPath, string(content), 1)
	require.NotNil(t, doc, "Failed to open document")

	// Analyze tag diagnostics
	diagnostics := publishDiagnostics.AnalyzeTagNameDiagnosticsForTest(server, doc)

	// Should have NO false positive diagnostics
	for _, diag := range diagnostics {
		msg := string(diag.Message.(protocol.String))
		if strings.Contains(msg, "my-element-a") && strings.Contains(msg, "Unknown") {
			t.Errorf("FALSE POSITIVE: Got 'unknown element' diagnostic for pnpm workspace sibling: %s", msg)
		}
		if strings.Contains(msg, "my-element-a") && strings.Contains(msg, "not imported") {
			t.Errorf("FALSE POSITIVE: Got 'missing import' diagnostic for pnpm workspace sibling: %s", msg)
		}
	}

	// Verify the element is in the registry
	element, exists := server.Element("my-element-a")
	assert.True(t, exists, "my-element-a should be loaded from pnpm workspace sibling")
	if exists {
		assert.Equal(t, "my-element-a", element.TagName)
	}
}
