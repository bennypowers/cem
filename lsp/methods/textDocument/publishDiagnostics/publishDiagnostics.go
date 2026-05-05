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
package publishDiagnostics

import (
	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	"github.com/bennypowers/glsp"
	protocol "github.com/bennypowers/glsp/protocol_3_17"
)

// ComputeDiagnostics analyzes a document and returns a non-nil slice of diagnostics.
func ComputeDiagnostics(ctx types.ServerContext, doc types.Document) []protocol.Diagnostic {
	diagnostics := []protocol.Diagnostic{}
	diagnostics = append(diagnostics, analyzeSlotDiagnostics(ctx, doc)...)
	diagnostics = append(diagnostics, analyzeTagNameDiagnostics(ctx, doc)...)
	diagnostics = append(diagnostics, analyzeAttributeDiagnostics(ctx, doc)...)
	diagnostics = append(diagnostics, analyzeAttributeValueDiagnostics(ctx, doc)...)
	return diagnostics
}

// PublishDiagnostics analyzes the document and publishes diagnostics via push notification
func PublishDiagnostics(ctx types.ServerContext, glspContext *glsp.Context, uri string) error {
	helpers.SafeDebugLog("[DIAGNOSTICS] Starting diagnostics for %s", uri)

	doc := ctx.Document(uri)
	if doc == nil {
		helpers.SafeDebugLog("[DIAGNOSTICS] No document found for %s", uri)
		return nil
	}

	diagnostics := ComputeDiagnostics(ctx, doc)
	helpers.SafeDebugLog("[DIAGNOSTICS] Found %d diagnostics for %s", len(diagnostics), uri)

	glspContext.Notify(protocol.ServerTextDocumentPublishDiagnostics, &protocol.PublishDiagnosticsParams{
		URI:         uri,
		Diagnostics: diagnostics,
	})
	return nil
}
