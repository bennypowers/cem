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
package resources

import (
	"context"
	"encoding/json"
	"sort"

	"bennypowers.dev/cem/mcp/types"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// ElementSummary represents a summary of element capabilities for discovery
type ElementSummary struct {
	TagName      string   `json:"tagName"`
	Name         string   `json:"name,omitempty"`
	Description  string   `json:"description,omitempty"`
	Package      string   `json:"package,omitempty"`
	Module       string   `json:"module,omitempty"`
	Capabilities []string `json:"capabilities"`

	// Capability counts for quick assessment
	AttributeCount   int `json:"attributeCount"`
	SlotCount        int `json:"slotCount"`
	EventCount       int `json:"eventCount"`
	CssPropertyCount int `json:"cssPropertyCount"`
	CssPartCount     int `json:"cssPartCount"`
	CssStateCount    int `json:"cssStateCount"`
}

// handleElementsResource provides all element declarations for discovery
func handleElementsResource(ctx context.Context, req *mcp.ReadResourceRequest, registry types.MCPContext) (*mcp.ReadResourceResult, error) {
	// Get all elements from registry
	elementMap := registry.AllElements()

	// Convert to element summaries for discovery
	elements := make([]ElementSummary, 0, len(elementMap))

	// Sort element keys for deterministic JSON output
	elementKeys := make([]string, 0, len(elementMap))
	for tagName := range elementMap {
		elementKeys = append(elementKeys, tagName)
	}
	sort.Strings(elementKeys)

	for _, tagName := range elementKeys {
		element := elementMap[tagName]
		summary := ElementSummary{
			TagName:          element.TagName(),
			Name:             element.Name(),
			Description:      element.Description(),
			Package:          extractPackageFromElement(element),
			Module:           element.Module(),
			AttributeCount:   len(element.Attributes()),
			SlotCount:        len(element.Slots()),
			EventCount:       len(element.Events()),
			CssPropertyCount: len(element.CssProperties()),
			CssPartCount:     len(element.CssParts()),
			CssStateCount:    len(element.CssStates()),
		}

		// Build capability summary
		summary.Capabilities = buildCapabilitySummary(element)

		elements = append(elements, summary)
	}

	// Build response following CEM elements structure
	elementsData := map[string]interface{}{
		"elements": elements,
		"metadata": map[string]interface{}{
			"totalElements": len(elements),
			"categories":    categorizeElementsByCapabilities(elements),
		},
	}

	contents, err := json.MarshalIndent(elementsData, "", "  ")
	if err != nil {
		return nil, err
	}

	return &mcp.ReadResourceResult{
		Contents: []*mcp.ResourceContents{{
			URI:      req.Params.URI,
			MIMEType: "application/json",
			Text:     string(contents),
		}},
	}, nil
}

// buildCapabilitySummary creates a list of capabilities for an element
func buildCapabilitySummary(element types.ElementInfo) []string {
	var capabilities []string

	if len(element.Attributes()) > 0 {
		capabilities = append(capabilities, "configurable")
	}

	if len(element.Slots()) > 0 {
		capabilities = append(capabilities, "content-slots")
	}

	if len(element.Events()) > 0 {
		capabilities = append(capabilities, "interactive")
	}

	if len(element.CssProperties()) > 0 {
		capabilities = append(capabilities, "themeable")
	}

	if len(element.CssParts()) > 0 {
		capabilities = append(capabilities, "styleable")
	}

	if len(element.CssStates()) > 0 {
		capabilities = append(capabilities, "stateful")
	}

	// Add semantic capabilities based on tag name patterns
	tagName := element.TagName()
	if containsAny(tagName, []string{"button", "input", "field", "form"}) {
		capabilities = append(capabilities, "form-element")
	}

	if containsAny(tagName, []string{"grid", "layout", "container", "flex"}) {
		capabilities = append(capabilities, "layout-element")
	}

	if containsAny(tagName, []string{"nav", "menu", "tab", "breadcrumb"}) {
		capabilities = append(capabilities, "navigation-element")
	}

	return capabilities
}

// categorizeElementsByCapabilities groups elements by their primary capabilities
func categorizeElementsByCapabilities(elements []ElementSummary) map[string]int {
	categories := map[string]int{
		"configurable":        0,
		"content-slots":       0,
		"interactive":         0,
		"themeable":           0,
		"styleable":           0,
		"stateful":            0,
		"form-elements":       0,
		"layout-elements":     0,
		"navigation-elements": 0,
	}

	for _, element := range elements {
		for _, capability := range element.Capabilities {
			if capability == "form-element" {
				categories["form-elements"]++
			} else if capability == "layout-element" {
				categories["layout-elements"]++
			} else if capability == "navigation-element" {
				categories["navigation-elements"]++
			} else if count, exists := categories[capability]; exists {
				categories[capability] = count + 1
			}
		}
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

// contains checks if a string contains a substring (case-insensitive)
func contains(s, substr string) bool {
	return len(s) >= len(substr) && indexOf(s, substr) >= 0
}

// indexOf returns the index of substr in s, or -1 if not found
func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
