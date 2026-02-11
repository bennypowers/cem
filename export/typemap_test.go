/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

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
package export

import "testing"

func TestMapCEMType(t *testing.T) {
	tests := []struct {
		input, want string
	}{
		{"string", "string"},
		{"number", "number"},
		{"boolean", "boolean"},
		{"", "unknown"},
		{"  ", "unknown"},
		{"HTMLElement", "HTMLElement"},
		{"string | number", "string | number"},
		{"'primary' | 'secondary'", "'primary' | 'secondary'"},
		{"Array<string>", "Array<string>"},
	}
	for _, tt := range tests {
		if got := MapCEMType(tt.input); got != tt.want {
			t.Errorf("MapCEMType(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestIsBooleanType(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"boolean", true},
		{"string", false},
		{"", false},
		{"boolean | undefined", false},
		{"  boolean  ", true},
	}
	for _, tt := range tests {
		if got := IsBooleanType(tt.input); got != tt.want {
			t.Errorf("IsBooleanType(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
