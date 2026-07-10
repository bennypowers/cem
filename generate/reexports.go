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
func resolveReExportedCEDefinitions(pkg *M.Package) {
	// Build lookup maps from all modules
	moduleByPath := make(map[string]*M.Module, len(pkg.Modules))
	for i := range pkg.Modules {
		moduleByPath[pkg.Modules[i].Path] = &pkg.Modules[i]
	}

	for i := range pkg.Modules {
		mod := &pkg.Modules[i]

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
				addCEExportForReExportedClass(mod, jsExport.Declaration.Name, jsExport.Declaration.Module, moduleByPath, seen)
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

		// Re-sort exports after adding new ones
		if len(mod.Exports) > originalExportCount {
			slices.SortStableFunc(mod.Exports, func(a M.Export, b M.Export) int {
				return cmp.Compare(a.GetStartByte(), b.GetStartByte())
			})
		}
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

// isTypeOnlyStatement checks if an import/export tree-sitter node is type-only
// by looking for a `type` anonymous child (statement-level) or checking that
// all specifiers have `type` children (per-specifier).
func isTypeOnlyStatement(root *ts.Node, nodeId int) bool {
	node := Q.GetDescendantById(root, nodeId)
	if node == nil {
		return false
	}
	return hasStatementLevelType(node) || allSpecifiersTypeOnly(node)
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

// allSpecifiersTypeOnly checks if every import_specifier/export_specifier
// has a `type` child, meaning all specifiers are type-only.
func allSpecifiersTypeOnly(node *ts.Node) bool {
	cursor := node.Walk()
	defer cursor.Close()
	var specifiers int
	var typeSpecifiers int
	for _, child := range node.NamedChildren(cursor) {
		walkSpecifiers(&child, &specifiers, &typeSpecifiers)
	}
	return specifiers > 0 && specifiers == typeSpecifiers
}

func walkSpecifiers(node *ts.Node, total, typed *int) {
	name := node.GrammarName()
	if name == "import_specifier" || name == "export_specifier" {
		*total++
		count := node.ChildCount()
		for i := uint(0); i < count; i++ {
			if node.Child(i).GrammarName() == "type" {
				*typed++
				break
			}
		}
		return
	}
	cursor := node.Walk()
	defer cursor.Close()
	for _, child := range node.NamedChildren(cursor) {
		walkSpecifiers(&child, total, typed)
	}
}
