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

	lipgloss "charm.land/lipgloss/v2"
	"github.com/stretchr/testify/assert"

	"bennypowers.dev/cem/internal/tui"
)

func TestColorizeDuration(t *testing.T) {
	tests := []struct {
		name  string
		d     time.Duration
		style lipgloss.Style
	}{
		{"fast green", 50 * time.Millisecond, tui.DurationFastStyle},
		{"zero green", 0, tui.DurationFastStyle},
		{"boundary 99ms green", 99 * time.Millisecond, tui.DurationFastStyle},
		{"boundary 100ms yellow", 100 * time.Millisecond, tui.DurationMediumStyle},
		{"medium yellow", 300 * time.Millisecond, tui.DurationMediumStyle},
		{"boundary 499ms yellow", 499 * time.Millisecond, tui.DurationMediumStyle},
		{"boundary 500ms red", 500 * time.Millisecond, tui.DurationSlowStyle},
		{"slow red", 2 * time.Second, tui.DurationSlowStyle},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			style := tui.ColorizeDuration(tc.d)
			assert.Equal(t, tc.style, style)
		})
	}
}
