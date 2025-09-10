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

// SuggestCssIntegrationArgs represents the arguments for the suggest_css_integration tool
type SuggestCssIntegrationArgs struct {
	TagName string            `json:"tagName"`
	Context string            `json:"context,omitempty"`
	Theme   string            `json:"theme,omitempty"`
	Options map[string]string `json:"options,omitempty"`
}

// handleSuggestCssIntegration provides CSS integration guidance for custom elements
func handleSuggestCssIntegration(ctx context.Context, req *mcp.CallToolRequest, registry types.Registry) (*mcp.CallToolResult, error) {
	// Parse args from request
	var cssArgs SuggestCssIntegrationArgs
	if req.Params.Arguments != nil {
		if argsData, err := json.Marshal(req.Params.Arguments); err != nil {
			return nil, fmt.Errorf("failed to marshal args: %w", err)
		} else if err := json.Unmarshal(argsData, &cssArgs); err != nil {
			return nil, fmt.Errorf("failed to unmarshal args: %w", err)
		}
	}

	// Get element information
	element, err := registry.ElementInfo(cssArgs.TagName)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: fmt.Sprintf("Element '%s' not found in registry", cssArgs.TagName),
				},
			},
		}, nil
	}

	var response strings.Builder
	response.WriteString(fmt.Sprintf("# CSS Integration Guide for `%s`\n\n", cssArgs.TagName))

	if element.Description() != "" {
		response.WriteString(fmt.Sprintf("**Element Description:** %s\n\n", element.Description()))
	}

	// Check available CSS features
	hasCustomProperties := len(element.CssProperties()) > 0
	hasCustomParts := len(element.CssParts()) > 0
	hasCustomStates := len(element.CssStates()) > 0

	if !hasCustomProperties && !hasCustomParts && !hasCustomStates {
		response.WriteString("❌ This element doesn't define CSS custom properties, parts, or states in the manifest.\n")
		response.WriteString("You can still style it using regular CSS selectors.\n\n")

		// Use template for basic styling guidance
		templateData := NewTemplateData(element, cssArgs)
		basicStyling, err := renderTemplate("basic_styling", templateData)
		if err != nil {
			return nil, fmt.Errorf("failed to render basic styling template: %w", err)
		}
		response.WriteString(basicStyling)

		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: response.String(),
				},
			},
		}, nil
	}

	// Prepare template data
	templateData := NewTemplateData(element, cssArgs)

	// Generate CSS custom properties section
	if hasCustomProperties {
		propsSection, err := renderTemplate("css_properties", templateData)
		if err != nil {
			return nil, fmt.Errorf("failed to render CSS properties template: %w", err)
		}
		response.WriteString(propsSection)
	}

	// Generate CSS parts section
	if hasCustomParts {
		partsSection, err := renderTemplate("css_parts", templateData)
		if err != nil {
			return nil, fmt.Errorf("failed to render CSS parts template: %w", err)
		}
		response.WriteString(partsSection)
	}

	// Generate CSS states section
	if hasCustomStates {
		statesSection, err := renderTemplate("css_states", templateData)
		if err != nil {
			return nil, fmt.Errorf("failed to render CSS states template: %w", err)
		}
		response.WriteString(statesSection)
	}

	// Generate theming guidance
	themingSection, err := renderTemplate("theming_guidance", templateData)
	if err != nil {
		return nil, fmt.Errorf("failed to render theming guidance template: %w", err)
	}
	response.WriteString(themingSection)

	// Generate responsive guidance if context suggests it
	if strings.Contains(strings.ToLower(cssArgs.Context), "responsive") || cssArgs.Options["responsive"] == "true" {
		responsiveSection, err := renderTemplate("responsive_guidance", templateData)
		if err != nil {
			return nil, fmt.Errorf("failed to render responsive guidance template: %w", err)
		}
		response.WriteString(responsiveSection)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: response.String(),
			},
		},
	}, nil
}
