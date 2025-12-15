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
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServeMarkdownAPI_ValidMarkdown(t *testing.T) {
	tests := []struct {
		name          string
		text          string
		shouldContain []string
		description   string
	}{
		{
			name:          "basic markdown",
			text:          "# Hello\n\nThis is **bold**",
			shouldContain: []string{"<h1>", "Hello", "<strong>bold</strong>"},
			description:   "Basic markdown should be converted correctly",
		},
		{
			name:          "empty text",
			text:          "",
			shouldContain: []string{},
			description:   "Empty text should return empty HTML",
		},
		{
			name:          "code blocks",
			text:          "```js\nconst x = 1;\n```",
			shouldContain: []string{"<code>", "const"},
			description:   "Code blocks should be converted",
		},
		{
			name:          "links",
			text:          "[example](https://example.com)",
			shouldContain: []string{"<a", "href=\"https://example.com\""},
			description:   "Links should be converted",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody := MarkdownRequest{Text: tt.text}
			body, err := json.Marshal(reqBody)
			if err != nil {
				t.Fatalf("Failed to marshal request: %v", err)
			}

			req := httptest.NewRequest(http.MethodPost, "/__cem/api/markdown", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			serveMarkdownAPI(rec, req)

			if rec.Code != http.StatusOK {
				t.Errorf("Expected status 200, got %d", rec.Code)
			}

			var resp MarkdownResponse
			if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			for _, expected := range tt.shouldContain {
				if !strings.Contains(resp.HTML, expected) {
					t.Errorf("%s: expected %q in output, got: %s", tt.description, expected, resp.HTML)
				}
			}
		})
	}
}

func TestServeMarkdownAPI_XSSSanitization(t *testing.T) {
	tests := []struct {
		name             string
		text             string
		shouldNotContain string
		description      string
	}{
		{
			name:             "script tag",
			text:             `<script>alert('xss')</script>`,
			shouldNotContain: "<script>",
			description:      "Script tags should be sanitized",
		},
		{
			name:             "javascript protocol",
			text:             `[click](javascript:alert('xss'))`,
			shouldNotContain: "javascript:",
			description:      "JavaScript protocol should be sanitized",
		},
		{
			name:             "event handler",
			text:             `<img src=x onerror="alert('xss')">`,
			shouldNotContain: "onerror",
			description:      "Event handlers should be sanitized",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody := MarkdownRequest{Text: tt.text}
			body, err := json.Marshal(reqBody)
			if err != nil {
				t.Fatalf("Failed to marshal request: %v", err)
			}

			req := httptest.NewRequest(http.MethodPost, "/__cem/api/markdown", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			serveMarkdownAPI(rec, req)

			if rec.Code != http.StatusOK {
				t.Errorf("Expected status 200, got %d", rec.Code)
			}

			var resp MarkdownResponse
			if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			if strings.Contains(resp.HTML, tt.shouldNotContain) {
				t.Errorf("%s: dangerous content not sanitized in output: %s", tt.description, resp.HTML)
			}
		})
	}
}

func TestServeMarkdownAPI_InvalidMethods(t *testing.T) {
	methods := []string{http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodPatch}

	for _, method := range methods {
		t.Run(method, func(t *testing.T) {
			req := httptest.NewRequest(method, "/__cem/api/markdown", nil)
			rec := httptest.NewRecorder()

			serveMarkdownAPI(rec, req)

			if rec.Code != http.StatusMethodNotAllowed {
				t.Errorf("Expected status 405 for %s, got %d", method, rec.Code)
			}
		})
	}
}

func TestServeMarkdownAPI_InvalidJSON(t *testing.T) {
	tests := []struct {
		name        string
		body        string
		description string
	}{
		{
			name:        "invalid json",
			body:        `{invalid}`,
			description: "Malformed JSON should return 400",
		},
		{
			name:        "wrong structure",
			body:        `{"wrong": "field"}`,
			description: "Wrong JSON structure should be handled gracefully",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/__cem/api/markdown", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			serveMarkdownAPI(rec, req)

			if rec.Code != http.StatusBadRequest && rec.Code != http.StatusOK {
				t.Errorf("Expected status 400 or 200 for %s, got %d", tt.description, rec.Code)
			}
		})
	}
}

func TestServeMarkdownAPI_EmptyBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/__cem/api/markdown", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	serveMarkdownAPI(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 for empty body, got %d", rec.Code)
	}
}
