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
	"fmt"
	"strings"

	"bennypowers.dev/cem/mcp/templates"
	"bennypowers.dev/cem/mcp/types"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// handleElementResource provides detailed information about a specific custom element
func handleElementResource(
	ctx context.Context,
	req *mcp.ReadResourceRequest,
	registry types.MCPContext,
) (*mcp.ReadResourceResult, error) {
	// Extract tag name from URI template: cem://element/{tagName}
	tagName, err := extractTagNameFromURI(req.Params.URI)
	if err != nil {
		return nil, fmt.Errorf("invalid element URI: %w", err)
	}

	// Find the element in the registry
	element, err := registry.ElementInfo(tagName)
	if err != nil {
		return nil, fmt.Errorf("element not found: %s: %w", tagName, err)
	}

	// Generate template-driven element documentation following Data + Context + LLM philosophy
	elementDoc, err := generateTemplateElementDocumentation(element)
	if err != nil {
		return nil, fmt.Errorf("failed to generate element documentation: %w", err)
	}

	return &mcp.ReadResourceResult{
		Contents: []*mcp.ResourceContents{{
			URI:      req.Params.URI,
			MIMEType: "text/markdown",
			Text:     elementDoc,
		}},
	}, nil
}

// extractTagNameFromURI extracts the tag name from a URI like cem://element/{tagName}
func extractTagNameFromURI(uri string) (string, error) {
	const prefix = "cem://element/"
	if !strings.HasPrefix(uri, prefix) {
		return "", fmt.Errorf("URI does not match element template: %s", uri)
	}

	tagName := strings.TrimPrefix(uri, prefix)
	if tagName == "" || tagName == "{tagName}" {
		return "", fmt.Errorf("no tag name specified in URI: %s", uri)
	}

	return tagName, nil
}

// generateTemplateElementDocumentation creates template-driven documentation following Data + Context + LLM philosophy
func generateTemplateElementDocumentation(element types.ElementInfo) (string, error) {
	// Prepare data for template rendering
	data := ElementTemplateData{
		TagName:       element.TagName(),
		Description:   element.Description(),
		Attributes:    convertAttributesForTemplate(element.Attributes()),
		Slots:         convertSlotsForTemplate(element.Slots()),
		Events:        convertEventsForTemplate(element.Events()),
		CssProperties: convertCSSPropertiesForTemplate(element.CssProperties()),
		CssParts:      convertCSSPartsForTemplate(element.CssParts()),
		CssStates:     convertCSSStatesForTemplate(element.CssStates()),
	}

	// Render the element examples template with rich context for LLM
	return templates.RenderTemplate("element_examples", data)
}

// ElementTemplateData represents the data structure for element templates
type ElementTemplateData struct {
	TagName       string
	Description   string
	Attributes    []AttributeTemplateData
	Slots         []SlotTemplateData
	Events        []EventTemplateData
	CssProperties []CSSPropertyTemplateData
	CssParts      []CSSPartTemplateData
	CssStates     []CSSStateTemplateData
}

// AttributeTemplateData represents attribute data for templates
type AttributeTemplateData struct {
	Name        string
	Description string
	Type        string
	Default     string
	Required    bool
	Values      []string
	Guidelines  []string
}

// SlotTemplateData represents slot data for templates
type SlotTemplateData struct {
	Name        string
	Description string
	Guidelines  []string
}

// EventTemplateData represents event data for templates
type EventTemplateData struct {
	Name        string
	Description string
	Type        string
	Guidelines  []string
}

// CSSPropertyTemplateData represents CSS property data for templates
type CSSPropertyTemplateData struct {
	Name        string
	Description string
	Syntax      string
	Initial     string
	Inherits    bool
	Guidelines  []string
}

// CSSPartTemplateData represents CSS part data for templates
type CSSPartTemplateData struct {
	Name        string
	Description string
	Guidelines  []string
}

// CSSStateTemplateData represents CSS state data for templates
type CSSStateTemplateData struct {
	Name        string
	Description string
	Guidelines  []string
}

// Conversion functions to prepare data for templates
func convertAttributesForTemplate(attrs []types.Attribute) []AttributeTemplateData {
	result := make([]AttributeTemplateData, len(attrs))
	for i, attr := range attrs {
		result[i] = AttributeTemplateData{
			Name:        attr.Name(),
			Description: attr.Description(),
			Type:        attr.Type(),
			Default:     attr.Default(),
			Required:    attr.Required(),
			Values:      attr.Values(),
			Guidelines:  attr.Guidelines(),
		}
	}
	return result
}

func convertSlotsForTemplate(slots []types.Slot) []SlotTemplateData {
	result := make([]SlotTemplateData, len(slots))
	for i, slot := range slots {
		result[i] = SlotTemplateData{
			Name:        slot.Name(),
			Description: slot.Description(),
			Guidelines:  slot.Guidelines(),
		}
	}
	return result
}

func convertEventsForTemplate(events []types.Event) []EventTemplateData {
	result := make([]EventTemplateData, len(events))
	for i, event := range events {
		result[i] = EventTemplateData{
			Name:        event.Name(),
			Description: event.Description(),
			Type:        event.Type(),
			Guidelines:  event.Guidelines(),
		}
	}
	return result
}

func convertCSSPropertiesForTemplate(props []types.CssProperty) []CSSPropertyTemplateData {
	result := make([]CSSPropertyTemplateData, len(props))
	for i, prop := range props {
		result[i] = CSSPropertyTemplateData{
			Name:        prop.Name(),
			Description: prop.Description(),
			Syntax:      prop.Syntax(),
			Initial:     prop.Initial(),
			Inherits:    prop.Inherits(),
			Guidelines:  prop.Guidelines(),
		}
	}
	return result
}

func convertCSSPartsForTemplate(parts []types.CssPart) []CSSPartTemplateData {
	result := make([]CSSPartTemplateData, len(parts))
	for i, part := range parts {
		result[i] = CSSPartTemplateData{
			Name:        part.Name(),
			Description: part.Description(),
			Guidelines:  part.Guidelines(),
		}
	}
	return result
}

func convertCSSStatesForTemplate(states []types.CssState) []CSSStateTemplateData {
	result := make([]CSSStateTemplateData, len(states))
	for i, state := range states {
		result[i] = CSSStateTemplateData{
			Name:        state.Name(),
			Description: state.Description(),
			Guidelines:  state.Guidelines(),
		}
	}
	return result
}
