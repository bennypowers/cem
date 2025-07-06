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

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	C "bennypowers.dev/cem/cmd/config"
	"bennypowers.dev/cem/list"
	M "bennypowers.dev/cem/manifest"
)

func readCfg() (*C.CemConfig, error) {
	cfg := C.CemConfig{}
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
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
	case "table", "tree":
		return format, nil
	// Add more supported formats here as the CLI grows
	default:
		return "", errors.New("unknown format: " + format + ". Supported formats: table, tree")
	}
}

func readPkg() (*M.Package, error) {
	cfg, err := readCfg()
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
	pterm.Debug.Printfln("Loaded manifest from %s", path)
	return M.UnmarshalPackage(json)
}

// validateTagCommandFlags returns the tag name, format, columns, deprecated filter, and loaded manifest package.
// It errors if --columns is passed but format is not "table".
func validateTagCommandFlags(cmd *cobra.Command) (tagName string, format string, columns []string, showDeprecated *bool, pkg *M.Package, err error) {
	tagName, err = requireTagName(cmd)
	if err != nil {
		return "", "", nil, nil, nil, err
	}
	pkg, err = readPkg()
	if err != nil {
		return "", "", nil, nil, nil, err
	}
	format, err = requireFormat(cmd)
	if err != nil {
		return "", "", nil, nil, nil, err
	}
	columns, err = cmd.Flags().GetStringArray("columns")
	if err != nil {
		return "", "", nil, nil, nil, err
	}
	if len(columns) > 0 && format != "table" {
		return "", "", nil, nil, nil, errors.New("--columns flag can only be used with --format table")
	}

	// Handle deprecated flag
	deprecatedFlag, err := cmd.Flags().GetString("deprecated")
	if err != nil {
		return "", "", nil, nil, nil, err
	}
	if deprecatedFlag != "" {
		switch deprecatedFlag {
		case "only":
			showDeprecated = &[]bool{true}[0]
		case "exclude":
			showDeprecated = &[]bool{false}[0]
		case "all":
			showDeprecated = nil
		default:
			return "", "", nil, nil, nil, errors.New("--deprecated must be one of: only, exclude, all")
		}
	}

	return tagName, format, columns, showDeprecated, pkg, nil
}

var listAttrsCmd = &cobra.Command{
	Use:     "attributes",
	Aliases: []string{"attrs"},
	Short:   "List attributes in the custom elements manifest by tag name",
	Long: `List all attributes for a given custom element tag

You must specify the tag name using the --tag-name flag. The output includes each attribute,
its corresponding DOM property, whether it reflects changes to the DOM, and a summary.

Examples:

  cem list attributes --tag-name my-button
  cem list attributes --tag-name my-button --format table --columns "DOM Property"
  cem list attributes --tag-name my-button --format table --columns "DOM Property" --columns Reflects
  cem list attributes --tag-name my-button --format table --columns "DOM Property" --columns Reflects --columns Summary
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagName, format, columns, showDeprecated, pkg, err := validateTagCommandFlags(cmd)
		if err != nil {
			return err
		}
		attrs, err := pkg.GetTagAttrsWithContext(tagName)
		if err != nil {
			return err
		}
		headers := []string{"Name", "DOM Property", "Reflects", "Summary"}
		title := "Attributes on " + tagName
		return list.RenderOutput(title, headers, attrs, columns, format, showDeprecated)
	},
}

var listSlotsCmd = &cobra.Command{
	Use:   "slots",
	Short: "List slots in the custom elements manifest by tag name",
	Long: `List all slots for a given custom element tag.

You must specify the tag name using the --tag-name flag. The output includes each slot name, and it's summary

Examples:

  cem list slots --tag-name my-button
  cem list slots --tag-name my-button --format table --columns Name
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagName, format, columns, showDeprecated, pkg, err := validateTagCommandFlags(cmd)
		if err != nil { return err }
		slots, err := pkg.GetTagSlotsWithContext(tagName)
		if err != nil {
			return err
		}
		headers := []string{"Name", "Summary"}
		title := "Slots on "+tagName
		return list.RenderOutput(title, headers, slots, columns, format, showDeprecated)
	},
}

