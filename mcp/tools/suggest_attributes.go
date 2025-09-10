package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"bennypowers.dev/cem/mcp/types"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// SuggestAttributesArgs represents the arguments for the suggest_attributes tool
type SuggestAttributesArgs struct {
	TagName string `json:"tagName"`
	Context string `json:"context,omitempty"`
	Partial string `json:"partial,omitempty"`
}

// handleSuggestAttributes provides intelligent attribute suggestions for custom elements
func handleSuggestAttributes(ctx context.Context, req *mcp.CallToolRequest, registry types.Registry) (*mcp.CallToolResult, error) {
	// Parse args from request
	var suggestArgs SuggestAttributesArgs
	if req.Params.Arguments != nil {
		if argsData, err := json.Marshal(req.Params.Arguments); err != nil {
			return nil, fmt.Errorf("failed to marshal args: %w", err)
		} else if err := json.Unmarshal(argsData, &suggestArgs); err != nil {
			return nil, fmt.Errorf("failed to unmarshal args: %w", err)
		}
	}

	// Get element information
	element, err := registry.ElementInfo(suggestArgs.TagName)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: fmt.Sprintf("Element '%s' not found in registry", suggestArgs.TagName),
				},
			},
		}, nil
	}

	var response strings.Builder
	response.WriteString(fmt.Sprintf("# Attribute Suggestions for `%s`\n\n", suggestArgs.TagName))

	if element.Description() != "" {
		response.WriteString(fmt.Sprintf("**Element Description:** %s\n\n", element.Description()))
	}

	attributes := element.Attributes()
	if len(attributes) == 0 {
		response.WriteString("❌ No attributes defined for this element in the manifest.\n")
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: response.String(),
				},
			},
		}, nil
	}

	// Filter attributes based on partial match if provided
	var filteredAttributes []types.Attribute
	if suggestArgs.Partial != "" {
		partialLower := strings.ToLower(suggestArgs.Partial)
		for _, attr := range attributes {
			if strings.Contains(strings.ToLower(attr.Name()), partialLower) {
				filteredAttributes = append(filteredAttributes, attr)
			}
		}
	} else {
		filteredAttributes = attributes
	}

	if len(filteredAttributes) == 0 {
		response.WriteString(fmt.Sprintf("❌ No attributes found matching '%s'\n", suggestArgs.Partial))
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: response.String(),
				},
			},
		}, nil
	}

	// Categorize attributes
	requiredAttrs := []types.Attribute{}
	optionalAttrs := []types.Attribute{}
	accessibilityAttrs := []types.Attribute{}

	for _, attr := range filteredAttributes {
		if attr.Required() {
			requiredAttrs = append(requiredAttrs, attr)
		} else {
			optionalAttrs = append(optionalAttrs, attr)
		}

		// Check if it's accessibility-related
		attrName := strings.ToLower(attr.Name())
		if strings.HasPrefix(attrName, "aria-") || attrName == "role" ||
			strings.Contains(attrName, "label") || strings.Contains(attrName, "accessible") {
			accessibilityAttrs = append(accessibilityAttrs, attr)
		}
	}

	// Show required attributes first
	if len(requiredAttrs) > 0 {
		response.WriteString("## ⚠️ Required Attributes\n\n")
		for _, attr := range requiredAttrs {
			response.WriteString(formatAttributeDetails(attr, suggestArgs.Context))
			response.WriteString("\n")
		}
	}

	// Show optional attributes
	if len(optionalAttrs) > 0 {
		response.WriteString("## 💡 Optional Attributes\n\n")
		for _, attr := range optionalAttrs {
			response.WriteString(formatAttributeDetails(attr, suggestArgs.Context))
			response.WriteString("\n")
		}
	}

	// Show accessibility attributes if context is accessibility
	if suggestArgs.Context == "accessibility" && len(accessibilityAttrs) > 0 {
		response.WriteString("## ♿ Accessibility Attributes\n\n")
		for _, attr := range accessibilityAttrs {
			response.WriteString(formatAttributeDetails(attr, suggestArgs.Context))
			response.WriteString("\n")
		}
	}

	// Add context-specific suggestions
	response.WriteString(generateContextualSuggestions(element, suggestArgs.Context))

	// Add usage examples
	response.WriteString(generateAttributeExamples(element, suggestArgs.Context))

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: response.String(),
			},
		},
	}, nil
}

// formatAttributeDetails creates a detailed description of an attribute
func formatAttributeDetails(attr types.Attribute, context string) string {
	var details strings.Builder

	// Attribute name and type
	details.WriteString(fmt.Sprintf("### `%s`", attr.Name()))
	if attr.Type() != "" {
		details.WriteString(fmt.Sprintf(" _%s_", attr.Type()))
	}
	details.WriteString("\n\n")

	// Description
	if attr.Description() != "" {
		details.WriteString(fmt.Sprintf("%s\n\n", attr.Description()))
	}

	// Type information
	if attr.Type() != "" {
		details.WriteString(fmt.Sprintf("**Type:** `%s`\n", attr.Type()))
	}

	// Default value
	if attr.Default() != "" {
		details.WriteString(fmt.Sprintf("**Default:** `%s`\n", attr.Default()))
	}

	// Required status
	if attr.Required() {
		details.WriteString("**Required:** Yes ⚠️\n")
	} else {
		details.WriteString("**Required:** No\n")
	}

	// Valid values (enum/union types)
	if len(attr.Values()) > 0 {
		details.WriteString("**Valid Values:**\n")
		for _, value := range attr.Values() {
			details.WriteString(fmt.Sprintf("- `%s`\n", value))
		}
	}

	// Usage examples from attribute
	if len(attr.Examples()) > 0 {
		details.WriteString("**Examples:**\n")
		for _, example := range attr.Examples() {
			details.WriteString(fmt.Sprintf("- `%s`\n", example))
		}
	}

	// Guidelines
	if len(attr.Guidelines()) > 0 {
		details.WriteString("**Guidelines:**\n")
		for _, guideline := range attr.Guidelines() {
			details.WriteString(fmt.Sprintf("- %s\n", guideline))
		}
	}

	return details.String()
}

