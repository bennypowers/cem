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
	"github.com/stretchr/testify/require"
)

// inline assertions: pure function with simple string equality
func TestSelectASCIILogo(t *testing.T) {
	tests := []struct {
		name   string
		width  int
		height int
		want   string
	}{
		{"narrow terminal", 40, 24, IC.LogoASCII40},
		{"medium wide enough tall enough", 70, 55, IC.LogoASCII60},
		{"medium wide but too short", 70, 40, IC.LogoASCII40},
		{"wide and tall", 120, 90, IC.LogoASCII100},
		{"wide but short", 120, 40, IC.LogoASCII40},
		{"wide medium height", 120, 55, IC.LogoASCII60},
		{"very narrow", 30, 24, IC.LogoASCII40},
		{"100 cols needs 84 rows for large", 100, 84, IC.LogoASCII100},
		{"100 cols with 49 rows gets medium", 100, 55, IC.LogoASCII60},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cmd.SelectASCIILogo(tt.width, tt.height)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLogoPNGEmbedded(t *testing.T) {
	require.GreaterOrEqual(t, len(IC.LogoPNG), 4, "logo PNG must be at least 4 bytes")
	assert.Equal(t, byte(0x89), IC.LogoPNG[0], "should start with PNG magic byte")
	assert.Equal(t, byte('P'), IC.LogoPNG[1])
	assert.Equal(t, byte('N'), IC.LogoPNG[2])
	assert.Equal(t, byte('G'), IC.LogoPNG[3])
}
