package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	DD "bennypowers.dev/cem/generate/demodiscovery"
	IC "bennypowers.dev/cem/internal/config"
	"bennypowers.dev/cem/internal/diagnostic"
	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/internal/sourcepos"
	W "bennypowers.dev/cem/internal/workspace"
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
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SetOut(os.Stdout)
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

		cfgFormat := IC.FormatFromPath(cfgFile)
		schemaErrs := IC.ValidateSchema(rawData, cfgFormat)

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

		allErrs := IC.DeduplicateErrors(schemaErrs, semanticErrs)

		if len(allErrs) > 0 {
			enrichPositions(allErrs, rawData, cfgFormat)
		}

		format, _ := cmd.Flags().GetString("format")
		switch format {
		case "json":
			return printConfigValidationJSON(cmd, cfgFile, allErrs)
		default:
			return printConfigValidationText(cmd, cfgFile, rawData, allErrs)
		}
	},
}

func splitBySeverity(errs []IC.ValidationError) (errors, warnings []IC.ValidationError) {
	for _, e := range errs {
		if e.Severity == IC.SeverityWarning {
			warnings = append(warnings, e)
		} else {
			errors = append(errors, e)
		}
	}
	return
}

func printConfigValidationJSON(cmd *cobra.Command, cfgFile string, errs []IC.ValidationError) error {
	errors, warnings := splitBySeverity(errs)

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

func enrichPositions(errs []IC.ValidationError, rawData []byte, format string) {
	posMap := sourcepos.BuildPositionMap(rawData, format)
	for i := range errs {
		pointer := sourcepos.FieldToJSONPointer(errs[i].Field)
		if pos, ok := sourcepos.Resolve(posMap, pointer); ok {
			errs[i].Line = pos.Line
			errs[i].Column = pos.Column
		}
	}
}

func printConfigValidationText(cmd *cobra.Command, cfgFile string, source []byte, errs []IC.ValidationError) error {
	errors, warnings := splitBySeverity(errs)

	if len(errors) == 0 && len(warnings) == 0 {
		cmd.Printf("config: %s\n\nvalid\n", cfgFile)
		return nil
	}

	w := cmd.OutOrStdout()
	diags := make([]diagnostic.Diagnostic, 0, len(errs))
	for _, e := range errs {
		diags = append(diags, diagnostic.Diagnostic{
			File:     cfgFile,
			Pos:      sourcepos.Position{Line: e.Line, Column: e.Column},
			Severity: e.Severity,
			Field:    e.Field,
			Message:  e.Message,
			Value:    e.Value,
			Source:   source,
		})
	}
	diagnostic.RenderAll(w, diags)

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
	configCmd.AddCommand(configPathCmd)

	configShowCmd.Flags().String("format", "yaml", "output format (yaml or json)")
	configValidateCmd.Flags().String("format", "text", "output format (text or json)")
}
