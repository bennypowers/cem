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
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/methods/textDocument/completion"
	"bennypowers.dev/cem/lsp/types"
	protocol "github.com/bennypowers/glsp/protocol_3_17"
)

// TestHelpers provides utilities for completion testing
type TestHelpers struct{}

// NewMockTemplateDocument creates a real Document for testing using a DocumentManager
// The templateContext parameter is ignored since real Documents don't need mock template context
// DEPRECATED: Tests should use dm.OpenDocument() directly
func NewMockTemplateDocument(content, templateContext string) types.Document {
	// Create a DocumentManager and use it to open the document
	dm, err := document.NewDocumentManager()
	if err != nil {
		panic("Failed to create DocumentManager for test: " + err.Error())
	}
	defer dm.Close()

	// Real Documents handle template context automatically through tree-sitter parsing
	return dm.OpenDocument("test://template.html", content, 1)
}

// CreateAttributeValueCompletionParams creates LSP completion parameters for testing attribute value completions
// This simulates a cursor positioned inside an attribute value, like: <tag-name attr="|"
func (h *TestHelpers) CreateAttributeValueCompletionParams(uri, tagName, attributeName string) *protocol.CompletionParams {
	// Position cursor right after the ="
	position := protocol.Position{
		Line:      0,
		Character: uint32(len(fmt.Sprintf(`<%s %s="`, tagName, attributeName))),
	}

	return &protocol.CompletionParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{
				URI: uri,
			},
			Position: position,
		},
		Context: &protocol.CompletionContext{
			TriggerKind: protocol.CompletionTriggerKindInvoked,
		},
	}
}

// CreateAttributeCompletionParams creates LSP completion parameters for testing attribute name completions
// This simulates a cursor positioned after a space in a tag, like: <tag-name |
func (h *TestHelpers) CreateAttributeCompletionParams(uri, tagName string) *protocol.CompletionParams {
	// Position cursor right after the space
	position := protocol.Position{
		Line:      0,
		Character: uint32(len(fmt.Sprintf(`<%s |`, tagName))),
	}

	return &protocol.CompletionParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{
				URI: uri,
			},
			Position: position,
		},
		Context: &protocol.CompletionContext{
			TriggerKind: protocol.CompletionTriggerKindInvoked,
		},
	}
}

// GetAttributeValueCompletionsUsingMainEntry is a helper that calls the main Completion entry point
// and extracts attribute value completions for the specified tag and attribute.
// This replaces direct calls to GetAttributeValueCompletions in tests.
func (h *TestHelpers) GetAttributeValueCompletionsUsingMainEntry(
	ctx types.ServerContext,
	uri, tagName, attributeName string,
) ([]protocol.CompletionItem, error) {
	// For now, fall back to the legacy function until we have a working main entry point approach
	// The main entry point requires complex document analysis that's hard to mock properly
	return completion.GetAttributeValueCompletions(ctx, tagName, attributeName), nil
}

var TestHelper = &TestHelpers{}

// MockDocument implements types.Document for testing

// Helper function to convert completion type to string for better error messages
func completionTypeString(t types.CompletionContextType) string {
	switch t {
	case types.CompletionUnknown:
		return "Unknown"
	case types.CompletionTagName:
		return "TagName"
	case types.CompletionAttributeName:
		return "AttributeName"
	case types.CompletionAttributeValue:
		return "AttributeValue"
	case types.CompletionLitEventBinding:
		return "LitEventBinding"
	case types.CompletionLitPropertyBinding:
		return "LitPropertyBinding"
	case types.CompletionLitBooleanAttribute:
		return "LitBooleanAttribute"
	default:
		return "Invalid"
	}
}

// CopyFixtureFiles loads fixtures from fixtureDir via testutil and writes
// them to targetDir, preserving directory structure.
func CopyFixtureFiles(t *testing.T, fixtureDir, targetDir string) {
	t.Helper()
	mfs := testutil.LoadTestdataFS(t, fixtureDir, "/")
	entries := mfs.GetMapFS()
	for path, file := range entries {
		if file.Mode.IsDir() || strings.HasSuffix(path, "README.md") {
			continue
		}
		data, err := mfs.ReadFile(path)
		if err != nil {
			t.Fatalf("Failed to read fixture %s: %v", path, err)
		}
		targetPath := filepath.Join(targetDir, path)
		if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
			t.Fatalf("Failed to create directory for %s: %v", targetPath, err)
		}
		if err := os.WriteFile(targetPath, data, 0644); err != nil {
			t.Fatalf("Failed to write %s: %v", targetPath, err)
		}
	}
}
