package completion

import (
	"testing"

	protocol "github.com/bennypowers/glsp/protocol_3_17"
	"github.com/stretchr/testify/assert"
)

// Inline: pure function, table-driven

func TestStartsWithIgnoreCase(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		prefix   string
		expected bool
	}{
		{
			name:     "matching case",
			s:        "my-element",
			prefix:   "my-",
			expected: true,
		},
		{
			name:     "different case",
			s:        "My-Element",
			prefix:   "my-",
			expected: true,
		},
		{
			name:     "no match",
			s:        "other-element",
			prefix:   "my-",
			expected: false,
		},
		{
			name:     "empty prefix",
			s:        "anything",
			prefix:   "",
			expected: true,
		},
		{
			name:     "prefix longer than string",
			s:        "ab",
			prefix:   "abc",
			expected: false,
		},
		{
			name:     "exact match",
			s:        "my-element",
			prefix:   "my-element",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := startsWithIgnoreCase(tt.s, tt.prefix)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDeduplicateCompletionItems(t *testing.T) {
	valueKind := protocol.CompletionItemKindValue

	tests := []struct {
		name     string
		items    []protocol.CompletionItem
		expected int // expected number of items after dedup
	}{
		{
			name:     "no duplicates",
			items: []protocol.CompletionItem{
				{Label: "alpha", Kind: &valueKind},
				{Label: "beta", Kind: &valueKind},
				{Label: "gamma", Kind: &valueKind},
			},
			expected: 3,
		},
		{
			name: "exact duplicates",
			items: []protocol.CompletionItem{
				{Label: "primary", Kind: &valueKind, Detail: &[]string{"Variant value"}[0]},
				{Label: "primary", Kind: &valueKind, Detail: &[]string{"Union type value"}[0]},
			},
			expected: 1,
		},
		{
			name: "different labels",
			items: []protocol.CompletionItem{
				{Label: "red", Kind: &valueKind},
				{Label: "blue", Kind: &valueKind},
			},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := deduplicateCompletionItems(tt.items)
			assert.Len(t, result, tt.expected)
		})
	}
}

func TestDeduplicateCompletionItems_PrefersTypeBased(t *testing.T) {
	valueKind := protocol.CompletionItemKindValue

	items := []protocol.CompletionItem{
		{Label: "primary", Kind: &valueKind, Detail: &[]string{"Variant value"}[0]},
		{Label: "primary", Kind: &valueKind, Detail: &[]string{"Union type value"}[0]},
	}

	result := deduplicateCompletionItems(items)
	assert.Len(t, result, 1)
	assert.Equal(t, "Union type value", *result[0].Detail)
}

func TestShouldPreferItem(t *testing.T) {
	tests := []struct {
		name       string
		newDetail  string
		existDetail string
		newLabel   string
		existLabel string
		expected   bool
	}{
		{
			name:       "prefer union type over variant",
			newDetail:  "Union type value",
			existDetail: "Variant value",
			newLabel:   "primary",
			existLabel: "primary",
			expected:   true,
		},
		{
			name:       "keep existing union type over variant",
			newDetail:  "Variant value",
			existDetail: "Union type value",
			newLabel:   "primary",
			existLabel: "primary",
			expected:   false,
		},
		{
			name:       "default label wins over non-default",
			newDetail:  "Some detail",
			existDetail: "Some detail",
			newLabel:   "value (default)",
			existLabel: "value",
			expected:   true,
		},
		{
			name:       "existing default label wins",
			newDetail:  "Some detail",
			existDetail: "Some detail",
			newLabel:   "value",
			existLabel: "value (default)",
			expected:   false,
		},
		{
			name:       "both same type, keep existing",
			newDetail:  "Variant value",
			existDetail: "Variant value",
			newLabel:   "primary",
			existLabel: "primary",
			expected:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newItem := protocol.CompletionItem{
				Label:  tt.newLabel,
				Detail: &tt.newDetail,
			}
			existItem := protocol.CompletionItem{
				Label:  tt.existLabel,
				Detail: &tt.existDetail,
			}
			result := shouldPreferItem(newItem, existItem)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestShouldPreferItem_NilDetails(t *testing.T) {
	newItem := protocol.CompletionItem{Label: "test"}
	existItem := protocol.CompletionItem{Label: "test"}

	// Neither has detail; should keep existing (return false)
	result := shouldPreferItem(newItem, existItem)
	assert.False(t, result)
}
