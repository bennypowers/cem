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
	"github.com/pterm/pterm"
)

var _ Renderable = (*RenderableDemo)(nil)

// Demo for custom elements.
type Demo struct {
	Description string           `json:"description,omitempty"`
	URL         string           `json:"url"`
	Source      *SourceReference `json:"source,omitempty"`
}

type RenderableDemo struct {
	Demo                     *Demo
	CustomElementDeclaration *CustomElementDeclaration
	Module                   *Module
	Package                  *Package
	ChildNodes               []Renderable
}

func NewRenderableDemo(
	demo *Demo,
	ced *CustomElementDeclaration,
	mod *Module,
	pkg *Package,
) *RenderableDemo {
	r := RenderableDemo{
		Demo:                     demo,
		CustomElementDeclaration: ced,
		Module:                   mod,
		Package:                  pkg,
	}
	return &r
}

func (x *RenderableDemo) Name() string {
	return x.Demo.URL
}

func (x *RenderableDemo) Label() string {
	return highlightIfDeprecated(x)
}

func (x *RenderableDemo) IsDeprecated() bool {
	return false
}

func (x *RenderableDemo) Deprecation() Deprecated {
	return nil
}

func (x *RenderableDemo) Children() []Renderable {
	return nil
}

func (x *RenderableDemo) GroupedChildren(p PredicateFunc) []pterm.TreeNode {
	return nil
}

func (x *RenderableDemo) ColumnHeadings() []string {
	return []string{
		"URL",
		"Source",
		"Description",
	}
}

// Renders a Demo as a table row.

func (x *RenderableDemo) ToTableRow() []string {
	if x.Demo == nil {
		return []string{}
	}
	href := ""
	if x.Demo.Source != nil {
		href = x.Demo.Source.Href
	}
	return []string{
		x.Demo.URL,
		href,
		x.Demo.Description,
	}
}

func (x *RenderableDemo) ToTreeNode(p PredicateFunc) pterm.TreeNode {
	return tn(x.Label(), x.GroupedChildren(p)...)
}

func (x *RenderableDemo) Sections() []Section {
	return []Section{}
}

func (x *RenderableDemo) Summary() string {
	return ""
}

func (x *RenderableDemo) Description() string {
	return x.Demo.Description
}
