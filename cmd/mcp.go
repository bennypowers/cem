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
	W "bennypowers.dev/cem/workspace"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
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
- Import path resolution
- Component compatibility checking`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// CRITICAL: Redirect all pterm output to stderr immediately to prevent MCP stdout contamination
		pterm.SetDefaultOutput(os.Stderr)

		ctx := cmd.Context()
		wctx := ctx.Value(W.WorkspaceContextKey).(W.WorkspaceContext)

		// Determine transport based on boolean flags
		transport := MCP.TransportStdio // default

		stdioFlag, _ := cmd.Flags().GetBool("stdio")
		tcpFlag, _ := cmd.Flags().GetBool("tcp")
		websocketFlag, _ := cmd.Flags().GetBool("websocket")

		// Check which transport flag is set
		flagCount := 0
		if stdioFlag {
			transport = MCP.TransportStdio
			flagCount++
		}
		if tcpFlag {
			transport = MCP.TransportTCP
			flagCount++
		}
		if websocketFlag {
			transport = MCP.TransportWebSocket
			flagCount++
		}

		// Ensure only one transport flag is set
		if flagCount > 1 {
			return fmt.Errorf("only one transport flag may be specified")
		}

		// Create and run the MCP server
		server, err := MCP.NewSimpleCEMServer(wctx)
		if err != nil {
			return err
		}

		// For now, only stdio transport is implemented
		if transport != MCP.TransportStdio {
			return fmt.Errorf("only stdio transport is currently implemented")
		}

		return server.Run(ctx)
	},
}

func init() {
	rootCmd.AddCommand(mcpCmd)
	mcpCmd.Flags().Bool("stdio", false, "Use stdio transport (default)")
	mcpCmd.Flags().Bool("tcp", false, "Use TCP transport")
	mcpCmd.Flags().Bool("websocket", false, "Use WebSocket transport")
}
