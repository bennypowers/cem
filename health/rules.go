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
	"bennypowers.dev/cem/validate"
)

// HealthRule evaluates a scoring category for a declaration.
type HealthRule interface {
	Evaluate(ctx *HealthContext) CategoryScore
	ID() string
	Category() string
	MaxPoints() int
}

// CategoryScore holds the result of evaluating a single scoring category.
type CategoryScore struct {
	ID        string    `json:"id"`
	Category  string    `json:"category"`
	Points    int       `json:"points"`
	MaxPoints int       `json:"maxPoints"`
	Status    string    `json:"status"` // "pass", "warn", "fail"
	Findings  []Finding `json:"findings"`
}

// Finding is an individual check result within a category.
type Finding struct {
	Check   string `json:"check"`
	Points  int    `json:"points"`
	Max     int    `json:"max"`
	Message string `json:"message,omitempty"`
}

// HealthContext provides context for health rule evaluation.
type HealthContext struct {
	Navigator   *validate.ManifestNavigator
	Module      validate.RawModule
	Declaration validate.RawDeclaration
	ModulePath  string
	DeclName    string
	DeclKind    string
}

// status computes the status string based on points and max.
func status(points, max int) string {
	if max == 0 {
		return "pass"
	}
	pct := float64(points) / float64(max) * 100
	if pct >= 80 {
		return "pass"
	}
	if pct >= 40 {
		return "warn"
	}
	return "fail"
}

// proportionalScore calculates proportional points: (count / total) * max.
// Returns max if total is 0 (vacuous truth).
func proportionalScore(count, total, max int) int {
	if total == 0 {
		return max
	}
	return min(int(float64(count)/float64(total)*float64(max)), max)
}

// defaultRules returns the standard set of health rules.
func defaultRules() []HealthRule {
	return []HealthRule{
		&DescriptionRule{},
		&AttributeDocRule{},
		&SlotDocRule{},
		&CSSDocRule{},
		&EventDocRule{},
		&DemoDocRule{},
	}
}
