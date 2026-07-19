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
package lsp_test

import (
	"testing"

	"bennypowers.dev/cem/internal/textutil"
	"bennypowers.dev/cem/lsp/document"
	"go.lsp.dev/protocol"
)

// TestUTF16ToByteOffset tests conversion from LSP UTF-16 positions to UTF-8 byte offsets
func TestUTF16ToByteOffset(t *testing.T) {
	tests := []struct {
		name         string
		text         string
		utf16Offset  uint32
		expectedByte uint
		description  string
	}{
		{
			name:         "ASCII only",
			text:         "hello world",
			utf16Offset:  6,
			expectedByte: 6,
			description:  "ASCII characters: 1 byte = 1 UTF-16 code unit",
		},
		{
			name:         "Emoji at start",
			text:         "😀 hello",
			utf16Offset:  3, // After emoji (2 UTF-16) and space (1 UTF-16)
			expectedByte: 5, // Emoji is 4 bytes + 1 space
			description:  "Emoji (U+1F600) is 4 bytes but 2 UTF-16 code units (surrogate pair)",
		},
		{
			name:         "Emoji in middle",
			text:         "hello 😀 world",
			utf16Offset:  8,  // "hello " (6) + emoji (2)
			expectedByte: 10, // "hello " (6 bytes) + emoji (4 bytes)
			description:  "Emoji in the middle of text",
		},
		{
			name:         "Multiple emoji",
			text:         "🔥🎉😀",
			utf16Offset:  4, // First two emoji (2+2 UTF-16)
			expectedByte: 8, // First two emoji (4+4 bytes)
			description:  "Multiple consecutive emoji",
		},
		{
			name:         "Chinese characters",
			text:         "你好世界",
			utf16Offset:  2, // First two characters
			expectedByte: 6, // Each Chinese char is 3 bytes in UTF-8
			description:  "Chinese characters (3 bytes each, 1 UTF-16 code unit each)",
		},
		{
			name:         "Mixed ASCII and emoji",
			text:         "Hello 🌍 World",
			utf16Offset:  8,  // "Hello " (6) + emoji (2)
			expectedByte: 10, // "Hello " (6 bytes) + emoji (4 bytes)
			description:  "Mixed ASCII and emoji",
		},
		{
			name:         "Spanish characters",
			text:         "Mañana",
			utf16Offset:  3, // "Mañ"
			expectedByte: 4, // "Ma" (2 bytes) + "ñ" (2 bytes)
			description:  "ñ is 2 bytes but 1 UTF-16 code unit",
		},
		{
			name:         "Zero offset",
			text:         "😀 test",
			utf16Offset:  0,
			expectedByte: 0,
			description:  "Zero offset should return zero",
		},
		{
			name:         "Offset beyond text",
			text:         "test",
			utf16Offset:  100,
			expectedByte: 4, // Should clamp to end of text
			description:  "Offset beyond text length",
		},
		{
			name:         "Replacement char U+FFFD (valid)",
			text:         "�x", // U+FFFD followed by ASCII
			utf16Offset:  1,    // after U+FFFD (1 UTF-16 unit)
			expectedByte: 3,    // U+FFFD is 3 bytes in UTF-8
			description:  "Valid U+FFFD should count as one UTF-16 unit, three bytes",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := textutil.UTF16ToByteOffset(tt.text, tt.utf16Offset)
			if result != tt.expectedByte {
				t.Errorf("%s: UTF16ToByteOffset(%q, %d) = %d, want %d",
					tt.description, tt.text, tt.utf16Offset, result, tt.expectedByte)
			}
		})
	}
}

