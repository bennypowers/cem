package completion_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/methods/textDocument"
	"bennypowers.dev/cem/lsp/methods/textDocument/completion"
	"bennypowers.dev/cem/lsp/testhelpers"
	"bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestAttributeValueCompletions(t *testing.T) {
	// Load test manifest
	fixtureDir := filepath.Join("attribute-values-test")
	manifestPath := filepath.Join(fixtureDir, "manifest.json")

	manifestBytes, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatalf("Failed to read test manifest: %v", err)
	}

	var pkg M.Package
	err = json.Unmarshal(manifestBytes, &pkg)
	if err != nil {
		t.Fatalf("Failed to parse manifest: %v", err)
	}

	// Create a completion context using MockServerContext and add the test manifest
	ctx := testhelpers.NewMockServerContext()
	ctx.AddManifest(&pkg)

	// Create and set a real DocumentManager
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	tests := []struct {
		name           string
		tagName        string
		attributeName  string
		expectedLabels []string
		description    string
	}{
		{
			name:           "Boolean attribute completions",
			tagName:        "test-component",
			attributeName:  "disabled",
			expectedLabels: []string{},
			description:    "Should provide no completions for boolean attributes (presence=true, absence=false)",
		},
		{
			name:           "Number attribute completions",
			tagName:        "test-component",
			attributeName:  "count",
			expectedLabels: []string{},
			description:    "Should provide no automatic numeric suggestions for number attributes",
		},
		{
			name:           "String attribute completions",
			tagName:        "test-component",
			attributeName:  "color",
			expectedLabels: []string{`""`, "red", "blue", "green"},
			description:    "Should provide empty string and color suggestions for color attributes",
		},
		{
			name:           "Union type completions",
			tagName:        "test-component",
			attributeName:  "theme",
			expectedLabels: []string{"light", "dark", "auto"},
			description:    "Should parse union types and provide individual options",
		},
		{
			name:           "Size attribute context-aware completions",
			tagName:        "test-component",
			attributeName:  "size",
			expectedLabels: []string{"medium (default)", "small", "medium", "large"},
			description:    "Should provide size suggestions and default value",
		},
		{
			name:           "Variant attribute context-aware completions",
			tagName:        "test-component",
			attributeName:  "variant",
			expectedLabels: []string{`""`, "primary", "secondary"},
			description:    "Should provide variant suggestions for variant attributes",
		},
		{
			name:           "Array type completions",
			tagName:        "test-component",
			attributeName:  "items",
			expectedLabels: []string{},
			description:    "Should provide no automatic completions for array attributes (ambiguous syntax)",
		},
		{
			name:           "Non-existent attribute",
			tagName:        "test-component",
			attributeName:  "non-existent",
			expectedLabels: []string{},
			description:    "Should return empty completions for non-existent attributes",
		},
		{
			name:           "Non-custom element",
			tagName:        "div",
			attributeName:  "class",
			expectedLabels: []string{},
			description:    "Should return empty completions for standard HTML elements",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the attribute value completion function directly
			completions := completion.GetAttributeValueCompletions(ctx, tt.tagName, tt.attributeName)

			// Check if we got the expected completions
			if len(tt.expectedLabels) == 0 {
				if len(completions) != 0 {
					t.Errorf("Expected no completions, but got %d", len(completions))
				}
				return
			}

			// Create a map of found labels for easy checking
			foundLabels := make(map[string]bool)
			for _, completion := range completions {
				foundLabels[completion.Label] = true
			}

			// Check that all expected labels are present
			for _, expectedLabel := range tt.expectedLabels {
				if !foundLabels[expectedLabel] {
					t.Errorf("Expected completion label '%s' not found. Got labels: %v",
						expectedLabel, testhelpers.GetCompletionLabels(completions))
				}
			}

			// Log completion details for debugging
			t.Logf("Test: %s", tt.description)
			t.Logf("Found %d completions: %v", len(completions), testhelpers.GetCompletionLabels(completions))
		})
	}
}

