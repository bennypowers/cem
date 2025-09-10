package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// GuidelinesResource provides extracted usage guidelines from manifests
type GuidelinesResource struct {
	Overview         string                    `json:"overview"`
	GeneralGuidelines []Guideline              `json:"generalGuidelines"`
	ElementGuidelines map[string][]Guideline   `json:"elementGuidelines"`
	AttributeGuidelines map[string][]Guideline `json:"attributeGuidelines"`
	AccessibilityGuidelines []Guideline        `json:"accessibilityGuidelines"`
	CssGuidelines    []Guideline              `json:"cssGuidelines"`
	BestPractices    []BestPractice           `json:"bestPractices"`
}

// Guideline represents a single usage guideline
type Guideline struct {
	Text        string   `json:"text"`
	Source      string   `json:"source"`
	Type        string   `json:"type"`
	Context     string   `json:"context,omitempty"`
	Examples    []string `json:"examples,omitempty"`
	Severity    string   `json:"severity"` // "must", "should", "recommended"
}

// BestPractice represents a design system best practice
type BestPractice struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	Examples    []string `json:"examples"`
	DoNot       []string `json:"doNot,omitempty"`
}

// handleGuidelinesResource extracts and provides usage guidelines from manifest data
func (s *Server) handleGuidelinesResource(ctx context.Context, uri string) (*mcp.ReadResourceResult, error) {
	elements := s.registry.GetAllElements()
	
	guidelines := &GuidelinesResource{
		Overview: s.generateGuidelinesOverview(len(elements)),
		GeneralGuidelines: s.extractGeneralGuidelines(),
		ElementGuidelines: make(map[string][]Guideline),
		AttributeGuidelines: make(map[string][]Guideline),
		AccessibilityGuidelines: s.generateAccessibilityGuidelines(),
		CssGuidelines: s.generateCssGuidelines(),
		BestPractices: s.generateBestPractices(),
	}
	
	// Extract guidelines from each element
	for tagName, element := range elements {
		elementGuidelines := s.extractElementGuidelines(element)
		if len(elementGuidelines) > 0 {
			guidelines.ElementGuidelines[tagName] = elementGuidelines
		}
		
		// Extract attribute-specific guidelines
		for _, attr := range element.Attributes() {
			attrGuidelines := s.extractAttributeGuidelines(attr, tagName)
			if len(attrGuidelines) > 0 {
				attrKey := fmt.Sprintf("%s.%s", tagName, attr.Name())
				guidelines.AttributeGuidelines[attrKey] = attrGuidelines
			}
		}
	}
	
	contents, err := json.MarshalIndent(guidelines, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal guidelines: %w", err)
	}

	return &mcp.ReadResourceResult{
		Contents: []*mcp.ResourceContents{{
			URI:      uri,
			MIMEType: "application/json",
			Text:     string(contents),
		}},
	}, nil
}

// generateGuidelinesOverview creates an overview of the guidelines system
func (s *Server) generateGuidelinesOverview(elementCount int) string {
	return fmt.Sprintf(`
This resource contains extracted usage guidelines and best practices from %d custom elements.
Guidelines are categorized by severity:
- "must": Required for proper functionality or accessibility
- "should": Strongly recommended for best practices
- "recommended": Suggested for optimal usage

Use these guidelines when generating HTML, validating usage, or providing recommendations.
`, elementCount)
}

// extractGeneralGuidelines provides general usage guidelines that apply to all elements
func (s *Server) extractGeneralGuidelines() []Guideline {
	return []Guideline{
		{
			Text:     "Always provide meaningful descriptions for custom elements",
			Source:   "Design System Standards",
			Type:     "documentation",
			Severity: "should",
			Examples: []string{
				`<my-button>Save Changes</my-button>`,
				`<my-input aria-label="Email address"></my-input>`,
			},
		},
		{
			Text:     "Use semantic slot names that describe content purpose",
			Source:   "Web Components Best Practices",
			Type:     "slots",
			Severity: "should",
			Examples: []string{
				`<my-card><h2 slot="header">Title</h2><p>Content</p></my-card>`,
			},
		},
		{
			Text:     "Validate attribute values against manifest definitions",
			Source:   "Type Safety Guidelines",
			Type:     "attributes",
			Severity: "must",
			Examples: []string{
				`<my-button variant="primary">Valid</my-button>`,
				`<!-- Don't use: <my-button variant="invalid"> -->`,
			},
		},
		{
			Text:     "Include accessibility attributes for interactive elements",
			Source:   "WCAG 2.1 Guidelines",
			Type:     "accessibility",
			Context:  "interactive",
			Severity: "must",
			Examples: []string{
				`<my-button role="button" aria-label="Close dialog">×</my-button>`,
				`<my-input aria-describedby="help-text" aria-required="true">`,
			},
		},
	}
}

