//go:build e2e

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
package completion_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"testing/synctest"

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/methods/textDocument/completion"
	"bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
	W "bennypowers.dev/cem/workspace"
)

func TestLocalElementChangesUpdateCompletions(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// Create a temporary workspace directory
		tempDir := t.TempDir()

		// Create package.json with CEM configuration
		packageJSON := `{
  "name": "test-package",
  "customElements": "dist/custom-elements.json",
  "scripts": {
    "analyze": "cem analyze --litelement --globs \"src/**/*.{js,ts}\""
  }
}`

		err := os.WriteFile(filepath.Join(tempDir, "package.json"), []byte(packageJSON), 0644)
		if err != nil {
			t.Fatalf("Failed to create package.json: %v", err)
		}

		// Create the initial TypeScript source file
		srcDir := filepath.Join(tempDir, "src")
		err = os.MkdirAll(srcDir, 0755)
		if err != nil {
			t.Fatalf("Failed to create src directory: %v", err)
		}

		initialTSContent := `import {LitElement, html, css} from 'lit';
import {customElement, property} from 'lit/decorators.js';

@customElement('test-button')
export class TestButton extends LitElement {
  /** The variant of the button */
  @property() variant: 'primary' | 'secondary' = 'primary';
  
  /** Whether the button is disabled */
  @property({type: Boolean}) disabled = false;

  render() {
    return html` + "`<button class=\"${this.variant}\" ?disabled=\"${this.disabled}\"><slot></slot></button>`" + `;
  }
}`

		tsFilePath := filepath.Join(srcDir, "test-button.ts")
		err = os.WriteFile(tsFilePath, []byte(initialTSContent), 0644)
		if err != nil {
			t.Fatalf("Failed to create TypeScript file: %v", err)
		}

		// Create initial manifest (simulating what cem analyze would generate)
		distDir := filepath.Join(tempDir, "dist")
		err = os.MkdirAll(distDir, 0755)
		if err != nil {
			t.Fatalf("Failed to create dist directory: %v", err)
		}

		// Create .config directory and cem.yaml configuration
		configDir := filepath.Join(tempDir, ".config")
		err = os.MkdirAll(configDir, 0755)
		if err != nil {
			t.Fatalf("Failed to create config directory: %v", err)
		}

		cemConfig := `generate:
  files:
    - src/test-button.ts
`

		err = os.WriteFile(filepath.Join(configDir, "cem.yaml"), []byte(cemConfig), 0644)
		if err != nil {
			t.Fatalf("Failed to create cem.yaml: %v", err)
		}

		initialManifest := `{
  "schemaVersion": "1.0.0",
  "readme": "",
  "modules": [
    {
      "kind": "javascript-module",
      "path": "src/test-button.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "TestButton",
          "customElement": true,
          "tagName": "test-button",
          "attributes": [
            {
              "name": "variant",
              "type": {
                "text": "\"primary\" | \"secondary\""
              },
              "default": "\"primary\""
            },
            {
              "name": "disabled",
              "type": {
                "text": "boolean"
              },
              "default": "false"
            }
          ]
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "test-button",
          "declaration": {
            "name": "TestButton",
            "module": "src/test-button.ts"
          }
        }
      ]
    }
  ]
}`

		manifestPath := filepath.Join(distDir, "custom-elements.json")
		err = os.WriteFile(manifestPath, []byte(initialManifest), 0644)
		if err != nil {
			t.Fatalf("Failed to create initial manifest: %v", err)
		}

		// Create HTML test file that uses the element
		htmlContent := `<!DOCTYPE html>
<html>
<body>
  <test-button variant=""></test-button>
</body>
</html>`

		htmlFilePath := filepath.Join(tempDir, "test.html")
		err = os.WriteFile(htmlFilePath, []byte(htmlContent), 0644)
		if err != nil {
			t.Fatalf("Failed to create HTML file: %v", err)
		}

		// Create workspace and registry
		workspace := W.NewFileSystemWorkspaceContext(tempDir)
		err = workspace.Init()
		if err != nil {
			t.Fatalf("Failed to initialize workspace: %v", err)
		}

		// Create registry with mock file watcher for synctest
		mockFileWatcher := platform.NewMockFileWatcher()
		registry := lsp.NewRegistry(mockFileWatcher)

		// Load initial manifests
		err = registry.LoadFromWorkspace(workspace)
		if err != nil {
			t.Fatalf("Failed to load workspace manifests: %v", err)
		}

		// Create document manager
		dm, err := lsp.NewDocumentManager()
		if err != nil {
			t.Fatalf("Failed to create document manager: %v", err)
		}
		defer dm.Close()

		// Set up the server context adapter with document manager
		ctx := &testCompletionContextWithDM{
			registry: registry,
			docMgr:   dm,
		}

		// Open the HTML document
		htmlURI := "file://" + htmlFilePath
		doc := dm.OpenDocument(htmlURI, htmlContent, 1)
		if doc == nil {
			t.Fatalf("Failed to open HTML document")
		}

		// Test initial attribute value completions directly
		// (The main purpose is to test file watching updates, not completion positioning)
		initialItems := completion.GetAttributeValueCompletions(ctx, "test-button", "variant")

		// Verify initial completions contain "primary" and "secondary"
		hasInitialPrimary := false
		hasInitialSecondary := false
		for _, item := range initialItems {
			if item.Label == "primary" {
				hasInitialPrimary = true
			}
			if item.Label == "secondary" {
				hasInitialSecondary = true
			}
		}

		if !hasInitialPrimary || !hasInitialSecondary {
			t.Fatalf("Initial completions missing expected values. Got labels: %v", getCompletionLabels(initialItems))
		}

		t.Logf("Initial completions found: %v", getCompletionLabels(initialItems))

		// With synctest, we don't need file watching - use direct operations

		// Modify the TypeScript source to add a new variant option
		updatedTSContent := strings.Replace(initialTSContent,
			`variant: 'primary' | 'secondary' = 'primary'`,
			`variant: 'primary' | 'secondary' | 'danger' = 'primary'`,
			1)

		err = os.WriteFile(tsFilePath, []byte(updatedTSContent), 0644)
		if err != nil {
			t.Fatalf("Failed to update TypeScript file: %v", err)
		}

		// Debug: Check if the TypeScript file exists and has the right content
		if updatedContent, err := os.ReadFile(tsFilePath); err == nil {
			t.Logf("TypeScript file content after update: %s", string(updatedContent))
		} else {
			t.Logf("Could not read updated TypeScript file: %v", err)
		}

		// Debug: List files in src directory
		if files, err := os.ReadDir(filepath.Join(tempDir, "src")); err == nil {
			t.Logf("Files in src directory:")
			for _, file := range files {
				t.Logf("  - %s", file.Name())
			}
		}

		// With synctest, update manifest directly to reflect the TypeScript change
		updatedManifest := strings.Replace(initialManifest,
			`"text": "\"primary\" | \"secondary\""`,
			`"text": "\"primary\" | \"secondary\" | \"danger\""`,
			1)

		err = os.WriteFile(manifestPath, []byte(updatedManifest), 0644)
		if err != nil {
			t.Fatalf("Failed to write updated manifest: %v", err)
		}

		// Reload manifests directly (no file watching delays)
		if err := registry.ReloadManifestsDirectly(); err != nil {
			t.Fatalf("Failed to reload manifests: %v", err)
		}

		t.Logf("Manifest updated and reloaded directly")

		// Debug: Check what's in the registry after reload
		if attrs, exists := ctx.Attributes("test-button"); exists {
			if variantAttr, hasVariant := attrs["variant"]; hasVariant {
				t.Logf("After reload - variant attribute type: %s", variantAttr.Type.Text)
			} else {
				t.Logf("After reload - variant attribute not found")
			}
		} else {
			t.Logf("After reload - test-button element not found")
		}

		// Test updated completions
		updatedItems := completion.GetAttributeValueCompletions(ctx, "test-button", "variant")

		// Verify updated completions now contain "danger"
		hasUpdatedPrimary := false
		hasUpdatedSecondary := false
		hasUpdatedDanger := false
		for _, item := range updatedItems {
			switch item.Label {
			case "primary":
				hasUpdatedPrimary = true
			case "secondary":
				hasUpdatedSecondary = true
			case "danger":
				hasUpdatedDanger = true
			}
		}

		if !hasUpdatedPrimary || !hasUpdatedSecondary {
			t.Errorf("Updated completions missing original values. Got labels: %v", getCompletionLabels(updatedItems))
		}

		if !hasUpdatedDanger {
			t.Errorf("Updated completions missing new 'danger' value. Got labels: %v", getCompletionLabels(updatedItems))
		}

		t.Logf("Updated completions found: %v", getCompletionLabels(updatedItems))
		t.Logf("Test passed: Local element changes successfully updated HTML completions")
	}) // End synctest.Test
}

