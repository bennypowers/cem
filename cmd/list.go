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

	A "github.com/IBM/fp-go/array"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	C "bennypowers.dev/cem/cmd/config"
	M "bennypowers.dev/cem/manifest"
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

func requireTagName(cmd *cobra.Command) (string, error) {
	tagName, err := cmd.Flags().GetString("tag-name")
	if err != nil {
		return "", err
	}
	if tagName == "" {
		return "", errors.New("must supply --tag-name")
	}
	return tagName, nil
}

func requireFormat(cmd *cobra.Command) (string, error) {
	format, err := cmd.Flags().GetString("format")
	if err != nil {
		return "", err
	}
	switch format {
	case "table":
		return format, nil
	// Add more supported formats here as the CLI grows
	default:
		return "", errors.New("unknown format: " + format)
	}
}

func requireTagFormatPkg(cmd *cobra.Command) (tagName string, format string, pkg *M.Package, err error) {
	tagName, err = requireTagName(cmd)
	if err != nil { return "", "", nil, err }
	pkg, err = readPkg()
	if err != nil { return "", "", nil, err }
	format, err = requireFormat(cmd)
	if err != nil { return "", "", nil, err }
	return tagName, format, pkg, err
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List items in the custom elements manifest like tag names, attributes, functions, etc",
	Long: `List various entities defined in the custom elements manifest.

This command provides subcommands for listing tag names, attributes, and other elements described in your manifest.
Use the available subcommands to explore your project's custom elements, their attributes, and more.

Examples:

  cem list tags
  cem list attributes --tag-name my-button
  cem list slots --tag-name my-button
  cem list css-custom-properties --tag-name my-button
  cem list css-custom-states --tag-name my-button
  cem list css-parts --tag-name my-button
  cem list events --tag-name my-button
  cem list methods --tag-name my-button
`,
}

var listTagsCmd = &cobra.Command{
	Use:   "tags",
	Aliases: []string{"elements", "tag-names"},
	Short: "List tag names in the custom elements manifest",
	Long: `List all custom element tag names

This command outputs a table with tag names and their corresponding source modules,
allowing you to quickly see which custom elements are available and where they are defined.

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
		}

		return nil
	},
}

var listAttrsCmd = &cobra.Command{
	Use:   "attrs",
	Aliases: []string{"attributes"},
	Short: "List attributes in the custom elements manifest by tag name",
	Long: `List all attributes for a given custom element tag

You must specify the tag name using the --tag-name flag. The output includes each attribute, its corresponding DOM property, and whether it reflects changes to the DOM.

Examples:

  cem list attrs --tag-name my-button --format table
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagName, format, pkg, err := requireTagFormatPkg(cmd)
		if err != nil { return err }
		attrs, err := pkg.GetTagAttrsWithContext(tagName)
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
		}

		return nil
	},
}

var listSlotsCmd = &cobra.Command{
	Use:   "slots",
	Short: "List slots in the custom elements manifest by tag name",
	Long: `List all slots for a given custom element tag.

You must specify the tag name using the --tag-name flag. The output includes each slot name, and it's summary

Examples:

  cem list slots --tag-name my-button --format table
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagName, format, pkg, err := requireTagFormatPkg(cmd)
		if err != nil { return err }
		slots, err := pkg.GetTagSlotsWithContext(tagName)
		if err != nil {
			return err
		}
		switch format {
		case "table":
			data := pterm.TableData{
			}
			pterm.DefaultTable.WithHasHeader().WithBoxed(true).WithData(data).Render()
		}

		return nil
	},
}

var listCssCustomPropertiesCmd = &cobra.Command{
	Use:   "css-custom-properties",
	Aliases: []string{"css-properties", "css-props", "css-custom-props"},
	Short: "List CSS custom properties used in a given tag name",
	Long: `List CSS custom properties used in a given tag name.

You must specify the tag name using the --tag-name flag. The output includes each css property by name, it's default value, and a summary

Examples:

  cem list css-custom-properties --tag-name my-button --format table
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagName, format, pkg, err := requireTagFormatPkg(cmd)
		if err != nil { return err }
		props, err := pkg.GetTagCssPropertiesWithContext(tagName)
		if err != nil { return err }
		switch format {
		case "table":
			data := pterm.TableData{
			}
			pterm.DefaultTable.WithHasHeader().WithBoxed(true).WithData(data).Render()
		}

		return nil
	},
}

