package resources

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"bennypowers.dev/cem/mcp/templates"
	V "bennypowers.dev/cem/validate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMultipleSchemaVersions_FixtureGolden(t *testing.T) {
	// Test different schema versions to ensure template functions work across versions
	schemaVersions := []string{
		"2.1.0",
		"2.1.1-speculative",
		"2.0.0",
		"1.0.0",
	}

	for _, version := range schemaVersions {
		t.Run(version, func(t *testing.T) {
			testSchemaVersionWithGolden(t, version)
		})
	}
}

func testSchemaVersionWithGolden(t *testing.T, version string) {
	// Load schema for this version
	schemaData, err := V.GetSchema(version)
	require.NoError(t, err, "Should be able to load schema %s", version)

	// Parse schema
	var schema map[string]interface{}
	err = json.Unmarshal(schemaData, &schema)
	require.NoError(t, err, "Should be able to parse schema %s", version)

	// Create context with this schema version using minimal test data
	context := ManifestContext{
		Overview:          fmt.Sprintf("# Component Context for Schema %s Test\n\nTest context for schema version %s.", version, version),
		SchemaVersion:     version,
		SchemaDefinitions: schema,
		ElementCount:      1,
		// Add minimal patterns to trigger schema sections
		SlotPatterns:        []SlotPattern{{Name: "default", UsageCount: 1, Description: "Used in 1 elements"}},
		CSSProperties:       []string{"--test-color"},
		SchemaVersions:      []string{version},
	}

	// Render template
	result, err := templates.RenderTemplate("manifest_context", context)
	require.NoError(t, err, "Template should render successfully for schema %s", version)

	// Compare with golden file using fixture/golden pattern
	goldenFile := fmt.Sprintf("schema_version_%s.golden.md", strings.ReplaceAll(version, ".", "_"))
	testWithGoldenFile(t, result, goldenFile)
}

func testWithGoldenFile(t *testing.T, actual, goldenFileName string) {
	goldenPath := filepath.Join("testdata", goldenFileName)

	if os.Getenv("UPDATE_GOLDEN") == "1" {
		err := os.WriteFile(goldenPath, []byte(actual), 0644)
		require.NoError(t, err, "Failed to update golden file %s", goldenPath)
		t.Logf("Updated golden file: %s", goldenPath)
		return
	}

	expected, err := os.ReadFile(goldenPath)
	if os.IsNotExist(err) {
		t.Fatalf("Golden file %s does not exist. Run with UPDATE_GOLDEN=1 to create it.", goldenPath)
	}
	require.NoError(t, err, "Failed to read golden file %s", goldenPath)

	assert.Equal(t, string(expected), actual, "Content should match golden file %s", goldenFileName)
}

func TestSchemaVersionBackwardsCompatibility_FixtureGolden(t *testing.T) {
	// Test that templates work with minimal data for older schemas
	olderVersions := []string{"1.0.0", "2.0.0"}

	for _, version := range olderVersions {
		t.Run("backwards_compatibility_"+version, func(t *testing.T) {
			testBackwardsCompatibilityWithGolden(t, version)
		})
	}
}

func testBackwardsCompatibilityWithGolden(t *testing.T, version string) {
	schemaData, err := V.GetSchema(version)
	require.NoError(t, err)

	var schema map[string]interface{}
	err = json.Unmarshal(schemaData, &schema)
	require.NoError(t, err)

	// Minimal context with no patterns - tests that template handles empty data gracefully
	context := ManifestContext{
		Overview:          fmt.Sprintf("# Component Context for Backwards Compatibility Test\n\nTesting schema %s with minimal data.", version),
		SchemaVersion:     version,
		SchemaDefinitions: schema,
		ElementCount:      1,
		SchemaVersions:    []string{version},
		// Intentionally minimal - no patterns to test adaptive behavior
	}

	// Should not fail even if some schema types don't exist in older versions
	result, err := templates.RenderTemplate("manifest_context", context)
	require.NoError(t, err, "Template should be backwards compatible with schema %s", version)

	// Compare with golden file using fixture/golden pattern
	goldenFile := fmt.Sprintf("backwards_compatibility_%s.golden.md", strings.ReplaceAll(version, ".", "_"))
	testWithGoldenFile(t, result, goldenFile)
}