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
	"os"
	"path/filepath"

	C "bennypowers.dev/cem/cmd/config"
	M "bennypowers.dev/cem/manifest"
	A "github.com/IBM/fp-go/array"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func readPkg() (*M.Package, error) {
	cfg := C.CemConfig{}
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	path := cfg.Generate.Output
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	path = filepath.Join(cwd, path)
	json, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return M.UnmarshalPackage(json)
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List items in the custom elements manifest like tag names, attributes, functions, etc",
	Long: `List various entities defined in the custom elements manifest (CEM) generated from your project.

This command provides subcommands for listing tag names, attributes, and other elements described in your manifest.
Use the available subcommands to explore your project's custom elements, their attributes, and more.

Examples:

  cem list tags
  cem list attrs --tag-name rh-button
`,
}

var listTagsCmd = &cobra.Command{
	Use:   "tags",
	Aliases: []string{"elements", "tag-names"},
	Short: "List tag names in the custom elements manifest",
	Long: `List all custom element tag names defined in your project's custom elements manifest (CEM).

This command outputs a table with tag names and their corresponding source modules, allowing you to quickly see which custom elements are available and where they are defined.

Example:

  cem list tags --format table
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		pkg, err := readPkg()
		if err != nil {
			return err
		}

		tags := pkg.GetAllTagNamesWithContext()
		format, err := cmd.Flags().GetString("format")
		if err != nil {
			return err
		}
		switch format {
		case "table":
			data := pterm.TableData{
				{"Tag", "Module"},
			}
			data = append(data, A.Map(func(r M.CustomElementWithContext) []string {
				return []string{r.TagName, r.Module.Path}
			})(tags)...)
			pterm.DefaultTable.WithHasHeader().WithBoxed(true).WithData(data).Render()
		default:
			return errors.New("Unknown format " + format)
		}

		return nil
	},
}

var listAttrsCmd = &cobra.Command{
	Use:   "attrs",
	Aliases: []string{"attributes"},
	Short: "List attributes in the custom elements manifest by tag name",
	Long: `List all attributes for a given custom element tag as described in the custom elements manifest (CEM).

You must specify the tag name using the --tag-name flag. The output includes each attribute, its corresponding DOM property, and whether it reflects changes to the DOM.

Examples:

  cem list attrs --tag-name rh-dialog --format table

This is useful for understanding the public API surface (attributes/properties) of a custom element in your project.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagName, err := cmd.Flags().GetString("tag-name")
		if err != nil {
			return err
		}
		if tagName == "" {
			return errors.New("cem list attrs: must supply --tag-name")
		}
		pkg, err := readPkg()
		if err != nil {
			return err
		}
		attrs, err := pkg.GetTagAttrsWithContext(tagName)
		if err != nil {
			return err
		}
		format, err := cmd.Flags().GetString("format")
		if err != nil {
			return err
		}
		switch format {
		case "table":
			data := pterm.TableData{
				{"Attribute", "DOM Property", "Reflects"},
			}
			data = append(data, A.Map(func(r M.AttributeWithContext) []string {
				return []string{
					r.Name,
					r.CustomElementField.Name,
					func() string {
						if r.CustomElementField.Reflects {
							return "✅"
						} else {
							return "❌"
						}
					}()}
			})(attrs)...)
			pterm.DefaultTable.WithHasHeader().WithBoxed(true).WithData(data).Render()
		default:
			return errors.New("Unknown format " + format)
		}

		return nil
	},
}

func init() {
	listCmd.AddCommand(listAttrsCmd)
	listCmd.AddCommand(listTagsCmd)
	rootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().StringP("format", "f", "table", "Output format")
	listAttrsCmd.Flags().StringP("tag-name", "t", "", "Tag name to list attributes for")
}
