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
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/methods/textDocument/publishDiagnostics"
	W "bennypowers.dev/cem/workspace"
)

func TestModuleGraphIntegration_DirectVsReExport(t *testing.T) {
	// Use the fixture directory with proper manifest
	fixtureDir, err := filepath.Abs(filepath.Join("..", "..", "..", "testdata", "integration", "direct-vs-reexport"))
	if err != nil {
		t.Fatalf("Failed to get absolute fixture path: %v", err)
	}

	// Create workspace and server
	workspace := W.NewFileSystemWorkspaceContext(fixtureDir)
	if err := workspace.Init(); err != nil {
		t.Fatalf("Failed to initialize workspace: %v", err)
	}

	server, err := lsp.NewServer(workspace, lsp.TransportStdio)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer func() { _ = server.Close() }()

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
	dm, err := document.NewDocumentManager()
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
