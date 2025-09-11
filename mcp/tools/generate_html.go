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

// handleGenerateHtml generates HTML structure for custom elements
func handleGenerateHtml(
	ctx context.Context,
	req *mcp.CallToolRequest,
	registry types.Registry,
) (*mcp.CallToolResult, error) {
	// Parse args from request
	var genArgs GenerateHtmlArgs
	if req.Params.Arguments != nil {
		if argsData, err := json.Marshal(req.Params.Arguments); err != nil {
			return nil, fmt.Errorf("failed to marshal args: %w", err)
		} else if err := json.Unmarshal(argsData, &genArgs); err != nil {
			return nil, fmt.Errorf("failed to unmarshal args: %w", err)
		}
	}

	// Get element information
	element, err := registry.ElementInfo(genArgs.TagName)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: fmt.Sprintf("Element '%s' not found in workspace", genArgs.TagName),
				},
			},
		}, nil
	}

	// Generate HTML structure first
	generatedHTML, err := generateHTMLStructure(element, genArgs)
	if err != nil {
		return nil, fmt.Errorf("failed to generate HTML structure: %w", err)
	}

	// Prepare template data
	templateData := prepareHTMLTemplateData(element, genArgs, generatedHTML)

	// Render the complete response using template
	response, err := renderHTMLTemplate("html_generation", templateData)
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
	templateData := HTMLGenerationData{
		ElementInfo: element,
		Content:     args.Content,
		Context:     args.Context,
		Options:     args.Options,
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
	return renderHTMLTemplate("html_structure", templateData)
}

// prepareHTMLTemplateData prepares all data for the main HTML generation template
func prepareHTMLTemplateData(element types.ElementInfo, args GenerateHtmlArgs, generatedHTML string) HTMLGenerationData {
	templateData := HTMLGenerationData{
		ElementInfo:   element,
		GeneratedHTML: generatedHTML,
		Content:       args.Content,
		Context:       args.Context,
		Options:       args.Options,
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
