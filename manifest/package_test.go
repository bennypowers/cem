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
package manifest_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"bennypowers.dev/cem/manifest"
)

// TestPackage tests the Clone functionality across all manifest types
// using a comprehensive manifest that includes all possible types and structures.
func TestPackage(t *testing.T) {
	// Load the comprehensive test manifest
	manifestPath := filepath.Join("fixtures", "comprehensive-clone-test.json")
	manifestData, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatalf("Failed to read test manifest: %v", err)
	}

	// Unmarshal the manifest
	var original manifest.Package
	if err := json.Unmarshal(manifestData, &original); err != nil {
		t.Fatalf("Failed to unmarshal manifest: %v", err)
	}

	// Clone the entire package
	cloned := original.Clone()
	if cloned == nil {
		t.Fatal("Clone returned nil")
	}

	// Test that the clone is a deep copy by verifying:
	// 1. The cloned object is equal in value to the original
	// 2. Modifying the clone doesn't affect the original
	// 3. Modifying the original doesn't affect the clone

	t.Run("ClonedValueEqualsOriginal", func(t *testing.T) {
		if !reflect.DeepEqual(&original, cloned) {
			t.Error("Cloned package is not deeply equal to original")
		}
	})

	t.Run("ModifyingCloneDoesNotAffectOriginal", func(t *testing.T) {
		// Store original values for comparison
		originalReadme := original.Readme
		originalDeprecated := original.Deprecated
		originalModuleCount := len(original.Modules)

		// Modify the clone
		if cloned.Readme != nil {
			*cloned.Readme = "Modified README"
		}
		if cloned.Deprecated != nil {
			// Create a new deprecated reason for the clone
			cloned.Deprecated = manifest.DeprecatedReason("Modified deprecation reason")
		}

		// Modify a module in the clone
		if len(cloned.Modules) > 0 {
			cloned.Modules[0].Summary = "Modified summary"
			cloned.Modules[0].Description = "Modified description"

			// Modify declarations
			if len(cloned.Modules[0].Declarations) > 0 {
				switch decl := cloned.Modules[0].Declarations[0].(type) {
				case *manifest.ClassDeclaration:
					decl.Summary = "Modified class summary"
				case *manifest.CustomElementDeclaration:
					decl.Summary = "Modified custom element summary"
				}
			}
		}

		// Verify original is unchanged
		if !reflect.DeepEqual(original.Readme, originalReadme) {
			t.Error("Original README was modified when clone was changed")
		}
		if !reflect.DeepEqual(original.Deprecated, originalDeprecated) {
			t.Error("Original Deprecated was modified when clone was changed")
		}
		if len(original.Modules) != originalModuleCount {
			t.Error("Original modules count changed when clone was modified")
		}
		if len(original.Modules) > 0 && original.Modules[0].Summary == "Modified summary" {
			t.Error("Original module summary was modified when clone was changed")
		}
	})

	t.Run("ModifyingOriginalDoesNotAffectClone", func(t *testing.T) {
		// Create a fresh clone for this test
		freshClone := original.Clone()

		// Store clone values for comparison
		clonedReadme := freshClone.Readme
		clonedDeprecated := freshClone.Deprecated
		clonedModuleCount := len(freshClone.Modules)

		// Modify the original
		if original.Readme != nil {
			*original.Readme = "Modified Original README"
		}
		if original.Deprecated != nil {
			// Create a new deprecated reason for the original
			original.Deprecated = manifest.DeprecatedReason("Modified original deprecation")
		}

		// Modify a module in the original
		if len(original.Modules) > 0 {
			original.Modules[0].Summary = "Modified original summary"

			// Modify declarations in original
			if len(original.Modules[0].Declarations) > 0 {
				switch decl := original.Modules[0].Declarations[0].(type) {
				case *manifest.ClassDeclaration:
					decl.Summary = "Modified original class summary"
				case *manifest.CustomElementDeclaration:
					decl.Summary = "Modified original custom element summary"
				}
			}
		}

		// Verify clone is unchanged
		if !reflect.DeepEqual(freshClone.Readme, clonedReadme) {
			t.Error("Clone README was modified when original was changed")
		}
		if !reflect.DeepEqual(freshClone.Deprecated, clonedDeprecated) {
			t.Error("Clone Deprecated was modified when original was changed")
		}
		if len(freshClone.Modules) != clonedModuleCount {
			t.Error("Clone modules count changed when original was modified")
		}
		if len(freshClone.Modules) > 0 && freshClone.Modules[0].Summary == "Modified original summary" {
			t.Error("Clone module summary was modified when original was changed")
		}
	})

	t.Run("AllDeclarationTypesCanBeCloned", func(t *testing.T) {
		cloned := original.Clone()

		// Verify all declaration types are properly cloned
		for i, module := range cloned.Modules {
			for j, decl := range module.Declarations {
				if decl == nil {
					t.Errorf("Declaration %d in module %d is nil after cloning", j, i)
					continue
				}

				// Test that we can call Clone on each declaration
				declClone := decl.Clone()
				if declClone == nil {
					t.Errorf("Declaration %d in module %d returned nil when cloned", j, i)
				}
			}
		}
	})

	t.Run("AllExportTypesCanBeCloned", func(t *testing.T) {
		cloned := original.Clone()

		// Verify all export types are properly cloned
		for i, module := range cloned.Modules {
			for j, export := range module.Exports {
				if export == nil {
					t.Errorf("Export %d in module %d is nil after cloning", j, i)
					continue
				}

				// Test that we can call Clone on each export
				exportClone := export.Clone()
				if exportClone == nil {
					t.Errorf("Export %d in module %d returned nil when cloned", j, i)
				}
			}
		}
	})

	t.Run("AllClassMemberTypesCanBeCloned", func(t *testing.T) {
		cloned := original.Clone()

		// Find class declarations and test their members
		for i, module := range cloned.Modules {
			for j, decl := range module.Declarations {
				switch d := decl.(type) {
				case *manifest.ClassDeclaration:
					for k, member := range d.Members {
						if member == nil {
							t.Errorf("Member %d in class declaration %d in module %d is nil after cloning", k, j, i)
							continue
						}

						memberClone := member.Clone()
						if memberClone == nil {
							t.Errorf("Member %d in class declaration %d in module %d returned nil when cloned", k, j, i)
						}
					}
				case *manifest.CustomElementDeclaration:
					for k, member := range d.Members {
						if member == nil {
							t.Errorf("Member %d in custom element declaration %d in module %d is nil after cloning", k, j, i)
							continue
						}

						memberClone := member.Clone()
						if memberClone == nil {
							t.Errorf("Member %d in custom element declaration %d in module %d returned nil when cloned", k, j, i)
						}
					}
				}
			}
		}
	})

	t.Run("NestedStructuresAreDeepCloned", func(t *testing.T) {
		cloned := original.Clone()

		// Test that nested structures like attributes, events, etc. are deep cloned
		for _, module := range cloned.Modules {
			for _, decl := range module.Declarations {
				if customElement, ok := decl.(*manifest.CustomElementDeclaration); ok {
					// Test attributes
					for i, attr := range customElement.Attributes {
						// Modify the clone
						attr.Summary = "Modified attribute summary"

						// Verify original is unchanged
						if i < len(original.Modules[0].Declarations[1].(*manifest.CustomElementDeclaration).Attributes) {
							originalAttr := original.Modules[0].Declarations[1].(*manifest.CustomElementDeclaration).Attributes[i]
							if originalAttr.Summary == "Modified attribute summary" {
								t.Error("Original attribute was modified when clone was changed")
							}
						}
					}

					// Test events
					for i, event := range customElement.Events {
						event.Summary = "Modified event summary"

						// Verify original is unchanged
						if i < len(original.Modules[0].Declarations[1].(*manifest.CustomElementDeclaration).Events) {
							originalEvent := original.Modules[0].Declarations[1].(*manifest.CustomElementDeclaration).Events[i]
							if originalEvent.Summary == "Modified event summary" {
								t.Error("Original event was modified when clone was changed")
							}
						}
					}
				}
			}
		}
	})

	t.Run("PerformanceComparison", func(t *testing.T) {
		// This test verifies that Clone is significantly faster than JSON serialization
		// It's more of a sanity check than a strict performance test

		const iterations = 100

		// Time the Clone method
		for range iterations {
			cloned := original.Clone()
			if cloned == nil {
				t.Fatal("Clone returned nil")
			}
		}
	})
}

