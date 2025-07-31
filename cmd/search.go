/*
Copyright © 2025 Benny Powers

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

	"github.com/spf13/cobra"

	"bennypowers.dev/cem/search"
	W "bennypowers.dev/cem/workspace"
)

var searchCmd = &cobra.Command{
	Use:   "search [pattern]",
	Short: "Search for tags, modules, attributes, and other elements by keyword or regex pattern",
	Long: `Search through the custom elements manifest for any element matching the given pattern.

The search pattern is treated as a regular expression by default, allowing for powerful
and flexible searches. If the regex is invalid, it falls back to literal string matching.

The search looks through names, descriptions, summaries, and labels of all manifest items
including tags, modules, attributes, slots, CSS properties, CSS states, CSS parts, events,
methods, demos, functions, variables, and more.

Examples:

  cem search button                    # Find anything containing "button"  
  cem search "^my-.*button$"           # Regex: elements starting with "my-" and ending with "button"
  cem search "click|hover"             # Regex: elements containing "click" OR "hover"
  cem search --format tree deprecated  # Search for "deprecated" and show as tree
  cem search "css.*property"           # Find CSS-related properties
  cem search "slot.*header"            # Find header-related slots
`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pattern := args[0]
		if pattern == "" {
			return errors.New("search pattern cannot be empty")
		}

		if ctx, err := W.GetWorkspaceContext(cmd); err != nil {
			return fmt.Errorf("project context not initialized: %w", err)
		} else {
			manifest, err := ctx.Manifest()
			if err != nil {
				return err
			}

			format, err := requireFormat(cmd, []string{"table", "tree"})
			if err != nil {
				return err
			}

			if s, err := search.RenderSearchResults(manifest, pattern, format); err != nil {
				return err
			} else {
				fmt.Println(s)
			}
			return nil
		}
	},
}

func init() {
	searchCmd.Flags().StringP("format", "f", "table", "Output format (table or tree)")
	rootCmd.AddCommand(searchCmd)
}
