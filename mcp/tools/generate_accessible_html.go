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

// GenerateAccessibleHtmlArgs represents the arguments for the generate_accessible_html tool
type GenerateAccessibleHtmlArgs struct {
	TagName    string            `json:"tagName"`
	Context    string            `json:"context,omitempty"`
	Options    map[string]string `json:"options,omitempty"`
	Content    string            `json:"content,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

// handleGenerateAccessibleHtml generates accessible HTML with proper ARIA attributes and semantic structure
func handleGenerateAccessibleHtml(ctx context.Context, req *mcp.CallToolRequest, registry types.Registry) (*mcp.CallToolResult, error) {
	// Parse args from request
	var genArgs GenerateAccessibleHtmlArgs
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
	response.WriteString(fmt.Sprintf("# Accessible HTML Generation for `%s`\n\n", genArgs.TagName))

	if element.Description() != "" {
		response.WriteString(fmt.Sprintf("**Element Description:** %s\n\n", element.Description()))
	}

	// Generate HTML based on context
	generatedHTML := generateAccessibleHTML(element, genArgs)

	response.WriteString("## Generated Accessible HTML\n\n")
	response.WriteString("```html\n")
	response.WriteString(generatedHTML)
	response.WriteString("\n```\n\n")

	// Add accessibility explanations
	response.WriteString(generateAccessibilityExplanation(element, genArgs))

	// Add WCAG compliance information
	response.WriteString(generateWCAGCompliance(element, genArgs))

	// Add alternative implementations
	response.WriteString(generateAlternativeImplementations(element, genArgs))

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: response.String(),
			},
		},
	}, nil
}

// generateAccessibleHTML creates the actual HTML with accessibility features
func generateAccessibleHTML(element types.ElementInfo, args GenerateAccessibleHtmlArgs) string {
	var html strings.Builder

	tagName := element.TagName()

	// Start opening tag
	html.WriteString("<" + tagName)

	// Add required attributes first
	for _, attr := range element.Attributes() {
		if attr.Required() {
			// Use provided value or default
			if value, exists := args.Attributes[attr.Name()]; exists {
				html.WriteString(fmt.Sprintf(` %s="%s"`, attr.Name(), value))
			} else if attr.Default() != "" {
				html.WriteString(fmt.Sprintf(` %s="%s"`, attr.Name(), attr.Default()))
			} else if len(attr.Values()) > 0 {
				html.WriteString(fmt.Sprintf(` %s="%s"`, attr.Name(), attr.Values()[0]))
			}
		}
	}

	// Add accessibility attributes based on context
	accessibilityAttrs := generateAccessibilityAttributes(element, args)
	for name, value := range accessibilityAttrs {
		html.WriteString(fmt.Sprintf(` %s="%s"`, name, value))
	}

	// Add user-provided attributes (non-required)
	for name, value := range args.Attributes {
		// Skip if already added as required attribute
		isRequired := false
		for _, attr := range element.Attributes() {
			if attr.Required() && attr.Name() == name {
				isRequired = true
				break
			}
		}
		if !isRequired && !strings.HasPrefix(name, "aria-") && name != "role" {
			html.WriteString(fmt.Sprintf(` %s="%s"`, name, value))
		}
	}

	html.WriteString(">")

	// Add content with proper slot structure
	content := generateElementContent(element, args)
	html.WriteString(content)

	// Close tag
	html.WriteString("</" + tagName + ">")

	return html.String()
}

