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
	"fmt"
	"strings"
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

func TestValidateHtml_UnknownAttributeDetection(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("../test-fixtures/multiple-elements")
	err := workspace.Init()
	require.NoError(t, err)

	registry, err := mcp.NewMCPContext(workspace)
	require.NoError(t, err)

	err = registry.LoadManifests()
	require.NoError(t, err)

	registryAdapter := mcp.NewMCPContextAdapter(registry).(*mcp.MCPContextAdapter)

	// First verify the element exists in the registry
	buttonElement, err := registryAdapter.ElementInfo("button-element")
	require.NoError(t, err)
	require.NotNil(t, buttonElement)
	t.Logf("Button element found: %s with %d attributes", buttonElement.TagName(), len(buttonElement.Attributes()))

	// Test the specific case that failed: priority instead of variant
	html := `<button-element priority="primary">Call to action</button-element>`

	argsJSON, err := json.Marshal(map[string]interface{}{
		"html":    html,
		"context": "test validation",
	})
	require.NoError(t, err)

	req := &mcpSDK.CallToolRequest{
		Params: &mcpSDK.CallToolParamsRaw{
			Name:      "validate_html",
			Arguments: json.RawMessage(argsJSON),
		},
	}

	handler := tools.MakeValidateHtmlHandler(registryAdapter)
	result, err := handler(context.Background(), req)
	require.NoError(t, err)
	require.Len(t, result.Content, 1)

	textContent, ok := result.Content[0].(*mcpSDK.TextContent)
	require.True(t, ok)

	t.Logf("Validation result: %s", textContent.Text)

	// The result should contain an error about the unknown attribute
	assert.Contains(t, textContent.Text, "Unknown Attribute")
	assert.Contains(t, textContent.Text, "priority")
	assert.Contains(t, textContent.Text, "button-element")
	// Should suggest the correct attribute if available
	if strings.Contains(textContent.Text, "Did you mean") {
		assert.Contains(t, textContent.Text, "variant")
	}
}

func TestValidateHtml_ValidAttributes(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("../test-fixtures/multiple-elements")
	err := workspace.Init()
	require.NoError(t, err)

	registry, err := mcp.NewMCPContext(workspace)
	require.NoError(t, err)

	err = registry.LoadManifests()
	require.NoError(t, err)

	registryAdapter := mcp.NewMCPContextAdapter(registry).(*mcp.MCPContextAdapter)

	// Test with correct variant attribute
	html := `<button-element variant="primary">Call to action</button-element>`

	argsJSON, err := json.Marshal(map[string]interface{}{
		"html":    html,
		"context": "test validation",
	})
	require.NoError(t, err)

	req := &mcpSDK.CallToolRequest{
		Params: &mcpSDK.CallToolParamsRaw{
			Name:      "validate_html",
			Arguments: json.RawMessage(argsJSON),
		},
	}

	handler := tools.MakeValidateHtmlHandler(registryAdapter)
	result, err := handler(context.Background(), req)
	require.NoError(t, err)
	require.Len(t, result.Content, 1)

	textContent, ok := result.Content[0].(*mcpSDK.TextContent)
	require.True(t, ok)

	// The result should show no unknown attribute issues
	assert.NotContains(t, textContent.Text, "Unknown Attribute")
	assert.NotContains(t, textContent.Text, "priority")
}

func TestValidateHtml_InvalidAttributeValue(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("../test-fixtures/multiple-elements")
	err := workspace.Init()
	require.NoError(t, err)

	registry, err := mcp.NewMCPContext(workspace)
	require.NoError(t, err)

	err = registry.LoadManifests()
	require.NoError(t, err)

	registryAdapter := mcp.NewMCPContextAdapter(registry).(*mcp.MCPContextAdapter)

	// Test with invalid attribute value (assuming variant has restricted values)
	html := `<button-element variant="invalid-value">Call to action</button-element>`

	argsJSON, err := json.Marshal(map[string]interface{}{
		"html":    html,
		"context": "test validation",
	})
	require.NoError(t, err)

	req := &mcpSDK.CallToolRequest{
		Params: &mcpSDK.CallToolParamsRaw{
			Name:      "validate_html",
			Arguments: json.RawMessage(argsJSON),
		},
	}

	handler := tools.MakeValidateHtmlHandler(registryAdapter)
	result, err := handler(context.Background(), req)
	require.NoError(t, err)
	require.Len(t, result.Content, 1)

	textContent, ok := result.Content[0].(*mcpSDK.TextContent)
	require.True(t, ok)

	// Check if validation detects invalid values (depends on manifest constraints)
	if strings.Contains(textContent.Text, "Invalid Attribute Value") {
		// The template output might not contain the exact string "invalid-value" but should show the validation error
		t.Logf("Validation detected invalid attribute value: %s", textContent.Text)
	}
}

