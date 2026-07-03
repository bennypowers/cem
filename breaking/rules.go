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

import M "bennypowers.dev/cem/manifest"

type BreakingRule interface {
	ID() string
	Check(base, head map[string]*M.CustomElementDeclaration) []Change
}

func AllRules() []BreakingRule {
	return []BreakingRule{
		&elementRemovedRule{},
		&elementAddedRule{},
		&attributeRemovedRule{},
		&attributeAddedRule{},
		&attributeTypeChangedRule{},
		&attributeDefaultChangedRule{},
		&slotRemovedRule{},
		&slotAddedRule{},
		&cssPropertyRemovedRule{},
		&cssPropertyAddedRule{},
		&cssPropertyDefaultChangedRule{},
		&cssPartRemovedRule{},
		&cssPartAddedRule{},
		&cssStateRemovedRule{},
		&cssStateAddedRule{},
		&eventRemovedRule{},
		&eventAddedRule{},
		&eventTypeChangedRule{},
		&methodRemovedRule{},
		&methodAddedRule{},
		&methodReturnTypeChangedRule{},
		&methodParameterChangedRule{},
		&fieldRemovedRule{},
		&fieldTypeChangedRule{},
	}
}
