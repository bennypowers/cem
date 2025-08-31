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
package publishDiagnostics_test

import (
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/methods/textDocument/publishDiagnostics"
	W "bennypowers.dev/cem/workspace"
)

func TestModuleGraphIntegration_ReExportResolution(t *testing.T) {
	// Create a temporary workspace with re-export scenario
	tempDir, err := os.MkdirTemp("", "module-graph-integration-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test files simulating my-tabs re-exporting my-tab
	if err := createTestFiles(tempDir); err != nil {
		t.Fatalf("Failed to create test files: %v", err)
	}

	// Create workspace context
	workspace := W.NewFileSystemWorkspaceContext(tempDir)
	err = workspace.Init()
	if err != nil {
		t.Fatalf("Failed to initialize workspace: %v", err)
	}

	// Create server and initialize with module graph support
	server, err := lsp.NewServer(workspace, lsp.TransportStdio)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	err = server.InitializeForTesting()
	if err != nil {
		t.Fatalf("Failed to initialize server: %v", err)
	}

	// Test HTML content that imports my-tabs.js and uses my-tab
	htmlContent := `<!DOCTYPE html>
<html>
<head>
    <script type="module">
        import './my-tabs.js';
    </script>
</head>
<body>
    <my-tabs>
        <my-tab slot="tab">Tab 1</my-tab>
        <my-tab slot="tab">Tab 2</my-tab>
    </my-tabs>
</body>
</html>`

	// Create document manager
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	// Create document
	doc := dm.OpenDocument("file://test.html", htmlContent, 1)
	if doc == nil {
		t.Fatal("Failed to open document")
	}

	// Run tag diagnostics
	diagnostics := publishDiagnostics.AnalyzeTagNameDiagnosticsForTest(server, doc)

	// Before module graph integration, this would show an error for my-tab
	// After integration, my-tab should be resolved via the re-export from my-tabs.js
	hasMyTabError := false
	for _, diagnostic := range diagnostics {
		if diagnostic.Message == "Custom element 'my-tab' is not imported. Add import from './my-tab.js'" ||
			diagnostic.Message == "Unknown custom element 'my-tab'. Did you mean 'my-tabs'?" {
			hasMyTabError = true
			t.Errorf("Found unexpected error for my-tab: %s", diagnostic.Message)
		}
	}

	if hasMyTabError {
		t.Error("Module graph integration failed: my-tab should be resolved via re-export from my-tabs.js")
	}

	// Test that my-tabs itself doesn't have errors
	hasMyTabsError := false
	for _, diagnostic := range diagnostics {
		if contains(diagnostic.Message, "my-tabs") {
			hasMyTabsError = true
			t.Errorf("Unexpected error for my-tabs: %s", diagnostic.Message)
		}
	}

	if hasMyTabsError {
		t.Error("my-tabs should not have any import errors")
	}
}

func TestModuleGraphIntegration_DirectVsReExport(t *testing.T) {
	// Create temporary workspace
	tempDir, err := os.MkdirTemp("", "direct-vs-reexport-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test files
	if err := createDirectVsReExportFiles(tempDir); err != nil {
		t.Fatalf("Failed to create test files: %v", err)
	}

	// Create workspace and server
	workspace := W.NewFileSystemWorkspaceContext(tempDir)
	if err := workspace.Init(); err != nil {
		t.Fatalf("Failed to initialize workspace: %v", err)
	}

	server, err := lsp.NewServer(workspace, lsp.TransportStdio)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	if err := server.InitializeForTesting(); err != nil {
		t.Fatalf("Failed to initialize server: %v", err)
	}

	// Test 1: Direct import should work
	directImportContent := `<!DOCTYPE html>
<html>
<head>
    <script type="module">
        import './button-element.js';
    </script>
</head>
<body>
    <my-button>Click me</my-button>
</body>
</html>`

	// Create document manager
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	// Test 1: Direct import should work
	doc1 := dm.OpenDocument("file://direct.html", directImportContent, 1)
	if doc1 == nil {
		t.Fatal("Failed to open document")
	}

	diagnostics1 := publishDiagnostics.AnalyzeTagNameDiagnosticsForTest(server, doc1)

	for _, diagnostic := range diagnostics1 {
		if contains(diagnostic.Message, "my-button") {
			t.Errorf("Direct import should not have errors for my-button: %s", diagnostic.Message)
		}
	}

	// Test 2: Re-export import should also work
	reExportContent := `<!DOCTYPE html>
<html>
<head>
    <script type="module">
        import './index.js';
    </script>
</head>
<body>
    <my-button>Click me</my-button>
</body>
</html>`

	doc2 := dm.OpenDocument("file://reexport.html", reExportContent, 1)
	if doc2 == nil {
		t.Fatal("Failed to open document")
	}

	diagnostics2 := publishDiagnostics.AnalyzeTagNameDiagnosticsForTest(server, doc2)

	for _, diagnostic := range diagnostics2 {
		if contains(diagnostic.Message, "my-button") {
			t.Errorf("Re-export import should not have errors for my-button: %s", diagnostic.Message)
		}
	}
}

// createTestFiles creates the test files for the my-tabs scenario
func createTestFiles(tempDir string) error {
	// Create package.json
	packageJSON := `{
  "name": "@my/elements",
  "version": "1.0.0",
  "customElements": "custom-elements.json"
}`
	if err := os.WriteFile(filepath.Join(tempDir, "package.json"), []byte(packageJSON), 0644); err != nil {
		return err
	}

	// Create custom-elements.json manifest
	manifest := `{
  "schemaVersion": "2.0.0",
  "modules": [
    {
      "path": "my-tab/my-tab.js",
      "declarations": [
        {
          "kind": "class",
          "name": "MyTab",
          "customElement": true,
          "tagName": "my-tab"
        }
      ]
    },
    {
      "path": "my-tabs/my-tabs.js", 
      "declarations": [
        {
          "kind": "class",
          "name": "MyTabs",
          "customElement": true,
          "tagName": "my-tabs"
        }
      ]
    }
  ]
}`
	if err := os.WriteFile(filepath.Join(tempDir, "custom-elements.json"), []byte(manifest), 0644); err != nil {
		return err
	}

	// Create component directories
	if err := os.MkdirAll(filepath.Join(tempDir, "my-tab"), 0755); err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Join(tempDir, "my-tabs"), 0755); err != nil {
		return err
	}

	// Create my-tab.js (direct export)
	myTabContent := `export class MyTab extends HTMLElement {
  // Tab implementation
}

customElements.define('my-tab', MyTab);`
	if err := os.WriteFile(filepath.Join(tempDir, "my-tab", "my-tab.js"), []byte(myTabContent), 0644); err != nil {
		return err
	}

	// Create my-tabs.js (re-exports my-tab)
	myTabsContent := `export class MyTabs extends HTMLElement {
  // Tabs container implementation
}

// Re-export MyTab for convenience - this is the key functionality we're testing
export { MyTab } from '../my-tab/my-tab.js';

customElements.define('my-tabs', MyTabs);`
	if err := os.WriteFile(filepath.Join(tempDir, "my-tabs", "my-tabs.js"), []byte(myTabsContent), 0644); err != nil {
		return err
	}

	return nil
}

// createDirectVsReExportFiles creates test files for direct vs re-export comparison
func createDirectVsReExportFiles(tempDir string) error {
	// Create button-element.js (direct export)
	buttonContent := `export class MyButton extends HTMLElement {
  // Button implementation
}

customElements.define('my-button', MyButton);`
	if err := os.WriteFile(filepath.Join(tempDir, "button-element.js"), []byte(buttonContent), 0644); err != nil {
		return err
	}

	// Create index.js (re-exports button)
	indexContent := `// Re-export button element
export { MyButton } from './button-element.js';`
	if err := os.WriteFile(filepath.Join(tempDir, "index.js"), []byte(indexContent), 0644); err != nil {
		return err
	}

	return nil
}

// contains checks if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr ||
		(len(s) > len(substr) &&
			(s[:len(substr)] == substr ||
				s[len(s)-len(substr):] == substr ||
				containsSubstring(s, substr))))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
