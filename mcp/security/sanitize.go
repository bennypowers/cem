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
)

// Maximum safe description length to prevent abuse and ensure good AI performance.
// This limit balances comprehensive documentation with processing efficiency.
// Descriptions longer than this are truncated with "..." appended.
// Can be customized via config file or --max-description-length flag.
const maxDescriptionLength = 2000

// Compiled regex for whitespace normalization to avoid recompiling on every call
var whitespaceRegex = regexp.MustCompile(`\s+`)

// SanitizeDescription sanitizes a description field for safe use in MCP responses
// Handles length limits and whitespace normalization
func SanitizeDescription(description string) string {
	return SanitizeDescriptionWithLength(description, maxDescriptionLength)
}

// SanitizeDescriptionWithLength sanitizes with a custom max length
func SanitizeDescriptionWithLength(description string, maxLength int) string {
	if description == "" {
		return ""
	}

	// Limit length to prevent abuse
	if len(description) > maxLength {
		description = description[:maxLength] + "..."
	}

	// Clean up excessive whitespace
	description = strings.TrimSpace(description)
	description = whitespaceRegex.ReplaceAllString(description, " ")

	return description
}
