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
package ephemeral_test

import (
	"slices"
	"testing"

	"bennypowers.dev/cem/lsp/ephemeral"
	M "bennypowers.dev/cem/manifest"
)

func makeTestPackage(tagName, className, summary string, attrs ...M.Attribute) *M.Package {
	pkg := &M.Package{
		SchemaVersion: "2.1.0",
		Modules: []M.Module{
			{
				Kind: "javascript-module",
				Path: "test.ts",
				Declarations: []M.Declaration{
					&M.CustomElementDeclaration{
						ClassDeclaration: M.ClassDeclaration{
							ClassLike: M.ClassLike{
								FullyQualified: M.FullyQualified{
									Name:    className,
									Summary: summary,
								},
							},
							Kind: "class",
						},
						CustomElement: M.CustomElement{
							TagName:       tagName,
							Attributes:    attrs,
							CustomElement: true,
						},
					},
				},
			},
		},
	}
	return pkg
}

func TestUpdate(t *testing.T) {
	r := ephemeral.NewRegistry()
	pkg := makeTestPackage("test-el", "TestEl", "A test element")

	r.Update("file:///test.ts", pkg)

	decl := r.FindCustomElementDeclaration("test-el")
	if decl == nil {
		t.Fatal("expected to find test-el declaration after Update")
	} else if decl.TagName != "test-el" {
		t.Errorf("expected tag name 'test-el', got %q", decl.TagName)
	}
	if decl.Summary != "A test element" {
		t.Errorf("expected summary 'A test element', got %q", decl.Summary)
	}
}

func TestRemove(t *testing.T) {
	r := ephemeral.NewRegistry()
	pkg := makeTestPackage("test-el", "TestEl", "A test element")

	r.Update("file:///test.ts", pkg)
	r.Remove("file:///test.ts")

	decl := r.FindCustomElementDeclaration("test-el")
	if decl != nil {
		t.Error("expected nil declaration after Remove")
	}

	el, ok := r.Element("test-el")
	if ok || el != nil {
		t.Error("expected Element to return false after Remove")
	}
}

func TestUpdateReplace(t *testing.T) {
	r := ephemeral.NewRegistry()

	// Add initial version
	pkg1 := makeTestPackage("test-el", "TestEl", "Version 1")
	r.Update("file:///test.ts", pkg1)

	// Replace with updated version
	pkg2 := makeTestPackage("test-el", "TestEl", "Version 2")
	r.Update("file:///test.ts", pkg2)

	decl := r.FindCustomElementDeclaration("test-el")
	if decl == nil {
		t.Fatal("expected to find test-el declaration after replacement")
	}
	if decl.Summary != "Version 2" {
		t.Errorf("expected summary 'Version 2', got %q", decl.Summary)
	}
}

func TestAllTagNames(t *testing.T) {
	r := ephemeral.NewRegistry()

	r.Update("file:///a.ts", makeTestPackage("el-a", "ElA", ""))
	r.Update("file:///b.ts", makeTestPackage("el-b", "ElB", ""))

	tags := r.AllTagNames()
	if len(tags) != 2 {
		t.Fatalf("expected 2 tag names, got %d", len(tags))
	}
	if !slices.Contains(tags, "el-a") {
		t.Error("expected tags to contain 'el-a'")
	}
	if !slices.Contains(tags, "el-b") {
		t.Error("expected tags to contain 'el-b'")
	}
}

func TestElement(t *testing.T) {
	r := ephemeral.NewRegistry()
	r.Update("file:///test.ts", makeTestPackage("test-el", "TestEl", ""))

	el, ok := r.Element("test-el")
	if !ok || el == nil {
		t.Fatal("expected to find element")
	}
	if el.TagName != "test-el" {
		t.Errorf("expected tag name 'test-el', got %q", el.TagName)
	}
}

