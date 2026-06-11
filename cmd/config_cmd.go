package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	IC "bennypowers.dev/cem/internal/config"
	"bennypowers.dev/cem/internal/platform"
	W "bennypowers.dev/cem/internal/workspace"
	DD "bennypowers.dev/cem/generate/demodiscovery"
	"bennypowers.dev/cem/serve/middleware/transform"
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
		ctx, err := W.GetWorkspaceContext(cmd)
		if err != nil {
			return fmt.Errorf("project context not initialized: %w", err)
		}
		cfgFile := ctx.ConfigFile()
		if cfgFile == "" {
			return fmt.Errorf("no config file found (searched %s)", ctx.Root())
		}

		rawData, err := os.ReadFile(cfgFile)
		if err != nil {
			return fmt.Errorf("failed to read config file: %w", err)
		}

		schemaErrs := IC.ValidateSchema(rawData)

		cfg, err := ctx.Config()
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		semanticErrs := IC.Validate(cfg, IC.ValidateOptions{
			CheckIO: true,
			Root:    ctx.Root(),
			IO:      &platform.OSFileSystem{},
			ValidateURLRewrites: func(rewrites []IC.URLRewrite) error {
				return transform.ValidateURLRewrites(rewrites)
			},
			ValidateDemoDiscovery: func(c *IC.CemConfig) error {
				return DD.ValidateDemoDiscoveryConfig(c, map[string]string{})
			},
		})

		allErrs := append(schemaErrs, semanticErrs...)

		format, _ := cmd.Flags().GetString("format")
		switch format {
		case "json":
			return printConfigValidationJSON(cmd, cfgFile, allErrs)
		default:
			return printConfigValidationText(cmd, cfgFile, allErrs)
		}
	},
}

func printConfigValidationJSON(cmd *cobra.Command, cfgFile string, errs []IC.ValidationError) error {
	var errors, warnings []IC.ValidationError
	for _, e := range errs {
		if e.Severity == IC.SeverityWarning {
			warnings = append(warnings, e)
		} else {
			errors = append(errors, e)
		}
	}

	result := struct {
		ConfigFile string               `json:"configFile"`
		Valid      bool                 `json:"valid"`
		Errors     []IC.ValidationError `json:"errors"`
		Warnings   []IC.ValidationError `json:"warnings"`
	}{
		ConfigFile: cfgFile,
		Valid:      len(errors) == 0,
		Errors:     errors,
		Warnings:   warnings,
	}
	if result.Errors == nil {
		result.Errors = []IC.ValidationError{}
	}
	if result.Warnings == nil {
		result.Warnings = []IC.ValidationError{}
	}

	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal results: %w", err)
	}
	cmd.Println(string(data))

	if len(errors) > 0 {
		return fmt.Errorf("config validation failed")
	}
	return nil
}

func printConfigValidationText(cmd *cobra.Command, cfgFile string, errs []IC.ValidationError) error {
	var errors, warnings []IC.ValidationError
	for _, e := range errs {
		if e.Severity == IC.SeverityWarning {
			warnings = append(warnings, e)
		} else {
			errors = append(errors, e)
		}
	}

	cmd.Printf("config: %s\n", cfgFile)

	if len(errors) == 0 && len(warnings) == 0 {
		cmd.Println("\nvalid")
		return nil
	}

	if len(errors) > 0 {
		cmd.Println("\nerrors:")
		for _, e := range errors {
			cmd.Printf("  %s\n", e.Error())
		}
	}

	if len(warnings) > 0 {
		cmd.Println("\nwarnings:")
		for _, e := range warnings {
			cmd.Printf("  %s\n", e.Error())
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("config validation failed")
	}
	return nil
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
	configValidateCmd.Flags().String("format", "text", "output format (text or json)")
}
