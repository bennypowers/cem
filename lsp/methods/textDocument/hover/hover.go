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
	"github.com/bennypowers/glsp"
	protocol "github.com/bennypowers/glsp/protocol_3_17"
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
		helpers.SafeDebugLog("[HOVER] Failed to get DocumentManager: %v", err)
		return nil, err
	}

	// Find the element at the cursor position
	element := doc.FindElementAtPosition(params.Position, dm)
	if element != nil {
		helpers.SafeDebugLog("[HOVER] Found element at position: tagName=%s, range=%+v\n", element.TagName, element.Range)

		// Look up the full declaration to get summary and description
		if decl := ctx.FindCustomElementDeclaration(element.TagName); decl != nil {
			helpers.SafeDebugLog("[HOVER] Found custom element declaration in registry: %s\n", element.TagName)

			// Create hover content with full element information including summary/description
			content := CreateElementHoverContentFromDeclaration(decl)
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
		helpers.SafeDebugLog("[HOVER] Found attribute at position: name=%s, tagName=%s, bindingPrefix=%s, range=%+v\n", attribute.Name, tagName, attribute.BindingPrefix, attribute.Range)

		switch attribute.BindingPrefix {
		case "@":
			if events, exists := ctx.Events(tagName); exists {
				if event, exists := events[attribute.Name]; exists {
					content := CreateEventHoverContent(event, tagName)
					return &protocol.Hover{
						Contents: protocol.MarkupContent{
							Kind:  protocol.MarkupKindMarkdown,
							Value: content,
						},
						Range: &attribute.Range,
					}, nil
				}
			}
		case ".":
			if fields, exists := ctx.Fields(tagName); exists {
				if field, exists := fields[attribute.Name]; exists {
					content := CreateFieldHoverContent(field, tagName)
					return &protocol.Hover{
						Contents: protocol.MarkupContent{
							Kind:  protocol.MarkupKindMarkdown,
							Value: content,
						},
						Range: &attribute.Range,
					}, nil
				}
			}
		default:
			if attrs, exists := ctx.Attributes(tagName); exists {
				if attr, exists := attrs[attribute.Name]; exists {
					content := CreateAttributeHoverContent(attr, tagName)
					return &protocol.Hover{
						Contents: protocol.MarkupContent{
							Kind:  protocol.MarkupKindMarkdown,
							Value: content,
						},
						Range: &attribute.Range,
					}, nil
				}
			}
		}
	}

	helpers.SafeDebugLog("[HOVER] No hover content found\n")
	return nil, nil
}

// CreateElementHoverContent creates markdown content for custom element hover
func CreateElementHoverContent(element *M.CustomElement) string {
	var content strings.Builder

	// 1. Tag name first (as title)
	fmt.Fprintf(&content, "## `<%s>`\n\n", element.TagName)

	// Note: For now we'll show basic info since the manifest structure doesn't
	// include summary/description at the CustomElement level. These would be
	// available on the ClassDeclaration that defines the custom element.

	// Show that it's a custom element
	content.WriteString("**Custom Element**\n\n")

	// Attributes, Events, Slots using shared formatters
	content.WriteString(formatAttributes(element.Attributes))
	content.WriteString(formatEvents(element.Events))
	content.WriteString(formatSlots(element.Slots))

	return content.String()
}

// CreateElementHoverContentFromDeclaration creates markdown content for custom element hover
// using the full CustomElementDeclaration which includes summary and description
func CreateElementHoverContentFromDeclaration(decl *M.CustomElementDeclaration) string {
	if decl == nil {
		return ""
	}

	var content strings.Builder

	// 1. Tag name first (as title)
	fmt.Fprintf(&content, "## `<%s>`\n\n", decl.TagName)

	// 2. Summary (if available)
	if decl.Summary != "" {
		fmt.Fprintf(&content, "**%s**\n\n", decl.Summary)
	}

	// 3. Description (if available and different from summary)
	if decl.Description != "" && decl.Description != decl.Summary {
		fmt.Fprintf(&content, "%s\n\n", decl.Description)
	}

	// 4. Attributes, Events, Slots using shared formatters
	content.WriteString(formatAttributes(decl.Attributes()))
	content.WriteString(formatEvents(decl.Events()))
	content.WriteString(formatSlots(decl.Slots()))

	return content.String()
}

// formatAttributes creates markdown content for an attributes list
func formatAttributes(attrs []M.Attribute) string {
	if len(attrs) == 0 {
		return ""
	}

	var content strings.Builder
	content.WriteString("### Attributes\n\n")
	for _, attr := range attrs {
		fmt.Fprintf(&content, "- **`%s`**", attr.Name)
		if attr.Type != nil && attr.Type.Text != "" {
			fmt.Fprintf(&content, " _%s_", attr.Type.Text)
		}
		if attr.InheritedFrom != nil {
			fmt.Fprintf(&content, " _(inherited from %s)_", attr.InheritedFrom.Name)
		}
		if attr.Description != "" {
			fmt.Fprintf(&content, " - %s", attr.Description)
		}
		content.WriteString("\n")
	}
	content.WriteString("\n")
	return content.String()
}

