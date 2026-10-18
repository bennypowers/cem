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
	"slices"

	M "bennypowers.dev/cem/manifest"
)

// dropOverloadImplementations removes the implementation signature of
// TypeScript overload sets from items. When several same-key items carry no
// body and exactly one carries a body, only the body-less overload
// signatures are callable by TypeScript consumers, so the body-carrying
// implementation is dropped to keep the manifest honest about the public
// API. The body-carrying item is identified position-independently: query
// match order is not source order, so slice order cannot be relied on.
// Same-key items that do not follow the overload shape, and lone body-less
// items (a stray overload signature with no implementation), are kept.
// mergeDocs copies documentation from the dropped implementation onto the
// first kept item, since authors commonly document an overload set only on
// its implementation. Items for which key returns ok=false never group.
func dropOverloadImplementations[T any](
	items []T,
	key func(T) (string, bool),
	hasBody func(T) bool,
	mergeDocs func(kept, impl T),
) []T {
	// indices into items, grouped by key; each group holds slice order
	groups := make(map[string][]int)
	for i, item := range items {
		if k, ok := key(item); ok {
			groups[k] = append(groups[k], i)
		}
	}
	drop := make([]bool, len(items))
	for _, indices := range groups {
		if len(indices) < 2 {
			continue
		}
		// exactly one body-carrying item among otherwise body-less same-key
		// items marks the overload implementation
		impl := -1
		for _, i := range indices {
			if !hasBody(items[i]) {
				continue
			}
			if impl >= 0 {
				impl = -1
				break
			}
			impl = i
		}
		if impl < 0 {
			continue
		}
		drop[impl] = true
		for _, i := range indices {
			if i != impl {
				mergeDocs(items[i], items[impl])
				break
			}
		}
	}
	if !slices.Contains(drop, true) {
		return items
	}
	kept := make([]T, 0, len(items)-1)
	for i, item := range items {
		if !drop[i] {
			kept = append(kept, item)
		}
	}
	return kept
}

// dropOverloadImplementationSignatures removes the implementation signature
// of TypeScript function overloads from a module's declarations, keeping
// only the callable overload signatures. Ambient `declare function`
// declarations are excluded earlier, at query-capture time. hasBody reports
// which function declarations carry a body.
func dropOverloadImplementationSignatures(
	declarations []M.Declaration,
	hasBody map[*M.FunctionDeclaration]bool,
) []M.Declaration {
	return dropOverloadImplementations(
		declarations,
		func(d M.Declaration) (string, bool) {
			fn, ok := d.(*M.FunctionDeclaration)
			if !ok {
				return "", false
			}
			return fn.Name(), true
		},
		func(d M.Declaration) bool {
			return hasBody[d.(*M.FunctionDeclaration)]
		},
		func(kept, impl M.Declaration) {
			mergeOverloadImplementationDocs(kept, impl)
		},
	)
}

// mergeOverloadImplementationDocs copies JSDoc-derived metadata from a
// dropped function overload implementation signature onto the first kept
// same-name overload signature.
func mergeOverloadImplementationDocs(kept, impl M.Declaration) {
	keptFn := kept.(*M.FunctionDeclaration)
	implFn := impl.(*M.FunctionDeclaration)
	if keptFn.Description == "" {
		keptFn.Description = implFn.Description
	}
	if keptFn.Summary == "" {
		keptFn.Summary = implFn.Summary
	}
	if keptFn.Deprecated == nil && implFn.Deprecated != nil {
		keptFn.Deprecated = implFn.Deprecated
	}
}

// dropOverloadImplementationMethods removes the implementation signature of
// TypeScript method overloads from a class's members, keeping only the
// callable overload signatures. Static and instance methods of the same
// name are distinct members and never form one overload set. Non-method
// members (fields, accessors) never group. hasBody reports which methods
// carry a body.
func dropOverloadImplementationMethods(
	members []M.ClassMember,
	hasBody map[*M.ClassMethod]bool,
) []M.ClassMember {
	return dropOverloadImplementations(
		members,
		func(m M.ClassMember) (string, bool) {
			method, ok := m.(*M.ClassMethod)
			if !ok {
				return "", false
			}
			return method.Name + "." + isStaticToTypeFlag(method.Static), true
		},
		func(m M.ClassMember) bool {
			return hasBody[m.(*M.ClassMethod)]
		},
		func(kept, impl M.ClassMember) {
			mergeOverloadMethodDocs(kept, impl)
		},
	)
}

// mergeOverloadMethodDocs copies JSDoc-derived metadata from a dropped
// method overload implementation signature onto the first kept same-name
// overload signature.
func mergeOverloadMethodDocs(kept, impl M.ClassMember) {
	keptMethod := kept.(*M.ClassMethod)
	implMethod := impl.(*M.ClassMethod)
	if keptMethod.Description == "" {
		keptMethod.Description = implMethod.Description
	}
	if keptMethod.Summary == "" {
		keptMethod.Summary = implMethod.Summary
	}
	if keptMethod.Deprecated == nil && implMethod.Deprecated != nil {
		keptMethod.Deprecated = implMethod.Deprecated
	}
}
