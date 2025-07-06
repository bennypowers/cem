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

var _ Deprecatable = (*Event)(nil)
var _ Renderable = (*RenderableEvent)(nil)

// Event emitted by a custom element.
type Event struct {
	FullyQualified
	Type          *Type      `json:"type,omitempty"`
	InheritedFrom *Reference `json:"inheritedFrom,omitempty"`
	Deprecated    Deprecated `json:"deprecated,omitempty"` // bool or string
}

func (x *Event) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

func (e *Event) UnmarshalJSON(data []byte) error {
	type Alias Event
	aux := &struct {
		Deprecated json.RawMessage `json:"deprecated"`
		*Alias
	}{
		Alias: (*Alias)(e),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if len(aux.Deprecated) > 0 && string(aux.Deprecated) != "null" {
		var dep Deprecated
		if !decodeDeprecatedField(&dep, aux.Deprecated) {
			return fmt.Errorf("invalid type for deprecated field")
		}
		e.Deprecated = dep
	}
	return nil
}

type RenderableEvent struct {
	Event                    *Event
	CustomElementDeclaration *CustomElementDeclaration
	CustomElementExport      *CustomElementExport
	JavaScriptModule         *JavaScriptModule
}

func (x *RenderableEvent) Name() string {
	return x.Event.Name
}

func (x *RenderableEvent) ColumnHeadings() []string {
	return []string{"Name", "Type", "Summary"}
}

// Renders an Event as a table row.
func (x *RenderableEvent) ToTableRow() []string {
	eventType := ""
	if x.Event.Type != nil {
		eventType = x.Event.Type.Text
	}
	return []string{
		highlightIfDeprecated(x),
		eventType,
		x.Event.Summary,
	}
}

func (x *RenderableEvent) ToTreeNode(pred PredicateFunc) pterm.TreeNode {
	label := highlightIfDeprecated(x)
	return pterm.TreeNode{Text: label}
}

func (x *RenderableEvent) Children() []Renderable {
	return nil // it's a leaf node
}

func (x *RenderableEvent) IsDeprecated() bool {
	return x.Event.Deprecated != nil
}

func (x *RenderableEvent) Deprecation() Deprecated {
	return x.Event.Deprecated
}

var _ Renderable = (*RenderableEvent)(nil)

func NewRenderableEvent(
	ev *Event,
	ced *CustomElementDeclaration,
	cee *CustomElementExport,
	mod *Module,

) *RenderableEvent {
	return &RenderableEvent{
		Event:                    ev,
		CustomElementDeclaration: ced,
		CustomElementExport:      cee,
		JavaScriptModule:         mod,
	}
}
