/*
Copyright © 2026 Benny Powers

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
package generate

import (
	"strings"
	"unicode"
	"unicode/utf8"

	M "bennypowers.dev/cem/manifest"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// typeParameterBinding pairs a generic type parameter's name with the node of
// its declaring identifier, e.g. the `Value` in `<Value extends string =
// string>`. The CEM schema cannot represent type parameter lists, so the
// parameter and return types that reference them instead carry a Type.source
// pointer to the parameter's declaration.
type typeParameterBinding struct {
	name     string
	nameNode *ts.Node
}

// typeParameterBindings reads a declaration's generic type parameter names in
// declaration order. The grammar wraps the `extends` keyword inside the
// constraint node and the `=` inside the default_type node (whose field is
// named `value`); only the names are needed to bind type texts to their
// declaring parameter.
func typeParameterBindings(node *ts.Node, code []byte) []typeParameterBinding {
	if node == nil || node.GrammarName() != "type_parameters" {
		return nil
	}
	var bindings []typeParameterBinding
	for i := range node.ChildCount() {
		child := node.Child(i)
		if child.GrammarName() != "type_parameter" {
			continue
		}
		name := child.ChildByFieldName("name")
		if name == nil {
			continue
		}
		bindings = append(bindings, typeParameterBinding{
			name:     name.Utf8Text(code),
			nameNode: name,
		})
	}
	return bindings
}

// bindingForTypeText returns the first binding in declaration order whose
// parameter name appears in the type text as a whole word, or nil. The match
// is best-effort and lexical: cem does not typecheck, so a same-named import
// or local type can also match.
func bindingForTypeText(bindings []typeParameterBinding, typeText string) *typeParameterBinding {
	for i := range bindings {
		if containsWord(typeText, bindings[i].name) {
			return &bindings[i]
		}
	}
	return nil
}

// annotateTypeParameterSources points each type's Source at the declaration
// of the generic type parameter it references, so consumers can resolve
// identifiers like the `Value` of `<Value extends string = string>` without
// typechecking. Types that reference none of the bindings, and declarations
// without generics, are left untouched. Errors are swallowed: the
// declaration-level source reference reports the same configuration problems.
func (mp *ModuleProcessor) annotateTypeParameterSources(bindings []typeParameterBinding, types ...*M.Type) {
	if len(bindings) == 0 {
		return
	}
	for _, typ := range types {
		if typ == nil || typ.Source != nil || typ.Text == "" {
			continue
		}
		binding := bindingForTypeText(bindings, typ.Text)
		if binding == nil {
			continue
		}
		sourceRef, err := mp.generateTypeParameterSourceReference(binding.nameNode)
		if err == nil && sourceRef != nil {
			typ.Source = sourceRef
		}
	}
}

// containsWord reports whether name appears in text delimited by non-word
// runes on both sides, so `Value` matches inside `Value[]` or
// `Record<Key, Value>` but not inside `ValueX` or `MyValue`.
func containsWord(text, name string) bool {
	if name == "" {
		return false
	}
	for offset := 0; ; {
		i := strings.Index(text[offset:], name)
		if i < 0 {
			return false
		}
		start, end := offset+i, offset+i+len(name)
		before, _ := utf8.DecodeLastRuneInString(text[:start])
		after, _ := utf8.DecodeRuneInString(text[end:])
		if !isWordRune(before) && !isWordRune(after) {
			return true
		}
		offset = end
	}
}

// isWordRune reports whether r can continue a TypeScript identifier.
func isWordRune(r rune) bool {
	return r == '_' || r == '$' || unicode.IsLetter(r) || unicode.IsDigit(r)
}
