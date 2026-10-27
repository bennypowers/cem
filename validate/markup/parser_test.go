/*
Copyright © 2026 Benny Powers <web@bennypowers.com>

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
package markup

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// isCustomElementTag is a pure function; a table is clearer than a golden.
func TestIsCustomElementTag(t *testing.T) {
	tests := []struct {
		tag  string
		want bool
	}{
		{"my-element", true},
		{"a-b", true},
		{"a-", true},
		{"a--b", true},
		{"my-element-", true},
		{"my_ns-element", false},
		{"my-element_", false},
		{"1-element", false},
		{"-element", false},
		{"element", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.tag, func(t *testing.T) {
			assert.Equal(t, tt.want, isCustomElementTag(tt.tag))
		})
	}
}

func TestJinjaHTMLRanges(t *testing.T) {
	tests := []struct {
		name      string
		source    string
		wantRange bool
	}{
		{"plain html", `<my-el></my-el>`, true},
		{"html region", `<my-el>{% if x %}<my-el>{% endif %}</my-el>`, true},
		{"unbalanced control flow", `{% if x %}<my-el bogus>`, true},
		{"comment", `{# <my-el bogus> #}`, false},
		{"expression string", `{{ "<my-el bogus>" }}`, false},
		{"raw block", `{% raw %}<my-el bogus>{% endraw %}`, false},
		{"set string", `{% set x = "<my-el bogus>" %}`, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ranges, err := jinjaHTMLRanges(tt.source)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRange, len(ranges) > 0)
		})
	}
}

func TestParseElementsJinjaWithoutHTMLRanges(t *testing.T) {
	elements, err := parseElements(`{% set x = "<my-el bogus>" %}`, formatJinja)
	require.NoError(t, err)
	assert.Empty(t, elements)
}

func TestParseElementsJinjaUnbalanced(t *testing.T) {
	elements, err := parseElements(`{% if x %}<my-el bogus>`, formatJinja)
	require.NoError(t, err)
	require.Len(t, elements, 1)
	assert.Equal(t, "my-el", elements[0].tagName)
}

// TestParseElementsManySiblings guards linear traversal of large sibling
// lists. It asserts correctness only; timing lives in
// BenchmarkParseElementsManySiblings.
func TestParseElementsManySiblings(t *testing.T) {
	const count = 20000
	source := strings.Repeat(`<div></div>`, count)

	elements, err := parseElements(source, formatHTML)
	require.NoError(t, err)
	require.Len(t, elements, count)
}

func BenchmarkParseElementsManySiblings(b *testing.B) {
	const count = 20000
	source := strings.Repeat(`<div></div>`, count)

	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		elements, err := parseElements(source, formatHTML)
		if err != nil {
			b.Fatal(err)
		}
		if len(elements) != count {
			b.Fatalf("got %d elements, want %d", len(elements), count)
		}
	}
}

func TestAttributeValueUnescape(t *testing.T) {
	elements, err := parseElements(`<my-el a="a&amp;b" b="prim&#97;ry" c="primar&#x79;">`, formatHTML)
	require.NoError(t, err)
	require.Len(t, elements, 1)
	require.Len(t, elements[0].attributes, 3)

	assert.Equal(t, "a&b", elements[0].attributes[0].value)
	assert.Equal(t, "primary", elements[0].attributes[1].value)
	assert.Equal(t, "primary", elements[0].attributes[2].value)
}

func TestPositionForNonASCII(t *testing.T) {
	elements, err := parseElements(`<p>éé</p><my-button bogus="1">`, formatHTML)
	require.NoError(t, err)

	var button *elementNode
	for _, element := range elements {
		if element.tagName == "my-button" {
			button = element
			break
		}
	}
	require.NotNil(t, button)
	require.Len(t, button.attributes, 1)
	// Columns are bytes: "<p>" (3) + "éé" (4) + "</p>" (4) + "<my-button " (11).
	assert.Equal(t, 23, button.attributes[0].column)
}