var listCssCustomPropertiesCmd = &cobra.Command{
	Use:   "css-custom-properties",
	Aliases: []string{"css-properties", "css-props", "css-custom-props"},
	Short: "List CSS custom properties used in a given tag name",
	Long: `List CSS custom properties used in a given tag name.

You must specify the tag name using the --tag-name flag. The output may include for each
css property by name, it's syntax, default value, and summary.

Examples:

  cem list css-custom-properties --tag-name my-button
  cem list css-custom-properties --tag-name my-button --format table --columns Syntax
  cem list css-custom-properties --tag-name my-button --format table --columns Syntax --columns Default
  cem list css-custom-properties --tag-name my-button --format table --columns Syntax --columns Default --columns Summary
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagName, format, columns, showDeprecated, pkg, err := validateTagCommandFlags(cmd)
		if err != nil { return err }
		props, err := pkg.GetTagCssPropertiesWithContext(tagName)
		if err != nil { return err }
		headers := []string{"Name", "Syntax", "Default", "Summary"}
		title := "CSS Custom Properties for " + tagName
		return list.RenderOutput(title, headers, props, columns, format, showDeprecated)
	},
}

var listCssCustomStatesCmd = &cobra.Command{
	Use:   "css-custom-states",
	Aliases: []string{"css-states"},
	Short: "List CSS custom states used in a given tag name",
	Long: `List CSS custom states used in a given tag name.

You must specify the tag name using the --tag-name flag. The output includes each css state by name,
and a summary

Examples:

  cem list css-custom-states --tag-name my-button
  cem list css-custom-states --tag-name my-button --format table --columns Name
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagName, format, columns, showDeprecated, pkg, err := validateTagCommandFlags(cmd)
		if err != nil { return err }
		props, err := pkg.GetTagCssStatesWithContext(tagName)
		if err != nil { return err }
		headers := []string{"Name", "Summary"}
		title := "CSS Custom States for " + tagName
		return list.RenderOutput(title, headers, props, columns, format, showDeprecated)
	},
}

var listCssPartsCmd = &cobra.Command{
	Use:   "css-parts",
	Aliases: []string{"parts","css-shadow-parts"},
	Short: "List CSS shadow parts for a given tag name",
	Long: `List CSS shadow parts for a given tag name.

You must specify the tag name using the --tag-name flag. The output includes each shadow part by name, and a summary

Examples:

  cem list css-parts --tag-name my-button --format table
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagName, format, columns, showDeprecated, pkg, err := validateTagCommandFlags(cmd)
		if err != nil { return err }
		parts, err := pkg.GetTagCssPartsWithContext(tagName)
		if err != nil { return err }
		headers := []string{"Name", "Summary"}
		title := "CSS Shadow Parts for " + tagName
		return list.RenderOutput(title, headers, parts, columns, format, showDeprecated)
	},
}

var listEventsCmd = &cobra.Command{
	Use:   "events",
	Short: "List events for a given tag name",
	Long: `List JavaScript events fired by a given element by tag name.

You must specify the tag name using the --tag-name flag. The output includes each event,
it's type, and a summary

Examples:

  cem list event --tag-name my-button
  cem list event --tag-name my-button --format table --columns Type
  cem list event --tag-name my-button --format table --columns Type --columns Summary
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagName, format, columns, showDeprecated, pkg, err := validateTagCommandFlags(cmd)
		if err != nil { return err }
		events, err := pkg.GetTagEventsWithContext(tagName)
		if err != nil { return err }
		headers := []string{"Name", "Type", "Summary"}
		title := "Events fired by " + tagName
		return list.RenderOutput(title, headers, events, columns, format, showDeprecated)
	},
}

var listMethodsCmd = &cobra.Command{
	Use:   "methods",
	Short: "List class methods for a given tag name",
	Long: `List DOM object methods for the class registered to a given tag name.

You must specify the tag name using the --tag-name flag. The output includes each method, it's return type, privacy, and summary

Examples:

  cem list methods --tag-name my-button --format table
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagName, format, columns, showDeprecated, pkg, err := validateTagCommandFlags(cmd)
		if err != nil { return err }
		methods, err := pkg.GetTagMethodsWithContext(tagName)
		if err != nil { return err }
		headers := []string{"Name", "Return Type", "Privacy", "Static", "Summary"}
		title := "Methods on " + tagName
		return list.RenderOutput(title, headers, methods, columns, format, showDeprecated)
	},
}

var listMembersCmd = &cobra.Command{
	Use:   "members",
	Short: "List all class members for a given tag name",
	Long: `List all class members (attributes, fields, methods, events, slots) for a given tag name grouped by type.

You must specify the tag name using the --tag-name flag. This command is particularly useful with the --tree format
to see all members organized by type, and with the --deprecated flag to filter deprecated members.

Examples:

  cem list members --tag-name my-button --tree
  cem list members --tag-name my-button --tree --deprecated only
  cem list members --tag-name my-button --tree --deprecated exclude
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagName, format, _, showDeprecated, pkg, err := validateTagCommandFlags(cmd)
		if err != nil { return err }
		
		if format == "table" {
			return errors.New("members command only supports --tree format. Use individual commands (attributes, methods, etc.) for table format")
		}

		allMembers, err := pkg.GetTagAllMembersWithContext(tagName)
		if err != nil { return err }
		
		title := "All Members for " + tagName
		return list.RenderTreeGrouped(title, allMembers, showDeprecated)
	},
}

