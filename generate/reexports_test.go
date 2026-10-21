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
	"slices"
	"testing"

	"bennypowers.dev/cem/internal/languages/typescript"
	M "bennypowers.dev/cem/manifest"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// parseStatement parses a single TS statement and returns the root node and tree.
func parseStatement(stmt string) (*ts.Node, *ts.Tree) {
	parser := typescript.BorrowParser()
	defer typescript.ReturnParser(parser)
	tree := parser.Parse([]byte(stmt), nil)
	if tree == nil {
		return nil, nil
	}
	root := tree.RootNode()
	return root, tree
}

// Inline: pure boolean classifier, table-driven cases are clearer than goldens

func TestIsTypeOnlyNode(t *testing.T) {
	tests := []struct {
		name string
		stmt string
		want bool
	}{
		{"import type named", "import type { Foo } from './foo.js';", true},
		{"import type named alias", "import type { Foo as Bar } from './foo.js';", true},
		{"import type namespace", "import type * as FooNS from './foo.js';", true},
		{"import type specifier", "import { type Foo } from './foo.js';", true},
		{"import type specifier alias", "import { type Foo as Bar } from './foo.js';", true},
		{"import type specifier multi", "import { type Foo, type Bar } from './foo.js';", true},
		{"side-effect import", "import './foo.js';", false},
		{"named import", "import { Foo } from './foo.js';", false},
		{"named import alias", "import { Foo as Bar } from './foo.js';", false},
		{"namespace import", "import * as FooNS from './foo.js';", false},
		{"mixed import", "import { Foo, type Bar } from './foo.js';", false},
		{"default plus type import", "import Foo, { type Bar } from './foo.js';", false},
		{"export type named", "export type { Foo } from './foo.js';", true},
		{"export type named alias", "export type { Foo as Bar } from './foo.js';", true},
		{"export type specifier", "export { type Foo } from './foo.js';", true},
		{"export type specifier alias", "export { type Foo as Bar } from './foo.js';", true},
		{"export type specifier multi", "export { type Foo, type Bar } from './foo.js';", true},
		{"named export", "export { Foo } from './foo.js';", false},
		{"named export alias", "export { Foo as Bar } from './foo.js';", false},
		{"mixed export", "export { Foo, type Bar } from './foo.js';", false},
		{"namespace export", "export * from './foo.js';", false},
		{"namespace alias export", "export * as FooNS from './foo.js';", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root, tree := parseStatement(tt.stmt)
			if tree == nil {
				t.Fatal("failed to parse statement")
			}
			defer tree.Close()
			cursor := root.Walk()
			defer cursor.Close()
			for _, child := range root.NamedChildren(cursor) {
				got := isTypeOnlyNode(&child)
				if got != tt.want {
					t.Errorf("isTypeOnlyNode(%q) = %v, want %v", tt.stmt, got, tt.want)
				}
			}
		})
	}
}

// Inline: pure function over constructed modules, table-driven comparisons
// are clearer than goldens (the fixture golden pins the sorted-path case)

// TestResolveReExportedCEDefinitionsStarChainOrderIndependence pins #450 item
// 2: star chains (a star-exports b, b star-exports c) must expand to the same
// export set regardless of module order, so re-sorting modules or generating
// from a different entry order cannot change the manifest.
func TestResolveReExportedCEDefinitionsStarChainOrderIndependence(t *testing.T) {
	makeChain := func() []M.Module {
		a := M.NewModule("src/a.js")
		b := M.NewModule("src/b.js")
		c := M.NewModule("src/c.js")

		c.Exports = append(c.Exports, &M.JavaScriptExport{
			Kind:        "js",
			Name:        "CThing",
			Declaration: M.NewReference("CThing", "", "src/c.js"),
		})

		b.Exports = append(b.Exports, &M.JavaScriptExport{
			Kind:        "js",
			Name:        "*",
			Declaration: M.NewReference("*", "", "src/c.js"),
		})
		b.Exports = append(b.Exports, &M.JavaScriptExport{
			Kind:        "js",
			Name:        "BThing",
			Declaration: M.NewReference("BThing", "", "src/b.js"),
		})
		b.ReExportAllSources = []string{"src/c.js"}

		a.Exports = append(a.Exports, &M.JavaScriptExport{
			Kind:        "js",
			Name:        "*",
			Declaration: M.NewReference("*", "", "src/b.js"),
		})
		a.Exports = append(a.Exports, &M.JavaScriptExport{
			Kind:        "js",
			Name:        "AThing",
			Declaration: M.NewReference("AThing", "", "src/a.js"),
		})
		a.ReExportAllSources = []string{"src/b.js"}

		return []M.Module{*a, *b, *c}
	}

	exportNames := func(pkg M.Package, path string) []string {
		var names []string
		for _, mod := range pkg.Modules {
			if mod.Path != path {
				continue
			}
			for _, e := range mod.Exports {
				switch ex := e.(type) {
				case *M.JavaScriptExport:
					names = append(names, ex.Name)
				case *M.CustomElementExport:
					names = append(names, ex.Name)
				}
			}
		}
		slices.Sort(names)
		return names
	}

	// src/a.js sits above both star sources in the chain, so a single pass in
	// path order expands only one link and drops CThing without the fixpoint
	forward := M.NewPackage(makeChain())
	resolveReExportedCEDefinitions(&forward)

	reversed := makeChain()
	slices.Reverse(reversed)
	reverse := M.NewPackage(reversed)
	resolveReExportedCEDefinitions(&reverse)

	want := []string{"*", "AThing", "BThing", "CThing"}
	if got := exportNames(forward, "src/a.js"); !slices.Equal(got, want) {
		t.Errorf("path order: src/a.js exports = %v, want %v", got, want)
	}
	if got := exportNames(reverse, "src/a.js"); !slices.Equal(got, want) {
		t.Errorf("reverse order: src/a.js exports = %v, want %v", got, want)
	}
}
