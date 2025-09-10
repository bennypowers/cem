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
		response.WriteString(generateBasicStylingGuidance(element, cssArgs))
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: response.String(),
				},
			},
		}, nil
	}

	// Generate CSS custom properties section
	if hasCustomProperties {
		response.WriteString(generateCustomPropertiesSection(element, cssArgs))
	}

	// Generate CSS parts section
	if hasCustomParts {
		response.WriteString(generateCssPartsSection(element, cssArgs))
	}

	// Generate CSS states section
	if hasCustomStates {
		response.WriteString(generateCssStatesSection(element, cssArgs))
	}

	// Generate theme examples
	response.WriteString(generateThemeExamples(element, cssArgs))

	// Generate responsive design guidance
	response.WriteString(generateResponsiveGuidance(element, cssArgs))

	// Generate accessibility considerations
	response.WriteString(generateCssAccessibilityGuidance(element, cssArgs))

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: response.String(),
			},
		},
	}, nil
}

// generateCustomPropertiesSection creates guidance for CSS custom properties
func generateCustomPropertiesSection(
	element types.ElementInfo,
	args SuggestCssIntegrationArgs,
) string {
	var section strings.Builder
	section.WriteString("## 🎨 CSS Custom Properties\n\n")

	properties := element.CssProperties()
	section.WriteString(fmt.Sprintf("This element supports %d CSS custom properties for theming:\n\n", len(properties)))

	for _, prop := range properties {
		section.WriteString(fmt.Sprintf("### `%s`\n", prop.Name()))
		if prop.Description() != "" {
			section.WriteString(fmt.Sprintf("%s\n\n", prop.Description()))
		}

		if prop.Syntax() != "" {
			section.WriteString(fmt.Sprintf("**Syntax:** `%s`\n", prop.Syntax()))
		}
		if prop.Initial() != "" {
			section.WriteString(fmt.Sprintf("**Default Value:** `%s`\n", prop.Initial()))
		}
		section.WriteString(fmt.Sprintf("**Inherits:** %t\n\n", prop.Inherits()))

		// Generate usage example
		section.WriteString("**Example:**\n")
		section.WriteString("```css\n")
		section.WriteString(fmt.Sprintf("%s {\n", element.TagName()))
		if prop.Initial() != "" {
			section.WriteString(fmt.Sprintf("  %s: %s; /* custom value */\n", prop.Name(), generateExampleValue(prop)))
		} else {
			section.WriteString(fmt.Sprintf("  %s: your-value;\n", prop.Name()))
		}
		section.WriteString("}\n")
		section.WriteString("```\n\n")
	}

	// Generate comprehensive theming example
	section.WriteString("### Complete Theming Example\n")
	section.WriteString("```css\n")
	section.WriteString(fmt.Sprintf("%s {\n", element.TagName()))
	for _, prop := range properties {
		exampleValue := generateExampleValue(prop)
		section.WriteString(fmt.Sprintf("  %s: %s;\n", prop.Name(), exampleValue))
	}
	section.WriteString("}\n")
	section.WriteString("```\n\n")

	return section.String()
}

// generateCssPartsSection creates guidance for CSS parts
func generateCssPartsSection(
	element types.ElementInfo,
	args SuggestCssIntegrationArgs,
) string {
	var section strings.Builder
	section.WriteString("## 🔧 CSS Parts\n\n")

	parts := element.CssParts()
	section.WriteString(fmt.Sprintf("This element exposes %d CSS parts for granular styling:\n\n", len(parts)))

	for _, part := range parts {
		section.WriteString(fmt.Sprintf("### `%s`\n", part.Name()))
		if part.Description() != "" {
			section.WriteString(fmt.Sprintf("%s\n\n", part.Description()))
		}

		section.WriteString("**Usage:**\n")
		section.WriteString("```css\n")
		section.WriteString(fmt.Sprintf("%s::%s {\n", element.TagName(), part.Name()))
		section.WriteString("  /* Style the " + part.Name() + " part */\n")
		section.WriteString(generateExamplePartStyles(part))
		section.WriteString("}\n")
		section.WriteString("```\n\n")
	}

	// Generate comprehensive parts styling example
	if len(parts) > 1 {
		section.WriteString("### Complete Parts Styling\n")
		section.WriteString("```css\n")
		for _, part := range parts {
			section.WriteString(fmt.Sprintf("%s::%s {\n", element.TagName(), part.Name()))
			section.WriteString(generateExamplePartStyles(part))
			section.WriteString("}\n\n")
		}
		section.WriteString("```\n\n")
	}

	return section.String()
}

