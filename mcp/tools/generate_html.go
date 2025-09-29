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
	"sort"

	"bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/mcp/helpers"
	"bennypowers.dev/cem/mcp/types"
	V "bennypowers.dev/cem/validate"
	"bennypowers.dev/cem/validations"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type SchemaDefinitionMap map[string]any
type AttributeMap map[string]string

// GenerateHtmlArgs represents the arguments for the generate_html tool
type GenerateHtmlArgs struct {
	TagName    string            `json:"tagName"`
	Context    string            `json:"context,omitempty"`
	Options    map[string]string `json:"options,omitempty"`
	Content    string            `json:"content,omitempty"`
	Attributes AttributeMap      `json:"attributes,omitempty"`
}

// HTMLGenerationTemplateData specific to HTML generation
type HTMLGenerationTemplateData struct {
	BaseTemplateData
	Content            string
	GeneratedHTML      string
	RequiredAttributes []AttributeWithValue
	OptionalAttributes []AttributeWithValue
	SortedAttributes   []AttributeWithValue
	Slots              []SlotWithContent
}

// NewHTMLGenerationTemplateData creates HTML generation template data
func NewHTMLGenerationTemplateData(element types.ElementInfo, context string, options map[string]string) HTMLGenerationTemplateData {
	return HTMLGenerationTemplateData{
		BaseTemplateData: NewBaseTemplateData(element, context, options),
	}
}

// NewHTMLGenerationTemplateDataWithSchema creates HTML generation template data with schema context
func NewHTMLGenerationTemplateDataWithSchema(element types.ElementInfo, context string, options map[string]string, schemaDefinitions any) HTMLGenerationTemplateData {
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
	templateData := prepareHTMLTemplateDataWithSchema(
		element,
		genArgs,
		generatedHTML,
		schemaDefinitions,
	)

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

	// All attributes are optional in custom elements manifests
	// (no Required field exists in manifest.Attribute)

	// Add user-provided attributes
	for name, value := range args.Attributes {
		// Find the attribute definition if it exists in the element manifest
		found := false
		for _, attr := range element.Attributes() {
			if attr.Name == name {
				templateData.OptionalAttributes = append(templateData.OptionalAttributes, AttributeWithValue{
					Attribute: attr,
					Value:     value,
				})
				found = true
				break
			}
		}

		// If not found in manifest but is a valid global HTML attribute, include it
		if !found && validations.IsGlobalAttribute(name) {
			attr := types.Attribute{
				FullyQualified: manifest.FullyQualified{Name: name},
			}
			templateData.OptionalAttributes = append(templateData.OptionalAttributes, AttributeWithValue{
				Attribute: attr,
				Value:     value,
			})
		}
	}

	// Sort attributes according to HTML conventions
	templateData.SortedAttributes = sortAttributesForHTML(templateData.OptionalAttributes, element.Attributes())

	// Prepare slot data with example content
	for _, slot := range element.Slots() {
		slotName := slot.Name
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

// prepareHTMLTemplateDataWithSchema prepares all data for the main HTML generation template with schema context
func prepareHTMLTemplateDataWithSchema(
	element types.ElementInfo,
	args GenerateHtmlArgs,
	generatedHTML string,
	schemaDefinitions SchemaDefinitionMap,
) HTMLGenerationTemplateData {
	templateData := HTMLGenerationTemplateData{
		BaseTemplateData: NewBaseTemplateDataWithSchema(element, args.Context, args.Options, schemaDefinitions),
		Content:          args.Content,
		GeneratedHTML:    generatedHTML,
	}

	// Add attribute data with values for the main template
	// All attributes are optional in custom elements manifests

	return templateData
}

// getAttributeValue determines the value to use for an attribute

// sortAttributesForHTML sorts attributes according to HTML conventions:
// 1. id (if exists)
// 2. class (if exists)
// 3. native HTML attributes (alphabetically)
// 4. custom element attributes (in manifest order)
func sortAttributesForHTML(attributes []AttributeWithValue, manifestAttrs []types.Attribute) []AttributeWithValue {
	if len(attributes) == 0 {
		return attributes
	}

	// Create map for manifest order lookup
	manifestOrder := make(map[string]int)
	for i, attr := range manifestAttrs {
		manifestOrder[attr.Name] = i
	}

	// Separate attributes into categories
	var idAttr, classAttr *AttributeWithValue
	var nativeAttrs, customAttrs []AttributeWithValue

	for i := range attributes {
		// Store the attribute value once to avoid repeated dereferencing
		attrValue := attributes[i]
		switch attrValue.Name {
		case "id":
			idAttr = &attributes[i]
		case "class":
			classAttr = &attributes[i]
		default:
			// Check if this attribute is defined in the manifest first
			// If so, it's a custom element attribute regardless of name collision
			isManifestAttr := false
			for _, manifestAttr := range manifestAttrs {
				if manifestAttr.Name == attrValue.Name {
					isManifestAttr = true
					break
				}
			}

			if isManifestAttr {
				customAttrs = append(customAttrs, attrValue)
			} else if validations.IsGlobalAttribute(attrValue.Name) {
				nativeAttrs = append(nativeAttrs, attrValue)
			} else {
				customAttrs = append(customAttrs, attrValue)
			}
		}
	}

	// Sort native attributes alphabetically
	sort.Slice(nativeAttrs, func(i, j int) bool {
		return nativeAttrs[i].Name < nativeAttrs[j].Name
	})

	// Sort custom attributes by manifest order
	sort.Slice(customAttrs, func(i, j int) bool {
		orderI, hasI := manifestOrder[customAttrs[i].Name]
		orderJ, hasJ := manifestOrder[customAttrs[j].Name]

		// If both have manifest order, use that
		if hasI && hasJ {
			return orderI < orderJ
		}
		// If only one has manifest order, it comes first
		if hasI {
			return true
		}
		if hasJ {
			return false
		}
		// If neither has manifest order, sort alphabetically
		return customAttrs[i].Name < customAttrs[j].Name
	})

	// Combine in order: id, class, native, custom
	var result []AttributeWithValue
	if idAttr != nil {
		result = append(result, *idAttr)
	}
	if classAttr != nil {
		result = append(result, *classAttr)
	}
	result = append(result, nativeAttrs...)
	result = append(result, customAttrs...)

	return result
}

// getSchemaDefinitions retrieves schema definitions for template context
func getSchemaDefinitions(registry types.MCPContext) (SchemaDefinitionMap, error) {
	// Get schema versions from manifests
	versions := registry.GetManifestSchemaVersions()
	if len(versions) == 0 {
		return make(map[string]any), nil
	}

	// Use the first version to load schema
	schemaVersion := versions[0]
	schemaData, err := V.GetSchema(schemaVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to load schema %s: %w", schemaVersion, err)
	}
	// Parse schema JSON - return the complete schema for template functions
	schema := make(map[string]any)
	if err := json.Unmarshal(schemaData, &schema); err != nil {
		return nil, fmt.Errorf("failed to parse schema JSON: %w", err)
	}
	// Return complete schema so template functions can navigate properly
	return schema, nil
}
