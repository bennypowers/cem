package resources

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"bennypowers.dev/cem/mcp/templates"
	"bennypowers.dev/cem/mcp/types"
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

// loadManifestFromFixture loads a manifest from a fixture file and creates a ManifestContext
func loadManifestFromFixture(t *testing.T, version string) ManifestContext {
	// Load manifest fixture file
	manifestPath := filepath.Join("..", "fixtures", "multiple-schemas", version, "custom-elements.json")
	manifestData, err := os.ReadFile(manifestPath)
	require.NoError(t, err, "Should be able to read manifest fixture for version %s", version)

	// Parse manifest to extract basic information for the overview
	var manifestJSON map[string]any
	err = json.Unmarshal(manifestData, &manifestJSON)
	require.NoError(t, err, "Should be able to parse manifest JSON for version %s", version)

	// Extract basic info from manifest for context
	var elementCount int
	if modules, ok := manifestJSON["modules"].([]any); ok {
		for _, module := range modules {
			if moduleMap, ok := module.(map[string]any); ok {
				if declarations, ok := moduleMap["declarations"].([]any); ok {
					for _, decl := range declarations {
						if declMap, ok := decl.(map[string]any); ok {
							if kind, ok := declMap["kind"].(string); ok && kind == "class" {
								if _, hasTagName := declMap["tagName"]; hasTagName {
									elementCount++
								}
							}
						}
					}
				}
			}
		}
	}

	// Load schema for this version
	schemaData, err := V.GetSchema(version)
	require.NoError(t, err, "Should be able to load schema %s", version)

	// Parse schema
	var schema map[string]any
	err = json.Unmarshal(schemaData, &schema)
	require.NoError(t, err, "Should be able to parse schema %s", version)

	// Create context using actual manifest data but simplified for testing
	context := ManifestContext{
		Overview:          fmt.Sprintf("# Component Context for Schema %s Test\n\nTest context for schema version %s from fixture file with %d elements.", version, version, elementCount),
		SchemaVersion:     version,
		SchemaDefinitions: schema,
		Elements:          []types.ElementInfo{}, // Empty for now - focus on schema version compatibility
		ElementCount:      elementCount,
		SchemaVersions:    []string{version},
	}

	return context
}

func testSchemaVersionWithGolden(t *testing.T, version string) {
	// Load manifest from fixture file instead of mocking
	context := loadManifestFromFixture(t, version)

	// Render template
	result, err := templates.RenderTemplate("manifest-context", context)
	require.NoError(t, err, "Template should render successfully for schema %s", version)

	// Compare with golden file using fixture/golden pattern
	goldenFile := fmt.Sprintf("schema_version_%s.golden.md", strings.ReplaceAll(version, ".", "_"))
	testWithGoldenFile(t, result, goldenFile)
}

func testWithGoldenFile(t *testing.T, actual, goldenFileName string) {
	goldenPath := filepath.Join("..", "fixtures", "multiple-schemas", goldenFileName)

	// Check for update mode via environment variable (internal test pattern)
	if updateGolden := os.Getenv("UPDATE_GOLDEN"); updateGolden != "" {
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
	// Load manifest from fixture file to test real data handling
	context := loadManifestFromFixture(t, version)

	// Modify context for backwards compatibility testing
	context.Overview = fmt.Sprintf("# Component Context for Backwards Compatibility Test\n\nTesting schema %s with real manifest data.", version)

	// Should not fail even if some schema types don't exist in older versions
	result, err := templates.RenderTemplate("manifest-context", context)
	require.NoError(t, err, "Template should be backwards compatible with schema %s", version)

	// Compare with golden file using fixture/golden pattern
	goldenFile := fmt.Sprintf("backwards_compatibility_%s.golden.md", strings.ReplaceAll(version, ".", "_"))
	testBackwardsCompatibilityWithGoldenFile(t, result, goldenFile)
}

func testBackwardsCompatibilityWithGoldenFile(t *testing.T, actual, goldenFileName string) {
	goldenPath := filepath.Join("..", "fixtures", "schema-backwards-compatibility", goldenFileName)

	// Check for update mode via environment variable (internal test pattern)
	if updateGolden := os.Getenv("UPDATE_GOLDEN"); updateGolden != "" {
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
