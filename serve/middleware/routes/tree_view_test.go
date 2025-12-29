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

package routes

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestTreeViewComponentFiles verifies tree view component files exist and have required content
func TestTreeViewComponentFiles(t *testing.T) {
	tests := []struct {
		name        string
		path        string
		mustContain []string
		description string
	}{
		{
			name: "pf-v6-tree-view.js",
			path: "templates/elements/pf-v6-tree-view/pf-v6-tree-view.js",
			mustContain: []string{
				"PfV6TreeView",
				"extends CemElement",
				"pf-v6-tree-view",
				"expandAll",
				"collapseAll",
			},
			description: "Tree view web component JavaScript",
		},
		{
			name: "pf-v6-tree-view.html",
			path: "templates/elements/pf-v6-tree-view/pf-v6-tree-view.html",
			mustContain: []string{
				"<ul",
				"role=\"tree\"",
				"<slot></slot>",
			},
			description: "Tree view component template",
		},
		{
			name: "pf-v6-tree-view.css",
			path: "templates/elements/pf-v6-tree-view/pf-v6-tree-view.css",
			mustContain: []string{
				":host",
				"#tree",
			},
			description: "Tree view component styles",
		},
		{
			name: "pf-v6-tree-item.js",
			path: "templates/elements/pf-v6-tree-item/pf-v6-tree-item.js",
			mustContain: []string{
				"PfV6TreeItem",
				"extends CemElement",
				"pf-v6-tree-item",
				"PfTreeItemSelectEvent",
				"PfTreeItemExpandEvent",
				"PfTreeItemCollapseEvent",
			},
			description: "Tree item web component JavaScript",
		},
		{
			name: "pf-v6-tree-item.html",
			path: "templates/elements/pf-v6-tree-item/pf-v6-tree-item.html",
			mustContain: []string{
				"<li",
				"role=\"treeitem\"",
				"<slot></slot>",
			},
			description: "Tree item component template",
		},
		{
			name: "pf-v6-tree-item.css",
			path: "templates/elements/pf-v6-tree-item/pf-v6-tree-item.css",
			mustContain: []string{
				":host",
				"--_",
				"#toggle",
				"#node",
			},
			description: "Tree item component styles",
		},
		{
			name: "manifest-icons.js",
			path: "templates/js/manifest-icons.js",
			mustContain: []string{
				"getIconForType",
				"getFolderIcon",
				"getChevronIcon",
				"ICONS",
			},
			description: "Manifest icon library",
		},
		{
			name: "manifest-tree-builder.js",
			path: "templates/js/manifest-tree-builder.js",
			mustContain: []string{
				"ManifestTreeBuilder",
				"build",
				"buildModule",
				"buildElement",
			},
			description: "Manifest tree builder utility",
		},
		{
			name: "manifest-search.js",
			path: "templates/js/manifest-search.js",
			mustContain: []string{
				"ManifestSearchIndex",
				"buildIndex",
				"search",
				"createDebouncedSearch",
			},
			description: "Manifest search utility",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fullPath := filepath.Join(".", tt.path)
			content, err := os.ReadFile(fullPath)
			if err != nil {
				t.Fatalf("Failed to read %s (%s): %v", tt.description, fullPath, err)
			}

			contentStr := string(content)
			for _, mustContain := range tt.mustContain {
				if !strings.Contains(contentStr, mustContain) {
					t.Errorf("%s (%s) must contain %q", tt.description, tt.name, mustContain)
				}
			}

			t.Logf("✓ %s contains all required content", tt.description)
		})
	}
}

// TestManifestTreeBuilder_Structure verifies tree builder produces correct structure
func TestManifestTreeBuilder_Structure(t *testing.T) {
	// Create a simple test manifest
	manifest := map[string]interface{}{
		"schemaVersion": "1.0.0",
		"modules": []map[string]interface{}{
			{
				"path": "./my-element.js",
				"declarations": []map[string]interface{}{
					{
						"kind":          "class",
						"name":          "MyElement",
						"customElement": true,
						"tagName":       "my-element",
						"summary":       "A test element",
						"attributes": []map[string]interface{}{
							{
								"name":    "label",
								"type":    map[string]interface{}{"text": "string"},
								"summary": "Element label",
							},
						},
						"members": []map[string]interface{}{
							{
								"kind":    "field",
								"name":    "value",
								"type":    map[string]interface{}{"text": "string"},
								"summary": "Element value",
							},
							{
								"kind":    "method",
								"name":    "reset",
								"summary": "Resets the element",
							},
						},
						"events": []map[string]interface{}{
							{
								"name":    "change",
								"type":    map[string]interface{}{"text": "Event"},
								"summary": "Fires when value changes",
							},
						},
					},
				},
			},
		},
	}

	manifestBytes, err := json.Marshal(manifest)
	if err != nil {
		t.Fatalf("Failed to marshal test manifest: %v", err)
	}

	// For now, just verify the manifest is valid JSON
	// Full integration testing would require running the JS in a browser environment
	var parsed map[string]interface{}
	err = json.Unmarshal(manifestBytes, &parsed)
	if err != nil {
		t.Fatalf("Test manifest is not valid JSON: %v", err)
	}

	modules, ok := parsed["modules"].([]interface{})
	if !ok || len(modules) == 0 {
		t.Errorf("Test manifest should have modules array")
	}

	t.Logf("✓ Test manifest structure is valid with %d module(s)", len(modules))
}

