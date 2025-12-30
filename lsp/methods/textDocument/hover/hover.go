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
package hover

import (
	"fmt"
	"strings"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// Hover handles textDocument/hover requests
func Hover(ctx types.ServerContext, context *glsp.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	uri := params.TextDocument.URI
	helpers.SafeDebugLog("[HOVER] Request for URI: %s, Position: line=%d, char=%d", uri, params.Position.Line, params.Position.Character)

	// Get the tracked document
	doc := ctx.Document(uri)
	if doc == nil {
		helpers.SafeDebugLog("[HOVER] No document found for URI: %s", uri)
		return nil, nil
	}
	helpers.SafeDebugLog("[HOVER] Found document for URI: %s", uri)

	dm, err := ctx.DocumentManager()
	if err != nil {
		helpers.SafeDebugLog("[COMPLETION] Failed to get DocumentManager: %v", err)
		return nil, err
	}

	// Find the element at the cursor position
	element := doc.FindElementAtPosition(params.Position, dm)
	if element != nil {
		helpers.SafeDebugLog("[HOVER] Found element at position: tagName=%s, range=%+v\n", element.TagName, element.Range)

		// Look up the element in our registry
		if customElement, exists := ctx.Element(element.TagName); exists {
			helpers.SafeDebugLog("[HOVER] Found custom element in registry: %s\n", element.TagName)

			// Create hover content with element information
			content := CreateElementHoverContent(customElement)
			result := &protocol.Hover{
				Contents: protocol.MarkupContent{
					Kind:  protocol.MarkupKindMarkdown,
					Value: content,
				},
				Range: &element.Range,
			}
			helpers.SafeDebugLog("[HOVER] Returning element hover content (length=%d)\n", len(content))
			return result, nil
		} else {
			helpers.SafeDebugLog("[HOVER] Element %s not found in registry\n", element.TagName)
		}
	} else {
		helpers.SafeDebugLog("[HOVER] No element found at position\n")
	}

	// Check if cursor is over an attribute

	attribute, tagName := doc.FindAttributeAtPosition(params.Position, dm)
	if attribute != nil && tagName != "" {
		helpers.SafeDebugLog("[HOVER] Found attribute at position: name=%s, tagName=%s, range=%+v\n", attribute.Name, tagName, attribute.Range)

		if attrs, exists := ctx.Attributes(tagName); exists {
			helpers.SafeDebugLog("[HOVER] Found %d attributes for element %s\n", len(attrs), tagName)

			if attr, exists := attrs[attribute.Name]; exists {
				helpers.SafeDebugLog("[HOVER] Found attribute %s in registry for element %s\n", attribute.Name, tagName)

				content := CreateAttributeHoverContent(attr, tagName)
				result := &protocol.Hover{
					Contents: protocol.MarkupContent{
						Kind:  protocol.MarkupKindMarkdown,
						Value: content,
					},
					Range: &attribute.Range,
				}
				helpers.SafeDebugLog("[HOVER] Returning attribute hover content (length=%d)\n", len(content))
				return result, nil
			} else {
				helpers.SafeDebugLog("[HOVER] Attribute %s not found in registry for element %s\n", attribute.Name, tagName)
			}
		} else {
			helpers.SafeDebugLog("[HOVER] No attributes found in registry for element %s\n", tagName)
		}
	} else {
		if attribute == nil {
			helpers.SafeDebugLog("[HOVER] No attribute found at position\n")
		} else {
			helpers.SafeDebugLog("[HOVER] Found attribute but no tagName: %+v\n", attribute)
		}
	}

	helpers.SafeDebugLog("[HOVER] No hover content found\n")
	return nil, nil
}

// CreateElementHoverContent creates markdown content for custom element hover
func CreateElementHoverContent(element *M.CustomElement) string {
	var content strings.Builder

	// 1. Tag name first (as title)
	content.WriteString(fmt.Sprintf("## `<%s>`\n\n", element.TagName))

	// Note: For now we'll show basic info since the manifest structure doesn't
	// include summary/description at the CustomElement level. These would be
	// available on the ClassDeclaration that defines the custom element.

	// Show that it's a custom element
	content.WriteString("**Custom Element**\n\n")

	// Attributes
	if len(element.Attributes) > 0 {
		content.WriteString("### Attributes\n\n")
		for _, attr := range element.Attributes {
			content.WriteString(fmt.Sprintf("- **`%s`**", attr.Name))
			if attr.Type != nil && attr.Type.Text != "" {
				content.WriteString(fmt.Sprintf(" _%s_", attr.Type.Text))
			}
			if attr.Description != "" {
				content.WriteString(fmt.Sprintf(" - %s", attr.Description))
			}
			content.WriteString("\n")
		}
		content.WriteString("\n")
	}

	// Events
	if len(element.Events) > 0 {
		content.WriteString("### Events\n\n")
		for _, event := range element.Events {
			content.WriteString(fmt.Sprintf("- **`%s`**", event.Name))
			if event.Type != nil && event.Type.Text != "" {
				content.WriteString(fmt.Sprintf(" _%s_", event.Type.Text))
			}
			if event.Description != "" {
				content.WriteString(fmt.Sprintf(" - %s", event.Description))
			}
			content.WriteString("\n")
		}
		content.WriteString("\n")
	}

	// Slots
	if len(element.Slots) > 0 {
		content.WriteString("### Slots\n\n")
		for _, slot := range element.Slots {
			content.WriteString(fmt.Sprintf("- **`%s`**", slot.Name))
			if slot.Description != "" {
				content.WriteString(fmt.Sprintf(" - %s", slot.Description))
			}
			content.WriteString("\n")
		}
		content.WriteString("\n")
	}

	return content.String()
}

// CreateAttributeHoverContent creates markdown content for attribute hover
func CreateAttributeHoverContent(attr *M.Attribute, tagName string) string {
	var content strings.Builder

	// Title
	content.WriteString(fmt.Sprintf("## `%s` attribute\n\n", attr.Name))
	content.WriteString(fmt.Sprintf("**On `<%s>` element**\n\n", tagName))

	// Type
	if attr.Type != nil && attr.Type.Text != "" {
		content.WriteString(fmt.Sprintf("**Type:** `%s`\n\n", attr.Type.Text))
	}

	// Description
	if attr.Description != "" {
		content.WriteString(fmt.Sprintf("%s\n\n", attr.Description))
	}

	// Default value
	if attr.Default != "" {
		content.WriteString(fmt.Sprintf("**Default:** `%s`\n\n", attr.Default))
	}

	// Deprecated warning
	if attr.IsDeprecated() {
		switch attr.Deprecated.(type) {
		case M.DeprecatedFlag:
			content.WriteString("⚠️ **Deprecated**\n\n")
		case M.DeprecatedReason:
			content.WriteString(fmt.Sprintf("⚠️ **Deprecated**: %s\n\n", attr.Deprecated))
		}
	}

	return content.String()
}

// CreateEventHoverContent creates markdown content for event hover
func CreateEventHoverContent(event *M.Event, tagName string) string {
	var content strings.Builder

	// Title
	content.WriteString(fmt.Sprintf("## `%s` event\n\n", event.Name))
	content.WriteString(fmt.Sprintf("**On `<%s>` element**\n\n", tagName))

	// Type (event detail type)
	if event.Type != nil && event.Type.Text != "" {
		content.WriteString(fmt.Sprintf("**Type:** `%s`\n\n", event.Type.Text))
	}

	// Description
	if event.Description != "" {
		content.WriteString(fmt.Sprintf("%s\n\n", event.Description))
	}

	// Deprecated warning
	if event.IsDeprecated() {
		switch event.Deprecated.(type) {
		case M.DeprecatedFlag:
			content.WriteString("⚠️ **Deprecated**\n\n")
		case M.DeprecatedReason:
			content.WriteString(fmt.Sprintf("⚠️ **Deprecated**: %s\n\n", event.Deprecated))
		}
	}

	return content.String()
}
