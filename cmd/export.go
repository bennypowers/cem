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

	C "bennypowers.dev/cem/cmd/config"
	"bennypowers.dev/cem/export"
	"bennypowers.dev/cem/workspace"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	exportCmd.Flags().String("format", "", "Framework to export: react, vue, or angular (exports all configured frameworks if omitted)")
	exportCmd.Flags().StringP("output", "o", "", "Output directory for generated files")
	exportCmd.Flags().String("strip-prefix", "", "Prefix to strip from tag names (e.g., 'demo-')")
	rootCmd.AddCommand(exportCmd)
}

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Generate framework wrapper components from a custom-elements manifest",
	Long: `Export generates framework-specific wrapper components (React, Vue, Angular) from
a custom-elements.json manifest. The generated wrappers provide native framework DX
with compile-time type checking and IDE autocomplete.

Frameworks can be configured in .config/cem.yaml under the 'export' key, or
selected via the --format flag for one-off usage.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := workspace.GetWorkspaceContext(cmd)
		if err != nil {
			return fmt.Errorf("getting workspace context: %w", err)
		}

		manifest, err := ctx.Manifest()
		if err != nil {
			return fmt.Errorf("reading manifest: %w", err)
		}
		if manifest == nil {
			fmt.Fprintln(os.Stderr, "Could not find custom-elements.json")
			os.Exit(1)
		}

		// Determine package name
		packageName := viper.GetString("packageName")
		if packageName == "" {
			if pkgJSON, err := ctx.PackageJSON(); err == nil && pkgJSON != nil {
				packageName = pkgJSON.Name
			}
		}

		// Build framework configs from YAML config
		frameworks := make(map[string]export.FrameworkExportConfig)

		var configExport map[string]C.FrameworkExportConfig
		if err := viper.UnmarshalKey("export", &configExport); err == nil {
			for name, cfg := range configExport {
				frameworks[name] = export.FrameworkExportConfig{
					Output:      cfg.Output,
					StripPrefix: cfg.StripPrefix,
					PackageName: cfg.PackageName,
					ModuleName:  cfg.ModuleName,
				}
			}
		}

		// Apply CLI flag overrides
		format, _ := cmd.Flags().GetString("format")
		output, _ := cmd.Flags().GetString("output")
		stripPrefix, _ := cmd.Flags().GetString("strip-prefix")

		if format != "" {
			// Single framework mode: use only the specified framework
			cfg, exists := frameworks[format]
			if !exists {
				cfg = export.FrameworkExportConfig{}
			}
			if output != "" {
				cfg.Output = output
			}
			if stripPrefix != "" {
				cfg.StripPrefix = stripPrefix
			}
			frameworks = map[string]export.FrameworkExportConfig{format: cfg}
		} else if output != "" || stripPrefix != "" {
			// Apply flags to all configured frameworks
			for name, cfg := range frameworks {
				if output != "" {
					cfg.Output = output
				}
				if stripPrefix != "" {
					cfg.StripPrefix = stripPrefix
				}
				frameworks[name] = cfg
			}
		}

		if len(frameworks) == 0 {
			return fmt.Errorf("no frameworks configured; use --format flag or add 'export' section to .config/cem.yaml")
		}

		return export.Export(export.Options{
			Manifest:    manifest,
			PackageName: packageName,
			Frameworks:  frameworks,
		})
	},
}
