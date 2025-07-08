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

// ClassDeclaration is a class.
type ClassDeclaration struct {
	ClassLike
	Kind       string     `json:"kind"`                 // 'class'
	Deprecated Deprecated `json:"deprecated,omitempty"` // bool or string
}

func (*ClassDeclaration) isDeclaration() {}

func (x *ClassDeclaration) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

func (x *ClassDeclaration) GetStartByte() uint { return x.StartByte }

func (c *ClassDeclaration) UnmarshalJSON(data []byte) error {
	type Alias ClassDeclaration
	aux := &struct {
		Deprecated json.RawMessage   `json:"deprecated"`
		Members    []json.RawMessage `json:"members"`
		*Alias
	}{
		Alias: (*Alias)(c),
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
				exp = je
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
	sum := ""
	if x.ClassDeclaration.Summary != "" {
		sum = pterm.Gray(x.ClassDeclaration.Summary)
	}
	return strings.TrimSpace(
		pterm.LightBlue("class") +
			" " +
			highlightIfDeprecated(x) +
			" " +
			sum,
	)
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
