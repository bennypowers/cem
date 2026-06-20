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
	testutil.LoadJSONFixture(t, "custom-element-attributes.json", &pkg)

	ced := pkg.Modules[0].Declarations[0].(*manifest.CustomElementDeclaration)
	attrs := ced.Attributes()
	renderable := manifest.NewRenderableAttribute(&attrs[0], ced, nil, &pkg.Modules[0])

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	rendered := list.FormatMarkdownTable(headers, [][]string{row})
	testutil.CheckGolden(t, "attribute-row", []byte(rendered), testutil.GoldenOptions{
		Dir:       "testdata",
		Extension: ".md",
	})
}

func TestSlotToTableRowGolden(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "custom-element-slots.json", &pkg)

	ced := pkg.Modules[0].Declarations[0].(*manifest.CustomElementDeclaration)
	slots := ced.Slots()
	renderable := manifest.NewRenderableSlot(&slots[0], ced, nil, &pkg.Modules[0])

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	rendered := list.FormatMarkdownTable(headers, [][]string{row})
	testutil.CheckGolden(t, "slot-row", []byte(rendered), testutil.GoldenOptions{
		Dir:       "testdata",
		Extension: ".md",
	})
}

func TestEventToTableRowGolden(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "custom-element-events.json", &pkg)

	ced := pkg.Modules[0].Declarations[0].(*manifest.CustomElementDeclaration)
	events := ced.Events()
	renderable := manifest.NewRenderableEvent(&events[0], ced, nil, &pkg.Modules[0])

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	rendered := list.FormatMarkdownTable(headers, [][]string{row})
	testutil.CheckGolden(t, "event-row", []byte(rendered), testutil.GoldenOptions{
		Dir:       "testdata",
		Extension: ".md",
	})
}

func TestCssPropertyToTableRowGolden(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "custom-element-css-parts-properties-states.json", &pkg)

	ced := pkg.Modules[0].Declarations[0].(*manifest.CustomElementDeclaration)
	props := ced.CssProperties()
	renderable := manifest.NewRenderableCssCustomProperty(&props[0], ced, nil, &pkg.Modules[0])

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	rendered := list.FormatMarkdownTable(headers, [][]string{row})
	testutil.CheckGolden(t, "css-property-row", []byte(rendered), testutil.GoldenOptions{
		Dir:       "testdata",
		Extension: ".md",
	})
}

func TestCssPartToTableRowGolden(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "custom-element-css-parts-properties-states.json", &pkg)

	ced := pkg.Modules[0].Declarations[0].(*manifest.CustomElementDeclaration)
	parts := ced.CssParts()
	renderable := manifest.NewRenderableCssPart(&parts[0], ced, nil, &pkg.Modules[0])

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	rendered := list.FormatMarkdownTable(headers, [][]string{row})
	testutil.CheckGolden(t, "css-part-row", []byte(rendered), testutil.GoldenOptions{
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
	testutil.CheckGolden(t, "class-method-row", []byte(rendered), testutil.GoldenOptions{
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
	testutil.CheckGolden(t, "class-field-row", []byte(rendered), testutil.GoldenOptions{
		Dir:       "testdata",
		Extension: ".md",
	})
}

func TestCustomElementFieldToTableRowGolden(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "custom-element-member-with-attribute.json", &pkg)

	ced := pkg.Modules[0].Declarations[0].(*manifest.CustomElementDeclaration)
	field := ced.Members[0].(*manifest.CustomElementField)
	renderable := manifest.NewRenderableCustomElementField(field, ced, nil, nil, &pkg.Modules[0], &pkg)

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	rendered := list.FormatMarkdownTable(headers, [][]string{row})
	testutil.CheckGolden(t, "custom-element-field-row", []byte(rendered), testutil.GoldenOptions{
		Dir:       "testdata",
		Extension: ".md",
	})
}

func TestCssCustomStateToTableRowGolden(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "custom-element-css-parts-properties-states.json", &pkg)

	ced := pkg.Modules[0].Declarations[0].(*manifest.CustomElementDeclaration)
	states := ced.CssStates()
	renderable := manifest.NewRenderableCssCustomState(&states[0], ced, nil, &pkg.Modules[0])

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	rendered := list.FormatMarkdownTable(headers, [][]string{row})
	testutil.CheckGolden(t, "css-state-row", []byte(rendered), testutil.GoldenOptions{
		Dir:       "testdata",
		Extension: ".md",
	})
}

func TestDemoToTableRowGolden(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "custom-element-demos.json", &pkg)

	ced := pkg.Modules[0].Declarations[0].(*manifest.CustomElementDeclaration)
	renderable := manifest.NewRenderableDemo(&ced.Demos[0], ced, &pkg.Modules[0], &pkg)

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	rendered := list.FormatMarkdownTable(headers, [][]string{row})
	testutil.CheckGolden(t, "demo-row", []byte(rendered), testutil.GoldenOptions{
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
	testutil.CheckGolden(t, "function-declaration-row", []byte(rendered), testutil.GoldenOptions{
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
	testutil.CheckGolden(t, "mixin-declaration-row", []byte(rendered), testutil.GoldenOptions{
		Dir:       "testdata",
		Extension: ".md",
	})
}

func TestCustomElementMixinDeclarationToTableRowGolden(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "custom-element-mixin.json", &pkg)

	cemd := pkg.Modules[0].Declarations[0].(*manifest.CustomElementMixinDeclaration)
	renderable := manifest.NewRenderableCustomElementMixinDeclaration(cemd, &pkg.Modules[0], &pkg)

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	rendered := list.FormatMarkdownTable(headers, [][]string{row})
	testutil.CheckGolden(t, "custom-element-mixin-declaration-row", []byte(rendered), testutil.GoldenOptions{
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
	testutil.CheckGolden(t, "variable-declaration-row", []byte(rendered), testutil.GoldenOptions{
		Dir:       "testdata",
		Extension: ".md",
	})
}

func TestModuleToTableRowGolden(t *testing.T) {
	var pkg manifest.Package
	testutil.LoadJSONFixture(t, "custom-element-members.json", &pkg)

	renderable := manifest.NewRenderableModule(&pkg.Modules[0], &pkg)

	headers := renderable.ColumnHeadings()
	row := renderable.ToTableRow()

	rendered := list.FormatMarkdownTable(headers, [][]string{row})
	testutil.CheckGolden(t, "module-row", []byte(rendered), testutil.GoldenOptions{
		Dir:       "testdata",
		Extension: ".md",
	})
}
