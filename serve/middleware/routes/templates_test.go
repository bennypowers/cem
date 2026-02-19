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

package routes

import (
	"slices"
	"strings"
	"testing"

	M "bennypowers.dev/cem/manifest"
)

func TestDeclarationAsKind(t *testing.T) {
	t.Run("class declaration", func(t *testing.T) {
		class := &M.ClassDeclaration{
			ClassLike: M.ClassLike{
				FullyQualified: M.FullyQualified{Name: "TestClass"},
			},
		}

		result := declarationAsKind[*M.ClassDeclaration](class)
		if result == nil {
			t.Fatal("Expected non-nil result for matching type")
		}
		if result.Name() != "TestClass" {
			t.Errorf("Expected Name='TestClass', got %q", result.Name())
		}
	})

	t.Run("custom element declaration", func(t *testing.T) {
		ce := &M.CustomElementDeclaration{
			ClassDeclaration: M.ClassDeclaration{
				ClassLike: M.ClassLike{
					FullyQualified: M.FullyQualified{Name: "MyElement"},
				},
			},
			CustomElement: M.CustomElement{
				TagName: "my-element",
			},
		}

		result := declarationAsKind[*M.CustomElementDeclaration](ce)
		if result == nil {
			t.Fatal("Expected non-nil result for matching type")
		} else if result.TagName != "my-element" {
			t.Errorf("Expected TagName='my-element', got %q", result.TagName)
		}
	})

	t.Run("function declaration", func(t *testing.T) {
		fn := &M.FunctionDeclaration{
			FullyQualified: M.FullyQualified{Name: "myFunction"},
		}

		result := declarationAsKind[*M.FunctionDeclaration](fn)
		if result == nil {
			t.Fatal("Expected non-nil result for matching type")
		} else if result.Name() != "myFunction" {
			t.Errorf("Expected Name='myFunction', got %q", result.Name())
		}
	})

	t.Run("variable declaration", func(t *testing.T) {
		v := &M.VariableDeclaration{
			PropertyLike: M.PropertyLike{
				FullyQualified: M.FullyQualified{Name: "myVar"},
			},
		}

		result := declarationAsKind[*M.VariableDeclaration](v)
		if result == nil {
			t.Fatal("Expected non-nil result for matching type")
		} else if result.Name() != "myVar" {
			t.Errorf("Expected Name='myVar', got %q", result.Name())
		}
	})

	t.Run("mixin declaration", func(t *testing.T) {
		mixin := &M.MixinDeclaration{
			FullyQualified: M.FullyQualified{Name: "MyMixin"},
		}

		result := declarationAsKind[*M.MixinDeclaration](mixin)
		if result == nil {
			t.Fatal("Expected non-nil result for matching type")
		} else if result.Name() != "MyMixin" {
			t.Errorf("Expected Name='MyMixin', got %q", result.Name())
		}
	})

	t.Run("wrong type returns nil", func(t *testing.T) {
		class := &M.ClassDeclaration{
			ClassLike: M.ClassLike{
				FullyQualified: M.FullyQualified{Name: "TestClass"},
			},
		}

		// Try to cast class to function - should return nil
		result := declarationAsKind[*M.FunctionDeclaration](class)
		if result != nil {
			t.Error("Expected nil for mismatched type")
		}
	})

	t.Run("nil input returns nil", func(t *testing.T) {
		result := declarationAsKind[*M.ClassDeclaration](nil)
		if result != nil {
			t.Error("Expected nil for nil input")
		}
	})
}

func TestTemplateStdlibFuncs(t *testing.T) {
	t.Run("string contains", func(t *testing.T) {
		// strings.Contains is used directly in templates
		if !strings.Contains("hello world", "world") {
			t.Error("Expected 'hello world' to contain 'world'")
		}
		if strings.Contains("hello world", "foo") {
			t.Error("Expected 'hello world' to not contain 'foo'")
		}
	})

	t.Run("slice elem", func(t *testing.T) {
		// slices.Contains is used as 'elem' in templates
		slice := []string{"foo", "bar", "baz"}
		if !slices.Contains(slice, "bar") {
			t.Error("Expected slice to contain 'bar'")
		}
		if slices.Contains(slice, "qux") {
			t.Error("Expected slice to not contain 'qux'")
		}
	})
}

func TestHasMethodMembers(t *testing.T) {
	t.Run("has methods", func(t *testing.T) {
		members := []M.ClassMember{
			&M.ClassField{
				PropertyLike: M.PropertyLike{
					FullyQualified: M.FullyQualified{Name: "field1"},
				},
			},
			&M.ClassMethod{
				FullyQualified: M.FullyQualified{Name: "method1"},
			},
		}
		if !hasMethodMembers(members) {
			t.Error("Expected to find methods in members")
		}
	})

	t.Run("no methods", func(t *testing.T) {
		members := []M.ClassMember{
			&M.ClassField{
				PropertyLike: M.PropertyLike{
					FullyQualified: M.FullyQualified{Name: "field1"},
				},
			},
			&M.ClassField{
				PropertyLike: M.PropertyLike{
					FullyQualified: M.FullyQualified{Name: "field2"},
				},
			},
		}
		if hasMethodMembers(members) {
			t.Error("Expected to not find methods in members")
		}
	})

	t.Run("empty members", func(t *testing.T) {
		if hasMethodMembers([]M.ClassMember{}) {
			t.Error("Expected no methods in empty slice")
		}
	})
}
