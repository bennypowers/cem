package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// GenerateAccessibleHtmlArgs represents the arguments for the generate_accessible_html tool
type GenerateAccessibleHtmlArgs struct {
	TagName    string            `json:"tagName"`
	Content    map[string]string `json:"content,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
	Context    string            `json:"context,omitempty"`
}

// handleGenerateAccessibleHtml generates semantically correct HTML with accessibility features
func (s *SimpleCEMServer) handleGenerateAccessibleHtml(ctx context.Context, req *mcp.CallToolRequest, args GenerateAccessibleHtmlArgs) (*mcp.CallToolResult, any, error) {
	element, err := s.registry.GetElementInfo(args.TagName)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Element '%s' not found in registry", args.TagName)},
			},
		}, nil, nil
	}

	var html strings.Builder
	html.WriteString(fmt.Sprintf("<%s", args.TagName))

	// Add validated attributes
	for attrName, attrValue := range args.Attributes {
		// Find the attribute in the manifest
		var found bool
		for _, attr := range element.Attributes() {
			if attr.Name() == attrName {
				found = true
				// Validate enum values if available
				if values := attr.Values(); len(values) > 0 {
					validValue := false
					for _, validVal := range values {
						if attrValue == validVal {
							validValue = true
							break
						}
					}
					if !validValue {
						return &mcp.CallToolResult{
							Content: []mcp.Content{
								&mcp.TextContent{Text: fmt.Sprintf("Invalid value '%s' for attribute '%s'. Valid values: %v", attrValue, attrName, values)},
							},
						}, nil, nil
					}
				}
				break
			}
		}
		
		if found {
			html.WriteString(fmt.Sprintf(` %s="%s"`, attrName, attrValue))
		} else {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{Text: fmt.Sprintf("Unknown attribute '%s' for element '%s'", attrName, args.TagName)},
				},
			}, nil, nil
		}
	}

	// Add accessibility attributes based on context
	if args.Context != "" {
		switch args.Context {
		case "form":
			// Add form-related ARIA attributes
			if _, hasRole := args.Attributes["role"]; !hasRole {
				html.WriteString(` role="group"`)
			}
		case "navigation":
			// Add navigation-related ARIA attributes
			if _, hasRole := args.Attributes["role"]; !hasRole {
				html.WriteString(` role="navigation"`)
			}
		case "content":
			// Add content-related ARIA attributes
			if _, hasRole := args.Attributes["role"]; !hasRole {
				html.WriteString(` role="main"`)
			}
		}
	}

	html.WriteString(">")

	// Add slot content in proper order (default slot last)
	var defaultContent string
	for _, slot := range element.Slots() {
		slotName := slot.Name()
		if slotContent, exists := args.Content[slotName]; exists {
			if slotName == "" {
				// Save default slot content for last
				defaultContent = slotContent
			} else {
				// Named slot content
				html.WriteString(fmt.Sprintf(`
  <span slot="%s">%s</span>`, slotName, slotContent))
			}
		}
	}

	// Add default slot content
	if defaultContent != "" {
		html.WriteString(fmt.Sprintf("\n  %s", defaultContent))
	}

	html.WriteString(fmt.Sprintf("\n</%s>", args.TagName))

	// Generate additional guidance
	var guidance strings.Builder
	guidance.WriteString(fmt.Sprintf("Generated accessible HTML for <%s>\n\n", args.TagName))
	guidance.WriteString("HTML:\n")
	guidance.WriteString(html.String())
	guidance.WriteString("\n\nAccessibility Features:\n")
	
	if args.Context != "" {
		guidance.WriteString(fmt.Sprintf("- Applied %s context ARIA roles\n", args.Context))
	}
	
	guidance.WriteString("- Validated all attributes against manifest definitions\n")
	guidance.WriteString("- Placed content in correct slots based on element structure\n")
	
	if len(element.CssProperties()) > 0 {
		guidance.WriteString("\nCSS Custom Properties available:\n")
		for _, prop := range element.CssProperties() {
			guidance.WriteString(fmt.Sprintf("- %s: %s\n", prop.Name(), prop.Description()))
		}
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: guidance.String()},
		},
	}, nil, nil
}