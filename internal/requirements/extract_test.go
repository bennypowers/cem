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

func TestExtractRules_ConditionThenRequirement(t *testing.T) {
	tests := []struct {
		name        string
		description string
		sourceAttr  string
		wantRules   []Rule
	}{
		{
			name:        "if attr is value, MUST set attr",
			description: "The button's theme variant. If you set `variant` to 'icon', you MUST also set `accessible-label`.",
			sourceAttr:  "variant",
			wantRules: []Rule{{
				Condition:   Condition{Type: ConditionAttrValue, AttrName: "variant", AttrValue: "icon"},
				Requirement: Requirement{Type: RequireAttrPresent, AttrName: "accessible-label"},
				SourceAttr:  "variant",
			}},
		},
		{
			name:        "when attr is set, MUST provide attr",
			description: "Shows a loading spinner. When `loading` is set, you MUST provide `accessible-label`.",
			sourceAttr:  "loading",
			wantRules: []Rule{{
				Condition:   Condition{Type: ConditionAttrPresent, AttrName: "loading"},
				Requirement: Requirement{Type: RequireAttrPresent, AttrName: "accessible-label"},
				SourceAttr:  "loading",
			}},
		},
		{
			name:        "if using attr, MUST set attr",
			description: "When using `icon`, you MUST set `accessible-label`.",
			sourceAttr:  "icon",
			wantRules: []Rule{{
				Condition:   Condition{Type: ConditionAttrPresent, AttrName: "icon"},
				Requirement: Requirement{Type: RequireAttrPresent, AttrName: "accessible-label"},
				SourceAttr:  "icon",
			}},
		},
		{
			name:        "if attr equals value, MUST NOT use attr",
			description: "If `variant` is 'compact', you MUST NOT use `icon`.",
			sourceAttr:  "variant",
			wantRules: []Rule{{
				Condition:   Condition{Type: ConditionAttrValue, AttrName: "variant", AttrValue: "compact"},
				Requirement: Requirement{Type: RequireAttrAbsent, AttrName: "icon"},
				SourceAttr:  "variant",
			}},
		},
		{
			name:        "double-quoted value",
			description: `If you set ` + "`variant`" + ` to "icon", you MUST also set ` + "`accessible-label`" + `.`,
			sourceAttr:  "variant",
			wantRules: []Rule{{
				Condition:   Condition{Type: ConditionAttrValue, AttrName: "variant", AttrValue: "icon"},
				Requirement: Requirement{Type: RequireAttrPresent, AttrName: "accessible-label"},
				SourceAttr:  "variant",
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rules := ExtractRules(tt.description, tt.sourceAttr)
			assertRulesMatch(t, tt.wantRules, rules)
		})
	}
}

func TestExtractRules_RequirementThenCondition(t *testing.T) {
	tests := []struct {
		name        string
		description string
		sourceAttr  string
		wantRules   []Rule
	}{
		{
			name:        "MUST set attr when attr is set",
			description: "You MUST set `accessible-label` when `loading` is set.",
			sourceAttr:  "loading",
			wantRules: []Rule{{
				Condition:   Condition{Type: ConditionAttrPresent, AttrName: "loading"},
				Requirement: Requirement{Type: RequireAttrPresent, AttrName: "accessible-label"},
				SourceAttr:  "loading",
			}},
		},
		{
			name:        "MUST provide attr when attr is value",
			description: "You MUST provide `accessible-label` when `variant` is 'icon'.",
			sourceAttr:  "variant",
			wantRules: []Rule{{
				Condition:   Condition{Type: ConditionAttrValue, AttrName: "variant", AttrValue: "icon"},
				Requirement: Requirement{Type: RequireAttrPresent, AttrName: "accessible-label"},
				SourceAttr:  "variant",
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rules := ExtractRules(tt.description, tt.sourceAttr)
			assertRulesMatch(t, tt.wantRules, rules)
		})
	}
}

func TestExtractRules_UnconditionalRequirement(t *testing.T) {
	tests := []struct {
		name        string
		description string
		sourceAttr  string
		wantRules   []Rule
	}{
		{
			name:        "MUST also set attr",
			description: "The icon name. You MUST also set `accessible-label`.",
			sourceAttr:  "icon",
			wantRules: []Rule{{
				Condition:   Condition{Type: ConditionAlways},
				Requirement: Requirement{Type: RequireAttrPresent, AttrName: "accessible-label"},
				SourceAttr:  "icon",
			}},
		},
		{
			name:        "MUST be used in conjunction with attr",
			description: "MUST be used in conjunction with `href`.",
			sourceAttr:  "download",
			wantRules: []Rule{{
				Condition:   Condition{Type: ConditionAlways},
				Requirement: Requirement{Type: RequireAttrPresent, AttrName: "href"},
				SourceAttr:  "download",
			}},
		},
		{
			name:        "MUST be used together with attr",
			description: "MUST be used together with `href`.",
			sourceAttr:  "download",
			wantRules: []Rule{{
				Condition:   Condition{Type: ConditionAlways},
				Requirement: Requirement{Type: RequireAttrPresent, AttrName: "href"},
				SourceAttr:  "download",
			}},
		},
		{
			name:        "REQUIRED keyword",
			description: "`accessible-label` is REQUIRED.",
			sourceAttr:  "accessible-label",
			wantRules:   nil, // no rule: it's about itself, not requiring another attr
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rules := ExtractRules(tt.description, tt.sourceAttr)
			assertRulesMatch(t, tt.wantRules, rules)
		})
	}
}

func TestExtractRules_NoMatch(t *testing.T) {
	tests := []struct {
		name        string
		description string
	}{
		{
			name:        "no RFC 2119 keywords",
			description: "The button's theme variant.",
		},
		{
			name:        "SHOULD keyword ignored",
			description: "If `variant` is 'icon', you SHOULD set `accessible-label`.",
		},
		{
			name:        "MAY keyword ignored",
			description: "You MAY set `accessible-label` when `loading` is set.",
		},
		{
			name:        "no backtick-quoted attrs",
			description: "If you set variant to icon, you MUST also set accessible-label.",
		},
		{
			name:        "value constraint not conditional requirement",
			description: "This attribute MUST be a valid URL.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rules := ExtractRules(tt.description, "test")
			if len(rules) != 0 {
				t.Errorf("expected no rules, got %d: %+v", len(rules), rules)
			}
		})
	}
}

