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
	"strings"
	"testing"

	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/methods/textDocument/publishDiagnostics"
	W "bennypowers.dev/cem/workspace"
)

func TestTransitiveIntegration_ThreeLevelDependency(t *testing.T) {
	// NOTE: The issue was that the test wasn't triggering the textDocument/didOpen
	// lifecycle which causes lazy module graph building. Fixed by properly opening
	// the document through the server.

	// Disable debug logging during this test to avoid race conditions with pterm global state
	helpers.SetDebugLoggingEnabled(false)
	defer helpers.SetDebugLoggingEnabled(false)

	// Use the fixture directory with predefined transitive dependency chain
	fixtureDir, err := filepath.Abs(filepath.Join("..", "..", "..", "test", "fixtures", "transitive-integration"))
	if err != nil {
		t.Fatalf("Failed to get absolute fixture path: %v", err)
	}
	t.Logf("DEBUG: Using fixture directory: %s", fixtureDir)

	// Create workspace context
	workspace := W.NewFileSystemWorkspaceContext(fixtureDir)
	err = workspace.Init()
	if err != nil {
		t.Fatalf("Failed to initialize workspace: %v", err)
	}

	// Create server and initialize with module graph support
	server, err := lsp.NewServer(workspace, lsp.TransportStdio)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer func() { _ = server.Close() }()

	t.Logf("DEBUG: About to call InitializeForTesting")
	err = server.InitializeForTesting()
	if err != nil {
		t.Fatalf("Failed to initialize server: %v", err)
	}
	t.Logf("DEBUG: InitializeForTesting completed")

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

	// Simulate the DidOpen lifecycle without the LSP notification part
	// This triggers the document processing and lazy module graph building

	// Get the document manager
	dm, err := server.DocumentManager()
	if err != nil {
		t.Fatalf("Failed to get document manager: %v", err)
	}

	// Open the document (this is what DidOpen does internally)
	doc := dm.OpenDocument("file:///test.html", htmlContent, 1)
	if doc == nil {
		t.Fatal("Failed to open document")
	}

	// Trigger the diagnostics processing pipeline which includes lazy module graph building
	// This is the key part that should cause BuildForImportPath to be called for './my-tabs.js'
	// We call the internal diagnostics analysis which triggers the lazy building

	// Debug: Check if server has module graph access and workspace root
	moduleGraph := server.ModuleGraph()
	if moduleGraph != nil {
		t.Logf("DEBUG: Server has module graph access")
		allModules := moduleGraph.GetAllModulePaths()
		t.Logf("DEBUG: Module graph contains %d modules before diagnostics: %v", len(allModules), allModules)

		// Check workspace root - this is critical for TypeScript file discovery
		workspaceRoot := server.Workspace().Root()
		t.Logf("DEBUG: Workspace root: %s", workspaceRoot)
		t.Logf("DEBUG: Fixture directory: %s", fixtureDir)

		// Check if TypeScript files exist at expected paths
		for _, module := range []string{"my-tabs.js", "my-tab.js", "my-icon.js"} {
			tsPath := strings.TrimSuffix(module, ".js") + ".ts"
			fullTsPath := filepath.Join(workspaceRoot, tsPath)
			t.Logf("DEBUG: Checking TypeScript file: %s (exists: %t)", fullTsPath, fileExists(fullTsPath))
		}
	} else {
		t.Logf("DEBUG: Server has NO module graph access!")
	}

	// Debug: Test the resolveImportPathToElements function manually
	t.Logf("DEBUG: Testing manual import resolution for './my-tabs.js'")

	diagnostics := publishDiagnostics.AnalyzeTagNameDiagnosticsForTest(server, doc)

	// Debug: Check module graph after diagnostics
	if moduleGraph != nil {
		allModules := moduleGraph.GetAllModulePaths()
		t.Logf("DEBUG: Module graph contains %d modules after diagnostics: %v", len(allModules), allModules)

		// Test BuildForImportPath manually to see if it works
		err := moduleGraph.BuildForImportPath("./my-tabs.js")
		if err != nil {
			t.Logf("DEBUG: BuildForImportPath('./my-tabs.js') failed: %v", err)
		} else {
			t.Logf("DEBUG: BuildForImportPath('./my-tabs.js') succeeded")
		}

		// Test building dependencies for each module separately to see if that fixes it
		err = moduleGraph.BuildForImportPath("./my-tab.js")
		if err != nil {
			t.Logf("DEBUG: BuildForImportPath('./my-tab.js') failed: %v", err)
		} else {
			t.Logf("DEBUG: BuildForImportPath('./my-tab.js') succeeded")
		}

		err = moduleGraph.BuildForImportPath("./my-icon.js")
		if err != nil {
			t.Logf("DEBUG: BuildForImportPath('./my-icon.js') failed: %v", err)
		} else {
			t.Logf("DEBUG: BuildForImportPath('./my-icon.js') succeeded")
		}

		// Test transitive resolution for the key case
		t.Logf("DEBUG: About to call GetTransitiveElements('my-tabs.js') - this should show BFS debug logs")
		transitiveElements := moduleGraph.GetTransitiveElements("my-tabs.js")
		t.Logf("DEBUG: GetTransitiveElements('my-tabs.js') returns: %v", transitiveElements)

		// Test the caching theory - create a fresh server and module graph
		t.Logf("DEBUG: Testing with fresh server to check caching theory")
		freshServer, err := lsp.NewServer(workspace, lsp.TransportStdio)
		if err != nil {
			t.Fatalf("Failed to create fresh server: %v", err)
		}
		defer func() { _ = freshServer.Close() }()

		err = freshServer.InitializeForTesting()
		if err != nil {
			t.Fatalf("Failed to initialize fresh server: %v", err)
		}

		freshModuleGraph := freshServer.ModuleGraph()
		if freshModuleGraph != nil {
			// Build dependencies on fresh graph
			err = freshModuleGraph.BuildForImportPath("./my-tabs.js")
			if err != nil {
				t.Logf("DEBUG: Fresh BuildForImportPath('./my-tabs.js') failed: %v", err)
			} else {
				t.Logf("DEBUG: Fresh BuildForImportPath('./my-tabs.js') succeeded")
			}

			// Test transitive resolution on fresh graph
			freshTransitive := freshModuleGraph.GetTransitiveElements("my-tabs.js")
			t.Logf("DEBUG: Fresh GetTransitiveElements('my-tabs.js') returns: %v", freshTransitive)
		}

		// Check what dependencies the dependency tracker knows about
		myTabsDeps := moduleGraph.GetModuleDependencies("my-tabs.js")
		myTabDeps := moduleGraph.GetModuleDependencies("my-tab.js")
		myIconDeps := moduleGraph.GetModuleDependencies("my-icon.js")
		t.Logf("DEBUG: Dependencies - my-tabs.js: %v", myTabsDeps)
		t.Logf("DEBUG: Dependencies - my-tab.js: %v", myTabDeps)
		t.Logf("DEBUG: Dependencies - my-icon.js: %v", myIconDeps)

		// Test direct module element lookup
		for _, module := range []string{"my-tabs.js", "my-tab.js", "my-icon.js"} {
			elements := moduleGraph.GetElementSources(module)
			t.Logf("DEBUG: GetElementSources('%s') returns: %v", module, elements)
		}

		// Check what module paths are stored in registry element definitions
		registry := server.Registry()
		if registry != nil {
			t.Logf("DEBUG: Registry element definitions:")
			for tagName := range registry.ElementDefinitions {
				elementDef, exists := registry.ElementDefinition(tagName)
				if exists {
					modulePath := elementDef.GetModulePath()
					t.Logf("DEBUG: Element '%s' -> module path '%s'", tagName, modulePath)
				}
			}

			// Test the registry manifest resolver directly
			manifestResolver := lsp.NewRegistryManifestResolver(registry)
			t.Logf("DEBUG: Testing manifest resolver GetElementsFromManifestModule:")
			for _, module := range []string{"my-tabs.js", "my-tab.js", "my-icon.js"} {
				elements := manifestResolver.GetElementsFromManifestModule(module)
				t.Logf("DEBUG: GetElementsFromManifestModule('%s') returns: %v", module, elements)
			}

			// Test manifest-based import path resolution (this might be what RHDS uses)
			t.Logf("DEBUG: Testing manifest-based import path resolution:")
			manifestModules := manifestResolver.FindManifestModulesForImportPath("./my-tabs.js")
			t.Logf("DEBUG: FindManifestModulesForImportPath('./my-tabs.js') returns: %v", manifestModules)

			// Manual simulation removed to see production debug logs
		}
	}

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
	// Use the fixture directory with predefined transitive dependency chain
	fixtureDir := filepath.Join("..", "..", "..", "test", "fixtures", "transitive-integration")

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

	dm, err := document.NewDocumentManager()
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

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

