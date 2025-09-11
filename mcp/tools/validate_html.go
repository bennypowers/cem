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

	// Prepare validation data
	validationData := collectValidationData(validateArgs, registry)

	// Render validation results using template
	response, err := renderValidationTemplate("html_validation_results", validationData)
	if err != nil {
		return nil, fmt.Errorf("failed to render validation template: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: response,
			},
		},
	}, nil
}

// collectValidationData gathers all validation information into structured data
func collectValidationData(args ValidateHtmlArgs, registry types.Registry) HTMLValidationData {
	html := args.Html
	elements := registry.AllElements()
	
	validationData := HTMLValidationData{
		Html:    html,
		Context: args.Context,
	}

	// Find custom elements in HTML
	for tagName, element := range elements {
		if strings.Contains(html, "<"+tagName) || strings.Contains(html, "</"+tagName) {
			validationData.FoundElements = append(validationData.FoundElements, ElementWithIssues{
				ElementInfo: element,
				UsageCount:  strings.Count(html, "<"+tagName),
			})
		}
	}

	// Collect manifest compliance issues
	validationData.ManifestIssues, validationData.ManifestFeatures = collectManifestValidation(html, elements)

	// Collect semantic structure issues
	validationData.SemanticIssues, validationData.SemanticSuggestions = collectSemanticValidation(html)

	// Handle specific element validation
	if args.TagName != "" {
		if element, err := registry.ElementInfo(args.TagName); err == nil {
			validationData.SpecificElement = collectSpecificElementValidation(html, args.TagName, element)
		} else {
			validationData.SpecificElement = &ElementValidationResult{
				TagName:      args.TagName,
				ElementFound: false,
			}
		}
	}

	return validationData
}

// collectManifestValidation checks custom element usage against manifests
func collectManifestValidation(html string, elements map[string]types.ElementInfo) ([]ValidationIssue, []ValidationFeature) {
	var issues []ValidationIssue
	var features []ValidationFeature

	for tagName, element := range elements {
		if strings.Contains(html, "<"+tagName) {
			// Extract element usages
			elementRegex := regexp.MustCompile(`<` + regexp.QuoteMeta(tagName) + `[^>]*>`)
			usages := elementRegex.FindAllString(html, -1)

			for _, usage := range usages {
				// Check required attributes
				for _, attr := range element.Attributes() {
					if attr.Required() && !strings.Contains(usage, attr.Name()+"=") {
						issues = append(issues, ValidationIssue{
							Type:      "missing-attribute",
							Element:   tagName,
							Attribute: attr.Name(),
							Priority:  "error",
						})
					}
				}

				// Check attribute values
				for _, attr := range element.Attributes() {
					if strings.Contains(usage, attr.Name()+"=") && len(attr.Values()) > 0 {
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
								issues = append(issues, ValidationIssue{
									Type:      "invalid-value",
									Element:   tagName,
									Attribute: attr.Name(),
									Actual:    value,
									Expected:  strings.Join(validValues, ", "),
									Priority:  "error",
								})
							}
						}
					}
				}
			}

			// Collect available features for guidance
			if len(element.Slots()) > 0 {
				features = append(features, ValidationFeature{
					Type:    "slots",
					Element: tagName,
					Details: getSlotNames(element.Slots()),
				})
			}

			if len(element.Guidelines()) > 0 {
				features = append(features, ValidationFeature{
					Type:    "guidelines",
					Element: tagName,
					Details: strings.Join(element.Guidelines(), "; "),
				})
			}

			cssCount := len(element.CssProperties()) + len(element.CssParts()) + len(element.CssStates())
			if cssCount > 0 {
				features = append(features, ValidationFeature{
					Type:    "css-apis",
					Element: tagName,
					Details: fmt.Sprintf("%d", cssCount),
				})
			}
		}
	}

	return issues, features
}

// collectSemanticValidation analyzes HTML for semantic structure
func collectSemanticValidation(html string) ([]ValidationIssue, []ValidationSuggestion) {
	var issues []ValidationIssue
	var suggestions []ValidationSuggestion

	// Check document structure
	hasMain := strings.Contains(html, "<main")
	if !hasMain {
		issues = append(issues, ValidationIssue{
			Type:     "semantic-issue",
			Message:  "Consider using <main> element for primary content",
			Priority: "warning",
		})
	}

	// Check for div abuse
	divCount := strings.Count(html, "<div")
	sectionCount := strings.Count(html, "<section")
	articleCount := strings.Count(html, "<article")

	if divCount > 5 && (sectionCount+articleCount) == 0 {
		suggestions = append(suggestions, ValidationSuggestion{
			Type:     "semantic-elements",
			Priority: "suggestion",
			Message:  "Consider using semantic elements like <section> or <article> instead of multiple <div> elements",
		})
	}

	// Check list usage
	if strings.Contains(html, "•") || strings.Contains(html, "1.") || strings.Contains(html, "-") {
		if !strings.Contains(html, "<ul") && !strings.Contains(html, "<ol") {
			suggestions = append(suggestions, ValidationSuggestion{
				Type:     "list-structure",
				Priority: "suggestion",
				Message:  "Consider using proper list elements (<ul> or <ol>) for list content",
			})
		}
	}

	return issues, suggestions
}

// collectSpecificElementValidation analyzes specific element usage
func collectSpecificElementValidation(html, tagName string, element types.ElementInfo) *ElementValidationResult {
	result := &ElementValidationResult{
		ElementInfo:  element,
		TagName:      tagName,
		ElementFound: strings.Contains(html, "<"+tagName),
	}

	if !result.ElementFound {
		return result
	}

	// Extract all usages
	elementRegex := regexp.MustCompile(`<` + regexp.QuoteMeta(tagName) + `[^>]*>`)
	usages := elementRegex.FindAllString(html, -1)

	for _, usage := range usages {
		elementUsage := ElementUsage{Html: usage}

		// Check for attribute issues
		for _, attr := range element.Attributes() {
			if attr.Required() && !strings.Contains(usage, attr.Name()+"=") {
				elementUsage.Issues = append(elementUsage.Issues, fmt.Sprintf("Missing required attribute: %s", attr.Name()))
			}
		}

		result.Usages = append(result.Usages, elementUsage)
	}

	return result
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
