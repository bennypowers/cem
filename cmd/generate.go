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

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a custom elements manifest",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one file argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		generate.Generate(expand(args))
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
