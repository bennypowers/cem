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
	"strings"
	"unicode"
	"unicode/utf8"
)

// DefaultMaxLength is the default maximum length for chunked descriptions.
const DefaultMaxLength = 2000

// rfc2119Keywords are the standard RFC 2119 requirement level keywords.
var rfc2119Keywords = []string{
	"MUST", "MUST NOT",
	"REQUIRED",
	"SHALL", "SHALL NOT",
	"SHOULD", "SHOULD NOT",
	"RECOMMENDED", "NOT RECOMMENDED",
	"MAY",
	"OPTIONAL",
	"AVOID",
}

// RFC2119Keywords returns a copy of the standard RFC 2119 keywords.
// Returns a copy to prevent accidental mutation of the shared slice.
func RFC2119Keywords() []string {
	return append([]string(nil), rfc2119Keywords...)
}

// Options configures semantic chunking behavior.
type Options struct {
	// MaxLength is the target maximum length (soft limit for sentence boundaries).
	// If zero, DefaultMaxLength is used.
	MaxLength int

	// PriorityKeywords are keywords to prioritize. If nil, RFC2119Keywords is used.
	PriorityKeywords []string

	// PreserveFirst ensures the first sentence is always included.
	// Set to true when using DefaultOptions(). Zero value (false) disables this behavior.
	PreserveFirst bool
}

