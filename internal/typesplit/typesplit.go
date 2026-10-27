/*
Copyright © 2026 Benny Powers <web@bennypowers.com>

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

// Package typesplit splits TypeScript type strings without a parser or cgo.
//
// It is a pure-Go approximation of the tree-sitter based splitter in
// internal/tstype, sufficient for manifest consumers that need to inspect
// union types without linking tree-sitter. It mirrors tree-sitter's treatment
// of the common top-level type forms: unions, generics, arrays, object types,
// function types, conditional types, and type operators such as readonly.
package typesplit

import "strings"

// SplitTopLevelUnion returns the top-level union members of typeText.
//
// It returns nil for empty input and a single-element slice for non-union
// types. Whitespace around members is trimmed, and empty members (for
// example from a leading pipe) are dropped.
func SplitTopLevelUnion(typeText string) []string {
	typeText = strings.TrimSpace(typeText)
	if typeText == "" {
		return nil
	}

	typeText = strings.TrimSpace(stripComments(typeText))
	if typeText == "" {
		return nil
	}

	scan := scanType(typeText)
	if !scan.rootMayBeUnion(typeText) {
		return []string{typeText}
	}

	parts := splitAt(typeText, scan.pipes)
	if len(parts) == 0 {
		return []string{typeText}
	}
	return parts
}

// scanResult records the top-level structure of a type string.
type scanResult struct {
	pipes            []int // byte offsets of top-level '|'
	firstPipe        int   // first top-level pipe, -1 if none
	topLevelArrow    int   // offset of a top-level "=>", -1 if none
	topLevelQuestion bool  // top-level '?'
	topLevelExtends  bool  // top-level "extends" keyword
	topLevelComma    bool  // top-level ','
	balanced         bool  // all brackets and quotes closed
	trailingPipe     bool  // the type ends with a top-level '|'
}

// rootMayBeUnion reports whether tree-sitter would parse the text as a
// top-level union. Forms whose grammar node wraps the union (readonly, function
// return types, and conditional types) are not unions at the root.
func (s scanResult) rootMayBeUnion(typeText string) bool {
	if !s.balanced || s.topLevelComma || s.trailingPipe {
		return false
	}

	words := leadingWords(typeText, 2)
	first, second := words[0], words[1]
	if first == "readonly" {
		return false
	}
	isConstructorType := first == "new" || (first == "abstract" && second == "new")
	if !isConstructorType && s.topLevelArrow >= 0 &&
		(s.firstPipe < 0 || s.topLevelArrow < s.firstPipe) {
		// The root is a function type, so a union in its return type is
		// not a top-level union.
		return false
	}
	if s.topLevelExtends && s.topLevelQuestion {
		// The root is a conditional type.
		return false
	}
	return true
}

// scanType walks typeText and records top-level delimiters. It is not a
// TypeScript parser; it only tracks bracket depth and string literals.
func scanType(s string) scanResult {
	res := scanResult{
		firstPipe:     -1,
		topLevelArrow: -1,
		balanced:      true,
	}

	var parens, brackets, braces, angles int
	atTopLevel := func() bool {
		return parens == 0 && brackets == 0 && braces == 0 && angles == 0
	}

	i := 0
	for i < len(s) {
		c := s[i]
		switch {
		case c == '\'' || c == '"':
			end, ok := skipQuoted(s, i, c)
			if !ok {
				res.balanced = false
				return res
			}
			i = end

		case c == '`':
			end, ok := skipTemplate(s, i)
			if !ok {
				res.balanced = false
				return res
			}
			i = end

		case c == '/' && i+1 < len(s) && s[i+1] == '/':
			for i < len(s) && s[i] != '\n' {
				i++
			}

		case c == '/' && i+1 < len(s) && s[i+1] == '*':
			i += 2
			for i+1 < len(s) && (s[i] != '*' || s[i+1] != '/') {
				i++
			}
			if i+1 >= len(s) {
				res.balanced = false
				return res
			}
			i += 2

		case c == '(':
			parens++
			i++
		case c == ')':
			parens--
			if parens < 0 {
				res.balanced = false
				return res
			}
			i++
		case c == '[':
			brackets++
			i++
		case c == ']':
			brackets--
			if brackets < 0 {
				res.balanced = false
				return res
			}
			i++
		case c == '{':
			braces++
			i++
		case c == '}':
			braces--
			if braces < 0 {
				res.balanced = false
				return res
			}
			i++
		case c == '<':
			angles++
			i++
		case c == '>':
			if angles > 0 {
				angles--
			}
			i++
		case c == '=' && i+1 < len(s) && s[i+1] == '>':
			if atTopLevel() {
				res.topLevelArrow = i
			}
			i += 2
		case c == '|':
			if atTopLevel() {
				if res.firstPipe < 0 {
					res.firstPipe = i
				}
				res.pipes = append(res.pipes, i)
			}
			i++
		case c == ',':
			if atTopLevel() {
				res.topLevelComma = true
			}
			i++
		case c == '?':
			if atTopLevel() {
				res.topLevelQuestion = true
			}
			i++
		default:
			if atTopLevel() && isIdentStart(c) {
				end := i + 1
				for end < len(s) && isIdentPart(s[end]) {
					end++
				}
				if s[i:end] == "extends" {
					res.topLevelExtends = true
				}
				i = end
				continue
			}
			i++
		}
	}

	if parens != 0 || brackets != 0 || braces != 0 || angles != 0 {
		res.balanced = false
		return res
	}

	if last := lastNonSpace(s); last >= 0 && s[last] == '|' && len(res.pipes) > 0 &&
		res.pipes[len(res.pipes)-1] == last {
		res.trailingPipe = true
	}
	return res
}

// splitAt splits s at the given top-level pipe offsets, trimming members and
// dropping empty ones.
func splitAt(s string, pipes []int) []string {
	parts := make([]string, 0, len(pipes)+1)
	start := 0
	for _, pos := range pipes {
		if part := strings.TrimSpace(s[start:pos]); part != "" {
			parts = append(parts, part)
		}
		start = pos + 1
	}
	if part := strings.TrimSpace(s[start:]); part != "" {
		parts = append(parts, part)
	}
	return parts
}

// stripComments removes line and block comments while preserving string and
// template literal contents.
func stripComments(s string) string {
	var b strings.Builder
	b.Grow(len(s))

	i := 0
	for i < len(s) {
		switch {
		case s[i] == '\'' || s[i] == '"':
			end, ok := skipQuoted(s, i, s[i])
			if !ok {
				b.WriteString(s[i:])
				return b.String()
			}
			b.WriteString(s[i:end])
			i = end

		case s[i] == '`':
			end, ok := skipTemplate(s, i)
			if !ok {
				b.WriteString(s[i:])
				return b.String()
			}
			b.WriteString(s[i:end])
			i = end

		case s[i] == '/' && i+1 < len(s) && s[i+1] == '/':
			for i < len(s) && s[i] != '\n' {
				i++
			}

		case s[i] == '/' && i+1 < len(s) && s[i+1] == '*':
			i += 2
			for i+1 < len(s) && (s[i] != '*' || s[i+1] != '/') {
				i++
			}
			if i+1 >= len(s) {
				return b.String()
			}
			i += 2
			b.WriteByte(' ')

		default:
			b.WriteByte(s[i])
			i++
		}
	}
	return b.String()
}

// skipQuoted returns the offset just past the closing quote.
func skipQuoted(s string, start int, quote byte) (int, bool) {
	for i := start + 1; i < len(s); i++ {
		switch s[i] {
		case '\\':
			i++
		case quote:
			return i + 1, true
		}
	}
	return len(s), false
}

// skipTemplate returns the offset just past the closing backtick. It tracks
// ${...} substitutions, including nested template literals, quoted strings, and
// braces inside substitution type expressions.
func skipTemplate(s string, start int) (int, bool) {
	for i := start + 1; i < len(s); {
		switch {
		case s[i] == '\\':
			i += 2
		case s[i] == '`':
			return i + 1, true
		case s[i] == '$' && i+1 < len(s) && s[i+1] == '{':
			end, ok := skipTemplateSubstitution(s, i+2)
			if !ok {
				return len(s), false
			}
			i = end
		default:
			i++
		}
	}
	return len(s), false
}

// skipTemplateSubstitution returns the offset just past the '}' that closes a
// ${...} substitution. start is the offset immediately after "${".
func skipTemplateSubstitution(s string, start int) (int, bool) {
	depth := 1
	for i := start; i < len(s); {
		switch {
		case s[i] == '\'' || s[i] == '"':
			end, ok := skipQuoted(s, i, s[i])
			if !ok {
				return len(s), false
			}
			i = end
		case s[i] == '`':
			end, ok := skipTemplate(s, i)
			if !ok {
				return len(s), false
			}
			i = end
		case s[i] == '/' && i+1 < len(s) && s[i+1] == '/':
			for i < len(s) && s[i] != '\n' {
				i++
			}
		case s[i] == '/' && i+1 < len(s) && s[i+1] == '*':
			i += 2
			for i+1 < len(s) && (s[i] != '*' || s[i+1] != '/') {
				i++
			}
			if i+1 >= len(s) {
				return len(s), false
			}
			i += 2
		case s[i] == '{':
			depth++
			i++
		case s[i] == '}':
			depth--
			i++
			if depth == 0 {
				return i, true
			}
		default:
			i++
		}
	}
	return len(s), false
}

// leadingWords returns up to n leading identifier-like words.
func leadingWords(s string, n int) []string {
	words := make([]string, 0, n)
	i := 0
	for len(words) < n {
		for i < len(s) && isSpace(s[i]) {
			i++
		}
		if i >= len(s) || !isIdentStart(s[i]) {
			break
		}
		end := i + 1
		for end < len(s) && isIdentPart(s[end]) {
			end++
		}
		words = append(words, s[i:end])
		i = end
	}
	for len(words) < n {
		words = append(words, "")
	}
	return words
}

// lastNonSpace returns the index of the last non-whitespace byte, or -1.
func lastNonSpace(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if !isSpace(s[i]) {
			return i
		}
	}
	return -1
}

// isIdentStart reports whether c can start a TypeScript identifier.
func isIdentStart(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c == '_' || c == '$'
}

// isIdentPart reports whether c can continue a TypeScript identifier.
func isIdentPart(c byte) bool {
	return isIdentStart(c) || c >= '0' && c <= '9'
}

// isSpace reports whether c is ASCII whitespace.
func isSpace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\r' || c == '\n'
}
