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
package chunking

import (
	"strings"
	"testing"
)

func TestSplitSentences(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: nil,
		},
		{
			name:     "single sentence",
			input:    "This is a sentence.",
			expected: []string{"This is a sentence."},
		},
		{
			name:     "multiple sentences",
			input:    "First sentence. Second sentence. Third sentence.",
			expected: []string{"First sentence.", "Second sentence.", "Third sentence."},
		},
		{
			name:     "sentences with exclamation",
			input:    "Hello! How are you? I am fine.",
			expected: []string{"Hello!", "How are you?", "I am fine."},
		},
		{
			name:     "no ending punctuation",
			input:    "This has no ending",
			expected: []string{"This has no ending"},
		},
		{
			name:     "trailing text after last sentence",
			input:    "First sentence. Some trailing text",
			expected: []string{"First sentence.", "Some trailing text"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := splitSentences(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("got %d sentences, want %d: %v", len(result), len(tt.expected), result)
				return
			}
			for i, s := range result {
				if s != tt.expected[i] {
					t.Errorf("sentence %d: got %q, want %q", i, s, tt.expected[i])
				}
			}
		})
	}
}

func TestContainsKeyword(t *testing.T) {
	tests := []struct {
		name     string
		sentence string
		keywords []string
		expected bool
	}{
		{
			name:     "contains MUST",
			sentence: "You MUST provide a value.",
			keywords: RFC2119Keywords,
			expected: true,
		},
		{
			name:     "contains SHOULD lowercase",
			sentence: "You should provide a value.",
			keywords: RFC2119Keywords,
			expected: true,
		},
		{
			name:     "contains MUST NOT",
			sentence: "You MUST NOT use this deprecated API.",
			keywords: RFC2119Keywords,
			expected: true,
		},
		{
			name:     "no keywords",
			sentence: "This is a regular sentence.",
			keywords: RFC2119Keywords,
			expected: false,
		},
		{
			name:     "keyword as part of word",
			sentence: "This is customary behavior.",
			keywords: []string{"MUST"},
			expected: false,
		},
		{
			name:     "keyword at start",
			sentence: "MUST be provided.",
			keywords: RFC2119Keywords,
			expected: true,
		},
		{
			name:     "keyword at end",
			sentence: "This is OPTIONAL",
			keywords: RFC2119Keywords,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := containsKeyword(tt.sentence, tt.keywords)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestChunk(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		opts     Options
		validate func(t *testing.T, result string)
	}{
		{
			name:  "empty string",
			input: "",
			opts:  DefaultOptions(),
			validate: func(t *testing.T, result string) {
				if result != "" {
					t.Errorf("got %q, want empty string", result)
				}
			},
		},
		{
			name:  "short text unchanged",
			input: "This is short.",
			opts:  DefaultOptions(),
			validate: func(t *testing.T, result string) {
				if result != "This is short." {
					t.Errorf("got %q, want %q", result, "This is short.")
				}
			},
		},
		{
			name:  "prioritizes RFC 2119 keywords",
			input: "This is an intro. You MUST follow this rule. Some filler. Another filler. More filler here.",
			opts:  Options{MaxLength: 80, PreserveFirst: true},
			validate: func(t *testing.T, result string) {
				if !strings.Contains(result, "MUST") {
					t.Errorf("result should contain MUST keyword: %q", result)
				}
				if !strings.Contains(result, "intro") {
					t.Errorf("result should contain first sentence: %q", result)
				}
			},
		},
		{
			name:  "preserves first sentence",
			input: "Important intro. Some other content. More content here. Even more content.",
			opts:  Options{MaxLength: 50, PreserveFirst: true},
			validate: func(t *testing.T, result string) {
				if !strings.Contains(result, "Important intro") {
					t.Errorf("result should contain first sentence: %q", result)
				}
			},
		},
		{
			name:  "respects max length as soft limit",
			input: "First. Second. Third. Fourth. Fifth. Sixth. Seventh. Eighth.",
			opts:  Options{MaxLength: 30, PreserveFirst: true},
			validate: func(t *testing.T, result string) {
				// Should include at least the first sentence, respecting sentence boundaries
				if !strings.Contains(result, "First") {
					t.Errorf("result should contain first sentence: %q", result)
				}
			},
		},
		{
			name:  "RFC 2119 over first sentence when both fit",
			input: "Intro here. You SHOULD do this. Regular sentence.",
			opts:  Options{MaxLength: 100, PreserveFirst: true},
			validate: func(t *testing.T, result string) {
				if !strings.Contains(result, "SHOULD") {
					t.Errorf("result should contain SHOULD: %q", result)
				}
				if !strings.Contains(result, "Intro") {
					t.Errorf("result should contain intro: %q", result)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Chunk(tt.input, tt.opts)
			tt.validate(t, result)
		})
	}
}

func TestExtractPriorityContent(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		keywords []string
		expected int // number of priority sentences
	}{
		{
			name:     "no priority content",
			input:    "Regular sentence. Another one.",
			keywords: RFC2119Keywords,
			expected: 0,
		},
		{
			name:     "one priority sentence",
			input:    "Regular sentence. You MUST do this. Another regular.",
			keywords: RFC2119Keywords,
			expected: 1,
		},
		{
			name:     "multiple priority sentences",
			input:    "You MUST do this. Regular. You SHOULD also do that. You MAY do this.",
			keywords: RFC2119Keywords,
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractPriorityContent(tt.input, tt.keywords)
			if len(result) != tt.expected {
				t.Errorf("got %d priority sentences, want %d: %v", len(result), tt.expected, result)
			}
		})
	}
}

func TestTruncateAtWord(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		maxLen   int
		expected string
	}{
		{
			name:     "short text unchanged",
			input:    "Short",
			maxLen:   10,
			expected: "Short",
		},
		{
			name:     "truncates at word boundary",
			input:    "This is a longer sentence that needs truncation",
			maxLen:   20,
			expected: "This is a longer...",
		},
		{
			name:     "no space found",
			input:    "Verylongwordwithoutspaces",
			maxLen:   10,
			expected: "Verylongwo...",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := truncateAtWord(tt.input, tt.maxLen)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}
