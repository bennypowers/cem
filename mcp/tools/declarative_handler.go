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
	"encoding/json"
	"fmt"

	"bennypowers.dev/cem/mcp/types"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// DeclarativeToolConfig holds the configuration for a declarative tool
type DeclarativeToolConfig struct {
	Name         string
	DataFetchers []types.DataFetcher
	Template     string
	ResponseType string
}

// handleDeclarativeTool is the generic handler for declarative tools
func handleDeclarativeTool(
	ctx context.Context,
	req *mcp.CallToolRequest,
	registry types.MCPContext,
	config DeclarativeToolConfig,
) (*mcp.CallToolResult, error) {
	// Step 1: Parse arguments from the request
	args, err := parseGenericArgs(req)
	if err != nil {
		return nil, fmt.Errorf("failed to parse tool arguments: %w", err)
	}

	// Step 2: Execute data fetchers to collect required data
	fetchedData, err := ExecuteDataFetchers(config.DataFetchers, registry, args)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}

	// Step 3: Prepare template data by combining args and fetched data
	templateData, err := prepareTemplateData(args, fetchedData, registry)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare template data: %w", err)
	}

	// Step 4: Determine template name (use explicit or default to tool name)
	templateName := config.Template
	if templateName == "" {
		templateName = config.Name
	}

	// Step 5: Render the template with the prepared data
	response, err := RenderTemplate(templateName, templateData)
	if err != nil {
		return nil, fmt.Errorf("failed to render template '%s': %w", templateName, err)
	}

	// Step 6: Return the response in the appropriate format
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: response,
			},
		},
	}, nil
}

// parseGenericArgs parses arguments from MCP request into a generic map
func parseGenericArgs(req *mcp.CallToolRequest) (map[string]interface{}, error) {
	args := make(map[string]interface{})

	if req.Params.Arguments != nil {
		// Convert the arguments to JSON and back to get a clean map
		argsData, err := json.Marshal(req.Params.Arguments)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal arguments: %w", err)
		}

		if err := json.Unmarshal(argsData, &args); err != nil {
			return nil, fmt.Errorf("failed to unmarshal arguments: %w", err)
		}
	}

	return args, nil
}

// prepareTemplateData combines arguments and fetched data for template rendering
func prepareTemplateData(args map[string]interface{}, fetchedData FetchedData, registry types.MCPContext) (interface{}, error) {
	// Create the base template data structure
	templateData := BaseTemplateData{
		Context: getStringArg(args, "context", ""),
		Options: make(map[string]string),
	}

	// Add element data if available
	if element, ok := fetchedData["element"]; ok {
		if elementInfo, ok := element.(types.ElementInfo); ok {
			templateData.Element = elementInfo
		}
	}

	// Add schema definitions if available
	if schema, ok := fetchedData["schema"]; ok {
		templateData.SchemaDefinitions = schema
	}

	// Create a combined data structure that includes both base template data
	// and individual fetched data for template access
	combinedData := struct {
		BaseTemplateData
		FetchedData FetchedData
		Args        map[string]interface{}
	}{
		BaseTemplateData: templateData,
		FetchedData:      fetchedData,
		Args:             args,
	}

	// If we have element data, use the existing pattern for compatibility
	if templateData.Element != nil {
		return NewBaseTemplateDataWithSchema(
			templateData.Element,
			templateData.Context,
			templateData.Options,
			templateData.SchemaDefinitions,
		), nil
	}

	return combinedData, nil
}

// getStringArg safely extracts a string argument with a default value
func getStringArg(args map[string]interface{}, key, defaultValue string) string {
	if value, ok := args[key]; ok {
		if str, ok := value.(string); ok {
			return str
		}
	}
	return defaultValue
}

// makeDeclarativeHandler creates a handler function for a declarative tool
func makeDeclarativeHandler(registry types.MCPContext, config DeclarativeToolConfig) mcp.ToolHandler {
	return func(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return handleDeclarativeTool(ctx, req, registry, config)
	}
}

// MakeDeclarativeHandler is the exported version for testing
func MakeDeclarativeHandler(registry types.MCPContext, config DeclarativeToolConfig) mcp.ToolHandler {
	return makeDeclarativeHandler(registry, config)
}