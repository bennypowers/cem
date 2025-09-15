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
package resources

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

// Using ResourceDefinition and ResourceFrontmatter from types package

// Resources returns all available resource definitions with their handlers
func Resources(registry types.MCPContext) ([]types.ResourceDefinition, error) {
	// Get the directory where this source file is located
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("failed to get current file path")
	}
	resourcesDir := filepath.Dir(filename)

	var resourceDefs []types.ResourceDefinition

	// Find all markdown files in the resources directory
	files, err := filepath.Glob(filepath.Join(resourcesDir, "*.md"))
	if err != nil {
		return nil, fmt.Errorf("failed to find resource markdown files: %w", err)
	}

	for _, file := range files {
		resourceDef, err := parseResourceDefinition(file, registry)
		if err != nil {
			return nil, fmt.Errorf("failed to parse resource definition from %s: %w", file, err)
		}
		resourceDefs = append(resourceDefs, resourceDef)
	}

	return resourceDefs, nil
}

// parseResourceDefinition parses a markdown file with frontmatter into a ResourceDefinition
func parseResourceDefinition(filename string, registry types.MCPContext) (types.ResourceDefinition, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return types.ResourceDefinition{}, fmt.Errorf("failed to read file: %w", err)
	}

	// Split frontmatter and markdown content
	parts := strings.SplitN(string(content), "---", 3)
	if len(parts) < 3 {
		return types.ResourceDefinition{}, fmt.Errorf("invalid markdown format: missing frontmatter")
	}

	// Parse YAML frontmatter
	var frontmatter types.ResourceFrontmatter
	if err := yaml.Unmarshal([]byte(parts[1]), &frontmatter); err != nil {
		return types.ResourceDefinition{}, fmt.Errorf("failed to parse frontmatter: %w", err)
	}

	// Extract description from markdown content
	description := strings.TrimSpace(parts[2])

	// Create resource definition first
	resourceDef := types.ResourceDefinition{
		URI:          frontmatter.URI,
		Name:         frontmatter.Name,
		MimeType:     frontmatter.MimeType,
		URITemplate:  frontmatter.URITemplate,
		Description:  description,
		DataFetchers: frontmatter.DataFetchers,
		Template:     frontmatter.Template,
		ResponseType: frontmatter.ResponseType,
	}

	// Get the corresponding handler
	handler, err := getResourceHandler(resourceDef, registry)
	if err != nil {
		return types.ResourceDefinition{}, fmt.Errorf("failed to get handler for resource %s: %w", frontmatter.Name, err)
	}

	// Set the handler and return the complete resource definition
	resourceDef.Handler = handler
	return resourceDef, nil
}

// getResourceHandler returns the appropriate handler function for a resource
func getResourceHandler(resourceDef types.ResourceDefinition, registry types.MCPContext) (mcp.ResourceHandler, error) {
	// Check if this is a declarative resource (has data fetchers)
	if len(resourceDef.DataFetchers) > 0 {
		config := DeclarativeResourceConfig{
			URI:          resourceDef.URI,
			Name:         resourceDef.Name,
			MimeType:     resourceDef.MimeType,
			URITemplate:  resourceDef.URITemplate,
			DataFetchers: resourceDef.DataFetchers,
			Template:     resourceDef.Template,
			ResponseType: resourceDef.ResponseType,
		}
		return MakeDeclarativeResourceHandler(registry, config), nil
	}

	// Fallback to legacy handlers for non-declarative resources
	switch resourceDef.Name {
	case "schema":
		return makeSchemaHandler(registry), nil
	case "packages":
		return makePackagesHandler(registry), nil
	case "elements":
		return makeElementsHandler(registry), nil
	case "element":
		return makeElementHandler(registry), nil
	case "guidelines":
		return makeGuidelinesHandler(registry), nil
	case "accessibility":
		return makeAccessibilityHandler(registry), nil
	default:
		return nil, fmt.Errorf("unknown resource: %s", resourceDef.Name)
	}
}

// Handler factory functions
func makeSchemaHandler(registry types.MCPContext) mcp.ResourceHandler {
	return func(ctx context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
		return handleSchemaResource(ctx, req, registry)
	}
}

func makePackagesHandler(registry types.MCPContext) mcp.ResourceHandler {
	return func(ctx context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
		return handlePackagesResource(ctx, req, registry)
	}
}

func makeElementsHandler(registry types.MCPContext) mcp.ResourceHandler {
	return func(ctx context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
		return handleElementsResource(ctx, req, registry)
	}
}

func makeElementHandler(registry types.MCPContext) mcp.ResourceHandler {
	return func(ctx context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
		return handleElementResource(ctx, req, registry)
	}
}

func makeGuidelinesHandler(registry types.MCPContext) mcp.ResourceHandler {
	return func(ctx context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
		return handleGuidelinesResource(ctx, req, registry)
	}
}

func makeAccessibilityHandler(registry types.MCPContext) mcp.ResourceHandler {
	return func(ctx context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
		return handleAccessibilityResource(ctx, req, registry)
	}
}
