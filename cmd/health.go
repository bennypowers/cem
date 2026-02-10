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
	"fmt"
	"math"
	"os"

	"bennypowers.dev/cem/health"
	"bennypowers.dev/cem/workspace"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	healthCmd.Flags().String("component", "", "Filter to a specific component by tag name or class name")
	healthCmd.Flags().StringArray("module", []string{}, "Filter to specific modules by path (can be repeated)")
	healthCmd.Flags().String("format", "text", "Output format: text, json, or markdown")
	healthCmd.Flags().Int("fail-below", 0, "Exit 1 if overall percentage is below this threshold (0-100)")
	healthCmd.Flags().StringArray("disable", []string{}, "Disable specific health categories by ID")
	rootCmd.AddCommand(healthCmd)
}

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Score documentation quality in a custom-elements manifest",
	Long:  `Analyze a custom-elements.json manifest and score the quality of its documentation.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, err := workspace.GetWorkspaceContext(cmd)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting workspace context: %v\n", err)
			os.Exit(1)
		}

		manifestPath := ctx.CustomElementsManifestPath()
		if len(args) > 0 {
			manifestPath = args[0]
		}

		if manifestPath == "" {
			fmt.Fprintln(os.Stderr, "Could not find custom-elements.json")
			os.Exit(1)
		}

		// Get flags
		component, _ := cmd.Flags().GetString("component")
		moduleFlags, _ := cmd.Flags().GetStringArray("module")
		format, _ := cmd.Flags().GetString("format")
		switch format {
		case "text", "json", "markdown":
		default:
			fmt.Fprintf(os.Stderr, "Invalid format %q: must be text, json, or markdown\n", format)
			os.Exit(1)
		}
		failBelow, _ := cmd.Flags().GetInt("fail-below")
		disableFlags, _ := cmd.Flags().GetStringArray("disable")

		// Merge config and flag values
		configDisabled := viper.GetStringSlice("health.disable")
		allDisabled := append(configDisabled, disableFlags...)

		configModules := viper.GetStringSlice("health.modules")
		allModules := append(configModules, moduleFlags...)

		configFailBelow := viper.GetInt("health.failBelow")
		if failBelow == 0 && configFailBelow > 0 {
			failBelow = configFailBelow
		}

		options := health.HealthOptions{
			Component: component,
			Modules:   allModules,
			Disable:   allDisabled,
		}

		result, err := health.Analyze(manifestPath, options)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error analyzing manifest: %v\n", err)
			os.Exit(1)
		}

		displayOptions := health.DisplayOptions{
			Format: format,
		}
		if err := health.PrintHealthResult(result, displayOptions); err != nil {
			fmt.Fprintf(os.Stderr, "Error displaying health results: %v\n", err)
			os.Exit(1)
		}

		// Check fail-below threshold
		if failBelow > 0 && result.OverallMax > 0 {
			pct := int(math.Round(float64(result.OverallScore) * 100 / float64(result.OverallMax)))
			if pct < failBelow {
				fmt.Fprintf(os.Stderr, "Health score %d%% (%d/%d) is below the --fail-below threshold of %d%%\n",
					pct, result.OverallScore, result.OverallMax, failBelow)
				os.Exit(1)
			}
		}
	},
}
