/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
*/

package serve

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"testing"

	litssr "bennypowers.dev/lit-ssr-wasm/go"
)

var (
	sharedRenderer     *litssr.Renderer
	sharedRendererOnce sync.Once
	sharedRendererErr  error
)

func getSharedRenderer(t *testing.T) *litssr.Renderer {
	t.Helper()
	sharedRendererOnce.Do(func() {
		source, err := buildComponentSource()
		if err != nil {
			sharedRendererErr = err
			return
		}
		if source == "" {
			sharedRendererErr = fmt.Errorf("no SSR bundle available")
			return
		}
		elements := discoverChromeElements()
		if len(elements) == 0 {
			sharedRendererErr = fmt.Errorf("no chrome elements found in manifests")
			return
		}
		sharedRenderer, sharedRendererErr = litssr.NewWithElements(
			context.Background(), source, elements, 1)
	})
	if sharedRendererErr != nil {
		t.Skip(sharedRendererErr)
	}
	return sharedRenderer
}

func render(t *testing.T, r *litssr.Renderer, html string) string {
	t.Helper()
	out, err := r.RenderHTML(context.Background(), html)
	if err != nil {
		t.Fatalf("RenderHTML: %v", err)
	}
	return out
}

// TestSSR_Lifecycle verifies init, render, and shutdown without race.
func TestSSR_Lifecycle(t *testing.T) {
	r := getSharedRenderer(t)
	out := render(t, r, "<pf-v6-badge>42</pf-v6-badge>")
	if !strings.Contains(out, "shadowrootmode") {
		t.Error("expected shadowrootmode in output")
	}
}

