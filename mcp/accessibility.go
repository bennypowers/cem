package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// AccessibilityResource provides comprehensive accessibility guidance
type AccessibilityResource struct {
	Overview            string                    `json:"overview"`
	WcagGuidelines      []WcagGuideline          `json:"wcagGuidelines"`
	ElementPatterns     map[string]A11yPattern   `json:"elementPatterns"`
	CommonPatterns      []A11yPattern            `json:"commonPatterns"`
	TestingGuidance     TestingGuidance          `json:"testingGuidance"`
	Tools               []A11yTool               `json:"tools"`
	Checklist           []ChecklistItem          `json:"checklist"`
}

// WcagGuideline represents a WCAG accessibility guideline
type WcagGuideline struct {
	Principle   string   `json:"principle"`
	Guideline   string   `json:"guideline"`
	Level       string   `json:"level"` // A, AA, AAA
	Description string   `json:"description"`
	Examples    []string `json:"examples"`
	Testing     string   `json:"testing"`
}

// A11yPattern represents an accessibility pattern for specific use cases
type A11yPattern struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	WhenToUse   string   `json:"whenToUse"`
	AriaRoles   []string `json:"ariaRoles,omitempty"`
	AriaProps   []string `json:"ariaProperties,omitempty"`
	Examples    []string `json:"examples"`
	CommonMistakes []string `json:"commonMistakes,omitempty"`
}

// TestingGuidance provides accessibility testing recommendations
type TestingGuidance struct {
	AutomatedTesting []TestingMethod `json:"automatedTesting"`
	ManualTesting    []TestingMethod `json:"manualTesting"`
	ScreenReaders    []ScreenReader  `json:"screenReaders"`
	KeyboardTesting  []string        `json:"keyboardTesting"`
}

// TestingMethod describes a testing approach
type TestingMethod struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tools       []string `json:"tools"`
	Steps       []string `json:"steps"`
}

// ScreenReader describes screen reader testing
type ScreenReader struct {
	Name         string   `json:"name"`
	Platform     string   `json:"platform"`
	TestingTips  []string `json:"testingTips"`
}

// A11yTool represents an accessibility tool
type A11yTool struct {
	Name        string `json:"name"`
	Type        string `json:"type"` // "browser-extension", "cli", "online", "desktop"
	Description string `json:"description"`
	Url         string `json:"url,omitempty"`
}

// ChecklistItem represents an accessibility checklist item
type ChecklistItem struct {
	Category    string `json:"category"`
	Requirement string `json:"requirement"`
	Level       string `json:"level"`
	HowToTest   string `json:"howToTest"`
}

// handleAccessibilityResource provides comprehensive accessibility guidance
func (s *Server) handleAccessibilityResource(ctx context.Context, uri string) (*mcp.ReadResourceResult, error) {
	resource := &AccessibilityResource{
		Overview:        s.generateAccessibilityOverview(),
		WcagGuidelines:  s.generateWcagGuidelines(),
		ElementPatterns: s.generateElementAccessibilityPatterns(),
		CommonPatterns:  s.generateCommonAccessibilityPatterns(),
		TestingGuidance: s.generateTestingGuidance(),
		Tools:           s.generateAccessibilityTools(),
		Checklist:       s.generateAccessibilityChecklist(),
	}

	contents, err := json.MarshalIndent(resource, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal accessibility resource: %w", err)
	}

	return &mcp.ReadResourceResult{
		Contents: []*mcp.ResourceContents{{
			URI:      uri,
			MIMEType: "application/json",
			Text:     string(contents),
		}},
	}, nil
}

// generateAccessibilityOverview creates an overview of accessibility requirements
func (s *Server) generateAccessibilityOverview() string {
	return `
Accessibility is a fundamental requirement for web applications. This resource provides comprehensive 
guidance for implementing accessible custom elements that work for all users, including those using 
assistive technologies.

Key Principles:
1. Perceivable - Information must be presentable to users in ways they can perceive
2. Operable - Interface components must be operable by all users
3. Understandable - Information and UI operation must be understandable
4. Robust - Content must be robust enough for interpretation by assistive technologies

This guidance follows WCAG 2.1 Level AA standards and provides patterns specific to custom elements.
`
}

