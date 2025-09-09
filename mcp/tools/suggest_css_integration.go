package mcp

import (
	"context"
	"fmt"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// SuggestCssIntegrationArgs represents the arguments for the suggest_css_integration tool
type SuggestCssIntegrationArgs struct {
	TagName     string `json:"tagName"`
	StyleTarget string `json:"styleTarget,omitempty"`
	Context     string `json:"context,omitempty"`
}

// handleSuggestCssIntegration provides CSS integration guidance for custom elements
func (s *SimpleCEMServer) handleSuggestCssIntegration(ctx context.Context, req *mcp.CallToolRequest, args SuggestCssIntegrationArgs) (*mcp.CallToolResult, any, error) {
	element, err := s.registry.GetElementInfo(args.TagName)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Element '%s' not found in registry", args.TagName)},
			},
		}, nil, nil
	}

	var response strings.Builder
	response.WriteString(fmt.Sprintf("CSS Integration Guide for <%s>:\n\n", args.TagName))

	// CSS Custom Properties
	cssProps := element.CssProperties()
	if len(cssProps) > 0 {
		response.WriteString("CSS Custom Properties (Variables):\n")
		for _, prop := range cssProps {
			response.WriteString(fmt.Sprintf("• %s\n", prop.Name()))
			if desc := prop.Description(); desc != "" {
				response.WriteString(fmt.Sprintf("  Description: %s\n", desc))
			}
			response.WriteString(fmt.Sprintf("  Syntax: %s\n", prop.Syntax()))
			response.WriteString(fmt.Sprintf("  Initial value: %s\n", prop.Initial()))
			response.WriteString(fmt.Sprintf("  Inherits: %t\n", prop.Inherits()))
			
			// Usage example
			response.WriteString("  Example:\n")
			response.WriteString(fmt.Sprintf("    %s {\n", args.TagName))
			response.WriteString(fmt.Sprintf("      %s: %s;\n", prop.Name(), s.getSampleValue(prop.Syntax())))
			response.WriteString("    }\n\n")
		}
	} else {
		response.WriteString("No CSS custom properties defined for this element.\n\n")
	}

	// CSS Parts
	cssParts := element.CssParts()
	if len(cssParts) > 0 {
		response.WriteString("CSS Parts (::part() selectors):\n")
		for _, part := range cssParts {
			response.WriteString(fmt.Sprintf("• %s\n", part.Name()))
			if desc := part.Description(); desc != "" {
				response.WriteString(fmt.Sprintf("  Description: %s\n", desc))
			}
			
			// Usage example
			response.WriteString("  Example:\n")
			response.WriteString(fmt.Sprintf("    %s::%s {\n", args.TagName, part.Name()))
			response.WriteString("      /* Style the internal part */\n")
			response.WriteString("      color: var(--primary-color);\n")
			response.WriteString("      font-weight: bold;\n")
			response.WriteString("    }\n\n")
		}
	} else {
		response.WriteString("No CSS parts defined for this element.\n\n")
	}

	// CSS States
	cssStates := element.CssStates()
	if len(cssStates) > 0 {
		response.WriteString("CSS Custom States (:--state selectors):\n")
		for _, state := range cssStates {
			response.WriteString(fmt.Sprintf("• %s\n", state.Name()))
			if desc := state.Description(); desc != "" {
				response.WriteString(fmt.Sprintf("  Description: %s\n", desc))
			}
			
			// Usage example
			response.WriteString("  Example:\n")
			response.WriteString(fmt.Sprintf("    %s:%s {\n", args.TagName, state.Name()))
			response.WriteString("      /* Style when element is in this state */\n")
			response.WriteString("      opacity: 0.5;\n")
			response.WriteString("    }\n\n")
		}
	} else {
		response.WriteString("No CSS custom states defined for this element.\n\n")
	}

	// Context-specific recommendations
	if args.Context != "" {
		response.WriteString(fmt.Sprintf("Recommendations for %s context:\n", args.Context))
		s.addCssContextRecommendations(&response, args.Context, args.TagName, element)
		response.WriteString("\n")
	}

	// Style target specific guidance
	if args.StyleTarget != "" {
		response.WriteString(fmt.Sprintf("Specific guidance for styling %s:\n", args.StyleTarget))
		s.addStyleTargetGuidance(&response, args.StyleTarget, args.TagName, element)
		response.WriteString("\n")
	}

	// General CSS integration best practices
	response.WriteString("CSS Integration Best Practices:\n")
	response.WriteString("• Use CSS custom properties for theming and consistent values\n")
	response.WriteString("• Style CSS parts instead of internal elements when possible\n")
	response.WriteString("• Respect the component's design system and intended styling API\n")
	response.WriteString("• Test styling with different themes and accessibility settings\n")
	response.WriteString("• Consider cascade and inheritance when setting custom properties\n")
	
	if len(cssProps) > 0 || len(cssParts) > 0 {
		response.WriteString("• This element provides official styling APIs - prefer these over global CSS\n")
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: response.String()},
		},
	}, nil, nil
}

