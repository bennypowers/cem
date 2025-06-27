/**
 * Copyright © 2025 Benny Powers <web@bennypowers.com>
 */
package cmd

import (
	"errors"
	"fmt"
	"os"

	G "bennypowers.dev/cem/generate"
	DS "github.com/bmatcuk/doublestar"
	"github.com/spf13/cobra"
)

func expand(globs []string) (files []string, errs error) {
	for _, pattern := range globs {
		matches, err := DS.Glob(pattern)
		if err != nil {
			errs = errors.Join(errs, err)
			continue
		}
		files = append(files, matches...)
	}
	return files, errs
}

func init() {
	// generateCmd represents the generate command
	var generateCmd = &cobra.Command{
		Use:   "generate [files or glob patterns]",
		Short: "Generates a custom elements manifest",
		Args: func(cmd *cobra.Command, args []string) error {
			// If we have args (i.e. files), that's fine
			// Or if no args, but files are configured, allow
			if len(args) > 0 || (len(args) == 0 && len(CemConfig.Generate.Files) > 0) {
				return nil
			}
			// Otherwise, error
			return errors.New("requires at least one file argument or a configured `generate.files` list")
		},
		RunE: func(cmd *cobra.Command, args []string) (errs error) {
			var err error
			CemConfig.Generate.Files, err = expand(append(CemConfig.Generate.Files, args...))
			if err != nil {
				errs = errors.Join(errs, err)
			}
			CemConfig.Generate.Exclude, err = expand(CemConfig.Generate.Exclude)
			if err != nil {
				errs = errors.Join(errs, err)
			}
			manifest, err := G.Generate(&CemConfig)
			if err != nil {
				errs = errors.Join(errs, err)
			}
			if manifest == nil {
				return errors.Join(errs, errors.New("manifest generation returned nil"))
			}
			if CemConfig.Generate.Output != "" {
				if err = os.WriteFile(CemConfig.Generate.Output, []byte(*manifest + "\n"), 0666); err != nil {
					errs = errors.Join(errs, err)
				} else {
					fmt.Println("\nWrote manifest to", CemConfig.Generate.Output)
				}
			} else {
				fmt.Println(*manifest)
			}
			return errs
		},
	}

	generateCmd.PersistentFlags().StringVarP(&CemConfig.SourceControlRootUrl,
																								"source-control-root-url",
																								"",
																								"",
																								"Canonical public source control URL corresponding to project root on primary branch. e.g. https://github.com/bennypowers/cem/tree/main/")

	generateCmd.PersistentFlags().StringVarP(&CemConfig.Generate.Output,
																						"output",
																						"o",
																						"",
																						"write custom elements manifest to this file")

	generateCmd.PersistentFlags().StringArrayVarP(&CemConfig.Generate.Exclude,
																								"exclude",
																								"e",
																								make([] string, 0),
																								"files or glob patterns to exclude")
	generateCmd.PersistentFlags().BoolVar(&CemConfig.Generate.NoDefaultExcludes,
																						"no-default-excludes",
																						false,
																						"do not exclude files by default (e.g. .d.ts files are included unless excluded explicitly)")

	generateCmd.PersistentFlags().StringVarP(&CemConfig.Generate.DesignTokens.Spec,
																								"design-tokens",
																								"t",
																								"",
																								"specifiers (relative paths or npm:@scope/package/path/file.json) to DTCG-format module design tokens")
	generateCmd.PersistentFlags().StringVarP(&CemConfig.Generate.DesignTokens.Prefix,
																						"design-tokens-prefix",
																						"p",
																						"",
																						"css custom property prefix for design tokens")

	generateCmd.PersistentFlags().StringVarP(&CemConfig.Generate.DemoDiscovery.FileGlob,
																								"demo-discovery-file-glob",
																								"",
																								"",
																								"Glob pattern for discovering demo files")
	generateCmd.PersistentFlags().StringVarP(&CemConfig.Generate.DemoDiscovery.URLPattern,
																								"demo-discovery-url-pattern",
																								"",
																								"",
																								"Go Regexp pattern with named capture groups for generating canonical demo urls")
	generateCmd.PersistentFlags().StringVarP(&CemConfig.Generate.DemoDiscovery.URLTemplate,
																								"demo-discovery-url-template",
																								"",
																								"",
																								"URL pattern string using {groupName} syntax to interpolate named captures from the URL pattern")

	rootCmd.AddCommand(generateCmd)
}
