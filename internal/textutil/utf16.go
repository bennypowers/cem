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
package textutil

import (
	"unicode/utf16"
	"unicode/utf8"
)

// UTF16ToByteOffset converts a UTF-16 code unit offset to a UTF-8 byte offset.
func UTF16ToByteOffset(text string, utf16Offset uint32) uint {
	var byteOffset uint
	var utf16Count uint32

	for byteOffset < uint(len(text)) && utf16Count < utf16Offset {
		r, size := utf8.DecodeRuneInString(text[byteOffset:])
		if r == utf8.RuneError && size == 1 {
			byteOffset++
			utf16Count++
			continue
		}

		byteOffset += uint(size)
		utf16Count += uint32(utf16.RuneLen(r))
	}

	return byteOffset
}

// ByteOffsetToUTF16 converts a UTF-8 byte offset to a UTF-16 code unit offset.
func ByteOffsetToUTF16(text string, byteOffset uint) uint32 {
	var currentByte uint
	var utf16Count uint32

	for currentByte < byteOffset && currentByte < uint(len(text)) {
		r, size := utf8.DecodeRuneInString(text[currentByte:])
		if r == utf8.RuneError && size == 1 {
			currentByte++
			utf16Count++
			continue
		}

		currentByte += uint(size)
		utf16Count += uint32(utf16.RuneLen(r))
	}

	return utf16Count
}

// UTF16ToByteOffsetBytes is like UTF16ToByteOffset but operates on []byte.
func UTF16ToByteOffsetBytes(text []byte, utf16Offset uint32) uint {
	var byteOffset uint
	var utf16Count uint32

	for byteOffset < uint(len(text)) && utf16Count < utf16Offset {
		r, size := utf8.DecodeRune(text[byteOffset:])
		if r == utf8.RuneError && size == 1 {
			byteOffset++
			utf16Count++
			continue
		}

		byteOffset += uint(size)
		utf16Count += uint32(utf16.RuneLen(r))
	}

	return byteOffset
}

// ByteOffsetToUTF16Bytes is like ByteOffsetToUTF16 but operates on []byte.
func ByteOffsetToUTF16Bytes(text []byte, byteOffset uint) uint32 {
	var currentByte uint
	var utf16Count uint32

	for currentByte < byteOffset && currentByte < uint(len(text)) {
		r, size := utf8.DecodeRune(text[currentByte:])
		if r == utf8.RuneError && size == 1 {
			currentByte++
			utf16Count++
			continue
		}

		currentByte += uint(size)
		utf16Count += uint32(utf16.RuneLen(r))
	}

	return utf16Count
}
