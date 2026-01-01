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
		SchemaVersion: "2.1.0",
		Modules:       modules,
	}
}

func (x *Package) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

// Clone creates a deep copy of the Package.
// This is significantly faster than JSON serialization (~5-10x performance improvement).
//
// Usage:
//
//	cloned := original.Clone()
//
// Performance: ~microseconds vs ~milliseconds for JSON round-trip
// Thread Safety: Safe for concurrent use (creates new instance)
func (p *Package) Clone() *Package {
	if p == nil {
		return nil
	}

	cloned := &Package{
		SchemaVersion: p.SchemaVersion,
		Readme:        cloneStringPtr(p.Readme),
	}

	if p.Deprecated != nil {
		cloned.Deprecated = p.Deprecated.Clone()
	}

	// Deep clone modules
	if len(p.Modules) > 0 {
		cloned.Modules = make([]Module, len(p.Modules))
		for i, module := range p.Modules {
			cloned.Modules[i] = *module.Clone()
		}
	} else {
		cloned.Modules = []Module{} // Maintain empty slice vs nil consistency
	}

	// Set backreferences from modules to package (same as UnmarshalJSON)
	for i := range cloned.Modules {
		cloned.Modules[i].Package = cloned
	}

	return cloned
}

// FindDeclaration resolves a Reference to its Declaration.
// Searches all modules in the package for a matching declaration.
// If ref.Module is specified, only searches in that module.
// Returns nil if not found or if package/reference is nil.
func (p *Package) FindDeclaration(ref Reference) Declaration {
	if p == nil || ref.Name == "" {
		return nil
	}

	for i := range p.Modules {
		mod := &p.Modules[i]

		// If module path specified, skip other modules
		if ref.Module != "" && mod.Path != ref.Module {
			continue
		}

		// Search declarations in this module
		for _, decl := range mod.Declarations {
			if decl.Name() == ref.Name {
				return decl
			}
		}
	}

	return nil
}

// FindCustomElementDeclaration finds the CustomElementDeclaration for a given tag name.
// This provides access to the full declaration including summary and description,
// which are not available on the CustomElement struct alone.
//
// Returns nil if the package is nil, the tag name is empty, or no matching element is found.
//
// Usage:
//
//	decl := pkg.FindCustomElementDeclaration("my-button")
//	if decl != nil {
//		fmt.Println(decl.Summary)
//		fmt.Println(decl.Description)
//	}
func (p *Package) FindCustomElementDeclaration(tagName string) *CustomElementDeclaration {
	if p == nil || tagName == "" {
		return nil
	}

	for i := range p.Modules {
		for j := range p.Modules[i].Declarations {
			// Check if this is a CustomElementDeclaration with matching tagName
			if customEl, ok := p.Modules[i].Declarations[j].(*CustomElementDeclaration); ok {
				if customEl.TagName == tagName {
					return customEl
				}
			}
		}
	}

	return nil
}

func (x *Package) UnmarshalJSON(data []byte) error {
	type Rest Package
	aux := &struct {
		Modules    []json.RawMessage `json:"modules"`
		Deprecated json.RawMessage   `json:"deprecated"`
		*Rest
	}{
		Rest: (*Rest)(x),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Handle deprecated field
	if len(aux.Deprecated) > 0 && string(aux.Deprecated) != "null" {
		var dep Deprecated
		if !decodeDeprecatedField(&dep, aux.Deprecated) {
			return fmt.Errorf("invalid type for deprecated field")
		}
		x.Deprecated = dep
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

	// Set backreferences from modules to package
	for i := range x.Modules {
		x.Modules[i].Package = x

		// Update declaration backreferences to point to module in slice (not local var)
		for j := range x.Modules[i].Declarations {
			switch decl := x.Modules[i].Declarations[j].(type) {
			case *ClassDeclaration:
				decl.Module = &x.Modules[i]
			case *CustomElementDeclaration:
				decl.Module = &x.Modules[i]
			case *MixinDeclaration:
				decl.Module = &x.Modules[i]
			case *CustomElementMixinDeclaration:
				decl.Module = &x.Modules[i]
			case *FunctionDeclaration:
				decl.Module = &x.Modules[i]
			case *VariableDeclaration:
				decl.Module = &x.Modules[i]
			}
		}
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
