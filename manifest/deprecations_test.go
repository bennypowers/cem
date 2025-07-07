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

func TestDeprecationMatrixTreeFiltering(t *testing.T) {
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
		rootNode := renderable.ToTreeNode(func(r Renderable) bool {
			return isDeprecated(r)
		})

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
			{"module", "mod-ok.js", "", false},

			// Class
			{"class", "ClassDeprecatedTrue", "", true},
			{"class", "ClassDeprecatedReason", "use ClassDeprecatedTrue", true},
			{"class", "ClassOk", "", false},

			// Mixin
			{"mixin", "MixinDeprecatedTrue", "", true},
			{"mixin", "MixinDeprecatedReason", "use MixinDeprecatedTrue", true},
			{"mixin", "MixinOk", "", false},

			// Function
			{"function", "funcDeprecatedTrue", "", true},
			{"function", "funcDeprecatedReason", "use funcDeprecatedTrue", true},
			{"function", "funcOk", "", false},

			// Custom Element class (by name)
			{"custom-element", "CETrue", "", true},
			{"custom-element", "CEReason", "use CETrue", true},
			{"custom-element", "CEOk", "", false},

			// Custom Element Exports
			{"custom-element-definition", "ce-true", "", true},
			{"custom-element-definition", "ce-reason", "use ce-true", true},
			{"custom-element-definition", "ce-ok", "", false},

			// Fields
			{"field", "fieldDeprecatedTrue", "", true},
			{"field", "fieldDeprecatedReason", "use something else", true},
			{"field", "fieldOk", "", false},

			// Methods
			{"method", "methodDeprecatedTrue", "", true},
			{"method", "methodDeprecatedReason", "do not use", true},
			{"method", "methodOk", "", false},

			// Properties
			{"property", "propDeprecatedTrue", "", true},
			{"property", "propDeprecatedReason", "moved", true},
			{"property", "propOk", "", false},

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

	t.Run("VisualTreeOutput", func(t *testing.T) {
		// To visually inspect output for debugging
		rootNode := renderable.ToTreeNode(isDeprecated)
		_, err := pterm.DefaultTree.WithRoot(rootNode).Srender()
		assert.NoError(t, err)
	})
}

func isDeprecated(r Renderable) bool {
	// Assumes each Renderable has a Deprecated() (bool, string) method or similar
	switch v := r.(type) {
	case interface{ Deprecated() (bool, string) }:
		dep, _ := v.Deprecated()
		return dep
	default:
		return false
	}
}

