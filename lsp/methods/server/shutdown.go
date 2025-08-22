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
	"fmt"

	"bennypowers.dev/cem/lsp/types"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// Shutdown handles the LSP shutdown request
func Shutdown(ctx types.ServerContext, context *glsp.Context) error {
	// Send shutdown message via LSP protocol
	context.Notify(protocol.ServerWindowShowMessage, &protocol.ShowMessageParams{
		Type:    protocol.MessageTypeInfo,
		Message: "CEM LSP Server shutting down...",
	})

	dm, _ := ctx.DocumentManager()
	if dm != nil {
		dm.Close()
	}

	if ws := ctx.Workspace(); ws != nil {
		if err := ws.Cleanup(); err != nil {
			context.Notify(protocol.ServerWindowShowMessage, &protocol.ShowMessageParams{
				Type:    protocol.MessageTypeWarning,
				Message: fmt.Sprintf("Warning: error during workspace cleanup: %v", err),
			})
		}
	}

	return nil
}