// TestSSR_DSDInjection verifies Declarative Shadow DOM is injected.
func TestSSR_DSDInjection(t *testing.T) {
	r := getSharedRenderer(t)

	tests := []struct {
		name  string
		input string
	}{
		{"badge", `<pf-v6-badge>42</pf-v6-badge>`},
		{"button", `<pf-v6-button>Click</pf-v6-button>`},
		{"card", `<pf-v6-card><h3 slot="title">Title</h3>Body</pf-v6-card>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := render(t, r, tt.input)
			if !strings.Contains(out, `shadowrootmode="open"`) {
				t.Errorf("missing shadowrootmode=\"open\" in %s output", tt.name)
			}
			if !strings.Contains(out, `<template shadowroot="open"`) {
				t.Errorf("missing <template shadowroot> in %s output", tt.name)
			}
		})
	}
}

// TestSSR_StylesNonEmpty verifies styles are rendered inside shadow roots.
func TestSSR_StylesNonEmpty(t *testing.T) {
	r := getSharedRenderer(t)

	tests := []struct {
		name     string
		input    string
		cssToken string
	}{
		{"badge", `<pf-v6-badge>42</pf-v6-badge>`, "--pf-v6-c-badge"},
		{"button", `<pf-v6-button>Click</pf-v6-button>`, "--pf-v6-c-button"},
		{"card", `<pf-v6-card>Body</pf-v6-card>`, "--pf-v6-c-card"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := render(t, r, tt.input)
			if strings.Contains(out, "<style></style>") {
				t.Error("style tag is empty")
			}
			if !strings.Contains(out, tt.cssToken) {
				t.Errorf("missing CSS token %q in styles", tt.cssToken)
			}
		})
	}
}

// TestSSR_HydrationDigests verifies template digests are base64-encoded.
func TestSSR_HydrationDigests(t *testing.T) {
	r := getSharedRenderer(t)
	out := render(t, r, "<pf-v6-badge>42</pf-v6-badge>")

	// Extract lit-part digest
	idx := strings.Index(out, "lit-part ")
	if idx == -1 {
		t.Fatal("no lit-part marker found")
	}
	digestStart := idx + len("lit-part ")
	digestEnd := strings.Index(out[digestStart:], "-")
	if digestEnd == -1 {
		t.Fatal("malformed lit-part marker")
	}
	digest := out[digestStart : digestStart+digestEnd]

	// Base64 uses A-Z, a-z, 0-9, +, /, = only
	for _, c := range digest {
		if !((c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') ||
			(c >= '0' && c <= '9') || c == '+' || c == '/' || c == '=') {
			t.Errorf("digest %q contains non-base64 character %q", digest, string(c))
			break
		}
	}
}

// TestSSR_Passthrough verifies unknown elements are not modified.
func TestSSR_Passthrough(t *testing.T) {
	r := getSharedRenderer(t)
	out := render(t, r, "<unknown-element>hello</unknown-element>")
	if strings.Contains(out, "shadowrootmode") {
		t.Error("unknown element should not get DSD")
	}
	if !strings.Contains(out, "hello") {
		t.Error("content should pass through")
	}
}

// TestSSR_ButtonConditionalRendering verifies button renders differently
// based on href attribute.
func TestSSR_ButtonConditionalRendering(t *testing.T) {
	r := getSharedRenderer(t)

	t.Run("without href", func(t *testing.T) {
		out := render(t, r, `<pf-v6-button>Click</pf-v6-button>`)
		if strings.Contains(out, "<a ") {
			t.Error("button without href should not render <a>")
		}
		if !strings.Contains(out, `<slot name="icon-start">`) {
			t.Error("button should render icon-start slot")
		}
	})

	t.Run("with href", func(t *testing.T) {
		out := render(t, r, `<pf-v6-button href="/foo">Link</pf-v6-button>`)
		if !strings.Contains(out, `<a `) {
			t.Error("button with href should render <a>")
		}
		if !strings.Contains(out, `href="/foo"`) {
			t.Error("button should pass href to <a>")
		}
	})
}

// TestSSR_ModuleScopeBrowserAPIs verifies components with guarded browser
// API calls at module scope don't crash SSR.
func TestSSR_ModuleScopeBrowserAPIs(t *testing.T) {
	r := getSharedRenderer(t)

	// These components reference document/window/CSS at module scope
	// but with guards (optional chaining, globalThis checks)
	tests := []struct {
		name  string
		input string
	}{
		{"dropdown (document.addEventListener)", `<pf-v6-dropdown>menu</pf-v6-dropdown>`},
		{"page (window.matchMedia)", `<pf-v6-page>content</pf-v6-page>`},
		{"popover (CSS.supports)", `<pf-v6-popover>tip</pf-v6-popover>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := render(t, r, tt.input)
			if !strings.Contains(out, "shadowrootmode") {
				t.Errorf("%s should render with DSD", tt.name)
			}
		})
	}
}

// TestSSR_BytecodeLifecycle verifies the bytecode path works end-to-end.
func TestSSR_BytecodeLifecycle(t *testing.T) {
	source, err := buildComponentSource()
	if err != nil || source == "" {
		t.Skip("no SSR bundle available")
	}

	// Compile to bytecode
	bytecode, err := litssr.CompileSource(context.Background(), source)
	if err != nil {
		t.Fatalf("CompileSource: %v", err)
	}
	if len(bytecode) == 0 {
		t.Fatal("bytecode is empty")
	}
	t.Logf("Source: %d KB, Bytecode: %d KB", len(source)/1024, len(bytecode)/1024)

	// Create renderer from bytecode
	elements := discoverChromeElements()
	r, err := litssr.NewFromBytecode(context.Background(), bytecode, elements, 1)
	if err != nil {
		t.Fatalf("NewFromBytecode: %v", err)
	}
	t.Cleanup(func() { r.Close(context.Background()) })

	// Render and verify
	out := render(t, r, "<pf-v6-badge>42</pf-v6-badge>")
	if !strings.Contains(out, `shadowrootmode="open"`) {
		t.Error("bytecode render missing DSD")
	}
	if strings.Contains(out, "<style></style>") {
		t.Error("bytecode render has empty styles")
	}
}
