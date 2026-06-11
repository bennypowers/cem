package config

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"bennypowers.dev/cem/serve/middleware/types"
)

// Tier 1 pure function tests: all validation depends on formal parameters,
// no I/O or side effects. Inline assertions are appropriate.

func TestValidate_NilConfig(t *testing.T) {
	errs := Validate(nil, ValidateOptions{})
	if errs != nil {
		t.Errorf("expected nil, got %v", errs)
	}
}

func TestValidate_EmptyConfig(t *testing.T) {
	errs := Validate(&CemConfig{}, ValidateOptions{})
	if errs != nil {
		t.Errorf("expected nil, got %v", errs)
	}
}

func TestValidate_RenderingMode(t *testing.T) {
	tests := []struct {
		mode    string
		wantErr bool
	}{
		{"", false},
		{"light", false},
		{"shadow", false},
		{"iframe", false},
		{"chromeless", false},
		{"dark", true},
		{"LIGHT", true},
		{"Light", true},
	}
	for _, tt := range tests {
		t.Run(tt.mode, func(t *testing.T) {
			cfg := &CemConfig{Serve: ServeConfig{Demos: DemosConfig{Rendering: tt.mode}}}
			errs := Validate(cfg, ValidateOptions{})
			if tt.wantErr && errs == nil {
				t.Errorf("expected error for mode %q", tt.mode)
			}
			if !tt.wantErr && errs != nil {
				t.Errorf("unexpected error for mode %q: %v", tt.mode, errs)
			}
			if tt.wantErr && errs != nil {
				if errs[0].Field != "serve.demos.rendering" {
					t.Errorf("expected field serve.demos.rendering, got %q", errs[0].Field)
				}
			}
		})
	}
}

func TestValidate_ESTarget(t *testing.T) {
	valid := []string{
		"", "es2015", "es2016", "es2017", "es2018", "es2019",
		"es2020", "es2021", "es2022", "es2023", "esnext",
	}
	invalid := []string{"es2014", "es2024", "es5", "es6", "latest", "ES2022", "ESNext", "ESNEXT"}

	for _, target := range valid {
		t.Run("valid_"+target, func(t *testing.T) {
			cfg := &CemConfig{Serve: ServeConfig{
				Transforms: TransformsConfig{
					TypeScript: TypeScriptTransformConfig{Target: target},
				},
			}}
			errs := Validate(cfg, ValidateOptions{})
			if errs != nil {
				t.Errorf("unexpected error for target %q: %v", target, errs)
			}
		})
	}
	for _, target := range invalid {
		t.Run("invalid_"+target, func(t *testing.T) {
			cfg := &CemConfig{Serve: ServeConfig{
				Transforms: TransformsConfig{
					TypeScript: TypeScriptTransformConfig{Target: target},
				},
			}}
			errs := Validate(cfg, ValidateOptions{})
			if errs == nil {
				t.Errorf("expected error for target %q", target)
			}
			if errs != nil && errs[0].Field != "serve.transforms.typescript.target" {
				t.Errorf("expected field serve.transforms.typescript.target, got %q", errs[0].Field)
			}
		})
	}
}

func TestValidate_Port(t *testing.T) {
	tests := []struct {
		name    string
		port    int
		wantErr bool
	}{
		{"zero", 0, false},
		{"one", 1, false},
		{"common", 8000, false},
		{"max", 65535, false},
		{"over_max", 65536, true},
		{"negative", -1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &CemConfig{Serve: ServeConfig{Port: tt.port}}
			errs := Validate(cfg, ValidateOptions{})
			if tt.wantErr && errs == nil {
				t.Errorf("expected error for port %d", tt.port)
			}
			if !tt.wantErr && errs != nil {
				t.Errorf("unexpected error for port %d: %v", tt.port, errs)
			}
		})
	}
}

func TestValidate_DesignTokensPrefix(t *testing.T) {
	tests := []struct {
		prefix  string
		wantErr bool
	}{
		{"", false},
		{"rh", false},
		{"my-ds", false},
		{"-rh", true},
		{"--rh", true},
		{"---", true},
	}
	for _, tt := range tests {
		t.Run(tt.prefix, func(t *testing.T) {
			cfg := &CemConfig{Generate: GenerateConfig{
				DesignTokens: DesignTokensConfig{Prefix: tt.prefix},
			}}
			errs := Validate(cfg, ValidateOptions{})
			if tt.wantErr && errs == nil {
				t.Errorf("expected error for prefix %q", tt.prefix)
			}
			if !tt.wantErr && errs != nil {
				t.Errorf("unexpected error for prefix %q: %v", tt.prefix, errs)
			}
		})
	}
}

