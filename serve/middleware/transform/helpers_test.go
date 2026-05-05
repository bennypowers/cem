/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
*/

package transform

import (
	"testing"
)

func TestIsRelativeImport(t *testing.T) {
	tests := []struct {
		name string
		path string
		want bool
	}{
		{
			name: "dot-slash relative",
			path: "./foo.js",
			want: true,
		},
		{
			name: "dot-dot-slash relative",
			path: "../bar.js",
			want: true,
		},
		{
			name: "scoped package",
			path: "@scope/pkg",
			want: false,
		},
		{
			name: "bare specifier without extension",
			path: "lodash",
			want: false,
		},
		{
			name: "bare specifier lit",
			path: "lit",
			want: false,
		},
		{
			name: "absolute path",
			path: "/usr/local/lib",
			want: false,
		},
		{
			name: "normalized relative with extension",
			path: "component.css",
			want: true,
		},
		{
			name: "normalized relative with subpath",
			path: "shared/utils.js",
			want: true,
		},
		{
			name: "deeply nested relative",
			path: "../../components/button.ts",
			want: true,
		},
		{
			name: "scoped package with subpath",
			path: "@rhds/elements/rh-button/rh-button.js",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isRelativeImport(tt.path)
			if got != tt.want {
				t.Errorf("isRelativeImport(%q) = %v, want %v", tt.path, got, tt.want)
			}
		})
	}
}

func TestResolveImport(t *testing.T) {
	tests := []struct {
		name       string
		sourcePath string
		importPath string
		want       string
	}{
		{
			name:       "same directory",
			sourcePath: "src/components/button.ts",
			importPath: "./utils.ts",
			want:       "src/components/utils.ts",
		},
		{
			name:       "parent directory",
			sourcePath: "src/components/button.ts",
			importPath: "../shared/styles.ts",
			want:       "src/shared/styles.ts",
		},
		{
			name:       "sibling file",
			sourcePath: "src/index.ts",
			importPath: "./app.ts",
			want:       "src/app.ts",
		},
		{
			name:       "deeply nested up",
			sourcePath: "src/a/b/c/deep.ts",
			importPath: "../../../root.ts",
			want:       "src/root.ts",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := resolveImport(tt.sourcePath, tt.importPath)
			if got != tt.want {
				t.Errorf("resolveImport(%q, %q) = %q, want %q", tt.sourcePath, tt.importPath, got, tt.want)
			}
		})
	}
}

func TestStringToTemplateLiteral(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "simple string no escapes needed",
			input: "hello world",
			want:  "hello world",
		},
		{
			name:  "string with backtick",
			input: "use `code` here",
			want:  "use \\`code\\` here",
		},
		{
			name:  "string with dollar-brace interpolation",
			input: "value is ${count}",
			want:  "value is \\${count}",
		},
		{
			name:  "dollar without brace is not escaped",
			input: "price is $5",
			want:  "price is $5",
		},
		{
			name:  "backslash is escaped",
			input: `path\to\file`,
			want:  `path\\to\\file`,
		},
		{
			name:  "closing script tag slash escaped",
			input: "</script>",
			want:  "<\\/script>",
		},
		{
			name:  "regular slash not escaped",
			input: "a/b/c",
			want:  "a/b/c",
		},
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
		{
			name:  "multiple escapes combined",
			input: "color: `red`; size: ${lg}; path: </end>",
			want:  "color: \\`red\\`; size: \\${lg}; path: <\\/end>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stringToTemplateLiteral(tt.input)
			if got != tt.want {
				t.Errorf("stringToTemplateLiteral(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
