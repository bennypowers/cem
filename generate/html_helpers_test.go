/*
Copyright © 2026 Benny Powers

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
package generate

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDedentYaml(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", ""},
		{"single line", "  hello  ", "hello"},
		{"single line no indent", "hello", "hello"},
		{"two lines same indent", "line1\n    line2", "line1\nline2"},
		{"two lines different indent", "line1\n    line2\n  line3", "line1\n  line2\nline3"},
		{"blank lines ignored", "line1\n    line2\n\n    line3", "line1\nline2\n\nline3"},
		{"tabs", "line1\n\t\tline2\n\t\tline3", "line1\nline2\nline3"},
		{"mixed indent picks minimum", "line1\n      deep\n    shallow", "line1\n  deep\nshallow"},
		{"already dedented", "line1\nline2\nline3", "line1\nline2\nline3"},
		{"trailing whitespace trimmed", "  \n    hello\n    world\n  ", "hello\nworld"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, dedentYaml(tc.in))
		})
	}
}

func TestUnescapeBackticks(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"no backticks", "no backticks"},
		{"one \\`backtick\\`", "one `backtick`"},
		{"\\`\\`\\`", "```"},
		{"", ""},
		{"already `unescaped`", "already `unescaped`"},
	}
	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			assert.Equal(t, tc.want, unescapeBackticks(tc.in))
		})
	}
}

func TestGetInnerComment(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"strips delimiters", "<!-- hello -->", "hello"},
		{"strips with newlines", "<!-- \n  hello\n  world\n -->", "hello\nworld"},
		{"no delimiters passthrough", "just text", "just text"},
		{"empty comment", "<!---->", ""},
		{"unclosed comment", "<!-- hello", "hello"},
		{"whitespace only", "<!--   -->", ""},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, getInnerComment(tc.in))
		})
	}
}

func TestIsYamlComment(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{"description key", "description: hello", true},
		{"summary key", "summary: hello", true},
		{"deprecated key", "deprecated: true", true},
		{"slot key", "slot:\n  description: hi", true},
		{"part key", "part:\n  description: hi", true},
		{"irrelevant yaml key", "foo: bar", false},
		{"invalid yaml", ": : : :", false},
		{"empty string", "", false},
		{"plain text", "just a description", false},
		{"not yaml at all", "this is not yaml: [broken", false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, isYamlComment(tc.in))
		})
	}
}

func TestParseYamlComment(t *testing.T) {
	t.Run("plain text returns description", func(t *testing.T) {
		doc, err := parseYamlComment("<!-- A simple button -->", "")
		require.NoError(t, err)
		assert.Equal(t, "A simple button", doc.Description)
	})

	t.Run("yaml description", func(t *testing.T) {
		doc, err := parseYamlComment("<!-- description: A button component -->", "")
		require.NoError(t, err)
		assert.Equal(t, "A button component", doc.Description)
	})

	t.Run("yaml with summary and deprecated", func(t *testing.T) {
		comment := "<!-- \n  description: A button\n  summary: Button\n  deprecated: Use new-button\n-->"
		doc, err := parseYamlComment(comment, "")
		require.NoError(t, err)
		assert.Equal(t, "A button", doc.Description)
		assert.Equal(t, "Button", doc.Summary)
		assert.Equal(t, "Use new-button", doc.Deprecated)
	})

	t.Run("slot kind extracts slot docs", func(t *testing.T) {
		comment := "<!-- slot:\n  description: Default slot content\n  summary: Default\n-->"
		doc, err := parseYamlComment(comment, "slot")
		require.NoError(t, err)
		assert.Equal(t, "Default slot content", doc.Description)
		assert.Equal(t, "Default", doc.Summary)
	})

	t.Run("part kind extracts part docs", func(t *testing.T) {
		comment := "<!-- part:\n  description: The button part\n-->"
		doc, err := parseYamlComment(comment, "part")
		require.NoError(t, err)
		assert.Equal(t, "The button part", doc.Description)
	})

	t.Run("kind mismatch falls through to top-level", func(t *testing.T) {
		comment := "<!-- description: top level\nslot:\n  description: slot desc -->"
		doc, err := parseYamlComment(comment, "part")
		require.NoError(t, err)
		assert.Equal(t, "top level", doc.Description)
	})

	t.Run("escaped backticks unescaped", func(t *testing.T) {
		comment := "<!-- Use \\`code\\` here -->"
		doc, err := parseYamlComment(comment, "")
		require.NoError(t, err)
		assert.Equal(t, "Use `code` here", doc.Description)
	})
}
