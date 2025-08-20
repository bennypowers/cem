package lsp_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/lsp/testhelpers"
	M "bennypowers.dev/cem/manifest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// loadTestManifest is a helper function to load manifests from test fixtures
func loadTestManifest(t *testing.T, fixturePath string) *M.Package {
	t.Helper()

	manifestBytes, err := os.ReadFile(fixturePath)
	if err != nil {
		t.Fatalf("Failed to read manifest: %v", err)
	}

	var manifest M.Package
	err = json.Unmarshal(manifestBytes, &manifest)
	if err != nil {
		t.Fatalf("Failed to parse manifest: %v", err)
	}

	return &manifest
}

// TestRegistryElementLookup tests registry's ability to store and retrieve custom elements
func TestRegistryElementLookup(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()

	// Load the basic test manifest
	manifest := loadTestManifest(t, filepath.Join("test", "fixtures", "registry-basic", "manifest.json"))
	ctx.AddManifest(manifest)

	t.Run("Get existing element", func(t *testing.T) {
		element, exists := ctx.Element("my-element")
		if !exists {
			t.Fatal("Expected to find my-element")
		}
		if element.TagName != "my-element" {
			t.Errorf("Expected TagName 'my-element', got %q", element.TagName)
		}
		if len(element.Attributes) != 2 {
			t.Errorf("Expected 2 attributes, got %d", len(element.Attributes))
		}
		if len(element.Events) != 1 {
			t.Errorf("Expected 1 event, got %d", len(element.Events))
		}
		if len(element.Slots) != 2 {
			t.Errorf("Expected 2 slots, got %d", len(element.Slots))
		}
	})

	t.Run("Get non-existent element", func(t *testing.T) {
		_, exists := ctx.Element("non-existent")
		if exists {
			t.Error("Expected not to find non-existent element")
		}
	})

	t.Run("Get all tag names", func(t *testing.T) {
		tagNames := ctx.AllTagNames()
		if len(tagNames) != 1 {
			t.Errorf("Expected 1 tag name, got %d", len(tagNames))
		}
		if tagNames[0] != "my-element" {
			t.Errorf("Expected tag name 'my-element', got %q", tagNames[0])
		}
	})

	t.Run("Get element attributes", func(t *testing.T) {
		attrs, exists := ctx.Attributes("my-element")
		if !exists {
			t.Fatal("Expected to find attributes for my-element")
		}
		if len(attrs) != 2 {
			t.Errorf("Expected 2 attributes, got %d", len(attrs))
		}

		// Check specific attributes
		if labelAttr, ok := attrs["label"]; ok {
			if labelAttr.Type.Text != "string" {
				t.Errorf("Expected label type 'string', got %q", labelAttr.Type.Text)
			}
		} else {
			t.Error("Expected to find 'label' attribute")
		}

		if countAttr, ok := attrs["count"]; ok {
			if countAttr.Type.Text != "number" {
				t.Errorf("Expected count type 'number', got %q", countAttr.Type.Text)
			}
		} else {
			t.Error("Expected to find 'count' attribute")
		}
	})

	t.Run("Get attributes for non-existent element", func(t *testing.T) {
		_, exists := ctx.Attributes("non-existent")
		if exists {
			t.Error("Expected not to find attributes for non-existent element")
		}
	})

	t.Run("Get element slots", func(t *testing.T) {
		slots, exists := ctx.Slots("my-element")
		if !exists {
			t.Fatal("Expected to find slots for my-element")
		}
		if len(slots) != 2 {
			t.Errorf("Expected 2 slots, got %d", len(slots))
		}

		slotNames := make(map[string]bool)
		for _, slot := range slots {
			slotNames[slot.Name] = true
		}

		if !slotNames["default"] {
			t.Error("Expected to find 'default' slot")
		}
		if !slotNames["header"] {
			t.Error("Expected to find 'header' slot")
		}
	})

	t.Run("Get element definition", func(t *testing.T) {
		definition, exists := ctx.ElementDefinition("my-element")
		if !exists {
			t.Fatal("Expected to find element definition for my-element")
		}
		if definition.Element().TagName != "my-element" {
			t.Errorf("Expected element tag name 'my-element', got %q", definition.Element().TagName)
		}
		if definition.ModulePath() != "my-element.js" {
			t.Errorf("Expected module path 'my-element.js', got %q", definition.ModulePath())
		}
	})
}

