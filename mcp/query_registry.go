package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// QueryRegistryArgs represents the arguments for the query_registry tool
type QueryRegistryArgs struct {
	TagName string `json:"tagName,omitempty"`
	Filter  string `json:"filter,omitempty"`
	Search  string `json:"search,omitempty"`
}

// handleQueryRegistry queries and explores the custom elements registry
func (s *Server) handleQueryRegistry(ctx context.Context, req *mcp.CallToolRequest, args QueryRegistryArgs) (*mcp.CallToolResult, any, error) {
	elements := s.registry.GetAllElements()

	// Handle specific element query
	if args.TagName != "" {
		if element, exists := elements[args.TagName]; exists {
			data, err := json.MarshalIndent(element, "", "  ")
			if err != nil {
				return nil, nil, fmt.Errorf("failed to marshal element: %w", err)
			}

			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{Text: string(data)},
				},
			}, nil, nil
		} else {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{Text: fmt.Sprintf("Element '%s' not found", args.TagName)},
				},
			}, nil, nil
		}
	}

	// Filter and search logic
	var matchingElements []*ElementInfo
	for tagName, element := range elements {
		match := true

		// Apply search filter
		if args.Search != "" {
			searchTerm := strings.ToLower(args.Search)
			elementText := strings.ToLower(fmt.Sprintf("%s %s", tagName, element.Description))
			
			// Also search in attribute and slot descriptions
			for _, attr := range element.Attributes() {
				elementText += " " + strings.ToLower(attr.Name() + " " + attr.Description())
			}
			for _, slot := range element.Slots() {
				elementText += " " + strings.ToLower(slot.Name() + " " + slot.Description())
			}
			
			if !strings.Contains(elementText, searchTerm) {
				match = false
			}
		}

		// Apply feature filter
		if match && args.Filter != "" {
			match = s.matchesFilter(element, args.Filter)
		}

		if match {
			matchingElements = append(matchingElements, element)
		}
	}

	// Build response
	var response strings.Builder

	if args.Search != "" || args.Filter != "" {
		response.WriteString(fmt.Sprintf("Registry Query Results (found %d elements):\n\n", len(matchingElements)))
		if args.Search != "" {
			response.WriteString(fmt.Sprintf("Search term: '%s'\n", args.Search))
		}
		if args.Filter != "" {
			response.WriteString(fmt.Sprintf("Filter: '%s'\n", args.Filter))
		}
		response.WriteString("\n")
	} else {
		response.WriteString(fmt.Sprintf("Custom Elements Registry (%d elements total):\n\n", len(matchingElements)))
	}

	if len(matchingElements) == 0 {
		response.WriteString("No elements found matching the criteria.")
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: response.String()},
			},
		}, nil, nil
	}

	// Summary format for multiple elements
	for _, element := range matchingElements {
		response.WriteString(fmt.Sprintf("<%s>\n", element.TagName))
		if element.Description != "" {
			response.WriteString(fmt.Sprintf("  Description: %s\n", element.Description))
		}
		
		// Capabilities summary
		capabilities := s.buildCapabilitiesSummary(element)
		if capabilities != "" {
			response.WriteString(fmt.Sprintf("  Capabilities: %s\n", capabilities))
		}
		
		// Accessibility features
		accessibilityFeatures := s.getAccessibilityFeatures(element)
		if len(accessibilityFeatures) > 0 {
			response.WriteString(fmt.Sprintf("  Accessibility: %s\n", strings.Join(accessibilityFeatures, ", ")))
		}
		
		response.WriteString("\n")
	}

	// Add usage guidance
	if len(matchingElements) > 1 {
		response.WriteString("Usage:\n")
		response.WriteString("• Use query_registry with tagName parameter for detailed element information\n")
		response.WriteString("• Use suggest_attributes tool for attribute guidance\n")
		response.WriteString("• Use generate_accessible_html tool for HTML generation\n")
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: response.String()},
		},
	}, nil, nil
}