// TestByteOffsetToUTF16 tests conversion from UTF-8 byte offsets to LSP UTF-16 positions
func TestByteOffsetToUTF16(t *testing.T) {
	tests := []struct {
		name          string
		text          string
		byteOffset    uint
		expectedUTF16 uint32
		description   string
	}{
		{
			name:          "ASCII only",
			text:          "hello world",
			byteOffset:    6,
			expectedUTF16: 6,
			description:   "ASCII characters: 1 byte = 1 UTF-16 code unit",
		},
		{
			name:          "Emoji at start",
			text:          "😀 hello",
			byteOffset:    5, // After emoji (4 bytes) and space (1 byte)
			expectedUTF16: 3, // Emoji (2 UTF-16) + space (1 UTF-16)
			description:   "Emoji (U+1F600) is 4 bytes but 2 UTF-16 code units",
		},
		{
			name:          "Chinese characters",
			text:          "你好世界",
			byteOffset:    6, // First two characters (3 bytes each)
			expectedUTF16: 2, // Each char is 1 UTF-16 code unit
			description:   "Chinese characters (3 bytes each, 1 UTF-16 code unit each)",
		},
		{
			name:          "Replacement char U+FFFD (valid)",
			text:          "�x",
			byteOffset:    3, // after U+FFFD's 3 bytes
			expectedUTF16: 1, // one UTF-16 code unit
			description:   "Valid U+FFFD should map 3 bytes -> 1 UTF-16 unit",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := textutil.ByteOffsetToUTF16(tt.text, tt.byteOffset)
			if result != tt.expectedUTF16 {
				t.Errorf("%s: ByteOffsetToUTF16(%q, %d) = %d, want %d",
					tt.description, tt.text, tt.byteOffset, result, tt.expectedUTF16)
			}
		})
	}
}

// TestIncrementalParseWithUTF16 tests incremental parsing with UTF-16 characters
// This is an integration test that exercises the position conversion functions
func TestIncrementalParseWithUTF16(t *testing.T) {
	tests := []struct {
		name        string
		initial     string
		change      protocol.TextDocumentContentChangeEvent // interface
		expected    string
		description string
	}{
		{
			name:    "Insert emoji in ASCII text",
			initial: "Hello World",
			change: &protocol.TextDocumentContentChangePartial{
				Range: protocol.Range{
					Start: protocol.Position{Line: 0, Character: 6}, // After "Hello "
					End:   protocol.Position{Line: 0, Character: 6},
				},
				Text: "😀 ",
			},
			expected:    "Hello 😀 World",
			description: "Inserting emoji should handle UTF-16 positions correctly",
		},
		{
			name:    "Replace ASCII with emoji",
			initial: "Hello XX World",
			change: &protocol.TextDocumentContentChangePartial{
				Range: protocol.Range{
					Start: protocol.Position{Line: 0, Character: 6}, // "XX" starts at position 6
					End:   protocol.Position{Line: 0, Character: 8}, // "XX" ends at position 8
				},
				Text: "😀",
			},
			expected:    "Hello 😀 World",
			description: "Replacing ASCII with emoji should handle UTF-16 correctly",
		},
		{
			name:    "Insert Chinese characters",
			initial: "Hello World",
			change: &protocol.TextDocumentContentChangePartial{
				Range: protocol.Range{
					Start: protocol.Position{Line: 0, Character: 6},
					End:   protocol.Position{Line: 0, Character: 11},
				},
				Text: "你好",
			},
			expected:    "Hello 你好",
			description: "Chinese characters should be positioned correctly",
		},
		{
			name:    "Multi-line with emoji",
			initial: "Line 1\nLine 2",
			change: &protocol.TextDocumentContentChangePartial{
				Range: protocol.Range{
					Start: protocol.Position{Line: 1, Character: 5}, // After "Line "
					End:   protocol.Position{Line: 1, Character: 6}, // Replace "2"
				},
				Text: "😀",
			},
			expected:    "Line 1\nLine 😀",
			description: "Multi-line edits with emoji should work",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dm, err := document.NewDocumentManager()
			if err != nil {
				t.Fatalf("Failed to create DocumentManager: %v", err)
			}
			defer dm.Close()

			uri := "file:///test.html"
			dm.OpenDocument(uri, tt.initial, 1)

			changes := []protocol.TextDocumentContentChangeEvent{tt.change}
			updated := dm.UpdateDocumentWithChanges(uri, tt.expected, 2, changes)
			if updated == nil {
				t.Fatal("UpdateDocumentWithChanges returned nil")
			}

			content, err := updated.Content()
			if err != nil {
				t.Fatalf("Failed to get content: %v", err)
			}
			if content != tt.expected {
				t.Errorf("%s:\n  got:  %q\n  want: %q", tt.description, content, tt.expected)
			}

			if tree := updated.Tree(); tree != nil {
				root := tree.RootNode()
				if root == nil {
					t.Error("Tree root is nil after incremental parse")
				} else if root.EndByte() < uint(len(tt.expected)) {
					t.Errorf("Tree doesn't span content: tree end %d < content length %d",
						root.EndByte(), len(tt.expected))
				}
			}
		})
	}
}

