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
	"path/filepath"
	"time"

	G "bennypowers.dev/cem/generate"
	W "bennypowers.dev/cem/workspace"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var start time.Time

// Uses the global ctx from root.go
var generateCmd = &cobra.Command{
	Use:   "generate [files or glob patterns]",
	Short: "Generates a custom elements manifest",
	RunE: func(cmd *cobra.Command, args []string) (errs error) {
		start = time.Now()
		ctx, err := W.GetWorkspaceContext(cmd)
		if err != nil {
			return fmt.Errorf("project context not initialized: %w", err)
		}

		// de-dupe globs
		allGlobs := append(args, viper.GetStringSlice("generate.files")...)
		seen := make(map[string]bool)
		uniqueGlobs := []string{}
		for _, glob := range allGlobs {
			if !seen[glob] {
				seen[glob] = true
				uniqueGlobs = append(uniqueGlobs, glob)
			}
		}

		files, err := expand(ctx, uniqueGlobs)
		if err != nil {
			errs = errors.Join(errs, err)
		}

		if len(files) == 0 {
			return errors.New("pass at least one file to generate")
		}

		exclude, err := expand(ctx, viper.GetStringSlice("generate.exclude"))
		if err != nil {
			errs = errors.Join(errs, err)
		}

		cfg, err := ctx.Config()
		if err != nil {
			errs = errors.Join(errs, err)
			return errs
		}

		cfg.Generate.Files = files
		cfg.Generate.Exclude = exclude

		// Check if watch mode is enabled
		watch, err := cmd.Flags().GetBool("watch")
		if err != nil {
			return err
		}

		if watch {
			return runWatchMode(ctx, uniqueGlobs)
		}

		// compute path to write custom elements manifest to
		// consider moving this to the context struct
		// if this is empty, we'll print to stdout instead
		outputPath, err := cmd.Flags().GetString("output")
		if err != nil {
			return err
		}

		if outputPath == "" {
			outputPath = cfg.Generate.Output
		}
		if outputPath == "" {
			pkgjson, err := ctx.PackageJSON()
			if err != nil {
				return err
			}
			if pkgjson != nil {
				outputPath = filepath.Join(ctx.Root(), pkgjson.CustomElements)
			}
		}

		// generate the manifest
		manifestStr, err := G.Generate(ctx)
		if err != nil {
			errs = errors.Join(errs, err)
		}

		if outputPath != "" {
			writer, err := ctx.OutputWriter(outputPath)
			if err != nil {
				errs = errors.Join(errs, err)
			} else {
				defer writer.Close()
				_, err := writer.Write([]byte(*manifestStr + "\n"))
				if err != nil {
					errs = errors.Join(errs, err)
				} else {
					end := time.Since(start)
					reloutputpath, err := filepath.Rel(ctx.Root(), outputPath)
					if err != nil {
						reloutputpath = outputPath
					}
					fmtstr := "Wrote manifest to %s in %s\n"
					args := []any{reloutputpath, G.ColorizeDuration(end).Sprint(end)}
					if errs != nil {
						pterm.Warning.Printf(fmtstr, args...)
					} else {
						pterm.Success.Printf(fmtstr, args...)
					}
				}
			}
		} else {
			fmt.Println(*manifestStr + "\n")
		}
		return errs
	},
}

// Use WorkspaceContext to expand globs
func expand(ctx W.WorkspaceContext, globs []string) (files []string, errs error) {
	for _, pattern := range globs {
		matches, err := ctx.Glob(pattern)
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
	generateCmd.Flags().BoolP("watch", "w", false, "watch files for changes and regenerate")
	viper.BindPFlag("generate.noDefaultExcludes", generateCmd.Flags().Lookup("no-default-excludes"))
	viper.BindPFlag("generate.output", generateCmd.Flags().Lookup("output"))
	viper.BindPFlag("generate.exclude", generateCmd.Flags().Lookup("exclude"))
	viper.BindPFlag("generate.designTokens.spec", generateCmd.Flags().Lookup("design-tokens"))
	viper.BindPFlag("generate.designTokens.prefix", generateCmd.Flags().Lookup("design-tokens-prefix"))
	viper.BindPFlag("generate.demoDiscovery.fileGlob", generateCmd.Flags().Lookup("demo-discovery-file-glob"))
	viper.BindPFlag("generate.demoDiscovery.urlPattern", generateCmd.Flags().Lookup("demo-discovery-url-pattern"))
	viper.BindPFlag("generate.demoDiscovery.urlTemplate", generateCmd.Flags().Lookup("demo-discovery-url-template"))
}

// runWatchMode starts the file watching mode - delegates to generate package
func runWatchMode(ctx W.WorkspaceContext, globs []string) error {
	session, err := G.NewWatchSession(ctx, globs)
	if err != nil {
		return err
	}
	defer session.Close()

	return session.RunWatch()
}
