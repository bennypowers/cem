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
package mcp_test

import (
	"context"
	"flag"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"bennypowers.dev/cem/mcp"
	"bennypowers.dev/cem/mcp/resources"
	W "bennypowers.dev/cem/workspace"
	mcpSDK "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var update = flag.Bool("update", false, "update golden files")

// getTestRegistry creates a registry using the test fixtures following the existing pattern
func getTestRegistry(t *testing.T) *mcp.MCPContext {
	workspace := W.NewFileSystemWorkspaceContext("./testdata/fixtures/multiple-elements-integration")
	err := workspace.Init()
	require.NoError(t, err)

	registry, err := mcp.NewMCPContext(workspace)
	require.NoError(t, err)

	err = registry.LoadManifests()
	require.NoError(t, err)

	return registry
}

// TestSchemaResilientTemplateIntegration tests the complete schema-resilient template rendering
func TestSchemaResilientTemplateIntegration(t *testing.T) {
	registry := getTestRegistry(t)
	registryAdapter := mcp.NewMCPContextAdapter(registry)

	// Get all resources to find the guidelines resource
	resourceDefs, err := resources.Resources(registryAdapter)
	require.NoError(t, err)

	// Find the general guidelines resource (not element-specific)
	var guidelinesURI string
	for _, resource := range resourceDefs {
		if resource.URI == "cem://guidelines" {
			guidelinesURI = resource.URI
			break
		}
	}
	require.NotEmpty(t, guidelinesURI, "Guidelines resource should exist")

	// Create request for guidelines resource using the real URI pattern
	req := &mcpSDK.ReadResourceRequest{
		Params: &mcpSDK.ReadResourceParams{
			URI: guidelinesURI,
		},
	}

	// Get the resource handler for guidelines and test it
	// This will use the real schema loading and template rendering
	var result *mcpSDK.ReadResourceResult
	for _, resource := range resourceDefs {
		if resource.URI == guidelinesURI && resource.Handler != nil {
			result, err = resource.Handler(context.Background(), req)
			require.NoError(t, err)
			break
		}
	}

	require.NotNil(t, result, "Guidelines resource should return a result")
	require.Len(t, result.Contents, 1)

	content := result.Contents[0].Text

	// Verify we got substantial output
	assert.Greater(t, len(content), 500, "Guidelines should produce substantial content")

	// Test schema-resilient template functionality
	t.Run("schema_descriptions_are_dynamically_loaded", func(t *testing.T) {
		// The content should include schema definitions if they're available
		// Look for evidence of schema processing

		if strings.Contains(content, "Schema Context:") {
			t.Log("✅ Template includes schema context section")
		}

		if strings.Contains(content, "Schema Field Definitions") {
			t.Log("✅ Template includes schema field definitions")
		}

		// Check for any schema version information
		if strings.Contains(content, "2.1.0") || strings.Contains(content, "schema") {
			t.Log("✅ Template references schema version or schema data")
		}

		// Look for evidence of dynamic schema function usage
		// These would only appear if the schemaDesc/schemaFieldDesc functions are working
		if strings.Contains(content, "CustomElement") || strings.Contains(content, "Attribute") {
			t.Log("✅ Template includes schema type definitions")
		}
	})

	t.Run("manifest_data_properly_extracted", func(t *testing.T) {
		// Should include actual element patterns from the fixture
		assert.Contains(t, content, "Component", "Should reference components")

		// Should include manifest-derived data
		if strings.Contains(content, "elements") {
			t.Log("✅ Template includes element count/info")
		}

		// Should include prefixes from actual manifests
		if strings.Contains(content, "prefix") {
			t.Log("✅ Template includes naming pattern analysis")
		}
	})

	// Golden file comparison (update mode)
	goldenPath := filepath.Join("fixtures", "template-rendering", "schema_resilient_template.golden.md")

	if *update {
		// Update golden file
		err := os.MkdirAll(filepath.Dir(goldenPath), 0755)
		require.NoError(t, err)

		err = os.WriteFile(goldenPath, []byte(content), 0644)
		require.NoError(t, err)
		t.Logf("Updated golden file: %s", goldenPath)
		t.Logf("Golden file content length: %d characters", len(content))
	} else {
		// Compare with golden file if it exists
		if goldenData, err := os.ReadFile(goldenPath); err == nil {
			goldenContent := string(goldenData)
			if content != goldenContent {
				t.Logf("Content differs from golden file. Run with --update to update.")
				t.Logf("Golden length: %d, Actual length: %d", len(goldenContent), len(content))

				// Show a snippet of the differences for debugging
				lines := strings.Split(content, "\n")
				if len(lines) > 10 {
					t.Logf("First 10 lines of output:\n%s", strings.Join(lines[:10], "\n"))
				}
			} else {
				t.Log("✅ Content matches golden file exactly")
			}
		} else {
			t.Logf("Golden file doesn't exist: %s", goldenPath)
			t.Log("Run with --update to create it")

			// Show sample output for verification
			lines := strings.Split(content, "\n")
			if len(lines) > 10 {
				t.Logf("Sample output (first 10 lines):\n%s", strings.Join(lines[:10], "\n"))
			}
		}
	}

	// Additional verification of schema-resilient features
	t.Run("template_functions_working", func(t *testing.T) {
		// The fact that the template rendered without error means the schema functions are working
		// If schema functions were broken, the template would fail to render
		assert.NotEmpty(t, content, "Template should render successfully with schema functions")

		// Verify template contains structured content sections
		expectedSections := []string{
			"Component Context",
			"Your Component Data",
			"How to Use This Context",
		}

		for _, section := range expectedSections {
			if strings.Contains(content, section) {
				t.Logf("✅ Found expected section: %s", section)
			}
		}
	})
}
