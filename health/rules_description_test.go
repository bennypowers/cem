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

func TestDescriptionRule(t *testing.T) {
	rule := &DescriptionRule{}

	t.Run("full score with rich description", func(t *testing.T) {
		decl := validate.RawDeclaration{
			"name":        "TestElement",
			"description": "A component for displaying data when the user provides input. It allows filtering and enables keyboard navigation with Tab and Enter. Must follow WCAG guidelines and should include aria labels for screen reader users.",
		}
		ctx := &HealthContext{Declaration: decl, DeclName: "TestElement"}
		result := rule.Evaluate(ctx)
		if result.Points != rule.MaxPoints() {
			t.Errorf("expected %d points, got %d", rule.MaxPoints(), result.Points)
		}
		if result.Status != "pass" {
			t.Errorf("expected status 'pass', got %q", result.Status)
		}
	})

	t.Run("zero score with no description", func(t *testing.T) {
		decl := validate.RawDeclaration{
			"name": "TestElement",
		}
		ctx := &HealthContext{Declaration: decl, DeclName: "TestElement"}
		result := rule.Evaluate(ctx)
		if result.Points != 0 {
			t.Errorf("expected 0 points, got %d", result.Points)
		}
		if result.Status != "fail" {
			t.Errorf("expected status 'fail', got %q", result.Status)
		}
	})

	t.Run("falls back to summary", func(t *testing.T) {
		decl := validate.RawDeclaration{
			"name":    "TestElement",
			"summary": "Provides a way to display things",
		}
		ctx := &HealthContext{Declaration: decl, DeclName: "TestElement"}
		result := rule.Evaluate(ctx)
		if result.Points == 0 {
			t.Error("expected some points when summary exists")
		}
	})

	t.Run("too long description", func(t *testing.T) {
		long := ""
		for len(long) <= 400 {
			long += "This is a very long description. "
		}
		decl := validate.RawDeclaration{
			"name":        "TestElement",
			"description": long,
		}
		ctx := &HealthContext{Declaration: decl, DeclName: "TestElement"}
		result := rule.Evaluate(ctx)
		// Should get "Has description" (5) but not "Optimal length" (0)
		found := false
		for _, f := range result.Findings {
			if f.Check == "Optimal length" {
				if f.Points != 0 {
					t.Errorf("expected 0 points for long description, got %d", f.Points)
				}
				found = true
			}
		}
		if !found {
			t.Error("expected 'Optimal length' finding")
		}
	})
}