// generateWcagGuidelines creates key WCAG guidelines relevant to custom elements
func (s *Server) generateWcagGuidelines() []WcagGuideline {
	return []WcagGuideline{
		{
			Principle:   "Perceivable",
			Guideline:   "1.1.1 Non-text Content",
			Level:       "A",
			Description: "All non-text content has a text alternative that serves the equivalent purpose",
			Examples: []string{
				`<my-icon alt="Settings" role="img"></my-icon>`,
				`<my-chart aria-label="Sales data from 2023"></my-chart>`,
				`<my-avatar alt="Profile picture of John Doe"></my-avatar>`,
			},
			Testing: "Verify all images, icons, and visual content have appropriate alternative text",
		},
		{
			Principle:   "Perceivable",
			Guideline:   "1.4.3 Contrast (Minimum)",
			Level:       "AA",
			Description: "Text and background colors have a contrast ratio of at least 4.5:1",
			Examples: []string{
				`/* Good contrast */`,
				`--text-color: #000000; --background-color: #ffffff; /* 21:1 ratio */`,
				`--primary-color: #0056b3; --text-on-primary: #ffffff; /* 4.5:1 ratio */`,
			},
			Testing: "Use contrast checking tools to verify color combinations meet minimum ratios",
		},
		{
			Principle:   "Operable",
			Guideline:   "2.1.1 Keyboard",
			Level:       "A",
			Description: "All functionality is available from a keyboard",
			Examples: []string{
				`<my-button tabindex="0" role="button">Keyboard accessible</my-button>`,
				`<my-dropdown role="combobox" aria-expanded="false" tabindex="0">`,
				`<!-- Ensure Enter and Space keys trigger actions -->`,
			},
			Testing: "Navigate using only keyboard (Tab, Enter, Space, Arrow keys)",
		},
		{
			Principle:   "Operable",
			Guideline:   "2.4.7 Focus Visible",
			Level:       "AA",
			Description: "Keyboard focus indicators are clearly visible",
			Examples: []string{
				`my-button:focus { outline: 2px solid #007acc; outline-offset: 2px; }`,
				`my-input:focus { box-shadow: 0 0 0 2px #007acc; }`,
			},
			Testing: "Tab through interface and verify focus indicators are clearly visible",
		},
		{
			Principle:   "Understandable",
			Guideline:   "3.2.2 On Input",
			Level:       "A",
			Description: "Changing input values doesn't cause unexpected context changes",
			Examples: []string{
				`<!-- Good: Explicit submission -->`,
				`<my-form><my-input><my-button type="submit">Submit</my-button></my-form>`,
				`<!-- Avoid: Auto-submission on input -->`,
			},
			Testing: "Verify form inputs don't automatically submit or navigate away",
		},
		{
			Principle:   "Robust",
			Guideline:   "4.1.2 Name, Role, Value",
			Level:       "A",
			Description: "UI components have accessible names, roles, and values",
			Examples: []string{
				`<my-toggle role="switch" aria-checked="false" aria-label="Enable notifications">`,
				`<my-slider role="slider" aria-valuemin="0" aria-valuemax="100" aria-valuenow="50">`,
				`<my-dialog role="dialog" aria-labelledby="dialog-title" aria-modal="true">`,
			},
			Testing: "Use screen readers to verify components are properly announced",
		},
	}
}

// generateElementAccessibilityPatterns creates patterns for specific element types
func (s *Server) generateElementAccessibilityPatterns() map[string]A11yPattern {
	patterns := make(map[string]A11yPattern)
	
	elements := s.registry.GetAllElements()
	for tagName, element := range elements {
		pattern := s.generateElementAccessibilityPattern(element)
		if pattern.Name != "" {
			patterns[tagName] = pattern
		}
	}
	
	return patterns
}

