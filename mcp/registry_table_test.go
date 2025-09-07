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

func TestRegistry_ConvertAttributes_TableDriven(t *testing.T) {
	workspace, cleanup := createTestWorkspace(t, "multiple-elements")
	defer cleanup()

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

			// Find the attribute using the new interface-based approach
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

func TestRegistry_ConvertSlots_TableDriven(t *testing.T) {
	workspace, cleanup := createTestWorkspace(t, "multiple-elements")
	defer cleanup()

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

func TestRegistry_ConvertCssProperties_TableDriven(t *testing.T) {
	workspace, cleanup := createTestWorkspace(t, "multiple-elements")
	defer cleanup()

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

func TestRegistry_ConvertCssParts_TableDriven(t *testing.T) {
	workspace, cleanup := createTestWorkspace(t, "multiple-elements")
	defer cleanup()

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

// Helper function for CSS parts
func findCssPart(parts []mcp.CssPart, name string) mcp.CssPart {
	for _, part := range parts {
		if part.Name() == name {
			return part
		}
	}
	return nil
}
