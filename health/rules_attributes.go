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

// AttributeDocRule scores attribute documentation quality (20 pts max).
type AttributeDocRule struct{}

func (r *AttributeDocRule) ID() string        { return "attributes" }
func (r *AttributeDocRule) Category() string   { return "Attribute documentation" }
func (r *AttributeDocRule) MaxPoints() int     { return 20 }

func (r *AttributeDocRule) Evaluate(ctx *HealthContext) CategoryScore {
	attrs, ok := ctx.Declaration.Attributes()
	total := 0
	if ok {
		total = len(attrs)
	}

	var findings []Finding
	totalPoints := 0

	// Check 1: All have descriptions (10 pts)
	{
		withDesc := 0
		if ok {
			for _, attr := range attrs {
				if attr.Description() != "" {
					withDesc++
				}
			}
		}
		pts := proportionalScore(withDesc, total, 10)
		msg := ""
		if pts < 10 && total > 0 {
			msg = "add descriptions to all attributes"
		}
		totalPoints += pts
		findings = append(findings, Finding{Check: "All have descriptions", Points: pts, Max: 10, Message: msg})
	}

	// Check 2: Descriptions explain purpose (5 pts)
	// Descriptions > 20 chars that aren't just type names
	{
		meaningful := 0
		if ok {
			for _, attr := range attrs {
				desc := attr.Description()
				typeName := attr.Type()
				if len(desc) > 20 && !strings.EqualFold(strings.TrimSpace(desc), strings.TrimSpace(typeName)) {
					meaningful++
				}
			}
		}
		pts := proportionalScore(meaningful, total, 5)
		msg := ""
		if pts < 5 && total > 0 {
			msg = "write meaningful descriptions (>20 chars) that explain purpose, not just type"
		}
		totalPoints += pts
		findings = append(findings, Finding{Check: "Descriptions explain purpose", Points: pts, Max: 5, Message: msg})
	}

	// Check 3: Constraints documented (5 pts)
	// Have default values or type info
	{
		constrained := 0
		if ok {
			for _, attr := range attrs {
				if attr.Default() != "" || attr.Type() != "" {
					constrained++
				}
			}
		}
		pts := proportionalScore(constrained, total, 5)
		msg := ""
		if pts < 5 && total > 0 {
			msg = "document type information or default values for attributes"
		}
		totalPoints += pts
		findings = append(findings, Finding{Check: "Constraints documented", Points: pts, Max: 5, Message: msg})
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
