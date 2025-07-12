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
	"slices"

	"github.com/spf13/cobra"

	"bennypowers.dev/cem/list"
	M "bennypowers.dev/cem/manifest"
)

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

func requireFormat(cmd *cobra.Command, supportedFormats []string) (string, error) {
	format, err := cmd.Flags().GetString("format")
	if err != nil {
		return "", err
	}
	if slices.Contains(supportedFormats, format) {
		return format, nil
	}
	return "", errors.New("unknown format: " + format)
}

// Helper to generate list subcommands for custom elements by section
func makeListSectionCmd(use, short, long string, includeSection string, aliases ...string) *cobra.Command {
	return &cobra.Command{
		Use:     use,
		Aliases: aliases,
		Short:   short,
		Long:    long,
		RunE: func(cmd *cobra.Command, args []string) error {
			if ctx, err := GetProjectContext(cmd); err != nil {
				return fmt.Errorf("project context not initialized: %w", err)
			} else {
				tagName, err := requireTagName(cmd)
				if err != nil {
					return err
				}
				manifest, err := ctx.Manifest()
				if err != nil {
					return err
				}
				format, err := requireFormat(cmd, []string{"table"})
				if err != nil {
					return err
				}
				columns, err := cmd.Flags().GetStringArray("columns")
				if err != nil {
					return err
				}
				switch format {
				case "table":
					ced, _, mod, err := manifest.FindCustomElementContext(tagName)
					if err != nil {
						return err
					}
					opts := list.RenderOptions{
						Columns:         columns,
						IncludeSections: []string{includeSection},
					}
					if s, err := list.Render(M.NewRenderableCustomElementDeclaration(ced, mod, manifest), opts); err != nil {
						return err
					} else {
						fmt.Println(s)
					}
				}
				return nil
			}
		},
	}
}

var listAttrsCmd = makeListSectionCmd(
	"attributes",
	"List attributes in the custom elements manifest by tag name",
	`List all attributes for a given custom element tag

You must specify the tag name using the --tag-name flag. The output includes each attribute,
its corresponding DOM property, whether it reflects changes to the DOM, and a summary.

Examples:

  cem list attributes --tag-name my-button
  cem list attributes --tag-name my-button --format table --columns "DOM Property"
  cem list attributes --tag-name my-button --format table --columns "DOM Property" --columns Reflects
  cem list attributes --tag-name my-button --format table --columns "DOM Property" --columns Reflects --columns Summary
`,
	"Attributes",
	"attrs",
)

var listSlotsCmd = makeListSectionCmd(
	"slots",
	"List slots in the custom elements manifest by tag name",
	`List all slots for a given custom element tag.

You must specify the tag name using the --tag-name flag. The output includes each slot name, and it's summary

Examples:

  cem list slots --tag-name my-button
  cem list slots --tag-name my-button --format table --columns Name
`,
	"Slots",
)

var listCssCustomPropertiesCmd = makeListSectionCmd(
	"css-custom-properties",
	"List CSS custom properties used in a given tag name",
	`List CSS custom properties used in a given tag name.

You must specify the tag name using the --tag-name flag. The output may include for each
css property by name, it's syntax, default value, and summary.

Examples:

  cem list css-custom-properties --tag-name my-button
  cem list css-custom-properties --tag-name my-button --format table --columns Syntax
  cem list css-custom-properties --tag-name my-button --format table --columns Syntax --columns Default
  cem list css-custom-properties --tag-name my-button --format table --columns Syntax --columns Default --columns Summary
`,
	"CSS Properties",
	"css-properties", "css-props", "css-custom-props",
)

var listCssCustomStatesCmd = makeListSectionCmd(
	"css-custom-states",
	"List CSS custom states used in a given tag name",
	`List CSS custom states used in a given tag name.

You must specify the tag name using the --tag-name flag. The output includes each css state by name,
and a summary

Examples:

  cem list css-custom-states --tag-name my-button
  cem list css-custom-states --tag-name my-button --format table --columns Name
`,
	"CSS States",
	"css-states",
)

var listCssPartsCmd = makeListSectionCmd(
	"css-parts",
	"List CSS shadow parts for a given tag name",
	`List CSS shadow parts for a given tag name.

You must specify the tag name using the --tag-name flag. The output includes each shadow part by name, and a summary

Examples:

  cem list css-parts --tag-name my-button --format table
`,
	"CSS Parts",
	"parts", "css-shadow-parts",
)

var listEventsCmd = makeListSectionCmd(
	"events",
	"List events for a given tag name",
	`List JavaScript events fired by a given element by tag name.

You must specify the tag name using the --tag-name flag. The output includes each event,
it's type, and a summary

Examples:

  cem list event --tag-name my-button
  cem list event --tag-name my-button --format table --columns Type
  cem list event --tag-name my-button --format table --columns Type --columns Summary
`,
	"Events",
)

