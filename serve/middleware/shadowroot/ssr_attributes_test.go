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
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"bennypowers.dev/cem/serve/middleware/shadowroot"
)

//go:embed ssr_testdata/**
var ssrTestFS embed.FS

var update = flag.Bool("update", false, "update golden files")

// TestSSR_ButtonVariants tests pf-v6-button SSR output with various attribute combinations
func TestSSR_ButtonVariants(t *testing.T) {
	tests := []struct {
		name       string
		html       string
		goldenFile string
	}{
		{
			name:       "button_without_href",
			html:       `<html><pf-v6-button variant="primary">Click me</pf-v6-button></html>`,
			goldenFile: "pf-v6-button-no-href.html",
		},
		{
			name:       "button_with_href",
			html:       `<html><pf-v6-button href="https://example.com" variant="primary">Go there</pf-v6-button></html>`,
			goldenFile: "pf-v6-button-with-href.html",
		},
		{
			name:       "button_disabled",
			html:       `<html><pf-v6-button aria-disabled="true" variant="secondary">Disabled</pf-v6-button></html>`,
			goldenFile: "pf-v6-button-disabled.html",
		},
		{
			name:       "button_with_aria",
			html:       `<html><pf-v6-button aria-label="Close dialog" variant="plain">✕</pf-v6-button></html>`,
			goldenFile: "pf-v6-button-with-aria.html",
		},
		{
			name:       "button_link_with_aria",
			html:       `<html><pf-v6-button href="/docs" aria-label="Documentation" variant="link">Docs</pf-v6-button></html>`,
			goldenFile: "pf-v6-button-link-with-aria.html",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := &mockLogger{}
			broadcaster := &mockBroadcaster{}

			mw := shadowroot.NewWithPrefix(logger, broadcaster, ssrTestFS, "ssr_testdata/templates/elements", realTemplateRenderer)

			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/html")
				w.Write([]byte(tt.html))
			})

			wrapped := mw(handler)
			req := httptest.NewRequest("GET", "/", nil)
			rec := httptest.NewRecorder()
			wrapped.ServeHTTP(rec, req)

			got := rec.Body.String()
			goldenPath := filepath.Join("ssr_testdata", "golden", tt.goldenFile)

			if *update {
				// Update golden file
				if err := os.MkdirAll(filepath.Dir(goldenPath), 0755); err != nil {
					t.Fatalf("Failed to create golden dir: %v", err)
				}
				if err := os.WriteFile(goldenPath, []byte(got), 0644); err != nil {
					t.Fatalf("Failed to write golden file: %v", err)
				}
				t.Logf("Updated golden file: %s", goldenPath)
				return
			}

			// Compare with golden file
			want, err := os.ReadFile(goldenPath)
			if err != nil {
				t.Fatalf("Failed to read golden file: %v\nRun with -update to create it", err)
			}

			if got != string(want) {
				t.Errorf("Output doesn't match golden file\nGot:\n%s\n\nWant:\n%s", got, string(want))
			}
		})
	}
}

// TestSSR_LabelVariants tests pf-v6-label SSR output with various attribute combinations
func TestSSR_LabelVariants(t *testing.T) {
	tests := []struct {
		name       string
		html       string
		goldenFile string
	}{
		{
			name:       "label_without_href",
			html:       `<html><pf-v6-label color="blue">Info</pf-v6-label></html>`,
			goldenFile: "pf-v6-label-no-href.html",
		},
		{
			name:       "label_with_href",
			html:       `<html><pf-v6-label href="/tags/important" color="red">Important</pf-v6-label></html>`,
			goldenFile: "pf-v6-label-with-href.html",
		},
		{
			name:       "label_compact",
			html:       `<html><pf-v6-label compact color="green">OK</pf-v6-label></html>`,
			goldenFile: "pf-v6-label-compact.html",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := &mockLogger{}
			broadcaster := &mockBroadcaster{}

			mw := shadowroot.NewWithPrefix(logger, broadcaster, ssrTestFS, "ssr_testdata/templates/elements", realTemplateRenderer)

			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/html")
				w.Write([]byte(tt.html))
			})

			wrapped := mw(handler)
			req := httptest.NewRequest("GET", "/", nil)
			rec := httptest.NewRecorder()
			wrapped.ServeHTTP(rec, req)

			got := rec.Body.String()
			goldenPath := filepath.Join("ssr_testdata", "golden", tt.goldenFile)

			if *update {
				// Update golden file
				if err := os.MkdirAll(filepath.Dir(goldenPath), 0755); err != nil {
					t.Fatalf("Failed to create golden dir: %v", err)
				}
				if err := os.WriteFile(goldenPath, []byte(got), 0644); err != nil {
					t.Fatalf("Failed to write golden file: %v", err)
				}
				t.Logf("Updated golden file: %s", goldenPath)
				return
			}

			// Compare with golden file
			want, err := os.ReadFile(goldenPath)
			if err != nil {
				t.Fatalf("Failed to read golden file: %v\nRun with -update to create it", err)
			}

			if got != string(want) {
				t.Errorf("Output doesn't match golden file\nGot:\n%s\n\nWant:\n%s", got, string(want))
			}
		})
	}
}

