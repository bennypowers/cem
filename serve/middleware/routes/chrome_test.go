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
	"flag"
	"html/template"
	"os"
	"path/filepath"
	"testing"
)

var update = flag.Bool("update", false, "update golden files")

// TestChromeRendering_BasicDemo verifies chrome template wraps demo HTML
func TestChromeRendering_BasicDemo(t *testing.T) {
	// Read demo partial
	demoPath := filepath.Join("testdata", "chrome-rendering", "basic-demo.html")
	demoHTML, err := os.ReadFile(demoPath)
	if err != nil {
		t.Fatalf("Failed to read demo fixture: %v", err)
	}

	// Render chrome with demo
	rendered, err := renderDemoChrome(ChromeData{
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

	if *update {
		err := os.WriteFile(goldenPath, []byte(rendered), 0644)
		if err != nil {
			t.Fatalf("Failed to update golden file: %v", err)
		}
		t.Log("Updated golden file:", goldenPath)
		return
	}

	expected, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("Failed to read golden file: %v", err)
	}

	if rendered != string(expected) {
		t.Errorf("Rendered chrome does not match golden file.\nGot:\n%s\n\nExpected:\n%s", rendered, string(expected))
		t.Log("Run 'make update' to update golden files")
	}
}

// TestChromeRendering_NoKnobs verifies chrome without knobs
func TestChromeRendering_NoKnobs(t *testing.T) {
	demoPath := filepath.Join("testdata", "chrome-rendering", "basic-demo.html")
	demoHTML, err := os.ReadFile(demoPath)
	if err != nil {
		t.Fatalf("Failed to read demo fixture: %v", err)
	}

	rendered, err := renderDemoChrome(ChromeData{
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

	if *update {
		err := os.WriteFile(goldenPath, []byte(rendered), 0644)
		if err != nil {
			t.Fatalf("Failed to update golden file: %v", err)
		}
		t.Log("Updated golden file:", goldenPath)
		return
	}

	expected, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("Failed to read golden file: %v", err)
	}

	if rendered != string(expected) {
		t.Errorf("Rendered chrome does not match golden file.\nGot:\n%s\n\nExpected:\n%s", rendered, string(expected))
		t.Log("Run 'make update' to update golden files")
	}
}

// TestChromeRendering_ShadowMode verifies ?shadow=true wraps demo
func TestChromeRendering_ShadowMode(t *testing.T) {
	demoPath := filepath.Join("testdata", "chrome-rendering", "basic-demo.html")
	demoHTML, err := os.ReadFile(demoPath)
	if err != nil {
		t.Fatalf("Failed to read demo fixture: %v", err)
	}

	rendered, err := renderDemoChrome(ChromeData{
		TagName:    "my-element",
		DemoTitle:  "Basic Example",
		DemoHTML:   template.HTML(demoHTML),
		ImportMap:  "{}",
		ShadowMode: true,
	})
	if err != nil {
		t.Fatalf("Failed to render chrome: %v", err)
	}

	goldenPath := filepath.Join("testdata", "chrome-rendering", "expected-shadow-mode.html")

	if *update {
		err := os.WriteFile(goldenPath, []byte(rendered), 0644)
		if err != nil {
			t.Fatalf("Failed to update golden file: %v", err)
		}
		t.Log("Updated golden file:", goldenPath)
		return
	}

	expected, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("Failed to read golden file: %v", err)
	}

	if rendered != string(expected) {
		t.Errorf("Rendered chrome does not match golden file.\nGot:\n%s\n\nExpected:\n%s", rendered, string(expected))
		t.Log("Run 'make update' to update golden files")
	}
}

// TestChromeRendering_MarkdownDescription verifies GFM rendering
func TestChromeRendering_MarkdownDescription(t *testing.T) {
	demoHTML := `<my-element></my-element>`

	rendered, err := renderDemoChrome(ChromeData{
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

	if *update {
		err := os.WriteFile(goldenPath, []byte(rendered), 0644)
		if err != nil {
			t.Fatalf("Failed to update golden file: %v", err)
		}
		t.Log("Updated golden file:", goldenPath)
		return
	}

	expected, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("Failed to read golden file: %v", err)
	}

	if rendered != string(expected) {
		t.Errorf("Rendered chrome does not match golden file.\nGot:\n%s\n\nExpected:\n%s", rendered, string(expected))
		t.Log("Run 'make update' to update golden files")
	}
}

// TestChromeRendering_WithNavigation verifies navigation drawer rendering
func TestChromeRendering_WithNavigation(t *testing.T) {
	// Read demo partial
	demoPath := filepath.Join("testdata", "chrome-rendering", "basic-demo.html")
	demoHTML, err := os.ReadFile(demoPath)
	if err != nil {
		t.Fatalf("Failed to read demo fixture: %v", err)
	}

	// Read manifest with multiple elements
	manifestPath := filepath.Join("testdata", "chrome-rendering", "manifest-with-navigation.json")
	manifestBytes, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatalf("Failed to read manifest fixture: %v", err)
	}

	// Build navigation HTML using the same function as production code
	navigationHTML, packageName, err := BuildSinglePackageNavigation(manifestBytes, "test-package")
	if err != nil {
		t.Fatalf("Failed to build navigation: %v", err)
	}

	// Render chrome with navigation
	rendered, err := renderDemoChrome(ChromeData{
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

	if *update {
		err := os.WriteFile(goldenPath, []byte(rendered), 0644)
		if err != nil {
			t.Fatalf("Failed to update golden file: %v", err)
		}
		t.Log("Updated golden file:", goldenPath)
		return
	}

	expected, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("Failed to read golden file: %v", err)
	}

	if rendered != string(expected) {
		t.Errorf("Rendered chrome does not match golden file.\nGot:\n%s\n\nExpected:\n%s", rendered, string(expected))
		t.Log("Run 'make update' to update golden files")
	}
}
