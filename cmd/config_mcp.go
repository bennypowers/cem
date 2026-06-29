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
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"bennypowers.dev/cem/internal/jsoncmerge"
	"bennypowers.dev/cem/internal/logging"
	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/internal/tui"
	W "bennypowers.dev/cem/internal/workspace"
	"charm.land/huh/v2"
	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

type aiTool string

const (
	toolClaudeCode    aiTool = "Claude Code"
	toolClaudeDesktop aiTool = "Claude Desktop"
	toolCursor        aiTool = "Cursor"
	toolVSCode        aiTool = "VS Code (Copilot)"
	toolZed           aiTool = "Zed"
	toolOther         aiTool = "Other (generic stdio)"
)

var knownTools = []aiTool{
	toolClaudeCode,
	toolClaudeDesktop,
	toolCursor,
	toolVSCode,
	toolZed,
	toolOther,
}

type mcpAction struct {
	tool    aiTool
	preview string
	apply   func() error
}

var configMCPCmd = &cobra.Command{
	Use:   "mcp",
	Short: "Generate MCP client configuration",
	Long: `Generate configuration snippets for AI tools (Claude Code, Claude Desktop, Cursor, etc.).

In interactive mode (TTY), presents a menu to select tools and configure options.
In non-interactive mode, use --tool to specify tools directly.`,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SetOut(os.Stdout)
		toolFlags, _ := cmd.Flags().GetStringSlice("tool")
		additionalPackages, _ := cmd.Flags().GetStringSlice("additional-packages")
		interactive := term.IsTerminal(int(os.Stdin.Fd()))
		osFS := platform.NewOSFileSystem()

		var selectedTools []aiTool

		if len(toolFlags) > 0 {
			for _, name := range toolFlags {
				t, err := parseToolName(name)
				if err != nil {
					return err
				}
				selectedTools = append(selectedTools, t)
			}
		} else if interactive {
			tools, err := promptToolSelection()
			if err != nil {
				return err
			}
			selectedTools = tools
		} else {
			return fmt.Errorf("no TTY detected; use --tool to specify tools, e.g.:\n  cem config mcp --tool \"Claude Desktop\" --tool \"Cursor\"")
		}

		if len(selectedTools) == 0 {
			return fmt.Errorf("no tools selected")
		}

		if interactive && len(additionalPackages) == 0 {
			hasManifest := projectHasManifest(cmd, osFS)
			if !hasManifest {
				pkgs, err := promptAdditionalPackages()
				if err != nil {
					return err
				}
				additionalPackages = pkgs
			}
		}

		cemPath := resolveCemPath(cmd)
		mcpArgs := buildMCPArgs(additionalPackages)

		var homeDir string
		resolveHomeDir := func() (string, error) {
			if homeDir == "" {
				var err error
				homeDir, err = os.UserHomeDir()
				if err != nil {
					return "", fmt.Errorf("unable to determine home directory: %w", err)
				}
			}
			return homeDir, nil
		}

		actions, err := buildActions(osFS, resolveHomeDir, selectedTools, cemPath, mcpArgs, interactive)
		if err != nil {
			return err
		}

		for i, action := range actions {
			if i > 0 {
				cmd.Println()
			}
			cmd.Print(action.preview)

			if interactive && action.apply != nil {
				confirmed, err := tui.Confirm("Apply?", true)
				if err != nil {
					return err
				}
				if confirmed {
					if err := action.apply(); err != nil {
						logging.Error("%v", err)
					} else {
						logging.Success("Done")
					}
				}
			}
		}

		return nil
	},
}

func init() {
	configMCPCmd.Flags().StringSlice("tool", nil, `AI tool(s) to configure (repeatable). Valid: "Claude Code", "Claude Desktop", "Cursor", "VS Code (Copilot)", "Zed", "Other"`)
	configMCPCmd.Flags().StringSliceP("additional-packages", "a", nil, "Additional packages to include in MCP args (npm:, jsr:, or URL specifiers)")
	configCmd.AddCommand(configMCPCmd)
}

func parseToolName(name string) (aiTool, error) {
	lower := strings.ToLower(strings.TrimSpace(name))
	for _, t := range knownTools {
		if strings.ToLower(string(t)) == lower {
			return t, nil
		}
	}
	switch lower {
	case "claude-code", "claudecode":
		return toolClaudeCode, nil
	case "claude-desktop", "claudedesktop", "desktop":
		return toolClaudeDesktop, nil
	case "cursor":
		return toolCursor, nil
	case "vscode", "vs-code", "copilot":
		return toolVSCode, nil
	case "zed":
		return toolZed, nil
	case "other", "generic", "stdio":
		return toolOther, nil
	}
	validNames := make([]string, len(knownTools))
	for i, t := range knownTools {
		validNames[i] = string(t)
	}
	return "", fmt.Errorf("unknown tool %q; valid tools: %s", name, strings.Join(validNames, ", "))
}

