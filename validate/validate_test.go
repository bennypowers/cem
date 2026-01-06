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
package validate

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
)

func TestValidateGolden(t *testing.T) {
	testCases := []struct {
		fixture string
		name    string
	}{
		{"valid-manifest", "valid_manifest"},
		{"invalid-class", "invalid_class"},
		{"invalid-mixin", "invalid_mixin"},
		{"invalid-variable", "invalid_variable"},
		{"invalid-function", "invalid_function"},
		{"invalid-field", "invalid_field"},
		{"invalid-method", "invalid_method"},
		{"invalid-event", "invalid_event"},
		{"event-without-type", "event_without_type"},
		{"invalid-attribute", "invalid_attribute"},
		{"invalid-slot", "invalid_slot"},
		{"invalid-css-part", "invalid_css_part"},
		{"invalid-css-custom-property", "invalid_css_custom_property"},
		{"unsupported-schema", "unsupported_schema"},
		{"warning-lifecycle-methods", "warning_lifecycle_methods"},
		{"warning-implementation-details", "warning_implementation_details"},
		{"warning-superclass-modules", "warning_superclass_modules"},
		{"warning-css-verbose", "warning_css_verbose"},
		{"schema-upgrade-2-1-0", "schema_upgrade_2_1_0"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Check if fixture exists
			fixturePath := filepath.Join("..", "cmd", "testdata", "fixtures", tc.fixture, "custom-elements.json")
			if _, err := os.Stat(fixturePath); os.IsNotExist(err) {
				t.Skipf("Fixture %s does not exist", fixturePath)
			}

			// Validate the manifest
			result, err := Validate(fixturePath, ValidationOptions{
				IncludeWarnings: true,
				DisabledRules:   []string{},
			})
			if err != nil {
				t.Fatalf("Validate returned error: %v", err)
			}

			// Clear the Path field since it contains absolute paths that vary by system
			result.Path = ""

			// Marshal to JSON with consistent formatting
			jsonBytes, err := json.MarshalIndent(result, "", "  ")
			if err != nil {
				t.Fatalf("Failed to marshal result to JSON: %v", err)
			}

			// Check against golden file
			testutil.CheckGolden(t, tc.name, jsonBytes, testutil.GoldenOptions{
				Dir:         "testdata/goldens",
				Extension:   ".json",
				UseJSONDiff: true,
			})
		})
	}
}

func TestValidateWithDisabledWarnings(t *testing.T) {
	fixturePath := filepath.Join("..", "cmd", "testdata", "fixtures", "warning-lifecycle-methods", "custom-elements.json")
	if _, err := os.Stat(fixturePath); os.IsNotExist(err) {
		t.Skip("warning-lifecycle-methods fixture does not exist")
	}

	// Test with warnings disabled by category
	result, err := Validate(fixturePath, ValidationOptions{
		IncludeWarnings: true,
		DisabledRules:   []string{"lifecycle"},
	})
	if err != nil {
		t.Fatalf("Validate returned error: %v", err)
	}

	// Should have no lifecycle warnings
	for _, warning := range result.Warnings {
		if warning.Category == "lifecycle" {
			t.Errorf("Found lifecycle warning when it should be disabled: %s", warning.Message)
		}
	}

	// Test with warnings disabled by specific ID
	result2, err := Validate(fixturePath, ValidationOptions{
		IncludeWarnings: true,
		DisabledRules:   []string{"private-underscore-methods"},
	})
	if err != nil {
		t.Fatalf("Validate returned error: %v", err)
	}

	// Should have fewer warnings than the original (which had all warnings)
	originalResult, err := Validate(fixturePath, ValidationOptions{
		IncludeWarnings: true,
		DisabledRules:   []string{},
	})
	if err != nil {
		t.Fatalf("Validate returned error: %v", err)
	}

	if len(result2.Warnings) >= len(originalResult.Warnings) {
		t.Errorf("Expected fewer warnings when disabling specific rule, got %d vs %d", len(result2.Warnings), len(originalResult.Warnings))
	}
}

// TestGetSchemaUpgrade210 verifies that 2.1.0 and 2.1.1 manifests are
// automatically upgraded to use the 2.1.1-speculative schema to work around
// discriminated union issues in the upstream CEM schema.
// See: https://github.com/webcomponents/custom-elements-manifest/issues/138
func TestGetSchemaUpgrade210(t *testing.T) {
	testCases := []struct {
		version    string
		expected   string
		checkTitle bool // Only 2.1.x schemas have title field
	}{
		{"2.1.0", "2.1.1-speculative", true},
		{"2.1.1", "2.1.1-speculative", true},
		{"2.1.1-speculative", "2.1.1-speculative", true},
		{"2.0.0", "2.0.0", false},
		{"1.0.0", "1.0.0", false},
	}

	for _, tc := range testCases {
		t.Run(tc.version, func(t *testing.T) {
			schemaData, err := getSchema(tc.version)
			if err != nil {
				t.Fatalf("getSchema(%q) returned error: %v", tc.version, err)
			}

			if !tc.checkTitle {
				// For older schemas, just verify we got valid JSON
				var schema map[string]interface{}
				if err := json.Unmarshal(schemaData, &schema); err != nil {
					t.Fatalf("Failed to parse schema JSON: %v", err)
				}
				return
			}

			// Parse the schema to verify which version was actually loaded
			var schema map[string]interface{}
			if err := json.Unmarshal(schemaData, &schema); err != nil {
				t.Fatalf("Failed to parse schema JSON: %v", err)
			}

			// Check the title field which contains the schema version (2.1.x only)
			title, ok := schema["title"].(string)
			if !ok {
				t.Fatalf("Schema missing 'title' field")
			}

			expectedTitle := "Custom Elements Manifest Schema " + tc.expected
			if title != expectedTitle {
				t.Errorf("getSchema(%q) loaded schema with title %q, expected %q", tc.version, title, expectedTitle)
			}
		})
	}
}
