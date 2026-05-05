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
package validate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSemverLessThan(t *testing.T) {
	tests := []struct {
		v1, v2 string
		want   bool
	}{
		{"1.0.0", "2.0.0", true},
		{"2.0.0", "1.0.0", false},
		{"1.0.0", "1.0.0", false},
		{"1.0.0", "1.1.0", true},
		{"1.0.0", "1.0.1", true},
		{"0.1.0", "1.0.0", true},
	}
	for _, tc := range tests {
		t.Run(tc.v1+" vs "+tc.v2, func(t *testing.T) {
			assert.Equal(t, tc.want, isSemverLessThan(tc.v1, tc.v2))
		})
	}
}

func TestFilterWarningsByConfig(t *testing.T) {
	warnings := []ValidationWarning{
		{ID: "missing-description", Category: "completeness"},
		{ID: "lifecycle-method", Category: "conventions"},
		{ID: "missing-summary", Category: "completeness"},
	}

	t.Run("no disabled rules", func(t *testing.T) {
		result := filterWarningsByConfig(warnings, nil)
		assert.Len(t, result, 3)
	})

	t.Run("disable by ID", func(t *testing.T) {
		result := filterWarningsByConfig(warnings, []string{"missing-description"})
		assert.Len(t, result, 2)
		assert.Equal(t, "lifecycle-method", result[0].ID)
	})

	t.Run("disable by category", func(t *testing.T) {
		result := filterWarningsByConfig(warnings, []string{"completeness"})
		assert.Len(t, result, 1)
		assert.Equal(t, "lifecycle-method", result[0].ID)
	})
}
