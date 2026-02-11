/*
Copyright © 2025 Benny Powers

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
package generate

import (
	"os"
	"path/filepath"
	"testing"

	M "bennypowers.dev/cem/manifest"
	Q "bennypowers.dev/cem/queries"
)

func TestExtractPackageName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"@scope/pkg/sub/path", "@scope/pkg"},
		{"@scope/pkg", "@scope/pkg"},
		{"pkg/sub/path", "pkg"},
		{"pkg", "pkg"},
		{"./relative", ""},
		{"../relative", ""},
		{".", ""},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := extractPackageName(tc.input)
			if got != tc.expected {
				t.Errorf("extractPackageName(%q) = %q, want %q", tc.input, got, tc.expected)
			}
		})
	}
}

func TestExtractSubpath(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"@scope/pkg/sub/path", "./sub/path"},
		{"@scope/pkg", "."},
		{"pkg/sub/path", "./sub/path"},
		{"pkg", "."},
		{"./relative", "."},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := extractSubpath(tc.input)
			if got != tc.expected {
				t.Errorf("extractSubpath(%q) = %q, want %q", tc.input, got, tc.expected)
			}
		})
	}
}

func TestParseJSDocTypedefs(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected map[string]string
	}{
		{
			name:    "single typedef",
			content: "/** @typedef {'light' | 'dark'} ThemeVariant */",
			expected: map[string]string{
				"ThemeVariant": "'light' | 'dark'",
			},
		},
		{
			name: "multiple typedefs",
			content: `/** @typedef {'light' | 'dark'} ThemeVariant */
/** @typedef {'sm' | 'md' | 'lg'} Size */`,
			expected: map[string]string{
				"ThemeVariant": "'light' | 'dark'",
				"Size":         "'sm' | 'md' | 'lg'",
			},
		},
		{
			name:     "no typedefs",
			content:  "// just a comment\nconst x = 1;",
			expected: map[string]string{},
		},
		{
			name:     "malformed typedef",
			content:  "/** @typedef ThemeVariant */",
			expected: map[string]string{},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := parseJSDocTypedefs([]byte(tc.content))
			if len(got) != len(tc.expected) {
				t.Errorf("parseJSDocTypedefs() returned %d entries, want %d", len(got), len(tc.expected))
			}
			for k, v := range tc.expected {
				if got[k] != v {
					t.Errorf("parseJSDocTypedefs()[%q] = %q, want %q", k, got[k], v)
				}
			}
		})
	}
}

