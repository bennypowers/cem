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
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
)

// testTemplates creates a template registry for testing (with nil context)
func testTemplates() *TemplateRegistry {
	return NewTemplateRegistry(nil)
}

// TestChromeRendering_BasicDemo verifies chrome template wraps demo HTML
func TestChromeRendering_BasicDemo(t *testing.T) {
	// Read demo partial
	demoHTML := testutil.LoadFixtureFile(t, filepath.Join("chrome-rendering", "basic-demo.html"))

	// Render chrome with demo
	rendered, err := renderDemo(testTemplates(), nil, ChromeData{
		TagName:      "my-element",
		DemoTitle:    "Basic Example",
		DemoHTML:     template.HTML(demoHTML),
		EnabledKnobs: "attributes properties slots css-properties",
		ImportMap:    "{}",
		Description:  "A simple example showing basic usage",
	})
	if err != nil {
		t.Fatalf("Failed to render chrome: %v", err)
	}

	// Compare with golden file
	goldenPath := filepath.Join("testdata", "chrome-rendering", "expected-basic.html")

	if *testutil.Update {
		err := os.WriteFile(goldenPath, []byte(rendered), 0644)
		if err != nil {
			t.Fatalf("Failed to update golden file: %v", err)
		}
		t.Log("Updated golden file:", goldenPath)
		return
	}

	expected := testutil.LoadFixtureFile(t, filepath.Join("chrome-rendering", "expected-basic.html"))

	if rendered != string(expected) {
		t.Errorf("Rendered chrome does not match golden file.\nGot:\n%s\n\nExpected:\n%s", rendered, string(expected))
		t.Log("Run 'make update' to update golden files")
	}
}

// TestChromeRendering_NoKnobs verifies chrome without knobs
func TestChromeRendering_NoKnobs(t *testing.T) {
	demoHTML := testutil.LoadFixtureFile(t, filepath.Join("chrome-rendering", "basic-demo.html"))

	rendered, err := renderDemo(testTemplates(), nil, ChromeData{
		TagName:      "my-element",
		DemoTitle:    "Basic Example",
		DemoHTML:     template.HTML(demoHTML),
		EnabledKnobs: "", // No knobs
		ImportMap:    "{}",
	})
	if err != nil {
		t.Fatalf("Failed to render chrome: %v", err)
	}

	goldenPath := filepath.Join("testdata", "chrome-rendering", "expected-no-knobs.html")

	if *testutil.Update {
		err := os.WriteFile(goldenPath, []byte(rendered), 0644)
		if err != nil {
			t.Fatalf("Failed to update golden file: %v", err)
		}
		t.Log("Updated golden file:", goldenPath)
		return
	}

	expected := testutil.LoadFixtureFile(t, filepath.Join("chrome-rendering", "expected-no-knobs.html"))

	if rendered != string(expected) {
		t.Errorf("Rendered chrome does not match golden file.\nGot:\n%s\n\nExpected:\n%s", rendered, string(expected))
		t.Log("Run 'make update' to update golden files")
	}
}

// TestChromeRendering_ShadowMode verifies ?shadow=true wraps demo
func TestChromeRendering_ShadowMode(t *testing.T) {
	demoHTML := testutil.LoadFixtureFile(t, filepath.Join("chrome-rendering", "basic-demo.html"))

	rendered, err := renderDemo(testTemplates(), nil, ChromeData{
		TagName:       "my-element",
		DemoTitle:     "Basic Example",
		DemoHTML:      template.HTML(demoHTML),
		ImportMap:     "{}",
		RenderingMode: "shadow",
	})
	if err != nil {
		t.Fatalf("Failed to render chrome: %v", err)
	}

	goldenPath := filepath.Join("testdata", "chrome-rendering", "expected-shadow-mode.html")

	if *testutil.Update {
		err := os.WriteFile(goldenPath, []byte(rendered), 0644)
		if err != nil {
			t.Fatalf("Failed to update golden file: %v", err)
		}
		t.Log("Updated golden file:", goldenPath)
		return
	}

	expected := testutil.LoadFixtureFile(t, filepath.Join("chrome-rendering", "expected-shadow-mode.html"))

	if rendered != string(expected) {
		t.Errorf("Rendered chrome does not match golden file.\nGot:\n%s\n\nExpected:\n%s", rendered, string(expected))
		t.Log("Run 'make update' to update golden files")
	}
}

// TestChromeRendering_MarkdownDescription verifies GFM rendering
func TestChromeRendering_MarkdownDescription(t *testing.T) {
	demoHTML := `<my-element></my-element>`

	rendered, err := renderDemo(testTemplates(), nil, ChromeData{
		TagName:     "my-element",
		DemoTitle:   "Test",
		DemoHTML:    template.HTML(demoHTML),
		ImportMap:   "{}",
		Description: "This is **bold** and this is *italic*\n\n- Item 1\n- Item 2",
	})
	if err != nil {
		t.Fatalf("Failed to render chrome: %v", err)
	}

	goldenPath := filepath.Join("testdata", "chrome-rendering", "expected-markdown.html")

	if *testutil.Update {
		err := os.WriteFile(goldenPath, []byte(rendered), 0644)
		if err != nil {
			t.Fatalf("Failed to update golden file: %v", err)
		}
		t.Log("Updated golden file:", goldenPath)
		return
	}

	expected := testutil.LoadFixtureFile(t, filepath.Join("chrome-rendering", "expected-markdown.html"))

	if rendered != string(expected) {
		t.Errorf("Rendered chrome does not match golden file.\nGot:\n%s\n\nExpected:\n%s", rendered, string(expected))
		t.Log("Run 'make update' to update golden files")
	}
}

