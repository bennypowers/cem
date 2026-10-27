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
package tstype

import (
	"testing"

	"bennypowers.dev/cem/internal/typesplit"
	"github.com/stretchr/testify/assert"
)

// TestPureSplitTopLevelUnionMatchesTreeSitter guards the cgo-free splitter
// used by the manifest package against drift from the tree-sitter parser.
func TestPureSplitTopLevelUnionMatchesTreeSitter(t *testing.T) {
	types := []string{
		"",
		"   ",
		"string",
		"'a' | 'b'",
		"'a' | 'b' | 'c'",
		"Extract<Foo, 'a' | 'b'>",
		"Extract<X, 'a'> | Exclude<Y, 'b'>",
		"(A | B) | C",
		"Array<string | number>",
		"string[] | Extract<Union, 'x' | 'y'> | null",
		"Record<string, 'a' | 'b'>[]",
		"Map<string, Set<'a' | 'b'>>",
		" 'a'  |  'b' ",
		"`${Side}-${Scale}`",
		"Map<string, string> | null",
		"'a|b' | 'c'",
		"'a' | `b|${C | D}` | 'e'",
		"(x: string) => void | string",
		"(e: Event) => void | Promise<void>",
		"(() => void) | string",
		"new () => Foo | Bar",
		"abstract new () => Foo | Bar",
		"A | new () => B",
		"readonly A | B",
		"readonly A[] | B",
		"readonlyThing | x",
		"T extends U ? X | Y : Z",
		"A | T extends U ? X : Y",
		"A & B | C",
		"{ a: string } | null",
		"// comment\nA | B",
		"| 'a' | 'b'",
		"'a' | 'b' |",
		"'a' | (",
	}
	for _, typeText := range types {
		t.Run(typeText, func(t *testing.T) {
			want := SplitTopLevelUnion(typeText)
			got := typesplit.SplitTopLevelUnion(typeText)
			assert.Equal(t, want, got)
		})
	}
}