func promptToolSelection() ([]aiTool, error) {
	options := make([]huh.Option[string], len(knownTools))
	for i, t := range knownTools {
		options[i] = huh.NewOption(string(t), string(t))
	}
	selected, err := tui.MultiSelect("Which AI tools would you like to configure?", options)
	if err != nil {
		return nil, fmt.Errorf("tool selection cancelled: %w", err)
	}
	tools := make([]aiTool, len(selected))
	for i, name := range selected {
		tools[i] = aiTool(name)
	}
	return tools, nil
}

func promptAdditionalPackages() ([]string, error) {
	logging.Info("No local custom-elements manifest detected.")
	logging.Info("If your project consumes a design system, enter package specifiers to load.")
	logging.Info("Examples: npm:@rhds/elements@2.0.0, jsr:@example/components, https://cdn.example.com/pkg/")

	input, err := tui.TextInput(
		"Additional packages (comma-separated, or press Enter to skip)",
		"npm:@example/elements@1.0.0",
	)
	if err != nil {
		return nil, err
	}
	input = strings.TrimSpace(input)
	if input == "" {
		return nil, nil
	}
	var packages []string
	for pkg := range strings.SplitSeq(input, ",") {
		pkg = strings.TrimSpace(pkg)
		if pkg != "" {
			packages = append(packages, pkg)
		}
	}
	return packages, nil
}

func projectHasManifest(cmd *cobra.Command, fsys platform.FileSystem) bool {
	ctx, err := W.GetWorkspaceContext(cmd)
	if err != nil {
		return false
	}
	manifestPath := ctx.CustomElementsManifestPath()
	if manifestPath == "" {
		return false
	}
	return fsys.Exists(manifestPath)
}

func resolveCemPath(cmd *cobra.Command) string {
	_, err := exec.LookPath("cem")
	if err != nil {
		absPath, exeErr := os.Executable()
		if exeErr != nil {
			return "cem"
		}
		cmd.PrintErrln("Warning: cem is not in your PATH. Using absolute path in config snippets.")
		cmd.PrintErrf("  To fix, add the directory containing cem to your PATH, or use:\n    %s\n\n", absPath)
		return absPath
	}
	return "cem"
}

func buildMCPArgs(additionalPackages []string) []string {
	args := []string{"mcp"}
	for _, pkg := range additionalPackages {
		args = append(args, "-a", pkg)
	}
	return args
}

func buildActions(fsys platform.FileSystem, resolveHomeDir func() (string, error), tools []aiTool, cemPath string, args []string, interactive bool) ([]mcpAction, error) {
	var actions []mcpAction
	for _, tool := range tools {
		action, err := buildAction(fsys, resolveHomeDir, tool, cemPath, args, interactive)
		if err != nil {
			return nil, err
		}
		actions = append(actions, action)
	}
	return actions, nil
}

func buildAction(fsys platform.FileSystem, resolveHomeDir func() (string, error), tool aiTool, cemPath string, args []string, interactive bool) (mcpAction, error) {
	switch tool {
	case toolClaudeCode:
		return claudeCodeAction(cemPath, args), nil
	case toolClaudeDesktop:
		homeDir, err := resolveHomeDir()
		if err != nil {
			return mcpAction{}, err
		}
		return claudeDesktopAction(fsys, homeDir, cemPath, args), nil
	case toolCursor:
		homeDir, err := resolveHomeDir()
		if err != nil {
			return mcpAction{}, err
		}
		return cursorAction(fsys, homeDir, cemPath, args, interactive), nil
	case toolVSCode:
		homeDir, err := resolveHomeDir()
		if err != nil {
			return mcpAction{}, err
		}
		return vscodeAction(fsys, homeDir, cemPath, args, interactive), nil
	case toolZed:
		return zedAction(fsys, interactive), nil
	default:
		return mcpAction{tool: tool, preview: genericSnippet(cemPath, args)}, nil
	}
}

func claudeCodeAction(cemPath string, args []string) mcpAction {
	cmdArgs := append([]string{"mcp", "add", "cem", "--"}, cemPath)
	cmdArgs = append(cmdArgs, args...)
	preview := fmt.Sprintf(`Claude Code

  Run: claude %s
`, strings.Join(cmdArgs, " "))
	return mcpAction{
		tool:    toolClaudeCode,
		preview: preview,
		apply: func() error {
			cmd := exec.Command("claude", cmdArgs...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				return fmt.Errorf("claude %s: %w", strings.Join(cmdArgs, " "), err)
			}
			return nil
		},
	}
}

