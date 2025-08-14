package hover_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/methods/textDocument/hover"
	"bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestHoverIntegrationWithDocumentChanges(t *testing.T) {
	fixturePath := filepath.Join("..", "..", "..", "test", "fixtures", "hover-integration")

	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	registry := lsp.NewRegistry()

	// Load the manifest manually for testing
	manifestPath := filepath.Join(fixturePath, "custom-elements.json")
	manifestBytes, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatalf("Failed to read manifest: %v", err)
	}

	var pkg M.Package
	err = json.Unmarshal(manifestBytes, &pkg)
	if err != nil {
		t.Fatalf("Failed to parse manifest: %v", err)
	}

	registry.AddManifest(&pkg)

	ctx := &testHoverContext{
		documentManager: dm,
		registry:        registry,
	}

	htmlFile := filepath.Join(fixturePath, "index.html")
	uri := "file://" + htmlFile

	initialContent := `<!DOCTYPE html>
<html>
<head>
  <title>Test Page</title>
</head>
<body>
  <test-element test-attr="hello"></test-element>
</body>
</html>`

	doc := dm.OpenDocument(uri, initialContent, 1)
	if doc == nil {
		t.Fatal("Failed to open document")
	}

	t.Run("hover works on original tag name", func(t *testing.T) {
		pos := protocol.Position{Line: 6, Character: 5}
		params := &protocol.HoverParams{
			TextDocumentPositionParams: protocol.TextDocumentPositionParams{
				TextDocument: protocol.TextDocumentIdentifier{URI: uri},
				Position:     pos,
			},
		}

		result, err := hover.Hover(ctx, nil, params)
		if err != nil {
			t.Fatalf("Hover failed: %v", err)
		}

		if result == nil {
			t.Fatal("Expected hover result, got nil")
		}

		content := result.Contents.(protocol.MarkupContent)
		if content.Value == "" {
			t.Error("Expected hover content, got empty string")
		}

		if !contains(content.Value, "test-element") {
			t.Errorf("Expected hover content to contain 'test-element', got: %s", content.Value)
		}
	})

	t.Run("hover still works after adding new line above", func(t *testing.T) {
		updatedContent := `<!DOCTYPE html>
<html>
<head>
  <title>Test Page</title>
</head>
<body>

  <test-element test-attr="hello"></test-element>
</body>
</html>`

		updatedDoc := dm.UpdateDocument(uri, updatedContent, 2)
		if updatedDoc == nil {
			t.Fatal("Failed to update document")
		}

		pos := protocol.Position{Line: 7, Character: 5}
		params := &protocol.HoverParams{
			TextDocumentPositionParams: protocol.TextDocumentPositionParams{
				TextDocument: protocol.TextDocumentIdentifier{URI: uri},
				Position:     pos,
			},
		}

		result, err := hover.Hover(ctx, nil, params)
		if err != nil {
			t.Fatalf("Hover failed after line addition: %v", err)
		}

		if result == nil {
			t.Fatal("Expected hover result after line addition, got nil")
		}

		content := result.Contents.(protocol.MarkupContent)
		if content.Value == "" {
			t.Error("Expected hover content after line addition, got empty string")
		}

		if !contains(content.Value, "test-element") {
			t.Errorf("Expected hover content to contain 'test-element' after line addition, got: %s", content.Value)
		}
	})

	t.Run("hover works after multiple edits", func(t *testing.T) {
		finalContent := `<!DOCTYPE html>
<html>
<head>
  <title>Test Page</title>
</head>
<body>
  <!-- Added comment -->

  <test-element test-attr="hello"></test-element>
  <!-- Another comment -->
</body>
</html>`

		finalDoc := dm.UpdateDocument(uri, finalContent, 3)
		if finalDoc == nil {
			t.Fatal("Failed to update document with final content")
		}

		pos := protocol.Position{Line: 8, Character: 5}
		params := &protocol.HoverParams{
			TextDocumentPositionParams: protocol.TextDocumentPositionParams{
				TextDocument: protocol.TextDocumentIdentifier{URI: uri},
				Position:     pos,
			},
		}

		result, err := hover.Hover(ctx, nil, params)
		if err != nil {
			t.Fatalf("Hover failed after multiple edits: %v", err)
		}

		if result == nil {
			t.Fatal("Expected hover result after multiple edits, got nil")
		}

		content := result.Contents.(protocol.MarkupContent)
		if content.Value == "" {
			t.Error("Expected hover content after multiple edits, got empty string")
		}

		if !contains(content.Value, "test-element") {
			t.Errorf("Expected hover content to contain 'test-element' after multiple edits, got: %s", content.Value)
		}

		if !contains(content.Value, "Custom Element") {
			t.Errorf("Expected hover content to contain 'Custom Element' after multiple edits, got: %s", content.Value)
		}
	})

	t.Run("hover on attribute works after edits", func(t *testing.T) {
		pos := protocol.Position{Line: 8, Character: 25}
		params := &protocol.HoverParams{
			TextDocumentPositionParams: protocol.TextDocumentPositionParams{
				TextDocument: protocol.TextDocumentIdentifier{URI: uri},
				Position:     pos,
			},
		}

		result, err := hover.Hover(ctx, nil, params)
		if err != nil {
			t.Fatalf("Hover on attribute failed: %v", err)
		}

		if result == nil {
			t.Fatal("Expected hover result on attribute, got nil")
		}

		content := result.Contents.(protocol.MarkupContent)
		if content.Value == "" {
			t.Error("Expected hover content on attribute, got empty string")
		}

		if !contains(content.Value, "test-attr") {
			t.Errorf("Expected hover content to contain 'test-attr', got: %s", content.Value)
		}
	})
}

