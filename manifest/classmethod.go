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
	"strconv"
	"strings"

	"github.com/pterm/pterm"
)

var _ Deprecatable = (*ClassMethod)(nil)
var _ Renderable = (*RenderableClassMethod)(nil)

// ClassMethod is a method.
type ClassMethod struct {
	FunctionLike
	FullyQualified
	Kind          string           `json:"kind"` // 'method'
	Static        bool             `json:"static,omitempty"`
	Privacy       Privacy          `json:"privacy,omitempty"` // 'public', 'private', 'protected'
	InheritedFrom *Reference       `json:"inheritedFrom,omitempty"`
	Source        *SourceReference `json:"source,omitempty"`
	Deprecated    Deprecated       `json:"deprecated,omitempty"` // bool or string
}

func (*ClassMethod) isClassMember() {}

func (x *ClassMethod) GetStartByte() uint { return x.StartByte }

func (x *ClassMethod) IsDeprecated() bool { return x.Deprecated != nil }

func (m *ClassMethod) UnmarshalJSON(data []byte) error {
	type Alias ClassMethod
	aux := &struct {
		Deprecated json.RawMessage   `json:"deprecated"`
		Parameters []json.RawMessage `json:"parameters"`
		*Alias
	}{
		Alias: (*Alias)(m),
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
	// Handle parameters
	m.Parameters = nil
	for _, p := range aux.Parameters {
		var param Parameter
		if err := json.Unmarshal(p, &param); err != nil {
			return fmt.Errorf("unmarshal parameter: %w", err)
		}
		m.Parameters = append(m.Parameters, param)
	}
	return nil
}

type RenderableClassMethod struct {
	Method                   *ClassMethod
	CustomElementDeclaration *CustomElementDeclaration
	CustomElementExport      *CustomElementExport
	ClassDeclaration         *ClassDeclaration
	JavaScriptExport         *JavaScriptExport
	JavaScriptModule         *JavaScriptModule
	Package                  *Package
	ChildNodes               []Renderable
}

func (x *RenderableClassMethod) Name() string {
	return x.Method.Name
}

func (x *RenderableClassMethod) Label() string {
	sum := ""
	if x.Method.Summary != "" {
		sum = pterm.Gray(x.Method.Summary)
	}
	return strings.TrimSpace(highlightIfDeprecated(x) + " " + sum)
}

func (x *RenderableClassMethod) IsDeprecated() bool {
	return x.Method.Deprecated != nil
}

func (x *RenderableClassMethod) Deprecation() Deprecated {
	return x.Method.Deprecated
}

func (x *RenderableClassMethod) Children() []Renderable {
	return x.ChildNodes
}

func (x *RenderableClassMethod) ColumnHeadings() []string {
	return []string{"Name", "Return Type", "Privacy", "Static", "Summary"}
}

// Renders an Event as a table row.
func (x *RenderableClassMethod) ToTableRow() []string {
	returnType := "void"
	privacy := string(x.Method.Privacy)
	if privacy == "" {
		privacy = "public"
	}
	if x.Method.Return != nil && x.Method.Return.Type != nil {
		returnType = x.Method.Return.Type.Text
	}
	return []string{
		highlightIfDeprecated(x),
		returnType,
		privacy,
		strconv.FormatBool(x.Method.Static),
		x.Method.Summary,
	}
}

func (x *RenderableClassMethod) ToTreeNode(pred PredicateFunc) pterm.TreeNode {
	return pterm.TreeNode{Text: x.Label()}
}

func NewRenderableClassMethod(
	method *ClassMethod,
	cd *ClassDeclaration,
	ce *JavaScriptExport,
	mod *Module,
	pkg *Package,
) *RenderableClassMethod {
	children := make([]Renderable, 0)
	// TODO: params
	return &RenderableClassMethod{
		Method:           method,
		ClassDeclaration: cd,
		JavaScriptExport: ce,
		JavaScriptModule: mod,
		Package:          pkg,
		ChildNodes:       children,
	}
}
