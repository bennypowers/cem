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

// Package requirements extracts and evaluates conditional requirements
// from custom element manifest descriptions using RFC 2119 keywords.
package requirements

// ConditionType describes what triggers a rule.
type ConditionType int

const (
	ConditionAlways       ConditionType = iota // unconditional requirement
	ConditionAttrPresent                       // "when `loading` is set"
	ConditionAttrValue                         // "if `variant` is 'icon'"
	ConditionAttrNotValue                      // "when `variant` is not 'default'"
)

// RequirementType describes what the rule demands.
type RequirementType int

const (
	RequireAttrPresent RequirementType = iota // "MUST set `accessible-label`"
	RequireAttrAbsent                         // "MUST NOT use `icon`"
)

// Condition describes when a rule applies.
type Condition struct {
	Type      ConditionType
	AttrName  string // attribute referenced in condition
	AttrValue string // specific value (only for ConditionAttrValue)
}

// Requirement describes what the rule demands.
type Requirement struct {
	Type     RequirementType
	AttrName string // attribute to require or forbid
}

// Rule is a single conditional requirement extracted from a description.
type Rule struct {
	Condition   Condition
	Requirement Requirement
	Source      string // original sentence for diagnostic messages
	SourceAttr  string // which attribute's description this came from
}

// Violation represents a rule that was not satisfied.
type Violation struct {
	Rule    Rule
	Message string
}

// ElementFacts describes the current state of an element instance in HTML.
type ElementFacts struct {
	Attributes map[string]string // attr name -> value ("" = present with no value)
	HasAttr    map[string]bool   // tracks attribute presence
}
