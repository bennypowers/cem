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
	"regexp"
	"strings"

	"bennypowers.dev/cem/lsp/types"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// CreateMissingImportAction creates a code action to add missing imports
func CreateMissingImportAction(
	ctx types.ServerContext,
	diagnostic *protocol.Diagnostic,
	data map[string]any,
	documentURI string,
) (*protocol.CodeAction, error) {
	// Parse the autofix data using type-safe approach
	autofixData, ok := types.AutofixDataFromMap(data)
	if !ok || autofixData.Type != types.DiagnosticTypeMissingImport {
		return nil, nil
	}

	// Get the document to analyze existing script tags (optional)
	doc := ctx.Document(documentURI)

	var importStatement string
	var insertPosition protocol.Position

	if strings.HasSuffix(documentURI, ".html") {
		// For HTML files, prioritize inline module scripts and head placement
		if doc != nil {
			// Detect the document's indentation pattern
			baseIndent, scriptIndent := detectIndentation(doc)

			// 1. Try to find existing inline module script (not external src)
			scriptPosition, hasInlineScript := doc.FindInlineModuleScript()
			if hasInlineScript {
				// Amend existing inline script tag by adding import statement inside it
				// Use the indentation detected from the specific script tag location
				scriptContentIndent := detectScriptTagIndentation(doc, scriptPosition)
				importStatement = fmt.Sprintf(`%simport "%s";`, scriptContentIndent, autofixData.ImportPath)
				insertPosition = scriptPosition
			} else {
				// 2. Try to find head section for new script placement
				dm, err := ctx.DocumentManager()
				if err != nil {
					return nil, err
				}
				headPosition, hasHead := doc.FindHeadInsertionPoint(dm)
				if hasHead {
					// Create new script tag inside head section with proper indentation
					importStatement = fmt.Sprintf(`%s<script type="module">
%simport "%s";
%s</script>`, baseIndent, scriptIndent, autofixData.ImportPath, baseIndent)
					insertPosition = headPosition
				} else {
					// 3. Fallback: Create new script tag at beginning of document (HTML partial)
					importStatement = fmt.Sprintf(`<script type="module">
%simport "%s";
</script>
`, scriptIndent, autofixData.ImportPath)
					insertPosition = protocol.Position{Line: 0, Character: 0}
				}
			}
		} else {
			// Fallback: Create new script tag when document is not available
			importStatement = fmt.Sprintf(`<script type="module">
  import "%s";
</script>
`, autofixData.ImportPath)
			insertPosition = protocol.Position{Line: 0, Character: 0}
		}
	} else {
		// TypeScript/JavaScript: Create appropriate import statement and find proper position
		if doc != nil {
			// Find the position after existing imports
			insertPosition = findImportInsertionPosition(doc)
			// Add a blank line before the import if inserting after existing imports
			if insertPosition.Line > 0 {
				importStatement = fmt.Sprintf(`
import "%s";`, autofixData.ImportPath)
			} else {
				importStatement = fmt.Sprintf(`import "%s";`, autofixData.ImportPath)
			}
		} else {
			// Fallback to beginning of file
			importStatement = fmt.Sprintf(`import "%s";`, autofixData.ImportPath)
			insertPosition = protocol.Position{Line: 0, Character: 0}
		}
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

	return &action, nil
}

// detectIndentation analyzes the document content to determine the indentation pattern
func detectIndentation(doc types.Document) (baseIndent string, scriptIndent string) {
	if doc == nil {
		// Default to 2 spaces (common in HTML)
		return "  ", "  "
	}

	content, err := doc.Content()
	if err != nil {
		// Default to 2 spaces (common in HTML)
		return "  ", "  "
	}

	lines := strings.Split(content, "\n")

	// Collect indentation samples from non-empty lines
	var indentLevels []int
	var hasTab bool

	// Regex to capture leading whitespace
	indentRegex := regexp.MustCompile(`^(\s+)`)

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue // Skip empty lines
		}

		matches := indentRegex.FindStringSubmatch(line)
		if len(matches) > 1 && matches[1] != "" {
			indent := matches[1]

			// Check if file uses tabs
			if strings.Contains(indent, "\t") {
				hasTab = true
				indentLevels = append(indentLevels, strings.Count(indent, "\t"))
			} else {
				// Count spaces
				spaceCount := len(indent)
				if spaceCount > 0 {
					indentLevels = append(indentLevels, spaceCount)
				}
			}
		}
	}

	if hasTab {
		// File uses tabs
		return "\t", "\t\t"
	}

	if len(indentLevels) > 0 {
		// Find the smallest non-zero indentation (likely base indentation unit)
		minIndent := indentLevels[0]
		for _, level := range indentLevels {
			if level > 0 && level < minIndent {
				minIndent = level
			}
		}

		// Use the detected minimum indentation as the base unit
		baseIndentStr := strings.Repeat(" ", minIndent)
		scriptIndentStr := strings.Repeat(" ", minIndent*2) // Double for script content

		return baseIndentStr, scriptIndentStr
	}

	// No indentation found, default to 2 spaces (common in HTML)
	return "  ", "  "
}

