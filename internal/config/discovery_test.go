package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"bennypowers.dev/cem/internal/config"
)

func TestFindConfigFile_YamlInDotConfig(t *testing.T) {
	root := t.TempDir()
	cfgDir := filepath.Join(root, ".config")
	os.MkdirAll(cfgDir, 0o755)
	os.WriteFile(filepath.Join(cfgDir, "cem.yaml"), []byte("generate:\n  files: ['*.ts']"), 0o644)

	got := config.FindConfigFile(root)
	want := filepath.Join(root, ".config", "cem.yaml")
	if got != want {
		t.Errorf("FindConfigFile() = %q, want %q", got, want)
	}
}

func TestFindConfigFile_YmlInDotConfig(t *testing.T) {
	root := t.TempDir()
	cfgDir := filepath.Join(root, ".config")
	os.MkdirAll(cfgDir, 0o755)
	os.WriteFile(filepath.Join(cfgDir, "cem.yml"), []byte("generate:\n  files: ['*.ts']"), 0o644)

	got := config.FindConfigFile(root)
	want := filepath.Join(root, ".config", "cem.yml")
	if got != want {
		t.Errorf("FindConfigFile() = %q, want %q", got, want)
	}
}

func TestFindConfigFile_JsonInDotConfig(t *testing.T) {
	root := t.TempDir()
	cfgDir := filepath.Join(root, ".config")
	os.MkdirAll(cfgDir, 0o755)
	os.WriteFile(filepath.Join(cfgDir, "cem.json"), []byte(`{"generate":{"files":["*.ts"]}}`), 0o644)

	got := config.FindConfigFile(root)
	want := filepath.Join(root, ".config", "cem.json")
	if got != want {
		t.Errorf("FindConfigFile() = %q, want %q", got, want)
	}
}

func TestFindConfigFile_JsoncInDotConfig(t *testing.T) {
	root := t.TempDir()
	cfgDir := filepath.Join(root, ".config")
	os.MkdirAll(cfgDir, 0o755)
	os.WriteFile(filepath.Join(cfgDir, "cem.jsonc"), []byte("// comment\n{\"generate\":{}}"), 0o644)

	got := config.FindConfigFile(root)
	want := filepath.Join(root, ".config", "cem.jsonc")
	if got != want {
		t.Errorf("FindConfigFile() = %q, want %q", got, want)
	}
}

func TestFindConfigFile_DotCemYaml(t *testing.T) {
	root := t.TempDir()
	os.WriteFile(filepath.Join(root, ".cem.yaml"), []byte("generate:\n  files: ['*.ts']"), 0o644)

	got := config.FindConfigFile(root)
	want := filepath.Join(root, ".cem.yaml")
	if got != want {
		t.Errorf("FindConfigFile() = %q, want %q", got, want)
	}
}

func TestFindConfigFile_DotCemYml(t *testing.T) {
	root := t.TempDir()
	os.WriteFile(filepath.Join(root, ".cem.yml"), []byte("generate:\n  files: ['*.ts']"), 0o644)

	got := config.FindConfigFile(root)
	want := filepath.Join(root, ".cem.yml")
	if got != want {
		t.Errorf("FindConfigFile() = %q, want %q", got, want)
	}
}

func TestFindConfigFile_DotCemJson(t *testing.T) {
	root := t.TempDir()
	os.WriteFile(filepath.Join(root, ".cem.json"), []byte(`{"generate":{}}`), 0o644)

	got := config.FindConfigFile(root)
	want := filepath.Join(root, ".cem.json")
	if got != want {
		t.Errorf("FindConfigFile() = %q, want %q", got, want)
	}
}

func TestFindConfigFile_DotCemJsonc(t *testing.T) {
	root := t.TempDir()
	os.WriteFile(filepath.Join(root, ".cem.jsonc"), []byte("// comment\n{}"), 0o644)

	got := config.FindConfigFile(root)
	want := filepath.Join(root, ".cem.jsonc")
	if got != want {
		t.Errorf("FindConfigFile() = %q, want %q", got, want)
	}
}