// generateCssStatesSection creates guidance for CSS states
func generateCssStatesSection(
	element types.ElementInfo,
	args SuggestCssIntegrationArgs,
) string {
	var section strings.Builder
	section.WriteString("## 🔄 CSS Custom States\n\n")

	states := element.CssStates()
	section.WriteString(fmt.Sprintf("This element supports %d custom states:\n\n", len(states)))

	for _, state := range states {
		section.WriteString(fmt.Sprintf("### `:%s`\n", state.Name()))
		if state.Description() != "" {
			section.WriteString(fmt.Sprintf("%s\n\n", state.Description()))
		}

		section.WriteString("**Usage:**\n")
		section.WriteString("```css\n")
		section.WriteString(fmt.Sprintf("%s:%s {\n", element.TagName(), state.Name()))
		section.WriteString("  /* Styles when element is in " + state.Name() + " state */\n")
		section.WriteString(generateExampleStateStyles(state))
		section.WriteString("}\n")
		section.WriteString("```\n\n")
	}

	return section.String()
}

// generateThemeExamples creates theme-specific styling examples
func generateThemeExamples(element types.ElementInfo, args SuggestCssIntegrationArgs) string {
	var section strings.Builder
	section.WriteString("## 🌈 Theme Examples\n\n")

	themes := []string{"light", "dark", "high-contrast"}
	if args.Theme != "" {
		themes = []string{args.Theme}
	}

	for _, theme := range themes {
		section.WriteString(fmt.Sprintf("### %s Theme\n", titleCase.String(theme)))
		section.WriteString("```css\n")

		// Theme-specific custom properties
		if len(element.CssProperties()) > 0 {
			section.WriteString(fmt.Sprintf("[data-theme=\"%s\"] %s {\n", theme, element.TagName()))
			for _, prop := range element.CssProperties() {
				value := generateThemeSpecificValue(prop, theme)
				section.WriteString(fmt.Sprintf("  %s: %s;\n", prop.Name(), value))
			}
			section.WriteString("}\n\n")
		}

		// Theme-specific part styling
		if len(element.CssParts()) > 0 {
			for _, part := range element.CssParts() {
				section.WriteString(fmt.Sprintf("[data-theme=\"%s\"] %s::%s {\n", theme, element.TagName(), part.Name()))
				section.WriteString(generateThemeSpecificPartStyles(part, theme))
				section.WriteString("}\n\n")
			}
		}

		section.WriteString("```\n\n")
	}

	return section.String()
}

// generateResponsiveGuidance creates responsive design guidance
func generateResponsiveGuidance(
	element types.ElementInfo,
	args SuggestCssIntegrationArgs,
) string {
	var section strings.Builder
	section.WriteString("## 📱 Responsive Design\n\n")

	section.WriteString("### Container Queries (Recommended)\n")
	section.WriteString("```css\n")
	section.WriteString(fmt.Sprintf("%s {\n", element.TagName()))
	section.WriteString("  container-type: inline-size;\n")
	section.WriteString("}\n\n")

	if len(element.CssProperties()) > 0 {
		section.WriteString(fmt.Sprintf("@container (min-width: 400px) {\n"))
		section.WriteString(fmt.Sprintf("  %s {\n", element.TagName()))
		for _, prop := range element.CssProperties() {
			if strings.Contains(strings.ToLower(prop.Name()), "size") ||
				strings.Contains(strings.ToLower(prop.Name()), "width") ||
				strings.Contains(strings.ToLower(prop.Name()), "space") {
				section.WriteString(fmt.Sprintf("    %s: larger-value;\n", prop.Name()))
			}
		}
		section.WriteString("  }\n")
		section.WriteString("}\n")
	}
	section.WriteString("```\n\n")

	section.WriteString("### Media Queries (Fallback)\n")
	section.WriteString("```css\n")
	section.WriteString("@media (max-width: 768px) {\n")
	section.WriteString(fmt.Sprintf("  %s {\n", element.TagName()))
	if len(element.CssProperties()) > 0 {
		for _, prop := range element.CssProperties() {
			if strings.Contains(strings.ToLower(prop.Name()), "size") ||
				strings.Contains(strings.ToLower(prop.Name()), "space") {
				section.WriteString(fmt.Sprintf("    %s: mobile-value;\n", prop.Name()))
			}
		}
	}
	section.WriteString("  }\n")
	section.WriteString("}\n")
	section.WriteString("```\n\n")

	return section.String()
}

