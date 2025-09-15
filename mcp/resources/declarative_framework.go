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

	"bennypowers.dev/cem/mcp/tools" // Only for template rendering and data fetchers
	"bennypowers.dev/cem/mcp/types"
	V "bennypowers.dev/cem/validate"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/tidwall/gjson"
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
type FetchedData map[string]interface{}

// BaseTemplateData provides common template data fields (ported from tools)
type BaseTemplateData struct {
	Element           types.ElementInfo
	Context           string
	Options           map[string]string
	SchemaDefinitions interface{}
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

	// Step 4: Determine template name (use explicit or default to resource name)
	templateName := config.Template
	if templateName == "" {
		templateName = config.Name
	}

	// Step 5: Render the template with the prepared data
	response, err := tools.RenderTemplate(templateName, templateData)
	if err != nil {
		return nil, fmt.Errorf("failed to render template '%s': %w", templateName, err)
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

// executeDataFetchers executes multiple data fetchers and returns combined results (ported from tools)
func executeDataFetchers(fetchers []types.DataFetcher, registry types.MCPContext, args map[string]interface{}) (FetchedData, error) {
	data := make(FetchedData)

	for _, fetcher := range fetchers {
		result, err := executeSingleFetcher(fetcher, registry, args)
		if err != nil {
			if fetcher.Required {
				return nil, fmt.Errorf("required data fetcher '%s' failed: %w", fetcher.Name, err)
			}
			// Skip non-required failed fetchers
			continue
		}
		data[fetcher.Name] = result
	}

	return data, nil
}

// executeSingleFetcher executes a single data fetcher (ported from tools)
func executeSingleFetcher(fetcher types.DataFetcher, registry types.MCPContext, args map[string]interface{}) (interface{}, error) {
	switch fetcher.Type {
	case "manifest_element":
		return fetchManifestElement(registry, args)
	case "schema_definitions":
		return fetchSchemaDefinitions(registry, fetcher.Path)
	case "attribute_collection":
		return fetchAttributeCollection(registry, args)
	case "element_collection":
		return fetchElementCollection(registry, fetcher.Path)
	default:
		return nil, fmt.Errorf("unknown data fetcher type: %s", fetcher.Type)
	}
}

// fetchManifestElement fetches element information from the registry (ported from tools)
func fetchManifestElement(registry types.MCPContext, args map[string]interface{}) (interface{}, error) {
	tagName, ok := args["tagName"].(string)
	if !ok || tagName == "" {
		return nil, fmt.Errorf("tagName is required for manifest_element fetcher")
	}

	return registry.ElementInfo(tagName)
}

// fetchSchemaDefinitions fetches schema definitions with gjson path support (ported from tools)
func fetchSchemaDefinitions(registry types.MCPContext, path string) (interface{}, error) {
	// Get schema versions from manifests
	versions := registry.GetManifestSchemaVersions()
	if len(versions) == 0 {
		return make(map[string]interface{}), nil
	}

	// Use the first version to load schema
	schemaVersion := versions[0]
	schemaData, err := getSchemaData(schemaVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to load schema %s: %w", schemaVersion, err)
	}

	// If no path specified, return the whole schema
	if path == "" {
		var schema map[string]interface{}
		if err := json.Unmarshal(schemaData, &schema); err != nil {
			return nil, fmt.Errorf("failed to parse schema JSON: %w", err)
		}
		return schema, nil
	}

	// Use gjson to extract the specific path
	result := gjson.GetBytes(schemaData, path)
	if !result.Exists() {
		return nil, fmt.Errorf("path '%s' not found in schema", path)
	}

	var extracted interface{}
	if err := json.Unmarshal([]byte(result.Raw), &extracted); err != nil {
		return nil, fmt.Errorf("failed to parse extracted schema data: %w", err)
	}

	return extracted, nil
}

// fetchAttributeCollection fetches attribute information (ported from tools)
func fetchAttributeCollection(registry types.MCPContext, args map[string]interface{}) (interface{}, error) {
	tagName, ok := args["tagName"].(string)
	if !ok || tagName == "" {
		return nil, fmt.Errorf("tagName is required for attribute_collection fetcher")
	}

	element, err := registry.ElementInfo(tagName)
	if err != nil {
		return nil, fmt.Errorf("failed to get element info: %w", err)
	}

	// Return the attributes directly
	return element.Attributes(), nil
}

// fetchElementCollection fetches multiple elements using gjson (ported from tools)
func fetchElementCollection(registry types.MCPContext, path string) (interface{}, error) {
	// Get all elements from registry
	allElements := registry.AllElements()

	// Convert to JSON for gjson processing
	elementsJSON, err := json.Marshal(allElements)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal elements: %w", err)
	}

	// If no path specified, return all elements
	if path == "" {
		return allElements, nil
	}

	// Use gjson to extract the specific path
	result := gjson.GetBytes(elementsJSON, path)
	if !result.Exists() {
		return nil, fmt.Errorf("path '%s' not found in elements", path)
	}

	var extracted interface{}
	if err := json.Unmarshal([]byte(result.Raw), &extracted); err != nil {
		return nil, fmt.Errorf("failed to parse extracted elements data: %w", err)
	}

	return extracted, nil
}

// getSchemaData retrieves schema data using the existing schema loading mechanism (ported from tools)
func getSchemaData(version string) ([]byte, error) {
	// Use the existing V.GetSchema function from the validate package
	return V.GetSchema(version)
}

// prepareTemplateData combines arguments and fetched data for template rendering (ported from tools)
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
		return tools.NewBaseTemplateDataWithSchema(
			templateData.Element,
			templateData.Context,
			templateData.Options,
			templateData.SchemaDefinitions,
		), nil
	}

	return combinedData, nil
}

