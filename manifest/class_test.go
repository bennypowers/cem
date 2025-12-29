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
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
)

// Helper to strip ANSI color codes
var ansiRegexp = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func stripANSI(s string) string {
	return ansiRegexp.ReplaceAllString(s, "")
}

func dumpTree(t *testing.T, node pterm.TreeNode, depth int) {
	t.Logf("%s%q\n", strings.Repeat("  ", depth), stripANSI(node.Text))
	for _, child := range node.Children {
		dumpTree(t, child, depth+1)
	}
}

func TestRenderableClassDeclaration(t *testing.T) {
	t.Run("ToTreeNode", func(t *testing.T) {
		manifestJSON, err := os.ReadFile(filepath.Join("fixtures", "class-member-grouping.json"))
		if err != nil {
			t.Fatal(err)
		}

		var pkg Package
		if err := json.Unmarshal([]byte(manifestJSON), &pkg); err != nil {
			t.Fatal(err)
		}

		renderable := NewRenderablePackage(&pkg)
		node := renderable.ToTreeNode(True)

		// Find the module node by partial match if necessary
		var moduleNode *pterm.TreeNode
		for _, child := range node.Children {
			if stripANSI(child.Text) == "src/my-module.js" || strings.Contains(stripANSI(child.Text), "my-module") {
				moduleNode = &child
				break
			}
		}
		if !assert.NotNil(t, moduleNode, "Module node should exist") {
			t.Logf("Available root children: %#v", node.Children)
			t.FailNow()
		}

		// Find the class node, ignoring ANSI color codes
		var classNode *pterm.TreeNode
		for _, child := range moduleNode.Children {
			text := stripANSI(child.Text)
			if strings.Contains(text, "MyClass") && strings.Contains(text, "class") {
				classNode = &child
				break
			}
		}
		if !assert.NotNil(t, classNode, "Class node should exist") {
			t.Logf("Available module children: %#v", moduleNode.Children)
			t.FailNow()
		}

		t.Run("RootNodeIsClass", func(t *testing.T) {
			assert.Contains(t, stripANSI(classNode.Text), "class")
			assert.Contains(t, stripANSI(classNode.Text), "MyClass")
		})

		t.Run("ChildrenAreGroupedByKind", func(t *testing.T) {
			kinds := []string{"Fields", "Methods"}
			groupLabels := make([]string, 0)
			for _, child := range classNode.Children {
				groupLabels = append(groupLabels, stripANSI(child.Text))
			}
			for _, kind := range kinds {
				assert.Containsf(t, groupLabels, kind, "expected group %q in class node children", kind)
			}
		})

		t.Run("FieldsGroupedTogether", func(t *testing.T) {
			var fieldsGroup *pterm.TreeNode
			for _, child := range classNode.Children {
				if stripANSI(child.Text) == "Fields" {
					fieldsGroup = &child
					break
				}
			}
			if assert.NotNil(t, fieldsGroup, "Fields group should exist") {
				fieldNames := []string{}
				for _, f := range fieldsGroup.Children {
					fieldNames = append(fieldNames, stripANSI(f.Text))
				}
				dumpTree(t, node, 0)
				assert.Subset(t, fieldNames, []string{"myField1", "myField2", "anotherField"})
			}
		})

		t.Run("MethodsGroupedTogether", func(t *testing.T) {
			var methodsGroup *pterm.TreeNode
			for _, child := range classNode.Children {
				if stripANSI(child.Text) == "Methods" {
					methodsGroup = &child
					break
				}
			}
			if assert.NotNil(t, methodsGroup, "Methods group should exist") {
				methodNames := []string{}
				for _, m := range methodsGroup.Children {
					methodNames = append(methodNames, stripANSI(m.Text))
				}
				assert.Subset(t, methodNames, []string{"myMethod1", "myMethod2"})
			}
		})

		t.Run("VisualTreeOutput", func(t *testing.T) {
			_, err := pterm.DefaultTree.WithRoot(node).Srender()
			assert.NoError(t, err)
		})
	})
}
