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

package routes

import (
	"bytes"
	"fmt"
	"html/template"
	"regexp"
	"strings"

	M "bennypowers.dev/cem/manifest"
	"golang.org/x/net/html"
)

// KnobType represents the type of knob control
type KnobType string

const (
	KnobTypeBoolean    KnobType = "boolean"
	KnobTypeString     KnobType = "string"
	KnobTypeNumber     KnobType = "number"
	KnobTypeEnum       KnobType = "enum"
	KnobTypeColor      KnobType = "color"
	KnobTypeUnknown    KnobType = "unknown"
)

// KnobCategory represents which category a knob belongs to
type KnobCategory string

const (
	KnobCategoryAttribute    KnobCategory = "attributes"
	KnobCategoryProperty     KnobCategory = "properties"
	KnobCategoryCSSProperty  KnobCategory = "css-properties"
)

// KnobsData represents all knobs for an element
type KnobsData struct {
	TagName          string
	ElementID        string // ID of the target element (for cem-serve-knob-group "for" attribute)
	AttributeKnobs   []KnobData
	PropertyKnobs    []KnobData
	CSSPropertyKnobs []KnobData
}

// KnobData represents a single knob control
type KnobData struct {
	Name         string
	Category     KnobCategory
	Type         KnobType
	CurrentValue string
	EnumValues   []string
	Summary      template.HTML // Brief summary for helper text
	Description  template.HTML // Full description for popover
	Default      string
}

// ElementInstance represents a discovered custom element instance in the demo HTML.
// Used by Phase 5b multi-instance knobs to track individual element instances.
type ElementInstance struct {
	Node        *html.Node
	TagName     string
	ID          string
	AriaLabel   string
	TextContent string
	Attributes  map[string]string
}

// ElementKnobGroup represents knobs for a single element instance.
// Used by Phase 5b to render multiple knob groups with labels and collapsible structure.
type ElementKnobGroup struct {
	TagName       string // Element tag name (e.g., "my-card")
	ElementID     string // ID of the element instance
	Label         string // Human-readable label for this instance (e.g., "#card-1", "User Settings")
	InstanceIndex int    // Per-tag-name instance index (0 for first rh-tab, 1 for second rh-tab, etc.)
	IsPrimary     bool   // Whether this is the primary instance (expanded by default)
	Knobs         *KnobsData
}

// GenerateKnobs generates knob controls from manifest and demo HTML
func GenerateKnobs(declaration *M.CustomElementDeclaration, demoHTML []byte, enabledKnobs string) (*KnobsData, error) {
	if declaration == nil {
		return nil, fmt.Errorf("declaration is nil")
	}

	// Parse demo HTML to extract current values
	currentValues := extractCurrentValues(declaration.TagName, demoHTML)

	// Try to discover element instances to get ID
	instances, err := discoverElementInstances(declaration.TagName, demoHTML)
	var elementID string
	if err == nil && len(instances) > 0 {
		// Use first instance's ID
		elementID = instances[0].ID
	}

	// Auto-generate ID if not present
	if elementID == "" {
		elementID = fmt.Sprintf("%s-0", declaration.TagName)
	}

	// Delegate to generateKnobsForInstance
	knobs, err := generateKnobsForInstance(declaration, currentValues, enabledKnobs)
	if err != nil {
		return nil, err
	}

	// Set ElementID on knobs
	knobs.ElementID = elementID

	return knobs, nil
}

// parseEnabledKnobs converts a space-separated list into a map
func parseEnabledKnobs(enabledKnobs string) map[KnobCategory]bool {
	enabled := make(map[KnobCategory]bool)
	if enabledKnobs == "" {
		return enabled
	}

	parts := strings.Fields(enabledKnobs)
	for _, part := range parts {
		switch KnobCategory(part) {
		case KnobCategoryAttribute:
			enabled[KnobCategoryAttribute] = true
		case KnobCategoryProperty:
			enabled[KnobCategoryProperty] = true
		case KnobCategoryCSSProperty:
			enabled[KnobCategoryCSSProperty] = true
		}
	}
	return enabled
}

