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
package mcp

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"bennypowers.dev/cem/mcp/security"
	"bennypowers.dev/cem/mcp/tools"
	"bennypowers.dev/cem/mcp/types"
	"bennypowers.dev/cem/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// newTestTemplateRenderer creates a template renderer for testing
func newTestTemplateRenderer() *tools.TemplateRenderer {
	return tools.NewTemplateRenderer()
}

// mockElementInfo implements types.ElementInfo for testing
type mockElementInfo struct {
	tagName     string
	description string
}

func (m *mockElementInfo) TagName() string                      { return m.tagName }
func (m *mockElementInfo) Name() string                         { return m.tagName }
func (m *mockElementInfo) Description() string                  { return m.description }
func (m *mockElementInfo) Module() string                       { return "" }
func (m *mockElementInfo) Package() string                      { return "" }
func (m *mockElementInfo) Guidelines() []string                 { return []string{} }
func (m *mockElementInfo) Attributes() []types.Attribute        { return []types.Attribute{} }
func (m *mockElementInfo) Slots() []types.Slot                  { return []types.Slot{} }
func (m *mockElementInfo) Events() []types.Event                { return []types.Event{} }
func (m *mockElementInfo) CssProperties() []types.CssProperty   { return []types.CssProperty{} }
func (m *mockElementInfo) CssParts() []types.CssPart            { return []types.CssPart{} }
func (m *mockElementInfo) CssStates() []types.CssState          { return []types.CssState{} }
func (m *mockElementInfo) Examples() []types.Example            { return []types.Example{} }
func (m *mockElementInfo) ItemsByKind(kind string) []types.Item { return []types.Item{} }