func TestAttributes(t *testing.T) {
	r := ephemeral.NewRegistry()
	attrs := []M.Attribute{
		{FullyQualified: M.FullyQualified{Name: "variant"}, Type: &M.Type{Text: "string"}},
		{FullyQualified: M.FullyQualified{Name: "disabled"}, Type: &M.Type{Text: "boolean"}},
	}
	r.Update("file:///test.ts", makeTestPackage("test-el", "TestEl", "", attrs...))

	attrMap, ok := r.Attributes("test-el")
	if !ok || attrMap == nil {
		t.Fatal("expected to find attributes")
	}
	if len(attrMap) != 2 {
		t.Fatalf("expected 2 attributes, got %d", len(attrMap))
	}
	if attrMap["variant"] == nil {
		t.Error("expected to find 'variant' attribute")
	}
	if attrMap["disabled"] == nil {
		t.Error("expected to find 'disabled' attribute")
	}
}

func TestSlots(t *testing.T) {
	r := ephemeral.NewRegistry()
	pkg := makeTestPackage("test-el", "TestEl", "")
	// Add slots to the custom element
	ce := pkg.Modules[0].Declarations[0].(*M.CustomElementDeclaration)
	ce.CustomElement.Slots = []M.Slot{
		{FullyQualified: M.FullyQualified{Name: "", Description: "Default slot"}},
		{FullyQualified: M.FullyQualified{Name: "header", Description: "Header slot"}},
	}
	r.Update("file:///test.ts", pkg)

	slots, ok := r.Slots("test-el")
	if !ok || slots == nil {
		t.Fatal("expected to find slots")
	}
	if len(slots) != 2 {
		t.Fatalf("expected 2 slots, got %d", len(slots))
	}
}

func TestElementDefinition(t *testing.T) {
	r := ephemeral.NewRegistry()
	r.Update("file:///test.ts", makeTestPackage("test-el", "TestEl", ""))

	def, ok := r.ElementDefinition("test-el")
	if !ok || def == nil {
		t.Fatal("expected to find element definition")
	}
	if def.ModulePath() != "file:///test.ts" {
		t.Errorf("expected module path 'file:///test.ts', got %q", def.ModulePath())
	}
	if def.PackageName() != "" {
		t.Errorf("expected empty package name, got %q", def.PackageName())
	}
}

func TestElementDescription(t *testing.T) {
	t.Run("summary only", func(t *testing.T) {
		r := ephemeral.NewRegistry()
		r.Update("file:///test.ts", makeTestPackage("test-el", "TestEl", "A test element"))

		desc, ok := r.ElementDescription("test-el")
		if !ok {
			t.Fatal("expected to find element description")
		}
		if desc != "A test element" {
			t.Errorf("expected description 'A test element', got %q", desc)
		}
	})

	t.Run("description preferred over summary", func(t *testing.T) {
		r := ephemeral.NewRegistry()
		pkg := makeTestPackage("test-el", "TestEl", "Short summary")
		// Set Description on the declaration — should be preferred over Summary
		ce := pkg.Modules[0].Declarations[0].(*M.CustomElementDeclaration)
		ce.Description = "Full description of the element"
		r.Update("file:///test.ts", pkg)

		desc, ok := r.ElementDescription("test-el")
		if !ok {
			t.Fatal("expected to find element description")
		}
		if desc != "Full description of the element" {
			t.Errorf("expected 'Full description of the element', got %q", desc)
		}
	})
}

func TestHas(t *testing.T) {
	r := ephemeral.NewRegistry()
	r.Update("file:///test.ts", makeTestPackage("test-el", "TestEl", ""))

	if !r.Has("test-el") {
		t.Error("expected Has('test-el') to return true")
	}
	if r.Has("nonexistent") {
		t.Error("expected Has('nonexistent') to return false")
	}
}

func TestNotFound(t *testing.T) {
	r := ephemeral.NewRegistry()

	if decl := r.FindCustomElementDeclaration("nope"); decl != nil {
		t.Error("expected nil for nonexistent tag")
	}
	if _, ok := r.Element("nope"); ok {
		t.Error("expected false for nonexistent element")
	}
	if _, ok := r.Attributes("nope"); ok {
		t.Error("expected false for nonexistent attributes")
	}
	if _, ok := r.Slots("nope"); ok {
		t.Error("expected false for nonexistent slots")
	}
	if _, ok := r.ElementDefinition("nope"); ok {
		t.Error("expected false for nonexistent definition")
	}
}