func TestValidate_MultipleErrors(t *testing.T) {
	cfg := &CemConfig{
		Serve: ServeConfig{
			Port:  -1,
			Demos: DemosConfig{Rendering: "invalid"},
			Transforms: TransformsConfig{
				TypeScript: TypeScriptTransformConfig{Target: "es5"},
			},
		},
		Generate: GenerateConfig{
			DesignTokens: DesignTokensConfig{Prefix: "--bad"},
		},
	}
	errs := Validate(cfg, ValidateOptions{})
	if len(errs) != 4 {
		t.Errorf("expected 4 errors, got %d: %v", len(errs), errs)
	}
}

func TestValidate_DelegatedURLRewrites(t *testing.T) {
	cfg := &CemConfig{
		Serve: ServeConfig{
			URLRewrites: []URLRewrite{{URLPattern: "/foo", URLTemplate: "/bar"}},
		},
	}

	t.Run("nil_validator_skips", func(t *testing.T) {
		errs := Validate(cfg, ValidateOptions{})
		if errs != nil {
			t.Errorf("expected nil with no validator, got %v", errs)
		}
	})

	t.Run("validator_called", func(t *testing.T) {
		called := false
		errs := Validate(cfg, ValidateOptions{
			ValidateURLRewrites: func(rewrites []URLRewrite) error {
				called = true
				return errors.New("bad pattern")
			},
		})
		if !called {
			t.Error("validator was not called")
		}
		if len(errs) != 1 || errs[0].Field != "serve.urlRewrites" {
			t.Errorf("unexpected errs: %v", errs)
		}
	})

	t.Run("validator_success", func(t *testing.T) {
		errs := Validate(cfg, ValidateOptions{
			ValidateURLRewrites: func(rewrites []URLRewrite) error {
				return nil
			},
		})
		if errs != nil {
			t.Errorf("expected nil, got %v", errs)
		}
	})

	t.Run("empty_rewrites_skips_validator", func(t *testing.T) {
		called := false
		errs := Validate(&CemConfig{}, ValidateOptions{
			ValidateURLRewrites: func(rewrites []URLRewrite) error {
				called = true
				return errors.New("should not be called")
			},
		})
		if called {
			t.Error("validator should not be called for empty rewrites")
		}
		if errs != nil {
			t.Errorf("expected nil, got %v", errs)
		}
	})
}

func TestValidate_DelegatedDemoDiscovery(t *testing.T) {
	cfg := &CemConfig{}

	t.Run("nil_validator_skips", func(t *testing.T) {
		errs := Validate(cfg, ValidateOptions{})
		if errs != nil {
			t.Errorf("expected nil, got %v", errs)
		}
	})

	t.Run("validator_called", func(t *testing.T) {
		errs := Validate(cfg, ValidateOptions{
			ValidateDemoDiscovery: func(c *CemConfig) error {
				return errors.New("bad template")
			},
		})
		if len(errs) != 1 || errs[0].Field != "generate.demoDiscovery" {
			t.Errorf("unexpected errs: %v", errs)
		}
	})
}

type mockFS struct {
	files    map[string]bool
	contents map[string][]byte
}

func (m *mockFS) Exists(path string) bool {
	return m.files[path]
}

func (m *mockFS) ReadFile(path string) ([]byte, error) {
	if data, ok := m.contents[path]; ok {
		return data, nil
	}
	return nil, fmt.Errorf("file not found: %s", path)
}