// matchesFilter checks if an element matches the given filter criteria
func (s *Server) matchesFilter(element *ElementInfo, filter string) bool {
	switch filter {
	case "has-slots":
		return len(element.Slots()) > 0
	case "has-events":
		return len(element.Events()) > 0
	case "has-attributes":
		return len(element.Attributes()) > 0
	case "has-css-properties":
		return len(element.CssProperties()) > 0
	case "has-css-parts":
		return len(element.CssParts()) > 0
	case "has-css-states":
		return len(element.CssStates()) > 0
	case "interactive":
		// Elements with events or form-related attributes are considered interactive
		if len(element.Events()) > 0 {
			return true
		}
		for _, attr := range element.Attributes() {
			attrName := strings.ToLower(attr.Name())
			if strings.Contains(attrName, "click") || strings.Contains(attrName, "disabled") || 
			   strings.Contains(attrName, "readonly") || strings.Contains(attrName, "selected") {
				return true
			}
		}
		return false
	case "form":
		// Elements that seem form-related
		tagName := strings.ToLower(element.TagName)
		if strings.Contains(tagName, "input") || strings.Contains(tagName, "button") || 
		   strings.Contains(tagName, "form") || strings.Contains(tagName, "field") {
			return true
		}
		for _, attr := range element.Attributes() {
			attrName := strings.ToLower(attr.Name())
			if strings.Contains(attrName, "value") || strings.Contains(attrName, "disabled") || 
			   strings.Contains(attrName, "required") || strings.Contains(attrName, "readonly") {
				return true
			}
		}
		return false
	case "layout":
		// Elements that seem layout-related
		tagName := strings.ToLower(element.TagName)
		return strings.Contains(tagName, "grid") || strings.Contains(tagName, "flex") || 
		       strings.Contains(tagName, "layout") || strings.Contains(tagName, "container") ||
		       strings.Contains(tagName, "card") || strings.Contains(tagName, "panel")
	case "navigation":
		// Elements that seem navigation-related
		tagName := strings.ToLower(element.TagName)
		return strings.Contains(tagName, "nav") || strings.Contains(tagName, "menu") || 
		       strings.Contains(tagName, "tab") || strings.Contains(tagName, "breadcrumb") ||
		       strings.Contains(tagName, "link")
	default:
		return true
	}
}

// buildCapabilitiesSummary creates a summary of element capabilities
func (s *Server) buildCapabilitiesSummary(element *ElementInfo) string {
	var capabilities []string
	
	if len(element.Attributes()) > 0 {
		capabilities = append(capabilities, fmt.Sprintf("%d attributes", len(element.Attributes())))
	}
	if len(element.Slots()) > 0 {
		capabilities = append(capabilities, fmt.Sprintf("%d slots", len(element.Slots())))
	}
	if len(element.Events()) > 0 {
		capabilities = append(capabilities, fmt.Sprintf("%d events", len(element.Events())))
	}
	if len(element.CssProperties()) > 0 {
		capabilities = append(capabilities, fmt.Sprintf("%d CSS properties", len(element.CssProperties())))
	}
	if len(element.CssParts()) > 0 {
		capabilities = append(capabilities, fmt.Sprintf("%d CSS parts", len(element.CssParts())))
	}
	if len(element.CssStates()) > 0 {
		capabilities = append(capabilities, fmt.Sprintf("%d CSS states", len(element.CssStates())))
	}
	
	return strings.Join(capabilities, ", ")
}

// getAccessibilityFeatures identifies accessibility-related features
func (s *Server) getAccessibilityFeatures(element *ElementInfo) []string {
	var features []string
	
	// Check for accessibility-related attributes
	for _, attr := range element.Attributes() {
		attrName := strings.ToLower(attr.Name())
		if strings.HasPrefix(attrName, "aria-") || attrName == "role" || attrName == "tabindex" {
			features = append(features, "ARIA support")
			break
		}
	}
	
	// Check for accessibility indicators in descriptions
	for _, attr := range element.Attributes() {
		desc := strings.ToLower(attr.Description())
		if strings.Contains(desc, "accessibility") || strings.Contains(desc, "screen reader") || 
		   strings.Contains(desc, "keyboard") || strings.Contains(desc, "focus") {
			features = append(features, "Accessibility documented")
			break
		}
	}
	
	// Check element description for accessibility mentions
	desc := strings.ToLower(element.Description)
	if strings.Contains(desc, "accessibility") || strings.Contains(desc, "accessible") || 
	   strings.Contains(desc, "screen reader") || strings.Contains(desc, "keyboard") {
		features = append(features, "Accessibility focused")
	}
	
	return features
}