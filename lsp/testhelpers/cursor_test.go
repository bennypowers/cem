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
package testhelpers

import (
	"testing"
)

// Inline: these test the cursor parser itself, which is a pure function.
// Fixture/golden pattern would be circular since the parser is part of the fixture loader.
func TestTSCursorParser(t *testing.T) {
	tests := []struct {
		name      string
		content   string
		wantLine  uint32
		wantChar  uint32
		wantClean string
		wantNil   bool
	}{
		{
			name: "JS comment in regular code",
			content: "const tpl = html`<test-element .prop>`;\n" +
				"/*                              ^cursor */\n",
			wantLine:  0,
			wantChar:  32,
			wantClean: "const tpl = html`<test-element .prop>`;\n",
		},
		{
			name:      "JS comment under import",
			content:   "import { html } from 'lit';\n/*        ^cursor */\nconst x = 1;\n",
			wantLine:  0,
			wantChar:  10,
			wantClean: "import { html } from 'lit';\nconst x = 1;\n",
		},
		{
			name: "HTML comment inside template literal",
			content: "import { html } from 'lit';\n" +
				"\n" +
				"export class MyComponent {\n" +
				"  render() {\n" +
				"    return html`\n" +
				"      <test-component prop=\"value\"></test-component>\n" +
				"      <!-- ^cursor -->\n" +
				"    `;\n" +
				"  }\n" +
				"}\n",
			wantLine: 5,
			wantChar: 11,
			wantClean: "import { html } from 'lit';\n" +
				"\n" +
				"export class MyComponent {\n" +
				"  render() {\n" +
				"    return html`\n" +
				"      <test-component prop=\"value\"></test-component>\n" +
				"    `;\n" +
				"  }\n" +
				"}\n",
		},
		{
			name: "HTML comment in template with interpolation",
			content: "const tpl = html`\n" +
				"  <my-el .prop=${this.val}>\n" +
				"    <span>${this.text}</span>\n" +
				"  </my-el>\n" +
				"  <!-- ^cursor -->\n" +
				"`;\n",
			wantLine: 3,
			wantChar: 7,
			wantClean: "const tpl = html`\n" +
				"  <my-el .prop=${this.val}>\n" +
				"    <span>${this.text}</span>\n" +
				"  </my-el>\n" +
				"`;\n",
		},
		{
			name: "cursor on child element inside slot",
			content: "const tpl = html`\n" +
				"  <my-card>\n" +
				"    <span slot=\"header\">Title</span>\n" +
				"    <!-- ^cursor -->\n" +
				"  </my-card>\n" +
				"`;\n",
			wantLine: 2,
			wantChar: 9,
			wantClean: "const tpl = html`\n" +
				"  <my-card>\n" +
				"    <span slot=\"header\">Title</span>\n" +
				"  </my-card>\n" +
				"`;\n",
		},
		{
			name: "cursor on element with slot attribute",
			content: "const tpl = html`\n" +
				"  <my-card>\n" +
				"    <h2 slot=\"header\">Title</h2>\n" +
				"    <!--     ^cursor -->\n" +
				"  </my-card>\n" +
				"`;\n",
			wantLine: 2,
			wantChar: 13,
			wantClean: "const tpl = html`\n" +
				"  <my-card>\n" +
				"    <h2 slot=\"header\">Title</h2>\n" +
				"  </my-card>\n" +
				"`;\n",
		},
		{
			name: "cursor after property binding with interpolation",
			content: "const tpl = html`\n" +
				"  <my-el .items=${this.data}>\n" +
				"  <!--          ^cursor -->\n" +
				"    ${this.data.map(item => html`<li>${item}</li>`)}\n" +
				"  </my-el>\n" +
				"`;\n",
			wantLine: 1,
			wantChar: 16,
			wantClean: "const tpl = html`\n" +
				"  <my-el .items=${this.data}>\n" +
				"    ${this.data.map(item => html`<li>${item}</li>`)}\n" +
				"  </my-el>\n" +
				"`;\n",
		},
		{
			name: "cursor on event binding",
			content: "const tpl = html`\n" +
				"  <my-button @click=${this.onClick}>\n" +
				"  <!--       ^cursor -->\n" +
				"    Click me\n" +
				"  </my-button>\n" +
				"`;\n",
			wantLine: 1,
			wantChar: 13,
			wantClean: "const tpl = html`\n" +
				"  <my-button @click=${this.onClick}>\n" +
				"    Click me\n" +
				"  </my-button>\n" +
				"`;\n",
		},
		{
			name: "cursor on boolean binding",
			content: "const tpl = html`\n" +
				"  <my-input ?disabled=${this.isDisabled}\n" +
				"  <!--      ^cursor -->\n" +
				"            type=\"text\">\n" +
				"`;\n",
			wantLine: 1,
			wantChar: 12,
			wantClean: "const tpl = html`\n" +
				"  <my-input ?disabled=${this.isDisabled}\n" +
				"            type=\"text\">\n" +
				"`;\n",
		},
		{
			name: "cursor between two interpolations",
			content: "const tpl = html`\n" +
				"  <div class=${this.cls}>\n" +
				"    <my-el .val=${this.v}></my-el>\n" +
				"    <!--   ^cursor -->\n" +
				"  </div>\n" +
				"`;\n",
			wantLine: 2,
			wantChar: 11,
			wantClean: "const tpl = html`\n" +
				"  <div class=${this.cls}>\n" +
				"    <my-el .val=${this.v}></my-el>\n" +
				"  </div>\n" +
				"`;\n",
		},
		{
			name: "nested template literals",
			content: "const tpl = html`\n" +
				"  <ul>\n" +
				"    ${items.map(i => html`\n" +
				"      <li>${i.name}</li>\n" +
				"      <!-- ^cursor -->\n" +
				"    `)}\n" +
				"  </ul>\n" +
				"`;\n",
			wantLine: 3,
			wantChar: 11,
			wantClean: "const tpl = html`\n" +
				"  <ul>\n" +
				"    ${items.map(i => html`\n" +
				"      <li>${i.name}</li>\n" +
				"    `)}\n" +
				"  </ul>\n" +
				"`;\n",
		},
		{
			name: "template with no interpolations",
			content: "const tpl = html`\n" +
				"  <my-element attr=\"value\"></my-element>\n" +
				"  <!--        ^cursor -->\n" +
				"`;\n",
			wantLine: 1,
			wantChar: 14,
			wantClean: "const tpl = html`\n" +
				"  <my-element attr=\"value\"></my-element>\n" +
				"`;\n",
		},
		{
			name: "CRLF line endings",
			content: "const tpl = html`\r\n" +
				"  <my-element attr=\"v\"></my-element>\r\n" +
				"  <!-- ^cursor -->\r\n" +
				"`;\r\n",
			wantLine: 1,
			wantChar: 7,
			wantClean: "const tpl = html`\r\n" +
				"  <my-element attr=\"v\"></my-element>\r\n" +
				"`;\r\n",
		},
		{
			name:    "no marker",
			content: "const x = 1;\n",
			wantNil: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleaned, cursor := TSCursorParser(tt.content)
			if tt.wantNil {
				if cursor != nil {
					t.Errorf("Expected nil cursor, got %+v", cursor)
				}
				return
			}
			if cursor == nil {
				t.Fatal("Expected cursor, got nil")
			}
			if cursor.Line != tt.wantLine {
				t.Errorf("Line: want %d, got %d", tt.wantLine, cursor.Line)
			}
			if cursor.Character != tt.wantChar {
				t.Errorf("Character: want %d, got %d", tt.wantChar, cursor.Character)
			}
			if cleaned != tt.wantClean {
				t.Errorf("Cleaned content mismatch.\nWant:\n%q\nGot:\n%q", tt.wantClean, cleaned)
			}
		})
	}
}