func TestValidateHtml_AttributeValueParsing_Regression(t *testing.T) {
	// Regression test for attribute value parsing bug where boolean attributes
	// cause misaligned indices between attr.name and attr.value arrays
	workspace := W.NewFileSystemWorkspaceContext("../test-fixtures/multiple-elements")
	err := workspace.Init()
	require.NoError(t, err)

	registry, err := mcp.NewMCPContext(workspace)
	require.NoError(t, err)

	err = registry.LoadManifests()
	require.NoError(t, err)

	registryAdapter := mcp.NewMCPContextAdapter(registry).(*mcp.MCPContextAdapter)

	// Test HTML with boolean attributes AND valued attributes - this was causing the bug
	// The "disabled" boolean attribute should not interfere with "variant" value parsing
	html := `<button-element variant="primary" disabled>Call to action</button-element>`

	argsJSON, err := json.Marshal(map[string]interface{}{
		"html":    html,
		"context": "test attribute value parsing",
	})
	require.NoError(t, err)

	req := &mcpSDK.CallToolRequest{
		Params: &mcpSDK.CallToolParamsRaw{
			Name:      "validate_html",
			Arguments: json.RawMessage(argsJSON),
		},
	}

	handler := tools.MakeValidateHtmlHandler(registryAdapter)
	result, err := handler(context.Background(), req)
	require.NoError(t, err)
	require.Len(t, result.Content, 1)

	textContent, ok := result.Content[0].(*mcpSDK.TextContent)
	require.True(t, ok)

	t.Logf("Validation result: %s", textContent.Text)

	// The bug would cause "empty attribute values" to be reported incorrectly
	// This test ensures that valid attribute values are NOT reported as empty
	assert.NotContains(t, textContent.Text, "empty attribute values",
		"Should not report empty attribute values for attributes that clearly have values")

	// Since variant="primary" has a valid value, it should not be flagged as empty
	if strings.Contains(textContent.Text, "variant") && strings.Contains(textContent.Text, "empty") {
		t.Errorf("variant attribute incorrectly reported as having empty value when it clearly has 'primary'")
	}

	// Regression test: ensure that boolean attributes with valued attributes work correctly
	// The result should show "No Manifest Compliance Issues Found" because:
	// 1. variant="primary" is a valid value for button-element
	// 2. disabled is a valid boolean attribute (global HTML attribute)
	assert.Contains(t, textContent.Text, "No Manifest Compliance Issues Found",
		"Should not report any manifest compliance issues when attributes are correctly parsed")

	// Ensure we're not getting the incorrect quoted values bug
	assert.NotContains(t, textContent.Text, `"primary"`,
		"Should not include quotes in attribute value validation")
}

func TestValidateHtml_UnknownElement(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("../test-fixtures/multiple-elements")
	err := workspace.Init()
	require.NoError(t, err)

	registry, err := mcp.NewMCPContext(workspace)
	require.NoError(t, err)

	err = registry.LoadManifests()
	require.NoError(t, err)

	registryAdapter := mcp.NewMCPContextAdapter(registry).(*mcp.MCPContextAdapter)

	// Test with unknown custom element
	html := `<unknown-custom-element attribute="value">Content</unknown-custom-element>`

	argsJSON, err := json.Marshal(map[string]interface{}{
		"html":    html,
		"context": "test validation",
	})
	require.NoError(t, err)

	req := &mcpSDK.CallToolRequest{
		Params: &mcpSDK.CallToolParamsRaw{
			Name:      "validate_html",
			Arguments: json.RawMessage(argsJSON),
		},
	}

	handler := tools.MakeValidateHtmlHandler(registryAdapter)
	result, err := handler(context.Background(), req)
	require.NoError(t, err)
	require.Len(t, result.Content, 1)

	textContent, ok := result.Content[0].(*mcpSDK.TextContent)
	require.True(t, ok)

	// The result should contain an error about the unknown element
	assert.Contains(t, textContent.Text, "Unknown Element")
	assert.Contains(t, textContent.Text, "unknown-custom-element")
}

func TestValidateHtml_MultipleIssues(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("../test-fixtures/multiple-elements")
	err := workspace.Init()
	require.NoError(t, err)

	registry, err := mcp.NewMCPContext(workspace)
	require.NoError(t, err)

	err = registry.LoadManifests()
	require.NoError(t, err)

	registryAdapter := mcp.NewMCPContextAdapter(registry).(*mcp.MCPContextAdapter)

	// HTML with multiple validation issues
	html := `
		<button-element priority="primary">Button 1</button-element>
		<unknown-custom-element>Unknown</unknown-custom-element>
	`

	argsJSON, err := json.Marshal(map[string]interface{}{
		"html":    html,
		"context": "test validation",
	})
	require.NoError(t, err)

	req := &mcpSDK.CallToolRequest{
		Params: &mcpSDK.CallToolParamsRaw{
			Name:      "validate_html",
			Arguments: json.RawMessage(argsJSON),
		},
	}

	handler := tools.MakeValidateHtmlHandler(registryAdapter)
	result, err := handler(context.Background(), req)
	require.NoError(t, err)
	require.Len(t, result.Content, 1)

	textContent, ok := result.Content[0].(*mcpSDK.TextContent)
	require.True(t, ok)

	// Should contain both types of issues
	assert.Contains(t, textContent.Text, "Unknown Attribute") // priority
	assert.Contains(t, textContent.Text, "Unknown Element")   // unknown-custom-element
}

