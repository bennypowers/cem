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

type attributeRemovedRule struct{}

func (*attributeRemovedRule) ID() string { return "attribute-removed" }

func (*attributeRemovedRule) Check(base, head map[string]*M.CustomElementDeclaration) []Change {
	var changes []Change
	for tag, baseEl := range base {
		headEl, ok := head[tag]
		if !ok {
			continue
		}
		headAttrs := indexByName(headEl.OwnAttributes(), attrName)
		for _, attr := range baseEl.OwnAttributes() {
			if _, ok := headAttrs[attr.Name]; !ok {
				changes = append(changes, Change{
					Rule:     "attribute-removed",
					Severity: Breaking,
					Element:  tag,
					Subject:  attr.Name,
					Message:  fmt.Sprintf("attribute %q removed from <%s>", attr.Name, tag),
				})
			}
		}
	}
	return changes
}

type attributeAddedRule struct{}

func (*attributeAddedRule) ID() string { return "attribute-added" }

func (*attributeAddedRule) Check(base, head map[string]*M.CustomElementDeclaration) []Change {
	var changes []Change
	for tag, headEl := range head {
		baseEl, ok := base[tag]
		if !ok {
			continue
		}
		baseAttrs := indexByName(baseEl.OwnAttributes(), attrName)
		for _, attr := range headEl.OwnAttributes() {
			if _, ok := baseAttrs[attr.Name]; !ok {
				changes = append(changes, Change{
					Rule:     "attribute-added",
					Severity: Safe,
					Element:  tag,
					Subject:  attr.Name,
					Message:  fmt.Sprintf("attribute %q added to <%s>", attr.Name, tag),
				})
			}
		}
	}
	return changes
}

type attributeTypeChangedRule struct{}

func (*attributeTypeChangedRule) ID() string { return "attribute-type-changed" }

func (*attributeTypeChangedRule) Check(base, head map[string]*M.CustomElementDeclaration) []Change {
	var changes []Change
	for tag, baseEl := range base {
		headEl, ok := head[tag]
		if !ok {
			continue
		}
		headAttrs := indexByName(headEl.OwnAttributes(), attrName)
		for _, baseAttr := range baseEl.OwnAttributes() {
			headAttr, ok := headAttrs[baseAttr.Name]
			if !ok {
				continue
			}
			baseType := typeText(baseAttr.Type)
			headType := typeText(headAttr.Type)
			if baseType != "" && headType != "" && baseType != headType {
				changes = append(changes, Change{
					Rule:     "attribute-type-changed",
					Severity: Breaking,
					Element:  tag,
					Subject:  baseAttr.Name,
					Message:  fmt.Sprintf("attribute %q type changed from %s to %s in <%s>", baseAttr.Name, baseType, headType, tag),
				})
			}
		}
	}
	return changes
}

type attributeDefaultChangedRule struct{}

func (*attributeDefaultChangedRule) ID() string { return "attribute-default-changed" }

func (*attributeDefaultChangedRule) Check(base, head map[string]*M.CustomElementDeclaration) []Change {
	var changes []Change
	for tag, baseEl := range base {
		headEl, ok := head[tag]
		if !ok {
			continue
		}
		headAttrs := indexByName(headEl.OwnAttributes(), attrName)
		for _, baseAttr := range baseEl.OwnAttributes() {
			headAttr, ok := headAttrs[baseAttr.Name]
			if !ok {
				continue
			}
			if baseAttr.Default != "" && headAttr.Default != "" && baseAttr.Default != headAttr.Default {
				changes = append(changes, Change{
					Rule:     "attribute-default-changed",
					Severity: Dangerous,
					Element:  tag,
					Subject:  baseAttr.Name,
					Message:  fmt.Sprintf("attribute %q default changed from %s to %s in <%s>", baseAttr.Name, baseAttr.Default, headAttr.Default, tag),
				})
			}
		}
	}
	return changes
}
