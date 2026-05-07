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
	"testing"
)

func TestDetectTagPrefix(t *testing.T) {
	tests := []struct {
		name     string
		tagNames []string
		expected string
	}{
		{
			name:     "common two-char prefix",
			tagNames: []string{"ex-button", "ex-accordion", "ex-textarea", "ex-card"},
			expected: "ex-",
		},
		{
			name:     "common single-char prefix",
			tagNames: []string{"x-foo", "x-bar"},
			expected: "x-",
		},
		{
			name:     "multi-segment tags still detect first-segment prefix",
			tagNames: []string{"ns-accordion", "ns-accordion-header", "ns-navigation-primary"},
			expected: "ns-",
		},
		{
			name:     "no common prefix",
			tagNames: []string{"my-button", "your-panel", "their-card"},
			expected: "",
		},
		{
			name:     "single element cannot determine prefix",
			tagNames: []string{"ex-button"},
			expected: "",
		},
		{
			name:     "empty list",
			tagNames: []string{},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DetectTagPrefix(tt.tagNames)
			if got != tt.expected {
				t.Errorf("DetectTagPrefix(%v) = %q, want %q", tt.tagNames, got, tt.expected)
			}
		})
	}
}

func TestAutoDeriveAliases(t *testing.T) {
	tests := []struct {
		name     string
		tagNames []string
		existing map[string]string
		expected map[string]string
	}{
		{
			name:     "strips detected prefix",
			tagNames: []string{"ex-button", "ex-accordion", "ex-textarea"},
			existing: map[string]string{},
			expected: map[string]string{
				"ex-button":    "button",
				"ex-accordion": "accordion",
				"ex-textarea":  "textarea",
			},
		},
		{
			name:     "explicit alias takes precedence",
			tagNames: []string{"ex-cta", "ex-button", "ex-textarea"},
			existing: map[string]string{"ex-cta": "call-to-action"},
			expected: map[string]string{
				"ex-cta":      "call-to-action",
				"ex-button":   "button",
				"ex-textarea": "textarea",
			},
		},
		{
			name:     "multi-segment remainder preserved",
			tagNames: []string{"ns-accordion", "ns-accordion-header", "ns-navigation-primary"},
			existing: map[string]string{},
			expected: map[string]string{
				"ns-accordion":          "accordion",
				"ns-accordion-header":   "accordion-header",
				"ns-navigation-primary": "navigation-primary",
			},
		},
		{
			name:     "no common prefix returns only existing aliases",
			tagNames: []string{"my-button", "your-panel"},
			existing: map[string]string{"my-button": "btn"},
			expected: map[string]string{"my-button": "btn"},
		},
		{
			name:     "empty tag names returns existing aliases",
			tagNames: []string{},
			existing: map[string]string{"ex-button": "button"},
			expected: map[string]string{"ex-button": "button"},
		},
		{
			name:     "single element no derivation",
			tagNames: []string{"ex-button"},
			existing: map[string]string{},
			expected: map[string]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AutoDeriveAliases(tt.tagNames, tt.existing)
			if len(got) != len(tt.expected) {
				t.Errorf("len = %d, want %d\ngot:  %v\nwant: %v", len(got), len(tt.expected), got, tt.expected)
				return
			}
			for k, want := range tt.expected {
				if gotVal, ok := got[k]; !ok {
					t.Errorf("missing key %q", k)
				} else if gotVal != want {
					t.Errorf("[%q] = %q, want %q", k, gotVal, want)
				}
			}
		})
	}
}
