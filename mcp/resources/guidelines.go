package resources

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"bennypowers.dev/cem/mcp/types"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// handleGuidelinesResource provides design system guidelines and best practices
func handleGuidelinesResource(ctx context.Context, req *mcp.ReadResourceRequest, registry types.Registry) (*mcp.ReadResourceResult, error) {
	// Generate comprehensive guidelines from the registry
	guidelines := generateDesignSystemGuidelines(registry)

	contents, err := json.MarshalIndent(guidelines, "", "  ")
	if err != nil {
		return nil, err
	}

	return &mcp.ReadResourceResult{
		Contents: []*mcp.ResourceContents{{
			URI:      req.Params.URI,
			MIMEType: "application/json",
			Text:     string(contents),
		}},
	}, nil
}

// generateDesignSystemGuidelines creates comprehensive design system guidelines
func generateDesignSystemGuidelines(registry types.Registry) map[string]interface{} {
	elementMap := registry.AllElements()

	// Convert map to slice
	elements := make([]types.ElementInfo, 0, len(elementMap))
	for _, element := range elementMap {
		elements = append(elements, element)
	}

	guidelines := map[string]interface{}{
		"overview":      generateOverviewGuidelines(),
		"naming":        generateNamingGuidelines(elements),
		"usage":         generateUsageGuidelines(elements),
		"patterns":      generatePatternGuidelines(elements),
		"accessibility": generateAccessibilityGuidelines(elements),
		"performance":   generatePerformanceGuidelines(),
		"theming":       generateThemingGuidelines(elements),
		"composition":   generateCompositionGuidelines(elements),
		"antiPatterns":  generateAntiPatterns(),
		"testing":       generateTestingGuidelines(),
	}

	return guidelines
}

// generateOverviewGuidelines creates high-level design system guidance
func generateOverviewGuidelines() map[string]interface{} {
	return map[string]interface{}{
		"principles": []string{
			"Consistency: Maintain consistent patterns across all components",
			"Accessibility: Ensure all components are accessible by default",
			"Performance: Optimize for fast loading and smooth interactions",
			"Flexibility: Design components to work in various contexts",
			"Semantic: Use meaningful element and attribute names",
		},
		"philosophy": map[string]interface{}{
			"componentFirst": "Design with components as the primary building blocks",
			"progressive":    "Support progressive enhancement and graceful degradation",
			"inclusive":      "Design for all users, abilities, and devices",
			"maintainable":   "Write code that is easy to understand and modify",
		},
	}
}

// generateNamingGuidelines creates guidelines for element and attribute naming
func generateNamingGuidelines(elements []types.ElementInfo) map[string]interface{} {
	// Analyze existing elements to extract naming patterns
	prefixes := extractCommonPrefixes(elements)

	return map[string]interface{}{
		"elements": map[string]interface{}{
			"format":   "kebab-case with meaningful prefixes",
			"pattern":  "prefix-component-variant",
			"examples": generateNamingExamples(prefixes),
			"prefixes": prefixes,
			"guidelines": []string{
				"Use descriptive names that indicate purpose",
				"Include namespace/brand prefix to avoid conflicts",
				"Use consistent naming patterns across related components",
				"Avoid abbreviations unless they're widely understood",
			},
		},
		"attributes": map[string]interface{}{
			"format": "kebab-case",
			"types": map[string]interface{}{
				"boolean": "Use boolean attributes for true/false values",
				"string":  "Use descriptive names for string attributes",
				"number":  "Include units in name when applicable (e.g., max-width)",
				"enum":    "Use clear, descriptive enum values",
			},
			"guidelines": []string{
				"Match HTML conventions where possible",
				"Use ARIA patterns for accessibility attributes",
				"Prefix custom attributes to avoid conflicts",
				"Be consistent with similar attributes across components",
			},
		},
	}
}

// generateUsageGuidelines creates component usage guidelines
func generateUsageGuidelines(elements []types.ElementInfo) map[string]interface{} {
	elementGuidelines := make(map[string]interface{})

	for _, element := range elements {
		elementGuidelines[element.TagName()] = generateElementUsageGuidelines(element)
	}

	return map[string]interface{}{
		"general": map[string]interface{}{
			"principles": []string{
				"Use semantic HTML within slots whenever possible",
				"Provide meaningful content for screen readers",
				"Ensure keyboard navigation works properly",
				"Test with various content lengths and types",
			},
			"patterns": []string{
				"Wrap components in semantic containers when needed",
				"Use proper heading hierarchy",
				"Provide skip links for complex interfaces",
				"Implement focus management for interactive components",
			},
		},
		"elements": elementGuidelines,
		"integration": map[string]interface{}{
			"frameworks": generateFrameworkGuidelines(),
			"bundling":   generateBundlingGuidelines(),
			"imports":    generateImportGuidelines(),
		},
	}
}