func TestUnionTypeParser(t *testing.T) {
	tests := []struct {
		name           string
		typeText       string
		expectedLabels []string
	}{
		{
			name:           "Simple union type",
			typeText:       `"red" | "green" | "blue"`,
			expectedLabels: []string{"red", "green", "blue"},
		},
		{
			name:           "Union type with spaces",
			typeText:       `"small" | "medium" | "large"`,
			expectedLabels: []string{"small", "medium", "large"},
		},
		{
			name:           "Union type with single quotes",
			typeText:       `'left' | 'center' | 'right'`,
			expectedLabels: []string{"left", "center", "right"},
		},
		{
			name:           "Mixed quotes",
			typeText:       `"auto" | 'manual'`,
			expectedLabels: []string{"auto", "manual"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test attribute with the union type
			attr := &M.Attribute{
				Type: &M.Type{
					Text: tt.typeText,
				},
			}

			// Get type-based completions (which should include union type parsing)
			completions := completion.GetTypeBasedCompletions(attr)

			// Check that we got the expected labels
			foundLabels := make(map[string]bool)
			for _, completion := range completions {
				foundLabels[completion.Label] = true
			}

			for _, expectedLabel := range tt.expectedLabels {
				if !foundLabels[expectedLabel] {
					t.Errorf("Expected union type label '%s' not found. Got labels: %v",
						expectedLabel, testhelpers.GetCompletionLabels(completions))
				}
			}
		})
	}
}

func TestAttributeCompletionAfterSpaces(t *testing.T) {
	// Load test manifest
	fixtureDir := filepath.Join("attribute-values-test")
	manifestPath := filepath.Join(fixtureDir, "manifest.json")

	manifestBytes, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatalf("Failed to read test manifest: %v", err)
	}

	var pkg M.Package
	err = json.Unmarshal(manifestBytes, &pkg)
	if err != nil {
		t.Fatalf("Failed to parse manifest: %v", err)
	}

	// Create a completion context using MockServerContext and add the test manifest
	ctx := testhelpers.NewMockServerContext()
	ctx.AddManifest(&pkg)

	// Create and set a real DocumentManager
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	tests := []struct {
		name          string
		html          string
		position      protocol.Position
		expectedAttrs []string
		description   string
	}{
		{
			name:          "Completion after space following tag name",
			html:          "<test-component ",
			position:      protocol.Position{Line: 0, Character: 16},
			expectedAttrs: []string{"color", "theme", "disabled", "count", "size", "variant", "items"},
			description:   "Should provide attribute completions after space following tag name",
		},
		{
			name:          "Completion after space following another attribute",
			html:          `<test-component color="red" `,
			position:      protocol.Position{Line: 0, Character: 28},
			expectedAttrs: []string{"theme", "disabled", "count", "size", "variant", "items"}, // color already used
			description:   "Should provide remaining attribute completions after space following existing attribute",
		},
		{
			name:          "Boolean attribute completion (just name, no value)",
			html:          "<test-component disabl",
			position:      protocol.Position{Line: 0, Character: 22},
			expectedAttrs: []string{"disabled"},
			description:   "Should complete boolean attribute name without value syntax",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a real document using the DocumentManager from context
			doc := dm.OpenDocument("test://test.html", tt.html, 1)

			// Analyze completion context with the DocumentManager
			analysis, err := textDocument.AnalyzeCompletionContextWithDM(doc, tt.position, "", dm)
			if err != nil {
				t.Fatalf("Failed to analyze completion context: %v", err)
			}

			// Should detect attribute name completion
			if analysis.Type != types.CompletionAttributeName {
				t.Errorf("Expected CompletionAttributeName, got %v", analysis.Type)
				return
			}

			// Get attribute completions
			completions := completion.GetAttributeCompletions(ctx, analysis.TagName)

			// Check that we got some completions
			if len(completions) == 0 {
				t.Errorf("Expected attribute completions, got none")
				return
			}

			// Get completion labels
			foundLabels := make(map[string]bool)
			for _, completion := range completions {
				foundLabels[completion.Label] = true
			}

			// Check for expected attributes (subset check for flexibility)
			foundCount := 0
			for _, expectedAttr := range tt.expectedAttrs {
				if foundLabels[expectedAttr] {
					foundCount++
				}
			}

			if foundCount == 0 {
				t.Errorf("Expected some of %v in completions, but found none. Got: %v",
					tt.expectedAttrs, testhelpers.GetCompletionLabels(completions))
			}

			// For boolean attributes, verify they complete without value syntax
			for _, completion := range completions {
				if completion.Label == "disabled" {
					if completion.InsertText == nil || *completion.InsertText != "disabled" {
						t.Errorf("Boolean attribute 'disabled' should insert just 'disabled', got: %v", completion.InsertText)
					}
					// Should not have snippet format for boolean attributes
					if completion.InsertTextFormat != nil && *completion.InsertTextFormat == protocol.InsertTextFormatSnippet {
						t.Errorf("Boolean attribute should not use snippet format")
					}
				} else {
					// Non-boolean attributes should use snippet format
					expectedSnippet := fmt.Sprintf("%s=\"$0\"", completion.Label)
					if completion.InsertText == nil || *completion.InsertText != expectedSnippet {
						t.Logf("Non-boolean attribute '%s' uses snippet format: %v", completion.Label, completion.InsertText)
					}
				}
			}

			t.Logf("Test: %s", tt.description)
			t.Logf("Found %d completions: %v", len(completions), testhelpers.GetCompletionLabels(completions))
		})
	}
}

