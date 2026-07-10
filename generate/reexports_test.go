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
	"testing"

	"bennypowers.dev/cem/internal/languages/typescript"
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
