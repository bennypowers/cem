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

	"bennypowers.dev/cem/lsp/types"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// createSlotAutofixAction creates a code action to fix an invalid slot name
func createSlotAutofixAction(diagnostic *protocol.Diagnostic, data map[string]any, documentURI string) *protocol.CodeAction {
	// Parse the autofix data using type-safe approach
	autofixData, ok := types.AutofixDataFromMap(data)
	if !ok || autofixData.Type != types.DiagnosticTypeSlotSuggestion {
		return nil
	}

	title := fmt.Sprintf("Change '%s' to '%s'", autofixData.Original, autofixData.Suggestion)
	kind := protocol.CodeActionKindQuickFix

	action := protocol.CodeAction{
		Title: title,
		Kind:  &kind,
		Edit: &protocol.WorkspaceEdit{
			Changes: map[string][]protocol.TextEdit{
				documentURI: {
					{
						Range:   autofixData.Range,
						NewText: autofixData.Suggestion,
					},
				},
			},
		},
		Diagnostics: []protocol.Diagnostic{*diagnostic},
	}

	return &action
}
