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
package codeAction

import (
	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// CodeActionContext provides the dependencies needed for code actions
type CodeActionContext interface {
	GetDocument(uri string) types.Document
}

// CodeAction handles textDocument/codeAction requests
func CodeAction(ctx CodeActionContext, context *glsp.Context, params *protocol.CodeActionParams) (any, error) {
	helpers.SafeDebugLog("[CODE_ACTION] Starting code action for %s", params.TextDocument.URI)

	var actions []protocol.CodeAction

	// Process diagnostics that have autofix suggestions
	for _, diagnostic := range params.Context.Diagnostics {
		if diagnostic.Source != nil && *diagnostic.Source == "cem-lsp" {
			if diagnostic.Data != nil {
				if dataMap, ok := diagnostic.Data.(map[string]any); ok {
					if actionType, exists := dataMap["type"]; exists {
						switch actionType {
						case "slot-suggestion":
							action := createSlotAutofixAction(&diagnostic, dataMap, params.TextDocument.URI)
							if action != nil {
								actions = append(actions, *action)
								helpers.SafeDebugLog("[CODE_ACTION] Created slot autofix action")
							}
						case "tag-suggestion":
							action := createTagAutofixAction(&diagnostic, dataMap, params.TextDocument.URI)
							if action != nil {
								actions = append(actions, *action)
								helpers.SafeDebugLog("[CODE_ACTION] Created tag autofix action")
							}
						case "missing-import":
							action := createMissingImportAction(&diagnostic, dataMap, params.TextDocument.URI)
							if action != nil {
								actions = append(actions, *action)
								helpers.SafeDebugLog("[CODE_ACTION] Created missing import action")
							}
						}
					}
				}
			}
		}
	}

	helpers.SafeDebugLog("[CODE_ACTION] Returning %d code actions", len(actions))
	return actions, nil
}