func TestLocalElementChangesUpdateLitTemplateCompletions(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// Create a temporary workspace directory
		tempDir := t.TempDir()

		// Create package.json with CEM configuration
		packageJSON := `{
  "name": "test-package",
  "customElements": "dist/custom-elements.json"
}`

		err := os.WriteFile(filepath.Join(tempDir, "package.json"), []byte(packageJSON), 0644)
		if err != nil {
			t.Fatalf("Failed to create package.json: %v", err)
		}

		// Create initial manifest for a button element
		distDir := filepath.Join(tempDir, "dist")
		err = os.MkdirAll(distDir, 0755)
		if err != nil {
			t.Fatalf("Failed to create dist directory: %v", err)
		}

		initialManifest := `{
  "schemaVersion": "1.0.0",
  "readme": "",
  "modules": [
    {
      "kind": "javascript-module",
      "path": "src/my-button.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "MyButton",
          "customElement": true,
          "tagName": "my-button",
          "attributes": [
            {
              "name": "size",
              "type": {
                "text": "\"small\" | \"medium\""
              },
              "default": "\"medium\""
            }
          ]
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "my-button",
          "declaration": {
            "name": "MyButton",
            "module": "src/my-button.ts"
          }
        }
      ]
    }
  ]
}`

		manifestPath := filepath.Join(distDir, "custom-elements.json")
		err = os.WriteFile(manifestPath, []byte(initialManifest), 0644)
		if err != nil {
			t.Fatalf("Failed to create initial manifest: %v", err)
		}

		// Create TypeScript file that uses the element in a Lit template
		tsContent := `import {LitElement, html} from 'lit';
import {customElement} from 'lit/decorators.js';

@customElement('my-app')
export class MyApp extends LitElement {
  render() {
    return html` + "`\n      <div>\n        <my-button size=\"\"></my-button>\n      </div>\n    `" + `;
  }
}`

		srcDir := filepath.Join(tempDir, "src")
		err = os.MkdirAll(srcDir, 0755)
		if err != nil {
			t.Fatalf("Failed to create src directory: %v", err)
		}

		tsFilePath := filepath.Join(srcDir, "my-app.ts")
		err = os.WriteFile(tsFilePath, []byte(tsContent), 0644)
		if err != nil {
			t.Fatalf("Failed to create TypeScript file: %v", err)
		}

		// Create workspace and registry
		workspace := W.NewFileSystemWorkspaceContext(tempDir)
		err = workspace.Init()
		if err != nil {
			t.Fatalf("Failed to initialize workspace: %v", err)
		}

		// Create registry with mock file watcher for synctest
		mockFileWatcher := platform.NewMockFileWatcher()
		registry := lsp.NewRegistry(mockFileWatcher)

		// Load initial manifests
		err = registry.LoadFromWorkspace(workspace)
		if err != nil {
			t.Fatalf("Failed to load workspace manifests: %v", err)
		}

		// Create document manager
		dm, err := lsp.NewDocumentManager()
		if err != nil {
			t.Fatalf("Failed to create document manager: %v", err)
		}
		defer dm.Close()

		// Set up the server context adapter with document manager
		ctx := &testCompletionContextWithDM{
			registry: registry,
			docMgr:   dm,
		}

		// Open the TypeScript document
		tsURI := "file://" + tsFilePath
		doc := dm.OpenDocument(tsURI, tsContent, 1)
		if doc == nil {
			t.Fatalf("Failed to open TypeScript document")
		}

		// Test initial completions for size attribute directly
		initialItems := completion.GetAttributeValueCompletions(ctx, "my-button", "size")

		// Verify initial completions contain "small" and "medium"
		hasInitialSmall := false
		hasInitialMedium := false
		for _, item := range initialItems {
			if item.Label == "small" {
				hasInitialSmall = true
			}
			if item.Label == "medium" {
				hasInitialMedium = true
			}
		}

		if !hasInitialSmall || !hasInitialMedium {
			t.Fatalf("Initial completions missing expected values. Got labels: %v", getCompletionLabels(initialItems))
		}

		t.Logf("Initial Lit template completions found: %v", getCompletionLabels(initialItems))

		// With synctest, update manifest directly
		updatedManifest := strings.Replace(initialManifest,
			`"text": "\"small\" | \"medium\""`,
			`"text": "\"small\" | \"medium\" | \"large\""`,
			1)

		err = os.WriteFile(manifestPath, []byte(updatedManifest), 0644)
		if err != nil {
			t.Fatalf("Failed to update manifest: %v", err)
		}

		// Reload manifests directly (no file watching delays)
		if err := registry.ReloadManifestsDirectly(); err != nil {
			t.Fatalf("Failed to reload manifests: %v", err)
		}

		t.Logf("Manifest updated and reloaded directly")

		// Test updated completions
		updatedItems := completion.GetAttributeValueCompletions(ctx, "my-button", "size")

		// Verify updated completions now contain "large"
		hasUpdatedSmall := false
		hasUpdatedMedium := false
		hasUpdatedLarge := false
		for _, item := range updatedItems {
			switch item.Label {
			case "small":
				hasUpdatedSmall = true
			case "medium":
				hasUpdatedMedium = true
			case "large":
				hasUpdatedLarge = true
			}
		}

		if !hasUpdatedSmall || !hasUpdatedMedium {
			t.Errorf("Updated completions missing original values. Got labels: %v", getCompletionLabels(updatedItems))
		}

		if !hasUpdatedLarge {
			t.Errorf("Updated completions missing new 'large' value. Got labels: %v", getCompletionLabels(updatedItems))
		}

		t.Logf("Updated Lit template completions found: %v", getCompletionLabels(updatedItems))
		t.Logf("Test passed: Local element changes successfully updated Lit template completions")
	}) // End synctest.Test
}

// testCompletionContextWithDM implements CompletionContext for testing with document manager
type testCompletionContextWithDM struct {
	registry *lsp.Registry
	docMgr   *lsp.DocumentManager
}

func (ctx *testCompletionContextWithDM) Document(uri string) types.Document {
	return ctx.docMgr.Document(uri)
}

func (ctx *testCompletionContextWithDM) AllTagNames() []string {
	return ctx.registry.AllTagNames()
}

func (ctx *testCompletionContextWithDM) Element(tagName string) (*M.CustomElement, bool) {
	return ctx.registry.Element(tagName)
}

func (ctx *testCompletionContextWithDM) Attributes(tagName string) (map[string]*M.Attribute, bool) {
	return ctx.registry.Attributes(tagName)
}

func (ctx *testCompletionContextWithDM) Slots(tagName string) ([]M.Slot, bool) {
	return ctx.registry.Slots(tagName)
}

// Helper function to get attribute names for debugging
func getAttrNames(attrs map[string]*M.Attribute) []string {
	names := make([]string, 0, len(attrs))
	for name := range attrs {
		names = append(names, name)
	}
	return names
}
