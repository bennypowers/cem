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

var _ Deprecatable = (*FunctionDeclaration)(nil)
var _ Renderable = (*RenderableFunctionDeclaration)(nil)

// Return value for functions.
type Return struct {
	Type        *Type  `json:"type,omitempty"`
	Summary     string `json:"summary,omitempty"`
	Description string `json:"description,omitempty"`
}

// Clone creates a deep copy of the Return structure.
func (r Return) Clone() Return {
	cloned := Return{
		Summary:     r.Summary,
		Description: r.Description,
	}

	if r.Type != nil {
		cloned.Type = r.Type.Clone()
	}

	return cloned
}

// FunctionLike is the common interface of functions and mixins.
type FunctionLike struct {
	StartByte  uint        `json:"-"`
	Parameters []Parameter `json:"parameters,omitempty"`
	Return     *Return     `json:"return,omitempty"`
}

// Clone creates a deep copy of the FunctionLike structure.
// Handles all parameters and return type with proper deep copying.
func (f FunctionLike) Clone() FunctionLike {
	cloned := FunctionLike{
		StartByte: f.StartByte,
	}

	if len(f.Parameters) > 0 {
		cloned.Parameters = make([]Parameter, len(f.Parameters))
		for i, param := range f.Parameters {
			cloned.Parameters[i] = param.Clone()
		}
	}

	if f.Return != nil {
		returnClone := f.Return.Clone()
		cloned.Return = &returnClone
	}

	return cloned
}

// CssPart describes a CSS part.
// FunctionDeclaration is a function.
type FunctionDeclaration struct {
	FunctionLike
	FullyQualified
	Deprecated Deprecated       `json:"deprecated,omitempty"` // bool or string
	Kind       string           `json:"kind"`                 // 'function'
	Source     *SourceReference `json:"source,omitempty"`
}

func (*FunctionDeclaration) isDeclaration() {}

// Clone creates a deep copy of the FunctionDeclaration.
// Handles the embedded FunctionLike and FullyQualified structures with parameters and return type.
//
// Performance: Efficient deep copying without JSON serialization overhead
// Thread Safety: Safe for concurrent use (creates new instance)
func (f *FunctionDeclaration) Clone() Declaration {
	if f == nil {
		return nil
	}

	cloned := &FunctionDeclaration{
		Kind: f.Kind,
	}

	if f.Deprecated != nil {
		cloned.Deprecated = f.Deprecated.Clone()
	}

	// Clone the embedded FunctionLike
	cloned.FunctionLike = f.FunctionLike.Clone()

	// Clone the embedded FullyQualified
	cloned.FullyQualified = f.FullyQualified.Clone()

	if f.Source != nil {
		source := f.Source.Clone()
		cloned.Source = &source
	}

	return cloned
}

func (x *FunctionDeclaration) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

func (x *FunctionDeclaration) GetStartByte() uint { return x.StartByte }

func (f *FunctionDeclaration) UnmarshalJSON(data []byte) error {
	type Rest FunctionDeclaration
	aux := &struct {
		Parameters []json.RawMessage `json:"parameters"`
		Deprecated json.RawMessage   `json:"deprecated"`
		*Rest
	}{
		Rest: (*Rest)(f),
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
		f.Deprecated = dep
	}
	// Unmarshal parameters using custom logic (since Parameter has its own UnmarshalJSON)
	f.Parameters = nil
	for _, p := range aux.Parameters {
		var param Parameter
		if err := json.Unmarshal(p, &param); err != nil {
			return fmt.Errorf("unmarshal parameter: %w", err)
		}
		f.Parameters = append(f.Parameters, param)
	}
	return nil
}

type RenderableFunctionDeclaration struct {
	FunctionDeclaration *FunctionDeclaration
	JavaScriptExport    *JavaScriptExport
	Module              *Module
	Package             *Package
	ChildNodes          []Renderable // parameters and return
}

func NewRenderableFunctionDeclaration(
	fd *FunctionDeclaration,
	mod *Module,
	pkg *Package,
) *RenderableFunctionDeclaration {
	// TODO use a map
	// TODO get je, cee from other modules
	var je *JavaScriptExport
	for i, exp := range mod.Exports {
		if eje, ok := exp.(*JavaScriptExport); ok {
			if eje.Declaration.Name == fd.Name && (eje.Declaration.Module == "" || eje.Declaration.Module == mod.Path) {
				je = mod.Exports[i].(*JavaScriptExport)
				break
			}
		}
	}
	// TODO: populate children with params, return
	children := make([]Renderable, 0)
	return &RenderableFunctionDeclaration{
		FunctionDeclaration: fd,
		JavaScriptExport:    je,
		Module:              mod,
		Package:             pkg,
		ChildNodes:          children,
	}
}

func (x *RenderableFunctionDeclaration) Name() string {
	return x.FunctionDeclaration.Name
	// label := pterm.LightBlue("function") + " " + highlightIfDeprecated(x.FunctionDeclaration.Name, x)
}

func (x *RenderableFunctionDeclaration) Label() string {
	return pterm.LightBlue("function") + " " + highlightIfDeprecated(x)
}

func (x *RenderableFunctionDeclaration) IsDeprecated() bool {
	return x.FunctionDeclaration.IsDeprecated()
}

func (x *RenderableFunctionDeclaration) Deprecation() Deprecated {
	return x.FunctionDeclaration.Deprecated
}

func (x *RenderableFunctionDeclaration) Children() []Renderable {
	return nil // it's a leaf node
}

func (x *RenderableFunctionDeclaration) ColumnHeadings() []string {
	return []string{"Name", "Return Type", "Summary"}
}

func (x *RenderableFunctionDeclaration) ToTableRow() []string {
	typeText := ""
	if x.FunctionDeclaration.Return != nil && x.FunctionDeclaration.Return.Type != nil {
		typeText = x.FunctionDeclaration.Return.Type.Text
	}
	return []string{
		highlightIfDeprecated(x),
		typeText,
		x.FunctionDeclaration.Summary,
	}
}

func (x *RenderableFunctionDeclaration) ToTreeNode(pred PredicateFunc) pterm.TreeNode {
	return tn(x.Label(), toTreeChildren(x.Children(), pred)...)
}
