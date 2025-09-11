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
	"strings"
	"testing"

	"bennypowers.dev/cem/mcp"
	W "bennypowers.dev/cem/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegistry_NewRegistry(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("./test-fixtures/multiple-elements")

	registry, err := mcp.NewRegistry(workspace)
	require.NoError(t, err, "Failed to create registry")
	assert.NotNil(t, registry, "Registry should not be nil")
}

func TestRegistry_LoadManifests(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("./test-fixtures/multiple-elements")
	err := workspace.Init()
	require.NoError(t, err)

	registry, err := mcp.NewRegistry(workspace)
	require.NoError(t, err, "Failed to create registry")

	err = registry.LoadManifests()
	require.NoError(t, err, "Failed to load manifests")
}

func TestRegistry_GetElementInfo(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("./test-fixtures/multiple-elements")
	err := workspace.Init()
	require.NoError(t, err)

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
			name:        "button element",
			tagName:     "button-element",
			expectError: false,
			checkFunc: func(t *testing.T, info *mcp.ElementInfo) {
				assert.Equal(t, "button-element", info.TagName)
				assert.Equal(t, "button-element", info.Name)
				assert.Equal(t, "", info.Description) // Description not available from CustomElement

				// Test type-specific accessors
				assert.Len(t, info.Attributes(), 3) // variant, size, disabled
				assert.Len(t, info.Slots(), 2)      // default, icon
				assert.Len(t, info.Events(), 1)     // button-click
				assert.Len(t, info.CssProperties(), 2) // --button-color, --button-padding
				assert.Len(t, info.CssParts(), 1)   // button

				// Check key attributes
				variantAttr := findAttribute(info.Attributes(), "variant")
				require.NotNil(t, variantAttr, "Should have variant attribute")
				assert.Equal(t, "\"primary\" | \"secondary\" | \"ghost\"", variantAttr.Type())
				assert.Equal(t, "\"primary\"", variantAttr.Default())
				assert.Contains(t, variantAttr.Values(), "primary")
				assert.Contains(t, variantAttr.Values(), "secondary")
				assert.Contains(t, variantAttr.Values(), "ghost")

				sizeAttr := findAttribute(info.Attributes(), "size")
				require.NotNil(t, sizeAttr, "Should have size attribute")
				assert.Equal(t, "\"small\" | \"medium\" | \"large\"", sizeAttr.Type())
				assert.Equal(t, "\"medium\"", sizeAttr.Default())

				disabledAttr := findAttribute(info.Attributes(), "disabled")
				require.NotNil(t, disabledAttr, "Should have disabled attribute")
				assert.Equal(t, "boolean", disabledAttr.Type())

				// Check slots
				defaultSlot := findSlot(info.Slots(), "")
				require.NotNil(t, defaultSlot, "Should have default slot")
				assert.Equal(t, "Button content", defaultSlot.Description())

				iconSlot := findSlot(info.Slots(), "icon")
				require.NotNil(t, iconSlot, "Should have icon slot")
				assert.Equal(t, "Button icon", iconSlot.Description())

				// Check CSS properties
				colorProp := findCssProperty(info.CssProperties(), "--button-color")
				require.NotNil(t, colorProp, "Should have --button-color property")
				assert.Equal(t, "blue", colorProp.Initial())
				assert.Equal(t, "<color>", colorProp.Syntax())

				paddingProp := findCssProperty(info.CssProperties(), "--button-padding")
				require.NotNil(t, paddingProp, "Should have --button-padding property")
				assert.Equal(t, "8px", paddingProp.Initial())
				assert.Equal(t, "<length>", paddingProp.Syntax())

				// Check CSS parts
				buttonPart := findCssPart(info.CssParts(), "button")
				require.NotNil(t, buttonPart, "Should have button part")
				assert.Equal(t, "The button element", buttonPart.Description())
			},
		},
		{
			name:        "card element",
			tagName:     "card-element",
			expectError: false,
			checkFunc: func(t *testing.T, info *mcp.ElementInfo) {
				assert.Equal(t, "card-element", info.TagName)
				assert.Equal(t, "card-element", info.Name)
				assert.Equal(t, "", info.Description) // Description not available from CustomElement

				// Test type-specific accessors
				assert.Len(t, info.Attributes(), 1)   // elevation
				assert.Len(t, info.Slots(), 3)        // default, header, footer
				assert.Len(t, info.Events(), 0)       // no events
				assert.Len(t, info.CssProperties(), 1) // --card-background
				assert.Len(t, info.CssParts(), 4)     // container, header, content, footer

				// Check elevation attribute
				elevationAttr := findAttribute(info.Attributes(), "elevation")
				require.NotNil(t, elevationAttr, "Should have elevation attribute")
				assert.Equal(t, "number", elevationAttr.Type())
				assert.Equal(t, "1", elevationAttr.Default())

				// Check slots
				defaultSlot := findSlot(info.Slots(), "")
				require.NotNil(t, defaultSlot, "Should have default slot")
				assert.Equal(t, "Card content", defaultSlot.Description())

				headerSlot := findSlot(info.Slots(), "header")
				require.NotNil(t, headerSlot, "Should have header slot")
				assert.Equal(t, "Card header", headerSlot.Description())

				footerSlot := findSlot(info.Slots(), "footer")
				require.NotNil(t, footerSlot, "Should have footer slot")
				assert.Equal(t, "Card footer", footerSlot.Description())

				// Check CSS properties
				backgroundProp := findCssProperty(info.CssProperties(), "--card-background")
				require.NotNil(t, backgroundProp, "Should have --card-background property")
				assert.Equal(t, "white", backgroundProp.Initial())
				assert.Equal(t, "<color>", backgroundProp.Syntax())

				// Check CSS parts
				containerPart := findCssPart(info.CssParts(), "container")
				require.NotNil(t, containerPart, "Should have container part")
				assert.Equal(t, "Card container", containerPart.Description())

				headerPart := findCssPart(info.CssParts(), "header")
				require.NotNil(t, headerPart, "Should have header part")
				assert.Equal(t, "Card header part", headerPart.Description())
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
	workspace := W.NewFileSystemWorkspaceContext("./test-fixtures/multiple-elements")
	err := workspace.Init()
	require.NoError(t, err)

	registry, err := mcp.NewRegistry(workspace)
	require.NoError(t, err, "Failed to create registry")

	err = registry.LoadManifests()
	require.NoError(t, err, "Failed to load manifests")

	elements := registry.GetAllElements()
	assert.Len(t, elements, 2)
	assert.Contains(t, elements, "button-element")
	assert.Contains(t, elements, "card-element")

	buttonElement := elements["button-element"]
	assert.Equal(t, "button-element", buttonElement.TagName)
	assert.Equal(t, "button-element", buttonElement.Name)
	assert.Equal(t, "", buttonElement.Description) // Description not available from CustomElement

	cardElement := elements["card-element"]
	assert.Equal(t, "card-element", cardElement.TagName)
	assert.Equal(t, "card-element", cardElement.Name)
	assert.Equal(t, "", cardElement.Description) // Description not available from CustomElement
}

func TestRegistry_GetManifestSchema(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("./test-fixtures/multiple-elements")
	err := workspace.Init()
	require.NoError(t, err)

	registry, err := mcp.NewRegistry(workspace)
	require.NoError(t, err, "Failed to create registry")

	schema, err := registry.GetManifestSchema()
	require.NoError(t, err, "Failed to get manifest schema")
	assert.NotNil(t, schema)

	// Check that schema has required fields from the actual custom elements manifest schema
	assert.Equal(t, "http://json-schema.org/draft-07/schema#", schema["$schema"])
	
	// The actual schema has a versioned title
	title, ok := schema["title"].(string)
	require.True(t, ok, "Schema should have a title")
	assert.Contains(t, title, "Custom Elements Manifest Schema")
	
	// The actual schema should have type at root level or in allOf/anyOf structure
	// Check for either direct type or complex schema structure
	if schemaType, exists := schema["type"]; exists {
		assert.Equal(t, "object", schemaType)
	} else {
		// Complex schema structure - just verify it's a valid schema
		assert.True(t, schema["$schema"] != nil, "Should have valid schema structure")
	}

	// Check that we got a valid custom elements manifest schema
	// The schema should have definitions for the manifest structure
	if definitions, ok := schema["definitions"].(map[string]interface{}); ok {
		assert.NotEmpty(t, definitions, "Schema should have definitions")
		
		// Look for manifest or package definition
		hasManifestDef := false
		for defName := range definitions {
			if strings.Contains(strings.ToLower(defName), "manifest") || 
			   strings.Contains(strings.ToLower(defName), "package") {
				hasManifestDef = true
				break
			}
		}
		assert.True(t, hasManifestDef, "Schema should have manifest-related definitions")
	} else {
		// If no definitions, check for $ref which indicates external schema reference
		if ref, ok := schema["$ref"].(string); ok {
			assert.NotEmpty(t, ref, "Schema should have valid $ref")
		} else {
			t.Errorf("Schema should have either definitions or $ref")
		}
	}
	
	// The rest of the test just verifies we got a valid schema
	// Since we're now using the real schema from V.GetSchema(), the exact structure
	// may vary but it should be a valid JSON schema
}

// Table-driven tests for detailed data conversion testing

func TestRegistry_AttributeConversion(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("./test-fixtures/multiple-elements")
	err := workspace.Init()
	require.NoError(t, err)

	registry, err := mcp.NewRegistry(workspace)
	require.NoError(t, err)

	err = registry.LoadManifests()
	require.NoError(t, err)

	tests := []struct {
		name             string
		tagName          string
		attrName         string
		expectedName     string
		expectedType     string
		expectedDesc     string
		expectedDefault  string
		expectedRequired bool
		expectedValues   []string
	}{
		{
			name:             "button variant attribute",
			tagName:          "button-element",
			attrName:         "variant",
			expectedName:     "variant",
			expectedType:     "\"primary\" | \"secondary\" | \"ghost\"",
			expectedDesc:     "Button variant",
			expectedDefault:  "\"primary\"",
			expectedRequired: false,
			expectedValues:   []string{"primary", "secondary", "ghost"},
		},
		{
			name:             "button size attribute",
			tagName:          "button-element",
			attrName:         "size",
			expectedName:     "size",
			expectedType:     "\"small\" | \"medium\" | \"large\"",
			expectedDesc:     "Button size",
			expectedDefault:  "\"medium\"",
			expectedRequired: false,
			expectedValues:   []string{"small", "medium", "large"},
		},
		{
			name:             "button disabled attribute",
			tagName:          "button-element",
			attrName:         "disabled",
			expectedName:     "disabled",
			expectedType:     "boolean",
			expectedDesc:     "Whether button is disabled",
			expectedDefault:  "",
			expectedRequired: false,
			expectedValues:   []string{},
		},
		{
			name:             "card elevation attribute",
			tagName:          "card-element",
			attrName:         "elevation",
			expectedName:     "elevation",
			expectedType:     "number",
			expectedDesc:     "Card elevation level",
			expectedDefault:  "1",
			expectedRequired: false,
			expectedValues:   []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			elementInfo, err := registry.GetElementInfo(test.tagName)
			require.NoError(t, err, "Failed to get element info for %s", test.tagName)

			// Find the attribute using the interface-based approach
			var foundAttr mcp.Attribute
			for _, attr := range elementInfo.Attributes() {
				if attr.Name() == test.attrName {
					foundAttr = attr
					break
				}
			}
			require.NotNil(t, foundAttr, "Attribute %s not found for element %s", test.attrName, test.tagName)

			assert.Equal(t, test.expectedName, foundAttr.Name())
			assert.Equal(t, test.expectedType, foundAttr.Type())
			assert.Equal(t, test.expectedDesc, foundAttr.Description())
			assert.Equal(t, test.expectedDefault, foundAttr.Default())
			assert.Equal(t, test.expectedRequired, foundAttr.Required())
			assert.ElementsMatch(t, test.expectedValues, foundAttr.Values())
		})
	}
}

