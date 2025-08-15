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
	"strings"

	"bennypowers.dev/cem/lsp/types"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// createMissingImportAction creates a code action to add missing imports
func createMissingImportAction(ctx CodeActionContext, diagnostic *protocol.Diagnostic, data map[string]any, documentURI string) *protocol.CodeAction {
	// Parse the autofix data using type-safe approach
	autofixData, ok := types.AutofixDataFromMap(data)
	if !ok || autofixData.Type != types.DiagnosticTypeMissingImport {
		return nil
	}

	// Get the document to analyze existing script tags
	doc := ctx.Document(documentURI)
	if doc == nil {
		return nil
	}

	var importStatement string
	var insertPosition protocol.Position

	if strings.HasSuffix(documentURI, ".html") {
		// For HTML files, try to amend existing script type="module" tags
		scriptPosition, hasExistingScript := doc.FindModuleScript()

		if hasExistingScript {
			// Amend existing script tag by adding import statement inside it
			importStatement = fmt.Sprintf(`	import '%s';`, autofixData.ImportPath)
			insertPosition = scriptPosition
		} else {
			// Create new script tag if no existing one found
			importStatement = fmt.Sprintf(`<script type="module">
	import '%s';
</script>`, autofixData.ImportPath)
			// TODO: Find proper position in HTML head section using tree-sitter parsing
			insertPosition = protocol.Position{Line: 0, Character: 0}
		}
	} else {
		// TypeScript/JavaScript: Create appropriate import statement
		importStatement = fmt.Sprintf(`import '%s';`, autofixData.ImportPath)
		// TODO: Find proper position with other imports using tree-sitter
		insertPosition = protocol.Position{Line: 0, Character: 0}
	}

	title := fmt.Sprintf("Add import for '%s'", autofixData.TagName)
	kind := protocol.CodeActionKindQuickFix

	action := protocol.CodeAction{
		Title: title,
		Kind:  &kind,
		Edit: &protocol.WorkspaceEdit{
			Changes: map[string][]protocol.TextEdit{
				documentURI: {
					{
						Range: protocol.Range{
							Start: insertPosition,
							End:   insertPosition,
						},
						NewText: importStatement + "\n",
					},
				},
			},
		},
		Diagnostics: []protocol.Diagnostic{*diagnostic},
	}

	return &action
}
