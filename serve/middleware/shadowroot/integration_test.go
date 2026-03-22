/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
*/

package shadowroot_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"bennypowers.dev/cem/serve/middleware/shadowroot"
)

type testLogger struct{}

func (testLogger) Debug(string, ...any) {}
func (testLogger) Info(string, ...any)  {}
func (testLogger) Warning(string, ...any) {}
func (testLogger) Error(string, ...any) {}

// mockSSR is a simple SSR renderer for testing that wraps custom elements
// with a <template shadowrootmode="open"> tag.
type mockSSR struct{}

func (m *mockSSR) RenderHTML(_ context.Context, html string) (string, error) {
	// Simple mock: replace <my-element> with DSD-injected version
	html = strings.ReplaceAll(html,
		"<my-element>",
		`<my-element><template shadowrootmode="open"><slot></slot></template>`)
	return html, nil
}

func (m *mockSSR) Close(_ context.Context) error { return nil }

func TestMiddleware_InjectsDSD(t *testing.T) {
	renderer := &mockSSR{}
	mw := shadowroot.New(testLogger{}, renderer)

	handler := mw(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("<html><body><my-element>hello</my-element></body></html>"))
	}))

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	body := rec.Body.String()
	if !strings.Contains(body, `shadowrootmode="open"`) {
		t.Errorf("expected DSD injection, got: %s", body)
	}
}

func TestMiddleware_SkipsNonHTML(t *testing.T) {
	renderer := &mockSSR{}
	mw := shadowroot.New(testLogger{}, renderer)

	handler := mw(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/javascript")
		w.Write([]byte("console.log('hello')"))
	}))

	req := httptest.NewRequest("GET", "/app.js", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Body.String() != "console.log('hello')" {
		t.Errorf("JS should pass through unchanged")
	}
}

func TestMiddleware_SkipsCemRoutes(t *testing.T) {
	renderer := &mockSSR{}
	mw := shadowroot.New(testLogger{}, renderer)

	handler := mw(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("<my-element>hello</my-element>"))
	}))

	req := httptest.NewRequest("GET", "/__cem/elements/foo/foo.js", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	body := rec.Body.String()
	if strings.Contains(body, "shadowrootmode") {
		t.Errorf("/__cem/ routes should not be SSR processed")
	}
}

func TestMiddleware_NilRenderer(t *testing.T) {
	mw := shadowroot.New(testLogger{}, nil)

	handler := mw(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("<my-element>hello</my-element>"))
	}))

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	if rec.Body.String() != "<my-element>hello</my-element>" {
		t.Errorf("nil renderer should pass through unchanged")
	}
}