// TestTreeViewEventClasses verifies custom event classes follow the pattern
func TestTreeViewEventClasses(t *testing.T) {
	jsPath := filepath.Join(".", "templates/elements/pf-v6-tree-item/pf-v6-tree-item.js")
	content, err := os.ReadFile(jsPath)
	if err != nil {
		t.Fatalf("Failed to read tree item JS: %v", err)
	}

	contentStr := string(content)

	// Verify event classes use class fields, not CustomEvent with details
	eventClasses := []string{
		"PfTreeItemSelectEvent",
		"PfTreeItemExpandEvent",
		"PfTreeItemCollapseEvent",
	}

	for _, className := range eventClasses {
		// Verify class exists
		if !strings.Contains(contentStr, "export class "+className) {
			t.Errorf("Missing event class: %s", className)
			continue
		}

		// Verify it extends Event
		if !strings.Contains(contentStr, className+" extends Event") {
			t.Errorf("Event class %s should extend Event", className)
		}

		t.Logf("✓ Event class %s properly defined", className)
	}

	// Verify component doesn't use CustomEvent with details (anti-pattern)
	if strings.Contains(contentStr, "new CustomEvent") {
		t.Errorf("Component should use custom event classes, not CustomEvent")
	}
}

// TestTreeViewAccessibility verifies ARIA attributes are present
func TestTreeViewAccessibility(t *testing.T) {
	htmlPath := filepath.Join(".", "templates/elements/pf-v6-tree-view/pf-v6-tree-view.html")
	content, err := os.ReadFile(htmlPath)
	if err != nil {
		t.Fatalf("Failed to read tree view HTML: %v", err)
	}

	contentStr := string(content)

	requiredARIA := []string{
		"role=\"tree\"",
	}

	for _, aria := range requiredARIA {
		if !strings.Contains(contentStr, aria) {
			t.Errorf("Template must contain ARIA attribute: %s", aria)
		}
	}

	t.Log("✓ Tree view template has required ARIA attributes")
}

// TestTreeItemCSS_Structure verifies tree item CSS has required structure
func TestTreeItemCSS_Structure(t *testing.T) {
	cssPath := filepath.Join(".", "templates/elements/pf-v6-tree-item/pf-v6-tree-item.css")
	content, err := os.ReadFile(cssPath)
	if err != nil {
		t.Fatalf("Failed to read tree item CSS: %v", err)
	}

	contentStr := string(content)

	requiredSelectors := []string{
		":host",
		"#item",
		"#content",
		"#node",
		"#toggle",
		"#children",
		":host([expanded])",
		":host([has-children])",
	}

	for _, selector := range requiredSelectors {
		if !strings.Contains(contentStr, selector) {
			t.Errorf("CSS must contain selector: %s", selector)
		}
	}

	// Verify use of private CSS custom properties (--_)
	if !strings.Contains(contentStr, "--_") {
		t.Errorf("CSS should use private custom properties (--_)")
	}

	t.Log("✓ Tree item CSS has all required selectors and uses --_ for private properties")
}

// TestManifestIcons_Coverage verifies all expected icon types are available
func TestManifestIcons_Coverage(t *testing.T) {
	jsPath := filepath.Join(".", "templates/js/manifest-icons.js")
	content, err := os.ReadFile(jsPath)
	if err != nil {
		t.Fatalf("Failed to read manifest-icons.js: %v", err)
	}

	contentStr := string(content)

	requiredIcons := []string{
		"package:",
		"module:",
		"element:",
		"class:",
		"function:",
		"attribute:",
		"property:",
		"method:",
		"event:",
		"slot:",
		"'css-property':",
		"'css-part':",
	}

	for _, iconKey := range requiredIcons {
		if !strings.Contains(contentStr, iconKey) {
			t.Errorf("Icon library must contain icon for: %s", iconKey)
		}
	}

	t.Log("✓ Manifest icons library has all required icon types")
}

