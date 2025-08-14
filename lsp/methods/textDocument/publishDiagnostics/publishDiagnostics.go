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
	M "bennypowers.dev/cem/manifest"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// DiagnosticsContext provides the dependencies needed for diagnostics
type DiagnosticsContext interface {
	Document(uri string) types.Document
	Slots(tagName string) ([]M.Slot, bool)
	AllTagNames() []string
	// For missing import diagnostics
	ElementDefinition(tagName string) (types.ElementDefinition, bool)
	ElementSource(tagName string) (string, bool) // Returns import path/package name
}

// PublishDiagnostics analyzes the document and publishes diagnostics
func PublishDiagnostics(ctx DiagnosticsContext, glspContext *glsp.Context, uri string) error {
	helpers.SafeDebugLog("[DIAGNOSTICS] Starting diagnostics for %s", uri)

	doc := ctx.Document(uri)
	if doc == nil {
		helpers.SafeDebugLog("[DIAGNOSTICS] No document found for %s", uri)
		return nil
	}

	diagnostics := []protocol.Diagnostic{}

	// Analyze slot attribute diagnostics
	slotDiagnostics := analyzeSlotDiagnostics(ctx, doc)
	diagnostics = append(diagnostics, slotDiagnostics...)

	// Analyze custom element tag name diagnostics
	tagDiagnostics := analyzeTagNameDiagnostics(ctx, doc)
	diagnostics = append(diagnostics, tagDiagnostics...)

	helpers.SafeDebugLog("[DIAGNOSTICS] Found %d diagnostics for %s", len(diagnostics), uri)

	// Publish diagnostics to the client
	params := &protocol.PublishDiagnosticsParams{
		URI:         uri,
		Diagnostics: diagnostics,
	}

	glspContext.Notify(protocol.ServerTextDocumentPublishDiagnostics, params)
	return nil
}
