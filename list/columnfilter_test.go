/*
Copyright © 2025 Benny Powers

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
package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveEmptyColumns(t *testing.T) {
	headers := []string{"Name", "Type", "Summary"}
	rows := [][]string{
		{"foo", "", "desc"},
		{"bar", "", "desc2"},
	}
	wantHeaders := []string{"Name", "Summary"}
	wantRows := [][]string{
		{"foo", "desc"},
		{"bar", "desc2"},
	}
	gotHeaders, gotRows := RemoveEmptyColumns(headers, rows)
	assert.Equal(t, wantHeaders, gotHeaders)
	assert.Equal(t, wantRows, gotRows)
}