func TestValidate_FilesystemChecks(t *testing.T) {
	fs := &mockFS{files: map[string]bool{
		"/root/exists.json": true,
		"/root/tokens.json": true,
	}}

	t.Run("skipped_when_disabled", func(t *testing.T) {
		cfg := &CemConfig{Serve: ServeConfig{
			ImportMap: types.ImportMapConfig{OverrideFile: "missing.json"},
		}}
		errs := Validate(cfg, ValidateOptions{CheckIO: false})
		if errs != nil {
			t.Errorf("expected nil when filesystem check disabled, got %v", errs)
		}
	})

	t.Run("skipped_when_no_fs", func(t *testing.T) {
		cfg := &CemConfig{Serve: ServeConfig{
			ImportMap: types.ImportMapConfig{OverrideFile: "missing.json"},
		}}
		errs := Validate(cfg, ValidateOptions{CheckIO: true})
		if errs != nil {
			t.Errorf("expected nil when FS is nil, got %v", errs)
		}
	})

	t.Run("importmap_override_found", func(t *testing.T) {
		cfg := &CemConfig{Serve: ServeConfig{
			ImportMap: types.ImportMapConfig{OverrideFile: "exists.json"},
		}}
		errs := Validate(cfg, ValidateOptions{CheckIO: true, Root: "/root", IO: fs})
		if errs != nil {
			t.Errorf("expected nil, got %v", errs)
		}
	})

	t.Run("importmap_override_missing", func(t *testing.T) {
		cfg := &CemConfig{Serve: ServeConfig{
			ImportMap: types.ImportMapConfig{OverrideFile: "missing.json"},
		}}
		errs := Validate(cfg, ValidateOptions{CheckIO: true, Root: "/root", IO: fs})
		if len(errs) != 1 || errs[0].Field != "serve.importMap.overrideFile" {
			t.Errorf("expected importMap error, got %v", errs)
		}
	})

	t.Run("design_tokens_found", func(t *testing.T) {
		cfg := &CemConfig{Generate: GenerateConfig{
			DesignTokens: DesignTokensConfig{Spec: "tokens.json"},
		}}
		errs := Validate(cfg, ValidateOptions{CheckIO: true, Root: "/root", IO: fs})
		if errs != nil {
			t.Errorf("expected nil, got %v", errs)
		}
	})

	t.Run("design_tokens_missing", func(t *testing.T) {
		cfg := &CemConfig{Generate: GenerateConfig{
			DesignTokens: DesignTokensConfig{Spec: "nope.json"},
		}}
		errs := Validate(cfg, ValidateOptions{CheckIO: true, Root: "/root", IO: fs})
		if len(errs) != 1 || errs[0].Field != "generate.designTokens.spec" {
			t.Errorf("expected designTokens error, got %v", errs)
		}
	})

	t.Run("npm_specifier_skips_filesystem", func(t *testing.T) {
		cfg := &CemConfig{Generate: GenerateConfig{
			DesignTokens: DesignTokensConfig{Spec: "npm:@rhds/tokens/tokens.json"},
		}}
		errs := Validate(cfg, ValidateOptions{CheckIO: true, Root: "/root", IO: fs})
		if errs != nil {
			t.Errorf("expected nil for npm specifier, got %v", errs)
		}
	})

	t.Run("jsr_specifier_skips_filesystem", func(t *testing.T) {
		cfg := &CemConfig{Generate: GenerateConfig{
			DesignTokens: DesignTokensConfig{Spec: "jsr:@example/tokens"},
		}}
		errs := Validate(cfg, ValidateOptions{CheckIO: true, Root: "/root", IO: fs})
		if errs != nil {
			t.Errorf("expected nil for jsr specifier, got %v", errs)
		}
	})

	t.Run("https_specifier_skips_filesystem", func(t *testing.T) {
		cfg := &CemConfig{Generate: GenerateConfig{
			DesignTokens: DesignTokensConfig{Spec: "https://cdn.example.com/tokens.json"},
		}}
		errs := Validate(cfg, ValidateOptions{CheckIO: true, Root: "/root", IO: fs})
		if errs != nil {
			t.Errorf("expected nil for https specifier, got %v", errs)
		}
	})

	t.Run("http_specifier_skips_filesystem", func(t *testing.T) {
		cfg := &CemConfig{Generate: GenerateConfig{
			DesignTokens: DesignTokensConfig{Spec: "http://cdn.example.com/tokens.json"},
		}}
		errs := Validate(cfg, ValidateOptions{CheckIO: true, Root: "/root", IO: fs})
		if errs != nil {
			t.Errorf("expected nil for http specifier, got %v", errs)
		}
	})

	t.Run("absolute_importmap_path", func(t *testing.T) {
		absFS := &mockFS{files: map[string]bool{"/abs/override.json": true}}
		cfg := &CemConfig{Serve: ServeConfig{
			ImportMap: types.ImportMapConfig{OverrideFile: "/abs/override.json"},
		}}
		errs := Validate(cfg, ValidateOptions{CheckIO: true, Root: "/root", IO: absFS})
		if errs != nil {
			t.Errorf("expected nil for absolute path, got %v", errs)
		}
	})

	t.Run("absolute_tokens_path", func(t *testing.T) {
		absFS := &mockFS{files: map[string]bool{"/abs/tokens.json": true}}
		cfg := &CemConfig{Generate: GenerateConfig{
			DesignTokens: DesignTokensConfig{Spec: "/abs/tokens.json"},
		}}
		errs := Validate(cfg, ValidateOptions{CheckIO: true, Root: "/root", IO: absFS})
		if errs != nil {
			t.Errorf("expected nil for absolute path, got %v", errs)
		}
	})
}

