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

package transform

import (
	"strings"
	"testing"
)

// TestTransformCSS_Basic tests basic CSS to constructable stylesheet transformation
func TestTransformCSS_Basic(t *testing.T) {
	cssContent := `:host {
  display: block;
  color: var(--color, red);
}`

	result := TransformCSS([]byte(cssContent), "test.css")

	// Should wrap in CSSStyleSheet
	if !strings.Contains(result, "new CSSStyleSheet()") {
		t.Error("Output missing CSSStyleSheet constructor")
	}

	// Should use replaceSync
	if !strings.Contains(result, "replaceSync") {
		t.Error("Output missing replaceSync call")
	}

	// Should export default
	if !strings.Contains(result, "export default") {
		t.Error("Output missing default export")
	}

	// Should contain the original CSS
	if !strings.Contains(result, ":host") {
		t.Error("Output missing original CSS content")
	}

	// Should escape backticks in CSS
	cssWithBacktick := "content: `test`;"
	escaped := TransformCSS([]byte(cssWithBacktick), "test.css")
	if strings.Contains(escaped, "content: `test`;") {
		t.Error("Backticks in CSS not escaped")
	}

	// Should only escape ${ not all $ (matches Lit behavior)
	cssWithDollar := "width: $var; height: ${expr};"
	dollarEscaped := TransformCSS([]byte(cssWithDollar), "test.css")
	// $var should NOT be escaped
	if !strings.Contains(dollarEscaped, "$var") {
		t.Error("Single $ incorrectly escaped (should only escape ${)")
	}
	// ${ should be escaped
	if strings.Contains(dollarEscaped, "${expr}") && !strings.Contains(dollarEscaped, "\\${") {
		t.Error("${ not properly escaped")
	}

	// Should escape </ to prevent script tag injection
	cssWithScriptTag := "content: '</script>';"
	scriptEscaped := TransformCSS([]byte(cssWithScriptTag), "test.css")
	if strings.Contains(scriptEscaped, "</script>") && !strings.Contains(scriptEscaped, "<\\/script>") {
		t.Error("</ not properly escaped")
	}
}
