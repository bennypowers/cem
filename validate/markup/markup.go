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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	"bennypowers.dev/cem/manifest"
)

// Errors returned by Validate and FindCustomElements.
var (
	ErrNilPackage        = errors.New("markup: nil manifest package")
	ErrUnsupportedFormat = errors.New("markup: unsupported format")
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

// Detail keys used in Issue.Details. Each rule reports a subset of these keys;
// consumers such as the MCP tool should use the constants rather than string
// literals so a rename is a compile error.
const (
	DetailAttribute  = "attribute"
	DetailSuggestion = "suggestion"
	DetailValue      = "value"
	DetailExpected   = "expected"
	DetailSlot       = "slot"
	DetailParent     = "parent"
	DetailGuideline  = "guideline"
	DetailReason     = "reason"
	DetailAttribute1 = "attribute1"
	DetailValue1     = "value1"
	DetailAttribute2 = "attribute2"
	DetailValue2     = "value2"
	DetailContent    = "content"
)

// Format selects the markup or template dialect parsed by [Validate] and
// [FindCustomElements]. The zero value means [FormatHTML].
type Format string

// Supported formats. Aliases are kept for callers that already pass the
// extension or dialect name.
const (
	FormatHTML       Format = "html"
	FormatHTM        Format = "htm"
	FormatTwig       Format = "twig"
	FormatNJK        Format = "njk"
	FormatJinja      Format = "jinja"
	FormatJinja2     Format = "jinja2"
	FormatJ2         Format = "j2"
	FormatLiquid     Format = "liquid"
	FormatHandlebars Format = "handlebars"
	FormatHBS        Format = "hbs"
)

// Options controls which validations run. Unknown element and deprecated
// element checks always run; the remaining rule families require a flag.
//
// Format selects the parser. The zero value means [FormatHTML]. The Jinja
// family covers [FormatTwig], [FormatNJK], [FormatJinja], [FormatJinja2], and
// [FormatJ2]. [FormatHBS] and [FormatHandlebars] are parsed as HTML on a
// best-effort basis without linking the Handlebars grammar, but their
// expression values keep template semantics and are not validated literally.
//
// ValidateAttrs enables attribute name, enum value, and attribute deprecation
// checks. ValidateSlots enables slot name and slot deprecation checks.
//
// Strict enables low-confidence findings, but only for a rule family whose
// category flag is also set: slot content guidelines require ValidateSlots,
// attribute conflicts require ValidateAttrs, and content/attribute redundancy
// requires ValidateSlots. Strict alone adds no findings.
type Options struct {
	Format        Format
	ValidateSlots bool
	ValidateAttrs bool
	Strict        bool
}

// Issue is a single markup validation finding. Line is a one-based line
// number and Column is a zero-based UTF-16 code-unit offset within that line,
// matching LSP positions.
//
// Element and the position fields identify the reported token per rule:
// element rules and the slot content/redundancy heuristics report the custom
// element's tag name; attribute rules report the attribute name on that
// element; unknown-slot and deprecated-slot report the slotted child's tag
// name and the position of its `slot` attribute. For those slot rules, the
// owning custom element is available as Details[DetailParent].
type Issue struct {
	Rule       string `json:"rule"`
	Severity   string `json:"severity"`
	Confidence string `json:"confidence"`
	Message    string `json:"message"`
	Element    string `json:"element"`
	Line       int    `json:"line"`
	Column     int    `json:"column"`

	// Details holds rule-specific structured data. Keys are the Detail*
	// constants and the map may be nil or empty. Each rule reports:
	//
	//   - unknown-attribute: DetailAttribute, plus DetailSuggestion when a close
	//     match exists.
	//   - invalid-attribute-value: DetailAttribute, DetailValue, DetailExpected.
	//   - deprecated-attribute: DetailAttribute, plus DetailReason when present.
	//   - deprecated-element: no details.
	//   - unknown-slot: DetailSlot, DetailParent.
	//   - deprecated-slot: DetailSlot, DetailParent, plus DetailReason when
	//     present.
	//   - slot-content-guideline: DetailSlot, DetailGuideline.
	//   - attribute-conflict: DetailAttribute1, DetailValue1, DetailAttribute2,
	//     DetailValue2, DetailReason.
	//   - content-attribute-redundancy: DetailAttribute, DetailValue, DetailSlot,
	//     DetailContent.
	Details map[string]string `json:"details,omitempty"`
}

// MarshalJSON leaves HTML characters unescaped because findings are
// code-oriented.
func (i Issue) MarshalJSON() ([]byte, error) {
	type issueAlias Issue

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(issueAlias(i)); err != nil {
		return nil, err
	}
	return bytes.TrimSuffix(buf.Bytes(), []byte{'\n'}), nil
}

// Element is a custom element occurrence found in markup. Line is one-based
// and Column is a zero-based UTF-16 code-unit offset within that line.
type Element struct {
	TagName string `json:"tagName"`
	Line    int    `json:"line"`
	Column  int    `json:"column"`
}

// utf16Column converts a zero-based UTF-8 byte column into a zero-based
// UTF-16 code-unit column, matching LSP positions. Astral characters count as
// two units and invalid bytes as one.
func utf16Column(source []byte, lineStart, byteColumn int) int {
	if lineStart < 0 {
		lineStart = 0
	}
	end := lineStart + byteColumn
	if end > len(source) {
		end = len(source)
	}

	units := 0
	for i := lineStart; i < end; {
		r, size := utf8.DecodeRune(source[i:end])
		if size == 0 {
			break
		}
		if r > 0xFFFF {
			units += 2
		} else {
			units++
		}
		i += size
	}
	return units
}

// FindCustomElements returns custom element occurrences in source order,
// without validating them. Tag names are normalized to ASCII lowercase.
func FindCustomElements(markup string, opts Options) ([]Element, error) {
	family, err := formatFamily(opts.Format)
	if err != nil {
		return nil, err
	}

	nodes, err := parseElements(markup, family)
	if err != nil {
		return nil, err
	}

	var elements []Element
	source := []byte(markup)
	var walk func([]*elementNode)
	walk = func(nodes []*elementNode) {
		for _, node := range nodes {
			if isCustomElementTag(node.tagName) && !node.dynamic {
				elements = append(elements, Element{
					TagName: node.tagName,
					Line:    node.line,
					Column:  utf16Column(source, node.lineStart, node.column),
				})
			}
			walk(node.children)
		}
	}
	walk(nodes)
	return elements, nil
}

// Validate parses markup and reports manifest compliance issues for custom
// elements found in it. It does not use an LSP document manager or workspace.
//
// A nil pkg is an error. Unsupported formats are an error. Malformed or
// unbalanced template control flow is tolerated: validation covers the markup
// that can be parsed.
func Validate(pkg *manifest.Package, markup string, opts Options) ([]Issue, error) {
	if pkg == nil {
		return nil, ErrNilPackage
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
		pkg:        pkg,
		opts:       opts,
		source:     []byte(markup),
		template:   family != formatHTML,
		handlebars: family == formatHandlebars,
	}
	return v.validate(elements), nil
}

type validator struct {
	pkg    *manifest.Package
	opts   Options
	source []byte
	// template marks Jinja-family or Handlebars input. Template expressions
	// are dynamic values, so enum checks skip them. Plain HTML keeps the
	// historical MCP behavior of validating the literal value.
	template   bool
	handlebars bool
	issues     []Issue
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

	// A tag name built from template syntax cannot be resolved against the
	// manifest, so skip the element itself but still inspect its children.
	if el.dynamic {
		v.walkChildren(el, nil)
		return
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
		v.validateAttributeCombinations(el, decl)
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
		if attr.nameDynamic {
			// The attribute name depends on template data, so it cannot be
			// matched against the manifest.
			continue
		}
		if isGlobalAttribute(attr.name) {
			continue
		}

		manifestAttr := findAttribute(attributes, attr.name)
		if manifestAttr == nil {
			message := fmt.Sprintf("Unknown attribute '%s' for element <%s>", attr.name, el.tagName)
			details := map[string]string{DetailAttribute: attr.name}
			if suggestion := suggestAttributeCorrection(attr.name, attributes); suggestion != "" {
				message += fmt.Sprintf(". Did you mean '%s'?", suggestion)
				details[DetailSuggestion] = suggestion
			}
			v.addAttribute(el, attr, RuleUnknownAttribute, SeverityError, ConfidenceHigh, message, details)
			continue
		}

		if manifestAttr.IsDeprecated() {
			message := fmt.Sprintf("Attribute '%s' on '%s' is deprecated", attr.name, el.tagName)
			details := map[string]string{DetailAttribute: attr.name}
			if reason := deprecationReason(manifestAttr.Deprecated); reason != "" {
				message += ": " + reason
				details[DetailReason] = reason
			}
			v.addAttribute(el, attr, RuleDeprecatedAttribute, SeverityWarning, ConfidenceHigh, message, details)
		}

		if !attr.dynamic && manifestAttr.IsEnum() && attr.value != "" &&
			(!v.template || !containsTemplateExpression(attr.value, v.handlebars)) {
			if !manifestAttr.IsValidValue(attr.value) {
				validValues := manifestAttr.EnumValues()
				message := fmt.Sprintf("Invalid attribute value: `%s=\"%s\"` in `<%s>`. Valid values: %v",
					attr.name, attr.value, el.tagName, validValues)
				v.addAttribute(el, attr, RuleInvalidAttributeValue, SeverityError, ConfidenceHigh, message, map[string]string{
					DetailAttribute: attr.name,
					DetailValue:     attr.value,
					DetailExpected:  fmt.Sprintf("%v", validValues),
				})
			}
		}
	}
}

func (v *validator) validateSlotAttribute(el, parent *elementNode) {
	slotAttr := findParsedAttribute(el.attributes, "slot")
	if slotAttr == nil || slotAttr.nameDynamic || slotAttr.value == "" || slotAttr.dynamic ||
		(v.template && containsTemplateExpression(slotAttr.value, v.handlebars)) {
		return
	}

	parentDecl := v.pkg.FindCustomElementDeclaration(parent.tagName)
	if parentDecl == nil {
		return
	}

	slots := parentDecl.Slots()
	if len(slots) == 0 && parentDecl.OwnSlots() == nil {
		return
	}

	for _, slot := range slots {
		if slot.Name != slotAttr.value {
			continue
		}
		if slot.IsDeprecated() {
			message := fmt.Sprintf("Slot '%s' on '%s' is deprecated", slot.Name, parent.tagName)
			details := map[string]string{DetailSlot: slot.Name, DetailParent: parent.tagName}
			if reason := deprecationReason(slot.Deprecated); reason != "" {
				message += ": " + reason
				details[DetailReason] = reason
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
		DetailSlot:   slotAttr.value,
		DetailParent: parent.tagName,
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
			DetailSlot:      slot.Name,
			DetailGuideline: description,
		})
	}
}

// validateAttributeCombinations checks for contradictory attribute pairs.
// Conflict truthiness follows the manifest type:
//   - exactly boolean: presence semantics, so any present value enables.
//   - a literal union containing both "true" and "false": value semantics,
//     where a static "false" disables and a dynamic value is unknown.
//   - any other declared type or no declaration: presence semantics, so any
//     present value enables.
func (v *validator) validateAttributeCombinations(el *elementNode, decl *manifest.CustomElementDeclaration) {
	attributes := make(map[string]parsedAttribute, len(el.attributes))
	for _, attr := range el.attributes {
		// A dynamic name may resolve to something else entirely.
		if attr.nameDynamic {
			continue
		}
		attributes[attr.name] = attr
	}

	manifestAttributes := decl.Attributes()
	enabled := func(name string) (value string, ok bool) {
		attr, present := attributes[name]
		if !present {
			return "", false
		}

		manifestAttr := findAttribute(manifestAttributes, name)
		switch {
		case isBooleanAttribute(manifestAttr):
			return attr.value, true
		case isTrueFalseUnion(manifestAttr):
			if attr.dynamic {
				// The value is unknown and may resolve to "false".
				return attr.value, false
			}
			return attr.value, !strings.EqualFold(strings.TrimSpace(attr.value), "false")
		default:
			// Presence semantics: any present value enables, including the
			// truthy string "false" on a plain string-typed attribute.
			return attr.value, true
		}
	}

	if loading, ok := attributes["loading"]; ok && loading.value == "eager" {
		if lazy, ok := enabled("lazy"); ok {
			v.add(el, RuleAttributeConflict, SeverityWarning, ConfidenceLow,
				"Eager loading conflicts with lazy loading", map[string]string{
					DetailAttribute1: "loading",
					DetailValue1:     loading.value,
					DetailAttribute2: "lazy",
					DetailValue2:     lazy,
					DetailReason:     "Eager loading conflicts with lazy loading",
				})
		}
	}

	if disabled, ok := enabled("disabled"); ok {
		if interactive, ok := enabled("interactive"); ok {
			v.add(el, RuleAttributeConflict, SeverityWarning, ConfidenceLow,
				"Disabled elements cannot be interactive", map[string]string{
					DetailAttribute1: "disabled",
					DetailValue1:     disabled,
					DetailAttribute2: "interactive",
					DetailValue2:     interactive,
					DetailReason:     "Disabled elements cannot be interactive",
				})
		}
	}
}

// isBooleanAttribute reports whether the manifest declares an attribute as a
// boolean. An absent declaration or type is not treated as boolean.
func isBooleanAttribute(attr *manifest.Attribute) bool {
	if attr == nil || attr.Type == nil {
		return false
	}
	return strings.EqualFold(strings.TrimSpace(attr.Type.Text), "boolean")
}

// isTrueFalseUnion reports whether the manifest type is a literal union that
// contains both "true" and "false", with quotes optional.
func isTrueFalseUnion(attr *manifest.Attribute) bool {
	if attr == nil || !attr.IsEnum() {
		return false
	}

	hasTrue, hasFalse := false, false
	for _, value := range attr.EnumValues() {
		normalized := strings.ToLower(strings.Trim(strings.TrimSpace(value), `"'`))
		switch normalized {
		case "true":
			hasTrue = true
		case "false":
			hasFalse = true
		}
	}
	return hasTrue && hasFalse
}

func (v *validator) validateContentAttributeRedundancy(el *elementNode, decl *manifest.CustomElementDeclaration) {
	attributes := make(map[string]string, len(el.attributes))
	for _, attr := range el.attributes {
		if attr.nameDynamic {
			continue
		}
		attributes[attr.name] = attr.value
	}

	for _, slot := range decl.Slots() {
		if slot.Name != "label" && slot.Name != "title" && slot.Name != "icon" {
			continue
		}
		value, ok := attributes[asciiLower(slot.Name)]
		if !ok {
			continue
		}
		// A global HTML attribute such as title is not the component's own
		// API, so it cannot be redundant with a same-named slot.
		if findAttribute(decl.Attributes(), slot.Name) == nil {
			continue
		}
		if !hasSlottedContent(el, slot.Name) {
			continue
		}
		v.add(el, RuleContentAttributeRedundancy, SeverityInfo, ConfidenceLow,
			fmt.Sprintf("Attribute '%s' on <%s> is redundant with the '%s' slot", slot.Name, el.tagName, slot.Name),
			map[string]string{
				DetailAttribute: slot.Name,
				DetailValue:     value,
				DetailSlot:      slot.Name,
				DetailContent:   "[slot content not analyzed - tree-sitter enhancement needed]",
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
		Column:     utf16Column(v.source, el.lineStart, el.column),
		Details:    details,
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
		Column:     utf16Column(v.source, attr.lineStart, attr.column),
		Details:    details,
	})
}

func findAttribute(attributes []manifest.Attribute, name string) *manifest.Attribute {
	for i := range attributes {
		// Markup names are ASCII-lowercased by the parser; manifest names may
		// be camelCase, so match case-insensitively like the LSP lookup did.
		if asciiLower(attributes[i].Name) == name {
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

// hasSlottedContent reports whether a direct child assigns content to the
// named slot. Slot assignment applies to direct children only.
func hasSlottedContent(el *elementNode, slotName string) bool {
	for _, child := range el.children {
		if child.dynamic {
			continue
		}
		slotAttr := findParsedAttribute(child.attributes, "slot")
		if slotAttr == nil || slotAttr.nameDynamic || slotAttr.dynamic {
			continue
		}
		if slotAttr.value == slotName {
			return true
		}
	}
	return false
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

func containsTemplateExpression(value string, handlebars bool) bool {
	if strings.Contains(value, "{{") {
		return true
	}
	if handlebars {
		return false
	}
	return strings.Contains(value, "{%") || strings.Contains(value, "{#")
}
