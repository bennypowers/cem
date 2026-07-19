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
	"encoding/json"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	"go.lsp.dev/protocol"
)

// CodeAction handles textDocument/codeAction requests
func CodeAction(ctx types.ServerContext, params *protocol.CodeActionParams) ([]protocol.CodeAction, error) {
	helpers.SafeDebugLog("[CODE_ACTION] Starting code action for %s", params.TextDocument.URI)

	var actions []protocol.CodeAction

	docURI := string(params.TextDocument.URI)

	for _, diagnostic := range params.Context.Diagnostics {
		source, hasSource := diagnostic.Source.Get()
		if !hasSource || source != "cem-lsp" {
			continue
		}
		if len(diagnostic.Data) == 0 {
			continue
		}
		var dataMap map[string]any
		if err := json.Unmarshal(diagnostic.Data, &dataMap); err != nil {
			continue
		}
		actionType, exists := dataMap["type"]
		if !exists {
			continue
		}
		switch actionType {
		case "slot-suggestion":
			action := createSlotAutofixAction(&diagnostic, dataMap, docURI)
			if action != nil {
				actions = append(actions, *action)
				helpers.SafeDebugLog("[CODE_ACTION] Created slot autofix action")
			}
		case "tag-suggestion":
			action := createTagAutofixAction(&diagnostic, dataMap, docURI)
			if action != nil {
				actions = append(actions, *action)
				helpers.SafeDebugLog("[CODE_ACTION] Created tag autofix action")
			}
		case "missing-import":
			action, err := CreateMissingImportAction(ctx, &diagnostic, dataMap, docURI)
			if err != nil {
				return nil, err
			}
			if action != nil {
				actions = append(actions, *action)
				helpers.SafeDebugLog("[CODE_ACTION] Created missing import action")
			}
		case "attribute-suggestion":
			action := createAttributeAutofixAction(&diagnostic, dataMap, docURI)
			if action != nil {
				actions = append(actions, *action)
				helpers.SafeDebugLog("[CODE_ACTION] Created attribute autofix action")
			}
		case "attribute-value-suggestion":
			action := createAttributeValueAutofixAction(&diagnostic, dataMap, docURI)
			if action != nil {
				actions = append(actions, *action)
				helpers.SafeDebugLog("[CODE_ACTION] Created attribute value autofix action")
			}
		case "css-ambiguous-comment":
			cssActions := createCSSAmbiguousCommentActions(&diagnostic, dataMap, docURI)
			actions = append(actions, cssActions...)
			helpers.SafeDebugLog("[CODE_ACTION] Created %d CSS ambiguous comment actions", len(cssActions))
		}
	}

	helpers.SafeDebugLog("[CODE_ACTION] Returning %d code actions", len(actions))
	return actions, nil
}
