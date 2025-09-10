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
	"bennypowers.dev/cem/mcp/resources"
	"bennypowers.dev/cem/mcp/tools"
	"bennypowers.dev/cem/types"
	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// ServerInfo represents basic server information
type ServerInfo struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

// Server implements an MCP server for custom elements
type Server struct {
	workspace types.WorkspaceContext
	registry  *Registry
	server    *mcp.Server
}

// NewServer creates a new CEM MCP server
func NewServer(workspace types.WorkspaceContext) (*Server, error) {
	helpers.SafeDebugLog("Creating CEM MCP server for workspace: %s", workspace.Root())

	registry, err := NewRegistry(workspace)
	if err != nil {
		return nil, fmt.Errorf("failed to create registry: %w", err)
	}

	// Create MCP server
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "cem",
		Version: "1.0.0",
	}, nil)

	cemServer := &Server{
		workspace: workspace,
		registry:  registry,
		server:    server,
	}

	// Add tools to the server
	if err := cemServer.setupTools(); err != nil {
		return nil, fmt.Errorf("failed to setup tools: %w", err)
	}

	// Add resources to the server
	if err := cemServer.setupResources(); err != nil {
		return nil, fmt.Errorf("failed to setup resources: %w", err)
	}

	return cemServer, nil
}

// GetInfo returns server information
func (s *Server) GetInfo() ServerInfo {
	return ServerInfo{
		Name:        "cem",
		Version:     "1.0.0",
		Description: "Custom Elements Manifest MCP Server - provides intelligent context about custom elements for AI systems",
	}
}

// Run starts the MCP server with stdio transport
func (s *Server) Run(ctx context.Context) error {
	helpers.SafeDebugLog("Starting CEM MCP server with stdio transport")

	// Initialize the registry
	err := s.registry.LoadManifests()
	if err != nil {
		return fmt.Errorf("failed to load manifests: %w", err)
	}

	// Run the MCP server
	return s.server.Run(ctx, &mcp.StdioTransport{})
}

// setupTools adds tools to the MCP server
func (s *Server) setupTools() error {
	// Create registry adapter for tools
	registryAdapter := NewRegistryAdapter(s.registry)

	// Get tool definitions from tools package
	toolDefs, err := tools.Tools(registryAdapter)
	if err != nil {
		return fmt.Errorf("failed to load tool definitions: %w", err)
	}

	// Register each tool with the MCP server
	for _, toolDef := range toolDefs {
		// Convert input schema from map to *jsonschema.Schema
		var inputSchema *jsonschema.Schema
		if toolDef.InputSchema != nil {
			schemaBytes, err := json.Marshal(toolDef.InputSchema)
			if err != nil {
				return fmt.Errorf("failed to marshal input schema for tool %s: %w", toolDef.Name, err)
			}

			var schema jsonschema.Schema
			if err := json.Unmarshal(schemaBytes, &schema); err != nil {
				return fmt.Errorf("failed to unmarshal input schema for tool %s: %w", toolDef.Name, err)
			}
			inputSchema = &schema
		}

		s.server.AddTool(&mcp.Tool{
			Name:        toolDef.Name,
			Description: toolDef.Description,
			InputSchema: inputSchema,
		}, toolDef.Handler)
	}

	return nil
}

// setupResources adds resources to the MCP server
func (s *Server) setupResources() error {
	// Create registry adapter for resources
	registryAdapter := NewRegistryAdapter(s.registry)

	// Get resource definitions from resources package
	resourceDefs, err := resources.Resources(registryAdapter)
	if err != nil {
		return fmt.Errorf("failed to load resource definitions: %w", err)
	}

	// Register each resource with the MCP server
	for _, resourceDef := range resourceDefs {
		s.server.AddResource(&mcp.Resource{
			URI:         resourceDef.URI,
			Name:        resourceDef.Name,
			MIMEType:    resourceDef.MimeType,
			Description: resourceDef.Description,
		}, resourceDef.Handler)
	}

	return nil
}
