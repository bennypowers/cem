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

type eventRemovedRule struct{}

func (*eventRemovedRule) ID() string { return "event-removed" }

func (*eventRemovedRule) Check(base, head map[string]*M.CustomElementDeclaration) []Change {
	var changes []Change
	for tag, baseEl := range base {
		headEl, ok := head[tag]
		if !ok {
			continue
		}
		headEvents := indexByName(headEl.OwnEvents(), eventName)
		for _, event := range baseEl.OwnEvents() {
			if _, ok := headEvents[event.Name]; !ok {
				changes = append(changes, Change{
					Rule:     "event-removed",
					Severity: Breaking,
					Element:  tag,
					Subject:  event.Name,
					Message:  fmt.Sprintf("event %q removed from <%s>", event.Name, tag),
				})
			}
		}
	}
	return changes
}

type eventAddedRule struct{}

func (*eventAddedRule) ID() string { return "event-added" }

func (*eventAddedRule) Check(base, head map[string]*M.CustomElementDeclaration) []Change {
	var changes []Change
	for tag, headEl := range head {
		baseEl, ok := base[tag]
		if !ok {
			continue
		}
		baseEvents := indexByName(baseEl.OwnEvents(), eventName)
		for _, event := range headEl.OwnEvents() {
			if _, ok := baseEvents[event.Name]; !ok {
				changes = append(changes, Change{
					Rule:     "event-added",
					Severity: Safe,
					Element:  tag,
					Subject:  event.Name,
					Message:  fmt.Sprintf("event %q added to <%s>", event.Name, tag),
				})
			}
		}
	}
	return changes
}

type eventTypeChangedRule struct{}

func (*eventTypeChangedRule) ID() string { return "event-type-changed" }

func (*eventTypeChangedRule) Check(base, head map[string]*M.CustomElementDeclaration) []Change {
	var changes []Change
	for tag, baseEl := range base {
		headEl, ok := head[tag]
		if !ok {
			continue
		}
		headEvents := indexByName(headEl.OwnEvents(), eventName)
		for _, baseEvent := range baseEl.OwnEvents() {
			headEvent, ok := headEvents[baseEvent.Name]
			if !ok {
				continue
			}
			baseType := typeText(baseEvent.Type)
			headType := typeText(headEvent.Type)
			if baseType != "" && headType != "" && baseType != headType {
				changes = append(changes, Change{
					Rule:     "event-type-changed",
					Severity: Breaking,
					Element:  tag,
					Subject:  baseEvent.Name,
					Message:  fmt.Sprintf("event %q type changed from %s to %s in <%s>", baseEvent.Name, baseType, headType, tag),
				})
			}
		}
	}
	return changes
}
