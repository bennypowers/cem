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
package security

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeDescription(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		description string
	}{
		{
			name:        "clean description",
			input:       "A button component for primary actions",
			expected:    "A button component for primary actions",
			description: "Normal description should pass through unchanged",
		},
		{
			name:        "template injection - variable access",
			input:       "A button component {{.Config}} for actions",
			expected:    "A button component for actions",
			description: "Template variable access should be removed",
		},
		{
			name:        "template injection - range construct",
			input:       "{{range .Items}}{{.Name}}{{end}} Available items",
			expected:    "Available items",
			description: "Range constructs should be removed",
		},
		{
			name:        "template injection - with construct",
			input:       "{{with .Settings}}{{.Theme}}{{end}} button styling",
			expected:    "button styling",
			description: "With constructs should be removed",
		},
		{
			name:        "template injection - define construct",
			input:       "{{define \"malicious\"}}evil{{end}} Safe content",
			expected:    "evil Safe content",
			description: "Define constructs should be removed but content preserved",
		},
		{
			name:        "template injection - template construct",
			input:       "{{template \"inject\" .}} Some content",
			expected:    "Some content",
			description: "Template constructs should be removed",
		},
		{
			name:        "template injection - block construct",
			input:       "{{block \"malicious\" .}}default{{end}} content",
			expected:    "default content",
			description: "Block constructs should be removed but content preserved",
		},
		{
			name:        "html injection",
			input:       "A button <script>alert('xss')</script> component",
			expected:    "A button &lt;script&gt;alert(&#39;xss&#39;)&lt;/script&gt; component",
			description: "HTML should be escaped",
		},
		{
			name:        "multiple injection attempts",
			input:       "{{.Secret}} A button <script>evil()</script> {{range .Items}}{{.}}{{end}}",
			expected:    "A button &lt;script&gt;evil()&lt;/script&gt;",
			description: "Multiple injection attempts should all be sanitized",
		},
		{
			name:        "empty description",
			input:       "",
			expected:    "",
			description: "Empty description should remain empty",
		},
		{
			name:        "very long description",
			input:       strings.Repeat("A", 3000),
			expected:    strings.Repeat("A", 2000) + "...",
			description: "Very long descriptions should be truncated",
		},
		{
			name:        "excessive whitespace",
			input:       "A   button    component\n\n\twith    whitespace",
			expected:    "A button component with whitespace",
			description: "Excessive whitespace should be normalized",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SanitizeDescription(test.input)
			assert.Equal(t, test.expected, result, test.description)
		})
	}
}

func TestSanitizeDescriptionPreservingMarkdown(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		description string
	}{
		{
			name:        "markdown code blocks",
			input:       "A button component with ```javascript\nconst x = 1;\n``` code example",
			expected:    "A button component with ```javascript const x = 1; ``` code example",
			description: "Markdown code blocks should be preserved but whitespace normalized",
		},
		{
			name:        "inline code",
			input:       "A button with `className` prop",
			expected:    "A button with `className` prop",
			description: "Inline code should be preserved",
		},
		{
			name:        "markdown with template injection",
			input:       "A button `{{.Config}}` with ```{{range .}}{{.}}{{end}}``` injection",
			expected:    "A button `` with `````` injection",
			description: "Template injection should be removed even with markdown",
		},
		{
			name:        "html in markdown",
			input:       "A button with `<div>safe</div>` and <script>evil</script>",
			expected:    "A button with `<div>safe</div>` and &lt;script&gt;evil&lt;/script&gt;",
			description: "HTML outside code should be escaped, inside code preserved",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SanitizeDescriptionPreservingMarkdown(test.input)
			assert.Equal(t, test.expected, result, test.description)
		})
	}
}

func TestDetectTemplateInjection(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    bool
		description string
	}{
		{
			name:        "clean text",
			input:       "A normal button component description",
			expected:    false,
			description: "Normal text should not be flagged",
		},
		{
			name:        "variable access",
			input:       "A button {{.Config}} component",
			expected:    true,
			description: "Variable access should be detected",
		},
		{
			name:        "range construct",
			input:       "{{range .Items}}item{{end}}",
			expected:    true,
			description: "Range constructs should be detected",
		},
		{
			name:        "with construct",
			input:       "{{with .Data}}value{{end}}",
			expected:    true,
			description: "With constructs should be detected",
		},
		{
			name:        "empty input",
			input:       "",
			expected:    false,
			description: "Empty input should not be flagged",
		},
		{
			name:        "nested braces",
			input:       "JavaScript: if (x) { return {y: 1}; }",
			expected:    false,
			description: "Regular braces should not be flagged",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := DetectTemplateInjection(test.input)
			assert.Equal(t, test.expected, result, test.description)
		})
	}
}

