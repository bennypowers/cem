package manifest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
)

func TestRenderableCustomElementDeclaration(t *testing.T) {
	t.Run("ToTreeNode", func(t *testing.T) {
		manifestJSON, err := os.ReadFile(filepath.Join("fixtures", "custom-element-member-grouping.json"))
		if err != nil {
			t.Fatal(err)
		}

		var pkg Package
		if err := json.Unmarshal([]byte(manifestJSON), &pkg); err != nil {
			t.Fatal(err)
		}

		renderable := NewRenderablePackage(&pkg)
		node := renderable.ToTreeNode(True)

		// Find the module node
		var moduleNode *pterm.TreeNode
		for _, child := range node.Children {
			if stripANSI(child.Text) == "src/my-module.js" {
				moduleNode = &child
				break
			}
		}
		if !assert.NotNil(t, moduleNode, "Module node should exist") {
			t.FailNow()
		}

		// Find the custom element class node
		var elemNode *pterm.TreeNode
		for _, child := range moduleNode.Children {
			t.Log(child.Text)
			text := stripANSI(child.Text)
			if strings.Contains(text, "<my-element>") {
				elemNode = &child
				break
			}
		}
		if !assert.NotNil(t, elemNode, "Custom element node should exist") {
			t.FailNow()
		}

		t.Run("RootNodeIsElement", func(t *testing.T) {
			assert.Contains(t, stripANSI(elemNode.Text), "my-element")
		})

		t.Run("ChildrenAreGroupedByKind", func(t *testing.T) {
			kinds := []string{
				"Attributes", "Slots", "Events", "Fields", "Methods",
				"CSS Properties", "Parts", "States",
			}
			groupLabels := make([]string, 0)
			for _, child := range elemNode.Children {
				groupLabels = append(groupLabels, stripANSI(child.Text))
			}
			for _, kind := range kinds {
				assert.Containsf(t, groupLabels, kind, "expected group %q in tree node children", kind)
			}
		})

		t.Run("VisualTreeOutput", func(t *testing.T) {
			_, err := pterm.DefaultTree.WithRoot(node).Srender()
			assert.NoError(t, err)
		})
	})
}
