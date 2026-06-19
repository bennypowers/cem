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
	"bennypowers.dev/cem/list"
	"bennypowers.dev/cem/manifest"
)

func TestAttributeToTableRowGolden(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "custom_element_attributes.json", &pkg)

	ced := pkg.Modules[0].Declarations[0].(*manifest.CustomElementDeclaration)
	attrs := ced.Attributes()
	renderable := manifest.NewRenderableAttribute(&attrs[0], ced, nil, &pkg.Modules[0])

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	rendered := list.FormatMarkdownTable(headers, [][]string{row})
	testutil.CheckGolden(t, "attribute_row", []byte(rendered), testutil.GoldenOptions{
		Dir:       "testdata",
		Extension: ".md",
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

	rendered := list.FormatMarkdownTable(headers, [][]string{row})
	testutil.CheckGolden(t, "slot_row", []byte(rendered), testutil.GoldenOptions{
		Dir:       "testdata",
		Extension: ".md",
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

	rendered := list.FormatMarkdownTable(headers, [][]string{row})
	testutil.CheckGolden(t, "event_row", []byte(rendered), testutil.GoldenOptions{
		Dir:       "testdata",
		Extension: ".md",
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

	rendered := list.FormatMarkdownTable(headers, [][]string{row})
	testutil.CheckGolden(t, "css_property_row", []byte(rendered), testutil.GoldenOptions{
		Dir:       "testdata",
		Extension: ".md",
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

	rendered := list.FormatMarkdownTable(headers, [][]string{row})
	testutil.CheckGolden(t, "css_part_row", []byte(rendered), testutil.GoldenOptions{
		Dir:       "testdata",
		Extension: ".md",
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

	rendered := list.FormatMarkdownTable(headers, [][]string{row})
	testutil.CheckGolden(t, "class_method_row", []byte(rendered), testutil.GoldenOptions{
		Dir:       "testdata",
		Extension: ".md",
	})
}

func TestClassFieldToTableRowGolden(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "class.json", &pkg)

	cd := pkg.Modules[0].Declarations[0].(*manifest.ClassDeclaration)
	field := cd.Members[0].(*manifest.ClassField)
	renderable := manifest.NewRenderableClassField(field, cd, nil, &pkg.Modules[0], &pkg)

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	rendered := list.FormatMarkdownTable(headers, [][]string{row})
	testutil.CheckGolden(t, "class_field_row", []byte(rendered), testutil.GoldenOptions{
		Dir:       "testdata",
		Extension: ".md",
	})
}

func TestCustomElementFieldToTableRowGolden(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "custom_element_member_with_attribute.json", &pkg)

	ced := pkg.Modules[0].Declarations[0].(*manifest.CustomElementDeclaration)
	field := ced.Members[0].(*manifest.CustomElementField)
	renderable := manifest.NewRenderableCustomElementField(field, ced, nil, nil, &pkg.Modules[0], &pkg)

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	rendered := list.FormatMarkdownTable(headers, [][]string{row})
	testutil.CheckGolden(t, "custom_element_field_row", []byte(rendered), testutil.GoldenOptions{
		Dir:       "testdata",
		Extension: ".md",
	})
}

func TestCssCustomStateToTableRowGolden(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "custom_element_css_parts_properties_states.json", &pkg)

	ced := pkg.Modules[0].Declarations[0].(*manifest.CustomElementDeclaration)
	states := ced.CssStates()
	renderable := manifest.NewRenderableCssCustomState(&states[0], ced, nil, &pkg.Modules[0])

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	rendered := list.FormatMarkdownTable(headers, [][]string{row})
	testutil.CheckGolden(t, "css_state_row", []byte(rendered), testutil.GoldenOptions{
		Dir:       "testdata",
		Extension: ".md",
	})
}

func TestDemoToTableRowGolden(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "custom_element_demos.json", &pkg)

	ced := pkg.Modules[0].Declarations[0].(*manifest.CustomElementDeclaration)
	renderable := manifest.NewRenderableDemo(&ced.Demos[0], ced, &pkg.Modules[0], &pkg)

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	rendered := list.FormatMarkdownTable(headers, [][]string{row})
	testutil.CheckGolden(t, "demo_row", []byte(rendered), testutil.GoldenOptions{
		Dir:       "testdata",
		Extension: ".md",
	})
}

func TestFunctionDeclarationToTableRowGolden(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "function.json", &pkg)

	fd := pkg.Modules[0].Declarations[0].(*manifest.FunctionDeclaration)
	renderable := manifest.NewRenderableFunctionDeclaration(fd, &pkg.Modules[0], &pkg)

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	rendered := list.FormatMarkdownTable(headers, [][]string{row})
	testutil.CheckGolden(t, "function_declaration_row", []byte(rendered), testutil.GoldenOptions{
		Dir:       "testdata",
		Extension: ".md",
	})
}

func TestMixinDeclarationToTableRowGolden(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "mixin.json", &pkg)

	md := pkg.Modules[0].Declarations[0].(*manifest.MixinDeclaration)
	renderable := manifest.NewRenderableMixinDeclaration(md, &pkg.Modules[0], &pkg)

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	rendered := list.FormatMarkdownTable(headers, [][]string{row})
	testutil.CheckGolden(t, "mixin_declaration_row", []byte(rendered), testutil.GoldenOptions{
		Dir:       "testdata",
		Extension: ".md",
	})
}

func TestCustomElementMixinDeclarationToTableRowGolden(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "custom_element_mixin.json", &pkg)

	cemd := pkg.Modules[0].Declarations[0].(*manifest.CustomElementMixinDeclaration)
	renderable := manifest.NewRenderableCustomElementMixinDeclaration(cemd, &pkg.Modules[0], &pkg)

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	rendered := list.FormatMarkdownTable(headers, [][]string{row})
	testutil.CheckGolden(t, "custom_element_mixin_declaration_row", []byte(rendered), testutil.GoldenOptions{
		Dir:       "testdata",
		Extension: ".md",
	})
}

func TestVariableDeclarationToTableRowGolden(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "variable.json", &pkg)

	vd := pkg.Modules[0].Declarations[0].(*manifest.VariableDeclaration)
	renderable := manifest.NewRenderableVariableDeclaration(vd, &pkg.Modules[0], &pkg)

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	rendered := list.FormatMarkdownTable(headers, [][]string{row})
	testutil.CheckGolden(t, "variable_declaration_row", []byte(rendered), testutil.GoldenOptions{
		Dir:       "testdata",
		Extension: ".md",
	})
}

func TestModuleToTableRowGolden(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "custom_element_members.json", &pkg)

	renderable := manifest.NewRenderableModule(&pkg.Modules[0], &pkg)

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	rendered := list.FormatMarkdownTable(headers, [][]string{row})
	testutil.CheckGolden(t, "module_row", []byte(rendered), testutil.GoldenOptions{
		Dir:       "testdata",
		Extension: ".md",
	})
}
