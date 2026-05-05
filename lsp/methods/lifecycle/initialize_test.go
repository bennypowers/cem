package lifecycle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseStringSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected []string
	}{
		{
			name:     "string slice",
			input:    []string{"foo", "bar", "baz"},
			expected: []string{"foo", "bar", "baz"},
		},
		{
			name:     "[]any with strings",
			input:    []any{"alpha", "beta"},
			expected: []string{"alpha", "beta"},
		},
		{
			name:     "[]any with mixed types - only strings kept",
			input:    []any{"hello", 42, true, "world"},
			expected: []string{"hello", "world"},
		},
		{
			name:     "non-slice returns nil",
			input:    "not a slice",
			expected: nil,
		},
		{
			name:     "nil returns nil",
			input:    nil,
			expected: nil,
		},
		{
			name:     "integer returns nil",
			input:    42,
			expected: nil,
		},
		{
			name:     "empty []any",
			input:    []any{},
			expected: []string{},
		},
		{
			name:     "empty []string",
			input:    []string{},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseStringSlice(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
