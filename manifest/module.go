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

// Module may expand in future; currently only JavaScriptModule.
type Module = JavaScriptModule

var _ Deprecatable = (*Module)(nil)
var _ Renderable = (*RenderableModule)(nil)

type JavaScriptModule struct {
	Kind         string        `json:"kind"` // 'javascript-module'
	Path         string        `json:"path"`
	Summary      string        `json:"summary,omitempty"`
	Description  string        `json:"description,omitempty"`
	Declarations []Declaration `json:"declarations,omitempty"`
	Exports      []Export      `json:"exports,omitempty"`
	Deprecated   Deprecated    `json:"deprecated,omitempty"` // bool or string
}

func NewModule(file string) *Module {
	return &Module{
		Kind: "javascript-module",
		Path: normalizePath(file),
	}
}

func (m *Module) UnmarshalJSON(data []byte) error {
	type Rest Module
	aux := &struct {
		Declarations []json.RawMessage `json:"declarations"`
		Exports      []json.RawMessage `json:"exports"`
		Deprecated   json.RawMessage   `json:"deprecated"`
		*Rest
	}{
		Rest: (*Rest)(m),
	}
	// Remove m.Declarations before unmarshaling so we control its population
	m.Declarations = nil

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if len(aux.Deprecated) > 0 && string(aux.Deprecated) != "null" {
		var dep Deprecated
		if !decodeDeprecatedField(&dep, aux.Deprecated) {
			return fmt.Errorf("invalid type for deprecated field")
		}
		m.Deprecated = dep
	}

	for _, d := range aux.Declarations {
		decl, err := unmarshalDeclaration(d)
		if err != nil {
			return fmt.Errorf("cannot unmarshal declaration: %w", err)
		}
		m.Declarations = append(m.Declarations, decl)
	}
	if m.Declarations == nil {
		m.Declarations = []Declaration{}
	}

	for _, e := range aux.Exports {
		export, err := unmarshalExport(e)
		if err != nil {
			return fmt.Errorf("cannot unmarshal export: %w", err)
		}
		m.Exports = append(m.Exports, export)
	}
	if m.Exports == nil {
		m.Exports = []Export{}
	}

	return nil
}

// Clone creates a deep copy of the JavaScriptModule.
// Handles all nested declarations and exports with proper deep copying.
//
// Usage:
//
//	cloned := original.Clone()
//
// Performance: Significantly faster than JSON serialization for complex modules
// Thread Safety: Safe for concurrent use (creates new instance)
func (m *JavaScriptModule) Clone() *JavaScriptModule {
	if m == nil {
		return nil
	}

	cloned := &JavaScriptModule{
		Kind:        m.Kind,
		Path:        m.Path,
		Summary:     m.Summary,
		Description: m.Description,
	}

	if m.Deprecated != nil {
		cloned.Deprecated = m.Deprecated.Clone()
	}

	// Clone declarations
	if len(m.Declarations) > 0 {
		cloned.Declarations = make([]Declaration, len(m.Declarations))
		for i, decl := range m.Declarations {
			cloned.Declarations[i] = decl.Clone()
		}
	} else {
		cloned.Declarations = []Declaration{} // Maintain consistency with unmarshal
	}

	// Clone exports
	if len(m.Exports) > 0 {
		cloned.Exports = make([]Export, len(m.Exports))
		for i, export := range m.Exports {
			cloned.Exports[i] = export.Clone()
		}
	} else {
		cloned.Exports = []Export{} // Maintain consistency
	}

	return cloned
}

func (x *Module) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

type RenderableModule struct {
	Path                 string
	Module               *Module
	Package              *Package
	CustomElementExports []CustomElementExport
	ChildNodes           []Renderable
}

func NewRenderableModule(
	mod *Module,
	pkg *Package,
) *RenderableModule {
	// TODO: populate children with declarations
	// TODO: populate exports with exports
	children := make([]Renderable, 0)
	exports := make([]CustomElementExport, 0)
	for i, decl := range mod.Declarations {
		switch decl.(type) {
		case *CustomElementDeclaration:
			ced := mod.Declarations[i].(*CustomElementDeclaration)
			children = append(children, NewRenderableCustomElementDeclaration(ced, mod, pkg))
		case *ClassDeclaration:
			cd := mod.Declarations[i].(*ClassDeclaration)
			children = append(children, NewRenderableClassDeclaration(cd, mod, pkg))
		case *FunctionDeclaration:
			fd := mod.Declarations[i].(*FunctionDeclaration)
			children = append(children, NewRenderableFunctionDeclaration(fd, mod, pkg))
		case *VariableDeclaration:
			vd := mod.Declarations[i].(*VariableDeclaration)
			children = append(children, NewRenderableVariableDeclaration(vd, mod, pkg))
		case *MixinDeclaration:
			md := mod.Declarations[i].(*MixinDeclaration)
			children = append(children, NewRenderableMixinDeclaration(md, mod, pkg))
		case *CustomElementMixinDeclaration:
			cemd := mod.Declarations[i].(*CustomElementMixinDeclaration)
			children = append(children, NewRenderableCustomElementMixinDeclaration(cemd, mod, pkg))
		}
	}
	return &RenderableModule{
		Path:                 mod.Path,
		Module:               mod,
		Package:              pkg,
		CustomElementExports: exports,
		ChildNodes:           children,
	}
}

func (x *RenderableModule) Name() string {
	return x.Module.Path
}

func (x *RenderableModule) Label() string {
	return pterm.LightBlue("module") + " " + highlightIfDeprecated(x)
}

func (x *RenderableModule) IsDeprecated() bool {
	return x.Module.IsDeprecated()
}

func (x *RenderableModule) Deprecation() Deprecated {
	return x.Module.Deprecated
}

func (x *RenderableModule) Children() []Renderable {
	return x.ChildNodes
}

func (x *RenderableModule) ColumnHeadings() []string {
	return []string{"Path", "Tag Names"}
}

func (x *RenderableModule) ToTableRow() []string {
	tags := make([]string, 0)
	for _, cee := range x.CustomElementExports {
		tags = append(tags, cee.Name)
	}
	return []string{
		x.Path,
		strings.Join(tags, ", "),
	}
}

func (x *RenderableModule) ToTreeNode(p PredicateFunc) pterm.TreeNode {
	return tn(x.Label(), toTreeChildren(x.Children(), p)...)
}
