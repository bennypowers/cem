/*
Copyright © 2026 Benny Powers

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

	M "bennypowers.dev/cem/manifest"
	"github.com/stretchr/testify/assert"
)

// Inline: pure function, table-driven

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
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := splitTopLevelUnion(tc.in)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestFindModuleBySpec(t *testing.T) {
	pkg := &M.Package{
		Modules: []M.Module{
			{Path: "src/button.js"},
			{Path: "src/utils/helpers.js"},
			{Path: "lib/types.js"},
			{Path: "src/index.js"},
			{Path: "src/deep/nested/thing.ts"},
		},
	}

	tests := []struct {
		name              string
		currentModulePath string
		importSpec        string
		wantPath          string
		wantNil           bool
	}{
		{
			"relative import same dir",
			"src/button.js",
			"./utils/helpers.js",
			"src/utils/helpers.js",
			false,
		},
		{
			"relative import without extension",
			"src/button.js",
			"./utils/helpers",
			"src/utils/helpers.js",
			false,
		},
		{
			"relative import parent dir",
			"src/utils/helpers.js",
			"../button.js",
			"src/button.js",
			false,
		},
		{
			"package-scoped import",
			"src/button.js",
			"@my/pkg/lib/types.js",
			"lib/types.js",
			false,
		},
		{
			"package-scoped without extension",
			"src/button.js",
			"@my/pkg/lib/types",
			"lib/types.js",
			false,
		},
		{
			"relative import index.js",
			"src/button.js",
			"./deep/nested/thing",
			"src/deep/nested/thing.ts",
			false,
		},
		{
			"nonexistent module",
			"src/button.js",
			"./nonexistent",
			"",
			true,
		},
		{
			"invalid package-scoped import",
			"src/button.js",
			"@only/scope",
			"",
			true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := findModuleBySpec(pkg, tc.currentModulePath, tc.importSpec)
			if tc.wantNil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, tc.wantPath, result.Path)
			}
		})
	}
}
