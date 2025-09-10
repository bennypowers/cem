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

// handleSuggestAttributes provides intelligent attribute suggestions for custom elements
func handleSuggestAttributes(ctx context.Context, req *mcp.CallToolRequest, registry types.Registry) (*mcp.CallToolResult, error) {
	// Parse args from request
	var suggestArgs SuggestAttributesArgs
	if req.Params.Arguments != nil {
		if argsData, err := json.Marshal(req.Params.Arguments); err != nil {
			return nil, fmt.Errorf("failed to marshal args: %w", err)
		} else if err := json.Unmarshal(argsData, &suggestArgs); err != nil {
			return nil, fmt.Errorf("failed to unmarshal args: %w", err)
		}
	}

	// Get element information
	element, err := registry.ElementInfo(suggestArgs.TagName)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: fmt.Sprintf("Element '%s' not found in registry", suggestArgs.TagName),
				},
			},
		}, nil
	}

	var response strings.Builder
	response.WriteString(fmt.Sprintf("# Attribute Suggestions for `%s`\n\n", suggestArgs.TagName))

	if element.Description() != "" {
		response.WriteString(fmt.Sprintf("**Element Description:** %s\n\n", element.Description()))
	}

	attributes := element.Attributes()
	if len(attributes) == 0 {
		response.WriteString("❌ No attributes defined for this element in the manifest.\n")
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: response.String(),
				},
			},
		}, nil
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
		response.WriteString(fmt.Sprintf("❌ No attributes found matching '%s'\n", suggestArgs.Partial))
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: response.String(),
				},
			},
		}, nil
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
		response.WriteString("## ⚠️ Required Attributes\n\n")
		for _, attr := range requiredAttrs {
			attrData := NewSingleAttributeData(attr, element.TagName(), suggestArgs.Context)
			attrDetails, err := renderAttributeTemplate("attribute_details", attrData)
			if err != nil {
				response.WriteString(fmt.Sprintf("Error rendering attribute details: %v\n", err))
			} else {
				response.WriteString(attrDetails)
				response.WriteString("\n")
			}
		}
	}

	// Show optional attributes
	if len(optionalAttrs) > 0 {
		response.WriteString("## 💡 Optional Attributes\n\n")
		for _, attr := range optionalAttrs {
			attrData := NewSingleAttributeData(attr, element.TagName(), suggestArgs.Context)
			attrDetails, err := renderAttributeTemplate("attribute_details", attrData)
			if err != nil {
				response.WriteString(fmt.Sprintf("Error rendering attribute details: %v\n", err))
			} else {
				response.WriteString(attrDetails)
				response.WriteString("\n")
			}
		}
	}

	// Show accessibility attributes if context is accessibility
	if suggestArgs.Context == "accessibility" && len(accessibilityAttrs) > 0 {
		response.WriteString("## ♿ Accessibility Attributes\n\n")
		for _, attr := range accessibilityAttrs {
			attrData := NewSingleAttributeData(attr, element.TagName(), suggestArgs.Context)
			attrDetails, err := renderAttributeTemplate("attribute_details", attrData)
			if err != nil {
				response.WriteString(fmt.Sprintf("Error rendering attribute details: %v\n", err))
			} else {
				response.WriteString(attrDetails)
				response.WriteString("\n")
			}
		}
	}

	// Add context-specific suggestions
	templateData := NewAttributeTemplateData(element, suggestArgs.Context)
	contextualSuggestions, err := renderTemplate("contextual_suggestions", templateData)
	if err != nil {
		response.WriteString(fmt.Sprintf("Error rendering contextual suggestions: %v\n", err))
	} else {
		response.WriteString(contextualSuggestions)
	}

	// Add usage examples
	attributeExamples, err := renderTemplate("attribute_examples", templateData)
	if err != nil {
		response.WriteString(fmt.Sprintf("Error rendering attribute examples: %v\n", err))
	} else {
		response.WriteString(attributeExamples)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: response.String(),
			},
		},
	}, nil
}
