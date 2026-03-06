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

import "fmt"

// Evaluate checks a set of rules against the element facts and returns violations.
func Evaluate(rules []Rule, facts ElementFacts) []Violation {
	var violations []Violation

	for _, rule := range rules {
		// For unconditional rules from an attribute's description,
		// implicitly require the source attribute to be present.
		if rule.Condition.Type == ConditionAlways && rule.SourceAttr != "" {
			if !facts.HasAttr[rule.SourceAttr] {
				continue
			}
		}
		if !conditionMet(rule.Condition, facts) {
			continue
		}
		if !requirementSatisfied(rule.Requirement, facts) {
			violations = append(violations, Violation{
				Rule:    rule,
				Message: violationMessage(rule),
			})
		}
	}

	return violations
}

func conditionMet(cond Condition, facts ElementFacts) bool {
	switch cond.Type {
	case ConditionAlways:
		return true
	case ConditionAttrPresent:
		return facts.HasAttr[cond.AttrName]
	case ConditionAttrValue:
		if !facts.HasAttr[cond.AttrName] {
			return false
		}
		return facts.Attributes[cond.AttrName] == cond.AttrValue
	case ConditionAttrNotValue:
		if !facts.HasAttr[cond.AttrName] {
			return false
		}
		return facts.Attributes[cond.AttrName] != cond.AttrValue
	default:
		return false
	}
}

func requirementSatisfied(req Requirement, facts ElementFacts) bool {
	switch req.Type {
	case RequireAttrPresent:
		return facts.HasAttr[req.AttrName]
	case RequireAttrAbsent:
		return !facts.HasAttr[req.AttrName]
	default:
		return true
	}
}

func violationMessage(rule Rule) string {
	switch rule.Requirement.Type {
	case RequireAttrPresent:
		switch rule.Condition.Type {
		case ConditionAlways:
			return fmt.Sprintf("Missing required attribute '%s'. %s", rule.Requirement.AttrName, rule.Source)
		case ConditionAttrPresent:
			return fmt.Sprintf("Attribute '%s' requires '%s' to also be set. %s",
				rule.Condition.AttrName, rule.Requirement.AttrName, rule.Source)
		case ConditionAttrValue:
			return fmt.Sprintf("When '%s' is '%s', attribute '%s' is required. %s",
				rule.Condition.AttrName, rule.Condition.AttrValue, rule.Requirement.AttrName, rule.Source)
		case ConditionAttrNotValue:
			return fmt.Sprintf("When '%s' is not '%s', attribute '%s' is required. %s",
				rule.Condition.AttrName, rule.Condition.AttrValue, rule.Requirement.AttrName, rule.Source)
		}
	case RequireAttrAbsent:
		switch rule.Condition.Type {
		case ConditionAlways:
			return fmt.Sprintf("Attribute '%s' must not be used. %s", rule.Requirement.AttrName, rule.Source)
		case ConditionAttrPresent:
			return fmt.Sprintf("Attribute '%s' must not be used when '%s' is set. %s",
				rule.Requirement.AttrName, rule.Condition.AttrName, rule.Source)
		case ConditionAttrValue:
			return fmt.Sprintf("Attribute '%s' must not be used when '%s' is '%s'. %s",
				rule.Requirement.AttrName, rule.Condition.AttrName, rule.Condition.AttrValue, rule.Source)
		case ConditionAttrNotValue:
			return fmt.Sprintf("Attribute '%s' must not be used when '%s' is not '%s'. %s",
				rule.Requirement.AttrName, rule.Condition.AttrName, rule.Condition.AttrValue, rule.Source)
		}
	}
	return rule.Source
}
