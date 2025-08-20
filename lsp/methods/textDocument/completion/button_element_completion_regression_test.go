package completion_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/methods/textDocument/completion"
	"bennypowers.dev/cem/lsp/testhelpers"
	"bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// TestButtonElementAttributeValueRegression is a regression test for button-element
// attribute value completions in the demo-project context
func TestButtonElementAttributeValueRegression(t *testing.T) {
	// Load the demo-project manifest to test with real data
	manifestPath := filepath.Join("..", "..", "..", "..", "demo-project", "custom-elements.json")
	
	manifestBytes, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Skipf("Demo project manifest not found: %v", err)
	}

	var pkg M.Package
	err = json.Unmarshal(manifestBytes, &pkg)
	if err != nil {
		t.Fatalf("Failed to parse demo-project manifest: %v", err)
	}

	// Create a completion context using MockServerContext and add the demo-project manifest
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
		name            string
		html            string
		position        protocol.Position
		expectedValues  []string
		description     string
	}{
		{
			name:     "button-element variant attribute values",
			html:     `<button-element variant="">`,
			position: protocol.Position{Line: 0, Character: 25}, // Inside the empty quotes
			expectedValues: []string{
				"primary",   // from manifest union type
				"secondary", 
				"success",
				"danger",
				"warning",
			},
			description: "Should provide variant completions for button-element from demo-project manifest",
		},
		{
			name:     "button-element size attribute values", 
			html:     `<button-element size="">`,
			position: protocol.Position{Line: 0, Character: 22}, // Inside the empty quotes
			expectedValues: []string{
				"small",
				"medium", 
				"large",
			},
			description: "Should provide size completions for button-element from demo-project manifest",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use real Document parsing via DocumentManager (not mock parsing)
			testURI := "test://button-element-test.html"
			doc := dm.OpenDocument(testURI, tt.html, 1)

			// Analyze completion context using real tree-sitter parsing
			analysis := doc.AnalyzeCompletionContextTS(tt.position, dm)

			// Should detect attribute value completion
			if analysis == nil || analysis.Type != types.CompletionAttributeValue {
				if analysis == nil {
					t.Errorf("Expected CompletionAttributeValue, but got nil analysis")
				} else {
					t.Errorf("Expected CompletionAttributeValue (%d), got %d", types.CompletionAttributeValue, analysis.Type)
				}
				return
			}

			// Verify tag name extraction
			if analysis.TagName != "button-element" {
				t.Errorf("Expected tag name 'button-element', got '%s'", analysis.TagName)
				return
			}


			// Get attribute value completions using the completion system
			completions := completion.GetAttributeValueCompletions(ctx, analysis.TagName, analysis.AttributeName)

			// Check that we got some completions
			if len(completions) == 0 {
				t.Errorf("❌ REGRESSION: Expected attribute value completions for %s %s, got none", 
					analysis.TagName, analysis.AttributeName)
				return
			}

			// Create a map of found labels for easy checking
			foundLabels := make(map[string]bool)
			for _, completion := range completions {
				foundLabels[completion.Label] = true
			}

			// Check that all expected values are present
			for _, expectedValue := range tt.expectedValues {
				if !foundLabels[expectedValue] {
					t.Errorf("❌ REGRESSION: Expected completion value '%s' not found. Got labels: %v",
						expectedValue, testhelpers.GetCompletionLabels(completions))
				}
			}

			t.Logf("✅ Test: %s", tt.description)
			t.Logf("Found %d completions: %v", len(completions), testhelpers.GetCompletionLabels(completions))
		})
	}
}

// TestButtonElementIntegrationContext tests the full integration with actual HTML context
func TestButtonElementIntegrationContext(t *testing.T) {
	// Load the demo-project manifest
	manifestPath := filepath.Join("..", "..", "..", "..", "demo-project", "custom-elements.json")
	
	manifestBytes, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Skipf("Demo project manifest not found: %v", err)
	}

	var pkg M.Package
	err = json.Unmarshal(manifestBytes, &pkg)
	if err != nil {
		t.Fatalf("Failed to parse demo-project manifest: %v", err)
	}

	// Create a completion context using MockServerContext and add the demo-project manifest
	ctx := testhelpers.NewMockServerContext()
	ctx.AddManifest(&pkg)

	// Create and set a real DocumentManager
	dm, err := lsp.NewDocumentManager()
	if err != nil {
		t.Fatalf("Failed to create DocumentManager: %v", err)
	}
	defer dm.Close()
	ctx.SetDocumentManager(dm)

	// Test with more realistic HTML context similar to the demo project
	htmlContent := `<!DOCTYPE html>
<html lang="en">
<head>
  <script type="module">
    import "cem-lsp-demo/components/card-element.js";
    import "cem-lsp-demo/components/button-element.js";
  </script>
</head>
<body>
  <main class="demo-section">
    <h2 class="demo-title">Button Elements</h2>
    <card-element>
      <button-element slot="footer" variant="">
    </card-element>
  </main>
</body>
</html>`

	// Position should be inside the variant="" attribute value (let me count the exact position)
	// Line 12: "      <button-element slot="footer" variant="">"
	//                                               ^
	//          012345678901234567890123456789012345678901234567890
	//                    1         2         3         4         5
	position := protocol.Position{Line: 12, Character: 45} // Inside the quotes

	// Use real Document parsing via DocumentManager (not mock parsing)
	testURI := "test://button-element-integration.html"
	doc := dm.OpenDocument(testURI, htmlContent, 1)

	// Analyze completion context using real tree-sitter parsing
	analysis := doc.AnalyzeCompletionContextTS(position, dm)

	// Should detect attribute value completion
	if analysis.Type != types.CompletionAttributeValue {
		t.Errorf("Expected CompletionAttributeValue, got %v for position %+v", analysis.Type, position)
		return
	}

	// Verify tag name extraction
	if analysis.TagName != "button-element" {
		t.Errorf("Expected tag name 'button-element', got '%s'", analysis.TagName)
		return
	}

	// Verify attribute name extraction
	if analysis.AttributeName != "variant" {
		t.Errorf("Expected attribute name 'variant', got '%s'", analysis.AttributeName)
		return
	}

	// Get attribute value completions
	completions := completion.GetAttributeValueCompletions(ctx, analysis.TagName, analysis.AttributeName)

	// Check that we got the expected completions
	if len(completions) == 0 {
		t.Errorf("❌ REGRESSION: Expected variant completions for button-element in full HTML context, got none")
		return
	}

	// Verify key completion values are present
	expectedValues := []string{"primary", "secondary", "success", "danger", "warning"}
	foundLabels := make(map[string]bool)
	for _, completion := range completions {
		foundLabels[completion.Label] = true
	}

	for _, expectedValue := range expectedValues {
		if !foundLabels[expectedValue] {
			t.Errorf("❌ REGRESSION: Expected completion value '%s' not found in full HTML context. Got labels: %v",
				expectedValue, testhelpers.GetCompletionLabels(completions))
		}
	}

	t.Logf("✅ Integration test passed with %d completions: %v", 
		len(completions), testhelpers.GetCompletionLabels(completions))
}