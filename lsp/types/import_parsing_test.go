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

func TestModuleGraph_ImportParsing_Simple(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "import-parsing-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create simple test files
	iconContent := `
export class MyIcon extends HTMLElement {}
customElements.define('my-icon', MyIcon);
`

	buttonContent := `
import './my-icon.js';
export class MyButton extends HTMLElement {}
customElements.define('my-button', MyButton);
`

	// Write files
	if err := os.WriteFile(filepath.Join(tempDir, "my-icon.js"), []byte(iconContent), 0644); err != nil {
		t.Fatalf("Failed to write my-icon.js: %v", err)
	}
	if err := os.WriteFile(filepath.Join(tempDir, "my-button.js"), []byte(buttonContent), 0644); err != nil {
		t.Fatalf("Failed to write my-button.js: %v", err)
	}

	// Build module graph
	mg := types.NewModuleGraph()
	err = mg.BuildFromWorkspace(tempDir)
	if err != nil {
		t.Fatalf("Failed to build module graph: %v", err)
	}

	// Check if my-icon element was detected
	// Use direct element tracking since no manifest resolver is configured
	iconElements := mg.GetTransitiveElementsDirect("my-icon.js")
	if len(iconElements) != 1 || iconElements[0] != "my-icon" {
		t.Errorf("Expected ['my-icon'] for my-icon.js, got %v", iconElements)
	}

	// Check if my-button has transitive dependency on my-icon
	// Use direct element tracking since no manifest resolver is configured
	buttonElements := mg.GetTransitiveElementsDirect("my-button.js")
	expectedElements := []string{"my-button", "my-icon"}

	if len(buttonElements) != 2 {
		t.Errorf("Expected 2 elements for my-button.js, got %d: %v", len(buttonElements), buttonElements)
	}

	for _, expected := range expectedElements {
		found := false
		for _, element := range buttonElements {
			if element == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected element '%s' not found in my-button.js transitive elements: %v", expected, buttonElements)
		}
	}
}