// TestChromeRendering_WithNavigation verifies navigation drawer rendering
func TestChromeRendering_WithNavigation(t *testing.T) {
	// Read demo partial
	demoHTML := testutil.LoadFixtureFile(t, filepath.Join("chrome-rendering", "basic-demo.html"))

	// Read manifest with multiple elements
	manifestBytes := testutil.LoadFixtureFile(t, filepath.Join("chrome-rendering", "manifest-with-navigation.json"))

	// Build navigation HTML using the same function as production code
	navigationHTML, packageName, err := BuildSinglePackageNavigation(testTemplates(), manifestBytes, "test-package")
	if err != nil {
		t.Fatalf("Failed to build navigation: %v", err)
	}

	// Render chrome with navigation
	rendered, err := renderDemo(testTemplates(), nil, ChromeData{
		TagName:        "demo-subset-one",
		DemoTitle:      "Demo with Navigation",
		DemoHTML:       template.HTML(demoHTML),
		EnabledKnobs:   "attributes properties",
		ImportMap:      "{}",
		Description:    "Testing navigation drawer rendering",
		PackageName:    packageName,
		NavigationHTML: navigationHTML,
	})
	if err != nil {
		t.Fatalf("Failed to render chrome: %v", err)
	}

	goldenPath := filepath.Join("testdata", "chrome-rendering", "expected-with-navigation.html")

	if *testutil.Update {
		err := os.WriteFile(goldenPath, []byte(rendered), 0644)
		if err != nil {
			t.Fatalf("Failed to update golden file: %v", err)
		}
		t.Log("Updated golden file:", goldenPath)
		return
	}

	expected := testutil.LoadFixtureFile(t, filepath.Join("chrome-rendering", "expected-with-navigation.html"))

	if rendered != string(expected) {
		t.Errorf("Rendered chrome does not match golden file.\nGot:\n%s\n\nExpected:\n%s", rendered, string(expected))
		t.Log("Run 'make update' to update golden files")
	}
}

// TestChromeRendering_Chromeless verifies chromeless mode renders minimal HTML
func TestChromeRendering_Chromeless(t *testing.T) {
	// Read demo partial
	demoHTML := testutil.LoadFixtureFile(t, filepath.Join("chrome-rendering", "basic-demo.html"))

	// Render in chromeless mode
	rendered, err := renderDemo(testTemplates(), nil, ChromeData{
		TagName:       "my-element",
		DemoTitle:     "Chromeless Example",
		DemoHTML:      template.HTML(demoHTML),
		ImportMap:     `{"imports":{"lit":"https://cdn.jsdelivr.net/npm/lit@3/index.js"}}`,
		RenderingMode: "chromeless",
	})
	if err != nil {
		t.Fatalf("Failed to render chromeless: %v", err)
	}

	// Compare with golden file
	goldenPath := filepath.Join("testdata", "chrome-rendering", "expected-chromeless.html")

	if *testutil.Update {
		err := os.WriteFile(goldenPath, []byte(rendered), 0644)
		if err != nil {
			t.Fatalf("Failed to update golden file: %v", err)
		}
		t.Log("Updated golden file:", goldenPath)
		return
	}

	expected := testutil.LoadFixtureFile(t, filepath.Join("chrome-rendering", "expected-chromeless.html"))

	if rendered != string(expected) {
		t.Errorf("Rendered chromeless does not match golden file.\nGot:\n%s\n\nExpected:\n%s", rendered, string(expected))
		t.Log("Run 'make update' to update golden files")
	}
}

// TestChromeRendering_NilManifest tests that chrome renders without crashing when Manifest is nil
// This is the case in workspace mode where there's no single manifest for the whole workspace
func TestChromeRendering_NilManifest(t *testing.T) {
	demoHTML := `<my-element id="example"></my-element>`

	// Render chrome with nil manifest (workspace mode scenario)
	rendered, err := renderDemo(testTemplates(), nil, ChromeData{
		TagName:     "my-element",
		DemoTitle:   "Workspace Mode Demo",
		DemoHTML:    template.HTML(demoHTML),
		ImportMap:   "{}",
		Description: "Testing nil manifest handling",
		PackageName: "workspace-package",
		Manifest:    nil, // Explicitly nil - this should not crash
	})
	if err != nil {
		t.Fatalf("Failed to render chrome with nil manifest: %v", err)
	}

	// Verify it rendered something
	if len(rendered) == 0 {
		t.Error("Rendered chrome is empty")
	}

	// Verify essential elements are present
	if !strings.Contains(rendered, "my-element") {
		t.Error("Rendered chrome doesn't contain element tag name")
	}
	if !strings.Contains(rendered, "Workspace Mode Demo") {
		t.Error("Rendered chrome doesn't contain demo title")
	}

	// Verify it doesn't have manifest tree element (should be skipped, though modulepreload in head is OK)
	if strings.Contains(rendered, `<pf-v6-tree-view`) {
		t.Error("Rendered chrome contains <pf-v6-tree-view> element when manifest is nil")
	}
}
