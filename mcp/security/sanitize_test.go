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
			name:        "html preserved",
			input:       "A button <script>alert('xss')</script> component",
			expected:    "A button <script>alert('xss')</script> component",
			description: "HTML is preserved (no XSS risk in MCP protocol)",
		},
		{
			name:        "template syntax preserved",
			input:       "A button component {{.Config}} for actions",
			expected:    "A button component {{.Config}} for actions",
			description: "Template syntax is preserved as literal text",
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
			expected:    strings.Repeat("A", 2000),
			description: "Very long descriptions should be hard-capped at maxLength",
		},
		{
			name:        "excessive whitespace",
			input:       "A   button    component\n\n\twith    whitespace",
			expected:    "A button component with whitespace",
			description: "Excessive whitespace should be normalized",
		},
		{
			name:        "utf8 safe truncation",
			input:       strings.Repeat("日本語", 1000), // 9 bytes per repetition, 9000 bytes total
			expected:    strings.Repeat("日本語", 222) + "..",  // 222*9 = 1998 bytes + ".." = 2000 bytes
			description: "UTF-8 truncation should not cut in middle of multi-byte character",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SanitizeDescription(test.input)
			assert.Equal(t, test.expected, result, test.description)
		})
	}
}

// Benchmark tests to ensure sanitization doesn't significantly impact performance
func BenchmarkSanitizeDescription(b *testing.B) {
	input := "A button component with some <script>alert('test')</script> content and lots of whitespace"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SanitizeDescription(input)
	}
}

// Integration test to verify sanitization works end-to-end
func TestSanitizationIntegration(t *testing.T) {
	testInputs := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "template syntax preserved",
			input:    "{{.Secrets}}",
			expected: "{{.Secrets}}",
		},
		{
			name:     "html preserved",
			input:    "<script>fetch('/admin/secrets').then(r=>r.text()).then(alert)</script>",
			expected: "<script>fetch('/admin/secrets').then(r=>r.text()).then(alert)</script>",
		},
		{
			name:     "whitespace normalized",
			input:    "A   button\n\n\twith   whitespace",
			expected: "A button with whitespace",
		},
	}

	for _, test := range testInputs {
		t.Run(test.name, func(t *testing.T) {
			result := SanitizeDescription(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}