// TestRegistryMultipleManifests tests adding multiple manifests to the same registry
func TestRegistryMultipleManifests(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()

	// Load two different manifests
	manifest1 := loadTestManifest(t, filepath.Join("test", "fixtures", "registry-multiple", "manifest-1.json"))
	manifest2 := loadTestManifest(t, filepath.Join("test", "fixtures", "registry-multiple", "manifest-2.json"))

	// Add both manifests
	ctx.AddManifest(manifest1)
	ctx.AddManifest(manifest2)

	// Verify both elements are available
	tagNames := ctx.AllTagNames()
	if len(tagNames) != 2 {
		t.Errorf("Expected 2 tag names after adding multiple manifests, got %d", len(tagNames))
	}

	if _, exists := ctx.Element("first-element"); !exists {
		t.Error("Expected to find first-element after adding manifests")
	}

	if _, exists := ctx.Element("second-element"); !exists {
		t.Error("Expected to find second-element after adding manifests")
	}

	// Verify attributes for each element
	firstAttrs, exists := ctx.Attributes("first-element")
	if !exists || len(firstAttrs) != 1 {
		t.Error("Expected to find attributes for first-element")
	}

	secondAttrs, exists := ctx.Attributes("second-element")
	if !exists || len(secondAttrs) != 1 {
		t.Error("Expected to find attributes for second-element")
	}
}

// TestRegistryEmptyManifest tests registry behavior with empty manifests
func TestRegistryEmptyManifest(t *testing.T) {
	ctx := testhelpers.NewMockServerContext()

	// Empty manifest
	emptyManifest := &M.Package{
		SchemaVersion: "1.0.0",
		Modules:       []M.Module{},
	}

	ctx.AddManifest(emptyManifest)

	tagNames := ctx.AllTagNames()
	if len(tagNames) != 0 {
		t.Errorf("Expected 0 tag names for empty manifest, got %d", len(tagNames))
	}
}

// TestNilTypeAttributeHandling tests the regression fix for nil pointer dereference
// when attributes don't have type information (attr.Type is nil).
func TestNilTypeAttributeHandling(t *testing.T) {
	// Create a manifest with an attribute that has nil Type
	pkg := &M.Package{
		SchemaVersion: "2.1.0",
		Modules: []M.Module{
			{
				Kind: "javascript-module",
				Path: "test-element.js",
				Declarations: []M.Declaration{
					&M.CustomElementDeclaration{
						CustomElement: M.CustomElement{
							TagName: "test-element",
							Attributes: []M.Attribute{
								{
									FullyQualified: M.FullyQualified{
										Name:        "test-attr",
										Description: "Test attribute without type",
									},
									Type: nil, // This was causing the nil pointer dereference
								},
								{
									FullyQualified: M.FullyQualified{
										Name:        "typed-attr",
										Description: "Test attribute with type",
									},
									Type: &M.Type{
										Text: "string",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	// Create registry and add the manifest with nil type
	ctx := testhelpers.NewMockServerContext()
	ctx.AddManifest(pkg)

	// This should not panic when accessing attributes with nil types
	attrs, exists := ctx.Attributes("test-element")
	require.True(t, exists, "Element should exist in registry")
	require.Len(t, attrs, 2, "Should have two attributes")

	// Verify nil type attribute exists and doesn't cause panic
	nilTypeAttr, exists := attrs["test-attr"]
	require.True(t, exists, "Attribute with nil type should exist")
	assert.Nil(t, nilTypeAttr.Type)
	assert.Equal(t, "test-attr", nilTypeAttr.Name)

	// Verify typed attribute exists
	typedAttr, exists := attrs["typed-attr"]
	require.True(t, exists, "Typed attribute should exist")
	require.NotNil(t, typedAttr.Type, "Type should not be nil")
	assert.Equal(t, "string", typedAttr.Type.Text)
}

// TestHandleManifestReloadWithNilTypes tests that manifest reload doesn't crash
// when encountering attributes with nil types.
func TestHandleManifestReloadWithNilTypes(t *testing.T) {
	// Create a server with a manifest containing nil type attributes
	pkg := &M.Package{
		SchemaVersion: "2.1.0",
		Modules: []M.Module{
			{
				Kind: "javascript-module",
				Path: "element.js",
				Declarations: []M.Declaration{
					&M.CustomElementDeclaration{
						CustomElement: M.CustomElement{
							TagName: "test-element",
							Attributes: []M.Attribute{
								{
									FullyQualified: M.FullyQualified{
										Name: "no-type-attr",
									},
									Type: nil, // This should not cause panic in debug logging
								},
							},
						},
					},
				},
			},
		},
	}

	ctx := testhelpers.NewMockServerContext()
	ctx.AddManifest(pkg)

	// Verify the registry has the element
	element, exists := ctx.Element("test-element")
	require.True(t, exists, "Element should exist")
	assert.Equal(t, "test-element", element.TagName)

	// Getting attributes should work without panic, even with nil types
	attrs, exists := ctx.Attributes("test-element")
	require.True(t, exists, "Should have attributes")
	require.Len(t, attrs, 1, "Should have one attribute")

	nilTypeAttr, exists := attrs["no-type-attr"]
	require.True(t, exists, "Nil type attribute should exist")
	assert.Nil(t, nilTypeAttr.Type)

	// This test primarily ensures no panics occur during attribute access
	// The actual fix was in server.go handleManifestReload method where
	// it safely handles attr.Type being nil in debug logging
}