// generateContextualSuggestions provides context-specific attribute recommendations
func generateContextualSuggestions(element types.ElementInfo, context string) string {
	var suggestions strings.Builder

	switch strings.ToLower(context) {
	case "accessibility":
		suggestions.WriteString("## ♿ Accessibility Recommendations\n\n")
		suggestions.WriteString("Consider these accessibility attributes:\n")
		suggestions.WriteString("- `role`: Define the element's purpose for screen readers\n")
		suggestions.WriteString("- `aria-label`: Provide accessible name if not obvious\n")
		suggestions.WriteString("- `aria-describedby`: Reference detailed description\n")
		suggestions.WriteString("- `aria-expanded`: For collapsible elements\n")
		suggestions.WriteString("- `aria-hidden`: Hide decorative elements from screen readers\n")
		suggestions.WriteString("- `tabindex`: Control keyboard navigation\n\n")

	case "form":
		suggestions.WriteString("## 📝 Form Context Recommendations\n\n")
		suggestions.WriteString("For form usage, consider:\n")
		suggestions.WriteString("- `name`: For form submission\n")
		suggestions.WriteString("- `required`: Mark required fields\n")
		suggestions.WriteString("- `disabled`: Disable when not available\n")
		suggestions.WriteString("- `aria-invalid`: Indicate validation state\n")
		suggestions.WriteString("- `aria-describedby`: Link to error messages\n\n")

	case "interactive":
		suggestions.WriteString("## 🖱️ Interactive Element Recommendations\n\n")
		suggestions.WriteString("For interactive elements, consider:\n")
		suggestions.WriteString("- `disabled`: Control interaction state\n")
		suggestions.WriteString("- `aria-pressed`: For toggle buttons\n")
		suggestions.WriteString("- `aria-expanded`: For expandable controls\n")
		suggestions.WriteString("- `aria-controls`: Reference controlled element\n\n")

	case "styling":
		if len(element.CssProperties()) > 0 {
			suggestions.WriteString("## 🎨 CSS Customization\n\n")
			suggestions.WriteString("This element supports CSS custom properties:\n")
			for _, prop := range element.CssProperties() {
				suggestions.WriteString(fmt.Sprintf("- `%s`: %s\n", prop.Name(), prop.Description()))
			}
			suggestions.WriteString("\n")
		}
	}

	return suggestions.String()
}

// generateAttributeExamples creates comprehensive usage examples
func generateAttributeExamples(element types.ElementInfo, context string) string {
	var examples strings.Builder
	examples.WriteString("## 📋 Usage Examples\n\n")

	tagName := element.TagName()

	// Basic example
	basicExample := fmt.Sprintf("<%s", tagName)

	// Add required attributes
	for _, attr := range element.Attributes() {
		if attr.Required() {
			if attr.Default() != "" {
				basicExample += fmt.Sprintf(` %s="%s"`, attr.Name(), attr.Default())
			} else if len(attr.Values()) > 0 {
				basicExample += fmt.Sprintf(` %s="%s"`, attr.Name(), attr.Values()[0])
			} else {
				basicExample += fmt.Sprintf(` %s="value"`, attr.Name())
			}
		}
	}
	basicExample += "></" + tagName + ">"

	examples.WriteString("### Basic Usage\n")
	examples.WriteString(fmt.Sprintf("```html\n%s\n```\n\n", basicExample))

	// Context-specific examples
	switch strings.ToLower(context) {
	case "accessibility":
		accessibleExample := fmt.Sprintf(`<%s role="button" aria-label="Accessible %s" tabindex="0">`, tagName, tagName)
		accessibleExample += "Content</" + tagName + ">"
		examples.WriteString("### Accessible Implementation\n")
		examples.WriteString(fmt.Sprintf("```html\n%s\n```\n\n", accessibleExample))

	case "form":
		formExample := fmt.Sprintf(`<form>
  <label for="my-%s">Label:</label>
  <%s id="my-%s" name="field-name" required>
    Content
  </%s>
</form>`, tagName, tagName, tagName, tagName)
		examples.WriteString("### Form Integration\n")
		examples.WriteString(fmt.Sprintf("```html\n%s\n```\n\n", formExample))
	}

	// Example with all optional attributes
	if len(element.Attributes()) > 1 {
		fullExample := fmt.Sprintf("<%s", tagName)
		for _, attr := range element.Attributes() {
			if attr.Default() != "" {
				fullExample += fmt.Sprintf(` %s="%s"`, attr.Name(), attr.Default())
			} else if len(attr.Values()) > 0 {
				fullExample += fmt.Sprintf(` %s="%s"`, attr.Name(), attr.Values()[0])
			}
		}
		fullExample += "></" + tagName + ">"

		examples.WriteString("### Full Configuration\n")
		examples.WriteString(fmt.Sprintf("```html\n%s\n```\n\n", fullExample))
	}

	return examples.String()
}
