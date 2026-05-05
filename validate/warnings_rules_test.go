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
package validate

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeCtx(members []map[string]any) *WarningContext {
	memberSlice := make([]any, len(members))
	for i, m := range members {
		memberSlice[i] = m
	}
	return &WarningContext{
		Declaration: RawDeclaration{"members": memberSlice},
		DeclName:    "TestClass",
		DeclKind:    "class",
		ModulePath:  "src/test.js",
	}
}

func TestPrivateMethodsRule_HashMethods(t *testing.T) {
	rule := &PrivateMethodsRule{}
	ctx := makeCtx([]map[string]any{
		{"name": "#private", "kind": "method"},
		{"name": "#field", "kind": "field"},
	})
	warnings := rule.Check(ctx)
	assert.Len(t, warnings, 2)
	assert.Equal(t, "private-hash-methods", warnings[0].ID)
	assert.Equal(t, "private-hash-fields", warnings[1].ID)
}

func TestPrivateMethodsRule_UnderscoreNotMarked(t *testing.T) {
	rule := &PrivateMethodsRule{}
	ctx := makeCtx([]map[string]any{
		{"name": "_internal", "kind": "method"},
		{"name": "_field", "kind": "field"},
	})
	warnings := rule.Check(ctx)
	assert.Len(t, warnings, 2)
	assert.Equal(t, "private-underscore-methods", warnings[0].ID)
	assert.Equal(t, "private-underscore-fields", warnings[1].ID)
}

func TestPrivateMethodsRule_UnderscoreMarkedPrivate(t *testing.T) {
	rule := &PrivateMethodsRule{}
	ctx := makeCtx([]map[string]any{
		{"name": "_internal", "kind": "method", "privacy": "private"},
		{"name": "_protected", "kind": "field", "privacy": "protected"},
	})
	warnings := rule.Check(ctx)
	assert.Empty(t, warnings)
}

func TestPrivateMethodsRule_NoMembers(t *testing.T) {
	rule := &PrivateMethodsRule{}
	ctx := &WarningContext{Declaration: RawDeclaration{}}
	assert.Empty(t, rule.Check(ctx))
}

func TestInternalMethodsRule(t *testing.T) {
	rule := &InternalMethodsRule{}

	t.Run("internal method names", func(t *testing.T) {
		ctx := makeCtx([]map[string]any{
			{"name": "init", "kind": "method"},
			{"name": "destroy", "kind": "method"},
		})
		warnings := rule.Check(ctx)
		assert.Len(t, warnings, 2)
	})

	t.Run("non-method member skipped", func(t *testing.T) {
		ctx := makeCtx([]map[string]any{
			{"name": "init", "kind": "field"},
		})
		warnings := rule.Check(ctx)
		assert.Empty(t, warnings)
	})

	t.Run("regular method name", func(t *testing.T) {
		ctx := makeCtx([]map[string]any{
			{"name": "handleClick", "kind": "method"},
		})
		assert.Empty(t, rule.Check(ctx))
	})
}

func TestSuperclassRule(t *testing.T) {
	rule := &SuperclassRule{}

	t.Run("builtin without global package", func(t *testing.T) {
		ctx := &WarningContext{
			Declaration: RawDeclaration{
				"superclass": map[string]any{"name": "HTMLElement"},
			},
			DeclName: "MyEl", DeclKind: "class", ModulePath: "test.js",
		}
		warnings := rule.Check(ctx)
		assert.Len(t, warnings, 1)
		assert.Contains(t, warnings[0].Message, "missing package field")
	})

	t.Run("builtin with wrong package", func(t *testing.T) {
		ctx := &WarningContext{
			Declaration: RawDeclaration{
				"superclass": map[string]any{"name": "HTMLElement", "package": "wrong"},
			},
			DeclName: "MyEl", DeclKind: "class", ModulePath: "test.js",
		}
		warnings := rule.Check(ctx)
		assert.Len(t, warnings, 1)
		assert.Contains(t, warnings[0].Message, `package is "wrong"`)
	})

	t.Run("builtin with global package", func(t *testing.T) {
		ctx := &WarningContext{
			Declaration: RawDeclaration{
				"superclass": map[string]any{"name": "HTMLElement", "package": "global:"},
			},
			DeclName: "MyEl", DeclKind: "class", ModulePath: "test.js",
		}
		assert.Empty(t, rule.Check(ctx))
	})

	t.Run("no superclass", func(t *testing.T) {
		ctx := &WarningContext{Declaration: RawDeclaration{}}
		assert.Empty(t, rule.Check(ctx))
	})
}

