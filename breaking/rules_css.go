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

// CSS Custom Properties

type cssPropertyRemovedRule struct{}

func (*cssPropertyRemovedRule) ID() string { return "css-custom-property-removed" }

func (*cssPropertyRemovedRule) Check(base, head map[string]*M.CustomElementDeclaration) []Change {
	var changes []Change
	for tag, baseEl := range base {
		headEl, ok := head[tag]
		if !ok {
			continue
		}
		headProps := indexByName(headEl.OwnCssProperties(), cssPropertyName)
		for _, prop := range baseEl.OwnCssProperties() {
			if _, ok := headProps[prop.Name]; !ok {
				changes = append(changes, Change{
					Rule:     "css-custom-property-removed",
					Severity: Breaking,
					Element:  tag,
					Subject:  prop.Name,
					Message:  fmt.Sprintf("CSS custom property %q removed from <%s>", prop.Name, tag),
				})
			}
		}
	}
	return changes
}

type cssPropertyAddedRule struct{}

func (*cssPropertyAddedRule) ID() string { return "css-custom-property-added" }

func (*cssPropertyAddedRule) Check(base, head map[string]*M.CustomElementDeclaration) []Change {
	var changes []Change
	for tag, headEl := range head {
		baseEl, ok := base[tag]
		if !ok {
			continue
		}
		baseProps := indexByName(baseEl.OwnCssProperties(), cssPropertyName)
		for _, prop := range headEl.OwnCssProperties() {
			if _, ok := baseProps[prop.Name]; !ok {
				changes = append(changes, Change{
					Rule:     "css-custom-property-added",
					Severity: Safe,
					Element:  tag,
					Subject:  prop.Name,
					Message:  fmt.Sprintf("CSS custom property %q added to <%s>", prop.Name, tag),
				})
			}
		}
	}
	return changes
}

type cssPropertyDefaultChangedRule struct{}

func (*cssPropertyDefaultChangedRule) ID() string { return "css-custom-property-default-changed" }

func (*cssPropertyDefaultChangedRule) Check(base, head map[string]*M.CustomElementDeclaration) []Change {
	var changes []Change
	for tag, baseEl := range base {
		headEl, ok := head[tag]
		if !ok {
			continue
		}
		headProps := indexByName(headEl.OwnCssProperties(), cssPropertyName)
		for _, baseProp := range baseEl.OwnCssProperties() {
			headProp, ok := headProps[baseProp.Name]
			if !ok {
				continue
			}
			if baseProp.Default != "" && headProp.Default != "" && baseProp.Default != headProp.Default {
				changes = append(changes, Change{
					Rule:     "css-custom-property-default-changed",
					Severity: Dangerous,
					Element:  tag,
					Subject:  baseProp.Name,
					Message:  fmt.Sprintf("CSS custom property %q default changed from %s to %s in <%s>", baseProp.Name, baseProp.Default, headProp.Default, tag),
				})
			}
		}
	}
	return changes
}

// CSS Parts

type cssPartRemovedRule struct{}

func (*cssPartRemovedRule) ID() string { return "css-part-removed" }

func (*cssPartRemovedRule) Check(base, head map[string]*M.CustomElementDeclaration) []Change {
	var changes []Change
	for tag, baseEl := range base {
		headEl, ok := head[tag]
		if !ok {
			continue
		}
		headParts := indexByName(headEl.OwnCssParts(), cssPartName)
		for _, part := range baseEl.OwnCssParts() {
			if _, ok := headParts[part.Name]; !ok {
				changes = append(changes, Change{
					Rule:     "css-part-removed",
					Severity: Breaking,
					Element:  tag,
					Subject:  part.Name,
					Message:  fmt.Sprintf("CSS part %q removed from <%s>", part.Name, tag),
				})
			}
		}
	}
	return changes
}

type cssPartAddedRule struct{}

func (*cssPartAddedRule) ID() string { return "css-part-added" }

func (*cssPartAddedRule) Check(base, head map[string]*M.CustomElementDeclaration) []Change {
	var changes []Change
	for tag, headEl := range head {
		baseEl, ok := base[tag]
		if !ok {
			continue
		}
		baseParts := indexByName(baseEl.OwnCssParts(), cssPartName)
		for _, part := range headEl.OwnCssParts() {
			if _, ok := baseParts[part.Name]; !ok {
				changes = append(changes, Change{
					Rule:     "css-part-added",
					Severity: Safe,
					Element:  tag,
					Subject:  part.Name,
					Message:  fmt.Sprintf("CSS part %q added to <%s>", part.Name, tag),
				})
			}
		}
	}
	return changes
}

// CSS States

type cssStateRemovedRule struct{}

func (*cssStateRemovedRule) ID() string { return "css-state-removed" }

func (*cssStateRemovedRule) Check(base, head map[string]*M.CustomElementDeclaration) []Change {
	var changes []Change
	for tag, baseEl := range base {
		headEl, ok := head[tag]
		if !ok {
			continue
		}
		headStates := indexByName(headEl.OwnCssStates(), cssStateName)
		for _, state := range baseEl.OwnCssStates() {
			if _, ok := headStates[state.Name]; !ok {
				changes = append(changes, Change{
					Rule:     "css-state-removed",
					Severity: Breaking,
					Element:  tag,
					Subject:  state.Name,
					Message:  fmt.Sprintf("CSS state %q removed from <%s>", state.Name, tag),
				})
			}
		}
	}
	return changes
}

type cssStateAddedRule struct{}

func (*cssStateAddedRule) ID() string { return "css-state-added" }

func (*cssStateAddedRule) Check(base, head map[string]*M.CustomElementDeclaration) []Change {
	var changes []Change
	for tag, headEl := range head {
		baseEl, ok := base[tag]
		if !ok {
			continue
		}
		baseStates := indexByName(baseEl.OwnCssStates(), cssStateName)
		for _, state := range headEl.OwnCssStates() {
			if _, ok := baseStates[state.Name]; !ok {
				changes = append(changes, Change{
					Rule:     "css-state-added",
					Severity: Safe,
					Element:  tag,
					Subject:  state.Name,
					Message:  fmt.Sprintf("CSS state %q added to <%s>", state.Name, tag),
				})
			}
		}
	}
	return changes
}
