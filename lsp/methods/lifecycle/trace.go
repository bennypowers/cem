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
	"bennypowers.dev/cem/internal/logging"
	"bennypowers.dev/cem/lsp/types"
	"go.lsp.dev/protocol"
)

// SetTrace handles the $/setTrace notification
func SetTrace(ctx types.ServerContext, params *protocol.SetTraceParams) error {
	switch params.Value {
	case protocol.TraceValueOff:
		ctx.DebugLog("Disabling debug logging via $/setTrace")
		logging.SetVerbosity(logging.VerbosityNormal)
	case protocol.TraceValueMessages:
		logging.SetVerbosity(logging.VerbosityDebug)
		ctx.DebugLog("Debug logging enabled via $/setTrace")
	case protocol.TraceValueVerbose:
		logging.SetVerbosity(logging.VerbosityTrace)
		ctx.DebugLog("Trace logging enabled via $/setTrace")
	default:
		ctx.DebugLog("Ignoring unknown $/setTrace value")
	}

	return nil
}
