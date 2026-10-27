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
package typesplit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSplitTopLevelUnion verifies splitting across nested and quoted type
// forms, including malformed input.
//
// The cases stay inline rather than in testdata goldens on purpose.
// SplitTopLevelUnion is a pure string -> []string function, so a golden would
// require inventing a serialization format for both the input and the expected
// members. The table also mirrors internal/tstype/union_test.go, which tests
// the tree-sitter backed implementation of the same signature; keeping both
// tables in the same shape is what makes a parity gap visible when reading
// them, and TestPureSplitTopLevelUnionMatchesTreeSitter in internal/tstype
// compares the two implementations over this corpus. Finally, each case runs
// as a named subtest, so a failure names the form that regressed, which a
// single golden diff would not. internal/jsoncmerge/merge_test.go documents
// the same exception for comparisons that are not a single golden document.
func TestSplitTopLevelUnion(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want []string
	}{
		{"empty", "", nil},
		{"whitespace only", "   ", nil},
		{"leading pipe", "| 'a' | 'b'", []string{"'a'", "'b'"}},
		{"trailing pipe", "'a' | 'b' |", []string{"'a' | 'b' |"}},
		{"unbalanced", "'a' | (", []string{"'a' | ("}},
		{"stray closing angle", "A > | B", []string{"A > | B"}},
		{"stray closing angles", "A >> | B", []string{"A >> | B"}},
		{"statement terminator", "'a' | 'b';", []string{"'a'", "'b'"}},
		{"statement terminator with comment", "'a' | 'b'; // c", []string{"'a'", "'b'"}},
		{"statement terminator single type", "A;", []string{"A"}},
		{"single type", "string", []string{"string"}},
		{"simple union", "'a' | 'b'", []string{"'a'", "'b'"}},
		{"three-way union", "'a' | 'b' | 'c'", []string{"'a'", "'b'", "'c'"}},
		{"generic no split", "Extract<Foo, 'a' | 'b'>", []string{"Extract<Foo, 'a' | 'b'>"}},
		{"union of generics", "Extract<X, 'a'> | Exclude<Y, 'b'>", []string{"Extract<X, 'a'>", "Exclude<Y, 'b'>"}},
		{"nested parens", "(A | B) | C", []string{"(A | B)", "C"}},
		{"array generic", "Array<string | number>", []string{"Array<string | number>"}},
		{"mixed with null", "string[] | Extract<Union, 'x' | 'y'> | null", []string{"string[]", "Extract<Union, 'x' | 'y'>", "null"}},
		{"square brackets", "Record<string, 'a' | 'b'>[]", []string{"Record<string, 'a' | 'b'>[]"}},
		{"deeply nested", "Map<string, Set<'a' | 'b'>>", []string{"Map<string, Set<'a' | 'b'>>"}},
		{"whitespace preserved", " 'a'  |  'b' ", []string{"'a'", "'b'"}},
		{"template literal passthrough", "`${Side}-${Scale}`", []string{"`${Side}-${Scale}`"}},
		{"nested template literal passthrough", "`${`a|b`}`", []string{"`${`a|b`}`"}},
		{"template substitution union", "`${A | B}`", []string{"`${A | B}`"}},
		{"nested template with union", "`${`a|b`}` | null", []string{"`${`a|b`}`", "null"}},
		{"quoted pipe in substitution", "`${'a|b'}` | C", []string{"`${'a|b'}`", "C"}},
		{"escaped backtick in template", "A | `a\\`b` | C", []string{"A", "`a\\`b`", "C"}},
		{"Map generic with null", "Map<string, string> | null", []string{"Map<string, string>", "null"}},

		// Malformed template literals and comments inside substitutions. An
		// unterminated construct must be returned as one unchanged member.
		{"unterminated substitution", "`${A | B", []string{"`${A | B"}},
		{"unterminated template", "`abc | D", []string{"`abc | D"}},
		{"unterminated string in substitution", "`${'a} | C", []string{"`${'a} | C"}},
		{"unterminated block comment in substitution", "`${A /* } | C", []string{"`${A /* } | C"}},
		{"block comment brace in substitution", "`${A /* } */}` | C", []string{"`${A /* } */}`", "C"}},
		{"line comment brace in substitution", "`${A // }\n}` | C", []string{"`${A // }\n}`", "C"}},

		// These cases mirror tree-sitter's parse result so that callers moving
		// off the cgo-backed parser observe no change.
		{"pipe inside string", "'a|b' | 'c'", []string{"'a|b'", "'c'"}},
		{"pipe inside template substitution", "'a' | `b|${C | D}` | 'e'", []string{"'a'", "`b|${C | D}`", "'e'"}},
		{"arrow return union is opaque", "(x: string) => void | string", []string{"(x: string) => void | string"}},
		{"arrow return union generic is opaque", "(e: Event) => void | Promise<void>", []string{"(e: Event) => void | Promise<void>"}},
		{"parenthesized arrow still splits", "(() => void) | string", []string{"(() => void)", "string"}},
		{"constructor type splits", "new () => Foo | Bar", []string{"new () => Foo", "Bar"}},
		{"abstract constructor splits", "abstract new () => Foo | Bar", []string{"abstract new () => Foo", "Bar"}},
		{"union member arrow splits", "A | new () => B", []string{"A", "new () => B"}},
		{"readonly is opaque", "readonly A | B", []string{"readonly A | B"}},
		{"readonly array is opaque", "readonly A[] | B", []string{"readonly A[] | B"}},
		{"readonly identifier splits", "readonlyThing | x", []string{"readonlyThing", "x"}},
		{"conditional type is opaque", "T extends U ? X | Y : Z", []string{"T extends U ? X | Y : Z"}},
		{"conditional union is opaque", "A | T extends U ? X : Y", []string{"A | T extends U ? X : Y"}},
		{"intersection members", "A & B | C", []string{"A & B", "C"}},
		{"object type member", "{ a: string } | null", []string{"{ a: string }", "null"}},
		{"comment stripped", "// comment\nA | B", []string{"A", "B"}},
		{"line comment between members", "A // note\n| B", []string{"A", "B"}},
		{"block comment between members", "A /* note */ | B", []string{"A", "B"}},
		{"leading comment before first member", "/* lead */ A | B", []string{"A", "B"}},
		{"comment only", "// only comment", []string{"// only comment"}},
		{"non union trailing comment", "A /* trailing */", []string{"A /* trailing */"}},
		{"trailing comment after pipe is opaque", "'a' | // trailing comment", []string{"'a' | // trailing comment"}},
		{"comment inside brace member preserved", "{ /* c */ a: string } | null", []string{"{ /* c */ a: string }", "null"}},
		{"comment inside generic preserved", "Map</* k */string, string> | null", []string{"Map</* k */string, string>", "null"}},

		// A function type's return type extends to the right, so pipes after a
		// non-constructor arrow belong to the return type, not the root union.
		{"function return union after member", "undefined | (x: B) => C | null", []string{"undefined", "(x: B) => C | null"}},
		{"function type absorbs return union", "() => A | () => B", []string{"() => A | () => B"}},
		{"function return absorbs multiple pipes", "A | (x) => B | C | D", []string{"A", "(x) => B | C | D"}},
		{"bare function return union", "(x) => A | B | C", []string{"(x) => A | B | C"}},
		{"intersection with function return", "A & (x) => B | C", []string{"A & (x) => B | C"}},
		{"member then intersection function return", "A | B & (x) => C | D", []string{"A", "B & (x) => C | D"}},
		{"constructor then function type", "new () => A | (x) => B | C", []string{"new () => A", "(x) => B | C"}},
		{"two constructors then member", "new () => A | new () => B | C", []string{"new () => A", "new () => B", "C"}},
		{"constructor then function absorbs", "A | new () => B | (x) => C | D", []string{"A", "new () => B", "(x) => C | D"}},
		{"constructor after intersection", "A & new () => B | C", []string{"A & new () => B", "C"}},

		// readonly is a prefix type operator whose operand is the whole union to
		// its right, so it absorbs every following top-level pipe.
		{"readonly absorbs return after intersection", "A & readonly B[] | C", []string{"A & readonly B[] | C"}},
		{"readonly absorbs after member", "A | readonly B[] | C", []string{"A", "readonly B[] | C"}},
		{"readonly absorbs multiple", "A | readonly B[] | C | D", []string{"A", "readonly B[] | C | D"}},
		{"readonly absorbs at root", "readonly B[] | C | D", []string{"readonly B[] | C | D"}},
		{"readonly conditional after member", "A & B | readonly C extends D ? E : F", []string{"A & B", "readonly C extends D ? E : F"}},
		{"readonly conditional at root", "readonly A extends B ? C : D", []string{"readonly A extends B ? C : D"}},

		// A conditional type's check type extends left across earlier union
		// members, so the root is conditional rather than a union.
		{"conditional absorbs preceding union", "A | B extends C ? D : E | F", []string{"A | B extends C ? D : E | F"}},

		{"keyof intersection splits", "keyof A & B | C", []string{"keyof A & B", "C"}},
		{"keyof member splits", "A | keyof B | C", []string{"A", "keyof B", "C"}},
		{"readonly keyof is opaque", "readonly keyof A | B", []string{"readonly keyof A | B"}},
		{"keyof constructor absorbs first pipe", "keyof new () => A | B", []string{"keyof new () => A | B"}},
		{"keyof constructor absorbs one level", "keyof new () => A | B | C", []string{"keyof new () => A | B", "C"}},
		{"keyof abstract constructor absorbs", "keyof abstract new () => A | B", []string{"keyof abstract new () => A | B"}},
		{"keyof constructor intersection splits", "keyof new () => A & B | C", []string{"keyof new () => A & B", "C"}},
		{"keyof constructor member absorbs first pipe", "A | keyof new () => B | C | D", []string{"A", "keyof new () => B | C", "D"}},
		{"keyof constructor conditional return scoped", "A | keyof new () => B extends C ? D : E", []string{"A", "keyof new () => B extends C ? D : E"}},
		{"nested keyof constructor absorbs", "keyof keyof new () => A | B", []string{"keyof keyof new () => A | B"}},
		{"second keyof constructor absorbs", "keyof new () => A | keyof new () => B | C", []string{"keyof new () => A | keyof new () => B | C"}},
		{"keyof constructor conditional at root", "A | keyof new () => B | C extends D ? E : F", []string{"A | keyof new () => B | C extends D ? E : F"}},
		{"nested keyof in constructor return", "keyof new () => keyof B | C", []string{"keyof new () => keyof B | C"}},
		{"nested keyof in constructor return after member", "A | keyof new () => keyof B | C", []string{"A", "keyof new () => keyof B | C"}},
		{"nested keyof constructor absorbs one level", "keyof new () => keyof A | B | C", []string{"keyof new () => keyof A | B", "C"}},
		{"keyof constructor conditional branch union", "keyof new () => A extends B ? C : D | E | F", []string{"keyof new () => A extends B ? C : D | E | F"}},
		{"keyof constructor conditional branch after member", "A | keyof new () => B extends C ? D : E | F | G", []string{"A", "keyof new () => B extends C ? D : E | F | G"}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := SplitTopLevelUnion(tc.in)
			assert.Equal(t, tc.want, got)
		})
	}
}
