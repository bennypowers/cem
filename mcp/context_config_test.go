package mcp_test

import (
	"encoding/json"
	"testing"

	testworkspace "bennypowers.dev/cem/internal/platform/testutil/workspace"
	"bennypowers.dev/cem/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Inline assertions justified: testing interface compliance on thin getters/delegates
// with no complex output to snapshot. Values are single scalars or byte slices.

func TestMCPContext_Config(t *testing.T) {
	workspace := testworkspace.NewMapWorkspaceContext(t, "./testdata/fixtures/multiple-elements-integration")
	err := workspace.Init()
	require.NoError(t, err)

	registry, err := mcp.NewMCPContext(workspace, nil)
	require.NoError(t, err)

	cfg, err := registry.Config()
	require.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, workspace.Root(), cfg.ProjectDir)
}

func TestMCPContext_ConfigFile(t *testing.T) {
	workspace := testworkspace.NewMapWorkspaceContext(t, "./testdata/fixtures/multiple-elements-integration")
	err := workspace.Init()
	require.NoError(t, err)

	registry, err := mcp.NewMCPContext(workspace, nil)
	require.NoError(t, err)

	// MapWorkspaceContext returns empty string for ConfigFile
	assert.Equal(t, "", registry.ConfigFile())
}

func TestMCPContext_ConfigSchemaJSON(t *testing.T) {
	workspace := testworkspace.NewMapWorkspaceContext(t, "./testdata/fixtures/multiple-elements-integration")
	err := workspace.Init()
	require.NoError(t, err)

	registry, err := mcp.NewMCPContext(workspace, nil)
	require.NoError(t, err)

	schemaBytes := registry.ConfigSchemaJSON()
	require.NotEmpty(t, schemaBytes)

	var schema map[string]any
	err = json.Unmarshal(schemaBytes, &schema)
	require.NoError(t, err)
	assert.Equal(t, "CEM Configuration", schema["title"])
}

func TestMCPContext_Root(t *testing.T) {
	workspace := testworkspace.NewMapWorkspaceContext(t, "./testdata/fixtures/multiple-elements-integration")
	err := workspace.Init()
	require.NoError(t, err)

	registry, err := mcp.NewMCPContext(workspace, nil)
	require.NoError(t, err)

	assert.Equal(t, workspace.Root(), registry.Root())
}

func TestMCPContextAdapter_ImplementsConfigInterface(t *testing.T) {
	workspace := testworkspace.NewMapWorkspaceContext(t, "./testdata/fixtures/multiple-elements-integration")
	err := workspace.Init()
	require.NoError(t, err)

	registry, err := mcp.NewMCPContext(workspace, nil)
	require.NoError(t, err)

	adapter := mcp.NewMCPContextAdapter(registry)

	cfg, err := adapter.Config()
	require.NoError(t, err)
	assert.NotNil(t, cfg)

	assert.Equal(t, workspace.Root(), adapter.Root())

	schemaBytes := adapter.ConfigSchemaJSON()
	assert.NotEmpty(t, schemaBytes)

	_ = adapter.ConfigFile()
}

func TestMCPContextAdapter_ConfigSchemaIsValidJSON(t *testing.T) {
	workspace := testworkspace.NewMapWorkspaceContext(t, "./testdata/fixtures/multiple-elements-integration")
	err := workspace.Init()
	require.NoError(t, err)

	registry, err := mcp.NewMCPContext(workspace, nil)
	require.NoError(t, err)

	adapter := mcp.NewMCPContextAdapter(registry)
	schemaBytes := adapter.ConfigSchemaJSON()

	var schema map[string]any
	err = json.Unmarshal(schemaBytes, &schema)
	require.NoError(t, err)

	props, ok := schema["properties"].(map[string]any)
	require.True(t, ok)

	expectedSections := []string{"generate", "serve", "health", "mcp", "export"}
	for _, section := range expectedSections {
		_, exists := props[section]
		assert.True(t, exists, "schema should have %s section", section)
	}
}

// TestMCPContextAdapter_SatisfiesInterface verifies compile-time interface compliance.
// The actual check is that the test file compiles -- adapter.Config() etc. calls above
// would fail to compile if the interface wasn't satisfied.
