/*
Copyright © 2026 Benny Powers <web@bennypowers.com>

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
package breaking

import (
	"fmt"
	"strings"

	M "bennypowers.dev/cem/manifest"
)

func publicMethods(ced *M.CustomElementDeclaration) map[string]*M.ClassMethod {
	methods := make(map[string]*M.ClassMethod)
	for _, member := range ced.Members {
		if m, ok := member.(*M.ClassMethod); ok && isPublic(m.Privacy) {
			methods[m.Name] = m
		}
	}
	return methods
}

func publicFields(ced *M.CustomElementDeclaration) map[string]*M.ClassField {
	fields := make(map[string]*M.ClassField)
	for _, member := range ced.Members {
		switch f := member.(type) {
		case *M.ClassField:
			if isPublic(f.Privacy) {
				fields[f.Name] = f
			}
		case *M.CustomElementField:
			if isPublic(f.Privacy) {
				fields[f.Name] = &f.ClassField
			}
		}
	}
	return fields
}

func paramSignature(m *M.ClassMethod) string {
	var params []string
	for _, p := range m.Parameters {
		s := p.Name
		if p.Type != nil && p.Type.Text != "" {
			s += ": " + p.Type.Text
		}
		params = append(params, s)
	}
	return "(" + strings.Join(params, ", ") + ")"
}

// Methods

type methodRemovedRule struct{}

func (*methodRemovedRule) ID() string { return "method-removed" }

func (*methodRemovedRule) Check(base, head map[string]*M.CustomElementDeclaration) []Change {
	var changes []Change
	for tag, baseEl := range base {
		headEl, ok := head[tag]
		if !ok {
			continue
		}
		headMethods := publicMethods(headEl)
		for name := range publicMethods(baseEl) {
			if _, ok := headMethods[name]; !ok {
				changes = append(changes, Change{
					Rule:     "method-removed",
					Severity: Breaking,
					Element:  tag,
					Subject:  name,
					Message:  fmt.Sprintf("method %q removed from <%s>", name, tag),
				})
			}
		}
	}
	return changes
}

type methodAddedRule struct{}

func (*methodAddedRule) ID() string { return "method-added" }

func (*methodAddedRule) Check(base, head map[string]*M.CustomElementDeclaration) []Change {
	var changes []Change
	for tag, headEl := range head {
		baseEl, ok := base[tag]
		if !ok {
			continue
		}
		baseMethods := publicMethods(baseEl)
		for name := range publicMethods(headEl) {
			if _, ok := baseMethods[name]; !ok {
				changes = append(changes, Change{
					Rule:     "method-added",
					Severity: Safe,
					Element:  tag,
					Subject:  name,
					Message:  fmt.Sprintf("method %q added to <%s>", name, tag),
				})
			}
		}
	}
	return changes
}

type methodReturnTypeChangedRule struct{}

func (*methodReturnTypeChangedRule) ID() string { return "method-return-type-changed" }

func (*methodReturnTypeChangedRule) Check(base, head map[string]*M.CustomElementDeclaration) []Change {
	var changes []Change
	for tag, baseEl := range base {
		headEl, ok := head[tag]
		if !ok {
			continue
		}
		headMethods := publicMethods(headEl)
		for name, baseMethod := range publicMethods(baseEl) {
			headMethod, ok := headMethods[name]
			if !ok {
				continue
			}
			baseReturn := ""
			headReturn := ""
			if baseMethod.Return != nil && baseMethod.Return.Type != nil {
				baseReturn = baseMethod.Return.Type.Text
			}
			if headMethod.Return != nil && headMethod.Return.Type != nil {
				headReturn = headMethod.Return.Type.Text
			}
			if baseReturn != "" && headReturn != "" && baseReturn != headReturn {
				changes = append(changes, Change{
					Rule:     "method-return-type-changed",
					Severity: Breaking,
					Element:  tag,
					Subject:  name,
					Message:  fmt.Sprintf("method %q return type changed from %s to %s in <%s>", name, baseReturn, headReturn, tag),
				})
			}
		}
	}
	return changes
}

type methodParameterChangedRule struct{}

func (*methodParameterChangedRule) ID() string { return "method-parameter-changed" }

func (*methodParameterChangedRule) Check(base, head map[string]*M.CustomElementDeclaration) []Change {
	var changes []Change
	for tag, baseEl := range base {
		headEl, ok := head[tag]
		if !ok {
			continue
		}
		headMethods := publicMethods(headEl)
		for name, baseMethod := range publicMethods(baseEl) {
			headMethod, ok := headMethods[name]
			if !ok {
				continue
			}
			baseSig := paramSignature(baseMethod)
			headSig := paramSignature(headMethod)
			if baseSig != headSig {
				changes = append(changes, Change{
					Rule:     "method-parameter-changed",
					Severity: Breaking,
					Element:  tag,
					Subject:  name,
					Message:  fmt.Sprintf("method %q parameters changed from %s to %s in <%s>", name, baseSig, headSig, tag),
				})
			}
		}
	}
	return changes
}

// Fields

type fieldRemovedRule struct{}

func (*fieldRemovedRule) ID() string { return "field-removed" }

func (*fieldRemovedRule) Check(base, head map[string]*M.CustomElementDeclaration) []Change {
	var changes []Change
	for tag, baseEl := range base {
		headEl, ok := head[tag]
		if !ok {
			continue
		}
		headFields := publicFields(headEl)
		for name := range publicFields(baseEl) {
			if _, ok := headFields[name]; !ok {
				changes = append(changes, Change{
					Rule:     "field-removed",
					Severity: Breaking,
					Element:  tag,
					Subject:  name,
					Message:  fmt.Sprintf("field %q removed from <%s>", name, tag),
				})
			}
		}
	}
	return changes
}

type fieldTypeChangedRule struct{}

func (*fieldTypeChangedRule) ID() string { return "field-type-changed" }

func (*fieldTypeChangedRule) Check(base, head map[string]*M.CustomElementDeclaration) []Change {
	var changes []Change
	for tag, baseEl := range base {
		headEl, ok := head[tag]
		if !ok {
			continue
		}
		headFields := publicFields(headEl)
		for name, baseField := range publicFields(baseEl) {
			headField, ok := headFields[name]
			if !ok {
				continue
			}
			baseType := typeText(baseField.Type)
			headType := typeText(headField.Type)
			if baseType != "" && headType != "" && baseType != headType {
				changes = append(changes, Change{
					Rule:     "field-type-changed",
					Severity: Dangerous,
					Element:  tag,
					Subject:  name,
					Message:  fmt.Sprintf("field %q type changed from %s to %s in <%s>", name, baseType, headType, tag),
				})
			}
		}
	}
	return changes
}
