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
	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// SetTrace handles the $/setTrace notification
func SetTrace(ctx types.ServerContext, context *glsp.Context, params *protocol.SetTraceParams) error {
	// Use LSP standard trace levels to control debug logging
	switch params.Value {
	case protocol.TraceValueOff:
		// Log before disabling to ensure message is emitted
		ctx.DebugLog("Disabling debug logging via $/setTrace")
		helpers.SetDebugLoggingEnabled(false)
	case protocol.TraceValueMessage:
		// Enable basic debug logging
		helpers.SetDebugLoggingEnabled(true)
		ctx.DebugLog("Basic debug logging enabled via $/setTrace")
	case protocol.TraceValueVerbose:
		// Enable verbose debug logging (same as messages for now)
		helpers.SetDebugLoggingEnabled(true)
		ctx.DebugLog("Verbose debug logging enabled via $/setTrace")
	default:
		// Unknown/unsupported value; ignore gracefully
		ctx.DebugLog("Ignoring unknown $/setTrace value")
	}

	return nil
}