// TestIncrementalParseAfterNonASCII verifies that edits positioned after
// non-ASCII content produce correct byte offsets. A wrong UTF-16-to-byte
// conversion would corrupt the tree or produce incorrect content.
func TestIncrementalParseAfterNonASCII(t *testing.T) {
	tests := []struct {
		name     string
		initial  string
		change   *protocol.TextDocumentContentChangePartial
		expected string
	}{
		{
			name:    "Edit after emoji on same line",
			initial: `<div>😀 old</div>`,
			change: &protocol.TextDocumentContentChangePartial{
				Range: protocol.Range{
					Start: protocol.Position{Line: 0, Character: 8}, // after "😀 " (5+2+1 = 8 UTF-16)
					End:   protocol.Position{Line: 0, Character: 11},
				},
				Text: "new",
			},
			expected: `<div>😀 new</div>`,
		},
		{
			name:    "Edit after Hebrew text",
			initial: `<div>שלום old</div>`,
			change: &protocol.TextDocumentContentChangePartial{
				Range: protocol.Range{
					Start: protocol.Position{Line: 0, Character: 10}, // after "שלום " (4+1+5 = 10 UTF-16)
					End:   protocol.Position{Line: 0, Character: 13},
				},
				Text: "new",
			},
			expected: `<div>שלום new</div>`,
		},
		{
			name:    "Edit after CJK on second line",
			initial: "<p>你好世界</p>\n<span>old</span>",
			change: &protocol.TextDocumentContentChangePartial{
				Range: protocol.Range{
					Start: protocol.Position{Line: 1, Character: 6},
					End:   protocol.Position{Line: 1, Character: 9},
				},
				Text: "new",
			},
			expected: "<p>你好世界</p>\n<span>new</span>",
		},
		{
			name:    "Edit after surrogate pair emoji",
			initial: `<b>🔥🎉 old</b>`,
			change: &protocol.TextDocumentContentChangePartial{
				Range: protocol.Range{
					Start: protocol.Position{Line: 0, Character: 8}, // after "🔥🎉 " (3+2+2+1 = 8 UTF-16)
					End:   protocol.Position{Line: 0, Character: 11},
				},
				Text: "new",
			},
			expected: `<b>🔥🎉 new</b>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dm, err := document.NewDocumentManager()
			if err != nil {
				t.Fatalf("Failed to create DocumentManager: %v", err)
			}
			defer dm.Close()

			uri := "file:///test-nonascii.html"
			doc := dm.OpenDocument(uri, tt.initial, 1)
			if doc == nil {
				t.Fatal("OpenDocument returned nil")
			}

			changes := []protocol.TextDocumentContentChangeEvent{tt.change}
			updated := dm.UpdateDocumentWithChanges(uri, tt.expected, 2, changes)
			if updated == nil {
				t.Fatal("UpdateDocumentWithChanges returned nil")
			}

			content, err := updated.Content()
			if err != nil {
				t.Fatalf("Failed to get content: %v", err)
			}
			if content != tt.expected {
				t.Errorf("Content mismatch:\n  got:  %q\n  want: %q", content, tt.expected)
			}

			if tree := updated.Tree(); tree != nil {
				root := tree.RootNode()
				if root == nil {
					t.Error("Tree root is nil after incremental parse")
				} else if root.EndByte() < uint(len(tt.expected)) {
					t.Errorf("Tree doesn't span content: tree end %d < content length %d",
						root.EndByte(), len(tt.expected))
				}
			}
		})
	}
}