func TestRegistry_SlotConversion(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("./test-fixtures/multiple-elements")
	err := workspace.Init()
	require.NoError(t, err)

	registry, err := mcp.NewRegistry(workspace)
	require.NoError(t, err)

	err = registry.LoadManifests()
	require.NoError(t, err)

	tests := []struct {
		name         string
		tagName      string
		slotName     string
		expectedName string
		expectedDesc string
	}{
		{
			name:         "button default slot",
			tagName:      "button-element",
			slotName:     "",
			expectedName: "",
			expectedDesc: "Button content",
		},
		{
			name:         "button icon slot",
			tagName:      "button-element",
			slotName:     "icon",
			expectedName: "icon",
			expectedDesc: "Button icon",
		},
		{
			name:         "card default slot",
			tagName:      "card-element",
			slotName:     "",
			expectedName: "",
			expectedDesc: "Card content",
		},
		{
			name:         "card header slot",
			tagName:      "card-element",
			slotName:     "header",
			expectedName: "header",
			expectedDesc: "Card header",
		},
		{
			name:         "card footer slot",
			tagName:      "card-element",
			slotName:     "footer",
			expectedName: "footer",
			expectedDesc: "Card footer",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			elementInfo, err := registry.GetElementInfo(test.tagName)
			require.NoError(t, err, "Failed to get element info for %s", test.tagName)

			slot := findSlot(elementInfo.Slots(), test.slotName)
			require.NotNil(t, slot, "Slot %s not found for element %s", test.slotName, test.tagName)

			assert.Equal(t, test.expectedName, slot.Name())
			assert.Equal(t, test.expectedDesc, slot.Description())
		})
	}
}

