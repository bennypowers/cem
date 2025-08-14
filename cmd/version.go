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
	"encoding/json"
	"fmt"

	"bennypowers.dev/cem/internal/version"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Long:  `Print version information for cem.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		output, err := cmd.Flags().GetString("output")
		if err != nil {
			return fmt.Errorf("error reading output flag: %v", err)
		}
		switch output {
		case "json":
			printVersionJSON()
		default:
			printVersionText()
		}
		return nil
	},
}

func printVersionText() {
	fmt.Printf("cem %s\n", version.GetVersion())
}

func printVersionJSON() {
	buildInfo := version.GetBuildInfo()
	output, err := json.MarshalIndent(buildInfo, "", "  ")
	if err != nil {
		fmt.Printf("error marshaling version info: %v\n", err)
		return
	}
	fmt.Println(string(output))
}

func init() {
	rootCmd.AddCommand(versionCmd)

	versionCmd.Flags().StringP("output", "o", "text", "Output format: text or json")
}