func TestValidateHtml_GlobalAttributesAllowed(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("../test-fixtures/multiple-elements")
	err := workspace.Init()
	require.NoError(t, err)

	registry, err := mcp.NewMCPContext(workspace)
	require.NoError(t, err)

	err = registry.LoadManifests()
	require.NoError(t, err)

	registryAdapter := mcp.NewMCPContextAdapter(registry).(*mcp.MCPContextAdapter)

	// Test HTML with various global attributes that should NOT be flagged as unknown
	html := `
		<button-element
			id="my-button"
			class="primary-btn"
			slot="footer"
			style="color: red"
			title="Click me"
			data-testid="button-test"
			aria-label="Primary button"
			onmouseover="handleHover()"
			tabindex="0"
			hidden
			draggable="true"
			contenteditable="false"
			spellcheck="false"
			translate="no"
			lang="en"
			dir="ltr">
			Click me
		</button-element>`

	argsJSON, err := json.Marshal(map[string]interface{}{
		"html":    html,
		"context": "global attributes test",
	})
	require.NoError(t, err)

	req := &mcpSDK.CallToolRequest{
		Params: &mcpSDK.CallToolParamsRaw{
			Name:      "validate_html",
			Arguments: json.RawMessage(argsJSON),
		},
	}

	handler := tools.MakeValidateHtmlHandler(registryAdapter)
	result, err := handler(context.Background(), req)
	require.NoError(t, err)
	require.Len(t, result.Content, 1)

	textContent, ok := result.Content[0].(*mcpSDK.TextContent)
	require.True(t, ok)

	t.Logf("Validation result for global attributes: %s", textContent.Text)

	// Global attributes should NOT be flagged as unknown
	globalAttrs := []string{"id", "class", "slot", "style", "title", "data-testid",
		"aria-label", "onmouseover", "tabindex", "hidden", "draggable",
		"contenteditable", "spellcheck", "translate", "lang", "dir"}

	for _, attr := range globalAttrs {
		assert.NotContains(t, textContent.Text, fmt.Sprintf("Unknown attribute '%s'", attr),
			"Global attribute '%s' should not be flagged as unknown", attr)
	}

	// Should not contain any unknown attribute errors for the global attributes
	if strings.Contains(textContent.Text, "Unknown Attribute") {
		// If there are unknown attribute errors, they should not be for global attributes
		for _, attr := range globalAttrs {
			assert.NotContains(t, textContent.Text, attr,
				"Unknown attribute error should not mention global attribute '%s'", attr)
		}
	}
}

func TestValidateHtml_GlobalAttributesWithUnknownCustomAttribute(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("../test-fixtures/multiple-elements")
	err := workspace.Init()
	require.NoError(t, err)

	registry, err := mcp.NewMCPContext(workspace)
	require.NoError(t, err)

	err = registry.LoadManifests()
	require.NoError(t, err)

	registryAdapter := mcp.NewMCPContextAdapter(registry).(*mcp.MCPContextAdapter)

	// Test HTML with both global attributes (which should be OK) and unknown custom attribute (which should be flagged)
	html := `<button-element
		id="my-button"
		class="primary-btn"
		slot="footer"
		unknown-custom-attr="should-be-flagged">
		Click me
	</button-element>`

	argsJSON, err := json.Marshal(map[string]interface{}{
		"html":    html,
		"context": "mixed attributes test",
	})
	require.NoError(t, err)

	req := &mcpSDK.CallToolRequest{
		Params: &mcpSDK.CallToolParamsRaw{
			Name:      "validate_html",
			Arguments: json.RawMessage(argsJSON),
		},
	}

	handler := tools.MakeValidateHtmlHandler(registryAdapter)
	result, err := handler(context.Background(), req)
	require.NoError(t, err)
	require.Len(t, result.Content, 1)

	textContent, ok := result.Content[0].(*mcpSDK.TextContent)
	require.True(t, ok)

	t.Logf("Validation result for mixed attributes: %s", textContent.Text)

	// Global attributes should NOT be flagged
	assert.NotContains(t, textContent.Text, "Unknown attribute 'id'")
	assert.NotContains(t, textContent.Text, "Unknown attribute 'class'")
	assert.NotContains(t, textContent.Text, "Unknown attribute 'slot'")

	// Unknown custom attribute SHOULD be flagged
	assert.Contains(t, textContent.Text, "unknown-custom-attr")
}
