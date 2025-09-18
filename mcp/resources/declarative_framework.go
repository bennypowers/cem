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
	"encoding/json"
	"fmt"
	"strings"

	"bennypowers.dev/cem/mcp/templates"
	"bennypowers.dev/cem/mcp/types"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// DeclarativeResourceConfig holds the configuration for a declarative resource
type DeclarativeResourceConfig struct {
	URI          string
	Name         string
	MimeType     string
	URITemplate  bool
	DataFetchers []types.DataFetcher
	Template     string
	ResponseType string
}

// FetchedData represents the result of a data fetcher (ported from tools)
type FetchedData map[string]any

// BaseTemplateData provides common template data fields (ported from tools)
type BaseTemplateData struct {
	Element           types.ElementInfo
	Context           string
	Options           map[string]string
	SchemaDefinitions any
}

// NewBaseTemplateDataWithSchema creates base template data with schema context
func NewBaseTemplateDataWithSchema(element types.ElementInfo, context string, options map[string]string, schemaDefinitions any) BaseTemplateData {
	return BaseTemplateData{
		Element:           element,
		Context:           context,
		Options:           options,
		SchemaDefinitions: schemaDefinitions,
	}
}

// handleDeclarativeResource is the generic handler for declarative resources
func handleDeclarativeResource(
	ctx context.Context,
	req *mcp.ReadResourceRequest,
	registry types.MCPContext,
	config DeclarativeResourceConfig,
) (*mcp.ReadResourceResult, error) {
	// Step 1: Parse URI parameters and extract arguments
	args, err := parseResourceURI(req.Params.URI, config.URI, config.URITemplate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse resource URI: %w", err)
	}

	// Add resource context
	args["context"] = "resource-access"

	// Step 2: Execute data fetchers to collect required data
	fetchedData, err := executeDataFetchers(config.DataFetchers, registry, args)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}

	// Step 3: Prepare template data by combining args and fetched data
	templateData, err := prepareTemplateData(args, fetchedData, registry)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare template data: %w", err)
	}

	// Step 4: Check if this should return JSON directly
	var response string
	if config.ResponseType == "json" {
		// Return JSON directly from the first data fetcher
		for _, data := range fetchedData {
			jsonBytes, err := json.MarshalIndent(data, "", "  ")
			if err != nil {
				return nil, fmt.Errorf("failed to marshal JSON response: %w", err)
			}
			response = string(jsonBytes)
			break // Use first fetcher result for JSON response
		}
	} else {
		// Step 5: Render the template with the prepared data
		templateName := config.Template
		if templateName == "" {
			templateName = config.Name
		}

		var err error
		response, err = templates.RenderTemplate(templateName, templateData)
		if err != nil {
			return nil, fmt.Errorf("failed to render template '%s': %w", templateName, err)
		}
	}

	// Step 6: Apply any sub-resource filtering if needed
	filteredResponse, err := applySubResourceFiltering(response, args, fetchedData)
	if err != nil {
		return nil, fmt.Errorf("failed to apply sub-resource filtering: %w", err)
	}

	// Step 7: Return the response as a resource
	return &mcp.ReadResourceResult{
		Contents: []*mcp.ResourceContents{{
			URI:      req.Params.URI,
			MIMEType: config.MimeType,
			Text:     filteredResponse,
		}},
	}, nil
}

// executeDataFetchers executes data fetchers using the path traversal system
func executeDataFetchers(fetchers []types.DataFetcher, registry types.MCPContext, args map[string]any) (FetchedData, error) {
	// Create data sources for path traversal
	sourceProvider := NewDataSourceProvider(registry)
	sources, err := sourceProvider.CreateDataSources(args)
	if err != nil {
		return nil, fmt.Errorf("failed to create data sources: %w", err)
	}

	// Create path traversal engine
	engine := NewPathTraversalEngine()

	data := make(FetchedData)

	for _, fetcher := range fetchers {
		// All fetchers must use the path-based system
		if fetcher.Source == "" {
			return nil, fmt.Errorf("fetcher '%s' must have Source configured for path-based execution", fetcher.Name)
		}

		result, err := engine.ExecuteDataFetcher(fetcher, sources)
		if err != nil {
			if fetcher.Required {
				return nil, fmt.Errorf("required data fetcher '%s' failed: %w", fetcher.Name, err)
			}
			// Skip non-required failed fetchers
			continue
		}

		if result != nil {
			data[fetcher.Name] = result
			// Make result available for subsequent fetchers
			sources[fetcher.Name] = result
		}
	}

	return data, nil
}






