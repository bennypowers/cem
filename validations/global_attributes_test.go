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
package validations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsGlobalAttribute(t *testing.T) {
	tests := []struct {
		name      string
		attribute string
		isGlobal  bool
		reason    string
	}{
		// Standard global attributes
		{"id attribute", "id", true, "id is a standard global attribute"},
		{"class attribute", "class", true, "class is a standard global attribute"},
		{"slot attribute", "slot", true, "slot is a standard global attribute"},
		{"style attribute", "style", true, "style is a standard global attribute"},
		{"title attribute", "title", true, "title is a standard global attribute"},
		{"lang attribute", "lang", true, "lang is a standard global attribute"},
		{"dir attribute", "dir", true, "dir is a standard global attribute"},
		{"hidden attribute", "hidden", true, "hidden is a standard global attribute"},
		{"tabindex attribute", "tabindex", true, "tabindex is a standard global attribute"},
		{"contenteditable attribute", "contenteditable", true, "contenteditable is a standard global attribute"},
		{"draggable attribute", "draggable", true, "draggable is a standard global attribute"},
		{"spellcheck attribute", "spellcheck", true, "spellcheck is a standard global attribute"},
		{"translate attribute", "translate", true, "translate is a standard global attribute"},

		// Case insensitive
		{"ID uppercase", "ID", true, "should be case insensitive"},
		{"Class mixed case", "Class", true, "should be case insensitive"},

		// Data attributes
		{"data attribute", "data-testid", true, "data-* attributes are always global"},
		{"data attribute complex", "data-my-custom-value", true, "data-* attributes are always global"},
		{"data attribute uppercase", "DATA-TEST", true, "data-* attributes should be case insensitive"},

		// ARIA attributes
		{"aria attribute", "aria-label", true, "aria-* attributes are always global"},
		{"aria attribute complex", "aria-describedby", true, "aria-* attributes are always global"},
		{"aria attribute uppercase", "ARIA-HIDDEN", true, "aria-* attributes should be case insensitive"},

		// Event handler attributes
		{"onclick event", "onclick", true, "on* event handlers are global"},
		{"onmouseover event", "onmouseover", true, "on* event handlers are global"},
		{"onload event", "onload", true, "on* event handlers are global"},
		{"event handler uppercase", "ONCLICK", true, "on* event handlers should be case insensitive"},

		// Non-global attributes
		{"custom attribute", "custom-attr", false, "custom attributes are not global"},
		{"element specific", "variant", false, "element-specific attributes are not global"},
		{"short on prefix", "on", false, "bare 'on' is not an event handler"},
		{"not data prefix", "data", false, "bare 'data' is not a data attribute"},
		{"not aria prefix", "aria", false, "bare 'aria' is not an aria attribute"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsGlobalAttribute(test.attribute)
			assert.Equal(t, test.isGlobal, result, "%s: %s", test.attribute, test.reason)
		})
	}
}

func TestGetGlobalAttributes(t *testing.T) {
	attrs := GetGlobalAttributes()

	// Should have loaded attributes
	assert.Greater(t, len(attrs), 0, "Should have loaded global attributes")

	// Should contain expected attributes
	expectedAttrs := []string{"id", "class", "slot", "style", "title", "hidden"}
	for _, attr := range expectedAttrs {
		assert.True(t, attrs[attr], "Should contain global attribute: %s", attr)
	}

	// Should be a copy (not the original map)
	attrs["test"] = true
	attrs2 := GetGlobalAttributes()
	assert.False(t, attrs2["test"], "Should return a copy of the map")
}

func TestGetGlobalAttributeCount(t *testing.T) {
	count := GetGlobalAttributeCount()
	assert.Greater(t, count, 20, "Should have loaded a reasonable number of global attributes")
	t.Logf("Loaded %d global attributes", count)
}
