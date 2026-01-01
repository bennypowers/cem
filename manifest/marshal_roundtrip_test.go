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
package manifest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// TestMarshalUnmarshalRoundTrip tests that we can unmarshal a comprehensive JSON
// manifest and then marshal it back to JSON without losing any data.
func TestMarshalUnmarshalRoundTrip(t *testing.T) {
	// Load the comprehensive golden file
	goldenPath := filepath.Join("testdata", "marshal-roundtrip-comprehensive.json")
	originalJSON, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("Failed to read golden file: %v", err)
	}

	// Unmarshal into Package struct
	var pkg Package
	if err := json.Unmarshal(originalJSON, &pkg); err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Marshal back to JSON
	roundTripJSON, err := json.Marshal(&pkg)
	if err != nil {
		t.Fatalf("Failed to marshal back to JSON: %v", err)
	}

	// Parse both JSONs for comparison (to handle formatting differences)
	var originalData, roundTripData interface{}

	if err := json.Unmarshal(originalJSON, &originalData); err != nil {
		t.Fatalf("Failed to parse original JSON: %v", err)
	}

	if err := json.Unmarshal(roundTripJSON, &roundTripData); err != nil {
		t.Fatalf("Failed to parse round-trip JSON: %v", err)
	}

	// Compare the data structures
	if !deepEqualIgnoreOmitEmpty(originalData, roundTripData) {
		t.Errorf("Round-trip failed.\nOriginal JSON:\n%s\n\nRound-trip JSON:\n%s",
			string(originalJSON), string(roundTripJSON))
	}
}

// TestMarshalUnmarshalSpecificCases tests specific edge cases and deprecation forms
func TestMarshalUnmarshalSpecificCases(t *testing.T) {
	testCases := []struct {
		name     string
		jsonData string
	}{
		{
			name: "PackageDeprecatedBool",
			jsonData: `{
				"schemaVersion": "2.1.0",
				"deprecated": true,
				"modules": []
			}`,
		},
		{
			name: "PackageDeprecatedReason",
			jsonData: `{
				"schemaVersion": "2.1.0", 
				"deprecated": "Use new-package instead",
				"modules": []
			}`,
		},
		{
			name: "FunctionParameterDeprecated",
			jsonData: `{
				"schemaVersion": "2.1.0",
				"modules": [{
					"kind": "javascript-module",
					"path": "test.js",
					"declarations": [{
						"kind": "function",
						"name": "testFn",
						"parameters": [{
							"name": "param1",
							"deprecated": true,
							"type": {"text": "string"}
						}, {
							"name": "param2", 
							"deprecated": "Use param3 instead",
							"type": {"text": "number"}
						}]
					}],
					"exports": []
				}]
			}`,
		},
		{
			name: "EmptySlicesNotNil",
			jsonData: `{
				"schemaVersion": "2.1.0",
				"modules": [{
					"kind": "javascript-module",
					"path": "test.js"
				}]
			}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Unmarshal
			var pkg Package
			if err := json.Unmarshal([]byte(tc.jsonData), &pkg); err != nil {
				t.Fatalf("Failed to unmarshal: %v", err)
			}

			// Marshal back
			roundTripJSON, err := json.Marshal(&pkg)
			if err != nil {
				t.Fatalf("Failed to marshal: %v", err)
			}

			// Parse for comparison
			var originalData, roundTripData interface{}
			if err := json.Unmarshal([]byte(tc.jsonData), &originalData); err != nil {
				t.Fatalf("Failed to parse original: %v", err)
			}
			if err := json.Unmarshal(roundTripJSON, &roundTripData); err != nil {
				t.Fatalf("Failed to parse round-trip: %v", err)
			}

			// Verify round-trip equality
			if !deepEqualIgnoreOmitEmpty(originalData, roundTripData) {
				t.Errorf("Round-trip failed for %s.\nOriginal: %s\nRound-trip: %s",
					tc.name, tc.jsonData, string(roundTripJSON))
			}

			// For empty slices test, ensure slices are not nil after unmarshal
			if tc.name == "EmptySlicesNotNil" {
				if pkg.Modules[0].Declarations == nil {
					t.Error("Declarations slice should not be nil after unmarshal")
				}
				if pkg.Modules[0].Exports == nil {
					t.Error("Exports slice should not be nil after unmarshal")
				}
			}
		})
	}
}

// deepEqualIgnoreOmitEmpty compares JSON structures while ignoring empty arrays/slices
// that would be omitted due to omitempty tags
func deepEqualIgnoreOmitEmpty(a, b interface{}) bool {
	// Normalize both structures to handle omitempty differences
	aNorm := normalizeOmitEmpty(a)
	bNorm := normalizeOmitEmpty(b)

	// Marshal to get canonical representation
	aJSON, err := json.Marshal(aNorm)
	if err != nil {
		return false
	}
	bJSON, err := json.Marshal(bNorm)
	if err != nil {
		return false
	}

	return string(aJSON) == string(bJSON)
}

// normalizeOmitEmpty recursively removes empty slices and arrays to simulate omitempty behavior
func normalizeOmitEmpty(v interface{}) interface{} {
	switch val := v.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{})
		for k, v := range val {
			normalized := normalizeOmitEmpty(v)
			// Only include non-empty values
			if !isEmpty(normalized) {
				result[k] = normalized
			}
		}
		return result
	case []interface{}:
		if len(val) == 0 {
			return nil // Empty slices become nil to match omitempty
		}
		result := make([]interface{}, len(val))
		for i, v := range val {
			result[i] = normalizeOmitEmpty(v)
		}
		return result
	default:
		return val
	}
}

// isEmpty checks if a value should be considered empty for omitempty purposes
func isEmpty(v interface{}) bool {
	if v == nil {
		return true
	}
	switch val := v.(type) {
	case []interface{}:
		return len(val) == 0
	case map[string]interface{}:
		return len(val) == 0
	case string:
		return val == ""
	default:
		return false
	}
}
