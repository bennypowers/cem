/*
Copyright © 2025 Benny Powers <web@bennypowers.com>
*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	A "github.com/IBM/fp-go/array"
	G "bennypowers.dev/cem/generate"
)

var expand = A.Chain(func (g string) []string {
	paths, err := filepath.Glob(g)
	if err != nil {
		log.Fatal(err)
	}
	return paths
})

var generateArgs G.GenerateArgs

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
			generateArgs.Files = expand(args)
			generateArgs.Exclude = expand(generateArgs.Exclude)
			manifest, err := G.Generate(generateArgs)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error generating manifest: %s", err)
			}
			if generateArgs.Output != "" {
				os.WriteFile(generateArgs.Output, []byte(*manifest), 0666)
				fmt.Println("\nWrote manifest to", generateArgs.Output)
			} else {
				fmt.Println(*manifest)
			}
		},
	}

	generateCmd.PersistentFlags().StringArrayVarP(&generateArgs.Exclude,
																								"exclude",
																								"e",
																								make([] string, 0),
																								"files or glob patterns to exclude")
	generateCmd.PersistentFlags().StringVarP(&generateArgs.DesignTokensSpec,
																								"design-tokens",
																								"t",
																								"",
																								"specifiers (relative paths or npm:@scope/package/path/file.json) to DTCG-format module design tokens")
	generateCmd.PersistentFlags().StringVarP(&generateArgs.DesignTokensPrefix,
																						"design-tokens-prefix",
																						"p",
																						"",
																						"css custom property prefix for design tokens")
	generateCmd.PersistentFlags().StringVarP(&generateArgs.Output,
																						"output",
																						"o",
																						"",
																						"write custom elements manifest to this file")
		generateCmd.PersistentFlags().BoolVar(&generateArgs.NoDefaultExcludes,
																						"no-default-excludes",
																						false,
																						"do not exclude files by default (e.g. .d.ts files are included unless excluded explicitly)",
																					)
	rootCmd.AddCommand(generateCmd)
}
