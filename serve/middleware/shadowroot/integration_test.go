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

package shadowroot_test

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"bennypowers.dev/cem/serve/middleware/shadowroot"
)

//go:embed integration_testdata/**
var integrationFS embed.FS

// realTemplateRenderer uses actual Go template rendering to test attribute handling
func realTemplateRenderer(elementName string, data interface{}) (string, error) {
	// Read template
	templatePath := fmt.Sprintf("integration_testdata/templates/elements/%s/%s.html", elementName, elementName)
	templateContent, err := integrationFS.ReadFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("template not found: %s", elementName)
	}

	// Read CSS
	cssPath := fmt.Sprintf("integration_testdata/templates/elements/%s/%s.css", elementName, elementName)
	cssContent, err := integrationFS.ReadFile(cssPath)
	if err != nil {
		cssContent = []byte("/* no styles */")
	}

	// Parse template
	tmpl, err := template.New(elementName).Parse(string(templateContent))
	if err != nil {
		return "", err
	}

	// Execute template with data
	var shadowBuf bytes.Buffer
	if err := tmpl.Execute(&shadowBuf, data); err != nil {
		return "", err
	}

	// Wrap in declarative shadow DOM
	return fmt.Sprintf(`<template shadowrootmode="open"><style>%s</style>%s</template>`,
		string(cssContent), shadowBuf.String()), nil
}

func TestIntegration_AttributeMap(t *testing.T) {
	logger := &mockLogger{}
	broadcaster := &mockBroadcaster{}

	mw := shadowroot.NewWithPrefix(logger, broadcaster, integrationFS, "integration_testdata/templates/elements", realTemplateRenderer)

	// Handler returns button with variant attribute
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(`<html><attr-button variant="primary" disabled></attr-button></html>`))
	})

	wrapped := mw(handler)
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	body := rec.Body.String()

	// Verify template used .Attributes map
	if !strings.Contains(body, `data-variant="primary"`) {
		t.Errorf("Expected template to use .Attributes map\nBody: %s", body)
	}

	// Verify template used boolean attribute
	if !strings.Contains(body, `data-disabled="true"`) {
		t.Errorf("Expected template to handle boolean attribute\nBody: %s", body)
	}
}

func TestIntegration_CamelCaseProperties(t *testing.T) {
	logger := &mockLogger{}
	broadcaster := &mockBroadcaster{}

	mw := shadowroot.NewWithPrefix(logger, broadcaster, integrationFS, "integration_testdata/templates/elements", realTemplateRenderer)

	// Handler returns button with aria-label
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(`<html><camel-button aria-label="Click me" data-test-id="btn-1"></camel-button></html>`))
	})

	wrapped := mw(handler)
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	body := rec.Body.String()

	// Verify template used .AriaLabel camelCase property
	if !strings.Contains(body, `aria-label="Click me"`) {
		t.Errorf("Expected template to use .AriaLabel property\nBody: %s", body)
	}

	// Verify template used .DataTestId camelCase property
	if !strings.Contains(body, `data-test-id="btn-1"`) {
		t.Errorf("Expected template to use .DataTestId property\nBody: %s", body)
	}
}

func TestIntegration_ConditionalTemplate(t *testing.T) {
	logger := &mockLogger{}
	broadcaster := &mockBroadcaster{}

	mw := shadowroot.NewWithPrefix(logger, broadcaster, integrationFS, "integration_testdata/templates/elements", realTemplateRenderer)

	tests := []struct {
		name        string
		html        string
		contains    string
		notContains string
	}{
		{
			name:     "with icon attribute",
			html:     `<html><conditional-button icon="check"></conditional-button></html>`,
			contains: `<span class="icon">check</span>`,
		},
		{
			name:        "without icon attribute",
			html:        `<html><conditional-button></conditional-button></html>`,
			notContains: `<span class="icon">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/html")
				_, _ = w.Write([]byte(tt.html))
			})

			wrapped := mw(handler)
			req := httptest.NewRequest("GET", "/", nil)
			rec := httptest.NewRecorder()
			wrapped.ServeHTTP(rec, req)

			body := rec.Body.String()

			if tt.contains != "" && !strings.Contains(body, tt.contains) {
				t.Errorf("Expected body to contain %q\nBody: %s", tt.contains, body)
			}

			if tt.notContains != "" && strings.Contains(body, tt.notContains) {
				t.Errorf("Expected body NOT to contain %q\nBody: %s", tt.notContains, body)
			}
		})
	}
}

func TestIntegration_NestedWithAttributes(t *testing.T) {
	logger := &mockLogger{}
	broadcaster := &mockBroadcaster{}

	mw := shadowroot.NewWithPrefix(logger, broadcaster, integrationFS, "integration_testdata/templates/elements", realTemplateRenderer)

	// Parent contains child with attributes in its shadow template
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(`<html><parent-element title="Parent"></parent-element></html>`))
	})

	wrapped := mw(handler)
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	body := rec.Body.String()

	// Verify parent's shadow root was injected
	if !strings.Contains(body, `<h2>Parent</h2>`) {
		t.Errorf("Expected parent template to use title attribute\nBody: %s", body)
	}

	// Verify child's shadow root was injected with attributes from parent's template
	if !strings.Contains(body, `data-variant="secondary"`) {
		t.Errorf("Expected child template to use variant attribute\nBody: %s", body)
	}

	// Count shadow roots - should be 2 (parent + child)
	count := strings.Count(body, `<template shadowrootmode="open">`)
	if count != 2 {
		t.Errorf("Expected 2 shadow roots, got %d\nBody: %s", count, body)
	}
}

func TestIntegration_MissingTemplate(t *testing.T) {
	logger := &mockLogger{}
	broadcaster := &mockBroadcaster{}

	mw := shadowroot.NewWithPrefix(logger, broadcaster, integrationFS, "integration_testdata/templates/elements", realTemplateRenderer)

	// Use an element that's NOT in the known elements set
	// It should be silently skipped (no warning/broadcast)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(`<html><unknown-element foo="bar"></unknown-element></html>`))
	})

	wrapped := mw(handler)
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	body := rec.Body.String()

	// Element should pass through unchanged (not in known elements)
	if strings.Contains(body, `<template shadowrootmode="open">`) {
		t.Error("Did not expect shadow root for unknown element")
	}

	// Element tag should remain
	if !strings.Contains(body, `<unknown-element`) {
		t.Error("Expected element tag to be preserved")
	}

	// No warnings/broadcasts for unknown elements (they're just skipped)
	if len(logger.warnings) > 0 {
		t.Errorf("Did not expect warnings for unknown element, got: %v", logger.warnings)
	}

	if len(broadcaster.broadcasts) > 0 {
		t.Errorf("Did not expect broadcasts for unknown element, got: %v", broadcaster.broadcasts)
	}
}
