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

var _ Deprecatable = (*ClassDeclaration)(nil)
var _ Renderable = (*RenderableClassDeclaration)(nil)
var _ GroupedRenderable = (*RenderableClassDeclaration)(nil)

// ClassLike is the common interface of classes and mixins.
type ClassLike struct {
	FullyQualified
	StartByte  uint             `json:"-"`
	Superclass *Reference       `json:"superclass,omitempty"`
	Mixins     []Reference      `json:"mixins,omitempty"`
	Members    []ClassMember    `json:"members,omitempty"`
	Source     *SourceReference `json:"source,omitempty"`
}

// Clone creates a deep copy of the ClassLike structure.
// Handles all nested references, mixins, and class members.
func (c ClassLike) Clone() ClassLike {
	cloned := ClassLike{
		FullyQualified: c.FullyQualified.Clone(),
		StartByte:      c.StartByte,
	}

	if c.Superclass != nil {
		superclass := c.Superclass.Clone()
		cloned.Superclass = &superclass
	}

	if len(c.Mixins) > 0 {
		cloned.Mixins = make([]Reference, len(c.Mixins))
		for i, mixin := range c.Mixins {
			cloned.Mixins[i] = mixin.Clone()
		}
	}

	if len(c.Members) > 0 {
		cloned.Members = make([]ClassMember, len(c.Members))
		for i, member := range c.Members {
			cloned.Members[i] = member.Clone()
		}
	}

	if c.Source != nil {
		source := c.Source.Clone()
		cloned.Source = &source
	}

	return cloned
}

// ClassDeclaration is a class.
type ClassDeclaration struct {
	ClassLike
	Kind       string     `json:"kind"`                 // 'class'
	Deprecated Deprecated `json:"deprecated,omitempty"` // bool or string
}

func (*ClassDeclaration) isDeclaration() {}

// Clone creates a deep copy of the ClassDeclaration.
// Handles all nested ClassMembers and properly clones the embedded ClassLike structure.
//
// Performance: Efficient deep copying without JSON serialization overhead
// Thread Safety: Safe for concurrent use (creates new instance)
func (c *ClassDeclaration) Clone() Declaration {
	if c == nil {
		return nil
	}

	cloned := &ClassDeclaration{
		Kind: c.Kind,
	}

	if c.Deprecated != nil {
		cloned.Deprecated = c.Deprecated.Clone()
	}

	// Clone the embedded ClassLike
	cloned.ClassLike = c.ClassLike.Clone()

	return cloned
}

func (x *ClassDeclaration) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

func (x *ClassDeclaration) GetStartByte() uint { return x.StartByte }

func (c *ClassDeclaration) UnmarshalJSON(data []byte) error {
	type Rest ClassDeclaration
	aux := &struct {
		Deprecated json.RawMessage   `json:"deprecated"`
		Members    []json.RawMessage `json:"members"`
		*Rest
	}{
		Rest: (*Rest)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	// Handle deprecated
	if len(aux.Deprecated) > 0 && string(aux.Deprecated) != "null" {
		var dep Deprecated
		if !decodeDeprecatedField(&dep, aux.Deprecated) {
			return fmt.Errorf("invalid type for deprecated field")
		}
		c.Deprecated = dep
	}
	// Handle members
	for _, m := range aux.Members {
		member, err := unmarshalClassMember(m)
		if err == nil && member != nil {
			c.Members = append(c.Members, member)
		}
	}
	return nil
}

type RenderableClassDeclaration struct {
	ClassDeclaration *ClassDeclaration
	JavaScriptExport *JavaScriptExport
	Module           *Module
	Package          *Package
	ChildNodes       []Renderable
	fields           []Renderable
	methods          []Renderable
}

func NewRenderableClassDeclaration(
	class *ClassDeclaration,
	mod *JavaScriptModule,
	pkg *Package,
) *RenderableClassDeclaration {
	// PERF: use a map
	var je *JavaScriptExport
	for i := range mod.Exports {
		exp := mod.Exports[i]
		if je, ok := exp.(*JavaScriptExport); ok {
			if je.Declaration.Name == class.Name && (je.Declaration.Module == "" || je.Declaration.Module == mod.Path) {
				break
			}
		}
	}
	r := RenderableClassDeclaration{
		ClassDeclaration: class,
		JavaScriptExport: je,
		Module:           mod,
		Package:          pkg,
	}

	for _, m := range class.Members {
		if field, ok := m.(*ClassField); ok {
			m := NewRenderableClassField(field, class, je, mod, pkg)
			r.ChildNodes = append(r.ChildNodes, m)
			r.fields = append(r.fields, m)
		} else if method, ok := m.(*ClassMethod); ok {
			m := NewRenderableClassMethod(method, class, je, mod, pkg)
			r.ChildNodes = append(r.ChildNodes, m)
			r.methods = append(r.methods, m)
		}
	}

	return &r
}

func (x *RenderableClassDeclaration) Name() string {
	return x.ClassDeclaration.Name
}

func (x *RenderableClassDeclaration) Label() string {
	return pterm.LightBlue("class") +
		" " +
		highlightIfDeprecated(x)
}

func (x *RenderableClassDeclaration) IsDeprecated() bool {
	return x.ClassDeclaration.IsDeprecated()
}

func (x *RenderableClassDeclaration) Deprecation() Deprecated {
	return x.ClassDeclaration.Deprecated
}

func (x *RenderableClassDeclaration) Children() []Renderable {
	return x.ChildNodes
}

func (x *RenderableClassDeclaration) GroupedChildren(p PredicateFunc) []pterm.TreeNode {
	var nodes []pterm.TreeNode
	fs := toTreeChildren(x.fields, p)
	ms := toTreeChildren(x.methods, p)
	if len(fs) > 0 {
		nodes = append(nodes, tn(pterm.Blue("Fields"), fs...))
	}
	if len(ms) > 0 {
		nodes = append(nodes, tn(pterm.Blue("Methods"), ms...))
	}
	return nodes
}

func (x *RenderableClassDeclaration) ColumnHeadings() []string {
	return []string{"Name", "Module Path", "Summary"}
}

func (x *RenderableClassDeclaration) ToTableRow() []string {
	modulePath := ""
	if x.Module != nil {
		modulePath = x.Module.Path
	}
	return []string{
		x.ClassDeclaration.Name,
		modulePath,
		x.ClassDeclaration.Summary,
	}
}

func (x *RenderableClassDeclaration) ToTreeNode(p PredicateFunc) pterm.TreeNode {
	return tn(x.Label(), x.GroupedChildren(p)...)
}

func (x *RenderableClassDeclaration) Sections() []Section {
	return []Section{
		{Title: "Fields", Items: x.fields},
		{Title: "Methods", Items: x.methods},
	}
}

func (x *RenderableClassDeclaration) Summary() string {
	return x.ClassDeclaration.Summary
}

func (x *RenderableClassDeclaration) Description() string {
	return x.ClassDeclaration.Description
}
