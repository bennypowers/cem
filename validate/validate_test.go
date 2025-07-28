package validate

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/nsf/jsondiff"
)

var update = flag.Bool("update", false, "update golden files")

func checkGolden(t *testing.T, name string, actual []byte) {
	t.Helper()
	goldenPath := filepath.Join("goldens", name+".json")
	if *update {
		if err := os.MkdirAll(filepath.Dir(goldenPath), 0755); err != nil {
			t.Fatalf("failed to create goldens directory: %v", err)
		}
		err := os.WriteFile(goldenPath, actual, 0644)
		if err != nil {
			t.Fatalf("failed to update golden file: %v", err)
		}
		return
	}

	expected, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("failed to read golden file %s: %v", goldenPath, err)
	}

	if string(expected) != string(actual) {
		options := jsondiff.DefaultConsoleOptions()
		diff, str := jsondiff.Compare(expected, actual, &options)
		if diff == jsondiff.FullMatch {
			t.Logf("Semantic match, string mismatch: %s", str)
		} else {
			t.Error(diff, str)
		}
	}
}

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
		{"invalid-attribute", "invalid_attribute"},
		{"invalid-slot", "invalid_slot"},
		{"invalid-css-part", "invalid_css_part"},
		{"invalid-css-custom-property", "invalid_css_custom_property"},
		{"unsupported-schema", "unsupported_schema"},
		{"warning-lifecycle-methods", "warning_lifecycle_methods"},
		{"warning-implementation-details", "warning_implementation_details"},
		{"warning-superclass-modules", "warning_superclass_modules"},
		{"warning-css-verbose", "warning_css_verbose"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Check if fixture exists
			fixturePath := filepath.Join("..", "cmd", "fixture", tc.fixture, "custom-elements.json")
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
			checkGolden(t, tc.name, jsonBytes)
		})
	}
}

func TestValidateWithDisabledWarnings(t *testing.T) {
	fixturePath := filepath.Join("..", "cmd", "fixture", "warning-lifecycle-methods", "custom-elements.json")
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
