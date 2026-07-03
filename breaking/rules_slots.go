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

type slotRemovedRule struct{}

func (*slotRemovedRule) ID() string { return "slot-removed" }

func (*slotRemovedRule) Check(base, head map[string]*M.CustomElementDeclaration) []Change {
	var changes []Change
	for tag, baseEl := range base {
		headEl, ok := head[tag]
		if !ok {
			continue
		}
		headSlots := indexByName(headEl.OwnSlots(), slotName)
		for _, slot := range baseEl.OwnSlots() {
			if _, ok := headSlots[slot.Name]; !ok {
				name := slot.Name
				if name == "" {
					name = "(default)"
				}
				changes = append(changes, Change{
					Rule:     "slot-removed",
					Severity: Breaking,
					Element:  tag,
					Subject:  slot.Name,
					Message:  fmt.Sprintf("slot %q removed from <%s>", name, tag),
				})
			}
		}
	}
	return changes
}

type slotAddedRule struct{}

func (*slotAddedRule) ID() string { return "slot-added" }

func (*slotAddedRule) Check(base, head map[string]*M.CustomElementDeclaration) []Change {
	var changes []Change
	for tag, headEl := range head {
		baseEl, ok := base[tag]
		if !ok {
			continue
		}
		baseSlots := indexByName(baseEl.OwnSlots(), slotName)
		for _, slot := range headEl.OwnSlots() {
			if _, ok := baseSlots[slot.Name]; !ok {
				name := slot.Name
				if name == "" {
					name = "(default)"
				}
				changes = append(changes, Change{
					Rule:     "slot-added",
					Severity: Safe,
					Element:  tag,
					Subject:  slot.Name,
					Message:  fmt.Sprintf("slot %q added to <%s>", name, tag),
				})
			}
		}
	}
	return changes
}
