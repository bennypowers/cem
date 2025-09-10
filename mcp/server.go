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
	"fmt"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/types"
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
	cemServer.setupTools()

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
func (s *Server) setupTools() {
	// Add query_registry tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "query_registry",
		Description: `Query the custom elements registry.
		Any time the user is writing HTML,
		query the registry for available custom elements`,
	}, func(ctx context.Context, req *mcp.CallToolRequest, args struct {
		TagName string `json:"tagName,omitempty" jsonschema:"Optional tag name to filter by"`
		Filter  string `json:"filter,omitempty" jsonschema:"Filter by element capabilities"`
		Search  string `json:"search,omitempty" jsonschema:"Search term for elements"`
	}) (*mcp.CallToolResult, any, error) {
		queryArgs := QueryRegistryArgs{
			TagName: args.TagName,
			Filter:  args.Filter,
			Search:  args.Search,
		}
		return s.handleQueryRegistry(ctx, req, queryArgs)
	})

	// Add validate_element tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "validate_html",
		Description: `Validate all HTML, by checking the use of custom elements in the code
		against the custom elements manifest descriptions (i.e. guidelines) for those elements`,
	}, func(ctx context.Context, req *mcp.CallToolRequest, args struct {
		HTML    string `json:"html" jsonschema:"The HTML to validate"`
		Context string `json:"context,omitempty" jsonschema:"Context for validation"`
		TagName string `json:"tagName,omitempty" jsonschema:"Specific element to validate"`
	}) (*mcp.CallToolResult, any, error) {
		validateArgs := ValidateHtmlArgs{
			Html:    args.HTML,
			Context: args.Context,
			TagName: args.TagName,
		}
		return s.handleValidateHtml(ctx, req, validateArgs)
	})

	// Add suggest_attributes tool
	mcp.AddTool(s.server, &mcp.Tool{
		Name:        "suggest_attributes",
		Description: "Get attribute suggestions for a custom element",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args struct {
		TagName string `json:"tagName" jsonschema:"The custom element tag name"`
		Context string `json:"context,omitempty" jsonschema:"Context for suggestions"`
		Partial string `json:"partial,omitempty" jsonschema:"Partial attribute name to filter"`
	}) (*mcp.CallToolResult, any, error) {
		suggestArgs := SuggestAttributesArgs{
			TagName: args.TagName,
			Context: args.Context,
			Partial: args.Partial,
		}
		return s.handleSuggestAttributes(ctx, req, suggestArgs)
	})
}