// getStringArg safely extracts a string argument with a default value (ported from tools)
func getStringArg(args map[string]interface{}, key, defaultValue string) string {
	if value, ok := args[key]; ok {
		if str, ok := value.(string); ok {
			return str
		}
	}
	return defaultValue
}

// parseResourceURI extracts parameters from the URI based on the template
func parseResourceURI(actualURI, templateURI string, isTemplate bool) (map[string]interface{}, error) {
	args := make(map[string]interface{})

	if !isTemplate {
		// Simple URI, no parameters to extract
		return args, nil
	}

	// Extract parameters from URI template
	// Expected patterns:
	// cem://element/{tagName}/attributes -> extracts tagName
	// cem://element/{tagName}/attributes/{name} -> extracts tagName and name
	// cem://element/{tagName}/css/parts/{name} -> extracts tagName and name

	// Remove scheme and split into parts
	actualParts := strings.Split(strings.TrimPrefix(actualURI, "cem://"), "/")
	templateParts := strings.Split(strings.TrimPrefix(templateURI, "cem://"), "/")

	// Match template parts with actual parts
	for i, templatePart := range templateParts {
		if i >= len(actualParts) {
			continue
		}

		actualPart := actualParts[i]

		// Check if this is a parameter (surrounded by {})
		if strings.HasPrefix(templatePart, "{") && strings.HasSuffix(templatePart, "}") {
			paramName := strings.Trim(templatePart, "{}")
			args[paramName] = actualPart
		}
	}

	// Handle sub-resource access (optional parameters beyond template)
	if len(actualParts) > len(templateParts) {
		// Additional path segments indicate sub-resource access
		extraParts := actualParts[len(templateParts):]
		if len(extraParts) > 0 {
			// Store the sub-resource name for filtering
			args["subResource"] = strings.Join(extraParts, "/")
		}
	}

	return args, nil
}

// applySubResourceFiltering filters the response for sub-resource access
func applySubResourceFiltering(response string, args map[string]interface{}, fetchedData FetchedData) (string, error) {
	// Check if this is a sub-resource request
	subResource, hasSubResource := args["subResource"].(string)
	if !hasSubResource || subResource == "" {
		return response, nil
	}

	// Apply filtering based on the sub-resource type and name
	// This would be expanded based on specific filtering needs
	// For now, return the full response (filtering can be implemented per resource type)
	return response, nil
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