/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
*/

package routes

import (
	"fmt"
	"strings"
	"testing"
)

func TestLevenshteinDistance(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want int
	}{
		{
			name: "identical strings",
			s1:   "kitten",
			s2:   "kitten",
			want: 0,
		},
		{
			name: "single substitution",
			s1:   "kitten",
			s2:   "sitten",
			want: 1,
		},
		{
			name: "single insertion",
			s1:   "cat",
			s2:   "cats",
			want: 1,
		},
		{
			name: "single deletion",
			s1:   "cats",
			s2:   "cat",
			want: 1,
		},
		{
			name: "completely different",
			s1:   "abc",
			s2:   "xyz",
			want: 3,
		},
		{
			name: "first string empty",
			s1:   "",
			s2:   "hello",
			want: 5,
		},
		{
			name: "second string empty",
			s1:   "hello",
			s2:   "",
			want: 5,
		},
		{
			name: "both empty",
			s1:   "",
			s2:   "",
			want: 0,
		},
		{
			name: "classic kitten/sitting",
			s1:   "kitten",
			s2:   "sitting",
			want: 3,
		},
		{
			name: "route-like paths",
			s1:   "/elements/button/demo/",
			s2:   "/elements/buton/demo/",
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := levenshteinDistance(tt.s1, tt.s2)
			if got != tt.want {
				t.Errorf("levenshteinDistance(%q, %q) = %d, want %d", tt.s1, tt.s2, got, tt.want)
			}
		})
	}
}

func TestPrettifyRoute(t *testing.T) {
	tests := []struct {
		name  string
		route string
		want  string
	}{
		{
			name:  "demo route with trailing slash",
			route: "/elements/accordion/demo/",
			want:  "Demo",
		},
		{
			name:  "named demo with trailing slash",
			route: "/elements/accordion/demo/accents/",
			want:  "Accents",
		},
		{
			name:  "demo file with extension",
			route: "/demo/basic.html",
			want:  "Basic",
		},
		{
			name:  "hyphenated name",
			route: "/demo/light-dom/",
			want:  "Light dom",
		},
		{
			name:  "underscored name",
			route: "/demo/my_component/",
			want:  "My component",
		},
		{
			name:  "root route",
			route: "/",
			want:  "Demo",
		},
		{
			name:  "just demo",
			route: "/demo",
			want:  "Demo",
		},
		{
			name:  "nested route",
			route: "/elements/avatar/demo/sizes/",
			want:  "Sizes",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := prettifyRoute(tt.route)
			if got != tt.want {
				t.Errorf("prettifyRoute(%q) = %q, want %q", tt.route, got, tt.want)
			}
		})
	}
}

func TestExtractLocalRoute(t *testing.T) {
	tests := []struct {
		name    string
		demoURL string
		want    string
	}{
		{
			name:    "full URL with scheme",
			demoURL: "https://example.com/elements/button/demo/",
			want:    "/elements/button/demo/",
		},
		{
			name:    "relative URL with dot-slash",
			demoURL: "./demo/basic.html",
			want:    "/demo/basic.html",
		},
		{
			name:    "relative URL without prefix",
			demoURL: "demo/sizes.html",
			want:    "/demo/sizes.html",
		},
		{
			name:    "already absolute path",
			demoURL: "/elements/accordion/demo/",
			want:    "/elements/accordion/demo/",
		},
		{
			name:    "directory style adds trailing slash",
			demoURL: "/elements/accordion/demo",
			want:    "/elements/accordion/demo/",
		},
		{
			name:    "file with extension keeps no trailing slash",
			demoURL: "/demo/basic.html",
			want:    "/demo/basic.html",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractLocalRoute(tt.demoURL, "")
			if got != tt.want {
				t.Errorf("extractLocalRoute(%q) = %q, want %q", tt.demoURL, got, tt.want)
			}
		})
	}
}

