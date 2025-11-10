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

package inject_test

import (
	"strings"
	"testing"

	"bennypowers.dev/cem/serve/middleware/inject"
)

// TestInjectScript_InjectsIntoHead tests basic injection into <head>
func TestInjectScript_InjectsIntoHead(t *testing.T) {
	html := `<html><head><title>Test</title></head><body>content</body></html>`
	script := `<script src="/test.js"></script>`

	result := inject.InjectScript(html, script)

	if !strings.Contains(result, script) {
		t.Error("Expected script to be present in result")
	}

	headEnd := strings.Index(result, "</head>")
	scriptPos := strings.Index(result, script)

	if scriptPos > headEnd {
		t.Error("Expected script to be before </head>")
	}
}

// TestInjectScript_InjectsIntoBodyIfNoHead tests fallback to <body>
func TestInjectScript_InjectsIntoBodyIfNoHead(t *testing.T) {
	html := `<html><body>content</body></html>`
	script := `<script src="/test.js"></script>`

	result := inject.InjectScript(html, script)

	if !strings.Contains(result, script) {
		t.Error("Expected script to be present in result")
	}

	// Script should be in body since there's no head
	bodyStart := strings.Index(result, "<body>")
	scriptPos := strings.Index(result, script)

	if scriptPos < bodyStart {
		t.Error("Expected script to be in body when no head exists")
	}
}

// TestInjectScript_HandlesMinimalHTML tests injection into minimal HTML
func TestInjectScript_HandlesMinimalHTML(t *testing.T) {
	html := `<html></html>`
	script := `<script src="/test.js"></script>`

	result := inject.InjectScript(html, script)

	if !strings.Contains(result, script) {
		t.Error("Expected script to be present in minimal HTML")
	}
}

// TestInjectScript_HandlesComplexHead tests injection into complex <head>
func TestInjectScript_HandlesComplexHead(t *testing.T) {
	html := `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Complex Page</title>
    <link rel="stylesheet" href="/style.css">
</head>
<body>
    <h1>Hello</h1>
</body>
</html>`
	script := `<script type="module" src="/test.js"></script>`

	result := inject.InjectScript(html, script)

	if !strings.Contains(result, script) {
		t.Error("Expected script to be present in complex HTML")
	}

	headEnd := strings.Index(result, "</head>")
	scriptPos := strings.Index(result, script)

	if scriptPos > headEnd {
		t.Error("Expected script to be before </head> in complex HTML")
	}
}

// TestInjectScript_PreservesExistingContent tests that existing content is preserved
func TestInjectScript_PreservesExistingContent(t *testing.T) {
	html := `<html>
<head>
    <title>Test Page</title>
    <meta charset="UTF-8">
</head>
<body>
    <h1>Heading</h1>
    <p>Paragraph</p>
</body>
</html>`
	script := `<script src="/test.js"></script>`

	result := inject.InjectScript(html, script)

	// Check that all original content is still present
	originalContent := []string{
		"<title>Test Page</title>",
		`<meta charset="UTF-8">`,
		"<h1>Heading</h1>",
		"<p>Paragraph</p>",
	}

	for _, content := range originalContent {
		if !strings.Contains(result, content) {
			t.Errorf("Expected original content '%s' to be preserved", content)
		}
	}
}

// TestInjectScript_MultipleScriptInjections tests injecting script multiple times
func TestInjectScript_MultipleScriptInjections(t *testing.T) {
	html := `<html><head></head><body>content</body></html>`
	script1 := `<script src="/test1.js"></script>`
	script2 := `<script src="/test2.js"></script>`

	result1 := inject.InjectScript(html, script1)
	result2 := inject.InjectScript(result1, script2)

	if !strings.Contains(result2, script1) {
		t.Error("Expected first script to be present")
	}
	if !strings.Contains(result2, script2) {
		t.Error("Expected second script to be present")
	}
}

