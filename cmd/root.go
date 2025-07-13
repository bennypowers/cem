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
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	W "bennypowers.dev/cem/workspace"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initialCWD string

type contextKey string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cem",
	Short: "Tool for generating and querying custom-elements manifests",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		initialCWD, err = os.Getwd()
		if err != nil {
			return fmt.Errorf("Unable to get current working directory: %v", err)
		}

		wctx, err := W.GetWorkspaceContext(cmd)
		if err != nil {
			return fmt.Errorf("Failed to create project context: %v", err)
		}

		// Store the project context in the Cobra context
		ctx := context.WithValue(cmd.Context(), W.WorkspaceContextKey, wctx)
		cmd.SetContext(ctx)

		rootDir := wctx.Root()
		if p, _ := cmd.Flags().GetString("package"); p == "" {
			viper.Set("package", rootDir)
		}
		viper.AddConfigPath(filepath.Join(rootDir, ".config"))
		viper.SetConfigType("yaml")
		viper.SetConfigName("cem")

		if viper.GetBool("verbose") {
			pterm.EnableDebugMessages()
		}
		pterm.Debug.Println("Using project directory: ", rootDir)

		cfgFile := wctx.ConfigFile()
		if cfgFile != "" {
			viper.SetConfigFile(cfgFile)
			viper.Set("configFile", cfgFile)
			if err := viper.ReadInConfig(); err != nil && !errors.Is(err, os.ErrNotExist) {
				return err
			} else {
				pterm.Debug.Println("Using config file: ", cfgFile)
			}
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
	rootCmd.PersistentFlags().String("config", "", "config file (default is $CWD/.config/cem.yaml)")
	rootCmd.PersistentFlags().StringP("package", "p", "", "deno-style package specifier e.g. npm:@scope/package, or path to package directory")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose logging output")

	viper.BindPFlag("configFile", rootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("package", rootCmd.PersistentFlags().Lookup("package"))
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("sourceControlRootUrl", rootCmd.PersistentFlags().Lookup("source-control-root-url"))

	rootCmd.PersistentFlags().String("project-dir", "", "Path to project directory (default: parent directory of .config/cem.yaml)")
	rootCmd.PersistentFlags().MarkDeprecated("project-dir", "Will be removed, use --package instead")
}