// attributeToKnob converts a manifest attribute to a knob
func attributeToKnob(attr M.Attribute, currentValues map[string]string) KnobData {
	knob := KnobData{
		Name:         attr.Name,
		Category:     KnobCategoryAttribute,
		Summary:      template.HTML(attr.Summary),
		Description:  template.HTML(attr.Description),
		Default:      attr.Default,
		CurrentValue: currentValues[attr.Name],
	}

	// Determine knob type from CEM type
	if attr.Type != nil {
		knob.Type, knob.EnumValues = parseType(attr.Type.Text)
	} else {
		knob.Type = KnobTypeString
	}

	return knob
}

// propertyToKnob converts a manifest class field to a knob
func propertyToKnob(field *M.ClassField, currentValues map[string]string) KnobData {
	knob := KnobData{
		Name:         field.Name,
		Category:     KnobCategoryProperty,
		Summary:      template.HTML(field.Summary),
		Description:  template.HTML(field.Description),
		Default:      field.Default,
		CurrentValue: currentValues["prop:"+field.Name],
	}

	// Determine knob type from CEM type
	if field.Type != nil {
		knob.Type, knob.EnumValues = parseType(field.Type.Text)
	} else {
		knob.Type = KnobTypeUnknown
	}

	return knob
}

// cssPropertyToKnob converts a CSS custom property to a knob
func cssPropertyToKnob(cssProp M.CssCustomProperty, currentValues map[string]string) KnobData {
	knob := KnobData{
		Name:         cssProp.Name,
		Category:     KnobCategoryCSSProperty,
		Summary:      template.HTML(cssProp.Summary),
		Description:  template.HTML(cssProp.Description),
		Default:      cssProp.Default,
		CurrentValue: currentValues["css:"+cssProp.Name],
		Type:         detectCSSPropertyType(cssProp.Default),
	}

	return knob
}

// parseType parses a TypeScript type string and returns the knob type and enum values
func parseType(typeText string) (KnobType, []string) {
	if typeText == "" {
		return KnobTypeString, nil
	}

	// Remove whitespace for easier parsing
	normalized := strings.ReplaceAll(typeText, " ", "")

	// Check for boolean
	if normalized == "boolean" {
		return KnobTypeBoolean, nil
	}

	// Check for number
	if normalized == "number" {
		return KnobTypeNumber, nil
	}

	// Check for string literal union (enum): 'a' | 'b' | 'c'
	if strings.Contains(normalized, "|") && strings.Contains(normalized, "'") {
		// Extract string literals
		re := regexp.MustCompile(`'([^']+)'`)
		matches := re.FindAllStringSubmatch(normalized, -1)
		if len(matches) > 0 {
			var enumValues []string
			for _, match := range matches {
				if len(match) > 1 {
					enumValues = append(enumValues, match[1])
				}
			}
			return KnobTypeEnum, enumValues
		}
	}

	// Default to string
	return KnobTypeString, nil
}

// detectCSSPropertyType detects the type of a CSS property from its default value
func detectCSSPropertyType(defaultValue string) KnobType {
	if defaultValue == "" {
		return KnobTypeString
	}

	// Check if it looks like a color (hex, rgb, hsl, named color)
	if strings.HasPrefix(defaultValue, "#") ||
		strings.HasPrefix(defaultValue, "rgb") ||
		strings.HasPrefix(defaultValue, "hsl") {
		return KnobTypeColor
	}

	// Default to string for CSS values
	return KnobTypeString
}

// extractCurrentValues parses demo HTML and extracts current attribute/property values
func extractCurrentValues(tagName string, demoHTML []byte) map[string]string {
	values := make(map[string]string)

	doc, err := html.Parse(bytes.NewReader(demoHTML))
	if err != nil {
		return values
	}

	// Find the custom element in the HTML
	var findElement func(*html.Node)
	findElement = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == tagName {
			// Extract attributes
			for _, attr := range n.Attr {
				values[attr.Key] = attr.Val
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findElement(c)
		}
	}
	findElement(doc)

	return values
}


// firstTextContent extracts the first non-empty text node from an element's subtree.
// This handles nested content like <my-card><h2>Title</h2></my-card> correctly.
func firstTextContent(n *html.Node) string {
	var text string
	var walk func(*html.Node) bool
	walk = func(node *html.Node) bool {
		if node.Type == html.TextNode {
			trimmed := strings.TrimSpace(node.Data)
			if trimmed != "" {
				text = trimmed
				return true // Found text, stop walking
			}
		}
		// Recurse through children
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			if walk(c) {
				return true // Bubble up the stop signal
			}
		}
		return false
	}
	walk(n)
	return text
}

