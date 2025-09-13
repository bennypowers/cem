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

	"bennypowers.dev/cem/lsp/types"
	mcpTypes "bennypowers.dev/cem/mcp/types"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// ValidationTemplateData specific to validation tools
type ValidationTemplateData struct {
	BaseTemplateData
	Html                         string
	FoundElements                []ElementWithIssues
	ManifestIssues               []ValidationIssue
	ManifestFeatures             []ValidationFeature
	SlotContentIssues            []SlotContentIssue
	AttributeConflicts           []AttributeConflict
	ContentAttributeRedundancies []ContentAttributeRedundancy
	SpecificElement              *ElementValidationResult
}

// NewValidationTemplateData creates validation template data
func NewValidationTemplateData(element mcpTypes.ElementInfo, context string, options map[string]string, html string) ValidationTemplateData {
	return ValidationTemplateData{
		BaseTemplateData: NewBaseTemplateData(element, context, options),
		Html:             html,
	}
}

// ValidateHtmlArgs represents the arguments for the validate_html tool
type ValidateHtmlArgs struct {
	Html    string `json:"html"`
	Context string `json:"context,omitempty"`
	TagName string `json:"tagName,omitempty"`
}

// handleValidateHtml validates HTML for accessibility and custom element compliance using tree-sitter
func handleValidateHtml(
	ctx context.Context,
	req *mcp.CallToolRequest,
	registry mcpTypes.MCPContext,
) (*mcp.CallToolResult, error) {
	// Parse args from request
	var validateArgs ValidateHtmlArgs
	if req.Params.Arguments != nil {
		if argsData, err := json.Marshal(req.Params.Arguments); err != nil {
			return nil, fmt.Errorf("failed to marshal args: %w", err)
		} else if err := json.Unmarshal(argsData, &validateArgs); err != nil {
			return nil, fmt.Errorf("failed to unmarshal args: %w", err)
		}
	}

	// Use tree-sitter to parse and validate HTML
	validationResult, err := validateHtmlWithTreeSitter(validateArgs.Html, registry)
	if err != nil {
		return nil, fmt.Errorf("failed to validate HTML: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: validationResult,
			},
		},
	}, nil
}

// validateHtmlWithTreeSitter performs HTML validation using tree-sitter parsing
func validateHtmlWithTreeSitter(html string, registry mcpTypes.MCPContext) (string, error) {
	// Get the shared document manager from registry
	dm := registry.DocumentManager()

	// Create a document from the HTML content
	doc := dm.OpenDocument("temp://validation.html", html, 1)

	// Prepare validation data using existing template structure
	data := ValidationTemplateData{
		Html:             html,
		BaseTemplateData: BaseTemplateData{Context: ""}, // Default context for general validation
	}

	// Use tree-sitter to find custom elements
	elements, err := doc.FindCustomElements(dm)
	if err != nil {
		return "", fmt.Errorf("failed to find custom elements: %w", err)
	}

	// Convert tree-sitter results to template-compatible format
	for _, element := range elements {
		// Check if element exists in registry
		registryElement, err := registry.ElementInfo(element.TagName)
		if err != nil {
			// Unknown element - add to manifest issues
			data.ManifestIssues = append(data.ManifestIssues, ValidationIssue{
				Type:    "unknown-element",
				Element: element.TagName,
				Message: fmt.Sprintf("Custom element '%s' not found in registry", element.TagName),
			})
		} else {
			// Add to found elements
			data.FoundElements = append(data.FoundElements, ElementWithIssues{
				ElementInfo: registryElement,
				UsageCount:  1, // Tree-sitter found it, so at least 1 usage
			})

			// Validate attributes using tree-sitter parsed data
			attributeIssues := validateElementAttributes(element, registryElement)
			data.ManifestIssues = append(data.ManifestIssues, attributeIssues...)

			// Validate slot content guidelines
			slotIssues := validateSlotContentGuidelines(element, registryElement)
			data.SlotContentIssues = append(data.SlotContentIssues, slotIssues...)

			// Validate attribute combinations
			conflicts := validateAttributeCombinations(element, registryElement)
			data.AttributeConflicts = append(data.AttributeConflicts, conflicts...)

			// Validate content/attribute redundancy
			redundancies := validateContentAttributeRedundancy(element, registryElement)
			data.ContentAttributeRedundancies = append(data.ContentAttributeRedundancies, redundancies...)

			// Add manifest features
			if registryElement.Description() != "" {
				data.ManifestFeatures = append(data.ManifestFeatures, ValidationFeature{
					Type:    "manifest-element",
					Details: fmt.Sprintf("Element <%s> properly registered in manifest", element.TagName),
					Element: element.TagName,
				})
			}
		}
	}

	return RenderTemplate("html_validation_results", data)
}

// validateElementAttributes validates attributes for a custom element using tree-sitter parsed data
func validateElementAttributes(element types.CustomElementMatch, registryElement mcpTypes.ElementInfo) []ValidationIssue {
	var issues []ValidationIssue

	// Check required attributes
	for _, attr := range registryElement.Attributes() {
		if attr.Required() {
			if _, exists := element.Attributes[attr.Name()]; !exists {
				issues = append(issues, ValidationIssue{
					Type:    "missing-required-attribute",
					Element: element.TagName,
					Message: fmt.Sprintf("Required attribute '%s' is missing", attr.Name()),
				})
			}
		}
	}

	// Validate attribute values against manifest constraints
	for attrName, attrMatch := range element.Attributes {
		// Find the attribute definition in the manifest
		var manifestAttr mcpTypes.Attribute
		var found bool
		for _, attr := range registryElement.Attributes() {
			if attr.Name() == attrName {
				manifestAttr = attr
				found = true
				break
			}
		}

		if !found {
			issues = append(issues, ValidationIssue{
				Type:    "unknown-attribute",
				Element: element.TagName,
				Message: fmt.Sprintf("Unknown attribute '%s' for element <%s>", attrName, element.TagName),
			})
			continue
		}

		// Validate attribute value if it has constraints
		if len(manifestAttr.Values()) > 0 && attrMatch.Value != "" {
			validValues := manifestAttr.Values()
			isValid := false
			for _, validValue := range validValues {
				if attrMatch.Value == validValue {
					isValid = true
					break
				}
			}
			if !isValid {
				issues = append(issues, ValidationIssue{
					Type:    "invalid-attribute-value",
					Element: element.TagName,
					Message: fmt.Sprintf("Invalid value '%s' for attribute '%s'. Expected one of: %v",
						attrMatch.Value, attrName, validValues),
				})
			}
		}
	}

	return issues
}

