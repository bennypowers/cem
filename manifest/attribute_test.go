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
package manifest_test

import (
	"testing"

	"bennypowers.dev/cem/manifest"
	"github.com/stretchr/testify/assert"
)

// TestAttributeEnumValidation covers enum detection and value validation.
//
// Inline assertions are used here: the behavior under test is a handful of
// predicates over a single string, so a golden file would obscure the cases
// rather than document them.
func TestAttributeEnumValidation(t *testing.T) {
	tests := []struct {
		name     string
		typeText string
		isEnum   bool
		values   []string
		valid    []string
		invalid  []string
	}{
		{
			name:     "single quotes",
			typeText: "'primary' | 'secondary' | 'danger'",
			isEnum:   true,
			values:   []string{"'primary'", "'secondary'", "'danger'"},
			valid:    []string{"'primary'"},
			invalid:  []string{"primary", "other"},
		},
		{
			name:     "double quotes unquoted on validation",
			typeText: `"a" | "b"`,
			isEnum:   true,
			values:   []string{`"a"`, `"b"`},
			valid:    []string{`"a"`, "a"},
			invalid:  []string{"c"},
		},
		{
			name:     "generic union splits at top level",
			typeText: "Array<string | number> | null",
			isEnum:   true,
			values:   []string{"Array<string | number>", "null"},
			valid:    []string{"null"},
			invalid:  []string{"string"},
		},
		{
			name:     "nested template literal is not an enum",
			typeText: "`${`a|b`}`",
			isEnum:   false,
			values:   nil,
			valid:    []string{"a"},
		},
		{
			name:     "nested template literal union splits only at top level",
			typeText: "`${`a|b`}` | null",
			isEnum:   true,
			values:   []string{"`${`a|b`}`", "null"},
			valid:    []string{"null"},
			invalid:  []string{"a"},
		},
		{
			name:     "template substitution union is not an enum",
			typeText: "`${A | B}`",
			isEnum:   false,
			values:   nil,
			valid:    []string{"A"},
		},
		{
			name:     "function return union splits only outer members",
			typeText: "undefined | (x: B) => C | null",
			isEnum:   true,
			values:   []string{"undefined", "(x: B) => C | null"},
			valid:    []string{"undefined"},
			invalid:  []string{"null", "(x: B) => C"},
		},
		{
			name:     "bare function type is not an enum",
			typeText: "() => A | () => B",
			isEnum:   false,
			values:   nil,
			valid:    []string{"anything"},
		},
		{
			name:     "readonly absorbs later members",
			typeText: "A | readonly B[] | C",
			isEnum:   true,
			values:   []string{"A", "readonly B[] | C"},
			valid:    []string{"A", "readonly B[] | C"},
			invalid:  []string{"C"},
		},
		{
			name:     "comment inside member is preserved",
			typeText: "{ /* c */ a: string } | null",
			isEnum:   true,
			values:   []string{"{ /* c */ a: string }", "null"},
			valid:    []string{"null"},
			invalid:  []string{"{   a: string }"},
		},
		{
			name:     "trailing statement terminator",
			typeText: "'a' | 'b';",
			isEnum:   true,
			values:   []string{"'a'", "'b'"},
			valid:    []string{"'a'", "'b'"},
			invalid:  []string{"b"},
		},
		{
			name:     "comments between members are not members",
			typeText: "A // note\n| B",
			isEnum:   true,
			values:   []string{"A", "B"},
			valid:    []string{"A", "B"},
			invalid:  []string{"// note"},
		},
		{
			name:     "non union",
			typeText: "string",
			isEnum:   false,
			values:   nil,
			valid:    []string{"anything"},
		},
		{
			name:     "readonly union is not an enum",
			typeText: "readonly A | B",
			isEnum:   false,
			values:   nil,
			valid:    []string{"A"},
		},
		{
			name:     "function return union is not an enum",
			typeText: "(e: Event) => void | Promise<void>",
			isEnum:   false,
			values:   nil,
			valid:    []string{"whatever"},
		},
		{
			name:     "conditional type is not an enum",
			typeText: "T extends U ? X | Y : Z",
			isEnum:   false,
			values:   nil,
			valid:    []string{"anything"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			attr := &manifest.Attribute{
				Type: &manifest.Type{Text: tc.typeText},
			}
			assert.Equal(t, tc.isEnum, attr.IsEnum())
			assert.Equal(t, tc.values, attr.EnumValues())
			for _, value := range tc.valid {
				assert.True(t, attr.IsValidValue(value), "expected %q to be valid", value)
				assert.NoError(t, attr.GetValidationError(value))
			}
			for _, value := range tc.invalid {
				assert.False(t, attr.IsValidValue(value), "expected %q to be invalid", value)
				assert.Error(t, attr.GetValidationError(value))
			}
		})
	}
}