// TestInjectScript_HandlesScriptWithAttributes tests complex script tags
func TestInjectScript_HandlesScriptWithAttributes(t *testing.T) {
	html := `<html><head></head><body>content</body></html>`
	script := `<script type="module" src="/test.js" defer crossorigin="anonymous"></script>`

	result := inject.InjectScript(html, script)

	if !strings.Contains(result, "type=\"module\"") {
		t.Error("Expected script type attribute to be preserved")
	}
	if !strings.Contains(result, "defer") {
		t.Error("Expected defer attribute to be preserved")
	}
	if !strings.Contains(result, "crossorigin=\"anonymous\"") {
		t.Error("Expected crossorigin attribute to be preserved")
	}
}

// TestInjectScript_HandlesInlineScript tests inline script injection
func TestInjectScript_HandlesInlineScript(t *testing.T) {
	html := `<html><head></head><body>content</body></html>`
	script := `<script>console.log('test');</script>`

	result := inject.InjectScript(html, script)

	if !strings.Contains(result, "console.log('test');") {
		t.Error("Expected inline script content to be preserved")
	}
}

// TestInjectScript_FallbackOnInvalidHTML tests fallback behavior on invalid HTML
func TestInjectScript_FallbackOnInvalidHTML(t *testing.T) {
	// Malformed HTML that might cause parsing issues
	html := `<html><head><title>Test</head><body>content</body></html>`
	script := `<script src="/test.js"></script>`

	result := inject.InjectScript(html, script)

	// Should still inject script somehow (via fallback)
	if !strings.Contains(result, "/test.js") {
		t.Error("Expected script to be injected even with malformed HTML")
	}
}

// TestInjectScript_HandlesEmptyScript tests injection of empty script
func TestInjectScript_HandlesEmptyScript(t *testing.T) {
	html := `<html><head></head><body>content</body></html>`
	script := ``

	result := inject.InjectScript(html, script)

	// Should not crash, result should be similar to original
	if !strings.Contains(result, "<html>") {
		t.Error("Expected HTML structure to be preserved with empty script")
	}
}

// TestInjectScript_PreservesDoctype tests that DOCTYPE is preserved
func TestInjectScript_PreservesDoctype(t *testing.T) {
	html := `<!DOCTYPE html>
<html><head></head><body>content</body></html>`
	script := `<script src="/test.js"></script>`

	result := inject.InjectScript(html, script)

	if !strings.Contains(result, "<!DOCTYPE html>") && !strings.Contains(result, "<!doctype html>") {
		t.Error("Expected DOCTYPE to be preserved")
	}
}

// TestInjectScript_HandlesNestedElements tests injection with nested elements
func TestInjectScript_HandlesNestedElements(t *testing.T) {
	html := `<html>
<head>
    <style>
        body { margin: 0; }
    </style>
    <script>
        var existing = true;
    </script>
</head>
<body>
    <div>
        <span>nested</span>
    </div>
</body>
</html>`
	script := `<script src="/new.js"></script>`

	result := inject.InjectScript(html, script)

	if !strings.Contains(result, "/new.js") {
		t.Error("Expected new script to be injected")
	}
	if !strings.Contains(result, "var existing = true;") {
		t.Error("Expected existing script to be preserved")
	}
	if !strings.Contains(result, "body { margin: 0; }") {
		t.Error("Expected existing style to be preserved")
	}
}

// TestInjectScript_HandlesUnicodeContent tests injection with Unicode characters
func TestInjectScript_HandlesUnicodeContent(t *testing.T) {
	html := `<html><head><title>测试页面</title></head><body>こんにちは</body></html>`
	script := `<script src="/test.js"></script>`

	result := inject.InjectScript(html, script)

	if !strings.Contains(result, "测试页面") {
		t.Error("Expected Unicode title to be preserved")
	}
	if !strings.Contains(result, "こんにちは") {
		t.Error("Expected Unicode body content to be preserved")
	}
	if !strings.Contains(result, "/test.js") {
		t.Error("Expected script to be injected")
	}
}

// TestInjectScript_HandlesHTMLEntities tests injection with HTML entities
func TestInjectScript_HandlesHTMLEntities(t *testing.T) {
	html := `<html><head></head><body>&lt;div&gt;&amp;&nbsp;</body></html>`
	script := `<script src="/test.js"></script>`

	result := inject.InjectScript(html, script)

	if !strings.Contains(result, "/test.js") {
		t.Error("Expected script to be injected")
	}
	// Entities might be normalized by HTML parser, just check script is there
}