// extractElementGuidelines extracts guidelines from element descriptions and metadata
func (s *Server) extractElementGuidelines(element *ElementInfo) []Guideline {
	var guidelines []Guideline
	
	description := element.Description
	if description == "" {
		return guidelines
	}
	
	// Look for guideline keywords in descriptions
	sentences := s.splitIntoSentences(description)
	
	for _, sentence := range sentences {
		severity := s.determineSeverity(sentence)
		if severity != "" {
			guideline := Guideline{
				Text:     sentence,
				Source:   fmt.Sprintf("Element: %s", element.TagName),
				Type:     "usage",
				Severity: severity,
			}
			
			// Add context based on element type
			if s.isFormElement(element) {
				guideline.Context = "form"
			} else if s.isLayoutElement(element) {
				guideline.Context = "layout"
			}
			
			guidelines = append(guidelines, guideline)
		}
	}
	
	return guidelines
}

// extractAttributeGuidelines extracts guidelines from attribute descriptions
func (s *Server) extractAttributeGuidelines(attr Attribute, elementName string) []Guideline {
	var guidelines []Guideline
	
	description := attr.Description()
	if description == "" {
		return guidelines
	}
	
	sentences := s.splitIntoSentences(description)
	
	for _, sentence := range sentences {
		severity := s.determineSeverity(sentence)
		if severity != "" {
			guideline := Guideline{
				Text:     sentence,
				Source:   fmt.Sprintf("Attribute: %s.%s", elementName, attr.Name()),
				Type:     "attribute",
				Severity: severity,
			}
			
			// Add examples for enum attributes
			if values := attr.Values(); len(values) > 0 {
				for _, value := range values {
					example := fmt.Sprintf(`%s="%s"`, attr.Name(), value)
					guideline.Examples = append(guideline.Examples, example)
				}
			}
			
			guidelines = append(guidelines, guideline)
		}
	}
	
	// Add required attribute guideline
	if attr.Required() {
		guidelines = append(guidelines, Guideline{
			Text:     fmt.Sprintf("The %s attribute is required", attr.Name()),
			Source:   fmt.Sprintf("Attribute: %s.%s", elementName, attr.Name()),
			Type:     "attribute",
			Severity: "must",
		})
	}
	
	return guidelines
}

// generateAccessibilityGuidelines creates accessibility-focused guidelines
func (s *Server) generateAccessibilityGuidelines() []Guideline {
	return []Guideline{
		{
			Text:     "Provide alternative text for images and icons",
			Source:   "WCAG 2.1 - 1.1.1 Non-text Content",
			Type:     "accessibility",
			Context:  "media",
			Severity: "must",
			Examples: []string{
				`<my-icon alt="Settings" role="img"></my-icon>`,
				`<my-avatar alt="User profile picture"></my-avatar>`,
			},
		},
		{
			Text:     "Ensure sufficient color contrast (4.5:1 for normal text)",
			Source:   "WCAG 2.1 - 1.4.3 Contrast",
			Type:     "accessibility",
			Context:  "styling",
			Severity: "must",
			Examples: []string{
				`/* Good: High contrast */`,
				`--text-color: #000000; --background-color: #ffffff;`,
			},
		},
		{
			Text:     "Make interactive elements keyboard accessible",
			Source:   "WCAG 2.1 - 2.1.1 Keyboard",
			Type:     "accessibility",
			Context:  "interactive",
			Severity: "must",
			Examples: []string{
				`<my-button tabindex="0" role="button">Accessible Button</my-button>`,
				`<!-- Ensure Enter and Space keys work -->`,
			},
		},
		{
			Text:     "Use semantic markup and ARIA labels appropriately",
			Source:   "WCAG 2.1 - 4.1.2 Name, Role, Value",
			Type:     "accessibility",
			Context:  "semantics",
			Severity: "should",
			Examples: []string{
				`<my-dialog role="dialog" aria-labelledby="dialog-title">`,
				`<my-nav role="navigation" aria-label="Main navigation">`,
			},
		},
		{
			Text:     "Provide clear focus indicators for interactive elements",
			Source:   "WCAG 2.1 - 2.4.7 Focus Visible",
			Type:     "accessibility",
			Context:  "styling",
			Severity: "must",
			Examples: []string{
				`my-button:focus { outline: 2px solid #007acc; }`,
				`my-input:focus { box-shadow: 0 0 0 2px #007acc; }`,
			},
		},
	}
}

