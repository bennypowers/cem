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
			valid:    []string{"primary", "secondary", "danger"},
			invalid:  []string{"other", "'primary'", `"primary"`},
		},
		{
			name:     "double quotes unquoted on validation",
			typeText: `"a" | "b"`,
			isEnum:   true,
			values:   []string{`"a"`, `"b"`},
			valid:    []string{"a", "b"},
			invalid:  []string{"c", `"a"`},
		},
		{
			name:     "both quote styles",
			typeText: `'a' | "b"`,
			isEnum:   true,
			values:   []string{"'a'", `"b"`},
			valid:    []string{"a", "b"},
			invalid:  []string{"c", "'a'", "'b'", `"b"`},
		},
		{
			name:     "numeric literals",
			typeText: "0 | 1 | -2 | 3.5 | 1e3 | 0xFF | 0b101 | 0o17 | 1_000 | 0x1_F",
			isEnum:   true,
			values:   []string{"0", "1", "-2", "3.5", "1e3", "0xFF", "0b101", "0o17", "1_000", "0x1_F"},
			valid:    []string{"0", "1", "-2", "3.5", "1e3", "0xFF", "0b101", "0o17", "1_000", "0x1_F"},
			invalid:  []string{"4", "one", "0x10", "+3", "010"},
		},
		{
			name:     "numeric literal and quoted numeral string",
			typeText: "1 | '2'",
			isEnum:   true,
			values:   []string{"1", "'2'"},
			valid:    []string{"1", "2"},
			invalid:  []string{"3", "'1'", "'2'"},
		},
		{
			name:     "malformed numeric members are not literals",
			typeText: "'a' | 0x | 1n | 1__0 | 1e",
			isEnum:   true,
			values:   []string{"'a'", "0x", "1n", "1__0", "1e"},
			valid:    []string{"a"},
			invalid:  []string{"0x", "1n", "1__0", "1e", "'a'"},
		},
		{
			name:     "boolean literals",
			typeText: "true | false",
			isEnum:   true,
			values:   []string{"true", "false"},
			valid:    []string{"true", "false"},
			invalid:  []string{"True", "1"},
		},
		{
			name:     "bare identifiers are not attribute values",
			typeText: "'sm' | null | undefined | SomeType | SOME_CONST",
			isEnum:   true,
			values:   []string{"'sm'", "null", "undefined", "SomeType", "SOME_CONST"},
			valid:    []string{"sm"},
			invalid:  []string{"null", "undefined", "SomeType", "SOME_CONST", "'sm'"},
		},
		{
			name:     "literal mixed with null",
			typeText: "'sm' | null",
			isEnum:   true,
			values:   []string{"'sm'", "null"},
			valid:    []string{"sm"},
			invalid:  []string{"null", "md", "'sm'"},
		},
		{
			name:     "empty string literal",
			typeText: "'' | 'a'",
			isEnum:   true,
			values:   []string{"''", "'a'"},
			valid:    []string{"", "a"},
			invalid:  []string{"b", "''"},
		},
		{
			name:     "escaped single quote",
			typeText: `'don\'t' | 'other'`,
			isEnum:   true,
			values:   []string{`'don\'t'`, "'other'"},
			valid:    []string{`don't`, "other"},
			invalid:  []string{"dont", `'don\'t'`},
		},
		{
			name:     "double quote inside single quoted literal",
			typeText: `'a"b' | 'c'`,
			isEnum:   true,
			values:   []string{`'a"b'`, "'c'"},
			valid:    []string{`a"b`, "c"},
			invalid:  []string{"a", `'a"b'`},
		},
		{
			name:     "intersection member is not a string literal",
			typeText: "'a' & 'b' | 'c'",
			isEnum:   true,
			values:   []string{"'a' & 'b'", "'c'"},
			valid:    []string{"c"},
			invalid:  []string{"a", "'a' & 'b'", "'c'"},
		},
		{
			name:     "standard and unicode escapes",
			typeText: `'a\tb' | '\u0041' | '\x42' | '\u{43}'`,
			isEnum:   true,
			values:   []string{`'a\tb'`, `'\u0041'`, `'\x42'`, `'\u{43}'`},
			valid:    []string{"a\tb", "A", "B", "C"},
			invalid:  []string{"atb", "u0041", "x42", "u{43}"},
		},
		{
			name:     "non-ascii hex escape is a code point",
			typeText: `'\xE9' | 'x'`,
			isEnum:   true,
			values:   []string{`'\xE9'`, "'x'"},
			valid:    []string{"é", "x"},
			invalid:  []string{"Ã©"},
		},
		{
			name:     "identity escape",
			typeText: `'C:\path' | 'other'`,
			isEnum:   true,
			values:   []string{`'C:\path'`, "'other'"},
			valid:    []string{`C:path`, "other"},
			invalid:  []string{`C:\path`, "path"},
		},
		{
			name:     "surrogate pair escape",
			typeText: `'\uD83D\uDE00' | 'x'`,
			isEnum:   true,
			values:   []string{`'\uD83D\uDE00'`, "'x'"},
			valid:    []string{"😀", "x"},
			invalid:  []string{"uD83DuDE00"},
		},
		{
			name:     "control escapes",
			typeText: `'\n\r\b\f\v\0' | 'x'`,
			isEnum:   true,
			values:   []string{`'\n\r\b\f\v\0'`, "'x'"},
			valid:    []string{"\n\r\b\f\v\x00", "x"},
			invalid:  []string{"n"},
		},
		{
			name:     "malformed and lone surrogate escapes fall back to identity",
			typeText: `'\xZZ' | '\u12' | '\u{41' | '\uD83D' | '\u{D800}' | 'x'`,
			isEnum:   true,
			values:   []string{`'\xZZ'`, `'\u12'`, `'\u{41'`, `'\uD83D'`, `'\u{D800}'`, "'x'"},
			valid:    []string{"xZZ", "u12", "u{41", "u{D800}", "x"},
			invalid:  []string{"uD83D", "D83D", "D800"},
		},
		{
			name:     "generic union splits at top level",
			typeText: "Array<string | number> | null",
			isEnum:   true,
			values:   []string{"Array<string | number>", "null"},
			valid:    nil,
			invalid:  []string{"null", "string"},
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
			valid:    nil,
			invalid:  []string{"a", "null"},
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
			valid:    nil,
			invalid:  []string{"undefined", "null", "(x: B) => C"},
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
			valid:    nil,
			invalid:  []string{"A", "readonly B[] | C", "C"},
		},
		{
			name:     "comment inside member is preserved",
			typeText: "{ /* c */ a: string } | null",
			isEnum:   true,
			values:   []string{"{ /* c */ a: string }", "null"},
			valid:    nil,
			invalid:  []string{"null", "{   a: string }"},
		},
		{
			name:     "trailing statement terminator",
			typeText: "'a' | 'b';",
			isEnum:   true,
			values:   []string{"'a'", "'b'"},
			valid:    []string{"a", "b"},
			invalid:  []string{"c", "'a'", "'b'"},
		},
		{
			name:     "comments between members are not members",
			typeText: "A // note\n| B",
			isEnum:   true,
			values:   []string{"A", "B"},
			valid:    nil,
			invalid:  []string{"A", "B", "// note"},
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

// TestAttributeDisplayEnumValues covers the user-facing list shown when a value
// is rejected. TypeScript source quoting is presentation detail, so the values
// are displayed the same way regardless of the source quote style.
//
// Inline table assertions are used because each case is a pure function of a
// type string; a golden would add indirection without documenting more.
func TestAttributeDisplayEnumValues(t *testing.T) {
	tests := []struct {
		name     string
		typeText string
		want     []string
	}{
		{
			name:     "single quotes",
			typeText: "'primary' | 'secondary'",
			want:     []string{`"primary"`, `"secondary"`},
		},
		{
			name:     "double quotes",
			typeText: `"primary" | "secondary"`,
			want:     []string{`"primary"`, `"secondary"`},
		},
		{
			name:     "mixed quote styles",
			typeText: `'a' | "b"`,
			want:     []string{`"a"`, `"b"`},
		},
		{
			name:     "non literals are omitted",
			typeText: "'sm' | null | undefined | SomeType",
			want:     []string{`"sm"`},
		},
		{
			name:     "numeric and boolean literals are not quoted",
			typeText: "0 | 1 | true | false",
			want:     []string{"0", "1", "true", "false"},
		},
		{
			name:     "empty string literal",
			typeText: "'' | 'a'",
			want:     []string{`""`, `"a"`},
		},
		{
			name:     "escaped single quote is decoded",
			typeText: `'don\'t' | 'other'`,
			want:     []string{`"don't"`, `"other"`},
		},
		{
			name:     "double quote inside single quoted literal",
			typeText: `'a"b' | 'c'`,
			want:     []string{`"a\"b"`, `"c"`},
		},
		{
			name:     "double quoted escape is decoded and re-quoted",
			typeText: `"a\"b" | "c"`,
			want:     []string{`"a\"b"`, `"c"`},
		},
		{
			name:     "unicode escapes decode the same regardless of quote style",
			typeText: `'A' | '\x42' | "C"`,
			want:     []string{`"A"`, `"B"`, `"C"`},
		},
		{
			name:     "non-ascii hex escape is decoded for display",
			typeText: `'\xE9' | 'x'`,
			want:     []string{`"é"`, `"x"`},
		},
		{
			name:     "surrogate pair decodes to one code point",
			typeText: `'😀' | 'x'`,
			want:     []string{`"😀"`, `"x"`},
		},
		{
			name:     "non union",
			typeText: "string",
			want:     nil,
		},
		{
			name:     "union without literal values",
			typeText: "null | undefined | SomeType",
			want:     nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			attr := &manifest.Attribute{
				Type: &manifest.Type{Text: tc.typeText},
			}
			assert.Equal(t, tc.want, attr.DisplayEnumValues())
		})
	}
}