// generatePatternGuidelines creates common usage pattern guidelines
func generatePatternGuidelines(elements []types.ElementInfo) map[string]interface{} {
	return map[string]interface{}{
		"layout": map[string]interface{}{
			"containers": []string{
				"Use layout components for consistent spacing",
				"Implement responsive design patterns",
				"Consider container queries for component-based responsiveness",
			},
			"grid": []string{
				"Use CSS Grid for two-dimensional layouts",
				"Implement consistent grid systems",
				"Support both fixed and flexible layouts",
			},
			"flexbox": []string{
				"Use Flexbox for one-dimensional layouts",
				"Implement consistent alignment patterns",
				"Support various content distributions",
			},
		},
		"forms":      generateFormPatterns(elements),
		"navigation": generateNavigationPatterns(elements),
		"feedback":   generateFeedbackPatterns(elements),
		"data":       generateDataPatterns(elements),
	}
}

// generateAccessibilityGuidelines creates accessibility-focused guidelines
func generateAccessibilityGuidelines(elements []types.ElementInfo) map[string]interface{} {
	return map[string]interface{}{
		"wcag": map[string]interface{}{
			"level": "AA compliance required",
			"principles": []string{
				"Perceivable: Provide text alternatives and sufficient contrast",
				"Operable: Ensure keyboard accessibility and reasonable time limits",
				"Understandable: Make text readable and predictable",
				"Robust: Maximize compatibility with assistive technologies",
			},
		},
		"aria": map[string]interface{}{
			"roles":      "Use appropriate ARIA roles for custom semantics",
			"properties": "Implement ARIA properties for relationships",
			"states":     "Manage ARIA states for dynamic content",
			"labels":     "Provide accessible names and descriptions",
		},
		"keyboard": map[string]interface{}{
			"navigation": "Support standard keyboard navigation patterns",
			"focus":      "Implement visible focus indicators",
			"shortcuts":  "Provide consistent keyboard shortcuts",
			"trapping":   "Implement focus trapping for modals",
		},
		"testing": []string{
			"Test with screen readers (NVDA, JAWS, VoiceOver)",
			"Verify keyboard-only navigation",
			"Check color contrast ratios",
			"Validate with automated accessibility tools",
		},
	}
}

// generatePerformanceGuidelines creates performance optimization guidelines
func generatePerformanceGuidelines() map[string]interface{} {
	return map[string]interface{}{
		"loading": map[string]interface{}{
			"lazy":       "Implement lazy loading for non-critical components",
			"bundling":   "Bundle related components together",
			"splitting":  "Use code splitting for large component libraries",
			"preloading": "Preload critical components",
		},
		"rendering": map[string]interface{}{
			"virtualization": "Use virtual scrolling for large lists",
			"memoization":    "Cache expensive calculations",
			"batching":       "Batch DOM updates when possible",
			"animations":     "Use CSS animations and transforms",
		},
		"memory": map[string]interface{}{
			"cleanup":    "Clean up event listeners and timers",
			"references": "Avoid memory leaks with proper cleanup",
			"observers":  "Disconnect observers when not needed",
		},
	}
}

// generateThemingGuidelines creates theming and styling guidelines
func generateThemingGuidelines(elements []types.ElementInfo) map[string]interface{} {
	// Extract CSS properties from all elements
	cssProperties := extractAllCSSProperties(elements)

	return map[string]interface{}{
		"tokens": map[string]interface{}{
			"colors":      "Use semantic color tokens (primary, secondary, etc.)",
			"typography":  "Implement consistent typography scales",
			"spacing":     "Use consistent spacing scales",
			"breakpoints": "Define standard responsive breakpoints",
		},
		"customProperties": map[string]interface{}{
			"naming":    "Use consistent naming conventions for CSS custom properties",
			"scoping":   "Scope properties appropriately (global vs component)",
			"fallbacks": "Provide fallback values for older browsers",
			"available": cssProperties,
		},
		"parts": map[string]interface{}{
			"usage":     "Use CSS parts for styling internal component structure",
			"naming":    "Use descriptive names for CSS parts",
			"stability": "Consider API stability when exposing parts",
		},
		"states": map[string]interface{}{
			"custom":      "Use CSS custom states for component state styling",
			"consistency": "Maintain consistent state naming across components",
		},
	}
}

