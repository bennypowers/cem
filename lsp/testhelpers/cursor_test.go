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
			name: "caret under property binding",
			content: "const tpl = html`<test-element .prop>`;\n" +
				"/*                              ^cursor */\n",
			wantLine:  0,
			wantChar:  32,
			wantClean: "const tpl = html`<test-element .prop>`;\n",
		},
		{
			name:      "caret under import",
			content:   "import { html } from 'lit';\n/*        ^cursor */\nconst x = 1;\n",
			wantLine:  0,
			wantChar:  10,
			wantClean: "import { html } from 'lit';\nconst x = 1;\n",
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
