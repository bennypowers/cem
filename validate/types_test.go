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
package validate_test

import (
	"testing"

	"bennypowers.dev/cem/validate"
	"github.com/stretchr/testify/assert"
)

func TestRawManifest_Modules(t *testing.T) {
	t.Run("valid modules", func(t *testing.T) {
		m := validate.RawManifest{"modules": []any{map[string]any{"path": "a.js"}}}
		modules, ok := m.Modules()
		assert.True(t, ok)
		assert.Len(t, modules, 1)
		assert.Equal(t, "a.js", modules[0].Path())
	})
	t.Run("missing key", func(t *testing.T) {
		m := validate.RawManifest{}
		_, ok := m.Modules()
		assert.False(t, ok)
	})
	t.Run("wrong type", func(t *testing.T) {
		m := validate.RawManifest{"modules": "not an array"}
		_, ok := m.Modules()
		assert.False(t, ok)
	})
}

func TestRawModule_Path(t *testing.T) {
	assert.Equal(t, "src/button.js", validate.RawModule{"path": "src/button.js"}.Path())
	assert.Equal(t, "", validate.RawModule{}.Path())
	assert.Equal(t, "", validate.RawModule{"path": 42}.Path())
}

func TestRawModule_Declarations(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		m := validate.RawModule{"declarations": []any{map[string]any{"name": "Foo"}}}
		decls, ok := m.Declarations()
		assert.True(t, ok)
		assert.Len(t, decls, 1)
		assert.Equal(t, "Foo", decls[0].Name())
	})
	t.Run("missing", func(t *testing.T) {
		_, ok := validate.RawModule{}.Declarations()
		assert.False(t, ok)
	})
}

func TestRawDeclaration_StringGetters(t *testing.T) {
	d := validate.RawDeclaration{
		"name":        "MyElement",
		"kind":        "class",
		"description": "A cool element",
		"summary":     "Cool",
		"tagName":     "my-element",
	}
	assert.Equal(t, "MyElement", d.Name())
	assert.Equal(t, "class", d.Kind())
	assert.Equal(t, "A cool element", d.Description())
	assert.Equal(t, "Cool", d.Summary())
	assert.Equal(t, "my-element", d.TagName())

	empty := validate.RawDeclaration{}
	assert.Equal(t, "", empty.Name())
	assert.Equal(t, "", empty.Kind())
	assert.Equal(t, "", empty.Description())
	assert.Equal(t, "", empty.Summary())
	assert.Equal(t, "", empty.TagName())
}

func TestRawDeclaration_IsCustomElement(t *testing.T) {
	assert.True(t, validate.RawDeclaration{"customElement": true}.IsCustomElement())
	assert.False(t, validate.RawDeclaration{"customElement": false}.IsCustomElement())
	assert.True(t, validate.RawDeclaration{"tagName": "x-el"}.IsCustomElement())
	assert.False(t, validate.RawDeclaration{}.IsCustomElement())
}

func TestRawDeclaration_Demos(t *testing.T) {
	d := validate.RawDeclaration{"demos": []any{map[string]any{"url": "http://demo"}}}
	demos, ok := d.Demos()
	assert.True(t, ok)
	assert.Len(t, demos, 1)
	assert.Equal(t, "http://demo", demos[0].URL())

	_, ok = validate.RawDeclaration{}.Demos()
	assert.False(t, ok)
}

func TestRawDeclaration_Members(t *testing.T) {
	d := validate.RawDeclaration{"members": []any{map[string]any{"name": "foo", "kind": "field"}}}
	members, ok := d.Members()
	assert.True(t, ok)
	assert.Len(t, members, 1)
	assert.Equal(t, "foo", members[0].Name())

	_, ok = validate.RawDeclaration{}.Members()
	assert.False(t, ok)
}

func TestRawDeclaration_Superclass(t *testing.T) {
	d := validate.RawDeclaration{"superclass": map[string]any{"name": "LitElement"}}
	sc, ok := d.Superclass()
	assert.True(t, ok)
	assert.Equal(t, "LitElement", sc.Name())

	_, ok = validate.RawDeclaration{}.Superclass()
	assert.False(t, ok)
}

