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
	"errors"
	"slices"
	"strings"

	"github.com/pterm/pterm"
)

type PredicateFunc func(Renderable) bool

type Renderable interface {
	Deprecatable
	Name() string
	Label() string
	Deprecation() Deprecated
	ColumnHeadings() []string
	ToTableRow() []string
	ToTreeNode(pred PredicateFunc) pterm.TreeNode
	Children() []Renderable
}

type GroupedRenderable interface {
	GroupedChildren(p PredicateFunc) []pterm.TreeNode
}

// Predicate: keep only deprecated nodes
func IsDeprecated(d Renderable) bool { return d.IsDeprecated() }

func True(d Renderable) bool         { return true }

func formatDeprecated(deprecated any) (label string) {
	if deprecated == nil {
		return ""
	}
	switch v := deprecated.(type) {
	case DeprecatedReason:
		return "(" + pterm.Red("DEPRECATED") + ": " + pterm.LightRed(v) + ")"
	default:
		return "(" + pterm.Red("DEPRECATED") + ")"
	}
}

func highlightIfDeprecated(x Renderable, fixes ...string) string {
	prefix:=""
	suffix:=""
	if len(fixes) > 0 {
		prefix = fixes[0]
	}
	if len(fixes) > 1 {
		suffix = fixes[1]
	}
	name := prefix + x.Name() + suffix
	if x == nil || !x.IsDeprecated() {
		return name
	}
	return name + " " + formatDeprecated(x.Deprecation())
}

func (x *Package) GetAllTagNames() (tags []string) {
	// Write index.html
	for _, m := range x.Modules {
		for _, decl := range m.Declarations {
			ced, ok := decl.(*CustomElementDeclaration)
			if !ok || ced.TagName == "" {
				continue
			}
			tags = append(tags, ced.TagName)
		}
	}
	return tags
}

// RenderableModules returns a slice of RenderableModule for all modules.
func (x *Package) RenderableModules() (modules []*RenderableModule) {
	ms := make(map[string]RenderableModule)
	for i := range x.Modules {
		module := &x.Modules[i]
		ms[module.Path] = RenderableModule{Path: module.Path, Module: module}
		for _, e := range module.Exports {
			if cee, ok := e.(*CustomElementExport); ok {
				c := ms[module.Path]
				c.CustomElementExports = append(ms[module.Path].CustomElementExports, *cee)
				ms[module.Path] = c
			}
		}
	}
	for i := range ms {
		m := ms[i]
		modules = append(modules, &m)
	}
	slices.SortStableFunc(modules, func(a *RenderableModule, b *RenderableModule) int {
		return strings.Compare(a.Path, b.Path)
	})
	return modules
}

// RenderableCustomElementDeclarations returns a slice of RenderableCustomElement for all custom elements in all modules.
func (x *Package) RenderableCustomElementDeclarations() (tags []*RenderableCustomElementDeclaration) {
	for _, m := range x.Modules {
		mrs := make(map[string]*RenderableCustomElementDeclaration)
		for _, d := range m.Declarations {
			if ced, ok := d.(*CustomElementDeclaration); ok {
				mrs[ced.TagName] = NewRenderableCustomElementDeclaration(ced, &m, x)
			}
		}
		for _, e := range m.Exports {
			if cee, ok := e.(*CustomElementExport); ok {
				r := mrs[cee.Name]
				r.CustomElementExport = cee
				mrs[cee.Name] = r
			}
		}
		for i := range mrs {
			tag := mrs[i]
			tags = append(tags, tag)
		}
	}
	slices.SortStableFunc(tags, func(a *RenderableCustomElementDeclaration, b *RenderableCustomElementDeclaration) int {
		return int(a.CustomElementDeclaration.StartByte - b.CustomElementDeclaration.StartByte)
	})
	return tags
}

// findCustomElementContext locates the first CustomElementDeclaration, CustomElementExport, and JavaScriptModule for a given tagName.
func (x *Package) findCustomElementContext(tagName string) (*CustomElementDeclaration, *CustomElementExport, *JavaScriptModule, error) {
	for _, m := range x.Modules {
		for _, d := range m.Declarations {
			if ced, ok := d.(*CustomElementDeclaration); ok && ced.TagName == tagName {
				var ceExport *CustomElementExport
				for _, e := range m.Exports {
					if cee, ok := e.(*CustomElementExport); ok && cee.Name == tagName {
						ceExport = cee
						break
					}
				}
				return ced, ceExport, &m, nil
			}
		}
	}
	return nil, nil, nil, errors.New("Tag not found: " + tagName)
}

