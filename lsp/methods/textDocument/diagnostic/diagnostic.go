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

// DocumentDiagnostic handles the textDocument/diagnostic request (LSP 3.17 pull model)
func DocumentDiagnostic(ctx types.ServerContext, context *glsp.Context, params *protocol.DocumentDiagnosticParams) (any, error) {
	helpers.SafeDebugLog("[DIAGNOSTICS] Pull diagnostic request for %s", params.TextDocument.URI)

	doc := ctx.Document(params.TextDocument.URI)
	if doc == nil {
		return protocol.RelatedFullDocumentDiagnosticReport{
			FullDocumentDiagnosticReport: protocol.FullDocumentDiagnosticReport{
				Kind:  string(protocol.DocumentDiagnosticReportKindFull),
				Items: []protocol.Diagnostic{},
			},
		}, nil
	}

	diagnostics := publishDiagnostics.ComputeDiagnostics(ctx, doc)

	helpers.SafeDebugLog("[DIAGNOSTICS] Pull diagnostics returning %d items for %s", len(diagnostics), params.TextDocument.URI)

	return protocol.RelatedFullDocumentDiagnosticReport{
		FullDocumentDiagnosticReport: protocol.FullDocumentDiagnosticReport{
			Kind:  string(protocol.DocumentDiagnosticReportKindFull),
			Items: diagnostics,
		},
	}, nil
}
