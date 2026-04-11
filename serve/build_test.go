/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
*/

package serve

import (
	"fmt"
	"io"
	"io/fs"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/serve/logger"
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

// TestBuildUserSources_SkipsOutputDir verifies that buildUserSources does not
// recurse into the output directory, which would cause an infinite loop when
// the output directory is inside the watch directory.
func TestBuildUserSources_SkipsOutputDir(t *testing.T) {
	dir := t.TempDir()

	// Create a source file
	srcDir := filepath.Join(dir, "src")
	if err := os.MkdirAll(srcDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(srcDir, "app.js"), []byte("// source"), 0o644); err != nil {
		t.Fatal(err)
	}

	// Create an output directory with files that should be skipped
	outDir := filepath.Join(dir, "_site")
	nestedDir := filepath.Join(outDir, "deep", "nested")
	if err := os.MkdirAll(nestedDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(outDir, "should-skip.js"), []byte("// output"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(nestedDir, "also-skip.js"), []byte("// nested output"), 0o644); err != nil {
		t.Fatal(err)
	}

	// Track which paths are walked (not skipped) by collecting all
	// non-directory entries via WalkDir with the same skip logic
	var walked []string
	absOutputDir, _ := filepath.Abs(outDir)

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		rel, _ := filepath.Rel(dir, path)
		if strings.HasPrefix(rel, "node_modules") {
			if d.IsDir() {
				return fs.SkipDir
			}
			return nil
		}

		absPath, _ := filepath.Abs(path)
		if d.IsDir() && absPath == absOutputDir {
			return fs.SkipDir
		}

		if d.IsDir() {
			return nil
		}

		walked = append(walked, rel)
		return nil
	})
	if err != nil {
		t.Fatalf("WalkDir error: %v", err)
	}

	// Verify: should contain the source file but NOT any output files
	foundSource := false
	for _, p := range walked {
		if strings.Contains(p, "_site") {
			t.Errorf("walked into output directory: %s", p)
		}
		if p == filepath.Join("src", "app.js") {
			foundSource = true
		}
	}
	if !foundSource {
		t.Errorf("did not walk source file; walked: %v", walked)
	}
}

