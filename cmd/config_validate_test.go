//go:build e2e

package cmd_test

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"
)

func projectRoot(t *testing.T) string {
	t.Helper()
	root, err := filepath.Abs("..")
	if err != nil {
		t.Fatalf("failed to resolve project root: %v", err)
	}
	return root
}

func TestConfigValidate_ValidConfig(t *testing.T) {
	root := projectRoot(t)
	stdout, stderr := runCemCommand(t, root, "config", "validate", "-p", "examples/kitchen-sink")
	if strings.Contains(stderr, "config validation failed") {
		t.Errorf("valid config should not fail: stderr=%s", stderr)
	}
	if !strings.Contains(stdout, "valid") {
		t.Errorf("expected 'valid' in stdout, got stdout=%q stderr=%q", stdout, stderr)
	}
}

func TestConfigValidate_JSONFormat(t *testing.T) {
	root := projectRoot(t)
	stdout, _ := runCemCommand(t, root, "config", "validate", "-p", "examples/kitchen-sink", "--format", "json")
	var result struct {
		ConfigFile string `json:"configFile"`
		Valid      bool   `json:"valid"`
		Errors     []any  `json:"errors"`
		Warnings   []any  `json:"warnings"`
	}
	if err := json.Unmarshal([]byte(stdout), &result); err != nil {
		t.Fatalf("JSON output should be parseable: %v\nstdout: %s", err, stdout)
	}
	if !result.Valid {
		t.Errorf("kitchen-sink config should be valid")
	}
	if result.ConfigFile == "" {
		t.Error("configFile should be set")
	}
}

func TestConfigValidate_NoConfig(t *testing.T) {
	_, stderr := runCemCommand(t, t.TempDir(), "config", "validate")
	if !strings.Contains(stderr, "no config file found") {
		t.Errorf("expected 'no config file found' error, got: %s", stderr)
	}
	if strings.Contains(stderr, "Usage:") {
		t.Error("SilenceUsage should prevent usage text on error")
	}
}

func TestConfigValidate_AllExamples(t *testing.T) {
	root := projectRoot(t)
	examples, _ := filepath.Glob(filepath.Join(root, "examples", "*"))
	for _, ex := range examples {
		name := filepath.Base(ex)
		t.Run(name, func(t *testing.T) {
			stdout, stderr := runCemCommand(t, root, "config", "validate", "-p", ex)
			if strings.Contains(stderr, "config validation failed") {
				t.Errorf("example %s should validate: stderr=%s stdout=%s", name, stderr, stdout)
			}
		})
	}
}