// detectScriptTagIndentation finds the indentation level for content inside an existing script tag
func detectScriptTagIndentation(doc types.Document, scriptPosition protocol.Position) string {
	if doc == nil {
		return "  " // Default to 2 spaces
	}

	content, err := doc.Content()
	if err != nil {
		return "  " // Default to 2 spaces
	}

	lines := strings.Split(content, "\n")

	// Collect indentation from all existing import statements to find the most consistent pattern
	var importIndents []string
	scriptEndLine := int(scriptPosition.Line)
	indentRegex := regexp.MustCompile(`^(\s*)`)

	// Find the actual script start by searching backwards for the opening script tag
	scriptStartLine := 0
	for i := scriptEndLine; i >= 0; i-- {
		if strings.Contains(lines[i], "<script") {
			scriptStartLine = i
			break
		}
	}

	for i := scriptStartLine; i <= scriptEndLine; i++ {
		line := lines[i]

		// Stop if we hit the closing script tag
		if strings.Contains(line, "</script>") {
			break
		}

		// Look for existing import statements
		trimmedLine := strings.TrimSpace(line)
		if strings.HasPrefix(trimmedLine, "import ") {
			// Extract the indentation from this import line
			matches := indentRegex.FindStringSubmatch(line)
			if len(matches) > 1 && matches[1] != "" {
				importIndents = append(importIndents, matches[1])
			}
		}
	}

	// If we found import statements, check if they're consistent
	if len(importIndents) > 0 {
		// Count occurrences of each indentation pattern
		indentCounts := make(map[string]int)
		for _, indent := range importIndents {
			indentCounts[indent]++
		}

		// If all imports use the same indentation, use it
		if len(indentCounts) == 1 {
			for indent := range indentCounts {
				return indent
			}
		}

		// If imports are inconsistent, prefer file's general indentation for consistency
		// This standardizes mixed indentation rather than perpetuating it
	}

	// If imports are inconsistent, prefer file's general indentation for consistency
	if len(importIndents) > 0 {
		baseIndent, _ := detectIndentation(doc)
		return baseIndent
	}
	
	// If no existing imports found, use the file's script indentation pattern
	_, scriptIndent := detectIndentation(doc)
	return scriptIndent
}

// findImportInsertionPosition finds the position to insert new imports after existing imports
func findImportInsertionPosition(doc types.Document) protocol.Position {
	content, err := doc.Content()
	if err != nil {
		return protocol.Position{Line: 0, Character: 0}
	}

	lines := strings.Split(content, "\n")
	
	// Find the last import statement
	lastImportLine := -1
	for i, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		// Look for import statements (both ES6 and CommonJS style)
		if strings.HasPrefix(trimmedLine, "import ") || strings.HasPrefix(trimmedLine, "from ") {
			lastImportLine = i
		}
		// Stop scanning when we hit the first non-import, non-comment, non-empty line
		if trimmedLine != "" && 
		   !strings.HasPrefix(trimmedLine, "import ") && 
		   !strings.HasPrefix(trimmedLine, "from ") &&
		   !strings.HasPrefix(trimmedLine, "//") &&
		   !strings.HasPrefix(trimmedLine, "/*") &&
		   !strings.HasPrefix(trimmedLine, "*") &&
		   !strings.HasPrefix(trimmedLine, "*/") {
			break
		}
	}
	
	if lastImportLine >= 0 {
		// Insert after the last import, with a blank line
		return protocol.Position{Line: uint32(lastImportLine + 1), Character: 0}
	}
	
	// No imports found, insert at the beginning
	return protocol.Position{Line: 0, Character: 0}
}
