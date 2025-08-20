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

	"bennypowers.dev/cem/lsp/types"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// Initialized handles the LSP initialized notification
func Initialized(ctx types.ServerContext, context *glsp.Context, params *protocol.InitializedParams) error {
	log.Printf("CEM LSP Server initialized")

	// Initialize manifests and start watching after successful LSP initialization
	if err := ctx.InitializeManifests(); err != nil {
		log.Printf("Warning: Failed to initialize manifests: %v", err)
		// Don't fail the LSP initialization if manifest loading fails
	}

	return nil
}
