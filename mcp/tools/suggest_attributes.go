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
	"fmt"
	"strings"

	"bennypowers.dev/cem/mcp/types"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// SuggestAttributesArgs represents the arguments for the suggest_attributes tool
type SuggestAttributesArgs struct {
	TagName string `json:"tagName"`
	Context string `json:"context,omitempty"`
	Partial string `json:"partial,omitempty"`
}

// AttributeSuggestionTemplateData specific to attribute suggestions
type AttributeSuggestionTemplateData struct {
	BaseTemplateData
	Attribute types.Attribute
	TagName   string
}

// NewAttributeSuggestionTemplateData creates attribute suggestion template data
func NewAttributeSuggestionTemplateData(element types.ElementInfo, attr types.Attribute, context string) AttributeSuggestionTemplateData {
	return AttributeSuggestionTemplateData{
		BaseTemplateData: NewBaseTemplateData(element, context, nil),
		Attribute:        attr,
		TagName:          element.TagName(),
	}
}

// handleSuggestAttributes provides intelligent attribute suggestions for custom elements
func handleSuggestAttributes(ctx context.Context, req *mcp.CallToolRequest, registry types.MCPContext) (*mcp.CallToolResult, error) {
	// Parse args from request
	suggestArgs, err := ParseToolArgs[SuggestAttributesArgs](req)
	if err != nil {
		return nil, err
	}

	// Get element information
	element, errorResponse, err := LookupElement(registry, suggestArgs.TagName)
	if err != nil {
		return nil, err
	}
	if errorResponse != nil {
		return errorResponse, nil
	}

	response := NewResponseBuilder()
	response.AddHeader(1, fmt.Sprintf("Attribute Suggestions for `%s`", suggestArgs.TagName))
	response.AddDescription(element)

	attributes := element.Attributes()
	if len(attributes) == 0 {
		response.AddSection("❌ No attributes defined for this element in the manifest.")
		return response.Build(), nil
	}

	// Filter attributes based on partial match if provided
	var filteredAttributes []types.Attribute
	if suggestArgs.Partial != "" {
		partialLower := strings.ToLower(suggestArgs.Partial)
		for _, attr := range attributes {
			if strings.Contains(strings.ToLower(attr.Name()), partialLower) {
				filteredAttributes = append(filteredAttributes, attr)
			}
		}
	} else {
		filteredAttributes = attributes
	}

	if len(filteredAttributes) == 0 {
		response.AddSection(fmt.Sprintf("❌ No attributes found matching '%s'", suggestArgs.Partial))
		return response.Build(), nil
	}

	// Categorize attributes
	requiredAttrs := []types.Attribute{}
	optionalAttrs := []types.Attribute{}
	accessibilityAttrs := []types.Attribute{}

	for _, attr := range filteredAttributes {
		if attr.Required() {
			requiredAttrs = append(requiredAttrs, attr)
		} else {
			optionalAttrs = append(optionalAttrs, attr)
		}

		// Check if it's accessibility-related
		attrName := strings.ToLower(attr.Name())
		if strings.HasPrefix(attrName, "aria-") || attrName == "role" ||
			strings.Contains(attrName, "label") || strings.Contains(attrName, "accessible") {
			accessibilityAttrs = append(accessibilityAttrs, attr)
		}
	}

	// Show required attributes first
	if len(requiredAttrs) > 0 {
		response.AddHeader(2, "⚠️ Required Attributes")
		for _, attr := range requiredAttrs {
			attrData := NewAttributeSuggestionTemplateData(element, attr, suggestArgs.Context)
			attrDetails, err := RenderTemplate("attribute_details", attrData)
			if err != nil {
				response.AddSection(fmt.Sprintf("Error rendering attribute details: %v", err))
			} else {
				response.AddSection(attrDetails)
			}
		}
	}

	// Show optional attributes
	if len(optionalAttrs) > 0 {
		response.AddHeader(2, "💡 Optional Attributes")
		for _, attr := range optionalAttrs {
			attrData := NewAttributeSuggestionTemplateData(element, attr, suggestArgs.Context)
			attrDetails, err := RenderTemplate("attribute_details", attrData)
			if err != nil {
				response.AddSection(fmt.Sprintf("Error rendering attribute details: %v", err))
			} else {
				response.AddSection(attrDetails)
			}
		}
	}

	// Show accessibility attributes if context is accessibility
	if suggestArgs.Context == "accessibility" && len(accessibilityAttrs) > 0 {
		response.AddHeader(2, "♿ Accessibility Attributes")
		for _, attr := range accessibilityAttrs {
			attrData := NewAttributeSuggestionTemplateData(element, attr, suggestArgs.Context)
			attrDetails, err := RenderTemplate("attribute_details", attrData)
			if err != nil {
				response.AddSection(fmt.Sprintf("Error rendering attribute details: %v", err))
			} else {
				response.AddSection(attrDetails)
			}
		}
	}

	// Add context-specific suggestions
	baseData := NewBaseTemplateData(element, suggestArgs.Context, nil)
	response.AddTemplateSection("contextual_suggestions", baseData)

	// Add usage examples
	response.AddTemplateSection("attribute_examples", baseData)

	return response.Build(), nil
}
