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

func indexByName[T any](items []T, name func(T) string) map[string]T {
	m := make(map[string]T, len(items))
	for _, item := range items {
		m[name(item)] = item
	}
	return m
}

func attrName(a M.Attribute) string           { return a.Name }
func slotName(s M.Slot) string                { return s.Name }
func eventName(e M.Event) string              { return e.Name }
func cssPropertyName(p M.CssCustomProperty) string { return p.Name }
func cssPartName(p M.CssPart) string          { return p.Name }
func cssStateName(s M.CssCustomState) string  { return s.Name }

func typeText(t *M.Type) string {
	if t == nil {
		return ""
	}
	return t.Text
}

func isPublic(privacy M.Privacy) bool {
	return privacy == "" || privacy == M.Public
}
