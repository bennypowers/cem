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

// Package validations provides shared validation utilities and data for HTML attributes,
// custom elements, and other validation concerns across LSP and MCP components.
package validations

import (
	_ "embed"
	"encoding/json"
	"strings"
)

//go:embed data/global_attributes.json
var globalAttributesJSON []byte

// MDNCompatData represents the structure of MDN browser compatibility data
type MDNCompatData struct {
	HTML struct {
		GlobalAttributes map[string]interface{} `json:"global_attributes"`
	} `json:"html"`
}

var globalAttributesCache map[string]bool

// Initialize global attributes from embedded MDN data
func init() {
	globalAttributesCache = make(map[string]bool)

	var mdnData MDNCompatData
	if err := json.Unmarshal(globalAttributesJSON, &mdnData); err == nil {
		for attrName := range mdnData.HTML.GlobalAttributes {
			// Convert data_attributes to data-* pattern
			if attrName == "data_attributes" {
				// We'll handle data-* pattern separately in IsGlobalAttribute
				continue
			}
			globalAttributesCache[attrName] = true
		}
	}
}

// IsGlobalAttribute checks if an attribute name is a global HTML attribute.
// This includes:
// - All standard global HTML attributes from MDN data (id, class, slot, style, etc.)
// - data-* attributes (always valid)
// - aria-* attributes (always valid)
// - Event handler attributes starting with "on" (always valid)
func IsGlobalAttribute(name string) bool {
	nameLower := strings.ToLower(name)

	// Check if it's in our global attributes list
	if globalAttributesCache[nameLower] {
		return true
	}

	// Check for data-* attributes (always valid)
	if strings.HasPrefix(nameLower, "data-") {
		return true
	}

	// Check for aria-* attributes (always valid)
	if strings.HasPrefix(nameLower, "aria-") {
		return true
	}

	// Check for event handler attributes starting with "on"
	if strings.HasPrefix(nameLower, "on") && len(nameLower) > 2 {
		return true
	}

	return false
}

// GetGlobalAttributes returns a copy of all loaded global attributes.
// This is useful for debugging or when you need the full list.
func GetGlobalAttributes() map[string]bool {
	result := make(map[string]bool)
	for k, v := range globalAttributesCache {
		result[k] = v
	}
	return result
}

// GetGlobalAttributeCount returns the number of loaded global attributes.
func GetGlobalAttributeCount() int {
	return len(globalAttributesCache)
}
