package tstype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitTopLevelUnion(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want []string
	}{
		{"empty", "", nil},
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
		{"Map generic with null", "Map<string, string> | null", []string{"Map<string, string>", "null"}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := SplitTopLevelUnion(tc.in)
			assert.Equal(t, tc.want, got)
		})
	}
}
