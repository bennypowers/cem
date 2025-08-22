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

	LSP "bennypowers.dev/cem/lsp"
	W "bennypowers.dev/cem/workspace"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// lspCmd represents the lsp command
var lspCmd = &cobra.Command{
	Use:   "lsp",
	Short: "Launch LSP server for custom elements",
	Long: `Launch a Language Server Protocol (LSP) server that provides IDE features
for custom elements in HTML files and TypeScript template literals.

The server reads custom-elements manifests from:
- Local package.json "customElements" reference
- Locally installed packages (via their package.json "customElements" field)
- Config-specified manifests
- Globally-configured manifests

Features provided:
- Hover information for custom element tags and attributes
- Autocomplete for tag names and attribute names/values
- Go-to-definition for custom elements
- Diagnostics for unknown elements and invalid attributes`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// CRITICAL: Redirect all pterm output to stderr immediately to prevent LSP stdout contamination
		pterm.SetDefaultOutput(os.Stderr)
		
		ctx := cmd.Context()
		wctx := ctx.Value(W.WorkspaceContextKey).(W.WorkspaceContext)

		// Determine transport based on boolean flags
		var transport LSP.TransportKind = LSP.TransportStdio // default

		stdioFlag, _ := cmd.Flags().GetBool("stdio")
		tcpFlag, _ := cmd.Flags().GetBool("tcp")
		websocketFlag, _ := cmd.Flags().GetBool("websocket")
		nodejsFlag, _ := cmd.Flags().GetBool("nodejs")

		// Check which transport flag is set
		flagCount := 0
		if stdioFlag {
			transport = LSP.TransportStdio
			flagCount++
		}
		if tcpFlag {
			transport = LSP.TransportTCP
			flagCount++
		}
		if websocketFlag {
			transport = LSP.TransportWebSocket
			flagCount++
		}
		if nodejsFlag {
			transport = LSP.TransportNodeJS
			flagCount++
		}

		// Ensure only one transport flag is set
		if flagCount > 1 {
			return fmt.Errorf("only one transport flag may be specified")
		}

		// Create and run the LSP server with debug mode enabled for stdio transport
		server, err := LSP.NewServer(wctx, transport)
		if err != nil {
			return err
		}
		return server.Run()
	},
}

func init() {
	rootCmd.AddCommand(lspCmd)
	lspCmd.Flags().Bool("stdio", false, "Use stdio transport (default)")
	lspCmd.Flags().Bool("tcp", false, "Use TCP transport")
	lspCmd.Flags().Bool("websocket", false, "Use WebSocket transport")
	lspCmd.Flags().Bool("nodejs", false, "Use Node.js transport")
}
