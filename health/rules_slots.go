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
	"strings"
)

// SlotDocRule scores slot documentation quality (15 pts max).
type SlotDocRule struct{}

func (r *SlotDocRule) ID() string        { return "slots" }
func (r *SlotDocRule) Category() string   { return "Slot documentation" }
func (r *SlotDocRule) MaxPoints() int     { return 15 }

func (r *SlotDocRule) Evaluate(ctx *HealthContext) CategoryScore {
	slots, ok := ctx.Declaration.Slots()
	total := 0
	if ok {
		total = len(slots)
	}

	var findings []Finding
	totalPoints := 0

	// Check 1: All have descriptions (5 pts)
	{
		withDesc := 0
		if ok {
			for _, slot := range slots {
				if slot.Description() != "" {
					withDesc++
				}
			}
		}
		pts := proportionalScore(withDesc, total, 5)
		msg := ""
		if pts < 5 && total > 0 {
			msg = "add descriptions to all slots"
		}
		totalPoints += pts
		findings = append(findings, Finding{Check: "All have descriptions", Points: pts, Max: 5, Message: msg})
	}

	// Check 2: Content type guidelines (5 pts)
	{
		contentTypeWords := []string{"element", "content", "text", "html", "node", "child", "children", "inline", "block"}
		withGuidelines := 0
		if ok {
			for _, slot := range slots {
				lower := strings.ToLower(slot.Description())
				for _, w := range contentTypeWords {
					if strings.Contains(lower, w) {
						withGuidelines++
						break
					}
				}
			}
		}
		pts := proportionalScore(withGuidelines, total, 5)
		msg := ""
		if pts < 5 && total > 0 {
			msg = "describe expected content types for slots (e.g., 'inline text', 'block elements')"
		}
		totalPoints += pts
		findings = append(findings, Finding{Check: "Content type guidelines", Points: pts, Max: 5, Message: msg})
	}

	// Check 3: A11y considerations (5 pts)
	{
		a11yWords := []string{"accessibility", "a11y", "aria", "screen reader", "wcag", "role", "label"}
		withA11y := 0
		if ok {
			for _, slot := range slots {
				lower := strings.ToLower(slot.Description())
				for _, w := range a11yWords {
					if strings.Contains(lower, w) {
						withA11y++
						break
					}
				}
			}
		}
		pts := proportionalScore(withA11y, total, 5)
		msg := ""
		if pts < 5 && total > 0 {
			msg = "mention accessibility considerations for slot content"
		}
		totalPoints += pts
		findings = append(findings, Finding{Check: "Accessibility considerations", Points: pts, Max: 5, Message: msg})
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
