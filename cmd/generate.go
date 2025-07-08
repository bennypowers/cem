/**
 * Copyright © 2025 Benny Powers <web@bennypowers.com>
 */
package cmd

import (
	"errors"
	"fmt"
	"os"
	"time"

	"bennypowers.dev/cem/cmd/config"
	G "bennypowers.dev/cem/generate"
	DS "github.com/bmatcuk/doublestar"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

var generateFiles []string

var start time.Time

var generateCmd = &cobra.Command{
	Use:   "generate [files or glob patterns]",
	Short: "Generates a custom elements manifest",
	Args: func(cmd *cobra.Command, args []string) error {
		// If we have args (i.e. files), that's fine
		// Or if no args, but files are configured, allow
		start = time.Now()
		var err error
		generateFiles, err = expand(viper.GetStringSlice("generate.files"))
		if err != nil {
			return err
		}
		if len(args) > 0 || (len(args) == 0 && len(viper.GetStringSlice("generate.files")) > 0) {
			return nil
		}
		// Otherwise, error
		return errors.New("requires at least one file argument or a configured `generate.files` list")
	},
	RunE: func(cmd *cobra.Command, args []string) (errs error) {
		files, err := expand(append(viper.GetStringSlice("generate.files"), args...))
		if err != nil {
			errs = errors.Join(errs, err)
		}
		exclude, err := expand(viper.GetStringSlice("generate.exclude"))
		if err != nil {
			errs = errors.Join(errs, err)
		}
		cfg := config.CemConfig{}
		viper.Unmarshal(&cfg)
		cfg.Generate.Files = files
		cfg.Generate.Exclude = exclude
		manifest, err := G.Generate(&cfg)
		if err != nil {
			errs = errors.Join(errs, err)
		}
		if manifest == nil {
			return errors.Join(errs, errors.New("manifest generation returned nil"))
		}
		if cfg.Generate.Output != "" {
			if err = os.WriteFile(cfg.Generate.Output, []byte(*manifest+"\n"), 0666); err != nil {
				errs = errors.Join(errs, err)
			} else {
				end := time.Since(start)
				pterm.Success.Printf("Wrote manifest to %s in %s", cfg.Generate.Output, G.ColorizeDuration(end).Sprint(end))
			}
		} else {
			fmt.Println(*manifest + "\n")
		}
		return errs
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().Bool("no-default-excludes", false, "do not exclude files by default (e.g. .d.ts files are included unless excluded explicitly)")
	generateCmd.Flags().StringP("output", "o", "", "write custom elements manifest to this file")
	generateCmd.Flags().StringArrayP("exclude", "e", make([]string, 0), "files or glob patterns to exclude")
	generateCmd.Flags().StringP("design-tokens", "t", "", "specifiers (relative paths or npm:@scope/package/path/file.json) to DTCG-format module design tokens")
	generateCmd.Flags().StringP("design-tokens-prefix", "p", "", "css custom property prefix for design tokens")
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
