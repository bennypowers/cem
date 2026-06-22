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
	"path/filepath"

	V "bennypowers.dev/cem/validate"
	"bennypowers.dev/cem/internal/workspace"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func validateWorkspace(cmd *cobra.Command) error {
	ctx, err := workspace.GetWorkspaceContext(cmd)
	if err != nil {
		return err
	}

	disableFlags, _ := cmd.Flags().GetStringArray("disable")
	format, _ := cmd.Flags().GetString("format")
	configDisabled := viper.GetStringSlice("warnings.disable")
	allDisabled := make([]string, 0, len(configDisabled)+len(disableFlags))
	allDisabled = append(allDisabled, configDisabled...)
	allDisabled = append(allDisabled, disableFlags...)

	var collected []*V.ValidationResult

	results := workspace.ForEachPackage(ctx.Root(), func(pkg workspace.PackageInfo) error {
		manifestPath := filepath.Join(pkg.Path, pkg.CustomElementsRef)
		options := V.ValidationOptions{
			IncludeWarnings: true,
			DisabledRules:   allDisabled,
		}
		result, err := V.Validate(manifestPath, options)
		if err != nil {
			return err
		}

		switch format {
		case "text", "":
			displayOptions := V.DisplayOptions{Format: format}
			if _, err := fmt.Fprintln(cmd.OutOrStdout(), "\n"+pkg.Name+":"); err != nil {
				return err
			}
			if err := V.PrintValidationResult(cmd.OutOrStdout(), manifestPath, result, displayOptions); err != nil {
				return err
			}
		case "json":
			result.Path = manifestPath
			collected = append(collected, result)
		default:
			return fmt.Errorf("invalid format %q: must be text or json", format)
		}

		if !result.IsValid {
			return fmt.Errorf("validation failed")
		}
		return nil
	})

	if format != "text" && format != "" {
		if err := V.PrintValidationResultsJSON(cmd.OutOrStdout(), collected); err != nil {
			return err
		}
	}

	return workspace.ReportResults("Validated manifests", results)
}

func init() {
	validateCmd.Flags().StringArray("disable", []string{}, "Disable specific warning rules or categories")
	validateCmd.Flags().String("format", "text", "Output format: text or json")
	rootCmd.AddCommand(validateCmd)
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate a custom-elements.json manifest",
	Long:  `Validate a custom-elements.json manifest against its JSON schema.`,
	Args:         cobra.MaximumNArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if workspace.ShouldUseWorkspaceMode(cmd) && len(args) == 0 {
			return validateWorkspace(cmd)
		}

		ctx, err := workspace.GetWorkspaceContext(cmd)
		if err != nil {
			return err
		}

		manifestPath := ctx.CustomElementsManifestPath()
		if len(args) > 0 {
			manifestPath = args[0]
		}

		if manifestPath == "" {
			return fmt.Errorf("could not find custom-elements.json")
		}

		disableFlags, _ := cmd.Flags().GetStringArray("disable")
		format, _ := cmd.Flags().GetString("format")

		configDisabled := viper.GetStringSlice("warnings.disable")
		allDisabled := append(configDisabled, disableFlags...)

		options := V.ValidationOptions{
			IncludeWarnings: true,
			DisabledRules:   allDisabled,
		}

		result, err := V.Validate(manifestPath, options)
		if err != nil {
			return err
		}

		switch format {
		case "json":
			result.Path = manifestPath
			if err := V.PrintValidationResultsJSON(cmd.OutOrStdout(), []*V.ValidationResult{result}); err != nil {
				return err
			}
		default:
			displayOptions := V.DisplayOptions{Format: format}
			if err := V.PrintValidationResult(cmd.OutOrStdout(), manifestPath, result, displayOptions); err != nil {
				return err
			}
		}

		if !result.IsValid {
			return fmt.Errorf("validation failed")
		}
		return nil
	},
}
