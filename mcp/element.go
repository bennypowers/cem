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
package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// handleElementResource provides detailed information about a specific element
func (s *Server) handleElementResource(ctx context.Context, tagName string) (*mcp.ReadResourceResult, error) {
	element, err := s.registry.GetElementInfo(tagName)
	if err != nil {
		return nil, fmt.Errorf("element '%s' not found: %w", tagName, err)
	}

	// Create enhanced element data with additional context
	elementData := s.buildEnhancedElementData(element)

	contents, err := json.MarshalIndent(elementData, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal element data: %w", err)
	}

	return &mcp.ReadResourceResult{
		Contents: []*mcp.ResourceContents{{
			URI:      fmt.Sprintf("cem://element/%s", tagName),
			MIMEType: "application/json",
			Text:     string(contents),
		}},
	}, nil
}

// EnhancedElementData provides rich element information for AI consumption
type EnhancedElementData struct {
	*ElementInfo

	// Additional AI-friendly metadata
	Categories        []string          `json:"categories"`
	UsagePatterns     []UsagePattern    `json:"usagePatterns"`
	AccessibilityInfo AccessibilityInfo `json:"accessibilityInfo"`
	CssIntegration    CssIntegration    `json:"cssIntegration"`
	Examples          []ElementExample  `json:"examples"`
	Guidelines        []string          `json:"guidelines"`
	RelatedElements   []string          `json:"relatedElements,omitempty"`
}

// UsagePattern describes common ways to use the element
type UsagePattern struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Example     string `json:"example"`
	Context     string `json:"context"`
}

// AccessibilityInfo provides accessibility guidance
type AccessibilityInfo struct {
	RecommendedRole   string   `json:"recommendedRole,omitempty"`
	AriaAttributes    []string `json:"ariaAttributes,omitempty"`
	KeyboardSupport   string   `json:"keyboardSupport,omitempty"`
	ScreenReaderNotes string   `json:"screenReaderNotes,omitempty"`
	FocusManagement   string   `json:"focusManagement,omitempty"`
	ColorContrast     string   `json:"colorContrast,omitempty"`
}

// CssIntegration provides CSS usage guidance
type CssIntegration struct {
	ThemingApproach    string            `json:"themingApproach,omitempty"`
	ResponsivePatterns []string          `json:"responsivePatterns,omitempty"`
	PropertyExamples   map[string]string `json:"propertyExamples,omitempty"`
	PartExamples       map[string]string `json:"partExamples,omitempty"`
	StateExamples      map[string]string `json:"stateExamples,omitempty"`
}

// ElementExample provides usage examples
type ElementExample struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Html        string `json:"html"`
	Context     string `json:"context"`
}

// buildEnhancedElementData creates comprehensive element data for AI consumption
func (s *Server) buildEnhancedElementData(element *ElementInfo) *EnhancedElementData {
	enhanced := &EnhancedElementData{
		ElementInfo:       element,
		Categories:        s.categorizeElement(element),
		UsagePatterns:     s.generateUsagePatterns(element),
		AccessibilityInfo: s.generateAccessibilityInfo(element),
		CssIntegration:    s.generateCssIntegration(element),
		Examples:          s.generateElementExamples(element),
		Guidelines:        s.extractGuidelines(element),
		RelatedElements:   s.findRelatedElements(element),
	}

	return enhanced
}

// categorizeElement assigns semantic categories to elements based on their features
func (s *Server) categorizeElement(element *ElementInfo) []string {
	var categories []string
	tagName := strings.ToLower(element.TagName)

	// Form elements
	if s.isFormElement(element) {
		categories = append(categories, "form")
	}

	// Layout elements
	if s.isLayoutElement(element) {
		categories = append(categories, "layout")
	}

	// Interactive elements
	if strings.Contains(tagName, "button") || strings.Contains(tagName, "link") {
		categories = append(categories, "interactive")
	}

	// Navigation elements
	if strings.Contains(tagName, "nav") || strings.Contains(tagName, "menu") {
		categories = append(categories, "navigation")
	}

	// Display elements
	if strings.Contains(tagName, "card") || strings.Contains(tagName, "panel") {
		categories = append(categories, "display")
	}

	// Content elements
	if strings.Contains(tagName, "text") || strings.Contains(tagName, "content") {
		categories = append(categories, "content")
	}

	// Media elements
	if strings.Contains(tagName, "image") || strings.Contains(tagName, "video") || strings.Contains(tagName, "audio") {
		categories = append(categories, "media")
	}

	// If no specific category matches, use generic
	if len(categories) == 0 {
		categories = append(categories, "generic")
	}

	return categories
}

