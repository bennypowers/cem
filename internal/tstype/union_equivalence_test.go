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
		"`${`a|b`}`",
		"`${A | B}`",
		"`${`a|b`}` | null",
		"`${'a|b'}` | C",
		"A | `a\\`b` | C",
		"undefined | (x: B) => C | null",
		"() => A | () => B",
		"A | (x) => B | C | D",
		"(x) => A | B | C",
		"A & (x) => B | C",
		"A | B & (x) => C | D",
		"new () => A | (x) => B | C",
		"new () => A | new () => B | C",
		"A | new () => B | (x) => C | D",
		"A & new () => B | C",
		"A | new () => B | C",
		"readonly B[] | C | D",
		"A | readonly B[] | C",
		"A | readonly B[] | C | D",
		"A & readonly B[] | C",
		"A & readonly B[] & C | D",
		"A & B | readonly C extends D ? E : F",
		"readonly A extends B ? C : D",
		"A | B extends C ? D : E | F",
		"A | T extends U ? X | Y : Z",
		"keyof A & B | C",
		"A | keyof B | C",
		"readonly keyof A | B",
		"keyof new () => A | B",
		"keyof new () => A | B | C",
		"keyof abstract new () => A | B",
		"keyof new () => A & B | C",
		"A | keyof new () => B | C | D",
		"keyof new () => A | readonly B | C",
		"keyof new () => A | (x) => B | C",
		"A | keyof new () => B extends C ? D : E",
		"new () => A | keyof new () => B extends C ? D : E",
		"A | keyof new () => B extends C ? D : E | F",
		"keyof keyof new () => A | B",
		"keyof new () => A | keyof new () => B | C",
		"A | keyof new () => B | C extends D ? E : F",
		"keyof new () => A & B | D extends E ? F : G",
		"keyof new () => A | B extends C ? D : E",
		"keyof new () => keyof B | C",
		"A | keyof new () => keyof B | C",
		"keyof new () => keyof A | B | C",
		"keyof new () => A extends B ? C : D | E | F",
		"A | keyof new () => B extends C ? D : E | F | G",
		"keyof new () => A extends B ? C | D : E | F",
		"{ /* c */ a: string } | null",
		"Map</* k */string, string> | null",
		"/* lead */ A | B",
		"A /* trailing */",
		"'a' | // trailing comment",
		"// only comment",
		"'a' | 'b';",
		"'a' | 'b'; // c",
	}
	for _, typeText := range types {
		t.Run(typeText, func(t *testing.T) {
			want := SplitTopLevelUnion(typeText)
			got := typesplit.SplitTopLevelUnion(typeText)
			assert.Equal(t, want, got)
		})
	}
}

// TestCommentMembersAreIntentionalDivergence documents that the pure splitter
// strips comments between union members, while tree-sitter exposes them as
// members. Comments are not types, so the pure result is the intended one.
func TestCommentMembersAreIntentionalDivergence(t *testing.T) {
	const typeText = "A // note\n| B"
	assert.Equal(t, []string{"A", "B"}, typesplit.SplitTopLevelUnion(typeText))
	assert.Equal(t, []string{"A", "// note", "B"}, SplitTopLevelUnion(typeText))
}
