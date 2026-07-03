/*
Copyright © 2026 Benny Powers <web@bennypowers.com>

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
package breaking

import (
	"cmp"
	"encoding/json"
	"fmt"
	"slices"

	M "bennypowers.dev/cem/manifest"
)

type Severity int

const (
	Safe Severity = iota
	Dangerous
	Breaking
)

func (s Severity) String() string {
	switch s {
	case Breaking:
		return "breaking"
	case Dangerous:
		return "dangerous"
	case Safe:
		return "safe"
	default:
		return "unknown"
	}
}

func (s Severity) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}

func (s *Severity) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	switch str {
	case "breaking":
		*s = Breaking
	case "dangerous":
		*s = Dangerous
	case "safe":
		*s = Safe
	default:
		return fmt.Errorf("unknown severity: %s", str)
	}
	return nil
}

type Change struct {
	Rule     string   `json:"rule"`
	Severity Severity `json:"severity"`
	Element  string   `json:"element"`
	Subject  string   `json:"subject,omitempty"`
	Message  string   `json:"message"`
}

type Result struct {
	Changes   []Change `json:"changes"`
	Breaking  int      `json:"breaking"`
	Dangerous int      `json:"dangerous"`
	Safe      int      `json:"safe"`
}

type Options struct {
	Disable []string
}

func Compare(base, head *M.Package, opts Options) *Result {
	disabled := make(map[string]bool, len(opts.Disable))
	for _, id := range opts.Disable {
		disabled[id] = true
	}

	rules := AllRules()

	baseElements := indexElements(base)
	headElements := indexElements(head)

	var changes []Change

	for _, rule := range rules {
		if disabled[rule.ID()] {
			continue
		}
		changes = append(changes, rule.Check(baseElements, headElements)...)
	}

	slices.SortFunc(changes, func(a, b Change) int {
		if c := cmp.Compare(b.Severity, a.Severity); c != 0 {
			return c
		}
		if c := cmp.Compare(a.Element, b.Element); c != 0 {
			return c
		}
		if c := cmp.Compare(a.Rule, b.Rule); c != 0 {
			return c
		}
		return cmp.Compare(a.Subject, b.Subject)
	})

	result := &Result{Changes: changes}
	for _, c := range changes {
		switch c.Severity {
		case Breaking:
			result.Breaking++
		case Dangerous:
			result.Dangerous++
		case Safe:
			result.Safe++
		}
	}

	return result
}

func indexElements(pkg *M.Package) map[string]*M.CustomElementDeclaration {
	elements := make(map[string]*M.CustomElementDeclaration)
	if pkg == nil {
		return elements
	}
	for i := range pkg.Modules {
		for _, decl := range pkg.Modules[i].Declarations {
			if ced, ok := decl.(*M.CustomElementDeclaration); ok && ced.TagName != "" {
				elements[ced.TagName] = ced
			}
		}
	}
	return elements
}
