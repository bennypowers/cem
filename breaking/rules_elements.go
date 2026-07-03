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
	"fmt"

	M "bennypowers.dev/cem/manifest"
)

type elementRemovedRule struct{}

func (*elementRemovedRule) ID() string { return "element-removed" }

func (*elementRemovedRule) Check(base, head map[string]*M.CustomElementDeclaration) []Change {
	var changes []Change
	for tag := range base {
		if _, ok := head[tag]; !ok {
			changes = append(changes, Change{
				Rule:     "element-removed",
				Severity: Breaking,
				Element:  tag,
				Message:  fmt.Sprintf("element <%s> removed", tag),
			})
		}
	}
	return changes
}

type elementAddedRule struct{}

func (*elementAddedRule) ID() string { return "element-added" }

func (*elementAddedRule) Check(base, head map[string]*M.CustomElementDeclaration) []Change {
	var changes []Change
	for tag := range head {
		if _, ok := base[tag]; !ok {
			changes = append(changes, Change{
				Rule:     "element-added",
				Severity: Safe,
				Element:  tag,
				Message:  fmt.Sprintf("element <%s> added", tag),
			})
		}
	}
	return changes
}
