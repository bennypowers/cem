/*
Copyright © 2025 Benny Powers

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

package manifest

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"bennypowers.dev/cem/internal/typesplit"
)

var _ Deprecatable = (*Attribute)(nil)
var _ Renderable = (*RenderableAttribute)(nil)

// Attribute for custom elements.
type Attribute struct {
	FullyQualified
	InheritedFrom *Reference `json:"inheritedFrom,omitempty"`
	Type          *Type      `json:"type,omitempty"`
	Default       string     `json:"default,omitempty"`
	FieldName     string     `json:"fieldName,omitempty"`
	Deprecated    Deprecated `json:"deprecated,omitempty"` // bool or string
	StartByte     uint       `json:"-"`
}

func (x *Attribute) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

func (a *Attribute) UnmarshalJSON(data []byte) error {
	type Rest Attribute
	aux := &struct {
		Deprecated json.RawMessage `json:"deprecated"`
		*Rest
	}{
		Rest: (*Rest)(a),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if len(aux.Deprecated) > 0 && string(aux.Deprecated) != "null" {
		var dep Deprecated
		if !decodeDeprecatedField(&dep, aux.Deprecated) {
			return fmt.Errorf("invalid type for deprecated field")
		}
		a.Deprecated = dep
	}
	return nil
}

// Clone creates a deep copy of the Attribute.
// Handles all embedded structures and references with proper deep copying.
//
// Performance: Efficient deep copying without JSON serialization overhead
// Thread Safety: Safe for concurrent use (creates new instance)
func (a Attribute) Clone() Attribute {
	cloned := Attribute{
		Default:   a.Default,
		FieldName: a.FieldName,
		StartByte: a.StartByte,
	}

	// Clone the embedded FullyQualified
	cloned.FullyQualified = a.FullyQualified.Clone()

	if a.Deprecated != nil {
		cloned.Deprecated = a.Deprecated.Clone()
	}

	if a.InheritedFrom != nil {
		inheritedFrom := a.InheritedFrom.Clone()
		cloned.InheritedFrom = &inheritedFrom
	}

	if a.Type != nil {
		cloned.Type = a.Type.Clone()
	}

	return cloned
}

// Validation methods for attribute values

// enumParts returns the union members if this attribute has a union type,
// or nil otherwise. Single parse, reusable for enum detection and validation.
func (a *Attribute) enumParts() []string {
	if a.Type == nil || a.Type.Text == "" {
		return nil
	}
	parts := typesplit.SplitTopLevelUnion(a.Type.Text)
	if len(parts) <= 1 {
		return nil
	}
	return parts
}

// IsEnum returns true if this attribute has enum/union type constraints
func (a *Attribute) IsEnum() bool {
	return a.enumParts() != nil
}

// IsClosedEnum reports whether every union member is a literal attribute value
// or the absence marker null/undefined. A union containing an open or
// unresolvable member (for example string, a type reference, or an object
// type) cannot be validated against a finite set of values.
func (a *Attribute) IsClosedEnum() bool {
	parts := a.enumParts()
	if parts == nil {
		return false
	}
	for _, part := range parts {
		member := strings.TrimSpace(part)
		if member == "null" || member == "undefined" {
			continue
		}
		if _, _, ok := literalAttributeValue(member); ok {
			continue
		}
		return false
	}
	return true
}

// EnumValues extracts enum values from union type definitions
func (a *Attribute) EnumValues() []string {
	return a.enumParts()
}

// DisplayEnumValues returns the literal values accepted by this attribute, in
// the form shown to users. String literals are decoded and re-quoted, so the
// source quote style and source escapes do not leak into diagnostics. Union
// members that are not literals cannot be attribute values and are omitted.
func (a *Attribute) DisplayEnumValues() []string {
	parts := a.enumParts()
	if parts == nil {
		return nil
	}

	var values []string
	for _, part := range parts {
		literal, quoted, ok := literalAttributeValue(strings.TrimSpace(part))
		if !ok {
			continue
		}
		if quoted {
			values = append(values, strconv.Quote(literal))
			continue
		}
		values = append(values, literal)
	}
	return values
}

// IsValidValue checks if a value is valid for this attribute
func (a *Attribute) IsValidValue(value string) bool {
	parts := a.enumParts()
	if parts == nil {
		return true
	}

	for _, part := range parts {
		member := strings.TrimSpace(part)
		literal, _, ok := literalAttributeValue(member)
		if !ok {
			continue
		}
		// Compare against both the literal value and the raw member for
		// compatibility with callers that pass the TypeScript source form.
		if value == literal || value == member {
			return true
		}
	}
	return false
}

// literalAttributeValue reports whether a union member can be supplied as an
// HTML attribute value, and returns that value. Only string literals, numeric
// literals, and boolean literals qualify. Bare identifiers such as null,
// undefined, and type or constant names are not attribute values. The quoted
// result reports whether the value came from a string literal.
func literalAttributeValue(member string) (value string, quoted, ok bool) {
	member = strings.TrimSpace(member)
	if value, ok := stringLiteralValue(member); ok {
		return value, true, true
	}
	if member == "true" || member == "false" {
		return member, false, true
	}
	if isNumericLiteral(member) {
		return member, false, true
	}
	return "", false, false
}

// stringLiteralValue reports whether member is a single- or double-quoted
// TypeScript string literal and returns its decoded value. Escaped delimiters
// are decoded; an unescaped delimiter inside the body means the member is not
// one string literal (for example an intersection).
func stringLiteralValue(member string) (string, bool) {
	if len(member) < 2 {
		return "", false
	}
	quote := member[0]
	if (quote != '\'' && quote != '"') || member[len(member)-1] != quote {
		return "", false
	}

	body := member[1 : len(member)-1]
	var value strings.Builder
	value.Grow(len(body))
	for i := 0; i < len(body); i++ {
		c := body[i]
		if c == quote {
			return "", false
		}
		if c != '\\' {
			value.WriteByte(c)
			continue
		}
		i++
		if i >= len(body) {
			return "", false
		}
		switch escaped := body[i]; escaped {
		case '\\', '\'', '"':
			value.WriteByte(escaped)
		case 'n':
			value.WriteByte('\n')
		case 'r':
			value.WriteByte('\r')
		case 't':
			value.WriteByte('\t')
		case 'b':
			value.WriteByte('\b')
		case 'f':
			value.WriteByte('\f')
		case 'v':
			value.WriteByte('\v')
		case '0':
			value.WriteByte(0)
		case 'x':
			decoded, next, ok := hexEscape(body, i+1, 2)
			if !ok {
				value.WriteByte(escaped)
				break
			}
			// \xHH names code point U+00HH, not a raw byte.
			value.WriteRune(rune(decoded))
			i = next - 1
		case 'u':
			decoded, next, ok := unicodeEscape(body, i+1)
			if !ok {
				value.WriteByte(escaped)
				break
			}
			// A UTF-16 surrogate pair is written as two consecutive \u
			// escapes; combine it into a single code point.
			if decoded >= 0xD800 && decoded <= 0xDBFF &&
				next+1 < len(body) && body[next] == '\\' && body[next+1] == 'u' {
				if low, lowNext, lowOK := unicodeEscape(body, next+2); lowOK &&
					low >= 0xDC00 && low <= 0xDFFF {
					decoded = 0x10000 + (decoded-0xD800)<<10 + (low - 0xDC00)
					next = lowNext
				}
			}
			if !utf8.ValidRune(decoded) {
				// A lone or mismatched surrogate has no UTF-8 form, so the
				// member cannot be a usable attribute literal.
				return "", false
			}
			value.WriteRune(decoded)
			i = next - 1
		default:
			// Identity escape: TypeScript keeps the escaped character.
			value.WriteByte(escaped)
		}
	}
	return value.String(), true
}

// hexEscape parses exactly n hexadecimal digits of s starting at start.
func hexEscape(s string, start, n int) (uint64, int, bool) {
	if start+n > len(s) {
		return 0, start, false
	}
	value, err := strconv.ParseUint(s[start:start+n], 16, 32)
	if err != nil {
		return 0, start, false
	}
	return value, start + n, true
}

// unicodeEscape parses the payload of a \u escape starting at start. Both the
// four-hex-digit form and the braced code point form are supported.
func unicodeEscape(s string, start int) (rune, int, bool) {
	if start < len(s) && s[start] == '{' {
		end := strings.IndexByte(s[start+1:], '}')
		if end < 0 {
			return 0, start, false
		}
		end += start + 1
		parsed, err := strconv.ParseUint(s[start+1:end], 16, 32)
		if err != nil {
			return 0, start, false
		}
		value := rune(parsed)
		if !utf8.ValidRune(value) {
			return 0, start, false
		}
		return value, end + 1, true
	}

	// Four-hex-digit escapes may name a UTF-16 surrogate half; the caller
	// combines a valid pair into one code point.
	parsed, next, ok := hexEscape(s, start, 4)
	if !ok {
		return 0, start, false
	}
	return rune(parsed), next, true
}

// isNumericLiteral reports whether s is a TypeScript numeric literal. Decimal,
// hexadecimal, binary, and octal forms are recognized, including numeric
// separators. Bigint suffixes and a leading plus sign are not literals.
func isNumericLiteral(s string) bool {
	i := 0
	if i < len(s) && s[i] == '-' {
		i++
	}
	if i >= len(s) {
		return false
	}

	if s[i] == '0' && i+1 < len(s) {
		base := 0
		switch s[i+1] {
		case 'x', 'X':
			base = 16
		case 'b', 'B':
			base = 2
		case 'o', 'O':
			base = 8
		}
		if base != 0 {
			return hasDigits(s[i+2:], base)
		}
	}

	intStart := i
	i, intDigits := scanDecimalDigits(s, i)
	if intDigits > 1 && s[intStart] == '0' {
		// A leading zero with more digits is a legacy octal literal, which
		// TypeScript rejects in modules.
		return false
	}

	fracDigits := 0
	if i < len(s) && s[i] == '.' {
		i++
		i, fracDigits = scanDecimalDigits(s, i)
	}
	if intDigits == 0 && fracDigits == 0 {
		return false
	}

	if i < len(s) && (s[i] == 'e' || s[i] == 'E') {
		i++
		if i < len(s) && (s[i] == '+' || s[i] == '-') {
			i++
		}
		var expDigits int
		i, expDigits = scanDecimalDigits(s, i)
		if expDigits == 0 {
			return false
		}
	}

	return i == len(s)
}

// scanDecimalDigits consumes decimal digits and numeric separators starting at
// i, returning the next offset and the digit count.
func scanDecimalDigits(s string, i int) (int, int) {
	count := 0
	for i < len(s) {
		switch {
		case isDecimalDigit(s[i]):
			count++
			i++
		case s[i] == '_' && count > 0 && i+1 < len(s) && isDecimalDigit(s[i+1]):
			i++
		default:
			return i, count
		}
	}
	return i, count
}

// hasDigits reports whether s is a non-empty digit sequence in the given base,
// allowing numeric separators between digits.
func hasDigits(s string, base int) bool {
	if s == "" {
		return false
	}
	count := 0
	lastDigit := false
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '_' {
			if !lastDigit || i+1 >= len(s) || !isDigitForBase(s[i+1], base) {
				return false
			}
			lastDigit = false
			continue
		}
		if !isDigitForBase(c, base) {
			return false
		}
		count++
		lastDigit = true
	}
	return count > 0
}

// isDigitForBase reports whether c is a digit in the given base.
func isDigitForBase(c byte, base int) bool {
	var digit int
	switch {
	case c >= '0' && c <= '9':
		digit = int(c - '0')
	case c >= 'a' && c <= 'f':
		digit = int(c-'a') + 10
	case c >= 'A' && c <= 'F':
		digit = int(c-'A') + 10
	default:
		return false
	}
	return digit < base
}

// isDecimalDigit reports whether c is a decimal digit.
func isDecimalDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// GetValidationError returns a descriptive error for invalid values
func (a *Attribute) GetValidationError(value string) error {
	if a.IsValidValue(value) {
		return nil
	}

	if !a.IsEnum() {
		return nil // Non-enum attributes don't have validation errors
	}

	enumValues := a.DisplayEnumValues()
	if len(enumValues) == 0 {
		// A union with no literal members has no accepted set to report.
		return fmt.Errorf("invalid value %q for attribute %q", value, a.Name)
	}
	return fmt.Errorf("invalid value %q for attribute %q. Valid values: %v",
		value, a.Name, enumValues)
}

// RenderableAttribute adds context and render/traversal methods.
type RenderableAttribute struct {
	Attribute                *Attribute
	CustomElementField       *CustomElementField
	CustomElementDeclaration *CustomElementDeclaration
	CustomElementExport      *CustomElementExport
	JavaScriptModule         *JavaScriptModule
	// Add more context fields as needed
}

func NewRenderableAttribute(
	attr *Attribute,
	ced *CustomElementDeclaration,
	cee *CustomElementExport,
	mod *Module,
) *RenderableAttribute {
	// Use O(1) map lookup instead of O(n) linear search
	field := ced.LookupAttributeField(attr.Name)
	return &RenderableAttribute{
		Attribute:                attr,
		CustomElementField:       field,
		CustomElementDeclaration: ced,
		CustomElementExport:      cee,
		JavaScriptModule:         mod,
	}
}

func (x *RenderableAttribute) Name() string {
	return x.Attribute.Name
}

func (x *RenderableAttribute) Label() string {
	label := highlightIfDeprecated(x)
	if x.CustomElementField != nil && x.CustomElementField.Reflects {
		label += " (reflects)"
	}
	label += summaryStyle.Render(" " + x.Attribute.Summary)
	return label
}

func (x *RenderableAttribute) IsDeprecated() bool {
	return x.Attribute != nil && x.Attribute.IsDeprecated()
}

func (x *RenderableAttribute) Deprecation() Deprecated {
	return x.Attribute.Deprecated
}

func (x *RenderableAttribute) Children() []Renderable {
	return nil // It's a leaf node
}

func (x *RenderableAttribute) ColumnHeadings() []string {
	return []string{
		"Name",
		"DOM Property",
		"Reflects",
		"Summary",
		"Default",
		"Type",
		"Inherited From",
	}
}

// Renders an Attribute as a table row.
func (x *RenderableAttribute) ToTableRow() []string {
	domProp := ""
	reflects := ""
	typeText := ""
	inheritedFrom := ""
	if x.CustomElementField != nil {
		domProp = x.CustomElementField.Name
		if x.CustomElementField.Reflects {
			reflects = "✅"
		}
	}
	if x.Attribute.Type != nil {
		typeText = x.Attribute.Type.Text
	}
	if x.Attribute.InheritedFrom != nil {
		inheritedFrom = x.Attribute.InheritedFrom.Name
	}
	return []string{
		highlightIfDeprecated(x),
		domProp,
		reflects,
		x.Attribute.Summary,
		x.Attribute.Default,
		typeText,
		inheritedFrom,
	}
}

func (x *RenderableAttribute) ToTreeNode(p PredicateFunc) TreeNode {
	return tn("attribute", x.Label())
}
