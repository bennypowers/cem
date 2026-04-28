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
package treesitter_test

import (
	"testing"

	"bennypowers.dev/cem/internal/treesitter"
)

func TestByteOffsetToPosition_UTF16(t *testing.T) {
	tests := []struct {
		name      string
		content   string
		offset    uint
		wantLine  uint32
		wantChar  uint32
	}{
		{"ASCII", "hello world", 6, 0, 6},
		{"emoji (4 bytes = 2 UTF-16 units)", "😀 hello", 4, 0, 2},
		{"after emoji and space", "😀 hello", 5, 0, 3},
		{"Chinese (3 bytes = 1 UTF-16 unit)", "你好world", 6, 0, 2},
		{"multiline ASCII", "line1\nline2\nline3", 14, 2, 2},
		{"multiline with emoji", "line1\n😀 hello", 10, 1, 2},
		{"offset at start", "hello", 0, 0, 0},
		{"offset at newline", "hello\nworld", 5, 0, 5},
		{"offset after newline", "hello\nworld", 6, 1, 0},
		{"empty content", "", 0, 0, 0},
		{"content ending with newline", "hello\n", 6, 1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pos := treesitter.ByteOffsetToPosition([]byte(tt.content), tt.offset)
			if pos.Line != tt.wantLine || pos.Character != tt.wantChar {
				t.Errorf("ByteOffsetToPosition(%q, %d) = {%d, %d}, want {%d, %d}",
					tt.content, tt.offset, pos.Line, pos.Character, tt.wantLine, tt.wantChar)
			}
		})
	}
}
