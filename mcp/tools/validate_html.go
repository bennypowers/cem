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
	"regexp"
	"strings"

	"bennypowers.dev/cem/mcp/types"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// ValidateHtmlArgs represents the arguments for the validate_html tool
type ValidateHtmlArgs struct {
	Html    string `json:"html"`
	Context string `json:"context,omitempty"`
	TagName string `json:"tagName,omitempty"`
}

// handleValidateHtml validates HTML for accessibility and custom element compliance
func handleValidateHtml(ctx context.Context, req *mcp.CallToolRequest, registry types.Registry) (*mcp.CallToolResult, error) {
	// Parse args from request
	var validateArgs ValidateHtmlArgs
	if req.Params.Arguments != nil {
		if argsData, err := json.Marshal(req.Params.Arguments); err != nil {
			return nil, fmt.Errorf("failed to marshal args: %w", err)
		} else if err := json.Unmarshal(argsData, &validateArgs); err != nil {
			return nil, fmt.Errorf("failed to unmarshal args: %w", err)
		}
	}

	var results strings.Builder
	results.WriteString("# HTML Validation Results\n\n")

	// Basic HTML structure analysis
	html := validateArgs.Html

	// Check for custom elements in the HTML
	elements := registry.AllElements()
	var foundElements []string

	for tagName := range elements {
		// Look for the custom element in the HTML
		if strings.Contains(html, "<"+tagName) || strings.Contains(html, "</"+tagName) {
			foundElements = append(foundElements, tagName)
		}
	}

	if len(foundElements) > 0 {
		results.WriteString("## Custom Elements Found\n\n")
		for _, tagName := range foundElements {
			element, _ := registry.ElementInfo(tagName)
			results.WriteString(fmt.Sprintf("- `%s`: %s\n", tagName, element.Description()))
		}
		results.WriteString("\n")
	}

	// Perform validation based on context
	switch strings.ToLower(validateArgs.Context) {
	case "semantic":
		results.WriteString(validateSemanticStructure(html))
	case "manifest-compliance":
		results.WriteString(validateManifestCompliance(html, elements))
	default:
		// Focus on manifest compliance and semantic structure
		results.WriteString(validateManifestCompliance(html, elements))
		results.WriteString(validateSemanticStructure(html))
	}

	// If specific tagName is provided, focus validation on that element
	if validateArgs.TagName != "" {
		if element, err := registry.ElementInfo(validateArgs.TagName); err == nil {
			results.WriteString(validateSpecificElement(html, validateArgs.TagName, element))
		} else {
			results.WriteString(fmt.Sprintf("## Element-Specific Validation\n\n❌ Element '%s' not found in registry\n\n", validateArgs.TagName))
		}
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: results.String(),
			},
		},
	}, nil
}

// validateManifestCompliance checks custom element usage against manifest guidelines only
func validateManifestCompliance(html string, elements map[string]types.ElementInfo) string {
	var results strings.Builder
	results.WriteString("## Manifest Compliance Validation\n\n")

	var issues []string
	var notes []string

	// Focus only on custom elements and their manifest-defined requirements
	for tagName, element := range elements {
		if strings.Contains(html, "<"+tagName) {
			// Extract the element usage from HTML
			elementRegex := regexp.MustCompile(`<` + regexp.QuoteMeta(tagName) + `[^>]*>`)
			usages := elementRegex.FindAllString(html, -1)

			for _, usage := range usages {
				// Check required attributes per manifest
				for _, attr := range element.Attributes() {
					if attr.Required() && !strings.Contains(usage, attr.Name()+"=") {
						issues = append(issues, fmt.Sprintf("Required attribute '%s' missing in <%s>", attr.Name(), tagName))
					}
				}

				// Check for proper attribute values per manifest constraints
				for _, attr := range element.Attributes() {
					if strings.Contains(usage, attr.Name()+"=") && len(attr.Values()) > 0 {
						// Extract attribute value
						attrRegex := regexp.MustCompile(attr.Name() + `=["']([^"']+)["']`)
						matches := attrRegex.FindStringSubmatch(usage)
						if len(matches) > 1 {
							value := matches[1]
							validValues := attr.Values()
							isValid := false
							for _, validValue := range validValues {
								if value == validValue {
									isValid = true
									break
								}
							}
							if !isValid {
								issues = append(issues, fmt.Sprintf("Invalid value '%s' for attribute '%s' in <%s>. Valid values: %s",
									value, attr.Name(), tagName, strings.Join(validValues, ", ")))
							}
						}
					}
				}
			}

			// Note available features for guidance
			if len(element.Slots()) > 0 {
				notes = append(notes, fmt.Sprintf("<%s> supports slots: %s", tagName, getSlotNames(element.Slots())))
			}

			// Note manifest guidelines if available
			if len(element.Guidelines()) > 0 {
				notes = append(notes, fmt.Sprintf("<%s> guidelines: %s", tagName, strings.Join(element.Guidelines(), "; ")))
			}
		}
	}

	// Report results
	if len(issues) > 0 {
		results.WriteString("### ❌ Manifest Compliance Issues\n\n")
		for _, issue := range issues {
			results.WriteString(fmt.Sprintf("- %s\n", issue))
		}
		results.WriteString("\n")
	} else {
		results.WriteString("### ✅ No Manifest Compliance Issues Found\n\n")
	}

	if len(notes) > 0 {
		results.WriteString("### 💡 Manifest-Defined Features\n\n")
		for _, note := range notes {
			results.WriteString(fmt.Sprintf("- %s\n", note))
		}
		results.WriteString("\n")
	}

	return results.String()
}

