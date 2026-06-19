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
package manifest

import (
	"encoding/json"
	"regexp"
	"strings"
	"testing"

	treeview "github.com/Digital-Shane/treeview/v2"
	"github.com/stretchr/testify/assert"
)

// Inline: integration test, scalar assertions (tree structure verification against fixture)

// Helper to strip ANSI color codes
var ansiRegexp = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func stripANSI(s string) string {
	return ansiRegexp.ReplaceAllString(s, "")
}

func dumpTree(t *testing.T, node TreeNode, depth int) {
	t.Logf("%s%q\n", strings.Repeat("  ", depth), stripANSI(node.Name()))
	for _, child := range node.Children() {
		dumpTree(t, child, depth+1)
	}
}

func TestRenderableClassDeclaration(t *testing.T) {
	t.Run("ToTreeNode", func(t *testing.T) {
		manifestJSON := loadFixture(t, "class-member-grouping.json")

		var pkg Package
		if err := json.Unmarshal([]byte(manifestJSON), &pkg); err != nil {
			t.Fatal(err)
		}

		renderable := NewRenderablePackage(&pkg)
		node := renderable.ToTreeNode(True)

		// Find the module node by partial match if necessary
		var moduleNode TreeNode
		for _, child := range node.Children() {
			if stripANSI(child.Name()) == "src/my-module.js" || strings.Contains(stripANSI(child.Name()), "my-module") {
				moduleNode = child
				break
			}
		}
		if !assert.NotNil(t, moduleNode, "Module node should exist") {
			t.Logf("Available root children: %#v", node.Children())
			t.FailNow()
		}

		// Find the class node, ignoring ANSI color codes
		var classNode TreeNode
		for _, child := range moduleNode.Children() {
			text := stripANSI(child.Name())
			if strings.Contains(text, "MyClass") && strings.Contains(text, "class") {
				classNode = child
				break
			}
		}
		if !assert.NotNil(t, classNode, "Class node should exist") {
			t.Logf("Available module children: %#v", moduleNode.Children())
			t.FailNow()
		}

		t.Run("RootNodeIsClass", func(t *testing.T) {
			assert.Contains(t, stripANSI(classNode.Name()), "class")
			assert.Contains(t, stripANSI(classNode.Name()), "MyClass")
		})

		t.Run("ChildrenAreGroupedByKind", func(t *testing.T) {
			kinds := []string{"Fields", "Methods"}
			groupLabels := make([]string, 0)
			for _, child := range classNode.Children() {
				groupLabels = append(groupLabels, stripANSI(child.Name()))
			}
			for _, kind := range kinds {
				assert.Containsf(t, groupLabels, kind, "expected group %q in class node children", kind)
			}
		})

		t.Run("FieldsGroupedTogether", func(t *testing.T) {
			var fieldsGroup TreeNode
			for _, child := range classNode.Children() {
				if stripANSI(child.Name()) == "Fields" {
					fieldsGroup = child
					break
				}
			}
			if assert.NotNil(t, fieldsGroup, "Fields group should exist") {
				fieldNames := []string{}
				for _, f := range fieldsGroup.Children() {
					fieldNames = append(fieldNames, stripANSI(f.Name()))
				}
				dumpTree(t, node, 0)
				assert.Subset(t, fieldNames, []string{"myField1", "myField2", "anotherField"})
			}
		})

		t.Run("MethodsGroupedTogether", func(t *testing.T) {
			var methodsGroup TreeNode
			for _, child := range classNode.Children() {
				if stripANSI(child.Name()) == "Methods" {
					methodsGroup = child
					break
				}
			}
			if assert.NotNil(t, methodsGroup, "Methods group should exist") {
				methodNames := []string{}
				for _, m := range methodsGroup.Children() {
					methodNames = append(methodNames, stripANSI(m.Name()))
				}
				assert.Subset(t, methodNames, []string{"myMethod1", "myMethod2"})
			}
		})

		t.Run("VisualTreeOutput", func(t *testing.T) {
			tree := treeview.NewTree([]TreeNode{node},
				treeview.WithExpandFunc[DisplayNode](func(_ TreeNode) bool { return true }),
			)
			model := treeview.NewTuiTreeModel(tree,
				treeview.WithTuiDisableNavBar[DisplayNode](true),
			)
			s := model.View().Content
			assert.NotEmpty(t, s)
		})
	})
}
