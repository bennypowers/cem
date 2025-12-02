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

package routes

import (
	"strings"
	"testing"
)

func TestMarkdownToHTML_XSSSanitization(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		shouldNotContain string
		description    string
	}{
		{
			name:           "script tag injection",
			input:          `<script>alert('xss')</script>`,
			shouldNotContain: "<script>",
			description:    "Script tags should be stripped to prevent XSS",
		},
		{
			name:           "img onerror injection",
			input:          `<img src=x onerror="alert('xss')">`,
			shouldNotContain: "onerror",
			description:    "Event handlers should be stripped from tags",
		},
		{
			name:           "javascript protocol in link",
			input:          `[click me](javascript:alert('xss'))`,
			shouldNotContain: "javascript:",
			description:    "JavaScript protocol in links should be stripped",
		},
		{
			name:           "onclick event handler",
			input:          `<a href="#" onclick="alert('xss')">click</a>`,
			shouldNotContain: "onclick",
			description:    "Event handlers should be stripped",
		},
		{
			name:           "iframe injection",
			input:          `<iframe src="evil.com"></iframe>`,
			shouldNotContain: "<iframe",
			description:    "iframes should be stripped",
		},
		{
			name:           "object tag injection",
			input:          `<object data="evil.swf"></object>`,
			shouldNotContain: "<object",
			description:    "Object tags should be stripped",
		},
		{
			name:           "embed tag injection",
			input:          `<embed src="evil.swf">`,
			shouldNotContain: "<embed",
			description:    "Embed tags should be stripped",
		},
		{
			name:           "data URI in image",
			input:          `<img src="data:text/html,<script>alert('xss')</script>">`,
			shouldNotContain: "data:text/html",
			description:    "Dangerous data URIs should be stripped",
		},
		{
			name:           "style tag injection",
			input:          `<style>body { background: url('javascript:alert(1)'); }</style>`,
			shouldNotContain: "<style>",
			description:    "Style tags should be stripped",
		},
		{
			name:           "mixed markdown and XSS",
			input:          `# Heading\n\n<script>alert('xss')</script>\n\nNormal text`,
			shouldNotContain: "<script>",
			description:    "XSS should be stripped even when mixed with valid markdown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := markdownToHTML(tt.input)
			if err != nil {
				t.Fatalf("markdownToHTML returned error: %v", err)
			}

			if strings.Contains(result, tt.shouldNotContain) {
				t.Errorf("%s: dangerous content not sanitized\nInput: %s\nOutput: %s\nShould not contain: %s",
					tt.description, tt.input, result, tt.shouldNotContain)
			}
		})
	}
}

func TestMarkdownToHTML_SafeMarkdown(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		shouldContain []string
		description string
	}{
		{
			name:  "basic markdown heading",
			input: "# Hello World",
			shouldContain: []string{"<h1>", "Hello World", "</h1>"},
			description: "Basic markdown should render correctly",
		},
		{
			name:  "emphasis and strong",
			input: "This is *italic* and **bold**",
			shouldContain: []string{"<em>italic</em>", "<strong>bold</strong>"},
			description: "Emphasis markdown should render correctly",
		},
		{
			name:  "code blocks",
			input: "```js\nconst x = 1;\n```",
			shouldContain: []string{"<code>", "const", "</code>"},
			description: "Code blocks should render correctly",
		},
		{
			name:  "safe links",
			input: "[example](https://example.com)",
			shouldContain: []string{"<a", "href=\"https://example.com\"", "example</a>"},
			description: "Safe links should render correctly",
		},
		{
			name:  "lists",
			input: "- Item 1\n- Item 2",
			shouldContain: []string{"<ul>", "<li>Item 1</li>", "<li>Item 2</li>", "</ul>"},
			description: "Lists should render correctly",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := markdownToHTML(tt.input)
			if err != nil {
				t.Fatalf("markdownToHTML returned error: %v", err)
			}

			for _, expected := range tt.shouldContain {
				if !strings.Contains(result, expected) {
					t.Errorf("%s: expected content not found\nInput: %s\nOutput: %s\nShould contain: %s",
						tt.description, tt.input, result, expected)
				}
			}
		})
	}
}

func TestMarkdownToHTML_EmptyInput(t *testing.T) {
	result, err := markdownToHTML("")
	if err != nil {
		t.Fatalf("markdownToHTML returned error for empty input: %v", err)
	}
	if result != "" {
		t.Errorf("Expected empty output for empty input, got: %s", result)
	}
}