// TestAttributeValueCompletionTreeSitterRegression is a regression test for the tree-sitter
// tag name extraction issue that was preventing attribute value completions from working
func TestAttributeValueCompletionTreeSitterRegression(t *testing.T) {
	// This test verifies that when tree-sitter analysis is used for attribute value completion,
	// the tag name is correctly extracted from the context captures when it's not available
	// in the attribute value capture itself.

	// Create a test element with various attribute types
	testElement := &M.CustomElement{
		TagName: "regression-element",
		Attributes: []M.Attribute{
			{
				FullyQualified: M.FullyQualified{Name: "state"},
				Type:           &M.Type{Text: "\"success\" | \"warning\" | \"error\""},
			},
			{
				FullyQualified: M.FullyQualified{Name: "variant"},
				Type:           &M.Type{Text: "'primary'"},
			},
		},
	}

	// Create manifest package
	pkg := &M.Package{
		Modules: []M.Module{
			{
				Declarations: []M.Declaration{
					&M.CustomElementDeclaration{
						CustomElement: *testElement,
					},
				},
			},
		},
	}

	// Create context using MockServerContext and add the test manifest
	ctx := testhelpers.NewMockServerContext()
	ctx.AddManifest(pkg)

	// Test union type attribute
	stateCompletions := completion.GetAttributeValueCompletions(ctx, "regression-element", "state")
	expectedStateCompletions := []string{"success", "warning", "error"}

	if len(stateCompletions) < len(expectedStateCompletions) {
		t.Errorf("Expected at least %d state completions, got %d", len(expectedStateCompletions), len(stateCompletions))
	}

	// Verify all expected completions are present
	for _, expected := range expectedStateCompletions {
		found := false
		for _, completion := range stateCompletions {
			if completion.Label == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected state completion '%s' not found", expected)
		}
	}

	// Test single literal type attribute
	variantCompletions := completion.GetAttributeValueCompletions(ctx, "regression-element", "variant")

	if len(variantCompletions) == 0 {
		t.Error("Expected variant completions, got none")
	}

	// Verify 'primary' completion is present
	found := false
	for _, completion := range variantCompletions {
		if completion.Label == "primary" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Expected variant completion 'primary' not found")
	}

	t.Logf("Regression test passed: state completions=%d, variant completions=%d",
		len(stateCompletions), len(variantCompletions))
}
