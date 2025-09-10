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

		// Generate usage guidance
		section.WriteString("**Usage:**\n")
		section.WriteString("```css\n")
		section.WriteString(fmt.Sprintf("%s {\n", element.TagName()))
		if prop.Initial() != "" {
			section.WriteString(fmt.Sprintf("  %s: %s; /* default value */\n", prop.Name(), prop.Initial()))
		} else {
			section.WriteString(fmt.Sprintf("  %s: /* set appropriate value based on %s */;\n", prop.Name(), prop.Syntax()))
		}
		section.WriteString("}\n")
		section.WriteString("```\n\n")
	}

	// Generate theming guidance
	section.WriteString("### Theming Guidance\n")
	section.WriteString("When customizing these properties, consider:\n")
	section.WriteString("- The element's intended purpose and context\n")
	section.WriteString("- Accessibility requirements (contrast, readability)\n")
	section.WriteString("- Consistency with your design system\n")
	section.WriteString("- The property's syntax requirements and constraints\n\n")

	section.WriteString("**Example Template:**\n")
	section.WriteString("```css\n")
	section.WriteString(fmt.Sprintf("%s {\n", element.TagName()))
	for _, prop := range properties {
		if prop.Initial() != "" {
			section.WriteString(fmt.Sprintf("  %s: %s; /* modify as needed */\n", prop.Name(), prop.Initial()))
		} else {
			section.WriteString(fmt.Sprintf("  %s: /* provide value matching syntax: %s */;\n", prop.Name(), prop.Syntax()))
		}
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
		if part.Description() != "" {
			section.WriteString(fmt.Sprintf("  /* %s */\n", part.Description()))
		}
		section.WriteString("  /* Add your styling here */\n")
		section.WriteString("}\n")
		section.WriteString("```\n\n")
	}

	// Generate comprehensive parts styling guidance
	if len(parts) > 1 {
		section.WriteString("### Styling Multiple Parts\n")
		section.WriteString("Consider how these parts work together in the element's design.\n")
		section.WriteString("Each part represents a distinct visual component that can be styled independently.\n\n")
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
		if state.Description() != "" {
			section.WriteString(fmt.Sprintf("  /* %s */\n", state.Description()))
		} else {
			section.WriteString(fmt.Sprintf("  /* Styles when element is in %s state */\n", state.Name()))
		}
		section.WriteString("  /* Add your styling here */\n")
		section.WriteString("}\n")
		section.WriteString("```\n\n")
	}

	return section.String()
}

// generateThemeExamples creates theme-specific styling guidance
func generateThemeExamples(element types.ElementInfo, args SuggestCssIntegrationArgs) string {
	var section strings.Builder
	section.WriteString("## 🌈 Theming Guidance\n\n")

	// Only show specific theme if requested
	if args.Theme != "" {
		section.WriteString(fmt.Sprintf("### %s Theme Considerations\n", strings.Title(args.Theme)))
	} else {
		section.WriteString("### Multi-Theme Support\n")
	}

	section.WriteString("When implementing themes for this element, consider:\n\n")

	if len(element.CssProperties()) > 0 {
		section.WriteString("**Custom Properties:** Each property has specific constraints\n")
		for _, prop := range element.CssProperties() {
			section.WriteString(fmt.Sprintf("- `%s`", prop.Name()))
			if prop.Syntax() != "" {
				section.WriteString(fmt.Sprintf(" (syntax: %s)", prop.Syntax()))
			}
			if prop.Description() != "" {
				section.WriteString(fmt.Sprintf(": %s", prop.Description()))
			}
			section.WriteString("\n")
		}
		section.WriteString("\n")
	}

	if len(element.CssParts()) > 0 {
		section.WriteString("**Styleable Parts:** Consider visual hierarchy and relationships\n")
		for _, part := range element.CssParts() {
			section.WriteString(fmt.Sprintf("- `%s`", part.Name()))
			if part.Description() != "" {
				section.WriteString(fmt.Sprintf(": %s", part.Description()))
			}
			section.WriteString("\n")
		}
		section.WriteString("\n")
	}

	if len(element.CssStates()) > 0 {
		section.WriteString("**Interactive States:** Consider user feedback and transitions\n")
		for _, state := range element.CssStates() {
			section.WriteString(fmt.Sprintf("- `:%s`", state.Name()))
			if state.Description() != "" {
				section.WriteString(fmt.Sprintf(": %s", state.Description()))
			}
			section.WriteString("\n")
		}
		section.WriteString("\n")
	}

	section.WriteString("**Best Practices:**\n")
	section.WriteString("- Use semantic color names and design tokens\n")
	section.WriteString("- Ensure sufficient contrast ratios for accessibility\n")
	section.WriteString("- Test themes with the element's intended content\n")
	section.WriteString("- Consider animation and transition preferences\n\n")

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
	// Use initial value from manifest if available - this is the most reliable
	if initial := prop.Initial(); initial != "" {
		return initial
	}

	// Use syntax to generate appropriate placeholder
	if syntax := prop.Syntax(); syntax != "" {
		return generateValueFromSyntax(syntax, prop.Name())
	}

	// If no manifest data available, provide semantic placeholder
	return fmt.Sprintf("/* your %s value */", prop.Name())
}

