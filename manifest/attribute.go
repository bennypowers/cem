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
	var field *CustomElementField
	// TODO: perf: use a map
	// reuse the one from TagRenderableAttributes, maybe refactor so it goes
	// from Field to Attr or something
	for _, f := range ced.Members {
		if cef, ok := f.(*CustomElementField); ok {
			if cef.Attribute == attr.Name {
				field = cef
				break
			}
		}
	}
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
