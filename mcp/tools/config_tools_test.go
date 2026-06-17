package tools_test

import (
	"context"
	"encoding/json"
	"testing"

	"bennypowers.dev/cem/mcp/tools"
	mcpSDK "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Inline assertions justified: testing tool handler wiring and output structure,
// not fixed content. Checking for key presence in dynamic schema-derived output.

func TestGenerateConfig_Registered(t *testing.T) {
	registry := getTestRegistry(t)

	toolDefs, err := tools.Tools(registry)
	require.NoError(t, err)

	names := make(map[string]bool)
	for _, d := range toolDefs {
		names[d.Name] = true
	}
	assert.True(t, names["generate_config"], "should register generate_config tool")
	assert.True(t, names["validate_config"], "should register validate_config tool")
}

func TestGenerateConfig_ReturnsSchemaAndGuidance(t *testing.T) {
	registry := getTestRegistry(t)

	handler := tools.MakeGenerateConfigHandler(registry)
	require.NotNil(t, handler)

	result, err := handler(context.Background(), &mcpSDK.CallToolRequest{
		Params: &mcpSDK.CallToolParamsRaw{
			Name:      "generate_config",
			Arguments: json.RawMessage(`{}`),
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, result.Content)

	text := result.Content[0].(*mcpSDK.TextContent).Text
	assert.Contains(t, text, "generate")
	assert.Contains(t, text, "serve")
	assert.Contains(t, text, "schema")
}

func TestGenerateConfig_WithFocus(t *testing.T) {
	registry := getTestRegistry(t)

	handler := tools.MakeGenerateConfigHandler(registry)

	result, err := handler(context.Background(), &mcpSDK.CallToolRequest{
		Params: &mcpSDK.CallToolParamsRaw{
			Name:      "generate_config",
			Arguments: json.RawMessage(`{"focus": "serve"}`),
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, result.Content)

	text := result.Content[0].(*mcpSDK.TextContent).Text
	assert.Contains(t, text, "serve")
	assert.Contains(t, text, "import")
}

func TestGenerateConfig_InvalidFocus(t *testing.T) {
	registry := getTestRegistry(t)

	handler := tools.MakeGenerateConfigHandler(registry)

	_, err := handler(context.Background(), &mcpSDK.CallToolRequest{
		Params: &mcpSDK.CallToolParamsRaw{
			Name:      "generate_config",
			Arguments: json.RawMessage(`{"focus": "nonexistent"}`),
		},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unknown config section")
}

func TestValidateConfig_NoConfigFile(t *testing.T) {
	registry := getTestRegistry(t)

	handler := tools.MakeValidateConfigHandler(registry)
	require.NotNil(t, handler)

	result, err := handler(context.Background(), &mcpSDK.CallToolRequest{
		Params: &mcpSDK.CallToolParamsRaw{
			Name:      "validate_config",
			Arguments: json.RawMessage(`{}`),
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, result.Content)

	text := result.Content[0].(*mcpSDK.TextContent).Text
	assert.Contains(t, text, "no config file")
}
