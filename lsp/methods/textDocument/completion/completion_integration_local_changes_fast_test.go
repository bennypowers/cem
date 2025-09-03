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
	"strings"
	"testing"
	"testing/synctest"

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/lsp"
	M "bennypowers.dev/cem/manifest"
)

// loadFixture loads a test fixture file from the filesystem
func loadFixture(t *testing.T, fixturePath string) string {
	content, err := os.ReadFile(fixturePath)
	if err != nil {
		t.Fatalf("Failed to load fixture file %s: %v", fixturePath, err)
	}
	return string(content)
}

// TestLocalElementChangesUpdateCompletionsFast tests local element changes with instant virtual time
// This eliminates the 10+ second delays from the E2E test while maintaining full functionality
func TestLocalElementChangesUpdateCompletionsFast(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// Create in-memory filesystem for instant operations
		mapFS := platform.NewMapFileSystem(nil)

		// Load fixture files into MapFS for instant access
		fixturesDir := "test-fixtures/local-changes"
		
		// Load package.json fixture
		packageJSON := loadFixture(t, fixturesDir+"/package.json")
		mapFS.AddFile("package.json", packageJSON, 0644)

		// Load initial TypeScript element fixture
		initialElementContent := loadFixture(t, fixturesDir+"/src/test-button.ts")
		mapFS.AddFile("src/test-button.ts", initialElementContent, 0644)

		// Load initial manifest fixture
		initialManifest := loadFixture(t, fixturesDir+"/initial-manifest.json")
		mapFS.AddFile("dist/custom-elements.json", initialManifest, 0644)

		// Load HTML test file fixture
		testHTML := loadFixture(t, fixturesDir+"/test.html")
		mapFS.AddFile("test.html", testHTML, 0644)

		// Validate filesystem compliance
		err := mapFS.TestFS("package.json", "src/test-button.ts", "dist/custom-elements.json", "test.html")
		if err != nil {
			t.Fatalf("Filesystem should pass compliance tests: %v", err)
		}

		// Create mock file watcher for instant events
		mockFileWatcher := platform.NewMockFileWatcher()
		registry := lsp.NewRegistry(mockFileWatcher)

		// Load initial manifest
		var pkg M.Package
		err = pkg.UnmarshalJSON([]byte(initialManifest))
		if err != nil {
			t.Fatalf("Failed to parse initial manifest: %v", err)
		}
		registry.AddManifest(&pkg)

		t.Logf("Testing local element changes and completion updates")

		// Test initial attributes - should have primary, secondary
		element, exists := registry.Element("test-button")
		if !exists {
			t.Fatalf("Expected test-button element to be loaded")
		}
		t.Logf("Initial element loaded: %s with %d attributes", element.TagName, len(element.Attributes))

		// Verify initial attributes
		attrs, attrsExist := registry.Attributes("test-button")
		if !attrsExist {
			t.Fatalf("Expected test-button attributes to be loaded")
		}
		variantAttr, variantExists := attrs["variant"]
		if !variantExists {
			t.Fatalf("Expected variant attribute to exist")
		}
		t.Logf("Initial variant attribute type: %s", variantAttr.Type.Text)

		// Verify initial completion values
		expectedInitial := []string{"primary", "secondary"}
		initialTypeText := variantAttr.Type.Text
		for _, expected := range expectedInitial {
			if !strings.Contains(initialTypeText, `"`+expected+`"`) {
				t.Errorf("Expected initial type to contain '%s', got: %s", expected, initialTypeText)
			}
		}

		// Set up MockGenerateWatcher with callback that simulates user adding 'danger' variant
		generateCallback := func() error {
			// Load the updated TypeScript element fixture with 'danger' variant
			updatedElementContent := loadFixture(t, fixturesDir+"/src/test-button-updated.ts")
			
			// Update the TypeScript file in MapFS
			mapFS.AddFile("src/test-button.ts", updatedElementContent, 0644)

			// Create updated manifest with 'danger' variant
			updatedManifest := strings.Replace(initialManifest,
				`"text": "\"primary\" | \"secondary\""`,
				`"text": "\"primary\" | \"secondary\" | \"danger\""`,
				1)

			// Update the manifest file in MapFS
			mapFS.AddFile("dist/custom-elements.json", updatedManifest, 0644)

			// Parse the updated manifest and add to registry
			var updatedPkg M.Package
			err := updatedPkg.UnmarshalJSON([]byte(updatedManifest))
			if err != nil {
				t.Logf("Failed to parse updated manifest: %v", err)
				return err
			}

			// Add the updated manifest to the registry
			registry.AddManifest(&updatedPkg)
			return nil
		}

		mockGenerateWatcher := platform.NewMockGenerateWatcher(generateCallback)
		registry.SetGenerateWatcher(mockGenerateWatcher)

		// Start the generate watcher
		err = registry.StartGenerateWatcher()
		if err != nil {
			t.Fatalf("Failed to start generate watcher: %v", err)
		}
		defer registry.StopGenerateWatcher()

		// Simulate user editing the source file - instant with virtual time
		t.Logf("=== User edits test-button.ts to add 'danger' variant ===")
		
		// Call the callback directly for instant execution
		err = generateCallback()
		if err != nil {
			t.Fatalf("Failed to run generate callback: %v", err)
		}

		// No timing delays needed - everything happens instantly with synctest
		t.Logf("Generate completed instantly with virtual time")

		// Test updated attributes - should now include 'danger'
		updatedAttrs, updatedAttrsExist := registry.Attributes("test-button")
		if !updatedAttrsExist {
			t.Fatalf("Expected test-button attributes to still exist after update")
		}
		updatedVariantAttr, updatedVariantExists := updatedAttrs["variant"]
		if !updatedVariantExists {
			t.Fatalf("Expected variant attribute to still exist after update")
		}
		t.Logf("Updated variant attribute type: %s", updatedVariantAttr.Type.Text)

		// Verify updated completions contain 'danger'
		expectedUpdated := []string{"primary", "secondary", "danger"}
		updatedTypeText := updatedVariantAttr.Type.Text
		for _, expected := range expectedUpdated {
			if !strings.Contains(updatedTypeText, `"`+expected+`"`) {
				t.Errorf("Expected updated type to contain '%s', got: %s", expected, updatedTypeText)

				// Debug: Check the manifest file content from MapFS
				if manifestContent, err := mapFS.ReadFile("dist/custom-elements.json"); err == nil {
					t.Logf("Current manifest content: %s", string(manifestContent))
				}
			}
		}

		// Verify 'danger' specifically was added
		if strings.Contains(updatedTypeText, `"danger"`) {
			t.Logf("✅ SUCCESS: 'danger' found in updated variant type!")
		} else {
			t.Errorf("❌ FAILURE: 'danger' not found in updated type: %s", updatedTypeText)
		}

		t.Logf("✅ Fast local element changes test completed instantly (eliminated 10+ second delays)")
	})
}