func claudeDesktopAction(fsys platform.FileSystem, homeDir string, cemPath string, args []string) mcpAction {
	paths := claudeDesktopConfigPaths(homeDir)
	snippet := mcpServersJSON(cemPath, args)
	var b strings.Builder
	b.WriteString("Claude Desktop\n\n")
	for _, p := range paths {
		fmt.Fprintf(&b, "  %s:\n    %s\n\n", p.label, p.path)
	}
	b.WriteString(indentJSON(snippet, "    "))
	b.WriteString("  Restart Claude Desktop after applying.\n\n")

	var targetPath string
	for _, p := range paths {
		if fsys.Exists(filepath.Dir(p.path)) {
			targetPath = p.path
			break
		}
	}

	var applyFn func() error
	if targetPath != "" {
		applyFn = func() error {
			return mergeJSONConfig(fsys, targetPath, "mcpServers", "cem", mcpServerEntry{
				Command: cemPath,
				Args:    args,
			})
		}
	}
	return mcpAction{tool: toolClaudeDesktop, preview: b.String(), apply: applyFn}
}

func cursorAction(fsys platform.FileSystem, homeDir string, cemPath string, args []string, interactive bool) mcpAction {
	snippet := mcpServersJSON(cemPath, args)
	projectPath := filepath.Join(".cursor", "mcp.json")
	globalPath := filepath.Join(homeDir, ".cursor", "mcp.json")
	targetPath := chooseScope(fsys, interactive, ".cursor", projectPath, globalPath)
	preview := fmt.Sprintf("Cursor\n\n  Write to %s:\n\n%s\n", targetPath, indentJSON(snippet, "    "))
	return mcpAction{
		tool:    toolCursor,
		preview: preview,
		apply: func() error {
			return mergeJSONConfig(fsys, targetPath, "mcpServers", "cem", mcpServerEntry{
				Command: cemPath,
				Args:    args,
			})
		},
	}
}

func vscodeAction(fsys platform.FileSystem, homeDir string, cemPath string, args []string, interactive bool) mcpAction {
	entry := vscodeServerEntry{Type: "stdio", Command: cemPath, Args: args}
	config := vscodeServersWrapper{
		Servers: map[string]vscodeServerEntry{"cem": entry},
	}
	snippet, _ := json.MarshalIndent(config, "", "  ")
	projectPath := filepath.Join(".vscode", "mcp.json")
	globalPath := vscodeGlobalConfigPath(homeDir)
	targetPath := chooseScope(fsys, interactive, ".vscode", projectPath, globalPath)
	preview := fmt.Sprintf(`VS Code (Copilot)

  Run: code --add-mcp '%s'

  Or write to %s:

%s
`, string(snippet), targetPath, indentJSON(string(snippet), "    "))
	return mcpAction{
		tool:    toolVSCode,
		preview: preview,
		apply: func() error {
			cmd := exec.Command("code", "--add-mcp", string(snippet))
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				return fmt.Errorf("code --add-mcp: %w", err)
			}
			return nil
		},
	}
}

func zedAction(fsys platform.FileSystem, interactive bool) mcpAction {
	projectPath := filepath.Join(".zed", "settings.json")
	globalPath := filepath.Join(xdg.ConfigHome, "zed", "settings.json")
	targetPath := chooseScope(fsys, interactive, ".zed", projectPath, globalPath)
	preview := fmt.Sprintf(`Zed
  The cem extension provides both LSP and MCP support.

  Add to %s:

    "auto_install_extensions": {
      "cem": true
    }

  Restart Zed after applying.
`, targetPath)
	return mcpAction{
		tool:    toolZed,
		preview: preview,
		apply: func() error {
			return mergeJSONConfig(fsys, targetPath, "auto_install_extensions", "cem", true)
		},
	}
}

func vscodeGlobalConfigPath(homeDir string) string {
	switch runtime.GOOS {
	case "darwin":
		return filepath.Join(homeDir, "Library", "Application Support", "Code", "User", "mcp.json")
	case "windows":
		return filepath.Join(os.Getenv("APPDATA"), "Code", "User", "mcp.json")
	default:
		return filepath.Join(xdg.ConfigHome, "Code", "User", "mcp.json")
	}
}

// chooseScope prompts for project vs global scope in interactive mode.
// If the project dir exists, defaults to project. Otherwise defaults to global.
// Non-interactive mode always uses project path.
func chooseScope(fsys platform.FileSystem, interactive bool, projectDir, projectPath, globalPath string) string {
	if !interactive {
		return projectPath
	}
	projectExists := fsys.Exists(projectDir)
	projectLabel := "Project (" + projectPath + ")"
	globalLabel := "Global (" + globalPath + ")"
	var options []string
	if projectExists {
		options = []string{projectLabel, globalLabel}
	} else {
		options = []string{globalLabel, projectLabel}
	}
	result, err := tui.Select("Configure for this project or globally?", tui.StringOptions(options...))
	if err != nil {
		if projectExists {
			return projectPath
		}
		return globalPath
	}
	if strings.HasPrefix(result, "Global") {
		return globalPath
	}
	return projectPath
}

