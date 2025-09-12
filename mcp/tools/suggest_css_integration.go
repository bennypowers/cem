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

// SuggestCssIntegrationArgs represents the arguments for the suggest_css_integration tool
type SuggestCssIntegrationArgs struct {
	TagName string            `json:"tagName"`
	Context string            `json:"context,omitempty"`
	Theme   string            `json:"theme,omitempty"`
	Options map[string]string `json:"options,omitempty"`
}

// CSSTemplateData specific to CSS integration tools
type CSSTemplateData struct {
	BaseTemplateData
	Theme string
}

// Template accessor methods for direct access in templates
func (c CSSTemplateData) TagName() string                    { return c.Element.TagName() }
func (c CSSTemplateData) CssProperties() []types.CssProperty { return c.Element.CssProperties() }
func (c CSSTemplateData) CssParts() []types.CssPart          { return c.Element.CssParts() }
func (c CSSTemplateData) CssStates() []types.CssState        { return c.Element.CssStates() }

// NewCSSTemplateData creates CSS-specific template data
func NewCSSTemplateData(element types.ElementInfo, context string, options map[string]string, theme string) CSSTemplateData {
	return CSSTemplateData{
		BaseTemplateData: NewBaseTemplateData(element, context, options),
		Theme:            theme,
	}
}

// handleSuggestCssIntegration provides CSS integration guidance for custom elements
func handleSuggestCssIntegration(ctx context.Context, req *mcp.CallToolRequest, registry types.MCPContext) (*mcp.CallToolResult, error) {
	// Parse args from request
	cssArgs, err := ParseToolArgs[SuggestCssIntegrationArgs](req)
	if err != nil {
		return nil, err
	}

	// Get element information
	element, errorResponse, err := LookupElement(registry, cssArgs.TagName)
	if err != nil {
		return nil, err
	}
	if errorResponse != nil {
		return errorResponse, nil
	}

	response := NewResponseBuilder()
	response.AddHeader(1, fmt.Sprintf("CSS Integration Guide for `%s`", cssArgs.TagName))
	response.AddDescription(element)

	// Check available CSS features
	hasCustomProperties := len(element.CssProperties()) > 0
	hasCustomParts := len(element.CssParts()) > 0
	hasCustomStates := len(element.CssStates()) > 0

	if !hasCustomProperties && !hasCustomParts && !hasCustomStates {
		response.AddSection("❌ This element doesn't define CSS custom properties, parts, or states in the manifest.\n")
		response.AddSection("You can still style it using regular CSS selectors.\n")

		// Use template for basic styling guidance
		templateData := NewCSSTemplateData(element, cssArgs.Context, cssArgs.Options, cssArgs.Theme)
		response.AddTemplateSection("basic_styling", templateData)

		return response.Build(), nil
	}

	// Prepare template data
	templateData := NewCSSTemplateData(element, cssArgs.Context, cssArgs.Options, cssArgs.Theme)

	// Generate CSS custom properties section
	if hasCustomProperties {
		response.AddTemplateSection("css_properties", templateData)
	}

	// Generate CSS parts section
	if hasCustomParts {
		response.AddTemplateSection("css_parts", templateData)
	}

	// Generate CSS states section
	if hasCustomStates {
		response.AddTemplateSection("css_states", templateData)
	}

	// Generate theming guidance
	response.AddTemplateSection("theming_guidance", templateData)

	// Generate responsive guidance if context suggests it
	if strings.Contains(strings.ToLower(cssArgs.Context), "responsive") || cssArgs.Options["responsive"] == "true" {
		response.AddTemplateSection("responsive_guidance", templateData)
	}

	return response.Build(), nil
}