func TestRawDeclaration_PropertyArrayGetters(t *testing.T) {
	keys := []string{"cssProperties", "cssParts", "cssStates", "attributes", "events", "slots"}
	getters := map[string]func(validate.RawDeclaration) ([]validate.RawProperty, bool){
		"cssProperties": validate.RawDeclaration.CSSProperties,
		"cssParts":      validate.RawDeclaration.CSSParts,
		"cssStates":     validate.RawDeclaration.CSSStates,
		"attributes":    validate.RawDeclaration.Attributes,
		"events":        validate.RawDeclaration.Events,
		"slots":         validate.RawDeclaration.Slots,
	}

	for _, key := range keys {
		t.Run(key, func(t *testing.T) {
			d := validate.RawDeclaration{key: []any{map[string]any{"name": "test"}}}
			props, ok := getters[key](d)
			assert.True(t, ok)
			assert.Len(t, props, 1)
			assert.Equal(t, "test", props[0].Name())

			_, ok = getters[key](validate.RawDeclaration{})
			assert.False(t, ok)
		})
	}
}

func TestRawDemo_Getters(t *testing.T) {
	d := validate.RawDemo{"url": "http://demo", "description": "A demo"}
	assert.Equal(t, "http://demo", d.URL())
	assert.Equal(t, "A demo", d.Description())

	empty := validate.RawDemo{}
	assert.Equal(t, "", empty.URL())
	assert.Equal(t, "", empty.Description())
}

func TestRawMember_Getters(t *testing.T) {
	m := validate.RawMember{
		"name":        "value",
		"kind":        "field",
		"description": "The value",
		"type":        map[string]any{"text": "string"},
		"static":      true,
		"privacy":     "public",
		"default":     "''",
	}
	assert.Equal(t, "value", m.Name())
	assert.Equal(t, "field", m.Kind())
	assert.Equal(t, "The value", m.Description())
	assert.Equal(t, "string", m.Type())
	assert.True(t, m.IsStatic())
	assert.Equal(t, "public", m.Privacy())
	assert.Equal(t, "''", m.Default())

	empty := validate.RawMember{}
	assert.Equal(t, "", empty.Name())
	assert.Equal(t, "", empty.Kind())
	assert.Equal(t, "", empty.Description())
	assert.Equal(t, "", empty.Type())
	assert.False(t, empty.IsStatic())
	assert.Equal(t, "", empty.Privacy())
	assert.Equal(t, "", empty.Default())
}

func TestRawProperty_Getters(t *testing.T) {
	p := validate.RawProperty{
		"name":        "--color",
		"description": "Main color",
		"type":        map[string]any{"text": "<color>"},
		"default":     "red",
	}
	assert.Equal(t, "--color", p.Name())
	assert.Equal(t, "Main color", p.Description())
	assert.Equal(t, "<color>", p.Type())
	assert.Equal(t, "red", p.Default())

	empty := validate.RawProperty{}
	assert.Equal(t, "", empty.Name())
	assert.Equal(t, "", empty.Description())
	assert.Equal(t, "", empty.Type())
	assert.Equal(t, "", empty.Default())
}

func TestRawSuperclass_Getters(t *testing.T) {
	s := validate.RawSuperclass{"name": "LitElement", "package": "lit", "module": "lit-element.js"}
	assert.Equal(t, "LitElement", s.Name())
	assert.Equal(t, "lit", s.Package())
	assert.Equal(t, "lit-element.js", s.Module())

	empty := validate.RawSuperclass{}
	assert.Equal(t, "", empty.Name())
	assert.Equal(t, "", empty.Package())
	assert.Equal(t, "", empty.Module())
}

func TestParseContext(t *testing.T) {
	t.Run("module", func(t *testing.T) {
		ctx := validate.ParseContext("/modules/2")
		assert.Equal(t, 2, ctx.ModuleIndex)
		assert.Equal(t, "module", ctx.Type)
	})
	t.Run("declaration", func(t *testing.T) {
		ctx := validate.ParseContext("/modules/0/declarations/1")
		assert.Equal(t, 0, ctx.ModuleIndex)
		assert.Equal(t, 1, ctx.DeclIndex)
		assert.Equal(t, "declaration", ctx.Type)
	})
	t.Run("member", func(t *testing.T) {
		ctx := validate.ParseContext("/modules/0/declarations/0/members/3")
		assert.Equal(t, 3, ctx.MemberIndex)
		assert.Equal(t, "member", ctx.Type)
	})
	t.Run("property", func(t *testing.T) {
		ctx := validate.ParseContext("/modules/0/declarations/0/cssProperties/1")
		assert.Equal(t, 1, ctx.PropertyIndex)
		assert.Equal(t, "cssProperties", ctx.PropertyType)
		assert.Equal(t, "property", ctx.Type)
	})
	t.Run("empty", func(t *testing.T) {
		ctx := validate.ParseContext("")
		assert.Equal(t, -1, ctx.ModuleIndex)
		assert.Equal(t, "", ctx.Type)
	})
}
