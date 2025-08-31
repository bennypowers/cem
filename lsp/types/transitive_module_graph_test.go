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
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
)

func TestModuleGraph_TransitiveElements_SingleLevel(t *testing.T) {
	// Enable debug logging for this test
	helpers.SetDebugLoggingEnabled(true)
	defer helpers.SetDebugLoggingEnabled(false)

	mg := types.NewModuleGraph()

	// my-button.js defines my-button element
	mg.AddDirectExport("my-button.js", "MyButton", "my-button")

	// my-card.js imports my-button.js and defines my-card
	mg.AddModuleDependency("my-card.js", "my-button.js")
	mg.AddDirectExport("my-card.js", "MyCard", "my-card")

	// When importing my-card.js, both my-card and my-button should be available
	// Use direct element tracking since no manifest resolver is configured
	transitiveElements := mg.GetTransitiveElementsDirect("my-card.js")
	expectedElements := []string{"my-card", "my-button"}

	if len(transitiveElements) != 2 {
		t.Errorf("Expected 2 transitive elements, got %d: %v", len(transitiveElements), transitiveElements)
	}

	for _, expected := range expectedElements {
		found := false
		for _, element := range transitiveElements {
			if element == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected element '%s' not found in transitive elements: %v", expected, transitiveElements)
		}
	}
}

func TestModuleGraph_TransitiveElements_TwoLevels(t *testing.T) {
	mg := types.NewModuleGraph()

	// my-icon.js defines my-icon element (base dependency)
	mg.AddDirectExport("my-icon.js", "MyIcon", "my-icon")

	// my-button.js imports my-icon.js and defines my-button
	mg.AddModuleDependency("my-button.js", "my-icon.js")
	mg.AddDirectExport("my-button.js", "MyButton", "my-button")

	// my-card.js imports my-button.js and defines my-card
	mg.AddModuleDependency("my-card.js", "my-button.js")
	mg.AddDirectExport("my-card.js", "MyCard", "my-card")

	// When importing my-card.js, all three elements should be available
	// Use direct element tracking since no manifest resolver is configured
	transitiveElements := mg.GetTransitiveElementsDirect("my-card.js")
	expectedElements := []string{"my-card", "my-button", "my-icon"}

	if len(transitiveElements) != 3 {
		t.Errorf("Expected 3 transitive elements, got %d: %v", len(transitiveElements), transitiveElements)
	}

	for _, expected := range expectedElements {
		found := false
		for _, element := range transitiveElements {
			if element == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected element '%s' not found in transitive elements: %v", expected, transitiveElements)
		}
	}
}

func TestModuleGraph_TransitiveElements_ThreeLevels(t *testing.T) {
	mg := types.NewModuleGraph()

	// Deep dependency chain: my-tabs -> my-tab -> my-button -> my-icon

	// my-icon.js (level 4 - deepest)
	mg.AddDirectExport("my-icon.js", "MyIcon", "my-icon")

	// my-button.js imports my-icon.js (level 3)
	mg.AddModuleDependency("my-button.js", "my-icon.js")
	mg.AddDirectExport("my-button.js", "MyButton", "my-button")

	// my-tab.js imports my-button.js (level 2)
	mg.AddModuleDependency("my-tab.js", "my-button.js")
	mg.AddDirectExport("my-tab.js", "MyTab", "my-tab")

	// my-tabs.js imports my-tab.js (level 1)
	mg.AddModuleDependency("my-tabs.js", "my-tab.js")
	mg.AddDirectExport("my-tabs.js", "MyTabs", "my-tabs")

	// When importing my-tabs.js, all four elements should be available
	// Use direct element tracking since no manifest resolver is configured
	transitiveElements := mg.GetTransitiveElementsDirect("my-tabs.js")
	expectedElements := []string{"my-tabs", "my-tab", "my-button", "my-icon"}

	if len(transitiveElements) != 4 {
		t.Errorf("Expected 4 transitive elements, got %d: %v", len(transitiveElements), transitiveElements)
	}

	for _, expected := range expectedElements {
		found := false
		for _, element := range transitiveElements {
			if element == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected element '%s' not found in transitive elements: %v", expected, transitiveElements)
		}
	}
}

