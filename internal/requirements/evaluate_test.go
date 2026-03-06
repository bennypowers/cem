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
package requirements

import (
	"testing"
)

func TestEvaluate_ConditionAttrValue(t *testing.T) {
	rules := []Rule{{
		Condition:   Condition{Type: ConditionAttrValue, AttrName: "variant", AttrValue: "icon"},
		Requirement: Requirement{Type: RequireAttrPresent, AttrName: "accessible-label"},
		Source:      "If you set `variant` to 'icon', you MUST also set `accessible-label`.",
		SourceAttr:  "variant",
	}}

	t.Run("violation when condition met and requirement missing", func(t *testing.T) {
		facts := ElementFacts{
			Attributes: map[string]string{"variant": "icon"},
			HasAttr:    map[string]bool{"variant": true},
		}
		violations := Evaluate(rules, facts)
		if len(violations) != 1 {
			t.Fatalf("expected 1 violation, got %d", len(violations))
		}
		if violations[0].Message == "" {
			t.Error("violation message should not be empty")
		}
	})

	t.Run("no violation when condition met and requirement satisfied", func(t *testing.T) {
		facts := ElementFacts{
			Attributes: map[string]string{"variant": "icon", "accessible-label": "Dog icon"},
			HasAttr:    map[string]bool{"variant": true, "accessible-label": true},
		}
		violations := Evaluate(rules, facts)
		if len(violations) != 0 {
			t.Fatalf("expected 0 violations, got %d: %+v", len(violations), violations)
		}
	})

	t.Run("no violation when condition not met", func(t *testing.T) {
		facts := ElementFacts{
			Attributes: map[string]string{"variant": "primary"},
			HasAttr:    map[string]bool{"variant": true},
		}
		violations := Evaluate(rules, facts)
		if len(violations) != 0 {
			t.Fatalf("expected 0 violations, got %d", len(violations))
		}
	})

	t.Run("no violation when condition attr absent", func(t *testing.T) {
		facts := ElementFacts{
			Attributes: map[string]string{},
			HasAttr:    map[string]bool{},
		}
		violations := Evaluate(rules, facts)
		if len(violations) != 0 {
			t.Fatalf("expected 0 violations, got %d", len(violations))
		}
	})
}

func TestEvaluate_ConditionAttrPresent(t *testing.T) {
	rules := []Rule{{
		Condition:   Condition{Type: ConditionAttrPresent, AttrName: "loading"},
		Requirement: Requirement{Type: RequireAttrPresent, AttrName: "accessible-label"},
		Source:      "When `loading` is set, you MUST provide `accessible-label`.",
		SourceAttr:  "loading",
	}}

	t.Run("violation when attr present but requirement missing", func(t *testing.T) {
		facts := ElementFacts{
			Attributes: map[string]string{"loading": ""},
			HasAttr:    map[string]bool{"loading": true},
		}
		violations := Evaluate(rules, facts)
		if len(violations) != 1 {
			t.Fatalf("expected 1 violation, got %d", len(violations))
		}
	})

	t.Run("no violation when attr absent", func(t *testing.T) {
		facts := ElementFacts{
			Attributes: map[string]string{},
			HasAttr:    map[string]bool{},
		}
		violations := Evaluate(rules, facts)
		if len(violations) != 0 {
			t.Fatalf("expected 0 violations, got %d", len(violations))
		}
	})
}

func TestEvaluate_RequireAttrAbsent(t *testing.T) {
	rules := []Rule{{
		Condition:   Condition{Type: ConditionAttrValue, AttrName: "variant", AttrValue: "compact"},
		Requirement: Requirement{Type: RequireAttrAbsent, AttrName: "icon"},
		Source:      "If `variant` is 'compact', you MUST NOT use `icon`.",
		SourceAttr:  "variant",
	}}

	t.Run("violation when forbidden attr present", func(t *testing.T) {
		facts := ElementFacts{
			Attributes: map[string]string{"variant": "compact", "icon": "star"},
			HasAttr:    map[string]bool{"variant": true, "icon": true},
		}
		violations := Evaluate(rules, facts)
		if len(violations) != 1 {
			t.Fatalf("expected 1 violation, got %d", len(violations))
		}
	})

	t.Run("no violation when forbidden attr absent", func(t *testing.T) {
		facts := ElementFacts{
			Attributes: map[string]string{"variant": "compact"},
			HasAttr:    map[string]bool{"variant": true},
		}
		violations := Evaluate(rules, facts)
		if len(violations) != 0 {
			t.Fatalf("expected 0 violations, got %d", len(violations))
		}
	})
}

func TestEvaluate_UnconditionalRequirement(t *testing.T) {
	rules := []Rule{{
		Condition:   Condition{Type: ConditionAlways},
		Requirement: Requirement{Type: RequireAttrPresent, AttrName: "href"},
		Source:      "MUST be used in conjunction with `href`.",
		SourceAttr:  "download",
	}}

	t.Run("violation when requirement missing", func(t *testing.T) {
		facts := ElementFacts{
			Attributes: map[string]string{"download": ""},
			HasAttr:    map[string]bool{"download": true},
		}
		violations := Evaluate(rules, facts)
		if len(violations) != 1 {
			t.Fatalf("expected 1 violation, got %d", len(violations))
		}
	})

	t.Run("no violation when requirement satisfied", func(t *testing.T) {
		facts := ElementFacts{
			Attributes: map[string]string{"download": "", "href": "/file.pdf"},
			HasAttr:    map[string]bool{"download": true, "href": true},
		}
		violations := Evaluate(rules, facts)
		if len(violations) != 0 {
			t.Fatalf("expected 0 violations, got %d", len(violations))
		}
	})

	t.Run("no violation when source attr not present", func(t *testing.T) {
		facts := ElementFacts{
			Attributes: map[string]string{},
			HasAttr:    map[string]bool{},
		}
		violations := Evaluate(rules, facts)
		if len(violations) != 0 {
			t.Fatalf("expected 0 violations, got %d", len(violations))
		}
	})
}
