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
package tools

import (
	"context"
	"embed"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"

	"bennypowers.dev/cem/mcp/types"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

//go:embed *.md
var toolDefinitions embed.FS

// Tools returns all available tool definitions with their handlers
func Tools(registry types.MCPContext) ([]types.ToolDefinition, error) {
	var toolDefs []types.ToolDefinition

	// Read all markdown files from embedded filesystem
	entries, err := toolDefinitions.ReadDir(".")
	if err != nil {
		return nil, fmt.Errorf("failed to read embedded tool definitions: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		toolDef, err := parseToolDefinition(entry.Name(), registry)
		if err != nil {
			return nil, fmt.Errorf("failed to parse tool definition from %s: %w", entry.Name(), err)
		}
		toolDefs = append(toolDefs, toolDef)
	}

	return toolDefs, nil
}

// parseToolDefinition parses a markdown file with frontmatter into a ToolDefinition
func parseToolDefinition(filename string, registry types.MCPContext) (types.ToolDefinition, error) {
	content, err := toolDefinitions.ReadFile(filename)
	if err != nil {
		return types.ToolDefinition{}, fmt.Errorf("failed to read file: %w", err)
	}

	// Split frontmatter and markdown content
	parts := strings.SplitN(string(content), "---", 3)
	if len(parts) < 3 {
		return types.ToolDefinition{}, fmt.Errorf("invalid markdown format: missing frontmatter")
	}

	// Parse YAML frontmatter
	var frontmatter types.Frontmatter
	if err := yaml.Unmarshal([]byte(parts[1]), &frontmatter); err != nil {
		return types.ToolDefinition{}, fmt.Errorf("failed to parse frontmatter: %w", err)
	}

	// Extract description from markdown content
	description := strings.TrimSpace(parts[2])

	// Create tool definition first
	toolDef := types.ToolDefinition{
		Name:         frontmatter.Name,
		Description:  description,
		InputSchema:  frontmatter.InputSchema,
		DataFetchers: frontmatter.DataFetchers,
		Template:     frontmatter.Template,
		ResponseType: frontmatter.ResponseType,
	}

	// Get the corresponding handler
	handler, err := getToolHandler(toolDef, registry)
	if err != nil {
		return types.ToolDefinition{}, fmt.Errorf("failed to get handler for tool %s: %w", frontmatter.Name, err)
	}

	// Set the handler and return the complete tool definition
	toolDef.Handler = handler
	return toolDef, nil
}

// getToolHandler returns the appropriate handler function for a tool
func getToolHandler(toolDef types.ToolDefinition, registry types.MCPContext) (mcp.ToolHandler, error) {
	// Only core action tools remain - progressive disclosure is handled by resources
	switch toolDef.Name {
	case "validate_html":
		return makeValidateHtmlHandler(registry), nil
	case "generate_html":
		return makeGenerateHtmlHandler(registry), nil
	default:
		return nil, fmt.Errorf("unknown tool: %s", toolDef.Name)
	}
}

// Handler factory functions
func makeValidateHtmlHandler(registry types.MCPContext) mcp.ToolHandler {
	return func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return handleValidateHtml(ctx, req, registry)
	}
}

// MakeValidateHtmlHandler is the exported version for testing
func MakeValidateHtmlHandler(registry types.MCPContext) mcp.ToolHandler {
	return makeValidateHtmlHandler(registry)
}

func makeGenerateHtmlHandler(registry types.MCPContext) mcp.ToolHandler {
	return func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return handleGenerateHtml(ctx, req, registry)
	}
}

// MakeGenerateHtmlHandler is the exported version for testing
func MakeGenerateHtmlHandler(registry types.MCPContext) mcp.ToolHandler {
	return makeGenerateHtmlHandler(registry)
}
