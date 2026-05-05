/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
*/

package serve

import (
	"testing"
)

func TestRewriteAttrPaths(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		attrPrefix string
		basePath   string
		want       string
	}{
		{
			name:       "rewrites href with absolute path",
			input:      `<link href="/style.css">`,
			attrPrefix: `href="`,
			basePath:   "/base",
			want:       `<link href="/base/style.css">`,
		},
		{
			name:       "does not rewrite relative path",
			input:      `<link href="style.css">`,
			attrPrefix: `href="`,
			basePath:   "/base",
			want:       `<link href="style.css">`,
		},
		{
			name:       "does not rewrite protocol-relative URL",
			input:      `<script src="//cdn.example.com/lib.js">`,
			attrPrefix: `src="`,
			basePath:   "/base",
			want:       `<script src="//cdn.example.com/lib.js">`,
		},
		{
			name:       "does not double-prefix",
			input:      `<link href="/base/style.css">`,
			attrPrefix: `href="`,
			basePath:   "/base",
			want:       `<link href="/base/style.css">`,
		},
		{
			name:       "rewrites multiple occurrences",
			input:      `<link href="/a.css"><link href="/b.css">`,
			attrPrefix: `href="`,
			basePath:   "/pfx",
			want:       `<link href="/pfx/a.css"><link href="/pfx/b.css">`,
		},
		{
			name:       "no matching prefix leaves string unchanged",
			input:      `<div class="foo">`,
			attrPrefix: `href="`,
			basePath:   "/base",
			want:       `<div class="foo">`,
		},
		{
			name:       "import with single quotes",
			input:      `import('/__cem/foo.js')`,
			attrPrefix: `import('`,
			basePath:   "/docs",
			want:       `import('/docs/__cem/foo.js')`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rewriteAttrPaths(tt.input, tt.attrPrefix, tt.basePath)
			if got != tt.want {
				t.Errorf("rewriteAttrPaths() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestLooksLikeScopeKey(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "valid scope key pattern",
			input: `/node_modules/lit/": {`,
			want:  true,
		},
		{
			name:  "scope key with whitespace before brace",
			input: `/node_modules/lit/"  :  {`,
			want:  true,
		},
		{
			name:  "not a scope key - no brace",
			input: `/node_modules/lit/": "value"`,
			want:  false,
		},
		{
			name:  "not a scope key - no colon",
			input: `/node_modules/lit/" {`,
			want:  false,
		},
		{
			name:  "not a scope key - no closing quote",
			input: `/node_modules/lit/`,
			want:  false,
		},
		{
			name:  "empty string",
			input: "",
			want:  false,
		},
		{
			name:  "scope key with newlines",
			input: "/path/\"\n:\n{",
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := looksLikeScopeKey(tt.input)
			if got != tt.want {
				t.Errorf("looksLikeScopeKey(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