// generateCssGuidelines creates CSS integration guidelines
func (s *Server) generateCssGuidelines() []Guideline {
	return []Guideline{
		{
			Text:     "Use CSS custom properties for consistent theming",
			Source:   "CSS Architecture Best Practices",
			Type:     "css",
			Context:  "theming",
			Severity: "should",
			Examples: []string{
				`:root { --primary-color: #007acc; }`,
				`my-button { --button-color: var(--primary-color); }`,
			},
		},
		{
			Text:     "Style CSS parts instead of internal elements when available",
			Source:   "Web Components Styling Guidelines",
			Type:     "css",
			Context:  "styling",
			Severity: "should",
			Examples: []string{
				`/* Good: Using CSS parts */`,
				`my-card::part(header) { color: blue; }`,
				`/* Avoid: Styling internal elements directly */`,
			},
		},
		{
			Text:     "Respect component design system and styling API",
			Source:   "Design System Guidelines",
			Type:     "css",
			Context:  "architecture",
			Severity: "must",
			Examples: []string{
				`/* Use provided custom properties */`,
				`my-element { --element-spacing: 16px; }`,
				`/* Don't override internal styles directly */`,
			},
		},
		{
			Text:     "Test styling with different themes and accessibility settings",
			Source:   "Inclusive Design Practices",
			Type:     "css",
			Context:  "testing",
			Severity: "should",
			Examples: []string{
				`@media (prefers-color-scheme: dark) { /* dark theme */ }`,
				`@media (prefers-reduced-motion: reduce) { /* respect motion preferences */ }`,
			},
		},
	}
}

// generateBestPractices creates design system best practices
func (s *Server) generateBestPractices() []BestPractice {
	return []BestPractice{
		{
			Title:       "Semantic HTML Structure",
			Description: "Use semantic HTML elements and proper element nesting for better accessibility and SEO",
			Category:    "accessibility",
			Examples: []string{
				`<main><my-content-area>Main content</my-content-area></main>`,
				`<nav><my-navigation>Menu items</my-navigation></nav>`,
				`<article><my-blog-post>Post content</my-blog-post></article>`,
			},
			DoNot: []string{
				`<div><my-content>All content in divs</my-content></div>`,
				`<!-- Missing semantic structure -->`,
			},
		},
		{
			Title:       "Progressive Enhancement",
			Description: "Ensure custom elements work with JavaScript disabled and provide fallback content",
			Category:    "architecture",
			Examples: []string{
				`<my-enhanced-button>`,
				`  <button type="submit">Fallback Button</button>`,
				`</my-enhanced-button>`,
			},
			DoNot: []string{
				`<my-button></my-button> <!-- No fallback content -->`,
			},
		},
		{
			Title:       "Consistent Naming Conventions",
			Description: "Use consistent, descriptive names for elements, attributes, and CSS properties",
			Category:    "naming",
			Examples: []string{
				`<design-system-button variant="primary">Good naming</design-system-button>`,
				`--button-primary-color: #007acc; /* Descriptive CSS property */`,
			},
			DoNot: []string{
				`<btn v="p">Bad naming</btn>`,
				`--c1: #007acc; /* Non-descriptive property */`,
			},
		},
		{
			Title:       "Responsive Design Integration",
			Description: "Design components that work across different screen sizes and devices",
			Category:    "responsive",
			Examples: []string{
				`@container (max-width: 480px) { my-card { --card-layout: stacked; } }`,
				`my-grid { --columns: clamp(1, calc(100vw / 300px), 4); }`,
			},
		},
		{
			Title:       "Performance Optimization",
			Description: "Optimize component loading and rendering for better user experience",
			Category:    "performance",
			Examples: []string{
				`<my-image loading="lazy" src="image.jpg">`,
				`<!-- Use appropriate loading strategies -->`,
			},
		},
	}
}

// Helper methods

// splitIntoSentences splits text into individual sentences
func (s *Server) splitIntoSentences(text string) []string {
	// Simple sentence splitting - could be enhanced with more sophisticated logic
	sentences := strings.Split(text, ". ")
	var result []string
	
	for _, sentence := range sentences {
		sentence = strings.TrimSpace(sentence)
		if sentence != "" {
			// Ensure sentence ends with period
			if !strings.HasSuffix(sentence, ".") && !strings.HasSuffix(sentence, "!") && !strings.HasSuffix(sentence, "?") {
				sentence += "."
			}
			result = append(result, sentence)
		}
	}
	
	return result
}

// determineSeverity determines the severity level of a guideline based on keywords
func (s *Server) determineSeverity(text string) string {
	lowerText := strings.ToLower(text)
	
	// Must/Required keywords
	if strings.Contains(lowerText, "must") || 
	   strings.Contains(lowerText, "required") || 
	   strings.Contains(lowerText, "essential") ||
	   strings.Contains(lowerText, "mandatory") {
		return "must"
	}
	
	// Should/Recommended keywords
	if strings.Contains(lowerText, "should") || 
	   strings.Contains(lowerText, "recommended") || 
	   strings.Contains(lowerText, "best practice") ||
	   strings.Contains(lowerText, "prefer") {
		return "should"
	}
	
	// Suggestion keywords
	if strings.Contains(lowerText, "consider") || 
	   strings.Contains(lowerText, "may") || 
	   strings.Contains(lowerText, "can") ||
	   strings.Contains(lowerText, "optional") {
		return "recommended"
	}
	
	return ""
}