// generateValueFromSyntax creates example values based on CSS syntax definitions
func generateValueFromSyntax(syntax, propName string) string {
	syntax = strings.TrimSpace(syntax)
	propName = strings.ToLower(propName)

	// Handle union types first (e.g., "small | medium | large")
	if strings.Contains(syntax, "|") {
		options := strings.Split(syntax, "|")
		for _, option := range options {
			option = strings.TrimSpace(option)
			// Return the first concrete value (not a generic type)
			if !strings.Contains(option, "<") && option != "" {
				return option
			}
		}
	}

	// For generic types, provide semantic placeholders
	syntax = strings.ToLower(syntax)
	switch {
	case strings.Contains(syntax, "<color>"):
		return "/* your color value */"
	case strings.Contains(syntax, "<length>"):
		return "/* your length value (e.g., 16px, 1rem) */"
	case strings.Contains(syntax, "<percentage>"):
		return "/* your percentage value (e.g., 50%) */"
	case strings.Contains(syntax, "<number>"):
		return "/* your number value */"
	case strings.Contains(syntax, "<integer>"):
		return "/* your integer value */"
	case strings.Contains(syntax, "<time>"):
		return "/* your time value (e.g., 0.3s, 200ms) */"
	case strings.Contains(syntax, "<angle>"):
		return "/* your angle value (e.g., 45deg, 0.5turn) */"
	case strings.Contains(syntax, "<string>"):
		return "/* your string value */"
	case strings.Contains(syntax, "<url>"):
		return "/* url('path/to/resource') */"
	default:
		// If syntax doesn't match known patterns, use it as-is if it looks like a value
		if syntax != "" && !strings.Contains(syntax, "<") {
			return syntax
		}
		return "/* your value here */"
	}
}

func generateExamplePartStyles(part types.CssPart) string {
	// Use description to guide styling suggestions
	if desc := part.Description(); desc != "" {
		return fmt.Sprintf("  /* %s */\n  /* Add your styles here */\n", desc)
	}

	// Fallback to basic comment
	return fmt.Sprintf("  /* Styles for %s part */\n  /* Add your styles here */\n", part.Name())
}

func generateExampleStateStyles(state types.CssState) string {
	// Use description to guide styling suggestions
	if desc := state.Description(); desc != "" {
		return fmt.Sprintf("  /* %s */\n  /* Add your styles here */\n", desc)
	}

	// Fallback to basic comment
	return fmt.Sprintf("  /* Styles for %s state */\n  /* Add your styles here */\n", state.Name())
}

func generateThemeSpecificValue(prop types.CssProperty, theme string) string {
	// Start with the base example value from manifest
	baseValue := generateExampleValue(prop)
	
	// If the property has a specific initial value, use theme variations of it
	if initial := prop.Initial(); initial != "" {
		return generateThemeVariation(initial, theme, prop.Name())
	}

	// Use syntax-aware theme generation
	if syntax := prop.Syntax(); syntax != "" {
		return generateThemeValueFromSyntax(syntax, theme, prop.Name())
	}

	// Fallback to name-based theming
	return generateThemeVariation(baseValue, theme, prop.Name())
}

// generateThemeVariation applies theme-specific modifications to a base value
func generateThemeVariation(baseValue, theme, propName string) string {
	// If baseValue is already a comment placeholder, enhance it with theme context
	if strings.Contains(baseValue, "/*") {
		return fmt.Sprintf("/* your %s value for %s theme */", propName, theme)
	}

	// If we have a concrete value, return it as-is for any theme
	// The element author knows best what values are appropriate
	return baseValue
}

// generateThemeValueFromSyntax creates theme-specific values based on CSS syntax
func generateThemeValueFromSyntax(syntax, theme, propName string) string {
	// For any syntax type, provide theme-aware placeholder
	if strings.Contains(strings.ToLower(syntax), "<color>") {
		return fmt.Sprintf("/* your %s color value for %s theme */", propName, theme)
	}

	// For non-color properties, return syntax-appropriate value with theme context
	baseValue := generateValueFromSyntax(syntax, propName)
	if strings.Contains(baseValue, "/*") {
		return strings.Replace(baseValue, "*/", fmt.Sprintf(" for %s theme */", theme), 1)
	}
	return baseValue
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
