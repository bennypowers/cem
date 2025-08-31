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
package types_test

import (
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/lsp/types"
)

func TestModuleGraph_BasicOperations(t *testing.T) {
	mg := types.NewModuleGraph()

	// Test adding direct exports
	mg.AddDirectExport("components/rh-tab.js", "RhTab", "rh-tab")
	mg.AddDirectExport("components/rh-button.js", "RhButton", "rh-button")

	// Test getting element sources
	sources := mg.GetElementSources("rh-tab")
	if len(sources) != 1 || sources[0] != "components/rh-tab.js" {
		t.Errorf("Expected ['components/rh-tab.js'], got %v", sources)
	}

	// Test getting all tag names
	allTagNames := mg.GetAllTagNames()
	expectedTags := []string{"rh-tab", "rh-button"}
	if len(allTagNames) != 2 {
		t.Errorf("Expected 2 tag names, got %d", len(allTagNames))
	}
	for _, expectedTag := range expectedTags {
		found := false
		for _, tag := range allTagNames {
			if tag == expectedTag {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected tag '%s' not found in %v", expectedTag, allTagNames)
		}
	}
}

func TestModuleGraph_ReExports(t *testing.T) {
	mg := types.NewModuleGraph()

	// Add a direct export
	mg.AddDirectExport("components/rh-tab.js", "RhTab", "rh-tab")

	// Add a re-export (rh-tabs.js re-exports rh-tab)
	mg.AddReExport("components/rh-tabs.js", "components/rh-tab.js", "RhTab", "rh-tab")

	// Test that rh-tab is now available from both modules
	sources := mg.GetElementSources("rh-tab")
	expectedSources := []string{"components/rh-tab.js", "components/rh-tabs.js"}

	if len(sources) != 2 {
		t.Errorf("Expected 2 sources for rh-tab, got %d: %v", len(sources), sources)
	}

	for _, expectedSource := range expectedSources {
		found := false
		for _, source := range sources {
			if source == expectedSource {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected source '%s' not found in %v", expectedSource, sources)
		}
	}
}

func TestModuleGraph_RealWorldScenario(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "module-graph-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test TypeScript files simulating the rh-tabs scenario
	rhTabFile := filepath.Join(tempDir, "rh-tab.ts")
	rhTabContent := `
export class RhTab extends HTMLElement {
  // Element implementation
}

customElements.define('rh-tab', RhTab);
`

	rhTabsFile := filepath.Join(tempDir, "rh-tabs.ts")
	rhTabsContent := `
export class RhTabs extends HTMLElement {
  // Tabs container implementation
}

// Re-export RhTab for convenience
export { RhTab } from './rh-tab.js';

customElements.define('rh-tabs', RhTabs);
`

	// Write test files
	if err := os.WriteFile(rhTabFile, []byte(rhTabContent), 0644); err != nil {
		t.Fatalf("Failed to write rh-tab.ts: %v", err)
	}
	if err := os.WriteFile(rhTabsFile, []byte(rhTabsContent), 0644); err != nil {
		t.Fatalf("Failed to write rh-tabs.ts: %v", err)
	}

	// Test module graph building
	mg := types.NewModuleGraph()
	err = mg.BuildFromWorkspace(tempDir)
	if err != nil {
		t.Fatalf("Failed to build module graph: %v", err)
	}

	// Verify that rh-tab was detected from rh-tab.ts
	tabSources := mg.GetElementSources("rh-tab")
	foundDirectExport := false
	for _, source := range tabSources {
		if source == "rh-tab.ts" {
			foundDirectExport = true
			break
		}
	}
	if !foundDirectExport {
		t.Errorf("Expected to find rh-tab direct export from rh-tab.ts, got sources: %v", tabSources)
	}

	// Verify that rh-tabs was detected from rh-tabs.ts
	tabsSources := mg.GetElementSources("rh-tabs")
	foundTabsExport := false
	for _, source := range tabsSources {
		if source == "rh-tabs.ts" {
			foundTabsExport = true
			break
		}
	}
	if !foundTabsExport {
		t.Errorf("Expected to find rh-tabs direct export from rh-tabs.ts, got sources: %v", tabsSources)
	}

	// Test that both elements are tracked
	allTagNames := mg.GetAllTagNames()
	hasRhTab := false
	hasRhTabs := false
	for _, tagName := range allTagNames {
		if tagName == "rh-tab" {
			hasRhTab = true
		}
		if tagName == "rh-tabs" {
			hasRhTabs = true
		}
	}

	if !hasRhTab {
		t.Error("Expected to find 'rh-tab' in all tag names")
	}
	if !hasRhTabs {
		t.Error("Expected to find 'rh-tabs' in all tag names")
	}
}

func TestModuleGraph_GetModuleExports(t *testing.T) {
	mg := types.NewModuleGraph()

	// Add exports to a module
	mg.AddDirectExport("components/multi-element.js", "ElementOne", "element-one")
	mg.AddDirectExport("components/multi-element.js", "ElementTwo", "element-two")

	// Get module exports
	exports := mg.GetModuleExports("components/multi-element.js")

	if len(exports) != 2 {
		t.Errorf("Expected 2 exports, got %d", len(exports))
	}

	// Verify export details
	foundElementOne := false
	foundElementTwo := false
	for _, export := range exports {
		if export.ElementName == "ElementOne" && export.TagName == "element-one" {
			foundElementOne = true
		}
		if export.ElementName == "ElementTwo" && export.TagName == "element-two" {
			foundElementTwo = true
		}
	}

	if !foundElementOne {
		t.Error("Expected to find ElementOne export")
	}
	if !foundElementTwo {
		t.Error("Expected to find ElementTwo export")
	}
}

func TestModuleGraph_EmptyGraph(t *testing.T) {
	mg := types.NewModuleGraph()

	// Test empty graph behavior
	allTagNames := mg.GetAllTagNames()
	if len(allTagNames) != 0 {
		t.Errorf("Expected empty tag names list, got %v", allTagNames)
	}

	sources := mg.GetElementSources("non-existent")
	if sources != nil {
		t.Errorf("Expected nil for non-existent element, got %v", sources)
	}

	exports := mg.GetModuleExports("non-existent.js")
	if exports != nil {
		t.Errorf("Expected nil for non-existent module, got %v", exports)
	}
}
