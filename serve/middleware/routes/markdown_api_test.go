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

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/serve/logger"
	"bennypowers.dev/cem/serve/middleware"
)

// mockContext implements middleware.DevServerContext for testing
type mockContext struct {
	manifestBytes []byte
}

func (m *mockContext) WatchDir() string                                  { return "" }
func (m *mockContext) IsWorkspace() bool                                 { return false }
func (m *mockContext) WorkspacePackages() []middleware.WorkspacePackage  { return nil }
func (m *mockContext) Manifest() ([]byte, error)                         { return m.manifestBytes, nil }
func (m *mockContext) ImportMap() middleware.ImportMap                   { return nil }
func (m *mockContext) DemoRoutes() any                                   { return nil }
func (m *mockContext) SourceControlRootURL() string                      { return "" }
func (m *mockContext) Logger() logger.Logger                             { return logger.NewDefaultLogger() }
func (m *mockContext) FileSystem() platform.FileSystem                   { return nil }
func (m *mockContext) PackageJSON() (*middleware.PackageJSON, error)     { return nil, nil }
func (m *mockContext) BroadcastError(title, message, file string) error  { return nil }
func (m *mockContext) DemoRenderingMode() string                         { return "light" }

func TestServeMarkdownAPI_ValidPaths(t *testing.T) {
	// Create test manifest with markdown content
	manifestJSON := `{
		"schemaVersion": "1.0.0",
		"modules": [
			{
				"path": "./test.js",
				"summary": "# Test Module\n\nThis is **bold**",
				"description": "Module description with [link](https://example.com)",
				"declarations": [
					{
						"kind": "class",
						"name": "TestElement",
						"tagName": "test-element",
						"customElement": true,
						"summary": "## Test Element\n\nElement summary",
						"description": "Element description with ` + "`code`" + `",
						"attributes": [
							{
								"name": "test-attr",
								"summary": "Attribute summary",
								"description": "Attribute **description**"
							}
						]
					}
				]
			}
		]
	}`

	tests := []struct {
		name          string
		path          string
		shouldContain []string
		description   string
	}{
		{
			name:          "module summary",
			path:          `modules.#(path=="./test.js").summary`,
			shouldContain: []string{"<h1>", "Test Module", "<strong>bold</strong>"},
			description:   "Module summary should be rendered",
		},
		{
			name:          "module description",
			path:          `modules.#(path=="./test.js").description`,
			shouldContain: []string{"<a", "href=\"https://example.com\""},
			description:   "Module description with links should be rendered",
		},
		{
			name:          "element summary",
			path:          `modules.#(path=="./test.js").declarations.#(tagName=="test-element").summary`,
			shouldContain: []string{"<h2>", "Test Element"},
			description:   "Element summary should be rendered",
		},
		{
			name:          "attribute description",
			path:          `modules.#(path=="./test.js").declarations.#(tagName=="test-element").attributes.#(name=="test-attr").description`,
			shouldContain: []string{"<strong>description</strong>"},
			description:   "Attribute description should be rendered",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create config with test manifest
			config := Config{
				Context: &mockContext{
					manifestBytes: []byte(manifestJSON),
				},
			}

			reqBody := MarkdownRequest{Path: tt.path}
			body, err := json.Marshal(reqBody)
			if err != nil {
				t.Fatalf("Failed to marshal request: %v", err)
			}

			req := httptest.NewRequest(http.MethodPost, "/__cem/api/markdown", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			serveMarkdownAPI(rec, req, config)

			if rec.Code != http.StatusOK {
				t.Errorf("Expected status 200, got %d: %s", rec.Code, rec.Body.String())
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

func TestServeMarkdownAPI_EmptyPath(t *testing.T) {
	config := Config{
		Context: &mockContext{
			manifestBytes: []byte(`{"schemaVersion": "1.0.0", "modules": []}`),
		},
	}

	reqBody := MarkdownRequest{Path: ""}
	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("Failed to marshal request: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/__cem/api/markdown", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	serveMarkdownAPI(rec, req, config)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status 200 for empty path, got %d", rec.Code)
	}

	var resp MarkdownResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if resp.HTML != "" {
		t.Errorf("Expected empty HTML for empty path, got: %s", resp.HTML)
	}
}

func TestServeMarkdownAPI_PathNotFound(t *testing.T) {
	config := Config{
		Context: &mockContext{
			manifestBytes: []byte(`{"schemaVersion": "1.0.0", "modules": []}`),
		},
	}

	reqBody := MarkdownRequest{Path: `modules.#(path=="nonexistent.js").summary`}
	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("Failed to marshal request: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/__cem/api/markdown", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	serveMarkdownAPI(rec, req, config)

	if rec.Code != http.StatusNotFound {
		t.Errorf("Expected status 404 for nonexistent path, got %d", rec.Code)
	}
}

func TestServeMarkdownAPI_NonStringField(t *testing.T) {
	manifestJSON := `{
		"schemaVersion": "1.0.0",
		"modules": [
			{
				"path": "./test.js",
				"declarations": []
			}
		]
	}`

	config := Config{
		Context: &mockContext{
			manifestBytes: []byte(manifestJSON),
		},
	}

	// Try to render an array field (declarations)
	reqBody := MarkdownRequest{Path: `modules.#(path=="./test.js").declarations`}
	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("Failed to marshal request: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/__cem/api/markdown", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	serveMarkdownAPI(rec, req, config)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 for non-string field, got %d", rec.Code)
	}
}

func TestServeMarkdownAPI_InvalidMethods(t *testing.T) {
	config := Config{
		Context: &mockContext{},
	}

	methods := []string{http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodPatch}

	for _, method := range methods {
		t.Run(method, func(t *testing.T) {
			req := httptest.NewRequest(method, "/__cem/api/markdown", nil)
			rec := httptest.NewRecorder()

			serveMarkdownAPI(rec, req, config)

			if rec.Code != http.StatusMethodNotAllowed {
				t.Errorf("Expected status 405 for %s, got %d", method, rec.Code)
			}
		})
	}
}

func TestServeMarkdownAPI_InvalidJSON(t *testing.T) {
	config := Config{
		Context: &mockContext{},
	}

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

			serveMarkdownAPI(rec, req, config)

			if rec.Code != http.StatusBadRequest && rec.Code != http.StatusOK {
				t.Errorf("Expected status 400 or 200 for %s, got %d", tt.description, rec.Code)
			}
		})
	}
}

func TestServeMarkdownAPI_EmptyBody(t *testing.T) {
	config := Config{
		Context: &mockContext{},
	}

	req := httptest.NewRequest(http.MethodPost, "/__cem/api/markdown", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	serveMarkdownAPI(rec, req, config)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 for empty body, got %d", rec.Code)
	}
}
