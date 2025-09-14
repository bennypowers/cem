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

	"bennypowers.dev/cem/mcp/helpers"
	"bennypowers.dev/cem/mcp/types"
	V "bennypowers.dev/cem/validate"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// GenerateHtmlArgs represents the arguments for the generate_html tool
type GenerateHtmlArgs struct {
	TagName    string            `json:"tagName"`
	Context    string            `json:"context,omitempty"`
	Options    map[string]string `json:"options,omitempty"`
	Content    string            `json:"content,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

// HTMLGenerationTemplateData specific to HTML generation
type HTMLGenerationTemplateData struct {
	BaseTemplateData
	Content            string
	GeneratedHTML      string
	RequiredAttributes []AttributeWithValue
	OptionalAttributes []AttributeWithValue
	Slots              []SlotWithContent
}

// NewHTMLGenerationTemplateData creates HTML generation template data
func NewHTMLGenerationTemplateData(element types.ElementInfo, context string, options map[string]string) HTMLGenerationTemplateData {
	return HTMLGenerationTemplateData{
		BaseTemplateData: NewBaseTemplateData(element, context, options),
	}
}

// NewHTMLGenerationTemplateDataWithSchema creates HTML generation template data with schema context
func NewHTMLGenerationTemplateDataWithSchema(element types.ElementInfo, context string, options map[string]string, schemaDefinitions interface{}) HTMLGenerationTemplateData {
	return HTMLGenerationTemplateData{
		BaseTemplateData: NewBaseTemplateDataWithSchema(element, context, options, schemaDefinitions),
	}
}

// handleGenerateHtml generates HTML structure for custom elements
func handleGenerateHtml(
	ctx context.Context,
	req *mcp.CallToolRequest,
	registry types.MCPContext,
) (*mcp.CallToolResult, error) {
	// Parse args from request
	genArgs, err := ParseToolArgs[GenerateHtmlArgs](req)
	if err != nil {
		return nil, err
	}

	// Get element information
	element, errorResponse, err := LookupElement(registry, genArgs.TagName)
	if err != nil {
		return nil, err
	}
	if errorResponse != nil {
		return errorResponse, nil
	}

	// Generate HTML structure first
	generatedHTML, err := generateHTMLStructure(element, genArgs)
	if err != nil {
		return nil, fmt.Errorf("failed to generate HTML structure: %w", err)
	}

	// Get schema definitions for rich context
	schemaDefinitions, err := getSchemaDefinitions(registry)
	if err != nil {
		return nil, fmt.Errorf("failed to get schema definitions: %w", err)
	}

	// Prepare template data with schema context
	templateData := prepareHTMLTemplateDataWithSchema(element, genArgs, generatedHTML, schemaDefinitions)

	// Render the complete response using template
	response, err := RenderTemplate("html_generation", templateData)
	if err != nil {
		return nil, fmt.Errorf("failed to render HTML generation template: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: response,
			},
		},
	}, nil
}

// generateHTMLStructure creates the actual HTML using template
func generateHTMLStructure(element types.ElementInfo, args GenerateHtmlArgs) (string, error) {
	// Prepare data for HTML structure template
	templateData := HTMLGenerationTemplateData{
		BaseTemplateData: NewBaseTemplateData(element, args.Context, args.Options),
		Content:          args.Content,
	}

	// Separate required and optional attributes
	for _, attr := range element.Attributes() {
		if attr.Required() {
			value := getAttributeValue(attr, args.Attributes)
			templateData.RequiredAttributes = append(templateData.RequiredAttributes, AttributeWithValue{
				Attribute: attr,
				Value:     value,
			})
		}
	}

	// Add user-provided optional attributes
	for name, value := range args.Attributes {
		isRequired := false
		for _, attr := range element.Attributes() {
			if attr.Required() && attr.Name() == name {
				isRequired = true
				break
			}
		}
		if !isRequired {
			// Find the attribute definition if it exists
			for _, attr := range element.Attributes() {
				if attr.Name() == name {
					templateData.OptionalAttributes = append(templateData.OptionalAttributes, AttributeWithValue{
						Attribute: attr,
						Value:     value,
					})
					break
				}
			}
		}
	}

	// Prepare slot data with example content
	for _, slot := range element.Slots() {
		slotName := slot.Name()
		var exampleContent, defaultContent string

		if slotName == "" {
			defaultContent = args.Content
			if defaultContent == "" {
				defaultContent = "Default content"
			}
		} else {
			exampleContent = helpers.TitleCaser.String(slotName) + " content"
		}

		templateData.Slots = append(templateData.Slots, SlotWithContent{
			Slot:           slot,
			ExampleContent: exampleContent,
			DefaultContent: defaultContent,
		})
	}

	// Use template to generate HTML structure
	return RenderTemplate("html_structure", templateData)
}

// prepareHTMLTemplateData prepares all data for the main HTML generation template
func prepareHTMLTemplateData(element types.ElementInfo, args GenerateHtmlArgs, generatedHTML string) HTMLGenerationTemplateData {
	templateData := HTMLGenerationTemplateData{
		BaseTemplateData: NewBaseTemplateData(element, args.Context, args.Options),
		Content:          args.Content,
		GeneratedHTML:    generatedHTML,
	}

	// Add attribute data with values for the main template
	for _, attr := range element.Attributes() {
		if attr.Required() {
			value := getAttributeValue(attr, args.Attributes)
			templateData.RequiredAttributes = append(templateData.RequiredAttributes, AttributeWithValue{
				Attribute: attr,
				Value:     value,
			})
		}
	}

	return templateData
}

// prepareHTMLTemplateDataWithSchema prepares all data for the main HTML generation template with schema context
func prepareHTMLTemplateDataWithSchema(element types.ElementInfo, args GenerateHtmlArgs, generatedHTML string, schemaDefinitions interface{}) HTMLGenerationTemplateData {
	templateData := HTMLGenerationTemplateData{
		BaseTemplateData: NewBaseTemplateDataWithSchema(element, args.Context, args.Options, schemaDefinitions),
		Content:          args.Content,
		GeneratedHTML:    generatedHTML,
	}

	// Add attribute data with values for the main template
	for _, attr := range element.Attributes() {
		if attr.Required() {
			value := getAttributeValue(attr, args.Attributes)
			templateData.RequiredAttributes = append(templateData.RequiredAttributes, AttributeWithValue{
				Attribute: attr,
				Value:     value,
			})
		}
	}

	return templateData
}

// getAttributeValue determines the value to use for an attribute
func getAttributeValue(attr types.Attribute, providedAttrs map[string]string) string {
	// Check if user provided a value
	if value, exists := providedAttrs[attr.Name()]; exists {
		return value
	}

	// Use default value if available
	if attr.Default() != "" {
		return attr.Default()
	}

	// Use first valid value if available
	if len(attr.Values()) > 0 {
		return attr.Values()[0]
	}

	// Fallback to placeholder
	return "value"
}

// getSchemaDefinitions retrieves schema definitions for template context
func getSchemaDefinitions(registry types.MCPContext) (map[string]interface{}, error) {
	// Get schema versions from manifests
	versions := registry.GetManifestSchemaVersions()
	if len(versions) == 0 {
		return make(map[string]interface{}), nil
	}

	// Use the first version to load schema
	schemaVersion := versions[0]
	schemaData, err := V.GetSchema(schemaVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to load schema %s: %w", schemaVersion, err)
	}
	// Parse schema JSON - return the complete schema for template functions
	var schema map[string]interface{}
	if err := json.Unmarshal(schemaData, &schema); err != nil {
		return nil, fmt.Errorf("failed to parse schema JSON: %w", err)
	}
	// Return complete schema so template functions can navigate properly
	return schema, nil
}
