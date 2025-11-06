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
	"strings"

	"github.com/pterm/pterm"
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

// IsEnum returns true if this attribute has enum/union type constraints
func (a *Attribute) IsEnum() bool {
	if a.Type == nil || a.Type.Text == "" {
		return false
	}
	return strings.Contains(a.Type.Text, "|")
}

// EnumValues extracts enum values from union type definitions
// Returns empty slice if not an enum type
func (a *Attribute) EnumValues() []string {
	if !a.IsEnum() {
		return nil
	}

	// Split on | and trim whitespace from each part
	parts := strings.Split(a.Type.Text, "|")
	values := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			values = append(values, trimmed)
		}
	}
	return values
}

// IsValidValue checks if a value is valid for this attribute
func (a *Attribute) IsValidValue(value string) bool {
	if !a.IsEnum() {
		return true // Non-enum attributes accept any value
	}

	enumValues := a.EnumValues()
	for _, validValue := range enumValues {
		// Check both quoted and unquoted versions
		// HTML attributes are unquoted but TypeScript values may be quoted
		unquoted := validValue
		if len(validValue) >= 2 && validValue[0] == '"' && validValue[len(validValue)-1] == '"' {
			unquoted = validValue[1 : len(validValue)-1]
		}
		if value == validValue || value == unquoted {
			return true
		}
	}
	return false
}

// GetValidationError returns a descriptive error for invalid values
func (a *Attribute) GetValidationError(value string) error {
	if a.IsValidValue(value) {
		return nil
	}

	if !a.IsEnum() {
		return nil // Non-enum attributes don't have validation errors
	}

	enumValues := a.EnumValues()
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
	label += pterm.Gray(" " + x.Attribute.Summary)
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
	}
}

// Renders an Attribute as a table row.
func (x *RenderableAttribute) ToTableRow() []string {
	domProp := ""
	reflects := ""
	typeText := ""
	if x.CustomElementField != nil {
		domProp = x.CustomElementField.Name
		if x.CustomElementField.Reflects {
			reflects = "✅"
		}
	}
	if x.Attribute.Type != nil {
		typeText = x.Attribute.Type.Text
	}
	return []string{
		highlightIfDeprecated(x),
		domProp,
		reflects,
		x.Attribute.Summary,
		x.Attribute.Default,
		typeText,
	}
}

func (x *RenderableAttribute) ToTreeNode(p PredicateFunc) pterm.TreeNode {
	return tn(x.Label())
}
