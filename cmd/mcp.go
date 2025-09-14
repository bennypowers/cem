/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

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
	"fmt"
	"os"

	MCP "bennypowers.dev/cem/mcp"
	"bennypowers.dev/cem/types"
	"bennypowers.dev/cem/workspace"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// mcpCmd represents the mcp command
var mcpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "Launch MCP server for custom elements",
	Long: `Launch a Model Context Protocol (MCP) server that provides intelligent context
about custom elements for AI systems, enabling smart HTML generation and component usage.

The server provides context from custom-elements manifests including:
- Local package.json "customElements" reference
- Locally installed packages (via their package.json "customElements" field)
- Config-specified manifests
- Globally-configured manifests

Resources provided:
- JSON schema for custom elements manifests
- Live manifest registry with real-time updates
- Individual element definitions with usage patterns
- CSS integration guides (custom properties, parts, states)

Tools provided:
- Element validation and HTML generation assistance
- Attribute and slot usage suggestions
- CSS integration guidance
- Registry querying and component discovery`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// CRITICAL: Redirect all pterm output to stderr immediately to prevent MCP stdout contamination
		pterm.SetDefaultOutput(os.Stderr)

		// Bind flags to viper
		if err := viper.BindPFlag("mcp.maxDescriptionLength", cmd.Flags().Lookup("max-description-length")); err != nil {
			return fmt.Errorf("failed to bind max-description-length flag: %w", err)
		}

		ctx := cmd.Context()
		wctx := ctx.Value(workspace.WorkspaceContextKey).(types.WorkspaceContext)

		// Get max description length from config/flag (flag overrides config, config overrides default)
		maxDescriptionLength := viper.GetInt("mcp.maxDescriptionLength")
		if maxDescriptionLength == 0 {
			maxDescriptionLength = 2000 // default
		}

		// Create and run the MCP server with stdio transport
		server, err := MCP.NewServerWithConfig(wctx, MCP.ServerConfig{
			MaxDescriptionLength: maxDescriptionLength,
		})
		if err != nil {
			return err
		}

		return server.Run(ctx)
	},
}

func init() {
	mcpCmd.Flags().IntP("max-description-length", "", 2000, "Maximum length for description fields before truncation")
	rootCmd.AddCommand(mcpCmd)
}
