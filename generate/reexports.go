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
	"regexp"
	"slices"
	"strings"

	M "bennypowers.dev/cem/manifest"
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

// typeSpecifierRe matches per-specifier type annotations where ALL specifiers are type-only.
// Matches: { type Foo }, { type Foo, type Bar }, { type Foo as Bar }, { type Foo as Bar, type Baz }
var typeSpecifierRe = regexp.MustCompile(`\{\s*(type\s+\w+(\s+as\s+\w+)?\s*,?\s*)+\}`)

// statementTextBefore extracts the statement text from the beginning of the line
// up to the given byte offset. Used to check for type-only import/export keywords
// that appear before the source specifier in the statement.
func statementTextBefore(code []byte, endByte uint) string {
	if int(endByte) > len(code) {
		endByte = uint(len(code))
	}
	start := int(endByte)
	for start > 0 && code[start-1] != '\n' {
		start--
	}
	return string(code[start:endByte])
}

// isTypeOnlyImport checks if an import statement is type-only and will be elided.
// Matches: import type { X }, import type * as Ns, import { type X }
// Does NOT match: import { X }, import './x.js', import { X, type Y } (has non-type specifier)
func isTypeOnlyImport(stmtText string) bool {
	trimmed := strings.TrimSpace(stmtText)
	// Statement-level: import type { X } or import type * as Ns
	if strings.HasPrefix(trimmed, "import type ") {
		return true
	}
	// Per-specifier: import { type X, type Y } (all specifiers are type-only)
	if typeSpecifierRe.MatchString(trimmed) {
		return true
	}
	return false
}

// isTypeOnlyExport checks if an export statement is type-only and will be elided.
// Matches: export type { X }, export { type X }
// Does NOT match: export { X }, export { X, type Y } (has non-type specifier)
func isTypeOnlyExport(stmtText string) bool {
	trimmed := strings.TrimSpace(stmtText)
	// Statement-level: export type { X }
	if strings.HasPrefix(trimmed, "export type ") {
		return true
	}
	// Per-specifier: export { type X, type Y } (all specifiers are type-only)
	if typeSpecifierRe.MatchString(trimmed) {
		return true
	}
	return false
}
