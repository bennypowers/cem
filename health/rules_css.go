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

// CSSDocRule scores CSS property and part documentation quality (15 pts max).
type CSSDocRule struct{}

func (r *CSSDocRule) ID() string        { return "css" }
func (r *CSSDocRule) Category() string   { return "CSS documentation" }
func (r *CSSDocRule) MaxPoints() int     { return 15 }

func (r *CSSDocRule) Evaluate(ctx *HealthContext) CategoryScore {
	cssProps, propsOk := ctx.Declaration.CSSProperties()
	cssParts, partsOk := ctx.Declaration.CSSParts()

	propsTotal := 0
	if propsOk {
		propsTotal = len(cssProps)
	}
	partsTotal := 0
	if partsOk {
		partsTotal = len(cssParts)
	}

	var findings []Finding
	totalPoints := 0

	// Check 1: CSS properties documented (5 pts)
	{
		withDesc := 0
		if propsOk {
			for _, prop := range cssProps {
				if prop.Description() != "" {
					withDesc++
				}
			}
		}
		pts := proportionalScore(withDesc, propsTotal, 5)
		msg := ""
		if pts < 5 && propsTotal > 0 {
			msg = "add descriptions to all CSS custom properties"
		}
		totalPoints += pts
		findings = append(findings, Finding{Check: "CSS properties documented", Points: pts, Max: 5, Message: msg})
	}

	// Check 2: CSS parts have descriptions (5 pts)
	{
		withDesc := 0
		if partsOk {
			for _, part := range cssParts {
				if part.Description() != "" {
					withDesc++
				}
			}
		}
		pts := proportionalScore(withDesc, partsTotal, 5)
		msg := ""
		if pts < 5 && partsTotal > 0 {
			msg = "add descriptions to all CSS parts"
		}
		totalPoints += pts
		findings = append(findings, Finding{Check: "CSS parts documented", Points: pts, Max: 5, Message: msg})
	}

	// Check 3: Design system notes (5 pts)
	{
		designWords := []string{"design", "token", "theme", "system"}
		combined := 0
		combinedTotal := propsTotal + partsTotal
		if propsOk {
			for _, prop := range cssProps {
				lower := strings.ToLower(prop.Description())
				for _, w := range designWords {
					if strings.Contains(lower, w) {
						combined++
						break
					}
				}
			}
		}
		if partsOk {
			for _, part := range cssParts {
				lower := strings.ToLower(part.Description())
				for _, w := range designWords {
					if strings.Contains(lower, w) {
						combined++
						break
					}
				}
			}
		}
		pts := proportionalScore(combined, combinedTotal, 5)
		msg := ""
		if pts < 5 && combinedTotal > 0 {
			msg = "reference design tokens or theme considerations in CSS descriptions"
		}
		totalPoints += pts
		findings = append(findings, Finding{Check: "Design system notes", Points: pts, Max: 5, Message: msg})
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
