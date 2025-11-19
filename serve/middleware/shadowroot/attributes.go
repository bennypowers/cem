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

package shadowroot

import (
	"strings"
	"unicode"

	"golang.org/x/net/html"
)

// attributesToTemplateData converts HTML attributes to template data
// Provides both raw map access (.Attributes["aria-label"]) and
// camelCased properties (.AriaLabel)
func attributesToTemplateData(attrs []html.Attribute) map[string]interface{} {
	data := make(map[string]interface{})
	attrMap := make(map[string]string)

	for _, attr := range attrs {
		// Store raw attribute (boolean attrs get empty string)
		value := attr.Val
		if value == "" {
			value = "" // Explicit: disabled → disabled=""
		}

		attrMap[attr.Key] = value

		// Also add camelCased version for convenient template access
		// aria-label → AriaLabel
		// data-foo-bar → DataFooBar
		camelKey := dashToCamelCase(attr.Key)
		data[camelKey] = value
	}

	// Make raw attributes available as .Attributes map
	data["Attributes"] = attrMap

	return data
}

// dashToCamelCase converts dash-case to PascalCase
// Examples:
//   - "aria-label" → "AriaLabel"
//   - "data-foo-bar" → "DataFooBar"
//   - "disabled" → "Disabled"
func dashToCamelCase(s string) string {
	if s == "" {
		return ""
	}

	var result strings.Builder
	capitalizeNext := true

	for _, ch := range s {
		if ch == '-' {
			capitalizeNext = true
			continue
		}

		if capitalizeNext {
			result.WriteRune(unicode.ToUpper(ch))
			capitalizeNext = false
		} else {
			result.WriteRune(ch)
		}
	}

	return result.String()
}
