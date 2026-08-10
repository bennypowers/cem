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
	"os"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
	"github.com/spf13/cobra"
	"golang.org/x/term"

	M "bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/search"
	W "bennypowers.dev/cem/internal/workspace"
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
  cem search -i                        # Interactive fuzzy search TUI
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("accepts at most 1 argument")
		}
		interactive, _ := cmd.Flags().GetBool("interactive")
		if !interactive && len(args) == 0 {
			return errors.New("requires a pattern argument (or use -i for interactive mode)")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		interactive, _ := cmd.Flags().GetBool("interactive")

		fsys := platform.NewOSFileSystem()
		if interactive {
			if !term.IsTerminal(int(os.Stdin.Fd())) || !term.IsTerminal(int(os.Stdout.Fd())) {
				return errors.New("interactive mode requires a terminal; omit -i to use non-interactive search")
			}
			var initialQuery string
			if len(args) > 0 {
				initialQuery = args[0]
			}
			if W.ShouldUseWorkspaceMode(cmd, fsys) {
				return searchInteractiveWorkspace(cmd, initialQuery)
			}
			ctx, err := W.GetWorkspaceContext(cmd)
			if err != nil {
				return fmt.Errorf("project context not initialized: %w", err)
			}
			manifest, err := ctx.Manifest()
			if err != nil {
				return err
			}
			roots := []M.Renderable{M.NewRenderablePackage(manifest)}
			model := search.NewInteractiveModel(roots)
			if initialQuery != "" {
				model = model.WithInitialQuery(initialQuery)
			}
			p := tea.NewProgram(model)
			if _, err := p.Run(); err != nil {
				return err
			}
			return nil
		}

		pattern := args[0]
		if pattern == "" {
			return errors.New("search pattern cannot be empty")
		}

		if W.ShouldUseWorkspaceMode(cmd, fsys) {
			return searchWorkspace(cmd, pattern)
		}

		ctx, err := W.GetWorkspaceContext(cmd)
		if err != nil {
			return fmt.Errorf("project context not initialized: %w", err)
		}
		manifest, err := ctx.Manifest()
		if err != nil {
			return err
		}

		format, err := requireFormat(cmd, []string{"table", "tree"})
		if err != nil {
			return err
		}

		s, err := search.RenderSearchResults(manifest, pattern, format)
		if err != nil {
			return err
		}
		if _, err := lipgloss.Fprintln(cmd.OutOrStdout(), s); err != nil {
			return err
		}
		return nil
	},
}

func searchInteractiveWorkspace(cmd *cobra.Command, initialQuery string) error {
	ctx, err := W.GetWorkspaceContext(cmd)
	if err != nil {
		return err
	}

	fsys := platform.NewOSFileSystem()
	var roots []M.Renderable
	results := W.ForEachPackage(ctx.Root(), fsys, func(pkg W.PackageInfo) error {
		pkgCtx := W.NewFileSystemWorkspaceContext(pkg.Path)
		if err := pkgCtx.Init(); err != nil {
			return err
		}
		manifest, err := pkgCtx.Manifest()
		if err != nil {
			return err
		}
		if manifest == nil {
			return nil
		}
		roots = append(roots, M.NewRenderablePackage(manifest))
		return nil
	})

	reportErr := W.ReportResults("Loaded manifests", results)

	if len(roots) == 0 {
		if reportErr != nil {
			return reportErr
		}
		return errors.New("no manifests found in workspace")
	}

	model := search.NewInteractiveModel(roots)
	if initialQuery != "" {
		model = model.WithInitialQuery(initialQuery)
	}
	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		return err
	}
	return reportErr
}

func searchWorkspace(cmd *cobra.Command, pattern string) error {
	ctx, err := W.GetWorkspaceContext(cmd)
	if err != nil {
		return err
	}

	format, err := requireFormat(cmd, []string{"table", "tree"})
	if err != nil {
		return err
	}

	fsys := platform.NewOSFileSystem()
	results := W.ForEachPackage(ctx.Root(), fsys, func(pkg W.PackageInfo) error {
		pkgCtx := W.NewFileSystemWorkspaceContext(pkg.Path)
		if err := pkgCtx.Init(); err != nil {
			return err
		}
		manifest, err := pkgCtx.Manifest()
		if err != nil {
			return err
		}
		if manifest == nil {
			return nil
		}
		s, err := search.RenderSearchResults(manifest, pattern, format)
		if err != nil {
			return err
		}
		if s != "" {
			if _, err := lipgloss.Fprintf(cmd.OutOrStdout(), "\n%s:\n%s\n", pkg.Name, s); err != nil {
				return err
			}
		}
		return nil
	})

	return W.ReportResults("Searched manifests", results)
}

func init() {
	searchCmd.Flags().StringP("format", "f", "table", "Output format (table or tree)")
	searchCmd.Flags().BoolP("interactive", "i", false, "Launch interactive fuzzy search TUI (requires a terminal)")
	searchCmd.MarkFlagsMutuallyExclusive("interactive", "format")
	rootCmd.AddCommand(searchCmd)
}
