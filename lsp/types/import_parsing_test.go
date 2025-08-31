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
	"testing"

	"bennypowers.dev/cem/lsp/types"
)

func TestModuleGraph_ImportParsing_Simple(t *testing.T) {
	// Create module graph and populate from manifest data (like production LSP server)
	mg := types.NewModuleGraph(nil) // No QueryManager needed for direct element tracking tests

	// Simulate manifest data for elements
	elementMap := map[string]interface{}{
		"my-icon": &types.MockElementDefinition{
			TagName:    "my-icon",
			ClassName:  "MyIcon",
			ModulePath: "my-icon.js",
		},
		"my-button": &types.MockElementDefinition{
			TagName:    "my-button",
			ClassName:  "MyButton",
			ModulePath: "my-button.js",
		},
	}

	// Populate module graph from manifest data (production approach)
	mg.PopulateFromManifests(elementMap)

	// Add dependency relationship (my-button.js imports my-icon.js)
	mg.AddModuleDependency("my-button.js", "my-icon.js")

	// Check if my-icon element was detected
	iconElements := mg.GetTransitiveElementsDirect("my-icon.js")
	if len(iconElements) != 1 || iconElements[0] != "my-icon" {
		t.Errorf("Expected ['my-icon'] for my-icon.js, got %v", iconElements)
	}

	// Check if my-button has transitive dependency on my-icon
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
