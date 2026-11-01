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

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/internal/version"
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

// ServerConfig contains configuration options for the MCP server
type ServerConfig struct {
	MaxDescriptionLength int
	// Additional packages to load manifests from (URLs, npm:, or jsr: specifiers)
	AdditionalPackages []string
	// FileSystem to use for filesystem operations. If nil, defaults to OSFileSystem.
	FileSystem platform.FileSystem
}

// Server implements an MCP server for custom elements
type Server struct {
	workspace types.WorkspaceContext
	registry  *MCPContext
	server    *mcp.Server
	config    ServerConfig
}

// NewServer creates a new CEM MCP server with default configuration
func NewServer(workspace types.WorkspaceContext) (*Server, error) {
	return NewServerWithConfig(workspace, ServerConfig{
		MaxDescriptionLength: 2000,
	})
}

// NewServerWithConfig creates a new CEM MCP server with custom configuration
func NewServerWithConfig(workspace types.WorkspaceContext, config ServerConfig) (*Server, error) {
	helpers.SafeDebugLog("Creating CEM MCP server for workspace: %s", workspace.Root())

	registry, err := NewMCPContext(workspace, config.FileSystem)
	if err != nil {
		return nil, fmt.Errorf("failed to create registry: %w", err)
	}

	// Create MCP server
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "cem",
		Version: version.GetVersion(),
	}, nil)

	cemServer := &Server{
		workspace: workspace,
		registry:  registry,
		server:    server,
		config:    config,
	}

	// Registration is shared with Register so both entry points expose the
	// same tools and resources. Unlike Register, this constructor defers
	// manifest loading to Run, preserving the existing lifecycle.
	if err := register(server, registry); err != nil {
		return nil, err
	}

	return cemServer, nil
}

// Register adds cem's tools and resources to an existing MCP server.
//
// Register must be called before server starts serving sessions: the SDK's
// registration methods mutate the server's handler tables.
//
// The caller owns the server and its lifecycle: Register does not start a
// transport or file watching. It creates its own MCPContext, loads the
// workspace manifests and any configured additional packages before
// registering handlers, so the mounted tools and resources are usable as
// soon as Register returns. The workspace must already be initialized.
//
// Tool names and resource URIs are registered as-is, replacing any existing
// entry with the same name or URI on server. Additional package failures are
// logged but do not fail registration, matching Run. If registration fails
// partway, server may already hold some cem tools or resources; callers
// should discard it.
//
// Callers that need manifest reloads or file watching on their own server
// cannot get them from Register. NewServerWithConfig and Run provide that
// lifecycle, but they own the SDK server and use stdio.
func Register(server *mcp.Server, workspace types.WorkspaceContext, config ServerConfig) error {
	if server == nil {
		return fmt.Errorf("server must not be nil")
	}
	if workspace == nil {
		return fmt.Errorf("workspace must not be nil")
	}

	registry, err := NewMCPContext(workspace, config.FileSystem)
	if err != nil {
		return fmt.Errorf("failed to create registry: %w", err)
	}

	if err := loadRegistryData(registry, config.AdditionalPackages, "config"); err != nil {
		return err
	}

	// Register does not watch for changes. NewMCPContext creates the file
	// watcher eagerly, so close it rather than hold an inotify descriptor and
	// goroutine for the server's lifetime.
	if err := registry.LSPRegistry().StopFileWatching(); err != nil {
		helpers.SafeDebugLog("Warning: Could not stop file watching: %v", err)
	}

	return register(server, registry)
}

// loadRegistryData loads manifests and additional packages into registry.
// Additional package failures are logged and do not fail the load.
func loadRegistryData(registry *MCPContext, additionalPackages []string, source string) error {
	if err := registry.LoadManifests(); err != nil {
		return fmt.Errorf("failed to load manifests: %w", err)
	}

	if len(additionalPackages) > 0 {
		helpers.SafeDebugLog("Loading %d additional packages from %s", len(additionalPackages), source)
		if err := registry.LSPRegistry().LoadAdditionalPackages(additionalPackages); err != nil {
			helpers.SafeDebugLog("Warning: Error loading additional packages: %v", err)
		}
	}

	return nil
}

// register adds every cem tool and resource to server using registry.
func register(server *mcp.Server, registry *MCPContext) error {
	if err := registerTools(server, registry); err != nil {
		return fmt.Errorf("failed to register tools: %w", err)
	}
	if err := registerResources(server, registry); err != nil {
		return fmt.Errorf("failed to register resources: %w", err)
	}
	return nil
}

// GetInfo returns server information
func (s *Server) GetInfo() ServerInfo {
	return ServerInfo{
		Name:        "cem",
		Version:     version.GetVersion(),
		Description: "Custom Elements Manifest MCP Server - provides intelligent context about custom elements for AI systems",
	}
}

// Registry returns the MCP registry for testing
// This is used by tests to verify stdout cleanliness
func (s *Server) Registry() *MCPContext {
	return s.registry
}

// Run starts the MCP server with stdio transport
func (s *Server) Run(ctx context.Context) error {
	helpers.SafeDebugLog("Starting CEM MCP server with stdio transport")

	// Initialize the registry
	if err := loadRegistryData(s.registry, s.config.AdditionalPackages, "CLI args"); err != nil {
		return err
	}

	// Start file watching for manifest and config changes
	if err := s.registry.LSPRegistry().StartFileWatching(func() {
		helpers.SafeDebugLog("File change detected, reloading manifests and config")
		s.registry.InvalidateConfig()
		if err := s.registry.LoadManifests(); err != nil {
			helpers.SafeDebugLog("Warning: Failed to reload manifests: %v", err)
		}
	}); err != nil {
		helpers.SafeDebugLog("Warning: Could not start file watching: %v", err)
	}

	// Run the MCP server
	return s.server.Run(ctx, &mcp.StdioTransport{})
}

// registerTools adds tools to the MCP server
func registerTools(server *mcp.Server, registry *MCPContext) error {
	// Create registry adapter for tools
	registryAdapter := NewMCPContextAdapter(registry)

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

		server.AddTool(&mcp.Tool{
			Name:        toolDef.Name,
			Description: toolDef.Description,
			InputSchema: inputSchema,
		}, toolDef.Handler)
	}

	return nil
}

// registerResources adds resources to the MCP server
func registerResources(server *mcp.Server, registry *MCPContext) error {
	// Create registry adapter for resources
	registryAdapter := NewMCPContextAdapter(registry)

	// Get resource definitions from resources package
	resourceDefs, err := resources.Resources(registryAdapter)
	if err != nil {
		return fmt.Errorf("failed to load resource definitions: %w", err)
	}

	// Register each resource with the MCP server
	for _, resourceDef := range resourceDefs {
		if resourceDef.URITemplate {
			server.AddResourceTemplate(&mcp.ResourceTemplate{
				URITemplate: resourceDef.URI,
				Name:        resourceDef.Name,
				MIMEType:    resourceDef.MimeType,
				Description: resourceDef.Description,
			}, resourceDef.Handler)
		} else {
			server.AddResource(&mcp.Resource{
				URI:         resourceDef.URI,
				Name:        resourceDef.Name,
				MIMEType:    resourceDef.MimeType,
				Description: resourceDef.Description,
			}, resourceDef.Handler)
		}
	}

	return nil
}
