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

	protocol "github.com/bennypowers/glsp/protocol_3_17"
)

func createAutofixAction(diagnostic *protocol.Diagnostic, dataMap map[string]any, uri string) *protocol.CodeAction {
	original, originalOk := dataMap["original"].(string)
	suggestion, suggestionOk := dataMap["suggestion"].(string)

	if !originalOk || !suggestionOk {
		return nil
	}

	kind := protocol.CodeActionKindQuickFix
	return &protocol.CodeAction{
		Title: fmt.Sprintf("Change '%s' to '%s'", original, suggestion),
		Kind:  &kind,
		Edit: &protocol.WorkspaceEdit{
			Changes: map[string][]protocol.TextEdit{
				uri: {{Range: diagnostic.Range, NewText: suggestion}},
			},
		},
		Diagnostics: []protocol.Diagnostic{*diagnostic},
	}
}

func createAttributeAutofixAction(diagnostic *protocol.Diagnostic, dataMap map[string]any, uri string) *protocol.CodeAction {
	return createAutofixAction(diagnostic, dataMap, uri)
}

func createAttributeValueAutofixAction(diagnostic *protocol.Diagnostic, dataMap map[string]any, uri string) *protocol.CodeAction {
	return createAutofixAction(diagnostic, dataMap, uri)
}
