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
	"encoding/json"
	"fmt"

	"bennypowers.dev/cem/lsp/helpers"
	W "bennypowers.dev/cem/workspace"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// SimpleCEMServer implements a simple MCP server for custom elements
type SimpleCEMServer struct {
	workspace W.WorkspaceContext
	registry  *Registry
	server    *mcp.Server
}

// NewSimpleCEMServer creates a new simple CEM MCP server
func NewSimpleCEMServer(workspace W.WorkspaceContext) (*SimpleCEMServer, error) {
	helpers.SafeDebugLog("Creating simple CEM MCP server for workspace: %s", workspace.Root())

	registry, err := NewRegistry(workspace)
	if err != nil {
		return nil, fmt.Errorf("failed to create registry: %w", err)
	}

	// Create MCP server
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "cem",
		Version: "1.0.0",
	}, nil)

	cemServer := &SimpleCEMServer{
		workspace: workspace,
		registry:  registry,
		server:    server,
	}

	// Add tools to the server
	cemServer.setupTools()

	return cemServer, nil
}

// GetInfo returns server information
func (s *SimpleCEMServer) GetInfo() ServerInfo {
	return ServerInfo{
		Name:        "cem",
		Version:     "1.0.0",
		Description: "Custom Elements Manifest MCP Server - provides intelligent context about custom elements for AI systems",
	}
}

// Run starts the MCP server with stdio transport
func (s *SimpleCEMServer) Run(ctx context.Context) error {
	helpers.SafeDebugLog("Starting simple CEM MCP server with stdio transport")

	// Initialize the registry
	err := s.registry.LoadManifests()
	if err != nil {
		return fmt.Errorf("failed to load manifests: %w", err)
	}

	// Run the MCP server
	return s.server.Run(ctx, &mcp.StdioTransport{})
}

// setupTools adds tools to the MCP server
func (s *SimpleCEMServer) setupTools() {
	// Add query_registry tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "query_registry",
		Description: "Query the custom elements registry",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args struct {
		TagName string `json:"tagName,omitempty" jsonschema:"Optional tag name to filter by"`
	}) (*mcp.CallToolResult, any, error) {
		return s.handleQueryRegistry(ctx, req, args)
	})

	// Add validate_element tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "validate_element",
		Description: "Validate custom element HTML usage against manifest definitions",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args struct {
		TagName string `json:"tagName" jsonschema:"The custom element tag name"`
		HTML    string `json:"html" jsonschema:"The HTML to validate"`
	}) (*mcp.CallToolResult, any, error) {
		return s.handleValidateElement(ctx, req, args)
	})

	// Add suggest_attributes tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "suggest_attributes",
		Description: "Get attribute suggestions for a custom element",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args struct {
		TagName string `json:"tagName" jsonschema:"The custom element tag name"`
	}) (*mcp.CallToolResult, any, error) {
		return s.handleSuggestAttributes(ctx, req, args)
	})
}

// Tool handlers

func (s *SimpleCEMServer) handleQueryRegistry(ctx context.Context, req *mcp.CallToolRequest, args struct {
	TagName string `json:"tagName,omitempty" jsonschema:"Optional tag name to filter by"`
}) (*mcp.CallToolResult, any, error) {
	elements := s.registry.GetAllElements()

	if args.TagName != "" {
		if element, exists := elements[args.TagName]; exists {
			data, err := json.MarshalIndent(element, "", "  ")
			if err != nil {
				return nil, nil, fmt.Errorf("failed to marshal element: %w", err)
			}

			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{Text: string(data)},
				},
			}, nil, nil
		} else {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{Text: fmt.Sprintf("Element '%s' not found", args.TagName)},
				},
			}, nil, nil
		}
	}

	// Return summary of all elements
	var tagNames []string
	for tagName := range elements {
		tagNames = append(tagNames, tagName)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("Found %d elements: %v", len(tagNames), tagNames)},
		},
	}, nil, nil
}

func (s *SimpleCEMServer) handleValidateElement(ctx context.Context, req *mcp.CallToolRequest, args struct {
	TagName string `json:"tagName" jsonschema:"The custom element tag name"`
	HTML    string `json:"html" jsonschema:"The HTML to validate"`
}) (*mcp.CallToolResult, any, error) {
	element, err := s.registry.GetElementInfo(args.TagName)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Element '%s' not found in registry", args.TagName)},
			},
		}, nil, nil
	}

	// Basic validation - in the future this could parse HTML and validate attributes/slots
	validationResult := fmt.Sprintf("Element '%s' found with %d attributes, %d slots, %d events. HTML validation would be performed here with: %s",
		element.TagName,
		len(element.Attributes()),
		len(element.Slots()),
		len(element.Events()),
		args.HTML)

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: validationResult},
		},
	}, nil, nil
}

func (s *SimpleCEMServer) handleSuggestAttributes(ctx context.Context, req *mcp.CallToolRequest, args struct {
	TagName string `json:"tagName" jsonschema:"The custom element tag name"`
}) (*mcp.CallToolResult, any, error) {
	element, err := s.registry.GetElementInfo(args.TagName)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Element '%s' not found in registry", args.TagName)},
			},
		}, nil, nil
	}

	attributes := element.Attributes()
	if len(attributes) == 0 {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Element '%s' has no attributes defined", args.TagName)},
			},
		}, nil, nil
	}

	// Build attribute suggestions
	suggestions := make([]map[string]interface{}, 0, len(attributes))
	for _, attr := range attributes {
		suggestion := map[string]interface{}{
			"name":        attr.Name(),
			"type":        attr.Type(),
			"description": attr.Description(),
			"required":    attr.Required(),
		}

		if attr.Default() != "" {
			suggestion["default"] = attr.Default()
		}

		if len(attr.Values()) > 0 {
			suggestion["values"] = attr.Values()
		}

		suggestions = append(suggestions, suggestion)
	}

	data, err := json.MarshalIndent(suggestions, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal attributes: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("Attributes for '%s':\n%s", args.TagName, string(data))},
		},
	}, nil, nil
}
