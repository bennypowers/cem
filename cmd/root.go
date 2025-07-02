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
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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

func resolveProjectDir(configPath, projectDirFlag string) (string, bool) {
	if projectDirFlag != "" {
		abs, err := expandPath(projectDirFlag)
		if err != nil {
			pterm.Fatal.Printf("Invalid --project-dir: %v", err)
		}
		return abs, true
	}
	configAbs, err := filepath.Abs(configPath)
	if err != nil {
		pterm.Fatal.Printf("Invalid --config: %v", err)
	}
	configDir := filepath.Dir(configAbs)
	base := filepath.Base(configDir)
	if base == ".config" || base == "config" {
		return filepath.Dir(configDir), true
	}
	// fallback: use current working directory
	cwd, err := os.Getwd()
	if err != nil {
		pterm.Fatal.Printf("Unable to get current working directory: %v", err)
	}
	if !strings.HasPrefix(configAbs, cwd) {
		pterm.Warning.Printf("Warning: --config is outside of current dir, guessing project root as %s\n", cwd)
	}
	return cwd, false
}

// expandPath expands ~, handles relative and absolute paths
func expandPath(path string) (string, error) {
	if path == "" {
		return "", nil
	}
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		// Support ~/ and ~
		if path == "~" {
			path = home
		} else if strings.HasPrefix(path, "~/") {
			path = filepath.Join(home, path[2:])
		}
		// Note: ~user/ is not supported (Go stdlib doesn't provide this)
	}
	return filepath.Abs(path)
}

func initConfig() {
	var err error
	cfgFile := viper.GetString("configFile")
	projectDir, shouldChange := resolveProjectDir(cfgFile, viper.GetString("projectDir"))
	viper.Set("projectDir", projectDir)
	viper.AddConfigPath(filepath.Join(projectDir, ".config"))
	viper.SetConfigType("yaml")
	viper.SetConfigName("cem")
	if shouldChange {
		// Search config in home directory with name ".snakes" (without extension).
		if err := os.Chdir(projectDir); err != nil {
			cobra.CheckErr(errors.Join(err, errors.New("Failed to change into project directory")))
		}
	}
	if viper.GetBool("verbose") {
		pterm.EnableDebugMessages()
	}
	pterm.Debug.Println("Using project directory: ", projectDir)
	if cfgFile != "" {
		// Use config file from the flag.
		cfgFile, err = expandPath(cfgFile)
		cobra.CheckErr(err)
	} else {
		// Search config in local project .config directory with name "cem.yaml"
		cfgFile, err = expandPath(filepath.Join(projectDir, ".config", "cem.yaml"))
		cobra.CheckErr(err)
	}
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		if err := viper.ReadInConfig(); err == nil {
			pterm.Debug.Println("Using config file: ", cfgFile)
		}
	}
	viper.Set("configFile", cfgFile)

	viper.AutomaticEnv()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().String("source-control-root-url", "", "Canonical public source control URL corresponding to project root on primary branch. e.g. https://github.com/bennypowers/cem/tree/main/")
	rootCmd.PersistentFlags().String("config", "", "config file (default is $CWD/.config/cem.yaml)")
	rootCmd.PersistentFlags().String("project-dir", "", "Path to project directory (default: parent directory of .config/cem.yaml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose logging output")
	viper.BindPFlag("configFile", rootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("projectDir", rootCmd.PersistentFlags().Lookup("project-dir"))
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("sourceControlRootUrl", rootCmd.PersistentFlags().Lookup("source-control-root-url"))
}