// generateElementAccessibilityPattern creates accessibility pattern for a specific element
func (s *Server) generateElementAccessibilityPattern(element *ElementInfo) A11yPattern {
	tagName := strings.ToLower(element.TagName)
	
	switch {
	case strings.Contains(tagName, "button"):
		return A11yPattern{
			Name:        "Button Pattern",
			Description: "Interactive button element with proper ARIA support",
			WhenToUse:   "For clickable actions and form submissions",
			AriaRoles:   []string{"button"},
			AriaProps:   []string{"aria-label", "aria-describedby", "aria-pressed"},
			Examples: []string{
				fmt.Sprintf(`<%s role="button" aria-label="Save document">Save</%s>`, element.TagName, element.TagName),
				fmt.Sprintf(`<%s role="button" aria-pressed="false">Toggle</%s>`, element.TagName, element.TagName),
			},
			CommonMistakes: []string{
				"Missing role='button' for custom button elements",
				"No keyboard support (Enter/Space keys)",
				"Missing aria-label for icon-only buttons",
			},
		}
		
	case strings.Contains(tagName, "input") || strings.Contains(tagName, "field"):
		return A11yPattern{
			Name:        "Form Input Pattern",
			Description: "Form input element with proper labeling and validation",
			WhenToUse:   "For collecting user input in forms",
			AriaProps:   []string{"aria-label", "aria-describedby", "aria-required", "aria-invalid"},
			Examples: []string{
				fmt.Sprintf(`<label for="email">Email:</label>
<%s id="email" type="email" aria-required="true" aria-describedby="email-help">
<div id="email-help">Enter your email address</div>`, element.TagName),
			},
			CommonMistakes: []string{
				"Missing or improper labeling",
				"No association between label and input",
				"Missing validation feedback",
			},
		}
		
	case strings.Contains(tagName, "dialog") || strings.Contains(tagName, "modal"):
		return A11yPattern{
			Name:        "Dialog Pattern",
			Description: "Modal dialog with focus management and escape handling",
			WhenToUse:   "For modal dialogs, alerts, and confirmations",
			AriaRoles:   []string{"dialog", "alertdialog"},
			AriaProps:   []string{"aria-labelledby", "aria-describedby", "aria-modal"},
			Examples: []string{
				fmt.Sprintf(`<%s role="dialog" aria-labelledby="dialog-title" aria-modal="true">
  <h2 id="dialog-title">Confirmation</h2>
  <p>Are you sure you want to delete this item?</p>
</%s>`, element.TagName, element.TagName),
			},
			CommonMistakes: []string{
				"Focus not trapped within dialog",
				"No escape key handling",
				"Missing aria-modal attribute",
				"Background content still accessible",
			},
		}
		
	default:
		// Generic pattern based on element features
		if len(element.Events()) > 0 {
			return A11yPattern{
				Name:        "Interactive Element Pattern",
				Description: "Interactive custom element with event handling",
				WhenToUse:   "For elements that respond to user interaction",
				AriaProps:   []string{"aria-label", "role", "tabindex"},
				Examples: []string{
					fmt.Sprintf(`<%s role="button" tabindex="0" aria-label="Interactive element">Content</%s>`, element.TagName, element.TagName),
				},
				CommonMistakes: []string{
					"Missing keyboard support",
					"No focus indicator",
					"Missing semantic role",
				},
			}
		}
	}
	
	return A11yPattern{}
}

// generateCommonAccessibilityPatterns creates common accessibility patterns
func (s *Server) generateCommonAccessibilityPatterns() []A11yPattern {
	return []A11yPattern{
		{
			Name:        "Skip Links",
			Description: "Navigation aids for keyboard users to skip repetitive content",
			WhenToUse:   "At the beginning of pages with navigation or repeated content",
			Examples: []string{
				`<a href="#main-content" class="skip-link">Skip to main content</a>`,
				`<main id="main-content"><my-content>Main page content</my-content></main>`,
			},
		},
		{
			Name:        "Focus Management",
			Description: "Proper focus handling for dynamic content and navigation",
			WhenToUse:   "When content changes dynamically or in single-page applications",
			Examples: []string{
				`// After navigation, focus the main heading`,
				`document.querySelector('h1').focus();`,
				`// For dialogs, focus the first interactive element`,
				`dialog.querySelector('button').focus();`,
			},
		},
		{
			Name:        "Live Regions",
			Description: "Announce dynamic content changes to screen readers",
			WhenToUse:   "For status updates, error messages, and dynamic content",
			AriaProps:   []string{"aria-live", "aria-atomic", "aria-relevant"},
			Examples: []string{
				`<div aria-live="polite" id="status">Status updates appear here</div>`,
				`<div aria-live="assertive" role="alert">Error messages here</div>`,
			},
		},
		{
			Name:        "Progressive Enhancement",
			Description: "Ensure functionality works without JavaScript",
			WhenToUse:   "For all interactive custom elements",
			Examples: []string{
				`<my-enhanced-form>`,
				`  <form action="/submit" method="post">`,
				`    <!-- Fallback form content -->`,
				`  </form>`,
				`</my-enhanced-form>`,
			},
		},
	}
}

