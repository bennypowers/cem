package resources

import (
	"context"
	"encoding/json"

	"bennypowers.dev/cem/mcp/types"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// handleAccessibilityResource provides comprehensive accessibility patterns and validation rules
func handleAccessibilityResource(ctx context.Context, req *mcp.ReadResourceRequest, registry types.Registry) (*mcp.ReadResourceResult, error) {
	// Generate comprehensive accessibility guidance
	accessibility := generateAccessibilityGuidance(registry)

	contents, err := json.MarshalIndent(accessibility, "", "  ")
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

// generateAccessibilityGuidance creates comprehensive accessibility guidance
func generateAccessibilityGuidance(registry types.Registry) map[string]interface{} {
	elementMap := registry.AllElements()

	// Convert map to slice
	elements := make([]types.ElementInfo, 0, len(elementMap))
	for _, element := range elementMap {
		elements = append(elements, element)
	}

	guidance := map[string]interface{}{
		"overview":       generateAccessibilityOverview(),
		"wcag":           generateWCAGGuidance(),
		"aria":           generateARIAGuidance(),
		"keyboard":       generateKeyboardGuidance(),
		"screenReaders":  generateScreenReaderGuidance(),
		"colorContrast":  generateColorContrastGuidance(),
		"forms":          generateFormAccessibilityGuidance(),
		"dynamicContent": generateDynamicContentGuidance(),
		"customElements": generateCustomElementAccessibilityGuidance(elements),
		"testing":        generateAccessibilityTestingGuidance(),
		"patterns":       generateAccessibilityPatterns(),
		"validation":     generateAccessibilityValidationRules(),
		"tools":          generateAccessibilityTools(),
		"checklist":      generateAccessibilityChecklist(),
	}

	return guidance
}

// generateAccessibilityOverview creates high-level accessibility principles
func generateAccessibilityOverview() map[string]interface{} {
	return map[string]interface{}{
		"principles": map[string]interface{}{
			"perceivable": map[string]interface{}{
				"description": "Information and user interface components must be presentable to users in ways they can perceive",
				"guidelines": []string{
					"Provide text alternatives for non-text content",
					"Provide captions and alternatives for multimedia",
					"Create content that can be presented in different ways without losing meaning",
					"Make it easier for users to see and hear content",
				},
			},
			"operable": map[string]interface{}{
				"description": "User interface components and navigation must be operable",
				"guidelines": []string{
					"Make all functionality available via keyboard",
					"Give users enough time to read and use content",
					"Do not use content that causes seizures or physical reactions",
					"Help users navigate and find content",
				},
			},
			"understandable": map[string]interface{}{
				"description": "Information and the operation of user interface must be understandable",
				"guidelines": []string{
					"Make text readable and understandable",
					"Make content appear and operate in predictable ways",
					"Help users avoid and correct mistakes",
				},
			},
			"robust": map[string]interface{}{
				"description": "Content must be robust enough to be interpreted reliably by a wide variety of user agents",
				"guidelines": []string{
					"Maximize compatibility with assistive technologies",
					"Use valid, semantic HTML",
					"Ensure content works across different browsers and devices",
				},
			},
		},
		"benefits": []string{
			"Improved usability for all users",
			"Better SEO and search engine compatibility",
			"Legal compliance and risk reduction",
			"Increased market reach and inclusivity",
			"Enhanced user experience across devices",
		},
	}
}

// generateWCAGGuidance creates WCAG 2.1 AA compliance guidance
func generateWCAGGuidance() map[string]interface{} {
	return map[string]interface{}{
		"level":   "AA",
		"version": "2.1",
		"categories": map[string]interface{}{
			"perceivable": map[string]interface{}{
				"1.1": "Text Alternatives - Provide text alternatives for non-text content",
				"1.2": "Time-based Media - Provide alternatives for time-based media",
				"1.3": "Adaptable - Create content that can be presented in different ways without losing information",
				"1.4": "Distinguishable - Make it easier for users to see and hear content",
			},
			"operable": map[string]interface{}{
				"2.1": "Keyboard Accessible - Make all functionality available from a keyboard",
				"2.2": "Enough Time - Provide users enough time to read and use content",
				"2.3": "Seizures and Physical Reactions - Do not design content that causes seizures",
				"2.4": "Navigable - Provide ways to help users navigate and find content",
				"2.5": "Input Modalities - Make it easier for users to operate functionality through various inputs",
			},
			"understandable": map[string]interface{}{
				"3.1": "Readable - Make text content readable and understandable",
				"3.2": "Predictable - Make Web pages appear and operate in predictable ways",
				"3.3": "Input Assistance - Help users avoid and correct mistakes",
			},
			"robust": map[string]interface{}{
				"4.1": "Compatible - Maximize compatibility with current and future assistive technologies",
			},
		},
		"successCriteria": generateWCAGSuccessCriteria(),
	}
}

// generateARIAGuidance creates comprehensive ARIA implementation guidance
func generateARIAGuidance() map[string]interface{} {
	return map[string]interface{}{
		"overview": "Accessible Rich Internet Applications (ARIA) provides semantic information about elements to assistive technologies",
		"roles": map[string]interface{}{
			"landmark": map[string]interface{}{
				"banner":        "Identifies the main header of a page",
				"navigation":    "Identifies navigation links",
				"main":          "Identifies the main content of a page",
				"contentinfo":   "Identifies a footer that contains information about the page",
				"search":        "Identifies search functionality",
				"form":          "Identifies a form",
				"complementary": "Identifies a supporting section of the document",
			},
			"widget": map[string]interface{}{
				"button":   "Identifies a clickable button",
				"checkbox": "Identifies a checkbox",
				"dialog":   "Identifies a dialog box",
				"menuitem": "Identifies an item in a menu",
				"slider":   "Identifies a slider control",
				"tab":      "Identifies a tab in a tab list",
				"textbox":  "Identifies a text input",
				"combobox": "Identifies a combobox (select with text input)",
			},
			"structure": map[string]interface{}{
				"heading":  "Identifies a heading",
				"list":     "Identifies a list",
				"listitem": "Identifies an item in a list",
				"table":    "Identifies a table",
				"row":      "Identifies a table row",
				"cell":     "Identifies a table cell",
				"group":    "Identifies a group of related elements",
			},
		},
		"properties": map[string]interface{}{
			"labeling": map[string]interface{}{
				"aria-label":       "Provides an accessible name when no visible label exists",
				"aria-labelledby":  "References elements that label the current element",
				"aria-describedby": "References elements that describe the current element",
			},
			"relationships": map[string]interface{}{
				"aria-controls": "References elements controlled by the current element",
				"aria-owns":     "References elements owned by the current element",
				"aria-flowto":   "References the next element in an alternative reading order",
			},
			"widget": map[string]interface{}{
				"aria-checked":  "Indicates the checked state of a checkbox or radio button",
				"aria-expanded": "Indicates if a collapsible element is expanded",
				"aria-selected": "Indicates the selected state of an option",
				"aria-pressed":  "Indicates the pressed state of a toggle button",
				"aria-hidden":   "Indicates whether an element is hidden from assistive technology",
			},
		},
		"states":        generateARIAStates(),
		"bestPractices": generateARIABestPractices(),
	}
}

// generateKeyboardGuidance creates keyboard accessibility guidance
func generateKeyboardGuidance() map[string]interface{} {
	return map[string]interface{}{
		"overview": "All interactive functionality must be available via keyboard",
		"navigation": map[string]interface{}{
			"tab":      "Use Tab to move forward through interactive elements",
			"shiftTab": "Use Shift+Tab to move backward through interactive elements",
			"arrow":    "Use arrow keys for directional navigation within components",
			"home":     "Use Home to move to the first item in a collection",
			"end":      "Use End to move to the last item in a collection",
			"pageUp":   "Use Page Up to move up one page in scrollable content",
			"pageDown": "Use Page Down to move down one page in scrollable content",
		},
		"activation": map[string]interface{}{
			"enter":  "Use Enter to activate buttons and links",
			"space":  "Use Space to activate buttons and toggle states",
			"escape": "Use Escape to close dialogs and cancel operations",
		},
		"focusManagement": map[string]interface{}{
			"visible":     "Always provide visible focus indicators",
			"logical":     "Ensure focus order follows a logical sequence",
			"trapping":    "Implement focus trapping in modal dialogs",
			"restoration": "Restore focus when closing overlays",
		},
		"tabindex": map[string]interface{}{
			"0":        "Makes element focusable in normal tab order",
			"-1":       "Makes element focusable but not in tab order (for programmatic focus)",
			"positive": "Avoid positive tabindex values as they disrupt natural tab order",
		},
		"patterns": generateKeyboardPatterns(),
	}
}

// generateScreenReaderGuidance creates screen reader compatibility guidance
func generateScreenReaderGuidance() map[string]interface{} {
	return map[string]interface{}{
		"overview": "Ensure content works well with screen readers and other assistive technologies",
		"testing": map[string]interface{}{
			"nvda":      "Test with NVDA (Windows, free)",
			"jaws":      "Test with JAWS (Windows, commercial)",
			"voiceOver": "Test with VoiceOver (macOS, built-in)",
			"orca":      "Test with Orca (Linux, free)",
			"talkback":  "Test with TalkBack (Android, built-in)",
		},
		"semantics": map[string]interface{}{
			"headings":  "Use proper heading hierarchy (h1-h6)",
			"landmarks": "Use landmark roles or semantic HTML elements",
			"lists":     "Use proper list markup (ul, ol, li)",
			"tables":    "Use proper table markup with headers",
			"forms":     "Associate labels with form controls",
		},
		"announcements": map[string]interface{}{
			"live":      "Use aria-live for dynamic content updates",
			"status":    "Use status role for status messages",
			"alert":     "Use alert role for important notifications",
			"assertive": "Use aria-live='assertive' for urgent updates",
			"polite":    "Use aria-live='polite' for non-urgent updates",
		},
		"content": map[string]interface{}{
			"descriptive":  "Provide descriptive text for images and media",
			"context":      "Provide sufficient context for all content",
			"instructions": "Include clear instructions for complex interactions",
		},
	}
}

// generateColorContrastGuidance creates color and visual accessibility guidance
func generateColorContrastGuidance() map[string]interface{} {
	return map[string]interface{}{
		"wcag": map[string]interface{}{
			"aa": map[string]interface{}{
				"normal": "4.5:1 minimum contrast ratio for normal text",
				"large":  "3:1 minimum contrast ratio for large text (18pt+ or 14pt+ bold)",
			},
			"aaa": map[string]interface{}{
				"normal": "7:1 enhanced contrast ratio for normal text",
				"large":  "4.5:1 enhanced contrast ratio for large text",
			},
		},
		"tools": []string{
			"WebAIM Contrast Checker",
			"Colour Contrast Analyser",
			"Stark (design tool plugin)",
			"Chrome DevTools contrast ratio",
		},
		"considerations": map[string]interface{}{
			"colorBlindness": "Don't rely solely on color to convey information",
			"focus":          "Ensure focus indicators have sufficient contrast",
			"states":         "Maintain contrast in all interactive states",
			"backgrounds":    "Consider background images and gradients",
		},
		"testing": []string{
			"Test in different lighting conditions",
			"Use color blindness simulators",
			"Check with high contrast mode enabled",
			"Verify with automatic contrast checking tools",
		},
	}
}

// generateFormAccessibilityGuidance creates form-specific accessibility guidance
func generateFormAccessibilityGuidance() map[string]interface{} {
	return map[string]interface{}{
		"labeling": map[string]interface{}{
			"explicit": "Use <label for='id'> to explicitly associate labels",
			"implicit": "Wrap form controls in <label> elements",
			"aria":     "Use aria-label when visible labels aren't appropriate",
			"required": "Clearly indicate required fields",
		},
		"validation": map[string]interface{}{
			"timing":       "Provide validation feedback at appropriate times",
			"association":  "Use aria-describedby to associate error messages",
			"instructions": "Provide clear instructions for fixing errors",
			"prevention":   "Help users avoid errors with clear guidance",
		},
		"grouping": map[string]interface{}{
			"fieldset":   "Use fieldset and legend for related form controls",
			"sections":   "Group related form sections logically",
			"navigation": "Provide a logical tab order through forms",
		},
		"help": map[string]interface{}{
			"instructions": "Provide clear instructions before form fields",
			"format":       "Explain expected format for complex inputs",
			"examples":     "Provide examples when helpful",
			"assistance":   "Offer contextual help when needed",
		},
	}
}

// generateDynamicContentGuidance creates dynamic content accessibility guidance
func generateDynamicContentGuidance() map[string]interface{} {
	return map[string]interface{}{
		"liveRegions": map[string]interface{}{
			"polite":    "Use for non-urgent updates that shouldn't interrupt",
			"assertive": "Use for urgent updates that should interrupt",
			"off":       "Use to temporarily disable announcements",
			"atomic":    "Use aria-atomic to control what gets announced",
		},
		"loading": map[string]interface{}{
			"states":   "Announce loading states to screen readers",
			"progress": "Use progress elements for longer operations",
			"spinners": "Provide text alternatives for loading spinners",
		},
		"updates": map[string]interface{}{
			"meaningful": "Only announce meaningful changes",
			"context":    "Provide sufficient context for updates",
			"timing":     "Don't overwhelm users with too many announcements",
		},
		"focus": map[string]interface{}{
			"management":   "Manage focus when content changes",
			"preservation": "Preserve user's place when possible",
			"restoration":  "Restore focus appropriately",
		},
	}
}

// generateCustomElementAccessibilityGuidance creates custom element specific guidance
func generateCustomElementAccessibilityGuidance(elements []types.ElementInfo) map[string]interface{} {
	elementGuidance := make(map[string]interface{})

	for _, element := range elements {
		elementGuidance[element.TagName()] = generateElementAccessibilityGuidance(element)
	}

	return map[string]interface{}{
		"general": map[string]interface{}{
			"semantics": "Ensure custom elements have appropriate implicit or explicit roles",
			"keyboard":  "Implement standard keyboard interaction patterns",
			"focus":     "Manage focus appropriately within shadow DOM",
			"aria":      "Use ARIA to enhance semantics when needed",
		},
		"shadowDOM": map[string]interface{}{
			"focus":      "Ensure focus works properly across shadow boundaries",
			"delegation": "Consider focus delegation for internal elements",
			"styling":    "Ensure focus indicators work in shadow DOM",
		},
		"slots": map[string]interface{}{
			"semantics":  "Preserve semantic meaning of slotted content",
			"roles":      "Consider how slotted content affects component semantics",
			"validation": "Validate that slotted content is appropriate",
		},
		"elements": elementGuidance,
	}
}

// generateAccessibilityTestingGuidance creates testing-specific guidance
func generateAccessibilityTestingGuidance() map[string]interface{} {
	return map[string]interface{}{
		"automated": map[string]interface{}{
			"tools": []string{
				"axe-core for comprehensive accessibility testing",
				"Lighthouse for accessibility audits",
				"Pa11y for command-line testing",
				"WAVE for browser-based testing",
			},
			"integration": "Integrate automated testing into CI/CD pipelines",
			"limitations": "Automated tools catch ~30% of issues - manual testing required",
		},
		"manual": map[string]interface{}{
			"keyboard":     "Test all functionality with keyboard only",
			"screenReader": "Test with actual screen reader software",
			"zoom":         "Test at 200% zoom and high contrast modes",
			"mobile":       "Test with mobile screen readers and voice control",
		},
		"users": map[string]interface{}{
			"inclusion": "Include users with disabilities in testing process",
			"feedback":  "Gather feedback from assistive technology users",
			"scenarios": "Test real-world usage scenarios",
		},
	}
}

// generateAccessibilityPatterns creates common accessibility patterns
func generateAccessibilityPatterns() map[string]interface{} {
	return map[string]interface{}{
		"button": map[string]interface{}{
			"role":     "Use button role for clickable elements",
			"keyboard": "Respond to Enter and Space key presses",
			"state":    "Use aria-pressed for toggle buttons",
			"disabled": "Use aria-disabled instead of disabled attribute when appropriate",
		},
		"dialog": map[string]interface{}{
			"role":     "Use dialog role for modal dialogs",
			"focus":    "Trap focus within dialog when open",
			"escape":   "Close dialog with Escape key",
			"labeling": "Use aria-labelledby or aria-label for dialog title",
		},
		"menu": map[string]interface{}{
			"role":      "Use menu and menuitem roles appropriately",
			"keyboard":  "Implement arrow key navigation",
			"state":     "Use aria-expanded for menu button state",
			"selection": "Use aria-selected for selected menu items",
		},
		"tabs": map[string]interface{}{
			"roles":    "Use tablist, tab, and tabpanel roles",
			"keyboard": "Implement arrow key navigation between tabs",
			"state":    "Use aria-selected for active tab",
			"controls": "Use aria-controls to associate tabs with panels",
		},
	}
}

// generateAccessibilityValidationRules creates validation rules for accessibility
func generateAccessibilityValidationRules() map[string]interface{} {
	return map[string]interface{}{
		"required": map[string]interface{}{
			"altText":  "All images must have alt text or be marked decorative",
			"labels":   "All form controls must have accessible labels",
			"headings": "Page must have proper heading hierarchy",
			"focus":    "All interactive elements must be keyboard accessible",
		},
		"recommended": map[string]interface{}{
			"landmarks": "Use landmark elements to structure page content",
			"skipLinks": "Provide skip links for keyboard navigation",
			"contrast":  "Ensure sufficient color contrast ratios",
			"resize":    "Content must be functional at 200% zoom",
		},
		"warnings": map[string]interface{}{
			"colorOnly": "Don't use color alone to convey information",
			"autoplay":  "Avoid auto-playing audio or video content",
			"timeouts":  "Provide warnings before session timeouts",
			"motion":    "Respect user preferences for reduced motion",
		},
	}
}

// generateAccessibilityTools creates list of accessibility tools
func generateAccessibilityTools() map[string]interface{} {
	return map[string]interface{}{
		"browser": map[string]interface{}{
			"devtools": "Chrome/Firefox DevTools accessibility features",
			"extensions": []string{
				"axe DevTools",
				"WAVE Evaluation Tool",
				"Accessibility Insights",
				"Lighthouse",
			},
		},
		"desktop": []string{
			"Colour Contrast Analyser",
			"Screen reader software (NVDA, JAWS, VoiceOver)",
			"Accessibility Inspector (macOS)",
		},
		"online": []string{
			"WebAIM Contrast Checker",
			"WAVE Web Accessibility Evaluation Tool",
			"Accessible Colors",
		},
		"development": []string{
			"axe-core JavaScript library",
			"Pa11y command line tool",
			"Lighthouse CI",
			"Playwright accessibility testing",
		},
	}
}

// generateAccessibilityChecklist creates comprehensive accessibility checklist
func generateAccessibilityChecklist() map[string]interface{} {
	return map[string]interface{}{
		"content": []string{
			"All images have appropriate alt text",
			"Headings are properly structured (h1-h6)",
			"Text has sufficient color contrast",
			"Links have descriptive text",
			"Content is readable and understandable",
		},
		"keyboard": []string{
			"All functionality is keyboard accessible",
			"Tab order is logical and intuitive",
			"Focus indicators are clearly visible",
			"No keyboard traps exist",
			"Keyboard shortcuts don't conflict",
		},
		"forms": []string{
			"All form controls have labels",
			"Required fields are clearly marked",
			"Error messages are descriptive and helpful",
			"Form validation is accessible",
			"Instructions are provided where needed",
		},
		"navigation": []string{
			"Multiple ways to navigate content",
			"Skip links are provided",
			"Breadcrumbs show current location",
			"Site search is available and accessible",
		},
		"media": []string{
			"Videos have captions",
			"Audio has transcripts",
			"Media doesn't auto-play with sound",
			"Controls are keyboard accessible",
		},
		"dynamic": []string{
			"Status changes are announced",
			"Loading states are communicated",
			"Error messages are announced",
			"Live regions are used appropriately",
		},
	}
}

// Helper functions for generating specific guidance sections

func generateWCAGSuccessCriteria() map[string]interface{} {
	return map[string]interface{}{
		"level_a": []string{
			"1.1.1 Non-text Content",
			"1.3.1 Info and Relationships",
			"1.3.2 Meaningful Sequence",
			"1.3.3 Sensory Characteristics",
			"1.4.1 Use of Color",
			"1.4.2 Audio Control",
			"2.1.1 Keyboard",
			"2.1.2 No Keyboard Trap",
			"2.1.4 Character Key Shortcuts",
			"2.2.1 Timing Adjustable",
			"2.2.2 Pause, Stop, Hide",
			"2.3.1 Three Flashes or Below Threshold",
			"2.4.1 Bypass Blocks",
			"2.4.2 Page Titled",
			"2.4.3 Focus Order",
			"2.4.4 Link Purpose (In Context)",
			"2.5.1 Pointer Gestures",
			"2.5.2 Pointer Cancellation",
			"2.5.3 Label in Name",
			"2.5.4 Motion Actuation",
			"3.1.1 Language of Page",
			"3.2.1 On Focus",
			"3.2.2 On Input",
			"3.3.1 Error Identification",
			"3.3.2 Labels or Instructions",
			"4.1.1 Parsing",
			"4.1.2 Name, Role, Value",
		},
		"level_aa": []string{
			"1.2.4 Captions (Live)",
			"1.2.5 Audio Description (Prerecorded)",
			"1.3.4 Orientation",
			"1.3.5 Identify Input Purpose",
			"1.4.3 Contrast (Minimum)",
			"1.4.4 Resize text",
			"1.4.5 Images of Text",
			"1.4.10 Reflow",
			"1.4.11 Non-text Contrast",
			"1.4.12 Text Spacing",
			"1.4.13 Content on Hover or Focus",
			"2.4.5 Multiple Ways",
			"2.4.6 Headings and Labels",
			"2.4.7 Focus Visible",
			"3.1.2 Language of Parts",
			"3.2.3 Consistent Navigation",
			"3.2.4 Consistent Identification",
			"3.3.3 Error Suggestion",
			"3.3.4 Error Prevention (Legal, Financial, Data)",
			"4.1.3 Status Messages",
		},
	}
}

func generateARIAStates() map[string]interface{} {
	return map[string]interface{}{
		"boolean": map[string]interface{}{
			"aria-checked":  "true/false/mixed for checkboxes and radio buttons",
			"aria-expanded": "true/false for collapsible elements",
			"aria-hidden":   "true/false to hide decorative elements",
			"aria-pressed":  "true/false/mixed for toggle buttons",
			"aria-selected": "true/false for selectable items",
			"aria-disabled": "true/false for disabled elements",
		},
		"string": map[string]interface{}{
			"aria-label":       "Accessible name when no visible label",
			"aria-placeholder": "Example of expected input",
			"aria-valuetext":   "Human-readable text alternative for values",
		},
		"reference": map[string]interface{}{
			"aria-labelledby":  "ID reference to labeling elements",
			"aria-describedby": "ID reference to describing elements",
			"aria-controls":    "ID reference to controlled elements",
			"aria-owns":        "ID reference to owned elements",
		},
	}
}

func generateARIABestPractices() []string {
	return []string{
		"Use semantic HTML first, ARIA second",
		"Don't change native semantics unnecessarily",
		"Ensure all interactive elements are keyboard accessible",
		"Use ARIA labels and descriptions appropriately",
		"Test with actual assistive technologies",
		"Keep ARIA simple and focused",
		"Validate ARIA usage with accessibility tools",
		"Update ARIA states dynamically as needed",
	}
}

func generateKeyboardPatterns() map[string]interface{} {
	return map[string]interface{}{
		"button": map[string]interface{}{
			"activation": "Enter or Space to activate",
			"focus":      "Tab to focus, visual focus indicator required",
		},
		"menu": map[string]interface{}{
			"navigation": "Arrow keys to navigate menu items",
			"activation": "Enter to select menu item",
			"escape":     "Escape to close menu",
		},
		"dialog": map[string]interface{}{
			"trapping": "Tab cycles through dialog elements only",
			"escape":   "Escape to close dialog",
			"focus":    "Focus moves to dialog when opened",
		},
		"tabs": map[string]interface{}{
			"navigation": "Arrow keys to navigate tabs",
			"activation": "Enter or Space to activate tab",
			"focus":      "Tab moves to tab panel content",
		},
	}
}

func generateElementAccessibilityGuidance(element types.ElementInfo) map[string]interface{} {
	guidance := map[string]interface{}{
		"tagName":     element.TagName(),
		"description": element.Description(),
		"considerations": []string{
			"Ensure appropriate ARIA roles if custom semantics are needed",
			"Implement keyboard accessibility for interactive elements",
			"Provide accessible names and descriptions",
			"Test with screen readers and other assistive technologies",
		},
	}

	// Add specific guidance based on element attributes and events
	if len(element.Attributes()) > 0 {
		guidance["attributes"] = "Review attribute usage for accessibility implications"
	}

	if len(element.Events()) > 0 {
		guidance["events"] = "Ensure event handlers are accessible to keyboard and assistive technology users"
	}

	if len(element.Slots()) > 0 {
		guidance["slots"] = "Ensure slotted content maintains semantic meaning and accessibility"
	}

	return guidance
}