func TestHoverWithTypeScriptTemplates(t *testing.T) {
	fixturePath := filepath.Join("test", "fixtures", "typescript-templates")

	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create document manager: %v", err)
	}
	defer dm.Close()

	registry := lsp.NewRegistry()

	// Load the manifest manually for testing
	manifestPath := filepath.Join(fixturePath, "manifest.json")
	manifestBytes, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatalf("Failed to read manifest: %v", err)
	}

	var pkg M.Package
	err = json.Unmarshal(manifestBytes, &pkg)
	if err != nil {
		t.Fatalf("Failed to parse manifest: %v", err)
	}

	registry.AddManifest(&pkg)

	ctx := &testHoverContext{
		documentManager: dm,
		registry:        registry,
	}

	// Read the TypeScript test file content
	tsFilePath := filepath.Join(fixturePath, "test.ts")
	tsContentBytes, err := os.ReadFile(tsFilePath)
	if err != nil {
		t.Fatalf("Failed to read TypeScript test file: %v", err)
	}
	initialContent := string(tsContentBytes)

	uri := "file:///test.ts"

	doc := dm.OpenDocument(uri, initialContent, 1)
	if doc == nil {
		t.Fatal("Failed to open TypeScript document")
	}

	t.Run("hover works on tag in template literal", func(t *testing.T) {
		pos := protocol.Position{Line: 5, Character: 10}
		params := &protocol.HoverParams{
			TextDocumentPositionParams: protocol.TextDocumentPositionParams{
				TextDocument: protocol.TextDocumentIdentifier{URI: uri},
				Position:     pos,
			},
		}

		result, err := hover.Hover(ctx, nil, params)
		if err != nil {
			t.Fatalf("Hover failed in TypeScript template: %v", err)
		}

		if result == nil {
			t.Fatal("Expected hover result in TypeScript template, got nil")
		}

		content := result.Contents.(protocol.MarkupContent)
		if !contains(content.Value, "test-component") {
			t.Errorf("Expected hover content to contain 'test-component' in TypeScript template, got: %s", content.Value)
		}
	})

	t.Run("hover still works after adding new line in template", func(t *testing.T) {
		updatedContent := "import { html } from 'lit';\n\nexport class MyComponent {\n  render() {\n    return html`\n\n      <test-component prop=\"value\"></test-component>\n    `;\n  }\n}"

		updatedDoc := dm.UpdateDocument(uri, updatedContent, 2)
		if updatedDoc == nil {
			t.Fatal("Failed to update TypeScript document")
		}

		pos := protocol.Position{Line: 6, Character: 10}
		params := &protocol.HoverParams{
			TextDocumentPositionParams: protocol.TextDocumentPositionParams{
				TextDocument: protocol.TextDocumentIdentifier{URI: uri},
				Position:     pos,
			},
		}

		result, err := hover.Hover(ctx, nil, params)
		if err != nil {
			t.Fatalf("Hover failed after adding line in TypeScript template: %v", err)
		}

		if result == nil {
			t.Fatal("Expected hover result after adding line in TypeScript template, got nil")
		}

		content := result.Contents.(protocol.MarkupContent)
		if !contains(content.Value, "test-component") {
			t.Errorf("Expected hover content to contain 'test-component' after line addition in TypeScript template, got: %s", content.Value)
		}
	})
}

// testHoverContext implements the HoverContext interface for testing
type testHoverContext struct {
	documentManager *lsp.DocumentManager
	registry        *lsp.Registry
}

func (ctx *testHoverContext) Document(uri string) types.Document {
	doc := ctx.documentManager.Document(uri)
	if doc == nil {
		return nil
	}
	return doc
}

func (ctx *testHoverContext) Element(tagName string) (*M.CustomElement, bool) {
	element, exists := ctx.registry.Elements[tagName]
	return element, exists
}

func (ctx *testHoverContext) Attributes(tagName string) (map[string]*M.Attribute, bool) {
	attrs, exists := ctx.registry.Attributes(tagName)
	return attrs, exists
}

func (ctx *testHoverContext) Slots(tagName string) ([]M.Slot, bool) {
	return ctx.registry.Slots(tagName)
}

func (ctx *testHoverContext) RawDocumentManager() any {
	return ctx.documentManager
}

// documentAdapter adapts lsp.Document to textDocument.Document interface
type documentAdapter struct {
	doc *lsp.Document
	dm  *lsp.DocumentManager
}

func (d *documentAdapter) FindElementAtPosition(position protocol.Position, dm any) *types.CustomElementMatch {
	element := d.doc.FindElementAtPosition(position, d.dm)
	if element == nil {
		return nil
	}
	// Already returns types.CustomElementMatch, no conversion needed
	return element
}

func (d *documentAdapter) FindAttributeAtPosition(position protocol.Position, dm any) (*types.AttributeMatch, string) {
	attr, tagName := d.doc.FindAttributeAtPosition(position, d.dm)
	if attr == nil {
		return nil, tagName
	}
	// Already returns types.AttributeMatch, no conversion needed
	return attr, tagName
}

// Content returns the document content
func (d *documentAdapter) Content() string {
	return d.doc.Content()
}

// Version returns the document version
func (d *documentAdapter) Version() int32 {
	return d.doc.Version()
}

// URI returns the document URI
func (d *documentAdapter) URI() string {
	return d.doc.URI()
}

// FindCustomElements returns custom elements
func (d *documentAdapter) FindCustomElements(dm any) ([]types.CustomElementMatch, error) {
	return d.doc.FindCustomElements(dm)
}

// AnalyzeCompletionContextTS returns completion analysis
func (d *documentAdapter) AnalyzeCompletionContextTS(position protocol.Position, dm any) *types.CompletionAnalysis {
	return d.doc.AnalyzeCompletionContextTS(position, dm)
}

func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
