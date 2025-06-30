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

	"bennypowers.dev/cem/cmd/config"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var cfgFile string
var projectDir string
var CemConfig config.CemConfig

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
		pterm.Error.Print(err)
		os.Exit(1)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Skip config loading for completion commands
	if isCompletionCommand() {
		return
	}
	
	cfg, err := config.LoadConfig(cfgFile, projectDir)
	if err != nil {
		pterm.Error.Print(err)
		os.Exit(1)
	}
	CemConfig = *cfg
	if err := os.Chdir(CemConfig.GetProjectDir()); err != nil {
		pterm.Error.Print("Failed to change into project directory: ", err)
		os.Exit(1)
	}
	pterm.Info.Print("Using project directory: ", CemConfig.GetProjectDir(), "\n")
	pterm.Info.Print("Using config file: ", cfg.GetConfigFile(), "\n")
}

// isCompletionCommand checks if the current command is a completion command
func isCompletionCommand() bool {
	if len(os.Args) < 2 {
		return false
	}
	// Check if the command is "completion" or any of its subcommands
	return os.Args[1] == "completion" || 
	       (len(os.Args) >= 3 && os.Args[1] == "completion" && 
	        (os.Args[2] == "bash" || os.Args[2] == "zsh" || os.Args[2] == "fish" || os.Args[2] == "powershell"))
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $CWD/.config/cem.yaml)")
	rootCmd.PersistentFlags().StringVar(&projectDir, "project-dir", "", "Path to project directory (default: parent directory of .config/cem.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&CemConfig.Verbose, "verbose", "v", false, "verbose logging output")
}

