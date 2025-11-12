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
	TagName         string
	AttributeKnobs  []KnobData
	PropertyKnobs   []KnobData
	CSSPropertyKnobs []KnobData
}

// KnobData represents a single knob control
type KnobData struct {
	Name         string
	Category     KnobCategory
	Type         KnobType
	CurrentValue string
	EnumValues   []string
	Description  template.HTML
	Default      string
}

// GenerateKnobs generates knob controls from manifest and demo HTML
func GenerateKnobs(declaration *M.CustomElementDeclaration, demoHTML []byte, enabledKnobs string) (*KnobsData, error) {
	if declaration == nil {
		return nil, fmt.Errorf("declaration is nil")
	}

	// Parse enabled knobs (space-separated list of categories)
	enabled := parseEnabledKnobs(enabledKnobs)

	knobs := &KnobsData{
		TagName: declaration.TagName,
	}

	// Parse demo HTML to extract current values
	currentValues := extractCurrentValues(declaration.TagName, demoHTML)

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
	// Also deduplicate by name, preferring attributes with more specific types
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

// RenderKnobsHTML renders the knobs HTML template
func RenderKnobsHTML(knobs *KnobsData) (template.HTML, error) {
	if knobs == nil {
		return "", nil
	}

	// Convert markdown descriptions to HTML for all knobs
	for i := range knobs.AttributeKnobs {
		if knobs.AttributeKnobs[i].Description != "" {
			knobs.AttributeKnobs[i].Description = template.HTML(markdownToHTML(string(knobs.AttributeKnobs[i].Description)))
		}
	}
	for i := range knobs.PropertyKnobs {
		if knobs.PropertyKnobs[i].Description != "" {
			knobs.PropertyKnobs[i].Description = template.HTML(markdownToHTML(string(knobs.PropertyKnobs[i].Description)))
		}
	}
	for i := range knobs.CSSPropertyKnobs {
		if knobs.CSSPropertyKnobs[i].Description != "" {
			knobs.CSSPropertyKnobs[i].Description = template.HTML(markdownToHTML(string(knobs.CSSPropertyKnobs[i].Description)))
		}
	}

	var buf bytes.Buffer
	err := KnobsTemplate.Execute(&buf, knobs)
	if err != nil {
		return "", err
	}

	return template.HTML(buf.String()), nil
}

// markdownToHTML converts markdown text to HTML
func markdownToHTML(text string) string {
	var buf bytes.Buffer
	// Use the same markdown renderer as chrome.go
	err := md.Convert([]byte(text), &buf)
	if err != nil {
		return text // Return original text on error
	}
	return buf.String()
}
