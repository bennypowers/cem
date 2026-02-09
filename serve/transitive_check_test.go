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

package serve

import (
	"net/http"
	"strings"
	"sync"
	"testing"

	importmappkg "bennypowers.dev/cem/serve/middleware/importmap"
)

func TestExtractPackageName(t *testing.T) {
	tests := []struct {
		specifier string
		want      string
	}{
		{"lit", "lit"},
		{"lit-html", "lit-html"},
		{"lit-html/directives/if-defined.js", "lit-html"},
		{"@lit/reactive-element", "@lit/reactive-element"},
		{"@lit/reactive-element/decorators.js", "@lit/reactive-element"},
		{"@scope/pkg/deep/path/file.js", "@scope/pkg"},
		{"single", "single"},
		{"@incomplete", ""},
	}

	for _, tt := range tests {
		t.Run(tt.specifier, func(t *testing.T) {
			got := extractPackageName(tt.specifier)
			if got != tt.want {
				t.Errorf("extractPackageName(%q) = %q, want %q", tt.specifier, got, tt.want)
			}
		})
	}
}

func TestIsBareSpecifier(t *testing.T) {
	tests := []struct {
		name      string
		specifier string
		want      bool
	}{
		{"bare package", "lit", true},
		{"bare subpath", "lit-html/directives/if-defined.js", true},
		{"bare scoped", "@lit/reactive-element", true},
		{"relative dot-slash", "./relative.js", false},
		{"relative parent", "../parent.js", false},
		{"absolute", "/absolute.js", false},
		{"http URL", "http://example.com/mod.js", false},
		{"https URL", "https://cdn.example.com/mod.js", false},
		{"empty string", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isBareSpecifier(tt.specifier)
			if got != tt.want {
				t.Errorf("isBareSpecifier(%q) = %v, want %v", tt.specifier, got, tt.want)
			}
		})
	}
}

func TestIsResolvedInImportMap(t *testing.T) {
	im := &importmappkg.ImportMap{
		Imports: map[string]string{
			"lit":                   "/node_modules/lit/index.js",
			"lit/":                  "/node_modules/lit/",
			"@patternfly/elements/": "/node_modules/@patternfly/elements/",
			"@patternfly/elements":  "/node_modules/@patternfly/elements/index.js",
		},
	}

	tests := []struct {
		name      string
		specifier string
		want      bool
	}{
		{"exact match", "lit", true},
		{"prefix match", "lit/decorators.js", true},
		{"scoped exact", "@patternfly/elements", true},
		{"scoped prefix", "@patternfly/elements/pf-button/pf-button.js", true},
		{"no match", "lit-html", false},
		{"no match subpath", "lit-html/directives/if-defined.js", false},
		{"partial name no match", "litz", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isResolvedInImportMap(tt.specifier, im)
			if got != tt.want {
				t.Errorf("isResolvedInImportMap(%q) = %v, want %v", tt.specifier, got, tt.want)
			}
		})
	}

	t.Run("nil import map", func(t *testing.T) {
		if isResolvedInImportMap("lit", nil) {
			t.Error("expected false for nil import map")
		}
	})

	t.Run("empty import map", func(t *testing.T) {
		empty := &importmappkg.ImportMap{Imports: map[string]string{}}
		if isResolvedInImportMap("lit", empty) {
			t.Error("expected false for empty import map")
		}
	})
}

func TestFormatTransitiveWarning(t *testing.T) {
	t.Run("single owner", func(t *testing.T) {
		msg := formatTransitiveWarning("lit-html/directives/if-defined.js", "lit-html", []string{"lit"})
		if !strings.Contains(msg, `Cannot resolve "lit-html/directives/if-defined.js"`) {
			t.Errorf("expected specifier in message, got: %s", msg)
		}
		if !strings.Contains(msg, `transitive dependency of "lit"`) {
			t.Errorf("expected owner in message, got: %s", msg)
		}
		if !strings.Contains(msg, `import from "lit" instead`) {
			t.Errorf("expected suggestion to import from owner, got: %s", msg)
		}
	})

	t.Run("multiple owners", func(t *testing.T) {
		msg := formatTransitiveWarning("shared-dep/util.js", "shared-dep", []string{"pkg-a", "pkg-b"})
		if !strings.Contains(msg, `transitive dependency of "pkg-a", "pkg-b"`) {
			t.Errorf("expected both owners in message, got: %s", msg)
		}
		if !strings.Contains(msg, "one of its dependents") {
			t.Errorf("expected generic suggestion for multiple owners, got: %s", msg)
		}
	})

	t.Run("no owners", func(t *testing.T) {
		msg := formatTransitiveWarning("unknown-pkg/foo.js", "unknown-pkg", nil)
		if !strings.Contains(msg, "not in the import map") {
			t.Errorf("expected unknown message, got: %s", msg)
		}
		if !strings.Contains(msg, `Add "unknown-pkg"`) {
			t.Errorf("expected add suggestion, got: %s", msg)
		}
	})
}