var listFieldsCmd = &cobra.Command{
	Use:   "fields",
	Aliases: []string{"properties"},
	Short: "List class fields for a given tag name",
	Long: `List class fields (properties) for the class registered to a given tag name.

You must specify the tag name using the --tag-name flag. The output includes each field, 
its type, privacy, whether it's static, and a summary.

Examples:

  cem list fields --tag-name my-button --format table
  cem list fields --tag-name my-button --tree
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagName, format, columns, showDeprecated, pkg, err := validateTagCommandFlags(cmd)
		if err != nil { return err }
		fields, err := pkg.GetTagFieldsWithContext(tagName)
		if err != nil { return err }
		headers := []string{"Name", "Type", "Privacy", "Static", "Summary"}
		title := "Fields on " + tagName
		return list.RenderOutput(title, headers, fields, columns, format, showDeprecated)
	},
}

var listTagsCmd = &cobra.Command{
	Use:   "tags",
	Aliases: []string{"elements", "tag-names"},
	Short: "List tag names in the custom elements manifest",
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
		pkg, err := readPkg()
		if err != nil {
			return err
		}
		tags := pkg.GetAllTagNamesWithContext()
		format, err := requireFormat(cmd)
		if err != nil {
			return err
		}
		columns, err := cmd.Flags().GetStringArray("columns")
		if err != nil {
			return err
		}
		switch format {
		case "table":
			headers := []string{"Tag", "Class", "Module", "Summary"}
			rows := list.MapToTableRows(tags)
			title := "Tags"
			return list.RenderTable(title, headers, rows, columns)
		}
		return nil
	},
}

var listModulesCmd = &cobra.Command{
	Use:   "modules",
	Aliases: []string{"files"},
	Short: "List modules (javascript files) in the custom elements manifest",
	Long: `Lists all modules

This command outputs a table with module path names and the custom elements they register, if any.
Allows you to quickly see which modules are available and whether they are pure.

Example:

  cem list modules
  cem list modules --format table --columns Name
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		pkg, err := readPkg()
		if err != nil {
			return err
		}
		tags := pkg.GetAllModulesWithContext()
		format, err := requireFormat(cmd)
		if err != nil {
			return err
		}
		columns, err := cmd.Flags().GetStringArray("columns")
		if err != nil {
			return err
		}
		switch format {
		case "table":
			headers := []string{"Name", "Custom Elements"}
			rows := list.MapToTableRows(tags)
			title := "Modules"
			return list.RenderTable(title, headers, rows, columns)
		}
		return nil
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List items in the custom elements manifest like tag names, attributes, functions, etc",
	Long: `List various entities defined in the custom elements manifest.

This command provides subcommands for listing tag names, attributes, and other elements described in your manifest.
Use the available subcommands to explore your project's custom elements, their attributes, and more.

Use --tree format for a hierarchical view grouped by member type.
Use --deprecated flag to filter deprecated members (only, exclude, or all).

Examples:

  cem list tags
  cem list modules
  cem list attributes --tag-name my-button
  cem list attributes --tag-name my-button --tree
  cem list attributes --tag-name my-button --deprecated only
  cem list members --tag-name my-button --tree
  cem list members --tag-name my-button --tree --deprecated exclude
  cem list slots --tag-name my-button
  cem list css-custom-properties --tag-name my-button
  cem list css-custom-states --tag-name my-button
  cem list css-parts --tag-name my-button
  cem list events --tag-name my-button
  cem list methods --tag-name my-button
  cem list fields --tag-name my-button
`,
}

func init() {
	listCmd.AddCommand(listTagsCmd)
	listCmd.AddCommand(listModulesCmd)
	listCmd.PersistentFlags().StringP("format", "f", "table", "Output format (table, tree)")
	listTagsCmd.Flags().StringArrayP("columns", "c", []string{}, "list of columns to display in the table")
	listModulesCmd.Flags().StringArrayP("columns", "c", []string{}, "list of columns to display in the table")
	for _, c := range []*cobra.Command{
		listAttrsCmd,
		listFieldsCmd,
		listSlotsCmd,
		listCssCustomPropertiesCmd,
		listCssCustomStatesCmd,
		listCssPartsCmd,
		listEventsCmd,
		listMethodsCmd,
		listMembersCmd,
	} {
		listCmd.AddCommand(c)
		c.Flags().StringP("tag-name", "t", "", "Tag name to list attributes for")
		c.Flags().StringArrayP("columns", "c", []string{}, "list of columns to display in the table")
		c.Flags().StringP("deprecated", "d", "", "Filter by deprecated status: 'only', 'exclude', or 'all' (default: all)")
	}
	rootCmd.AddCommand(listCmd)
}