// discoverElementInstances finds all instances of a custom element in demo HTML
func discoverElementInstances(tagName string, demoHTML []byte) ([]ElementInstance, error) {
	var instances []ElementInstance

	doc, err := html.Parse(bytes.NewReader(demoHTML))
	if err != nil {
		return nil, fmt.Errorf("failed to parse demo HTML: %w", err)
	}

	// Walk the DOM tree in source order
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == tagName {
			instance := ElementInstance{
				Node:       n,
				TagName:    tagName,
				Attributes: make(map[string]string),
			}

			// Extract attributes
			for _, attr := range n.Attr {
				instance.Attributes[attr.Key] = attr.Val

				// Save important attributes for labeling
				switch attr.Key {
				case "id":
					instance.ID = attr.Val
				case "aria-label":
					instance.AriaLabel = attr.Val
				}
			}

			// Extract the first non-empty descendant text node, trimmed
			if text := firstTextContent(n); text != "" {
				instance.TextContent = text
			}

			instances = append(instances, instance)
		}

		// Recurse through children
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(doc)

	return instances, nil
}

// generateElementLabel creates a label for an element instance using the priority system:
// 1. ID (#id)
// 2. Text content (trimmed, max 20 chars)
// 3. aria-label
// 4. Fallback: "tag-name No. N"
func generateElementLabel(instance ElementInstance, tagName string, index int) string {
	// Priority 1: ID
	if instance.ID != "" {
		return "#" + instance.ID
	}

	// Priority 2: Text content (trimmed, max 20 chars)
	if instance.TextContent != "" {
		text := instance.TextContent
		runes := []rune(text)
		if len(runes) > 20 {
			// Truncate to 20 runes (preserves multi-byte characters like emoji/CJK)
			truncated := string(runes[:20])
			return strings.TrimSpace(truncated) + "…"
		}
		return text
	}

	// Priority 3: aria-label
	if instance.AriaLabel != "" {
		return instance.AriaLabel
	}

	// Priority 4: Fallback - "tag-name No. N" (1-indexed for display)
	return fmt.Sprintf("%s No. %d", tagName, index+1)
}

// GenerateMultiInstanceKnobs generates knobs for all instances of an element in demo HTML.
// This is the Phase 5b implementation that supports multiple element instances.
// Each instance gets its own ElementKnobGroup with a unique label based on:
// 1. ID attribute (#id)
// 2. Text content (trimmed to 20 chars)
// 3. aria-label
// 4. Fallback (tag-name No. N)
// The first instance is marked as primary (expanded by default in UI).
func GenerateMultiInstanceKnobs(declaration *M.CustomElementDeclaration, demoHTML []byte, enabledKnobs string) ([]ElementKnobGroup, error) {
	if declaration == nil {
		return nil, fmt.Errorf("declaration is nil")
	}

	// Discover all instances in demo HTML
	instances, err := discoverElementInstances(declaration.TagName, demoHTML)
	if err != nil {
		return nil, err
	}

	var groups []ElementKnobGroup

	for i, instance := range instances {
		// Generate label for this instance
		label := generateElementLabel(instance, declaration.TagName, i)

		// Generate or use existing ID
		elementID := instance.ID
		if elementID == "" {
			// Auto-generate ID if not present: tagname-instance-index
			elementID = fmt.Sprintf("%s-%d", declaration.TagName, i)
		}

		// Extract current values for this specific instance
		currentValues := instance.Attributes

		// Generate knobs for this instance using existing GenerateKnobs logic
		knobs, err := generateKnobsForInstance(declaration, currentValues, enabledKnobs)
		if err != nil {
			return nil, fmt.Errorf("failed to generate knobs for instance %d: %w", i, err)
		}

		// Set ElementID on knobs
		knobs.ElementID = elementID

		group := ElementKnobGroup{
			TagName:       declaration.TagName,
			ElementID:     elementID,
			Label:         label,
			InstanceIndex: i,
			IsPrimary:     i == 0, // First instance is primary
			Knobs:         knobs,
		}

		groups = append(groups, group)
	}

	return groups, nil
}

