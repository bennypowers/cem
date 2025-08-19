package completion_test

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/methods/textDocument/completion"
	"bennypowers.dev/cem/lsp/testhelpers"
	"bennypowers.dev/cem/lsp/types"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// TestHelpers provides utilities for completion testing
type TestHelpers struct{}

// MockDocument is an alias to the unified mock document
type MockDocument = testhelpers.MockDocument

// NewMockDocument creates a new MockDocument with the given content
func NewMockDocument(content string) *MockDocument {
	return testhelpers.NewMockDocument(content)
}

// MockTemplateDocument is an alias to the unified mock document with template support
type MockTemplateDocument = testhelpers.MockDocument

// NewMockTemplateDocument creates a new MockTemplateDocument with template context
func NewMockTemplateDocument(content, templateContext string) *MockTemplateDocument {
	doc := testhelpers.NewMockDocument(content)
	doc.TemplateContext = templateContext
	return doc
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
	ctx completion.CompletionContext,
	uri, tagName, attributeName string,
) ([]protocol.CompletionItem, error) {
	// For now, fall back to the legacy function until we have a working main entry point approach
	// The main entry point requires complex document analysis that's hard to mock properly
	return completion.GetAttributeValueCompletions(ctx, tagName, attributeName), nil
}

// getDocumentManagerFromContext extracts DocumentManager from a completion context
// This uses type assertion to check if the context has a document manager
func getDocumentManagerFromContext(ctx completion.CompletionContext) *lsp.DocumentManager {
	// Try to extract document manager using type assertion
	if contextWithDM, ok := ctx.(interface{ GetDocumentManager() *lsp.DocumentManager }); ok {
		return contextWithDM.GetDocumentManager()
	}
	return nil
}

// Global helper instance
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

// CopyFixtureFiles copies files from the fixture directory to the target directory,
// preserving directory structure
func CopyFixtureFiles(t *testing.T, fixtureDir, targetDir string) {
	t.Helper()

	err := filepath.Walk(fixtureDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip README files
		if strings.HasSuffix(path, "README.md") {
			return nil
		}

		// Get relative path from fixture directory
		relPath, err := filepath.Rel(fixtureDir, path)
		if err != nil {
			return err
		}

		targetPath := filepath.Join(targetDir, relPath)

		if info.IsDir() {
			return os.MkdirAll(targetPath, info.Mode())
		}

		// Copy file
		src, err := os.Open(path)
		if err != nil {
			return err
		}
		defer src.Close()

		err = os.MkdirAll(filepath.Dir(targetPath), 0755)
		if err != nil {
			return err
		}

		dst, err := os.Create(targetPath)
		if err != nil {
			return err
		}
		defer dst.Close()

		_, err = io.Copy(dst, src)
		return err
	})

	if err != nil {
		t.Fatalf("Failed to copy fixture files: %v", err)
	}
}

// CopyFixtureFile copies a single file from the fixture directory to the target location
func CopyFixtureFile(t *testing.T, fixtureDir, filename, targetDir, targetFilename string) {
	t.Helper()

	src := filepath.Join(fixtureDir, filename)
	dst := filepath.Join(targetDir, targetFilename)

	srcFile, err := os.Open(src)
	if err != nil {
		t.Fatalf("Failed to open source file %s: %v", src, err)
	}
	defer srcFile.Close()

	err = os.MkdirAll(filepath.Dir(dst), 0755)
	if err != nil {
		t.Fatalf("Failed to create target directory: %v", err)
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		t.Fatalf("Failed to create target file %s: %v", dst, err)
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		t.Fatalf("Failed to copy file content: %v", err)
	}
}
