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
	"embed"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"bennypowers.dev/cem/serve/middleware/shadowroot"
)

//go:embed testdata/**
var testFS embed.FS

// mockLogger implements types.Logger for testing
type mockLogger struct {
	warnings []string
	errors   []string
}

func (m *mockLogger) Debug(msg string, args ...any)   {}
func (m *mockLogger) Info(msg string, args ...any)    {}
func (m *mockLogger) Warning(msg string, args ...any) { m.warnings = append(m.warnings, msg) }
func (m *mockLogger) Error(msg string, args ...any)   { m.errors = append(m.errors, msg) }

// mockBroadcaster implements shadowroot.ErrorBroadcaster for testing
type mockBroadcaster struct {
	broadcasts []struct {
		title    string
		message  string
		filename string
	}
}

func (m *mockBroadcaster) BroadcastError(title, message, filename string) {
	m.broadcasts = append(m.broadcasts, struct {
		title    string
		message  string
		filename string
	}{title, message, filename})
}

// mockRenderer implements ShadowRootRenderer for testing
func mockRenderer(elementName string, data interface{}) (string, error) {
	// Simple mock: returns a basic shadow root template
	return `<template shadowrootmode="open"><style>/* styles */</style><slot></slot></template>`, nil
}

// mockRendererWithNested returns shadow roots that may contain nested custom elements
func mockRendererWithNested(elementName string, data interface{}) (string, error) {
	switch elementName {
	case "outer-element":
		// Contains an inner-element
		return `<template shadowrootmode="open"><style>/* outer */</style><inner-element></inner-element></template>`, nil
	case "inner-element":
		// Simple shadow root
		return `<template shadowrootmode="open"><style>/* inner */</style><slot></slot></template>`, nil
	default:
		return `<template shadowrootmode="open"><slot></slot></template>`, nil
	}
}

func TestShadowRootInjection_CustomElement(t *testing.T) {
	logger := &mockLogger{}
	broadcaster := &mockBroadcaster{}

	mw := shadowroot.NewWithPrefix(logger, broadcaster, testFS, "testdata/templates/elements", mockRenderer)

	// Handler that returns HTML with a custom element
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(`<html><body><test-button></test-button></body></html>`))
	})

	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	wrapped.ServeHTTP(rec, req)

	body := rec.Body.String()

	// Verify shadow root was injected
	if !strings.Contains(body, `<template shadowrootmode="open">`) {
		t.Error("Expected shadow root template to be injected")
	}

	if !strings.Contains(body, `<test-button>`) {
		t.Error("Expected custom element tag to remain")
	}

	if !strings.Contains(body, `<slot></slot>`) {
		t.Error("Expected shadow root content to be present")
	}
}