func TestValidationErrors(t *testing.T) {
	t.Run("nil_for_empty", func(t *testing.T) {
		if err := ValidationErrors(nil); err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})

	t.Run("formats_single", func(t *testing.T) {
		err := ValidationErrors([]ValidationError{{Field: "f", Message: "m", Value: "v"}})
		if err == nil {
			t.Fatal("expected error")
		}
		want := "config validation failed:\n  f: m (got \"v\")"
		if err.Error() != want {
			t.Errorf("got %q, want %q", err.Error(), want)
		}
	})

	t.Run("formats_multiple", func(t *testing.T) {
		err := ValidationErrors([]ValidationError{
			{Field: "a", Message: "bad"},
			{Field: "b", Message: "worse", Value: "x"},
		})
		if err == nil {
			t.Fatal("expected error")
		}
		want := "config validation failed:\n  a: bad\n  b: worse (got \"x\")"
		if err.Error() != want {
			t.Errorf("got %q, want %q", err.Error(), want)
		}
	})
}

// Tier 1: param extraction and consistency are pure functions.
// Inline assertions justified: testing pure param extraction logic.

func TestValidate_DemoDiscoveryParamConsistency(t *testing.T) {
	t.Run("both_empty_ok", func(t *testing.T) {
		cfg := &CemConfig{}
		errs := Validate(cfg, ValidateOptions{})
		for _, e := range errs {
			if e.Field == "generate.demoDiscovery" {
				t.Errorf("unexpected demoDiscovery error: %v", e)
			}
		}
	})

	t.Run("pattern_without_template_warns", func(t *testing.T) {
		cfg := &CemConfig{Generate: GenerateConfig{
			DemoDiscovery: DemoDiscoveryConfig{URLPattern: "/elements/:tag/demo/:demo.html"},
		}}
		errs := Validate(cfg, ValidateOptions{})
		found := false
		for _, e := range errs {
			if e.Field == "generate.demoDiscovery" && e.Severity == SeverityWarning {
				found = true
			}
		}
		if !found {
			t.Error("expected warning for pattern without template")
		}
	})

	t.Run("template_without_pattern_warns", func(t *testing.T) {
		cfg := &CemConfig{Generate: GenerateConfig{
			DemoDiscovery: DemoDiscoveryConfig{URLTemplate: "/elements/{{.tag}}/demo/{{.demo}}/"},
		}}
		errs := Validate(cfg, ValidateOptions{})
		found := false
		for _, e := range errs {
			if e.Field == "generate.demoDiscovery" && e.Severity == SeverityWarning {
				found = true
			}
		}
		if !found {
			t.Error("expected warning for template without pattern")
		}
	})

	t.Run("matching_params_ok", func(t *testing.T) {
		cfg := &CemConfig{Generate: GenerateConfig{
			DemoDiscovery: DemoDiscoveryConfig{
				URLPattern:  "/elements/:tag/demo/:demo.html",
				URLTemplate: "/elements/{{.tag}}/demo/{{.demo}}/",
			},
		}}
		errs := Validate(cfg, ValidateOptions{})
		for _, e := range errs {
			if e.Field == "generate.demoDiscovery.urlTemplate" || e.Field == "generate.demoDiscovery.urlPattern" {
				t.Errorf("unexpected param error: %v", e)
			}
		}
	})

	t.Run("template_references_unknown_param", func(t *testing.T) {
		cfg := &CemConfig{Generate: GenerateConfig{
			DemoDiscovery: DemoDiscoveryConfig{
				URLPattern:  "/elements/:tag/demo/:demo.html",
				URLTemplate: "/elements/{{.tag}}/demo/{{.unknown}}/",
			},
		}}
		errs := Validate(cfg, ValidateOptions{})
		found := false
		for _, e := range errs {
			if e.Field == "generate.demoDiscovery.urlTemplate" && e.Severity == SeverityError {
				found = true
			}
		}
		if !found {
			t.Error("expected error for template referencing unknown param")
		}
	})

	t.Run("pattern_defines_unused_param", func(t *testing.T) {
		cfg := &CemConfig{Generate: GenerateConfig{
			DemoDiscovery: DemoDiscoveryConfig{
				URLPattern:  "/elements/:tag/demo/:demo.html",
				URLTemplate: "/elements/{{.tag}}/",
			},
		}}
		errs := Validate(cfg, ValidateOptions{})
		found := false
		for _, e := range errs {
			if e.Field == "generate.demoDiscovery.urlPattern" && e.Severity == SeverityWarning {
				found = true
			}
		}
		if !found {
			t.Error("expected warning for pattern defining unused param")
		}
	})
}