func TestCheckBareSpecifiers(t *testing.T) {
	t.Run("transitive dep triggers warning", func(t *testing.T) {
		var warnings []string
		var broadcasts int
		var mu sync.Mutex

		server := &Server{
			logger: &testLogger{
				warningFn: func(format string, args ...any) {
					mu.Lock()
					defer mu.Unlock()
					warnings = append(warnings, format)
				},
				debugFn: func(string, ...any) {},
			},
			importMap: &importmappkg.ImportMap{
				Imports: map[string]string{
					"lit":  "/node_modules/lit/index.js",
					"lit/": "/node_modules/lit/",
				},
			},
			importMapGraph: importmappkg.NewTestDependencyGraph(map[string][]string{
				"lit-html": {"lit"},
			}),
			wsManager: &testWSManager{
				broadcastFn: func(msg []byte) error {
					mu.Lock()
					defer mu.Unlock()
					broadcasts++
					return nil
				},
			},
		}

		code := []byte(`import { html } from 'lit';
import { ifDefined } from 'lit-html/directives/if-defined.js';
import './relative.js';
`)

		server.checkBareSpecifiers("demo/my-element.ts", code)

		mu.Lock()
		defer mu.Unlock()

		if len(warnings) != 1 {
			t.Fatalf("expected 1 warning, got %d: %v", len(warnings), warnings)
		}
		if broadcasts != 1 {
			t.Fatalf("expected 1 broadcast, got %d", broadcasts)
		}
	})

	t.Run("direct dep does not trigger warning", func(t *testing.T) {
		var warnings []string
		var mu sync.Mutex

		server := &Server{
			logger: &testLogger{
				warningFn: func(format string, args ...any) {
					mu.Lock()
					defer mu.Unlock()
					warnings = append(warnings, format)
				},
				debugFn: func(string, ...any) {},
			},
			importMap: &importmappkg.ImportMap{
				Imports: map[string]string{
					"lit":  "/node_modules/lit/index.js",
					"lit/": "/node_modules/lit/",
				},
			},
		}

		code := []byte(`import { html } from 'lit';
import { css } from 'lit/decorators.js';
`)

		server.checkBareSpecifiers("demo/my-element.ts", code)

		mu.Lock()
		defer mu.Unlock()

		if len(warnings) != 0 {
			t.Fatalf("expected 0 warnings, got %d: %v", len(warnings), warnings)
		}
	})

	t.Run("deduplicates warnings", func(t *testing.T) {
		var warnings []string
		var mu sync.Mutex

		server := &Server{
			logger: &testLogger{
				warningFn: func(format string, args ...any) {
					mu.Lock()
					defer mu.Unlock()
					warnings = append(warnings, format)
				},
				debugFn: func(string, ...any) {},
			},
			importMap: &importmappkg.ImportMap{
				Imports: map[string]string{
					"lit":  "/node_modules/lit/index.js",
					"lit/": "/node_modules/lit/",
				},
			},
			importMapGraph: importmappkg.NewTestDependencyGraph(map[string][]string{
				"lit-html": {"lit"},
			}),
			wsManager: &testWSManager{
				broadcastFn: func(msg []byte) error { return nil },
			},
		}

		code := []byte(`import { ifDefined } from 'lit-html/directives/if-defined.js';`)

		server.checkBareSpecifiers("file1.ts", code)
		server.checkBareSpecifiers("file2.ts", code)

		mu.Lock()
		defer mu.Unlock()

		if len(warnings) != 1 {
			t.Fatalf("expected 1 warning (deduped), got %d", len(warnings))
		}
	})

	t.Run("relative and absolute imports ignored", func(t *testing.T) {
		var warnings []string
		var mu sync.Mutex

		server := &Server{
			logger: &testLogger{
				warningFn: func(format string, args ...any) {
					mu.Lock()
					defer mu.Unlock()
					warnings = append(warnings, format)
				},
				debugFn: func(string, ...any) {},
			},
			importMap: &importmappkg.ImportMap{
				Imports: map[string]string{},
			},
		}

		code := []byte(`import './local.js';
import '../parent.js';
import '/absolute.js';
`)

		server.checkBareSpecifiers("demo/my-element.ts", code)

		mu.Lock()
		defer mu.Unlock()

		if len(warnings) != 0 {
			t.Fatalf("expected 0 warnings for relative/absolute imports, got %d", len(warnings))
		}
	})
}

// testLogger implements the Logger interface for testing
type testLogger struct {
	warningFn func(string, ...any)
	debugFn   func(string, ...any)
}

func (l *testLogger) Info(format string, args ...any)    {}
func (l *testLogger) Error(format string, args ...any)   {}
func (l *testLogger) Warning(format string, args ...any) { l.warningFn(format, args...) }
func (l *testLogger) Debug(format string, args ...any)   { l.debugFn(format, args...) }

// testWSManager implements WebSocketManager for testing
type testWSManager struct {
	broadcastFn func([]byte) error
}

func (m *testWSManager) HandleConnection(w http.ResponseWriter, r *http.Request) {}
func (m *testWSManager) Broadcast(msg []byte) error                              { return m.broadcastFn(msg) }
func (m *testWSManager) BroadcastToPages(msg []byte, pageURLs []string) error    { return nil }
func (m *testWSManager) BroadcastShutdown() error                                { return nil }
func (m *testWSManager) CloseAll() error                                         { return nil }
func (m *testWSManager) ConnectionCount() int                                    { return 0 }
func (m *testWSManager) SetLogger(logger Logger)                                 {}