func TestModuleGraph_TransitiveElements_DiamondDependency(t *testing.T) {
	mg := types.NewModuleGraph()

	// Diamond dependency pattern:
	//     my-form.js
	//     /        \
	// my-input.js  my-button.js
	//     \        /
	//     my-icon.js

	// my-icon.js (shared dependency)
	mg.AddDirectExport("my-icon.js", "MyIcon", "my-icon")

	// my-input.js and my-button.js both import my-icon.js
	mg.AddModuleDependency("my-input.js", "my-icon.js")
	mg.AddDirectExport("my-input.js", "MyInput", "my-input")

	mg.AddModuleDependency("my-button.js", "my-icon.js")
	mg.AddDirectExport("my-button.js", "MyButton", "my-button")

	// my-form.js imports both my-input.js and my-button.js
	mg.AddModuleDependency("my-form.js", "my-input.js")
	mg.AddModuleDependency("my-form.js", "my-button.js")
	mg.AddDirectExport("my-form.js", "MyForm", "my-form")

	// When importing my-form.js, all elements should be available (no duplicates)
	// Use direct element tracking since no manifest resolver is configured
	transitiveElements := mg.GetTransitiveElementsDirect("my-form.js")
	expectedElements := []string{"my-form", "my-input", "my-button", "my-icon"}

	if len(transitiveElements) != 4 {
		t.Errorf("Expected 4 unique transitive elements, got %d: %v", len(transitiveElements), transitiveElements)
	}

	// Check for duplicates
	seen := make(map[string]bool)
	for _, element := range transitiveElements {
		if seen[element] {
			t.Errorf("Found duplicate element '%s' in transitive elements", element)
		}
		seen[element] = true
	}

	for _, expected := range expectedElements {
		found := false
		for _, element := range transitiveElements {
			if element == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected element '%s' not found in transitive elements: %v", expected, transitiveElements)
		}
	}
}

func TestModuleGraph_TransitiveElements_CircularDependency(t *testing.T) {
	mg := types.NewModuleGraph()

	// Circular dependency: my-a.js -> my-b.js -> my-c.js -> my-a.js
	mg.AddDirectExport("my-a.js", "MyA", "my-a")
	mg.AddDirectExport("my-b.js", "MyB", "my-b")
	mg.AddDirectExport("my-c.js", "MyC", "my-c")

	mg.AddModuleDependency("my-a.js", "my-b.js")
	mg.AddModuleDependency("my-b.js", "my-c.js")
	mg.AddModuleDependency("my-c.js", "my-a.js") // Creates cycle

	// Should handle circular dependencies without infinite loop
	// Use direct element tracking since no manifest resolver is configured
	transitiveElements := mg.GetTransitiveElementsDirect("my-a.js")
	expectedElements := []string{"my-a", "my-b", "my-c"}

	if len(transitiveElements) != 3 {
		t.Errorf("Expected 3 transitive elements, got %d: %v", len(transitiveElements), transitiveElements)
	}

	for _, expected := range expectedElements {
		found := false
		for _, element := range transitiveElements {
			if element == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected element '%s' not found in transitive elements: %v", expected, transitiveElements)
		}
	}
}

func TestModuleGraph_TransitiveElements_DepthLimit(t *testing.T) {
	mg := types.NewModuleGraph()

	// Create a deep chain: my-1 -> my-2 -> my-3 -> ... -> my-10
	for i := 1; i <= 10; i++ {
		modulePath := fmt.Sprintf("my-%d.js", i)
		elementName := fmt.Sprintf("My%d", i)
		tagName := fmt.Sprintf("my-%d", i)

		mg.AddDirectExport(modulePath, elementName, tagName)

		if i > 1 {
			prevModulePath := fmt.Sprintf("my-%d.js", i-1)
			mg.AddModuleDependency(prevModulePath, modulePath)
		}
	}

	// Should limit depth to prevent performance issues
	// Use direct element tracking since no manifest resolver is configured
	transitiveElements := mg.GetTransitiveElementsDirect("my-1.js")

	// Should include at least 5 levels (configurable depth limit)
	if len(transitiveElements) < 5 {
		t.Errorf("Expected at least 5 transitive elements for depth limit test, got %d", len(transitiveElements))
	}

	// Should not include all 10 due to depth limit
	if len(transitiveElements) >= 10 {
		t.Errorf("Expected depth limit to prevent all 10 elements, got %d", len(transitiveElements))
	}
}

