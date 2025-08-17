package publishDiagnostics_test

import (
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/methods/textDocument/publishDiagnostics"
	"bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
	W "bennypowers.dev/cem/workspace"
)

// testAdapter implements the context interfaces needed for diagnostics testing
type testAdapter struct {
	registry *lsp.Registry
	docMgr   *lsp.DocumentManager
}

func (t *testAdapter) Document(uri string) types.Document {
	return t.docMgr.Document(uri)
}

func (t *testAdapter) AllTagNames() []string {
	return t.registry.AllTagNames()
}

func (t *testAdapter) ElementDefinition(tagName string) (types.ElementDefinition, bool) {
	return t.registry.ElementDefinition(tagName)
}

func (t *testAdapter) ElementSource(tagName string) (string, bool) {
	if def, exists := t.registry.ElementDefinition(tagName); exists {
		// Try package name first, then fall back to module path
		if packageName := def.PackageName(); packageName != "" {
			return packageName + "/" + def.ModulePath(), true
		}
		return def.ModulePath(), true
	}
	return "", false
}

func (t *testAdapter) Slots(tagName string) ([]M.Slot, bool) {
	return t.registry.Slots(tagName)
}

func (t *testAdapter) Attributes(tagName string) (map[string]*M.Attribute, bool) {
	return t.registry.Attributes(tagName)
}

func TestTagDiagnostics_WithImports(t *testing.T) {
	// Setup fixture workspace
	fixtureDir, err := filepath.Abs("test-fixtures/tag-diagnostics")
	if err != nil {
		t.Fatalf("Failed to get fixture path: %v", err)
	}

	// Create workspace context
	workspace := W.NewFileSystemWorkspaceContext(fixtureDir)
	err = workspace.Init()
	if err != nil {
		t.Fatalf("Failed to initialize workspace: %v", err)
	}

	// Create server (this loads manifests automatically)
	server, err := lsp.NewServer(workspace)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	// Initialize the server
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

	// Create adapter (simplified context for diagnostics)
	adapter := &testAdapter{
		registry: server.Registry(),
		docMgr:   dm,
	}

	// Read the HTML file content from disk
	htmlPath := filepath.Join(fixtureDir, "with-imports.html")
	content, err := os.ReadFile(htmlPath)
	if err != nil {
		t.Fatalf("Failed to read HTML file: %v", err)
	}

	// Open document
	uri := "file://" + htmlPath
	doc := dm.OpenDocument(uri, string(content), 1)
	if doc == nil {
		t.Fatal("Failed to open document")
	}

	// Analyze tag diagnostics
	diagnostics := publishDiagnostics.AnalyzeTagNameDiagnosticsForTest(adapter, doc)

	// Debug: check element sources
	allTagNames := adapter.AllTagNames()
	t.Logf("Available elements: %v", allTagNames)
	for _, tagName := range allTagNames {
		if source, hasSource := adapter.ElementSource(tagName); hasSource {
			t.Logf("Element '%s' has source: '%s'", tagName, source)
		}
	}

	// Should have 0 diagnostics when imports are present
	// Debug: Check if script tags are being parsed
	if doc != nil {
		scriptTags := doc.GetScriptTags()
		t.Logf("Document has %d script tags parsed", len(scriptTags))
		for i, tag := range scriptTags {
			t.Logf("Script %d: IsModule=%t, %d imports", i, tag.IsModule, len(tag.Imports))
			for j, imp := range tag.Imports {
				t.Logf("  Import %d: %s (%s)", j, imp.ImportPath, imp.Type)
			}
		}
	}

	if len(diagnostics) != 0 {
		t.Errorf("Expected 0 diagnostics when imports are present, got %d", len(diagnostics))
		for i, diag := range diagnostics {
			t.Errorf("Unexpected diagnostic %d: %s", i, diag.Message)
		}
	}
}

func TestTagDiagnostics_TypeScriptImports(t *testing.T) {
	// Setup fixture workspace
	fixtureDir, err := filepath.Abs("test-fixtures/tag-diagnostics")
	if err != nil {
		t.Fatalf("Failed to get fixture path: %v", err)
	}

	// Create workspace context
	workspace := W.NewFileSystemWorkspaceContext(fixtureDir)
	err = workspace.Init()
	if err != nil {
		t.Fatalf("Failed to initialize workspace: %v", err)
	}

	// Create server (this loads manifests automatically)
	server, err := lsp.NewServer(workspace)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	// Initialize the server
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

	// Create adapter
	adapter := &testAdapter{
		registry: server.Registry(),
		docMgr:   dm,
	}

	// Read the TypeScript file content from disk
	tsPath := filepath.Join(fixtureDir, "typescript-imports.ts")
	content, err := os.ReadFile(tsPath)
	if err != nil {
		t.Fatalf("Failed to read TypeScript file: %v", err)
	}

	// Open document
	uri := "file://" + tsPath
	doc := dm.OpenDocument(uri, string(content), 1)
	if doc == nil {
		t.Fatal("Failed to open document")
	}

	// Debug: check element sources
	allTagNames := adapter.AllTagNames()
	t.Logf("Available elements: %v", allTagNames)
	for _, tagName := range allTagNames {
		if source, hasSource := adapter.ElementSource(tagName); hasSource {
			t.Logf("Element '%s' has source: '%s'", tagName, source)
		}
	}

	// Analyze tag diagnostics
	diagnostics := publishDiagnostics.AnalyzeTagNameDiagnosticsForTest(adapter, doc)

	// Should have 0 diagnostics for TypeScript file with imports
	if len(diagnostics) != 0 {
		t.Errorf("Expected 0 diagnostics for TypeScript file with imports, got %d", len(diagnostics))
		for i, diag := range diagnostics {
			t.Errorf("Diagnostic %d: %s", i, diag.Message)
		}
	}
}

