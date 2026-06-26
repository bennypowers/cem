/*
Copyright © 2026 Benny Powers <web@bennypowers.com>

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
package cmd_test

import (
	"path/filepath"
	"testing"
	"testing/fstest"

	"bennypowers.dev/cem/cmd"
	IC "bennypowers.dev/cem/internal/config"
	"bennypowers.dev/cem/internal/platform/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// inline assertions: pure string->string function, table-driven with simple equality checks
func TestNormalizeGitURL(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want string
	}{
		{
			name: "ssh github",
			raw:  "git@github.com:bennypowers/cem.git",
			want: "https://github.com/bennypowers/cem/",
		},
		{
			name: "https github with .git",
			raw:  "https://github.com/bennypowers/cem.git",
			want: "https://github.com/bennypowers/cem/",
		},
		{
			name: "https github without .git",
			raw:  "https://github.com/bennypowers/cem",
			want: "https://github.com/bennypowers/cem/",
		},
		{
			name: "ssh gitlab",
			raw:  "git@gitlab.com:group/project.git",
			want: "https://gitlab.com/group/project/",
		},
		{
			name: "https gitlab",
			raw:  "https://gitlab.com/group/project.git",
			want: "https://gitlab.com/group/project/",
		},
		{
			name: "ssh bitbucket",
			raw:  "git@bitbucket.org:team/repo.git",
			want: "https://bitbucket.org/team/repo/",
		},
		{
			name: "ssh:// protocol",
			raw:  "ssh://git@github.com/bennypowers/cem.git",
			want: "https://github.com/bennypowers/cem/",
		},
		{
			name: "ssh:// with port",
			raw:  "ssh://git@github.com:22/bennypowers/cem.git",
			want: "https://github.com/bennypowers/cem/",
		},
		{
			name: "empty",
			raw:  "",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cmd.NormalizeGitURL(tt.raw)
			assert.Equal(t, tt.want, got)
		})
	}
}

// inline assertions: pure fs.FS->[]string function, MapFS fixtures with simple slice equality
func TestDetectSourceFiles(t *testing.T) {
	tests := []struct {
		name  string
		files map[string]*fstest.MapFile
		want  []string
	}{
		{
			name: "single directory",
			files: map[string]*fstest.MapFile{
				"src/button.ts":    {},
				"src/card.ts":      {},
				"src/icon/icon.ts": {},
			},
			want: []string{"src/**/*.ts"},
		},
		{
			name: "multiple directories",
			files: map[string]*fstest.MapFile{
				"elements/button/button.ts": {},
				"base/base-element.ts":      {},
				"mixins/resizable.ts":       {},
			},
			want: []string{"base/**/*.ts", "elements/**/*.ts", "mixins/**/*.ts"},
		},
		{
			name: "skips .d.ts files",
			files: map[string]*fstest.MapFile{
				"src/button.ts":   {},
				"src/button.d.ts": {},
			},
			want: []string{"src/**/*.ts"},
		},
		{
			name: "skips node_modules",
			files: map[string]*fstest.MapFile{
				"src/button.ts":                     {},
				"node_modules/@lit/element/index.ts": {},
			},
			want: []string{"src/**/*.ts"},
		},
		{
			name: "no ts files",
			files: map[string]*fstest.MapFile{
				"src/readme.md": {},
			},
			want: nil,
		},
		{
			name: "skips dist and test",
			files: map[string]*fstest.MapFile{
				"src/button.ts":       {},
				"dist/button.ts":      {},
				"test/button.test.ts": {},
			},
			want: []string{"src/**/*.ts"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fsys := fstest.MapFS(tt.files)
			got, err := cmd.DetectSourceFiles(fsys)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

// inline assertions: pure function returning two strings, simple equality checks
func TestDetectDemoFiles(t *testing.T) {
	tests := []struct {
		name        string
		files       map[string]*fstest.MapFile
		wantGlob    string
		wantPattern string
	}{
		{
			name: "kitchen-sink layout",
			files: map[string]*fstest.MapFile{
				"elements/my-button/demo/basic.html":    {},
				"elements/my-button/demo/variants.html": {},
				"elements/my-card/demo/basic.html":      {},
			},
			wantGlob:    "elements/**/demo/*.html",
			wantPattern: "elements/:tag/demo/:demo.html",
		},
		{
			name: "src layout",
			files: map[string]*fstest.MapFile{
				"src/button/demo/basic.html": {},
				"src/card/demo/basic.html":   {},
			},
			wantGlob:    "src/**/demo/*.html",
			wantPattern: "src/:tag/demo/:demo.html",
		},
		{
			name: "no demo files",
			files: map[string]*fstest.MapFile{
				"src/button.ts": {},
			},
			wantGlob:    "",
			wantPattern: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fsys := fstest.MapFS(tt.files)
			glob, pattern, err := cmd.DetectDemoFiles(fsys)
			require.NoError(t, err)
			assert.Equal(t, tt.wantGlob, glob)
			assert.Equal(t, tt.wantPattern, pattern)
		})
	}
}

// inline assertions: pure function returning two slices, simple equality checks
func TestDetectCSSPatterns(t *testing.T) {
	tests := []struct {
		name           string
		files          map[string]*fstest.MapFile
		sourcePatterns []string
		wantInclude    []string
		wantExclude    []string
	}{
		{
			name: "colocated CSS",
			files: map[string]*fstest.MapFile{
				"elements/button/button.ts":  {},
				"elements/button/button.css": {},
				"elements/card/card.ts":      {},
				"elements/card/card.css":     {},
			},
			sourcePatterns: []string{"elements/**/*.ts"},
			wantInclude:    []string{"elements/**/*.css"},
			wantExclude:    []string{"demo/**/*.css", "**/*.min.css"},
		},
		{
			name: "no CSS files",
			files: map[string]*fstest.MapFile{
				"elements/button/button.ts": {},
			},
			sourcePatterns: []string{"elements/**/*.ts"},
			wantInclude:    nil,
			wantExclude:    nil,
		},
		{
			name: "multiple source dirs",
			files: map[string]*fstest.MapFile{
				"src/button/button.ts":  {},
				"src/button/button.css": {},
				"lib/card/card.ts":      {},
				"lib/card/card.css":     {},
			},
			sourcePatterns: []string{"src/**/*.ts", "lib/**/*.ts"},
			wantInclude:    []string{"lib/**/*.css", "src/**/*.css"},
			wantExclude:    []string{"demo/**/*.css", "**/*.min.css"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fsys := fstest.MapFS(tt.files)
			include, exclude := cmd.DetectCSSPatterns(tt.sourcePatterns, fsys)
			assert.Equal(t, tt.wantInclude, include)
			assert.Equal(t, tt.wantExclude, exclude)
		})
	}
}

// inline assertions: pure struct methods with simple string returns
func TestFieldValue(t *testing.T) {
	tests := []struct {
		name      string
		existing  string
		detected  string
		fallback  string
		wantValue string
	}{
		{"existing wins", "elements/**/*.ts", "src/**/*.ts", "**/*.ts", "elements/**/*.ts"},
		{"detected when no existing", "", "src/**/*.ts", "**/*.ts", "src/**/*.ts"},
		{"fallback when nothing else", "", "", "**/*.ts", "**/*.ts"},
		{"existing matches detected", "src/**/*.ts", "src/**/*.ts", "**/*.ts", "src/**/*.ts"},
		{"all empty", "", "", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fv := cmd.FieldValue{
				Existing: tt.existing,
				Detected: tt.detected,
				Fallback: tt.fallback,
			}
			assert.Equal(t, tt.wantValue, fv.Value())
		})
	}
}