var listCssCustomStatesCmd = &cobra.Command{
	Use:   "css-custom-states",
	Aliases: []string{"css-states"},
	Short: "List CSS custom states used in a given tag name",
	Long: `List CSS custom states used in a given tag name.

You must specify the tag name using the --tag-name flag. The output includes each css state by name, and a summary

Examples:

  cem list css-custom-states --tag-name my-button --format table
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagName, format, pkg, err := requireTagFormatPkg(cmd)
		if err != nil { return err }
		props, err := pkg.GetTagCssPropertiesWithContext(tagName)
		if err != nil { return err }
		switch format {
		case "table":
			data := pterm.TableData{
			}
			pterm.DefaultTable.WithHasHeader().WithBoxed(true).WithData(data).Render()
		}

		return nil
	},
}

var listCssPartsCmd = &cobra.Command{
	Use:   "css-parts",
	Aliases: []string{"css-shadow-parts"},
	Short: "List CSS shadow parts for a given tag name",
	Long: `List CSS shadow parts for a given tag name.

You must specify the tag name using the --tag-name flag. The output includes each shadow part by name, and a summary

Examples:

  cem list css-parts --tag-name my-button --format table
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagName, format, pkg, err := requireTagFormatPkg(cmd)
		if err != nil { return err }
		parts, err := pkg.GetTagCssPartsWithContext(tagName)
		if err != nil { return err }
		switch format {
		case "table":
			data := pterm.TableData{
			}
			pterm.DefaultTable.WithHasHeader().WithBoxed(true).WithData(data).Render()
		}

		return nil
	},
}

var listEventsCmd = &cobra.Command{
	Use:   "events",
	Short: "List events for a given tag name",
	Long: `List JavaScript events fired by a given element by tag name.

You must specify the tag name using the --tag-name flag. The output includes each event, it's type, and a summary

Examples:

  cem list event --tag-name my-button --format table
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagName, format, pkg, err := requireTagFormatPkg(cmd)
		if err != nil { return err }
		events, err := pkg.GetTagEventsWithContext(tagName)
		if err != nil { return err }
		switch format {
		case "table":
			data := pterm.TableData{
			}
			pterm.DefaultTable.WithHasHeader().WithBoxed(true).WithData(data).Render()
		}

		return nil
	},
}

var listMethodsCmd = &cobra.Command{
	Use:   "methods",
	Short: "List class methods for a given tag name",
	Long: `List DOM object methods for the class registered to a given tag name.

You must specify the tag name using the --tag-name flag. The output includes each method, it's return type, and a summary

Examples:

  cem list methods --tag-name my-button --format table
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagName, format, pkg, err := requireTagFormatPkg(cmd)
		if err != nil { return err }
		methods, err := pkg.GetTagMethodsWithContext(tagName)
		if err != nil { return err }
		switch format {
		case "table":
			data := pterm.TableData{
			}
			pterm.DefaultTable.WithHasHeader().WithBoxed(true).WithData(data).Render()
		}

		return nil
	},
}

func init() {
	listCmd.AddCommand(listTagsCmd)
	listCmd.PersistentFlags().StringP("format", "f", "table", "Output format")
	for _, c := range []*cobra.Command{
		listAttrsCmd,
		listSlotsCmd,
		listCssCustomPropertiesCmd,
		listCssCustomStatesCmd,
		listCssPartsCmd,
		listEventsCmd,
		listMethodsCmd,
	} {
		listCmd.AddCommand(c)
		c.Flags().StringP("tag-name", "t", "", "Tag name to list attributes for")
	}
	rootCmd.AddCommand(listCmd)
}
