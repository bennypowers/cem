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
package mcp

import (
	"context"
	"testing"

	"bennypowers.dev/cem/mcp/resources"
	W "bennypowers.dev/cem/workspace"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewServer(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("./test-fixtures/basic")

	server, err := NewServer(workspace)
	require.NoError(t, err)
	require.NotNil(t, server)

	assert.Equal(t, workspace, server.workspace)
	assert.NotNil(t, server.registry)
	assert.NotNil(t, server.server)
}

func TestServer_GetInfo(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("./test-fixtures/basic")
	server, err := NewServer(workspace)
	require.NoError(t, err)

	info := server.GetInfo()

	assert.Equal(t, "cem", info.Name)
	assert.Equal(t, "1.0.0", info.Version)
	assert.Contains(t, info.Description, "Custom Elements Manifest")
}

func TestDynamicSchemaVersionDetection(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("./test-fixtures/basic")

	// Initialize workspace
	err := workspace.Init()
	require.NoError(t, err)

	// Create registry and load manifests
	registry, err := NewMCPContext(workspace)
	require.NoError(t, err)

	err = registry.LoadManifests()
	require.NoError(t, err)

	// Create registry adapter
	registryAdapter := NewMCPContextAdapter(registry)

	// Test GetManifestSchemaVersions method
	versions := registryAdapter.GetManifestSchemaVersions()
	t.Logf("Detected schema versions: %v", versions)

	// Should detect the schema version from the test fixture
	// The test fixture in ./test-fixtures/basic has schemaVersion "2.1.1"
	assert.NotEmpty(t, versions, "Should detect at least one schema version")

	if len(versions) > 0 {
		// Verify it's a valid version string
		assert.NotEmpty(t, versions[0], "Schema version should not be empty")
		t.Logf("First detected version: %s", versions[0])
	}

	// Test the schema resource with dynamic detection
	resourceDefs, err := resources.Resources(registryAdapter)
	require.NoError(t, err)

	// Find the schema resource
	var schemaResourceDef *resources.ResourceDefinition
	for _, resourceDef := range resourceDefs {
		if resourceDef.Name == "schema" {
			schemaResourceDef = &resourceDef
			break
		}
	}

	require.NotNil(t, schemaResourceDef, "Schema resource should be available")

	// Test the schema resource handler
	req := &mcp.ReadResourceRequest{
		Params: &mcp.ReadResourceParams{
			URI: "cem://schema",
		},
	}

	result, err := schemaResourceDef.Handler(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Len(t, result.Contents, 1)

	// Verify we get a valid JSON schema response
	content := result.Contents[0]
	assert.Equal(t, "cem://schema", content.URI)
	assert.Equal(t, "application/json", content.MIMEType)
	assert.NotEmpty(t, content.Text, "Schema content should not be empty")
	assert.Contains(t, content.Text, "json-schema.org", "Should contain valid JSON schema")

	t.Logf("Schema resource returned %d bytes of JSON schema", len(content.Text))
}