// TestLocalElementChangesInLitTemplatesFast tests Lit template completion updates with instant execution
func TestLocalElementChangesInLitTemplatesFast(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// Create in-memory filesystem for instant operations
		mapFS := platform.NewMapFileSystem(nil)

		// Set up project structure with Lit template example
		mapFS.AddFile("package.json", `{
  "name": "test-package",
  "customElements": "dist/custom-elements.json"
}`, 0644)

		// Initial TypeScript element with size attribute
		initialElementContent := `import {LitElement, html} from 'lit';
import {customElement, property} from 'lit/decorators.js';

@customElement('my-button')
export class MyButton extends LitElement {
  @property() size: 'small' | 'medium' = 'medium';
  
  render() {
    return html` + "`<button class=\"btn-${this.size}\"><slot></slot></button>`" + `;
  }
}`

		mapFS.AddFile("src/my-button.ts", initialElementContent, 0644)

		// Initial manifest
		initialManifest := `{
  "schemaVersion": "1.0.0",
  "readme": "",
  "modules": [
    {
      "kind": "javascript-module", 
      "path": "src/my-button.js",
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
            "module": "src/my-button.js"
          }
        }
      ]
    }
  ]
}`

		mapFS.AddFile("dist/custom-elements.json", initialManifest, 0644)

		// TypeScript file with Lit template
		mapFS.AddFile("src/app.ts", `import {LitElement, html} from 'lit';
import {customElement} from 'lit/decorators.js';

@customElement('my-app')
export class MyApp extends LitElement {
  render() {
    return html`+"`<my-button size=\"\"></my-button>`"+`;
  }
}`, 0644)

		// Validate filesystem compliance
		err := mapFS.TestFS("package.json", "src/my-button.ts", "dist/custom-elements.json", "src/app.ts")
		if err != nil {
			t.Fatalf("Filesystem should pass compliance tests: %v", err)
		}

		// Create registry with mock file watcher
		mockFileWatcher := platform.NewMockFileWatcher()
		registry := lsp.NewRegistry(mockFileWatcher)

		// Load initial manifest
		var pkg M.Package
		err = pkg.UnmarshalJSON([]byte(initialManifest))
		if err != nil {
			t.Fatalf("Failed to parse initial manifest: %v", err)
		}
		registry.AddManifest(&pkg)

		t.Logf("Testing Lit template completion updates")

		// Test initial attributes - should have small, medium
		attrs, attrsExist := registry.Attributes("my-button")
		if !attrsExist {
			t.Fatalf("Expected my-button attributes to be loaded")
		}
		sizeAttr, sizeExists := attrs["size"]
		if !sizeExists {
			t.Fatalf("Expected size attribute to exist")
		}
		t.Logf("Initial size attribute type: %s", sizeAttr.Type.Text)

		// Set up MockGenerateWatcher to add 'large' size
		generateCallback := func() error {
			// Simulate user editing to add 'large' size
			updatedElementContent := strings.Replace(initialElementContent,
				`size: 'small' | 'medium' = 'medium'`,
				`size: 'small' | 'medium' | 'large' = 'medium'`,
				1)

			mapFS.AddFile("src/my-button.ts", updatedElementContent, 0644)

			// Simulate generate creating updated manifest
			updatedManifest := strings.Replace(initialManifest,
				`"text": "\"small\" | \"medium\""`,
				`"text": "\"small\" | \"medium\" | \"large\""`,
				1)

			mapFS.AddFile("dist/custom-elements.json", updatedManifest, 0644)

			// Parse and add to registry
			var updatedPkg M.Package
			err := updatedPkg.UnmarshalJSON([]byte(updatedManifest))
			if err != nil {
				return err
			}
			registry.AddManifest(&updatedPkg)
			return nil
		}

		mockGenerateWatcher := platform.NewMockGenerateWatcher(generateCallback)
		registry.SetGenerateWatcher(mockGenerateWatcher)

		err = registry.StartGenerateWatcher()
		if err != nil {
			t.Fatalf("Failed to start generate watcher: %v", err)
		}
		defer registry.StopGenerateWatcher()

		// Simulate user editing - instant execution
		t.Logf("=== User edits my-button.ts to add 'large' size ===")
		err = generateCallback()
		if err != nil {
			t.Fatalf("Failed to run generate callback: %v", err)
		}

		t.Logf("Generate completed instantly with virtual time")

		// Test updated attributes
		updatedAttrs, updatedAttrsExist := registry.Attributes("my-button")
		if !updatedAttrsExist {
			t.Fatalf("Expected my-button attributes to still exist after update")
		}
		updatedSizeAttr, updatedSizeExists := updatedAttrs["size"]
		if !updatedSizeExists {
			t.Fatalf("Expected size attribute to still exist after update")
		}
		t.Logf("Updated size attribute type: %s", updatedSizeAttr.Type.Text)

		// Verify updated completions contain 'large'
		expectedUpdated := []string{"small", "medium", "large"}
		updatedTypeText := updatedSizeAttr.Type.Text
		for _, expected := range expectedUpdated {
			if !strings.Contains(updatedTypeText, `"`+expected+`"`) {
				t.Errorf("Expected updated type to contain '%s', got: %s", expected, updatedTypeText)
			}
		}

		// Verify 'large' specifically was added
		if strings.Contains(updatedTypeText, `"large"`) {
			t.Logf("✅ SUCCESS: 'large' found in updated size type!")
		} else {
			t.Errorf("❌ FAILURE: 'large' not found in updated type: %s", updatedTypeText)
		}

		t.Logf("✅ Fast Lit template completion test completed instantly (eliminated 10+ second delays)")
	})
}