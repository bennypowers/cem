/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
*/

package serve

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSiteRoot(t *testing.T) {
	tests := []struct {
		name     string
		config   BuildConfig
		expected string
	}{
		{
			name:     "no base path",
			config:   BuildConfig{OutputDir: "dist"},
			expected: "dist",
		},
		{
			name:     "with base path",
			config:   BuildConfig{OutputDir: "dist", BasePath: "/docs"},
			expected: filepath.Join("dist", "/docs"),
		},
		{
			name:     "empty base path",
			config:   BuildConfig{OutputDir: "/tmp/out", BasePath: ""},
			expected: "/tmp/out",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.config.siteRoot()
			if got != tt.expected {
				t.Errorf("siteRoot() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestRewriteBasePath(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		basePath string
		want     string
	}{
		{
			name:     "rewrites href",
			input:    `<link href="/__cem/style.css">`,
			basePath: "/docs",
			want:     `<link href="/docs/__cem/style.css">`,
		},
		{
			name:     "rewrites src",
			input:    `<script src="/__cem/bundle.js"></script>`,
			basePath: "/docs",
			want:     `<script src="/docs/__cem/bundle.js"></script>`,
		},
		{
			name:     "rewrites import map JSON",
			input:    `"lit": "/node_modules/lit/index.js"`,
			basePath: "/docs",
			want:     `"lit": "/docs/node_modules/lit/index.js"`,
		},
		{
			name:     "rewrites dynamic imports",
			input:    `import('/__cem/health-badges.js')`,
			basePath: "/docs",
			want:     `import('/docs/__cem/health-badges.js')`,
		},
		{
			name:     "does not double-prefix",
			input:    `<link href="/docs/__cem/style.css">`,
			basePath: "/docs",
			want:     `<link href="/docs/__cem/style.css">`,
		},
		{
			name:     "preserves external URLs",
			input:    `<a href="https://example.com">`,
			basePath: "/docs",
			want:     `<a href="https://example.com">`,
		},
		{
			name:     "preserves protocol-relative URLs",
			input:    `<script src="//cdn.example.com/lib.js">`,
			basePath: "/docs",
			want:     `<script src="//cdn.example.com/lib.js">`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := string(rewriteBasePath([]byte(tt.input), tt.basePath))
			if got != tt.want {
				t.Errorf("rewriteBasePath() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestParseNodeModulePath(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantPkg    string
		wantSubpath string
	}{
		{
			name:        "bare package",
			input:       `lit/index.js"`,
			wantPkg:     "lit",
			wantSubpath: "index.js",
		},
		{
			name:        "scoped package",
			input:       `@lit/reactive-element/reactive-element.js"`,
			wantPkg:     "@lit/reactive-element",
			wantSubpath: "reactive-element.js",
		},
		{
			name:        "bare package root",
			input:       `lit"`,
			wantPkg:     "lit",
			wantSubpath: "",
		},
		{
			name:        "scoped package root",
			input:       `@lit/reactive-element"`,
			wantPkg:     "@lit/reactive-element",
			wantSubpath: "",
		},
		{
			name:        "nested subpath",
			input:       `lit/directives/class-map.js"`,
			wantPkg:     "lit",
			wantSubpath: "directives/class-map.js",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pkg, subpath := parseNodeModulePath(tt.input)
			if pkg != tt.wantPkg {
				t.Errorf("pkg = %q, want %q", pkg, tt.wantPkg)
			}
			if subpath != tt.wantSubpath {
				t.Errorf("subpath = %q, want %q", subpath, tt.wantSubpath)
			}
		})
	}
}

func TestCdnURL(t *testing.T) {
	tests := []struct {
		mode    string
		pkg     string
		subpath string
		want    string
	}{
		{"esm", "lit", "index.js", "https://esm.sh/lit/index.js"},
		{"esm", "lit", "", "https://esm.sh/lit"},
		{"esm", "@lit/reactive-element", "reactive-element.js", "https://esm.sh/@lit/reactive-element/reactive-element.js"},
		{"jspm", "lit", "index.js", "https://ga.jspm.io/npm:lit/index.js"},
		{"jspm", "lit", "", "https://ga.jspm.io/npm:lit"},
		{"unpkg", "lit", "index.js", "https://unpkg.com/lit/index.js"},
		{"unpkg", "@lit/reactive-element", "", "https://unpkg.com/@lit/reactive-element"},
	}

	for _, tt := range tests {
		t.Run(tt.mode+"/"+tt.pkg, func(t *testing.T) {
			got := cdnURL(tt.mode, tt.pkg, tt.subpath)
			if got != tt.want {
				t.Errorf("cdnURL() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestRewriteNodeModulesToCDN(t *testing.T) {
	input := `{
    "imports": {
      "lit": "/node_modules/lit/index.js",
      "lit/decorators.js": "/node_modules/lit/decorators.js",
      "@lit/reactive-element": "/node_modules/@lit/reactive-element/reactive-element.js"
    }
  }`

	tests := []struct {
		mode string
		want []string
	}{
		{
			mode: "esm",
			want: []string{
				`"https://esm.sh/lit/index.js"`,
				`"https://esm.sh/lit/decorators.js"`,
				`"https://esm.sh/@lit/reactive-element/reactive-element.js"`,
			},
		},
		{
			mode: "unpkg",
			want: []string{
				`"https://unpkg.com/lit/index.js"`,
				`"https://unpkg.com/@lit/reactive-element/reactive-element.js"`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.mode, func(t *testing.T) {
			got := rewriteNodeModulesToCDN(input, tt.mode)
			for _, w := range tt.want {
				if !strings.Contains(got, w) {
					t.Errorf("expected %q in output, got:\n%s", w, got)
				}
			}
			if strings.Contains(got, "/node_modules/") {
				t.Errorf("output still contains /node_modules/:\n%s", got)
			}
		})
	}
}

func TestWriteWebSocketStub(t *testing.T) {
	dir := t.TempDir()
	s := &Server{}
	config := BuildConfig{OutputDir: dir}

	if err := s.writeWebSocketStub(config); err != nil {
		t.Fatalf("writeWebSocketStub() error: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(dir, "__cem", "websocket-client.js"))
	if err != nil {
		t.Fatalf("ReadFile() error: %v", err)
	}

	content := string(data)
	if !strings.Contains(content, "export class CEMReloadClient") {
		t.Error("stub should export CEMReloadClient")
	}
	if !strings.Contains(content, "init() {}") {
		t.Error("stub should have no-op init()")
	}
}

func TestBuildLightdomCSS(t *testing.T) {
	dir := t.TempDir()
	s := &Server{}
	config := BuildConfig{OutputDir: dir}

	if err := s.buildLightdomCSS(config); err != nil {
		t.Fatalf("buildLightdomCSS() error: %v", err)
	}

	// The output may or may not exist depending on whether there are
	// lightdom CSS files in the embedded FS. Just verify no error.
	outPath := filepath.Join(dir, "__cem", "lightdom.css")
	if _, err := os.Stat(outPath); err == nil {
		data, _ := os.ReadFile(outPath)
		if len(data) == 0 {
			t.Error("lightdom.css should not be empty if it exists")
		}
	}
}

func TestBuildConfig_SiteRoot_WithBasePath(t *testing.T) {
	config := BuildConfig{OutputDir: "dist", BasePath: "/docs/api"}
	got := config.siteRoot()
	want := filepath.Join("dist", "/docs/api")
	if got != want {
		t.Errorf("siteRoot() = %q, want %q", got, want)
	}
}
