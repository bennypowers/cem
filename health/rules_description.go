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
	"regexp"
	"strings"
)

// wordPattern compiles a regex for whole-word matching with word boundaries.
func wordPattern(word string) *regexp.Regexp {
	return regexp.MustCompile(`\b` + regexp.QuoteMeta(word) + `\b`)
}

// containsWord checks for a whole-word match using a pre-compiled pattern.
func containsWord(text string, pattern *regexp.Regexp) bool {
	return pattern.MatchString(text)
}

// DescriptionRule scores element description quality (25 pts max).
type DescriptionRule struct{}

func (r *DescriptionRule) ID() string       { return "description" }
func (r *DescriptionRule) Category() string { return "Element description" }
func (r *DescriptionRule) MaxPoints() int   { return 25 }

func (r *DescriptionRule) Evaluate(ctx *HealthContext) CategoryScore {
	desc := ctx.Declaration.Description()
	if desc == "" {
		desc = ctx.Declaration.Summary()
	}

	var findings []Finding
	totalPoints := 0

	// Check 1: Has description (5 pts)
	{
		pts := 0
		msg := ""
		if desc != "" {
			pts = 5
		} else {
			msg = "add a description to explain what this declaration does"
		}
		totalPoints += pts
		findings = append(findings, Finding{Check: "Has description", Points: pts, Max: 5, Message: msg})
	}

	// Check 2: Optimal length (3 pts)
	{
		pts := 0
		msg := ""
		if desc == "" {
			msg = "add a description"
		} else if len(desc) < 20 {
			msg = "description is too short; provide at least 20 characters"
		} else if len(desc) > 400 {
			msg = "description is longer than 400 characters; consider being more concise"
		} else {
			pts = 3
		}
		totalPoints += pts
		findings = append(findings, Finding{Check: "Optimal length", Points: pts, Max: 3, Message: msg})
	}

	// Check 3: Explains purpose/context (5 pts)
	{
		pts := 0
		msg := ""
		lower := strings.ToLower(desc)
		purposeWords := []*regexp.Regexp{
			wordPattern("for"), wordPattern("when"), wordPattern("use"),
			wordPattern("provides"), wordPattern("allows"), wordPattern("enables"),
		}
		found := 0
		for _, w := range purposeWords {
			if containsWord(lower, w) {
				found++
			}
		}
		if found >= 2 {
			pts = 5
		} else if found == 1 {
			pts = 3
		}
		if pts < 5 {
			msg = "describe purpose and context using words like 'for', 'when', 'provides', 'allows'"
		}
		totalPoints += pts
		findings = append(findings, Finding{Check: "Explains purpose/context", Points: pts, Max: 5, Message: msg})
	}

	// Check 4: Uses RFC 2119 keywords (5 pts)
	{
		pts := 0
		msg := ""
		lower := strings.ToLower(desc)
		rfcWords := []*regexp.Regexp{
			wordPattern("must"), wordPattern("should"), wordPattern("avoid"),
			wordPattern("shall"), wordPattern("required"), wordPattern("recommended"),
		}
		found := 0
		for _, w := range rfcWords {
			if containsWord(lower, w) {
				found++
			}
		}
		if found >= 2 {
			pts = 5
		} else if found == 1 {
			pts = 3
		}
		if pts < 5 {
			msg = "use RFC 2119 keywords (MUST, SHOULD, AVOID) to clarify requirements"
		}
		totalPoints += pts
		findings = append(findings, Finding{Check: "RFC 2119 keywords", Points: pts, Max: 5, Message: msg})
	}

	// Check 5: Accessibility notes (4 pts)
	{
		pts := 0
		msg := ""
		lower := strings.ToLower(desc)
		a11yWords := []*regexp.Regexp{
			wordPattern("accessibility"), wordPattern("a11y"), wordPattern("aria"),
			wordPattern("screen reader"), wordPattern("wcag"),
		}
		for _, w := range a11yWords {
			if containsWord(lower, w) {
				pts = 4
				break
			}
		}
		if pts == 0 {
			msg = "mention accessibility considerations (ARIA, screen reader, WCAG)"
		}
		totalPoints += pts
		findings = append(findings, Finding{Check: "Accessibility notes", Points: pts, Max: 4, Message: msg})
	}

	// Check 6: Keyboard interaction (3 pts)
	{
		pts := 0
		msg := ""
		lower := strings.ToLower(desc)
		kbWords := []*regexp.Regexp{
			wordPattern("keyboard"), wordPattern("key"), wordPattern("enter"),
			wordPattern("space"), wordPattern("tab"), wordPattern("focus"),
		}
		for _, w := range kbWords {
			if containsWord(lower, w) {
				pts = 3
				break
			}
		}
		if pts == 0 {
			msg = "document keyboard interaction (Enter, Space, Tab, focus)"
		}
		totalPoints += pts
		findings = append(findings, Finding{Check: "Keyboard interaction", Points: pts, Max: 3, Message: msg})
	}

	return CategoryScore{
		ID:        r.ID(),
		Category:  r.Category(),
		Points:    totalPoints,
		MaxPoints: r.MaxPoints(),
		Status:    status(totalPoints, r.MaxPoints()),
		Findings:  findings,
	}
}
