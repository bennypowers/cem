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
package security

import (
	"regexp"
	"strings"
	"unicode/utf8"

	"bennypowers.dev/cem/mcp/chunking"
)

// Maximum safe description length to prevent abuse and ensure good AI performance.
// This limit balances comprehensive documentation with processing efficiency.
// Descriptions longer than this are semantically chunked to preserve sentence
// boundaries and prioritize RFC 2119 keywords, with a hard safety cap enforced.
// Can be customized via config file or --max-description-length flag.
const maxDescriptionLength = 2000

// Compiled regex for whitespace normalization to avoid recompiling on every call
var whitespaceRegex = regexp.MustCompile(`\s+`)

// SanitizeDescription sanitizes a description field for safe use in MCP responses
// Handles length limits and whitespace normalization
func SanitizeDescription(description string) string {
	return SanitizeDescriptionWithLength(description, maxDescriptionLength)
}

// SanitizeDescriptionWithLength sanitizes with a custom max length.
// Uses semantic chunking to preserve sentence boundaries and prioritize
// RFC 2119 keywords (MUST, SHOULD, etc.) when truncating long descriptions.
func SanitizeDescriptionWithLength(description string, maxLength int) string {
	if description == "" {
		return ""
	}

	// Clean up excessive whitespace first
	description = strings.TrimSpace(description)
	description = whitespaceRegex.ReplaceAllString(description, " ")

	// Use semantic chunking for length limiting
	// This preserves sentence boundaries and prioritizes important content
	if len(description) > maxLength {
		description = chunking.Chunk(description, chunking.Options{
			MaxLength:        maxLength,
			PriorityKeywords: chunking.RFC2119Keywords(),
			PreserveFirst:    true,
		})
		// Hard safety cap: ensure we never exceed maxLength bytes.
		// Chunking can exceed the limit in edge cases (e.g., long single sentence).
		if len(description) > maxLength {
			cut := maxLength
			for cut > 0 && !utf8.ValidString(description[:cut]) {
				cut--
			}
			description = description[:cut]
		}
	}

	return description
}
