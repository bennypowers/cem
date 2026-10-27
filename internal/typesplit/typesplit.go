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

import (
	"slices"
	"strings"
)

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

	scan := scanType(typeText)
	if scan.trailingSemicolon && scan.lastSemicolon >= 0 {
		// A top-level ';' terminates the type expression; anything after it is
		// not part of the type.
		trimmed := strings.TrimSpace(typeText[:scan.lastSemicolon])
		if trimmed == "" {
			return []string{typeText}
		}
		typeText = trimmed
		scan = scanType(typeText)
	}
	if !scan.balanced || scan.topLevelComma || scan.trailingPipe {
		return []string{typeText}
	}

	pipes := scan.separators()
	if len(pipes) == 0 {
		return []string{typeText}
	}

	parts := splitAt(typeText, pipes)
	if len(parts) == 0 {
		return []string{typeText}
	}
	return parts
}

// scanResult records the top-level structure of a type string.
type scanResult struct {
	pipes         []int // byte offsets of top-level '|'
	firstAbsorber int   // first offset where a construct absorbs later pipes, -1 if none
	topLevelComma bool  // top-level ','
	balanced      bool  // all brackets and quotes closed
	trailingPipe  bool  // the type ends with a '|'

	topLevelQuestion bool // top-level '?'
	firstExtends     int  // offset of the first top-level "extends", -1 if none

	lastSemicolon     int  // offset of the last top-level ';', -1 if none
	trailingSemicolon bool // the type ends with a top-level ';'

	// keyofAbsorbedPipes are top-level pipes that tree-sitter parses inside
	// the operand of a keyof over a constructor type, so they are not root
	// separators.
	keyofAbsorbedPipes []int
}

// separators returns the top-level pipes that split the root union. Constructs
// whose grammar node wraps a union (readonly, function return types, and
// conditional types) absorb every pipe at or after their first offset, so those
// pipes are not root separators.
func (s scanResult) separators() []int {
	pipes := make([]int, 0, len(s.pipes))
	for _, pos := range s.pipes {
		if slices.Contains(s.keyofAbsorbedPipes, pos) {
			continue
		}
		if s.firstAbsorber >= 0 && pos >= s.firstAbsorber {
			continue
		}
		pipes = append(pipes, pos)
	}
	return pipes
}

// tokenKind classifies the last significant top-level token.
type tokenKind uint8

const (
	tokenOther tokenKind = iota
	tokenPipe
	tokenSemicolon
)

