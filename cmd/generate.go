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
	"os"
	"time"

	G "bennypowers.dev/cem/generate"
	M "bennypowers.dev/cem/manifest"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var generateFiles []string
var start time.Time

// Uses the global ctx from root.go
var generateCmd = &cobra.Command{
	Use:   "generate [files or glob patterns]",
	Short: "Generates a custom elements manifest",
	RunE: func(cmd *cobra.Command, args []string) (errs error) {
		start = time.Now()

		ctx, err := GetInitializedProjectContext(cmd)
		if err != nil {
			return fmt.Errorf("project context not initialized: %w", err)
		}
		cfg, err := ctx.Config()
		if err != nil {
			return fmt.Errorf("config not initialized: %w", err)
		}

		if noDefaultExcludes, _ := cmd.Flags().GetBool("no-default-excludes"); noDefaultExcludes {
			cfg.Generate.NoDefaultExcludes = noDefaultExcludes
		}
		if output, _ := cmd.Flags().GetString("output"); output != "" {
			cfg.Generate.Output = output
		}
		if exclude, _ := cmd.Flags().GetStringSlice("exclude"); exclude != nil {
			excludes, err := expand(ctx, exclude)
			if err != nil {
				errs = errors.Join(errs, err)
			} else {
				cfg.Generate.Exclude = excludes
			}
		}
		if tokens, _ := cmd.Flags().GetString("design-tokens"); tokens != "" {
			cfg.Generate.DesignTokens.Spec = tokens
		}
		if prefix, _ := cmd.Flags().GetString("design-tokens-prefix"); prefix != "" {
			cfg.Generate.DesignTokens.Prefix = prefix
		}
		if glob, _ := cmd.Flags().GetString("demo-discovery-file-glob"); glob != "" {
			cfg.Generate.DemoDiscovery.FileGlob = glob
		}
		if pattern, _ := cmd.Flags().GetString("demo-discovery-url-pattern"); pattern != "" {
			cfg.Generate.DemoDiscovery.URLPattern = pattern
		}
		if template, _ := cmd.Flags().GetString("demo-discovery-url-template"); template != "" {
			cfg.Generate.DemoDiscovery.URLTemplate = template
		}

		viperGenerateFiles := viper.GetStringSlice("generate.files")

		combined := append(viperGenerateFiles, args...)
		files, err := expand(ctx, combined)
		if err != nil {
			errs = errors.Join(errs, err)
		}
		cfg.Generate.Files = append(cfg.Generate.Files, combined...)

		fmt.Fprintf(
			os.Stderr,
			`
cfg=%+v\n`,
			cfg,
		)
		if len(files) < 1 {
			return errors.New("requires at least one file argument or a configured `generate.files` list")
		}

		manifestStr, err := G.Generate(ctx, cfg)
		if err != nil {
			errs = errors.Join(errs, err)
		}
		if manifestStr == nil {
			return errors.Join(errs, errors.New("manifest generation returned nil"))
		}
		if cfg.Generate.Output != "" {
			writer, err := ctx.OutputWriter(cfg.Generate.Output)
			if err != nil {
				errs = errors.Join(errs, err)
			} else {
				defer writer.Close()
				_, err := writer.Write([]byte(*manifestStr + "\n"))
				if err != nil {
					errs = errors.Join(errs, err)
				} else {
					end := time.Since(start)
					outputPath := cfg.Generate.Output
					pterm.Success.Printf("Wrote manifest to %s in %s", outputPath, G.ColorizeDuration(end).Sprint(end))
				}
			}
		} else {
			fmt.Println(*manifestStr + "\n")
		}
		return errs
	},
}

// Use ProjectContext to expand globs
func expand(ctx M.ProjectContext, globs []string) (files []string, errs error) {
	for _, pattern := range globs {
		matches, err := ctx.ListFiles(pattern)
		if err != nil {
			errs = errors.Join(errs, err)
			continue
		}
		files = append(files, matches...)
	}
	return files, errs
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().Bool("no-default-excludes", false, "do not exclude files by default (e.g. .d.ts files are included unless excluded explicitly)")
	generateCmd.Flags().StringP("output", "o", "", "write custom elements manifest to this file")
	generateCmd.Flags().StringArrayP("exclude", "e", make([]string, 0), "files or glob patterns to exclude")
	generateCmd.Flags().String("design-tokens", "", "specifiers (relative paths or npm:@scope/package/path/file.json) to DTCG-format module design tokens")
	generateCmd.Flags().String("design-tokens-prefix", "", "css custom property prefix for design tokens")
	generateCmd.Flags().String("demo-discovery-file-glob", "", "Glob pattern for discovering demo files")
	generateCmd.Flags().String("demo-discovery-url-pattern", "", "Go Regexp pattern with named capture groups for generating canonical demo urls")
	generateCmd.Flags().String("demo-discovery-url-template", "", "URL pattern string using {groupName} syntax to interpolate named captures from the URL pattern")
	viper.BindPFlag("generate.noDefaultExcludes", generateCmd.Flags().Lookup("no-default-excludes"))
	viper.BindPFlag("generate.output", generateCmd.Flags().Lookup("output"))
	viper.BindPFlag("generate.exclude", generateCmd.Flags().Lookup("exclude"))
	viper.BindPFlag("generate.designTokens.spec", generateCmd.Flags().Lookup("design-tokens"))
	viper.BindPFlag("generate.designTokens.prefix", generateCmd.Flags().Lookup("design-tokens-prefix"))
	viper.BindPFlag("generate.demoDiscovery.fileGlob", generateCmd.Flags().Lookup("demo-discovery-file-glob"))
	viper.BindPFlag("generate.demoDiscovery.urlPattern", generateCmd.Flags().Lookup("demo-discovery-url-pattern"))
	viper.BindPFlag("generate.demoDiscovery.urlTemplate", generateCmd.Flags().Lookup("demo-discovery-url-template"))
}
