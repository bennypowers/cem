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

	ts "github.com/tree-sitter/go-tree-sitter"
)

// Inline: pure node-to-struct and string matching over single statements;
// table-driven cases are clearer than goldens (same rationale as
// TestParameterFromNode). Source-href emission for referencing types is
// covered by the generics golden fixtures.

func TestTypeParameterBindings(t *testing.T) {
	tests := []struct {
		name      string
		stmt      string
		wantNames []string
	}{
		{
			name:      "constraint and default",
			stmt:      "function identity<Value extends string = string>(value: Value): Value { return value; }",
			wantNames: []string{"Value"},
		},
		{
			name:      "bare parameter",
			stmt:      "function f<T>(value: T) { return value; }",
			wantNames: []string{"T"},
		},
		{
			name:      "multiple parameters in declaration order",
			stmt:      "function f<K extends string = string, V = unknown>(k: K, v: V) { return v; }",
			wantNames: []string{"K", "V"},
		},
		{
			name:      "class declaration",
			stmt:      "class C<T extends HTMLElement> {}",
			wantNames: []string{"T"},
		},
		{
			name:      "class method",
			stmt:      "class C { m<T>(x: T): T { return x; } }",
			wantNames: []string{"T"},
		},
		{
			name:      "arrow function with trailing comma",
			stmt:      "const f = <T,>(value: T) => value;",
			wantNames: []string{"T"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node, tree := typeParametersNode(t, tt.stmt)
			defer tree.Close()
			got := typeParameterBindings(node, []byte(tt.stmt))
			if len(got) != len(tt.wantNames) {
				t.Fatalf("typeParameterBindings(%q) = %v names, want %v", tt.stmt, got, tt.wantNames)
			}
			for i := range got {
				if got[i].name != tt.wantNames[i] {
					t.Errorf("typeParameterBindings(%q)[%d].name = %q, want %q", tt.stmt, i, got[i].name, tt.wantNames[i])
				}
				if got[i].nameNode == nil {
					t.Errorf("typeParameterBindings(%q)[%d].nameNode = nil, want the declaring identifier", tt.stmt, i)
				}
			}
		})
	}

	t.Run("nil node", func(t *testing.T) {
		if got := typeParameterBindings(nil, nil); got != nil {
			t.Errorf("typeParameterBindings(nil) = %v, want nil", got)
		}
	})
}

func TestBindingForTypeText(t *testing.T) {
	bindings := []typeParameterBinding{
		{name: "Key"},
		{name: "Value"},
	}
	tests := []struct {
		name     string
		typeText string
		want     string // "" for no match
	}{
		{"exact name", "Value", "Value"},
		{"array type", "Value[]", "Value"},
		{"generic argument", "Record<Key, Value>", "Key"},
		{"first declared wins", "Record<Key, Value>", "Key"},
		{"prefix is not a word", "ValueX", ""},
		{"suffix is not a word", "MyValue", ""},
		{"unrelated type", "number", ""},
		{"empty text", "", ""},
		{"string-literal content is ignored", "\"Value\"", ""},
		{"quoted record keys are ignored", "Record<'Key', Value>", "Value"},
		{"template-literal content is ignored", "`prefix-${Value}`", ""},
		{"name outside quotes still matches", "Value | \"Value\"", "Value"},
		{"escaped quote in string-literal still matches trailing name", `'it\'s' | Value`, "Value"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bindingForTypeText(bindings, tt.typeText)
			if tt.want == "" {
				if got != nil {
					t.Errorf("bindingForTypeText(%q) = %q, want nil", tt.typeText, got.name)
				}
				return
			}
			if got == nil {
				t.Fatalf("bindingForTypeText(%q) = nil, want %q", tt.typeText, tt.want)
			}
			if got.name != tt.want {
				t.Errorf("bindingForTypeText(%q) = %q, want %q", tt.typeText, got.name, tt.want)
			}
		})
	}
}

func TestContainsWord(t *testing.T) {
	tests := []struct {
		name string
		text string
		word string
		want bool
	}{
		{"whole word", "Value", "Value", true},
		{"inside array", "Value[]", "Value", true},
		{"inside generic args", "Record<Key, Value>", "Value", true},
		{"at start", "Value | null", "Value", true},
		{"at end", "null | Value", "Value", true},
		{"prefix extends word", "ValueX", "Value", false},
		{"suffix extends word", "MyValue", "Value", false},
		{"underscore continues word", "Value_Type", "Value", false},
		{"dollar continues word", "$Value", "Value", false},
		{"absent", "number", "Value", false},
		{"empty word", "anything", "", false},
		{"non-ascii identifier boundary", "Välue Value", "Value", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsWord(tt.text, tt.word); got != tt.want {
				t.Errorf("containsWord(%q, %q) = %v, want %v", tt.text, tt.word, got, tt.want)
			}
		})
	}
}

func TestStripQuotedSegments(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{"no quotes", "Record<Key, Value>", "Record<Key, Value>"},
		{"double-quoted string", "\"Value\"", ""},
		{"single-quoted key", "Record<'Key', Value>", "Record<, Value>"},
		{"backtick template", "`prefix-${Value}`", ""},
		{"unterminated quote keeps prefix", "Value | \"unterminated", "Value | "},
		{"mixed quotes", "'a' | \"b\" | `c`", " |  | "},
		{"escaped closing quote does not close the span", `'it\'s' | Value`, " | Value"},
		{"escaped backslash before closing quote", `'a\\' | Value`, " | Value"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stripQuotedSegments(tt.text); got != tt.want {
				t.Errorf("stripQuotedSegments(%q) = %q, want %q", tt.text, got, tt.want)
			}
		})
	}
}

// typeParametersNode parses stmt and returns its first type_parameters node
// and the tree keeping it alive; the caller must Close the tree after using
// the node.
func typeParametersNode(t *testing.T, stmt string) (*ts.Node, *ts.Tree) {
	t.Helper()
	root, tree := parseStatement(stmt)
	if tree == nil {
		t.Fatal("failed to parse statement")
	}
	var find func(n *ts.Node) *ts.Node
	find = func(n *ts.Node) *ts.Node {
		if n.GrammarName() == "type_parameters" {
			return n
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
		t.Fatalf("no type_parameters node in %q", stmt)
	}
	return node, tree
}
