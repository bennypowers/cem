package resources_test

import (
	"context"
	"encoding/json"
	"testing"

	IC "bennypowers.dev/cem/internal/config"
	"bennypowers.dev/cem/mcp"
	"bennypowers.dev/cem/mcp/resources"
	"bennypowers.dev/cem/mcp/types"
	mcpSDK "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Inline assertions justified: testing resource handler wiring and output structure,
// not fixed content. Checking for key presence in dynamic schema-derived output.

func findResource(t *testing.T, defs []types.ResourceDefinition, name string) *types.ResourceDefinition {
	t.Helper()
	for _, def := range defs {
		if def.Name == name {
			return &def
		}
	}
	t.Fatalf("resource %q not found", name)
	return nil
}

func TestConfigResource_Registered(t *testing.T) {
	registry := getTestRegistry(t)
	adapter := mcp.NewMCPContextAdapter(registry)

	defs, err := resources.Resources(adapter)
	require.NoError(t, err)

	names := make(map[string]bool)
	for _, d := range defs {
		names[d.Name] = true
	}

	assert.True(t, names["config"], "should register cem://config resource")
	assert.True(t, names["config-schema"], "should register cem://config/schema resource")
	assert.True(t, names["config-schema-section"], "should register cem://config/schema/{section} resource")
}

func TestConfigResource_ReturnsResolvedConfig(t *testing.T) {
	registry := getTestRegistry(t)
	adapter := mcp.NewMCPContextAdapter(registry)

	defs, err := resources.Resources(adapter)
	require.NoError(t, err)

	res := findResource(t, defs, "config")
	require.Equal(t, "cem://config", res.URI)
	require.Equal(t, "application/json", res.MimeType)

	result, err := res.Handler(context.Background(), &mcpSDK.ReadResourceRequest{
		Params: &mcpSDK.ReadResourceParams{URI: "cem://config"},
	})
	require.NoError(t, err)
	require.Len(t, result.Contents, 1)

	var cfg IC.CemConfig
	err = json.Unmarshal([]byte(result.Contents[0].Text), &cfg)
	require.NoError(t, err, "should return valid CemConfig JSON")
	assert.NotEmpty(t, cfg.ProjectDir)
}

func TestConfigSchemaResource_ReturnsOverview(t *testing.T) {
	registry := getTestRegistry(t)
	adapter := mcp.NewMCPContextAdapter(registry)

	defs, err := resources.Resources(adapter)
	require.NoError(t, err)

	res := findResource(t, defs, "config-schema")
	require.Equal(t, "cem://config/schema", res.URI)

	result, err := res.Handler(context.Background(), &mcpSDK.ReadResourceRequest{
		Params: &mcpSDK.ReadResourceParams{URI: "cem://config/schema"},
	})
	require.NoError(t, err)
	require.Len(t, result.Contents, 1)

	text := result.Contents[0].Text
	assert.Contains(t, text, "generate")
	assert.Contains(t, text, "serve")
	assert.Contains(t, text, "health")
	assert.Contains(t, text, "cem://config/schema/generate")
	assert.Contains(t, text, "cem://config/schema/serve")
}

func TestConfigSchemaSectionResource_ReturnsSection(t *testing.T) {
	registry := getTestRegistry(t)
	adapter := mcp.NewMCPContextAdapter(registry)

	defs, err := resources.Resources(adapter)
	require.NoError(t, err)

	res := findResource(t, defs, "config-schema-section")
	require.True(t, res.URITemplate)

	tests := []struct {
		section  string
		contains string
	}{
		{"generate", "files"},
		{"serve", "port"},
		{"health", "failBelow"},
		{"mcp", "maxDescriptionLength"},
		{"export", "output"},
	}

	for _, tc := range tests {
		t.Run(tc.section, func(t *testing.T) {
			result, err := res.Handler(context.Background(), &mcpSDK.ReadResourceRequest{
				Params: &mcpSDK.ReadResourceParams{URI: "cem://config/schema/" + tc.section},
			})
			require.NoError(t, err)
			require.Len(t, result.Contents, 1)

			var section map[string]any
			err = json.Unmarshal([]byte(result.Contents[0].Text), &section)
			require.NoError(t, err, "section should be valid JSON")

			props, ok := section["properties"].(map[string]any)
			if ok {
				assert.Contains(t, props, tc.contains)
			} else {
				// export uses additionalProperties
				assert.Contains(t, result.Contents[0].Text, tc.contains)
			}
		})
	}
}

func TestConfigSchemaSectionResource_UnknownSection(t *testing.T) {
	registry := getTestRegistry(t)
	adapter := mcp.NewMCPContextAdapter(registry)

	defs, err := resources.Resources(adapter)
	require.NoError(t, err)

	res := findResource(t, defs, "config-schema-section")

	_, err = res.Handler(context.Background(), &mcpSDK.ReadResourceRequest{
		Params: &mcpSDK.ReadResourceParams{URI: "cem://config/schema/nonexistent"},
	})
	assert.Error(t, err)
}
