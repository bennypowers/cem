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

// TestJinjaHTMLRanges asserts the return shape of one pure range function;
// exact byte ranges are the point, so an inline table is clearer than a golden.
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
		{"mixed html and expression string", `<div>{{ "<my-el bogus>" }}</div>`, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ranges, _, err := analyzeJinjaRanges(tt.source, false)
			require.NoError(t, err)
			assert.Equal(t, tt.wantRange, len(ranges) > 0)

			// Text ranges must never widen to cover template syntax.
			for _, r := range ranges {
				text := tt.source[r.StartByte:r.EndByte]
				assert.NotContains(t, text, "{{", "range covers an expression")
				assert.NotContains(t, text, "{%", "range covers a statement")
			}
			if tt.name == "mixed html and expression string" {
				for _, r := range ranges {
					assert.NotContains(t, tt.source[r.StartByte:r.EndByte], "my-el",
						"markup inside an expression must not become a text range")
				}
			}
		})
	}
}

// Handlebars expressions in tag position are not attributes, and Handlebars
// comments are not markup. Inline assertions keep the parsed names visible.
func TestParseElementsHandlebarsMasksExpressions(t *testing.T) {
	tests := []struct {
		name      string
		source    string
		wantAttrs []string
	}{
		{
			name:      "conditional attribute",
			source:    `<my-button {{#if primary}}variant="primary"{{/if}}>Save</my-button>`,
			wantAttrs: []string{"variant"},
		},
		{
			name:      "comment containing markup",
			source:    `<my-button>{{!-- <my-bad bogus> --}}Save</my-button>`,
			wantAttrs: []string{},
		},
		{
			name:      "helper with a string argument",
			source:    `<my-button variant="{{t "save"}}">Save</my-button>`,
			wantAttrs: []string{"variant"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			elements, err := parseElements(tt.source, formatHandlebars)
			require.NoError(t, err)
			require.Len(t, elements, 1)
			assert.Equal(t, "my-button", elements[0].tagName)

			names := make([]string, len(elements[0].attributes))
			for i, attr := range elements[0].attributes {
				names[i] = attr.name
				assert.NotContains(t, attr.name, "{{", "template syntax leaked into an attribute name")
				assert.NotContains(t, attr.name, "}}", "template syntax leaked into an attribute name")
			}
			assert.Equal(t, tt.wantAttrs, names)
		})
	}
}

// This is an empty-result edge case; a golden would only pin [].
func TestParseElementsJinjaWithoutHTMLRanges(t *testing.T) {
	elements, err := parseElements(`{% set x = "<my-el bogus>" %}`, formatJinja)
	require.NoError(t, err)
	assert.Empty(t, elements)
}

// A golden would pin the full parsed node tree for a one-element edge case;
// the asserted subset states the behavior directly.
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

// The assertion is about the dynamic flag on each parsed attribute, which a
// golden of validation issues cannot observe.
func TestParseElementsDynamicValues(t *testing.T) {
	elements, err := parseElements(
		`<my-el a="{{ x }}-suffix" b="pre-{{ x }}" c="{{ x }}" d="static">`,
		formatJinja,
	)
	require.NoError(t, err)
	require.Len(t, elements, 1)

	attributes := elements[0].attributes
	require.Len(t, attributes, 4)
	assert.True(t, attributes[0].dynamic, "expression plus suffix")
	assert.True(t, attributes[1].dynamic, "static prefix plus expression")
	assert.True(t, attributes[2].dynamic, "pure expression")
	assert.False(t, attributes[3].dynamic, "static value")
}

// This pins decoded attribute values, a parser detail that validation goldens
// only observe indirectly.
func TestAttributeValueUnescape(t *testing.T) {
	elements, err := parseElements(
		`<my-el a="a&amp;b" b="prim&#97;ry" c="primar&#x79;" d=a&amp;b e="&amp;amp;" slot="i&amp;con">`,
		formatHTML,
	)
	require.NoError(t, err)
	require.Len(t, elements, 1)
	require.Len(t, elements[0].attributes, 6)

	assert.Equal(t, "a&b", elements[0].attributes[0].value)     // named
	assert.Equal(t, "primary", elements[0].attributes[1].value) // decimal
	assert.Equal(t, "primary", elements[0].attributes[2].value) // hex
	assert.Equal(t, "a&b", elements[0].attributes[3].value)     // unquoted
	assert.Equal(t, "&amp;", elements[0].attributes[4].value)   // decoded once
	assert.Equal(t, "i&con", elements[0].attributes[5].value)   // slot value
}

// The parser keeps zero-based UTF-8 byte columns; the exported API converts
// them to UTF-16 at the boundary. The assertion stays inline because it checks
// an internal parser detail, not a validation finding a golden could capture.
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
	assert.Equal(t, 22, button.attributes[0].column)
}