// Tests for the signal-based approach: phrasing variations the old regex couldn't handle

func TestExtractRules_PassiveVoice(t *testing.T) {
	tests := []struct {
		name        string
		description string
		sourceAttr  string
		wantRules   []Rule
	}{
		{
			name:        "attr MUST be provided when condition",
			description: "`accessible-label` MUST be provided when `variant` is 'icon'.",
			sourceAttr:  "variant",
			wantRules: []Rule{{
				Condition:   Condition{Type: ConditionAttrValue, AttrName: "variant", AttrValue: "icon"},
				Requirement: Requirement{Type: RequireAttrPresent, AttrName: "accessible-label"},
				SourceAttr:  "variant",
			}},
		},
		{
			name:        "attr MUST be set when condition present",
			description: "`accessible-label` MUST be set when `loading` is present.",
			sourceAttr:  "loading",
			wantRules: []Rule{{
				Condition:   Condition{Type: ConditionAttrPresent, AttrName: "loading"},
				Requirement: Requirement{Type: RequireAttrPresent, AttrName: "accessible-label"},
				SourceAttr:  "loading",
			}},
		},
		{
			name:        "attr is REQUIRED when condition",
			description: "`accessible-label` is REQUIRED when `icon` is set.",
			sourceAttr:  "icon",
			wantRules: []Rule{{
				Condition:   Condition{Type: ConditionAttrPresent, AttrName: "icon"},
				Requirement: Requirement{Type: RequireAttrPresent, AttrName: "accessible-label"},
				SourceAttr:  "icon",
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rules := ExtractRules(tt.description, tt.sourceAttr)
			assertRulesMatch(t, tt.wantRules, rules)
		})
	}
}

func TestExtractRules_ArbitraryVerbs(t *testing.T) {
	tests := []struct {
		name        string
		description string
		sourceAttr  string
		wantRules   []Rule
	}{
		{
			name:        "MUST configure attr",
			description: "You MUST configure `accessible-label` when `icon` is set.",
			sourceAttr:  "icon",
			wantRules: []Rule{{
				Condition:   Condition{Type: ConditionAttrPresent, AttrName: "icon"},
				Requirement: Requirement{Type: RequireAttrPresent, AttrName: "accessible-label"},
				SourceAttr:  "icon",
			}},
		},
		{
			name:        "MUST define attr",
			description: "If `variant` is 'icon', you MUST define `accessible-label`.",
			sourceAttr:  "variant",
			wantRules: []Rule{{
				Condition:   Condition{Type: ConditionAttrValue, AttrName: "variant", AttrValue: "icon"},
				Requirement: Requirement{Type: RequireAttrPresent, AttrName: "accessible-label"},
				SourceAttr:  "variant",
			}},
		},
		{
			name:        "MUST specify attr",
			description: "When using `icon`, you MUST specify `accessible-label`.",
			sourceAttr:  "icon",
			wantRules: []Rule{{
				Condition:   Condition{Type: ConditionAttrPresent, AttrName: "icon"},
				Requirement: Requirement{Type: RequireAttrPresent, AttrName: "accessible-label"},
				SourceAttr:  "icon",
			}},
		},
		{
			name:        "MUST ensure attr is present",
			description: "When `loading` is set, you MUST ensure `accessible-label` is present.",
			sourceAttr:  "loading",
			wantRules: []Rule{{
				Condition:   Condition{Type: ConditionAttrPresent, AttrName: "loading"},
				Requirement: Requirement{Type: RequireAttrPresent, AttrName: "accessible-label"},
				SourceAttr:  "loading",
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rules := ExtractRules(tt.description, tt.sourceAttr)
			assertRulesMatch(t, tt.wantRules, rules)
		})
	}
}

func TestExtractRules_NegatedConditions(t *testing.T) {
	tests := []struct {
		name        string
		description string
		sourceAttr  string
		wantRules   []Rule
	}{
		{
			name:        "when attr is not value, MUST set attr",
			description: "When `variant` is not 'default', you MUST set `accessible-label`.",
			sourceAttr:  "variant",
			wantRules: []Rule{{
				Condition:   Condition{Type: ConditionAttrNotValue, AttrName: "variant", AttrValue: "default"},
				Requirement: Requirement{Type: RequireAttrPresent, AttrName: "accessible-label"},
				SourceAttr:  "variant",
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rules := ExtractRules(tt.description, tt.sourceAttr)
			assertRulesMatch(t, tt.wantRules, rules)
		})
	}
}

func TestExtractRules_MultipleValues(t *testing.T) {
	tests := []struct {
		name        string
		description string
		sourceAttr  string
		wantRules   []Rule
	}{
		{
			name:        "if attr is value or value, MUST set attr",
			description: "If `variant` is 'icon' or 'loading', you MUST set `accessible-label`.",
			sourceAttr:  "variant",
			wantRules: []Rule{
				{
					Condition:   Condition{Type: ConditionAttrValue, AttrName: "variant", AttrValue: "icon"},
					Requirement: Requirement{Type: RequireAttrPresent, AttrName: "accessible-label"},
					SourceAttr:  "variant",
				},
				{
					Condition:   Condition{Type: ConditionAttrValue, AttrName: "variant", AttrValue: "loading"},
					Requirement: Requirement{Type: RequireAttrPresent, AttrName: "accessible-label"},
					SourceAttr:  "variant",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rules := ExtractRules(tt.description, tt.sourceAttr)
			assertRulesMatch(t, tt.wantRules, rules)
		})
	}
}

func TestExtractRules_PresenceVariations(t *testing.T) {
	tests := []struct {
		name        string
		description string
		sourceAttr  string
		wantRules   []Rule
	}{
		{
			name:        "attr is defined",
			description: "When `icon` is defined, you MUST set `accessible-label`.",
			sourceAttr:  "icon",
			wantRules: []Rule{{
				Condition:   Condition{Type: ConditionAttrPresent, AttrName: "icon"},
				Requirement: Requirement{Type: RequireAttrPresent, AttrName: "accessible-label"},
				SourceAttr:  "icon",
			}},
		},
		{
			name:        "attr is specified",
			description: "When `icon` is specified, `accessible-label` is REQUIRED.",
			sourceAttr:  "icon",
			wantRules: []Rule{{
				Condition:   Condition{Type: ConditionAttrPresent, AttrName: "icon"},
				Requirement: Requirement{Type: RequireAttrPresent, AttrName: "accessible-label"},
				SourceAttr:  "icon",
			}},
		},
		{
			name:        "attr is provided",
			description: "If `icon` is provided, you MUST also set `accessible-label`.",
			sourceAttr:  "icon",
			wantRules: []Rule{{
				Condition:   Condition{Type: ConditionAttrPresent, AttrName: "icon"},
				Requirement: Requirement{Type: RequireAttrPresent, AttrName: "accessible-label"},
				SourceAttr:  "icon",
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rules := ExtractRules(tt.description, tt.sourceAttr)
			assertRulesMatch(t, tt.wantRules, rules)
		})
	}
}

func TestExtractSignals(t *testing.T) {
	t.Run("extracts attr refs", func(t *testing.T) {
		signals := extractSignals("If `variant` is 'icon', you MUST set `accessible-label`.")
		if len(signals.attrRefs) != 2 {
			t.Fatalf("expected 2 attr refs, got %d: %+v", len(signals.attrRefs), signals.attrRefs)
		}
		if signals.attrRefs[0].name != "variant" {
			t.Errorf("expected first attr ref 'variant', got %q", signals.attrRefs[0].name)
		}
		if signals.attrRefs[1].name != "accessible-label" {
			t.Errorf("expected second attr ref 'accessible-label', got %q", signals.attrRefs[1].name)
		}
	})

	t.Run("extracts quoted values", func(t *testing.T) {
		signals := extractSignals("If `variant` is 'icon', you MUST set `label`.")
		if len(signals.quotedValues) != 1 {
			t.Fatalf("expected 1 quoted value, got %d", len(signals.quotedValues))
		}
		if signals.quotedValues[0].value != "icon" {
			t.Errorf("expected value 'icon', got %q", signals.quotedValues[0].value)
		}
	})

	t.Run("extracts MUST keyword", func(t *testing.T) {
		signals := extractSignals("You MUST set `label`.")
		if signals.mustKeyword == nil {
			t.Fatal("expected MUST keyword")
		}
		if signals.mustKeyword.negated {
			t.Error("expected non-negated keyword")
		}
	})

	t.Run("extracts MUST NOT keyword", func(t *testing.T) {
		signals := extractSignals("You MUST NOT set `label`.")
		if signals.mustKeyword == nil {
			t.Fatal("expected MUST NOT keyword")
		}
		if !signals.mustKeyword.negated {
			t.Error("expected negated keyword")
		}
	})

	t.Run("extracts conditional marker", func(t *testing.T) {
		signals := extractSignals("If `variant` is 'icon', you MUST set `label`.")
		if len(signals.conditionals) == 0 {
			t.Fatal("expected conditional marker")
		}
	})

	t.Run("extracts negation near condition", func(t *testing.T) {
		signals := extractSignals("When `variant` is not 'default', you MUST set `label`.")
		if len(signals.negations) == 0 {
			t.Fatal("expected negation")
		}
	})

	t.Run("extracts multiple quoted values with or", func(t *testing.T) {
		signals := extractSignals("If `variant` is 'icon' or 'loading', you MUST set `label`.")
		if len(signals.quotedValues) != 2 {
			t.Fatalf("expected 2 quoted values, got %d: %+v", len(signals.quotedValues), signals.quotedValues)
		}
	})
}

func assertRulesMatch(t *testing.T, want, got []Rule) {
	t.Helper()
	if len(want) != len(got) {
		t.Fatalf("expected %d rules, got %d\nwant: %+v\ngot:  %+v", len(want), len(got), want, got)
	}
	for i := range want {
		w, g := want[i], got[i]
		if w.Condition.Type != g.Condition.Type {
			t.Errorf("rule[%d] condition type: want %v, got %v", i, w.Condition.Type, g.Condition.Type)
		}
		if w.Condition.AttrName != g.Condition.AttrName {
			t.Errorf("rule[%d] condition attr: want %q, got %q", i, w.Condition.AttrName, g.Condition.AttrName)
		}
		if w.Condition.AttrValue != g.Condition.AttrValue {
			t.Errorf("rule[%d] condition value: want %q, got %q", i, w.Condition.AttrValue, g.Condition.AttrValue)
		}
		if w.Requirement.Type != g.Requirement.Type {
			t.Errorf("rule[%d] requirement type: want %v, got %v", i, w.Requirement.Type, g.Requirement.Type)
		}
		if w.Requirement.AttrName != g.Requirement.AttrName {
			t.Errorf("rule[%d] requirement attr: want %q, got %q", i, w.Requirement.AttrName, g.Requirement.AttrName)
		}
		if w.SourceAttr != g.SourceAttr {
			t.Errorf("rule[%d] source attr: want %q, got %q", i, w.SourceAttr, g.SourceAttr)
		}
		if g.Source == "" {
			t.Errorf("rule[%d] source sentence should not be empty", i)
		}
	}
}