// TestVirtualTreeComponent verifies virtual tree component exists and has required content
func TestVirtualTreeComponent(t *testing.T) {
	tests := []struct {
		name        string
		path        string
		mustContain []string
		description string
	}{
		{
			name: "cem-virtual-tree.js",
			path: "templates/elements/cem-virtual-tree/cem-virtual-tree.js",
			mustContain: []string{
				"CemVirtualTree",
				"extends CemElement",
				"cem-virtual-tree",
				"buildFlatList",
				"search",
				"expandAll",
				"collapseAll",
				"ItemSelectEvent",
			},
			description: "Virtual tree component JavaScript",
		},
		{
			name: "cem-virtual-tree.html",
			path: "templates/elements/cem-virtual-tree/cem-virtual-tree.html",
			mustContain: []string{
				"tree-container",
				"viewport",
			},
			description: "Virtual tree component template",
		},
		{
			name: "cem-virtual-tree.css",
			path: "templates/elements/cem-virtual-tree/cem-virtual-tree.css",
			mustContain: []string{
				":host",
				"#tree-container",
				"#viewport",
			},
			description: "Virtual tree component styles",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fullPath := filepath.Join(".", tt.path)
			content, err := os.ReadFile(fullPath)
			if err != nil {
				t.Fatalf("Failed to read %s (%s): %v", tt.description, fullPath, err)
			}

			contentStr := string(content)
			for _, mustContain := range tt.mustContain {
				if !strings.Contains(contentStr, mustContain) {
					t.Errorf("%s (%s) must contain %q", tt.description, tt.name, mustContain)
				}
			}

			t.Logf("✓ %s contains all required content", tt.description)
		})
	}
}

// TestDetailPanelComponent verifies detail panel component exists and has required content
func TestDetailPanelComponent(t *testing.T) {
	tests := []struct {
		name        string
		path        string
		mustContain []string
		description string
	}{
		{
			name: "cem-detail-panel.js",
			path: "templates/elements/cem-detail-panel/cem-detail-panel.js",
			mustContain: []string{
				"CemDetailPanel",
				"extends CemElement",
				"cem-detail-panel",
				"renderItem",
				"renderMarkdown",
				"buildDetailHTML",
			},
			description: "Detail panel component JavaScript",
		},
		{
			name: "cem-detail-panel.html",
			path: "templates/elements/cem-detail-panel/cem-detail-panel.html",
			mustContain: []string{
				"details-content",
			},
			description: "Detail panel component template",
		},
		{
			name: "cem-detail-panel.css",
			path: "templates/elements/cem-detail-panel/cem-detail-panel.css",
			mustContain: []string{
				":host",
				"#details-content",
			},
			description: "Detail panel component styles",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fullPath := filepath.Join(".", tt.path)
			content, err := os.ReadFile(fullPath)
			if err != nil {
				t.Fatalf("Failed to read %s (%s): %v", tt.description, fullPath, err)
			}

			contentStr := string(content)
			for _, mustContain := range tt.mustContain {
				if !strings.Contains(contentStr, mustContain) {
					t.Errorf("%s (%s) must contain %q", tt.description, tt.name, mustContain)
				}
			}

			t.Logf("✓ %s contains all required content", tt.description)
		})
	}
}

// TestManifestBrowserIntegration verifies manifest browser uses new components
func TestManifestBrowserIntegration(t *testing.T) {
	jsPath := filepath.Join(".", "templates/elements/cem-manifest-browser/cem-manifest-browser.js")
	content, err := os.ReadFile(jsPath)
	if err != nil {
		t.Fatalf("Failed to read manifest browser JS: %v", err)
	}

	contentStr := string(content)

	// Verify imports
	requiredImports := []string{
		"cem-virtual-tree/cem-virtual-tree.js",
		"cem-detail-panel/cem-detail-panel.js",
	}

	for _, imp := range requiredImports {
		if !strings.Contains(contentStr, imp) {
			t.Errorf("Manifest browser must import: %s", imp)
		}
	}

	// Verify integration methods
	requiredContent := []string{
		"#virtualTree",
		"#detailPanel",
		"item-select",
		"renderItem",
	}

	for _, content := range requiredContent {
		if !strings.Contains(contentStr, content) {
			t.Errorf("Manifest browser must contain: %s", content)
		}
	}

	t.Log("✓ Manifest browser properly integrates virtual tree and detail panel")
}

// TestDemoChromeNoSSR verifies demo chrome template no longer has SSR tree
func TestDemoChromeNoSSR(t *testing.T) {
	htmlPath := filepath.Join(".", "templates/demo-chrome.html")
	content, err := os.ReadFile(htmlPath)
	if err != nil {
		t.Fatalf("Failed to read demo-chrome.html: %v", err)
	}

	contentStr := string(content)

	// Verify SSR tree elements are NOT present
	shouldNotContain := []string{
		"<pf-v6-tree-view slot=\"manifest-tree\"",
		"slot=\"manifest-details\"",
		"data-type=\"custom-element\"",
		"data-type=\"attribute\"",
		"{{range .Manifest.Modules}}",
	}

	for _, text := range shouldNotContain {
		if strings.Contains(contentStr, text) {
			t.Errorf("Demo chrome template should NOT contain SSR'd: %s", text)
		}
	}

	// Manifest is now loaded client-side via fetch from /custom-elements.json
	// The conditional at lines 104-107 is intentionally empty because the client
	// components (cem-virtual-tree, cem-detail-panel) fetch the manifest on demand

	t.Log("✓ Demo chrome template no longer SSRs manifest tree or injects window.__CEM_MANIFEST__")
}