// getSampleValue returns sample values for different CSS syntax types
func (s *SimpleCEMServer) getSampleValue(syntax string) string {
	switch {
	case strings.Contains(syntax, "color"):
		return "#007acc"
	case strings.Contains(syntax, "length"):
		return "16px"
	case strings.Contains(syntax, "percentage"):
		return "100%"
	case strings.Contains(syntax, "number"):
		return "1.5"
	case strings.Contains(syntax, "string"):
		return "\"value\""
	case strings.Contains(syntax, "url"):
		return "url(./image.png)"
	case strings.Contains(syntax, "time"):
		return "300ms"
	case strings.Contains(syntax, "angle"):
		return "45deg"
	default:
		return "value"
	}
}

// addCssContextRecommendations adds context-specific CSS recommendations
func (s *SimpleCEMServer) addCssContextRecommendations(response *strings.Builder, context, tagName string, element *ElementInfo) {
	switch context {
	case "theme":
		response.WriteString("• Use CSS custom properties to create consistent theming\n")
		response.WriteString("• Define color schemes using CSS custom properties\n")
		response.WriteString("• Consider dark mode variants for all custom properties\n")
		response.WriteString("• Group related properties under common naming conventions\n")
		
		if len(element.CssProperties()) > 0 {
			response.WriteString("\nExample theme integration:\n")
			response.WriteString(":root {\n")
			for _, prop := range element.CssProperties() {
				response.WriteString(fmt.Sprintf("  %s: %s; /* Light theme */\n", prop.Name(), prop.Initial()))
			}
			response.WriteString("}\n\n")
			response.WriteString("[data-theme=\"dark\"] {\n")
			for _, prop := range element.CssProperties() {
				response.WriteString(fmt.Sprintf("  %s: %s; /* Dark theme variant */\n", prop.Name(), s.getDarkModeValue(prop.Initial())))
			}
			response.WriteString("}\n")
		}

	case "responsive":
		response.WriteString("• Use container queries for component-level responsive design\n")
		response.WriteString("• Leverage CSS custom properties for responsive values\n")
		response.WriteString("• Consider mobile-first approach with progressive enhancement\n")
		response.WriteString("• Test component at different viewport sizes\n")

	case "dark-mode":
		response.WriteString("• Override CSS custom properties for dark mode variants\n")
		response.WriteString("• Ensure sufficient color contrast in dark themes\n")
		response.WriteString("• Test with system dark mode preferences\n")
		response.WriteString("• Consider reduced motion preferences\n")
	}
}

// addStyleTargetGuidance adds guidance for specific styling targets
func (s *SimpleCEMServer) addStyleTargetGuidance(response *strings.Builder, target, tagName string, element *ElementInfo) {
	switch target {
	case "element":
		response.WriteString("• Style the element container using standard CSS selectors\n")
		response.WriteString("• Avoid styling internal structure directly\n")
		response.WriteString("• Use CSS custom properties to communicate values to the component\n")
		
	case "parts":
		if len(element.CssParts()) > 0 {
			response.WriteString("• Use ::part() pseudo-element to style internal component parts\n")
			response.WriteString("• Parts provide controlled access to internal styling\n")
			response.WriteString("• Respect the component's styling API and design intentions\n")
		} else {
			response.WriteString("• This element doesn't expose CSS parts for styling\n")
			response.WriteString("• Use CSS custom properties or element-level styling instead\n")
		}
		
	case "states":
		if len(element.CssStates()) > 0 {
			response.WriteString("• Use :--state pseudo-class selectors for state-based styling\n")
			response.WriteString("• States provide semantic styling hooks for different component states\n")
			response.WriteString("• Combine with other selectors for complex state styling\n")
		} else {
			response.WriteString("• This element doesn't expose custom CSS states\n")
			response.WriteString("• Use standard pseudo-classes (:hover, :focus, etc.) instead\n")
		}
		
	case "properties":
		if len(element.CssProperties()) > 0 {
			response.WriteString("• Set CSS custom properties to customize component appearance\n")
			response.WriteString("• Properties cascade and can be inherited from parent elements\n")
			response.WriteString("• Use CSS calc() and other functions with custom properties\n")
		} else {
			response.WriteString("• This element doesn't define CSS custom properties\n")
			response.WriteString("• Use standard CSS properties on the element selector\n")
		}
	}
}

// getDarkModeValue suggests dark mode variants for CSS values
func (s *SimpleCEMServer) getDarkModeValue(lightValue string) string {
	// Basic dark mode value suggestions
	switch lightValue {
	case "white", "#ffffff", "#fff":
		return "#1a1a1a"
	case "black", "#000000", "#000":
		return "#ffffff"
	case "blue", "#007acc":
		return "#4db8ff"
	default:
		return "/* dark mode variant */"
	}
}