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
	Q "bennypowers.dev/cem/internal/treesitter"
	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	"encoding/json"
	"go.lsp.dev/protocol"
)

func analyzeCssDiagnostics(ctx types.ServerContext, doc types.Document) []protocol.Diagnostic {
	return AnalyzeCssDiagnosticsForTest(ctx, doc)
}

// AnalyzeCssDiagnosticsForTest is the exported version for testing
func AnalyzeCssDiagnosticsForTest(ctx types.ServerContext, doc types.Document) []protocol.Diagnostic {
	var diagnostics []protocol.Diagnostic

	if doc.Language() != "css" {
		return diagnostics
	}

	content, err := doc.Content()
	if err != nil {
		return diagnostics
	}

	tree, releaseTree := doc.AcquireTree()
	if tree == nil {
		return diagnostics
	}
	defer releaseTree()

	queryManager, err := ctx.QueryManager()
	if err != nil {
		helpers.SafeDebugLog("[DIAGNOSTICS] Failed to get QueryManager: %v", err)
		return diagnostics
	}

	matcher, err := Q.GetCachedQueryMatcher(queryManager, "css", "cssCustomProperties")
	if err != nil {
		helpers.SafeDebugLog("[DIAGNOSTICS] Failed to create CSS query matcher: %v", err)
		return diagnostics
	}
	defer matcher.Close()

	contentBytes := []byte(content)
	for captures := range matcher.ParentCaptures(tree.RootNode(), contentBytes, "cssProperty") {
		comments, hasComment := captures["comment"]
		properties := captures["property"]
		fns := captures["fn"]
		if !hasComment || len(comments) == 0 || len(properties) <= 1 {
			continue
		}
		commentRange := doc.ByteRangeToProtocolRange(content, comments[0].StartByte, comments[0].EndByte)
		deleteRange := protocol.Range{
			Start: protocol.Position{Line: commentRange.Start.Line, Character: 0},
			End:   protocol.Position{Line: commentRange.End.Line + 1, Character: 0},
		}
		propertyData := make([]map[string]any, 0, len(properties))
		for i, prop := range properties {
			entry := map[string]any{
				"name": prop.Text,
			}
			if i < len(fns) {
				insertRange := doc.ByteRangeToProtocolRange(content, fns[i].StartByte, fns[i].StartByte)
				entry["insertPosition"] = map[string]any{
					"line":      float64(insertRange.Start.Line),
					"character": float64(insertRange.Start.Character),
				}
			}
			propertyData = append(propertyData, entry)
		}
		dataMap := map[string]any{
			"type":        string(types.DiagnosticTypeCSSAmbiguousComment),
			"commentText": comments[0].Text,
			"deleteRange": map[string]any{
				"start": map[string]any{
					"line":      float64(deleteRange.Start.Line),
					"character": float64(deleteRange.Start.Character),
				},
				"end": map[string]any{
					"line":      float64(deleteRange.End.Line),
					"character": float64(deleteRange.End.Character),
				},
			},
			"properties": propertyData,
		}
		data, _ := json.Marshal(dataMap)
		diagnostics = append(diagnostics, protocol.Diagnostic{
			Range:    commentRange,
			Severity: protocol.DiagnosticSeverityWarning,
			Source:   protocol.NewOptional("cem-lsp"),
			Message:  protocol.String("Ambiguous comment ignored: more than one var() call in declaration."),
			Data:     data,
		})
	}

	return diagnostics
}
