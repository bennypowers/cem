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
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := SplitTopLevelUnion(tc.in)
			assert.Equal(t, tc.want, got)
		})
	}
}
