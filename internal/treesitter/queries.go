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
package treesitter

import (
	"fmt"

	"bennypowers.dev/cem/internal/textutil"
	ts "github.com/tree-sitter/go-tree-sitter"
)

type NoCaptureError struct {
	Capture string
	Query   string
}

func (e *NoCaptureError) Error() string {
	return fmt.Sprintf("No nodes for capture %s in query %s", e.Capture, e.Query)
}

type ParentNodeCaptures struct {
	NodeId   uintptr
	Captures CaptureMap
}

type CaptureInfo struct {
	NodeId    int
	Text      string
	StartByte uint
	EndByte   uint
}

type CaptureMap = map[string][]CaptureInfo

func GetDescendantById(root *ts.Node, id int) *ts.Node {
	c := root.Walk()
	defer c.Close()
	var find func(node *ts.Node) *ts.Node
	find = func(node *ts.Node) *ts.Node {
		if int(node.Id()) == id {
			return node
		}
		for i := range int(node.ChildCount()) {
			child := node.Child(uint(i))
			if child == nil {
				continue
			}
			if res := find(child); res != nil {
				return res
			}
		}
		return nil
	}

	return find(root)
}

// Position represents a line/character position
type Position struct {
	Line      uint32
	Character uint32
}

// Range represents a start/end range
type Range struct {
	Start Position
	End   Position
}

// NodeToRange converts a tree-sitter node to a Range using byte-to-position conversion
func NodeToRange(node *ts.Node, content []byte) Range {
	start := ByteOffsetToPosition(content, node.StartByte())
	end := ByteOffsetToPosition(content, node.EndByte())
	return Range{
		Start: start,
		End:   end,
	}
}

// ByteOffsetToPosition converts a byte offset to line/character position.
// Character is in UTF-16 code units per the LSP specification.
func ByteOffsetToPosition(content []byte, offset uint) Position {
	line := uint32(0)
	lineStart := uint(0)

	for i := uint(0); i < offset && i < uint(len(content)); i++ {
		if content[i] == '\n' {
			line++
			lineStart = i + 1
		}
	}

	return Position{
		Line:      line,
		Character: textutil.ByteOffsetToUTF16Bytes(content[lineStart:], offset-lineStart),
	}
}
