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
	"regexp"
	"strings"

	M "bennypowers.dev/cem/manifest"
)

// Signal types extracted from a sentence

type attrRef struct {
	name  string
	start int // byte position in sentence
	end   int
}

type quotedValue struct {
	value string
	start int
	end   int
}

type mustKeywordSignal struct {
	negated bool // MUST NOT / SHALL NOT
	start   int
	end     int
}

type conditionalSignal struct {
	start int
	end   int
}

type negationSignal struct {
	start int
	end   int
}

type signals struct {
	attrRefs     []attrRef
	quotedValues []quotedValue
	mustKeyword  *mustKeywordSignal
	conditionals []conditionalSignal
	negations    []negationSignal
}

// Compiled regexes for signal extraction

var (
	// Backtick-quoted names: `variant`, `accessible-label`
	attrRefRegex = regexp.MustCompile("`([a-zA-Z][a-zA-Z0-9-]*)`")

	// Single or double quoted values: 'icon', "icon"
	quotedValueRegex = regexp.MustCompile(`'([^']+)'|"([^"]+)"`)

	// MUST-level RFC 2119 keywords (case-insensitive)
	mustKeywordRegex = regexp.MustCompile(`(?i)\b(MUST\s+NOT|SHALL\s+NOT|MUST|SHALL|REQUIRED)\b`)

	// Conditional markers
	conditionalRegex = regexp.MustCompile(`(?i)\b(if|when|whenever|where)\b`)

	// Negation words (excluding "NOT" after MUST/SHALL, which is handled separately)
	negationRegex = regexp.MustCompile(`(?i)\b(not|never|without)\b`)

	// Sentence-ending punctuation
	sentenceEndRegex = regexp.MustCompile(`[.!?](?:\s+|$)`)
)

// extractSignals finds all meaningful signals in a sentence.
func extractSignals(sentence string) signals {
	var s signals

	// Extract attr refs
	for _, match := range attrRefRegex.FindAllStringSubmatchIndex(sentence, -1) {
		s.attrRefs = append(s.attrRefs, attrRef{
			name:  sentence[match[2]:match[3]],
			start: match[0],
			end:   match[1],
		})
	}

	// Extract quoted values
	for _, match := range quotedValueRegex.FindAllStringSubmatchIndex(sentence, -1) {
		value := ""
		if match[2] >= 0 {
			value = sentence[match[2]:match[3]] // single-quoted
		} else if match[4] >= 0 {
			value = sentence[match[4]:match[5]] // double-quoted
		}
		// Skip values that look like they're inside backticks (attr refs)
		if isInsideBackticks(sentence, match[0]) {
			continue
		}
		s.quotedValues = append(s.quotedValues, quotedValue{
			value: value,
			start: match[0],
			end:   match[1],
		})
	}

	// Extract MUST keyword
	if match := mustKeywordRegex.FindStringSubmatchIndex(sentence); match != nil {
		keyword := strings.ToUpper(strings.TrimSpace(sentence[match[2]:match[3]]))
		negated := strings.Contains(keyword, "NOT")
		s.mustKeyword = &mustKeywordSignal{
			negated: negated,
			start:   match[0],
			end:     match[1],
		}
	}

	// Extract conditionals
	for _, match := range conditionalRegex.FindAllStringSubmatchIndex(sentence, -1) {
		s.conditionals = append(s.conditionals, conditionalSignal{
			start: match[0],
			end:   match[1],
		})
	}

	// Extract negations (but not ones that are part of MUST NOT / SHALL NOT)
	for _, match := range negationRegex.FindAllStringSubmatchIndex(sentence, -1) {
		// Skip if this "not" is part of the MUST NOT keyword
		if s.mustKeyword != nil && match[0] >= s.mustKeyword.start && match[1] <= s.mustKeyword.end {
			continue
		}
		s.negations = append(s.negations, negationSignal{
			start: match[0],
			end:   match[1],
		})
	}

	return s
}

func isInsideBackticks(s string, pos int) bool {
	// Check if position is between a pair of backticks
	backticksBefore := strings.Count(s[:pos], "`")
	return backticksBefore%2 == 1
}

// splitSentences splits text into sentences, preserving the ending punctuation.
func splitSentences(text string) []string {
	text = strings.TrimSpace(text)
	if text == "" {
		return nil
	}

	var sentences []string
	indices := sentenceEndRegex.FindAllStringIndex(text, -1)

	if len(indices) == 0 {
		return []string{text}
	}

	start := 0
	for _, idx := range indices {
		end := idx[1]
		sentence := strings.TrimSpace(text[start:end])
		if sentence != "" {
			sentences = append(sentences, sentence)
		}
		start = end
	}

	if start < len(text) {
		remaining := strings.TrimSpace(text[start:])
		if remaining != "" {
			sentences = append(sentences, remaining)
		}
	}

	return sentences
}

// ExtractRules extracts conditional requirement rules from a description string.
func ExtractRules(description, sourceAttr string) []Rule {
	var rules []Rule

	sentences := splitSentences(description)
	for _, sentence := range sentences {
		rules = append(rules, inferRules(sentence, sourceAttr)...)
	}

	return rules
}