func TestImplementationDetailsRule(t *testing.T) {
	rule := &ImplementationDetailsRule{}

	t.Run("static implementation fields", func(t *testing.T) {
		ctx := makeCtx([]map[string]any{
			{"name": "styles", "kind": "field", "static": true},
			{"name": "observedAttributes", "kind": "field", "static": true},
		})
		warnings := rule.Check(ctx)
		assert.Len(t, warnings, 2)
	})

	t.Run("non-static field ignored", func(t *testing.T) {
		ctx := makeCtx([]map[string]any{
			{"name": "styles", "kind": "field", "static": false},
		})
		assert.Empty(t, rule.Check(ctx))
	})

	t.Run("method ignored", func(t *testing.T) {
		ctx := makeCtx([]map[string]any{
			{"name": "styles", "kind": "method", "static": true},
		})
		assert.Empty(t, rule.Check(ctx))
	})

	t.Run("regular static field ignored", func(t *testing.T) {
		ctx := makeCtx([]map[string]any{
			{"name": "VERSION", "kind": "field", "static": true},
		})
		assert.Empty(t, rule.Check(ctx))
	})
}

func TestLifecycleMethodsRule(t *testing.T) {
	rule := &LifecycleMethodsRule{}

	t.Run("web component lifecycle", func(t *testing.T) {
		ctx := makeCtx([]map[string]any{
			{"name": "connectedCallback", "kind": "method"},
			{"name": "disconnectedCallback", "kind": "method"},
		})
		warnings := rule.Check(ctx)
		assert.Len(t, warnings, 2)
		assert.Equal(t, "lifecycle-web-components", warnings[0].ID)
	})

	t.Run("lit lifecycle", func(t *testing.T) {
		ctx := makeCtx([]map[string]any{
			{"name": "firstUpdated", "kind": "method"},
			{"name": "updated", "kind": "method"},
		})
		warnings := rule.Check(ctx)
		assert.Len(t, warnings, 2)
		assert.Equal(t, "lifecycle-lit-methods", warnings[0].ID)
	})

	t.Run("render in lit element", func(t *testing.T) {
		ctx := makeCtx([]map[string]any{
			{"name": "render", "kind": "method"},
		})
		ctx.IsLitElement = true
		warnings := rule.Check(ctx)
		assert.Len(t, warnings, 1)
		assert.Equal(t, "lifecycle-lit-render", warnings[0].ID)
	})

	t.Run("render in non-lit element", func(t *testing.T) {
		ctx := makeCtx([]map[string]any{
			{"name": "render", "kind": "method"},
		})
		ctx.IsLitElement = false
		assert.Empty(t, rule.Check(ctx))
	})

	t.Run("form callbacks", func(t *testing.T) {
		ctx := makeCtx([]map[string]any{
			{"name": "formDisabledCallback", "kind": "method"},
		})
		warnings := rule.Check(ctx)
		assert.Len(t, warnings, 1)
		assert.Equal(t, "lifecycle-form-callbacks", warnings[0].ID)
	})

	t.Run("non-method members skipped", func(t *testing.T) {
		ctx := makeCtx([]map[string]any{
			{"name": "connectedCallback", "kind": "field"},
		})
		assert.Empty(t, rule.Check(ctx))
	})

	t.Run("regular method no warning", func(t *testing.T) {
		ctx := makeCtx([]map[string]any{
			{"name": "handleClick", "kind": "method"},
		})
		assert.Empty(t, rule.Check(ctx))
	})
}

func TestVerboseDefaultsRule(t *testing.T) {
	rule := &VerboseDefaultsRule{}

	t.Run("verbose member default", func(t *testing.T) {
		longDefault := strings.Repeat("x", DefaultLengthThreshold+1)
		ctx := makeCtx([]map[string]any{
			{"name": "data", "kind": "field", "default": longDefault},
		})
		warnings := rule.Check(ctx)
		assert.Len(t, warnings, 1)
	})

	t.Run("short member default ok", func(t *testing.T) {
		ctx := makeCtx([]map[string]any{
			{"name": "data", "kind": "field", "default": "short"},
		})
		assert.Empty(t, rule.Check(ctx))
	})

	t.Run("verbose CSS property default", func(t *testing.T) {
		longDefault := strings.Repeat("x", DefaultLengthThreshold+1)
		ctx := &WarningContext{
			Declaration: RawDeclaration{
				"cssProperties": []any{map[string]any{"name": "--color", "default": longDefault}},
			},
			DeclName: "MyEl", DeclKind: "class", ModulePath: "test.js",
		}
		warnings := rule.Check(ctx)
		assert.Len(t, warnings, 1)
	})
}
