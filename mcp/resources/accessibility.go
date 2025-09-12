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
	"fmt"
	"strings"

	"bennypowers.dev/cem/mcp/types"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)


// handleAccessibilityResource provides comprehensive accessibility patterns and validation rules
func handleAccessibilityResource(ctx context.Context, req *mcp.ReadResourceRequest, registry types.Registry) (*mcp.ReadResourceResult, error) {
	// Generate comprehensive accessibility guidance
	accessibility := generateAccessibilityGuidance(registry)

	contents, err := json.MarshalIndent(accessibility, "", "  ")
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

// generateAccessibilityGuidance creates manifest-specific accessibility guidance
func generateAccessibilityGuidance(registry types.Registry) map[string]interface{} {
	elementMap := registry.AllElements()

	// Convert map to slice
	elements := make([]types.ElementInfo, 0, len(elementMap))
	for _, element := range elementMap {
		elements = append(elements, element)
	}

	// Analyze elements for accessibility information in their manifest data
	accessibilityAnalysis := analyzeManifestAccessibility(elements)

	guidance := map[string]interface{}{
		"manifestAnalysis": accessibilityAnalysis,
		"customElements":   generateCustomElementAccessibilityGuidance(elements),
		"fallbackGuidance": generateFallbackAccessibilityGuidance(),
	}

	return guidance
}


// generateCustomElementAccessibilityGuidance creates manifest-based accessibility guidance
func generateCustomElementAccessibilityGuidance(elements []types.ElementInfo) map[string]interface{} {
	elementGuidance := make(map[string]interface{})

	for _, element := range elements {
		elementGuidance[element.TagName()] = generateElementAccessibilityGuidance(element)
	}

	return elementGuidance
}





func generateElementAccessibilityGuidance(element types.ElementInfo) map[string]interface{} {
	guidance := map[string]interface{}{
		"tagName":     element.TagName(),
		"description": element.Description(),
	}

	// Extract accessibility-related information from manifest data
	accessibilityInfo := extractManifestAccessibilityInfo(element)
	if len(accessibilityInfo) > 0 {
		guidance["manifestAccessibilityInfo"] = accessibilityInfo
	}

	// Analyze attributes for potential accessibility concerns
	if len(element.Attributes()) > 0 {
		attributeConcerns := analyzeAttributesForAccessibility(element.Attributes())
		if len(attributeConcerns) > 0 {
			guidance["attributeConcerns"] = attributeConcerns
		}
	}

	// Analyze events for accessibility patterns
	if len(element.Events()) > 0 {
		eventConcerns := analyzeEventsForAccessibility(element.Events())
		if len(eventConcerns) > 0 {
			guidance["eventConcerns"] = eventConcerns
		}
	}

	// Analyze slots for accessibility implications
	if len(element.Slots()) > 0 {
		slotConcerns := analyzeSlotsForAccessibility(element.Slots())
		if len(slotConcerns) > 0 {
			guidance["slotConcerns"] = slotConcerns
		}
	}

	return guidance
}


// analyzeManifestAccessibility analyzes all elements for accessibility information
func analyzeManifestAccessibility(elements []types.ElementInfo) map[string]interface{} {
	totalElements := len(elements)
	elementsWithA11yInfo := 0
	elementsWithA11yAttributes := 0
	elementsWithA11yEvents := 0

	for _, element := range elements {
		if hasAccessibilityInfo(element) {
			elementsWithA11yInfo++
		}
		if hasAccessibilityAttributes(element) {
			elementsWithA11yAttributes++
		}
		if hasAccessibilityEvents(element) {
			elementsWithA11yEvents++
		}
	}

	return map[string]interface{}{
		"summary": map[string]interface{}{
			"totalElements":               totalElements,
			"elementsWithA11yInfo":        elementsWithA11yInfo,
			"elementsWithA11yAttributes":  elementsWithA11yAttributes,
			"elementsWithA11yEvents":      elementsWithA11yEvents,
			"coveragePercentage":          float64(elementsWithA11yInfo) / float64(totalElements) * 100,
		},
	}
}

// extractManifestAccessibilityInfo extracts accessibility-related info from element description and attributes
func extractManifestAccessibilityInfo(element types.ElementInfo) []string {
	var info []string

	// Check description for accessibility keywords
	desc := strings.ToLower(element.Description())
	accessibilityKeywords := []string{"aria", "accessible", "screen reader", "keyboard", "focus", "role", "label"}
	for _, keyword := range accessibilityKeywords {
		if strings.Contains(desc, keyword) {
			info = append(info, fmt.Sprintf("Description mentions %s", keyword))
		}
	}

	return info
}

// analyzeAttributesForAccessibility analyzes attributes for accessibility patterns
func analyzeAttributesForAccessibility(attributes []types.Attribute) []string {
	var concerns []string

	if len(attributes) > 0 {
		concerns = append(concerns, "hasAttributes")
	}

	return concerns
}

// analyzeEventsForAccessibility analyzes events for accessibility patterns
func analyzeEventsForAccessibility(events []types.Event) []string {
	var concerns []string

	if len(events) > 0 {
		concerns = append(concerns, "hasEvents")
	}

	return concerns
}

// analyzeSlotsForAccessibility analyzes slots for accessibility implications
func analyzeSlotsForAccessibility(slots []types.Slot) []string {
	var concerns []string

	if len(slots) > 0 {
		concerns = append(concerns, "hasSlots")
	}

	return concerns
}

// hasAccessibilityInfo checks if element has accessibility-related information
func hasAccessibilityInfo(element types.ElementInfo) bool {
	desc := strings.ToLower(element.Description())
	accessibilityKeywords := []string{"aria", "accessible", "screen reader", "keyboard", "focus", "role", "label"}
	for _, keyword := range accessibilityKeywords {
		if strings.Contains(desc, keyword) {
			return true
		}
	}
	return false
}

// hasAccessibilityAttributes checks if element has accessibility-related attributes
func hasAccessibilityAttributes(element types.ElementInfo) bool {
	// This would need actual implementation based on attribute interface
	return len(element.Attributes()) > 0
}

// hasAccessibilityEvents checks if element has accessibility-related events
func hasAccessibilityEvents(element types.ElementInfo) bool {
	// This would need actual implementation based on event interface
	return len(element.Events()) > 0
}

// generateFallbackAccessibilityGuidance provides guidance when no manifest accessibility info exists
func generateFallbackAccessibilityGuidance() map[string]interface{} {
	return map[string]interface{}{
		"requiresManualAudit": true,
	}
}
