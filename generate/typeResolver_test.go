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

func TestParseGenericType(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		wantName string
		wantArgs []string
		wantOk   bool
	}{
		{"simple generic", "Array<string>", "Array", []string{"string"}, true},
		{"two args", "Extract<Foo, Bar>", "Extract", []string{"Foo", "Bar"}, true},
		{"nested union arg", "Extract<'a' | 'b' | 'c', 'a' | 'c'>", "Extract", []string{"'a' | 'b' | 'c'", "'a' | 'c'"}, true},
		{"nested generic arg", "Readonly<Array<string>>", "Readonly", []string{"Array<string>"}, true},
		{"deeply nested", "Map<string, Set<'a' | 'b'>>", "Map", []string{"string", "Set<'a' | 'b'>"}, true},
		{"not generic - identifier", "string", "", nil, false},
		{"not generic - array", "string[]", "", nil, false},
		{"not generic - union", "'a' | 'b'", "", nil, false},
		{"not generic - literal", "'hello'", "", nil, false},
		{"empty", "", "", nil, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			name, args, ok := parseGenericType(tc.in)
			assert.Equal(t, tc.wantOk, ok)
			if ok {
				assert.Equal(t, tc.wantName, name)
				assert.Equal(t, tc.wantArgs, args)
			}
		})
	}
}

func TestResolveUtilityTypeText(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		aliases map[string]string
		want    string
	}{
		{
			"Extract with literals",
			"Extract<'a' | 'b' | 'c', 'a' | 'c'>",
			nil,
			"'a' | 'c'",
		},
		{
			"Extract with alias",
			"Extract<ColorPalette, 'lightest' | 'darkest'>",
			map[string]string{"ColorPalette": "'light' | 'lighter' | 'lightest' | 'dark' | 'darker' | 'darkest'"},
			"'lightest' | 'darkest'",
		},
		{
			"Exclude with literals",
			"Exclude<'a' | 'b' | 'c', 'b'>",
			nil,
			"'a' | 'c'",
		},
		{
			"Exclude removes multiple",
			"Exclude<'a' | 'b' | 'c', 'a' | 'c'>",
			nil,
			"'b'",
		},
		{
			"Array<string> normalizes",
			"Array<string>",
			nil,
			"string[]",
		},
		{
			"Array<union> normalizes",
			"Array<string | number>",
			nil,
			"(string | number)[]",
		},
		{
			"Readonly unwraps",
			"Readonly<'a' | 'b'>",
			nil,
			"'a' | 'b'",
		},
		{
			"Required unwraps",
			"Required<'a' | 'b'>",
			nil,
			"'a' | 'b'",
		},
		{
			"NonNullable filters",
			"NonNullable<'a' | null | undefined | 'b'>",
			nil,
			"'a' | 'b'",
		},
		{
			"NonNullable alias",
			"NonNullable<MaybeColor>",
			map[string]string{"MaybeColor": "'red' | 'blue' | null"},
			"'red' | 'blue'",
		},
		{
			"Readonly with alias resolution",
			"Readonly<Variant>",
			map[string]string{"Variant": "'primary' | 'secondary'"},
			"'primary' | 'secondary'",
		},
		{
			"string[] passthrough",
			"string[]",
			nil,
			"string[]",
		},
		{
			"unknown generic passthrough",
			"Partial<Config>",
			nil,
			"Partial<Config>",
		},
		{
			"Extract empty result",
			"Extract<'a' | 'b', 'x'>",
			nil,
			"never",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pkg := &M.Package{
				Modules: []M.Module{{Path: "test.ts"}},
			}
			typeAliases := make(moduleTypeAliasesMap)
			if tc.aliases != nil {
				typeAliases["test.ts"] = tc.aliases
			}
			imports := make(moduleImportsMap)
			ctx := newResolutionContext()
			resolved, _ := resolveTypeText(tc.in, &pkg.Modules[0], pkg, typeAliases, imports, nil, ctx)
			assert.Equal(t, tc.want, resolved)
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