func TestMaliciousManifestMitigation(t *testing.T) {
	// Create a temporary directory for our test
	tempDir, err := os.MkdirTemp("", "mcp_security_test")
	require.NoError(t, err)
	defer func() {
		if err := os.RemoveAll(tempDir); err != nil {
			t.Logf("Failed to remove temp dir: %v", err)
		}
	}()

	// Create a malicious manifest
	maliciousManifest := map[string]interface{}{
		"schemaVersion": "1.0.0",
		"readme":        "",
		"modules": []map[string]interface{}{
			{
				"kind": "javascript-module",
				"path": "src/malicious-button.js",
				"declarations": []map[string]interface{}{
					{
						"kind":          "class",
						"description":   "A button component {{.DatabasePassword}} with {{range .Secrets}}{{.}}{{end}} injection attempts",
						"name":          "MaliciousButton",
						"tagName":       "malicious-button",
						"customElement": true,
						"attributes": []map[string]interface{}{
							{
								"name":        "variant",
								"description": "Button style {{.AdminToken}} and <script>fetch('/admin/secrets')</script> injection",
								"type":        map[string]interface{}{"text": "string"},
							},
						},
						"slots": []map[string]interface{}{
							{
								"name":        "",
								"description": "Button content {{with .Config}}{{.APIKey}}{{end}} injection",
							},
						},
						"events": []map[string]interface{}{
							{
								"name":        "click",
								"description": "Click event {{template \"../../../etc/passwd\" .}} injection",
							},
						},
						"cssProperties": []map[string]interface{}{
							{
								"name":        "--button-color",
								"description": "Button color {{define \"malicious\"}}{{.}}{{end}} injection",
								"syntax":      "<color>",
							},
						},
						"cssParts": []map[string]interface{}{
							{
								"name":        "button",
								"description": "Button element {{block \"evil\" .}}default{{end}} injection",
							},
						},
						"cssStates": []map[string]interface{}{
							{
								"name":        "disabled",
								"description": "Disabled state {{$secret := .Config.Secret}}{{$secret}} injection",
							},
						},
					},
				},
			},
		},
	}

	// Write the malicious manifest
	manifestPath := filepath.Join(tempDir, "custom-elements.json")
	manifestBytes, err := json.MarshalIndent(maliciousManifest, "", "  ")
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(manifestPath, manifestBytes, 0644))

	// Create a package.json pointing to the manifest
	packageJSON := map[string]interface{}{
		"name":           "malicious-package",
		"version":        "1.0.0",
		"customElements": "custom-elements.json",
	}
	packageBytes, err := json.MarshalIndent(packageJSON, "", "  ")
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(filepath.Join(tempDir, "package.json"), packageBytes, 0644))

	// Create workspace and registry
	wctx := workspace.NewFileSystemWorkspaceContext(tempDir)
	require.NoError(t, wctx.Init())

	registry, err := NewMCPContext(wctx)
	require.NoError(t, err)

	// Load the malicious manifest
	require.NoError(t, registry.LoadManifests())

	// Get element info for the malicious button
	element, err := registry.GetElementInfo("malicious-button")
	require.NoError(t, err)
	require.NotNil(t, element)

	// Test that the description is sanitized
	t.Run("element description sanitized", func(t *testing.T) {
		description := element.Description
		assert.NotContains(t, description, "{{", "Template syntax should be removed from element description")
		assert.NotContains(t, description, "}}", "Template syntax should be removed from element description")

		// Note: Element description might be empty because it's not populated from manifest description
		// The important thing is that no template injection is present
		t.Logf("Element description: '%s'", description)
	})

	// Test that attributes are sanitized
	t.Run("attribute descriptions sanitized", func(t *testing.T) {
		attributes := element.Attributes()
		require.Len(t, attributes, 1)

		attr := attributes[0]
		description := attr.Description()
		assert.NotContains(t, description, "{{", "Template syntax should be removed from attribute description")
		assert.NotContains(t, description, "}}", "Template syntax should be removed from attribute description")
		assert.Contains(t, description, "&lt;script&gt;", "HTML should be escaped in attribute description")
	})

	// Test that slots are sanitized
	t.Run("slot descriptions sanitized", func(t *testing.T) {
		slots := element.Slots()
		require.Len(t, slots, 1)

		slot := slots[0]
		description := slot.Description()
		assert.NotContains(t, description, "{{", "Template syntax should be removed from slot description")
		assert.NotContains(t, description, "}}", "Template syntax should be removed from slot description")
	})

	// Test that events are sanitized
	t.Run("event descriptions sanitized", func(t *testing.T) {
		events := element.Events()
		require.Len(t, events, 1)

		event := events[0]
		description := event.Description()
		assert.NotContains(t, description, "{{", "Template syntax should be removed from event description")
		assert.NotContains(t, description, "}}", "Template syntax should be removed from event description")
	})

	// Test that CSS properties are sanitized
	t.Run("css property descriptions sanitized", func(t *testing.T) {
		props := element.CssProperties()
		require.Len(t, props, 1)

		prop := props[0]
		description := prop.Description()
		assert.NotContains(t, description, "{{", "Template syntax should be removed from CSS property description")
		assert.NotContains(t, description, "}}", "Template syntax should be removed from CSS property description")
	})

	// Test that CSS parts are sanitized
	t.Run("css part descriptions sanitized", func(t *testing.T) {
		parts := element.CssParts()
		require.Len(t, parts, 1)

		part := parts[0]
		description := part.Description()
		assert.NotContains(t, description, "{{", "Template syntax should be removed from CSS part description")
		assert.NotContains(t, description, "}}", "Template syntax should be removed from CSS part description")
	})

	// Test that CSS states are sanitized
	t.Run("css state descriptions sanitized", func(t *testing.T) {
		states := element.CssStates()
		require.Len(t, states, 1)

		state := states[0]
		description := state.Description()
		assert.NotContains(t, description, "{{", "Template syntax should be removed from CSS state description")
		assert.NotContains(t, description, "}}", "Template syntax should be removed from CSS state description")
	})
}

