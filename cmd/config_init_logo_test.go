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
package cmd_test

import (
	"testing"

	"bennypowers.dev/cem/cmd"
	IC "bennypowers.dev/cem/internal/config"
	"github.com/stretchr/testify/assert"
)

// inline assertions: pure function with simple string equality
func TestSelectASCIILogo(t *testing.T) {
	tests := []struct {
		name  string
		width int
		want  string
	}{
		{"narrow terminal", 40, IC.LogoASCII40},
		{"medium terminal", 70, IC.LogoASCII60},
		{"wide terminal", 120, IC.LogoASCII100},
		{"very narrow", 30, IC.LogoASCII40},
		{"boundary 59", 59, IC.LogoASCII40},
		{"boundary 60", 60, IC.LogoASCII60},
		{"boundary 99", 99, IC.LogoASCII60},
		{"boundary 100", 100, IC.LogoASCII100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cmd.SelectASCIILogo(tt.width)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLogoPNGEmbedded(t *testing.T) {
	assert.NotEmpty(t, IC.LogoPNG, "logo PNG should be embedded")
	assert.Equal(t, byte(0x89), IC.LogoPNG[0], "should start with PNG magic byte")
	assert.Equal(t, byte('P'), IC.LogoPNG[1])
	assert.Equal(t, byte('N'), IC.LogoPNG[2])
	assert.Equal(t, byte('G'), IC.LogoPNG[3])
}
