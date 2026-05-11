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

import "bytes"

// StripFrontmatter removes YAML frontmatter (delimited by "---") from content,
// returning only the content after the closing delimiter.
// If no frontmatter is present, returns the original content unchanged.
func StripFrontmatter(content []byte) []byte {
	trimmed := bytes.TrimLeft(content, " \t\r\n")

	var openLen int
	switch {
	case bytes.HasPrefix(trimmed, []byte("---\n")):
		openLen = 4
	case bytes.HasPrefix(trimmed, []byte("---\r\n")):
		openLen = 5
	default:
		return content
	}

	rest := trimmed[openLen:]

	// Check if rest itself starts with closing delimiter (empty frontmatter)
	if bytes.HasPrefix(rest, []byte("---")) {
		after := rest[3:]
		switch {
		case len(after) == 0:
			return after
		case after[0] == '\n':
			return after[1:]
		case after[0] == '\r' && len(after) > 1 && after[1] == '\n':
			return after[2:]
		}
	}

	for i := range len(rest) {
		if rest[i] != '\n' {
			continue
		}
		line := rest[i+1:]
		if !bytes.HasPrefix(line, []byte("---")) {
			continue
		}
		after := line[3:]
		switch {
		case len(after) == 0:
			return after
		case after[0] == '\n':
			return after[1:]
		case after[0] == '\r' && len(after) > 1 && after[1] == '\n':
			return after[2:]
		default:
			continue
		}
	}

	// No closing delimiter found: stray "---" line before HTML content.
	// Strip it so raw delimiters don't appear in rendered output.
	trimmedRest := bytes.TrimLeft(rest, " \t")
	if len(trimmedRest) > 0 && trimmedRest[0] == '<' {
		return rest
	}

	return content
}
