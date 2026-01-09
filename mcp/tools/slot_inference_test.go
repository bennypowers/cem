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
package tools

import (
	"testing"

	"bennypowers.dev/cem/manifest"
	"github.com/stretchr/testify/assert"
)

func TestInferRelevantSlots(t *testing.T) {
	tests := []struct {
		name        string
		slots       []manifest.Slot
		context     string
		expectedTop string // Name of the slot expected to be first
	}{
		{
			name: "empty context returns original order",
			slots: []manifest.Slot{
				{FullyQualified: manifest.FullyQualified{Name: "header"}},
				{FullyQualified: manifest.FullyQualified{Name: "footer"}},
			},
			context:     "",
			expectedTop: "header",
		},
		{
			name: "context matches slot name",
			slots: []manifest.Slot{
				{FullyQualified: manifest.FullyQualified{Name: "header"}},
				{FullyQualified: manifest.FullyQualified{Name: "footer"}},
				{FullyQualified: manifest.FullyQualified{Name: "image"}},
			},
			context:     "add an image to the card",
			expectedTop: "image",
		},
		{
			name: "context matches slot description",
			slots: []manifest.Slot{
				{FullyQualified: manifest.FullyQualified{Name: "header", Description: "The header section"}},
				{FullyQualified: manifest.FullyQualified{Name: "media", Description: "Place for images and videos"}},
				{FullyQualified: manifest.FullyQualified{Name: "footer", Description: "The footer section"}},
			},
			context:     "I need to display a video",
			expectedTop: "media",
		},
		{
			name: "default slot gets boost",
			slots: []manifest.Slot{
				{FullyQualified: manifest.FullyQualified{Name: "header"}},
				{FullyQualified: manifest.FullyQualified{Name: ""}}, // default slot
				{FullyQualified: manifest.FullyQualified{Name: "footer"}},
			},
			context:     "some generic content",
			expectedTop: "", // default slot should be first
		},
		{
			name: "name match beats description match",
			slots: []manifest.Slot{
				{FullyQualified: manifest.FullyQualified{Name: "other", Description: "Contains button content"}},
				{FullyQualified: manifest.FullyQualified{Name: "button"}},
			},
			context:     "add a button",
			expectedTop: "button",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := inferRelevantSlots(tt.slots, tt.context)
			assert.Equal(t, len(tt.slots), len(result), "Should return same number of slots")
			if len(result) > 0 {
				assert.Equal(t, tt.expectedTop, result[0].Name, "Top slot should match expected")
			}
		})
	}
}

func TestInferRelevantSlots_EmptySlots(t *testing.T) {
	result := inferRelevantSlots(nil, "some context")
	assert.Nil(t, result, "Should return nil for nil slots")

	result = inferRelevantSlots([]manifest.Slot{}, "some context")
	assert.Equal(t, 0, len(result), "Should return empty for empty slots")
}
