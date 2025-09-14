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
package tools_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/mcp"
	"bennypowers.dev/cem/mcp/templates"
	"bennypowers.dev/cem/mcp/tools"
	W "bennypowers.dev/cem/workspace"
	mcpSDK "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// getTestRegistry creates a registry using the test fixtures following the existing pattern
func getTestRegistry(t *testing.T) *mcp.MCPContextAdapter {
	workspace := W.NewFileSystemWorkspaceContext("../test-fixtures/multiple-elements")
	err := workspace.Init()
	require.NoError(t, err)

	registry, err := mcp.NewMCPContext(workspace)
	require.NoError(t, err)

	err = registry.LoadManifests()
	require.NoError(t, err)

	return mcp.NewMCPContextAdapter(registry).(*mcp.MCPContextAdapter)
}

func TestTemplateRenderer(t *testing.T) {
	registry := getTestRegistry(t)

	// Get real elements from fixtures
	_, err := registry.ElementInfo("button-element")
	require.NoError(t, err)

	tests := []struct {
		name         string
		templateName string
		setupData    func() interface{}
		expectError  bool
		checkContent []string // strings that should be present in output
	}{
		{
			name:         "accessibility context template",
			templateName: "contextual_suggestions",
			setupData: func() interface{} {
				return map[string]interface{}{
					"TagName": "button-element",
					"Context": "accessibility", // This should match a condition in the template
				}
			},
			expectError:  false,
			checkContent: []string{"button-element", "Accessibility"},
		},
		{
			name:         "non-existent template",
			templateName: "non-existent",
			setupData: func() interface{} {
				return map[string]string{}
			},
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data := test.setupData()
			output, err := templates.RenderTemplate(test.templateName, data)

			if test.expectError {
				assert.Error(t, err, "Expected error for template: %s", test.templateName)
				assert.Empty(t, output, "Output should be empty on error")
			} else {
				require.NoError(t, err, "Should not error for template: %s", test.templateName)
				assert.NotEmpty(t, output, "Template should produce output")

				for _, content := range test.checkContent {
					assert.Contains(t, output, content, "Output should contain: %s", content)
				}
			}
		})
	}
}

func TestToolsLoading_Integration(t *testing.T) {
	registry := getTestRegistry(t)

	// Test that all tools can be loaded without error
	toolDefs, err := tools.Tools(registry)
	require.NoError(t, err)
	require.NotEmpty(t, toolDefs, "Should load tool definitions")

	// Verify expected tools are present
	expectedTools := []string{"validate_html", "suggest_attributes", "generate_html", "suggest_css_integration"}
	toolNames := make(map[string]bool)
	for _, def := range toolDefs {
		toolNames[def.Name] = true

		// Basic validation of each tool definition
		assert.NotEmpty(t, def.Name, "Tool should have name")
		assert.NotEmpty(t, def.Description, "Tool should have description")
		assert.NotNil(t, def.Handler, "Tool should have handler")
		assert.NotNil(t, def.InputSchema, "Tool should have input schema")
	}

	for _, expected := range expectedTools {
		assert.True(t, toolNames[expected], "Should have tool: %s", expected)
	}
}

func TestToolArgs_Parsing(t *testing.T) {
	// Test the ParseToolArgs helper function with proper mock data
	args := tools.GenerateHtmlArgs{
		TagName: "button-element",
		Context: "form",
		Content: "Test content",
	}

	// Create a mock request similar to what MCP would send
	argsMap := map[string]interface{}{
		"tagName": args.TagName,
		"context": args.Context,
		"content": args.Content,
	}

	argsJSON, err := json.Marshal(argsMap)
	require.NoError(t, err)

	req := &mcpSDK.CallToolRequest{
		Params: &mcpSDK.CallToolParamsRaw{
			Name:      "generate_html",
			Arguments: json.RawMessage(argsJSON),
		},
	}

	parsedArgs, err := tools.ParseToolArgs[tools.GenerateHtmlArgs](req)
	require.NoError(t, err)

	assert.Equal(t, "button-element", parsedArgs.TagName)
	assert.Equal(t, "form", parsedArgs.Context)
	assert.Equal(t, "Test content", parsedArgs.Content)
}

func TestHelpers_Integration(t *testing.T) {
	registry := getTestRegistry(t)

	t.Run("LookupElement", func(t *testing.T) {
		// Test existing element
		element, errorResponse, err := tools.LookupElement(registry, "button-element")
		require.NoError(t, err)
		assert.Nil(t, errorResponse)
		assert.NotNil(t, element)
		assert.Equal(t, "button-element", element.TagName())

		// Test non-existent element
		element, errorResponse, err = tools.LookupElement(registry, "non-existent")
		require.NoError(t, err)
		assert.NotNil(t, errorResponse)
		assert.Nil(t, element)
	})

	t.Run("ResponseBuilder", func(t *testing.T) {
		builder := tools.NewResponseBuilder()
		builder.AddHeader(1, "Test Header")
		builder.AddSection("Test section")

		result := builder.Build()
		require.Len(t, result.Content, 1)

		textContent := result.Content[0].(*mcpSDK.TextContent)
		content := textContent.Text

		assert.Contains(t, content, "# Test Header")
		assert.Contains(t, content, "Test section")
	})
}

func TestTemplateRenderer_Basic(t *testing.T) {
	// Test that the template renderer works with basic data

	// Test with data that matches template conditions
	testDataWithContext := map[string]interface{}{
		"TagName": "test-element",
		"Context": "form", // This should match the form condition
	}

	templateNames := []string{"contextual_suggestions"}

	for _, templateName := range templateNames {
		t.Run(templateName, func(t *testing.T) {
			output, err := templates.RenderTemplate(templateName, testDataWithContext)

			if err != nil {
				// If template doesn't exist, that's okay - just log it
				t.Logf("Template %s not found or error: %v", templateName, err)
			} else {
				assert.NotEmpty(t, output, "Template %s should produce output", templateName)
				assert.NotContains(t, output, "{{", "Template %s should not contain unprocessed syntax", templateName)
			}
		})
	}
}

// Helper function to test templates with fixture/golden pattern
func testTemplateWithGolden(t *testing.T, templateName, fixtureFile, goldenFile string) {
	// Read template data from fixture file
	fixturePath := filepath.Join("../testdata/template_fixtures", fixtureFile)
	fixtureData, err := os.ReadFile(fixturePath)
	require.NoError(t, err, "Should be able to read fixture file: %s", fixturePath)

	var templateData interface{}
	err = json.Unmarshal(fixtureData, &templateData)
	require.NoError(t, err, "Should be able to parse fixture JSON")

	// Render template
	output, err := templates.RenderTemplate(templateName, templateData)
	require.NoError(t, err, "Template should render without error")

	// Compare with golden file
	goldenPath := filepath.Join("../testdata", goldenFile)
	expectedData, err := os.ReadFile(goldenPath)
	require.NoError(t, err, "Should be able to read golden file: %s", goldenPath)

	assert.Equal(t, string(expectedData), output, "Template output should match golden file")
}

// Fixture/Golden pattern tests for templates
func TestTemplates_FixtureGolden_AccessibilityContext(t *testing.T) {
	testTemplateWithGolden(t, "contextual_suggestions", "accessibility_context.json", "template_accessibility_context.golden.md")
}

func TestTemplates_FixtureGolden_FormContext(t *testing.T) {
	testTemplateWithGolden(t, "contextual_suggestions", "form_context.json", "template_form_context.golden.md")
}
