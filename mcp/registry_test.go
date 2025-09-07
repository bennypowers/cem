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
package mcp_test

import (
	"testing"

	"bennypowers.dev/cem/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegistry_NewRegistry(t *testing.T) {
	workspace, cleanup := createTestWorkspace(t, "basic")
	defer cleanup()

	registry, err := mcp.NewRegistry(workspace)
	require.NoError(t, err, "Failed to create registry")
	assert.NotNil(t, registry, "Registry should not be nil")
}

func TestRegistry_LoadManifests(t *testing.T) {
	workspace, cleanup := createTestWorkspace(t, "basic")
	defer cleanup()

	registry, err := mcp.NewRegistry(workspace)
	require.NoError(t, err, "Failed to create registry")

	err = registry.LoadManifests()
	require.NoError(t, err, "Failed to load manifests")
}

func TestRegistry_GetElementInfo(t *testing.T) {
	workspace, cleanup := createTestWorkspace(t, "basic")
	defer cleanup()

	registry, err := mcp.NewRegistry(workspace)
	require.NoError(t, err, "Failed to create registry")

	err = registry.LoadManifests()
	require.NoError(t, err, "Failed to load manifests")

	tests := []struct {
		name        string
		tagName     string
		expectError bool
		checkFunc   func(t *testing.T, info *mcp.ElementInfo)
	}{
		{
			name:        "existing element",
			tagName:     "test-element",
			expectError: false,
			checkFunc: func(t *testing.T, info *mcp.ElementInfo) {
				assert.Equal(t, "test-element", info.TagName)
				assert.Equal(t, "test-element", info.Name)
				// Test unified Items array - should contain all item types
				assert.Len(t, info.Items, 8) // 2 attrs + 2 slots + 1 event + 1 CSS prop + 1 CSS part + 1 CSS state

				// Test type-specific accessors
				assert.Len(t, info.Attributes(), 2)
				assert.Len(t, info.Slots(), 2)
				assert.Len(t, info.Events(), 1)
				assert.Len(t, info.CssProperties(), 1)
				assert.Len(t, info.CssParts(), 1)
				assert.Len(t, info.CssStates(), 1)

				// Check attribute details
				variantAttr := findAttribute(info.Attributes(), "variant")
				require.NotNil(t, variantAttr, "Should have variant attribute")
				assert.Equal(t, "\"primary\" | \"secondary\"", variantAttr.Type())
				assert.Equal(t, "\"primary\"", variantAttr.Default())
				assert.Contains(t, variantAttr.Values(), "primary")
				assert.Contains(t, variantAttr.Values(), "secondary")

				disabledAttr := findAttribute(info.Attributes(), "disabled")
				require.NotNil(t, disabledAttr, "Should have disabled attribute")
				assert.Equal(t, "boolean", disabledAttr.Type())

				// Check slots
				defaultSlot := findSlot(info.Slots(), "")
				require.NotNil(t, defaultSlot, "Should have default slot")
				headerSlot := findSlot(info.Slots(), "header")
				require.NotNil(t, headerSlot, "Should have header slot")

				// Check CSS properties
				colorProp := findCssProperty(info.CssProperties(), "--test-color")
				require.NotNil(t, colorProp, "Should have --test-color property")
				assert.Equal(t, "blue", colorProp.Initial())
				assert.Equal(t, "<color>", colorProp.Syntax())
			},
		},
		{
			name:        "non-existent element",
			tagName:     "non-existent",
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			info, err := registry.GetElementInfo(test.tagName)

			if test.expectError {
				assert.Error(t, err)
				assert.Nil(t, info)
			} else {
				require.NoError(t, err)
				require.NotNil(t, info)
				if test.checkFunc != nil {
					test.checkFunc(t, info)
				}
			}
		})
	}
}

func TestRegistry_GetAllElements(t *testing.T) {
	workspace, cleanup := createTestWorkspace(t, "basic")
	defer cleanup()

	registry, err := mcp.NewRegistry(workspace)
	require.NoError(t, err, "Failed to create registry")

	err = registry.LoadManifests()
	require.NoError(t, err, "Failed to load manifests")

	elements := registry.GetAllElements()
	assert.Len(t, elements, 1)
	assert.Contains(t, elements, "test-element")

	testElement := elements["test-element"]
	assert.Equal(t, "test-element", testElement.TagName)
	assert.Equal(t, "test-element", testElement.Name)
}

func TestRegistry_GetManifestSchema(t *testing.T) {
	workspace, cleanup := createTestWorkspace(t, "basic")
	defer cleanup()

	registry, err := mcp.NewRegistry(workspace)
	require.NoError(t, err, "Failed to create registry")

	schema, err := registry.GetManifestSchema()
	require.NoError(t, err, "Failed to get manifest schema")
	assert.NotNil(t, schema)

	// Check that schema has required fields
	assert.Equal(t, "http://json-schema.org/draft-07/schema#", schema["$schema"])
	assert.Equal(t, "Custom Elements Manifest", schema["title"])
	assert.Equal(t, "object", schema["type"])

	properties, ok := schema["properties"].(map[string]interface{})
	require.True(t, ok, "Schema should have properties")
	assert.Contains(t, properties, "schemaVersion")
	assert.Contains(t, properties, "modules")

	required, ok := schema["required"].([]string)
	require.True(t, ok, "Schema should have required array")
	assert.Contains(t, required, "schemaVersion")
	assert.Contains(t, required, "modules")
}

// Helper functions for testing

func findAttribute(attributes []mcp.Attribute, name string) mcp.Attribute {
	for _, attr := range attributes {
		if attr.Name() == name {
			return attr
		}
	}
	return nil
}

func findSlot(slots []mcp.Slot, name string) mcp.Slot {
	for _, slot := range slots {
		if slot.Name() == name {
			return slot
		}
	}
	return nil
}

func findCssProperty(props []mcp.CssProperty, name string) mcp.CssProperty {
	for _, prop := range props {
		if prop.Name() == name {
			return prop
		}
	}
	return nil
}
