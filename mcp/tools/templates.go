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
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"path/filepath"
	"strings"

	"bennypowers.dev/cem/mcp/helpers"
	"bennypowers.dev/cem/mcp/security"
	"bennypowers.dev/cem/mcp/types"
)

//go:embed templates/*.md
var templateFiles embed.FS

// TemplateRenderer provides a unified interface for rendering templates with different data types
type TemplateRenderer struct {
	funcMap        template.FuncMap
	securityPolicy security.SecurityPolicy
}

// NewTemplateRenderer creates a new template renderer with unified function map
func NewTemplateRenderer() *TemplateRenderer {
	return NewSecureTemplateRenderer(security.DefaultSecurityPolicy())
}

// NewSecureTemplateRenderer creates a new template renderer with security controls
func NewSecureTemplateRenderer(policy security.SecurityPolicy) *TemplateRenderer {
	return &TemplateRenderer{
		securityPolicy: policy,
		funcMap: createSecureFuncMap(),
	}
}

// createSecureFuncMap creates a restricted function map for template security
func createSecureFuncMap() template.FuncMap {
	return template.FuncMap{
		"title": helpers.TitleCaser.String,
		"len": func(slice interface{}) int {
			switch s := slice.(type) {
			case []types.Attribute:
				return len(s)
			case []types.Slot:
				return len(s)
			case []types.Event:
				return len(s)
			case []types.CssProperty:
				return len(s)
			case []types.CssPart:
				return len(s)
			case []types.CssState:
				return len(s)
			case []ElementWithIssues:
				return len(s)
			case []ValidationIssue:
				return len(s)
			case []ValidationFeature:
				return len(s)
			case []ValidationSuggestion:
				return len(s)
			case []ElementUsage:
				return len(s)
			case []SlotContentIssue:
				return len(s)
			case []AttributeConflict:
				return len(s)
			case []ContentAttributeRedundancy:
				return len(s)
			case []AttributeWithValue:
				return len(s)
			case []SlotWithContent:
				return len(s)
			case []string:
				return len(s)
			default:
				return 0
			}
		},
		"index": func(slice interface{}, i int) interface{} {
			switch s := slice.(type) {
			case []string:
				if i >= 0 && i < len(s) {
					return s[i]
				}
			}
			return ""
		},
		"gt": func(a, b int) bool {
			return a > b
		},
		"join": func(slice []string, sep string) string {
			return strings.Join(slice, sep)
		},
		"add": func(a, b int) int {
			return a + b
		},
	}
}

// Render renders a template with the given data using the unified template system
func (tr *TemplateRenderer) Render(templateName string, data interface{}) (string, error) {
	// Security: Validate template name to prevent path traversal
	if strings.Contains(templateName, "..") || strings.Contains(templateName, "/") {
		return "", fmt.Errorf("invalid template name: %s", templateName)
	}

	// Create template with unified function map
	tmpl := template.New(templateName).Funcs(tr.funcMap)

	// Load template content
	templatePath := filepath.Join("templates", templateName+".md")
	content, err := templateFiles.ReadFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to read template %s: %w", templateName, err)
	}

	// Note: Template files are trusted and expected to contain Go template syntax
	// Security validation is applied to user data, not template structure
	templateContent := string(content)

	// Parse template
	tmpl, err = tmpl.Parse(templateContent)
	if err != nil {
		return "", fmt.Errorf("failed to parse template %s: %w", templateName, err)
	}

	// Execute template with security context
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template %s: %w", templateName, err)
	}

	// Return the rendered result
	// Note: Security is enforced at the data input level (description sanitization)
	// rather than output validation to avoid false positives
	return buf.String(), nil
}

// Global template renderer instance
var globalRenderer = NewTemplateRenderer()

// TemplateDataProvider defines the interface for all template data types
type TemplateDataProvider interface {
	// Core element information
	GetElement() types.ElementInfo
	GetContext() string
	GetOptions() map[string]string
}

// BaseTemplateData provides common template data fields
type BaseTemplateData struct {
	Element types.ElementInfo
	Context string
	Options map[string]string
}

func (b BaseTemplateData) GetElement() types.ElementInfo { return b.Element }
func (b BaseTemplateData) GetContext() string            { return b.Context }
func (b BaseTemplateData) GetOptions() map[string]string { return b.Options }

// ElementWithCapabilities wraps an element with its capability strings
type ElementWithCapabilities struct {
	types.ElementInfo
	Capabilities []string
}

// AttributeWithValue pairs an attribute with its assigned value
type AttributeWithValue struct {
	types.Attribute
	Value string
}

// SlotWithContent pairs a slot with example content
type SlotWithContent struct {
	types.Slot
	ExampleContent string
	DefaultContent string
}

// HTMLValidationData represents data for custom element validation templates
type HTMLValidationData struct {
	Html                         string
	Context                      string
	FoundElements                []ElementWithIssues
	ManifestIssues               []ValidationIssue
	ManifestFeatures             []ValidationFeature
	SlotContentIssues            []SlotContentIssue
	AttributeConflicts           []AttributeConflict
	ContentAttributeRedundancies []ContentAttributeRedundancy
	SpecificElement              *ElementValidationResult
}

// ElementWithIssues pairs an element with its validation results
type ElementWithIssues struct {
	types.ElementInfo
	UsageCount int
}

// ValidationIssue represents a specific validation problem
type ValidationIssue struct {
	Type      string // "missing-attribute", "invalid-value", "semantic-issue"
	Element   string
	Attribute string
	Expected  string
	Actual    string
	Message   string
	Priority  string // "error", "warning", "info"
}

// ValidationFeature represents available manifest features
type ValidationFeature struct {
	Type    string // "slots", "guidelines", "css-apis"
	Element string
	Details string
}

// ValidationSuggestion represents improvement suggestions
type ValidationSuggestion struct {
	Type     string // "semantic-elements", "list-structure", "document-structure"
	Priority string // "error", "warning", "suggestion"
	Message  string
}

// ElementValidationResult represents specific element validation
type ElementValidationResult struct {
	types.ElementInfo
	TagName      string
	ElementFound bool
	Usages       []ElementUsage
}

// ElementUsage represents a specific usage of an element in HTML
type ElementUsage struct {
	Html   string
	Issues []string
}

// SlotContentIssue represents content guideline violations in slots
type SlotContentIssue struct {
	SlotName         string
	ElementTagName   string
	Guideline        string
	ViolationMessage string
}

// AttributeConflict represents conflicting attribute combinations
type AttributeConflict struct {
	ElementTagName string
	Attribute1     string
	Value1         string
	Attribute2     string
	Value2         string
	ConflictReason string
}

// ContentAttributeRedundancy represents cases where slot content overrides attributes
type ContentAttributeRedundancy struct {
	ElementTagName string
	AttributeName  string
	AttributeValue string
	SlotName       string
	SlotContent    string
}

// NewBaseTemplateData creates base template data
func NewBaseTemplateData(element types.ElementInfo, context string, options map[string]string) BaseTemplateData {
	return BaseTemplateData{
		Element: element,
		Context: context,
		Options: options,
	}
}

// RenderTemplate renders a template using the unified renderer (new API)
func RenderTemplate(templateName string, data interface{}) (string, error) {
	return globalRenderer.Render(templateName, data)
}
