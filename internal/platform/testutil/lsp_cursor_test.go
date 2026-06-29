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
package testutil

import (
	"testing"

	protocol "github.com/bennypowers/glsp/protocol_3_17"
)

func TestExtractHTMLCursorMarker(t *testing.T) {
	tests := []struct {
		name      string
		content   string
		wantLine  uint32
		wantChar  uint32
		wantClean string
		wantNil   bool
	}{
		{
			name:      "edge case line 0",
			content:   "<test-element attr=\"hello\"></test-element>\n<!-- ^cursor -->\n",
			wantLine:  0,
			wantChar:  5,
			wantClean: "<test-element attr=\"hello\"></test-element>\n",
		},
		{
			name: "caret under nested element",
			content: `<!DOCTYPE html>
<html>
<body>
  <test-element></test-element>
  <!--       ^cursor -->
</body>
</html>`,
			wantLine:  3,
			wantChar:  13,
			wantClean: "<!DOCTYPE html>\n<html>\n<body>\n  <test-element></test-element>\n</body>\n</html>",
		},
		{
			name: "caret under attribute value",
			content: `<!DOCTYPE html>
<html>
<body>
  <test-element test-attr="hello"></test-element>
  <!--                     ^cursor -->
</body>
</html>`,
			wantLine:  3,
			wantChar:  27,
			wantClean: "<!DOCTYPE html>\n<html>\n<body>\n  <test-element test-attr=\"hello\"></test-element>\n</body>\n</html>",
		},
		{
			name:    "no marker",
			content: "<test-element></test-element>\n",
			wantNil: true,
		},
		{
			name:    "marker on first line has no line above",
			content: "<!-- ^cursor -->\n<test-element></test-element>\n",
			wantNil: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleaned, cursor := extractHTMLCursorMarker(tt.content)
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

func TestExtractCursorFrontmatter(t *testing.T) {
	tests := []struct {
		name      string
		content   string
		wantLine  uint32
		wantChar  uint32
		wantClean string
		wantNil   bool
	}{
		{
			name: "basic frontmatter",
			content: `---
cursor:
  line: 2
  character: 10
---
<html>
<body>
  <test-element></test-element>
</body>
</html>
`,
			wantLine:  2,
			wantChar:  10,
			wantClean: "<html>\n<body>\n  <test-element></test-element>\n</body>\n</html>\n",
		},
		{
			name: "no cursor key in frontmatter",
			content: `---
title: test
---
<html></html>
`,
			wantNil: true,
		},
		{
			name:    "no frontmatter",
			content: "<html></html>\n",
			wantNil: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleaned, cursor := extractCursorFrontmatter(tt.content)
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

func TestExtractCursor_HTMLMarkerOverFrontmatter(t *testing.T) {
	content := "<test-element></test-element>\n<!-- ^cursor -->\n"
	cleaned, cursor := extractCursor(content, "html")
	if cursor == nil {
		t.Fatal("Expected cursor from HTML marker, got nil")
	}
	if cursor.Line != 0 {
		t.Errorf("Line: want 0, got %d", cursor.Line)
	}
	_ = cleaned
}

func TestExtractCursor_FrontmatterFallbackForTS(t *testing.T) {
	content := `---
cursor:
  line: 1
  character: 5
---
const x = 1;
const y = 2;
`
	cleaned, cursor := extractCursor(content, "ts")
	if cursor == nil {
		t.Fatal("Expected cursor from frontmatter, got nil")
	}
	if cursor.Line != 1 || cursor.Character != 5 {
		t.Errorf("Want (1,5), got (%d,%d)", cursor.Line, cursor.Character)
	}
	if cleaned != "const x = 1;\nconst y = 2;\n" {
		t.Errorf("Frontmatter not stripped: %q", cleaned)
	}
}

func TestExtractCursor_NilWhenNoMarkerOrFrontmatter(t *testing.T) {
	_, cursor := extractCursor("<test-element></test-element>\n", "html")
	if cursor != nil {
		t.Errorf("Expected nil cursor, got %+v", cursor)
	}
}

func TestLSPFixture_ParseCursor(t *testing.T) {
	fixture := &LSPFixture{
		InputContent: "some content",
		InputHTML:    "some content",
	}
	mockParser := func(content string) (string, *protocol.Position) {
		return "cleaned", &protocol.Position{Line: 3, Character: 7}
	}
	fixture.ParseCursor(mockParser)
	if fixture.Cursor == nil {
		t.Fatal("Expected Cursor to be set")
	}
	if fixture.Cursor.Line != 3 || fixture.Cursor.Character != 7 {
		t.Errorf("Want (3,7), got (%d,%d)", fixture.Cursor.Line, fixture.Cursor.Character)
	}
	if fixture.InputContent != "cleaned" {
		t.Errorf("InputContent not updated: %q", fixture.InputContent)
	}
}

func TestLSPFixture_ParseCursor_SkipsWhenAlreadySet(t *testing.T) {
	fixture := &LSPFixture{
		InputContent: "original",
		Cursor:       &protocol.Position{Line: 1, Character: 2},
	}
	called := false
	mockParser := func(content string) (string, *protocol.Position) {
		called = true
		return "changed", &protocol.Position{Line: 99, Character: 99}
	}
	fixture.ParseCursor(mockParser)
	if called {
		t.Error("Parser should not be called when Cursor is already set")
	}
	if fixture.InputContent != "original" {
		t.Error("InputContent should not change")
	}
}
