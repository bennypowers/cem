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
	"time"

	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
)

// Inline: pure function, table-driven

func TestColorizeDuration(t *testing.T) {
	tests := []struct {
		name  string
		d     time.Duration
		color pterm.Color
	}{
		{"fast green", 50 * time.Millisecond, pterm.FgGreen},
		{"zero green", 0, pterm.FgGreen},
		{"boundary 99ms green", 99 * time.Millisecond, pterm.FgGreen},
		{"boundary 100ms yellow", 100 * time.Millisecond, pterm.FgYellow},
		{"medium yellow", 300 * time.Millisecond, pterm.FgYellow},
		{"boundary 499ms yellow", 499 * time.Millisecond, pterm.FgYellow},
		{"boundary 500ms red", 500 * time.Millisecond, pterm.FgRed},
		{"slow red", 2 * time.Second, pterm.FgRed},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			style := ColorizeDuration(tc.d)
			expected := pterm.NewStyle(tc.color)
			assert.Equal(t, expected, style)
		})
	}
}
