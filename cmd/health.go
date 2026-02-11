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
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"

	"bennypowers.dev/cem/health"
	M "bennypowers.dev/cem/manifest"
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
	Use:   "health [source files...]",
	Short: "Score documentation quality in a custom-elements manifest",
	Long: `Analyze a custom-elements.json manifest and score the quality of its documentation.

Source file paths can be passed as positional arguments to filter the report
to modules matching those files. File extensions are mapped to their compiled
output equivalents (.ts -> .js, .tsx -> .js, etc.) and resolved through the
package.json exports map, replicating the same logic used by cem generate.`,
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, err := workspace.GetWorkspaceContext(cmd)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting workspace context: %v\n", err)
			os.Exit(1)
		}

		manifestPath := ctx.CustomElementsManifestPath()
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
		allDisabled := make([]string, 0, len(configDisabled)+len(disableFlags))
		allDisabled = append(allDisabled, configDisabled...)
		allDisabled = append(allDisabled, disableFlags...)

		configModules := viper.GetStringSlice("health.modules")
		allModules := make([]string, 0, len(configModules)+len(moduleFlags))
		allModules = append(allModules, configModules...)
		allModules = append(allModules, moduleFlags...)

		// Resolve positional file args to module paths
		if len(args) > 0 {
			pkgJSON, _ := ctx.PackageJSON()

			// Determine the package path prefix to strip from file args.
			// File args (e.g. from git diff) are repo-root-relative, but
			// module paths in the manifest are package-root-relative.
			pkgPrefix := ""
			if p, _ := cmd.Flags().GetString("package"); p != "" && p != "." {
				pkgPrefix = filepath.ToSlash(filepath.Clean(p)) + "/"
			}

			for _, file := range args {
				rel := filepath.ToSlash(file)
				rel = strings.TrimPrefix(rel, pkgPrefix)
				normalized := M.NormalizeSourcePath(rel)
				resolved, err := M.ResolveExportPath(pkgJSON, normalized)
				if err != nil {
					if errors.Is(err, M.ErrNotExported) {
						// Not exported — use normalized path as-is (same as generate)
						allModules = append(allModules, normalized)
					}
					// Other files (README, tests, etc.) — silently skip
					continue
				}
				allModules = append(allModules, resolved)
			}
		}

		configFailBelow := viper.GetInt("health.failBelow")
		if failBelow == 0 && configFailBelow > 0 {
			failBelow = configFailBelow
		}

		options := health.Options{
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
