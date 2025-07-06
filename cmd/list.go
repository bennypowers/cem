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
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"slices"

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

func readPkg() (pkg *M.Package, err error) {
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
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &pkg); err != nil {
		return pkg, err
	}
	pterm.Debug.Printfln("Loaded manifest from %s", path)
	return pkg, err
}

// validateTagCommandFlags returns the tag name, format, and loaded manifest package.
// It errors if --columns is passed but format is not "table".
func validateTagCommandFlags(cmd *cobra.Command) (tagName string, format string, columns []string, pkg *M.Package, err error) {
	tagName, err = requireTagName(cmd)
	if err != nil {
		return "", "", nil, nil, err
	}
	pkg, err = readPkg()
	if err != nil {
		return "", "", nil, nil, err
	}
	format, err = requireFormat(cmd, []string{"table"})
	if err != nil {
		return "", "", nil, nil, err
	}
	columns, err = cmd.Flags().GetStringArray("columns")
	if err != nil {
		return "", "", nil, nil, err
	}
	if len(columns) > 0 && format != "table" {
		return "", "", nil, nil, errors.New("--columns flag can only be used with --format table")
	}
	return tagName, format, columns, pkg, nil
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
		tagName, format, columns, pkg, err := validateTagCommandFlags(cmd)
		if err != nil {
			return err
		}
		attrs, err := pkg.TagRenderableAttributes(tagName)
		if err != nil {
			return err
		}
		switch format {
		case "table":
			headers := []string{"Name", "DOM Property", "Reflects", "Summary"}
			rows := list.MapToTableRows(attrs)
			title := "Attributes on " + tagName
			return list.RenderTable(title, headers, rows, columns)
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

  cem list slots --tag-name my-button
  cem list slots --tag-name my-button --format table --columns Name
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagName, format, columns, pkg, err := validateTagCommandFlags(cmd)
		if err != nil {
			return err
		}
		slots, err := pkg.TagRenderableSlots(tagName)
		if err != nil {
			return err
		}
		switch format {
		case "table":
			headers := []string{"Name", "Summary"}
			rows := list.MapToTableRows(slots)
			title := "Slots on " + tagName
			return list.RenderTable(title, headers, rows, columns)
		}

		return nil
	},
}

var listCssCustomPropertiesCmd = &cobra.Command{
	Use:     "css-custom-properties",
	Aliases: []string{"css-properties", "css-props", "css-custom-props"},
	Short:   "List CSS custom properties used in a given tag name",
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
		tagName, format, columns, pkg, err := validateTagCommandFlags(cmd)
		if err != nil {
			return err
		}
		props, err := pkg.TagRenderableCssProperties(tagName)
		if err != nil {
			return err
		}
		switch format {
		case "table":
			headers := []string{"Name", "Syntax", "Default", "Summary"}
			rows := list.MapToTableRows(props)
			title := "CSS Custom Properties for " + tagName
			return list.RenderTable(title, headers, rows, columns)
		}
		return nil
	},
}

var listCssCustomStatesCmd = &cobra.Command{
	Use:     "css-custom-states",
	Aliases: []string{"css-states"},
	Short:   "List CSS custom states used in a given tag name",
	Long: `List CSS custom states used in a given tag name.

You must specify the tag name using the --tag-name flag. The output includes each css state by name,
and a summary

Examples:

  cem list css-custom-states --tag-name my-button
  cem list css-custom-states --tag-name my-button --format table --columns Name
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagName, format, columns, pkg, err := validateTagCommandFlags(cmd)
		if err != nil {
			return err
		}
		props, err := pkg.TagRenderableCssStates(tagName)
		if err != nil {
			return err
		}
		switch format {
		case "table":
			headers := []string{"Name", "Summary"}
			rows := list.MapToTableRows(props)
			title := "CSS Custom States for " + tagName
			return list.RenderTable(title, headers, rows, columns)
		}

		return nil
	},
}

var listCssPartsCmd = &cobra.Command{
	Use:     "css-parts",
	Aliases: []string{"parts", "css-shadow-parts"},
	Short:   "List CSS shadow parts for a given tag name",
	Long: `List CSS shadow parts for a given tag name.

You must specify the tag name using the --tag-name flag. The output includes each shadow part by name, and a summary

Examples:

  cem list css-parts --tag-name my-button --format table
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagName, format, columns, pkg, err := validateTagCommandFlags(cmd)
		if err != nil {
			return err
		}
		parts, err := pkg.TagRenderableCssParts(tagName)
		if err != nil {
			return err
		}
		switch format {
		case "table":
			headers := []string{"Name", "Summary"}
			rows := list.MapToTableRows(parts)
			title := "CSS Shadow Parts for " + tagName
			return list.RenderTable(title, headers, rows, columns)
		}
		return nil
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
		tagName, format, columns, pkg, err := validateTagCommandFlags(cmd)
		if err != nil {
			return err
		}
		events, err := pkg.TagRenderableEvents(tagName)
		if err != nil {
			return err
		}
		switch format {
		case "table":
			headers := []string{"Name", "Type", "Summary"}
			rows := list.MapToTableRows(events)
			title := "Events fired by " + tagName
			return list.RenderTable(title, headers, rows, columns)
		}
		return nil
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
		tagName, format, columns, pkg, err := validateTagCommandFlags(cmd)
		if err != nil {
			return err
		}
		methods, err := pkg.TagRenderableMethods(tagName)
		if err != nil {
			return err
		}
		switch format {
		case "table":
			headers := []string{"Name", "Return Type", "Privacy", "Static", "Summary"}
			rows := list.MapToTableRows(methods)
			title := "Methods on " + tagName
			return list.RenderTable(title, headers, rows, columns)
		}

		return nil
	},
}

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
		pkg, err := readPkg()
		if err != nil {
			return err
		}
		tags := pkg.RenderableCustomElementDeclarations()
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
			headers := []string{"Tag", "Class", "Module", "Summary"}
			rows := list.MapToTableRows(tags)
			title := "Tags"
			return list.RenderTable(title, headers, rows, columns)
		}
		return nil
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
		pkg, err := readPkg()
		if err != nil {
			return err
		}
		tags := pkg.RenderableModules()
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
		pkg, err := readPkg()
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
			return list.RenderTree(title, M.NewRenderablePackage(pkg), pred)
		}
		return nil
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
