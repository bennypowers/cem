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

var _ Deprecatable = (*Package)(nil)
var _ Renderable = (*RenderablePackage)(nil)

// Package is the top-level interface of a custom elements manifest file.
type Package struct {
	SchemaVersion string     `json:"schemaVersion"`
	Readme        *string    `json:"readme,omitempty"`
	Modules       []Module   `json:"modules"`
	Deprecated    Deprecated `json:"deprecated,omitempty"` // bool or string
}

func NewPackage(modules []Module) Package {
	return Package{
		SchemaVersion: "1.0.0",
		Modules:       modules,
	}
}

func (x *Package) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

func (x *Package) UnmarshalJSON(data []byte) error {
	type Alias Package
	aux := &struct {
		Modules []json.RawMessage `json:"modules"`
		*Alias
	}{
		Alias: (*Alias)(x),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	x.Modules = nil
	for _, m := range aux.Modules {
		var mod Module
		if err := json.Unmarshal(m, &mod); err == nil {
			x.Modules = append(x.Modules, mod)
		} else {
			return fmt.Errorf("cannot unmarshal module: %w", err)
		}
	}
	if x.Modules == nil {
		x.Modules = []Module{}
	}
	return nil
}

type RenderablePackage struct {
	Package    *Package
	ChildNodes []Renderable
}

func NewRenderablePackage(pkg *Package) *RenderablePackage {
	if pkg == nil {
		return nil
	}
	var children []Renderable
	for i := range pkg.Modules {
		children = append(children, NewRenderableModule(&pkg.Modules[i], pkg))
	}
	return &RenderablePackage{
		Package:    pkg,
		ChildNodes: children,
	}
}

func (x *RenderablePackage) Name() string {
	// TODO: out of band package name
	return "<root>"
}

func (x *RenderablePackage) Label() string {
	return x.Name() // TODO: out of band package name
}

func (x *RenderablePackage) IsDeprecated() bool {
	return x.Package.IsDeprecated()
}

func (x *RenderablePackage) Deprecation() Deprecated {
	return x.Package.Deprecated
}

func (x *RenderablePackage) Children() []Renderable {
	if x == nil || x.ChildNodes == nil {
		return make([]Renderable, 0)
	}
	return x.ChildNodes
}

func (x *RenderablePackage) ColumnHeadings() []string {
	return []string{}
}

func (x *RenderablePackage) ToTableRow() []string {
	return []string{}
}

func (x *RenderablePackage) ToTreeNode(p PredicateFunc) pterm.TreeNode {
	return tn(x.Label(), toTreeChildren(x.Children(), p)...)
}

