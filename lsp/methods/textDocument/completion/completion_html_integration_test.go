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
	"encoding/json"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/methods/textDocument"
	"bennypowers.dev/cem/lsp/methods/textDocument/completion"
	"bennypowers.dev/cem/lsp/testhelpers"
	M "bennypowers.dev/cem/manifest"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// TestHTMLFileCompletion tests tag name completion in actual HTML files (not TypeScript templates)
// This is the critical test case that should expose the issue
func TestHTMLFileCompletion(t *testing.T) {
	// Load test manifest
	fs := testutil.NewFixtureFS(t, "slot-completions-test", "/test")
	manifestBytes, err := fs.ReadFile("/test/manifest.json")
	if err != nil {
		t.Fatalf("Failed to read test manifest: %v", err)
	}

	var pkg M.Package
	err = json.Unmarshal(manifestBytes, &pkg)
	if err != nil {
		t.Fatalf("Failed to parse manifest: %v", err)
	}

	// Create context and add the test manifest
	ctx := testhelpers.NewMockServerContext()
	ctx.AddManifest(&pkg)

	// Debug: Check what tags were loaded
	allTags := ctx.AllTagNames()
	t.Logf("Loaded %d tags: %v", len(allTags), allTags)

	// Create and set a real DocumentManager
	dm, err := document.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	t.Run("HTML File - Minimal Tag", func(t *testing.T) {
		// Test minimal HTML with just a < character
		htmlContent := "<"
		htmlURI := "file:///test/test.html"

		// Open HTML document
		doc := dm.OpenDocument(htmlURI, htmlContent, 1)
		if doc == nil {
			t.Fatal("Failed to open HTML document")
		}
		t.Logf("Opened HTML document: %s, Language: %s", htmlURI, doc.Language())

		// Add document to context
		ctx.AddDocument(htmlURI, doc)

		// Create completion params
		params := &protocol.CompletionParams{
			TextDocumentPositionParams: protocol.TextDocumentPositionParams{
				TextDocument: protocol.TextDocumentIdentifier{URI: htmlURI},
				Position:     protocol.Position{Line: 0, Character: 1}, // After the <
			},
			Context: &protocol.CompletionContext{
				TriggerCharacter: stringPtr("<"),
			},
		}

		// Request completions
		t.Logf("Requesting completions at position (0, 1) after '<'")
		result, err := completion.Completion(ctx, &glsp.Context{}, params)
		if err != nil {
			t.Fatalf("Completion failed: %v", err)
		}

		items, ok := result.([]protocol.CompletionItem)
		if !ok {
			t.Fatalf("Expected []CompletionItem, got %T", result)
		}

		t.Logf("Received %d completion items", len(items))
		for i, item := range items {
			t.Logf("  [%d] %s", i, item.Label)
		}

		// Verify we got custom element completions
		if len(items) == 0 {
			t.Error("❌ FAILED: Expected tag name completions, got none")
			t.Error("This confirms the bug - HTML file completions are broken!")
		} else {
			expectedTags := []string{"card-element", "dialog-element", "my-custom-element"}
			foundTags := make(map[string]bool)
			for _, item := range items {
				foundTags[item.Label] = true
			}

			for _, expectedTag := range expectedTags {
				if foundTags[expectedTag] {
					t.Logf("✅ Found expected tag: %s", expectedTag)
				} else {
					t.Errorf("❌ Missing expected tag: %s", expectedTag)
				}
			}
		}
	})

	t.Run("HTML File - Full Document", func(t *testing.T) {
		// Test with a full HTML document structure
		htmlContent := `<!DOCTYPE html>
<html>
<head>
  <title>Test</title>
</head>
<body>
  <
</body>
</html>`

		htmlURI := "file:///test/full-test.html"

		// Debug: Print the content with line numbers
		t.Log("HTML Content:")
		lines := []string{
			"0: <!DOCTYPE html>",
			"1: <html>",
			"2: <head>",
			"3:   <title>Test</title>",
			"4: </head>",
			"5: <body>",
			"6:   <",
			"7: </body>",
			"8: </html>",
		}
		for _, line := range lines {
			t.Log(line)
		}
		t.Log("Cursor position: Line 6, Character 3 (after the '<')")

		// Open HTML document
		doc := dm.OpenDocument(htmlURI, htmlContent, 1)
		if doc == nil {
			t.Fatal("Failed to open HTML document")
		}

		// Add document to context
		ctx.AddDocument(htmlURI, doc)

		// Create completion params (cursor after < in body)
		params := &protocol.CompletionParams{
			TextDocumentPositionParams: protocol.TextDocumentPositionParams{
				TextDocument: protocol.TextDocumentIdentifier{URI: htmlURI},
				Position:     protocol.Position{Line: 6, Character: 3}, // After the < in body
			},
			Context: &protocol.CompletionContext{
				TriggerCharacter: stringPtr("<"),
			},
		}

		// Debug: Check what context analysis returns
		dm2, _ := ctx.DocumentManager()
		analysis, err := textDocument.AnalyzeCompletionContextWithDM(doc, params.Position, "<", dm2)
		if err != nil {
			t.Logf("Analysis error: %v", err)
		} else {
			t.Logf("Analysis result: Type=%d, TagName=%q, AttributeName=%q",
				analysis.Type, analysis.TagName, analysis.AttributeName)
			// Check what prefix is extracted
			prefix := doc.CompletionPrefix(analysis)
			t.Logf("Completion prefix: %q", prefix)
		}

		// Request completions
		t.Log("Requesting completions...")
		result, err := completion.Completion(ctx, &glsp.Context{}, params)
		if err != nil {
			t.Fatalf("Completion failed: %v", err)
		}

		items, ok := result.([]protocol.CompletionItem)
		if !ok {
			t.Fatalf("Expected []CompletionItem, got %T", result)
		}

		t.Logf("Full HTML document: Received %d completion items", len(items))

		// Verify we got completions
		if len(items) < 3 {
			t.Errorf("❌ FAILED: Expected at least 3 completions, got %d", len(items))
			t.Error("This is the critical bug - full HTML documents don't get completions!")
			t.Error("Enable debug logging with -vvv to see what's happening in the completion analysis")
		} else {
			t.Logf("✅ SUCCESS: Got %d completions in full HTML document", len(items))
		}
	})

	t.Run("HTML File - Partial Tag Name", func(t *testing.T) {
		// Test with partial tag name
		htmlContent := "<my-c"
		htmlURI := "file:///test/partial-test.html"

		// Open HTML document
		doc := dm.OpenDocument(htmlURI, htmlContent, 1)
		if doc == nil {
			t.Fatal("Failed to open HTML document")
		}

		// Add document to context
		ctx.AddDocument(htmlURI, doc)

		// Create completion params
		params := &protocol.CompletionParams{
			TextDocumentPositionParams: protocol.TextDocumentPositionParams{
				TextDocument: protocol.TextDocumentIdentifier{URI: htmlURI},
				Position:     protocol.Position{Line: 0, Character: 5}, // After "my-c"
			},
		}

		// Request completions
		result, err := completion.Completion(ctx, &glsp.Context{}, params)
		if err != nil {
			t.Fatalf("Completion failed: %v", err)
		}

		items, ok := result.([]protocol.CompletionItem)
		if !ok {
			t.Fatalf("Expected []CompletionItem, got %T", result)
		}

		t.Logf("Partial tag name: Received %d completion items", len(items))

		// Should get only my-custom-element
		foundMyCustomElement := false
		for _, item := range items {
			if item.Label == "my-custom-element" {
				foundMyCustomElement = true
			}
		}

		if !foundMyCustomElement {
			t.Error("Expected to find 'my-custom-element' in completions for partial tag 'my-c'")
		}
	})
}