func TestScanTypeAliasesFromFile(t *testing.T) {
	// Create a temp .ts file with type aliases
	dir := t.TempDir()
	tsFile := filepath.Join(dir, "types.ts")
	content := `export type Variant = 'primary' | 'secondary';
export type Size = 'sm' | 'md' | 'lg';
`
	if err := os.WriteFile(tsFile, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	qm, err := Q.NewQueryManager(Q.GenerateQueries())
	if err != nil {
		t.Fatal(err)
	}
	defer qm.Close()

	resolver := &ExternalTypeResolver{queryManager: qm}
	aliases, err := resolver.scanTypeAliasesFromFile(tsFile)
	if err != nil {
		t.Fatal(err)
	}

	if aliases["Variant"] != "'primary' | 'secondary'" {
		t.Errorf("Variant = %q, want %q", aliases["Variant"], "'primary' | 'secondary'")
	}
	if aliases["Size"] != "'sm' | 'md' | 'lg'" {
		t.Errorf("Size = %q, want %q", aliases["Size"], "'sm' | 'md' | 'lg'")
	}
}

func TestResolveType_DTS(t *testing.T) {
	// Set up a temp directory mimicking node_modules with .d.ts
	dir := t.TempDir()
	pkgDir := filepath.Join(dir, "node_modules", "@tokens", "core")
	if err := os.MkdirAll(pkgDir, 0755); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(filepath.Join(pkgDir, "package.json"), []byte(`{
		"name": "@tokens/core",
		"types": "index.d.ts"
	}`), 0644); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(filepath.Join(pkgDir, "index.d.ts"), []byte(
		`export type ThemeVariant = 'light' | 'dark' | 'high-contrast';`,
	), 0644); err != nil {
		t.Fatal(err)
	}

	qm, err := Q.NewQueryManager(Q.GenerateQueries())
	if err != nil {
		t.Fatal(err)
	}
	defer qm.Close()

	resolver := &ExternalTypeResolver{
		projectRoot:  dir,
		queryManager: qm,
	}

	def, pkg, found := resolver.ResolveType("@tokens/core", "ThemeVariant")
	if !found {
		t.Fatal("expected to find ThemeVariant")
	}
	if pkg != "@tokens/core" {
		t.Errorf("package = %q, want %q", pkg, "@tokens/core")
	}
	if def != "'light' | 'dark' | 'high-contrast'" {
		t.Errorf("definition = %q, want %q", def, "'light' | 'dark' | 'high-contrast'")
	}
}

func TestResolveType_JSDocFallback(t *testing.T) {
	// Set up a temp directory mimicking node_modules with .js only
	dir := t.TempDir()
	pkgDir := filepath.Join(dir, "node_modules", "@tokens", "colors")
	if err := os.MkdirAll(pkgDir, 0755); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(filepath.Join(pkgDir, "package.json"), []byte(`{
		"name": "@tokens/colors",
		"main": "index.js"
	}`), 0644); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(filepath.Join(pkgDir, "index.js"), []byte(
		`/** @typedef {'red' | 'blue' | 'green'} ColorVariant */
export {};`,
	), 0644); err != nil {
		t.Fatal(err)
	}

	qm, err := Q.NewQueryManager(Q.GenerateQueries())
	if err != nil {
		t.Fatal(err)
	}
	defer qm.Close()

	resolver := &ExternalTypeResolver{
		projectRoot:  dir,
		queryManager: qm,
	}

	def, pkg, found := resolver.ResolveType("@tokens/colors", "ColorVariant")
	if !found {
		t.Fatal("expected to find ColorVariant")
	}
	if pkg != "@tokens/colors" {
		t.Errorf("package = %q, want %q", pkg, "@tokens/colors")
	}
	if def != "'red' | 'blue' | 'green'" {
		t.Errorf("definition = %q, want %q", def, "'red' | 'blue' | 'green'")
	}
}

func TestResolveType_CacheHit(t *testing.T) {
	dir := t.TempDir()
	pkgDir := filepath.Join(dir, "node_modules", "my-pkg")
	if err := os.MkdirAll(pkgDir, 0755); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(filepath.Join(pkgDir, "package.json"), []byte(`{
		"name": "my-pkg",
		"types": "index.d.ts"
	}`), 0644); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(filepath.Join(pkgDir, "index.d.ts"), []byte(
		`export type MyType = 'a' | 'b';`,
	), 0644); err != nil {
		t.Fatal(err)
	}

	qm, err := Q.NewQueryManager(Q.GenerateQueries())
	if err != nil {
		t.Fatal(err)
	}
	defer qm.Close()

	resolver := &ExternalTypeResolver{
		projectRoot:  dir,
		queryManager: qm,
	}

	// First call populates cache
	def1, _, found1 := resolver.ResolveType("my-pkg", "MyType")
	if !found1 {
		t.Fatal("first call: expected to find MyType")
	}

	// Second call should use cache (even if we remove the file)
	if err := os.Remove(filepath.Join(pkgDir, "index.d.ts")); err != nil {
		t.Fatal(err)
	}

	def2, _, found2 := resolver.ResolveType("my-pkg", "MyType")
	if !found2 {
		t.Fatal("second call: expected to find MyType from cache")
	}
	if def1 != def2 {
		t.Errorf("cached result differs: %q vs %q", def1, def2)
	}
}

func TestResolveType_NotFound(t *testing.T) {
	resolver := &ExternalTypeResolver{
		projectRoot: t.TempDir(),
	}

	_, _, found := resolver.ResolveType("nonexistent-pkg", "SomeType")
	if found {
		t.Error("expected not to find type from nonexistent package")
	}
}

func TestResolveType_RelativeImport(t *testing.T) {
	resolver := &ExternalTypeResolver{
		projectRoot: t.TempDir(),
	}

	_, _, found := resolver.ResolveType("./local", "SomeType")
	if found {
		t.Error("expected not to resolve relative imports")
	}
}

func TestResolveImportSubpath(t *testing.T) {
	tests := []struct {
		name     string
		pkg      *M.PackageJSON
		subpath  string
		expected string
		wantErr  bool
	}{
		{
			name:     "types field for root",
			pkg:      &M.PackageJSON{Types: "index.d.ts"},
			subpath:  ".",
			expected: "index.d.ts",
		},
		{
			name:     "typings field for root",
			pkg:      &M.PackageJSON{Typings: "./dist/index.d.ts"},
			subpath:  ".",
			expected: "dist/index.d.ts",
		},
		{
			name:     "main field for root",
			pkg:      &M.PackageJSON{Main: "./dist/index.js"},
			subpath:  ".",
			expected: "dist/index.js",
		},
		{
			name: "exports map with types condition",
			pkg: &M.PackageJSON{
				Exports: map[string]any{
					".": map[string]any{
						"types":   "./dist/index.d.ts",
						"default": "./dist/index.js",
					},
				},
			},
			subpath:  ".",
			expected: "dist/index.d.ts",
		},
		{
			name: "exports map with subpath",
			pkg: &M.PackageJSON{
				Exports: map[string]any{
					"./lib/types": "./dist/lib/types.js",
				},
			},
			subpath:  "./lib/types",
			expected: "dist/lib/types.js",
		},
		{
			name: "exports map with wildcard",
			pkg: &M.PackageJSON{
				Exports: map[string]any{
					"./lib/*": "./dist/lib/*.js",
				},
			},
			subpath:  "./lib/types",
			expected: "dist/lib/types.js",
		},
		{
			name:    "no resolution possible",
			pkg:     &M.PackageJSON{},
			subpath: ".",
			wantErr: true,
		},
		{
			name:    "nil package",
			pkg:     nil,
			subpath: ".",
			wantErr: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := M.ResolveImportSubpath(tc.pkg, tc.subpath)
			if tc.wantErr {
				if err == nil {
					t.Errorf("expected error, got %q", got)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.expected {
				t.Errorf("got %q, want %q", got, tc.expected)
			}
		})
	}
}
