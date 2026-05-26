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

	"github.com/stretchr/testify/assert"
)

// Inline: pure function, table-driven

func TestIsQuoted(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"'single'", true},
		{`"double"`, true},
		{"'mismatched\"", false},
		{"\"mismatched'", false},
		{"unquoted", false},
		{"", false},
		{"'", true},
		{"''", true},
		{`""`, true},
		{"'a", false},
		{"a'", false},
	}
	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			assert.Equal(t, tc.want, isQuoted(tc.in))
		})
	}
}

func TestIsNumericLiteral(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"0", true},
		{"42", true},
		{"-5", true},
		{"1.5", true},
		{"-0.5", true},
		{"100", true},
		{"", false},
		{"-", false},
		{".", false},
		{"-.", false},
		{"1.2.3", false},
		{"abc", false},
		{"1a", false},
		{"-1.5", true},
		{"--1", false},
		{"1-", false},
	}
	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			assert.Equal(t, tc.want, isNumericLiteral(tc.in))
		})
	}
}

func TestReconstructTemplate(t *testing.T) {
	tests := []struct {
		name string
		td   *templateDef
		want string
	}{
		{
			"single expression",
			&templateDef{statics: []string{"pre-", ""}, exprs: []string{"Side"}},
			"`pre-${Side}`",
		},
		{
			"two expressions",
			&templateDef{statics: []string{"a-", "-", ""}, exprs: []string{"X", "Y"}},
			"`a-${X}-${Y}`",
		},
		{
			"no static prefix",
			&templateDef{statics: []string{"", "-end"}, exprs: []string{"X"}},
			"`${X}-end`",
		},
		{
			"only expression",
			&templateDef{statics: []string{"", ""}, exprs: []string{"X"}},
			"`${X}`",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, reconstructTemplate(tc.td))
		})
	}
}

func TestValidateExpressionValues(t *testing.T) {
	tests := []struct {
		name       string
		resolved   string
		wantValues []string
		wantBail   string
	}{
		{
			"quoted strings",
			"'a' | 'b' | 'c'",
			[]string{"a", "b", "c"},
			"",
		},
		{
			"numeric literals",
			"1 | 2 | 3",
			[]string{"1", "2", "3"},
			"",
		},
		{
			"mixed quoted and numeric",
			"'a' | 1 | 'b'",
			[]string{"a", "1", "b"},
			"",
		},
		{
			"undefined and null skipped",
			"'a' | undefined | null | 'b'",
			[]string{"a", "b"},
			"",
		},
		{
			"broad primitive collapses to string",
			"'a' | string",
			nil,
			"string",
		},
		{
			"number primitive collapses",
			"'a' | number",
			nil,
			"string",
		},
		{
			"any primitive collapses",
			"any",
			nil,
			"string",
		},
		{
			"boolean primitive collapses",
			"boolean",
			nil,
			"string",
		},
		{
			"unknown primitive collapses",
			"unknown",
			nil,
			"string",
		},
		{
			"unresolved identifier bails",
			"SomeType",
			nil,
			"unresolvable",
		},
		{
			"only undefined/null is unresolvable",
			"undefined | null",
			nil,
			"unresolvable",
		},
		{
			"empty is unresolvable",
			"",
			nil,
			"unresolvable",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			values, bail := validateExpressionValues(tc.resolved)
			assert.Equal(t, tc.wantValues, values)
			assert.Equal(t, tc.wantBail, bail)
		})
	}
}

func TestResolveTextParts(t *testing.T) {
	tests := []struct {
		name    string
		def     string
		aliases map[string]aliasDefinition
		want    string
	}{
		{"empty", "", nil, ""},
		{"quoted passthrough", "'hello'", nil, "'hello'"},
		{"primitive passthrough", "string", nil, "string"},
		{"number primitive", "number", nil, "number"},
		{"unknown identifier passthrough", "UnknownType", nil, "UnknownType"},
		{
			"alias resolution",
			"MyType",
			map[string]aliasDefinition{"MyType": {text: "'a' | 'b'"}},
			"'a' | 'b'",
		},
		{
			"union resolution",
			"'a' | MyType",
			map[string]aliasDefinition{"MyType": {text: "'b' | 'c'"}},
			"'a' | 'b' | 'c'",
		},
		{
			"circular reference stops",
			"A",
			map[string]aliasDefinition{"A": {text: "A"}},
			"A",
		},
		{
			"nested alias chain",
			"Outer",
			map[string]aliasDefinition{
				"Outer": {text: "Inner"},
				"Inner": {text: "'val'"},
			},
			"'val'",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			visited := make(map[string]bool)
			got := resolveTextParts(tc.def, tc.aliases, visited)
			assert.Equal(t, tc.want, got)
		})
	}
}
