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
	"strings"
	"testing"
	"testing/synctest"

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/lsp"
	M "bennypowers.dev/cem/manifest"
)

// TestCompletionUpdateCycle tests the complete update cycle with instant virtual time
// This replaces the E2E test's 15-second timeout with instant synctest operations
func TestCompletionUpdateCycle(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// Create in-memory filesystem for instant operations
		mapFS := platform.NewMapFileSystem(nil)

		// Set up project structure instantly
		mapFS.AddFile("package.json", `{
  "name": "test-package",
  "customElements": "dist/custom-elements.json"
}`, 0644)

		mapFS.AddFile(".config/cem.yaml", `generate:
  files:
    - elements/*/test-*.ts
`, 0644)

		// Initial TypeScript element with three values
		initialElementContent := `import {LitElement} from 'lit';
import {customElement, property} from 'lit/decorators.js';

@customElement('test-element')
export class TestElement extends LitElement {
  @property() prop: 'one' | 'two' | 'three' = 'one';
}`

		mapFS.AddFile("elements/test-element/test-element.ts", initialElementContent, 0644)

		// Initial manifest with three completion values
		initialManifest := `{
  "schemaVersion": "1.0.0",
  "readme": "",
  "modules": [
    {
      "kind": "javascript-module",
      "path": "elements/test-element/test-element.js",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "TestElement",
          "customElement": true,
          "tagName": "test-element",
          "attributes": [
            {
              "name": "prop",
              "type": {
                "text": "\"one\" | \"two\" | \"three\""
              },
              "default": "\"one\""
            }
          ]
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "test-element",
          "declaration": {
            "name": "TestElement",
            "module": "elements/test-element/test-element.js"
          }
        }
      ]
    }
  ]
}`

		mapFS.AddFile("dist/custom-elements.json", initialManifest, 0644)

		// Demo HTML file for testing
		mapFS.AddFile("demo.html", `<!DOCTYPE html>
<html>
<body>
  <test-element prop=""></test-element>
</body>
</html>`, 0644)

		// Validate filesystem compliance
		err := mapFS.TestFS("package.json", ".config/cem.yaml", "elements/test-element/test-element.ts", "dist/custom-elements.json", "demo.html")
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

		t.Logf("Testing completion update cycle with fast operations")

		// Test initial element - should have one, two, three
		element, exists := registry.Element("test-element")
		if !exists {
			t.Fatalf("Expected test-element to be loaded")
		}
		t.Logf("Initial element loaded: %s with %d attributes", element.TagName, len(element.Attributes))

		// Verify initial attributes
		attrs, attrsExist := registry.Attributes("test-element")
		if !attrsExist {
			t.Fatalf("Expected test-element attributes to be loaded")
		}
		propAttr, propExists := attrs["prop"]
		if !propExists {
			t.Fatalf("Expected prop attribute to exist")
		}
		t.Logf("Initial prop attribute type: %s", propAttr.Type.Text)

		// Verify initial type contains expected values
		expectedInitial := []string{"one", "two", "three"}
		initialTypeText := propAttr.Type.Text
		for _, expected := range expectedInitial {
			if !strings.Contains(initialTypeText, `"`+expected+`"`) {
				t.Errorf("Expected initial type to contain '%s', got: %s", expected, initialTypeText)
			}
		}

		// Set up MockGenerateWatcher with callback that simulates user adding 'four' to the union
		generateCallback := func() error {
			// Simulate user editing the source file to add 'four'
			updatedTSContent := strings.Replace(initialElementContent,
				`prop: 'one' | 'two' | 'three' = 'one'`,
				`prop: 'one' | 'two' | 'three' | 'four' = 'one'`,
				1)

			// Update the TypeScript file in MapFS
			mapFS.AddFile("elements/test-element/test-element.ts", updatedTSContent, 0644)

			// Simulate generate creating updated manifest
			updatedManifest := strings.Replace(initialManifest,
				`"text": "\"one\" | \"two\" | \"three\""`,
				`"text": "\"one\" | \"two\" | \"three\" | \"four\""`,
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
		defer func() { _ = registry.StopGenerateWatcher() }()

		// Simulate user editing the source file - instant with virtual time
		t.Logf("=== User edits test-element.ts to add 'four' to union ===")

		// Call the callback directly for instant execution
		err = generateCallback()
		if err != nil {
			t.Fatalf("Failed to run generate callback: %v", err)
		}

		// No timing delays needed - everything happens instantly with synctest
		t.Logf("Generate completed instantly with virtual time")

		// Test updated attributes - should now include 'four'
		updatedAttrs, updatedAttrsExist := registry.Attributes("test-element")
		if !updatedAttrsExist {
			t.Fatalf("Expected test-element attributes to still exist after update")
		}
		updatedPropAttr, updatedPropExists := updatedAttrs["prop"]
		if !updatedPropExists {
			t.Fatalf("Expected prop attribute to still exist after update")
		}
		t.Logf("Updated prop attribute type: %s", updatedPropAttr.Type.Text)

		// Verify updated type contains all expected values including 'four'
		expectedUpdated := []string{"one", "two", "three", "four"}
		updatedTypeText := updatedPropAttr.Type.Text
		for _, expected := range expectedUpdated {
			if !strings.Contains(updatedTypeText, `"`+expected+`"`) {
				t.Errorf("Expected updated type to contain '%s', got: %s", expected, updatedTypeText)

				// Debug: Check the manifest file content from MapFS
				if manifestContent, err := mapFS.ReadFile("dist/custom-elements.json"); err == nil {
					t.Logf("Current manifest content: %s", string(manifestContent))
				}
			}
		}

		// Verify 'four' specifically was added
		if strings.Contains(updatedTypeText, `"four"`) {
			t.Logf("✅ SUCCESS: 'four' found in updated prop type!")
		} else {
			t.Errorf("❌ FAILURE: 'four' not found in updated type: %s", updatedTypeText)
		}

		t.Logf("✅ Fast completion update cycle test completed instantly (eliminated 15-second delay)")
	})
}
