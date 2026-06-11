//go:build e2e

package cmd_test

import (
	"path/filepath"
	"strings"
	"testing"
)

func configMCPGolden(name string) string {
	return filepath.Join("testdata", "goldens", "config-mcp", name)
}

func TestConfigMCP_ClaudeCode(t *testing.T) {
	root := projectRoot(t)
	stdout, _ := runCemCommand(t, root, "config", "mcp", "--tool", "Claude Code")
	compareGolden(t, configMCPGolden("claude-code.txt"), stdout, false)
}

func TestConfigMCP_ClaudeDesktop(t *testing.T) {
	root := projectRoot(t)
	stdout, _ := runCemCommand(t, root, "config", "mcp", "--tool", "Claude Desktop")
	compareGolden(t, configMCPGolden("claude-desktop.txt"), stdout, false)
}

func TestConfigMCP_Cursor(t *testing.T) {
	root := projectRoot(t)
	stdout, _ := runCemCommand(t, root, "config", "mcp", "--tool", "Cursor")
	compareGolden(t, configMCPGolden("cursor.txt"), stdout, false)
}

func TestConfigMCP_VSCode(t *testing.T) {
	root := projectRoot(t)
	stdout, _ := runCemCommand(t, root, "config", "mcp", "--tool", "VS Code (Copilot)")
	compareGolden(t, configMCPGolden("vscode.txt"), stdout, false)
}

func TestConfigMCP_Zed(t *testing.T) {
	root := projectRoot(t)
	stdout, _ := runCemCommand(t, root, "config", "mcp", "--tool", "Zed")
	compareGolden(t, configMCPGolden("zed.txt"), stdout, false)
}

func TestConfigMCP_Other(t *testing.T) {
	root := projectRoot(t)
	stdout, _ := runCemCommand(t, root, "config", "mcp", "--tool", "Other")
	compareGolden(t, configMCPGolden("other.txt"), stdout, false)
}

func TestConfigMCP_MultipleTools(t *testing.T) {
	root := projectRoot(t)
	stdout, _ := runCemCommand(t, root, "config", "mcp", "--tool", "Claude Code", "--tool", "Cursor")
	compareGolden(t, configMCPGolden("multiple.txt"), stdout, false)
}

func TestConfigMCP_WithAdditionalPackages(t *testing.T) {
	root := projectRoot(t)
	stdout, _ := runCemCommand(t, root, "config", "mcp", "--tool", "Cursor", "-a", "npm:@rhds/elements@2.0.0")
	compareGolden(t, configMCPGolden("with-packages.txt"), stdout, false)
}

func TestConfigMCP_ShortNames(t *testing.T) {
	root := projectRoot(t)
	for _, tc := range []struct {
		short    string
		expected string
	}{
		{"claude-code", "Claude Code"},
		{"desktop", "Claude Desktop"},
		{"cursor", "Cursor"},
		{"vscode", "VS Code (Copilot)"},
		{"copilot", "VS Code (Copilot)"},
		{"zed", "Zed"},
		{"generic", "Other (generic stdio)"},
	} {
		t.Run(tc.short, func(t *testing.T) {
			stdout, _ := runCemCommand(t, root, "config", "mcp", "--tool", tc.short)
			if !strings.HasPrefix(stdout, tc.expected) {
				t.Errorf("--tool %q: expected output starting with %q, got %q", tc.short, tc.expected, stdout[:min(len(stdout), 40)])
			}
		})
	}
}

func TestConfigMCP_NoTTYNoFlags(t *testing.T) {
	root := projectRoot(t)
	_, stderr := runCemCommand(t, root, "config", "mcp")
	if !strings.Contains(stderr, "no TTY detected") {
		t.Errorf("expected TTY error, got: %s", stderr)
	}
}

func TestConfigMCP_InvalidTool(t *testing.T) {
	root := projectRoot(t)
	_, stderr := runCemCommand(t, root, "config", "mcp", "--tool", "invalid-tool")
	if !strings.Contains(stderr, "unknown tool") {
		t.Errorf("expected unknown tool error, got: %s", stderr)
	}
}
