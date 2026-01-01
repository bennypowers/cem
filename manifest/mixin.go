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

var _ Deprecatable = (*MixinDeclaration)(nil)
var _ Deprecatable = (*CustomElementMixinDeclaration)(nil)
var _ Renderable = (*RenderableMixinDeclaration)(nil)
var _ Renderable = (*RenderableCustomElementMixinDeclaration)(nil)

// MixinDeclaration describes a class mixin.
type MixinDeclaration struct {
	ClassLike
	FunctionLike
	FullyQualified
	Deprecated Deprecated `json:"deprecated,omitempty"`
	Kind       string     `json:"kind"` // 'mixin'
}

// CustomElementMixinDeclaration extends MixinDeclaration and CustomElement.
type CustomElementMixinDeclaration struct {
	MixinDeclaration
	CustomElementDeclaration
	FullyQualified
}

func (*MixinDeclaration) isDeclaration() {}

func (m *MixinDeclaration) Name() string {
	if m == nil {
		return ""
	}
	return m.FullyQualified.Name
}

// Clone creates a deep copy of the MixinDeclaration.
// Handles all embedded structures including ClassLike, FunctionLike, and FullyQualified.
//
// Performance: Efficient deep copying without JSON serialization overhead
// Thread Safety: Safe for concurrent use (creates new instance)
func (m *MixinDeclaration) Clone() Declaration {
	if m == nil {
		return nil
	}

	cloned := &MixinDeclaration{
		Kind: m.Kind,
	}

	if m.Deprecated != nil {
		cloned.Deprecated = m.Deprecated.Clone()
	}

	// Clone embedded structures
	cloned.ClassLike = m.ClassLike.Clone()
	cloned.FunctionLike = m.FunctionLike.Clone()
	cloned.FullyQualified = m.FullyQualified.Clone()

	return cloned
}

func (x *MixinDeclaration) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

func (x *MixinDeclaration) GetStartByte() uint { return x.FunctionLike.StartByte }

func (m *MixinDeclaration) UnmarshalJSON(data []byte) error {
	type Rest MixinDeclaration
	aux := &struct {
		Deprecated json.RawMessage   `json:"deprecated"`
		Members    []json.RawMessage `json:"members"`
		*Rest
	}{
		Rest: (*Rest)(m),
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
		m.Deprecated = dep
	}
	// Handle members
	m.Members = nil
	for _, mem := range aux.Members {
		member, err := unmarshalClassMember(mem)
		if err == nil && member != nil {
			m.Members = append(m.Members, member)
		}
	}
	if m.Members == nil {
		m.Members = []ClassMember{}
	}
	return nil
}

type RenderableMixinDeclaration struct {
	MixinDeclaration *MixinDeclaration
	JavaScriptExport *JavaScriptExport
	Module           *Module
	Package          *Package
	ChildNodes       []Renderable
}

func NewRenderableMixinDeclaration(
	md *MixinDeclaration,
	mod *Module,
	pkg *Package,
) *RenderableMixinDeclaration {
	// TODO: calculate export
	var export *JavaScriptExport
	// TODO: populate children with params, return
	children := make([]Renderable, 0)
	return &RenderableMixinDeclaration{
		MixinDeclaration: md,
		JavaScriptExport: export,
		Module:           mod,
		Package:          pkg,
		ChildNodes:       children,
	}
}

func (*CustomElementMixinDeclaration) isDeclaration() {}

func (c *CustomElementMixinDeclaration) Name() string {
	if c == nil {
		return ""
	}
	return c.FullyQualified.Name
}

// Clone creates a deep copy of the CustomElementMixinDeclaration.
// Handles all embedded structures including MixinDeclaration, CustomElementDeclaration, and FullyQualified.
//
// Performance: Efficient deep copying without JSON serialization overhead
// Thread Safety: Safe for concurrent use (creates new instance)
func (c *CustomElementMixinDeclaration) Clone() Declaration {
	if c == nil {
		return nil
	}

	cloned := &CustomElementMixinDeclaration{}

	// Clone the embedded MixinDeclaration
	if mixinDecl := c.MixinDeclaration.Clone(); mixinDecl != nil {
		if md, ok := mixinDecl.(*MixinDeclaration); ok {
			cloned.MixinDeclaration = *md
		}
	}

	// Clone the embedded CustomElementDeclaration
	if customElementDecl := c.CustomElementDeclaration.Clone(); customElementDecl != nil {
		if ced, ok := customElementDecl.(*CustomElementDeclaration); ok {
			cloned.CustomElementDeclaration = *ced
		}
	}

	// Clone the embedded FullyQualified
	cloned.FullyQualified = c.FullyQualified.Clone()

	return cloned
}

