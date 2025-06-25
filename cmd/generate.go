/**
 * Copyright © 2025 Benny Powers <web@bennypowers.com>
 */
package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	G "bennypowers.dev/cem/generate"
	A "github.com/IBM/fp-go/array"
	DS "github.com/bmatcuk/doublestar"
	"github.com/spf13/cobra"
)

var expand = A.Chain(func (g string) []string {
	paths, err := DS.Glob(g)
	if err != nil {
		log.Fatal(err)
	}
	return paths
})

func init() {
	// generateCmd represents the generate command
	var generateCmd = &cobra.Command{
		Use:   "generate [files or glob patterns]",
		Short: "Generates a custom elements manifest",
		Args: func(cmd *cobra.Command, args []string) error {
				if len(args) < 1 {
					return errors.New("requires at least one file argument")
				}
				return nil
			},
		Run: func(cmd *cobra.Command, args []string) {
			CemConfig.Generate.Files = expand(args)
			CemConfig.Generate.Exclude = expand(CemConfig.Generate.Exclude)
			manifest, err := G.Generate(&CemConfig)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error generating manifest: %s", err)
			}
			if CemConfig.Generate.Output != "" {
				os.WriteFile(CemConfig.Generate.Output, []byte(*manifest), 0666)
				fmt.Println("\nWrote manifest to", CemConfig.Generate.Output)
			} else {
				fmt.Println(*manifest)
			}
		},
	}

	generateCmd.PersistentFlags().StringArrayVarP(&CemConfig.Generate.Exclude,
																								"exclude",
																								"e",
																								make([] string, 0),
																								"files or glob patterns to exclude")
	generateCmd.PersistentFlags().StringVarP(&CemConfig.Generate.DesignTokensSpec,
																								"design-tokens",
																								"t",
																								"",
																								"specifiers (relative paths or npm:@scope/package/path/file.json) to DTCG-format module design tokens")
	generateCmd.PersistentFlags().StringVarP(&CemConfig.Generate.DesignTokensPrefix,
																						"design-tokens-prefix",
																						"p",
																						"",
																						"css custom property prefix for design tokens")
	generateCmd.PersistentFlags().StringVarP(&CemConfig.Generate.Output,
																						"output",
																						"o",
																						"",
																						"write custom elements manifest to this file")
		generateCmd.PersistentFlags().BoolVar(&CemConfig.Generate.NoDefaultExcludes,
																						"no-default-excludes",
																						false,
																						"do not exclude files by default (e.g. .d.ts files are included unless excluded explicitly)",
																					)
	rootCmd.AddCommand(generateCmd)
}
