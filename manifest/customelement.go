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
	"errors"
	"fmt"

	"github.com/pterm/pterm"
)

var _ Deprecatable = (*CustomElementDeclaration)(nil)
var _ Renderable = (*RenderableCustomElementDeclaration)(nil)
var _ GroupedRenderable = (*RenderableCustomElementDeclaration)(nil)

// Demo for custom elements.
type Demo struct {
	Description string           `json:"description,omitempty"`
	URL         string           `json:"url"`
	Source      *SourceReference `json:"source,omitempty"`
}

// CustomElementDeclaration extends ClassDeclaration and CustomElement.
type CustomElementDeclaration struct {
	ClassDeclaration
	CustomElement
}

// CustomElement adds fields to classes/mixins for custom elements.
type CustomElement struct {
	TagName       string              `json:"tagName,omitempty"`
	Attributes    []Attribute         `json:"attributes,omitempty"`
	Events        []Event             `json:"events,omitempty"`
	Slots         []Slot              `json:"slots,omitempty"`
	CssParts      []CssPart           `json:"cssParts,omitempty"`
	CssProperties []CssCustomProperty `json:"cssProperties,omitempty"`
	CssStates     []CssCustomState    `json:"cssStates,omitempty"`
	Demos         []Demo              `json:"demos,omitempty"`
	CustomElement bool                `json:"customElement"`
}

func (*CustomElementDeclaration) isDeclaration() {}

func (x *CustomElementDeclaration) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

func (x *CustomElementDeclaration) GetStartByte() uint {
	return x.StartByte
}

func (c *CustomElementDeclaration) UnmarshalJSON(data []byte) (errs error) {
	type Rest CustomElementDeclaration
	aux := &struct {
		Deprecated json.RawMessage   `json:"deprecated"`
		Members    []json.RawMessage `json:"members"`
		*Rest
	}{
		Rest: (*Rest)(c),
	}

	// First, unmarshal into the "rest" (fills ClassDeclaration and CustomElement)
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Second, unmarshal directly into the embedded CustomElement
	// This ensures CustomElement fields (TagName, Attributes, etc) are populated.
	if err := json.Unmarshal(data, &c.CustomElement); err != nil {
		return err
	}

	// Handle deprecated (custom logic)
	if len(aux.Deprecated) > 0 && string(aux.Deprecated) != "null" {
		var dep Deprecated
		if !decodeDeprecatedField(&dep, aux.Deprecated) {
			return fmt.Errorf("invalid type for deprecated field")
		}
		c.Deprecated = dep
	}

	// Handle members (custom logic)
	for _, m := range aux.Members {
		member, err := unmarshalClassMember(m)
		if member != nil {
			c.Members = append(c.Members, member)
		}
		if err != nil {
			errs = errors.Join(errs, err)
		}
	}

	return errs
}

type RenderableCustomElementDeclaration struct {
	CustomElementDeclaration *CustomElementDeclaration
	CustomElementExport      *CustomElementExport
	JavaScriptExport         *JavaScriptExport
	Module                   *Module
	Package                  *Package
	ChildNodes               []Renderable
	attributes               []Renderable
	events                   []Renderable
	slots                    []Renderable
	cssparts                 []Renderable
	cssprops                 []Renderable
	cssstates                []Renderable
	fields                   []Renderable
	methods                  []Renderable
}

func NewRenderableCustomElementDeclaration(
	ced *CustomElementDeclaration,
	mod *Module,
	pkg *Package,
) *RenderableCustomElementDeclaration {
	// TODO use a map
	// TODO get je, cee from other modules
	var cee *CustomElementExport
	var je *JavaScriptExport
	for i, exp := range mod.Exports {
		if ecee, ok := exp.(*CustomElementExport); ok {
			if ecee != nil && ecee.Declaration != nil && ecee.Declaration.Name == ced.Name && (ecee.Declaration.Module == "" || ecee.Declaration.Module == mod.Path) {
				cee = mod.Exports[i].(*CustomElementExport)
			}
		}
		if eje, ok := exp.(*JavaScriptExport); ok {
			if eje.Declaration.Name == ced.Name && (eje.Declaration.Module == "" || eje.Declaration.Module == mod.Path) {
				je = mod.Exports[i].(*JavaScriptExport)
			}
		}
		if je != nil && cee != nil {
			break
		}
	}
	r := RenderableCustomElementDeclaration{
		CustomElementDeclaration: ced,
		CustomElementExport:      cee,
		Module:                   mod,
	}
	for i := range ced.Attributes {
		m := NewRenderableAttribute(&ced.Attributes[i], ced, cee, mod)
		r.ChildNodes = append(r.ChildNodes, m)
		r.attributes = append(r.attributes, m)
	}
	for i := range ced.Events {
		m := NewRenderableEvent(&ced.Events[i], ced, cee, mod)
		r.ChildNodes = append(r.ChildNodes, m)
		r.events = append(r.events, m)
	}
	for i := range ced.Slots {
		m := NewRenderableSlot(&ced.Slots[i], ced, cee, mod)
		r.ChildNodes = append(r.ChildNodes, m)
		r.slots = append(r.slots, m)
	}
	for i := range ced.CssParts {
		m := NewRenderableCssPart(&ced.CssParts[i], ced, cee, mod)
		r.ChildNodes = append(r.ChildNodes, m)
		r.cssparts = append(r.cssparts, m)
	}
	for i := range ced.CssProperties {
		m := NewRenderableCssCustomProperty(&ced.CssProperties[i], ced, cee, mod)
		r.ChildNodes = append(r.ChildNodes, m)
		r.cssprops = append(r.cssprops, m)
	}
	for i := range ced.CssStates {
		m := NewRenderableCssCustomState(&ced.CssStates[i], ced, cee, mod)
		r.ChildNodes = append(r.ChildNodes, m)
		r.cssstates = append(r.cssstates, m)
	}
	for _, m := range ced.Members {
		if field, ok := m.(*ClassField); ok {
			m := NewRenderableClassField(field, &ced.ClassDeclaration, je, mod, pkg)
			r.ChildNodes = append(r.ChildNodes, m)
			r.fields = append(r.fields, m)
		} else if cef, ok := m.(*CustomElementField); ok {
			m := NewRenderableCustomElementField(cef, ced, je, cee, mod, pkg)
			r.ChildNodes = append(r.ChildNodes, m)
			r.fields = append(r.fields, m)
		} else if method, ok := m.(*ClassMethod); ok {
			m := NewRenderableClassMethod(method, &ced.ClassDeclaration, je, mod, pkg)
			r.ChildNodes = append(r.ChildNodes, m)
			r.methods = append(r.methods, m)
		}
	}
	return &r
}

