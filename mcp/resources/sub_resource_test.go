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
package resources

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubResourceFiltering_FilterArrayByName(t *testing.T) {
	engine := NewPathTraversalEngine()

	// Test data - array of attributes
	testData := []any{
		map[string]any{
			"name":        "variant",
			"type":        "string",
			"description": "Button variant",
		},
		map[string]any{
			"name":        "disabled",
			"type":        "boolean",
			"description": "Disabled state",
		},
	}

	tests := []struct {
		name       string
		targetName string
		expectName string
		expectNil  bool
	}{
		{
			name:       "find existing attribute",
			targetName: "variant",
			expectName: "variant",
			expectNil:  false,
		},
		{
			name:       "find other existing attribute",
			targetName: "disabled",
			expectName: "disabled",
			expectNil:  false,
		},
		{
			name:       "non-existent attribute",
			targetName: "nonexistent",
			expectNil:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := engine.filterArrayByName(testData, tt.targetName)
			require.NoError(t, err, "Should not error")

			if tt.expectNil {
				assert.Nil(t, result, "Should return nil for non-existent item")
			} else {
				require.NotNil(t, result, "Should return result")
				resultMap, ok := result.(map[string]any)
				require.True(t, ok, "Result should be a map")
				assert.Equal(t, tt.expectName, resultMap["name"], "Should return correct item")
			}
		})
	}
}

func TestSubResourceAccess_URIParsing(t *testing.T) {
	tests := []struct {
		name           string
		actualURI      string
		templateURI    string
		expectedArgs   map[string]any
		expectedSubRes string
	}{
		{
			name:        "no sub-resource",
			actualURI:   "cem://element/button-element/attributes",
			templateURI: "cem://element/{tagName}/attributes",
			expectedArgs: map[string]any{
				"tagName": "button-element",
			},
			expectedSubRes: "",
		},
		{
			name:        "single sub-resource",
			actualURI:   "cem://element/button-element/attributes/variant",
			templateURI: "cem://element/{tagName}/attributes",
			expectedArgs: map[string]any{
				"tagName":     "button-element",
				"subResource": "variant",
			},
			expectedSubRes: "variant",
		},
		{
			name:        "nested sub-resource",
			actualURI:   "cem://element/button-element/css/parts/button",
			templateURI: "cem://element/{tagName}/css/parts",
			expectedArgs: map[string]any{
				"tagName":     "button-element",
				"subResource": "button",
			},
			expectedSubRes: "button",
		},
		{
			name:        "deep sub-resource path",
			actualURI:   "cem://element/button-element/attributes/variant/type",
			templateURI: "cem://element/{tagName}/attributes",
			expectedArgs: map[string]any{
				"tagName":     "button-element",
				"subResource": "variant/type",
			},
			expectedSubRes: "variant/type",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args, err := parseResourceURI(tt.actualURI, tt.templateURI, true)
			require.NoError(t, err, "URI parsing should not error")

			// Check tagName is extracted correctly
			assert.Equal(t, tt.expectedArgs["tagName"], args["tagName"], "tagName should be extracted")

			// Check sub-resource extraction
			if tt.expectedSubRes == "" {
				assert.NotContains(t, args, "subResource", "Should not have subResource for full collection")
			} else {
				assert.Equal(t, tt.expectedSubRes, args["subResource"], "subResource should be extracted correctly")
			}
		})
	}
}
