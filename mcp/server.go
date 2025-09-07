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

	W "bennypowers.dev/cem/workspace"
)

// MCPServer implements the MCP server interface for custom elements
type MCPServer struct {
	workspace   W.WorkspaceContext
	registry    *Registry
	initialized bool
}

// ServerInfo represents basic server information
type ServerInfo struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

// Resource represents an MCP resource
type Resource struct {
	URI      string `json:"uri"`
	Name     string `json:"name"`
	MimeType string `json:"mimeType"`
	Contents string `json:"contents"`
}

// Tool represents an MCP tool
type Tool struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	InputSchema map[string]interface{} `json:"inputSchema"`
}

// ToolResult represents the result of a tool call
type ToolResult struct {
	Content []map[string]interface{} `json:"content"`
}

// NewMCPServer creates a new MCP server instance
func NewMCPServer(workspace W.WorkspaceContext) (*MCPServer, error) {
	registry, err := NewRegistry(workspace)
	if err != nil {
		return nil, fmt.Errorf("failed to create registry: %w", err)
	}

	return &MCPServer{
		workspace: workspace,
		registry:  registry,
	}, nil
}

// Initialize initializes the MCP server
func (s *MCPServer) Initialize(ctx context.Context) error {
	err := s.registry.LoadManifests()
	if err != nil {
		return fmt.Errorf("failed to load manifests: %w", err)
	}

	s.initialized = true
	return nil
}

// GetServerInfo returns basic server information
func (s *MCPServer) GetServerInfo() ServerInfo {
	return ServerInfo{
		Name:        "cem",
		Version:     "1.0.0",
		Description: "Custom Elements Manifest MCP Server - provides intelligent context about custom elements for AI systems",
	}
}

// ListResources returns available MCP resources
func (s *MCPServer) ListResources(ctx context.Context) ([]Resource, error) {
	if !s.initialized {
		return nil, fmt.Errorf("server not initialized")
	}

	resources := []Resource{
		{
			URI:      "cem://schema",
			Name:     "Custom Elements Manifest Schema",
			MimeType: "application/json",
		},
		{
			URI:      "cem://registry",
			Name:     "Element Registry",
			MimeType: "application/json",
		},
	}

	// Add element-specific resources
	elements := s.registry.GetAllElements()
	for tagName := range elements {
		resources = append(resources, Resource{
			URI:      fmt.Sprintf("cem://element/%s", tagName),
			Name:     fmt.Sprintf("Element: %s", tagName),
			MimeType: "application/json",
		})
	}

	return resources, nil
}

// GetResource returns the content of a specific resource
func (s *MCPServer) GetResource(ctx context.Context, uri string) (*Resource, error) {
	if !s.initialized {
		return nil, fmt.Errorf("server not initialized")
	}

	switch uri {
	case "cem://schema":
		schema, err := s.registry.GetManifestSchema()
		if err != nil {
			return nil, fmt.Errorf("failed to get schema: %w", err)
		}

		contents, err := json.MarshalIndent(schema, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("failed to marshal schema: %w", err)
		}

		return &Resource{
			URI:      uri,
			Name:     "Custom Elements Manifest Schema",
			MimeType: "application/json",
			Contents: string(contents),
		}, nil

	case "cem://registry":
		elements := s.registry.GetAllElements()

		contents, err := json.MarshalIndent(elements, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("failed to marshal registry: %w", err)
		}

		return &Resource{
			URI:      uri,
			Name:     "Element Registry",
			MimeType: "application/json",
			Contents: string(contents),
		}, nil

	default:
		return nil, fmt.Errorf("resource not found: %s", uri)
	}
}

// ListTools returns available MCP tools
func (s *MCPServer) ListTools(ctx context.Context) ([]Tool, error) {
	if !s.initialized {
		return nil, fmt.Errorf("server not initialized")
	}

	tools := []Tool{
		{
			Name:        "validate_element",
			Description: "Validate custom element HTML usage against manifest definitions",
			InputSchema: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"tagName": map[string]interface{}{
						"type":        "string",
						"description": "The custom element tag name",
					},
					"html": map[string]interface{}{
						"type":        "string",
						"description": "The HTML to validate",
					},
				},
				"required": []string{"tagName", "html"},
			},
		},
		{
			Name:        "query_registry",
			Description: "Query the custom elements registry",
			InputSchema: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"tagName": map[string]interface{}{
						"type":        "string",
						"description": "Optional tag name to filter by",
					},
				},
			},
		},
	}

	return tools, nil
}

// CallTool executes a tool with the given arguments
func (s *MCPServer) CallTool(ctx context.Context, name string, args map[string]interface{}) (*ToolResult, error) {
	if !s.initialized {
		return nil, fmt.Errorf("server not initialized")
	}

	switch name {
	case "validate_element":
		return s.validateElement(ctx, args)
	case "query_registry":
		return s.queryRegistry(ctx, args)
	default:
		return nil, fmt.Errorf("unknown tool: %s", name)
	}
}

// validateElement validates custom element HTML usage
func (s *MCPServer) validateElement(ctx context.Context, args map[string]interface{}) (*ToolResult, error) {
	tagName, ok := args["tagName"].(string)
	if !ok {
		return nil, fmt.Errorf("tagName is required and must be a string")
	}

	_, ok = args["html"].(string)
	if !ok {
		return nil, fmt.Errorf("html is required and must be a string")
	}

	// Basic validation logic
	element, err := s.registry.GetElementInfo(tagName)
	if err != nil {
		return &ToolResult{
			Content: []map[string]interface{}{
				{
					"type": "text",
					"text": fmt.Sprintf("Element '%s' not found in registry", tagName),
				},
			},
		}, nil
	}

	// For now, just return basic element info
	return &ToolResult{
		Content: []map[string]interface{}{
			{
				"type": "text",
				"text": fmt.Sprintf("Element '%s' found. HTML validation would be performed here.", element.TagName),
			},
		},
	}, nil
}

// queryRegistry queries the custom elements registry
func (s *MCPServer) queryRegistry(ctx context.Context, args map[string]interface{}) (*ToolResult, error) {
	elements := s.registry.GetAllElements()

	if tagName, ok := args["tagName"].(string); ok {
		if element, exists := elements[tagName]; exists {
			data, err := json.MarshalIndent(element, "", "  ")
			if err != nil {
				return nil, fmt.Errorf("failed to marshal element: %w", err)
			}

			return &ToolResult{
				Content: []map[string]interface{}{
					{
						"type": "text",
						"text": string(data),
					},
				},
			}, nil
		} else {
			return &ToolResult{
				Content: []map[string]interface{}{
					{
						"type": "text",
						"text": fmt.Sprintf("Element '%s' not found", tagName),
					},
				},
			}, nil
		}
	}

	// Return all elements
	var tagNames []string
	for tagName := range elements {
		tagNames = append(tagNames, tagName)
	}

	data, err := json.MarshalIndent(tagNames, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal tag names: %w", err)
	}

	return &ToolResult{
		Content: []map[string]interface{}{
			{
				"type": "text",
				"text": fmt.Sprintf("Found %d elements: %s", len(tagNames), string(data)),
			},
		},
	}, nil
}
