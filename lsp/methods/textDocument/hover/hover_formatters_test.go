package hover

import (
	"testing"

	M "bennypowers.dev/cem/manifest"
	"github.com/stretchr/testify/assert"
)

// Inline: pure function, table-driven
// Markdown formatting output is compared as strings. These could be goldens
// but each case is a single-line expected string, making table-driven clearer.

func TestFormatAttributes(t *testing.T) {
	tests := []struct {
		name     string
		attrs    []M.Attribute
		expected string
	}{
		{
			name:     "empty list",
			attrs:    nil,
			expected: "",
		},
		{
			name: "single attr",
			attrs: []M.Attribute{
				{FullyQualified: M.FullyQualified{Name: "disabled"}},
			},
			expected: "### Attributes\n\n- **`disabled`**\n\n",
		},
		{
			name: "attr with type",
			attrs: []M.Attribute{
				{
					FullyQualified: M.FullyQualified{Name: "variant"},
					Type:           &M.Type{Text: "string"},
				},
			},
			expected: "### Attributes\n\n- **`variant`** _string_\n\n",
		},
		{
			name: "attr with inheritance",
			attrs: []M.Attribute{
				{
					FullyQualified: M.FullyQualified{Name: "color"},
					InheritedFrom:  &M.Reference{Name: "BaseElement"},
				},
			},
			expected: "### Attributes\n\n- **`color`** _(inherited from BaseElement)_\n\n",
		},
		{
			name: "attr with description",
			attrs: []M.Attribute{
				{
					FullyQualified: M.FullyQualified{Name: "size", Description: "The size of the element"},
				},
			},
			expected: "### Attributes\n\n- **`size`** - The size of the element\n\n",
		},
		{
			name: "attr with type, inheritance, and description",
			attrs: []M.Attribute{
				{
					FullyQualified: M.FullyQualified{Name: "label", Description: "Accessible label"},
					Type:           &M.Type{Text: "string"},
					InheritedFrom:  &M.Reference{Name: "LabelMixin"},
				},
			},
			expected: "### Attributes\n\n- **`label`** _string_ _(inherited from LabelMixin)_ - Accessible label\n\n",
		},
		{
			name: "multiple attributes",
			attrs: []M.Attribute{
				{FullyQualified: M.FullyQualified{Name: "disabled"}},
				{FullyQualified: M.FullyQualified{Name: "variant", Description: "Visual variant"}, Type: &M.Type{Text: "string"}},
			},
			expected: "### Attributes\n\n- **`disabled`**\n- **`variant`** _string_ - Visual variant\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatAttributes(tt.attrs)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatEvents(t *testing.T) {
	tests := []struct {
		name     string
		events   []M.Event
		expected string
	}{
		{
			name:     "empty list",
			events:   nil,
			expected: "",
		},
		{
			name: "single event",
			events: []M.Event{
				{FullyQualified: M.FullyQualified{Name: "click"}},
			},
			expected: "### Events\n\n- **`click`**\n\n",
		},
		{
			name: "event with type",
			events: []M.Event{
				{
					FullyQualified: M.FullyQualified{Name: "change"},
					Type:           &M.Type{Text: "CustomEvent"},
				},
			},
			expected: "### Events\n\n- **`change`** _CustomEvent_\n\n",
		},
		{
			name: "event with inheritance and description",
			events: []M.Event{
				{
					FullyQualified: M.FullyQualified{Name: "select", Description: "Fired when selection changes"},
					Type:           &M.Type{Text: "Event"},
					InheritedFrom:  &M.Reference{Name: "SelectMixin"},
				},
			},
			expected: "### Events\n\n- **`select`** _Event_ _(inherited from SelectMixin)_ - Fired when selection changes\n\n",
		},
		{
			name: "multiple events",
			events: []M.Event{
				{FullyQualified: M.FullyQualified{Name: "click"}},
				{FullyQualified: M.FullyQualified{Name: "change"}, Type: &M.Type{Text: "CustomEvent"}},
			},
			expected: "### Events\n\n- **`click`**\n- **`change`** _CustomEvent_\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatEvents(tt.events)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatSlots(t *testing.T) {
	tests := []struct {
		name     string
		slots    []M.Slot
		expected string
	}{
		{
			name:     "empty list",
			slots:    nil,
			expected: "",
		},
		{
			name: "named slot",
			slots: []M.Slot{
				{FullyQualified: M.FullyQualified{Name: "header"}},
			},
			expected: "### Slots\n\n- **`header`**\n\n",
		},
		{
			name: "inherited slot",
			slots: []M.Slot{
				{
					FullyQualified: M.FullyQualified{Name: "content"},
					InheritedFrom:  &M.Reference{Name: "LayoutMixin"},
				},
			},
			expected: "### Slots\n\n- **`content`** _(inherited from LayoutMixin)_\n\n",
		},
		{
			name: "slot with description",
			slots: []M.Slot{
				{
					FullyQualified: M.FullyQualified{Name: "footer", Description: "Footer content area"},
				},
			},
			expected: "### Slots\n\n- **`footer`** - Footer content area\n\n",
		},
		{
			name: "default slot (empty name)",
			slots: []M.Slot{
				{FullyQualified: M.FullyQualified{Name: "", Description: "Default slot"}},
			},
			expected: "### Slots\n\n- **``** - Default slot\n\n",
		},
		{
			name: "multiple slots",
			slots: []M.Slot{
				{FullyQualified: M.FullyQualified{Name: "header"}},
				{FullyQualified: M.FullyQualified{Name: "footer", Description: "Footer area"}, InheritedFrom: &M.Reference{Name: "LayoutMixin"}},
			},
			expected: "### Slots\n\n- **`header`**\n- **`footer`** _(inherited from LayoutMixin)_ - Footer area\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatSlots(tt.slots)
			assert.Equal(t, tt.expected, result)
		})
	}
}