// generateKnobsForInstance is a helper that generates knobs for a single instance
// It's similar to GenerateKnobs but takes pre-extracted currentValues
func generateKnobsForInstance(declaration *M.CustomElementDeclaration, currentValues map[string]string, enabledKnobs string) (*KnobsData, error) {
	// Parse enabled knobs (space-separated list of categories)
	enabled := parseEnabledKnobs(enabledKnobs)

	knobs := &KnobsData{
		TagName: declaration.TagName,
	}

	// Build set of property names for deduplication
	propertyNames := make(map[string]bool)
	if enabled[KnobCategoryProperty] {
		for _, member := range declaration.Members {
			if field, ok := member.(*M.ClassField); ok {
				if field.Privacy == M.Public || field.Privacy == "" {
					propertyNames[field.Name] = true
				}
			}
		}
	}

	// Generate attribute knobs (skip if same-named property exists)
	if enabled[KnobCategoryAttribute] {
		seenAttrs := make(map[string]*M.Attribute)
		for i := range declaration.Attributes {
			attr := &declaration.Attributes[i]
			// Skip attribute if a property with the same name exists
			if propertyNames[attr.Name] {
				continue
			}

			// Deduplicate: keep attribute with more specific type
			if existing, exists := seenAttrs[attr.Name]; exists {
				// Prefer the one with a type over null
				if attr.Type != nil && (existing.Type == nil || existing.Type.Text == "") {
					seenAttrs[attr.Name] = attr
				}
				// Otherwise keep the first one
				continue
			}
			seenAttrs[attr.Name] = attr
		}

		// Convert deduplicated attributes to knobs
		for _, attr := range seenAttrs {
			knob := attributeToKnob(*attr, currentValues)
			knobs.AttributeKnobs = append(knobs.AttributeKnobs, knob)
		}
	}

	// Generate property knobs (from class fields)
	if enabled[KnobCategoryProperty] {
		for _, member := range declaration.Members {
			if field, ok := member.(*M.ClassField); ok {
				if field.Privacy == M.Public || field.Privacy == "" {
					knob := propertyToKnob(field, currentValues)
					knobs.PropertyKnobs = append(knobs.PropertyKnobs, knob)
				}
			}
		}
	}

	// Generate CSS custom property knobs
	if enabled[KnobCategoryCSSProperty] {
		for _, cssProp := range declaration.CssProperties {
			knob := cssPropertyToKnob(cssProp, currentValues)
			knobs.CSSPropertyKnobs = append(knobs.CSSPropertyKnobs, knob)
		}
	}

	return knobs, nil
}

// GenerateKnobsForAllElements discovers all custom elements in demo HTML and generates knobs for each.
// This supports compositional components (e.g., accordion + accordion-header + accordion-panel).
// Returns knob groups in depth-first source order (parent, then children left-to-right).
func GenerateKnobsForAllElements(pkg *M.Package, demoHTML []byte, enabledKnobs string) ([]ElementKnobGroup, error) {
	// Build a map of tag name -> declaration for quick lookup
	// Use RenderableCustomElementDeclarations to get ALL elements in package, not just ones with demos
	declarations := make(map[string]*M.CustomElementDeclaration)
	for _, renderableDecl := range pkg.RenderableCustomElementDeclarations() {
		declarations[renderableDecl.CustomElementDeclaration.TagName] = renderableDecl.CustomElementDeclaration
	}

	// Discover all custom element instances in depth-first source order
	instances, err := discoverAllCustomElementInstances(demoHTML)
	if err != nil {
		return nil, err
	}

	// Generate knobs for each instance in order
	var allGroups []ElementKnobGroup
	instanceCounts := make(map[string]int) // Track instance index per tag name

	for _, instance := range instances {
		declaration, ok := declarations[instance.TagName]
		if !ok {
			// Skip elements not in manifest
			continue
		}

		// Generate label for this instance
		instanceIndex := instanceCounts[instance.TagName]
		label := generateElementLabel(instance, instance.TagName, instanceIndex)
		instanceCounts[instance.TagName]++

		// Generate knobs for this instance
		knobs, err := generateKnobsForInstance(declaration, instance.Attributes, enabledKnobs)
		if err != nil {
			return nil, fmt.Errorf("failed to generate knobs for %s instance %d: %w", instance.TagName, instanceIndex, err)
		}

		// Generate or use existing ID
		elementID := instance.ID
		if elementID == "" {
			// Auto-generate ID if not present: tagname-instance-index
			elementID = fmt.Sprintf("%s-%d", instance.TagName, instanceIndex)
		}

		// Set ElementID on knobs
		knobs.ElementID = elementID

		group := ElementKnobGroup{
			TagName:       instance.TagName,
			ElementID:     elementID,
			Label:         label,
			InstanceIndex: instanceIndex,
			IsPrimary:     len(allGroups) == 0, // First instance overall is primary
			Knobs:         knobs,
		}

		allGroups = append(allGroups, group)
	}

	return allGroups, nil
}

