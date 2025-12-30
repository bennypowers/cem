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
package lifecycle

import (
	"os"

	"bennypowers.dev/cem/internal/logging"
	"bennypowers.dev/cem/internal/version"
	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// Initialize handles the LSP initialize request
func Initialize(ctx types.ServerContext, context *glsp.Context, params *protocol.InitializeParams) (any, error) {
	// Configure centralized logger for LSP mode
	logging.SetLSPContext(context)

	// Send initialization message with version
	serverVersion := version.GetVersion()
	logging.Info("CEM LSP Server initializing (version %s)...", serverVersion)

	// Debug info about binary location
	if executable, err := os.Executable(); err == nil {
		logging.Debug("LSP server binary: %s", executable)
	}

	// Log workspace information from LSP client
	if params.RootURI != nil {
		helpers.SafeDebugLog("[INITIALIZE] LSP client provided root URI: %s", *params.RootURI)
	}
	if len(params.WorkspaceFolders) > 0 {
		for i, folder := range params.WorkspaceFolders {
			helpers.SafeDebugLog("[INITIALIZE] LSP client workspace folder %d: %s (%s)", i, folder.URI, folder.Name)
		}
	}

	// Update workspace context with information from LSP client
	if err := ctx.UpdateWorkspaceFromLSP(params.RootURI, params.WorkspaceFolders); err != nil {
		helpers.SafeDebugLog("[INITIALIZE] Warning: Failed to update workspace from LSP parameters: %v", err)
		// Don't fail initialization, just log the warning
	}

	// Define server capabilities
	openClose := true
	changeKind := protocol.TextDocumentSyncKindIncremental

	capabilities := protocol.ServerCapabilities{
		HoverProvider: &protocol.HoverOptions{},
		CompletionProvider: &protocol.CompletionOptions{
			TriggerCharacters: []string{
				"<",
				"=",
				"\"",
				"@",
				".",
				"?",
			},
			ResolveProvider: &[]bool{true}[0],
		},
		DefinitionProvider: &protocol.DefinitionOptions{},
		ReferencesProvider: &protocol.ReferenceOptions{},
		CodeActionProvider: &protocol.CodeActionOptions{
			CodeActionKinds: []protocol.CodeActionKind{
				protocol.CodeActionKindQuickFix,
			},
		},
		WorkspaceSymbolProvider: &protocol.WorkspaceSymbolOptions{},
		TextDocumentSync: &protocol.TextDocumentSyncOptions{
			OpenClose: &openClose,
			Change:    &changeKind,
		},
	}

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    "cem-lsp",
			Version: &serverVersion,
		},
	}, nil
}
