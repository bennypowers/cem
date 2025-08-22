//go:build e2e

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
package completion_test

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	G "bennypowers.dev/cem/generate"
	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/methods/textDocument/completion"
	M "bennypowers.dev/cem/manifest"
	W "bennypowers.dev/cem/workspace"
)

// TestServerLevelIntegration tests the complete server integration including
// manifest reloading behavior exactly as it would work in real usage
func TestServerLevelIntegration(t *testing.T) {
	// Create a temporary workspace directory that mimics a real project
	tempDir := t.TempDir()

	// Create package.json
	packageJSON := `{
  "name": "test-project", 
  "customElements": "dist/custom-elements.json"
}`

	err := os.WriteFile(filepath.Join(tempDir, "package.json"), []byte(packageJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to create package.json: %v", err)
	}

	// Create source directory and TypeScript file
	srcDir := filepath.Join(tempDir, "src")
	err = os.MkdirAll(srcDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create src directory: %v", err)
	}

	initialTSContent := `import { LitElement, html } from 'lit';
import { customElement, property } from 'lit/decorators.js';

@customElement('test-alert')
export class TestAlert extends LitElement {
  /** The state of the alert */
  @property() state: 'info' | 'success' | 'warning' = 'info';

  render() {
    return html` + "`<div class=\"alert alert-${this.state}\"><slot></slot></div>`" + `;
  }
}`

	tsFilePath := filepath.Join(srcDir, "test-alert.ts")
	err = os.WriteFile(tsFilePath, []byte(initialTSContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create TypeScript file: %v", err)
	}

	// Create dist directory
	distDir := filepath.Join(tempDir, "dist")
	err = os.MkdirAll(distDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create dist directory: %v", err)
	}

	// Create .config directory and cem.yaml configuration
	configDir := filepath.Join(tempDir, ".config")
	err = os.MkdirAll(configDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create config directory: %v", err)
	}

	cemConfig := `generate:
  files:
    - src/test-alert.ts
`

	err = os.WriteFile(filepath.Join(configDir, "cem.yaml"), []byte(cemConfig), 0644)
	if err != nil {
		t.Fatalf("Failed to create cem.yaml: %v", err)
	}

	// Create initial manifest by running generate manually
	workspace := W.NewFileSystemWorkspaceContext(tempDir)
	err = workspace.Init()
	if err != nil {
		t.Fatalf("Failed to initialize workspace: %v", err)
	}

	// Run initial generation to create the manifest
	generateSession, err := G.NewGenerateSession(workspace)
	if err != nil {
		t.Fatalf("Failed to create generate session: %v", err)
	}
	defer generateSession.Close()

	pkg, err := generateSession.GenerateFullManifest(context.Background())
	if err != nil {
		t.Fatalf("Failed to generate initial manifest: %v", err)
	}

	// Write initial manifest to file
	manifestStr, err := M.SerializeToString(pkg)
	if err != nil {
		t.Fatalf("Failed to serialize manifest: %v", err)
	}

	manifestPath := filepath.Join(distDir, "custom-elements.json")
	err = os.WriteFile(manifestPath, []byte(manifestStr), 0644)
	if err != nil {
		t.Fatalf("Failed to write manifest: %v", err)
	}

	// Now create a full LSP server instance - this simulates real usage
	server, err := lsp.NewServer(workspace, lsp.TransportStdio)
	if err != nil {
		t.Fatalf("Failed to create LSP server: %v", err)
	}
	defer server.Close()

	// Initialize the server (this loads manifests and starts file watching)
	err = server.InitializeForTesting()
	if err != nil {
		t.Fatalf("Failed to initialize server: %v", err)
	}

	// Create document manager
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	// Set up the server context adapter with document manager
	ctx := &testCompletionContextWithDM{
		registry: server.Registry(), // We need to expose this for testing
		docMgr:   dm,
	}

	// Create HTML file that uses the element
	htmlContent := `<!DOCTYPE html>
<html>
<body>
  <test-alert state=""></test-alert>
</body>
</html>`

	htmlFilePath := filepath.Join(tempDir, "index.html")
	err = os.WriteFile(htmlFilePath, []byte(htmlContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create HTML file: %v", err)
	}

	// Open the HTML document in document manager
	htmlURI := "file://" + htmlFilePath
	doc := dm.OpenDocument(htmlURI, htmlContent, 1)
	if doc == nil {
		t.Fatalf("Failed to open HTML document")
	}

	// Test initial completions - should have info, success, warning
	initialItems := completion.GetAttributeValueCompletions(ctx, "test-alert", "state")

	t.Logf("Initial completions: %v", getCompletionLabels(initialItems))

	// Verify initial completions contain expected values
	hasInfo := false
	hasSuccess := false
	hasWarning := false
	for _, item := range initialItems {
		switch item.Label {
		case "info":
			hasInfo = true
		case "success":
			hasSuccess = true
		case "warning":
			hasWarning = true
		}
	}

	if !hasInfo || !hasSuccess || !hasWarning {
		t.Fatalf("Initial completions missing expected values. Got: %v", getCompletionLabels(initialItems))
	}

	// Now simulate the user editing the TypeScript file to add 'error'
	t.Logf("=== User edits test-alert.ts to add 'error' ===")

	updatedTSContent := strings.Replace(initialTSContent,
		`state: 'info' | 'success' | 'warning' = 'info'`,
		`state: 'info' | 'success' | 'warning' | 'error' = 'info'`,
		1)

	err = os.WriteFile(tsFilePath, []byte(updatedTSContent), 0644)
	if err != nil {
		t.Fatalf("Failed to update TypeScript file: %v", err)
	}

	t.Logf("File updated, waiting for generate watcher to process changes...")

	// Wait for the generate watcher to detect changes and regenerate
	time.Sleep(5 * time.Second)

	// Test updated completions - should now include 'error'
	updatedItems := completion.GetAttributeValueCompletions(ctx, "test-alert", "state")

	t.Logf("Updated completions: %v", getCompletionLabels(updatedItems))

	// Verify updated completions contain 'error'
	hasUpdatedInfo := false
	hasUpdatedSuccess := false
	hasUpdatedWarning := false
	hasError := false
	for _, item := range updatedItems {
		switch item.Label {
		case "info":
			hasUpdatedInfo = true
		case "success":
			hasUpdatedSuccess = true
		case "warning":
			hasUpdatedWarning = true
		case "error":
			hasError = true
		}
	}

	if !hasUpdatedInfo || !hasUpdatedSuccess || !hasUpdatedWarning {
		t.Errorf("Updated completions missing original values. Got: %v", getCompletionLabels(updatedItems))
	}

	if !hasError {
		t.Errorf("Updated completions missing 'error' value. Got: %v", getCompletionLabels(updatedItems))

		// Debug: Check what's in the registry
		if attrs, exists := ctx.Attributes("test-alert"); exists {
			if stateAttr, hasState := attrs["state"]; hasState {
				t.Logf("Registry shows state attribute type: %s", stateAttr.Type.Text)
			}
		}

		// Debug: Check the manifest file content
		manifestContent, err := os.ReadFile(manifestPath)
		if err == nil {
			t.Logf("Current manifest content: %s", string(manifestContent))
		}
	} else {
		t.Logf("✅ SUCCESS: 'error' found in completions!")
	}
}
