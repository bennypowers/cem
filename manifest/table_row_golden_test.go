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
	"testing"

	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/manifest"
	"github.com/pterm/pterm"
)

func TestAttributeToTableRowGolden(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "custom_element_attributes.json", &pkg)

	ced := pkg.Modules[0].Declarations[0].(*manifest.CustomElementDeclaration)
	attrs := ced.Attributes()
	renderable := manifest.NewRenderableAttribute(&attrs[0], ced, nil, &pkg.Modules[0])

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	table, _ := pterm.DefaultTable.WithHasHeader(true).WithData([][]string{headers, row}).Srender()
	testutil.CheckGolden(t, "attribute_row", []byte(table), testutil.GoldenOptions{
		Dir:       "fixtures",
		Extension: ".md",
		StripANSI: true,
	})
}

func TestSlotToTableRowGolden(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "custom_element_slots.json", &pkg)

	ced := pkg.Modules[0].Declarations[0].(*manifest.CustomElementDeclaration)
	slots := ced.Slots()
	renderable := manifest.NewRenderableSlot(&slots[0], ced, nil, &pkg.Modules[0])

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	table, _ := pterm.DefaultTable.WithHasHeader(true).WithData([][]string{headers, row}).Srender()
	testutil.CheckGolden(t, "slot_row", []byte(table), testutil.GoldenOptions{
		Dir:       "fixtures",
		Extension: ".md",
		StripANSI: true,
	})
}

func TestEventToTableRowGolden(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "custom_element_events.json", &pkg)

	ced := pkg.Modules[0].Declarations[0].(*manifest.CustomElementDeclaration)
	events := ced.Events()
	renderable := manifest.NewRenderableEvent(&events[0], ced, nil, &pkg.Modules[0])

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	table, _ := pterm.DefaultTable.WithHasHeader(true).WithData([][]string{headers, row}).Srender()
	testutil.CheckGolden(t, "event_row", []byte(table), testutil.GoldenOptions{
		Dir:       "fixtures",
		Extension: ".md",
		StripANSI: true,
	})
}

func TestCssPropertyToTableRowGolden(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "custom_element_css_parts_properties_states.json", &pkg)

	ced := pkg.Modules[0].Declarations[0].(*manifest.CustomElementDeclaration)
	props := ced.CssProperties()
	renderable := manifest.NewRenderableCssCustomProperty(&props[0], ced, nil, &pkg.Modules[0])

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	table, _ := pterm.DefaultTable.WithHasHeader(true).WithData([][]string{headers, row}).Srender()
	testutil.CheckGolden(t, "css_property_row", []byte(table), testutil.GoldenOptions{
		Dir:       "fixtures",
		Extension: ".md",
		StripANSI: true,
	})
}

func TestCssPartToTableRowGolden(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "custom_element_css_parts_properties_states.json", &pkg)

	ced := pkg.Modules[0].Declarations[0].(*manifest.CustomElementDeclaration)
	parts := ced.CssParts()
	renderable := manifest.NewRenderableCssPart(&parts[0], ced, nil, &pkg.Modules[0])

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	table, _ := pterm.DefaultTable.WithHasHeader(true).WithData([][]string{headers, row}).Srender()
	testutil.CheckGolden(t, "css_part_row", []byte(table), testutil.GoldenOptions{
		Dir:       "fixtures",
		Extension: ".md",
		StripANSI: true,
	})
}

func TestClassMethodToTableRowGolden(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "class.json", &pkg)

	cd := pkg.Modules[0].Declarations[0].(*manifest.ClassDeclaration)
	method := &manifest.ClassMethod{
		FullyQualified: manifest.FullyQualified{
			Name:    "myMethod",
			Summary: "A test method.",
		},
	}
	renderable := manifest.NewRenderableClassMethod(method, cd, nil, &pkg.Modules[0], &pkg)

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	table, _ := pterm.DefaultTable.WithHasHeader(true).WithData([][]string{headers, row}).Srender()
	testutil.CheckGolden(t, "class_method_row", []byte(table), testutil.GoldenOptions{
		Dir:       "fixtures",
		Extension: ".md",
		StripANSI: true,
	})
}
