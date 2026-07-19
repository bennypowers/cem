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
package testhelpers

import (
	"strings"

	"bennypowers.dev/cem/internal/textutil"
	"go.lsp.dev/protocol"
	ts "github.com/tree-sitter/go-tree-sitter"
	tsCss "github.com/tree-sitter/tree-sitter-css/bindings/go"
	tsTypescript "github.com/tree-sitter/tree-sitter-typescript/bindings/go"
	"golang.org/x/net/html"
)

// TSCursorParser extracts cursor position from <!-- ^cursor --> comments
// inside tagged template literals in TypeScript, or /* ^cursor */ comments
// in regular TS code. Compatible with testutil.CursorParser.
func TSCursorParser(content string) (string, *protocol.Position) {
	lang := ts.NewLanguage(tsTypescript.LanguageTypescript())
	parser := ts.NewParser()
	defer parser.Close()
	if err := parser.SetLanguage(lang); err != nil {
		return content, nil
	}
	tree := parser.Parse([]byte(content), nil)
	if tree == nil {
		return content, nil
	}
	defer tree.Close()

	// Try HTML comments inside template string fragments first
	cleaned, cursor := findCursorInTemplateStrings(tree.RootNode(), content)
	if cursor != nil {
		return cleaned, cursor
	}
	// Fall back to /* ^cursor */ comments in regular TS code
	return treeSitterCursorParser(content, tree)
}

// CSSCursorParser extracts cursor position from /* ^cursor */ comments in CSS.
// Compatible with testutil.CursorParser.
func CSSCursorParser(content string) (string, *protocol.Position) {
	lang := ts.NewLanguage(tsCss.Language())
	parser := ts.NewParser()
	defer parser.Close()
	if err := parser.SetLanguage(lang); err != nil {
		return content, nil
	}
	tree := parser.Parse([]byte(content), nil)
	if tree == nil {
		return content, nil
	}
	defer tree.Close()
	return treeSitterCursorParser(content, tree)
}

// findCursorInTemplateStrings searches template_string nodes for HTML
// <!-- ^cursor --> comments. Uses the Go HTML tokenizer on each
// string_fragment's content to find comment nodes.
func findCursorInTemplateStrings(root *ts.Node, content string) (string, *protocol.Position) {
	var fragments []*ts.Node
	collectStringFragments(root, &fragments)

	for _, frag := range fragments {
		fragStart := int(frag.StartByte())
		fragEnd := int(frag.EndByte())
		fragContent := content[fragStart:fragEnd]

		tokenizer := html.NewTokenizer(strings.NewReader(fragContent))
		htmlOffset := 0
		for {
			tt := tokenizer.Next()
			if tt == html.ErrorToken {
				break
			}
			raw := string(tokenizer.Raw())
			if tt == html.CommentToken && strings.Contains(raw, "^cursor") {
				caretInRaw := strings.Index(raw, "^")
				if caretInRaw < 0 {
					htmlOffset += len(raw)
					continue
				}
				// Map back to original source position
				globalOffset := fragStart + htmlOffset + caretInRaw
				caretLine, caretChar := byteOffsetToPosition(content, globalOffset)
				if caretLine == 0 {
					htmlOffset += len(raw)
					continue
				}
				// Strip the entire line containing the comment (LF and CRLF)
				lineStart := fragStart + htmlOffset
				for lineStart > 0 && content[lineStart-1] != '\n' {
					lineStart--
				}
				lineEnd := fragStart + htmlOffset + len(raw)
				if lineEnd < len(content) && content[lineEnd] == '\r' {
					lineEnd++
				}
				if lineEnd < len(content) && content[lineEnd] == '\n' {
					lineEnd++
				}
				cleaned := content[:lineStart] + content[lineEnd:]
				return cleaned, &protocol.Position{
					Line:      uint32(caretLine - 1),
					Character: caretChar,
				}
			}
			htmlOffset += len(raw)
		}
	}
	return content, nil
}

func collectStringFragments(node *ts.Node, out *[]*ts.Node) {
	if node.Kind() == "string_fragment" {
		parent := node.Parent()
		if parent != nil && parent.Kind() == "template_string" {
			*out = append(*out, node)
		}
	}
	for i := range int(node.ChildCount()) {
		child := node.Child(uint(i))
		if child != nil {
			collectStringFragments(child, out)
		}
	}
}

func treeSitterCursorParser(content string, tree *ts.Tree) (string, *protocol.Position) {
	node := findCursorComment(tree.RootNode(), []byte(content))
	if node == nil {
		return content, nil
	}

	nodeText := content[node.StartByte():node.EndByte()]
	caretInNode := strings.Index(nodeText, "^")
	if caretInNode < 0 {
		return content, nil
	}

	caretRow := node.StartPosition().Row
	if caretRow == 0 {
		return content, nil
	}

	lineStart := int(node.StartByte())
	for lineStart > 0 && content[lineStart-1] != '\n' {
		lineStart--
	}
	caretByteCol := int(node.StartByte()) - lineStart + caretInNode
	caretChar := textutil.ByteOffsetToUTF16(content[lineStart:], uint(caretByteCol))

	lineEnd := int(node.EndByte())
	if lineEnd < len(content) && content[lineEnd] == '\r' {
		lineEnd++
	}
	if lineEnd < len(content) && content[lineEnd] == '\n' {
		lineEnd++
	}
	cleaned := content[:lineStart] + content[lineEnd:]

	return cleaned, &protocol.Position{
		Line:      uint32(caretRow - 1),
		Character: caretChar,
	}
}

// byteOffsetToPosition converts a byte offset to an LSP-compatible
// (line, character) pair where character is in UTF-16 code units.
func byteOffsetToPosition(content string, offset int) (line int, character uint32) {
	lineStart := 0
	for i := range offset {
		if i >= len(content) {
			break
		}
		if content[i] == '\n' {
			line++
			lineStart = i + 1
		}
	}
	byteCol := uint(offset - lineStart)
	character = textutil.ByteOffsetToUTF16(content[lineStart:], byteCol)
	return line, character
}

func findCursorComment(node *ts.Node, source []byte) *ts.Node {
	if node.Kind() == "comment" {
		text := string(source[node.StartByte():node.EndByte()])
		if strings.Contains(text, "^cursor") {
			return node
		}
	}
	for i := range int(node.ChildCount()) {
		child := node.Child(uint(i))
		if child == nil {
			continue
		}
		if found := findCursorComment(child, source); found != nil {
			return found
		}
	}
	return nil
}