// generateCssAccessibilityGuidance creates accessibility-focused CSS guidance
func generateCssAccessibilityGuidance(
	element types.ElementInfo,
	args SuggestCssIntegrationArgs,
) string {
	var section strings.Builder
	section.WriteString("## ♿ Accessibility Considerations\n\n")

	section.WriteString("### Color Contrast\n")
	section.WriteString("Ensure sufficient contrast ratios when customizing colors:\n")
	section.WriteString("```css\n")
	section.WriteString(fmt.Sprintf("%s {\n", element.TagName()))
	if len(element.CssProperties()) > 0 {
		for _, prop := range element.CssProperties() {
			if strings.Contains(strings.ToLower(prop.Name()), "color") ||
				strings.Contains(strings.ToLower(prop.Name()), "background") {
				section.WriteString(fmt.Sprintf("  %s: color-contrast(your-color vs white, black);\n", prop.Name()))
			}
		}
	}
	section.WriteString("}\n")
	section.WriteString("```\n\n")

	section.WriteString("### Focus Indicators\n")
	section.WriteString("```css\n")
	section.WriteString(fmt.Sprintf("%s:focus {\n", element.TagName()))
	section.WriteString("  outline: 2px solid var(--focus-color, #005fcc);\n")
	section.WriteString("  outline-offset: 2px;\n")
	section.WriteString("}\n\n")

	section.WriteString(fmt.Sprintf("%s:focus:not(:focus-visible) {\n", element.TagName()))
	section.WriteString("  outline: none;\n")
	section.WriteString("}\n")
	section.WriteString("```\n\n")

	section.WriteString("### Reduced Motion\n")
	section.WriteString("```css\n")
	section.WriteString("@media (prefers-reduced-motion: reduce) {\n")
	section.WriteString(fmt.Sprintf("  %s {\n", element.TagName()))
	if len(element.CssProperties()) > 0 {
		for _, prop := range element.CssProperties() {
			if strings.Contains(strings.ToLower(prop.Name()), "animation") ||
				strings.Contains(strings.ToLower(prop.Name()), "transition") {
				section.WriteString(fmt.Sprintf("    %s: none;\n", prop.Name()))
			}
		}
	}
	section.WriteString("  }\n")
	section.WriteString("}\n")
	section.WriteString("```\n\n")

	return section.String()
}

// Helper functions for generating example values

func generateExampleValue(prop types.CssProperty) string {
	propName := strings.ToLower(prop.Name())
	switch {
	case strings.Contains(propName, "color"):
		return "#007acc"
	case strings.Contains(propName, "size") || strings.Contains(propName, "width") || strings.Contains(propName, "height"):
		return "16px"
	case strings.Contains(propName, "space") || strings.Contains(propName, "padding") || strings.Contains(propName, "margin"):
		return "8px"
	case strings.Contains(propName, "font"):
		return "Inter, sans-serif"
	case strings.Contains(propName, "weight"):
		return "500"
	case strings.Contains(propName, "radius"):
		return "4px"
	case strings.Contains(propName, "shadow"):
		return "0 2px 4px rgba(0,0,0,0.1)"
	default:
		if prop.Initial() != "" {
			return "custom-value"
		}
		return "your-value"
	}
}

