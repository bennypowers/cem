/*
Copyright © 2025 Benny Powers <web@bennypowers.com>
*/
package cmd

import (
	"errors"
	"fmt"
	"sync"

	"github.com/spf13/cobra"
)

var wg sync.WaitGroup

func analyze(file string) {
	defer wg.Done()
	fmt.Println(file)
}

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
		for _, file := range args {
			wg.Add(1)
			go analyze(file)
		}
		wg.Wait()
		fmt.Println("Done!")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
