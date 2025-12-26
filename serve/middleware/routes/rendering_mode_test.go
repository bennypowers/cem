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
	"html/template"
	"strings"
	"testing"
)

// TestRenderingMode_Default verifies light mode is used by default (when RenderingMode is empty)
func TestRenderingMode_Default(t *testing.T) {
	demoHTML := `<my-element></my-element>`

	rendered, err := renderDemo(testTemplates(), nil, ChromeData{
		TagName:       "my-element",
		DemoTitle:     "Default Rendering",
		DemoHTML:      template.HTML(demoHTML),
		ImportMap:     "{}",
		RenderingMode: "", // Empty should default to light DOM
	})
	if err != nil {
		t.Fatalf("Failed to render chrome: %v", err)
	}

	// Light mode: no <template shadowrootmode="open">
	if strings.Contains(rendered, `<template shadowrootmode="open">`) {
		t.Error("Default rendering mode should not use shadow DOM")
	}

	// Should contain demo HTML directly in cem-serve-demo
	if !strings.Contains(rendered, `<cem-serve-demo id="demo"><my-element></my-element></cem-serve-demo>`) {
		t.Error("Default rendering mode should render demo in light DOM")
	}
}

// TestRenderingMode_Light verifies explicit light mode
func TestRenderingMode_Light(t *testing.T) {
	demoHTML := `<my-element></my-element>`

	rendered, err := renderDemo(testTemplates(), nil, ChromeData{
		TagName:       "my-element",
		DemoTitle:     "Light Rendering",
		DemoHTML:      template.HTML(demoHTML),
		ImportMap:     "{}",
		RenderingMode: "light",
	})
	if err != nil {
		t.Fatalf("Failed to render chrome: %v", err)
	}

	// Light mode: no <template shadowrootmode="open">
	if strings.Contains(rendered, `<template shadowrootmode="open">`) {
		t.Error("Light rendering mode should not use shadow DOM")
	}

	// Should contain demo HTML directly
	if !strings.Contains(rendered, `<cem-serve-demo id="demo"><my-element></my-element></cem-serve-demo>`) {
		t.Error("Light rendering mode should render demo in light DOM")
	}
}

// TestRenderingMode_Shadow verifies shadow mode wraps demo in declarative shadow DOM
func TestRenderingMode_Shadow(t *testing.T) {
	demoHTML := `<my-element></my-element>`

	rendered, err := renderDemo(testTemplates(), nil, ChromeData{
		TagName:       "my-element",
		DemoTitle:     "Shadow Rendering",
		DemoHTML:      template.HTML(demoHTML),
		ImportMap:     "{}",
		RenderingMode: "shadow",
	})
	if err != nil {
		t.Fatalf("Failed to render chrome: %v", err)
	}

	// Shadow mode: should have <template shadowrootmode="open">
	if !strings.Contains(rendered, `<template shadowrootmode="open">`) {
		t.Error("Shadow rendering mode should use declarative shadow DOM")
	}

	// Should contain demo HTML inside template
	if !strings.Contains(rendered, `<template shadowrootmode="open"><my-element></my-element></template>`) {
		t.Error("Shadow rendering mode should wrap demo in shadow DOM template")
	}
}

// TestRenderingMode_IframeFallback verifies iframe mode falls back with error message
func TestRenderingMode_IframeFallback(t *testing.T) {
	demoHTML := `<my-element></my-element>`

	rendered, err := renderDemo(testTemplates(), nil, ChromeData{
		TagName:       "my-element",
		DemoTitle:     "Iframe Rendering",
		DemoHTML:      template.HTML(demoHTML),
		ImportMap:     "{}",
		RenderingMode: "iframe",
	})
	if err != nil {
		t.Fatalf("Failed to render chrome: %v", err)
	}

	// Iframe mode not implemented - should show error message
	if !strings.Contains(rendered, "Error: iframe rendering mode is not yet implemented") {
		t.Error("Iframe rendering mode should show error message")
	}

	// Should NOT contain the actual demo HTML
	if strings.Contains(rendered, `<my-element></my-element>`) {
		t.Error("Iframe rendering mode should not render the demo (not implemented)")
	}
}

// TestRenderingMode_InvalidValue verifies unknown rendering modes gracefully degrade
func TestRenderingMode_InvalidValue(t *testing.T) {
	demoHTML := `<my-element></my-element>`

	// Invalid value should degrade to light mode (template should handle gracefully)
	rendered, err := renderDemo(testTemplates(), nil, ChromeData{
		TagName:       "my-element",
		DemoTitle:     "Invalid Rendering",
		DemoHTML:      template.HTML(demoHTML),
		ImportMap:     "{}",
		RenderingMode: "invalid-mode",
	})
	if err != nil {
		t.Fatalf("Failed to render chrome: %v", err)
	}

	// Should fallback to light mode (else branch)
	if strings.Contains(rendered, `<template shadowrootmode="open">`) {
		t.Error("Invalid rendering mode should fallback to light DOM")
	}

	// Should contain demo HTML in light DOM
	if !strings.Contains(rendered, `<cem-serve-demo id="demo"><my-element></my-element></cem-serve-demo>`) {
		t.Error("Invalid rendering mode should render demo in light DOM")
	}
}