func TestRegistry_CssPropertyConversion(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("./test-fixtures/multiple-elements")
	err := workspace.Init()
	require.NoError(t, err)

	registry, err := mcp.NewRegistry(workspace)
	require.NoError(t, err)

	err = registry.LoadManifests()
	require.NoError(t, err)

	tests := []struct {
		name             string
		tagName          string
		propName         string
		expectedName     string
		expectedSyntax   string
		expectedInherits bool
		expectedInitial  string
		expectedDesc     string
	}{
		{
			name:             "button color property",
			tagName:          "button-element",
			propName:         "--button-color",
			expectedName:     "--button-color",
			expectedSyntax:   "<color>",
			expectedInherits: false,
			expectedInitial:  "blue",
			expectedDesc:     "Button color",
		},
		{
			name:             "button padding property",
			tagName:          "button-element",
			propName:         "--button-padding",
			expectedName:     "--button-padding",
			expectedSyntax:   "<length>",
			expectedInherits: false,
			expectedInitial:  "8px",
			expectedDesc:     "Button padding",
		},
		{
			name:             "card background property",
			tagName:          "card-element",
			propName:         "--card-background",
			expectedName:     "--card-background",
			expectedSyntax:   "<color>",
			expectedInherits: false,
			expectedInitial:  "white",
			expectedDesc:     "Card background color",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			elementInfo, err := registry.GetElementInfo(test.tagName)
			require.NoError(t, err, "Failed to get element info for %s", test.tagName)

			prop := findCssProperty(elementInfo.CssProperties(), test.propName)
			require.NotNil(t, prop, "CSS property %s not found for element %s", test.propName, test.tagName)

			assert.Equal(t, test.expectedName, prop.Name())
			assert.Equal(t, test.expectedSyntax, prop.Syntax())
			assert.Equal(t, test.expectedInherits, prop.Inherits())
			assert.Equal(t, test.expectedInitial, prop.Initial())
			assert.Equal(t, test.expectedDesc, prop.Description())
		})
	}
}