// TestSSR_FormFieldGroupVariants tests pf-v6-form-field-group SSR output
func TestSSR_FormFieldGroupVariants(t *testing.T) {
	tests := []struct {
		name       string
		html       string
		goldenFile string
	}{
		{
			name:       "field_group_not_expandable",
			html:       `<html><pf-v6-form-field-group toggle-text="Section Title" description="Section description"><p>Content</p></pf-v6-form-field-group></html>`,
			goldenFile: "pf-v6-form-field-group-not-expandable.html",
		},
		{
			name:       "field_group_expandable_collapsed",
			html:       `<html><pf-v6-form-field-group expandable toggle-text="Section Title"><p>Content</p></pf-v6-form-field-group></html>`,
			goldenFile: "pf-v6-form-field-group-expandable-collapsed.html",
		},
		{
			name:       "field_group_expandable_expanded",
			html:       `<html><pf-v6-form-field-group expandable expanded toggle-text="Section Title"><p>Content</p></pf-v6-form-field-group></html>`,
			goldenFile: "pf-v6-form-field-group-expandable-expanded.html",
		},
		{
			name:       "field_group_no_toggle",
			html:       `<html><pf-v6-form-field-group><p>Content without header</p></pf-v6-form-field-group></html>`,
			goldenFile: "pf-v6-form-field-group-no-toggle.html",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := &mockLogger{}
			broadcaster := &mockBroadcaster{}

			mw := shadowroot.NewWithPrefix(logger, broadcaster, ssrTestFS, "ssr_testdata/templates/elements", realTemplateRenderer)

			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/html")
				w.Write([]byte(tt.html))
			})

			wrapped := mw(handler)
			req := httptest.NewRequest("GET", "/", nil)
			rec := httptest.NewRecorder()
			wrapped.ServeHTTP(rec, req)

			got := rec.Body.String()
			goldenPath := filepath.Join("ssr_testdata", "golden", tt.goldenFile)

			if *update {
				// Update golden file
				if err := os.MkdirAll(filepath.Dir(goldenPath), 0755); err != nil {
					t.Fatalf("Failed to create golden dir: %v", err)
				}
				if err := os.WriteFile(goldenPath, []byte(got), 0644); err != nil {
					t.Fatalf("Failed to write golden file: %v", err)
				}
				t.Logf("Updated golden file: %s", goldenPath)
				return
			}

			// Compare with golden file
			want, err := os.ReadFile(goldenPath)
			if err != nil {
				t.Fatalf("Failed to read golden file: %v\nRun with -update to create it", err)
			}

			if got != string(want) {
				t.Errorf("Output doesn't match golden file\nGot:\n%s\n\nWant:\n%s", got, string(want))
			}
		})
	}
}

// TestSSR_AttributeInheritance tests that host attributes are correctly preserved
func TestSSR_AttributeInheritance(t *testing.T) {
	logger := &mockLogger{}
	broadcaster := &mockBroadcaster{}

	mw := shadowroot.NewWithPrefix(logger, broadcaster, ssrTestFS, "ssr_testdata/templates/elements", realTemplateRenderer)

	// Test that data-* and class attributes on host are preserved
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<html><pf-v6-button class="custom-class" data-testid="my-button" variant="primary">Click</pf-v6-button></html>`))
	})

	wrapped := mw(handler)
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	body := rec.Body.String()

	// Host attributes should be preserved
	if !strings.Contains(body, `class="custom-class"`) {
		t.Error("Expected host class attribute to be preserved")
	}

	if !strings.Contains(body, `data-testid="my-button"`) {
		t.Error("Expected host data-testid attribute to be preserved")
	}

	// Shadow root should still be injected
	if !strings.Contains(body, `<template shadowrootmode="open">`) {
		t.Error("Expected shadow root to be injected")
	}
}

// mockLogger and mockBroadcaster implementations
type mockLogger struct {
	warnings []string
}

func (m *mockLogger) Debug(format string, args ...interface{}) {}
func (m *mockLogger) Info(format string, args ...interface{})  {}
func (m *mockLogger) Warn(format string, args ...interface{}) {
	m.warnings = append(m.warnings, fmt.Sprintf(format, args...))
}
func (m *mockLogger) Error(format string, args ...interface{}) {}

type mockBroadcaster struct {
	broadcasts []string
}

func (m *mockBroadcaster) Broadcast(event string, data interface{}) {
	m.broadcasts = append(m.broadcasts, event)
}