// generateCompositionGuidelines creates component composition guidelines
func generateCompositionGuidelines(elements []types.ElementInfo) map[string]interface{} {
	return map[string]interface{}{
		"slots": map[string]interface{}{
			"design":     "Design flexible slot APIs for maximum reusability",
			"naming":     "Use descriptive slot names",
			"defaults":   "Provide sensible default content",
			"validation": "Consider content validation for slots",
		},
		"nesting": map[string]interface{}{
			"hierarchy":   "Maintain semantic hierarchy in nested components",
			"context":     "Consider how components interact when nested",
			"performance": "Be mindful of nested component performance",
		},
		"communication": map[string]interface{}{
			"events":     "Use custom events for component communication",
			"properties": "Design clear property APIs",
			"context":    "Consider context providers for shared state",
		},
	}
}

// generateAntiPatterns creates guidelines about what to avoid
func generateAntiPatterns() map[string]interface{} {
	return map[string]interface{}{
		"avoid": []string{
			"Breaking semantic HTML structure",
			"Ignoring accessibility requirements",
			"Creating overly complex component APIs",
			"Mixing presentation and business logic",
			"Using non-standard naming conventions",
			"Forgetting keyboard navigation",
			"Hardcoding styles without theming support",
			"Creating components that can't be styled",
			"Ignoring performance implications",
		},
		"alternatives": map[string]interface{}{
			"semantics":   "Always use appropriate semantic HTML",
			"a11y":        "Build accessibility in from the start",
			"api":         "Keep component APIs simple and focused",
			"separation":  "Separate concerns clearly",
			"naming":      "Follow established naming conventions",
			"keyboard":    "Always consider keyboard users",
			"theming":     "Support theming and customization",
			"styling":     "Provide multiple styling options",
			"performance": "Consider performance from the beginning",
		},
	}
}

// generateTestingGuidelines creates component testing guidelines
func generateTestingGuidelines() map[string]interface{} {
	return map[string]interface{}{
		"types": map[string]interface{}{
			"unit":          "Test individual component functionality",
			"integration":   "Test component interactions",
			"accessibility": "Test with accessibility tools",
			"visual":        "Test visual appearance across browsers",
			"performance":   "Test loading and runtime performance",
		},
		"tools": []string{
			"Web Test Runner for unit testing",
			"Playwright for integration testing",
			"axe-core for accessibility testing",
			"Lighthouse for performance testing",
		},
		"practices": []string{
			"Test component APIs thoroughly",
			"Include accessibility in all tests",
			"Test keyboard navigation",
			"Test with real content and edge cases",
			"Automate visual regression testing",
		},
	}
}

// Helper functions for generating specific guidelines

func extractCommonPrefixes(elements []types.ElementInfo) []string {
	prefixCount := make(map[string]int)

	for _, element := range elements {
		tagName := element.TagName()
		if parts := strings.Split(tagName, "-"); len(parts) > 1 {
			prefix := parts[0]
			prefixCount[prefix]++
		}
	}

	// Return prefixes used by multiple elements
	var prefixes []string
	for prefix, count := range prefixCount {
		if count > 1 {
			prefixes = append(prefixes, prefix)
		}
	}

	return prefixes
}

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

func generateElementUsageGuidelines(element types.ElementInfo) map[string]interface{} {
	guidelines := map[string]interface{}{
		"description": element.Description(),
		"attributes":  generateAttributeGuidelines(element.Attributes()),
		"slots":       generateSlotUsageGuidelines(element.Slots()),
		"examples":    generateUsageExamples(element),
	}

	return guidelines
}

func generateAttributeGuidelines(attributes []types.Attribute) []map[string]interface{} {
	guidelines := make([]map[string]interface{}, len(attributes))

	for i, attr := range attributes {
		guidelines[i] = map[string]interface{}{
			"name":        attr.Name(),
			"description": attr.Description(),
			"type":        attr.Type(),
			"required":    attr.Required(),
			"guidance":    generateAttributeSpecificGuidance(attr),
		}
	}

	return guidelines
}