// scanType walks typeText and records top-level delimiters. It is not a
// TypeScript parser; it only tracks bracket depth and string literals.
func scanType(s string) scanResult {
	res := scanResult{
		firstAbsorber: -1,
		firstExtends:  -1,
		lastSemicolon: -1,
		balanced:      true,
	}

	var parens, brackets, braces, angles int
	atTopLevel := func() bool {
		return parens == 0 && brackets == 0 && braces == 0 && angles == 0
	}

	sawNew := false
	keyofPending := false
	keyofCtor := false
	ctorArrow := -1
	ctorAmp := false
	firstPipeAfterCtor := -1
	extendsInKeyofCtor := false
	conditionalCtorArrow := -1
	lastToken := tokenOther

	i := 0
	for i < len(s) {
		c := s[i]
		top := atTopLevel()

		// A keyof operand starts with an identifier. Any other top-level
		// token means it is not a constructor type.
		if top && keyofPending && !isSpace(c) && !isIdentStart(c) && !isCommentStart(s, i) {
			keyofPending = false
		}

		switch {
		case c == '\'' || c == '"':
			end, ok := skipQuoted(s, i, c)
			if !ok {
				res.balanced = false
				return res
			}
			lastToken = tokenOther
			i = end

		case c == '`':
			end, ok := skipTemplate(s, i)
			if !ok {
				res.balanced = false
				return res
			}
			lastToken = tokenOther
			i = end

		case isCommentStart(s, i):
			end, ok := skipComment(s, i)
			if !ok {
				res.balanced = false
				return res
			}
			i = end

		case c == '(':
			parens++
			lastToken = tokenOther
			i++
		case c == ')':
			parens--
			if parens < 0 {
				res.balanced = false
				return res
			}
			lastToken = tokenOther
			i++
		case c == '[':
			brackets++
			lastToken = tokenOther
			i++
		case c == ']':
			brackets--
			if brackets < 0 {
				res.balanced = false
				return res
			}
			lastToken = tokenOther
			i++
		case c == '{':
			braces++
			lastToken = tokenOther
			i++
		case c == '}':
			braces--
			if braces < 0 {
				res.balanced = false
				return res
			}
			lastToken = tokenOther
			i++
		case c == '<':
			angles++
			lastToken = tokenOther
			i++
		case c == '>':
			if angles == 0 {
				res.balanced = false
				return res
			}
			angles--
			lastToken = tokenOther
			i++
		case c == '=' && i+1 < len(s) && s[i+1] == '>':
			if top {
				if keyofCtor && sawNew && ctorArrow < 0 {
					ctorArrow = i
				}
				// A non-constructor function type's return type extends to
				// the right, so it absorbs later pipes. Constructor types do
				// not: tree-sitter splits their return type at the root.
				if !sawNew && (res.firstAbsorber < 0 || i < res.firstAbsorber) {
					res.firstAbsorber = i
				}
				sawNew = false
			}
			lastToken = tokenOther
			i += 2
		case c == '&':
			if top && keyofCtor && ctorArrow >= 0 && firstPipeAfterCtor < 0 {
				ctorAmp = true
			}
			lastToken = tokenOther
			i++
		case c == '|':
			if top {
				if keyofCtor && ctorArrow >= 0 && firstPipeAfterCtor < 0 {
					firstPipeAfterCtor = i
					if !ctorAmp {
						res.keyofAbsorbedPipes = append(res.keyofAbsorbedPipes, i)
					}
				}
				res.pipes = append(res.pipes, i)
				sawNew = false
			}
			lastToken = tokenPipe
			i++
		case c == ',':
			if top {
				res.topLevelComma = true
			}
			lastToken = tokenOther
			i++
		case c == '?':
			if top {
				res.topLevelQuestion = true
			}
			lastToken = tokenOther
			i++
		case c == ';':
			if top {
				res.lastSemicolon = i
				lastToken = tokenSemicolon
			} else {
				lastToken = tokenOther
			}
			i++
		default:
			if top && isIdentStart(c) {
				end := i + 1
				for end < len(s) && isIdentPart(s[end]) {
					end++
				}
				word := s[i:end]
				switch {
				case word == "keyof":
					switch {
					case keyofCtor && ctorArrow >= 0 && firstPipeAfterCtor < 0:
						// Nested in the constructor's return type: the outer
						// keyof operand still owns the following pipe.
					case !keyofPending:
						keyofPending = true
						keyofCtor = false
						ctorArrow = -1
						ctorAmp = false
						firstPipeAfterCtor = -1
					}
				case keyofPending:
					switch word {
					case "abstract":
						// Keep waiting for the new keyword.
					case "new":
						keyofPending = false
						keyofCtor = true
					default:
						keyofPending = false
					}
				}
				switch word {
				case "new":
					sawNew = true
				case "readonly":
					// readonly is a prefix operator whose operand is the
					// whole union to its right, so it absorbs later pipes.
					if res.firstAbsorber < 0 || i < res.firstAbsorber {
						res.firstAbsorber = i
					}
				case "extends":
					if res.firstExtends < 0 {
						res.firstExtends = i
						// A conditional directly in a keyof constructor's
						// return type is scoped inside the keyof operand.
						if keyofCtor && ctorArrow >= 0 && i > ctorArrow &&
							firstPipeAfterCtor < 0 && !ctorAmp {
							extendsInKeyofCtor = true
							conditionalCtorArrow = ctorArrow
						}
					}
				}
				lastToken = tokenOther
				i = end
				continue
			}
			if !isSpace(c) {
				lastToken = tokenOther
			}
			i++
		}
	}

	if parens != 0 || brackets != 0 || braces != 0 || angles != 0 {
		res.balanced = false
		return res
	}

	// A conditional type wraps everything to its left into its check type, so
	// it absorbs earlier pipes. A readonly or function absorber that starts
	// before the extends keyword still wins.
	if res.topLevelQuestion && res.firstExtends >= 0 && !extendsInKeyofCtor &&
		(res.firstAbsorber < 0 || res.firstAbsorber > res.firstExtends) {
		res.firstAbsorber = 0
	}

	// A conditional in a keyof constructor's return type owns every following
	// pipe, including unions in its branches.
	if conditionalCtorArrow >= 0 {
		for _, pos := range res.pipes {
			if pos > conditionalCtorArrow && !slices.Contains(res.keyofAbsorbedPipes, pos) {
				res.keyofAbsorbedPipes = append(res.keyofAbsorbedPipes, pos)
			}
		}
	}

	res.trailingPipe = lastToken == tokenPipe
	res.trailingSemicolon = lastToken == tokenSemicolon
	return res
}

// isCommentStart reports whether a line or block comment starts at i.
func isCommentStart(s string, i int) bool {
	return i+1 < len(s) && s[i] == '/' && (s[i+1] == '/' || s[i+1] == '*')
}

// skipComment returns the offset just past the comment starting at i. It
// reports false for a block comment that reaches the end of s.
func skipComment(s string, i int) (int, bool) {
	if !isCommentStart(s, i) {
		return i, false
	}
	second := s[i+1]
	i += 2
	if second == '/' {
		for i < len(s) && s[i] != '\n' {
			i++
		}
		return i, true
	}
	for i+1 < len(s) && (s[i] != '*' || s[i+1] != '/') {
		i++
	}
	if i+1 >= len(s) {
		return len(s), false
	}
	return i + 2, true
}

// splitAt splits s at the given top-level pipe offsets, trimming members and
// dropping empty ones.
func splitAt(s string, pipes []int) []string {
	parts := make([]string, 0, len(pipes)+1)
	start := 0
	for _, pos := range pipes {
		if part := trimTopLevelComments(s[start:pos]); part != "" {
			parts = append(parts, part)
		}
		start = pos + 1
	}
	if part := trimTopLevelComments(s[start:]); part != "" {
		parts = append(parts, part)
	}
	return parts
}

// trimTopLevelComments removes line and block comments at the edges of a union
// member while preserving comments inside the member's type expression and
// inside string and template literals. Returned text is trimmed of whitespace.
func trimTopLevelComments(s string) string {
	first, last := -1, -1

	i := 0
	for i < len(s) {
		switch {
		case isSpace(s[i]):
			i++

		case s[i] == '\'' || s[i] == '"':
			end, ok := skipQuoted(s, i, s[i])
			if !ok {
				end = len(s)
			}
			if first < 0 {
				first = i
			}
			last = end
			i = end

		case s[i] == '`':
			end, ok := skipTemplate(s, i)
			if !ok {
				end = len(s)
			}
			if first < 0 {
				first = i
			}
			last = end
			i = end

		case isCommentStart(s, i):
			i, _ = skipComment(s, i)

		default:
			if first < 0 {
				first = i
			}
			last = i + 1
			i++
		}
	}

	if first < 0 || last <= first {
		return ""
	}
	return strings.TrimSpace(s[first:last])
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
		case isCommentStart(s, i):
			end, ok := skipComment(s, i)
			if !ok {
				return len(s), false
			}
			i = end
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
