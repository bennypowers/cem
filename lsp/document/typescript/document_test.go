package typescript

import (
	"sort"
	"testing"

	Q "bennypowers.dev/cem/internal/treesitter"
	"github.com/stretchr/testify/assert"
)

func TestGetCaptureMapKeys(t *testing.T) {
	tests := []struct {
		name     string
		input    Q.CaptureMap
		expected []string
	}{
		{
			name: "populated map",
			input: Q.CaptureMap{
				"tag.name":   {{Text: "div"}},
				"attr.name":  {{Text: "class"}},
				"attr.value": {{Text: "foo"}},
			},
			expected: []string{"attr.name", "attr.value", "tag.name"},
		},
		{
			name:     "empty map",
			input:    Q.CaptureMap{},
			expected: []string{},
		},
		{
			name: "single key",
			input: Q.CaptureMap{
				"context": {{Text: "something"}},
			},
			expected: []string{"context"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getCaptureMapKeys(tt.input)
			sort.Strings(result)
			sort.Strings(tt.expected)
			assert.Equal(t, tt.expected, result)
		})
	}
}
