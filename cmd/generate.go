/*
Copyright © 2025 Benny Powers <web@bennypowers.com>
*/
package cmd

import (
	"errors"
	"log"
	"path/filepath"

	generate "github.com/bennypowers/cemgen/lib"
	"github.com/repeale/fp-go"
	"github.com/spf13/cobra"
)

var expand = fp.FlatMap(func (g string) []string {
	paths, err := filepath.Glob(g)
	if err != nil {
		log.Fatal(err)
	}
	return paths
})

var excludes []string

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
			generate.Generate(files, exclude)
		},
	}

	generateCmd.PersistentFlags().StringArrayVarP(&excludes,
																								"exclude",
																								"e",
																								make([] string, 0),
																								"files or glob patterns to exclude")
	rootCmd.AddCommand(generateCmd)
}