// prepareTemplateData combines arguments and fetched data for template rendering (ported from tools)
func prepareTemplateData(args map[string]any, fetchedData FetchedData, registry types.MCPContext) (any, error) {
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
		Args        map[string]any
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

// getStringArg safely extracts a string argument with a default value (ported from tools)
func getStringArg(args map[string]any, key, defaultValue string) string {
	if value, ok := args[key]; ok {
		if str, ok := value.(string); ok {
			return str
		}
	}
	return defaultValue
}

// parseResourceURI extracts parameters from the URI based on the template
func parseResourceURI(actualURI, templateURI string, isTemplate bool) (map[string]any, error) {
	args := make(map[string]any)

	if !isTemplate {
		return args, nil
	}

	// Optimize: Remove scheme prefix once for both URIs
	const scheme = "cem://"
	actualPath := strings.TrimPrefix(actualURI, scheme)
	templatePath := strings.TrimPrefix(templateURI, scheme)

	// Split paths into segments
	actualParts := strings.Split(actualPath, "/")
	templateParts := strings.Split(templatePath, "/")

	// Extract parameters from matching template segments
	for i, templatePart := range templateParts {
		if i >= len(actualParts) {
			break
		}

		// Check for parameter syntax: {paramName}
		if len(templatePart) > 2 && templatePart[0] == '{' && templatePart[len(templatePart)-1] == '}' {
			paramName := templatePart[1 : len(templatePart)-1]
			args[paramName] = actualParts[i]
		}
	}

	// Handle sub-resource access beyond template
	if len(actualParts) > len(templateParts) {
		extraParts := actualParts[len(templateParts):]
		args["subResource"] = strings.Join(extraParts, "/")
	}

	return args, nil
}

// applySubResourceFiltering filters the response for sub-resource access using the path traversal system
func applySubResourceFiltering(
	response string,
	args map[string]any,
	fetchedData FetchedData,
) (string, error) {
	// Check if this is a sub-resource request
	subResource, hasSubResource := args["subResource"].(string)
	if !hasSubResource || subResource == "" {
		return response, nil
	}

	// Use the path traversal engine to filter the response based on sub-resource
	engine := NewPathTraversalEngine()

	// Try to filter each data source for an item matching the sub-resource name
	for dataKey, data := range fetchedData {
		// Create a temporary data source map for path resolution
		sources := map[string]any{
			"data": data,
			"args": args,
		}

		// Try to resolve using path expression to find item by name
		// Format: data.#(name==subResourceName) for arrays with name field
		filteredResult, err := engine.ResolvePath("data", fmt.Sprintf("#(name==%s)", subResource), sources)
		if err == nil && filteredResult != nil {
			// Re-render the filtered result as JSON
			return renderFilteredResponse(filteredResult, dataKey)
		}

		// Also try direct path access for nested objects
		// Format: data.subResourceName
		directResult, err := engine.ResolvePath("data", subResource, sources)
		if err == nil && directResult != nil {
			return renderFilteredResponse(directResult, dataKey)
		}
	}

	// If no filtering was applied, return the original response
	return response, nil
}

// renderFilteredResponse renders a filtered single item as JSON
func renderFilteredResponse(filteredItem any, dataKey string) (string, error) {
	// For now, just return the filtered item as JSON
	// In the future, this could use templates for formatted output
	jsonBytes, err := json.MarshalIndent(filteredItem, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal filtered response: %w", err)
	}
	return string(jsonBytes), nil
}

// makeDeclarativeResourceHandler creates a handler function for a declarative resource
func makeDeclarativeResourceHandler(registry types.MCPContext, config DeclarativeResourceConfig) mcp.ResourceHandler {
	return func(ctx context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
		return handleDeclarativeResource(ctx, req, registry, config)
	}
}

// MakeDeclarativeResourceHandler is the exported version for testing
func MakeDeclarativeResourceHandler(registry types.MCPContext, config DeclarativeResourceConfig) mcp.ResourceHandler {
	return makeDeclarativeResourceHandler(registry, config)
}