func generateExamplePartStyles(part types.CssPart) string {
	partName := strings.ToLower(part.Name())
	switch {
	case strings.Contains(partName, "button"):
		return "  background: #007acc;\n  color: white;\n  border-radius: 4px;\n"
	case strings.Contains(partName, "input"):
		return "  border: 1px solid #ccc;\n  padding: 8px;\n  border-radius: 4px;\n"
	case strings.Contains(partName, "text"):
		return "  font-size: 16px;\n  color: #333;\n"
	case strings.Contains(partName, "icon"):
		return "  width: 16px;\n  height: 16px;\n  fill: currentColor;\n"
	default:
		return "  /* Custom styles for " + part.Name() + " */\n"
	}
}

func generateExampleStateStyles(state types.CssState) string {
	stateName := strings.ToLower(state.Name())
	switch {
	case strings.Contains(stateName, "active"):
		return "  background: #005fcc;\n  transform: scale(0.98);\n"
	case strings.Contains(stateName, "disabled"):
		return "  opacity: 0.5;\n  pointer-events: none;\n"
	case strings.Contains(stateName, "hover"):
		return "  background: #0066cc;\n"
	case strings.Contains(stateName, "loading"):
		return "  opacity: 0.7;\n  cursor: wait;\n"
	default:
		return "  /* Styles for " + state.Name() + " state */\n"
	}
}

func generateThemeSpecificValue(prop types.CssProperty, theme string) string {
	propName := strings.ToLower(prop.Name())
	switch theme {
	case "dark":
		switch {
		case strings.Contains(propName, "background"):
			return "#1a1a1a"
		case strings.Contains(propName, "color") && !strings.Contains(propName, "background"):
			return "#ffffff"
		case strings.Contains(propName, "border"):
			return "#333"
		default:
			return generateExampleValue(prop)
		}
	case "high-contrast":
		switch {
		case strings.Contains(propName, "background"):
			return "#000000"
		case strings.Contains(propName, "color") && !strings.Contains(propName, "background"):
			return "#ffffff"
		case strings.Contains(propName, "border"):
			return "#ffffff"
		default:
			return generateExampleValue(prop)
		}
	default: // light theme
		return generateExampleValue(prop)
	}
}

func generateThemeSpecificPartStyles(part types.CssPart, theme string) string {
	base := generateExamplePartStyles(part)
	switch theme {
	case "dark":
		return strings.ReplaceAll(strings.ReplaceAll(base, "#007acc", "#66b3ff"), "#333", "#ccc")
	case "high-contrast":
		return strings.ReplaceAll(strings.ReplaceAll(base, "#007acc", "#ffffff"), "#ccc", "#ffffff")
	default:
		return base
	}
}

// generateBasicStylingGuidance provides guidance when no CSS features are defined
func generateBasicStylingGuidance(
	element types.ElementInfo,
	args SuggestCssIntegrationArgs,
) string {
	var section strings.Builder
	section.WriteString("## Basic Styling Guidance\n\n")

	section.WriteString("Since this element doesn't define CSS custom properties or parts, ")
	section.WriteString("you can style it using standard CSS selectors:\n\n")

	section.WriteString("```css\n")
	section.WriteString(fmt.Sprintf("%s {\n", element.TagName()))
	section.WriteString("  /* Basic styling */\n")
	section.WriteString("  display: block;\n")
	section.WriteString("  padding: 1rem;\n")
	section.WriteString("  border: 1px solid #ccc;\n")
	section.WriteString("  border-radius: 4px;\n")
	section.WriteString("}\n\n")

	section.WriteString(fmt.Sprintf("%s:hover {\n", element.TagName()))
	section.WriteString("  border-color: #007acc;\n")
	section.WriteString("}\n\n")

	section.WriteString(fmt.Sprintf("%s:focus {\n", element.TagName()))
	section.WriteString("  outline: 2px solid #005fcc;\n")
	section.WriteString("  outline-offset: 2px;\n")
	section.WriteString("}\n")
	section.WriteString("```\n\n")

	return section.String()
}
