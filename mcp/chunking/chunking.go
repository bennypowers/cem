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
		PriorityKeywords: rfc2119Keywords,
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
		// Use proper UTF-8 decoding to handle multi-byte characters
		var beforeOK bool
		if idx == 0 {
			beforeOK = true
		} else {
			r, _ := utf8.DecodeLastRuneInString(upper[:idx])
			beforeOK = r == utf8.RuneError || !unicode.IsLetter(r)
		}

		afterIdx := idx + len(kwUpper)
		var afterOK bool
		if afterIdx >= len(upper) {
			afterOK = true
		} else {
			r, _ := utf8.DecodeRuneInString(upper[afterIdx:])
			afterOK = r == utf8.RuneError || !unicode.IsLetter(r)
		}

		if beforeOK && afterOK {
			return true
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

	// Build result: priority sentences first
	var result []string
	currentLen := 0

	// Add priority sentences
	for _, s := range prioritySentences {
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

	// Add first sentence if not already included and PreserveFirst is true
	if opts.PreserveFirst && firstSentence != "" {
		if !slices.Contains(result, firstSentence) {
			// Insert first sentence at the beginning
			// Space needed after first sentence if result has content
			spaceNeeded := 0
			if len(result) > 0 {
				spaceNeeded = 1
			}
			if currentLen+spaceNeeded+len(firstSentence) <= opts.MaxLength {
				result = append([]string{firstSentence}, result...)
				currentLen += spaceNeeded + len(firstSentence)
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

// truncateAtWord truncates text at a word boundary, then appends "...".
// The maxLen parameter specifies the maximum content length before the ellipsis.
// Note: The returned string may exceed maxLen by up to 3 bytes due to the "..." suffix.
// Callers requiring a strict byte limit should apply an additional hard cap.
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