var listMethodsCmd = makeListSectionCmd(
	"methods",
	"List class methods for a given tag name",
	`List DOM object methods for the class registered to a given tag name.

You must specify the tag name using the --tag-name flag. The output includes each method, it's return type, privacy, and summary

Examples:

  cem list methods --tag-name my-button --format table
`,
	"Methods",
)

var listTagsCmd = &cobra.Command{
	Use:     "tags",
	Aliases: []string{"elements", "tag-names"},
	Short:   "List tag names in the custom elements manifest",
	Long: `List all custom element tag names

This command outputs a table with tag names and their corresponding source modules,
allowing you to quickly see which custom elements are available and where they are defined.

Example:

  cem list tags
  cem list tags --format table --columns Class
  cem list tags --format table --columns Class --columns Module
  cem list tags --format table --columns Class --columns Module --columns Summary
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if ctx, err := GetProjectContext(cmd); err != nil {
			return fmt.Errorf("project context not initialized: %w", err)
		} else {
			manifest, err := ctx.Manifest()
			if err != nil {
				return err
			}
			format, err := requireFormat(cmd, []string{"table"})
			if err != nil {
				return err
			}
			columns, err := cmd.Flags().GetStringArray("columns")
			if err != nil {
				return err
			}
			switch format {
			case "table":
				opts := list.RenderOptions{Columns: columns}
				if s, err := list.Render(M.NewRenderablePackage(manifest), opts); err != nil {
					return err
				} else {
					fmt.Println(s)
				}
			}
			return nil
		}
	},
}

var listModulesCmd = &cobra.Command{
	Use:     "modules",
	Aliases: []string{"files"},
	Short:   "List modules (javascript files) in the custom elements manifest",
	Long: `Lists all modules

This command outputs a table with module path names and the custom elements they register, if any.
Allows you to quickly see which modules are available and whether they are pure.

Example:

  cem list modules
  cem list modules --format table --columns Name
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if ctx, err := GetProjectContext(cmd); err != nil {
			return fmt.Errorf("project context not initialized: %w", err)
		} else {
			manifest, err := ctx.Manifest()
			if err != nil {
				return err
			}
			format, err := requireFormat(cmd, []string{"table"})
			if err != nil {
				return err
			}
			columns, err := cmd.Flags().GetStringArray("columns")
			if err != nil {
				return err
			}
			switch format {
			case "table":
				opts := list.RenderOptions{Columns: columns}
				if s, err := list.RenderModulesTable(manifest, opts); err != nil {
					return err
				} else {
					fmt.Println(s)
				}
			}
			return nil
		}
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List items in the custom elements manifest like tag names, attributes, functions, etc",
	Long: `List various entities defined in the custom elements manifest.

This command provides subcommands for listing tag names, attributes, and other elements described in your manifest.
Use the available subcommands to explore your project's custom elements, their attributes, and more.
Without subcommands, lists the entire custom elements manifest. Use the --deprecated flag to filter
the list, showing only itemas that should not be used.

Examples:

	cem list
	cem list --deprecated --format tree
  cem list tags
  cem list modules
  cem list attributes --tag-name my-button
  cem list slots --tag-name my-button
  cem list css-custom-properties --tag-name my-button
  cem list css-custom-states --tag-name my-button
  cem list css-parts --tag-name my-button
  cem list events --tag-name my-button
  cem list methods --tag-name my-button
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if ctx, err := GetProjectContext(cmd); err != nil {
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
			deprecated, err := cmd.Flags().GetBool("deprecated")
			if err != nil {
				return err
			}
			if deprecated && format != "tree" {
				return errors.New("--deprecated currently only supported with --format tree")
			}
			switch format {
			case "tree":
				title := "Manifest"
				pred := M.True
				if deprecated {
					title = "Deprecations"
					pred = M.IsDeprecated
				}
				if s, err := list.RenderTree(title, M.NewRenderablePackage(manifest), pred); err != nil {
					return err
				} else {
					fmt.Println(s)
				}
			case "table":
				opts := list.RenderOptions{}
				if s, err := list.Render(M.NewRenderablePackage(manifest), opts); err != nil {
					return err
				} else {
					fmt.Println(s)
				}
			}
			return nil
		}
	},
}

func init() {
	listCmd.AddCommand(listTagsCmd)
	listCmd.AddCommand(listModulesCmd)
	listCmd.PersistentFlags().StringP("format", "f", "table", "Output format")
	listCmd.PersistentFlags().Bool("deprecated", false, "Filter the results, showing only deprecated items")
	listTagsCmd.Flags().StringArrayP("columns", "c", []string{}, "list of columns to display in the table")
	listModulesCmd.Flags().StringArrayP("columns", "c", []string{}, "list of columns to display in the table")
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
		c.Flags().StringArrayP("columns", "c", []string{}, "list of columns to display in the table")
	}
	rootCmd.AddCommand(listCmd)
}
