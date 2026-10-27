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
	"log/slog"

	"bennypowers.dev/cem/manifest"
	mcpTypes "bennypowers.dev/cem/mcp/types"
	"bennypowers.dev/cem/validate/markup"
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
	_ context.Context,
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
	// Prepare validation data using existing template structure
	data := ValidationTemplateData{
		Html:                         html,
		BaseTemplateData:             BaseTemplateData{Context: ""}, // Default context for general validation
		FoundElements:                []ElementWithIssues{},
		ManifestIssues:               []ValidationIssue{},
		ManifestFeatures:             []ValidationFeature{},
		SlotContentIssues:            []SlotContentIssue{},
		AttributeConflicts:           []AttributeConflict{},
		ContentAttributeRedundancies: []ContentAttributeRedundancy{},
	}

	// Discovery uses the same normalized parser as validation, so uppercase
	// markup produces both a FoundElements entry and compliance issues.
	elements, err := markup.FindCustomElements(html, markup.Options{Format: "html"})
	if err != nil {
		return "", fmt.Errorf("failed to find custom elements: %w", err)
	}

	for _, element := range elements {
		registryElement, err := registry.ElementInfo(element.TagName)
		if err != nil {
			// Unknown elements are reported by markup.Validate.
			continue
		}

		data.FoundElements = append(data.FoundElements, ElementWithIssues{
			ElementInfo: registryElement,
			UsageCount:  1, // Tree-sitter found it, so at least 1 usage
		})

		if registryElement.Description() != "" {
			data.ManifestFeatures = append(data.ManifestFeatures, ValidationFeature{
				Type:    "manifest-element",
				Details: fmt.Sprintf("Element <%s> properly registered in manifest", element.TagName),
				Element: element.TagName,
			})
		}
	}

	// Discovery and validation each parse the markup once. Sharing a parse
	// would require a richer result API; two passes are bounded and keep the
	// public Validate signature small. registryManifest is rebuilt per call
	// from the registry's cached element map instead of being cached here, so
	// reloaded manifests are always reflected and no lifecycle owner is
	// introduced.
	issues, err := markup.Validate(registryManifest(registry), html, markup.Options{
		Format:        "html",
		ValidateAttrs: true,
		ValidateSlots: true,
		Strict:        true,
	})
	if err != nil {
		return "", fmt.Errorf("failed to validate markup: %w", err)
	}
	applyMarkupIssues(&data, issues)

	return RenderTemplate("html_validation_results", data)
}

// registryManifest builds a manifest package view over the registry's element
// declarations so the markup validator can resolve tags without depending on
// the MCP context or an LSP document manager. Declarations keep their original
// module and package backreferences, so inherited members remain visible.
func registryManifest(registry mcpTypes.MCPContext) *manifest.Package {
	mod := manifest.Module{
		Kind: "javascript-module",
		Path: "mcp://registry",
	}
	for _, element := range registry.AllElements() {
		if decl := element.Declaration(); decl != nil {
			mod.Declarations = append(mod.Declarations, decl)
		}
	}
	return &manifest.Package{
		SchemaVersion: "2.1.0",
		Modules:       []manifest.Module{mod},
	}
}

func applyMarkupIssues(data *ValidationTemplateData, issues []markup.Issue) {
	for _, issue := range issues {
		details := issue.Details()
		switch issue.Rule {
		case markup.RuleUnknownElement:
			data.ManifestIssues = append(data.ManifestIssues, ValidationIssue{
				Type:    "unknown-element",
				Element: issue.Element,
				Message: issue.Message,
			})
		case markup.RuleUnknownAttribute:
			data.ManifestIssues = append(data.ManifestIssues, ValidationIssue{
				Type:    "unknown-attribute",
				Element: issue.Element,
				Message: issue.Message,
			})
		case markup.RuleInvalidAttributeValue:
			data.ManifestIssues = append(data.ManifestIssues, ValidationIssue{
				Type:      "invalid-attribute-value",
				Element:   issue.Element,
				Attribute: details[markup.DetailAttribute],
				Actual:    details[markup.DetailValue],
				Expected:  details[markup.DetailExpected],
				Message:   issue.Message,
			})
		case markup.RuleSlotContentGuideline:
			data.SlotContentIssues = append(data.SlotContentIssues, SlotContentIssue{
				SlotName:         details[markup.DetailSlot],
				ElementTagName:   issue.Element,
				Guideline:        details[markup.DetailGuideline],
				ViolationMessage: issue.Message,
			})
		case markup.RuleAttributeConflict:
			data.AttributeConflicts = append(data.AttributeConflicts, AttributeConflict{
				ElementTagName: issue.Element,
				Attribute1:     details[markup.DetailAttribute1],
				Value1:         details[markup.DetailValue1],
				Attribute2:     details[markup.DetailAttribute2],
				Value2:         details[markup.DetailValue2],
				ConflictReason: details[markup.DetailReason],
			})
		case markup.RuleContentAttributeRedundancy:
			data.ContentAttributeRedundancies = append(data.ContentAttributeRedundancies, ContentAttributeRedundancy{
				ElementTagName: issue.Element,
				AttributeName:  details[markup.DetailAttribute],
				AttributeValue: details[markup.DetailValue],
				SlotName:       details[markup.DetailSlot],
				SlotContent:    details[markup.DetailContent],
			})
		case markup.RuleDeprecatedElement, markup.RuleDeprecatedAttribute, markup.RuleDeprecatedSlot, markup.RuleUnknownSlot:
			// Intentionally not rendered. The validate_html tool has never
			// surfaced these findings, so dropping them here keeps existing
			// MCP output stable. The exported markup API reports them for
			// other consumers.
		default:
			// A rule added to validate/markup without a mapping here would
			// otherwise disappear from the tool output silently.
			slog.Warn("unhandled markup rule in validate_html output", "rule", issue.Rule)
		}
	}
}
