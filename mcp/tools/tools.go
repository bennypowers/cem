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
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"gopkg.in/yaml.v3"

	"bennypowers.dev/cem/mcp/types"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Tools returns all available tool definitions with their handlers
func Tools(registry types.MCPContext) ([]types.ToolDefinition, error) {
	// Get the directory where this source file is located
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("failed to get current file path")
	}
	toolsDir := filepath.Dir(filename)

	var toolDefs []types.ToolDefinition

	// Find all markdown files in the tools directory
	files, err := filepath.Glob(filepath.Join(toolsDir, "*.md"))
	if err != nil {
		return nil, fmt.Errorf("failed to find tool markdown files: %w", err)
	}

	for _, file := range files {
		toolDef, err := parseToolDefinition(file, registry)
		if err != nil {
			return nil, fmt.Errorf("failed to parse tool definition from %s: %w", file, err)
		}
		toolDefs = append(toolDefs, toolDef)
	}

	return toolDefs, nil
}

// parseToolDefinition parses a markdown file with frontmatter into a ToolDefinition
func parseToolDefinition(filename string, registry types.MCPContext) (types.ToolDefinition, error) {
	content, err := os.ReadFile(filename)
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
	// Check if this is a declarative tool (has data fetchers)
	if len(toolDef.DataFetchers) > 0 {
		config := DeclarativeToolConfig{
			Name:         toolDef.Name,
			DataFetchers: toolDef.DataFetchers,
			Template:     toolDef.Template,
			ResponseType: toolDef.ResponseType,
		}
		return MakeDeclarativeHandler(registry, config), nil
	}

	// Fallback to legacy handlers for non-declarative tools
	switch toolDef.Name {
	case "validate_html":
		return makeValidateHtmlHandler(registry), nil
	case "suggest_attributes":
		return makeSuggestAttributesHandler(registry), nil
	case "generate_html":
		return makeGenerateHtmlHandler(registry), nil
	case "suggest_css_integration":
		return makeSuggestCssIntegrationHandler(registry), nil
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

func makeSuggestAttributesHandler(registry types.MCPContext) mcp.ToolHandler {
	return func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return handleSuggestAttributes(ctx, req, registry)
	}
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

func makeSuggestCssIntegrationHandler(registry types.MCPContext) mcp.ToolHandler {
	return func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return handleSuggestCssIntegration(ctx, req, registry)
	}
}

// MakeElementDetailsHandler is the exported version for testing - now uses declarative framework
func MakeElementDetailsHandler(registry types.MCPContext) mcp.ToolHandler {
	config := DeclarativeToolConfig{
		Name: "element_details",
		DataFetchers: []types.DataFetcher{
			{Name: "element", Type: "manifest_element", Path: "", Required: true},
			{Name: "schema", Type: "schema_definitions", Path: "definitions", Required: false},
		},
		Template:     "",
		ResponseType: "",
	}
	return MakeDeclarativeHandler(registry, config)
}

// MakeElementAttributesHandler is the exported version for testing - now uses declarative framework
func MakeElementAttributesHandler(registry types.MCPContext) mcp.ToolHandler {
	config := DeclarativeToolConfig{
		Name: "element_attributes",
		DataFetchers: []types.DataFetcher{
			{Name: "element", Type: "manifest_element", Path: "", Required: true},
			{Name: "schema", Type: "schema_definitions", Path: "definitions", Required: false},
			{Name: "attributes", Type: "attribute_collection", Path: "", Required: true},
		},
		Template:     "",
		ResponseType: "",
	}
	return MakeDeclarativeHandler(registry, config)
}

// MakeElementSlotsHandler is the exported version for testing - now uses declarative framework
func MakeElementSlotsHandler(registry types.MCPContext) mcp.ToolHandler {
	config := DeclarativeToolConfig{
		Name: "element_slots",
		DataFetchers: []types.DataFetcher{
			{Name: "element", Type: "manifest_element", Path: "", Required: true},
			{Name: "schema", Type: "schema_definitions", Path: "definitions", Required: false},
		},
		Template:     "",
		ResponseType: "",
	}
	return MakeDeclarativeHandler(registry, config)
}

// MakeElementEventsHandler is the exported version for testing - now uses declarative framework
func MakeElementEventsHandler(registry types.MCPContext) mcp.ToolHandler {
	config := DeclarativeToolConfig{
		Name: "element_events",
		DataFetchers: []types.DataFetcher{
			{Name: "element", Type: "manifest_element", Path: "", Required: true},
			{Name: "schema", Type: "schema_definitions", Path: "definitions", Required: false},
		},
		Template:     "",
		ResponseType: "",
	}
	return MakeDeclarativeHandler(registry, config)
}

// MakeElementStylingHandler is the exported version for testing - now uses declarative framework
func MakeElementStylingHandler(registry types.MCPContext) mcp.ToolHandler {
	config := DeclarativeToolConfig{
		Name: "element_styling",
		DataFetchers: []types.DataFetcher{
			{Name: "element", Type: "manifest_element", Path: "", Required: true},
			{Name: "schema", Type: "schema_definitions", Path: "definitions", Required: false},
		},
		Template:     "",
		ResponseType: "",
	}
	return MakeDeclarativeHandler(registry, config)
}

// Argument types for testing compatibility
type ElementDetailsArgs struct {
	TagName string `json:"tagName"`
	Context string `json:"context,omitempty"`
}

type ElementAttributesArgs struct {
	TagName string `json:"tagName"`
	Context string `json:"context,omitempty"`
}

type ElementSlotsArgs struct {
	TagName string `json:"tagName"`
	Context string `json:"context,omitempty"`
}

type ElementEventsArgs struct {
	TagName string `json:"tagName"`
	Context string `json:"context,omitempty"`
}

type ElementStylingArgs struct {
	TagName string `json:"tagName"`
	Context string `json:"context,omitempty"`
}
