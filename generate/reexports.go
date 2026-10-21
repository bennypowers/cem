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
package generate

import (
	"cmp"
	"slices"

	Q "bennypowers.dev/cem/internal/treesitter"
	M "bennypowers.dev/cem/manifest"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// resolveReExportedCEDefinitions adds custom-element-definition exports to modules
// that re-export custom element classes from other modules.
//
// This handles three cases:
//  1. Named re-exports: export { ChildElement } from './child.js'
//     → adds CE export for child-element if ChildElement is a CE in child.js
//  2. Namespace re-exports: export * from './child.js'
//     → copies all JS and CE exports from child.js
//  3. Side-effect imports: import './child.js'
//     → adds CE exports for all CEs defined in child.js
//
// Star re-exports chain (a star-exports b, b star-exports c), so a single pass
// over modules expands only one link per module whenever a re-exporter sorts
// before its source, making the expanded set depend on module order. The
// expansion therefore runs as a reverse-dependency worklist: every module is
// seeded once, and a module whose exports grew re-enqueues exactly the modules
// that read them, so each chain link is expanded where it is discovered
// instead of restarting a whole-package sweep per link. This computes the same
// least fixpoint as repeating full sweeps: each expansion copies its sources'
// exports as they currently stand, exports only ever grow and are
// deduplicated by key, and the queue drains (even for cyclic star re-exports)
// once no module can absorb anything more.
func resolveReExportedCEDefinitions(pkg *M.Package) {
	// Build lookup maps from all modules
	moduleByPath := make(map[string]*M.Module, len(pkg.Modules))
	for i := range pkg.Modules {
		moduleByPath[pkg.Modules[i].Path] = &pkg.Modules[i]
	}

	dependents := dependentsBySource(pkg.Modules)

	// Seed the worklist with every module in package order. A module whose
	// exports grew re-enqueues its dependents, deduplicated: a module already
	// queued is not re-added and absorbs all pending growth in one expansion.
	queue := make([]*M.Module, 0, len(pkg.Modules))
	queued := make(map[*M.Module]bool, len(pkg.Modules))
	for i := range pkg.Modules {
		mod := &pkg.Modules[i]
		queue = append(queue, mod)
		queued[mod] = true
	}
	for len(queue) > 0 {
		mod := queue[0]
		queue = queue[1:]
		delete(queued, mod)
		if !expandModuleReExports(mod, moduleByPath) {
			continue
		}
		for _, dependent := range dependents[mod.Path] {
			if !queued[dependent] {
				queued[dependent] = true
				queue = append(queue, dependent)
			}
		}
	}
}

// dependentsBySource maps each module path to the modules that read its
// exports and may therefore absorb growth: modules listing it in
// ReExportAllSources (star re-exporters) or SideEffectImports (side-effect
// importers), and modules holding a JS export whose declaration names it as
// the source module (named and namespace-alias re-exporters, whose expansion
// reads the source's exports). The map is built once, before any expansion.
func dependentsBySource(modules []M.Module) map[string][]*M.Module {
	dependents := make(map[string][]*M.Module, len(modules))
	add := func(sourcePath string, mod *M.Module) {
		ds := dependents[sourcePath]
		if slices.Contains(ds, mod) {
			return
		}
		dependents[sourcePath] = append(ds, mod)
	}
	for i := range modules {
		mod := &modules[i]
		for _, sourcePath := range mod.ReExportAllSources {
			add(sourcePath, mod)
		}
		for _, sourcePath := range mod.SideEffectImports {
			add(sourcePath, mod)
		}
		for _, export := range mod.Exports {
			jsExport, ok := export.(*M.JavaScriptExport)
			if !ok || jsExport.Declaration == nil || jsExport.Declaration.Module == mod.Path {
				continue
			}
			add(jsExport.Declaration.Module, mod)
		}
	}
	return dependents
}

// expandModuleReExports runs one expansion pass over a single module and
// reports whether it appended any exports.
func expandModuleReExports(mod *M.Module, moduleByPath map[string]*M.Module) bool {
	// Track original export count to know if we added anything
	originalExportCount := len(mod.Exports)

	// Build set of existing export keys to avoid duplicates
	seen := exportKeySet(mod.Exports)

	// Case 1: Named/namespace-alias re-exports - check existing JS exports that reference other modules
	// Take a snapshot of current exports length to avoid iterating over newly added exports
	existingExports := originalExportCount
	for j := range existingExports {
		jsExport, ok := mod.Exports[j].(*M.JavaScriptExport)
		if !ok || jsExport.Declaration == nil {
			continue
		}
		// Skip exports that reference the same module (not a re-export)
		if jsExport.Declaration.Module == mod.Path {
			continue
		}
		if jsExport.Declaration.Name == "*" {
			// Namespace alias: export * as Ns from './module.js' - add all CE exports
			addAllCEExports(mod, jsExport.Declaration.Module, moduleByPath, seen)
		} else {
			className := jsExport.Declaration.Name
			sourcePath := jsExport.Declaration.Module
			// `export { default } from './x.js'` names its declaration
			// "default", which never matches a CE class; resolve through
			// x's own default export to find the underlying class. Chained
			// default re-exports (x itself re-exports y's default) carry
			// the same "default" name, so follow the chain until a real
			// declaration name appears, guarding against re-export cycles.
			visited := make(map[string]bool)
			for className == "default" {
				if visited[sourcePath] {
					break
				}
				visited[sourcePath] = true
				sourceMod, ok := moduleByPath[sourcePath]
				if !ok {
					break
				}
				next := ""
				for _, export := range sourceMod.Exports {
					if js, ok := export.(*M.JavaScriptExport); ok && js.Name == "default" && js.Declaration != nil {
						next = js.Declaration.Name
						sourcePath = js.Declaration.Module
						break
					}
				}
				if next == "" {
					break
				}
				className = next
			}
			addCEExportForReExportedClass(mod, className, sourcePath, moduleByPath, seen)
		}
	}

	// Case 2: Namespace re-exports - export * from './module.js'
	for _, sourcePath := range mod.ReExportAllSources {
		sourceMod, ok := moduleByPath[sourcePath]
		if !ok {
			continue
		}
		for _, export := range sourceMod.Exports {
			switch e := export.(type) {
			case *M.JavaScriptExport:
				if e.Declaration == nil {
					continue
				}
				// ES export * does not re-export the default binding; only an
				// explicit `export { default } from` (Case 1) surfaces it
				if e.Name == "default" {
					continue
				}
				key := exportKey{e.Kind, e.Name}
				if seen[key] {
					continue
				}
				ref := M.NewReference(e.Declaration.Name, e.Declaration.Package, e.Declaration.Module)
				mod.Exports = append(mod.Exports, &M.JavaScriptExport{
					Kind:        "js",
					Name:        e.Name,
					Declaration: ref,
				})
				seen[key] = true
			case *M.CustomElementExport:
				if e.Declaration == nil {
					continue
				}
				key := exportKey{e.Kind, e.Name}
				if seen[key] {
					continue
				}
				ref := M.NewReference(e.Declaration.Name, e.Declaration.Package, e.Declaration.Module)
				mod.Exports = append(mod.Exports, M.NewCustomElementExport(
					e.Name,
					ref,
					0,
					nil,
				))
				seen[key] = true
			}
		}
	}

	// Case 3: Side-effect imports - import './module.js'
	for _, sourcePath := range mod.SideEffectImports {
		sourceMod, ok := moduleByPath[sourcePath]
		if !ok {
			continue
		}
		for _, export := range sourceMod.Exports {
			if ceExport, ok := export.(*M.CustomElementExport); ok && ceExport.Declaration != nil {
				key := exportKey{ceExport.Kind, ceExport.Name}
				if seen[key] {
					continue
				}
				ref := M.NewReference(ceExport.Declaration.Name, ceExport.Declaration.Package, ceExport.Declaration.Module)
				mod.Exports = append(mod.Exports, M.NewCustomElementExport(
					ceExport.Name,
					ref,
					0,
					nil,
				))
				seen[key] = true
			}
		}
	}

	// Re-sort exports after adding new ones. Star-copied exports carry
	// StartByte 0, tying with each other (and with a module's own
	// zero-position exports), so kind then name breaks the tie and the
	// sequence does not depend on the order modules were expanded in.
	if len(mod.Exports) > originalExportCount {
		slices.SortStableFunc(mod.Exports, compareExports)
	}
	return len(mod.Exports) > originalExportCount
}

// compareExports orders exports by source position, tie-breaking equal
// StartByte groups (star-copied exports all carry 0) by kind then name so
// the sequence is deterministic regardless of expansion order.
func compareExports(a, b M.Export) int {
	aStart, aKind, aName := exportSortKeys(a)
	bStart, bKind, bName := exportSortKeys(b)
	if c := cmp.Compare(aStart, bStart); c != 0 {
		return c
	}
	if c := cmp.Compare(aKind, bKind); c != 0 {
		return c
	}
	return cmp.Compare(aName, bName)
}

// exportSortKeys extracts an export's position, kind, and name sort keys.
func exportSortKeys(e M.Export) (start uint, kind, name string) {
	switch ex := e.(type) {
	case *M.JavaScriptExport:
		return ex.StartByte, ex.Kind, ex.Name
	case *M.CustomElementExport:
		return ex.StartByte, ex.Kind, ex.Name
	default:
		return e.GetStartByte(), "", ""
	}
}

// exportKey uniquely identifies an export by kind and name.
type exportKey struct {
	kind string
	name string
}

// exportKeySet builds a set of export keys from existing exports.
func exportKeySet(exports []M.Export) map[exportKey]bool {
	seen := make(map[exportKey]bool, len(exports))
	for _, e := range exports {
		switch ex := e.(type) {
		case *M.JavaScriptExport:
			seen[exportKey{ex.Kind, ex.Name}] = true
		case *M.CustomElementExport:
			seen[exportKey{ex.Kind, ex.Name}] = true
		}
	}
	return seen
}

// addAllCEExports copies all custom-element-definition exports from the source module,
// skipping any that already exist in the module's exports.
func addAllCEExports(mod *M.Module, sourceModulePath string, moduleByPath map[string]*M.Module, seen map[exportKey]bool) {
	sourceMod, ok := moduleByPath[sourceModulePath]
	if !ok {
		return
	}
	for _, export := range sourceMod.Exports {
		if ceExport, ok := export.(*M.CustomElementExport); ok && ceExport.Declaration != nil {
			key := exportKey{ceExport.Kind, ceExport.Name}
			if seen[key] {
				continue
			}
			ref := M.NewReference(ceExport.Declaration.Name, ceExport.Declaration.Package, ceExport.Declaration.Module)
			mod.Exports = append(mod.Exports, M.NewCustomElementExport(
				ceExport.Name,
				ref,
				0,
				nil,
			))
			seen[key] = true
		}
	}
}

// addCEExportForReExportedClass checks if a re-exported class is a custom element
// in the source module, and if so adds a custom-element-definition export.
func addCEExportForReExportedClass(mod *M.Module, className string, sourceModulePath string, moduleByPath map[string]*M.Module, seen map[exportKey]bool) {
	sourceMod, ok := moduleByPath[sourceModulePath]
	if !ok {
		return
	}
	for _, decl := range sourceMod.Declarations {
		if ce, ok := decl.(*M.CustomElementDeclaration); ok && ce.Name() == className {
			key := exportKey{"custom-element-definition", ce.TagName}
			if seen[key] {
				return
			}
			ref := M.NewReference(className, "", sourceModulePath)
			mod.Exports = append(mod.Exports, M.NewCustomElementExport(
				ce.TagName,
				ref,
				0,
				nil,
			))
			seen[key] = true
			return
		}
	}
}

// findAncestor walks up from a node to find the nearest ancestor with the given grammar name.
func findAncestor(root *ts.Node, nodeId int, grammarName string) *ts.Node {
	node := Q.GetDescendantById(root, nodeId)
	for node != nil {
		if node.GrammarName() == grammarName {
			return node
		}
		node = node.Parent()
	}
	return nil
}

// isTypeOnlyNode checks if an import/export tree-sitter node is type-only
// by looking for a `type` anonymous child (statement-level) or checking that
// all bindings are type-only (per-specifier).
func isTypeOnlyNode(node *ts.Node) bool {
	if node == nil {
		return false
	}
	return hasStatementLevelType(node) || allBindingsTypeOnly(node)
}

// isDefaultExportStatement checks whether an export_statement node is a
// default export (`export default ...`). Inline default declarations
// (`export default class C {}`, `export default function f() {}`) also match
// the declarations query's generic export pattern, so that pattern's handler
// skips them and leaves default exports to the exports query.
func isDefaultExportStatement(node *ts.Node) bool {
	if node == nil {
		return false
	}
	return hasAnonymousChild(node, "default")
}

// hasFromClause checks whether an import or export statement node carries a
// `from` clause, i.e. it imports from or re-exports another module.
func hasFromClause(node *ts.Node) bool {
	if node == nil {
		return false
	}
	return hasAnonymousChild(node, "from")
}

// hasAnonymousChild reports whether the node has a child token with the
// given grammar name, e.g. anonymous keyword tokens like `default` or `from`.
func hasAnonymousChild(node *ts.Node, name string) bool {
	count := node.ChildCount()
	for i := range count {
		if node.Child(i).GrammarName() == name {
			return true
		}
	}
	return false
}

// hasStatementLevelType checks for `import type` or `export type` at statement level.
func hasStatementLevelType(node *ts.Node) bool {
	count := node.ChildCount()
	for i := uint(0); i+1 < count; i++ {
		child := node.Child(i)
		next := node.Child(i + 1)
		if child.GrammarName() == "import" && next.GrammarName() == "type" {
			return true
		}
		if child.GrammarName() == "export" && next.GrammarName() == "type" {
			return true
		}
	}
	return false
}

// allBindingsTypeOnly checks if every binding in the import/export is type-only.
// Returns false if any default binding, namespace binding, or non-type specifier exists.
func allBindingsTypeOnly(node *ts.Node) bool {
	cursor := node.Walk()
	defer cursor.Close()
	var bindings int
	var typeBindings int
	for _, child := range node.NamedChildren(cursor) {
		classifyBindings(&child, &bindings, &typeBindings)
	}
	return bindings > 0 && bindings == typeBindings
}

func classifyBindings(node *ts.Node, total, typed *int) {
	switch node.GrammarName() {
	case "import_specifier", "export_specifier":
		*total++
		if isTypeOnlySpecifier(node) {
			*typed++
		}
	case "identifier":
		// Default binding (import Foo, { ... }) -- never type-only
		*total++
	case "namespace_import", "namespace_export":
		// import * as Ns / export * as Ns -- never type-only at per-binding level
		*total++
	case "import_clause", "named_imports", "export_clause":
		cursor := node.Walk()
		defer cursor.Close()
		for _, child := range node.NamedChildren(cursor) {
			classifyBindings(&child, total, typed)
		}
	}
}

// isTypeOnlySpecifier reports whether an import/export specifier node
// carries a per-specifier `type` modifier (e.g. `export { type Foo }` or
// `import { type Foo }`), which the compiler elides.
func isTypeOnlySpecifier(node *ts.Node) bool {
	if node == nil {
		return false
	}
	switch node.GrammarName() {
	case "import_specifier", "export_specifier":
	default:
		return false
	}
	count := node.ChildCount()
	for i := range count {
		if node.Child(i).GrammarName() == "type" {
			return true
		}
	}
	return false
}
