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
package health

import (
	"testing"

	"bennypowers.dev/cem/validate"
)

func TestAttributeDocRule(t *testing.T) {
	rule := &AttributeDocRule{}

	t.Run("full score with well-documented attributes", func(t *testing.T) {
		decl := validate.RawDeclaration{
			"name": "TestElement",
			"attributes": []any{
				map[string]any{
					"name":        "value",
					"description": "The current value of the input field for form submission",
					"type":        map[string]any{"text": "string"},
					"default":     "''",
				},
			},
		}
		ctx := &HealthContext{Declaration: decl}
		result := rule.Evaluate(ctx)
		if result.Points != rule.MaxPoints() {
			t.Errorf("expected %d points, got %d", rule.MaxPoints(), result.Points)
		}
	})

	t.Run("full score with no attributes (vacuous)", func(t *testing.T) {
		decl := validate.RawDeclaration{
			"name": "TestElement",
		}
		ctx := &HealthContext{Declaration: decl}
		result := rule.Evaluate(ctx)
		if result.Points != rule.MaxPoints() {
			t.Errorf("expected %d points (vacuous), got %d", rule.MaxPoints(), result.Points)
		}
	})

	t.Run("zero score with undocumented attributes", func(t *testing.T) {
		decl := validate.RawDeclaration{
			"name": "TestElement",
			"attributes": []any{
				map[string]any{"name": "foo"},
				map[string]any{"name": "bar"},
			},
		}
		ctx := &HealthContext{Declaration: decl}
		result := rule.Evaluate(ctx)
		if result.Points != 0 {
			t.Errorf("expected 0 points, got %d", result.Points)
		}
	})

	t.Run("partial score with some documented", func(t *testing.T) {
		decl := validate.RawDeclaration{
			"name": "TestElement",
			"attributes": []any{
				map[string]any{
					"name":        "documented",
					"description": "This attribute controls something specific in the component",
					"type":        map[string]any{"text": "string"},
				},
				map[string]any{
					"name": "undocumented",
				},
			},
		}
		ctx := &HealthContext{Declaration: decl}
		result := rule.Evaluate(ctx)
		if result.Points == 0 || result.Points == rule.MaxPoints() {
			t.Errorf("expected partial score, got %d", result.Points)
		}
	})
}