// DefaultOptions returns the default chunking options.
func DefaultOptions() Options {
	return Options{
		MaxLength:        DefaultMaxLength,
		PriorityKeywords: RFC2119Keywords(),
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

// isWordChar returns true if the rune is considered part of a word
// (letter, digit, or underscore).
func isWordChar(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_'
}

// containsKeyword checks if a sentence contains any of the given keywords.
// Keywords are matched case-insensitively as whole words.
// A word boundary is defined as start/end of string or a non-word character
// (where word characters are letters, digits, and underscores).
func containsKeyword(sentence string, keywords []string) bool {
	upper := strings.ToUpper(sentence)
	for _, kw := range keywords {
		kwUpper := strings.ToUpper(kw)
		// Iterate over all occurrences of the keyword
		offset := 0
		for offset < len(upper) {
			idx := strings.Index(upper[offset:], kwUpper)
			if idx == -1 {
				break
			}
			// Adjust idx to be relative to the full string
			idx += offset

			// Verify it's a whole word (not part of a larger word)
			// Use proper UTF-8 decoding to handle multi-byte characters
			var beforeOK bool
			if idx == 0 {
				beforeOK = true
			} else {
				r, _ := utf8.DecodeLastRuneInString(upper[:idx])
				beforeOK = r == utf8.RuneError || !isWordChar(r)
			}

			afterIdx := idx + len(kwUpper)
			var afterOK bool
			if afterIdx >= len(upper) {
				afterOK = true
			} else {
				r, _ := utf8.DecodeRuneInString(upper[afterIdx:])
				afterOK = r == utf8.RuneError || !isWordChar(r)
			}

			if beforeOK && afterOK {
				return true
			}

			// Move past this occurrence to check for more
			offset = idx + len(kwUpper)
		}
	}
	return false
}

// Chunk returns semantically-chunked text respecting sentence boundaries.
// Priority order: RFC 2119 keywords > first sentence > remaining sentences.
//
// Note: If a priority sentence exceeds MaxLength, it will still be included
// to ensure meaningful content is retained. For strict length enforcement,
// use the hard cap in sanitize.go. When no priority keywords exist and
// PreserveFirst is false, may return empty string if all sentences exceed MaxLength.
func Chunk(text string, opts Options) string {
	if text == "" {
		return ""
	}

	// Apply defaults
	if opts.MaxLength <= 0 {
		opts.MaxLength = DefaultMaxLength
	}
	if opts.PriorityKeywords == nil {
		opts.PriorityKeywords = rfc2119Keywords
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

	// Build result
	var result []string
	currentLen := 0

	// When PreserveFirst is true, always start with the first sentence to ensure
	// it's included even when priority sentences would otherwise consume all space.
	firstSentenceIncluded := false
	if opts.PreserveFirst && firstSentence != "" {
		if len(firstSentence) <= opts.MaxLength {
			result = append(result, firstSentence)
			currentLen = len(firstSentence)
			firstSentenceIncluded = true
		} else {
			// First sentence is too long - truncate it
			truncated := truncateAtWord(firstSentence, opts.MaxLength)
			result = append(result, truncated)
			currentLen = len(truncated)
			firstSentenceIncluded = true
		}
	}

	// Add priority sentences that fit (skipping first sentence if already included)
	for _, s := range prioritySentences {
		if firstSentenceIncluded && s == firstSentence {
			continue // Already included as first sentence
		}
		// Account for space delimiter only when joining with existing sentences
		spaceNeeded := 0
		if len(result) > 0 {
			spaceNeeded = 1
		}
		if currentLen+spaceNeeded+len(s) > opts.MaxLength && len(result) > 0 {
			break
		}
		result = append(result, s)
		currentLen += spaceNeeded + len(s)
	}

	// If PreserveFirst was false and we have no content yet, try priority sentences
	if !opts.PreserveFirst && len(result) == 0 && len(prioritySentences) > 0 {
		s := prioritySentences[0]
		if len(s) <= opts.MaxLength {
			result = append(result, s)
			currentLen = len(s)
		} else {
			truncated := truncateAtWord(s, opts.MaxLength)
			result = append(result, truncated)
			currentLen = len(truncated)
		}
	}

	// Fill remaining space with other sentences
	for _, s := range otherSentences {
		// Need a space if we have existing content
		if len(result) == 0 {
			// Edge case: no priority, PreserveFirst false, all sentences too long
			// Add first other sentence with truncation if needed
			if len(s) <= opts.MaxLength {
				result = append(result, s)
				currentLen = len(s)
			} else {
				result = append(result, truncateAtWord(s, opts.MaxLength))
				break
			}
			continue
		}
		if currentLen+1+len(s) > opts.MaxLength {
			break
		}
		result = append(result, s)
		currentLen += 1 + len(s)
	}

	return strings.Join(result, " ")
}

// TruncateToByteLimit truncates text to fit within maxLen bytes,
// using rune-aware truncation to avoid splitting multi-byte UTF-8 characters.
// Does not append any suffix. Returns empty string if maxLen <= 0.
func TruncateToByteLimit(text string, maxLen int) string {
	if maxLen <= 0 {
		return ""
	}
	if len(text) <= maxLen {
		return text
	}
	runes := []rune(text)
	left, right := 0, len(runes)
	for left < right {
		mid := (left + right + 1) / 2
		if len(string(runes[:mid])) <= maxLen {
			left = mid
		} else {
			right = mid - 1
		}
	}
	if left > 0 {
		return string(runes[:left])
	}
	return ""
}

// truncateAtWord truncates text at a word boundary, then appends "...".
// The maxLen parameter specifies the maximum content length before the ellipsis.
// Note: The returned string may exceed maxLen by up to 3 bytes due to the "..." suffix.
// Callers requiring a strict byte limit should apply an additional hard cap.
// Uses rune-aware truncation to avoid splitting multi-byte UTF-8 characters.
func truncateAtWord(text string, maxLen int) string {
	if len(text) <= maxLen {
		return text
	}

	// Find a rune-safe prefix that fits in maxLen bytes
	truncated := TruncateToByteLimit(text, maxLen)

	// Find last space within the safe prefix for word boundary
	lastSpace := strings.LastIndex(truncated, " ")
	if lastSpace > 0 {
		truncated = truncated[:lastSpace]
	}

	return strings.TrimSpace(truncated) + "..."
}

// ExtractPriorityContent extracts sentences containing priority keywords.
func ExtractPriorityContent(text string, keywords []string) []string {
	if keywords == nil {
		keywords = rfc2119Keywords
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
