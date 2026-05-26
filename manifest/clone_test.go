/*
Copyright © 2026 Benny Powers

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
package manifest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Inline: verifying deep-copy independence

func TestFullyQualifiedClone(t *testing.T) {
	orig := FullyQualified{Name: "x", Summary: "s", Description: "d"}
	cloned := orig.Clone()
	orig.Name = "changed"
	assert.Equal(t, "x", cloned.Name)
	assert.Equal(t, "s", cloned.Summary)
	assert.Equal(t, "d", cloned.Description)
}

func TestReferenceClone(t *testing.T) {
	orig := Reference{Name: "Foo", Package: "pkg", Module: "mod.js"}
	cloned := orig.Clone()
	orig.Name = "changed"
	assert.Equal(t, "Foo", cloned.Name)
	assert.Equal(t, "pkg", cloned.Package)
}

func TestSourceReferenceClone(t *testing.T) {
	orig := SourceReference{Href: "http://example.com"}
	cloned := orig.Clone()
	orig.Href = "changed"
	assert.Equal(t, "http://example.com", cloned.Href)
}

func TestTypeClone(t *testing.T) {
	orig := &Type{
		Text: "string",
		References: []TypeReference{
			{Reference: Reference{Name: "Ref1"}, Start: 0, End: 6},
		},
		Source: &SourceReference{Href: "http://src"},
	}
	cloned := orig.Clone()

	orig.Text = "changed"
	orig.References[0].Name = "changed"
	orig.Source.Href = "changed"

	assert.Equal(t, "string", cloned.Text)
	assert.Equal(t, "Ref1", cloned.References[0].Name)
	assert.Equal(t, "http://src", cloned.Source.Href)
}

func TestTypeCloneNil(t *testing.T) {
	var orig *Type
	assert.Nil(t, orig.Clone())
}

func TestTypeReferenceClone(t *testing.T) {
	orig := TypeReference{Reference: Reference{Name: "R"}, Start: 1, End: 5}
	cloned := orig.Clone()
	orig.Name = "changed"
	assert.Equal(t, "R", cloned.Name)
	assert.Equal(t, 1, cloned.Start)
}

func TestPropertyLikeClone(t *testing.T) {
	orig := PropertyLike{
		FullyQualified: FullyQualified{Name: "prop"},
		Type:           &Type{Text: "string"},
		Default:        "hello",
		Deprecated:     DeprecatedFlag(true),
	}
	cloned := orig.Clone()
	orig.Name = "changed"
	orig.Type.Text = "changed"
	assert.Equal(t, "prop", cloned.Name)
	assert.Equal(t, "string", cloned.Type.Text)
}

func TestParameterClone(t *testing.T) {
	orig := Parameter{
		PropertyLike: PropertyLike{
			FullyQualified: FullyQualified{Name: "arg"},
			Type:           &Type{Text: "number"},
		},
		Optional: true,
		Rest:     false,
	}
	cloned := orig.Clone()
	orig.Name = "changed"
	assert.Equal(t, "arg", cloned.Name)
	assert.True(t, cloned.Optional)
}

func TestReturnClone(t *testing.T) {
	orig := Return{
		Type:        &Type{Text: "void"},
		Description: "returns nothing",
	}
	cloned := orig.Clone()
	orig.Type.Text = "changed"
	assert.Equal(t, "void", cloned.Type.Text)
	assert.Equal(t, "returns nothing", cloned.Description)
}

func TestFunctionLikeClone(t *testing.T) {
	orig := FunctionLike{
		Parameters: []Parameter{{PropertyLike: PropertyLike{FullyQualified: FullyQualified{Name: "a"}}}},
		Return:     &Return{Type: &Type{Text: "string"}},
	}
	cloned := orig.Clone()
	orig.Parameters[0].Name = "changed"
	orig.Return.Type.Text = "changed"
	assert.Equal(t, "a", cloned.Parameters[0].Name)
	assert.Equal(t, "string", cloned.Return.Type.Text)
}

func TestAttributeClone(t *testing.T) {
	orig := Attribute{
		FullyQualified: FullyQualified{Name: "disabled"},
		Type:           &Type{Text: "boolean"},
		Default:        "false",
		FieldName:      "disabled",
		InheritedFrom:  &Reference{Name: "Base"},
		Deprecated:     DeprecatedReason("use new-attr"),
	}
	cloned := orig.Clone()
	orig.Name = "changed"
	orig.Type.Text = "changed"
	orig.InheritedFrom.Name = "changed"
	assert.Equal(t, "disabled", cloned.Name)
	assert.Equal(t, "boolean", cloned.Type.Text)
	assert.Equal(t, "Base", cloned.InheritedFrom.Name)
	assert.Equal(t, DeprecatedReason("use new-attr"), cloned.Deprecated)
}

func TestSlotClone(t *testing.T) {
	orig := Slot{
		FullyQualified: FullyQualified{Name: "default"},
		InheritedFrom:  &Reference{Name: "Base"},
		Deprecated:     DeprecatedFlag(true),
	}
	cloned := orig.Clone()
	orig.Name = "changed"
	orig.InheritedFrom.Name = "changed"
	assert.Equal(t, "default", cloned.Name)
	assert.Equal(t, "Base", cloned.InheritedFrom.Name)
}

func TestEventClone(t *testing.T) {
	orig := Event{
		FullyQualified: FullyQualified{Name: "click"},
		Type:           &Type{Text: "CustomEvent"},
		InheritedFrom:  &Reference{Name: "Base"},
		Deprecated:     DeprecatedReason("old"),
	}
	cloned := orig.Clone()
	orig.Name = "changed"
	orig.Type.Text = "changed"
	orig.InheritedFrom.Name = "changed"
	assert.Equal(t, "click", cloned.Name)
	assert.Equal(t, "CustomEvent", cloned.Type.Text)
	assert.Equal(t, "Base", cloned.InheritedFrom.Name)
}

func TestDemoClone(t *testing.T) {
	orig := Demo{
		Description: "demo",
		URL:         "http://example.com",
		Source:      &SourceReference{Href: "src"},
	}
	cloned := orig.Clone()
	orig.Description = "changed"
	orig.Source.Href = "changed"
	assert.Equal(t, "demo", cloned.Description)
	assert.Equal(t, "src", cloned.Source.Href)
}

func TestCssCustomPropertyClone(t *testing.T) {
	orig := CssCustomProperty{
		FullyQualified: FullyQualified{Name: "--color"},
		Default:        "red",
		Syntax:         "<color>",
		InheritedFrom:  &Reference{Name: "Base"},
		Deprecated:     DeprecatedFlag(true),
	}
	cloned := orig.Clone()
	orig.Name = "changed"
	orig.InheritedFrom.Name = "changed"
	assert.Equal(t, "--color", cloned.Name)
	assert.Equal(t, "Base", cloned.InheritedFrom.Name)
	assert.Equal(t, "red", cloned.Default)
}

func TestCssCustomStateClone(t *testing.T) {
	orig := CssCustomState{
		FullyQualified: FullyQualified{Name: ":--active"},
		InheritedFrom:  &Reference{Name: "Base"},
		Deprecated:     DeprecatedReason("renamed"),
	}
	cloned := orig.Clone()
	orig.Name = "changed"
	orig.InheritedFrom.Name = "changed"
	assert.Equal(t, ":--active", cloned.Name)
	assert.Equal(t, "Base", cloned.InheritedFrom.Name)
}

func TestCssPartClone(t *testing.T) {
	orig := CssPart{
		FullyQualified: FullyQualified{Name: "button"},
		InheritedFrom:  &Reference{Name: "Base"},
		Deprecated:     DeprecatedFlag(true),
	}
	cloned := orig.Clone()
	orig.Name = "changed"
	orig.InheritedFrom.Name = "changed"
	assert.Equal(t, "button", cloned.Name)
	assert.Equal(t, "Base", cloned.InheritedFrom.Name)
}

func TestClassFieldClone(t *testing.T) {
	orig := &ClassField{
		PropertyLike: PropertyLike{
			FullyQualified: FullyQualified{Name: "value"},
			Type:           &Type{Text: "string"},
		},
		Kind:          "field",
		Static:        true,
		Privacy:       "public",
		InheritedFrom: &Reference{Name: "Base"},
		Source:        &SourceReference{Href: "src.ts"},
	}
	cloned := orig.Clone().(*ClassField)
	orig.Name = "changed"
	orig.Type.Text = "changed"
	orig.InheritedFrom.Name = "changed"
	assert.Equal(t, "value", cloned.Name)
	assert.Equal(t, "string", cloned.Type.Text)
	assert.Equal(t, "Base", cloned.InheritedFrom.Name)
	assert.True(t, cloned.Static)
}

func TestCustomElementFieldClone(t *testing.T) {
	orig := &CustomElementField{
		ClassField: ClassField{
			PropertyLike: PropertyLike{
				FullyQualified: FullyQualified{Name: "value"},
			},
			Kind: "field",
		},
		Attribute: "value",
		Reflects:  true,
	}
	cloned := orig.Clone().(*CustomElementField)
	orig.Name = "changed"
	assert.Equal(t, "value", cloned.Name)
	assert.Equal(t, "value", cloned.Attribute)
	assert.True(t, cloned.Reflects)
}

func TestClassMethodClone(t *testing.T) {
	orig := &ClassMethod{
		FunctionLike: FunctionLike{
			Parameters: []Parameter{
				{PropertyLike: PropertyLike{FullyQualified: FullyQualified{Name: "x"}}},
			},
		},
		FullyQualified: FullyQualified{Name: "doThing"},
		Kind:           "method",
		Static:         false,
		Privacy:        "public",
		InheritedFrom:  &Reference{Name: "Base"},
		Source:         &SourceReference{Href: "src.ts"},
		Deprecated:     DeprecatedReason("old"),
	}
	cloned := orig.Clone().(*ClassMethod)
	orig.FunctionLike.Parameters[0].Name = "changed"
	orig.Name = "changed"
	assert.Equal(t, "doThing", cloned.Name)
	assert.Equal(t, "x", cloned.Parameters[0].Name)
	assert.Equal(t, "Base", cloned.InheritedFrom.Name)
}

func TestJavaScriptExportClone(t *testing.T) {
	orig := &JavaScriptExport{
		Kind:        "js",
		Name:        "MyClass",
		Declaration: &Reference{Name: "MyClass", Module: "mod.js"},
		Deprecated:  DeprecatedFlag(true),
	}
	cloned := orig.Clone().(*JavaScriptExport)
	orig.Name = "changed"
	orig.Declaration.Name = "changed"
	assert.Equal(t, "MyClass", cloned.Name)
	assert.Equal(t, "MyClass", cloned.Declaration.Name)
}

func TestCustomElementExportClone(t *testing.T) {
	orig := &CustomElementExport{
		Kind:        "custom-element-definition",
		Name:        "my-el",
		Declaration: &Reference{Name: "MyEl"},
		Deprecated:  DeprecatedReason("old"),
	}
	cloned := orig.Clone().(*CustomElementExport)
	orig.Name = "changed"
	orig.Declaration.Name = "changed"
	assert.Equal(t, "my-el", cloned.Name)
	assert.Equal(t, "MyEl", cloned.Declaration.Name)
}

func TestFunctionDeclarationClone(t *testing.T) {
	orig := &FunctionDeclaration{
		FunctionLike: FunctionLike{
			Parameters: []Parameter{
				{PropertyLike: PropertyLike{FullyQualified: FullyQualified{Name: "a"}}},
			},
			Return: &Return{Type: &Type{Text: "void"}},
		},
		FullyQualified: FullyQualified{Name: "doStuff"},
		Kind:           "function",
		Deprecated:     DeprecatedFlag(true),
	}
	cloned := orig.Clone().(*FunctionDeclaration)
	orig.FullyQualified.Name = "changed"
	orig.FunctionLike.Parameters[0].Name = "changed"
	assert.Equal(t, "doStuff", cloned.FullyQualified.Name)
	assert.Equal(t, "a", cloned.Parameters[0].Name)
}

func TestVariableDeclarationClone(t *testing.T) {
	orig := &VariableDeclaration{
		PropertyLike: PropertyLike{
			FullyQualified: FullyQualified{Name: "myVar"},
			Type:           &Type{Text: "string"},
		},
		Kind:   "variable",
		Source: &SourceReference{Href: "src.ts"},
	}
	cloned := orig.Clone().(*VariableDeclaration)
	orig.FullyQualified.Name = "changed"
	orig.Source.Href = "changed"
	assert.Equal(t, "myVar", cloned.FullyQualified.Name)
	assert.Equal(t, "src.ts", cloned.Source.Href)
}

func TestDeprecatedClone(t *testing.T) {
	t.Run("flag", func(t *testing.T) {
		orig := DeprecatedFlag(true)
		cloned := orig.Clone()
		assert.Equal(t, orig, cloned)
	})
	t.Run("reason", func(t *testing.T) {
		orig := DeprecatedReason("use new thing")
		cloned := orig.Clone()
		assert.Equal(t, orig, cloned)
	})
}
