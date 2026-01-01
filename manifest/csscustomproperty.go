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
	InheritedFrom *Reference `json:"inheritedFrom,omitempty"`
	StartByte     uint       `json:"-"`
	Default       string     `json:"default,omitempty"`
	Syntax        string     `json:"syntax,omitempty"`
	Deprecated    Deprecated `json:"deprecated,omitempty"` // bool or string
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
	type Rest CssCustomProperty
	aux := &struct {
		Deprecated json.RawMessage `json:"deprecated"`
		*Rest
	}{
		Rest: (*Rest)(c),
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

// Clone creates a deep copy of the CssCustomProperty.
// Handles the embedded FullyQualified structure and all custom property fields.
//
// Performance: Efficient deep copying without JSON serialization overhead
// Thread Safety: Safe for concurrent use (creates new instance)
func (c CssCustomProperty) Clone() CssCustomProperty {
	cloned := CssCustomProperty{
		StartByte: c.StartByte,
		Default:   c.Default,
		Syntax:    c.Syntax,
	}

	// Clone the embedded FullyQualified
	cloned.FullyQualified = c.FullyQualified.Clone()

	if c.InheritedFrom != nil {
		inheritedFrom := *c.InheritedFrom
		cloned.InheritedFrom = &inheritedFrom
	}

	if c.Deprecated != nil {
		cloned.Deprecated = c.Deprecated.Clone()
	}

	return cloned
}

type RenderableCssCustomProperty struct {
	CssCustomProperty        *CssCustomProperty
	CustomElementDeclaration *CustomElementDeclaration
	CustomElementExport      *CustomElementExport
	JavaScriptModule         *JavaScriptModule
}

func (x *RenderableCssCustomProperty) Name() string {
	return x.CssCustomProperty.Name
}

func (x *RenderableCssCustomProperty) Label() string {
	return highlightIfDeprecated(x) + " " + pterm.Gray(x.CssCustomProperty.Summary)
}

func (x *RenderableCssCustomProperty) IsDeprecated() bool {
	return x.CssCustomProperty.Deprecated != nil
}

func (x *RenderableCssCustomProperty) Deprecation() Deprecated {
	return x.CssCustomProperty.Deprecated
}

func (x *RenderableCssCustomProperty) Children() []Renderable {
	return nil // it's a leaf node
}

func (x *RenderableCssCustomProperty) ColumnHeadings() []string {
	return []string{"Name", "Syntax", "Default", "Summary", "Inherited From"}
}

// Renders a CSS CssCustomProperty as a table row.
func (x *RenderableCssCustomProperty) ToTableRow() []string {
	inheritedFrom := ""
	if x.CssCustomProperty.InheritedFrom != nil {
		inheritedFrom = x.CssCustomProperty.InheritedFrom.Name
	}
	return []string{
		highlightIfDeprecated(x),
		x.CssCustomProperty.Syntax,
		x.CssCustomProperty.Default,
		x.CssCustomProperty.Summary,
		inheritedFrom,
	}
}

func (x *RenderableCssCustomProperty) ToTreeNode(pred PredicateFunc) pterm.TreeNode {
	return pterm.TreeNode{Text: x.Label()}
}

func NewRenderableCssCustomProperty(
	prop *CssCustomProperty,
	ced *CustomElementDeclaration,
	cee *CustomElementExport,
	mod *Module,
) *RenderableCssCustomProperty {
	return &RenderableCssCustomProperty{
		CssCustomProperty:        prop,
		CustomElementDeclaration: ced,
		CustomElementExport:      cee,
		JavaScriptModule:         mod,
	}
}
