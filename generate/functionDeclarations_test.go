/*
Copyright © 2026 Benny Powers

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
package generate

import (
	"testing"

	M "bennypowers.dev/cem/manifest"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// Inline: pure node-to-struct mapping over single statements; table-driven
// cases are clearer than goldens (same rationale as TestIsTypeOnlyNode).
// Multi-parameter ordering and query integration are covered by the
// project-function-params golden fixtures.

func TestParameterFromNode(t *testing.T) {
	tests := []struct {
		name  string
		stmt  string
		param int
		want  M.Parameter
	}{
		{
			name:  "typed required",
			stmt:  "function f(value: number) {}",
			param: 0,
			want:  M.Parameter{PropertyLike: M.PropertyLike{FullyQualified: M.FullyQualified{Name: "value"}, Type: &M.Type{Text: "number"}}},
		},
		{
			name:  "optional marker",
			stmt:  "function f(value?: number) {}",
			param: 0,
			want:  M.Parameter{Optional: true, PropertyLike: M.PropertyLike{FullyQualified: M.FullyQualified{Name: "value"}, Type: &M.Type{Text: "number"}}},
		},
		{
			name:  "typed default",
			stmt:  "function f(value: number = 1) {}",
			param: 0,
			want:  M.Parameter{Optional: true, PropertyLike: M.PropertyLike{FullyQualified: M.FullyQualified{Name: "value"}, Type: &M.Type{Text: "number"}, Default: "1"}},
		},
		{
			name:  "untyped default",
			stmt:  "function f(value = 1) {}",
			param: 0,
			want:  M.Parameter{Optional: true, PropertyLike: M.PropertyLike{FullyQualified: M.FullyQualified{Name: "value"}, Default: "1"}},
		},
		{
			name:  "optional with default",
			stmt:  "function f(value?: number = 1) {}",
			param: 0,
			want:  M.Parameter{Optional: true, PropertyLike: M.PropertyLike{FullyQualified: M.FullyQualified{Name: "value"}, Type: &M.Type{Text: "number"}, Default: "1"}},
		},
		{
			name:  "untyped no default",
			stmt:  "function f(value) {}",
			param: 0,
			want:  M.Parameter{PropertyLike: M.PropertyLike{FullyQualified: M.FullyQualified{Name: "value"}}},
		},
		{
			name:  "rest",
			stmt:  "function f(...args: string[]) {}",
			param: 0,
			want:  M.Parameter{Rest: true, PropertyLike: M.PropertyLike{FullyQualified: M.FullyQualified{Name: "args"}, Type: &M.Type{Text: "string[]"}}},
		},
		{
			name:  "object type annotation",
			stmt:  "function f(options: { enabled?: boolean } = {}) {}",
			param: 0,
			want:  M.Parameter{Optional: true, PropertyLike: M.PropertyLike{FullyQualified: M.FullyQualified{Name: "options"}, Type: &M.Type{Text: "{ enabled?: boolean }"}, Default: "{}"}},
		},
		{
			name:  "second parameter",
			stmt:  "function f(first: string, second: number) {}",
			param: 1,
			want:  M.Parameter{PropertyLike: M.PropertyLike{FullyQualified: M.FullyQualified{Name: "second"}, Type: &M.Type{Text: "number"}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node, tree := paramNode(t, tt.stmt, tt.param)
			defer tree.Close()
			got, ok := parameterFromNode(node, []byte(tt.stmt))
			if !ok {
				t.Fatalf("parameterFromNode(%q, param %d) reported false, want true", tt.stmt, tt.param)
			}
			if got.Name != tt.want.Name || got.Optional != tt.want.Optional ||
				got.Rest != tt.want.Rest || got.Default != tt.want.Default {
				t.Errorf("parameterFromNode(%q, param %d) = %+v, want %+v", tt.stmt, tt.param, got, tt.want)
			}
			if tt.want.Type == nil {
				if got.Type != nil {
					t.Errorf("parameterFromNode(%q, param %d) type = %v, want nil", tt.stmt, tt.param, got.Type)
				}
			} else if got.Type == nil || got.Type.Text != tt.want.Type.Text {
				t.Errorf("parameterFromNode(%q, param %d) type = %v, want %v", tt.stmt, tt.param, got.Type, tt.want.Type)
			}
		})
	}

	t.Run("skipped shapes", func(t *testing.T) {
		for _, stmt := range []string{
			"function f(this: Window, value: number) {}",
			"function f({ a, b }: Options) {}",
			"function f([a, b]: string[]) {}",
		} {
			node, tree := paramNode(t, stmt, 0)
			got, ok := parameterFromNode(node, []byte(stmt))
			tree.Close()
			if ok {
				t.Errorf("parameterFromNode(%q, param 0) = %+v, want skipped", stmt, got)
			}
		}
	})

	t.Run("nil node", func(t *testing.T) {
		if got, ok := parameterFromNode(nil, nil); ok {
			t.Errorf("parameterFromNode(nil) = %+v, want skipped", got)
		}
	})
}

// paramNode parses stmt and returns its i'th formal parameter node and the
// tree keeping it alive; the caller must Close the tree after using the node
func paramNode(t *testing.T, stmt string, i int) (*ts.Node, *ts.Tree) {
	t.Helper()
	root, tree := parseStatement(stmt)
	if tree == nil {
		t.Fatal("failed to parse statement")
	}
	var find func(n *ts.Node) *ts.Node
	find = func(n *ts.Node) *ts.Node {
		if n.GrammarName() == "formal_parameters" {
			params := make([]*ts.Node, 0, n.ChildCount())
			for j := range n.ChildCount() {
				if c := n.Child(j); c.GrammarName() == "required_parameter" || c.GrammarName() == "optional_parameter" {
					params = append(params, c)
				}
			}
			if i < len(params) {
				return params[i]
			}
			return nil
		}
		for j := range n.ChildCount() {
			if found := find(n.Child(j)); found != nil {
				return found
			}
		}
		return nil
	}
	node := find(root)
	if node == nil {
		tree.Close()
		t.Fatalf("no parameter %d in %q", i, stmt)
	}
	return node, tree
}
