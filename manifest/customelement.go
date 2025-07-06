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

var _ Renderable = (*RenderableCustomElementDeclaration)(nil)
var _ Deprecatable = (*CustomElementDeclaration)(nil)

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
	type Alias CustomElementDeclaration
	aux := &struct {
		Deprecated json.RawMessage   `json:"deprecated"`
		Members    []json.RawMessage `json:"members"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}

	// First, unmarshal into the alias (fills ClassDeclaration and CustomElement)
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
	name                     string // tag name
	CustomElementDeclaration *CustomElementDeclaration
	CustomElementExport      *CustomElementExport
	JavaScriptExport         *JavaScriptExport
	Module                   *Module
	Package                  *Package
	ChildNodes               []Renderable
}

func (x *RenderableCustomElementDeclaration) Name() string {
	return x.CustomElementDeclaration.TagName
	// pterm.LightBlue("<" + x.Name() + ">"),
}

func (x *RenderableCustomElementDeclaration) ColumnHeadings() []string {
	return []string{
		"Tag", // (tag name),
		"Class", // (class name),
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

func (x *RenderableCustomElementDeclaration) ToTreeNode(pred PredicateFunc) pterm.TreeNode {
	// TODO: hmmm
	label := "<"+highlightIfDeprecated(x)+">"
	ft := filterRenderableTree(x, pred)
	children := make([]pterm.TreeNode, 0)
	for _, mem := range ft.Children() {
		children = append(children, mem.ToTreeNode(pred))
	}
	return pterm.TreeNode{
		Text: label,
		Children: children,
	}
}

func (x *RenderableCustomElementDeclaration) Children() []Renderable {
	return x.ChildNodes
}

func (x *RenderableCustomElementDeclaration) IsDeprecated() bool {
	return x.CustomElementDeclaration.IsDeprecated()
}

func (x *RenderableCustomElementDeclaration) Deprecation() Deprecated {
	return x.CustomElementDeclaration.Deprecated
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
			if ecee.Declaration.Name == ced.Name && (ecee.Declaration.Module == "" || ecee.Declaration.Module == mod.Path) {
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
	children := make([]Renderable, 0)
	for i := range ced.Attributes {
		children = append(children, NewRenderableAttribute(&ced.Attributes[i], ced, cee, mod))
	}
	for i := range ced.Events {
		children = append(children, NewRenderableEvent(&ced.Events[i], ced, cee, mod))
	}
	for i := range ced.Slots {
		children = append(children, NewRenderableSlot(&ced.Slots[i], ced, cee, mod))
	}
	for i := range ced.CssParts {
		children = append(children, NewRenderableCssPart(&ced.CssParts[i], ced, cee, mod))
	}
	for i := range ced.CssProperties {
		children = append(children, NewRenderableCssCustomProperty(&ced.CssProperties[i], ced, cee, mod))
	}
	for i := range ced.CssStates {
		children = append(children, NewRenderableCssCustomState(&ced.CssStates[i], ced, cee, mod))
	}
	for _, m := range ced.Members {
		if field, ok := m.(*ClassField); ok {
			children = append(children, NewRenderableClassField(field, &ced.ClassDeclaration, je, mod, pkg))
		} else if cef, ok := m.(*CustomElementField); ok {
			children = append(children, NewRenderableCustomElementField(cef, ced, je, cee, mod, pkg))
		} else if method, ok := m.(*ClassMethod); ok {
			children = append(children, NewRenderableClassMethod(method, &ced.ClassDeclaration, je, mod, pkg))
		}
	}
	return &RenderableCustomElementDeclaration{
		name:                   ced.Name,
		CustomElementDeclaration:  ced,
		CustomElementExport:       cee,
		Module:                    mod,
		ChildNodes:                children,
	}
}

