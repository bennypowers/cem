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

	"atomicgo.dev/keyboard/keys"

	W "bennypowers.dev/cem/internal/workspace"
	"github.com/adrg/xdg"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/tailscale/hujson" // JSONC validation only
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

		actions := buildActions(selectedTools, cemPath, mcpArgs)

		for i, action := range actions {
			if i > 0 {
				cmd.Println()
			}
			cmd.Print(action.preview)

			if interactive && action.apply != nil {
				confirmed, err := pterm.DefaultInteractiveConfirm.
					WithDefaultValue(true).
					WithDefaultText("Apply?").
					Show()
				if err != nil {
					return err
				}
				if confirmed {
					if err := action.apply(); err != nil {
						pterm.Error.Printfln("%v", err)
					} else {
						pterm.Success.Println("Done")
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
	options := make([]string, len(knownTools))
	for i, t := range knownTools {
		options[i] = string(t)
	}
	selected, err := pterm.DefaultInteractiveMultiselect.
		WithOptions(options).
		WithDefaultText("Which AI tools would you like to configure?\n  space: select, enter: confirm, left/right: none/all").
		WithKeySelect(keys.Space).
		WithKeyConfirm(keys.Enter).
		WithFilter(false).
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

func buildActions(tools []aiTool, cemPath string, args []string) []mcpAction {
	var actions []mcpAction
	for _, tool := range tools {
		actions = append(actions, buildAction(tool, cemPath, args))
	}
	return actions
}

func buildAction(tool aiTool, cemPath string, args []string) mcpAction {
	switch tool {
	case toolClaudeCode:
		return claudeCodeAction(cemPath, args)
	case toolClaudeDesktop:
		return claudeDesktopAction(cemPath, args)
	case toolCursor:
		return cursorAction(cemPath, args)
	case toolVSCode:
		return vscodeAction(cemPath, args)
	case toolZed:
		return zedAction(cemPath, args)
	default:
		return mcpAction{tool: tool, preview: genericSnippet(cemPath, args)}
	}
}

func claudeCodeAction(cemPath string, args []string) mcpAction {
	cmdArgs := append([]string{"mcp", "add", "cem", "--"}, cemPath)
	cmdArgs = append(cmdArgs, args...)
	var b strings.Builder
	b.WriteString("Claude Code\n\n")
	fmt.Fprintf(&b, "  Run: claude %s\n\n", strings.Join(cmdArgs, " "))
	return mcpAction{
		tool:    toolClaudeCode,
		preview: b.String(),
		apply: func() error {
			cmd := exec.Command("claude", cmdArgs...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			return cmd.Run()
		},
	}
}

func claudeDesktopAction(cemPath string, args []string) mcpAction {
	paths := claudeDesktopConfigPaths()
	snippet := mcpServersJSON(cemPath, args)
	var b strings.Builder
	b.WriteString("Claude Desktop\n\n")
	for _, p := range paths {
		fmt.Fprintf(&b, "  %s:\n", p.label)
		fmt.Fprintf(&b, "    %s\n\n", p.path)
	}
	b.WriteString(indentJSON(snippet, "    "))
	b.WriteString("\n")

	var targetPath string
	for _, p := range paths {
		if _, err := os.Stat(filepath.Dir(p.path)); err == nil {
			targetPath = p.path
			break
		}
	}

	var applyFn func() error
	if targetPath != "" {
		applyFn = func() error {
			return mergeJSONConfig(targetPath, "mcpServers", "cem", mcpServerEntry{
				Command: cemPath,
				Args:    args,
			})
		}
	}
	return mcpAction{tool: toolClaudeDesktop, preview: b.String(), apply: applyFn}
}

func cursorAction(cemPath string, args []string) mcpAction {
	snippet := mcpServersJSON(cemPath, args)
	targetPath := filepath.Join(".cursor", "mcp.json")
	var b strings.Builder
	fmt.Fprintf(&b, "Cursor\n\n  Write to %s:\n\n%s\n", targetPath, indentJSON(snippet, "    "))
	return mcpAction{
		tool:    toolCursor,
		preview: b.String(),
		apply: func() error {
			return mergeJSONConfig(targetPath, "mcpServers", "cem", mcpServerEntry{
				Command: cemPath,
				Args:    args,
			})
		},
	}
}

func vscodeAction(cemPath string, args []string) mcpAction {
	entry := vscodeServerEntry{Type: "stdio", Command: cemPath, Args: args}
	config := vscodeServersWrapper{
		Servers: map[string]vscodeServerEntry{"cem": entry},
	}
	snippet, _ := json.MarshalIndent(config, "", "  ")
	cmdArgs := fmt.Sprintf("--add-mcp '%s'", string(snippet))
	targetPath := filepath.Join(".vscode", "mcp.json")
	var b strings.Builder
	fmt.Fprintf(&b, "VS Code (Copilot)\n\n  Run: code %s\n\n", cmdArgs)
	fmt.Fprintf(&b, "  Or write to %s:\n\n%s\n", targetPath, indentJSON(string(snippet), "    "))
	return mcpAction{
		tool:    toolVSCode,
		preview: b.String(),
		apply: func() error {
			cmd := exec.Command("code", "--add-mcp", string(snippet))
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			return cmd.Run()
		},
	}
}

func zedAction(cemPath string, args []string) mcpAction {
	config := zedServersWrapper{
		ContextServers: map[string]zedServerEntry{
			"cem": {Command: cemPath, Args: args},
		},
	}
	snippet, _ := json.MarshalIndent(config, "", "  ")
	targetPath := zedConfigPath()
	var b strings.Builder
	fmt.Fprintf(&b, "Zed\n\n  Merge into %s:\n\n%s\n", targetPath, indentJSON(string(snippet), "    "))
	return mcpAction{
		tool:    toolZed,
		preview: b.String(),
		apply: func() error {
			return mergeJSONConfig(targetPath, "context_servers", "cem", zedServerEntry{
				Command: cemPath,
				Args:    args,
			})
		},
	}
}

func genericSnippet(cemPath string, args []string) string {
	config := mcpServerEntry{Command: cemPath, Args: args}
	snippet, _ := json.MarshalIndent(config, "", "  ")
	return fmt.Sprintf("Other (generic stdio)\n\n  Configure your MCP client with:\n\n%s\n", indentJSON(string(snippet), "    "))
}

// mergeJSONConfig inserts topKey.subKey = value into a JSON/JSONC file.
// Preserves all existing comments, formatting, and content via byte-level splicing.
// Uses hujson only for JSONC validation.
// Creates the file and parent directories if they don't exist.
// Backs up existing files before writing.
func mergeJSONConfig(path, topKey, subKey string, value any) error {
	data, err := os.ReadFile(path)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to read %s: %w", path, err)
	}
	if err == nil {
		if err := backupFile(path, data); err != nil {
			return err
		}
	} else {
		data = []byte("{}\n")
	}

	result, err := mergeJSONCBytes(data, topKey, subKey, value)
	if err != nil {
		return fmt.Errorf("failed to merge into %s: %w", path, err)
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	return os.WriteFile(path, result, 0o644)
}

// mergeJSONCBytes inserts or replaces topKey.subKey in raw JSONC bytes,
// preserving comments and formatting via byte-level splicing.
// Uses hujson AST for key location to avoid matching inside comments/strings.
func mergeJSONCBytes(data []byte, topKey, subKey string, value any) ([]byte, error) {
	v, err := hujson.Parse(data)
	if err != nil {
		return nil, fmt.Errorf("invalid JSONC: %w", err)
	}
	v.UpdateOffsets()

	topNode := v.Find("/" + topKey)

	if topNode == nil {
		rootObj, ok := v.Value.(*hujson.Object)
		if !ok {
			return nil, fmt.Errorf("root is not an object")
		}
		closeBrace := rootCloseOffset(&v, rootObj)
		topKeyStr := fmt.Sprintf("%q", topKey)
		innerJSON, _ := json.MarshalIndent(map[string]any{subKey: value}, "  ", "  ")
		newBlock := fmt.Sprintf("  %s: %s", topKeyStr, string(innerJSON))
		return insertBeforeClose(data, closeBrace, newBlock, "")
	}

	subNode := v.Find("/" + topKey + "/" + subKey)

	if subNode != nil {
		valueJSON, _ := json.MarshalIndent(value, "    ", "  ")
		return spliceBytes(data, subNode.StartOffset, subNode.EndOffset, valueJSON), nil
	}

	topObj, ok := topNode.Value.(*hujson.Object)
	if !ok {
		return nil, fmt.Errorf("%q is not an object", topKey)
	}
	closeBrace := objCloseOffset(topNode, topObj)
	valueJSON, _ := json.MarshalIndent(value, "    ", "  ")
	newMember := fmt.Sprintf("    %q: %s", subKey, string(valueJSON))
	return insertBeforeClose(data, closeBrace, newMember, "  ")
}

func rootCloseOffset(v *hujson.Value, obj *hujson.Object) int {
	return v.EndOffset - len("}") - len(obj.AfterExtra) - len(v.AfterExtra)
}

func objCloseOffset(v *hujson.Value, obj *hujson.Object) int {
	return v.EndOffset - len("}") - len(obj.AfterExtra)
}

// insertBeforeClose inserts content before a closing brace, handling commas.
// Skips trailing whitespace, line comments (//), and block comments (/* */).
func insertBeforeClose(data []byte, closeBrace int, content, closeIndent string) ([]byte, error) {
	lastSig := closeBrace - 1
	for lastSig >= 0 {
		ch := data[lastSig]
		if ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
			lastSig--
			continue
		}
		if ch == '/' && lastSig > 0 && data[lastSig-1] == '/' {
			for lastSig > 0 && data[lastSig-1] != '\n' {
				lastSig--
			}
			lastSig--
			continue
		}
		if ch == '/' && lastSig >= 2 && data[lastSig-1] == '*' {
			j := lastSig - 2
			for j >= 1 {
				if data[j] == '/' && data[j+1] == '*' {
					lastSig = j - 1
					break
				}
				j--
			}
			if j < 1 {
				break
			}
			continue
		}
		break
	}

	if lastSig < 0 || data[lastSig] == '{' || data[lastSig] == '[' {
		inserted := "\n" + content + "\n" + closeIndent
		return spliceBytes(data, closeBrace, closeBrace, []byte(inserted)), nil
	}

	if data[lastSig] == ',' {
		inserted := "\n" + content + "\n" + closeIndent
		return spliceBytes(data, lastSig+1, closeBrace, []byte(inserted)), nil
	}

	inserted := ",\n" + content + "\n" + closeIndent
	return spliceBytes(data, lastSig+1, closeBrace, []byte(inserted)), nil
}

func spliceBytes(data []byte, start, end int, insert []byte) []byte {
	var buf []byte
	buf = append(buf, data[:start]...)
	buf = append(buf, insert...)
	buf = append(buf, data[end:]...)
	return buf
}

func backupFile(path string, data []byte) error {
	backup, err := os.CreateTemp("", filepath.Base(path)+".backup-*")
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
	pterm.Info.Printfln("Backed up %s to %s", path, backup.Name())
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

type zedServerEntry struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

type zedServersWrapper struct {
	ContextServers map[string]zedServerEntry `json:"context_servers"`
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
