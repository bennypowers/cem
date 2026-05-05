/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
*/

package routes

import (
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
			got := extractLocalRoute(tt.demoURL)
			if got != tt.want {
				t.Errorf("extractLocalRoute(%q) = %q, want %q", tt.demoURL, got, tt.want)
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