// generateUsagePatterns creates common usage patterns for the element
func (s *Server) generateUsagePatterns(element *ElementInfo) []UsagePattern {
	var patterns []UsagePattern
	tagName := element.TagName

	// Basic usage pattern
	basicExample := fmt.Sprintf("<%s", tagName)

	// Add required attributes
	for _, attr := range element.Attributes() {
		if attr.Required() {
			if defaultVal := attr.Default(); defaultVal != "" {
				value := strings.Trim(defaultVal, "\"'")
				basicExample += fmt.Sprintf(` %s="%s"`, attr.Name(), value)
			} else if len(attr.Values()) > 0 {
				basicExample += fmt.Sprintf(` %s="%s"`, attr.Name(), attr.Values()[0])
			}
		}
	}

	// Add default slot content if available
	hasDefaultSlot := false
	for _, slot := range element.Slots() {
		if slot.Name() == "" {
			hasDefaultSlot = true
			break
		}
	}

	if hasDefaultSlot {
		basicExample += ">Content</" + tagName + ">"
	} else {
		basicExample += "></" + tagName + ">"
	}

	patterns = append(patterns, UsagePattern{
		Name:        "Basic Usage",
		Description: "Standard element usage with minimal configuration",
		Example:     basicExample,
		Context:     "general",
	})

	// Form context pattern
	if s.isFormElement(element) {
		patterns = append(patterns, UsagePattern{
			Name:        "Form Integration",
			Description: "Usage within forms with proper accessibility",
			Example:     s.generateFormExample(element),
			Context:     "form",
		})
	}

	// Layout context pattern
	if s.isLayoutElement(element) {
		patterns = append(patterns, UsagePattern{
			Name:        "Layout Pattern",
			Description: "Usage in layout contexts with responsive design",
			Example:     s.generateLayoutExample(element),
			Context:     "layout",
		})
	}

	return patterns
}

// generateAccessibilityInfo creates accessibility guidance
func (s *Server) generateAccessibilityInfo(element *ElementInfo) AccessibilityInfo {
	info := AccessibilityInfo{}
	tagName := strings.ToLower(element.TagName)

	// Recommend roles based on element type
	switch {
	case strings.Contains(tagName, "button"):
		info.RecommendedRole = "button"
		info.KeyboardSupport = "Space and Enter keys should activate the button"
		info.FocusManagement = "Must be focusable and have visible focus indicator"
	case strings.Contains(tagName, "input") || strings.Contains(tagName, "field"):
		info.AriaAttributes = []string{"aria-label", "aria-describedby", "aria-required"}
		info.KeyboardSupport = "Standard input keyboard navigation"
		info.ScreenReaderNotes = "Ensure proper labeling for screen readers"
	case strings.Contains(tagName, "nav") || strings.Contains(tagName, "menu"):
		info.RecommendedRole = "navigation"
		info.AriaAttributes = []string{"aria-label", "aria-current"}
		info.KeyboardSupport = "Arrow key navigation for menu items"
	case strings.Contains(tagName, "dialog") || strings.Contains(tagName, "modal"):
		info.RecommendedRole = "dialog"
		info.AriaAttributes = []string{"aria-labelledby", "aria-describedby", "aria-modal"}
		info.FocusManagement = "Focus should be trapped within the dialog"
		info.KeyboardSupport = "Escape key should close the dialog"
	}

	// Check for accessibility-related attributes
	var ariaAttrs []string
	for _, attr := range element.Attributes() {
		attrName := attr.Name()
		if strings.HasPrefix(attrName, "aria-") || attrName == "role" {
			ariaAttrs = append(ariaAttrs, attrName)
		}
	}
	if len(ariaAttrs) > 0 {
		info.AriaAttributes = append(info.AriaAttributes, ariaAttrs...)
	}

	// Color contrast guidance for themed elements
	if len(element.CssProperties()) > 0 {
		info.ColorContrast = "Ensure sufficient color contrast (4.5:1) when customizing colors"
	}

	return info
}

// generateCssIntegration creates CSS integration guidance
func (s *Server) generateCssIntegration(element *ElementInfo) CssIntegration {
	integration := CssIntegration{
		PropertyExamples: make(map[string]string),
		PartExamples:     make(map[string]string),
		StateExamples:    make(map[string]string),
	}

	// CSS Custom Properties examples
	if len(element.CssProperties()) > 0 {
		integration.ThemingApproach = "Use CSS custom properties for consistent theming"
		for _, prop := range element.CssProperties() {
			integration.PropertyExamples[prop.Name()] = s.generatePropertyExample(prop)
		}
	}

	// CSS Parts examples
	for _, part := range element.CssParts() {
		integration.PartExamples[part.Name()] = fmt.Sprintf("%s::%s { /* Style the %s */ }",
			element.TagName, part.Name(), part.Description())
	}

	// CSS States examples
	for _, state := range element.CssStates() {
		integration.StateExamples[state.Name()] = fmt.Sprintf("%s:%s { /* Style when %s */ }",
			element.TagName, state.Name(), state.Description())
	}

	// Responsive patterns
	if len(element.CssProperties()) > 0 {
		integration.ResponsivePatterns = []string{
			"Use container queries for component-level responsive design",
			"Leverage CSS custom properties for responsive values",
			"Consider mobile-first approach with progressive enhancement",
		}
	}

	return integration
}

