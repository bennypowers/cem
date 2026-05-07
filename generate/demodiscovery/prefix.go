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
package demodiscovery

import (
	"maps"
	"strings"
)

// DetectTagPrefix finds the common first-segment prefix shared by all tag names.
// Returns "prefix-" if all tags start with the same segment before the first hyphen,
// or "" if fewer than 2 tags or no common prefix.
func DetectTagPrefix(tagNames []string) string {
	if len(tagNames) < 2 {
		return ""
	}
	first := tagNames[0]
	idx := strings.Index(first, "-")
	if idx < 0 {
		return ""
	}
	prefix := first[:idx+1]
	for _, tag := range tagNames[1:] {
		if !strings.HasPrefix(tag, prefix) {
			return ""
		}
	}
	return prefix
}

// AutoDeriveAliases returns a merged alias map: explicit aliases from existingAliases
// plus auto-derived aliases for any tag name that lacks one.
// Auto-derivation strips the common tag prefix detected from tagNames.
func AutoDeriveAliases(tagNames []string, existingAliases map[string]string) map[string]string {
	result := make(map[string]string, len(existingAliases))
	maps.Copy(result, existingAliases)
	prefix := DetectTagPrefix(tagNames)
	if prefix == "" {
		return result
	}
	for _, tag := range tagNames {
		if _, ok := result[tag]; ok {
			continue
		}
		result[tag] = strings.TrimPrefix(tag, prefix)
	}
	return result
}
