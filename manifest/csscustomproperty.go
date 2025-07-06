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

var _ Deprecatable = (*CssCustomProperty)(nil)
var _ Renderable = (*RenderableCssCustomProperty)(nil)

// CssCustomProperty describes a CSS custom property.
type CssCustomProperty struct {
	FullyQualified
	StartByte  uint       `json:"-"`
	Default    string     `json:"default,omitempty"`
	Syntax     string     `json:"syntax,omitempty"`
	Deprecated Deprecated `json:"deprecated,omitempty"` // bool or string
}

func (x *CssCustomProperty) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

func (x *CssCustomProperty) Deprecation() Deprecated {
	return x.Deprecated
}

func (c *CssCustomProperty) UnmarshalJSON(data []byte) error {
	type Alias CssCustomProperty
	aux := &struct {
		Deprecated json.RawMessage `json:"deprecated"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if len(aux.Deprecated) > 0 && string(aux.Deprecated) != "null" {
		var dep Deprecated
		if !decodeDeprecatedField(&dep, aux.Deprecated) {
			return fmt.Errorf("invalid type for deprecated field")
		}
		c.Deprecated = dep
	}
	return nil
}

type RenderableCssCustomProperty struct {
	name                     string
	CssCustomProperty        *CssCustomProperty
	CustomElementDeclaration *CustomElementDeclaration
	CustomElementExport      *CustomElementExport
	JavaScriptModule         *JavaScriptModule
}

func (x *RenderableCssCustomProperty) Name() string {
	return x.CssCustomProperty.Name
}

func (x *RenderableCssCustomProperty) ColumnHeadings() []string {
	return []string{"Name", "Syntax", "Default", "Summary"}
}

// Renders a CSS CssCustomProperty as a table row.
func (x *RenderableCssCustomProperty) ToTableRow() []string {
	return []string{
		x.name,
		x.CssCustomProperty.Syntax,
		x.CssCustomProperty.Default,
		x.CssCustomProperty.Summary,
	}
}

func (x *RenderableCssCustomProperty) ToTreeNode(pred PredicateFunc) pterm.TreeNode {
	label := highlightIfDeprecated(x)
	return pterm.TreeNode{Text: label}
}

func (x *RenderableCssCustomProperty) Children() []Renderable {
	return nil // it's a leaf node
}

func (x *RenderableCssCustomProperty) IsDeprecated() bool {
	return x.CssCustomProperty.Deprecated != nil
}

func (x *RenderableCssCustomProperty) Deprecation() Deprecated {
	return x.CssCustomProperty.Deprecated
}

func NewRenderableCssCustomProperty(
  prop *CssCustomProperty,
	ced *CustomElementDeclaration,
	cee *CustomElementExport,
	mod *Module,
) *RenderableCssCustomProperty {
	return  &RenderableCssCustomProperty{
		name:                     prop.Name,
		CssCustomProperty:        prop,
		CustomElementDeclaration: ced,
		CustomElementExport:      cee,
		JavaScriptModule:         mod,
	}
}
