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
	"strconv"
	"strings"
)

// Raw types represent the untyped JSON manifest data
type RawManifest map[string]any
type RawModule map[string]any
type RawDeclaration map[string]any
type RawMember map[string]any
type RawProperty map[string]any
type RawSuperclass map[string]any

// CSS property length threshold for warnings
const CSSDefaultLengthThreshold = 200

// Helper methods for type-safe access to manifest data

func (m RawManifest) GetModules() ([]RawModule, bool) {
	modules, ok := m["modules"].([]any)
	if !ok {
		return nil, false
	}
	result := make([]RawModule, len(modules))
	for i, mod := range modules {
		if module, ok := mod.(map[string]any); ok {
			result[i] = RawModule(module)
		}
	}
	return result, true
}

func (m RawModule) GetPath() string {
	if path, ok := m["path"].(string); ok {
		return path
	}
	return ""
}

func (m RawModule) GetDeclarations() ([]RawDeclaration, bool) {
	declarations, ok := m["declarations"].([]any)
	if !ok {
		return nil, false
	}
	result := make([]RawDeclaration, len(declarations))
	for i, decl := range declarations {
		if declaration, ok := decl.(map[string]any); ok {
			result[i] = RawDeclaration(declaration)
		}
	}
	return result, true
}

func (d RawDeclaration) GetName() string {
	if name, ok := d["name"].(string); ok {
		return name
	}
	return ""
}

func (d RawDeclaration) GetKind() string {
	if kind, ok := d["kind"].(string); ok {
		return kind
	}
	return ""
}

func (d RawDeclaration) GetMembers() ([]RawMember, bool) {
	members, ok := d["members"].([]any)
	if !ok {
		return nil, false
	}
	result := make([]RawMember, len(members))
	for i, mem := range members {
		if member, ok := mem.(map[string]any); ok {
			result[i] = RawMember(member)
		}
	}
	return result, true
}

func (d RawDeclaration) GetSuperclass() (RawSuperclass, bool) {
	if superclass, ok := d["superclass"].(map[string]any); ok {
		return RawSuperclass(superclass), true
	}
	return nil, false
}

func (d RawDeclaration) GetCSSProperties() ([]RawProperty, bool) {
	return d.getPropertyArray("cssProperties")
}

func (d RawDeclaration) getPropertyArray(key string) ([]RawProperty, bool) {
	props, ok := d[key].([]any)
	if !ok {
		return nil, false
	}
	result := make([]RawProperty, len(props))
	for i, prop := range props {
		if property, ok := prop.(map[string]any); ok {
			result[i] = RawProperty(property)
		}
	}
	return result, true
}

func (m RawMember) GetName() string {
	if name, ok := m["name"].(string); ok {
		return name
	}
	return ""
}

func (m RawMember) GetKind() string {
	if kind, ok := m["kind"].(string); ok {
		return kind
	}
	return ""
}

func (m RawMember) IsStatic() bool {
	if isStatic, ok := m["static"].(bool); ok {
		return isStatic
	}
	return false
}

func (m RawMember) GetPrivacy() string {
	if privacy, ok := m["privacy"].(string); ok {
		return privacy
	}
	return ""
}

func (p RawProperty) GetName() string {
	if name, ok := p["name"].(string); ok {
		return name
	}
	return ""
}

func (p RawProperty) GetDefault() string {
	if defaultVal, ok := p["default"].(string); ok {
		return defaultVal
	}
	return ""
}

func (s RawSuperclass) GetName() string {
	if name, ok := s["name"].(string); ok {
		return name
	}
	return ""
}

func (s RawSuperclass) GetModule() string {
	if module, ok := s["module"].(string); ok {
		return module
	}
	return ""
}

// Context represents the parsing context for validation issues
type Context struct {
	Type          string // "module", "declaration", "member", "property"
	ModuleIndex   int
	DeclIndex     int
	MemberIndex   int
	PropertyIndex int
	PropertyType  string // "attributes", "events", etc.
}

// ParseContext extracts context information from a JSON path
func ParseContext(location string) Context {
	parts := splitLocation(location)
	ctx := Context{
		ModuleIndex:   -1,
		DeclIndex:     -1,
		MemberIndex:   -1,
		PropertyIndex: -1,
	}

	for i, part := range parts {
		if isIndex(part) {
			index, _ := strconv.Atoi(part)
			if i > 0 {
				parentKey := parts[i-1]
				switch parentKey {
				case "modules":
					ctx.ModuleIndex = index
					ctx.Type = "module"
				case "declarations":
					ctx.DeclIndex = index
					ctx.Type = "declaration"
				case "members":
					ctx.MemberIndex = index
					ctx.Type = "member"
				case "attributes", "events", "slots", "cssProperties", "cssParts", "cssStates":
					ctx.PropertyIndex = index
					ctx.PropertyType = parentKey
					ctx.Type = "property"
				}
			}
		}
	}

	return ctx
}

func splitLocation(location string) []string {
	if location == "" || location == "root" {
		return []string{}
	}
	return strings.Split(strings.TrimPrefix(location, "/"), "/")
}

func isIndex(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}