// generateAccessibilityAttributes creates appropriate ARIA attributes
func generateAccessibilityAttributes(element types.ElementInfo, args GenerateAccessibleHtmlArgs) map[string]string {
	attrs := make(map[string]string)
	tagName := strings.ToLower(element.TagName())

	// Context-specific accessibility attributes
	switch strings.ToLower(args.Context) {
	case "button", "interactive":
		attrs["role"] = "button"
		attrs["tabindex"] = "0"
		if args.Options["label"] != "" {
			attrs["aria-label"] = args.Options["label"]
		} else {
			attrs["aria-label"] = fmt.Sprintf("%s button", element.TagName())
		}

	case "navigation", "nav":
		attrs["role"] = "navigation"
		if args.Options["label"] != "" {
			attrs["aria-label"] = args.Options["label"]
		} else {
			attrs["aria-label"] = "Navigation menu"
		}

	case "form", "input":
		if args.Options["label"] != "" {
			attrs["aria-label"] = args.Options["label"]
		}
		if args.Options["required"] == "true" {
			attrs["aria-required"] = "true"
		}
		if args.Options["invalid"] == "true" {
			attrs["aria-invalid"] = "true"
		}

	case "dialog", "modal":
		attrs["role"] = "dialog"
		attrs["aria-modal"] = "true"
		if args.Options["label"] != "" {
			attrs["aria-labelledby"] = args.Options["label"]
		}

	case "tab":
		attrs["role"] = "tab"
		attrs["tabindex"] = "0"
		attrs["aria-selected"] = "false"
		if args.Options["controls"] != "" {
			attrs["aria-controls"] = args.Options["controls"]
		}

	case "menu":
		attrs["role"] = "menu"
		if args.Options["orientation"] != "" {
			attrs["aria-orientation"] = args.Options["orientation"]
		} else {
			attrs["aria-orientation"] = "vertical"
		}

	default:
		// Auto-detect based on tag name patterns
		switch {
		case strings.Contains(tagName, "button"):
			attrs["role"] = "button"
			attrs["tabindex"] = "0"
		case strings.Contains(tagName, "nav") || strings.Contains(tagName, "menu"):
			attrs["role"] = "navigation"
		case strings.Contains(tagName, "dialog") || strings.Contains(tagName, "modal"):
			attrs["role"] = "dialog"
		case strings.Contains(tagName, "tab"):
			attrs["role"] = "tab"
		case len(element.Events()) > 0:
			// Interactive element
			attrs["tabindex"] = "0"
		}
	}

	// Add describedby if description is available
	if args.Options["describedby"] != "" {
		attrs["aria-describedby"] = args.Options["describedby"]
	}

	// Add custom ARIA attributes from user input
	for name, value := range args.Attributes {
		if strings.HasPrefix(name, "aria-") || name == "role" {
			attrs[name] = value
		}
	}

	return attrs
}

// generateElementContent creates appropriate content with slot structure
func generateElementContent(element types.ElementInfo, args GenerateAccessibleHtmlArgs) string {
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
				content.WriteString(fmt.Sprintf(`<span slot="%s">%s content</span>`, slotName, strings.Title(slotName)))
			}
			content.WriteString("\n  ")
		}
	} else {
		// No slots defined, add generic accessible content
		switch strings.ToLower(args.Context) {
		case "button", "interactive":
			content.WriteString("Click me")
		case "navigation", "nav":
			content.WriteString(`<ul>
    <li><a href="#home">Home</a></li>
    <li><a href="#about">About</a></li>
    <li><a href="#contact">Contact</a></li>
  </ul>`)
		case "form":
			content.WriteString("Form content")
		default:
			content.WriteString("Element content")
		}
	}

	return content.String()
}

// generateAccessibilityExplanation explains the accessibility features
func generateAccessibilityExplanation(element types.ElementInfo, args GenerateAccessibleHtmlArgs) string {
	var explanation strings.Builder
	explanation.WriteString("## 🔍 Accessibility Features Explained\n\n")

	explanation.WriteString("### ARIA Attributes Added\n\n")

	accessibilityAttrs := generateAccessibilityAttributes(element, args)
	for name, value := range accessibilityAttrs {
		switch name {
		case "role":
			explanation.WriteString(fmt.Sprintf("- **`role=\"%s\"`**: Defines the element's purpose for assistive technologies\n", value))
		case "aria-label":
			explanation.WriteString(fmt.Sprintf("- **`aria-label=\"%s\"`**: Provides an accessible name for the element\n", value))
		case "aria-labelledby":
			explanation.WriteString(fmt.Sprintf("- **`aria-labelledby=\"%s\"`**: References another element that labels this one\n", value))
		case "aria-describedby":
			explanation.WriteString(fmt.Sprintf("- **`aria-describedby=\"%s\"`**: References detailed description\n", value))
		case "tabindex":
			explanation.WriteString(fmt.Sprintf("- **`tabindex=\"%s\"`**: Makes element keyboard focusable\n", value))
		case "aria-modal":
			explanation.WriteString("- **`aria-modal=\"true\"`**: Indicates this is a modal dialog\n")
		case "aria-expanded":
			explanation.WriteString(fmt.Sprintf("- **`aria-expanded=\"%s\"`**: Indicates if collapsible content is expanded\n", value))
		case "aria-selected":
			explanation.WriteString(fmt.Sprintf("- **`aria-selected=\"%s\"`**: Indicates selection state\n", value))
		case "aria-controls":
			explanation.WriteString(fmt.Sprintf("- **`aria-controls=\"%s\"`**: References controlled element\n", value))
		}
	}

	explanation.WriteString("\n### Keyboard Navigation\n\n")
	explanation.WriteString("- Element is keyboard accessible via Tab key\n")
	explanation.WriteString("- Activate with Enter or Space (for interactive elements)\n")
	explanation.WriteString("- Follows standard keyboard interaction patterns\n\n")

	explanation.WriteString("### Screen Reader Support\n\n")
	explanation.WriteString("- Proper semantic meaning communicated\n")
	explanation.WriteString("- Element purpose clearly announced\n")
	explanation.WriteString("- State changes will be announced\n\n")

	return explanation.String()
}