func TestValidate_OutputMismatch(t *testing.T) {
	t.Run("no_output_no_error", func(t *testing.T) {
		cfg := &CemConfig{}
		fs := &mockFS{files: map[string]bool{}, contents: map[string][]byte{}}
		errs := Validate(cfg, ValidateOptions{CheckIO: true, Root: "/root", IO: fs})
		for _, e := range errs {
			if e.Field == "generate.output" {
				t.Errorf("unexpected output error: %v", e)
			}
		}
	})

	t.Run("matching_output_ok", func(t *testing.T) {
		cfg := &CemConfig{Generate: GenerateConfig{Output: "custom-elements.json"}}
		fs := &mockFS{
			files:    map[string]bool{"/root/package.json": true},
			contents: map[string][]byte{"/root/package.json": []byte(`{"customElements": "custom-elements.json"}`)},
		}
		errs := Validate(cfg, ValidateOptions{CheckIO: true, Root: "/root", IO: fs})
		for _, e := range errs {
			if e.Field == "generate.output" {
				t.Errorf("unexpected output error: %v", e)
			}
		}
	})

	t.Run("mismatching_output_warns", func(t *testing.T) {
		cfg := &CemConfig{Generate: GenerateConfig{Output: "custom-elements.json"}}
		fs := &mockFS{
			files:    map[string]bool{"/root/package.json": true},
			contents: map[string][]byte{"/root/package.json": []byte(`{"customElements": "manifest.json"}`)},
		}
		errs := Validate(cfg, ValidateOptions{CheckIO: true, Root: "/root", IO: fs})
		found := false
		for _, e := range errs {
			if e.Field == "generate.output" && e.Severity == SeverityWarning {
				found = true
			}
		}
		if !found {
			t.Error("expected warning for output mismatch")
		}
	})

	t.Run("no_package_json_ok", func(t *testing.T) {
		cfg := &CemConfig{Generate: GenerateConfig{Output: "custom-elements.json"}}
		fs := &mockFS{files: map[string]bool{}, contents: map[string][]byte{}}
		errs := Validate(cfg, ValidateOptions{CheckIO: true, Root: "/root", IO: fs})
		for _, e := range errs {
			if e.Field == "generate.output" {
				t.Errorf("unexpected output error: %v", e)
			}
		}
	})

	t.Run("no_customElements_field_ok", func(t *testing.T) {
		cfg := &CemConfig{Generate: GenerateConfig{Output: "custom-elements.json"}}
		fs := &mockFS{
			files:    map[string]bool{"/root/package.json": true},
			contents: map[string][]byte{"/root/package.json": []byte(`{"name": "my-pkg"}`)},
		}
		errs := Validate(cfg, ValidateOptions{CheckIO: true, Root: "/root", IO: fs})
		for _, e := range errs {
			if e.Field == "generate.output" {
				t.Errorf("unexpected output error: %v", e)
			}
		}
	})
}

func TestExtractURLPatternParams(t *testing.T) {
	tests := []struct {
		pattern string
		want    string
	}{
		{"/elements/:tag/demo/:demo.html", "demo,tag"},
		{"/api/:version/:resource", "resource,version"},
		{"/:path*", "path"},
		{"", ""},
		{"/no/params/here", ""},
	}
	for _, tt := range tests {
		t.Run(tt.pattern, func(t *testing.T) {
			got := extractURLPatternParams(tt.pattern)
			result := strings.Join(got, ",")
			if result != tt.want {
				t.Errorf("extractURLPatternParams(%q) = %q, want %q", tt.pattern, result, tt.want)
			}
		})
	}
}

func TestExtractTemplateParams(t *testing.T) {
	tests := []struct {
		tmpl string
		want string
	}{
		{"/elements/{{.tag}}/demo/{{.demo}}/", "demo,tag"},
		{"{{.a}}{{.b}}{{.c}}", "a,b,c"},
		{"{{.tag | lower}}", "tag"},
		{"no templates here", ""},
		{"{{.tag}}{{.tag}}", "tag"},
		{"{{if .tag}}{{.demo}}{{end}}", "demo,tag"},
		{"{{if .a}}{{.b}}{{else}}{{.c}}{{end}}", "a,b,c"},
		{"{{range .items}}{{.name}}{{end}}", "items,name"},
	}
	for _, tt := range tests {
		t.Run(tt.tmpl, func(t *testing.T) {
			got := extractTemplateParams(tt.tmpl)
			result := strings.Join(got, ",")
			if result != tt.want {
				t.Errorf("extractTemplateParams(%q) = %q, want %q", tt.tmpl, result, tt.want)
			}
		})
	}
}