func TestTagDiagnostics_MissingImports(t *testing.T) {
	// Setup fixture workspace
	fixtureDir, err := filepath.Abs("test-fixtures/tag-diagnostics")
	if err != nil {
		t.Fatalf("Failed to get fixture path: %v", err)
	}

	// Create workspace context
	workspace := W.NewFileSystemWorkspaceContext(fixtureDir)
	err = workspace.Init()
	if err != nil {
		t.Fatalf("Failed to initialize workspace: %v", err)
	}

	// Create server (this loads manifests automatically)
	server, err := lsp.NewServer(workspace)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	// Initialize the server
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

	// Create adapter
	adapter := &testAdapter{
		registry: server.Registry(),
		docMgr:   dm,
	}

	// Read the HTML file content from disk
	htmlPath := filepath.Join(fixtureDir, "missing-imports.html")
	content, err := os.ReadFile(htmlPath)
	if err != nil {
		t.Fatalf("Failed to read HTML file: %v", err)
	}

	// Open document
	uri := "file://" + htmlPath
	doc := dm.OpenDocument(uri, string(content), 1)
	if doc == nil {
		t.Fatal("Failed to open document")
	}

	// Analyze tag diagnostics
	diagnostics := publishDiagnostics.AnalyzeTagNameDiagnosticsForTest(adapter, doc)

	// Should have diagnostics for missing imports and typos
	if len(diagnostics) == 0 {
		t.Error("Expected diagnostics for missing imports, got none")
		return
	}

	// Log all diagnostics found
	foundMessages := make(map[string]bool)
	for i, diag := range diagnostics {
		t.Logf("Found diagnostic %d: %s", i, diag.Message)
		foundMessages[diag.Message] = true
	}

	// Verify we get appropriate error messages
	// Note: Exact messages will depend on manifest content and available elements
	if len(foundMessages) == 0 {
		t.Error("No diagnostic messages found")
	}
}

func TestTagDiagnostics_IgnoreComment(t *testing.T) {
	// Setup fixture workspace
	fixtureDir, err := filepath.Abs("test-fixtures/tag-diagnostics")
	if err != nil {
		t.Fatalf("Failed to get fixture path: %v", err)
	}

	// Create temporary file with ignore comment
	tempDir, err := os.MkdirTemp("", "tag-diagnostics-ignore-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Write file with ignore comment
	ignoreFile := filepath.Join(tempDir, "ignore-test.html")
	ignoreContent := `<!-- cem-lsp ignore missing-import -->
<unknown-element>Test</unknown-element>
<another-missing-element>Test</another-missing-element>`

	err = os.WriteFile(ignoreFile, []byte(ignoreContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write ignore test file: %v", err)
	}

	// Create workspace context (use fixture dir for manifests)
	workspace := W.NewFileSystemWorkspaceContext(fixtureDir)
	err = workspace.Init()
	if err != nil {
		t.Fatalf("Failed to initialize workspace: %v", err)
	}

	// Create server (this loads manifests automatically)
	server, err := lsp.NewServer(workspace)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	// Initialize the server
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

	// Create adapter
	adapter := &testAdapter{
		registry: server.Registry(),
		docMgr:   dm,
	}

	// Open document with ignore comment
	uri := "file://" + ignoreFile
	doc := dm.OpenDocument(uri, ignoreContent, 1)
	if doc == nil {
		t.Fatal("Failed to open document")
	}

	// Analyze tag diagnostics
	diagnostics := publishDiagnostics.AnalyzeTagNameDiagnosticsForTest(adapter, doc)

	// Should have 0 diagnostics when ignore comment is present
	if len(diagnostics) != 0 {
		t.Errorf("Expected 0 diagnostics with ignore comment, got %d", len(diagnostics))
		for i, diag := range diagnostics {
			t.Errorf("Unexpected diagnostic %d: %s", i, diag.Message)
		}
	}
}