func TestRegistry_CssPartConversion(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("./test-fixtures/multiple-elements")
	err := workspace.Init()
	require.NoError(t, err)

	registry, err := mcp.NewRegistry(workspace)
	require.NoError(t, err)

	err = registry.LoadManifests()
	require.NoError(t, err)

	tests := []struct {
		name         string
		tagName      string
		partName     string
		expectedName string
		expectedDesc string
	}{
		{
			name:         "button part",
			tagName:      "button-element",
			partName:     "button",
			expectedName: "button",
			expectedDesc: "The button element",
		},
		{
			name:         "card container part",
			tagName:      "card-element",
			partName:     "container",
			expectedName: "container",
			expectedDesc: "Card container",
		},
		{
			name:         "card header part",
			tagName:      "card-element",
			partName:     "header",
			expectedName: "header",
			expectedDesc: "Card header part",
		},
		{
			name:         "card content part",
			tagName:      "card-element",
			partName:     "content",
			expectedName: "content",
			expectedDesc: "Card content part",
		},
		{
			name:         "card footer part",
			tagName:      "card-element",
			partName:     "footer",
			expectedName: "footer",
			expectedDesc: "Card footer part",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			elementInfo, err := registry.GetElementInfo(test.tagName)
			require.NoError(t, err, "Failed to get element info for %s", test.tagName)

			part := findCssPart(elementInfo.CssParts(), test.partName)
			require.NotNil(t, part, "CSS part %s not found for element %s", test.partName, test.tagName)

			assert.Equal(t, test.expectedName, part.Name())
			assert.Equal(t, test.expectedDesc, part.Description())
		})
	}
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

func findCssPart(parts []mcp.CssPart, name string) mcp.CssPart {
	for _, part := range parts {
		if part.Name() == name {
			return part
		}
	}
	return nil
}
