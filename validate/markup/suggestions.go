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
package markup

import (
	"bennypowers.dev/cem/manifest"
	"github.com/agext/levenshtein"
)

// suggestAttributeCorrection returns the closest manifest attribute name when
// the typed name is plausibly a typo. The threshold matches the MCP tool's
// historical behavior: distance at most half the typed name's length.
func suggestAttributeCorrection(typed string, attributes []manifest.Attribute) string {
	if len(attributes) == 0 {
		return ""
	}

	bestMatch := ""
	minDistance := len(typed) + 1
	for _, attr := range attributes {
		distance := levenshtein.Distance(typed, attr.Name, nil)
		if distance < minDistance && distance <= len(typed)/2 {
			bestMatch = attr.Name
			minDistance = distance
		}
	}
	return bestMatch
}
