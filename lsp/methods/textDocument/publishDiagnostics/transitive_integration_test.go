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

func TestTransitiveIntegration_ThreeLevelDependency(t *testing.T) {
	// Create a temporary workspace with transitive dependency chain
	tempDir, err := os.MkdirTemp("", "transitive-integration-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test files demonstrating: my-tabs -> my-tab -> my-icon
	if err := createTransitiveTestWorkspace(tempDir); err != nil {
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

	// Test HTML content that imports my-tabs.js and uses all three elements
	htmlContent := `<!DOCTYPE html>
<html>
<head>
    <script type="module">
        import './my-tabs.js';
    </script>
</head>
<body>
    <my-tabs>
        <my-tab slot="tab">
            Tab 1 <my-icon name="star"></my-icon>
        </my-tab>
        <my-tab slot="tab">
            Tab 2 <my-icon name="heart"></my-icon>
        </my-tab>
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

	// With transitive dependency resolution, ALL elements should be available
	// None of my-tabs, my-tab, or my-icon should have import errors
	problematicElements := []string{"my-tabs", "my-tab", "my-icon"}

	for _, diagnostic := range diagnostics {
		for _, element := range problematicElements {
			if contains(diagnostic.Message, element) &&
				(contains(diagnostic.Message, "not imported") || contains(diagnostic.Message, "Unknown custom element")) {
				t.Errorf("Transitive dependency failed: %s should be available via import './my-tabs.js' but got error: %s", element, diagnostic.Message)
			}
		}
	}

	// Log all diagnostics for debugging
	if len(diagnostics) > 0 {
		t.Logf("Found %d diagnostics:", len(diagnostics))
		for i, diagnostic := range diagnostics {
			t.Logf("  [%d]: %s", i+1, diagnostic.Message)
		}
	} else {
		t.Log("No diagnostics found - transitive dependencies working correctly!")
	}
}

func TestTransitiveIntegration_PartialImport(t *testing.T) {
	// Create a temporary workspace
	tempDir, err := os.MkdirTemp("", "partial-import-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test files
	if err := createTransitiveTestWorkspace(tempDir); err != nil {
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

	// Test importing only my-tab (which should bring in my-icon transitively)
	htmlContent := `<!DOCTYPE html>
<html>
<head>
    <script type="module">
        import './my-tab.js';
    </script>
</head>
<body>
    <my-tab>
        Tab content <my-icon name="star"></my-icon>
    </my-tab>
    <my-tabs>This should error since my-tabs.js is not imported</my-tabs>
</body>
</html>`

	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	doc := dm.OpenDocument("file://test.html", htmlContent, 1)
	if doc == nil {
		t.Fatal("Failed to open document")
	}

	diagnostics := publishDiagnostics.AnalyzeTagNameDiagnosticsForTest(server, doc)

	// my-tab and my-icon should be available (no errors)
	hasMyTabsError := false

	for _, diagnostic := range diagnostics {
		// Check for my-tabs first (longer string) to avoid substring matching issues
		if contains(diagnostic.Message, "my-tabs") &&
			(contains(diagnostic.Message, "not imported") || contains(diagnostic.Message, "Unknown custom element")) {
			hasMyTabsError = true
			t.Log("Expected error for my-tabs - it was not imported")
		} else if contains(diagnostic.Message, "'my-tab'") &&
			(contains(diagnostic.Message, "not imported") || contains(diagnostic.Message, "Unknown custom element")) {
			t.Errorf("my-tab should be available via direct import: %s", diagnostic.Message)
		} else if contains(diagnostic.Message, "'my-icon'") &&
			(contains(diagnostic.Message, "not imported") || contains(diagnostic.Message, "Unknown custom element")) {
			t.Errorf("my-icon should be available via transitive import from my-tab: %s", diagnostic.Message)
		}
	}

	if !hasMyTabsError {
		t.Error("Expected my-tabs to have an import error since it wasn't imported")
	}
}

// createTransitiveTestWorkspace creates a realistic workspace with transitive dependencies
func createTransitiveTestWorkspace(tempDir string) error {
	// my-icon.js (base dependency)
	iconContent := `
export class MyIcon extends HTMLElement {
  connectedCallback() {
    this.innerHTML = '<svg><!-- icon --></svg>';
  }
}

customElements.define('my-icon', MyIcon);
`

	// my-tab.js (imports my-icon)
	tabContent := `
import './my-icon.js';

export class MyTab extends HTMLElement {
  connectedCallback() {
    this.innerHTML = '<div class="tab">Tab content</div>';
  }
}

customElements.define('my-tab', MyTab);
`

	// my-tabs.js (imports my-tab, which transitively imports my-icon)
	tabsContent := `
import './my-tab.js';

export class MyTabs extends HTMLElement {
  connectedCallback() {
    this.innerHTML = '<div class="tabs">Tabs container</div>';
  }
}

customElements.define('my-tabs', MyTabs);
`

	// Write test files
	files := map[string]string{
		"my-icon.js": iconContent,
		"my-tab.js":  tabContent,
		"my-tabs.js": tabsContent,
	}

	for filename, content := range files {
		if err := os.WriteFile(filepath.Join(tempDir, filename), []byte(content), 0644); err != nil {
			return err
		}
	}

	return nil
}
