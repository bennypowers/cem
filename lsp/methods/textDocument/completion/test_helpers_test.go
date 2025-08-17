package completion_test

import (
	"fmt"

	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/methods/textDocument/completion"
	"bennypowers.dev/cem/lsp/types"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// TestHelpers provides utilities for completion testing
type TestHelpers struct{}

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
type MockDocument struct {
	content string
}

func (m *MockDocument) Content() string {
	return m.content
}

func (m *MockDocument) FindElementAtPosition(position protocol.Position, dm any) *types.CustomElementMatch {
	// Simple mock implementation - not used in most tests
	return nil
}

func (m *MockDocument) FindAttributeAtPosition(position protocol.Position, dm any) (*types.AttributeMatch, string) {
	// Simple mock implementation - not used in most tests
	return nil, ""
}

func (m *MockDocument) TemplateContext(position protocol.Position) string {
	return "" // HTML documents don't have template context
}

func (m *MockDocument) Version() int32 {
	return 1
}

func (m *MockDocument) URI() string {
	return "test://mock"
}

func (m *MockDocument) FindCustomElements(dm any) ([]types.CustomElementMatch, error) {
	return nil, nil
}

func (m *MockDocument) AnalyzeCompletionContextTS(position protocol.Position, dm any) *types.CompletionAnalysis {
	return nil
}

func (m *MockDocument) GetScriptTags() []types.ScriptTag {
	return nil
}

func (m *MockDocument) FindModuleScript() (protocol.Position, bool) {
	return protocol.Position{}, false
}

// MockTemplateDocument extends MockDocument with template context support
type MockTemplateDocument struct {
	content       string
	isLitTemplate bool
}

func (m *MockTemplateDocument) Content() string {
	return m.content
}

func (m *MockTemplateDocument) FindElementAtPosition(position protocol.Position, dm any) *types.CustomElementMatch {
	return nil
}

func (m *MockTemplateDocument) FindAttributeAtPosition(position protocol.Position, dm any) (*types.AttributeMatch, string) {
	return nil, ""
}

func (m *MockTemplateDocument) TemplateContext(position protocol.Position) string {
	if m.isLitTemplate {
		return "html"
	}
	return "innerHTML"
}

func (m *MockTemplateDocument) Version() int32 {
	return 1
}

func (m *MockTemplateDocument) URI() string {
	return "test://mock-template"
}

func (m *MockTemplateDocument) FindCustomElements(dm any) ([]types.CustomElementMatch, error) {
	return nil, nil
}

func (m *MockTemplateDocument) AnalyzeCompletionContextTS(position protocol.Position, dm any) *types.CompletionAnalysis {
	return nil
}

func (m *MockTemplateDocument) GetScriptTags() []types.ScriptTag {
	return nil
}

func (m *MockTemplateDocument) FindModuleScript() (protocol.Position, bool) {
	return protocol.Position{}, false
}

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
