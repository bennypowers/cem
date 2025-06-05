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
	generate "bennypowers.dev/cem/generate"
)

var expand = A.Chain(func (g string) []string {
	paths, err := filepath.Glob(g)
	if err != nil {
		log.Fatal(err)
	}
	return paths
})

var excludes []string
var output string

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
			files := expand(args)
			exclude := expand(excludes)
			err, manifest := generate.Generate(files, exclude)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error generating manifest: %s", err)
			}
			if output != "" {
				os.WriteFile(output, []byte(*manifest), 0666)
				fmt.Println("Wrote manifest to", output)
			} else {
				fmt.Println(manifest)
			}
		},
	}

	generateCmd.PersistentFlags().StringArrayVarP(&excludes,
																								"exclude",
																								"e",
																								make([] string, 0),
																								"files or glob patterns to exclude")
	generateCmd.PersistentFlags().StringVarP(&output,
																						"output",
																						"o",
																						"",
																						"write custom elements manifest to this file")
	rootCmd.AddCommand(generateCmd)
}