func TestGetInjectionPatterns(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    []string
		description string
	}{
		{
			name:        "variable access",
			input:       "{{.Config}} test",
			expected:    []string{"variable_access"},
			description: "Should detect variable access pattern",
		},
		{
			name:        "range construct",
			input:       "{{range .Items}}{{.}}{{end}}",
			expected:    []string{"range_construct", "variable_access"},
			description: "Should detect range and variable access patterns",
		},
		{
			name:        "multiple patterns",
			input:       "{{with .Data}}{{range .Items}}{{template \"x\" .}}{{end}}{{end}}",
			expected:    []string{"range_construct", "with_construct", "template_construct"},
			description: "Should detect multiple patterns",
		},
		{
			name:        "dollar variable",
			input:       "{{$var := .Config}}{{$var}}",
			expected:    []string{"dollar_variable"},
			description: "Should detect dollar variable pattern",
		},
		{
			name:        "clean text",
			input:       "Normal description",
			expected:    []string{},
			description: "Should detect no patterns in clean text",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := GetInjectionPatterns(test.input)
			assert.ElementsMatch(t, test.expected, result, test.description)
		})
	}
}

func TestSanitizeWithPolicy(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		policy      SecurityPolicy
		expected    string
		description string
	}{
		{
			name:        "default policy",
			input:       "A button {{.Config}} with `code` and <script>evil</script>",
			policy:      DefaultSecurityPolicy(),
			expected:    "A button with `code` and &lt;script&gt;evil&lt;/script&gt;",
			description: "Default policy should preserve markdown and escape HTML",
		},
		{
			name:        "strict policy",
			input:       "A button {{.Config}} with `code` and <script>evil</script>",
			policy:      StrictSecurityPolicy(),
			expected:    "A button with `code` and &lt;script&gt;evil&lt;/script&gt;",
			description: "Strict policy should escape everything",
		},
		{
			name:        "short length limit",
			input:       "A very long description that exceeds the policy limit",
			policy:      SecurityPolicy{MaxDescriptionLength: 10, AllowMarkdown: false, StrictMode: true},
			expected:    "A very lon...",
			description: "Should respect custom length limits",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SanitizeWithPolicy(test.input, test.policy)
			assert.Equal(t, test.expected, result, test.description)
		})
	}
}

func TestSecurityPolicies(t *testing.T) {
	t.Run("default policy", func(t *testing.T) {
		policy := DefaultSecurityPolicy()
		assert.Equal(t, 2000, policy.MaxDescriptionLength)
		assert.True(t, policy.AllowMarkdown)
		assert.False(t, policy.StrictMode)
	})

	t.Run("strict policy", func(t *testing.T) {
		policy := StrictSecurityPolicy()
		assert.Equal(t, 500, policy.MaxDescriptionLength)
		assert.False(t, policy.AllowMarkdown)
		assert.True(t, policy.StrictMode)
	})
}

// Benchmark tests to ensure sanitization doesn't significantly impact performance
func BenchmarkSanitizeDescription(b *testing.B) {
	input := "A button component {{.Config}} with some <script>alert('test')</script> content"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SanitizeDescription(input)
	}
}

func BenchmarkDetectTemplateInjection(b *testing.B) {
	input := "A button component {{.Config}} with {{range .Items}}{{.}}{{end}} injection"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DetectTemplateInjection(input)
	}
}

// Integration test to verify sanitization works end-to-end
func TestTemplateInjectionMitigation(t *testing.T) {
	maliciousInputs := []string{
		"{{.Secrets}}",
		"{{range .Config}}{{.}}{{end}}",
		"{{with .Database}}{{.Password}}{{end}}",
		"{{template \"../../../etc/passwd\" .}}",
		"{{define \"malicious\"}}{{.}}{{end}}{{template \"malicious\" .}}",
		"<script>fetch('/admin/secrets').then(r=>r.text()).then(alert)</script>",
	}

	for _, input := range maliciousInputs {
		t.Run("mitigate_"+input[:min(20, len(input))], func(t *testing.T) {
			// Test basic sanitization
			sanitized := SanitizeDescription(input)
			assert.NotContains(t, sanitized, "{{", "Template syntax should be removed")
			assert.NotContains(t, sanitized, "}}", "Template syntax should be removed")

			// Test detection - only check for template injection if input contains template syntax
			if strings.Contains(input, "{{") {
				assert.True(t, DetectTemplateInjection(input), "Template injection should be detected")
			}
			assert.False(t, DetectTemplateInjection(sanitized), "Sanitized text should not trigger detection")

			// Test that HTML is escaped
			if strings.Contains(input, "<script>") {
				assert.Contains(t, sanitized, "&lt;script&gt;", "Script tags should be escaped")
			}
		})
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
