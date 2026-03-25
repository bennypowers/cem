package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/internal/config"
)

func TestLoadConfig_Yaml(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "cem.yaml")
	os.WriteFile(path, []byte(`generate:
  files:
    - "src/**/*.ts"
  output: custom-elements.json
`), 0o644)

	cfg, err := config.LoadConfig(path)
	if err != nil {
		t.Fatalf("LoadConfig() error: %v", err)
	}
	if len(cfg.Generate.Files) != 1 || cfg.Generate.Files[0] != "src/**/*.ts" {
		t.Errorf("Generate.Files = %v, want [src/**/*.ts]", cfg.Generate.Files)
	}
	if cfg.Generate.Output != "custom-elements.json" {
		t.Errorf("Generate.Output = %q, want custom-elements.json", cfg.Generate.Output)
	}
}

func TestLoadConfig_Json(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "cem.json")
	os.WriteFile(path, []byte(`{
  "generate": {
    "files": ["src/**/*.ts"],
    "output": "custom-elements.json"
  }
}`), 0o644)

	cfg, err := config.LoadConfig(path)
	if err != nil {
		t.Fatalf("LoadConfig() error: %v", err)
	}
	if len(cfg.Generate.Files) != 1 || cfg.Generate.Files[0] != "src/**/*.ts" {
		t.Errorf("Generate.Files = %v, want [src/**/*.ts]", cfg.Generate.Files)
	}
}

func TestLoadConfig_Jsonc(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "cem.jsonc")
	os.WriteFile(path, []byte(`{
  // Generate configuration
  "generate": {
    "files": ["src/**/*.ts"],
    "output": "custom-elements.json"
  }
}`), 0o644)

	cfg, err := config.LoadConfig(path)
	if err != nil {
		t.Fatalf("LoadConfig() error: %v", err)
	}
	if len(cfg.Generate.Files) != 1 || cfg.Generate.Files[0] != "src/**/*.ts" {
		t.Errorf("Generate.Files = %v, want [src/**/*.ts]", cfg.Generate.Files)
	}
}

func TestLoadConfig_JsoncTrailingCommas(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "cem.jsonc")
	os.WriteFile(path, []byte(`{
  "generate": {
    "files": [
      "src/**/*.ts",
      "lib/**/*.ts",
    ],
    "output": "custom-elements.json",
  },
}`), 0o644)

	cfg, err := config.LoadConfig(path)
	if err != nil {
		t.Fatalf("LoadConfig() error: %v", err)
	}
	if len(cfg.Generate.Files) != 2 {
		t.Errorf("Generate.Files has %d entries, want 2", len(cfg.Generate.Files))
	}
}

func TestLoadConfig_MissingFile(t *testing.T) {
	_, err := config.LoadConfig("/nonexistent/path/cem.yaml")
	if err == nil {
		t.Error("LoadConfig() expected error for missing file, got nil")
	}
}

func TestLoadConfig_InvalidYaml(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "cem.yaml")
	os.WriteFile(path, []byte("{{{{invalid yaml"), 0o644)

	_, err := config.LoadConfig(path)
	if err == nil {
		t.Error("LoadConfig() expected error for invalid yaml, got nil")
	}
}

func TestLoadConfig_InvalidJson(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "cem.json")
	os.WriteFile(path, []byte("{not valid json"), 0o644)

	_, err := config.LoadConfig(path)
	if err == nil {
		t.Error("LoadConfig() expected error for invalid json, got nil")
	}
}

func TestLoadConfig_EmptyFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "cem.yaml")
	os.WriteFile(path, []byte(""), 0o644)

	cfg, err := config.LoadConfig(path)
	if err != nil {
		t.Fatalf("LoadConfig() error: %v", err)
	}
	// Should return a valid zero-value config
	if cfg == nil {
		t.Fatal("LoadConfig() returned nil config for empty file")
	}
}