// generateElementExamples creates comprehensive usage examples
func (s *Server) generateElementExamples(element *ElementInfo) []ElementExample {
	var examples []ElementExample

	// Basic example
	examples = append(examples, ElementExample{
		Title:       "Basic Usage",
		Description: fmt.Sprintf("Standard %s element with default configuration", element.TagName),
		Html:        s.generateBasicExample(element),
		Context:     "general",
	})

	// Accessible example
	examples = append(examples, ElementExample{
		Title:       "Accessible Implementation",
		Description: "Element configured with accessibility best practices",
		Html:        s.generateAccessibleExample(element),
		Context:     "accessibility",
	})

	// Themed example if CSS properties are available
	if len(element.CssProperties()) > 0 {
		examples = append(examples, ElementExample{
			Title:       "Custom Styling",
			Description: "Element with custom theme using CSS properties",
			Html:        s.generateThemedExample(element),
			Context:     "styling",
		})
	}

	return examples
}

// Helper methods for generating examples
func (s *Server) generateBasicExample(element *ElementInfo) string {
	// Implementation details for basic example generation
	return fmt.Sprintf("<%s>Basic content</%s>", element.TagName, element.TagName)
}

func (s *Server) generateAccessibleExample(element *ElementInfo) string {
	// Implementation details for accessible example generation
	return fmt.Sprintf("<%s role=\"button\" aria-label=\"Accessible %s\">Content</%s>",
		element.TagName, element.TagName, element.TagName)
}

func (s *Server) generateThemedExample(element *ElementInfo) string {
	// Implementation details for themed example generation
	return fmt.Sprintf("<%s style=\"%s: custom-value;\">Themed content</%s>",
		element.TagName, element.CssProperties()[0].Name(), element.TagName)
}

func (s *Server) generatePropertyExample(prop CssProperty) string {
	return fmt.Sprintf("%s: %s; /* %s */", prop.Name(), prop.Initial(), prop.Description())
}

func (s *Server) isFormElement(element *ElementInfo) bool {
	tagName := strings.ToLower(element.TagName)
	return strings.Contains(tagName, "input") || strings.Contains(tagName, "button") ||
		strings.Contains(tagName, "field") || strings.Contains(tagName, "form")
}

func (s *Server) isLayoutElement(element *ElementInfo) bool {
	tagName := strings.ToLower(element.TagName)
	return strings.Contains(tagName, "grid") || strings.Contains(tagName, "flex") ||
		strings.Contains(tagName, "layout") || strings.Contains(tagName, "card")
}

func (s *Server) generateFormExample(element *ElementInfo) string {
	return fmt.Sprintf(`<form>
  <label for="example">Label:</label>
  <%s id="example" name="example" required>
    Content
  </%s>
</form>`, element.TagName, element.TagName)
}

func (s *Server) generateLayoutExample(element *ElementInfo) string {
	return fmt.Sprintf(`<section class="layout">
  <%s>
    <h2>Section Title</h2>
    <p>Section content</p>
  </%s>
</section>`, element.TagName, element.TagName)
}

func (s *Server) extractGuidelines(element *ElementInfo) []string {
	var guidelines []string

	// Extract from element description
	desc := element.Description
	if desc != "" {
		// Look for guideline keywords
		if strings.Contains(strings.ToLower(desc), "should") ||
			strings.Contains(strings.ToLower(desc), "must") ||
			strings.Contains(strings.ToLower(desc), "recommended") {
			guidelines = append(guidelines, desc)
		}
	}

	// Extract from attribute guidelines
	for _, attr := range element.Attributes() {
		attrGuidelines := attr.Guidelines()
		guidelines = append(guidelines, attrGuidelines...)
	}

	return guidelines
}

func (s *Server) findRelatedElements(element *ElementInfo) []string {
	// Simple implementation - could be enhanced with more sophisticated matching
	var related []string
	tagName := element.TagName

	elements := s.registry.GetAllElements()
	for otherTagName := range elements {
		if otherTagName != tagName {
			// Check for name similarity
			if s.areElementsRelated(tagName, otherTagName) {
				related = append(related, otherTagName)
			}
		}
	}

	return related
}

func (s *Server) areElementsRelated(tagName1, tagName2 string) bool {
	// Simple heuristic for related elements
	commonPrefixes := []string{"button", "input", "form", "card", "nav", "menu"}

	for _, prefix := range commonPrefixes {
		if strings.HasPrefix(tagName1, prefix) && strings.HasPrefix(tagName2, prefix) {
			return true
		}
	}

	return false
}