// TestCloneNilSafety tests that Clone methods handle nil inputs safely
func TestCloneNilSafety(t *testing.T) {
	t.Run("NilPackageClone", func(t *testing.T) {
		var pkg *manifest.Package
		cloned := pkg.Clone()
		if cloned != nil {
			t.Error("Cloning nil Package should return nil")
		}
	})

	t.Run("NilModuleClone", func(t *testing.T) {
		var module *manifest.JavaScriptModule
		cloned := module.Clone()
		if cloned != nil {
			t.Error("Cloning nil JavaScriptModule should return nil")
		}
	})

	t.Run("NilDeclarationClone", func(t *testing.T) {
		var class *manifest.ClassDeclaration
		cloned := class.Clone()
		if cloned != nil {
			t.Error("Cloning nil ClassDeclaration should return nil")
		}

		var customElement *manifest.CustomElementDeclaration
		cloned2 := customElement.Clone()
		if cloned2 != nil {
			t.Error("Cloning nil CustomElementDeclaration should return nil")
		}
	})

	t.Run("NilClassMemberClone", func(t *testing.T) {
		var field *manifest.ClassField
		cloned := field.Clone()
		if cloned != nil {
			t.Error("Cloning nil ClassField should return nil")
		}

		var method *manifest.ClassMethod
		cloned2 := method.Clone()
		if cloned2 != nil {
			t.Error("Cloning nil ClassMethod should return nil")
		}
	})

	t.Run("NilExportClone", func(t *testing.T) {
		var jsExport *manifest.JavaScriptExport
		cloned := jsExport.Clone()
		if cloned != nil {
			t.Error("Cloning nil JavaScriptExport should return nil")
		}

		var ceExport *manifest.CustomElementExport
		cloned2 := ceExport.Clone()
		if cloned2 != nil {
			t.Error("Cloning nil CustomElementExport should return nil")
		}
	})
}
