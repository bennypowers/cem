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
	viper.Unmarshal(&cfg)
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
	Long: `TODO: say more`,
}

var listTagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "List tag names in the custom elements manifest",
	Long: `TODO: say more`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := C.CemConfig{}
		viper.Unmarshal(&cfg)
		path := cfg.Generate.Output
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		path = filepath.Join(cwd, path)
		json, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		pkg, err := M.UnmarshalPackage(json)
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
	Short: "List attributes in the custom elements manifest by tag name",
	Long: `TODO: say more`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tagName, err := cmd.Flags().GetString("tag-name")
		if tagName == "" {
			return errors.New("cem list attrs: must supply --tag-name")
		}
		if err != nil {
			return err
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
					func() string { if r.CustomElementField.Reflects {return "✅" } else {return "❌"}}()}
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
