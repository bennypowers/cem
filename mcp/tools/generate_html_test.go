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
	"context"
	"encoding/json"
	"testing"

	"bennypowers.dev/cem/mcp/tools"
	mcpSDK "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateHTMLStructure(t *testing.T) {
	registry := getTestRegistry(t)

	tests := []struct {
		name          string
		tagName       string
		content       string
		attributes    map[string]string
		expectError   bool
		checkHTML     []string // strings that should be present in HTML output
	}{
		{
			name:        "basic element generation",
			tagName:     "button-element",
			content:     "Click me",
			expectError: false,
			checkHTML:   []string{"&lt;button-element", "&lt;/button-element>", "Click me"},
		},
		{
			name:        "element with attributes",
			tagName:     "button-element",
			content:     "Submit",
			attributes:  map[string]string{"variant": "primary", "size": "large"},
			expectError: false,
			checkHTML:   []string{"&lt;button-element", "variant=\"primary\"", "size=\"large\"", "Submit", "&lt;/button-element>"},
		},
		{
			name:        "element with slots - default content",
			tagName:     "button-element",
			content:     "Button Text",
			expectError: false,
			checkHTML:   []string{"&lt;button-element", "&lt;/button-element>", "Button Text"},
		},
		{
			name:        "non-existent element",
			tagName:     "non-existent-element",
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Get element info
			element, errorResponse, err := tools.LookupElement(registry, test.tagName)
			if test.expectError {
				// Should either have lookup error or error response
				if err == nil && errorResponse == nil {
					t.Fatalf("Expected error for non-existent element %s", test.tagName)
				}
				return
			}

			require.NoError(t, err)
			require.Nil(t, errorResponse)
			require.NotNil(t, element)

			// Test generateHTMLStructure function directly
			args := tools.GenerateHtmlArgs{
				TagName:    test.tagName,
				Content:    test.content,
				Attributes: test.attributes,
				Context:    "test",
			}

			html, err := tools.GenerateHTMLStructureForTesting(element, args)
			require.NoError(t, err, "generateHTMLStructure should not fail")
			assert.NotEmpty(t, html, "Generated HTML should not be empty")

			// Verify expected content
			for _, expected := range test.checkHTML {
				assert.Contains(t, html, expected, "Generated HTML should contain: %s", expected)
			}

			// Verify no template syntax remains
			assert.NotContains(t, html, "{{", "Generated HTML should not contain unprocessed template syntax")
			assert.NotContains(t, html, "}}", "Generated HTML should not contain unprocessed template syntax")

			// Verify proper tag structure (HTML encoded)
			assert.Contains(t, html, "&lt;"+test.tagName, "Should contain opening tag")
			assert.Contains(t, html, "&lt;/"+test.tagName+">", "Should contain closing tag")
		})
	}
}

func TestGenerateHTMLTool_Integration(t *testing.T) {
	registry := getTestRegistry(t)

	tests := []struct {
		name        string
		args        tools.GenerateHtmlArgs
		expectError bool
		checkOutput []string
	}{
		{
			name: "successful HTML generation",
			args: tools.GenerateHtmlArgs{
				TagName: "button-element",
				Content: "Test Button",
				Context: "form",
			},
			expectError: false,
			checkOutput: []string{"button-element", "Test Button", "Generated HTML"},
		},
		{
			name: "element not found",
			args: tools.GenerateHtmlArgs{
				TagName: "unknown-element",
				Context: "test",
			},
			expectError: false, // Tool returns error response, not Go error
			checkOutput: []string{"not found", "unknown-element"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Create MCP request
			argsJSON, err := json.Marshal(test.args)
			require.NoError(t, err)

			req := &mcpSDK.CallToolRequest{
				Params: &mcpSDK.CallToolParamsRaw{
					Name:      "generate_html",
					Arguments: json.RawMessage(argsJSON),
				},
			}

			// Call the tool handler
			result, err := tools.HandleGenerateHtmlForTesting(context.Background(), req, registry)

			if test.expectError {
				assert.Error(t, err, "Expected error for test: %s", test.name)
				return
			}

			require.NoError(t, err, "Tool should not return Go error")
			require.NotNil(t, result, "Tool should return result")
			require.Len(t, result.Content, 1, "Result should have content")

			textContent, ok := result.Content[0].(*mcpSDK.TextContent)
			require.True(t, ok, "Content should be text")

			output := textContent.Text
			assert.NotEmpty(t, output, "Tool output should not be empty")

			// Check expected output strings
			for _, expected := range test.checkOutput {
				assert.Contains(t, output, expected, "Output should contain: %s", expected)
			}
		})
	}
}

func TestGenerateHTMLRegression_TagNameAccess(t *testing.T) {
	// Regression test for the specific error in the transcript:
	// "can't evaluate field TagName in type tools.HTMLGenerationTemplateData"

	registry := getTestRegistry(t)

	// Get button-element which we know exists in fixtures
	element, errorResponse, err := tools.LookupElement(registry, "button-element")
	require.NoError(t, err)
	require.Nil(t, errorResponse)
	require.NotNil(t, element)

	// Test the specific template data structure that was failing
	templateData := tools.HTMLGenerationTemplateData{
		BaseTemplateData: tools.NewBaseTemplateData(element, "test", nil),
		Content:          "test content",
	}

	// This should NOT fail with "can't evaluate field TagName"
	html, err := tools.RenderTemplate("html_structure", templateData)
	require.NoError(t, err, "html_structure template should render without TagName field error")
	assert.NotEmpty(t, html, "Template should produce HTML")

	// Verify the fix worked - should contain the actual tag name
	assert.Contains(t, html, "button-element", "Generated HTML should contain the correct tag name")
	assert.Contains(t, html, "&lt;button-element", "Should contain opening tag")
	assert.Contains(t, html, "&lt;/button-element>", "Should contain closing tag")
}