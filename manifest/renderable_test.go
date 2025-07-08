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
package manifest_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/manifest"
	"github.com/stretchr/testify/assert"
)

func loadFixture(t *testing.T, path string) *manifest.Package {
	t.Helper()
	bytes, err := os.ReadFile(filepath.Join("fixtures", path))
	assert.NoError(t, err)
	var pkg manifest.Package
	err = json.Unmarshal(bytes, &pkg)
	assert.NoError(t, err)
	return &pkg
}

func TestRenderableAttributeToTableRow(t *testing.T) {
	pkg := loadFixture(t, "custom_element_attributes.json")
	ced := pkg.Modules[0].Declarations[0].(*manifest.CustomElementDeclaration)
	renderable := manifest.NewRenderableAttribute(&ced.Attributes[0], ced, nil, &pkg.Modules[0])
	row := renderable.ToTableRow()
	expected := []string{"variant", "", "", "", "", "\"primary\" | \"secondary\""}
	assert.Equal(t, expected, row)
}

func TestRenderableSlotToTableRow(t *testing.T) {
	pkg := loadFixture(t, "custom_element_slots.json")
	ced := pkg.Modules[0].Declarations[0].(*manifest.CustomElementDeclaration)
	renderable := manifest.NewRenderableSlot(&ced.Slots[0], ced, nil, &pkg.Modules[0])
	row := renderable.ToTableRow()
	expected := []string{"<default>", ""}
	assert.Equal(t, expected, row)
}

func TestRenderableEventToTableRow(t *testing.T) {
	pkg := loadFixture(t, "custom_element_events.json")
	ced := pkg.Modules[0].Declarations[0].(*manifest.CustomElementDeclaration)
	renderable := manifest.NewRenderableEvent(&ced.Events[0], ced, nil, &pkg.Modules[0])
	row := renderable.ToTableRow()
	expected := []string{"my-card-selected", "CustomEvent", ""}
	assert.Equal(t, expected, row)
}

func TestRenderableCssCustomPropertyToTableRow(t *testing.T) {
	pkg := loadFixture(t, "custom_element_css_parts_properties_states.json")
	ced := pkg.Modules[0].Declarations[0].(*manifest.CustomElementDeclaration)
	renderable := manifest.NewRenderableCssCustomProperty(&ced.CssProperties[0], ced, nil, &pkg.Modules[0])
	row := renderable.ToTableRow()
	expected := []string{"--my-card-color", "", "", ""}
	assert.Equal(t, expected, row)
}

func TestRenderableCssPartToTableRow(t *testing.T) {
	pkg := loadFixture(t, "custom_element_css_parts_properties_states.json")
	ced := pkg.Modules[0].Declarations[0].(*manifest.CustomElementDeclaration)
	renderable := manifest.NewRenderableCssPart(&ced.CssParts[0], ced, nil, &pkg.Modules[0])
	row := renderable.ToTableRow()
	expected := []string{"header", ""}
	assert.Equal(t, expected, row)
}