func TestTSCursorParser_NonASCII(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		wantLine uint32
		wantChar uint32
	}{
		{
			name: "emoji in attribute value",
			content: "const tpl = html`\n" +
				"  <my-el label=\"\U0001F600 hello\"></my-el>\n" +
				"  <!--                   ^cursor -->\n" +
				"`;\n",
			wantLine: 1,
			wantChar: 25,
		},
		{
			name: "emoji in tag name cursor after emoji",
			// x-\U0001F600el is a legal custom element name (emoji not first char)
			// cursor points at the attribute after the emoji tag
			content: "const tpl = html`\n" +
				"  <x-\U0001F600el attr=\"v\"></x-\U0001F600el>\n" +
				"  <!--       ^cursor -->\n" +
				"`;\n",
			wantLine: 1,
			wantChar: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, cursor := TSCursorParser(tt.content)
			if cursor == nil {
				t.Fatal("Expected cursor, got nil")
			}
			if cursor.Line != tt.wantLine {
				t.Errorf("Line: want %d, got %d", tt.wantLine, cursor.Line)
			}
			if cursor.Character != tt.wantChar {
				t.Errorf("Character: want %d, got %d", tt.wantChar, cursor.Character)
			}
		})
	}
}

func TestCSSCursorParser(t *testing.T) {
	content := `my-element::part(button) {
/*                ^cursor */
  color: red;
}
`
	cleaned, cursor := CSSCursorParser(content)
	if cursor == nil {
		t.Fatal("Expected cursor, got nil")
	}
	if cursor.Line != 0 {
		t.Errorf("Line: want 0, got %d", cursor.Line)
	}
	if cursor.Character != 18 {
		t.Errorf("Character: want 18, got %d", cursor.Character)
	}
	wantClean := "my-element::part(button) {\n  color: red;\n}\n"
	if cleaned != wantClean {
		t.Errorf("Cleaned content mismatch.\nWant:\n%q\nGot:\n%q", wantClean, cleaned)
	}
}
