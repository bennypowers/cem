package tstype

import (
	"strings"

	"bennypowers.dev/cem/internal/languages/typescript"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// ParseTypeValue parses typeText as a TypeScript type via tree-sitter and
// returns the value node from `type _ = <typeText>`. Caller must close the
// returned tree.
func ParseTypeValue(typeText string) (valueNode *ts.Node, source []byte, tree *ts.Tree, ok bool) {
	typeText = strings.TrimSpace(typeText)
	if typeText == "" {
		return nil, nil, nil, false
	}

	source = []byte("type _ = " + typeText)
	parser := typescript.BorrowParser()
	defer typescript.ReturnParser(parser)

	tree = parser.Parse(source, nil)
	if tree == nil {
		return nil, nil, nil, false
	}

	root := tree.RootNode()
	cursor := root.Walk()
	defer cursor.Close()

	for _, child := range root.NamedChildren(cursor) {
		if child.GrammarName() == "type_alias_declaration" {
			if vn := child.ChildByFieldName("value"); vn != nil {
				return vn, source, tree, true
			}
		}
	}

	tree.Close()
	return nil, nil, nil, false
}

// SplitTopLevelUnion parses typeText as a TypeScript type using tree-sitter
// and splits top-level union members. Returns nil for empty input, a
// single-element slice for non-union types.
func SplitTopLevelUnion(typeText string) []string {
	typeText = strings.TrimSpace(typeText)
	if typeText == "" {
		return nil
	}

	valueNode, source, tree, ok := ParseTypeValue(typeText)
	if !ok {
		return []string{typeText}
	}
	defer tree.Close()

	if valueNode.GrammarName() == "union_type" {
		return FlattenUnionType(valueNode, source)
	}
	return []string{typeText}
}

// FlattenUnionType recursively collects leaf type texts from a
// left-recursive union_type tree-sitter node.
func FlattenUnionType(node *ts.Node, source []byte) []string {
	var parts []string
	cursor := node.Walk()
	defer cursor.Close()

	for _, child := range node.NamedChildren(cursor) {
		if child.GrammarName() == "union_type" {
			parts = append(parts, FlattenUnionType(&child, source)...)
		} else {
			text := strings.TrimSpace(child.Utf8Text(source))
			if text != "" {
				parts = append(parts, text)
			}
		}
	}

	return parts
}