// Inline assertion: the contract is a single error string, so a golden would
// duplicate the implementation without adding coverage.
func TestAttributeValidationErrorUsesDisplayValues(t *testing.T) {
	attr := &manifest.Attribute{
		FullyQualified: manifest.FullyQualified{Name: "variant"},
		Type:           &manifest.Type{Text: "'primary' | 'secondary' | 'tertiary'"},
	}

	err := attr.GetValidationError("bogus")
	assert.EqualError(t, err,
		`invalid value "bogus" for attribute "variant". Valid values: ["primary" "secondary" "tertiary"]`)
}

// TestAttributeIsClosedEnum covers the predicate markup uses to decide whether
// a finite set of values can be validated. Inline table assertions are
// appropriate because each case is a pure predicate over a type string.
func TestAttributeIsClosedEnum(t *testing.T) {
	tests := []struct {
		typeText string
		want     bool
	}{
		{typeText: "'a' | 'b'", want: true},
		{typeText: "'sm' | null", want: true},
		{typeText: "'sm' | undefined", want: true},
		{typeText: "true | false", want: true},
		{typeText: "0 | 1", want: true},
		{typeText: "'a' | string", want: false},
		{typeText: "number | 'any'", want: false},
		{typeText: "string | undefined", want: false},
		{typeText: "Array<string> | null", want: false},
		{typeText: "A | B", want: false},
		{typeText: "string", want: false},
	}

	for _, tc := range tests {
		t.Run(tc.typeText, func(t *testing.T) {
			attr := &manifest.Attribute{Type: &manifest.Type{Text: tc.typeText}}
			assert.Equal(t, tc.want, attr.IsClosedEnum())
		})
	}
}

// A union with no literal members has no accepted set to report, so the error
// must not end with an empty list.
func TestAttributeValidationErrorOmitsEmptyDisplayValues(t *testing.T) {
	attr := &manifest.Attribute{
		FullyQualified: manifest.FullyQualified{Name: "items"},
		Type:           &manifest.Type{Text: "Array<string | number> | null"},
	}

	err := attr.GetValidationError("anything")
	assert.EqualError(t, err, `invalid value "anything" for attribute "items"`)
}
