package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsCustomElementTag(t *testing.T) {
	tests := []struct {
		name     string
		tagName  string
		expected bool
	}{
		{"valid custom element", "my-element", true},
		{"standard HTML element", "div", false},
		{"leading hyphen", "-leading-hyphen", false},
		{"uppercase start", "Uppercase-element", false},
		{"minimum valid", "a-b", true},
		{"empty string", "", false},
		{"no hyphen", "nohyphen", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, IsCustomElementTag(tt.tagName))
		})
	}
}
