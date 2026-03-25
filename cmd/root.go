/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

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
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"

	IC "bennypowers.dev/cem/internal/config"
	"bennypowers.dev/cem/internal/logging"
	W "bennypowers.dev/cem/workspace"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cem",
	Short: "Tool for generating and querying custom-elements manifests",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		_, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("unable to get current working directory: %v", err)
		}

		wctx, err := W.GetWorkspaceContext(cmd)
		if err != nil {
			return fmt.Errorf("failed to create project context: %v", err)
		}

		// Store the project context in the Cobra context
		ctx := context.WithValue(cmd.Context(), W.WorkspaceContextKey, wctx)
		cmd.SetContext(ctx)

		rootDir := wctx.Root()
		if p, _ := cmd.Flags().GetString("package"); p == "" {
			viper.Set("package", rootDir)
		}

		// Handle verbose and quiet flags (mutually exclusive)
		verbose := viper.GetBool("verbose")
		quiet := viper.GetBool("quiet")

		if verbose && quiet {
			return fmt.Errorf("cannot use both --verbose and --quiet flags together")
		}

		if verbose {
			pterm.EnableDebugMessages()
		}

		// Configure quiet mode to suppress INFO and DEBUG messages
		if quiet {
			logging.SetQuietEnabled(true)
		}

		pterm.Debug.Printfln("Project directory: %q", rootDir)

		// Load config into viper using the discovered config file path.
		// The workspace context already discovered the file via internal/config.FindConfigFile.
		cfgFile := wctx.ConfigFile()
		if cfgFile != "" {
			format := IC.FormatFromPath(cfgFile)
			if format == "jsonc" {
				// Viper doesn't support JSONC natively. Read the file,
				// strip comments, and feed clean JSON to viper.
				data, readErr := os.ReadFile(cfgFile)
				if readErr != nil {
					return fmt.Errorf("failed to read config file %s: %w", cfgFile, readErr)
				}
				stripped := IC.StripComments(data)
				viper.SetConfigType("json")
				if err := viper.ReadConfig(bytes.NewReader(stripped)); err != nil {
					return fmt.Errorf("failed to parse config file %s: %w", cfgFile, err)
				}
			} else {
				viper.SetConfigFile(cfgFile)
				if format != "" {
					viper.SetConfigType(format)
				}
				if err := viper.ReadInConfig(); err != nil && !errors.Is(err, os.ErrNotExist) {
					return err
				}
			}
			viper.Set("configFile", cfgFile)
			pterm.Debug.Printfln("Config file: %q", cfgFile)
		}
		viper.AutomaticEnv()
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("source-control-root-url", "", "Canonical public source control URL corresponding to project root on primary branch. e.g. https://github.com/bennypowers/cem/tree/main/")
	rootCmd.PersistentFlags().String("config", "", "config file (searches .config/cem.{yaml,yml,json,jsonc} then .cem.{yaml,yml,json,jsonc})")
	rootCmd.PersistentFlags().StringP("package", "p", "", "package specifier: npm:@scope/package, URL (https://cdn.example.com/pkg/), or local path")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose logging output")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "quiet output (only warnings and errors)")

	_ = viper.BindPFlag("configFile", rootCmd.PersistentFlags().Lookup("config"))
	_ = viper.BindPFlag("package", rootCmd.PersistentFlags().Lookup("package"))
	_ = viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	_ = viper.BindPFlag("quiet", rootCmd.PersistentFlags().Lookup("quiet"))
	_ = viper.BindPFlag("sourceControlRootUrl", rootCmd.PersistentFlags().Lookup("source-control-root-url"))

	rootCmd.PersistentFlags().String("project-dir", "", "Path to project directory (default: parent directory of .config/cem.yaml)")
	_ = rootCmd.PersistentFlags().MarkDeprecated("project-dir", "Will be removed, use --package instead")
}
