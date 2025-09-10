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

// QueryRegistryArgs represents the arguments for the query_registry tool
type QueryRegistryArgs struct {
	TagName string `json:"tagName,omitempty"`
	Filter  string `json:"filter,omitempty"`
	Search  string `json:"search,omitempty"`
}

// handleQueryRegistry queries and explores the custom elements registry
func handleQueryRegistry(ctx context.Context, req *mcp.CallToolRequest, registry types.Registry) (*mcp.CallToolResult, error) {
	// Parse args from request
	var queryArgs QueryRegistryArgs
	if req.Params.Arguments != nil {
		if argsData, err := json.Marshal(req.Params.Arguments); err != nil {
			return nil, fmt.Errorf("failed to marshal args: %w", err)
		} else if err := json.Unmarshal(argsData, &queryArgs); err != nil {
			return nil, fmt.Errorf("failed to unmarshal args: %w", err)
		}
	}

	elements := registry.AllElements()

	// Handle specific element query
	if queryArgs.TagName != "" {
		if element, exists := elements[queryArgs.TagName]; exists {
			templateData := NewAttributeTemplateData(element, "")
			elementDetails, err := renderTemplate("element_details", templateData)
			if err != nil {
				return nil, fmt.Errorf("failed to render element details: %w", err)
			}

			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{
						Text: elementDetails,
					},
				},
			}, nil
		} else {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					&mcp.TextContent{
						Text: fmt.Sprintf("Element '%s' not found in registry", queryArgs.TagName),
					},
				},
			}, nil
		}
	}

	// Filter elements based on criteria
	var filteredElements map[string]types.ElementInfo
	if queryArgs.Filter != "" || queryArgs.Search != "" {
		filteredElements = make(map[string]types.ElementInfo)
		for tagName, element := range elements {
			if matchesFilter(element, queryArgs.Filter, queryArgs.Search) {
				filteredElements[tagName] = element
			}
		}
	} else {
		filteredElements = elements
	}

	// Build comprehensive registry information
	registryData := RegistryQueryData{
		Elements:   filteredElements,
		Categories: categorizeElementsWithCapabilities(filteredElements),
		Filter:     queryArgs.Filter,
		Search:     queryArgs.Search,
	}

	result, err := renderRegistryTemplate("registry_query_results", registryData)
	if err != nil {
		return nil, fmt.Errorf("failed to render registry query results: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: result,
			},
		},
	}, nil
}

// matchesFilter checks if an element matches the given filter and search criteria
func matchesFilter(element types.ElementInfo, filter, search string) bool {
	// Search criteria (case insensitive text search)
	if search != "" {
		searchLower := strings.ToLower(search)
		if !strings.Contains(strings.ToLower(element.TagName()), searchLower) &&
			!strings.Contains(strings.ToLower(element.Description()), searchLower) &&
			!strings.Contains(strings.ToLower(element.Name()), searchLower) {
			return false
		}
	}

	// Filter criteria
	if filter != "" {
		switch strings.ToLower(filter) {
		case "has-slots":
			if len(element.Slots()) == 0 {
				return false
			}
		case "has-events":
			if len(element.Events()) == 0 {
				return false
			}
		case "has-css-properties":
			if len(element.CssProperties()) == 0 {
				return false
			}
		case "has-css-parts":
			if len(element.CssParts()) == 0 {
				return false
			}
		case "interactive":
			hasEvents := len(element.Events()) > 0
			hasInteractiveAttrs := false
			for _, attr := range element.Attributes() {
				attrName := strings.ToLower(attr.Name())
				if strings.Contains(attrName, "click") || strings.Contains(attrName, "disabled") {
					hasInteractiveAttrs = true
					break
				}
			}
			if !hasEvents && !hasInteractiveAttrs {
				return false
			}
		case "form":
			tagName := strings.ToLower(element.TagName())
			isFormRelated := strings.Contains(tagName, "input") || strings.Contains(tagName, "button") ||
				strings.Contains(tagName, "field") || strings.Contains(tagName, "form")
			if !isFormRelated {
				return false
			}
		}
	}

	return true
}

// categorizeElementsWithCapabilities groups elements by their primary capabilities and adds capability info
func categorizeElementsWithCapabilities(elements map[string]types.ElementInfo) map[string][]ElementWithCapabilities {
	categories := map[string][]ElementWithCapabilities{
		"Form Elements":        {},
		"Interactive Elements": {},
		"Layout Elements":      {},
		"Content Elements":     {},
		"Media Elements":       {},
		"Navigation Elements":  {},
		"Other Elements":       {},
	}

	for _, element := range elements {
		tagName := strings.ToLower(element.TagName())
		category := "Other Elements"

		// Categorize based on tag name patterns
		switch {
		case strings.Contains(tagName, "input") || strings.Contains(tagName, "button") ||
			strings.Contains(tagName, "field") || strings.Contains(tagName, "form"):
			category = "Form Elements"
		case len(element.Events()) > 0 || strings.Contains(tagName, "button"):
			category = "Interactive Elements"
		case strings.Contains(tagName, "grid") || strings.Contains(tagName, "layout") ||
			strings.Contains(tagName, "container") || strings.Contains(tagName, "flex"):
			category = "Layout Elements"
		case strings.Contains(tagName, "text") || strings.Contains(tagName, "content") ||
			strings.Contains(tagName, "paragraph") || strings.Contains(tagName, "heading"):
			category = "Content Elements"
		case strings.Contains(tagName, "image") || strings.Contains(tagName, "video") ||
			strings.Contains(tagName, "audio") || strings.Contains(tagName, "media"):
			category = "Media Elements"
		case strings.Contains(tagName, "nav") || strings.Contains(tagName, "menu") ||
			strings.Contains(tagName, "breadcrumb") || strings.Contains(tagName, "tab"):
			category = "Navigation Elements"
		}

		// Build capability list
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

		elementWithCaps := ElementWithCapabilities{
			ElementInfo:  element,
			Capabilities: capabilities,
		}

		categories[category] = append(categories[category], elementWithCaps)
	}

	return categories
}
