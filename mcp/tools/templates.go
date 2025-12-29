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
	"embed"

	"bennypowers.dev/cem/mcp/templates"
	"bennypowers.dev/cem/mcp/types"
)

//go:embed templates/*.md
var toolsTemplateFiles embed.FS

func init() {
	// Register tools templates with the global template pool
	templates.RegisterTemplateSource("tools", &toolsTemplateFiles)
}

// Use unified template data types from types package
type TemplateDataProvider = types.TemplateDataProvider
type BaseTemplateData = types.BaseTemplateData

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

// Use unified template data constructors from types package
var NewBaseTemplateData = types.NewBaseTemplateData
var NewBaseTemplateDataWithSchema = types.NewBaseTemplateDataWithSchema

// RenderTemplate renders a template using the global thread-safe pool
func RenderTemplate(templateName string, data interface{}) (string, error) {
	return templates.RenderTemplate(templateName, data)
}
