package config_test

import (
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/internal/config"
)

func fixture(name string) string {
	return filepath.Join("testdata", name)
}

func TestLoadConfig_Yaml(t *testing.T) {
	cfg, err := config.LoadConfig(fixture("valid.yaml"))
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
	cfg, err := config.LoadConfig(fixture("valid.json"))
	if err != nil {
		t.Fatalf("LoadConfig() error: %v", err)
	}
	if len(cfg.Generate.Files) != 1 || cfg.Generate.Files[0] != "src/**/*.ts" {
		t.Errorf("Generate.Files = %v, want [src/**/*.ts]", cfg.Generate.Files)
	}
}

func TestLoadConfig_Jsonc(t *testing.T) {
	cfg, err := config.LoadConfig(fixture("valid.jsonc"))
	if err != nil {
		t.Fatalf("LoadConfig() error: %v", err)
	}
	if len(cfg.Generate.Files) != 1 || cfg.Generate.Files[0] != "src/**/*.ts" {
		t.Errorf("Generate.Files = %v, want [src/**/*.ts]", cfg.Generate.Files)
	}
}

func TestLoadConfig_JsoncTrailingCommas(t *testing.T) {
	cfg, err := config.LoadConfig(fixture("trailing-commas.jsonc"))
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
	_, err := config.LoadConfig(fixture("invalid.yaml"))
	if err == nil {
		t.Error("LoadConfig() expected error for invalid yaml, got nil")
	}
}

func TestLoadConfig_InvalidJson(t *testing.T) {
	_, err := config.LoadConfig(fixture("invalid.json"))
	if err == nil {
		t.Error("LoadConfig() expected error for invalid json, got nil")
	}
}

func TestLoadConfig_InvalidJsonc(t *testing.T) {
	_, err := config.LoadConfig(fixture("invalid.jsonc"))
	if err == nil {
		t.Error("LoadConfig() expected error for invalid jsonc, got nil")
	}
}

func TestLoadConfig_EmptyYaml(t *testing.T) {
	cfg, err := config.LoadConfig(fixture("empty.yaml"))
	if err != nil {
		t.Fatalf("LoadConfig() error: %v", err)
	}
	if cfg == nil {
		t.Fatal("LoadConfig() returned nil config for empty file")
	}
}

func TestLoadConfig_EmptyJson(t *testing.T) {
	cfg, err := config.LoadConfig(fixture("empty.json"))
	if err != nil {
		t.Fatalf("LoadConfig() error: %v", err)
	}
	if cfg == nil {
		t.Fatal("LoadConfig() returned nil config for empty json file")
	}
}

func TestLoadConfig_EmptyJsonc(t *testing.T) {
	cfg, err := config.LoadConfig(fixture("empty.jsonc"))
	if err != nil {
		t.Fatalf("LoadConfig() error: %v", err)
	}
	if cfg == nil {
		t.Fatal("LoadConfig() returned nil config for empty jsonc file")
	}
}

func TestLoadConfig_JsoncOnlyComments(t *testing.T) {
	cfg, err := config.LoadConfig(fixture("comments-only.jsonc"))
	if err != nil {
		t.Fatalf("LoadConfig() error: %v", err)
	}
	if cfg == nil {
		t.Fatal("LoadConfig() returned nil config for comment-only jsonc file")
	}
}

func TestLoadConfig_UnsupportedFormat(t *testing.T) {
	_, err := config.LoadConfig(fixture("unsupported.toml"))
	if err == nil {
		t.Error("LoadConfig() expected error for unsupported format, got nil")
	}
}

func TestLoadConfig_AllFields(t *testing.T) {
	cfg, err := config.LoadConfig(fixture("all-fields.yaml"))
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
