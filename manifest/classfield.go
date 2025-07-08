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

var _ Deprecatable = (*ClassField)(nil)
var _ Deprecatable = (*CustomElementField)(nil)
var _ Renderable = (*RenderableClassField)(nil)
var _ Renderable = (*RenderableCustomElementField)(nil)

// ClassField is a class field.
type ClassField struct {
	PropertyLike
	Deprecatable  `json:"-"`
	Kind          string           `json:"kind"` // 'field'
	Static        bool             `json:"static,omitempty"`
	Privacy       Privacy          `json:"privacy,omitempty"` // 'public', 'private', 'protected'
	InheritedFrom *Reference       `json:"inheritedFrom,omitempty"`
	Source        *SourceReference `json:"source,omitempty"`
}

func (*ClassField) isClassMember() {}

func (x *ClassField) IsDeprecated() bool {
	return x != nil && x.Deprecated != nil
}

func (f *ClassField) GetStartByte() uint { return f.StartByte }

func (f *ClassField) UnmarshalJSON(data []byte) error {
	type Alias ClassField
	aux := &struct {
		Deprecated json.RawMessage `json:"deprecated"`
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if len(aux.Deprecated) != 0 && string(aux.Deprecated) != "null" {
		// Try bool
		var b bool
		if err := json.Unmarshal(aux.Deprecated, &b); err == nil {
			f.Deprecated = DeprecatedFlag(b)
			return nil
		}
		// Try string
		var s string
		if err := json.Unmarshal(aux.Deprecated, &s); err == nil {
			f.Deprecated = DeprecatedReason(s)
			return nil
		} else {
			return err
		}
		// Ignore other types
	}
	return nil
}

type RenderableClassField struct {
	name             string
	ClassField       *ClassField
	ClassDeclaration *ClassDeclaration
	JavaScriptExport *JavaScriptExport
	Module           *Module
	Package          *Package
}

func NewRenderableClassField(
	field *ClassField,
	cd *ClassDeclaration,
	ce *JavaScriptExport,
	mod *Module,
	pkg *Package,
) *RenderableClassField {
	return &RenderableClassField{
		name:             field.Name,
		ClassField:       field,
		ClassDeclaration: cd,
		JavaScriptExport: ce,
		Module:           mod,
		Package:          pkg,
	}
}

func (x *RenderableClassField) Name() string {
	return x.ClassField.Name
	// "class field " + x.Name()
}

func (x *RenderableClassField) Label() string {
	sum := ""
	if x.ClassField.Summary != "" {
		sum = pterm.Gray(x.ClassField.Summary)
	}
	return strings.TrimSpace(highlightIfDeprecated(x) + " " + sum)
}

func (x *RenderableClassField) IsDeprecated() bool {
	return x.ClassField.IsDeprecated()
}

func (x *RenderableClassField) Deprecation() Deprecated {
	return x.ClassField.Deprecated
}

func (x *RenderableClassField) Children() []Renderable {
	return nil // it's a leaf node
}

func (x *RenderableClassField) ColumnHeadings() []string {
	return []string{"Name", "Type", "Summary"}
}

func (x *RenderableClassField) ToTableRow() []string {
	typeText := ""
	if x.ClassField.Type != nil {
		typeText = x.ClassField.Type.Text
	}
	return []string{
		highlightIfDeprecated(x),
		typeText,
		x.ClassField.Summary,
	}
}

func (x *RenderableClassField) ToTreeNode(pred PredicateFunc) pterm.TreeNode {
	return pterm.TreeNode{Text: x.Label()}
}

// CustomElementField extends ClassField with attribute/reflects.
type CustomElementField struct {
	ClassField
	Attribute string `json:"attribute,omitempty"`
	Reflects  bool   `json:"reflects,omitempty"`
}

func (x *CustomElementField) isClassMember() {}

func (x *CustomElementField) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

func (x *CustomElementField) GetStartByte() uint { return x.ClassField.StartByte }

func (f *CustomElementField) UnmarshalJSON(data []byte) error {
	type Alias CustomElementField
	aux := &struct {
		Deprecated json.RawMessage `json:"deprecated"`
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if len(aux.Deprecated) > 0 && string(aux.Deprecated) != "null" {
		var dep Deprecated
		if !decodeDeprecatedField(&dep, aux.Deprecated) {
			return fmt.Errorf("invalid type for deprecated field")
		}
		f.Deprecated = dep
	}
	return nil
}

type RenderableCustomElementField struct {
	name                     string
	CustomElementField       *CustomElementField
	CustomElementDeclaration *CustomElementDeclaration
	JavaScriptExport         *JavaScriptExport
	CustomElementExport      *CustomElementExport
	Module                   *Module
	Package                  *Package
}

func NewRenderableCustomElementField(
	field *CustomElementField,
	ced *CustomElementDeclaration,
	je *JavaScriptExport,
	cee *CustomElementExport,
	mod *Module,
	pkg *Package,
) *RenderableCustomElementField {
	return &RenderableCustomElementField{
		name:                     field.Name,
		CustomElementField:       field,
		CustomElementDeclaration: ced,
		JavaScriptExport:         je,
		CustomElementExport:      cee,
		Module:                   mod,
		Package:                  pkg,
	}
}

func (x *RenderableCustomElementField) Name() string {
	return x.CustomElementField.Name
	// return "custom element field " + x.CustomElementField.Name,
}

func (x *RenderableCustomElementField) Label() string {
	return highlightIfDeprecated(x) + " " + pterm.Gray(x.CustomElementField.Summary)
}

func (x *RenderableCustomElementField) IsDeprecated() bool {
	return x.CustomElementField.IsDeprecated()
}

func (x *RenderableCustomElementField) Deprecation() Deprecated {
	return x.CustomElementField.Deprecated
}

func (x *RenderableCustomElementField) Children() []Renderable {
	return nil // it's a leaf node
}

func (x *RenderableCustomElementField) ColumnHeadings() []string {
	return []string{"Name", "Type", "Summary"}
}

func (x *RenderableCustomElementField) ToTableRow() []string {
	typeText := ""
	if x.CustomElementField.Type != nil {
		typeText = x.CustomElementField.Type.Text
	}
	return []string{
		highlightIfDeprecated(x),
		typeText,
		x.CustomElementField.Summary,
	}
}

func (x *RenderableCustomElementField) ToTreeNode(pred PredicateFunc) pterm.TreeNode {
	return pterm.TreeNode{Text: x.Label()}
}
