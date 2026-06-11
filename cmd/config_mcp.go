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
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	W "bennypowers.dev/cem/internal/workspace"
	"github.com/adrg/xdg"
	"github.com/pterm/pterm"
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
			hasManifest := projectHasManifest(cmd)
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

		for i, tool := range selectedTools {
			if i > 0 {
				cmd.Println()
			}
			cmd.Print(configSnippet(tool, cemPath, mcpArgs))
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
	// Allow short names
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
	options := make([]string, len(knownTools))
	for i, t := range knownTools {
		options[i] = string(t)
	}
	selected, err := pterm.DefaultInteractiveMultiselect.
		WithOptions(options).
		WithDefaultText("Which AI tools would you like to configure?").
		Show()
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
	pterm.Info.Println("No local custom-elements manifest detected.")
	pterm.Info.Println("If your project consumes a design system, enter package specifiers to load.")
	pterm.Info.Println("Examples: npm:@rhds/elements@2.0.0, jsr:@example/components, https://cdn.example.com/pkg/")

	input, err := pterm.DefaultInteractiveTextInput.
		WithDefaultText("Additional packages (comma-separated, or press Enter to skip)").
		Show()
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

func projectHasManifest(cmd *cobra.Command) bool {
	ctx, err := W.GetWorkspaceContext(cmd)
	if err != nil {
		return false
	}
	manifestPath := ctx.CustomElementsManifestPath()
	if manifestPath == "" {
		return false
	}
	_, err = os.Stat(manifestPath)
	return err == nil
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

func configSnippet(tool aiTool, cemPath string, args []string) string {
	switch tool {
	case toolClaudeCode:
		return claudeCodeSnippet(cemPath, args)
	case toolClaudeDesktop:
		return claudeDesktopSnippet(cemPath, args)
	case toolCursor:
		return cursorSnippet(cemPath, args)
	case toolVSCode:
		return vscodeSnippet(cemPath, args)
	case toolZed:
		return zedSnippet(cemPath, args)
	case toolOther:
		return genericSnippet(cemPath, args)
	default:
		return ""
	}
}

func claudeCodeSnippet(cemPath string, args []string) string {
	var b strings.Builder
	b.WriteString("Claude Code\n\n")
	b.WriteString("  Run this command:\n\n")
	fmt.Fprintf(&b, "    claude mcp add cem -- %s", cemPath)
	for _, a := range args {
		fmt.Fprintf(&b, " %s", a)
	}
	b.WriteString("\n")
	return b.String()
}

type mcpServerEntry struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

type mcpServersWrapper struct {
	MCPServers map[string]mcpServerEntry `json:"mcpServers"`
}

func claudeDesktopSnippet(cemPath string, args []string) string {
	snippet := mcpServersJSON(cemPath, args)
	var b strings.Builder
	b.WriteString("Claude Desktop\n\n")
	for _, p := range claudeDesktopConfigPaths() {
		fmt.Fprintf(&b, "  %s:\n", p.label)
		fmt.Fprintf(&b, "    %s\n\n", p.path)
	}
	b.WriteString(indentJSON(snippet, "    "))
	b.WriteString("\n")
	return b.String()
}

func cursorSnippet(cemPath string, args []string) string {
	snippet := mcpServersJSON(cemPath, args)
	return fmt.Sprintf("Cursor\n\n  Add to .cursor/mcp.json in your project root:\n\n%s\n", indentJSON(snippet, "    "))
}

type vscodeServerEntry struct {
	Type    string   `json:"type"`
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

type vscodeServersWrapper struct {
	Servers map[string]vscodeServerEntry `json:"servers"`
}

func vscodeSnippet(cemPath string, args []string) string {
	config := vscodeServersWrapper{
		Servers: map[string]vscodeServerEntry{
			"cem": {Type: "stdio", Command: cemPath, Args: args},
		},
	}
	snippet, _ := json.MarshalIndent(config, "", "  ")
	var b strings.Builder
	b.WriteString("VS Code (Copilot)\n\n")
	fmt.Fprintf(&b, "  Run this command:\n\n    code --add-mcp '%s'\n\n", string(snippet))
	b.WriteString("  Or add to .vscode/mcp.json in your project root:\n\n")
	b.WriteString(indentJSON(string(snippet), "    "))
	b.WriteString("\n")
	return b.String()
}

type zedServerEntry struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

type zedServersWrapper struct {
	ContextServers map[string]zedServerEntry `json:"context_servers"`
}

func zedSnippet(cemPath string, args []string) string {
	config := zedServersWrapper{
		ContextServers: map[string]zedServerEntry{
			"cem": {Command: cemPath, Args: args},
		},
	}
	snippet, _ := json.MarshalIndent(config, "", "  ")
	return fmt.Sprintf("Zed\n\n  Add to %s:\n\n%s\n", zedConfigPath(), indentJSON(string(snippet), "    "))
}

func genericSnippet(cemPath string, args []string) string {
	config := mcpServerEntry{Command: cemPath, Args: args}
	snippet, _ := json.MarshalIndent(config, "", "  ")
	return fmt.Sprintf("Other (generic stdio)\n\n  Configure your MCP client with:\n\n%s\n", indentJSON(string(snippet), "    "))
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

func claudeDesktopConfigPaths() []configPathInfo {
	home, _ := os.UserHomeDir()
	switch runtime.GOOS {
	case "darwin":
		return []configPathInfo{
			{"macOS", filepath.Join(home, "Library", "Application Support", "Claude", "claude_desktop_config.json")},
		}
	case "windows":
		return []configPathInfo{
			{"Windows", filepath.Join(os.Getenv("APPDATA"), "Claude", "claude_desktop_config.json")},
		}
	default:
		return []configPathInfo{
			{"Linux (native)", filepath.Join(xdg.ConfigHome, "Claude", "claude_desktop_config.json")},
			{"Linux (Flatpak)", filepath.Join(home, ".var", "app", "com.anthropic.Claude", "config", "Claude", "claude_desktop_config.json")},
		}
	}
}

func zedConfigPath() string {
	return filepath.Join(xdg.ConfigHome, "zed", "settings.json")
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