// generateTestingGuidance creates comprehensive testing guidance
func (s *Server) generateTestingGuidance() TestingGuidance {
	return TestingGuidance{
		AutomatedTesting: []TestingMethod{
			{
				Name:        "Automated Accessibility Testing",
				Description: "Use automated tools to catch common accessibility issues",
				Tools:       []string{"axe-core", "Pa11y", "Lighthouse", "WAVE"},
				Steps: []string{
					"Install accessibility testing tools in CI/CD pipeline",
					"Run automated scans on all pages and components",
					"Address all violations found by automated tools",
					"Note: Automated tools catch ~30% of issues, manual testing is essential",
				},
			},
			{
				Name:        "Unit Testing Accessibility",
				Description: "Test accessibility features in component unit tests",
				Tools:       []string{"@testing-library/jest-dom", "jest-axe"},
				Steps: []string{
					"Test ARIA attributes are properly set",
					"Verify keyboard event handlers work correctly",
					"Test focus management in interactive components",
					"Validate accessible names and descriptions",
				},
			},
		},
		ManualTesting: []TestingMethod{
			{
				Name:        "Keyboard Navigation Testing",
				Description: "Test all functionality using only keyboard",
				Steps: []string{
					"Navigate using Tab key to move between interactive elements",
					"Use Enter and Space to activate buttons and controls",
					"Use arrow keys for menu and list navigation",
					"Verify all interactive elements are reachable",
					"Ensure focus indicators are visible",
					"Test Escape key for closing dialogs and menus",
				},
			},
			{
				Name:        "Screen Reader Testing",
				Description: "Test with actual screen reader software",
				Steps: []string{
					"Navigate using screen reader commands",
					"Verify content is announced correctly",
					"Test landmark navigation",
					"Check form labeling and validation messages",
					"Verify dynamic content updates are announced",
				},
			},
		},
		ScreenReaders: []ScreenReader{
			{
				Name:        "NVDA",
				Platform:    "Windows",
				TestingTips: []string{
					"Free and widely used",
					"Good for testing basic functionality",
					"Use browse mode (virtual buffer) for content",
					"Use focus mode for interactive elements",
				},
			},
			{
				Name:        "JAWS",
				Platform:    "Windows",
				TestingTips: []string{
					"Most popular commercial screen reader",
					"Test if possible due to market share",
					"Similar navigation to NVDA",
				},
			},
			{
				Name:        "VoiceOver",
				Platform:    "macOS/iOS",
				TestingTips: []string{
					"Built into Apple devices",
					"Use rotor for navigation",
					"Test gesture controls on mobile",
				},
			},
		},
		KeyboardTesting: []string{
			"Tab - Move to next interactive element",
			"Shift+Tab - Move to previous interactive element",
			"Enter - Activate buttons and links",
			"Space - Activate buttons, check checkboxes",
			"Arrow keys - Navigate within components (menus, tabs)",
			"Escape - Close dialogs, cancel operations",
			"Home/End - Move to beginning/end of lists",
		},
	}
}

// generateAccessibilityTools creates a list of recommended accessibility tools
func (s *Server) generateAccessibilityTools() []A11yTool {
	return []A11yTool{
		{
			Name:        "axe DevTools",
			Type:        "browser-extension",
			Description: "Browser extension for automated accessibility testing",
			Url:         "https://www.deque.com/axe/devtools/",
		},
		{
			Name:        "WAVE",
			Type:        "browser-extension",
			Description: "Web Accessibility Evaluation Tool",
			Url:         "https://wave.webaim.org/",
		},
		{
			Name:        "Lighthouse",
			Type:        "browser-extension",
			Description: "Google's web performance and accessibility auditing tool",
			Url:         "https://developers.google.com/web/tools/lighthouse",
		},
		{
			Name:        "Pa11y",
			Type:        "cli",
			Description: "Command line accessibility testing tool",
			Url:         "https://pa11y.org/",
		},
		{
			Name:        "Color Contrast Analyser",
			Type:        "desktop",
			Description: "Tool for testing color contrast ratios",
			Url:         "https://www.tpgi.com/color-contrast-checker/",
		},
		{
			Name:        "WebAIM Contrast Checker",
			Type:        "online",
			Description: "Online color contrast checking tool",
			Url:         "https://webaim.org/resources/contrastchecker/",
		},
	}
}

// generateAccessibilityChecklist creates a comprehensive accessibility checklist
func (s *Server) generateAccessibilityChecklist() []ChecklistItem {
	return []ChecklistItem{
		{
			Category:    "Keyboard Navigation",
			Requirement: "All interactive elements are keyboard accessible",
			Level:       "A",
			HowToTest:   "Navigate using only keyboard (Tab, Enter, Space, arrows)",
		},
		{
			Category:    "Focus Management",
			Requirement: "Focus indicators are clearly visible",
			Level:       "AA",
			HowToTest:   "Tab through interface and verify focus indicators",
		},
		{
			Category:    "Alternative Text",
			Requirement: "All images have appropriate alternative text",
			Level:       "A",
			HowToTest:   "Check alt attributes on images and icons",
		},
		{
			Category:    "Color Contrast",
			Requirement: "Text has sufficient contrast (4.5:1 minimum)",
			Level:       "AA",
			HowToTest:   "Use color contrast checking tools",
		},
		{
			Category:    "Form Labels",
			Requirement: "All form inputs have associated labels",
			Level:       "A",
			HowToTest:   "Verify label elements or aria-label attributes",
		},
		{
			Category:    "Semantic Markup",
			Requirement: "Content uses appropriate semantic HTML",
			Level:       "A",
			HowToTest:   "Review HTML structure for semantic elements",
		},
		{
			Category:    "ARIA Usage",
			Requirement: "ARIA attributes are used correctly",
			Level:       "A",
			HowToTest:   "Validate ARIA syntax and usage patterns",
		},
		{
			Category:    "Dynamic Content",
			Requirement: "Screen readers are notified of content changes",
			Level:       "A",
			HowToTest:   "Test with screen readers or check aria-live regions",
		},
	}
}