func generateSlotUsageGuidelines(slots []types.Slot) []map[string]interface{} {
	guidelines := make([]map[string]interface{}, len(slots))

	for i, slot := range slots {
		guidelines[i] = map[string]interface{}{
			"name":        slot.Name(),
			"description": slot.Description(),
			"guidance":    generateSlotSpecificGuidance(slot),
		}
	}

	return guidelines
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

func generateAttributeSpecificGuidance(attr types.Attribute) []string {
	guidance := []string{}

	if attr.Required() {
		guidance = append(guidance, "This attribute is required")
	}

	if attr.Default() != "" {
		guidance = append(guidance, fmt.Sprintf("Default value: %s", attr.Default()))
	}

	// Add type-specific guidance
	attrType := strings.ToLower(attr.Type())
	if strings.Contains(attrType, "boolean") {
		guidance = append(guidance, "Use as a boolean attribute (present = true, absent = false)")
	}

	return guidance
}

func generateSlotSpecificGuidance(slot types.Slot) []string {
	guidance := []string{
		"Use semantic HTML elements when possible",
		"Consider accessibility when adding content",
	}

	return guidance
}

func generateFormPatterns(elements []types.ElementInfo) map[string]interface{} {
	return map[string]interface{}{
		"validation": []string{
			"Provide real-time validation feedback",
			"Use appropriate input types",
			"Include clear error messages",
		},
		"labeling": []string{
			"Associate labels with form controls",
			"Use placeholder text appropriately",
			"Provide help text when needed",
		},
		"grouping": []string{
			"Group related form fields",
			"Use fieldsets for complex forms",
			"Implement logical tab order",
		},
	}
}

func generateNavigationPatterns(elements []types.ElementInfo) map[string]interface{} {
	return map[string]interface{}{
		"structure": []string{
			"Use consistent navigation patterns",
			"Implement breadcrumbs for deep navigation",
			"Provide clear navigation hierarchy",
		},
		"interaction": []string{
			"Support keyboard navigation",
			"Indicate current location",
			"Use appropriate focus management",
		},
	}
}

func generateFeedbackPatterns(elements []types.ElementInfo) map[string]interface{} {
	return map[string]interface{}{
		"types": []string{
			"Success messages for completed actions",
			"Error messages for failed actions",
			"Warning messages for important information",
			"Info messages for general guidance",
		},
		"delivery": []string{
			"Use appropriate timing for messages",
			"Provide multiple ways to access feedback",
			"Ensure messages are accessible",
		},
	}
}

func generateDataPatterns(elements []types.ElementInfo) map[string]interface{} {
	return map[string]interface{}{
		"display": []string{
			"Use appropriate components for data types",
			"Implement consistent formatting",
			"Support different data densities",
		},
		"interaction": []string{
			"Provide sorting and filtering options",
			"Implement pagination for large datasets",
			"Support data export when appropriate",
		},
	}
}

func generateFrameworkGuidelines() map[string]interface{} {
	return map[string]interface{}{
		"vanilla": "Import and use directly in HTML",
		"lit":     "Use within Lit templates with proper property binding",
		"react":   "Import and use as JSX elements",
		"vue":     "Register as custom elements for use in templates",
		"angular": "Use CUSTOM_ELEMENTS_SCHEMA for custom elements",
	}
}

func generateBundlingGuidelines() map[string]interface{} {
	return map[string]interface{}{
		"strategy":     "Bundle related components together",
		"splitting":    "Use dynamic imports for lazy loading",
		"optimization": "Tree-shake unused components",
	}
}

func generateImportGuidelines() map[string]interface{} {
	return map[string]interface{}{
		"es6":      "Use ES6 import syntax when possible",
		"cdn":      "Consider CDN imports for simpler projects",
		"bundling": "Import from package entry points",
	}
}

func extractAllCSSProperties(elements []types.ElementInfo) []string {
	propertySet := make(map[string]bool)

	for _, element := range elements {
		for _, prop := range element.CssProperties() {
			propertySet[prop.Name()] = true
		}
	}

	properties := make([]string, 0, len(propertySet))
	for prop := range propertySet {
		properties = append(properties, prop)
	}

	return properties
}
