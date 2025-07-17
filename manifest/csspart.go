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

var _ Deprecatable = (*CssPart)(nil)
var _ Renderable = (*RenderableCssPart)(nil)

// CssPart describes a CSS part.
type CssPart struct {
	FullyQualified
	Deprecated Deprecated `json:"deprecated,omitempty"` // bool or string
}

func (x *CssPart) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

func (c *CssPart) UnmarshalJSON(data []byte) error {
	type Rest CssPart
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

// RenderableCssPart adds context and render/traversal methods.
type RenderableCssPart struct {
	CssPart                  *CssPart
	CustomElementDeclaration *CustomElementDeclaration
	CustomElementExport      *CustomElementExport
	JavaScriptModule         *JavaScriptModule
	// Add more context fields as needed
}

func NewRenderableCssPart(
	part *CssPart,
	ced *CustomElementDeclaration,
	cee *CustomElementExport,
	mod *Module,
) *RenderableCssPart {
	return &RenderableCssPart{
		CssPart:                  part,
		CustomElementDeclaration: ced,
		CustomElementExport:      cee,
		JavaScriptModule:         mod,
	}
}

func (x *RenderableCssPart) Name() string {
	return x.CssPart.Name
}

func (x *RenderableCssPart) Label() string {
	return highlightIfDeprecated(x) + " " + pterm.Gray(x.CssPart.Summary)
}

func (x *RenderableCssPart) IsDeprecated() bool {
	return x.CssPart != nil && x.CssPart.IsDeprecated()
}

func (x *RenderableCssPart) Deprecation() Deprecated {
	return x.CssPart.Deprecated
}

func (x *RenderableCssPart) Children() []Renderable {
	return nil // It's a leaf node
}

func (x *RenderableCssPart) ColumnHeadings() []string {
	return []string{
		"Name",
		"Summary",
	}
}

// Renders a CssPart as a table row.
func (x *RenderableCssPart) ToTableRow() []string {
	return []string{
		highlightIfDeprecated(x),
		x.CssPart.Summary,
	}
}

func (x *RenderableCssPart) ToTreeNode(pred PredicateFunc) pterm.TreeNode {
	return pterm.TreeNode{Text: x.Label()}
}
