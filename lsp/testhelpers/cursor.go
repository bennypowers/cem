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
	protocol "github.com/bennypowers/glsp/protocol_3_17"
	ts "github.com/tree-sitter/go-tree-sitter"
	tsCss "github.com/tree-sitter/tree-sitter-css/bindings/go"
	tsTypescript "github.com/tree-sitter/tree-sitter-typescript/bindings/go"
)

// TSCursorParser extracts cursor position from /* ^cursor */ comments in TypeScript.
// Compatible with testutil.CursorParser.
func TSCursorParser(content string) (string, *protocol.Position) {
	lang := ts.NewLanguage(tsTypescript.LanguageTypescript())
	return treeSitterCursorParser(content, lang)
}

// CSSCursorParser extracts cursor position from /* ^cursor */ comments in CSS.
// Compatible with testutil.CursorParser.
func CSSCursorParser(content string) (string, *protocol.Position) {
	lang := ts.NewLanguage(tsCss.Language())
	return treeSitterCursorParser(content, lang)
}

func treeSitterCursorParser(content string, lang *ts.Language) (string, *protocol.Position) {
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

	// Find the start of the line containing the comment for UTF-16 conversion
	lineStart := int(node.StartByte())
	for lineStart > 0 && content[lineStart-1] != '\n' {
		lineStart--
	}
	caretByteCol := int(node.StartByte()) - lineStart + caretInNode
	caretChar := textutil.ByteOffsetToUTF16(content[lineStart:], uint(caretByteCol))

	lineEnd := int(node.EndByte())
	if lineEnd < len(content) && content[lineEnd] == '\n' {
		lineEnd++
	}
	cleaned := content[:lineStart] + content[lineEnd:]

	return cleaned, &protocol.Position{
		Line:      uint32(caretRow - 1),
		Character: caretChar,
	}
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
