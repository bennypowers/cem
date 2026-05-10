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

func createCSSAmbiguousCommentActions(diagnostic *protocol.Diagnostic, data map[string]any, documentURI string) []protocol.CodeAction {
	commentText, _ := data["commentText"].(string)
	if commentText == "" {
		return nil
	}

	deleteRange, ok := parseRange(data["deleteRange"])
	if !ok {
		return nil
	}

	properties, ok := data["properties"].([]any)
	if !ok || len(properties) == 0 {
		return nil
	}

	var actions []protocol.CodeAction
	kind := protocol.CodeActionKindQuickFix

	for _, prop := range properties {
		propMap, ok := prop.(map[string]any)
		if !ok {
			continue
		}
		name, _ := propMap["name"].(string)
		if name == "" {
			continue
		}
		insertPos, ok := parsePosition(propMap["insertPosition"])
		if !ok {
			continue
		}

		actions = append(actions, protocol.CodeAction{
			Title: fmt.Sprintf("Associate comment with `%s`", name),
			Kind:  &kind,
			Edit: &protocol.WorkspaceEdit{
				Changes: map[string][]protocol.TextEdit{
					documentURI: {
						{Range: deleteRange, NewText: ""},
						{Range: protocol.Range{Start: insertPos, End: insertPos}, NewText: commentText + " "},
					},
				},
			},
			Diagnostics: []protocol.Diagnostic{*diagnostic},
		})
	}

	return actions
}

func parseRange(v any) (protocol.Range, bool) {
	m, ok := v.(map[string]any)
	if !ok {
		return protocol.Range{}, false
	}
	start, ok1 := parsePosition(m["start"])
	end, ok2 := parsePosition(m["end"])
	if !ok1 || !ok2 {
		return protocol.Range{}, false
	}
	return protocol.Range{Start: start, End: end}, true
}

func parsePosition(v any) (protocol.Position, bool) {
	m, ok := v.(map[string]any)
	if !ok {
		return protocol.Position{}, false
	}
	line, ok1 := m["line"].(float64)
	char, ok2 := m["character"].(float64)
	if !ok1 || !ok2 {
		return protocol.Position{}, false
	}
	return protocol.Position{Line: uint32(line), Character: uint32(char)}, true
}
