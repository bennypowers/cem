package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// RegistryResource represents a summary of the element registry
type RegistryResource struct {
	TotalElements int                           `json:"totalElements"`
	Elements      map[string]ElementSummary     `json:"elements"`
	Capabilities  RegistryCapabilities          `json:"capabilities"`
	Packages      map[string]PackageInfo        `json:"packages,omitempty"`
}

// ElementSummary provides a summary of an element's capabilities
type ElementSummary struct {
	TagName       string   `json:"tagName"`
	Description   string   `json:"description"`
	Attributes    int      `json:"attributes"`
	Slots         int      `json:"slots"`
	Events        int      `json:"events"`
	CssProperties int      `json:"cssProperties"`
	CssParts      int      `json:"cssParts"`
	CssStates     int      `json:"cssStates"`
	Categories    []string `json:"categories"`
	Accessibility []string `json:"accessibility"`
}

// RegistryCapabilities provides aggregate statistics
type RegistryCapabilities struct {
	TotalAttributes    int `json:"totalAttributes"`
	TotalSlots         int `json:"totalSlots"`
	TotalEvents        int `json:"totalEvents"`
	TotalCssProperties int `json:"totalCssProperties"`
	TotalCssParts      int `json:"totalCssParts"`
	TotalCssStates     int `json:"totalCssStates"`
}

// PackageInfo provides information about packages containing elements
type PackageInfo struct {
	Name        string   `json:"name"`
	Version     string   `json:"version,omitempty"`
	Description string   `json:"description,omitempty"`
	Elements    []string `json:"elements"`
}

// handleRegistryResource provides a comprehensive overview of the element registry
func (s *SimpleCEMServer) handleRegistryResource(ctx context.Context, uri string) (*mcp.Resource, error) {
	elements := s.registry.GetAllElements()
	
	registryData := &RegistryResource{
		TotalElements: len(elements),
		Elements:      make(map[string]ElementSummary),
		Capabilities:  RegistryCapabilities{},
	}

	// Build element summaries and aggregate statistics
	for tagName, element := range elements {
		summary := ElementSummary{
			TagName:       element.TagName,
			Description:   element.Description,
			Attributes:    len(element.Attributes()),
			Slots:         len(element.Slots()),
			Events:        len(element.Events()),
			CssProperties: len(element.CssProperties()),
			CssParts:      len(element.CssParts()),
			CssStates:     len(element.CssStates()),
			Categories:    s.categorizeElement(element),
			Accessibility: s.getAccessibilityFeatures(element),
		}
		
		registryData.Elements[tagName] = summary
		
		// Aggregate statistics
		registryData.Capabilities.TotalAttributes += summary.Attributes
		registryData.Capabilities.TotalSlots += summary.Slots
		registryData.Capabilities.TotalEvents += summary.Events
		registryData.Capabilities.TotalCssProperties += summary.CssProperties
		registryData.Capabilities.TotalCssParts += summary.CssParts
		registryData.Capabilities.TotalCssStates += summary.CssStates
	}

	contents, err := json.MarshalIndent(registryData, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal registry data: %w", err)
	}

	return &mcp.Resource{
		URI:      uri,
		Name:     "Custom Elements Registry",
		MimeType: "application/json",
		Text:     string(contents),
	}, nil
}

// categorizeElement determines categories for an element based on its features
func (s *SimpleCEMServer) categorizeElement(element *ElementInfo) []string {
	var categories []string
	tagName := element.TagName
	
	// Basic categorization by name patterns
	switch {
	case containsAny(tagName, []string{"button", "btn"}):
		categories = append(categories, "interactive", "form")
	case containsAny(tagName, []string{"input", "field", "form"}):
		categories = append(categories, "form", "interactive")
	case containsAny(tagName, []string{"card", "panel", "container"}):
		categories = append(categories, "layout", "content")
	case containsAny(tagName, []string{"nav", "menu", "tab", "breadcrumb"}):
		categories = append(categories, "navigation", "interactive")
	case containsAny(tagName, []string{"grid", "flex", "layout"}):
		categories = append(categories, "layout")
	case containsAny(tagName, []string{"dialog", "modal", "popup", "tooltip"}):
		categories = append(categories, "overlay", "interactive")
	case containsAny(tagName, []string{"icon", "image", "avatar"}):
		categories = append(categories, "media", "display")
	case containsAny(tagName, []string{"text", "heading", "title", "label"}):
		categories = append(categories, "typography", "content")
	}
	
	// Categorization by features
	if len(element.Events()) > 0 {
		categories = appendUnique(categories, "interactive")
	}
	
	if len(element.Slots()) > 0 {
		categories = appendUnique(categories, "content")
	}
	
	if len(element.CssProperties()) > 0 {
		categories = appendUnique(categories, "themeable")
	}
	
	if len(element.CssParts()) > 0 {
		categories = appendUnique(categories, "styleable")
	}
	
	// Check for form-related attributes
	for _, attr := range element.Attributes() {
		attrName := attr.Name()
		if containsAny(attrName, []string{"value", "disabled", "required", "readonly"}) {
			categories = appendUnique(categories, "form")
			break
		}
	}
	
	// Default category if none found
	if len(categories) == 0 {
		categories = append(categories, "component")
	}
	
	return categories
}

// containsAny checks if a string contains any of the given substrings
func containsAny(s string, substrings []string) bool {
	for _, substr := range substrings {
		if contains(s, substr) {
			return true
		}
	}
	return false
}

// contains checks if string contains substring (case-insensitive)
func contains(s, substr string) bool {
	return len(s) >= len(substr) && 
		   (s == substr || 
		    (len(s) > len(substr) && findSubstring(s, substr)))
}

// findSubstring performs case-insensitive substring search
func findSubstring(s, substr string) bool {
	if len(substr) == 0 {
		return true
	}
	if len(s) < len(substr) {
		return false
	}
	
	// Simple case-insensitive search
	for i := 0; i <= len(s)-len(substr); i++ {
		match := true
		for j := 0; j < len(substr); j++ {
			c1, c2 := s[i+j], substr[j]
			// Convert to lowercase
			if c1 >= 'A' && c1 <= 'Z' {
				c1 = c1 + 32
			}
			if c2 >= 'A' && c2 <= 'Z' {
				c2 = c2 + 32
			}
			if c1 != c2 {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

// appendUnique appends a string to a slice if it's not already present
func appendUnique(slice []string, item string) []string {
	for _, existing := range slice {
		if existing == item {
			return slice
		}
	}
	return append(slice, item)
}