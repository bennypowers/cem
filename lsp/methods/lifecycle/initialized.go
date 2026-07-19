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

// Initialized handles the LSP initialized notification.
// Manifest loading now happens during Initialize (the request) to avoid
// races with AsyncHandler's concurrent dispatch.
func Initialized(_ types.ServerContext, _ *protocol.InitializedParams) error {
	logging.Info("CEM LSP is early software. Report issues at: https://github.com/bennypowers/cem/issues")
	return nil
}
