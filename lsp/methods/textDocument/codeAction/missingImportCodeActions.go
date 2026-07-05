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

	"bennypowers.dev/cem/internal/languages/typescript"
	"bennypowers.dev/cem/lsp/types"
	protocol "github.com/bennypowers/glsp/protocol_3_17"
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
				result := findLastImportPosition(doc, scriptPosition)
				scriptContentIndent := detectScriptTagIndentation(doc, scriptPosition)
				if result.sameLineScript && result.scriptTag != nil {
					st := result.scriptTag
					tagRange := st.Range
					docContent, _ := doc.Content()
					docLines := strings.Split(docContent, "\n")
					tagLine := docLines[tagRange.Start.Line]
					openTag := tagLine[tagRange.Start.Character:st.ContentRange.Start.Character]
					contentStart := int(st.ContentRange.Start.Character)
					contentEnd := int(st.ContentRange.End.Character)
					insertChar := int(result.position.Character)
					beforeImport := strings.TrimSpace(tagLine[contentStart:insertChar])
					afterImport := strings.TrimSpace(tagLine[insertChar:contentEnd])
					tagIndent := strings.Repeat(" ", int(tagRange.Start.Character))
					var replacement string
					if afterImport == "" {
						replacement = fmt.Sprintf("%s\n%s%s\n%simport \"%s\";\n%s</script>",
							openTag, scriptContentIndent, beforeImport,
							scriptContentIndent, autofixData.ImportPath,
							tagIndent)
					} else {
						replacement = fmt.Sprintf("%s\n%s%s\n%simport \"%s\";\n%s%s\n%s</script>",
							openTag, scriptContentIndent, beforeImport,
							scriptContentIndent, autofixData.ImportPath,
							scriptContentIndent, afterImport,
							tagIndent)
					}
					action := buildCodeAction(autofixData.TagName, documentURI, diagnostic,
						tagRange, replacement)
					return &action, nil
				}
				importStatement = fmt.Sprintf(`%simport "%s";`, scriptContentIndent, autofixData.ImportPath)
				insertPosition = result.position
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

	editRange := protocol.Range{Start: insertPosition, End: insertPosition}
	action := buildCodeAction(autofixData.TagName, documentURI, diagnostic,
		editRange, importStatement+"\n")
	return &action, nil
}

func buildCodeAction(tagName, documentURI string, diagnostic *protocol.Diagnostic,
	editRange protocol.Range, newText string,
) protocol.CodeAction {
	title := fmt.Sprintf("Add import for '%s'", tagName)
	kind := protocol.CodeActionKindQuickFix
	preferred := true
	return protocol.CodeAction{
		Title:       title,
		Kind:        &kind,
		IsPreferred: &preferred,
		Edit: &protocol.WorkspaceEdit{
			Changes: map[string][]protocol.TextEdit{
				documentURI: {{Range: editRange, NewText: newText}},
			},
		},
		Diagnostics: []protocol.Diagnostic{*diagnostic},
	}
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

	// If no existing imports found, try to detect the base indentation level
	// inside the script tag rather than using the doubled scriptIndent
	baseIndent, _ := detectIndentation(doc)
	return baseIndent
}

// importInsertionResult holds the result of finding the import insertion point.
type importInsertionResult struct {
	position       protocol.Position
	sameLineScript bool
	scriptTag      *types.ScriptTag
}

// findLastImportPosition finds the position after the last import statement
// in an inline script tag by parsing the script content as TypeScript.
func findLastImportPosition(doc types.Document, scriptPosition protocol.Position) importInsertionResult {
	fallback := importInsertionResult{position: scriptPosition}
	if doc == nil {
		return fallback
	}

	var scriptTag *types.ScriptTag
	scriptTags := doc.ScriptTags()
	for i := range scriptTags {
		st := &scriptTags[i]
		if st.IsModule && st.Src == "" &&
			scriptPosition.Line >= st.Range.Start.Line && scriptPosition.Line <= st.Range.End.Line {
			scriptTag = st
			break
		}
	}
	if scriptTag == nil {
		return fallback
	}

	content, err := doc.Content()
	if err != nil {
		return fallback
	}

	lines := strings.Split(content, "\n")
	startLine := int(scriptTag.ContentRange.Start.Line)
	endLine := int(scriptTag.ContentRange.End.Line)
	if startLine >= len(lines) {
		return fallback
	}
	sameLineScript := startLine == endLine

	var scriptContent strings.Builder
	for i := startLine; i <= endLine && i < len(lines); i++ {
		line := lines[i]
		if i == startLine && i == endLine {
			startChar := min(int(scriptTag.ContentRange.Start.Character), len(line))
			endChar := min(int(scriptTag.ContentRange.End.Character), len(line))
			line = line[startChar:endChar]
		} else if i == startLine {
			startChar := min(int(scriptTag.ContentRange.Start.Character), len(line))
			line = line[startChar:]
		} else if i == endLine {
			endChar := min(int(scriptTag.ContentRange.End.Character), len(line))
			line = line[:endChar]
		}
		if i > startLine {
			scriptContent.WriteByte('\n')
		}
		scriptContent.WriteString(line)
	}

	parser := typescript.BorrowParser()
	defer typescript.ReturnParser(parser)

	tree := parser.Parse([]byte(scriptContent.String()), nil)
	if tree == nil {
		return fallback
	}
	defer tree.Close()

	root := tree.RootNode()
	var lastImportEnd protocol.Position
	found := false

	for i := range root.ChildCount() {
		child := root.Child(i)
		if child == nil {
			continue
		}
		if child.Kind() == "import_statement" {
			end := child.EndPosition()
			if sameLineScript {
				lastImportEnd = protocol.Position{
					Line:      uint32(startLine),
					Character: scriptTag.ContentRange.Start.Character + uint32(end.Column),
				}
			} else {
				lastImportEnd = protocol.Position{
					Line:      uint32(startLine) + uint32(end.Row) + 1,
					Character: 0,
				}
			}
			found = true
		}
	}

	if found {
		return importInsertionResult{
			position:       lastImportEnd,
			sameLineScript: sameLineScript,
			scriptTag:      scriptTag,
		}
	}

	return fallback
}

// findImportInsertionPosition finds the position to insert new imports after
// existing imports using tree-sitter to correctly handle multi-line imports.
func findImportInsertionPosition(doc types.Document) protocol.Position {
	tree, release := doc.AcquireTree()
	if tree == nil {
		return protocol.Position{Line: 0, Character: 0}
	}
	defer release()

	root := tree.RootNode()
	var lastImportEnd protocol.Position
	found := false

	for i := range root.ChildCount() {
		child := root.Child(i)
		if child == nil {
			continue
		}
		if child.Kind() == "import_statement" {
			end := child.EndPosition()
			lastImportEnd = protocol.Position{
				Line:      uint32(end.Row + 1),
				Character: 0,
			}
			found = true
		}
	}

	if found {
		return lastImportEnd
	}

	return protocol.Position{Line: 0, Character: 0}
}
