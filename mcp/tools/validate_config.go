package tools

import (
	"context"
	"encoding/json"
	"fmt"

	IC "bennypowers.dev/cem/internal/config"
	"bennypowers.dev/cem/internal/sourcepos"
	mcpTypes "bennypowers.dev/cem/mcp/types"
	"github.com/modelcontextprotocol/go-sdk/mcp"

	DD "bennypowers.dev/cem/generate/demodiscovery"
	"bennypowers.dev/cem/serve/middleware/transform"
)

type validateConfigResult struct {
	ConfigFile string                 `json:"configFile"`
	Valid      bool                   `json:"valid"`
	Errors     []IC.ValidationError   `json:"errors"`
	Warnings   []IC.ValidationError   `json:"warnings"`
}

func handleValidateConfig(
	_ context.Context,
	_ *mcp.CallToolRequest,
	registry mcpTypes.MCPContext,
) (*mcp.CallToolResult, error) {
	cfgFile := registry.ConfigFile()
	if cfgFile == "" {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "no config file found"},
			},
		}, nil
	}

	rawData, err := registry.FileSystem().ReadFile(cfgFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	cfgFormat := IC.FormatFromPath(cfgFile)
	schemaErrs := IC.ValidateSchema(rawData, cfgFormat)

	cfg, err := registry.Config()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	semanticErrs := IC.Validate(cfg, IC.ValidateOptions{
		CheckIO: true,
		Root:    registry.Root(),
		IO:      registry.FileSystem(),
		ValidateURLRewrites: func(rewrites []IC.URLRewrite) error {
			return transform.ValidateURLRewrites(rewrites)
		},
		ValidateDemoDiscovery: func(c *IC.CemConfig) error {
			return DD.ValidateDemoDiscoveryConfig(c, map[string]string{})
		},
	})

	allErrs := IC.DeduplicateErrors(schemaErrs, semanticErrs)

	if len(allErrs) > 0 {
		posMap := sourcepos.BuildPositionMap(rawData, cfgFormat)
		for i := range allErrs {
			pointer := sourcepos.FieldToJSONPointer(allErrs[i].Field)
			if pos, ok := sourcepos.Resolve(posMap, pointer); ok {
				allErrs[i].Line = pos.Line
				allErrs[i].Column = pos.Column
			}
		}
	}

	var errors, warnings []IC.ValidationError
	for _, e := range allErrs {
		if e.Severity == IC.SeverityWarning {
			warnings = append(warnings, e)
		} else {
			errors = append(errors, e)
		}
	}
	if errors == nil {
		errors = []IC.ValidationError{}
	}
	if warnings == nil {
		warnings = []IC.ValidationError{}
	}

	result := validateConfigResult{
		ConfigFile: cfgFile,
		Valid:      len(errors) == 0,
		Errors:     errors,
		Warnings:   warnings,
	}

	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal results: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(data)},
		},
	}, nil
}

func makeValidateConfigHandler(registry mcpTypes.MCPContext) mcp.ToolHandler {
	return func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return handleValidateConfig(ctx, req, registry)
	}
}

func MakeValidateConfigHandler(registry mcpTypes.MCPContext) mcp.ToolHandler {
	return makeValidateConfigHandler(registry)
}
