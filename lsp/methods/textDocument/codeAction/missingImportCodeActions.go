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
func createMissingImportAction(diagnostic *protocol.Diagnostic, data map[string]any, documentURI string) *protocol.CodeAction {
	// Parse the autofix data using type-safe approach
	autofixData, ok := types.AutofixDataFromMap(data)
	if !ok || autofixData.Type != types.DiagnosticTypeMissingImport {
		return nil
	}

	var importStatement string
	var insertPosition protocol.Position

	// Determine if this is a package import or local module import
	isPackageImport := !strings.HasPrefix(autofixData.ImportPath, "./") &&
		!strings.HasPrefix(autofixData.ImportPath, "../") &&
		!strings.HasPrefix(autofixData.ImportPath, "/")

	if strings.HasSuffix(documentURI, ".html") {
		if isPackageImport {
			// HTML: Add script tag for package import - use CDN or importmap
			// For packages, suggest using import maps or CDN
			importStatement = fmt.Sprintf(`<!-- Add to import map: "%s": "path/to/%s" -->
<script type="module" src="%s"></script>`, autofixData.ImportPath, autofixData.ImportPath, autofixData.ImportPath)
		} else {
			// HTML: Add script tag for local module
			importStatement = fmt.Sprintf(`<script type="module" src="%s"></script>`, autofixData.ImportPath)
		}
		// TODO: Find proper position in HTML head section using tree-sitter parsing
		insertPosition = protocol.Position{Line: 0, Character: 0}
	} else {
		// TypeScript/JavaScript: Create appropriate import statement
		if isPackageImport {
			// Package import: import "package-name"
			importStatement = fmt.Sprintf(`import "%s";`, autofixData.ImportPath)
		} else {
			// Local module import: import "./local-element.js"
			importStatement = fmt.Sprintf(`import "%s";`, autofixData.ImportPath)
		}
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
