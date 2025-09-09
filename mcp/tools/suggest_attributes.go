package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// SuggestAttributesArgs represents the arguments for the suggest_attributes tool
type SuggestAttributesArgs struct {
	TagName string `json:"tagName"`
	Context string `json:"context,omitempty"`
	Partial string `json:"partial,omitempty"`
}

// handleSuggestAttributes provides intelligent attribute suggestions for custom elements
func (s *SimpleCEMServer) handleSuggestAttributes(ctx context.Context, req *mcp.CallToolRequest, args SuggestAttributesArgs) (*mcp.CallToolResult, any, error) {
	element, err := s.registry.GetElementInfo(args.TagName)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Element '%s' not found in registry", args.TagName)},
			},
		}, nil, nil
	}

	attributes := element.Attributes()
	if len(attributes) == 0 {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Element '%s' has no attributes defined", args.TagName)},
			},
		}, nil, nil
	}

	// Filter attributes by partial match if provided
	var filteredAttributes []Attribute
	if args.Partial != "" {
		for _, attr := range attributes {
			if strings.Contains(attr.Name(), args.Partial) {
				filteredAttributes = append(filteredAttributes, attr)
			}
		}
	} else {
		filteredAttributes = attributes
	}

	if len(filteredAttributes) == 0 {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("No attributes found matching '%s' for element '%s'", args.Partial, args.TagName)},
			},
		}, nil, nil
	}

	// Build comprehensive attribute suggestions
	var response strings.Builder
	response.WriteString(fmt.Sprintf("Attribute suggestions for <%s>:\n\n", args.TagName))

	// Group attributes by priority
	var requiredAttrs, optionalAttrs []Attribute
	for _, attr := range filteredAttributes {
		if attr.Required() {
			requiredAttrs = append(requiredAttrs, attr)
		} else {
			optionalAttrs = append(optionalAttrs, attr)
		}
	}

	// Show required attributes first
	if len(requiredAttrs) > 0 {
		response.WriteString("Required Attributes:\n")
		for _, attr := range requiredAttrs {
			s.formatAttributeSuggestion(&response, attr, args.Context, true)
		}
		response.WriteString("\n")
	}

	// Show optional attributes
	if len(optionalAttrs) > 0 {
		response.WriteString("Optional Attributes:\n")
		for _, attr := range optionalAttrs {
			s.formatAttributeSuggestion(&response, attr, args.Context, false)
		}
	}

	// Add usage examples
	response.WriteString("\nUsage Examples:\n")
	
	// Basic example with required attributes
	response.WriteString(fmt.Sprintf("<%s", args.TagName))
	for _, attr := range requiredAttrs {
		if defaultVal := attr.Default(); defaultVal != "" {
			value := strings.Trim(defaultVal, "\"'")
			response.WriteString(fmt.Sprintf(` %s="%s"`, attr.Name(), value))
		} else if len(attr.Values()) > 0 {
			response.WriteString(fmt.Sprintf(` %s="%s"`, attr.Name(), attr.Values()[0]))
		} else {
			response.WriteString(fmt.Sprintf(` %s="value"`, attr.Name()))
		}
	}
	response.WriteString(fmt.Sprintf("></%s>\n", args.TagName))

	// Context-specific recommendations
	if args.Context != "" {
		response.WriteString(fmt.Sprintf("\nRecommendations for %s context:\n", args.Context))
		s.addContextualRecommendations(&response, args.Context, filteredAttributes)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: response.String()},
		},
	}, nil, nil
}

// formatAttributeSuggestion formats a single attribute suggestion with details
func (s *SimpleCEMServer) formatAttributeSuggestion(response *strings.Builder, attr Attribute, context string, required bool) {
	// Attribute name and type
	response.WriteString(fmt.Sprintf("• %s", attr.Name()))
	if required {
		response.WriteString(" (REQUIRED)")
	}
	response.WriteString(fmt.Sprintf(" : %s\n", attr.Type()))
	
	// Description
	if desc := attr.Description(); desc != "" {
		response.WriteString(fmt.Sprintf("  %s\n", desc))
	}
	
	// Default value
	if defaultVal := attr.Default(); defaultVal != "" {
		response.WriteString(fmt.Sprintf("  Default: %s\n", defaultVal))
	}
	
	// Enum values
	if values := attr.Values(); len(values) > 0 {
		response.WriteString(fmt.Sprintf("  Valid values: %v\n", values))
		
		// Add explanations for enum values based on context
		if context != "" && len(values) <= 5 {
			for _, value := range values {
				explanation := s.explainEnumValue(attr.Name(), value, context)
				if explanation != "" {
					response.WriteString(fmt.Sprintf("    • %s: %s\n", value, explanation))
				}
			}
		}
	}
	
	// Usage guidelines from descriptions
	if guidelines := attr.Guidelines(); len(guidelines) > 0 {
		response.WriteString("  Guidelines:\n")
		for _, guideline := range guidelines {
			response.WriteString(fmt.Sprintf("    - %s\n", guideline))
		}
	}
	
	response.WriteString("\n")
}

// explainEnumValue provides contextual explanations for enum values
func (s *SimpleCEMServer) explainEnumValue(attrName, value, context string) string {
	// Basic explanations based on common patterns
	switch {
	case attrName == "variant" || attrName == "type":
		switch value {
		case "primary":
			return "Main action or primary content"
		case "secondary":
			return "Supporting action or secondary content"
		case "danger", "error":
			return "Destructive actions or error states"
		case "success":
			return "Positive actions or success states"
		}
	case attrName == "size":
		switch value {
		case "small", "sm":
			return "Compact size for dense layouts"
		case "medium", "md":
			return "Standard size for most use cases"
		case "large", "lg":
			return "Prominent size for emphasis"
		}
	case attrName == "state" || attrName == "status":
		switch value {
		case "loading":
			return "Element is in a loading state"
		case "disabled":
			return "Element is not interactive"
		case "active":
			return "Element is currently active"
		}
	}
	
	return ""
}

// addContextualRecommendations adds context-specific attribute recommendations
func (s *SimpleCEMServer) addContextualRecommendations(response *strings.Builder, context string, attributes []Attribute) {
	switch context {
	case "form":
		response.WriteString("- Consider accessibility attributes like aria-label or aria-describedby\n")
		response.WriteString("- Use appropriate form validation attributes\n")
		response.WriteString("- Ensure proper labeling for screen readers\n")
	case "navigation":
		response.WriteString("- Add role='navigation' for accessibility\n")
		response.WriteString("- Consider aria-current for active navigation items\n")
		response.WriteString("- Use semantic markup for menu structures\n")
	case "content":
		response.WriteString("- Ensure proper heading hierarchy\n")
		response.WriteString("- Add landmark roles where appropriate\n")
		response.WriteString("- Consider aria-expanded for collapsible content\n")
	}
}