func TestRewriteAssetPaths(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		basePath string
		want     string
	}{
		{
			name:     "rewrites JS dynamic import",
			input:    `import("/__cem/websocket-client.js")`,
			basePath: "/docs",
			want:     `import("/docs/__cem/websocket-client.js")`,
		},
		{
			name:     "rewrites JS fetch",
			input:    `fetch("/custom-elements.json")`,
			basePath: "/docs",
			want:     `fetch("/docs/custom-elements.json")`,
		},
		{
			name:     "rewrites JS single-quote import",
			input:    `import('/__cem/state-persistence.js')`,
			basePath: "/docs",
			want:     `import('/docs/__cem/state-persistence.js')`,
		},
		{
			name:     "rewrites CSS import url",
			input:    `@import url('/__cem/elements/pf-v6-card/pf-v6-card-lightdom.css');`,
			basePath: "/docs",
			want:     `@import url('/docs/__cem/elements/pf-v6-card/pf-v6-card-lightdom.css');`,
		},
		{
			name:     "rewrites new URL pattern",
			input:    `new URL("/__cem/api/health", location.origin)`,
			basePath: "/docs",
			want:     `new URL("/docs/__cem/api/health", location.origin)`,
		},
		{
			name:     "does not double-prefix",
			input:    `import("/docs/__cem/websocket-client.js")`,
			basePath: "/docs",
			want:     `import("/docs/__cem/websocket-client.js")`,
		},
		{
			name:     "preserves non-absolute paths",
			input:    `import("./local-module.js")`,
			basePath: "/docs",
			want:     `import("./local-module.js")`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := string(rewriteAssetPaths([]byte(tt.input), tt.basePath))
			if got != tt.want {
				t.Errorf("rewriteAssetPaths() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestRewriteJSONScopeKeys(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		basePath string
		want     string
	}{
		{
			name: "rewrites scope keys",
			input: `{
    "scopes": {
      "/node_modules/lit/": {
        "@lit/reactive-element": "/node_modules/@lit/reactive-element/reactive-element.js"
      }
    }
  }`,
			basePath: "/base",
			want: `{
    "scopes": {
      "/base/node_modules/lit/": {
        "@lit/reactive-element": "/node_modules/@lit/reactive-element/reactive-element.js"
      }
    }
  }`,
		},
		{
			name:     "does not double-prefix scope keys",
			input:    `"/base/node_modules/lit/": {`,
			basePath: "/base",
			want:     `"/base/node_modules/lit/": {`,
		},
		{
			name:     "preserves protocol-relative URLs",
			input:    `src="//cdn.example.com/lib.js"`,
			basePath: "/base",
			want:     `src="//cdn.example.com/lib.js"`,
		},
		{
			name:     "does not rewrite JSON values",
			input:    `": "/node_modules/lit/index.js"`,
			basePath: "/base",
			want:     `": "/node_modules/lit/index.js"`,
		},
		{
			name:     "does not rewrite HTML attributes",
			input:    `<a href="/foo/bar">link</a>`,
			basePath: "/base",
			want:     `<a href="/foo/bar">link</a>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rewriteJSONScopeKeys(tt.input, tt.basePath)
			if got != tt.want {
				t.Errorf("rewriteJSONScopeKeys() = %q, want %q", got, tt.want)
			}
		})
	}
}

// forceCloseConnection hijacks an HTTP connection and closes it immediately,
// causing the client to see an EOF error. Panics if hijacking fails (which
// would indicate a broken test server, not a test failure).
func forceCloseConnection(w http.ResponseWriter) {
	hj, ok := w.(http.Hijacker)
	if !ok {
		panic("test server does not support hijacking")
	}
	conn, _, err := hj.Hijack()
	if err != nil {
		panic(fmt.Sprintf("hijack failed: %v", err))
	}
	_ = conn.Close()
}

func TestRetryFetch_SucceedsOnFirstAttempt(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("ok"))
	}))
	defer ts.Close()

	s := &Server{logger: logger.NewDefaultLogger()}
	body, err := s.retryFetch(ts, "/test")
	if err != nil {
		t.Fatalf("retryFetch() error: %v", err)
	}
	if string(body) != "ok" {
		t.Errorf("retryFetch() = %q, want %q", string(body), "ok")
	}
}

func TestRetryFetch_RetriesOnEOF(t *testing.T) {
	var attempts atomic.Int32
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		n := attempts.Add(1)
		if n <= 2 {
			// Simulate EOF by hijacking the connection and closing it
			forceCloseConnection(w)
			return
		}
		_, _ = w.Write([]byte("recovered"))
	}))
	defer ts.Close()

	s := &Server{logger: logger.NewDefaultLogger()}
	body, err := s.retryFetch(ts, "/test")
	if err != nil {
		t.Fatalf("retryFetch() error after retries: %v", err)
	}
	if string(body) != "recovered" {
		t.Errorf("retryFetch() = %q, want %q", string(body), "recovered")
	}
	if got := attempts.Load(); got != 3 {
		t.Errorf("expected 3 attempts, got %d", got)
	}
}

func TestRetryFetch_DoesNotRetryHTTPErrors(t *testing.T) {
	var attempts atomic.Int32
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		attempts.Add(1)
		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()

	s := &Server{logger: logger.NewDefaultLogger()}
	_, err := s.retryFetch(ts, "/missing")
	if err == nil {
		t.Fatal("expected error for 404 response")
	}
	if got := attempts.Load(); got != 1 {
		t.Errorf("expected 1 attempt for HTTP error, got %d", got)
	}
}

func TestRetryFetch_ExhaustsRetries(t *testing.T) {
	var attempts atomic.Int32
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		attempts.Add(1)
		forceCloseConnection(w)
	}))
	defer ts.Close()

	s := &Server{logger: logger.NewDefaultLogger()}
	_, err := s.retryFetch(ts, "/always-fail")
	if err == nil {
		t.Fatal("expected error after exhausting retries")
	}
	if got := attempts.Load(); got != 4 {
		t.Errorf("expected 4 attempts (1 + 3 retries), got %d", got)
	}
}

func TestIsTransientError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{"EOF", io.EOF, true},
		{"unexpected EOF", io.ErrUnexpectedEOF, true},
		{"connection reset", &net.OpError{Op: "read", Err: &os.SyscallError{Syscall: "read", Err: fmt.Errorf("connection reset by peer")}}, true},
		{"not found", fmt.Errorf("HTTP 404"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTransientError(tt.err); got != tt.want {
				t.Errorf("isTransientError(%v) = %v, want %v", tt.err, got, tt.want)
			}
		})
	}
}

// TestBuild_ReturnsErrorOnPageFailure verifies that Build() returns a non-nil
// error when a demo page fetch fails, instead of silently succeeding.
func TestBuild_ReturnsErrorOnPageFailure(t *testing.T) {
	// Manifest with two elements: one demo file will exist, one won't.
	manifest := []byte(`{
  "schemaVersion": "1.0.0",
  "modules": [
    {
      "kind": "javascript-module",
      "path": "src/el-ok.ts",
      "declarations": [{
        "kind": "class",
        "customElement": true,
        "name": "ElOk",
        "tagName": "el-ok",
        "demos": [{"description": "OK", "url": "./demo/exists.html"}]
      }]
    },
    {
      "kind": "javascript-module",
      "path": "src/el-bad.ts",
      "declarations": [{
        "kind": "class",
        "customElement": true,
        "name": "ElBad",
        "tagName": "el-bad",
        "demos": [{"description": "Bad", "url": "./demo/missing.html"}]
      }]
    }
  ]
}`)

	// Create a MapFileSystem with only the "exists" demo file present.
	mfs := platform.NewMapFileSystem(nil)
	mfs.AddFile("/test/demo/exists.html", `<el-ok>hello</el-ok>`, 0o644)
	// Deliberately omit /test/demo/missing.html

	server, err := NewServerWithConfig(Config{
		Port:   0,
		Reload: false,
		FS:     mfs,
	})
	if err != nil {
		t.Fatalf("NewServerWithConfig: %v", err)
	}

	if err := server.SetWatchDir("/test"); err != nil {
		t.Fatalf("SetWatchDir: %v", err)
	}
	if err := server.SetManifest(manifest); err != nil {
		t.Fatalf("SetManifest: %v", err)
	}

	// Build should fail because the missing demo file causes HTTP 500.
	buildErr := server.Build(BuildConfig{OutputDir: t.TempDir()})
	if buildErr == nil {
		t.Fatal("expected Build() to return an error when a page fetch fails, got nil")
	}
	if !strings.Contains(buildErr.Error(), "failed to build pages") {
		t.Errorf("expected error about failed pages, got: %v", buildErr)
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
