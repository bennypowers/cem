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
	"os"

	V "bennypowers.dev/cem/validate"
	"bennypowers.dev/cem/workspace"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	validateCmd.Flags().BoolP("verbose", "v", false, "Show detailed information including schema version")
	validateCmd.Flags().StringArray("disable", []string{}, "Disable specific warning rules or categories")
	validateCmd.Flags().String("format", "text", "Output format: text or json")
	rootCmd.AddCommand(validateCmd)
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate a custom-elements.json manifest",
	Long:  `Validate a custom-elements.json manifest against its JSON schema.`,
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
		verbose, _ := cmd.Flags().GetBool("verbose")
		disableFlags, _ := cmd.Flags().GetStringArray("disable")
		format, _ := cmd.Flags().GetString("format")

		// Merge config and flag disabled rules
		configDisabled := viper.GetStringSlice("warnings.disable")
		allDisabled := append(configDisabled, disableFlags...)

		// Set up validation options
		options := V.ValidationOptions{
			IncludeWarnings: true, // Always include warnings, but filter them
			DisabledRules:   allDisabled,
		}

		// Validate the manifest
		result, err := V.Validate(manifestPath, options)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error validating manifest: %v\n", err)
			os.Exit(1)
		}

		// Print the results
		displayOptions := V.DisplayOptions{
			Verbose: verbose,
			Format:  format,
		}
		V.PrintValidationResult(manifestPath, result, displayOptions)

		// Exit with error code if validation failed
		if !result.IsValid {
			os.Exit(1)
		}
	},
}
