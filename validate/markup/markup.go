/*
Copyright © 2026 Benny Powers <web@bennypowers.com>

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
package markup

import (
	"fmt"
	"strings"

	"bennypowers.dev/cem/manifest"
)

// Rule identifies the validation rule that produced an Issue.
const (
	RuleUnknownElement             = "unknown-element"
	RuleUnknownAttribute           = "unknown-attribute"
	RuleInvalidAttributeValue      = "invalid-attribute-value"
	RuleUnknownSlot                = "unknown-slot"
	RuleDeprecatedElement          = "deprecated-element"
	RuleDeprecatedAttribute        = "deprecated-attribute"
	RuleDeprecatedSlot             = "deprecated-slot"
	RuleSlotContentGuideline       = "slot-content-guideline"
	RuleAttributeConflict          = "attribute-conflict"
	RuleContentAttributeRedundancy = "content-attribute-redundancy"
)

// Severity values reported on an Issue.
const (
	SeverityError   = "error"
	SeverityWarning = "warning"
	SeverityInfo    = "info"
)

// Confidence values reported on an Issue. Manifest-backed findings are high
// confidence; prose-derived or heuristic findings are low confidence.
const (
	ConfidenceHigh = "high"
	ConfidenceLow  = "low"
)

// Options controls which validations run.
//
// Format selects the parser. Supported values are "html" (the default), and
// the Jinja-family template formats "twig", "njk", "jinja", "jinja2", "j2",
// and "liquid". "hbs" and "handlebars" are parsed as HTML on a best-effort
// basis without linking the Handlebars grammar.
//
// ValidateAttrs enables attribute name, enum value, and deprecation checks.
// ValidateSlots enables slot name, slot deprecation, and slot content checks.
// Strict enables low-confidence findings inferred from prose or heuristics,
// such as slot content guidelines, attribute conflicts, and redundant
// content/attribute pairs.
type Options struct {
	Format        string
	ValidateSlots bool
	ValidateAttrs bool
	Strict        bool
}

// Issue is a single markup validation finding. Line is a one-based line
// number and Column is a one-based rune offset within that line. Tree-sitter's
// byte columns are converted so non-ASCII text does not shift positions.
type Issue struct {
	Rule       string `json:"rule"`
	Severity   string `json:"severity"`
	Confidence string `json:"confidence"`
	Message    string `json:"message"`
	Element    string `json:"element"`
	Line       int    `json:"line"`
	Column     int    `json:"column"`

	details map[string]string
}

// Details returns rule-specific structured data for consumers that render
// findings themselves. Keys vary by rule and the map may be nil or empty.
func (i Issue) Details() map[string]string {
	if len(i.details) == 0 {
		return nil
	}
	out := make(map[string]string, len(i.details))
	for k, v := range i.details {
		out[k] = v
	}
	return out
}

// Validate parses markup and reports manifest compliance issues for custom
// elements found in it. It does not use an LSP document manager or workspace.
//
// A nil pkg is an error. Unsupported formats are an error. Malformed or
// unbalanced template control flow is tolerated: validation covers the markup
// that can be parsed.
func Validate(pkg *manifest.Package, markup string, opts Options) ([]Issue, error) {
	if pkg == nil {
		return nil, fmt.Errorf("markup: nil manifest package")
	}

	family, err := formatFamily(opts.Format)
	if err != nil {
		return nil, err
	}

	elements, err := parseElements(markup, family)
	if err != nil {
		return nil, err
	}

	v := validator{
		pkg:      pkg,
		opts:     opts,
		template: family == formatJinja,
	}
	return v.validate(elements), nil
}

type validator struct {
	pkg  *manifest.Package
	opts Options
	// template marks Jinja-family input. Template expressions are dynamic
	// values, so enum checks skip them. Plain HTML keeps the historical MCP
	// behavior of validating the literal value.
	template bool
	issues   []Issue
}

func (v *validator) validate(elements []*elementNode) []Issue {
	for _, el := range elements {
		v.validateElement(el, nil)
	}
	return v.issues
}

func (v *validator) validateElement(el *elementNode, parent *elementNode) {
	// A slot attribute is consumed by the element's immediate parent, which
	// may be a standard element such as a span.
	if v.opts.ValidateSlots && parent != nil {
		v.validateSlotAttribute(el, parent)
	}

	if !isCustomElementTag(el.tagName) {
		v.walkChildren(el, nil)
		return
	}

	decl := v.pkg.FindCustomElementDeclaration(el.tagName)
	if decl == nil {
		v.add(el, RuleUnknownElement, SeverityError, ConfidenceHigh,
			fmt.Sprintf("Custom element '%s' not found in registry", el.tagName), nil)
		v.walkChildren(el, nil)
		return
	}

	if decl.IsDeprecated() {
		message := fmt.Sprintf("Custom element '%s' is deprecated", el.tagName)
		if reason := deprecationReason(decl.Deprecated); reason != "" {
			message += ": " + reason
		}
		v.add(el, RuleDeprecatedElement, SeverityWarning, ConfidenceHigh, message, nil)
	}

	if v.opts.ValidateAttrs {
		v.validateAttributes(el, decl)
	}
	if v.opts.ValidateSlots && v.opts.Strict {
		v.validateSlotContentGuidelines(el, decl)
	}
	if v.opts.Strict && v.opts.ValidateAttrs {
		v.validateAttributeCombinations(el)
	}
	if v.opts.Strict && v.opts.ValidateSlots {
		v.validateContentAttributeRedundancy(el, decl)
	}

	v.walkChildren(el, el)
}

func (v *validator) walkChildren(el *elementNode, customParent *elementNode) {
	for _, child := range el.children {
		v.validateElement(child, customParent)
	}
}

func (v *validator) validateAttributes(el *elementNode, decl *manifest.CustomElementDeclaration) {
	attributes := decl.Attributes()
	for _, attr := range el.attributes {
		if isGlobalAttribute(attr.name) {
			continue
		}

		manifestAttr := findAttribute(attributes, attr.name)
		if manifestAttr == nil {
			message := fmt.Sprintf("Unknown attribute '%s' for element <%s>", attr.name, el.tagName)
			details := map[string]string{"attribute": attr.name}
			if suggestion := suggestAttributeCorrection(attr.name, attributes); suggestion != "" {
				message += fmt.Sprintf(". Did you mean '%s'?", suggestion)
				details["suggestion"] = suggestion
			}
			v.addAttribute(el, attr, RuleUnknownAttribute, SeverityError, ConfidenceHigh, message, details)
			continue
		}

		if manifestAttr.IsDeprecated() {
			message := fmt.Sprintf("Attribute '%s' on '%s' is deprecated", attr.name, el.tagName)
			details := map[string]string{"attribute": attr.name}
			if reason := deprecationReason(manifestAttr.Deprecated); reason != "" {
				message += ": " + reason
				details["reason"] = reason
			}
			v.addAttribute(el, attr, RuleDeprecatedAttribute, SeverityWarning, ConfidenceHigh, message, details)
		}

		if manifestAttr.IsEnum() && attr.value != "" && (!v.template || !containsTemplateExpression(attr.value)) {
			if !manifestAttr.IsValidValue(attr.value) {
				validValues := manifestAttr.EnumValues()
				message := fmt.Sprintf("Invalid attribute value: `%s=\"%s\"` in `<%s>`. Valid values: %v",
					attr.name, attr.value, el.tagName, validValues)
				v.addAttribute(el, attr, RuleInvalidAttributeValue, SeverityError, ConfidenceHigh, message, map[string]string{
					"attribute": attr.name,
					"value":     attr.value,
					"expected":  fmt.Sprintf("%v", validValues),
				})
			}
		}
	}
}

func (v *validator) validateSlotAttribute(el, parent *elementNode) {
	slotAttr := findParsedAttribute(el.attributes, "slot")
	if slotAttr == nil || slotAttr.value == "" || (v.template && containsTemplateExpression(slotAttr.value)) {
		return
	}

	parentDecl := v.pkg.FindCustomElementDeclaration(parent.tagName)
	if parentDecl == nil {
		return
	}

	slots := parentDecl.Slots()
	if len(slots) == 0 {
		return
	}

	for _, slot := range slots {
		if slot.Name != slotAttr.value {
			continue
		}
		if slot.IsDeprecated() {
			message := fmt.Sprintf("Slot '%s' on '%s' is deprecated", slot.Name, parent.tagName)
			details := map[string]string{"slot": slot.Name, "parent": parent.tagName}
			if reason := deprecationReason(slot.Deprecated); reason != "" {
				message += ": " + reason
				details["reason"] = reason
			}
			v.addAttribute(el, *slotAttr, RuleDeprecatedSlot, SeverityWarning, ConfidenceHigh, message, details)
		}
		return
	}

	names := make([]string, len(slots))
	for i, slot := range slots {
		names[i] = slot.Name
	}
	message := fmt.Sprintf("Unknown slot '%s' for element <%s>. Available slots: '%s'",
		slotAttr.value, parent.tagName, strings.Join(names, "', '"))
	v.addAttribute(el, *slotAttr, RuleUnknownSlot, SeverityError, ConfidenceHigh, message, map[string]string{
		"slot":   slotAttr.value,
		"parent": parent.tagName,
	})
}

func (v *validator) validateSlotContentGuidelines(el *elementNode, decl *manifest.CustomElementDeclaration) {
	for _, slot := range decl.Slots() {
		description := slot.Description
		if description == "" {
			continue
		}

		lower := strings.ToLower(description)
		if !strings.Contains(lower, "only") &&
			!strings.Contains(lower, "must") &&
			!strings.Contains(lower, "should not") &&
			!strings.Contains(lower, "avoid") {
			continue
		}
		if !strings.Contains(lower, "interactive") {
			continue
		}

		message := "Potential interactive content guideline - check slot content manually"
		v.add(el, RuleSlotContentGuideline, SeverityWarning, ConfidenceLow, message, map[string]string{
			"slot":      slot.Name,
			"guideline": description,
			"violation": message,
		})
	}
}

func (v *validator) validateAttributeCombinations(el *elementNode) {
	attributes := make(map[string]string, len(el.attributes))
	for _, attr := range el.attributes {
		attributes[attr.name] = attr.value
	}

	if loading, ok := attributes["loading"]; ok {
		if lazy, ok := attributes["lazy"]; ok && loading == "eager" && (lazy == "true" || lazy == "") {
			v.add(el, RuleAttributeConflict, SeverityWarning, ConfidenceLow,
				"Eager loading conflicts with lazy loading", map[string]string{
					"attribute1": "loading",
					"value1":     loading,
					"attribute2": "lazy",
					"value2":     lazy,
					"reason":     "Eager loading conflicts with lazy loading",
				})
		}
	}

	if disabled, ok := attributes["disabled"]; ok {
		if interactive, ok := attributes["interactive"]; ok &&
			(disabled == "true" || disabled == "") && (interactive == "true" || interactive == "") {
			v.add(el, RuleAttributeConflict, SeverityWarning, ConfidenceLow,
				"Disabled elements cannot be interactive", map[string]string{
					"attribute1": "disabled",
					"value1":     disabled,
					"attribute2": "interactive",
					"value2":     interactive,
					"reason":     "Disabled elements cannot be interactive",
				})
		}
	}
}

func (v *validator) validateContentAttributeRedundancy(el *elementNode, decl *manifest.CustomElementDeclaration) {
	attributes := make(map[string]string, len(el.attributes))
	for _, attr := range el.attributes {
		attributes[attr.name] = attr.value
	}

	for _, slot := range decl.Slots() {
		if slot.Name != "label" && slot.Name != "title" && slot.Name != "icon" {
			continue
		}
		value, ok := attributes[slot.Name]
		if !ok {
			continue
		}
		v.add(el, RuleContentAttributeRedundancy, SeverityInfo, ConfidenceLow,
			fmt.Sprintf("Attribute '%s' on <%s> is redundant with the '%s' slot", slot.Name, el.tagName, slot.Name),
			map[string]string{
				"attribute": slot.Name,
				"value":     value,
				"slot":      slot.Name,
				"content":   "[slot content not analyzed - tree-sitter enhancement needed]",
			})
	}
}

func (v *validator) add(el *elementNode, rule, severity, confidence, message string, details map[string]string) {
	v.issues = append(v.issues, Issue{
		Rule:       rule,
		Severity:   severity,
		Confidence: confidence,
		Message:    message,
		Element:    el.tagName,
		Line:       el.line,
		Column:     el.column,
		details:    details,
	})
}

func (v *validator) addAttribute(el *elementNode, attr parsedAttribute, rule, severity, confidence, message string, details map[string]string) {
	v.issues = append(v.issues, Issue{
		Rule:       rule,
		Severity:   severity,
		Confidence: confidence,
		Message:    message,
		Element:    el.tagName,
		Line:       attr.line,
		Column:     attr.column,
		details:    details,
	})
}

func findAttribute(attributes []manifest.Attribute, name string) *manifest.Attribute {
	for i := range attributes {
		if attributes[i].Name == name {
			return &attributes[i]
		}
	}
	return nil
}

func findParsedAttribute(attributes []parsedAttribute, name string) *parsedAttribute {
	for i := range attributes {
		if attributes[i].name == name {
			return &attributes[i]
		}
	}
	return nil
}

func deprecationReason(deprecated manifest.Deprecated) string {
	if deprecated == nil {
		return ""
	}
	if reason, ok := deprecated.Value().(string); ok {
		return reason
	}
	return ""
}

func containsTemplateExpression(value string) bool {
	return strings.Contains(value, "{{") ||
		strings.Contains(value, "{%") ||
		strings.Contains(value, "${")
}
