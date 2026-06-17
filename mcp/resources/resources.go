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
	"embed"
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"

	IC "bennypowers.dev/cem/internal/config"
	"bennypowers.dev/cem/mcp/types"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

//go:embed *.md
var resourceDefinitions embed.FS

// Using ResourceDefinition and ResourceFrontmatter from types package

// Resources returns all available resource definitions with their handlers
func Resources(registry types.MCPContext) ([]types.ResourceDefinition, error) {
	var resourceDefs []types.ResourceDefinition

	// Read all markdown files from embedded filesystem
	entries, err := resourceDefinitions.ReadDir(".")
	if err != nil {
		return nil, fmt.Errorf("failed to read embedded resource definitions: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		resourceDef, err := parseResourceDefinition(entry.Name(), registry)
		if err != nil {
			return nil, fmt.Errorf("failed to parse resource definition from %s: %w", entry.Name(), err)
		}
		resourceDefs = append(resourceDefs, resourceDef)
	}

	return resourceDefs, nil
}

// parseResourceDefinition parses a markdown file with frontmatter into a ResourceDefinition
func parseResourceDefinition(filename string, registry types.MCPContext) (types.ResourceDefinition, error) {
	content, err := resourceDefinitions.ReadFile(filename)
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
	switch resourceDef.Name {
	case "config-schema":
		return makeConfigSchemaOverviewHandler(registry), nil
	case "config-schema-section":
		return makeConfigSchemaSectionHandler(registry), nil
	}

	if len(resourceDef.DataFetchers) == 0 {
		return nil, fmt.Errorf("resource %s is missing data fetchers configuration", resourceDef.Name)
	}

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

func makeConfigSchemaOverviewHandler(_ types.MCPContext) mcp.ResourceHandler {
	var b strings.Builder
	b.WriteString("# CEM Configuration Schema\n\n")
	b.WriteString("## Fields\n\n")
	for _, s := range IC.SchemaSections() {
		fmt.Fprintf(&b, "- **%s** -- %s\n  `cem://config/schema/%s`\n", s.Key, s.Description, s.Key)
	}

	text := b.String()

	return func(_ context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
		return &mcp.ReadResourceResult{
			Contents: []*mcp.ResourceContents{{
				URI:      req.Params.URI,
				MIMEType: "text/markdown",
				Text:     text,
			}},
		}, nil
	}
}

func makeConfigSchemaSectionHandler(_ types.MCPContext) mcp.ResourceHandler {
	schemaBytes := IC.SchemaJSON()
	var schema map[string]any
	if err := json.Unmarshal(schemaBytes, &schema); err != nil {
		return func(_ context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
			return nil, fmt.Errorf("failed to parse config schema: %w", err)
		}
	}

	props, _ := schema["properties"].(map[string]any)

	return func(_ context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
		uri := req.Params.URI
		section := strings.TrimPrefix(uri, "cem://config/schema/")
		if !IC.IsSchemaSection(section) {
			return nil, fmt.Errorf("unknown config section %q", section)
		}

		sectionSchema, ok := props[section]
		if !ok {
			return nil, fmt.Errorf("section %q not found in config schema", section)
		}

		data, err := json.MarshalIndent(sectionSchema, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("failed to marshal section schema: %w", err)
		}

		return &mcp.ReadResourceResult{
			Contents: []*mcp.ResourceContents{{
				URI:      uri,
				MIMEType: "application/json",
				Text:     string(data),
			}},
		}, nil
	}
}