func TestLoadConfig_JsoncOnlyComments(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "cem.jsonc")
	os.WriteFile(path, []byte("// just a comment\n/* nothing else */"), 0o644)

	cfg, err := config.LoadConfig(path)
	if err != nil {
		t.Fatalf("LoadConfig() error: %v", err)
	}
	if cfg == nil {
		t.Fatal("LoadConfig() returned nil config for comment-only jsonc file")
	}
}

func TestLoadConfig_UnsupportedFormat(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "cem.toml")
	os.WriteFile(path, []byte("[generate]\nfiles = ['*.ts']"), 0o644)

	_, err := config.LoadConfig(path)
	if err == nil {
		t.Error("LoadConfig() expected error for unsupported format, got nil")
	}
}

func TestLoadConfig_EmptyJson(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "cem.json")
	os.WriteFile(path, []byte(""), 0o644)

	cfg, err := config.LoadConfig(path)
	if err != nil {
		t.Fatalf("LoadConfig() error: %v", err)
	}
	if cfg == nil {
		t.Fatal("LoadConfig() returned nil config for empty json file")
	}
}

func TestLoadConfig_EmptyJsonc(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "cem.jsonc")
	os.WriteFile(path, []byte(""), 0o644)

	cfg, err := config.LoadConfig(path)
	if err != nil {
		t.Fatalf("LoadConfig() error: %v", err)
	}
	if cfg == nil {
		t.Fatal("LoadConfig() returned nil config for empty jsonc file")
	}
}

func TestLoadConfig_InvalidJsonc(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "cem.jsonc")
	os.WriteFile(path, []byte("// comment\n{not valid json even after stripping}"), 0o644)

	_, err := config.LoadConfig(path)
	if err == nil {
		t.Error("LoadConfig() expected error for invalid jsonc, got nil")
	}
}

func TestLoadConfig_AllFields(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "cem.yaml")
	os.WriteFile(path, []byte(`generate:
  files:
    - "src/**/*.ts"
  exclude:
    - "**/*.test.ts"
  output: custom-elements.json
  noDefaultExcludes: true
  designTokens:
    spec: tokens/tokens.json
    prefix: my-prefix
  demoDiscovery:
    fileGlob: elements/*/demo/*.html
    urlPattern: /elements/:tag/demo/:demo.html
    urlTemplate: /elements/{{.tag}}/demo/{{.demo}}/
serve:
  port: 3000
  transforms:
    css:
      enabled: true
      include:
        - "src/**/*.css"
mcp:
  maxDescriptionLength: 1500
health:
  failBelow: 80
  disable:
    - demos
sourceControlRootUrl: https://github.com/example/repo/tree/main/
`), 0o644)

	cfg, err := config.LoadConfig(path)
	if err != nil {
		t.Fatalf("LoadConfig() error: %v", err)
	}
	if cfg.Generate.Output != "custom-elements.json" {
		t.Errorf("Generate.Output = %q", cfg.Generate.Output)
	}
	if !cfg.Generate.NoDefaultExcludes {
		t.Error("Generate.NoDefaultExcludes should be true")
	}
	if cfg.Generate.DesignTokens.Prefix != "my-prefix" {
		t.Errorf("DesignTokens.Prefix = %q", cfg.Generate.DesignTokens.Prefix)
	}
	if cfg.Generate.DemoDiscovery.FileGlob != "elements/*/demo/*.html" {
		t.Errorf("DemoDiscovery.FileGlob = %q", cfg.Generate.DemoDiscovery.FileGlob)
	}
	if cfg.Serve.Port != 3000 {
		t.Errorf("Serve.Port = %d", cfg.Serve.Port)
	}
	if cfg.MCP.MaxDescriptionLength != 1500 {
		t.Errorf("MCP.MaxDescriptionLength = %d", cfg.MCP.MaxDescriptionLength)
	}
	if cfg.Health.FailBelow != 80 {
		t.Errorf("Health.FailBelow = %d", cfg.Health.FailBelow)
	}
	if cfg.SourceControlRootUrl != "https://github.com/example/repo/tree/main/" {
		t.Errorf("SourceControlRootUrl = %q", cfg.SourceControlRootUrl)
	}
}