func (x *RenderableCustomElementDeclaration) Name() string {
	return x.CustomElementDeclaration.TagName
}

func (x *RenderableCustomElementDeclaration) Label() string {
	return highlightIfDeprecated(x, "<", ">")
}

func (x *RenderableCustomElementDeclaration) IsDeprecated() bool {
	return x.CustomElementDeclaration.IsDeprecated()
}

func (x *RenderableCustomElementDeclaration) Deprecation() Deprecated {
	return x.CustomElementDeclaration.Deprecated
}

func (x *RenderableCustomElementDeclaration) Children() []Renderable {
	return x.ChildNodes
}

func (x *RenderableCustomElementDeclaration) GroupedChildren(p PredicateFunc) []pterm.TreeNode {
	var cs []pterm.TreeNode

	if attrs := toTreeChildren(x.attributes, p); len(attrs) > 0 {
		cs = append(cs, tn(pterm.Blue("Attributes"), attrs...))
	}
	if slots := toTreeChildren(x.slots, p); len(slots) > 0 {
		cs = append(cs, tn(pterm.Blue("Slots"), slots...))
	}
	if events := toTreeChildren(x.events, p); len(events) > 0 {
		cs = append(cs, tn(pterm.Blue("Events"), events...))
	}
	if fields := toTreeChildren(x.fields, p); len(fields) > 0 {
		cs = append(cs, tn(pterm.Blue("Fields"), fields...))
	}
	if methods := toTreeChildren(x.methods, p); len(methods) > 0 {
		cs = append(cs, tn(pterm.Blue("Methods"), methods...))
	}
	if cssprops := toTreeChildren(x.cssprops, p); len(cssprops) > 0 {
		cs = append(cs, tn(pterm.Blue("CSS Properties"), cssprops...))
	}
	if cssparts := toTreeChildren(x.cssparts, p); len(cssparts) > 0 {
		cs = append(cs, tn(pterm.Blue("Parts"), cssparts...))
	}
	if cssstates := toTreeChildren(x.cssstates, p); len(cssstates) > 0 {
		cs = append(cs, tn(pterm.Blue("States"), cssstates...))
	}

	return cs
}

func (x *RenderableCustomElementDeclaration) ColumnHeadings() []string {
	return []string{
		"Tag",    // (tag name),
		"Class",  // (class name),
		"Module", //(module path),
		"Summary",
	}
}

// Renders a CustomElement as a table row.

func (x *RenderableCustomElementDeclaration) ToTableRow() []string {
	modulePath := ""
	if x.Module != nil {
		modulePath = x.Module.Path
	}
	return []string{
		highlightIfDeprecated(x),
		x.CustomElementDeclaration.Name,
		modulePath,
		x.CustomElementDeclaration.Summary,
	}
}

func (x *RenderableCustomElementDeclaration) ToTreeNode(p PredicateFunc) pterm.TreeNode {
	return tn(x.Label(), x.GroupedChildren(p)...)
}

func (x *RenderableCustomElementDeclaration) Sections() []Section {
	return []Section{
		{Title: "Attributes", Items: x.attributes},
		{Title: "Slots", Items: x.slots},
		{Title: "Events", Items: x.events},
		{Title: "Fields", Items: x.fields},
		{Title: "Methods", Items: x.methods},
		{Title: "CSS Properties", Items: x.cssprops},
		{Title: "CSS Parts", Items: x.cssparts},
		{Title: "CSS States", Items: x.cssstates},
	}
}

func (x *RenderableCustomElementDeclaration) Summary() string {
	return x.CustomElementDeclaration.Summary
}

func (x *RenderableCustomElementDeclaration) Description() string {
	return x.CustomElementDeclaration.Description
}