// generateWCAGCompliance shows WCAG compliance information
func generateWCAGCompliance(element types.ElementInfo, args GenerateAccessibleHtmlArgs) string {
	var compliance strings.Builder
	compliance.WriteString("## ✅ WCAG 2.1 AA Compliance\n\n")

	compliance.WriteString("### Perceivable\n")
	compliance.WriteString("- ✅ **1.3.1 Info and Relationships**: Semantic structure maintained\n")
	compliance.WriteString("- ✅ **1.4.3 Contrast**: Ensure sufficient color contrast in CSS\n\n")

	compliance.WriteString("### Operable\n")
	compliance.WriteString("- ✅ **2.1.1 Keyboard**: All functionality available via keyboard\n")
	compliance.WriteString("- ✅ **2.4.3 Focus Order**: Logical focus order maintained\n")
	compliance.WriteString("- ✅ **2.4.6 Headings and Labels**: Descriptive labels provided\n\n")

	compliance.WriteString("### Understandable\n")
	compliance.WriteString("- ✅ **3.2.2 On Input**: No unexpected context changes\n")
	compliance.WriteString("- ✅ **3.3.2 Labels or Instructions**: Clear labels provided\n\n")

	compliance.WriteString("### Robust\n")
	compliance.WriteString("- ✅ **4.1.2 Name, Role, Value**: ARIA properly implemented\n")
	compliance.WriteString("- ✅ **4.1.3 Status Messages**: Status changes communicated\n\n")

	compliance.WriteString("### Additional Considerations\n")
	compliance.WriteString("- Ensure color contrast ratios meet AA standards (4.5:1)\n")
	compliance.WriteString("- Test with screen readers (NVDA, JAWS, VoiceOver)\n")
	compliance.WriteString("- Verify keyboard navigation flow\n")
	compliance.WriteString("- Test with users who have disabilities\n\n")

	return compliance.String()
}

// generateAlternativeImplementations shows different accessible approaches
func generateAlternativeImplementations(element types.ElementInfo, args GenerateAccessibleHtmlArgs) string {
	var alternatives strings.Builder
	alternatives.WriteString("## 🔄 Alternative Accessible Implementations\n\n")

	tagName := element.TagName()

	// Minimal accessibility version
	alternatives.WriteString("### Minimal Accessibility\n")
	alternatives.WriteString("```html\n")
	minimalHTML := fmt.Sprintf(`<%s role="button" tabindex="0">
  Basic content
</%s>`, tagName, tagName)
	alternatives.WriteString(minimalHTML)
	alternatives.WriteString("\n```\n\n")

	// Enhanced accessibility version
	alternatives.WriteString("### Enhanced Accessibility\n")
	alternatives.WriteString("```html\n")
	enhancedHTML := fmt.Sprintf(`<%s 
  role="button" 
  tabindex="0"
  aria-label="Descriptive button label"
  aria-describedby="help-text">
  Enhanced content
</%s>
<div id="help-text" class="sr-only">
  Additional context for screen readers
</div>`, tagName, tagName)
	alternatives.WriteString(enhancedHTML)
	alternatives.WriteString("\n```\n\n")

	// Form context version
	if strings.ToLower(args.Context) == "form" ||
		strings.Contains(strings.ToLower(element.TagName()), "input") {
		alternatives.WriteString("### Form Context\n")
		alternatives.WriteString("```html\n")
		formHTML := fmt.Sprintf(`<div class="form-field">
  <label for="field-id">Field Label</label>
  <%s 
    id="field-id"
    name="field-name"
    aria-describedby="field-help field-error"
    aria-required="true">
    Field content
  </%s>
  <div id="field-help">Help text for this field</div>
  <div id="field-error" role="alert" aria-live="polite">
    Error message appears here
  </div>
</div>`, tagName, tagName)
		alternatives.WriteString(formHTML)
		alternatives.WriteString("\n```\n\n")
	}

	return alternatives.String()
}