func TestExtractLocalRoute_PrefixStripping(t *testing.T) {
	tests := []struct {
		name   string
		url    string
		prefix string
		want   string
	}{
		{
			name:   "strips deployment prefix from full URL",
			url:    "https://example.com/cem/examples/kitchen-sink/demo-button/variants/",
			prefix: "/cem/examples/kitchen-sink",
			want:   "/demo-button/variants/",
		},
		{
			name:   "no prefix leaves path unchanged",
			url:    "https://example.com/demo-button/variants/",
			prefix: "",
			want:   "/demo-button/variants/",
		},
		{
			name:   "prefix that does not match leaves path unchanged",
			url:    "/other/path/demo/",
			prefix: "/cem/examples",
			want:   "/other/path/demo/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractLocalRoute(tt.url, tt.prefix)
			if got != tt.want {
				t.Errorf("extractLocalRoute(%q, %q) = %q, want %q", tt.url, tt.prefix, got, tt.want)
			}
		})
	}
}

func TestSlugify(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "simple lowercase",
			input: "demo",
			want:  "demo",
		},
		{
			name:  "spaces to hyphens",
			input: "Light dom",
			want:  "light-dom",
		},
		{
			name:  "special characters removed",
			input: "Hello, World!",
			want:  "hello-world",
		},
		{
			name:  "hyphens preserved",
			input: "my-component",
			want:  "my-component",
		},
		{
			name:  "uppercase lowered",
			input: "Accents",
			want:  "accents",
		},
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
		{
			name:  "numbers preserved",
			input: "version 2",
			want:  "version-2",
		},
		{
			name:  "consecutive spaces collapse",
			input: "hello  world",
			want:  "hello-world",
		},
		{
			name:  "leading/trailing special chars trimmed",
			input: " hello ",
			want:  "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slugify(tt.input)
			if got != tt.want {
				t.Errorf("slugify(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestDemoTitleFromRoute(t *testing.T) {
	tests := []struct {
		name     string
		entry    DemoRouteEntry
		expected string
	}{
		{
			name:     "named route",
			entry:    DemoRouteEntry{LocalRoute: "/elements/accordion/demo/accents/", FilePath: "demo/accents.html"},
			expected: "Accents",
		},
		{
			name:     "generic demo route falls back to filename",
			entry:    DemoRouteEntry{LocalRoute: "/demo/", FilePath: "demo/basic.html"},
			expected: "basic.html",
		},
		{
			name:     "nested demo route",
			entry:    DemoRouteEntry{LocalRoute: "/elements/tabs/demo/vertical/", FilePath: "demo/vertical.html"},
			expected: "Vertical",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := demoTitleFromRoute(&tt.entry)
			if result != tt.expected {
				t.Errorf("demoTitleFromRoute() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestFormatPackageRoutingErrors(t *testing.T) {
	errs := []packageRoutingError{
		{PackageName: "@scope/element", PackagePath: "/path/to/element", Error: fmt.Errorf("missing manifest")},
	}
	result := formatPackageRoutingErrors(errs)
	if result == nil {
		t.Fatal("expected non-nil error")
	}
	msg := result.Error()
	if !strings.Contains(msg, "@scope/element") {
		t.Errorf("error should contain package name, got: %s", msg)
	}
	if !strings.Contains(msg, "missing manifest") {
		t.Errorf("error should contain original error, got: %s", msg)
	}
}

func TestFormatRouteConflictsError(t *testing.T) {
	conflicts := map[string][]routeConflict{
		"/demo/basic": {
			{PackageName: "pkg-a", PackagePath: "/a", FilePath: "demo/basic.html"},
			{PackageName: "pkg-b", PackagePath: "/b", FilePath: "demo/basic.html"},
		},
	}
	result := formatRouteConflictsError(conflicts)
	if result == nil {
		t.Fatal("expected non-nil error")
	}
	msg := result.Error()
	if !strings.Contains(msg, "/demo/basic") {
		t.Errorf("error should contain route, got: %s", msg)
	}
	if !strings.Contains(msg, "pkg-a") || !strings.Contains(msg, "pkg-b") {
		t.Errorf("error should contain both package names, got: %s", msg)
	}
}
