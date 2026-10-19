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

// methodGroupKey groups a declaration's public methods: static and
// instance methods of the same name are distinct members, so both name and
// staticity form the key.
type methodGroupKey struct {
	name   string
	static bool
}

// publicMethods groups a declaration's public methods by name and
// staticity. Manifests carry one member per overload signature, so a group
// maps to the set of its overloads; a group counts as present when any
// overload remains.
func publicMethods(ced *M.CustomElementDeclaration) map[methodGroupKey][]*M.ClassMethod {
	methods := make(map[methodGroupKey][]*M.ClassMethod)
	for _, member := range ced.Members {
		if m, ok := member.(*M.ClassMethod); ok && isPublic(m.Privacy) {
			key := methodGroupKey{name: m.Name, static: m.Static}
			methods[key] = append(methods[key], m)
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

// paramSignature renders a method's parameter list for overload pairing and
// change reporting. Rest and optional markers are part of the signature, so
// an optional-to-required or rest-to-non-rest transition does not pair as an
// unchanged signature.
func paramSignature(m *M.ClassMethod) string {
	var params []string
	for _, p := range m.Parameters {
		param := typeText(p.Type)
		if p.Rest {
			param = "..." + param
		}
		if p.Optional {
			param += "?"
		}
		params = append(params, param)
	}
	return "(" + strings.Join(params, ", ") + ")"
}

// overloadPair couples a base overload with the head overload it compares
// against.
type overloadPair struct {
	base *M.ClassMethod
	head *M.ClassMethod
}

// overloadPairs pairs each base overload with the head overload it compares
// against for the parameter and return-type rules. A method with a single
// overload on each side pairs directly, preserving the classic
// single-method comparison; overload sets pair per overload by parameter
// signature, each head overload pairing with at most one base overload, and
// base overloads without an unused matching head entry come back as
// unmatched: that call shape is gone from the set, which the caller
// reports.
func overloadPairs(base, head []*M.ClassMethod) (pairs []overloadPair, unmatched []*M.ClassMethod) {
	if len(base) == 1 && len(head) == 1 {
		return []overloadPair{{base[0], head[0]}}, nil
	}
	used := make([]bool, len(head))
	for _, b := range base {
		baseSig := paramSignature(b)
		match := -1
		for i, h := range head {
			if !used[i] && paramSignature(h) == baseSig {
				match = i
				break
			}
		}
		if match < 0 {
			unmatched = append(unmatched, b)
			continue
		}
		used[match] = true
		pairs = append(pairs, overloadPair{base: b, head: head[match]})
	}
	return pairs, unmatched
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
		for key := range publicMethods(baseEl) {
			if _, ok := headMethods[key]; !ok {
				changes = append(changes, Change{
					Rule:     "method-removed",
					Severity: Breaking,
					Element:  tag,
					Subject:  key.name,
					Message:  fmt.Sprintf("method %q removed from <%s>", key.name, tag),
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
		for key := range publicMethods(headEl) {
			if _, ok := baseMethods[key]; !ok {
				changes = append(changes, Change{
					Rule:     "method-added",
					Severity: Safe,
					Element:  tag,
					Subject:  key.name,
					Message:  fmt.Sprintf("method %q added to <%s>", key.name, tag),
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
		for key, baseOverloads := range publicMethods(baseEl) {
			headOverloads, ok := headMethods[key]
			if !ok {
				continue
			}
			// the return type is compared on the overload paired by
			// overloadPairs; unmatched base overloads carry no head
			// counterpart to compare against and are reported by
			// method-parameter-changed instead
			pairs, _ := overloadPairs(baseOverloads, headOverloads)
			for _, pair := range pairs {
				baseReturn := returnTypeText(pair.base)
				headReturn := returnTypeText(pair.head)
				if baseReturn != "" && headReturn != "" && baseReturn != headReturn {
					changes = append(changes, Change{
						Rule:     "method-return-type-changed",
						Severity: Breaking,
						Element:  tag,
						Subject:  key.name,
						Message:  fmt.Sprintf("method %q return type changed from %s to %s in <%s>", key.name, baseReturn, headReturn, tag),
					})
				}
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
		for key, baseOverloads := range publicMethods(baseEl) {
			headOverloads, ok := headMethods[key]
			if !ok {
				continue
			}
			pairs, unmatched := overloadPairs(baseOverloads, headOverloads)
			for _, pair := range pairs {
				// only the direct single-overload pairing can differ;
				// overload-set pairs are matched by parameter signature
				baseSig := paramSignature(pair.base)
				headSig := paramSignature(pair.head)
				if baseSig == headSig {
					continue
				}
				changes = append(changes, Change{
					Rule:     "method-parameter-changed",
					Severity: Breaking,
					Element:  tag,
					Subject:  key.name,
					Message:  fmt.Sprintf("method %q parameters changed from %s to %s in <%s>", key.name, baseSig, headSig, tag),
				})
			}
			for _, baseMethod := range unmatched {
				changes = append(changes, Change{
					Rule:     "method-parameter-changed",
					Severity: Breaking,
					Element:  tag,
					Subject:  key.name,
					Message:  fmt.Sprintf("method %q overload parameters %s removed from <%s>", key.name, paramSignature(baseMethod), tag),
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

type fieldAddedRule struct{}

func (*fieldAddedRule) ID() string { return "field-added" }

func (*fieldAddedRule) Check(base, head map[string]*M.CustomElementDeclaration) []Change {
	var changes []Change
	for tag, headEl := range head {
		baseEl, ok := base[tag]
		if !ok {
			continue
		}
		baseFields := publicFields(baseEl)
		for name := range publicFields(headEl) {
			if _, ok := baseFields[name]; !ok {
				changes = append(changes, Change{
					Rule:     "field-added",
					Severity: Safe,
					Element:  tag,
					Subject:  name,
					Message:  fmt.Sprintf("field %q added to <%s>", name, tag),
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
