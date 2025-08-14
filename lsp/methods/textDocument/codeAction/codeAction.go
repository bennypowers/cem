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
	"fmt"

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
				if dataMap, ok := diagnostic.Data.(map[string]interface{}); ok {
					if actionType, exists := dataMap["type"]; exists && actionType == "slot-suggestion" {
						action := createSlotAutofixAction(&diagnostic, dataMap, params.TextDocument.URI)
						if action != nil {
							actions = append(actions, *action)
							helpers.SafeDebugLog("[CODE_ACTION] Created slot autofix action")
						}
					}
				}
			}
		}
	}

	helpers.SafeDebugLog("[CODE_ACTION] Returning %d code actions", len(actions))
	return actions, nil
}

// createSlotAutofixAction creates a code action to fix an invalid slot name
func createSlotAutofixAction(diagnostic *protocol.Diagnostic, data map[string]interface{}, documentURI string) *protocol.CodeAction {
	original, originalOk := data["original"].(string)
	suggestion, suggestionOk := data["suggestion"].(string)
	rangeData, rangeOk := data["range"]

	if !originalOk || !suggestionOk || !rangeOk {
		return nil
	}

	// Convert range data back to protocol.Range
	var fixRange protocol.Range
	if rangeMap, ok := rangeData.(map[string]interface{}); ok {
		if start, startOk := rangeMap["start"].(map[string]interface{}); startOk {
			if line, lineOk := start["line"].(float64); lineOk {
				fixRange.Start.Line = uint32(line)
			}
			if char, charOk := start["character"].(float64); charOk {
				fixRange.Start.Character = uint32(char)
			}
		}
		if end, endOk := rangeMap["end"].(map[string]interface{}); endOk {
			if line, lineOk := end["line"].(float64); lineOk {
				fixRange.End.Line = uint32(line)
			}
			if char, charOk := end["character"].(float64); charOk {
				fixRange.End.Character = uint32(char)
			}
		}
	} else {
		// Fallback: use diagnostic range if data range parsing fails
		fixRange = diagnostic.Range
	}

	title := fmt.Sprintf("Change '%s' to '%s'", original, suggestion)
	kind := protocol.CodeActionKindQuickFix

	action := protocol.CodeAction{
		Title: title,
		Kind:  &kind,
		Edit: &protocol.WorkspaceEdit{
			Changes: map[string][]protocol.TextEdit{
				documentURI: {
					{
						Range:   fixRange,
						NewText: suggestion,
					},
				},
			},
		},
		Diagnostics: []protocol.Diagnostic{*diagnostic},
	}

	return &action
}