func TestShadowRootInjection_NonHTML(t *testing.T) {
	logger := &mockLogger{}
	broadcaster := &mockBroadcaster{}

	mw := shadowroot.NewWithPrefix(logger, broadcaster, testFS, "testdata/templates/elements", mockRenderer)

	// Handler that returns JSON
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"key": "value"}`))
	})

	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	wrapped.ServeHTTP(rec, req)

	body := rec.Body.String()

	// JSON should be unchanged (no shadow root injection)
	if body != `{"key": "value"}` {
		t.Errorf("Expected JSON to be unchanged, got: %s", body)
	}
}

func TestShadowRootInjection_NonCustomElement(t *testing.T) {
	logger := &mockLogger{}
	broadcaster := &mockBroadcaster{}

	mw := shadowroot.NewWithPrefix(logger, broadcaster, testFS, "testdata/templates/elements", mockRenderer)

	// Handler with standard HTML elements (no hyphens)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(`<html><body><div><p>Hello</p></div></body></html>`))
	})

	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	wrapped.ServeHTTP(rec, req)

	body := rec.Body.String()

	// No shadow roots should be injected for standard elements
	if strings.Contains(body, `<template shadowrootmode="open">`) {
		t.Error("Did not expect shadow root for standard HTML elements")
	}

	if body != `<html><body><div><p>Hello</p></div></body></html>` {
		t.Errorf("Expected HTML to be unchanged, got: %s", body)
	}
}

func TestShadowRootInjection_UnknownCustomElement(t *testing.T) {
	logger := &mockLogger{}
	broadcaster := &mockBroadcaster{}

	mw := shadowroot.NewWithPrefix(logger, broadcaster, testFS, "testdata/templates/elements", mockRenderer)

	// Handler with unknown custom element (has hyphen but not in templates)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(`<html><body><unknown-element></unknown-element></body></html>`))
	})

	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	wrapped.ServeHTTP(rec, req)

	body := rec.Body.String()

	// No shadow root should be injected for unknown elements
	if strings.Contains(body, `<template shadowrootmode="open">`) {
		t.Error("Did not expect shadow root for unknown custom element")
	}

	// Element should remain unchanged
	if !strings.Contains(body, `<unknown-element></unknown-element>`) {
		t.Error("Expected unknown element to pass through unchanged")
	}
}

func TestShadowRootInjection_Recursive(t *testing.T) {
	logger := &mockLogger{}
	broadcaster := &mockBroadcaster{}

	mw := shadowroot.NewWithPrefix(logger, broadcaster, testFS, "testdata/templates/elements", mockRendererWithNested)

	// Handler with nested custom elements
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(`<html><body><outer-element></outer-element></body></html>`))
	})

	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	wrapped.ServeHTTP(rec, req)

	body := rec.Body.String()

	// Count shadow root templates - should be 2 (outer + inner)
	count := strings.Count(body, `<template shadowrootmode="open">`)
	if count != 2 {
		t.Errorf("Expected 2 shadow roots (outer + inner), got %d\nBody: %s", count, body)
	}

	// Verify outer element's shadow root
	if !strings.Contains(body, `/* outer */`) {
		t.Error("Expected outer element's styles")
	}

	// Verify inner element's shadow root (nested within outer)
	if !strings.Contains(body, `/* inner */`) {
		t.Error("Expected inner element's styles")
	}

	// Verify both elements are present
	if !strings.Contains(body, `<outer-element>`) {
		t.Error("Expected outer-element tag")
	}

	if !strings.Contains(body, `<inner-element>`) {
		t.Error("Expected inner-element tag (nested in outer's shadow root)")
	}
}

func TestAttributeExtraction(t *testing.T) {
	logger := &mockLogger{}
	broadcaster := &mockBroadcaster{}

	// Renderer that checks attributes are passed correctly
	var capturedData interface{}
	renderer := func(elementName string, data interface{}) (string, error) {
		capturedData = data
		return `<template shadowrootmode="open"><slot></slot></template>`, nil
	}

	mw := shadowroot.NewWithPrefix(logger, broadcaster, testFS, "testdata/templates/elements", renderer)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(`<html><test-button aria-label="Click me" data-foo="bar" disabled></test-button></html>`))
	})

	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	wrapped.ServeHTTP(rec, req)

	// Verify attributes were captured
	if capturedData == nil {
		t.Fatal("Expected data to be passed to renderer")
	}

	dataMap, ok := capturedData.(map[string]interface{})
	if !ok {
		t.Fatalf("Expected data to be map[string]interface{}, got %T", capturedData)
	}

	// Check raw attributes map
	attrs, ok := dataMap["Attributes"].(map[string]string)
	if !ok {
		t.Fatal("Expected Attributes map")
	}

	// All attributes should be PascalCase in the Attributes map
	if attrs["AriaLabel"] != "Click me" {
		t.Errorf("Expected AriaLabel='Click me', got %q", attrs["AriaLabel"])
	}

	if attrs["DataFoo"] != "bar" {
		t.Errorf("Expected DataFoo='bar', got %q", attrs["DataFoo"])
	}

	// Boolean attribute should have empty value
	if attrs["Disabled"] != "" {
		t.Errorf("Expected Disabled='', got %q", attrs["Disabled"])
	}
}

func TestSelfClosingTags(t *testing.T) {
	logger := &mockLogger{}
	broadcaster := &mockBroadcaster{}

	mw := shadowroot.NewWithPrefix(logger, broadcaster, testFS, "testdata/templates/elements", mockRenderer)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(`<html><body><test-icon /></body></html>`))
	})

	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	wrapped.ServeHTTP(rec, req)

	body := rec.Body.String()

	// Shadow root should be injected for self-closing custom elements
	if !strings.Contains(body, `<template shadowrootmode="open">`) {
		t.Error("Expected shadow root for self-closing custom element")
	}

	if !strings.Contains(body, `<test-icon`) {
		t.Error("Expected test-icon element")
	}
}

func TestPassThrough(t *testing.T) {
	logger := &mockLogger{}
	broadcaster := &mockBroadcaster{}

	mw := shadowroot.NewWithPrefix(logger, broadcaster, testFS, "testdata/templates/elements", mockRenderer)

	// Handler with mixed content
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = io.WriteString(w, `<!DOCTYPE html>
<html lang="en">
<head><title>Test</title></head>
<body>
  <!-- Comment -->
  <h1>Hello</h1>
  <test-button>Click</test-button>
  <p>Text</p>
</body>
</html>`)
	})

	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	wrapped.ServeHTTP(rec, req)

	body := rec.Body.String()

	// Verify all content is preserved
	if !strings.Contains(body, `<!DOCTYPE html>`) {
		t.Error("Expected DOCTYPE to be preserved")
	}

	if !strings.Contains(body, `<!-- Comment -->`) {
		t.Error("Expected comment to be preserved")
	}

	if !strings.Contains(body, `<h1>Hello</h1>`) {
		t.Error("Expected standard elements to be preserved")
	}

	if !strings.Contains(body, `<test-button>`) {
		t.Error("Expected custom element to be present")
	}

	if !strings.Contains(body, `<template shadowrootmode="open">`) {
		t.Error("Expected shadow root to be injected")
	}
}