// TagRenderableAttributes returns attributes for a given tag name with context.
func (x *Package) TagRenderableAttributes(tagName string) (attrs []*RenderableAttribute, err error) {
	ced, ceExport, m, err := x.findCustomElementContext(tagName)
	if err != nil {
		return nil, err
	}
	fieldMap := make(map[string]CustomElementField)
	for _, member := range ced.Members {
		if cef, ok := member.(*CustomElementField); ok {
			fieldMap[cef.Attribute] = *cef
		}
	}
	attrMap := make(map[string]RenderableAttribute)
	for _, attr := range ced.Attributes {
		attrCopy := attr // Create a copy of the loop variable
		var field *CustomElementField
		if f, ok := fieldMap[attrCopy.Name]; ok {
			field = &f
		}
		attrMap[attrCopy.Name] = RenderableAttribute{
			Attribute:                &attrCopy,
			CustomElementDeclaration: ced,
			CustomElementField:       field,
			CustomElementExport:      ceExport,
			JavaScriptModule:         m,
		}
	}
	for _, awc := range attrMap {
		attrs = append(attrs, &awc)
	}
	slices.SortStableFunc(attrs, func(a *RenderableAttribute, b *RenderableAttribute) int {
		if a.Attribute != nil && b.Attribute != nil {
			return int(a.Attribute.StartByte - b.Attribute.StartByte)
		}
		return 0
	})
	return attrs, nil
}

// TagRenderableSlots returns slots for a given tag name with context.
func (x *Package) TagRenderableSlots(tagName string) (slots []*RenderableSlot, err error) {
	ced, ceExport, m, err := x.findCustomElementContext(tagName)
	if err != nil {
		return nil, err
	}
	for i := range ced.Slots {
		slot := &ced.Slots[i]
		slots = append(slots, &RenderableSlot{
			Slot:                     slot,
			CustomElementDeclaration: ced,
			CustomElementExport:      ceExport,
			JavaScriptModule:         m,
		})
	}
	return slots, nil
}

// TagRenderableCssProperties returns CSS custom properties for a given tag name with context.
func (x *Package) TagRenderableCssProperties(tagName string) (props []*RenderableCssCustomProperty, err error) {
	ced, ceExport, m, err := x.findCustomElementContext(tagName)
	if err != nil {
		return nil, err
	}
	for i := range ced.CssProperties {
		prop := &ced.CssProperties[i]
		props = append(props, &RenderableCssCustomProperty{
			CssCustomProperty:        prop,
			CustomElementDeclaration: ced,
			CustomElementExport:      ceExport,
			JavaScriptModule:         m,
		})
	}
	return props, nil
}

// TagRenderableCssStates returns CSS custom states for a given tag name with context.
func (x *Package) TagRenderableCssStates(tagName string) (states []*RenderableCssCustomState, err error) {
	ced, ceExport, m, err := x.findCustomElementContext(tagName)
	if err != nil {
		return nil, err
	}
	for i := range ced.CssStates {
		state := &ced.CssStates[i]
		states = append(states, &RenderableCssCustomState{
			CssCustomState:           state,
			CustomElementDeclaration: ced,
			CustomElementExport:      ceExport,
			JavaScriptModule:         m,
		})
	}
	return states, nil
}

// TagRenderableCssParts returns CSS shadow parts for a given tag name with context.
func (x *Package) TagRenderableCssParts(tagName string) (parts []*RenderableCssPart, err error) {
	ced, ceExport, m, err := x.findCustomElementContext(tagName)
	if err != nil {
		return nil, err
	}
	for i := range ced.CssParts {
		part := &ced.CssParts[i]
		parts = append(parts, &RenderableCssPart{
			CssPart:                  part,
			CustomElementDeclaration: ced,
			CustomElementExport:      ceExport,
			JavaScriptModule:         m,
		})
	}
	return parts, nil
}

// TagRenderableEvents returns events for a given tag name with context.
func (x *Package) TagRenderableEvents(tagName string) (events []*RenderableEvent, err error) {
	ced, ceExport, m, err := x.findCustomElementContext(tagName)
	if err != nil {
		return nil, err
	}
	for i := range ced.Events {
		event := &ced.Events[i]
		events = append(events, &RenderableEvent{
			Event:                    event,
			CustomElementDeclaration: ced,
			CustomElementExport:      ceExport,
			JavaScriptModule:         m,
		})
	}
	return events, nil
}

// TagRenderableClassMethods returns methods for a given tag name with context.
func (x *Package) TagRenderableClassMethods(tagName string) (methods []*RenderableClassMethod, err error) {
	ced, ceExport, m, err := x.findCustomElementContext(tagName)
	if err != nil {
		return nil, err
	}
	for _, member := range ced.Members {
		if method, ok := member.(*ClassMethod); ok {
			methods = append(methods, &RenderableClassMethod{
				Method:                   method,
				CustomElementDeclaration: ced,
				CustomElementExport:      ceExport,
				JavaScriptModule:         m,
			})
		}
	}
	return methods, nil
}

func toTreeChildren(xs []Renderable, p PredicateFunc) (nodes []pterm.TreeNode) {
	for _, n := range xs {
		var children []pterm.TreeNode
			// If this Renderable knows how to group its children, use that
			if gr, ok := n.(GroupedRenderable); ok {
					children = append(nodes, gr.GroupedChildren(p)...)
			} else {
				// Otherwise, use the default recursion+filtering
				children = toTreeChildren(n.Children(), p)
			}

			if p(n) || len(children) > 0 {
					node := n.ToTreeNode(p)
					node.Children = children
					nodes = append(nodes, node)
			}
	}
	return nodes
}

func tn(text string, children... pterm.TreeNode) pterm.TreeNode {
	return pterm.TreeNode{Text: text, Children: children}
}
