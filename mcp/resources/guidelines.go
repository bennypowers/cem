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
package resources

import (
	"context"
	"fmt"
	"strings"

	"bennypowers.dev/cem/mcp/templates"
	"bennypowers.dev/cem/mcp/types"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// handleGuidelinesResource provides design system guidelines and best practices
func handleGuidelinesResource(ctx context.Context, req *mcp.ReadResourceRequest, registry types.MCPContext) (*mcp.ReadResourceResult, error) {
	// Collect guidelines data from the registry
	guidelinesData := collectGuidelinesData(registry)

	// Render comprehensive guidelines using templates
	var response strings.Builder

	// Render each section using the shared template system
	overview, err := templates.RenderTemplate("guidelines_overview", guidelinesData)
	if err != nil {
		return nil, fmt.Errorf("failed to render guidelines overview: %w", err)
	}
	response.WriteString(overview)
	response.WriteString("\n\n")

	elementGuidelines, err := templates.RenderTemplate("element_guidelines", guidelinesData)
	if err != nil {
		return nil, fmt.Errorf("failed to render element guidelines: %w", err)
	}
	response.WriteString(elementGuidelines)
	response.WriteString("\n\n")

	accessibilityGuidelines, err := templates.RenderTemplate("accessibility_guidelines", guidelinesData)
	if err != nil {
		return nil, fmt.Errorf("failed to render accessibility guidelines: %w", err)
	}
	response.WriteString(accessibilityGuidelines)
	response.WriteString("\n\n")

	cssGuidelines, err := templates.RenderTemplate("css_guidelines", guidelinesData)
	if err != nil {
		return nil, fmt.Errorf("failed to render CSS guidelines: %w", err)
	}
	response.WriteString(cssGuidelines)
	response.WriteString("\n\n")

	patternGuidelines, err := templates.RenderTemplate("pattern_guidelines", guidelinesData)
	if err != nil {
		return nil, fmt.Errorf("failed to render pattern guidelines: %w", err)
	}
	response.WriteString(patternGuidelines)

	return &mcp.ReadResourceResult{
		Contents: []*mcp.ResourceContents{{
			URI:      req.Params.URI,
			MIMEType: "text/markdown",
			Text:     response.String(),
		}},
	}, nil
}

// collectGuidelinesData gathers all guidelines information into structured data
func collectGuidelinesData(registry types.MCPContext) GuidelinesData {
	elementMap := registry.AllElements()

	// Convert map to slice for processing
	elements := make([]types.ElementInfo, 0, len(elementMap))
	for _, element := range elementMap {
		elements = append(elements, element)
	}

	return GuidelinesData{
		Overview:                generateOverview(len(elements)),
		Principles:              generatePrinciples(),
		Philosophy:              generatePhilosophy(),
		GeneralGuidelines:       generateGeneralGuidelines(),
		ElementGuidelines:       generateElementGuidelines(elements),
		NamingGuidelines:        generateNamingGuidelines(registry),
		AccessibilityGuidelines: generateAccessibilityGuidelines(),
		ThemingGuidelines:       generateThemingGuidelines(registry),
		PerformanceGuidelines:   generatePerformanceGuidelines(),
		CompositionGuidelines:   generateCompositionGuidelines(),
		IntegrationGuidelines:   generateIntegrationGuidelines(),
		TestingGuidelines:       generateTestingGuidelines(),
		AntiPatterns:            generateAntiPatterns(),
		BestPractices:           generateBestPractices(),
		FrameworkGuidelines:     generateFrameworkGuidelines(),
		LayoutPatterns:          generateLayoutPatterns(),
		DataPatterns:            generateDataPatterns(),
		FormPatterns:            generateFormPatterns(),
		NavigationPatterns:      generateNavigationPatterns(),
		FeedbackPatterns:        generateFeedbackPatterns(),
		PatternsByCategory:      generatePatternsByCategory(),
		ProgressiveEnhancement:  generateProgressiveEnhancement(),
		ErrorHandlingPatterns:   generateErrorHandlingPatterns(),
		SecurityGuidelines:      generateSecurityGuidelines(),
		ColorGuidelines:         generateColorGuidelines(),
		ResponsiveGuidelines:    generateResponsiveGuidelines(),
		CSSBestPractices:        generateCSSBestPractices(),
	}
}

// generateOverview creates the overview description
func generateOverview(elementCount int) string {
	return fmt.Sprintf("Comprehensive design system guidelines extracted from %d custom elements and their manifests. These guidelines ensure consistent, accessible, and maintainable component usage across your application.", elementCount)
}

// generatePrinciples creates core design principles
func generatePrinciples() []string {
	return []string{
		"Consistency: Maintain consistent patterns across all components",
		"Accessibility: Ensure all components are accessible by default",
		"Performance: Optimize for fast loading and smooth interactions",
		"Flexibility: Design components to work in various contexts",
		"Semantic: Use meaningful element and attribute names",
	}
}

// generatePhilosophy creates design philosophy mapping
func generatePhilosophy() map[string]string {
	return map[string]string{
		"componentFirst": "Design with components as the primary building blocks",
		"progressive":    "Support progressive enhancement and graceful degradation",
		"inclusive":      "Design for all users, abilities, and devices",
		"maintainable":   "Write code that is easy to understand and modify",
	}
}

// generateGeneralGuidelines creates general guideline categories
func generateGeneralGuidelines() []GuidelineCategory {
	return []GuidelineCategory{
		{
			Category: "semantic",
			Guidelines: []string{
				"Use semantic HTML within slots whenever possible",
				"Provide meaningful content for screen readers",
				"Ensure keyboard navigation works properly",
				"Test with various content lengths and types",
			},
		},
		{
			Category: "documentation",
			Guidelines: []string{
				"Always provide meaningful descriptions for custom elements",
				"Use semantic slot names that describe content purpose",
				"Include accessibility attributes for interactive elements",
			},
		},
	}
}

// generateElementGuidelines creates element-specific guidelines
func generateElementGuidelines(elements []types.ElementInfo) map[string]ElementGuideline {
	elementGuidelines := make(map[string]ElementGuideline)

	for _, element := range elements {
		// Create attribute guidelines
		var attrGuidelines []AttributeGuideline
		for _, attr := range element.Attributes() {
			guidance := []string{}
			if attr.Required() {
				guidance = append(guidance, "This attribute is required")
			}
			if attr.Default() != "" {
				guidance = append(guidance, fmt.Sprintf("Default value: %s", attr.Default()))
			}
			if len(attr.Values()) > 0 {
				guidance = append(guidance, fmt.Sprintf("Valid values: %s", strings.Join(attr.Values(), ", ")))
			}

			attrGuidelines = append(attrGuidelines, AttributeGuideline{
				Attribute: attr,
				Guidance:  guidance,
			})
		}

		// Create slot guidelines
		var slotGuidelines []SlotGuideline
		for _, slot := range element.Slots() {
			slotGuidelines = append(slotGuidelines, SlotGuideline{
				Slot: slot,
				Guidance: []string{
					"Use semantic HTML elements when possible",
					"Consider accessibility when adding content",
				},
			})
		}

		// Generate usage examples
		examples := generateUsageExamples(element)

		elementGuidelines[element.TagName()] = ElementGuideline{
			ElementInfo: element,
			Guidelines:  []string{"Follow manifest-defined constraints and patterns"},
			Examples:    examples,
			Attributes:  attrGuidelines,
			Slots:       slotGuidelines,
		}
	}

	return elementGuidelines
}

// generateNamingGuidelines creates naming convention guidelines
func generateNamingGuidelines(registry types.MCPContext) NamingGuidelines {
	prefixes := registry.CommonPrefixes()

	return NamingGuidelines{
		Elements: NamingCategory{
			Format:  "kebab-case with meaningful prefixes",
			Pattern: "prefix-component-variant",
			Guidelines: []string{
				"Use descriptive names that indicate purpose",
				"Include namespace/brand prefix to avoid conflicts",
				"Use consistent naming patterns across related components",
				"Avoid abbreviations unless they're widely understood",
			},
			Examples: generateNamingExamples(prefixes),
		},
		Attributes: NamingCategory{
			Format: "kebab-case",
			Guidelines: []string{
				"Match HTML conventions where possible",
				"Use ARIA patterns for accessibility attributes",
				"Prefix custom attributes to avoid conflicts",
				"Be consistent with similar attributes across components",
			},
		},
	}
}

// generateAccessibilityGuidelines creates accessibility guidelines
func generateAccessibilityGuidelines() AccessibilityGuidelines {
	return AccessibilityGuidelines{
		WCAG: WCAGGuidelines{
			Level: "AA compliance required",
			Principles: []string{
				"Perceivable: Provide text alternatives and sufficient contrast",
				"Operable: Ensure keyboard accessibility and reasonable time limits",
				"Understandable: Make text readable and predictable",
				"Robust: Maximize compatibility with assistive technologies",
			},
		},
		ARIA: map[string]string{
			"roles":      "Use appropriate ARIA roles for custom semantics",
			"properties": "Implement ARIA properties for relationships",
			"states":     "Manage ARIA states for dynamic content",
			"labels":     "Provide accessible names and descriptions",
		},
		Keyboard: map[string]string{
			"navigation": "Support standard keyboard navigation patterns",
			"focus":      "Implement visible focus indicators",
			"shortcuts":  "Provide consistent keyboard shortcuts",
			"trapping":   "Implement focus trapping for modals",
		},
		Testing: []string{
			"Test with screen readers (NVDA, JAWS, VoiceOver)",
			"Verify keyboard-only navigation",
			"Check color contrast ratios",
			"Validate with automated accessibility tools",
		},
	}
}

// generateThemingGuidelines creates CSS theming guidelines
func generateThemingGuidelines(registry types.MCPContext) ThemingGuidelines {
	cssProperties := registry.AllCSSProperties()

	return ThemingGuidelines{
		Tokens: map[string]string{
			"colors":      "Use semantic color tokens (primary, secondary, etc.)",
			"typography":  "Implement consistent typography scales",
			"spacing":     "Use consistent spacing scales",
			"breakpoints": "Define standard responsive breakpoints",
		},
		CustomProperties: map[string]interface{}{
			"naming":    "Use consistent naming conventions for CSS custom properties",
			"scoping":   "Scope properties appropriately (global vs component)",
			"fallbacks": "Provide fallback values for older browsers",
			"available": cssProperties,
		},
		Parts: map[string]string{
			"usage":     "Use CSS parts for styling internal component structure",
			"naming":    "Use descriptive names for CSS parts",
			"stability": "Consider API stability when exposing parts",
		},
		States: map[string]string{
			"custom":      "Use CSS custom states for component state styling",
			"consistency": "Maintain consistent state naming across components",
		},
	}
}

// generatePerformanceGuidelines creates performance guidelines
func generatePerformanceGuidelines() PerformanceGuidelines {
	return PerformanceGuidelines{
		Loading: map[string]string{
			"lazy":       "Implement lazy loading for non-critical components",
			"bundling":   "Bundle related components together",
			"splitting":  "Use code splitting for large component libraries",
			"preloading": "Preload critical components",
		},
		Rendering: map[string]string{
			"virtualization": "Use virtual scrolling for large lists",
			"memoization":    "Cache expensive calculations",
			"batching":       "Batch DOM updates when possible",
			"animations":     "Use CSS animations and transforms",
		},
		Memory: map[string]string{
			"cleanup":    "Clean up event listeners and timers",
			"references": "Avoid memory leaks with proper cleanup",
			"observers":  "Disconnect observers when not needed",
		},
	}
}

// generateCompositionGuidelines creates composition guidelines
func generateCompositionGuidelines() CompositionGuidelines {
	return CompositionGuidelines{
		Slots: map[string]string{
			"design":     "Design flexible slot APIs for maximum reusability",
			"naming":     "Use descriptive slot names",
			"defaults":   "Provide sensible default content",
			"validation": "Consider content validation for slots",
		},
		Nesting: map[string]string{
			"hierarchy":   "Maintain semantic hierarchy in nested components",
			"context":     "Consider how components interact when nested",
			"performance": "Be mindful of nested component performance",
		},
		Communication: map[string]string{
			"events":     "Use custom events for component communication",
			"properties": "Design clear property APIs",
			"context":    "Consider context providers for shared state",
		},
	}
}

// generateIntegrationGuidelines creates integration guidelines
func generateIntegrationGuidelines() IntegrationGuidelines {
	return IntegrationGuidelines{
		Frameworks: map[string]string{
			"vanilla": "Import and use directly in HTML",
			"lit":     "Use within Lit templates with proper property binding",
			"react":   "Import and use as JSX elements",
			"vue":     "Register as custom elements for use in templates",
			"angular": "Use CUSTOM_ELEMENTS_SCHEMA for custom elements",
		},
		Bundling: map[string]string{
			"strategy":     "Bundle related components together",
			"splitting":    "Use dynamic imports for lazy loading",
			"optimization": "Tree-shake unused components",
		},
		Imports: map[string]string{
			"es6":      "Use ES6 import syntax when possible",
			"cdn":      "Consider CDN imports for simpler projects",
			"bundling": "Import from package entry points",
		},
	}
}

// generateTestingGuidelines creates testing guidelines
func generateTestingGuidelines() TestingGuidelines {
	return TestingGuidelines{
		Types: map[string]string{
			"unit":          "Test individual component functionality",
			"integration":   "Test component interactions",
			"accessibility": "Test with accessibility tools",
			"visual":        "Test visual appearance across browsers",
			"performance":   "Test loading and runtime performance",
		},
		Tools: []string{
			"Web Test Runner for unit testing",
			"Playwright for integration testing",
			"axe-core for accessibility testing",
			"Lighthouse for performance testing",
		},
		Practices: []string{
			"Test component APIs thoroughly",
			"Include accessibility in all tests",
			"Test keyboard navigation",
			"Test with real content and edge cases",
			"Automate visual regression testing",
		},
	}
}

// generateAntiPatterns creates anti-pattern guidelines
func generateAntiPatterns() []string {
	return []string{
		"Breaking semantic HTML structure",
		"Ignoring accessibility requirements",
		"Creating overly complex component APIs",
		"Mixing presentation and business logic",
		"Using non-standard naming conventions",
		"Forgetting keyboard navigation",
		"Hardcoding styles without theming support",
		"Creating components that can't be styled",
		"Ignoring performance implications",
	}
}

// generateBestPractices creates best practice guidelines
func generateBestPractices() []string {
	return []string{
		"Always use appropriate semantic HTML",
		"Build accessibility in from the start",
		"Keep component APIs simple and focused",
		"Separate concerns clearly",
		"Follow established naming conventions",
		"Always consider keyboard users",
		"Support theming and customization",
		"Provide multiple styling options",
		"Consider performance from the beginning",
	}
}

// generateFrameworkGuidelines creates framework integration guidelines
func generateFrameworkGuidelines() map[string]string {
	return map[string]string{
		"vanilla": "Import and use directly in HTML",
		"lit":     "Use within Lit templates with proper property binding",
		"react":   "Import and use as JSX elements",
		"vue":     "Register as custom elements for use in templates",
		"angular": "Use CUSTOM_ELEMENTS_SCHEMA for custom elements",
	}
}

// generateLayoutPatterns creates layout pattern guidelines
func generateLayoutPatterns() LayoutPatterns {
	return LayoutPatterns{
		Containers: []string{
			"Use layout components for consistent spacing",
			"Implement responsive design patterns",
			"Consider container queries for component-based responsiveness",
		},
		Grid: []string{
			"Use CSS Grid for two-dimensional layouts",
			"Implement consistent grid systems",
			"Support both fixed and flexible layouts",
		},
		Flexbox: []string{
			"Use Flexbox for one-dimensional layouts",
			"Implement consistent alignment patterns",
			"Support various content distributions",
		},
	}
}

// generateDataPatterns creates data pattern guidelines
func generateDataPatterns() DataPatterns {
	return DataPatterns{
		Display: []string{
			"Use appropriate components for data types",
			"Implement consistent formatting",
			"Support different data densities",
		},
		Interaction: []string{
			"Provide sorting and filtering options",
			"Implement pagination for large datasets",
			"Support data export when appropriate",
		},
	}
}

// generateFormPatterns creates form pattern guidelines
func generateFormPatterns() FormPatterns {
	return FormPatterns{
		Validation: []string{
			"Provide real-time validation feedback",
			"Use appropriate input types",
			"Include clear error messages",
		},
		Labeling: []string{
			"Associate labels with form controls",
			"Use placeholder text appropriately",
			"Provide help text when needed",
		},
		Grouping: []string{
			"Group related form fields",
			"Use fieldsets for complex forms",
			"Implement logical tab order",
		},
	}
}

// generateNavigationPatterns creates navigation pattern guidelines
func generateNavigationPatterns() NavigationPatterns {
	return NavigationPatterns{
		Structure: []string{
			"Use consistent navigation patterns",
			"Implement breadcrumbs for deep navigation",
			"Provide clear navigation hierarchy",
		},
		Interaction: []string{
			"Support keyboard navigation",
			"Indicate current location",
			"Use appropriate focus management",
		},
	}
}

// generateFeedbackPatterns creates feedback pattern guidelines
func generateFeedbackPatterns() FeedbackPatterns {
	return FeedbackPatterns{
		Types: []string{
			"Success messages for completed actions",
			"Error messages for failed actions",
			"Warning messages for important information",
			"Info messages for general guidance",
		},
		Delivery: []string{
			"Use appropriate timing for messages",
			"Provide multiple ways to access feedback",
			"Ensure messages are accessible",
		},
	}
}

// generatePatternsByCategory creates categorized patterns
func generatePatternsByCategory() map[string][]string {
	return map[string][]string{
		"semantic": {
			"Use appropriate semantic elements",
			"Maintain document outline hierarchy",
			"Provide meaningful content structure",
		},
		"interactive": {
			"Ensure keyboard accessibility",
			"Provide clear focus indicators",
			"Support assistive technologies",
		},
		"layout": {
			"Use consistent spacing patterns",
			"Implement responsive design",
			"Consider container queries",
		},
	}
}

// generateProgressiveEnhancement creates progressive enhancement guidelines
func generateProgressiveEnhancement() []string {
	return []string{
		"Ensure functionality works without JavaScript",
		"Provide fallback content for custom elements",
		"Layer enhancements progressively",
		"Test with various levels of support",
	}
}

// generateErrorHandlingPatterns creates error handling guidelines
func generateErrorHandlingPatterns() []string {
	return []string{
		"Provide clear error messages",
		"Implement graceful degradation",
		"Handle network failures appropriately",
		"Log errors for debugging purposes",
	}
}

// generateSecurityGuidelines creates security guidelines
func generateSecurityGuidelines() []string {
	return []string{
		"Sanitize user input appropriately",
		"Validate data before processing",
		"Follow secure coding practices",
		"Consider XSS prevention measures",
	}
}

// generateColorGuidelines creates color guidelines
func generateColorGuidelines() map[string]string {
	return map[string]string{
		"contrast":      "Ensure sufficient color contrast (4.5:1 for normal text)",
		"independence":  "Don't rely on color alone to convey information",
		"theming":       "Support light and dark themes",
		"accessibility": "Consider color blindness and visual impairments",
	}
}

// generateResponsiveGuidelines creates responsive design guidelines
func generateResponsiveGuidelines() map[string]string {
	return map[string]string{
		"mobile":      "Design mobile-first responsive layouts",
		"breakpoints": "Use consistent breakpoint systems",
		"containers":  "Consider container queries for components",
		"flexibility": "Design for various screen sizes and orientations",
	}
}

// generateCSSBestPractices creates CSS best practices
func generateCSSBestPractices() []string {
	return []string{
		"Use CSS custom properties for theming",
		"Implement consistent naming conventions",
		"Minimize specificity conflicts",
		"Use logical properties for internationalization",
		"Consider CSS containment for performance",
	}
}

// Helper functions that are still needed

func generateNamingExamples(prefixes []string) []string {
	examples := []string{
		"my-button",
		"my-card",
		"my-dialog",
	}

	if len(prefixes) > 0 {
		prefix := prefixes[0]
		examples = []string{
			prefix + "-button",
			prefix + "-card",
			prefix + "-dialog",
		}
	}

	return examples
}

func generateUsageExamples(element types.ElementInfo) []string {
	tagName := element.TagName()

	examples := []string{
		fmt.Sprintf("<%s></%s>", tagName, tagName),
	}

	// Add examples with attributes
	if attrs := element.Attributes(); len(attrs) > 0 {
		attrExample := fmt.Sprintf("<%s %s=\"example\"></%s>", tagName, attrs[0].Name(), tagName)
		examples = append(examples, attrExample)
	}

	// Add examples with slots
	if slots := element.Slots(); len(slots) > 0 {
		slotName := slots[0].Name()
		if slotName == "" {
			slotName = "default"
		}

		var slotExample string
		if slotName == "default" {
			slotExample = fmt.Sprintf("<%s>Content</%s>", tagName, tagName)
		} else {
			slotExample = fmt.Sprintf("<%s><span slot=\"%s\">Content</span></%s>", tagName, slotName, tagName)
		}
		examples = append(examples, slotExample)
	}

	return examples
}
