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
	"os"
	"path/filepath"
	"strings"

	M "bennypowers.dev/cem/manifest"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initialCWD string
var ctx M.ProjectContext

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cem",
	Short: "Generate a custom elements manifest",
	Long: `Scans your projects TypeScript sources and identifies custom elements.
Generates a custom elements manifest file (custom-elements.json) describing your modules.
Supports projects written with Lit`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Helper to create the correct ProjectContext
func resolveProjectContext(configPath, packageFlag string) (M.ProjectContext, error) {
	var ctx M.ProjectContext
	if packageFlag != "" {
		if isLikelyPath(packageFlag) {
			ctx = M.NewLocalFSProjectContext(packageFlag)
		} else {
			ctx = M.NewRemoteProjectContext(packageFlag)
		}
	} else {
		ctx = M.NewLocalFSProjectContext(filepath.Dir(configPath))
	}
	if err := ctx.Init(); err != nil {
		return nil, err
	} else {
		return ctx, nil
	}
}

// Simple heuristic: if it starts with ./, ../, /, or ~, or doesn't contain ":" or "@", treat as path
func isLikelyPath(spec string) bool {
	return strings.HasPrefix(spec, ".") ||
		strings.HasPrefix(spec, "/") ||
		strings.HasPrefix(spec, "~") ||
		(!strings.Contains(spec, ":") && !strings.Contains(spec, "@"))
}

func initConfig() {
	var err error
	// initialCWD still used for output display
	initialCWD, err = os.Getwd()
	if err != nil {
		pterm.Fatal.Printf("Unable to get current working directory: %v", err)
	}

	cfgFile := viper.GetString("configFile")
	packageFlag := viper.GetString("package")

	ctx, err = resolveProjectContext(cfgFile, packageFlag)
	if err != nil {
		pterm.Fatal.Printf("Failed to create project context: %v", err)
	}

	// Use projectContext for config file/dir resolution
	rootDir := ctx.Root()
	viper.Set("package", rootDir)
	viper.AddConfigPath(filepath.Join(rootDir, ".config"))
	viper.SetConfigType("yaml")
	viper.SetConfigName("cem")

	// No need for os.Chdir! All path access now through context

	if viper.GetBool("verbose") {
		pterm.EnableDebugMessages()
	}
	pterm.Debug.Println("Using project directory: ", rootDir)

	// Use context for config file existence, reading, etc.
	cfgFile, err = ctx.ConfigFile()
	cobra.CheckErr(err)

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		viper.Set("configFile", cfgFile)
		if err := viper.ReadInConfig(); err == nil {
			pterm.Debug.Println("Using config file: ", cfgFile)
		}
	}
	viper.AutomaticEnv()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().String("source-control-root-url", "", "Canonical public source control URL corresponding to project root on primary branch. e.g. https://github.com/bennypowers/cem/tree/main/")
	rootCmd.PersistentFlags().String("config", "", "config file (default is $CWD/.config/cem.yaml)")
	rootCmd.PersistentFlags().String("project-dir", "", "Path to project directory (default: parent directory of .config/cem.yaml)")
	rootCmd.PersistentFlags().MarkDeprecated("project-dir", "Will be removed, use --package instead")
	rootCmd.PersistentFlags().StringP("package", "p", "", "deno-style package specifier e.g. npm:@scope/package, or path to package directory")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose logging output")
	viper.BindPFlag("configFile", rootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("projectDir", rootCmd.PersistentFlags().Lookup("project-dir"))
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("sourceControlRootUrl", rootCmd.PersistentFlags().Lookup("source-control-root-url"))
}
