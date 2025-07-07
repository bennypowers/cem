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

		// Print all root children
		for i, child := range node.Children {
			t.Logf("Root child[%d]: %q", i, child.Text)
		}

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

		// Print all module children
		for i, child := range moduleNode.Children {
			t.Logf("Module child[%d]: %q", i, child.Text)
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
				assert.Subset(t, fieldNames, []string{"field myField1", "field myField2", "field anotherField"})
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
