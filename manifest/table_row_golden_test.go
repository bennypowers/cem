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
	"flag"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"bennypowers.dev/cem/manifest"
	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
)

var update = flag.Bool("update", false, "update golden files")
var ansiRegexp = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func stripANSI(s string) string {
	return ansiRegexp.ReplaceAllString(s, "")
}

func loadTestFixture(t *testing.T, path string) *manifest.Package {
	t.Helper()
	bytes, err := os.ReadFile(filepath.Join("fixtures", path))
	assert.NoError(t, err)
	var pkg manifest.Package
	err = json.Unmarshal(bytes, &pkg)
	assert.NoError(t, err)
	return &pkg
}

func checkGolden(t *testing.T, name string, actual []byte) {
	t.Helper()
	goldenPath := filepath.Join("fixtures", name+".md")
	if *update {
		err := os.WriteFile(goldenPath, actual, 0644)
		assert.NoError(t, err, "failed to update golden file")
		return
	}

	expected, err := os.ReadFile(goldenPath)
	assert.NoError(t, err, "failed to read golden file")
	assert.Equal(t, string(expected), string(actual))
}

func TestAttributeToTableRowGolden(t *testing.T) {
	pkg := loadTestFixture(t, "custom_element_attributes.json")
	ced := pkg.Modules[0].Declarations[0].(*manifest.CustomElementDeclaration)
	renderable := manifest.NewRenderableAttribute(&ced.Attributes[0], ced, nil, &pkg.Modules[0])

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	table, _ := pterm.DefaultTable.WithHasHeader(true).WithData([][]string{headers, row}).Srender()
	checkGolden(t, "attribute_row", []byte(stripANSI(table)))
}

func TestSlotToTableRowGolden(t *testing.T) {
	pkg := loadTestFixture(t, "custom_element_slots.json")
	ced := pkg.Modules[0].Declarations[0].(*manifest.CustomElementDeclaration)
	renderable := manifest.NewRenderableSlot(&ced.Slots[0], ced, nil, &pkg.Modules[0])

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	table, _ := pterm.DefaultTable.WithHasHeader(true).WithData([][]string{headers, row}).Srender()
	checkGolden(t, "slot_row", []byte(stripANSI(table)))
}

func TestEventToTableRowGolden(t *testing.T) {
	pkg := loadTestFixture(t, "custom_element_events.json")
	ced := pkg.Modules[0].Declarations[0].(*manifest.CustomElementDeclaration)
	renderable := manifest.NewRenderableEvent(&ced.Events[0], ced, nil, &pkg.Modules[0])

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	table, _ := pterm.DefaultTable.WithHasHeader(true).WithData([][]string{headers, row}).Srender()
	checkGolden(t, "event_row", []byte(stripANSI(table)))
}

func TestCssPropertyToTableRowGolden(t *testing.T) {
	pkg := loadTestFixture(t, "custom_element_css_parts_properties_states.json")
	ced := pkg.Modules[0].Declarations[0].(*manifest.CustomElementDeclaration)
	renderable := manifest.NewRenderableCssCustomProperty(&ced.CssProperties[0], ced, nil, &pkg.Modules[0])

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	table, _ := pterm.DefaultTable.WithHasHeader(true).WithData([][]string{headers, row}).Srender()
	checkGolden(t, "css_property_row", []byte(stripANSI(table)))
}

func TestCssPartToTableRowGolden(t *testing.T) {
	pkg := loadTestFixture(t, "custom_element_css_parts_properties_states.json")
	ced := pkg.Modules[0].Declarations[0].(*manifest.CustomElementDeclaration)
	renderable := manifest.NewRenderableCssPart(&ced.CssParts[0], ced, nil, &pkg.Modules[0])

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	table, _ := pterm.DefaultTable.WithHasHeader(true).WithData([][]string{headers, row}).Srender()
	checkGolden(t, "css_part_row", []byte(stripANSI(table)))
}

func TestClassMethodToTableRowGolden(t *testing.T) {
	pkg := loadTestFixture(t, "class.json")
	cd := pkg.Modules[0].Declarations[0].(*manifest.ClassDeclaration)
	method := &manifest.ClassMethod{
		FullyQualified: manifest.FullyQualified{
			Name:    "myMethod",
			Summary: "A test method.",
		},
	}
	renderable := manifest.NewRenderableClassMethod(method, cd, nil, &pkg.Modules[0], pkg)

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	table, _ := pterm.DefaultTable.WithHasHeader(true).WithData([][]string{headers, row}).Srender()
	checkGolden(t, "class_method_row", []byte(stripANSI(table)))
}
