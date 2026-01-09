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

// Package chunking provides semantic-aware text chunking for MCP context management.
// It prioritizes RFC 2119 keywords and preserves sentence boundaries.
package chunking

import (
	"regexp"
	"slices"
	"strings"
	"unicode"
)

// DefaultMaxLength is the default maximum length for chunked descriptions.
const DefaultMaxLength = 2000

// RFC2119Keywords are the standard RFC 2119 requirement level keywords.
var RFC2119Keywords = []string{
	"MUST", "MUST NOT",
	"REQUIRED",
	"SHALL", "SHALL NOT",
	"SHOULD", "SHOULD NOT",
	"RECOMMENDED", "NOT RECOMMENDED",
	"MAY",
	"OPTIONAL",
	"AVOID",
}

// Options configures semantic chunking behavior.
type Options struct {
	// MaxLength is the target maximum length (soft limit for sentence boundaries).
	// If zero, DefaultMaxLength is used.
	MaxLength int

	// PriorityKeywords are keywords to prioritize. If nil, RFC2119Keywords is used.
	PriorityKeywords []string

	// PreserveFirst ensures the first sentence is always included.
	// Default is true.
	PreserveFirst bool
}

// DefaultOptions returns the default chunking options.
func DefaultOptions() Options {
	return Options{
		MaxLength:        DefaultMaxLength,
		PriorityKeywords: RFC2119Keywords,
		PreserveFirst:    true,
	}
}

// sentenceEndRegex matches sentence-ending punctuation followed by space or end of string.
var sentenceEndRegex = regexp.MustCompile(`[.!?](?:\s+|$)`)

// splitSentences splits text into sentences, preserving the ending punctuation.
func splitSentences(text string) []string {
	text = strings.TrimSpace(text)
	if text == "" {
		return nil
	}

	var sentences []string
	indices := sentenceEndRegex.FindAllStringIndex(text, -1)

	if len(indices) == 0 {
		// No sentence-ending punctuation found; treat entire text as one sentence
		return []string{text}
	}

	start := 0
	for _, idx := range indices {
		end := idx[1]
		sentence := strings.TrimSpace(text[start:end])
		if sentence != "" {
			sentences = append(sentences, sentence)
		}
		start = end
	}

	// Handle any remaining text after the last sentence-ending punctuation
	if start < len(text) {
		remaining := strings.TrimSpace(text[start:])
		if remaining != "" {
			sentences = append(sentences, remaining)
		}
	}

	return sentences
}

// containsKeyword checks if a sentence contains any of the given keywords.
// Keywords are matched case-insensitively as whole words.
func containsKeyword(sentence string, keywords []string) bool {
	upper := strings.ToUpper(sentence)
	for _, kw := range keywords {
		kwUpper := strings.ToUpper(kw)
		// Check for whole word match
		idx := strings.Index(upper, kwUpper)
		if idx == -1 {
			continue
		}
		// Verify it's a whole word (not part of a larger word)
		beforeOK := idx == 0 || !unicode.IsLetter(rune(upper[idx-1]))
		afterIdx := idx + len(kwUpper)
		afterOK := afterIdx >= len(upper) || !unicode.IsLetter(rune(upper[afterIdx]))
		if beforeOK && afterOK {
			return true
		}
	}
	return false
}

// Chunk returns semantically-chunked text respecting sentence boundaries.
// Priority order: RFC 2119 keywords > first sentence > remaining sentences.
func Chunk(text string, opts Options) string {
	if text == "" {
		return ""
	}

	// Apply defaults
	if opts.MaxLength <= 0 {
		opts.MaxLength = DefaultMaxLength
	}
	if opts.PriorityKeywords == nil {
		opts.PriorityKeywords = RFC2119Keywords
	}

	// If text is already short enough, return as-is
	if len(text) <= opts.MaxLength {
		return text
	}

	sentences := splitSentences(text)
	if len(sentences) == 0 {
		// Fallback: truncate at word boundary
		return truncateAtWord(text, opts.MaxLength)
	}

	// Categorize sentences
	var prioritySentences []string
	var firstSentence string
	var otherSentences []string

	for i, s := range sentences {
		if i == 0 {
			firstSentence = s
		}
		if containsKeyword(s, opts.PriorityKeywords) {
			prioritySentences = append(prioritySentences, s)
		} else if i > 0 {
			otherSentences = append(otherSentences, s)
		}
	}

	// Build result: priority sentences first
	var result []string
	currentLen := 0

	// Add priority sentences
	for _, s := range prioritySentences {
		if currentLen+len(s)+1 > opts.MaxLength && len(result) > 0 {
			break
		}
		result = append(result, s)
		currentLen += len(s) + 1 // +1 for space
	}

	// Add first sentence if not already included and PreserveFirst is true
	if opts.PreserveFirst && firstSentence != "" {
		if !slices.Contains(result, firstSentence) {
			// Insert first sentence at the beginning
			if currentLen+len(firstSentence)+1 <= opts.MaxLength {
				result = append([]string{firstSentence}, result...)
				currentLen += len(firstSentence) + 1
			} else if len(result) == 0 {
				// No content yet and first sentence is too long - truncate it
				truncated := truncateAtWord(firstSentence, opts.MaxLength)
				result = append(result, truncated)
				currentLen = len(truncated)
			}
		}
	}

	// Fill remaining space with other sentences
	for _, s := range otherSentences {
		if currentLen+len(s)+1 > opts.MaxLength {
			break
		}
		result = append(result, s)
		currentLen += len(s) + 1
	}

	return strings.Join(result, " ")
}

// truncateAtWord truncates text at a word boundary, not exceeding maxLen.
func truncateAtWord(text string, maxLen int) string {
	if len(text) <= maxLen {
		return text
	}

	// Find last space before maxLen
	truncated := text[:maxLen]
	lastSpace := strings.LastIndex(truncated, " ")
	if lastSpace > 0 {
		truncated = truncated[:lastSpace]
	}

	return strings.TrimSpace(truncated) + "..."
}

// ExtractPriorityContent extracts sentences containing priority keywords.
func ExtractPriorityContent(text string, keywords []string) []string {
	if keywords == nil {
		keywords = RFC2119Keywords
	}

	sentences := splitSentences(text)
	var priority []string

	for _, s := range sentences {
		if containsKeyword(s, keywords) {
			priority = append(priority, s)
		}
	}

	return priority
}