// validateSemanticStructure checks for proper semantic HTML usage
func validateSemanticStructure(html string) string {
	var results strings.Builder
	results.WriteString("## Semantic Structure Validation\n\n")

	var issues []string
	var suggestions []string

	// Check for proper document structure
	hasMain := strings.Contains(html, "<main")

	if !hasMain {
		issues = append(issues, "Consider using <main> element for primary content")
	}

	// Check for div abuse
	divCount := strings.Count(html, "<div")
	sectionCount := strings.Count(html, "<section")
	articleCount := strings.Count(html, "<article")

	if divCount > 5 && (sectionCount+articleCount) == 0 {
		suggestions = append(suggestions, "Consider using semantic elements like <section> or <article> instead of multiple <div> elements")
	}

	// Check for proper list usage
	if strings.Contains(html, "•") || strings.Contains(html, "1.") || strings.Contains(html, "-") {
		if !strings.Contains(html, "<ul") && !strings.Contains(html, "<ol") {
			suggestions = append(suggestions, "Consider using proper list elements (<ul> or <ol>) for list content")
		}
	}

	// Report results
	if len(issues) > 0 {
		results.WriteString("### ❌ Semantic Issues\n\n")
		for _, issue := range issues {
			results.WriteString(fmt.Sprintf("- %s\n", issue))
		}
		results.WriteString("\n")
	}

	if len(suggestions) > 0 {
		results.WriteString("### 💡 Semantic Suggestions\n\n")
		for _, suggestion := range suggestions {
			results.WriteString(fmt.Sprintf("- %s\n", suggestion))
		}
		results.WriteString("\n")
	}

	if len(issues) == 0 && len(suggestions) == 0 {
		results.WriteString("### ✅ Good Semantic Structure\n\n")
	}

	return results.String()
}

// validateSpecificElement performs focused validation on a specific element
func validateSpecificElement(html string, tagName string, element types.ElementInfo) string {
	var results strings.Builder
	results.WriteString(fmt.Sprintf("## Element-Specific Validation: `%s`\n\n", tagName))

	if !strings.Contains(html, "<"+tagName) {
		results.WriteString("❌ Element not found in HTML\n\n")
		return results.String()
	}

	results.WriteString(fmt.Sprintf("**Description**: %s\n\n", element.Description()))

	// Extract all usages of this element
	elementRegex := regexp.MustCompile(`<` + regexp.QuoteMeta(tagName) + `[^>]*>`)
	usages := elementRegex.FindAllString(html, -1)

	results.WriteString(fmt.Sprintf("**Found %d usage(s):**\n\n", len(usages)))

	for i, usage := range usages {
		results.WriteString(fmt.Sprintf("%d. `%s`\n", i+1, usage))

		// Validate attributes for this usage
		var attrIssues []string
		for _, attr := range element.Attributes() {
			if attr.Required() && !strings.Contains(usage, attr.Name()+"=") {
				attrIssues = append(attrIssues, fmt.Sprintf("Missing required attribute: %s", attr.Name()))
			}
		}

		if len(attrIssues) > 0 {
			results.WriteString("   - Issues: " + strings.Join(attrIssues, "; ") + "\n")
		} else {
			results.WriteString("   - ✅ No issues found\n")
		}
	}

	results.WriteString("\n")

	// Show available features
	if len(element.Attributes()) > 0 {
		results.WriteString("**Available Attributes:**\n")
		for _, attr := range element.Attributes() {
			required := ""
			if attr.Required() {
				required = " (required)"
			}
			results.WriteString(fmt.Sprintf("- `%s`: %s%s\n", attr.Name(), attr.Description(), required))
		}
		results.WriteString("\n")
	}

	return results.String()
}

// getSlotNames extracts slot names from the slots slice
func getSlotNames(slots []types.Slot) string {
	var names []string
	for _, slot := range slots {
		name := slot.Name()
		if name == "" {
			name = "default"
		}
		names = append(names, name)
	}
	return strings.Join(names, ", ")
}
