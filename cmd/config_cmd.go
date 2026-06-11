package cmd

import (
	"encoding/json"
	"fmt"

	W "bennypowers.dev/cem/internal/workspace"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage CEM configuration",
	Long:  "Commands for creating, validating, and inspecting CEM configuration files.",
}

var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a CEM configuration file",
	Long:  "Interactive wizard that detects project settings and generates .config/cem.yaml.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("not yet implemented")
	},
}

var configValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate the configuration file",
	Long:  "Check the config file for invalid values, unreachable patterns, and missing files.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("not yet implemented")
	},
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Print the resolved configuration",
	Long:  "Print the fully resolved and merged config, showing workspace inheritance.",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := W.GetWorkspaceContext(cmd)
		if err != nil {
			return fmt.Errorf("project context not initialized: %w", err)
		}
		cfg, err := ctx.Config()
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}
		if cfg == nil {
			return fmt.Errorf("no configuration found")
		}

		format, _ := cmd.Flags().GetString("format")
		switch format {
		case "json":
			data, err := json.MarshalIndent(cfg, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to marshal config: %w", err)
			}
			cmd.Println(string(data))
		default:
			data, err := yaml.Marshal(cfg)
			if err != nil {
				return fmt.Errorf("failed to marshal config: %w", err)
			}
			cmd.Print(string(data))
		}
		return nil
	},
}

var configMCPCmd = &cobra.Command{
	Use:   "mcp",
	Short: "Generate MCP client configuration",
	Long:  "Generate configuration snippets for AI tools (Claude Code, Claude Desktop, Cursor, etc.).",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("not yet implemented")
	},
}

var configPathCmd = &cobra.Command{
	Use:   "path",
	Short: "Print the config file path",
	Long:  "Print the resolved config file path. Useful for scripting: $EDITOR $(cem config path)",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := W.GetWorkspaceContext(cmd)
		if err != nil {
			return fmt.Errorf("project context not initialized: %w", err)
		}
		cfgFile := ctx.ConfigFile()
		if cfgFile == "" {
			root := ctx.Root()
			return fmt.Errorf("no config file found (searched %s)", root)
		}
		cmd.Println(cfgFile)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configInitCmd)
	configCmd.AddCommand(configValidateCmd)
	configCmd.AddCommand(configShowCmd)
	configCmd.AddCommand(configMCPCmd)
	configCmd.AddCommand(configPathCmd)

	configShowCmd.Flags().String("format", "yaml", "output format (yaml or json)")
}
