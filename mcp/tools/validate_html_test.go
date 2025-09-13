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
	"testing"

	"bennypowers.dev/cem/mcp"
	"bennypowers.dev/cem/mcp/tools"
	W "bennypowers.dev/cem/workspace"
	mcpSDK "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateHTML_Args(t *testing.T) {
	// Test basic argument parsing for validate HTML tool
	tests := []struct {
		name        string
		args        tools.ValidateHtmlArgs
		expectValid bool
	}{
		{
			name: "valid arguments",
			args: tools.ValidateHtmlArgs{
				Html:    `<button-element variant="primary">Click me</button-element>`,
				Context: "form",
			},
			expectValid: true,
		},
		{
			name: "empty HTML",
			args: tools.ValidateHtmlArgs{
				Html:    "",
				Context: "empty",
			},
			expectValid: true, // Should handle gracefully
		},
		{
			name: "HTML without custom elements",
			args: tools.ValidateHtmlArgs{
				Html:    `<div><p>Regular HTML content</p></div>`,
				Context: "standard",
			},
			expectValid: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Basic validation that args are well-formed
			if test.expectValid {
				assert.NotNil(t, test.args.Html, "HTML should be provided")
			}
		})
	}
}

func TestValidateHtmlArgs_Parsing(t *testing.T) {
	tests := []struct {
		name        string
		input       map[string]interface{}
		expectError bool
		checkFunc   func(t *testing.T, args tools.ValidateHtmlArgs)
	}{
		{
			name: "complete args",
			input: map[string]interface{}{
				"html":    "<button-element>Test</button-element>",
				"context": "validation-test",
				"tagName": "button-element",
			},
			expectError: false,
			checkFunc: func(t *testing.T, args tools.ValidateHtmlArgs) {
				assert.Equal(t, "<button-element>Test</button-element>", args.Html)
				assert.Equal(t, "validation-test", args.Context)
				assert.Equal(t, "button-element", args.TagName)
			},
		},
		{
			name: "minimal args",
			input: map[string]interface{}{
				"html": "<div>content</div>",
			},
			expectError: false,
			checkFunc: func(t *testing.T, args tools.ValidateHtmlArgs) {
				assert.Equal(t, "<div>content</div>", args.Html)
				assert.Empty(t, args.Context)
				assert.Empty(t, args.TagName)
			},
		},
		{
			name: "empty HTML",
			input: map[string]interface{}{
				"html":    "",
				"context": "empty-test",
			},
			expectError: false,
			checkFunc: func(t *testing.T, args tools.ValidateHtmlArgs) {
				assert.Empty(t, args.Html)
				assert.Equal(t, "empty-test", args.Context)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			argsJSON, err := json.Marshal(test.input)
			require.NoError(t, err)

			req := &mcpSDK.CallToolRequest{
				Params: &mcpSDK.CallToolParamsRaw{
					Name:      "validate_html",
					Arguments: json.RawMessage(argsJSON),
				},
			}

			args, err := tools.ParseToolArgs[tools.ValidateHtmlArgs](req)

			if test.expectError {
				assert.Error(t, err, "Expected error for test: %s", test.name)
			} else {
				assert.NoError(t, err, "Should not error for test: %s", test.name)

				if test.checkFunc != nil {
					test.checkFunc(t, args)
				}
			}
		})
	}
}

func TestValidationTemplateData_WithFixtures(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("../test-fixtures/multiple-elements")
	err := workspace.Init()
	require.NoError(t, err)

	registry, err := mcp.NewMCPContext(workspace)
	require.NoError(t, err)

	err = registry.LoadManifests()
	require.NoError(t, err)

	registryAdapter := mcp.NewMCPContextAdapter(registry).(*mcp.MCPContextAdapter)

	buttonElement, err := registryAdapter.ElementInfo("button-element")
	require.NoError(t, err)

	html := `<button-element variant="primary">Test</button-element>`
	context := "testing"
	options := map[string]string{"framework": "lit"}

	data := tools.NewValidationTemplateData(buttonElement, context, options, html)

	assert.Equal(t, buttonElement, data.Element, "Element should be set")
	assert.Equal(t, context, data.Context, "Context should be set")
	assert.Equal(t, options, data.Options, "Options should be set")
	assert.Equal(t, html, data.Html, "HTML should be set")

	// Test interface compliance
	assert.Equal(t, buttonElement, data.GetElement())
	assert.Equal(t, context, data.GetContext())
	assert.Equal(t, options, data.GetOptions())
}
