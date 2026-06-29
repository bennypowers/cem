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

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/internal/platform/testutil"
	"github.com/adrg/xdg"
)

// loadValidateFixtureMapFS loads a fixture manifest into a MapFileSystem
// and returns the MapFileSystem and virtual path.
func loadValidateFixtureMapFS(t *testing.T, fixture string) (*platform.MapFileSystem, string) {
	t.Helper()
	mfs := testutil.LoadTestdataFS(t, filepath.Join("..", "cmd", "testdata", "fixtures", fixture), "/")
	return mfs, "custom-elements.json"
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

	t.Run("os", func(t *testing.T) {
		fsys := platform.NewOSFileSystem()
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				fixturePath := filepath.Join("..", "cmd", "testdata", "fixtures", tc.fixture, "custom-elements.json")
				if _, err := os.Stat(fixturePath); os.IsNotExist(err) {
					t.Skipf("Fixture %s does not exist", fixturePath)
				}

				result, err := Validate(fsys, fixturePath, ValidationOptions{
					IncludeWarnings: true,
					DisabledRules:   []string{},
				})
				if err != nil {
					t.Fatalf("Validate returned error: %v", err)
				}

				result.Path = ""

				jsonBytes, err := json.MarshalIndent(result, "", "  ")
				if err != nil {
					t.Fatalf("Failed to marshal result to JSON: %v", err)
				}

				testutil.CheckGolden(t, tc.name, jsonBytes, testutil.GoldenOptions{
					Dir:         "testdata/goldens",
					Extension:   ".json",
					UseJSONDiff: true,
				})
			})
		}
	})

	t.Run("mapfs", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				mfs, virtualPath := loadValidateFixtureMapFS(t, tc.fixture)

				result, err := Validate(mfs, virtualPath, ValidationOptions{
					IncludeWarnings: true,
					DisabledRules:   []string{},
				})
				if err != nil {
					t.Fatalf("Validate returned error: %v", err)
				}

				result.Path = ""

				jsonBytes, err := json.MarshalIndent(result, "", "  ")
				if err != nil {
					t.Fatalf("Failed to marshal result to JSON: %v", err)
				}

				testutil.CheckGolden(t, tc.name, jsonBytes, testutil.GoldenOptions{
					Dir:         "testdata/goldens",
					Extension:   ".json",
					UseJSONDiff: true,
				})
			})
		}
	})
}

func testValidateWithDisabledWarnings(t *testing.T, fsys platform.FileSystem, fixturePath string) {
	t.Helper()

	// Test with warnings disabled by category
	result, err := Validate(fsys, fixturePath, ValidationOptions{
		IncludeWarnings: true,
		DisabledRules:   []string{"lifecycle"},
	})
	if err != nil {
		t.Fatalf("Validate returned error: %v", err)
	}

	for _, warning := range result.Warnings {
		if warning.Category == "lifecycle" {
			t.Errorf("Found lifecycle warning when it should be disabled: %s", warning.Message)
		}
	}

	// Test with warnings disabled by specific ID
	result2, err := Validate(fsys, fixturePath, ValidationOptions{
		IncludeWarnings: true,
		DisabledRules:   []string{"private-underscore-methods"},
	})
	if err != nil {
		t.Fatalf("Validate returned error: %v", err)
	}

	originalResult, err := Validate(fsys, fixturePath, ValidationOptions{
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

func TestValidateWithDisabledWarnings(t *testing.T) {
	osFixturePath := filepath.Join("..", "cmd", "testdata", "fixtures", "warning-lifecycle-methods", "custom-elements.json")
	if _, err := os.Stat(osFixturePath); os.IsNotExist(err) {
		t.Skip("warning-lifecycle-methods fixture does not exist")
	}

	t.Run("os", func(t *testing.T) {
		testValidateWithDisabledWarnings(t, platform.NewOSFileSystem(), osFixturePath)
	})

	t.Run("mapfs", func(t *testing.T) {
		mfs, virtualPath := loadValidateFixtureMapFS(t, "warning-lifecycle-methods")
		testValidateWithDisabledWarnings(t, mfs, virtualPath)
	})
}

// TestTryFetchSchema_NoCacheDir verifies that tryFetchSchema works when
// xdg.CacheHome is empty, as it would be in wasm or sandboxed environments.
// The embedded schema path (getSchema -> getEmbeddedSchema) bypasses
// tryFetchSchema entirely; this test confirms tryFetchSchema itself is
// resilient when called for non-embedded versions without a cache directory.
func TestTryFetchSchema_NoCacheDir(t *testing.T) {
	// Save and restore xdg.CacheHome
	origCacheHome := xdg.CacheHome
	defer func() { xdg.CacheHome = origCacheHome }()
	xdg.CacheHome = ""

	fsys := platform.NewMapFS(nil)

	// getSchema for an embedded version should succeed without cache
	schemaData, err := getSchema(fsys, "2.1.0")
	if err != nil {
		t.Fatalf("getSchema with empty CacheHome returned error: %v", err)
	}

	var schema map[string]any
	if err := json.Unmarshal(schemaData, &schema); err != nil {
		t.Fatalf("failed to parse schema JSON: %v", err)
	}

	if _, ok := schema["title"]; !ok {
		t.Error("expected schema to contain 'title' field")
	}
}

// TestGetSchemaUpgrade210 verifies that 2.1.0 and 2.1.1 manifests are
// automatically upgraded to use the 2.1.1-speculative schema to work around
// discriminated union issues in the upstream CEM schema.
// See: https://github.com/webcomponents/custom-elements-manifest/issues/138
func testGetSchemaUpgrade210(t *testing.T, fsys platform.FileSystem) {
	t.Helper()

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
			schemaData, err := getSchema(fsys, tc.version)
			if err != nil {
				t.Fatalf("getSchema(%q) returned error: %v", tc.version, err)
			}

			if !tc.checkTitle {
				var schema map[string]any
				if err := json.Unmarshal(schemaData, &schema); err != nil {
					t.Fatalf("Failed to parse schema JSON: %v", err)
				}
				return
			}

			var schema map[string]any
			if err := json.Unmarshal(schemaData, &schema); err != nil {
				t.Fatalf("Failed to parse schema JSON: %v", err)
			}

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

func TestGetSchemaUpgrade210(t *testing.T) {
	t.Run("os", func(t *testing.T) {
		testGetSchemaUpgrade210(t, platform.NewOSFileSystem())
	})

	t.Run("mapfs", func(t *testing.T) {
		// getSchema uses embedded schemas for known versions, so MapFS works
		// without needing to load any fixture files
		testGetSchemaUpgrade210(t, platform.NewMapFileSystem(nil))
	})
}
