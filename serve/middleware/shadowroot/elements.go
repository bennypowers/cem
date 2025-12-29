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
	"embed"
	"fmt"

	"bennypowers.dev/cem/set"
)

// buildKnownElements scans the embedded templates filesystem and builds
// a set of element names that have shadow root templates
func buildKnownElements(templatesFS embed.FS) set.Set[string] {
	return buildKnownElementsWithPrefix(templatesFS, "templates/elements")
}

// buildKnownElementsWithPrefix scans the embedded filesystem with a custom prefix
func buildKnownElementsWithPrefix(templatesFS embed.FS, prefix string) set.Set[string] {
	knownElements := set.NewSet[string]()

	// Read the elements directory
	entries, err := templatesFS.ReadDir(prefix)
	if err != nil {
		// If we can't read the directory, return empty set
		return knownElements
	}

	// Scan each subdirectory for element templates
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		elementName := entry.Name()

		// Verify element has an HTML template file
		templatePath := fmt.Sprintf("%s/%s/%s.html", prefix, elementName, elementName)
		if _, err := templatesFS.Open(templatePath); err == nil {
			knownElements.Add(elementName)
		}
	}

	return knownElements
}