// discoverAllCustomElementInstances finds all custom element instances in HTML
// in depth-first source order (parent first, then children left-to-right).
func discoverAllCustomElementInstances(demoHTML []byte) ([]ElementInstance, error) {
	doc, err := html.Parse(bytes.NewReader(demoHTML))
	if err != nil {
		return nil, fmt.Errorf("failed to parse demo HTML: %w", err)
	}

	var instances []ElementInstance

	// Walk DOM tree in depth-first order
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && strings.Contains(n.Data, "-") {
			// This is a custom element
			instance := ElementInstance{
				Node:       n,
				TagName:    n.Data,
				Attributes: make(map[string]string),
			}

			// Extract attributes
			for _, attr := range n.Attr {
				instance.Attributes[attr.Key] = attr.Val

				// Save important attributes for labeling
				switch attr.Key {
				case "id":
					instance.ID = attr.Val
				case "aria-label":
					instance.AriaLabel = attr.Val
				}
			}

			// Extract first non-empty text content
			if text := firstTextContent(n); text != "" {
				instance.TextContent = text
			}

			instances = append(instances, instance)
		}

		// Recurse through children (depth-first)
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(doc)

	return instances, nil
}

// RenderKnobsHTML renders the knobs HTML template.
// Takes a slice of ElementKnobGroup (can be single or multiple instances).
// Each group is rendered in a pf-v6-card with tag name and instance label.
// Knobs are organized by category (attributes, properties, CSS properties) within each group.
func RenderKnobsHTML(knobGroups []ElementKnobGroup) (template.HTML, error) {
	if len(knobGroups) == 0 {
		return "", nil
	}

	// Convert markdown summary and description to HTML for all knobs in all groups
	for gi := range knobGroups {
		if knobGroups[gi].Knobs == nil {
			continue
		}

		for i := range knobGroups[gi].Knobs.AttributeKnobs {
			if knobGroups[gi].Knobs.AttributeKnobs[i].Summary != "" {
				knobGroups[gi].Knobs.AttributeKnobs[i].Summary = template.HTML(markdownToHTML(string(knobGroups[gi].Knobs.AttributeKnobs[i].Summary)))
			}
			if knobGroups[gi].Knobs.AttributeKnobs[i].Description != "" {
				knobGroups[gi].Knobs.AttributeKnobs[i].Description = template.HTML(markdownToHTML(string(knobGroups[gi].Knobs.AttributeKnobs[i].Description)))
			}
		}
		for i := range knobGroups[gi].Knobs.PropertyKnobs {
			if knobGroups[gi].Knobs.PropertyKnobs[i].Summary != "" {
				knobGroups[gi].Knobs.PropertyKnobs[i].Summary = template.HTML(markdownToHTML(string(knobGroups[gi].Knobs.PropertyKnobs[i].Summary)))
			}
			if knobGroups[gi].Knobs.PropertyKnobs[i].Description != "" {
				knobGroups[gi].Knobs.PropertyKnobs[i].Description = template.HTML(markdownToHTML(string(knobGroups[gi].Knobs.PropertyKnobs[i].Description)))
			}
		}
		for i := range knobGroups[gi].Knobs.CSSPropertyKnobs {
			if knobGroups[gi].Knobs.CSSPropertyKnobs[i].Summary != "" {
				knobGroups[gi].Knobs.CSSPropertyKnobs[i].Summary = template.HTML(markdownToHTML(string(knobGroups[gi].Knobs.CSSPropertyKnobs[i].Summary)))
			}
			if knobGroups[gi].Knobs.CSSPropertyKnobs[i].Description != "" {
				knobGroups[gi].Knobs.CSSPropertyKnobs[i].Description = template.HTML(markdownToHTML(string(knobGroups[gi].Knobs.CSSPropertyKnobs[i].Description)))
			}
		}
	}

	var buf bytes.Buffer
	err := KnobsTemplate.Execute(&buf, knobGroups)
	if err != nil {
		return "", err
	}

	return template.HTML(buf.String()), nil
}
