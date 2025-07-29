/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

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
package validate

import (
	"fmt"
	"strings"
)

// ManifestNavigator provides type-safe navigation through manifest data
type ManifestNavigator struct {
	manifest RawManifest
	modules  []RawModule
}

// NewManifestNavigator creates a new navigator for the given manifest
func NewManifestNavigator(manifest map[string]any) *ManifestNavigator {
	nav := &ManifestNavigator{
		manifest: RawManifest(manifest),
	}

	if modules, ok := nav.manifest.GetModules(); ok {
		nav.modules = modules
	}

	return nav
}

// GetModule returns the module at the given index
func (n *ManifestNavigator) GetModule(index int) (RawModule, bool) {
	if index < 0 || index >= len(n.modules) {
		return nil, false
	}
	return n.modules[index], true
}

// GetDeclaration returns the declaration at the given module and declaration indices
func (n *ManifestNavigator) GetDeclaration(moduleIndex, declIndex int) (RawDeclaration, bool) {
	module, ok := n.GetModule(moduleIndex)
	if !ok {
		return nil, false
	}

	declarations, ok := module.GetDeclarations()
	if !ok || declIndex < 0 || declIndex >= len(declarations) {
		return nil, false
	}

	return declarations[declIndex], true
}

// GetMember returns the member at the given indices
func (n *ManifestNavigator) GetMember(moduleIndex, declIndex, memberIndex int) (RawMember, bool) {
	decl, ok := n.GetDeclaration(moduleIndex, declIndex)
	if !ok {
		return nil, false
	}

	members, ok := decl.GetMembers()
	if !ok || memberIndex < 0 || memberIndex >= len(members) {
		return nil, false
	}

	return members[memberIndex], true
}

// GetProperty returns the property at the given indices for the specified property type
func (n *ManifestNavigator) GetProperty(moduleIndex, declIndex, propIndex int, propType string) (RawProperty, bool) {
	decl, ok := n.GetDeclaration(moduleIndex, declIndex)
	if !ok {
		return nil, false
	}

	properties, ok := decl.getPropertyArray(propType)
	if !ok || propIndex < 0 || propIndex >= len(properties) {
		return nil, false
	}

	return properties[propIndex], true
}

// BuildContextualNames builds human-readable names for validation issues
func (n *ManifestNavigator) BuildContextualNames(ctx Context) (module, declaration, member, property string) {
	if ctx.ModuleIndex >= 0 {
		if mod, ok := n.GetModule(ctx.ModuleIndex); ok {
			module = mod.GetPath()
		}
	}

	if ctx.DeclIndex >= 0 {
		if decl, ok := n.GetDeclaration(ctx.ModuleIndex, ctx.DeclIndex); ok {
			declaration = formatDeclaration(decl, ctx.DeclIndex)
		}
	}

	if ctx.MemberIndex >= 0 {
		if mem, ok := n.GetMember(ctx.ModuleIndex, ctx.DeclIndex, ctx.MemberIndex); ok {
			member = formatMember(mem, ctx.MemberIndex)
		}
	}

	if ctx.PropertyIndex >= 0 && ctx.PropertyType != "" {
		if prop, ok := n.GetProperty(ctx.ModuleIndex, ctx.DeclIndex, ctx.PropertyIndex, ctx.PropertyType); ok {
			property = formatProperty(prop, ctx.PropertyType, ctx.PropertyIndex)
		}
	}

	return
}

func formatDeclaration(decl RawDeclaration, index int) string {
	name := decl.GetName()
	kind := decl.GetKind()

	if name != "" {
		if kind != "" {
			return fmt.Sprintf("%s %s", kind, name)
		}
		return name
	}

	if kind != "" {
		return fmt.Sprintf("%s[%d]", kind, index)
	}

	return fmt.Sprintf("declaration[%d]", index)
}

func formatMember(member RawMember, index int) string {
	name := member.GetName()
	kind := member.GetKind()

	if name != "" {
		if kind != "" {
			return fmt.Sprintf("%s %s", kind, name)
		}
		return name
	}

	if kind != "" {
		return fmt.Sprintf("%s[%d]", kind, index)
	}

	return fmt.Sprintf("member[%d]", index)
}

func formatProperty(prop RawProperty, propType string, index int) string {
	name := prop.GetName()
	singularName := getSingularPropertyName(propType)

	if name != "" {
		return fmt.Sprintf("%s %s", singularName, name)
	}

	return fmt.Sprintf("%s[%d]", singularName, index)
}

func getSingularPropertyName(propType string) string {
	switch propType {
	case "cssProperties":
		return "CSS property"
	case "cssParts":
		return "CSS part"
	case "cssStates":
		return "CSS state"
	default:
		return strings.TrimSuffix(propType, "s")
	}
}
