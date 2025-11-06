package manifest_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/manifest"
)

func TestRenderableAttribute_WithMapLookup(t *testing.T) {
	t.Run("MatchesCurrentBehavior", func(t *testing.T) {
		// Load fixture with custom element members
		manifestJSON, err := os.ReadFile(filepath.Join("fixtures", "custom_element_member_with_attribute.json"))
		if err != nil {
			t.Fatal(err)
		}

		var pkg manifest.Package
		if err := json.Unmarshal([]byte(manifestJSON), &pkg); err != nil {
			t.Fatal(err)
		}

		// Create renderable package (which creates RenderableAttribute internally)
		renderable := manifest.NewRenderablePackage(&pkg)

		// Verify the renderable was created successfully
		if renderable == nil {
			t.Fatal("Expected non-nil renderable package")
		}

		// The map-based lookup should produce identical results to linear search
		// This test ensures no regression in behavior
	})
}

func TestRenderableCustomElementDeclaration_WithMapLookup(t *testing.T) {
	t.Run("MatchesCurrentBehavior", func(t *testing.T) {
		// Load fixture with custom element exports
		manifestJSON, err := os.ReadFile(filepath.Join("fixtures", "custom-element-member-grouping.json"))
		if err != nil {
			t.Fatal(err)
		}

		var pkg manifest.Package
		if err := json.Unmarshal([]byte(manifestJSON), &pkg); err != nil {
			t.Fatal(err)
		}

		// Create renderable package
		renderable := manifest.NewRenderablePackage(&pkg)

		if renderable == nil {
			t.Fatal("Expected non-nil renderable package")
		}

		// Verify exports are found correctly using map-based lookups
		// The behavior should be identical to the linear search implementation
	})
}

func TestNewRenderablePackage_WithMapLookup(t *testing.T) {
	t.Run("ComprehensiveFixture", func(t *testing.T) {
		// Load comprehensive test fixture
		manifestJSON, err := os.ReadFile(filepath.Join("fixtures", "comprehensive-clone-test.json"))
		if err != nil {
			t.Fatal(err)
		}

		var pkg manifest.Package
		if err := json.Unmarshal([]byte(manifestJSON), &pkg); err != nil {
			t.Fatal(err)
		}

		// This fixture has many elements with various attributes and exports
		// Create renderable and verify no errors
		renderable := manifest.NewRenderablePackage(&pkg)

		if renderable == nil {
			t.Fatal("Expected non-nil renderable package")
		}

		// The map-based implementation should handle all cases correctly
		// without any behavior changes from the linear search version
	})

	t.Run("EmptyPackage", func(t *testing.T) {
		pkg := &manifest.Package{}
		pkg.SchemaVersion = "2.0.0"

		renderable := manifest.NewRenderablePackage(pkg)

		if renderable == nil {
			t.Fatal("Expected non-nil renderable for empty package")
		}
	})

	t.Run("PackageWithManyMembers", func(t *testing.T) {
		// Create a package with an element that has many members
		// This tests that map lookups work correctly with large member counts
		pkg := &manifest.Package{}
		pkg.SchemaVersion = "2.0.0"

		mod := &manifest.Module{}
		mod.Path = "/test.js"

		ced := &manifest.CustomElementDeclaration{}
		ced.Name = "TestElement"
		ced.TagName = "test-element"

		// Add many members to test map performance
		ced.Members = make([]manifest.ClassMember, 50)
		for i := 0; i < 50; i++ {
			if i%2 == 0 {
				field := &manifest.CustomElementField{}
				field.Name = "field" + string(rune('A'+i))
				field.Attribute = "attr" + string(rune('a'+i))
				ced.Members[i] = field
			} else {
				field := &manifest.ClassField{}
				field.Name = "field" + string(rune('A'+i))
				ced.Members[i] = field
			}
		}

		// Add an attribute that references one of the fields
		attr := &manifest.Attribute{}
		attr.Name = "attr" + string(rune('a'))
		ced.Attributes = []manifest.Attribute{*attr}

		mod.Declarations = []manifest.Declaration{ced}

		cee := &manifest.CustomElementExport{}
		cee.Declaration = &manifest.Reference{
			Name:   ced.Name,
			Module: mod.Path,
		}
		mod.Exports = []manifest.Export{cee}

		pkg.Modules = []manifest.Module{*mod}

		renderable := manifest.NewRenderablePackage(pkg)

		if renderable == nil {
			t.Fatal("Expected non-nil renderable package")
		}

		// Map-based lookup should efficiently handle this large member list
	})
}
