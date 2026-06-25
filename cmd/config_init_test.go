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
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"

	"bennypowers.dev/cem/cmd"
	"bennypowers.dev/cem/internal/platform/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
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
		name           string
		existing       string
		detected       string
		fallback       string
		wantValue      string
		wantAnnotation string
	}{
		{
			name:           "existing wins",
			existing:       "elements/**/*.ts",
			detected:       "src/**/*.ts",
			fallback:       "**/*.ts",
			wantValue:      "elements/**/*.ts",
			wantAnnotation: "detected: src/**/*.ts",
		},
		{
			name:           "detected when no existing",
			existing:       "",
			detected:       "src/**/*.ts",
			fallback:       "**/*.ts",
			wantValue:      "src/**/*.ts",
			wantAnnotation: "",
		},
		{
			name:           "fallback when nothing else",
			existing:       "",
			detected:       "",
			fallback:       "**/*.ts",
			wantValue:      "**/*.ts",
			wantAnnotation: "",
		},
		{
			name:           "no annotation when existing matches detected",
			existing:       "src/**/*.ts",
			detected:       "src/**/*.ts",
			fallback:       "**/*.ts",
			wantValue:      "src/**/*.ts",
			wantAnnotation: "",
		},
		{
			name:           "all empty",
			existing:       "",
			detected:       "",
			fallback:       "",
			wantValue:      "",
			wantAnnotation: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fv := cmd.FieldValue{
				Existing: tt.existing,
				Detected: tt.detected,
				Fallback: tt.fallback,
			}
			assert.Equal(t, tt.wantValue, fv.Value())
			assert.Equal(t, tt.wantAnnotation, fv.Annotation())
		})
	}
}


func configInitGolden(name string) string {
	return filepath.Join("testdata", "goldens", "config-init", name)
}

func TestInitConfigMarshalYAML_Full(t *testing.T) {
	cfg := cmd.InitConfig{
		SourceControlRootUrl: "https://github.com/example/project/tree/main/",
		Generate: &cmd.InitGenerateConfig{
			Files:  []string{"elements/**/*.ts", "base/**/*.ts"},
			Output: "custom-elements.json",
			DesignTokens: &cmd.InitDesignTokensConfig{
				Spec:   "design-tokens.json",
				Prefix: "demo",
			},
			DemoDiscovery: &cmd.InitDemoDiscoveryConfig{
				FileGlob:    "elements/**/demo/*.html",
				URLPattern:  "elements/:tag/demo/:demo.html",
				URLTemplate: "https://example.com/{{.tag}}/{{.demo}}/",
			},
		},
		Serve: &cmd.InitServeConfig{
			Port: 8000,
			ImportMap: &cmd.InitImportMapConfig{
				Generate: true,
			},
			Transforms: &cmd.InitTransformsConfig{
				CSS: &cmd.InitCSSConfig{
					Enabled: true,
					Include: []string{"elements/**/*.css"},
					Exclude: []string{"demo/**/*.css", "**/*.min.css"},
				},
			},
			Demos: &cmd.InitDemosConfig{
				Rendering: "shadow",
			},
		},
	}

	got, err := yaml.Marshal(cfg)
	require.NoError(t, err)

	golden := configInitGolden("full.yaml")
	if *testutil.Update {
		require.NoError(t, os.WriteFile(golden, got, 0o644))
	}

	want, err := os.ReadFile(golden)
	require.NoError(t, err)
	assert.Equal(t, string(want), string(got))
}

func TestInitConfigMarshalYAML_Minimal(t *testing.T) {
	cfg := cmd.InitConfig{
		Generate: &cmd.InitGenerateConfig{
			Files:  []string{"src/**/*.ts"},
			Output: "custom-elements.json",
		},
	}

	got, err := yaml.Marshal(cfg)
	require.NoError(t, err)

	golden := configInitGolden("minimal.yaml")
	if *testutil.Update {
		require.NoError(t, os.WriteFile(golden, got, 0o644))
	}

	want, err := os.ReadFile(golden)
	require.NoError(t, err)
	assert.Equal(t, string(want), string(got))
}