func TestModuleGraph_RealWorldWorkspaceScenario(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "transitive-module-graph-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test files that represent the transitive dependency scenario
	if err := createTransitiveTestFiles(tempDir); err != nil {
		t.Fatalf("Failed to create test files: %v", err)
	}

	// Test module graph building with real files
	mg := types.NewModuleGraph()
	err = mg.BuildFromWorkspace(tempDir)
	if err != nil {
		t.Fatalf("Failed to build module graph: %v", err)
	}

	// Verify transitive dependencies were detected
	// Use direct element tracking since no manifest resolver is configured
	transitiveElements := mg.GetTransitiveElementsDirect("my-tabs.ts")

	// Should include my-tabs, my-tab, and my-icon
	expectedElements := []string{"my-tabs", "my-tab", "my-icon"}

	for _, expected := range expectedElements {
		found := false
		for _, element := range transitiveElements {
			if element == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected element '%s' not found in transitive elements from real workspace: %v", expected, transitiveElements)
		}
	}
}

// createTransitiveTestFiles creates realistic test files demonstrating transitive dependencies
func createTransitiveTestFiles(tempDir string) error {
	// my-icon.ts (base dependency)
	iconContent := `
export class MyIcon extends HTMLElement {
  // Icon implementation
}

customElements.define('my-icon', MyIcon);
`

	// my-tab.ts (imports my-icon)
	tabContent := `
import './my-icon.ts';

export class MyTab extends HTMLElement {
  // Tab implementation using MyIcon
}

customElements.define('my-tab', MyTab);
`

	// my-tabs.ts (imports my-tab, which transitively imports my-icon)
	tabsContent := `
import './my-tab.ts';

export class MyTabs extends HTMLElement {
  // Tabs container implementation
}

customElements.define('my-tabs', MyTabs);
`

	// Write test files
	files := map[string]string{
		"my-icon.ts": iconContent,
		"my-tab.ts":  tabContent,
		"my-tabs.ts": tabsContent,
	}

	for filename, content := range files {
		if err := os.WriteFile(filepath.Join(tempDir, filename), []byte(content), 0644); err != nil {
			return err
		}
	}

	return nil
}

// TestModuleGraph_RecursiveDependencyBuilding tests the new recursive dependency building functionality
// that ensures complete transitive dependency trees are built for manifest modules
func TestModuleGraph_RecursiveDependencyBuilding(t *testing.T) {
	// Create temporary directory for test files
	tempDir, err := os.MkdirTemp("", "module-graph-recursive-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test files with transitive dependency chain
	if err := createTransitiveTestFiles(tempDir); err != nil {
		t.Fatalf("Failed to create test files: %v", err)
	}

	// Create module graph with workspace root
	mg := types.NewModuleGraphWithDependencies(
		&types.OSFileParser{},
		&types.DefaultExportParser{},
		&types.NoOpManifestResolver{}, // Use no-op since we're testing file-based dependency building
		&types.NoOpMetricsCollector{},
	)
	mg.SetWorkspaceRoot(tempDir)

	// Test that buildDependenciesForManifestModule builds the complete transitive tree
	// This should recursively find: my-tabs.ts -> my-tab.ts -> my-icon.ts
	err = mg.BuildForImportPath("./my-tabs.ts")
	if err != nil {
		t.Fatalf("Failed to build dependencies for my-tabs.ts: %v", err)
	}

	// Verify that all transitive dependencies were built
	tabsDeps := mg.GetModuleDependencies("my-tabs.ts")
	tabDeps := mg.GetModuleDependencies("my-tab.ts")
	iconDeps := mg.GetModuleDependencies("my-icon.ts")

	// Check the complete dependency chain
	if len(tabsDeps) != 1 || tabsDeps[0] != "my-tab.ts" {
		t.Errorf("Expected my-tabs.ts to depend on [my-tab.ts], got %v", tabsDeps)
	}

	if len(tabDeps) != 1 || tabDeps[0] != "my-icon.ts" {
		t.Errorf("Expected my-tab.ts to depend on [my-icon.ts], got %v", tabDeps)
	}

	if len(iconDeps) != 0 {
		t.Errorf("Expected my-icon.ts to have no dependencies, got %v", iconDeps)
	}

	t.Logf("✅ Recursive dependency building successfully created complete transitive chain:")
	t.Logf("  my-tabs.ts -> %v", tabsDeps)
	t.Logf("  my-tab.ts -> %v", tabDeps)
	t.Logf("  my-icon.ts -> %v", iconDeps)
}
