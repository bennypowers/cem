/*
Copyright © 2026 Benny Powers <web@bennypowers.com>

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

	"bennypowers.dev/cem/breaking"
	"bennypowers.dev/cem/internal/logging"
	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/internal/workspace"
	M "bennypowers.dev/cem/manifest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	breakingCmd.Flags().String("base", "", "Git ref for baseline (default: latest semver tag)")
	breakingCmd.Flags().String("head", "", "Git ref for head (default: working tree manifest)")
	breakingCmd.Flags().String("format", "text", "Output format: text, json, or markdown")
	breakingCmd.Flags().String("fail-on", "", "Exit 1 if changes at this severity or above: breaking or dangerous")
	breakingCmd.Flags().StringArray("disable", []string{}, "Disable specific breaking change rules (can be repeated)")
	rootCmd.AddCommand(breakingCmd)
}

var breakingCmd = &cobra.Command{
	Use:   "breaking [old-manifest.json new-manifest.json]",
	Short: "Detect breaking API changes between two custom-elements manifests",
	Long: `Compare two custom-elements.json manifests and detect breaking changes
in component APIs. Changes are classified as breaking, dangerous, or safe.

When given two file arguments, compares them directly. When given --base
(and optionally --head), resolves manifests from git refs. With no arguments,
auto-detects the baseline from the latest semver tag.`,
	Args:         cobra.MaximumNArgs(2),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		format, _ := cmd.Flags().GetString("format")
		switch format {
		case "text", "json", "markdown":
		default:
			return fmt.Errorf("invalid format %q: must be text, json, or markdown", format)
		}

		failOn, _ := cmd.Flags().GetString("fail-on")
		if failOn != "" && failOn != "breaking" && failOn != "dangerous" {
			return fmt.Errorf("invalid --fail-on value %q: must be breaking or dangerous", failOn)
		}

		disableFlags, _ := cmd.Flags().GetStringArray("disable")
		configDisabled := viper.GetStringSlice("breaking.disable")
		allDisabled := make([]string, 0, len(configDisabled)+len(disableFlags))
		allDisabled = append(allDisabled, configDisabled...)
		allDisabled = append(allDisabled, disableFlags...)

		var basePkg, headPkg *M.Package
		var err error

		fsys := platform.NewOSFileSystem()

		switch {
		case len(args) == 2:
			basePkg, err = loadManifestFile(fsys, args[0])
			if err != nil {
				return fmt.Errorf("failed to load base manifest: %w", err)
			}
			headPkg, err = loadManifestFile(fsys, args[1])
			if err != nil {
				return fmt.Errorf("failed to load head manifest: %w", err)
			}

		default:
			ctx, err := workspace.GetWorkspaceContext(cmd)
			if err != nil {
				return err
			}
			manifestPath := ctx.CustomElementsManifestPath()
			if manifestPath == "" {
				return fmt.Errorf("could not find custom-elements.json")
			}
			workDir := ctx.Root()

			baseRef, _ := cmd.Flags().GetString("base")
			headRef, _ := cmd.Flags().GetString("head")

			if baseRef == "" {
				baseRef, err = breaking.FindLatestSemverTag(workDir)
				if err != nil {
					return fmt.Errorf("no --base specified and %w; provide --base or pass two manifest files", err)
				}
				logging.Info("Using baseline: %s", baseRef)
			}

			basePkg, err = breaking.ResolveManifestFromGit(baseRef, manifestPath, workDir)
			if err != nil {
				return err
			}

			if headRef != "" {
				headPkg, err = breaking.ResolveManifestFromGit(headRef, manifestPath, workDir)
				if err != nil {
					return err
				}
			} else {
				headPkg, err = loadManifestFile(fsys, manifestPath)
				if err != nil {
					return fmt.Errorf("failed to load current manifest: %w", err)
				}
			}
		}

		result := breaking.Compare(basePkg, headPkg, breaking.Options{
			Disable: allDisabled,
		})

		if err := breaking.PrintResult(cmd.OutOrStdout(), result, breaking.DisplayOptions{Format: format}); err != nil {
			return err
		}

		if failOn == "breaking" && result.Breaking > 0 {
			return fmt.Errorf("%d breaking change(s) detected", result.Breaking)
		}
		if failOn == "dangerous" && (result.Breaking > 0 || result.Dangerous > 0) {
			return fmt.Errorf("%d breaking and %d dangerous change(s) detected", result.Breaking, result.Dangerous)
		}

		return nil
	},
}

func loadManifestFile(fsys platform.FileSystem, path string) (*M.Package, error) {
	data, err := fsys.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var pkg M.Package
	if err := json.Unmarshal(data, &pkg); err != nil {
		return nil, fmt.Errorf("failed to parse manifest %s: %w", path, err)
	}
	return &pkg, nil
}
