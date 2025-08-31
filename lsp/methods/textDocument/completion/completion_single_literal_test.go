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
package completion_test

import (
	"testing"

	"bennypowers.dev/cem/lsp/methods/textDocument/completion"
	M "bennypowers.dev/cem/manifest"
)

// TestSingleLiteralTypeCompletion tests completion for single literal types like 'promo'
// This handles the edge case where a property has a single literal type that could
// potentially become a union in the future (e.g., 'promo' -> 'promo' | 'standard')
func TestSingleLiteralTypeCompletion(t *testing.T) {
	tests := []struct {
		name                string
		attributeType       string
		expectedCompletions []string
		description         string
	}{
		{
			name:                "Single quoted literal type",
			attributeType:       "'promo'",
			expectedCompletions: []string{"promo"},
			description:         "Should provide completion for single literal type with single quotes",
		},
		{
			name:                "Single double-quoted literal type",
			attributeType:       "\"promo\"",
			expectedCompletions: []string{"promo"},
			description:         "Should provide completion for single literal type with double quotes",
		},
		{
			name:                "Union with single quoted literals",
			attributeType:       "'promo' | 'standard'",
			expectedCompletions: []string{"promo", "standard"},
			description:         "Should provide completions for union of single-quoted literals",
		},
		{
			name:                "Union with mixed quotes",
			attributeType:       "'promo' | \"standard\" | 'featured'",
			expectedCompletions: []string{"promo", "standard", "featured"},
			description:         "Should handle mixed quote styles in union types",
		},
		{
			name:                "Complex union type",
			attributeType:       "\"red\" | \"green\" | \"blue\"",
			expectedCompletions: []string{"red", "green", "blue"},
			description:         "Should handle traditional union type with double quotes",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test attribute with the specified type
			attr := &M.Attribute{
				FullyQualified: M.FullyQualified{Name: "variant"},
				Type:           &M.Type{Text: tt.attributeType},
			}

			// Get type-based completions
			completions := completion.GetTypeBasedCompletions(attr)

			// Extract completion labels
			var labels []string
			for _, completion := range completions {
				labels = append(labels, completion.Label)
			}

			// Check that we got the expected completions
			if len(labels) != len(tt.expectedCompletions) {
				t.Errorf("Expected %d completions, got %d. Expected: %v, Got: %v",
					len(tt.expectedCompletions), len(labels), tt.expectedCompletions, labels)
			}

			// Check that all expected completions are present
			for _, expected := range tt.expectedCompletions {
				found := false
				for _, label := range labels {
					if label == expected {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected completion '%s' not found. Got: %v", expected, labels)
				}
			}

			t.Logf("Test: %s", tt.description)
			t.Logf("Type: %s -> Completions: %v", tt.attributeType, labels)
		})
	}
}
