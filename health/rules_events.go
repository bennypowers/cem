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

// EventDocRule scores event documentation quality (15 pts max).
type EventDocRule struct{}

func (r *EventDocRule) ID() string        { return "events" }
func (r *EventDocRule) Category() string   { return "Event documentation" }
func (r *EventDocRule) MaxPoints() int     { return 15 }

func (r *EventDocRule) Evaluate(ctx *HealthContext) CategoryScore {
	events, ok := ctx.Declaration.Events()
	total := 0
	if ok {
		total = len(events)
	}

	var findings []Finding
	totalPoints := 0

	// Check 1: All have descriptions (5 pts)
	{
		withDesc := 0
		if ok {
			for _, event := range events {
				if event.Description() != "" {
					withDesc++
				}
			}
		}
		pts := proportionalScore(withDesc, total, 5)
		msg := ""
		if pts < 5 && total > 0 {
			msg = "add descriptions to all events"
		}
		totalPoints += pts
		findings = append(findings, Finding{Check: "All have descriptions", Points: pts, Max: 5, Message: msg})
	}

	// Check 2: Trigger conditions (5 pts)
	{
		triggerWords := []string{"when", "fired", "trigger", "dispatched", "emitted", "after", "before"}
		withTrigger := 0
		if ok {
			for _, event := range events {
				lower := strings.ToLower(event.Description())
				for _, w := range triggerWords {
					if strings.Contains(lower, w) {
						withTrigger++
						break
					}
				}
			}
		}
		pts := proportionalScore(withTrigger, total, 5)
		msg := ""
		if pts < 5 && total > 0 {
			msg = "describe when events are fired (e.g., 'Fired when the user clicks...')"
		}
		totalPoints += pts
		findings = append(findings, Finding{Check: "Trigger conditions", Points: pts, Max: 5, Message: msg})
	}

	// Check 3: Event detail described (5 pts)
	{
		detailWords := []string{"detail", "data", "contains", "includes", "payload", "property", "value"}
		withDetail := 0
		if ok {
			for _, event := range events {
				lower := strings.ToLower(event.Description())
				for _, w := range detailWords {
					if strings.Contains(lower, w) {
						withDetail++
						break
					}
				}
			}
		}
		pts := proportionalScore(withDetail, total, 5)
		msg := ""
		if pts < 5 && total > 0 {
			msg = "describe event detail/data shape"
		}
		totalPoints += pts
		findings = append(findings, Finding{Check: "Event detail described", Points: pts, Max: 5, Message: msg})
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
