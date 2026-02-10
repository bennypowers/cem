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

// DemoDocRule scores demo documentation quality (10 pts max).
type DemoDocRule struct{}

func (r *DemoDocRule) ID() string       { return "demos" }
func (r *DemoDocRule) Category() string { return "Demos" }
func (r *DemoDocRule) MaxPoints() int   { return 10 }

func (r *DemoDocRule) Evaluate(ctx *HealthContext) CategoryScore {
	demos, ok := ctx.Declaration.Demos()

	var findings []Finding
	totalPoints := 0

	// Check 1: Has at least one demo (5 pts)
	{
		pts := 0
		msg := ""
		hasDemo := false
		if ok {
			for _, demo := range demos {
				if demo.URL() != "" {
					hasDemo = true
					break
				}
			}
		}
		if hasDemo {
			pts = 5
		} else {
			msg = "add at least one demo with a URL"
		}
		totalPoints += pts
		findings = append(findings, Finding{Check: "Has demo", Points: pts, Max: 5, Message: msg})
	}

	// Check 2: Demo has description (5 pts)
	{
		pts := 0
		msg := ""
		hasDescribedDemo := false
		if ok {
			for _, demo := range demos {
				if demo.Description() != "" {
					hasDescribedDemo = true
					break
				}
			}
		}
		if hasDescribedDemo {
			pts = 5
		} else if ok && len(demos) > 0 {
			msg = "add descriptions to demos"
		} else {
			msg = "add demos with descriptions"
		}
		totalPoints += pts
		findings = append(findings, Finding{Check: "Demo has description", Points: pts, Max: 5, Message: msg})
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
