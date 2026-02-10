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

func TestDemoDocRule(t *testing.T) {
	rule := &DemoDocRule{}

	t.Run("full score with described demo", func(t *testing.T) {
		decl := validate.RawDeclaration{
			"name": "TestElement",
			"demos": []any{
				map[string]any{
					"url":         "demos/basic.html",
					"description": "Basic usage example",
				},
			},
		}
		ctx := &HealthContext{Declaration: decl}
		result := rule.Evaluate(ctx)
		if result.Points != rule.MaxPoints() {
			t.Errorf("expected %d points, got %d", rule.MaxPoints(), result.Points)
		}
	})

	t.Run("partial score with demo but no description", func(t *testing.T) {
		decl := validate.RawDeclaration{
			"name": "TestElement",
			"demos": []any{
				map[string]any{
					"url": "demos/basic.html",
				},
			},
		}
		ctx := &HealthContext{Declaration: decl}
		result := rule.Evaluate(ctx)
		if result.Points != 5 {
			t.Errorf("expected 5 points (has demo, no description), got %d", result.Points)
		}
	})

	t.Run("zero score with no demos", func(t *testing.T) {
		decl := validate.RawDeclaration{
			"name": "TestElement",
		}
		ctx := &HealthContext{Declaration: decl}
		result := rule.Evaluate(ctx)
		if result.Points != 0 {
			t.Errorf("expected 0 points, got %d", result.Points)
		}
	})
}
