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
	"bennypowers.dev/cem/cmd/config"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var verbose bool
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
		pterm.Fatal.Print(err)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in local project .config directory with name "cem.yaml"
		viper.AddConfigPath("./.config")
		viper.SetConfigType("yaml")
		viper.SetConfigName("cem.yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if CemConfig.Verbose {
			pterm.EnableDebugMessages()
		} else {
			pterm.DisableDebugMessages()
		}
		pterm.Info.Print("Using config file:", viper.ConfigFileUsed(), "\n")
	}
	// Bind to struct
	if err := viper.Unmarshal(&CemConfig); err != nil {
		pterm.Error.Println("Unable to decode into CemConfig struct:", err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $CWD/.config/cem.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&CemConfig.Verbose, "verbose", "v", false, "verbose logging output")
}