func (x *RenderableMixinDeclaration) Name() string {
	return x.MixinDeclaration.Name()
}

func (x *RenderableMixinDeclaration) Label() string {
	return pterm.LightBlue("mixin") + " " + highlightIfDeprecated(x)
}

func (x *RenderableMixinDeclaration) IsDeprecated() bool {
	return x.MixinDeclaration.IsDeprecated()
}

func (x *RenderableMixinDeclaration) Deprecation() Deprecated {
	return x.MixinDeclaration.Deprecated
}

func (x *RenderableMixinDeclaration) Children() []Renderable {
	return x.ChildNodes
}

func (x *RenderableMixinDeclaration) ColumnHeadings() []string {
	// TODO: elaborate
	return []string{"Name"}
}

func (x *RenderableMixinDeclaration) ToTableRow() []string {
	// TODO: maybe this gets a section, not a row?
	return []string{
		highlightIfDeprecated(x),
	}
}

func (x *RenderableMixinDeclaration) ToTreeNode(p PredicateFunc) pterm.TreeNode {
	// TODO: grouped params, return, class stuff as children
	return tn(x.Label(), toTreeChildren(x.Children(), p)...)
}

func (x *CustomElementMixinDeclaration) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

func (x *CustomElementMixinDeclaration) GetStartByte() uint { return x.FunctionLike.StartByte }

func (m *CustomElementMixinDeclaration) UnmarshalJSON(data []byte) error {
	type Rest CustomElementMixinDeclaration
	aux := &struct {
		Members    []json.RawMessage `json:"members"`
		Deprecated json.RawMessage   `json:"deprecated"`
		*Rest
	}{
		Rest: (*Rest)(m),
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
		m.Deprecated = dep
	}

	m.Members = nil
	for _, mem := range aux.Members {
		member, err := unmarshalClassMember(mem)
		if err != nil {
			return fmt.Errorf("cannot unmarshal member: %w", err)
		}
		m.Members = append(m.Members, member)
	}
	return nil
}

type RenderableCustomElementMixinDeclaration struct {
	TagName                       string
	CustomElementMixinDeclaration *CustomElementMixinDeclaration
	JavaScriptExport              *JavaScriptExport
	Module                        *Module
	Package                       *Package
	ChildNodes                    []Renderable
}

func NewRenderableCustomElementMixinDeclaration(
	cemd *CustomElementMixinDeclaration,
	mod *Module,
	pkg *Package,
) *RenderableCustomElementMixinDeclaration {
	// TODO: calculate export
	var export *JavaScriptExport
	// TODO: populate children with params, return, class stuff
	// Preferably reuse stuff from NewRenderableMixinDeclaration
	children := make([]Renderable, 0)
	return &RenderableCustomElementMixinDeclaration{
		TagName:          cemd.TagName,
		JavaScriptExport: export,
		Module:           mod,
		Package:          pkg,
		ChildNodes:       children,
	}
}

func (x *RenderableCustomElementMixinDeclaration) IsDeprecated() bool {
	return x.CustomElementMixinDeclaration.IsDeprecated()
}

func (x *RenderableCustomElementMixinDeclaration) Deprecation() Deprecated {
	return x.CustomElementMixinDeclaration.Deprecated
}

func (x *RenderableCustomElementMixinDeclaration) Children() []Renderable {
	return x.ChildNodes
}

func (x *RenderableCustomElementMixinDeclaration) Name() string {
	return x.CustomElementMixinDeclaration.Name()
}

func (x *RenderableCustomElementMixinDeclaration) Label() string {
	return pterm.LightBlue("custom element mixin") + " " + highlightIfDeprecated(x)
}

func (x *RenderableCustomElementMixinDeclaration) ColumnHeadings() []string {
	// TODO: elaborate
	return []string{"Name"}
}

func (x *RenderableCustomElementMixinDeclaration) ToTableRow() []string {
	// TODO: maybe this gets a section, not a row?
	return []string{
		highlightIfDeprecated(x),
	}
}

func (x *RenderableCustomElementMixinDeclaration) ToTreeNode(p PredicateFunc) pterm.TreeNode {
	// TODO: group children in constructor
	return tn(x.Label(), toTreeChildren(x.Children(), p)...)
}
