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
package textutil_test

import (
	"testing"
	"unicode/utf8"

	"bennypowers.dev/cem/internal/textutil"
)

func TestUTF16ToByteOffset(t *testing.T) {
	tests := []struct {
		name         string
		text         string
		utf16Offset  uint32
		expectedByte uint
	}{
		{"ASCII only", "hello world", 6, 6},
		{"Emoji at start", "😀 hello", 3, 5},
		{"Emoji in middle", "hello 😀 world", 8, 10},
		{"Multiple emoji", "🔥🎉😀", 4, 8},
		{"Chinese characters", "你好世界", 2, 6},
		{"Mixed ASCII and emoji", "Hello 🌍 World", 8, 10},
		{"Spanish ñ", "Mañana", 3, 4},
		{"Zero offset", "😀 test", 0, 0},
		{"Offset beyond text", "test", 100, 4},
		{"Valid U+FFFD", "�x", 1, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := textutil.UTF16ToByteOffset(tt.text, tt.utf16Offset)
			if result != tt.expectedByte {
				t.Errorf("UTF16ToByteOffset(%q, %d) = %d, want %d",
					tt.text, tt.utf16Offset, result, tt.expectedByte)
			}
		})
	}
}

func TestByteOffsetToUTF16(t *testing.T) {
	tests := []struct {
		name          string
		text          string
		byteOffset    uint
		expectedUTF16 uint32
	}{
		{"ASCII only", "hello world", 6, 6},
		{"Emoji at start", "😀 hello", 5, 3},
		{"Chinese characters", "你好世界", 6, 2},
		{"Valid U+FFFD", "�x", 3, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := textutil.ByteOffsetToUTF16(tt.text, tt.byteOffset)
			if result != tt.expectedUTF16 {
				t.Errorf("ByteOffsetToUTF16(%q, %d) = %d, want %d",
					tt.text, tt.byteOffset, result, tt.expectedUTF16)
			}
		})
	}
}

func TestUTF16ToByteOffsetBytes(t *testing.T) {
	tests := []struct {
		name         string
		text         string
		utf16Offset  uint32
		expectedByte uint
	}{
		{"ASCII only", "hello world", 6, 6},
		{"Emoji at start", "😀 hello", 3, 5},
		{"Chinese characters", "你好世界", 2, 6},
		{"Zero offset", "😀 test", 0, 0},
		{"Offset beyond text", "test", 100, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := textutil.UTF16ToByteOffsetBytes([]byte(tt.text), tt.utf16Offset)
			if result != tt.expectedByte {
				t.Errorf("UTF16ToByteOffsetBytes(%q, %d) = %d, want %d",
					tt.text, tt.utf16Offset, result, tt.expectedByte)
			}
		})
	}
}

func TestByteOffsetToUTF16Bytes(t *testing.T) {
	tests := []struct {
		name          string
		text          string
		byteOffset    uint
		expectedUTF16 uint32
	}{
		{"ASCII only", "hello world", 6, 6},
		{"Emoji at start", "😀 hello", 5, 3},
		{"Chinese characters", "你好世界", 6, 2},
		{"Valid U+FFFD", "�x", 3, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := textutil.ByteOffsetToUTF16Bytes([]byte(tt.text), tt.byteOffset)
			if result != tt.expectedUTF16 {
				t.Errorf("ByteOffsetToUTF16Bytes(%q, %d) = %d, want %d",
					tt.text, tt.byteOffset, result, tt.expectedUTF16)
			}
		})
	}
}

func TestRoundTrip(t *testing.T) {
	texts := []string{
		"hello world",
		"😀 hello 🌍 world",
		"你好世界 hello",
		"Mañana 🔥🎉",
	}

	for _, text := range texts {
		t.Run(text, func(t *testing.T) {
			// Only test at rune-aligned byte offsets; mid-rune offsets
			// snap forward and can't round-trip.
			for byteOff := 0; byteOff < len(text); {
				utf16 := textutil.ByteOffsetToUTF16(text, uint(byteOff))
				back := textutil.UTF16ToByteOffset(text, utf16)
				if back != uint(byteOff) {
					t.Errorf("round-trip failed: byte %d -> utf16 %d -> byte %d",
						byteOff, utf16, back)
				}
				_, size := utf8.DecodeRuneInString(text[byteOff:])
				byteOff += size
			}
		})
	}
}
