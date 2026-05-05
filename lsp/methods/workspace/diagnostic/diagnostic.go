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
package diagnostic

import (
	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/methods/textDocument/publishDiagnostics"
	"bennypowers.dev/cem/lsp/types"
	"github.com/bennypowers/glsp"
	protocol "github.com/bennypowers/glsp/protocol_3_17"
)

// WorkspaceDiagnostic handles the workspace/diagnostic request (LSP 3.17 pull model)
func WorkspaceDiagnostic(ctx types.ServerContext, context *glsp.Context, params *protocol.WorkspaceDiagnosticParams) (*protocol.WorkspaceDiagnosticReport, error) {
	helpers.SafeDebugLog("[DIAGNOSTICS] Workspace diagnostic request")

	items := []protocol.WorkspaceDocumentDiagnosticReport{}
	for _, doc := range ctx.AllDocuments() {
		docURI := doc.URI()
		diagnostics := publishDiagnostics.ComputeDiagnostics(ctx, doc)

		items = append(items, protocol.WorkspaceFullDocumentDiagnosticReport{
			FullDocumentDiagnosticReport: protocol.FullDocumentDiagnosticReport{
				Kind:  string(protocol.DocumentDiagnosticReportKindFull),
				Items: diagnostics,
			},
			URI:     docURI,
			Version: nil,
		})
	}

	helpers.SafeDebugLog("[DIAGNOSTICS] Workspace diagnostics returning %d documents", len(items))

	return &protocol.WorkspaceDiagnosticReport{
		Items: items,
	}, nil
}
