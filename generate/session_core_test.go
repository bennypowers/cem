/*
Copyright © 2025 Benny Powers

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
package generate

import (
	"testing"

	M "bennypowers.dev/cem/manifest"
)

func TestDeepCopyManifest(t *testing.T) {
	// Create a test GenerateSession (just for the deepCopyManifest method)
	gs := &GenerateSession{}

	// Test nil case
	if result := gs.deepCopyManifest(nil); result != nil {
		t.Error("deepCopyManifest should return nil for nil input")
	}

	// Create a test manifest with nested data
	original := &M.Package{
		SchemaVersion: "2.1.0",
		Modules: []M.Module{
			{
				Kind: "javascript-module",
				Path: "test/module.js",
				Summary: "Test module",
				Description: "A test module for deep copy verification",
			},
		},
	}

	// Test deep copy
	copy := gs.deepCopyManifest(original)
	if copy == nil {
		t.Fatal("deepCopyManifest should not return nil for valid input")
	}

	// Verify the copy has the same content
	if copy.SchemaVersion != original.SchemaVersion {
		t.Errorf("SchemaVersion mismatch: got %s, want %s", copy.SchemaVersion, original.SchemaVersion)
	}

	if len(copy.Modules) != len(original.Modules) {
		t.Errorf("Modules length mismatch: got %d, want %d", len(copy.Modules), len(original.Modules))
	}

	if len(copy.Modules) > 0 && len(original.Modules) > 0 {
		if copy.Modules[0].Path != original.Modules[0].Path {
			t.Errorf("Module path mismatch: got %s, want %s", copy.Modules[0].Path, original.Modules[0].Path)
		}
	}

	// Verify it's actually a deep copy (different memory addresses)
	if copy == original {
		t.Error("deepCopyManifest should return a different instance")
	}

	if len(copy.Modules) > 0 && len(original.Modules) > 0 {
		// Note: We can't easily test slice element address differences due to how Go handles slice copying,
		// but the JSON round-trip ensures deep copying of all nested structures
	}

	// Test that modifying the copy doesn't affect the original
	if len(copy.Modules) > 0 {
		copy.Modules[0].Summary = "Modified summary"
		if original.Modules[0].Summary == "Modified summary" {
			t.Error("Modifying copy should not affect original (not properly deep copied)")
		}
	}
}

func TestGetInMemoryManifest_Performance(t *testing.T) {
	gs := &GenerateSession{}

	// Test nil manifest
	if result := gs.GetInMemoryManifest(); result != nil {
		t.Error("GetInMemoryManifest should return nil when no manifest is set")
	}

	// Set a test manifest
	original := &M.Package{
		SchemaVersion: "2.1.0",
		Modules: []M.Module{
			{
				Kind: "javascript-module",
				Path: "test/module.js",
				Summary: "Test module",
			},
		},
	}

	gs.mu.Lock()
	gs.inMemoryManifest = original
	gs.mu.Unlock()

	// Get a copy (shallow for performance)
	copy := gs.GetInMemoryManifest()
	if copy == nil {
		t.Fatal("GetInMemoryManifest should return a copy of the manifest")
	}

	// Verify it's a different instance (thread-safe at the package level)
	if copy == original {
		t.Error("GetInMemoryManifest should return a different instance")
	}

	// Verify content is preserved
	if copy.SchemaVersion != original.SchemaVersion {
		t.Errorf("Content mismatch after copy: got %s, want %s", copy.SchemaVersion, original.SchemaVersion)
	}
}

func TestGetInMemoryManifestDeep_ThreadSafety(t *testing.T) {
	gs := &GenerateSession{}

	// Test nil manifest
	if result := gs.GetInMemoryManifestDeep(); result != nil {
		t.Error("GetInMemoryManifestDeep should return nil when no manifest is set")
	}

	// Set a test manifest
	original := &M.Package{
		SchemaVersion: "2.1.0",
		Modules: []M.Module{
			{
				Kind: "javascript-module",
				Path: "test/module.js",
				Summary: "Test module",
			},
		},
	}

	gs.mu.Lock()
	gs.inMemoryManifest = original
	gs.mu.Unlock()

	// Get a deep copy
	copy := gs.GetInMemoryManifestDeep()
	if copy == nil {
		t.Fatal("GetInMemoryManifestDeep should return a copy of the manifest")
	}

	// Verify it's a different instance (thread-safe)
	if copy == original {
		t.Error("GetInMemoryManifestDeep should return a different instance for thread safety")
	}

	// Verify content is preserved
	if copy.SchemaVersion != original.SchemaVersion {
		t.Errorf("Content mismatch after deep copy: got %s, want %s", copy.SchemaVersion, original.SchemaVersion)
	}

	// Test that modifying the deep copy doesn't affect the original
	if len(copy.Modules) > 0 {
		copy.Modules[0].Summary = "Modified summary"
		if original.Modules[0].Summary == "Modified summary" {
			t.Error("Modifying deep copy should not affect original (not properly deep copied)")
		}
	}
}

func TestModuleIndex_Performance(t *testing.T) {
	gs := &GenerateSession{}
	
	// Create test modules
	modules := []M.Module{
		{Kind: "javascript-module", Path: "test/module1.js"},
		{Kind: "javascript-module", Path: "test/module2.js"},
		{Kind: "javascript-module", Path: "test/module3.js"},
	}
	
	pkg := M.NewPackage(modules)
	
	// Set manifest and build index
	gs.mu.Lock()
	gs.inMemoryManifest = &pkg
	gs.moduleIndex = make(map[string]*M.Module)
	gs.rebuildModuleIndex()
	gs.mu.Unlock()
	
	// Test O(1) lookup
	module := gs.GetModuleByPath("test/module2.js")
	if module == nil {
		t.Fatal("GetModuleByPath should find existing module")
	}
	
	if module.Path != "test/module2.js" {
		t.Errorf("Wrong module returned: got %s, want test/module2.js", module.Path)
	}
	
	// Test non-existent module
	module = gs.GetModuleByPath("test/nonexistent.js")
	if module != nil {
		t.Error("GetModuleByPath should return nil for non-existent module")
	}
}

func TestMergeModulesIntoManifest_WithIndex(t *testing.T) {
	gs := &GenerateSession{
		moduleIndex: make(map[string]*M.Module),
	}
	
	// Test with empty manifest
	updatedModules := []M.Module{
		{Kind: "javascript-module", Path: "test/new.js", Summary: "New module"},
	}
	
	gs.MergeModulesIntoManifest(updatedModules)
	
	if gs.inMemoryManifest == nil {
		t.Fatal("Manifest should be created")
	}
	
	if len(gs.inMemoryManifest.Modules) != 1 {
		t.Errorf("Expected 1 module, got %d", len(gs.inMemoryManifest.Modules))
	}
	
	// Verify index was built
	module := gs.GetModuleByPath("test/new.js")
	if module == nil {
		t.Error("Module should be findable in index")
	}
	
	// Test updating existing module
	existingModules := []M.Module{
		{Kind: "javascript-module", Path: "test/existing.js", Summary: "Original"},
	}
	
	pkg := M.NewPackage(existingModules)
	gs.mu.Lock()
	gs.inMemoryManifest = &pkg
	gs.rebuildModuleIndex()
	gs.mu.Unlock()
	
	// Update the module
	updatedModules = []M.Module{
		{Kind: "javascript-module", Path: "test/existing.js", Summary: "Updated"},
	}
	
	gs.MergeModulesIntoManifest(updatedModules)
	
	// Verify update worked
	module = gs.GetModuleByPath("test/existing.js")
	if module == nil {
		t.Fatal("Updated module should be findable")
	}
	
	if module.Summary != "Updated" {
		t.Errorf("Module summary should be updated: got %s, want Updated", module.Summary)
	}
	
	// Should still only have one module
	if len(gs.inMemoryManifest.Modules) != 1 {
		t.Errorf("Should still have 1 module after update, got %d", len(gs.inMemoryManifest.Modules))
	}
}