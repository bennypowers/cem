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

var _ Deprecatable = (*Slot)(nil)
var _ Renderable = (*RenderableSlot)(nil)

// Slot in a custom element.
type Slot struct {
	FullyQualified
	Deprecated Deprecated `json:"deprecated,omitempty"` // bool or string
}

func (x *Slot) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

func (s *Slot) UnmarshalJSON(data []byte) error {
	type Alias Slot
	aux := &struct {
		Deprecated json.RawMessage `json:"deprecated"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if len(aux.Deprecated) > 0 && string(aux.Deprecated) != "null" {
		var dep Deprecated
		if !decodeDeprecatedField(&dep, aux.Deprecated) {
			return fmt.Errorf("invalid type for deprecated field")
		}
		s.Deprecated = dep
	}
	return nil
}

// RenderableSlot adds context and render/traversal methods.
type RenderableSlot struct {
	Slot                     *Slot
	CustomElementDeclaration *CustomElementDeclaration
	CustomElementExport      *CustomElementExport
	JavaScriptModule         *JavaScriptModule
	// Add more context fields as needed
}

func NewRenderableSlot(
  slot *Slot,
	ced *CustomElementDeclaration,
	cee *CustomElementExport,
	mod *Module,
) *RenderableSlot {
	return &RenderableSlot{
		Slot:                     slot,
		CustomElementDeclaration: ced,
		CustomElementExport:      cee,
		JavaScriptModule:         mod,
	}
}

func (x *RenderableSlot) Name() string {
	slotName := x.Slot.Name
	if slotName == "" {
		slotName = "<default>"
	}
	return slotName
}

func (x *RenderableSlot) Label() string {
	return highlightIfDeprecated(x)
}

func (x *RenderableSlot) IsDeprecated() bool {
	return x.Slot != nil && x.Slot.IsDeprecated()
}

func (x *RenderableSlot) Deprecation() Deprecated {
	return x.Slot.Deprecated
}

func (x *RenderableSlot) Children() []Renderable {
	return nil // It's a leaf node
}
func (x *RenderableSlot) ColumnHeadings() []string {
	return []string{
		"Name",
		"Summary",
	}
}

// Renders a Slot as a table row.
func (x *RenderableSlot) ToTableRow() []string {
	return []string{
		highlightIfDeprecated(x),
		x.Slot.Summary,
	}
}

func (x *RenderableSlot) ToTreeNode(p PredicateFunc) pterm.TreeNode {
	return tn(x.Label())
}