func TestTemplateRenderingWithMaliciousData(t *testing.T) {
	// This test simulates what would happen if malicious data somehow made it through to template rendering
	// It verifies that our template security controls and html/template auto-escaping prevent execution

	// Create a template renderer for testing
	templateRenderer := newTestTemplateRenderer()

	// Test data that simulates malicious content that bypassed input sanitization
	// (this shouldn't happen due to our input sanitization, but tests defense in depth)
	maliciousTestCases := []struct {
		name         string
		templateName string
		data         interface{}
		expectError  bool
		description  string
	}{
		{
			name:         "html_injection_in_template_data",
			templateName: "html_validation_results",
			data: tools.HTMLValidationData{
				FoundElements: []tools.ElementWithIssues{
					{
						ElementInfo: &mockElementInfo{
							tagName:     "malicious-element",
							description: "<script>alert('xss')</script>Button component",
						},
						UsageCount: 1,
					},
				},
			},
			expectError: false,
			description: "HTML should be auto-escaped by html/template",
		},
		{
			name:         "javascript_in_template_data",
			templateName: "html_validation_results",
			data: tools.HTMLValidationData{
				FoundElements: []tools.ElementWithIssues{
					{
						ElementInfo: &mockElementInfo{
							tagName:     "test-element",
							description: "javascript:alert('xss')",
						},
						UsageCount: 1,
					},
				},
			},
			expectError: false,
			description: "JavaScript URLs in text content are preserved (not a security issue in this context)",
		},
		{
			name:         "template_syntax_in_data",
			templateName: "html_validation_results",
			data: tools.HTMLValidationData{
				FoundElements: []tools.ElementWithIssues{
					{
						ElementInfo: &mockElementInfo{
							tagName:     "test-element",
							description: "{{.Config}} template injection attempt",
						},
						UsageCount: 1,
					},
				},
			},
			expectError: false,
			description: "Template syntax in data should be rendered as literal text",
		},
		{
			name:         "invalid_template_name",
			templateName: "../../../etc/passwd",
			data:         map[string]interface{}{},
			expectError:  true,
			description:  "Path traversal in template name should be rejected",
		},
		{
			name:         "directory_traversal_template",
			templateName: "../../../../sensitive_file",
			data:         map[string]interface{}{},
			expectError:  true,
			description:  "Directory traversal attempts should be blocked",
		},
	}

	for _, tc := range maliciousTestCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := templateRenderer.Render(tc.templateName, tc.data)

			if tc.expectError {
				assert.Error(t, err, tc.description)
				assert.Empty(t, result, "Should not return content on error")
			} else {
				assert.NoError(t, err, tc.description)
				assert.NotEmpty(t, result, "Should return rendered content")

				// Verify that dangerous content is properly escaped
				if strings.Contains(tc.templateName, "html_validation_results") {
					// Check that script tags are escaped
					assert.NotContains(t, result, "<script>", "Script tags should be escaped")
					// Note: javascript: in text content is not escaped by html/template and that's okay

					// Check that HTML is escaped by html/template
					if data, ok := tc.data.(tools.HTMLValidationData); ok {
						if len(data.FoundElements) > 0 {
							element := data.FoundElements[0]
							desc := element.Description()
							if strings.Contains(desc, "<script>") {
								assert.Contains(t, result, "&lt;script&gt;", "HTML should be escaped")
							}
							if strings.Contains(desc, "{{.Config}}") {
								// The template syntax should appear literally in the output
								assert.Contains(t, result, "{{.Config}}", "Template syntax should be literal")
								assert.NotContains(t, result, "actual_config_value", "Template should not execute")
							}
						}
					}
				}
			}
		})
	}
}

func TestSecurityMetrics(t *testing.T) {
	// Test that security functions perform adequately
	t.Run("sanitization performance", func(t *testing.T) {
		maliciousInput := strings.Repeat("{{.Secret}} ", 1000) + "Normal content"

		// Should complete in reasonable time
		start := func() {
			for i := 0; i < 100; i++ {
				security.SanitizeDescription(maliciousInput)
			}
		}

		// Just verify it doesn't panic or hang
		start()
	})

	t.Run("detection accuracy", func(t *testing.T) {
		// Test with various injection patterns
		injectionTests := []string{
			"{{.Config}}",
			"{{range .Items}}{{.}}{{end}}",
			"{{with .Data}}{{.Secret}}{{end}}",
			"{{template \"malicious\" .}}",
			"{{define \"evil\"}}{{.}}{{end}}",
			"{{block \"bad\" .}}default{{end}}",
		}

		for _, test := range injectionTests {
			assert.True(t, security.DetectTemplateInjection(test), "Should detect injection in: %s", test)
		}

		// Test false positives
		safeTests := []string{
			"Normal description",
			"JavaScript: {key: value}",
			"CSS: .class { color: red; }",
			"Code example: if (x) { return true; }",
		}

		for _, test := range safeTests {
			assert.False(t, security.DetectTemplateInjection(test), "Should not detect injection in: %s", test)
		}
	})
}
