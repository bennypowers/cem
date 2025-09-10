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

	"bennypowers.dev/cem/mcp/types"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// handleRegistryResource provides the complete registry state with all elements
func handleRegistryResource(ctx context.Context, req *mcp.ReadResourceRequest, registry types.Registry) (*mcp.ReadResourceResult, error) {
	// Get all elements from registry
	elementMap := registry.AllElements()

	// Convert map to slice
	elements := make([]types.ElementInfo, 0, len(elementMap))
	for _, element := range elementMap {
		elements = append(elements, element)
	}

	// Convert to a format suitable for AI consumption
	registryData := map[string]interface{}{
		"elements": convertElementsForAI(elements),
		"metadata": map[string]interface{}{
			"totalElements": len(elements),
			"packages":      getPackageList(elements),
		},
	}

	contents, err := json.MarshalIndent(registryData, "", "  ")
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

// convertElementsForAI converts element info to AI-friendly format
func convertElementsForAI(elements []types.ElementInfo) []map[string]interface{} {
	result := make([]map[string]interface{}, len(elements))

	for i, element := range elements {
		result[i] = map[string]interface{}{
			"tagName":       element.TagName(),
			"description":   element.Description(),
			"attributes":    convertAttributesForAI(element.Attributes()),
			"slots":         convertSlotsForAI(element.Slots()),
			"events":        convertEventsForAI(element.Events()),
			"cssProperties": convertCSSPropertiesForAI(element.CssProperties()),
			"cssParts":      convertCSSPartsForAI(element.CssParts()),
			"cssStates":     convertCSSStatesForAI(element.CssStates()),
			"examples":      convertExamplesForAI(element.Examples()),
		}
	}

	return result
}

// convertAttributesForAI converts attributes to AI-friendly format
func convertAttributesForAI(attributes []types.Attribute) []map[string]interface{} {
	result := make([]map[string]interface{}, len(attributes))

	for i, attr := range attributes {
		result[i] = map[string]interface{}{
			"name":        attr.Name(),
			"description": attr.Description(),
			"type":        attr.Type(),
			"default":     attr.Default(),
			"required":    attr.Required(),
		}
	}

	return result
}

// convertSlotsForAI converts slots to AI-friendly format
func convertSlotsForAI(slots []types.Slot) []map[string]interface{} {
	result := make([]map[string]interface{}, len(slots))

	for i, slot := range slots {
		result[i] = map[string]interface{}{
			"name":        slot.Name(),
			"description": slot.Description(),
		}
	}

	return result
}

// convertEventsForAI converts events to AI-friendly format
func convertEventsForAI(events []types.Event) []map[string]interface{} {
	result := make([]map[string]interface{}, len(events))

	for i, event := range events {
		result[i] = map[string]interface{}{
			"name":        event.Name(),
			"description": event.Description(),
			"type":        event.Type(),
		}
	}

	return result
}

// convertCSSPropertiesForAI converts CSS properties to AI-friendly format
func convertCSSPropertiesForAI(props []types.CssProperty) []map[string]interface{} {
	result := make([]map[string]interface{}, len(props))

	for i, prop := range props {
		result[i] = map[string]interface{}{
			"name":        prop.Name(),
			"description": prop.Description(),
			"syntax":      prop.Syntax(),
			"initial":     prop.Initial(),
			"inherits":    prop.Inherits(),
		}
	}

	return result
}

// convertCSSPartsForAI converts CSS parts to AI-friendly format
func convertCSSPartsForAI(parts []types.CssPart) []map[string]interface{} {
	result := make([]map[string]interface{}, len(parts))

	for i, part := range parts {
		result[i] = map[string]interface{}{
			"name":        part.Name(),
			"description": part.Description(),
		}
	}

	return result
}

// convertCSSStatesForAI converts CSS states to AI-friendly format
func convertCSSStatesForAI(states []types.CssState) []map[string]interface{} {
	result := make([]map[string]interface{}, len(states))

	for i, state := range states {
		result[i] = map[string]interface{}{
			"name":        state.Name(),
			"description": state.Description(),
		}
	}

	return result
}

// convertExamplesForAI converts examples to AI-friendly format
func convertExamplesForAI(examples []types.Example) []map[string]interface{} {
	result := make([]map[string]interface{}, len(examples))

	for i, example := range examples {
		result[i] = map[string]interface{}{
			"title":       example.Title(),
			"description": example.Description(),
			"code":        example.Code(),
		}
	}

	return result
}

// getPackageList extracts unique package names from elements
func getPackageList(elements []types.ElementInfo) []string {
	packageSet := make(map[string]bool)

	// For now, we'll extract package info from tag names or descriptions
	// This could be enhanced based on actual package metadata in the registry
	for _, element := range elements {
		// Simple heuristic: extract potential package from tag name prefix
		tagName := element.TagName()
		if parts := splitTagName(tagName); len(parts) > 1 {
			packageSet[parts[0]] = true
		}
	}

	packages := make([]string, 0, len(packageSet))
	for pkg := range packageSet {
		packages = append(packages, pkg)
	}

	return packages
}

// splitTagName splits a tag name on hyphens
func splitTagName(tagName string) []string {
	result := []string{}
	current := ""

	for _, char := range tagName {
		if char == '-' {
			if current != "" {
				result = append(result, current)
				current = ""
			}
		} else {
			current += string(char)
		}
	}

	if current != "" {
		result = append(result, current)
	}

	return result
}
