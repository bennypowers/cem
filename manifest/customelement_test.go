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
	"io/fs"
	"strings"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
	treeview "github.com/Digital-Shane/treeview/v2"
	"github.com/stretchr/testify/assert"
)

func TestRenderableCustomElementDeclaration(t *testing.T) {
	t.Run("ToTreeNode", func(t *testing.T) {
		fixtureFS := testutil.NewFixtureFS(t, "", "/")
		manifestJSON, err := fs.ReadFile(fixtureFS, "/custom-element-member-grouping.json")
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
		var moduleNode TreeNode
		for _, child := range node.Children() {
			if stripANSI(child.Name()) == "module src/my-module.js" {
				moduleNode = child
				break
			}
		}
		if !assert.NotNil(t, moduleNode, "Module node should exist") {
			t.FailNow()
		}

		// Find the custom element class node
		var elemNode TreeNode
		for _, child := range moduleNode.Children() {
			text := stripANSI(child.Name())
			if strings.Contains(text, "<my-element>") {
				elemNode = child
				break
			}
		}
		if !assert.NotNil(t, elemNode, "Custom element node should exist") {
			dumpTree(t, node, 0)
			t.FailNow()
		}

		t.Run("RootNodeIsElement", func(t *testing.T) {
			assert.Contains(t, stripANSI(elemNode.Name()), "my-element")
		})

		t.Run("ChildrenAreGroupedByKind", func(t *testing.T) {
			kinds := []string{
				"Attributes", "Slots", "Events", "Fields", "Methods",
				"CSS Properties", "Parts", "States",
			}
			groupLabels := make([]string, 0)
			for _, child := range elemNode.Children() {
				groupLabels = append(groupLabels, stripANSI(child.Name()))
			}
			for _, kind := range kinds {
				assert.Containsf(t, groupLabels, kind, "expected group %q in tree node children", kind)
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