// validateSlotContentGuidelines validates slotted content against manifest guidelines
func validateSlotContentGuidelines(element types.CustomElementMatch, registryElement mcpTypes.ElementInfo) []SlotContentIssue {
	var issues []SlotContentIssue

	// Check if element has slots defined in manifest
	slots := registryElement.Slots()
	if len(slots) == 0 {
		return issues
	}

	// For each slot, check if there are content guidelines in the description
	for _, slot := range slots {
		description := slot.Description()
		if description == "" {
			continue
		}

		// Look for guideline keywords in slot descriptions
		// This is a basic implementation that can be enhanced
		if strings.Contains(strings.ToLower(description), "only") ||
			strings.Contains(strings.ToLower(description), "must") ||
			strings.Contains(strings.ToLower(description), "should not") ||
			strings.Contains(strings.ToLower(description), "avoid") {

			// Found potential content guidelines - this is where we would
			// analyze the actual slotted content using tree-sitter
			// For now, create a placeholder issue to show the structure works
			if strings.Contains(strings.ToLower(description), "interactive") {
				issues = append(issues, SlotContentIssue{
					SlotName:         slot.Name(),
					ElementTagName:   element.TagName,
					Guideline:        description,
					ViolationMessage: "Potential interactive content guideline - check slot content manually",
				})
			}
		}
	}

	return issues
}

// validateAttributeCombinations detects conflicting attribute combinations
func validateAttributeCombinations(element types.CustomElementMatch, registryElement mcpTypes.ElementInfo) []AttributeConflict {
	var conflicts []AttributeConflict

	// Look for common contradictory attribute patterns
	attributes := element.Attributes

	// Check for loading/lazy conflicts
	if loadingAttr, hasLoading := attributes["loading"]; hasLoading {
		if lazyAttr, hasLazy := attributes["lazy"]; hasLazy {
			if loadingAttr.Value == "eager" && (lazyAttr.Value == "true" || lazyAttr.Value == "") {
				conflicts = append(conflicts, AttributeConflict{
					ElementTagName: element.TagName,
					Attribute1:     "loading",
					Value1:         loadingAttr.Value,
					Attribute2:     "lazy",
					Value2:         lazyAttr.Value,
					ConflictReason: "Eager loading conflicts with lazy loading",
				})
			}
		}
	}

	// Check for disabled/interactive conflicts
	if disabledAttr, hasDisabled := attributes["disabled"]; hasDisabled {
		if interactiveAttr, hasInteractive := attributes["interactive"]; hasInteractive {
			if (disabledAttr.Value == "true" || disabledAttr.Value == "") &&
				(interactiveAttr.Value == "true" || interactiveAttr.Value == "") {
				conflicts = append(conflicts, AttributeConflict{
					ElementTagName: element.TagName,
					Attribute1:     "disabled",
					Value1:         disabledAttr.Value,
					Attribute2:     "interactive",
					Value2:         interactiveAttr.Value,
					ConflictReason: "Disabled elements cannot be interactive",
				})
			}
		}
	}

	return conflicts
}

// validateContentAttributeRedundancy detects when slot content overrides attributes
func validateContentAttributeRedundancy(element types.CustomElementMatch, registryElement mcpTypes.ElementInfo) []ContentAttributeRedundancy {
	var redundancies []ContentAttributeRedundancy

	// Check for common attribute/slot combinations that might be redundant
	attributes := element.Attributes
	slots := registryElement.Slots()

	// Look for attributes that have corresponding slots
	for _, slot := range slots {
		slotName := slot.Name()

		// Check for common patterns where attributes might be overridden by slots
		if slotName == "label" {
			if labelAttr, hasLabel := attributes["label"]; hasLabel {
				redundancies = append(redundancies, ContentAttributeRedundancy{
					ElementTagName: element.TagName,
					AttributeName:  "label",
					AttributeValue: labelAttr.Value,
					SlotName:       slotName,
					SlotContent:    "[slot content not analyzed - tree-sitter enhancement needed]",
				})
			}
		}

		if slotName == "title" {
			if titleAttr, hasTitle := attributes["title"]; hasTitle {
				redundancies = append(redundancies, ContentAttributeRedundancy{
					ElementTagName: element.TagName,
					AttributeName:  "title",
					AttributeValue: titleAttr.Value,
					SlotName:       slotName,
					SlotContent:    "[slot content not analyzed - tree-sitter enhancement needed]",
				})
			}
		}

		if slotName == "icon" {
			if iconAttr, hasIcon := attributes["icon"]; hasIcon {
				redundancies = append(redundancies, ContentAttributeRedundancy{
					ElementTagName: element.TagName,
					AttributeName:  "icon",
					AttributeValue: iconAttr.Value,
					SlotName:       slotName,
					SlotContent:    "[slot content not analyzed - tree-sitter enhancement needed]",
				})
			}
		}
	}

	return redundancies
}