// inferRules uses signal extraction + positional inference to find rules.
func inferRules(sentence, sourceAttr string) []Rule {
	s := extractSignals(sentence)

	// Must have a MUST-level keyword
	if s.mustKeyword == nil {
		return nil
	}

	// Must have at least one backtick-quoted attr ref
	if len(s.attrRefs) == 0 {
		return nil
	}

	// Determine if there's a conditional clause
	hasConditional := len(s.conditionals) > 0

	// Find condition and requirement via positional inference
	condAttr, condValues, condNegated := findCondition(s, hasConditional)
	reqAttr := findRequirement(s, condAttr)

	// If we couldn't determine a requirement attr, nothing to do
	if reqAttr == "" {
		return nil
	}

	// Don't create rules where the requirement is the same as the source
	if reqAttr == sourceAttr {
		return nil
	}

	// Determine requirement type
	reqType := RequireAttrPresent
	if s.mustKeyword.negated {
		reqType = RequireAttrAbsent
	}

	// Build condition
	if condAttr == "" {
		// Unconditional: "You MUST set `attr`"
		return []Rule{{
			Condition:   Condition{Type: ConditionAlways},
			Requirement: Requirement{Type: reqType, AttrName: reqAttr},
			Source:      sentence,
			SourceAttr:  sourceAttr,
		}}
	}

	// Multiple values: "if `attr` is 'a' or 'b'" -> one rule per value
	if len(condValues) > 1 {
		var rules []Rule
		for _, v := range condValues {
			condType := ConditionAttrValue
			if condNegated {
				condType = ConditionAttrNotValue
			}
			rules = append(rules, Rule{
				Condition:   Condition{Type: condType, AttrName: condAttr, AttrValue: v},
				Requirement: Requirement{Type: reqType, AttrName: reqAttr},
				Source:      sentence,
				SourceAttr:  sourceAttr,
			})
		}
		return rules
	}

	// Single value condition
	if len(condValues) == 1 {
		condType := ConditionAttrValue
		if condNegated {
			condType = ConditionAttrNotValue
		}
		return []Rule{{
			Condition:   Condition{Type: condType, AttrName: condAttr, AttrValue: condValues[0]},
			Requirement: Requirement{Type: reqType, AttrName: reqAttr},
			Source:      sentence,
			SourceAttr:  sourceAttr,
		}}
	}

	// Presence condition (no value)
	return []Rule{{
		Condition:   Condition{Type: ConditionAttrPresent, AttrName: condAttr},
		Requirement: Requirement{Type: reqType, AttrName: reqAttr},
		Source:      sentence,
		SourceAttr:  sourceAttr,
	}}
}

// findCondition determines the condition attribute, its values, and whether it's negated.
func findCondition(s signals, hasConditional bool) (string, []string, bool) {
	if !hasConditional {
		return "", nil, false
	}

	condPos := s.conditionals[0].start

	// The condition attr is the one closest to the conditional marker
	condRef := nearestAttrRef(s.attrRefs, condPos)
	if condRef == nil {
		return "", nil, false
	}

	// Find quoted values near the condition attr
	var values []string
	for _, qv := range s.quotedValues {
		if isNear(qv.start, condRef.start, condRef.end, 80) {
			values = append(values, qv.value)
		}
	}

	// Check for negation near the condition
	negated := false
	for _, neg := range s.negations {
		// Negation between the condition attr and the closest quoted value,
		// or between the conditional and the value
		if neg.start > condRef.start-20 && neg.start < condRef.end+40 {
			negated = true
			break
		}
		if len(values) > 0 {
			// Also check if negation is between condRef and the first value
			for _, qv := range s.quotedValues {
				if neg.start > condRef.end && neg.start < qv.start {
					negated = true
					break
				}
			}
		}
	}

	// If no quoted values, check if there are presence words near the attr
	// to distinguish "when `attr` is set" (presence) from just "when `attr`" (also presence)
	// Both mean the same thing: attr presence condition

	return condRef.name, values, negated
}

// findRequirement determines the requirement attribute.
func findRequirement(s signals, condAttr string) string {
	// The requirement attr is the one on the MUST side that isn't the condition
	// Typical patterns:
	//   "If `condAttr` is X, MUST set `reqAttr`"  -> reqAttr is after MUST
	//   "`reqAttr` MUST be set when `condAttr`"    -> reqAttr is before MUST

	candidates := make([]attrRef, 0, len(s.attrRefs))
	for _, ref := range s.attrRefs {
		if ref.name != condAttr {
			candidates = append(candidates, ref)
		}
	}

	if len(candidates) == 0 {
		return ""
	}

	if len(candidates) == 1 {
		return candidates[0].name
	}

	// Multiple candidates: pick the one nearest to the MUST keyword
	best := nearestAttrRef(candidates, s.mustKeyword.start)
	if best != nil {
		return best.name
	}

	return candidates[0].name
}

// nearestAttrRef finds the attr ref nearest to a given position.
func nearestAttrRef(refs []attrRef, pos int) *attrRef {
	if len(refs) == 0 {
		return nil
	}
	best := &refs[0]
	bestDist := abs(refs[0].start - pos)
	for i := 1; i < len(refs); i++ {
		d := abs(refs[i].start - pos)
		if d < bestDist {
			bestDist = d
			best = &refs[i]
		}
	}
	return best
}

func isNear(pos, refStart, refEnd, threshold int) bool {
	if pos >= refStart && pos <= refEnd {
		return true
	}
	return abs(pos-refStart) < threshold || abs(pos-refEnd) < threshold
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// ExtractElementRules extracts all rules from all descriptions on an element declaration.
func ExtractElementRules(decl *M.CustomElementDeclaration) []Rule {
	if decl == nil {
		return nil
	}

	var rules []Rule

	for _, attr := range decl.Attributes() {
		desc := attr.Description
		if desc == "" {
			desc = attr.Summary
		}
		if desc != "" {
			rules = append(rules, ExtractRules(desc, attr.Name)...)
		}
	}

	for _, slot := range decl.Slots() {
		desc := slot.Description
		if desc == "" {
			desc = slot.Summary
		}
		if desc != "" {
			rules = append(rules, ExtractRules(desc, slot.Name)...)
		}
	}

	return rules
}
