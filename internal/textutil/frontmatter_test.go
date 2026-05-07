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
package textutil

import (
	"testing"
)

func TestStripFrontmatter(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "no frontmatter",
			input: "<my-element>content</my-element>",
			want:  "<my-element>content</my-element>",
		},
		{
			name:  "simple frontmatter",
			input: "---\ndescription: A demo\n---\n<my-element>content</my-element>\n",
			want:  "<my-element>content</my-element>\n",
		},
		{
			name:  "frontmatter with multiple fields",
			input: "---\ndescription: Default accordion\nurl: /demo/\nfor: rh-accordion\n---\n<rh-accordion>\n  <rh-accordion-header>Item</rh-accordion-header>\n</rh-accordion>\n",
			want:  "<rh-accordion>\n  <rh-accordion-header>Item</rh-accordion-header>\n</rh-accordion>\n",
		},
		{
			name:  "frontmatter with CRLF line endings",
			input: "---\r\ndescription: A demo\r\n---\r\n<my-element>content</my-element>\r\n",
			want:  "<my-element>content</my-element>\r\n",
		},
		{
			name:  "frontmatter with leading whitespace",
			input: "\n\n---\ndescription: A demo\n---\n<my-element>content</my-element>\n",
			want:  "<my-element>content</my-element>\n",
		},
		{
			name:  "no closing delimiter",
			input: "---\ndescription: broken\n<my-element>content</my-element>\n",
			want:  "---\ndescription: broken\n<my-element>content</my-element>\n",
		},
		{
			name:  "empty input",
			input: "",
			want:  "",
		},
		{
			name:  "only frontmatter no content",
			input: "---\ndescription: Just metadata\n---\n",
			want:  "",
		},
		{
			name:  "closing delimiter followed by lone CR",
			input: "---\nfield: value\n---\r<content>",
			want:  "\r<content>",
		},
		{
			name:  "triple dash in HTML content not stripped",
			input: "<div>---</div>\n<my-element>content</my-element>\n",
			want:  "<div>---</div>\n<my-element>content</my-element>\n",
		},
		{
			name:  "frontmatter then body with triple dash",
			input: "---\ndescription: A demo\n---\n<div>---</div>\n<my-element>content</my-element>\n",
			want:  "<div>---</div>\n<my-element>content</my-element>\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := string(StripFrontmatter([]byte(tt.input)))
			if got != tt.want {
				t.Errorf("StripFrontmatter() =\n%q\nwant:\n%q", got, tt.want)
			}
		})
	}
}
