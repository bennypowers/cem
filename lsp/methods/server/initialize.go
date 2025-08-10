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
package server

import (
	"log"

	"bennypowers.dev/cem/internal/version"
	"bennypowers.dev/cem/lsp/helpers"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// Initialize handles the LSP initialize request
func Initialize(ctx ServerContext, context *glsp.Context, params *protocol.InitializeParams) (any, error) {
	log.Printf("CEM LSP Server initializing...")

	// Set up debug logging (disabled by default, enabled via $/setTrace)
	helpers.SetDebugLogger(func(format string, args ...any) {
		ctx.DebugLog(format, args...)
	})

	// Define server capabilities
	openClose := true
	changeKind := protocol.TextDocumentSyncKindIncremental
	serverVersion := version.GetVersion()

	capabilities := protocol.ServerCapabilities{
		HoverProvider: &protocol.HoverOptions{},
		CompletionProvider: &protocol.CompletionOptions{
			TriggerCharacters: []string{
				"<",
				" ",
				"=",
				"\"",
				"@",
				".",
				"?",
			},
		},
		DefinitionProvider: &protocol.DefinitionOptions{},
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