// formatEvents creates markdown content for an events list
func formatEvents(events []M.Event) string {
	if len(events) == 0 {
		return ""
	}

	var content strings.Builder
	content.WriteString("### Events\n\n")
	for _, event := range events {
		fmt.Fprintf(&content, "- **`%s`**", event.Name)
		if event.Type != nil && event.Type.Text != "" {
			fmt.Fprintf(&content, " _%s_", event.Type.Text)
		}
		if event.InheritedFrom != nil {
			fmt.Fprintf(&content, " _(inherited from %s)_", event.InheritedFrom.Name)
		}
		if event.Description != "" {
			fmt.Fprintf(&content, " - %s", event.Description)
		}
		content.WriteString("\n")
	}
	content.WriteString("\n")
	return content.String()
}

// formatSlots creates markdown content for a slots list
func formatSlots(slots []M.Slot) string {
	if len(slots) == 0 {
		return ""
	}

	var content strings.Builder
	content.WriteString("### Slots\n\n")
	for _, slot := range slots {
		fmt.Fprintf(&content, "- **`%s`**", slot.Name)
		if slot.InheritedFrom != nil {
			fmt.Fprintf(&content, " _(inherited from %s)_", slot.InheritedFrom.Name)
		}
		if slot.Description != "" {
			fmt.Fprintf(&content, " - %s", slot.Description)
		}
		content.WriteString("\n")
	}
	content.WriteString("\n")
	return content.String()
}

// CreateAttributeHoverContent creates markdown content for attribute hover
func CreateAttributeHoverContent(attr *M.Attribute, tagName string) string {
	var content strings.Builder

	// Title
	fmt.Fprintf(&content, "## `%s` attribute\n\n", attr.Name)
	fmt.Fprintf(&content, "**On `<%s>` element**\n\n", tagName)

	// Inheritance info
	if attr.InheritedFrom != nil {
		fmt.Fprintf(&content, "_Inherited from %s_\n\n", attr.InheritedFrom.Name)
	}

	// Type
	if attr.Type != nil && attr.Type.Text != "" {
		fmt.Fprintf(&content, "**Type:** `%s`\n\n", attr.Type.Text)
	}

	// Description
	if attr.Description != "" {
		fmt.Fprintf(&content, "%s\n\n", attr.Description)
	}

	// Default value
	if attr.Default != "" {
		fmt.Fprintf(&content, "**Default:** `%s`\n\n", attr.Default)
	}

	// Deprecated warning
	if attr.IsDeprecated() {
		switch attr.Deprecated.(type) {
		case M.DeprecatedFlag:
			content.WriteString("⚠️ **Deprecated**\n\n")
		case M.DeprecatedReason:
			fmt.Fprintf(&content, "⚠️ **Deprecated**: %s\n\n", attr.Deprecated)
		}
	}

	return content.String()
}

// CreateFieldHoverContent creates markdown content for class field hover
func CreateFieldHoverContent(field *M.ClassField, tagName string) string {
	var content strings.Builder

	fmt.Fprintf(&content, "## `%s` property\n\n", field.Name)
	fmt.Fprintf(&content, "**On `<%s>` element**\n\n", tagName)

	if field.InheritedFrom != nil {
		fmt.Fprintf(&content, "_Inherited from %s_\n\n", field.InheritedFrom.Name)
	}

	if field.Type != nil && field.Type.Text != "" {
		fmt.Fprintf(&content, "**Type:** `%s`\n\n", field.Type.Text)
	}

	if field.Description != "" {
		fmt.Fprintf(&content, "%s\n\n", field.Description)
	}

	if field.Default != "" {
		fmt.Fprintf(&content, "**Default:** `%s`\n\n", field.Default)
	}

	if field.IsDeprecated() {
		switch field.Deprecated.(type) {
		case M.DeprecatedFlag:
			content.WriteString("⚠️ **Deprecated**\n\n")
		case M.DeprecatedReason:
			fmt.Fprintf(&content, "⚠️ **Deprecated**: %s\n\n", field.Deprecated)
		}
	}

	return content.String()
}

// CreateEventHoverContent creates markdown content for event hover
func CreateEventHoverContent(event *M.Event, tagName string) string {
	var content strings.Builder

	// Title
	fmt.Fprintf(&content, "## `%s` event\n\n", event.Name)
	fmt.Fprintf(&content, "**On `<%s>` element**\n\n", tagName)

	// Inheritance info
	if event.InheritedFrom != nil {
		fmt.Fprintf(&content, "_Inherited from %s_\n\n", event.InheritedFrom.Name)
	}

	// Type (event detail type)
	if event.Type != nil && event.Type.Text != "" {
		fmt.Fprintf(&content, "**Type:** `%s`\n\n", event.Type.Text)
	}

	// Description
	if event.Description != "" {
		fmt.Fprintf(&content, "%s\n\n", event.Description)
	}

	// Deprecated warning
	if event.IsDeprecated() {
		switch event.Deprecated.(type) {
		case M.DeprecatedFlag:
			content.WriteString("⚠️ **Deprecated**\n\n")
		case M.DeprecatedReason:
			fmt.Fprintf(&content, "⚠️ **Deprecated**: %s\n\n", event.Deprecated)
		}
	}

	return content.String()
}
