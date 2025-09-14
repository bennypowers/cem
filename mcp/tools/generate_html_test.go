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
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/mcp/tools"
	mcpSDK "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateHTML_WithFixtures(t *testing.T) {
	registry := getTestRegistry(t)

	// Find all test fixture directories
	fixturesDir := "test-fixtures/generate-html"
	fixtures, err := os.ReadDir(fixturesDir)
	require.NoError(t, err, "Should be able to read fixtures directory")

	for _, fixture := range fixtures {
		if !fixture.IsDir() {
			continue
		}

		fixtureName := fixture.Name()
		t.Run(fixtureName, func(t *testing.T) {
			// Load input
			inputPath := filepath.Join(fixturesDir, fixtureName, "input.json")
			inputData, err := os.ReadFile(inputPath)
			require.NoError(t, err, "Should be able to read input fixture")

			var args tools.GenerateHtmlArgs
			err = json.Unmarshal(inputData, &args)
			require.NoError(t, err, "Should be able to parse input JSON")

			// Create MCP request
			argsJSON, err := json.Marshal(args)
			require.NoError(t, err)

			req := &mcpSDK.CallToolRequest{
				Params: &mcpSDK.CallToolParamsRaw{
					Name:      "generate_html",
					Arguments: json.RawMessage(argsJSON),
				},
			}

			// Call the actual MCP tool through public API
			handler := tools.MakeGenerateHtmlHandler(registry)
			result, err := handler(context.Background(), req)

			// Handle error cases (like element-not-found)
			if fixtureName == "element-not-found" {
				require.NoError(t, err, "Tool should not return Go error, but MCP error response")
				require.NotNil(t, result, "Should return error response")
				require.Len(t, result.Content, 1, "Should have error content")

				textContent, ok := result.Content[0].(*mcpSDK.TextContent)
				require.True(t, ok, "Content should be text")

				// Check against expected error response
				expectedPath := filepath.Join(fixturesDir, fixtureName, "expected-response.md")
				expectedData, err := os.ReadFile(expectedPath)
				require.NoError(t, err, "Should be able to read expected response")

				assert.Contains(t, textContent.Text, string(expectedData), "Should contain expected error message")
				return
			}

			// For successful cases
			require.NoError(t, err, "Tool should not return error")
			require.NotNil(t, result, "Should return result")
			require.Len(t, result.Content, 1, "Should have content")

			textContent, ok := result.Content[0].(*mcpSDK.TextContent)
			require.True(t, ok, "Content should be text")

			output := textContent.Text
			assert.NotEmpty(t, output, "Tool output should not be empty")

			// Check that output contains expected HTML (if expected-html.html exists)
			expectedHTMLPath := filepath.Join(fixturesDir, fixtureName, "expected-html.html")
			if _, err := os.Stat(expectedHTMLPath); err == nil {
				expectedHTML, err := os.ReadFile(expectedHTMLPath)
				require.NoError(t, err, "Should be able to read expected HTML")

				assert.Contains(t, output, string(expectedHTML), "Output should contain expected HTML structure")
			}

			// General validation for successful HTML generation
			assert.Contains(t, output, args.TagName, "Should contain element tag name")
			if args.Content != "" {
				assert.Contains(t, output, args.Content, "Should contain provided content")
			}

			// Verify no template syntax remains
			assert.NotContains(t, output, "{{", "Output should not contain unprocessed template syntax")
			assert.NotContains(t, output, "}}", "Output should not contain unprocessed template syntax")
		})
	}
}

func TestGenerateHTMLTool_Integration(t *testing.T) {
	registry := getTestRegistry(t)

	// Quick integration test to verify tool handler exists and basic functionality
	handler := tools.MakeGenerateHtmlHandler(registry)
	assert.NotNil(t, handler, "Tool should have handler")

	// Test successful case
	args := tools.GenerateHtmlArgs{
		TagName: "button-element",
		Content: "Test Button",
		Context: "integration-test",
	}

	argsJSON, err := json.Marshal(args)
	require.NoError(t, err)

	req := &mcpSDK.CallToolRequest{
		Params: &mcpSDK.CallToolParamsRaw{
			Name:      "generate_html",
			Arguments: json.RawMessage(argsJSON),
		},
	}

	result, err := handler(context.Background(), req)
	require.NoError(t, err, "Tool should not return error")
	require.NotNil(t, result, "Tool should return result")
	require.Len(t, result.Content, 1, "Result should have content")

	textContent, ok := result.Content[0].(*mcpSDK.TextContent)
	require.True(t, ok, "Content should be text")

	output := textContent.Text
	assert.NotEmpty(t, output, "Tool output should not be empty")
	assert.Contains(t, output, "button-element", "Should contain element name")
	assert.Contains(t, output, "Test Button", "Should contain provided content")
}

func TestGenerateHTMLRegression_TagNameAccess(t *testing.T) {
	// Regression test for the specific error in the transcript:
	// "can't evaluate field TagName in type tools.HTMLGenerationTemplateData"

	registry := getTestRegistry(t)

	// Test that the html_structure template can access TagName correctly
	args := tools.GenerateHtmlArgs{
		TagName: "button-element",
		Content: "test content",
		Context: "regression-test",
	}

	argsJSON, err := json.Marshal(args)
	require.NoError(t, err)

	req := &mcpSDK.CallToolRequest{
		Params: &mcpSDK.CallToolParamsRaw{
			Name:      "generate_html",
			Arguments: json.RawMessage(argsJSON),
		},
	}

	// Call through the full MCP tool chain
	handler := tools.MakeGenerateHtmlHandler(registry)
	result, err := handler(context.Background(), req)

	// This should NOT fail with "can't evaluate field TagName"
	require.NoError(t, err, "Tool should render without TagName field error")
	require.NotNil(t, result, "Should return result")
	require.Len(t, result.Content, 1, "Should have content")

	textContent, ok := result.Content[0].(*mcpSDK.TextContent)
	require.True(t, ok, "Content should be text")

	output := textContent.Text
	assert.NotEmpty(t, output, "Template should produce output")

	// Verify the fix worked - should contain the actual tag name
	assert.Contains(t, output, "button-element", "Generated HTML should contain the correct tag name")
	assert.Contains(t, output, "&amp;lt;button-element", "Should contain opening tag (double-encoded)")
	assert.Contains(t, output, "&amp;lt;/button-element&gt;", "Should contain closing tag (double-encoded)")
	assert.Contains(t, output, "test content", "Should contain provided content")
}
