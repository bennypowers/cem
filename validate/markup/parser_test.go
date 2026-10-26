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
package markup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// isCustomElementTag is a pure function; a table is clearer than a golden.
func TestIsCustomElementTag(t *testing.T) {
	tests := []struct {
		tag  string
		want bool
	}{
		{"my-element", true},
		{"a-b", true},
		{"a-", true},
		{"a--b", true},
		{"my-element-", true},
		{"my_ns-element", false},
		{"my-element_", false},
		{"1-element", false},
		{"-element", false},
		{"element", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.tag, func(t *testing.T) {
			assert.Equal(t, tt.want, isCustomElementTag(tt.tag))
		})
	}
}