func TestFindConfigFile_Precedence_DotConfigWinsOverDotFile(t *testing.T) {
	root := t.TempDir()
	cfgDir := filepath.Join(root, ".config")
	os.MkdirAll(cfgDir, 0o755)
	os.WriteFile(filepath.Join(cfgDir, "cem.yaml"), []byte("# primary"), 0o644)
	os.WriteFile(filepath.Join(root, ".cem.yaml"), []byte("# secondary"), 0o644)

	got := config.FindConfigFile(root)
	want := filepath.Join(root, ".config", "cem.yaml")
	if got != want {
		t.Errorf("FindConfigFile() = %q, want %q (should prefer .config/ over dotfile)", got, want)
	}
}

func TestFindConfigFile_Precedence_YamlWinsOverYml(t *testing.T) {
	root := t.TempDir()
	cfgDir := filepath.Join(root, ".config")
	os.MkdirAll(cfgDir, 0o755)
	os.WriteFile(filepath.Join(cfgDir, "cem.yaml"), []byte("# yaml"), 0o644)
	os.WriteFile(filepath.Join(cfgDir, "cem.yml"), []byte("# yml"), 0o644)

	got := config.FindConfigFile(root)
	want := filepath.Join(root, ".config", "cem.yaml")
	if got != want {
		t.Errorf("FindConfigFile() = %q, want %q (should prefer .yaml over .yml)", got, want)
	}
}

func TestFindConfigFile_NoneFound(t *testing.T) {
	root := t.TempDir()

	got := config.FindConfigFile(root)
	if got != "" {
		t.Errorf("FindConfigFile() = %q, want empty string", got)
	}
}

func TestFormatFromPath_Yaml(t *testing.T) {
	if got := config.FormatFromPath("foo/cem.yaml"); got != "yaml" {
		t.Errorf("FormatFromPath(cem.yaml) = %q, want yaml", got)
	}
}

func TestFormatFromPath_Yml(t *testing.T) {
	if got := config.FormatFromPath("foo/cem.yml"); got != "yaml" {
		t.Errorf("FormatFromPath(cem.yml) = %q, want yaml", got)
	}
}

func TestFormatFromPath_Json(t *testing.T) {
	if got := config.FormatFromPath("foo/cem.json"); got != "json" {
		t.Errorf("FormatFromPath(cem.json) = %q, want json", got)
	}
}

func TestFormatFromPath_Jsonc(t *testing.T) {
	if got := config.FormatFromPath("foo/cem.jsonc"); got != "jsonc" {
		t.Errorf("FormatFromPath(cem.jsonc) = %q, want jsonc", got)
	}
}

func TestFormatFromPath_Unknown(t *testing.T) {
	if got := config.FormatFromPath("foo/cem.toml"); got != "" {
		t.Errorf("FormatFromPath(cem.toml) = %q, want empty", got)
	}
}

func TestConfigPaths_Length(t *testing.T) {
	if len(config.ConfigPaths) != 8 {
		t.Errorf("ConfigPaths has %d entries, want 8", len(config.ConfigPaths))
	}
}

func TestConfigPaths_Order(t *testing.T) {
	expected := []string{
		".config/cem.yaml",
		".config/cem.yml",
		".config/cem.json",
		".config/cem.jsonc",
		".cem.yaml",
		".cem.yml",
		".cem.json",
		".cem.jsonc",
	}
	for i, want := range expected {
		if i >= len(config.ConfigPaths) {
			t.Fatalf("ConfigPaths too short: got %d entries, want at least %d", len(config.ConfigPaths), i+1)
		}
		if config.ConfigPaths[i] != want {
			t.Errorf("ConfigPaths[%d] = %q, want %q", i, config.ConfigPaths[i], want)
		}
	}
}