func genericSnippet(cemPath string, args []string) string {
	config := mcpServerEntry{Command: cemPath, Args: args}
	snippet, _ := json.MarshalIndent(config, "", "  ")
	return fmt.Sprintf(`Other (generic stdio)

  Configure your MCP client with:

%s
`, indentJSON(string(snippet), "    "))
}

// mergeJSONConfig inserts topKey.subKey = value into a JSON/JSONC file.
// Preserves all existing comments, formatting, and content via byte-level splicing.
// Uses hujson only for JSONC validation.
// Creates the file and parent directories if they don't exist.
// Backs up existing files before writing.
func mergeJSONConfig(fsys platform.FileSystem, path, topKey, subKey string, value any) error {
	data, err := fsys.ReadFile(path)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return fmt.Errorf("failed to read %s: %w", path, err)
	}
	if err == nil {
		if err := backupFile(fsys, path, data); err != nil {
			return err
		}
	} else {
		data = []byte("{}\n")
	}

	result, err := jsoncmerge.Merge(data, topKey, subKey, value)
	if err != nil {
		return fmt.Errorf("failed to merge into %s: %w", path, err)
	}

	dir := filepath.Dir(path)
	if err := fsys.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}
	tmp, err := fsys.CreateTemp(dir, ".cem-config-*.tmp")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	tmpPath := tmp.Name()
	if _, err := tmp.Write(result); err != nil {
		_ = tmp.Close()
		_ = fsys.Remove(tmpPath)
		return fmt.Errorf("failed to write config: %w", err)
	}
	if err := tmp.Close(); err != nil {
		_ = fsys.Remove(tmpPath)
		return fmt.Errorf("failed to close temp file: %w", err)
	}
	if err := fsys.Rename(tmpPath, path); err != nil {
		_ = fsys.Remove(tmpPath)
		return fmt.Errorf("failed to rename temp file: %w", err)
	}
	return nil
}

func backupFile(fsys platform.FileSystem, path string, data []byte) error {
	backup, err := fsys.CreateTemp(filepath.Dir(path), filepath.Base(path)+".backup-*")
	if err != nil {
		return fmt.Errorf("failed to create backup: %w", err)
	}
	if _, err := backup.Write(data); err != nil {
		_ = backup.Close()
		return fmt.Errorf("failed to write backup: %w", err)
	}
	if err := backup.Close(); err != nil {
		return fmt.Errorf("failed to close backup file: %w", err)
	}
	logging.Info("Backed up %s to %s", path, backup.Name())
	return nil
}

type mcpServerEntry struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

type mcpServersWrapper struct {
	MCPServers map[string]mcpServerEntry `json:"mcpServers"`
}

type vscodeServerEntry struct {
	Type    string   `json:"type"`
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

type vscodeServersWrapper struct {
	Servers map[string]vscodeServerEntry `json:"servers"`
}

func mcpServersJSON(cemPath string, args []string) string {
	config := mcpServersWrapper{
		MCPServers: map[string]mcpServerEntry{
			"cem": {Command: cemPath, Args: args},
		},
	}
	data, _ := json.MarshalIndent(config, "", "  ")
	return string(data)
}

type configPathInfo struct {
	label string
	path  string
}

func claudeDesktopConfigPaths(homeDir string) []configPathInfo {
	switch runtime.GOOS {
	case "darwin":
		return []configPathInfo{
			{"macOS", filepath.Join(homeDir, "Library", "Application Support", "Claude", "claude_desktop_config.json")},
		}
	case "windows":
		return []configPathInfo{
			{"Windows", filepath.Join(os.Getenv("APPDATA"), "Claude", "claude_desktop_config.json")},
		}
	default:
		return []configPathInfo{
			{"Linux (native)", filepath.Join(xdg.ConfigHome, "Claude", "claude_desktop_config.json")},
			{"Linux (Flatpak)", filepath.Join(homeDir, ".var", "app", "com.anthropic.Claude", "config", "Claude", "claude_desktop_config.json")},
		}
	}
}

func indentJSON(jsonStr string, prefix string) string {
	var b strings.Builder
	for line := range strings.SplitSeq(jsonStr, "\n") {
		b.WriteString(prefix)
		b.WriteString(line)
		b.WriteString("\n")
	}
	return b.String()
}
