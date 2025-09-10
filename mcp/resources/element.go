package resources

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"bennypowers.dev/cem/mcp/types"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var titleCase = cases.Title(language.English)

// handleElementResource provides detailed information about a specific custom element
func handleElementResource(
	ctx context.Context,
	req *mcp.ReadResourceRequest,
	registry types.Registry,
) (*mcp.ReadResourceResult, error) {
	// Extract tag name from URI template: cem://element/{tagName}
	tagName, err := extractTagNameFromURI(req.Params.URI)
	if err != nil {
		return nil, fmt.Errorf("invalid element URI: %w", err)
	}

	// Find the element in the registry
	element, err := registry.ElementInfo(tagName)
	if err != nil {
		return nil, fmt.Errorf("element not found: %s: %w", tagName, err)
	}

	// Create comprehensive element documentation
	elementDoc := createElementDocumentation(element)

	contents, err := json.MarshalIndent(elementDoc, "", "  ")
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

// extractTagNameFromURI extracts the tag name from a URI like cem://element/{tagName}
func extractTagNameFromURI(uri string) (string, error) {
	const prefix = "cem://element/"
	if !strings.HasPrefix(uri, prefix) {
		return "", fmt.Errorf("URI does not match element template: %s", uri)
	}

	tagName := strings.TrimPrefix(uri, prefix)
	if tagName == "" || tagName == "{tagName}" {
		return "", fmt.Errorf("no tag name specified in URI: %s", uri)
	}

	return tagName, nil
}

// createElementDocumentation creates comprehensive documentation for an element
func createElementDocumentation(element types.ElementInfo) map[string]any {
	doc := map[string]any{
		"tagName":       element.TagName(),
		"description":   element.Description(),
		"attributes":    createAttributeDocumentation(element.Attributes()),
		"slots":         createSlotDocumentation(element.Slots()),
		"events":        createEventDocumentation(element.Events()),
		"cssProperties": createCSSPropertyDocumentation(element.CssProperties()),
		"cssParts":      createCSSPartDocumentation(element.CssParts()),
		"cssStates":     createCSSStateDocumentation(element.CssStates()),
		"examples":      createExampleDocumentation(element.Examples()),
		"usagePatterns": createUsagePatterns(element),
		"accessibility": createAccessibilityGuidance(element),
		"integration":   createIntegrationGuidance(element),
	}

	return doc
}

// createAttributeDocumentation creates detailed attribute documentation
func createAttributeDocumentation(attributes []types.Attribute) []map[string]any {
	result := make([]map[string]any, len(attributes))

	for i, attr := range attributes {
		attrDoc := map[string]any{
			"name":        attr.Name(),
			"description": attr.Description(),
			"type":        attr.Type(),
			"default":     attr.Default(),
			"required":    attr.Required(),
			"examples":    generateAttributeExamples(attr),
			"validation":  generateAttributeValidation(attr),
		}

		result[i] = attrDoc
	}

	return result
}

// createSlotDocumentation creates detailed slot documentation
func createSlotDocumentation(slots []types.Slot) []map[string]any {
	result := make([]map[string]any, len(slots))

	for i, slot := range slots {
		slotDoc := map[string]any{
			"name":        slot.Name(),
			"description": slot.Description(),
			"examples":    generateSlotExamples(slot),
			"guidelines":  generateSlotGuidelines(slot),
		}

		result[i] = slotDoc
	}

	return result
}

// createEventDocumentation creates detailed event documentation
func createEventDocumentation(events []types.Event) []map[string]any {
	result := make([]map[string]any, len(events))

	for i, event := range events {
		eventDoc := map[string]any{
			"name":        event.Name(),
			"description": event.Description(),
			"type":        event.Type(),
			"examples":    generateEventExamples(event),
			"usage":       generateEventUsage(event),
		}

		result[i] = eventDoc
	}

	return result
}

// createCSSPropertyDocumentation creates detailed CSS property documentation
func createCSSPropertyDocumentation(props []types.CssProperty) []map[string]any {
	result := make([]map[string]any, len(props))

	for i, prop := range props {
		propDoc := map[string]any{
			"name":        prop.Name(),
			"description": prop.Description(),
			"syntax":      prop.Syntax(),
			"initial":     prop.Initial(),
			"inherits":    prop.Inherits(),
			"examples":    generateCSSPropertyExamples(prop),
			"integration": generateCSSPropertyIntegration(prop),
		}

		result[i] = propDoc
	}

	return result
}

// createCSSPartDocumentation creates detailed CSS parts documentation
func createCSSPartDocumentation(parts []types.CssPart) []map[string]any {
	result := make([]map[string]any, len(parts))

	for i, part := range parts {
		partDoc := map[string]any{
			"name":        part.Name(),
			"description": part.Description(),
			"examples":    generateCSSPartExamples(part),
			"selectors":   generateCSSPartSelectors(part),
		}

		result[i] = partDoc
	}

	return result
}

// createCSSStateDocumentation creates detailed CSS states documentation
func createCSSStateDocumentation(states []types.CssState) []map[string]any {
	result := make([]map[string]any, len(states))

	for i, state := range states {
		stateDoc := map[string]any{
			"name":        state.Name(),
			"description": state.Description(),
			"examples":    generateCSSStateExamples(state),
			"selectors":   generateCSSStateSelectors(state),
		}

		result[i] = stateDoc
	}

	return result
}

// createExampleDocumentation creates enhanced example documentation
func createExampleDocumentation(examples []types.Example) []map[string]any {
	result := make([]map[string]any, len(examples))

	for i, example := range examples {
		exampleDoc := map[string]any{
			"title":       example.Title(),
			"description": example.Description(),
			"code":        example.Code(),
			"context":     generateExampleContext(example),
			"variations":  generateExampleVariations(example),
		}

		result[i] = exampleDoc
	}

	return result
}

// Helper functions for generating additional documentation

func generateAttributeExamples(attr types.Attribute) []string {
	// Generate example attribute usage based on type and name
	examples := []string{
		fmt.Sprintf(`<%s %s="%s"></%s>`, "my-element", attr.Name(), getExampleValue(attr), "my-element"),
	}

	if attr.Default() != "" {
		examples = append(examples, fmt.Sprintf("<!-- Default value: %s -->", attr.Default()))
	}

	return examples
}

func generateAttributeValidation(attr types.Attribute) map[string]any {
	validation := map[string]any{
		"type":     attr.Type(),
		"required": attr.Required(),
	}

	if attr.Default() != "" {
		validation["default"] = attr.Default()
	}

	return validation
}

func generateSlotExamples(slot types.Slot) []string {
	slotName := slot.Name()
	if slotName == "" {
		slotName = "default"
	}

	examples := []string{
		fmt.Sprintf(`<my-element>\n  <div slot="%s">Content for %s slot</div>\n</my-element>`, slotName, slotName),
	}

	return examples
}

func generateSlotGuidelines(slot types.Slot) []string {
	guidelines := []string{
		"Use semantic HTML elements when possible",
		"Consider accessibility when adding slotted content",
	}

	if slot.Description() != "" {
		guidelines = append(guidelines, fmt.Sprintf("Purpose: %s", slot.Description()))
	}

	return guidelines
}

func generateEventExamples(event types.Event) []string {
	eventName := event.Name()

	examples := []string{
		fmt.Sprintf(`element.addEventListener('%s', (e) => {\n  console.log('Event details:', e.detail);\n});`, eventName),
		fmt.Sprintf(`<my-element on%s="handleEvent(event)"></my-element>`, titleCase.String(eventName)),
	}

	return examples
}

func generateEventUsage(event types.Event) map[string]any {
	usage := map[string]any{
		"bubbles":    true,  // Most custom events bubble
		"cancelable": false, // Most custom events are not cancelable
		"detail":     event.Type(),
	}

	return usage
}

func generateCSSPropertyExamples(prop types.CssProperty) []string {
	propName := prop.Name()
	initialValue := prop.Initial()
	if initialValue == "" {
		initialValue = "value"
	}

	examples := []string{
		fmt.Sprintf(`my-element {\n  %s: %s;\n}`, propName, initialValue),
		fmt.Sprintf(`:root {\n  %s: %s; /* Global override */\n}`, propName, initialValue),
	}

	return examples
}

func generateCSSPropertyIntegration(prop types.CssProperty) map[string]any {
	integration := map[string]any{
		"syntax":   prop.Syntax(),
		"inherits": prop.Inherits(),
		"initial":  prop.Initial(),
	}

	return integration
}

func generateCSSPartExamples(part types.CssPart) []string {
	partName := part.Name()

	examples := []string{
		fmt.Sprintf(`my-element::%s {\n  /* Styles for %s part */\n}`, partName, partName),
	}

	return examples
}

func generateCSSPartSelectors(part types.CssPart) []string {
	partName := part.Name()

	selectors := []string{
		fmt.Sprintf("my-element::%s", partName),
		fmt.Sprintf("my-element::%s:hover", partName),
		fmt.Sprintf("my-element::%s:focus", partName),
	}

	return selectors
}

func generateCSSStateExamples(state types.CssState) []string {
	stateName := state.Name()

	examples := []string{
		fmt.Sprintf(`my-element:%s {\n  /* Styles when in %s state */\n}`, stateName, stateName),
	}

	return examples
}

func generateCSSStateSelectors(state types.CssState) []string {
	stateName := state.Name()

	selectors := []string{
		fmt.Sprintf("my-element:%s", stateName),
		fmt.Sprintf("my-element:%s:hover", stateName),
	}

	return selectors
}

func generateExampleContext(example types.Example) map[string]any {
	context := map[string]any{
		"title":       example.Title(),
		"description": example.Description(),
		"useCase":     "General usage example",
	}

	return context
}

func generateExampleVariations(
	example types.Example,
) []string {
	variations := []string{
		"Basic usage",
		"With additional attributes",
		"In different contexts",
	}

	return variations
}

func createUsagePatterns(element types.ElementInfo) map[string]any {
	patterns := map[string]any{
		"common": []string{
			fmt.Sprintf("<%s></%s>", element.TagName(), element.TagName()),
		},
		"withAttributes": generateCommonAttributePatterns(element),
		"withSlots":      generateCommonSlotPatterns(element),
	}

	return patterns
}

func createAccessibilityGuidance(
	element types.ElementInfo,
) map[string]any {
	guidance := map[string]any{
		"considerations": []string{
			"Ensure proper ARIA attributes when needed",
			"Maintain keyboard accessibility",
			"Provide appropriate labels and descriptions",
		},
		"patterns": []string{
			"Use semantic HTML within slots",
			"Consider focus management",
		},
	}

	return guidance
}

func createIntegrationGuidance(element types.ElementInfo) map[string]any {
	guidance := map[string]any{
		"frameworks": map[string]any{
			"vanilla": fmt.Sprintf("document.createElement('%s')", element.TagName()),
			"lit":     fmt.Sprintf("html`<%s></%s>`", element.TagName(), element.TagName()),
			"react":   fmt.Sprintf("<%s></%s>", element.TagName(), element.TagName()),
		},
		"imports": []string{
			"Add appropriate import statements",
			"Ensure element definition is loaded",
		},
	}

	return guidance
}

func generateCommonAttributePatterns(element types.ElementInfo) []string {
	patterns := []string{}
	attributes := element.Attributes()

	if len(attributes) > 0 {
		// Generate pattern with first few attributes
		attrStr := ""
		for i, attr := range attributes {
			if i >= 3 { // Limit to first 3 attributes
				break
			}
			if i > 0 {
				attrStr += " "
			}
			attrStr += fmt.Sprintf(`%s="%s"`, attr.Name(), getExampleValue(attr))
		}

		patterns = append(patterns, fmt.Sprintf("<%s %s></%s>", element.TagName(), attrStr, element.TagName()))
	}

	return patterns
}

func generateCommonSlotPatterns(element types.ElementInfo) []string {
	patterns := []string{}
	slots := element.Slots()

	if len(slots) > 0 {
		for _, slot := range slots {
			slotName := slot.Name()
			if slotName == "" {
				slotName = "default"
			}

			content := fmt.Sprintf(`<div slot="%s">Content</div>`, slotName)
			if slotName == "default" {
				content = "<div>Default content</div>"
			}

			patterns = append(patterns, fmt.Sprintf("<%s>\n  %s\n</%s>", element.TagName(), content, element.TagName()))
		}
	}

	return patterns
}

func getExampleValue(attr types.Attribute) string {
	attrType := attr.Type()

	switch {
	case strings.Contains(strings.ToLower(attrType), "boolean"):
		return "true"
	case strings.Contains(strings.ToLower(attrType), "number"):
		return "42"
	case strings.Contains(strings.ToLower(attrType), "string"):
		if attr.Default() != "" {
			return attr.Default()
		}
		return "example"
	default:
		if attr.Default() != "" {
			return attr.Default()
		}
		return "value"
	}
}
