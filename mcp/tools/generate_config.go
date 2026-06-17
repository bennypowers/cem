package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	IC "bennypowers.dev/cem/internal/config"
	mcpTypes "bennypowers.dev/cem/mcp/types"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type GenerateConfigArgs struct {
	Focus string `json:"focus,omitempty"`
}

func handleGenerateConfig(
	_ context.Context,
	req *mcp.CallToolRequest,
	registry mcpTypes.MCPContext,
) (*mcp.CallToolResult, error) {
	var args GenerateConfigArgs
	if req.Params.Arguments != nil {
		if err := json.Unmarshal(req.Params.Arguments, &args); err != nil {
			return nil, fmt.Errorf("failed to parse args: %w", err)
		}
	}

	if args.Focus != "" && !IC.IsSchemaSection(args.Focus) {
		return nil, fmt.Errorf("unknown config section %q", args.Focus)
	}

	var b strings.Builder

	cfg, err := registry.Config()
	if err != nil || cfg == nil {
		b.WriteString("## Current config\n\nNo config file found.\n\n")
	} else {
		cfgJSON, err := json.MarshalIndent(cfg, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("failed to marshal config: %w", err)
		}
		fmt.Fprintf(&b, "## Current config\n\nFile: `%s`\n\n```json\n%s\n```\n\n", registry.ConfigFile(), string(cfgJSON))
	}

	schemaBytes := registry.ConfigSchemaJSON()
	var schema map[string]any
	if err := json.Unmarshal(schemaBytes, &schema); err != nil {
		return nil, fmt.Errorf("failed to parse config schema: %w", err)
	}
	if args.Focus != "" {
		if props, ok := schema["properties"].(map[string]any); ok {
			if section, ok := props[args.Focus]; ok {
				sectionJSON, err := json.MarshalIndent(section, "", "  ")
				if err != nil {
					return nil, fmt.Errorf("failed to marshal schema section: %w", err)
				}
				fmt.Fprintf(&b, "## Schema: %s\n\n```json\n%s\n```\n\n", args.Focus, string(sectionJSON))
			}
		}
	} else {
		schemaJSON, err := json.MarshalIndent(schema, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("failed to marshal schema: %w", err)
		}
		fmt.Fprintf(&b, "## Schema\n\n```json\n%s\n```\n\n", string(schemaJSON))
	}

	b.WriteString("## Guidance\n\n")
	writeGuidance(&b, args.Focus)

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: b.String()},
		},
	}, nil
}

var configGuidance = map[string][]string{
	"generate": {
		"Look for source files containing custom element definitions (classes extending HTMLElement, LitElement, etc.) to set `generate.files` globs.",
		"Check if `**/*.d.ts` files should be included (unusual); if so, set `generate.noDefaultExcludes: true`.",
		"Check package.json for a `customElements` field pointing to the manifest output path, or set `generate.output` explicitly.",
		"If the project uses design tokens (DTCG JSON format), set `generate.designTokens.spec` to the token file path.",
		"Look for demo HTML files to configure `generate.demoDiscovery.fileGlob` and URL patterns.",
	},
	"serve": {
		"If the project has bare import specifiers (e.g. `import 'lit'`), enable `serve.importMap.generate: true` to auto-resolve from node_modules.",
		"For dependencies not in node_modules or needing custom resolution, add entries to `serve.importMap.override.imports`.",
		"Check for .ts source files to enable `serve.transforms.typescript.enabled: true`.",
		"If components import CSS files as modules, enable `serve.transforms.css.enabled: true` and set include globs.",
		"Check if the project needs URL rewrites (e.g. mapping `/dist/` paths to `/src/`).",
		"Choose a demo rendering mode: light (default), shadow (encapsulated), iframe (isolated), or chromeless (no dev server UI).",
	},
	"health": {
		"Set `health.failBelow` to enforce a minimum documentation quality score (0-100).",
		"Disable checks not relevant to the project with `health.disable` (available: description, attributes, slots, css, events, demos).",
	},
	"mcp": {
		"Adjust `mcp.maxDescriptionLength` if element descriptions are being truncated (default: 2000 chars).",
	},
	"export": {
		"Configure framework exports (React, Vue, etc.) with output directory, package name, and prefix stripping.",
	},
}

func writeGuidance(b *strings.Builder, focus string) {
	if focus != "" {
		if prompts, ok := configGuidance[focus]; ok {
			fmt.Fprintf(b, "### %s\n\n", focus)
			for _, p := range prompts {
				fmt.Fprintf(b, "- %s\n", p)
			}
			b.WriteString("\n")
		} else {
			fmt.Fprintf(b, "No specific guidance for `%s`. See `cem://config/schema/%s` for available fields.\n\n", focus, focus)
		}
		return
	}
	for _, s := range IC.SchemaSections() {
		if prompts, ok := configGuidance[s.Key]; ok {
			fmt.Fprintf(b, "### %s\n\n", s.Key)
			for _, p := range prompts {
				fmt.Fprintf(b, "- %s\n", p)
			}
			b.WriteString("\n")
		}
	}
}

func makeGenerateConfigHandler(registry mcpTypes.MCPContext) mcp.ToolHandler {
	return func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return handleGenerateConfig(ctx, req, registry)
	}
}

func MakeGenerateConfigHandler(registry mcpTypes.MCPContext) mcp.ToolHandler {
	return makeGenerateConfigHandler(registry)
}
