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
