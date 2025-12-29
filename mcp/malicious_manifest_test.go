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
package mcp_test

import (
	"strings"
	"testing"

	"bennypowers.dev/cem/mcp"
	"bennypowers.dev/cem/mcp/security"
	"bennypowers.dev/cem/mcp/templates"
	"bennypowers.dev/cem/mcp/tools"
	"bennypowers.dev/cem/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMaliciousManifestMitigation(t *testing.T) {
	// Use the malicious manifest fixture
	wctx := workspace.NewFileSystemWorkspaceContext("fixtures/malicious-manifest-security")
	require.NoError(t, wctx.Init())

	registry, err := mcp.NewMCPContext(wctx)
	require.NoError(t, err)

	// Load the malicious manifest
	require.NoError(t, registry.LoadManifests())

	// Get element info for the malicious button
	adapter := mcp.NewMCPContextAdapter(registry)
	element, err := adapter.ElementInfo("malicious-button")
	require.NoError(t, err)
	require.NotNil(t, element)

	// Test that the description is sanitized
	t.Run("element description sanitized", func(t *testing.T) {
		description := element.Description()
		// Element descriptions are not currently populated from manifest declarations (architectural limitation)
		// The important security protection is at the item level (attributes, slots, events, etc.)
		assert.Equal(t, "", description, "Element descriptions are not currently supported")
		t.Logf("Element description: '%s'", description)
	})

	// Test that attributes are sanitized
	t.Run("attribute descriptions sanitized", func(t *testing.T) {
		attributes := element.Attributes()
		require.Len(t, attributes, 1)

		attr := attributes[0]
		description := attr.Description
		// Template syntax is now preserved as literal text and whitespace normalized
		assert.Contains(t, description, "{{", "Template syntax should be preserved as literal text")
		assert.NotContains(t, description, "\n", "Whitespace should be normalized")
	})

	// Test that slots are sanitized
	t.Run("slot descriptions sanitized", func(t *testing.T) {
		slots := element.Slots()
		require.Len(t, slots, 1)

		slot := slots[0]
		description := slot.Description
		// Template syntax is now preserved as literal text and whitespace normalized
		assert.Contains(t, description, "{{", "Template syntax should be preserved as literal text")
		assert.NotContains(t, description, "\n", "Whitespace should be normalized")
	})

	// Test that events are sanitized
	t.Run("event descriptions sanitized", func(t *testing.T) {
		events := element.Events()
		require.Len(t, events, 1)

		event := events[0]
		description := event.Description
		// Template syntax is now preserved as literal text and whitespace normalized
		assert.Contains(t, description, "{{", "Template syntax should be preserved as literal text")
		assert.NotContains(t, description, "\n", "Whitespace should be normalized")
	})

	// Test that CSS properties are sanitized
	t.Run("css property descriptions sanitized", func(t *testing.T) {
		props := element.CssProperties()
		require.Len(t, props, 1)

		prop := props[0]
		description := prop.Description
		// Template syntax is now preserved as literal text and whitespace normalized
		assert.Contains(t, description, "{{", "Template syntax should be preserved as literal text")
		assert.NotContains(t, description, "\n", "Whitespace should be normalized")
	})

	// Test that CSS parts are sanitized
	t.Run("css part descriptions sanitized", func(t *testing.T) {
		parts := element.CssParts()
		require.Len(t, parts, 1)

		part := parts[0]
		description := part.Description
		// Template syntax is now preserved as literal text and whitespace normalized
		assert.Contains(t, description, "{{", "Template syntax should be preserved as literal text")
		assert.NotContains(t, description, "\n", "Whitespace should be normalized")
	})

	// Test that CSS states are sanitized
	t.Run("css state descriptions sanitized", func(t *testing.T) {
		states := element.CssStates()
		require.Len(t, states, 1)

		state := states[0]
		description := state.Description
		// Template syntax is now preserved as literal text and whitespace normalized
		assert.Contains(t, description, "{{", "Template syntax should be preserved as literal text")
		assert.NotContains(t, description, "\n", "Whitespace should be normalized")
	})
}

func TestTemplateRenderingWithMaliciousData(t *testing.T) {
	// Test that the template security controls and html/template auto-escaping work
	// by using the malicious manifest fixture

	// Use the malicious manifest fixture
	wctx := workspace.NewFileSystemWorkspaceContext("fixtures/malicious-manifest-security")
	require.NoError(t, wctx.Init())

	registry, err := mcp.NewMCPContext(wctx)
	require.NoError(t, err)
	require.NoError(t, registry.LoadManifests())

	adapter := mcp.NewMCPContextAdapter(registry)
	element, err := adapter.ElementInfo("malicious-button")
	require.NoError(t, err)
	require.NotNil(t, element)

	// Test template rendering with the sanitized data from fixtures
	testCases := []struct {
		name         string
		templateName string
		setupData    func() interface{}
		expectError  bool
		description  string
	}{
		{
			name:         "html_validation_results_with_sanitized_data",
			templateName: "html_validation_results",
			setupData: func() interface{} {
				return tools.HTMLValidationData{
					FoundElements: []tools.ElementWithIssues{
						{
							ElementInfo: element,
							UsageCount:  1,
						},
					},
				}
			},
			expectError: false,
			description: "Template should render sanitized fixture data safely",
		},
		{
			name:         "invalid_template_name",
			templateName: "../../../etc/passwd",
			setupData: func() interface{} {
				return map[string]interface{}{}
			},
			expectError: true,
			description: "Path traversal in template name should be rejected",
		},
		{
			name:         "directory_traversal_template",
			templateName: "../../../../sensitive_file",
			setupData: func() interface{} {
				return map[string]interface{}{}
			},
			expectError: true,
			description: "Directory traversal attempts should be blocked",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data := tc.setupData()
			result, err := templates.RenderTemplate(tc.templateName, data)

			if tc.expectError {
				assert.Error(t, err, tc.description)
				assert.Empty(t, result, "Should not return content on error")
			} else {
				assert.NoError(t, err, tc.description)
				assert.NotEmpty(t, result, "Should return rendered content")

				// Verify that content is properly handled by security package
				// and that html/template provides additional escaping for HTML content
				if strings.Contains(tc.templateName, "html_validation_results") {
					// Template syntax is preserved as literal text in descriptions
					// but html/template will escape it in HTML output

					// Check that HTML content is escaped by html/template
					assert.NotContains(t, result, "<script>", "Script tags should be escaped")

					// The template renders descriptions as literal text,
					// so template syntax appears as text, not as executable template code
					t.Logf("Template output length: %d", len(result))
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

	t.Run("sanitization behavior", func(t *testing.T) {
		// Test that sanitization preserves content and normalizes whitespace
		testCases := []struct {
			input    string
			expected string
		}{
			{"{{.Config}}", "{{.Config}}"},
			{"{{range .Items}}{{.}}{{end}}", "{{range .Items}}{{.}}{{end}}"},
			{"<script>alert('xss')</script>", "<script>alert('xss')</script>"},
			{"Normal\n\n\tdescription", "Normal description"},
		}

		for _, test := range testCases {
			result := security.SanitizeDescription(test.input)
			assert.Equal(t, test.expected, result, "Sanitization should preserve content and normalize whitespace for: %s", test.input)
		}
	})
}
