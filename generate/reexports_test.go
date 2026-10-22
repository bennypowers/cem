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
	"fmt"
	"slices"
	"strconv"
	"testing"
	"time"

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
// export set and the same export SEQUENCE regardless of module order, so
// re-sorting modules or generating from a different entry order cannot change
// the manifest at all. Star-copied exports carry StartByte 0, so their
// relative order needs the deterministic kind/name tie-break; without it the
// order is whatever order the copies were appended in, which depends on
// module order.
func TestResolveReExportedCEDefinitionsStarChainOrderIndependence(t *testing.T) {
	// StartByte values mirror real source positions: each module's own
	// exports carry their statement offsets, while exports copied by star
	// expansion always carry StartByte 0 and tie with each other.
	makeChain := func() []M.Module {
		a := M.NewModule("src/a.js")
		b := M.NewModule("src/b.js")
		c := M.NewModule("src/c.js")

		c.Exports = append(c.Exports, &M.JavaScriptExport{
			Kind:        "js",
			Name:        "CThing",
			Declaration: M.NewReference("CThing", "", "src/c.js"),
			StartByte:   40,
		})

		b.Exports = append(b.Exports, &M.JavaScriptExport{
			Kind:        "js",
			Name:        "*",
			Declaration: M.NewReference("*", "", "src/c.js"),
			StartByte:   10,
		})
		b.Exports = append(b.Exports, &M.JavaScriptExport{
			Kind:        "js",
			Name:        "BThing",
			Declaration: M.NewReference("BThing", "", "src/b.js"),
			StartByte:   30,
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
			StartByte:   50,
		})
		a.ReExportAllSources = []string{"src/b.js"}

		return []M.Module{*a, *b, *c}
	}

	makeCycle := func() []M.Module {
		p := M.NewModule("src/p.js")
		q := M.NewModule("src/q.js")

		p.Exports = append(p.Exports, &M.JavaScriptExport{
			Kind:        "js",
			Name:        "*",
			Declaration: M.NewReference("*", "", "src/q.js"),
		})
		p.Exports = append(p.Exports, &M.JavaScriptExport{
			Kind:        "js",
			Name:        "PThing",
			Declaration: M.NewReference("PThing", "", "src/p.js"),
		})
		p.ReExportAllSources = []string{"src/q.js"}

		q.Exports = append(q.Exports, &M.JavaScriptExport{
			Kind:        "js",
			Name:        "*",
			Declaration: M.NewReference("*", "", "src/p.js"),
		})
		q.Exports = append(q.Exports, &M.JavaScriptExport{
			Kind:        "js",
			Name:        "QThing",
			Declaration: M.NewReference("QThing", "", "src/q.js"),
		})
		q.ReExportAllSources = []string{"src/p.js"}

		return []M.Module{*p, *q}
	}

	exportSequence := func(pkg M.Package, path string) []string {
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
		return names
	}

	exportNames := func(pkg M.Package, path string) []string {
		names := exportSequence(pkg, path)
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

	// the sequence, not just the set, must not depend on module order: the
	// zero-StartByte star copies tie, so only the kind/name tie-break makes
	// their relative order deterministic
	wantSequence := []string{"*", "BThing", "CThing", "AThing"}
	if got := exportSequence(forward, "src/a.js"); !slices.Equal(got, wantSequence) {
		t.Errorf("path order: src/a.js export sequence = %v, want %v", got, wantSequence)
	}
	if got := exportSequence(reverse, "src/a.js"); !slices.Equal(got, wantSequence) {
		t.Errorf("reverse order: src/a.js export sequence = %v, want %v", got, wantSequence)
	}

	// Cyclic star re-exports (p star-exports q, q star-exports p) must
	// terminate: exports only ever grow and are deduplicated by key, so the
	// fixpoint reaches a pass that adds nothing. A dedup regression would
	// hang rather than fail an assertion, so the call runs under a
	// completion guard.
	cycle := M.NewPackage(makeCycle())
	done := make(chan struct{})
	go func() {
		resolveReExportedCEDefinitions(&cycle)
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("resolveReExportedCEDefinitions did not terminate on cyclic star re-exports")
	}

	wantCycle := []string{"*", "PThing", "QThing"}
	if got := exportNames(cycle, "src/p.js"); !slices.Equal(got, wantCycle) {
		t.Errorf("cyclic: src/p.js exports = %v, want %v", got, wantCycle)
	}
	if got := exportNames(cycle, "src/q.js"); !slices.Equal(got, wantCycle) {
		t.Errorf("cyclic: src/q.js exports = %v, want %v", got, wantCycle)
	}
}

// BenchmarkResolveReExportedCEDefinitions measures the worklist fixpoint's
// scaling on synthetic forward star chains (module i star-exports module i+1,
// so every re-exporter sorts before its source in path order — the worst
// shape for whole-sweep repetition, where each sweep advances only one chain
// link). Modules are rebuilt per iteration because the fixpoint mutates the
// package in place. No test asserts on these times; the benchmark exists to
// document the scaling curve.
func BenchmarkResolveReExportedCEDefinitions(b *testing.B) {
	for _, size := range []int{50, 100, 200} {
		b.Run(strconv.Itoa(size), func(b *testing.B) {
			newModules := func() []M.Module {
				modules := make([]M.Module, 0, size)
				for i := range size {
					path := fmt.Sprintf("src/m%04d.js", i)
					mod := M.NewModule(path)
					if i < size-1 {
						next := fmt.Sprintf("src/m%04d.js", i+1)
						mod.Exports = append(mod.Exports, &M.JavaScriptExport{
							Kind:        "js",
							Name:        "*",
							Declaration: M.NewReference("*", "", next),
						})
						mod.ReExportAllSources = []string{next}
					}
					mod.Exports = append(mod.Exports, &M.JavaScriptExport{
						Kind:        "js",
						Name:        fmt.Sprintf("Thing%d", i),
						Declaration: M.NewReference(fmt.Sprintf("Thing%d", i), "", path),
					})
					modules = append(modules, *mod)
				}
				return modules
			}
			b.ResetTimer()
			for b.Loop() {
				pkg := M.NewPackage(newModules())
				resolveReExportedCEDefinitions(&pkg)
			}
		})
	}
}
