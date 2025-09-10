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
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var titleCase = cases.Title(language.English)

// GenerateHtmlArgs represents the arguments for the generate_html tool
type GenerateHtmlArgs struct {
	TagName    string            `json:"tagName"`
	Context    string            `json:"context,omitempty"`
	Options    map[string]string `json:"options,omitempty"`
	Content    string            `json:"content,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

// handleGenerateHtml generates HTML structure for custom elements
func handleGenerateHtml(
	ctx context.Context,
	req *mcp.CallToolRequest,
	registry types.Registry,
) (*mcp.CallToolResult, error) {
	// Parse args from request
	var genArgs GenerateHtmlArgs
	if req.Params.Arguments != nil {
		if argsData, err := json.Marshal(req.Params.Arguments); err != nil {
			return nil, fmt.Errorf("failed to marshal args: %w", err)
		} else if err := json.Unmarshal(argsData, &genArgs); err != nil {
			return nil, fmt.Errorf("failed to unmarshal args: %w", err)
		}
	}

	// Get element information
	element, err := registry.ElementInfo(genArgs.TagName)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: fmt.Sprintf("Element '%s' not found in registry", genArgs.TagName),
				},
			},
		}, nil
	}

	var response strings.Builder
	response.WriteString(fmt.Sprintf("# HTML Generation for `%s`\n\n", genArgs.TagName))

	if element.Description() != "" {
		response.WriteString(fmt.Sprintf("**Element Description:** %s\n\n", element.Description()))
	}

	// Generate HTML structure
	generatedHTML := generateHTML(element, genArgs)

	response.WriteString("## Generated HTML\n\n")
	response.WriteString("```html\n")
	response.WriteString(generatedHTML)
	response.WriteString("\n```\n\n")

	// Add usage notes
	response.WriteString(generateUsageNotes(element, genArgs))

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: response.String(),
			},
		},
	}, nil
}

// generateHTML creates HTML structure for the element
func generateHTML(element types.ElementInfo, args GenerateHtmlArgs) string {
	var html strings.Builder

	tagName := element.TagName()

	// Start opening tag
	html.WriteString("<" + tagName)

	// Add required attributes with provided or default values
	for _, attr := range element.Attributes() {
		if attr.Required() {
			if value, exists := args.Attributes[attr.Name()]; exists {
				html.WriteString(fmt.Sprintf(` %s="%s"`, attr.Name(), value))
			} else if attr.Default() != "" {
				html.WriteString(fmt.Sprintf(` %s="%s"`, attr.Name(), attr.Default()))
			} else if len(attr.Values()) > 0 {
				html.WriteString(fmt.Sprintf(` %s="%s"`, attr.Name(), attr.Values()[0]))
			}
		}
	}

	// Add user-provided attributes (excluding required ones already added)
	for name, value := range args.Attributes {
		isRequired := false
		for _, attr := range element.Attributes() {
			if attr.Required() && attr.Name() == name {
				isRequired = true
				break
			}
		}
		if !isRequired {
			html.WriteString(fmt.Sprintf(` %s="%s"`, name, value))
		}
	}

	html.WriteString(">")

	// Add content with slot structure
	content := generateElementContent(element, args)
	html.WriteString(content)

	// Close tag
	html.WriteString("</" + tagName + ">")

	return html.String()
}

// generateElementContent creates content based on element slots
func generateElementContent(element types.ElementInfo, args GenerateHtmlArgs) string {
	if args.Content != "" {
		return args.Content
	}

	var content strings.Builder

	// Generate content based on available slots
	slots := element.Slots()
	if len(slots) > 0 {
		for _, slot := range slots {
			slotName := slot.Name()
			if slotName == "" {
				// Default slot
				content.WriteString("Default content")
			} else {
				content.WriteString(fmt.Sprintf(`<span slot="%s">%s content</span>`, slotName, titleCase.String(slotName)))
			}
			content.WriteString("\n  ")
		}
	} else {
		// No slots defined, add basic content
		content.WriteString("Element content")
	}

	return content.String()
}

// generateUsageNotes provides basic usage information
func generateUsageNotes(
	element types.ElementInfo,
	args GenerateHtmlArgs,
) string {
	var notes strings.Builder
	notes.WriteString("## Usage Notes\n\n")

	// Required attributes
	requiredAttrs := []string{}
	for _, attr := range element.Attributes() {
		if attr.Required() {
			requiredAttrs = append(requiredAttrs, attr.Name())
		}
	}
	if len(requiredAttrs) > 0 {
		notes.WriteString("### Required Attributes\n")
		for _, attrName := range requiredAttrs {
			notes.WriteString(fmt.Sprintf("- `%s`\n", attrName))
		}
		notes.WriteString("\n")
	}

	// Available slots
	slots := element.Slots()
	if len(slots) > 0 {
		notes.WriteString("### Available Slots\n")
		for _, slot := range slots {
			slotName := slot.Name()
			if slotName == "" {
				notes.WriteString("- Default slot\n")
			} else {
				notes.WriteString(fmt.Sprintf("- `%s`", slotName))
				if slot.Description() != "" {
					notes.WriteString(fmt.Sprintf(": %s", slot.Description()))
				}
				notes.WriteString("\n")
			}
		}
		notes.WriteString("\n")
	}

	// Accessibility considerations
	notes.WriteString("### Accessibility Considerations\n")
	notes.WriteString("- This element may have built-in accessibility features via ElementInternals\n")
	notes.WriteString("- Check the element's shadow DOM for implemented accessibility patterns\n")
	notes.WriteString("- Avoid adding redundant ARIA attributes that may conflict with built-in semantics\n")
	notes.WriteString("- Test with screen readers to verify the element's accessibility implementation\n\n")

	return notes.String()
}
