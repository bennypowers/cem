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
	"strings"
	"testing"

	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
)

func TestDeprecated(t *testing.T) {
	t.Run("deprecated-module", func(t *testing.T) {

		manifestPath := filepath.Join("testdata", "deprecated-module.json")
		raw, err := os.ReadFile(manifestPath)
		if err != nil {
			t.Fatalf("failed to read manifest fixture: %v", err)
		}

		var pkg Package
		if err := json.Unmarshal(raw, &pkg); err != nil {
			t.Fatalf("failed to unmarshal manifest: %v", err)
		}

		renderable := NewRenderablePackage(&pkg)
		assert.True(t, renderable.Children()[0].(*RenderableModule).IsDeprecated())
	})

	t.Run("deprecations", func(t *testing.T) {
		manifestPath := filepath.Join("testdata", "deprecations.json")
		raw, err := os.ReadFile(manifestPath)
		if err != nil {
			t.Fatalf("failed to read manifest fixture: %v", err)
		}

		var pkg Package
		if err := json.Unmarshal(raw, &pkg); err != nil {
			t.Fatalf("failed to unmarshal manifest: %v", err)
		}

		renderable := NewRenderablePackage(&pkg)

		t.Run("Tree includes only deprecated nodes with --deprecated", func(t *testing.T) {
			// True: show only deprecated nodes
			rootNode := renderable.ToTreeNode(IsDeprecated)

			dumpTree(t, rootNode, 0)

			// Each deprecatable kind must have a true, a string, and a non-deprecated version.
			// We'll check presence/absence for each.
			type Check struct {
				kind, name, reason string
				expect             bool // Should this node be present in the deprecated tree?
			}
			checks := []Check{
				// Module
				{"module", "mod-deprecated-true.js", "", true},
				{"module", "mod-deprecated-reason.js", "use mod-deprecated-true.js", true},
				{"module", "mod-ok.js", "", true},

				// Class
				{"class", "ClassDeprecatedTrue", "", true},
				{"class", "ClassDeprecatedReason", "use ClassDeprecatedTrue", true},
				{"class", "ClassOk", "", true},

				// Mixin
				{"mixin", "MixinDeprecatedTrue", "", true},
				{"mixin", "MixinDeprecatedReason", "use MixinDeprecatedTrue", true},
				{"mixin", "MixinOk", "", false},

				// Function
				{"function", "funcDeprecatedTrue", "", true},
				{"function", "funcDeprecatedReason", "use funcDeprecatedTrue", true},
				{"function", "funcOk", "", false},

				// Fields
				{"field", "fieldDeprecatedTrue", "", true},
				{"field", "fieldDeprecatedReason", "use something else", true},
				{"field", "fieldOk", "", false},

				// Methods
				{"method", "methodDeprecatedTrue", "", true},
				{"method", "methodDeprecatedReason", "do not use", true},
				{"method", "methodOk", "", false},

				// Attributes
				{"attribute", "attrDeprecatedTrue", "", true},
				{"attribute", "attrDeprecatedReason", "don't use", true},
				{"attribute", "attrOk", "", false},

				// Events
				{"event", "eventDeprecatedTrue", "", true},
				{"event", "eventDeprecatedReason", "old", true},
				{"event", "eventOk", "", false},

				// Slots
				{"slot", "slotDeprecatedTrue", "", true},
				{"slot", "slotDeprecatedReason", "use slot-new", true},
				{"slot", "slotOk", "", false},

				// CSS Parts
				{"cssPart", "partDeprecatedTrue", "", true},
				{"cssPart", "partDeprecatedReason", "not recommended", true},
				{"cssPart", "partOk", "", false},

				// CSS Properties
				{"cssProperty", "--cssDeprecatedTrue", "", true},
				{"cssProperty", "--cssDeprecatedReason", "renamed", true},
				{"cssProperty", "--cssOk", "", false},
			}

			// We'll flatten the tree to a list of text labels for easier searching.
			var flat []string
			var flatten func(node *pterm.TreeNode)
			flatten = func(node *pterm.TreeNode) {
				flat = append(flat, stripANSI(node.Text))
				for i := range node.Children {
					flatten(&node.Children[i])
				}
			}
			flatten(&rootNode)

			for _, check := range checks {
				var found bool
				for _, label := range flat {
					if check.reason != "" {
						// If a reason is provided, match on name and reason
						if strings.Contains(label, check.name) && strings.Contains(label, check.reason) {
							found = true
							break
						}
					} else {
						if strings.Contains(label, check.name) {
							found = true
							break
						}
					}
				}
				assert.Equalf(t, check.expect, found,
					"kind=%s name=%s reason=%s expect=%t", check.kind, check.name, check.reason, check.expect)
			}
		})
	})

	t.Run("VisualTreeOutput", func(t *testing.T) {
		manifestPath := filepath.Join("testdata", "deprecations.json")
		raw, err := os.ReadFile(manifestPath)
		if err != nil {
			t.Fatalf("failed to read manifest fixture: %v", err)
		}

		var pkg Package
		if err := json.Unmarshal(raw, &pkg); err != nil {
			t.Fatalf("failed to unmarshal manifest: %v", err)
		}

		renderable := NewRenderablePackage(&pkg)

		// To visually inspect output for debugging
		rootNode := renderable.ToTreeNode(IsDeprecated)
		s, err := pterm.DefaultTree.WithRoot(rootNode).Srender()
		t.Log(s)
		assert.NoError(t, err)
	})
}