// inline assertions: pure function with simple int return
func TestDetectIndent(t *testing.T) {
	tests := []struct {
		name string
		data string
		want int
	}{
		{"two space", "key:\n  value: x\n", 2},
		{"four space", "key:\n    value: x\n", 4},
		{"tab-like 8", "key:\n        value: x\n", 8},
		{"empty", "", 2},
		{"no indent", "key: value\n", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, cmd.DetectIndent([]byte(tt.data)))
		})
	}
}

func TestMarshalConfigYAML_Full(t *testing.T) {
	cfg := &IC.CemConfig{
		SourceControlRootUrl: "https://github.com/example/project/tree/main/",
		Generate: IC.GenerateConfig{
			Files:  []string{"elements/**/*.ts", "base/**/*.ts"},
			Output: "custom-elements.json",
			DesignTokens: IC.DesignTokensConfig{
				Spec:   "design-tokens.json",
				Prefix: "demo",
			},
			DemoDiscovery: IC.DemoDiscoveryConfig{
				FileGlob:    "elements/**/demo/*.html",
				URLPattern:  "elements/:tag/demo/:demo.html",
				URLTemplate: "https://example.com/{{.tag}}/{{.demo}}/",
			},
		},
		Serve: IC.ServeConfig{
			Port: 8000,
			Demos: IC.DemosConfig{
				Rendering: "shadow",
			},
			Transforms: IC.TransformsConfig{
				CSS: IC.CSSTransformConfig{
					Enabled: true,
					Include: []string{"elements/**/*.css"},
					Exclude: []string{"demo/**/*.css", "**/*.min.css"},
				},
			},
		},
	}
	cfg.Serve.ImportMap.Generate = true

	got, err := cmd.MarshalConfig(cfg, "yaml", nil)
	require.NoError(t, err)

	goldenDir := filepath.Join("testdata", "goldens", "config-init")
	fs := testutil.LoadTestdataFS(t, goldenDir, goldenDir)
	testutil.CheckGolden(t, "full", got, testutil.GoldenOptions{
		Dir:       goldenDir,
		Extension: ".yaml",
		FS:        fs,
	})
}

func TestMarshalConfigYAML_Minimal(t *testing.T) {
	cfg := &IC.CemConfig{
		Generate: IC.GenerateConfig{
			Files:  []string{"src/**/*.ts"},
			Output: "custom-elements.json",
		},
	}

	got, err := cmd.MarshalConfig(cfg, "yaml", nil)
	require.NoError(t, err)

	goldenDir := filepath.Join("testdata", "goldens", "config-init")
	fs := testutil.LoadTestdataFS(t, goldenDir, goldenDir)
	testutil.CheckGolden(t, "minimal", got, testutil.GoldenOptions{
		Dir:       goldenDir,
		Extension: ".yaml",
		FS:        fs,
	})
}
