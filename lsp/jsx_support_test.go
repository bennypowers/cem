package lsp_test

import (
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/document"
	"bennypowers.dev/cem/lsp/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestJSXDocumentSupport(t *testing.T) {
	tests := []struct {
		name     string
		uri      string
		expected string
	}{
		{
			name:     "TSX file should be detected as TSX language",
			uri:      "file:///path/to/App.tsx",
			expected: "tsx",
		},
		{
			name:     "JSX file should be detected as TSX language",
			uri:      "file:///path/to/Component.jsx",
			expected: "tsx",
		},
		{
			name:     "Regular TypeScript should remain typescript",
			uri:      "file:///path/to/regular.ts",
			expected: "typescript",
		},
		{
			name:     "Regular JavaScript should remain typescript",
			uri:      "file:///path/to/regular.js",
			expected: "typescript",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This test will fail initially as JSX support doesn't exist yet
			dm, err := document.NewDocumentManager()
			require.NoError(t, err, "Should create document manager")
			defer dm.Close()

			// Open document and check language detection
			doc := dm.OpenDocument(tt.uri, "// test content", 1)

			// Check language detection using the interface method
			assert.Equal(t, tt.expected, doc.Language(), "Document language should match expected")
		})
	}
}

func TestJSXCustomElementDetection(t *testing.T) {
	// Setup test registry with JSX fixtures
	fixtureDir := filepath.Join("test", "fixtures", "jsx-support-test")

	// Create mock file watcher for testing
	mockWatcher := platform.NewMockFileWatcher()
	registry := lsp.NewRegistry(mockWatcher)

	manifestPath := filepath.Join(fixtureDir, "custom-elements.json")
	registry.AddManifestPath(manifestPath)

	dm, err := document.NewDocumentManager()
	require.NoError(t, err, "Should create document manager")
	defer dm.Close()

	t.Run("TSX file custom element detection", func(t *testing.T) {
		// Read fixture content
		tsxPath := filepath.Join(fixtureDir, "App.tsx")
		content, err := os.ReadFile(tsxPath)
		require.NoError(t, err, "Should read TSX fixture file")

		// This will fail initially - JSX parsing not implemented yet
		doc := dm.OpenDocument("file:///"+tsxPath, string(content), 1)
		elements, err := doc.FindCustomElements(dm)
		require.NoError(t, err, "Should find custom elements in TSX")

		// Expected elements: 5 total (2 my-card, 3 ui-button)
		assert.Len(t, elements, 5, "Should find 5 custom elements")

		// Verify element detection
		elementNames := make([]string, len(elements))
		for i, elem := range elements {
			elementNames[i] = elem.TagName
		}

		assert.Contains(t, elementNames, "my-card", "Should detect my-card element")
		assert.Contains(t, elementNames, "ui-button", "Should detect ui-button elements")
	})

	t.Run("JSX file custom element detection", func(t *testing.T) {
		// Read fixture content
		jsxPath := filepath.Join(fixtureDir, "Component.jsx")
		content, err := os.ReadFile(jsxPath)
		require.NoError(t, err, "Should read JSX fixture file")

		// This will fail initially - JSX parsing not implemented yet
		doc := dm.OpenDocument("file:///"+jsxPath, string(content), 1)
		elements, err := doc.FindCustomElements(dm)
		require.NoError(t, err, "Should find custom elements in TSX")

		// Expected elements: 8 total (4 my-card, 4 ui-button)
		assert.Len(t, elements, 8, "Should find 8 custom elements")
	})
}

func TestJSXAttributeDetection(t *testing.T) {
	// Setup test registry
	fixtureDir := filepath.Join("test", "fixtures", "jsx-support-test")

	// Create mock file watcher for testing
	mockWatcher := platform.NewMockFileWatcher()
	registry := lsp.NewRegistry(mockWatcher)

	manifestPath := filepath.Join(fixtureDir, "custom-elements.json")
	registry.AddManifestPath(manifestPath)

	dm, err := document.NewDocumentManager()
	require.NoError(t, err, "Should create document manager")
	defer dm.Close()

	// Read fixture content
	attributeTestPath := filepath.Join(fixtureDir, "AttributeTest.tsx")
	content, err := os.ReadFile(attributeTestPath)
	require.NoError(t, err, "Should read attribute test fixture")

	// This will fail initially - JSX parsing not implemented yet
	doc := dm.OpenDocument("file:///"+attributeTestPath, string(content), 1)
	elements, err := doc.FindCustomElements(dm)
	require.NoError(t, err, "Should parse JSX without error")

	// Find my-card element (order is not guaranteed)
	var myCard *types.CustomElementMatch
	for i := range elements {
		if elements[i].TagName == "my-card" {
			myCard = &elements[i]
			break
		}
	}

	require.NotNil(t, myCard, "Should find my-card element in parsed elements")

	// Verify attributes are detected
	expectedAttrs := []string{"title", "variant", "disabled"}
	for _, expectedAttr := range expectedAttrs {
		_, exists := myCard.Attributes[expectedAttr]
		assert.True(t, exists, "Should detect %s attribute", expectedAttr)
	}
}

func TestJSXCompletionContext(t *testing.T) {
	// Setup test environment
	dm, err := document.NewDocumentManager()
	require.NoError(t, err, "Should create document manager")
	defer dm.Close()

	fixtureDir := filepath.Join("test", "fixtures", "jsx-support-test")

	tests := []struct {
		name         string
		fixtureFile  string
		position     protocol.Position
		expectedType types.CompletionContextType
		expectedTag  string
		expectedAttr string
	}{
		{
			name:         "Tag name completion in JSX",
			fixtureFile:  "TagCompletion.tsx",
			position:     protocol.Position{Line: 2, Character: 13},
			expectedType: types.CompletionTagName,
		},
		{
			name:         "Attribute name completion in JSX",
			fixtureFile:  "AttributeNameCompletion.tsx",
			position:     protocol.Position{Line: 2, Character: 18},
			expectedType: types.CompletionAttributeName,
			expectedTag:  "my-card",
		},
		{
			name:         "Attribute value completion in JSX",
			fixtureFile:  "AttributeValueCompletion.tsx",
			position:     protocol.Position{Line: 2, Character: 27},
			expectedType: types.CompletionAttributeValue,
			expectedTag:  "my-card",
			expectedAttr: "variant",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read fixture content
			fixturePath := filepath.Join(fixtureDir, tt.fixtureFile)
			content, err := os.ReadFile(fixturePath)
			require.NoError(t, err, "Should read fixture file")

			// This will fail initially - JSX completion context not implemented
			doc := dm.OpenDocument("file:///"+fixturePath, string(content), 1)
			analysis := doc.AnalyzeCompletionContextTS(tt.position, dm)

			if analysis != nil {
				assert.Equal(t, tt.expectedType, analysis.Type, "Completion type should match")
				if tt.expectedTag != "" {
					assert.Equal(t, tt.expectedTag, analysis.TagName, "Tag name should match")
				}
				if tt.expectedAttr != "" {
					assert.Equal(t, tt.expectedAttr, analysis.AttributeName, "Attribute name should match")
				}
			}
		})
	}
}
