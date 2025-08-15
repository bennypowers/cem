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
	"fmt"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// analyzeTagNameDiagnostics finds invalid custom element tag names and suggests corrections
func analyzeTagNameDiagnostics(ctx DiagnosticsContext, doc types.Document) []protocol.Diagnostic {
	return AnalyzeTagNameDiagnosticsForTest(ctx, doc)
}

// AnalyzeTagNameDiagnosticsForTest is the exported version for testing
func AnalyzeTagNameDiagnosticsForTest(ctx DiagnosticsContext, doc types.Document) []protocol.Diagnostic {
	var diagnostics []protocol.Diagnostic

	// Get document content to search for custom element tags
	content := ""
	if docWithContent, ok := doc.(interface{ Content() string }); ok {
		content = docWithContent.Content()
	} else {
		return diagnostics
	}

	helpers.SafeDebugLog("[DIAGNOSTICS] Analyzing custom element tag names in document (length=%d)", len(content))

	// Find all custom element tag names in the document
	tagMatches := findCustomElementTags(content)
	helpers.SafeDebugLog("[DIAGNOSTICS] Found %d custom element tags", len(tagMatches))

	// Parse script imports to get actually imported elements
	importedElements := parseScriptImports(content, ctx, doc)
	helpers.SafeDebugLog("[DIAGNOSTICS] Imported elements from script tags: %v", importedElements)

	// Get all available tag names from manifest for fallback
	allAvailableTagNames := ctx.AllTagNames()
	helpers.SafeDebugLog("[DIAGNOSTICS] All available tag names in manifests: %v", allAvailableTagNames)

	for _, match := range tagMatches {
		tagName := match.Value
		helpers.SafeDebugLog("[DIAGNOSTICS] Checking tag '%s'", tagName)

		// Check if the tag exists in the manifest
		isValid := slices.Contains(allAvailableTagNames, tagName)

		// Check if the tag is imported (available in current context)
		isImported := slices.Contains(importedElements, tagName)

		if !isValid {
			// Check if this element exists but might need an import
			if _, exists := ctx.ElementDefinition(tagName); exists {
				if importPath, hasSource := ctx.ElementSource(tagName); hasSource {
					// Element exists but may need import - create missing import diagnostic
					// Note: importPath currently contains module path, needs package.json resolution
					helpers.SafeDebugLog("[DIAGNOSTICS] Element '%s' exists but may need import from '%s'", tagName, importPath)

					severity := protocol.DiagnosticSeverityError
					source := "cem-lsp"
					var diagnostic protocol.Diagnostic
					diagnostic.Range = match.Range
					diagnostic.Severity = &severity
					diagnostic.Source = &source
					diagnostic.Message = fmt.Sprintf("Custom element '%s' is not imported. Add import from '%s'", tagName, importPath)

					// Add missing import autofix data
					autofixData := &types.AutofixData{
						Type:       types.DiagnosticTypeMissingImport,
						Original:   tagName,
						Suggestion: fmt.Sprintf("import '%s'", importPath),
						Range:      match.Range,
						ImportPath: importPath,
						TagName:    tagName,
					}
					diagnostic.Data = autofixData.ToMap()

					diagnostics = append(diagnostics, diagnostic)
					helpers.SafeDebugLog("[DIAGNOSTICS] Added missing import diagnostic for tag '%s'", tagName)
					continue
				}
			}
			helpers.SafeDebugLog("[DIAGNOSTICS] Invalid custom element tag '%s'", tagName)

			// Find closest match using Levenshtein distance
			closestMatch, distance := findClosestMatch(tagName, allAvailableTagNames, 3)

			severity := protocol.DiagnosticSeverityError
			source := "cem-lsp"
			var diagnostic protocol.Diagnostic
			diagnostic.Range = match.Range
			diagnostic.Severity = &severity
			diagnostic.Source = &source

			if closestMatch != "" && distance <= 2 {
				diagnostic.Message = fmt.Sprintf("Unknown custom element '%s'. Did you mean '%s'?", tagName, closestMatch)

				// Add code action data for quick fix
				autofixData := &types.AutofixData{
					Type:       types.DiagnosticTypeTagSuggestion,
					Original:   tagName,
					Suggestion: closestMatch,
					Range:      match.Range,
				}
				diagnostic.Data = autofixData.ToMap()
			} else {
				if len(allAvailableTagNames) > 0 {
					// For large distances, avoid showing random elements which might not be helpful
					if len(allAvailableTagNames) <= 5 {
						// Only show elements if there are very few available
						// Sort for consistent output in tests and user experience
						sortedNames := make([]string, len(allAvailableTagNames))
						copy(sortedNames, allAvailableTagNames)
						slices.Sort(sortedNames)
						availableList := strings.Join(sortedNames, "', '")
						diagnostic.Message = fmt.Sprintf("Unknown custom element '%s'. Available elements: '%s'", tagName, availableList)
					} else {
						// For projects with many elements, suggest checking documentation instead
						diagnostic.Message = fmt.Sprintf("Unknown custom element '%s'. Check available elements in your project's manifest or documentation.", tagName)
					}
				} else {
					diagnostic.Message = fmt.Sprintf("Unknown custom element '%s'. No custom elements found in manifest.", tagName)
				}
			}

			diagnostics = append(diagnostics, diagnostic)
			helpers.SafeDebugLog("[DIAGNOSTICS] Added diagnostic for invalid tag '%s'", tagName)
		} else if isValid && !isImported {
			// Element exists in manifest but is not imported
			if importPath, hasSource := ctx.ElementSource(tagName); hasSource {
				severity := protocol.DiagnosticSeverityError
				source := "cem-lsp"
				var diagnostic protocol.Diagnostic
				diagnostic.Range = match.Range
				diagnostic.Severity = &severity
				diagnostic.Source = &source
				diagnostic.Message = fmt.Sprintf("Custom element '%s' is not imported. Add import from '%s'", tagName, importPath)

				// Add missing import autofix data
				autofixData := &types.AutofixData{
					Type:       types.DiagnosticTypeMissingImport,
					Original:   tagName,
					Suggestion: fmt.Sprintf("import '%s'", importPath),
					Range:      match.Range,
					ImportPath: importPath,
					TagName:    tagName,
				}
				diagnostic.Data = autofixData.ToMap()

				diagnostics = append(diagnostics, diagnostic)
				helpers.SafeDebugLog("[DIAGNOSTICS] Added missing import diagnostic for valid but unimported tag '%s'", tagName)
			}
		}
	}

	return diagnostics
}

// parseScriptImports parses HTML script tags and TypeScript imports and returns custom element tag names that are imported
func parseScriptImports(content string, ctx DiagnosticsContext, doc types.Document) []string {
	var importedElements []string

	// Check for ignore comment first
	if strings.Contains(content, "// cem-lsp ignore missing-import") ||
		strings.Contains(content, "<!-- cem-lsp ignore missing-import -->") {
		helpers.SafeDebugLog("[DIAGNOSTICS] Found ignore comment, skipping import parsing")
		// Return all available elements to effectively disable missing import checks
		return ctx.AllTagNames()
	}

	// Parse TypeScript imports directly (for .ts files)
	importedElements = append(importedElements, parseTypeScriptImports(content, ctx)...)

	// Parse module scripts for static and dynamic imports (for HTML files)
	importedElements = append(importedElements, parseModuleScriptImports(content, ctx, doc)...)

	// Parse non-module scripts with src attributes (for HTML files)
	importedElements = append(importedElements, parseNonModuleScriptImports(content, ctx)...)

	// Remove duplicates
	seen := make(map[string]bool)
	var uniqueImports []string
	for _, elem := range importedElements {
		if !seen[elem] {
			seen[elem] = true
			uniqueImports = append(uniqueImports, elem)
		}
	}

	return uniqueImports
}

// parseTypeScriptImports parses TypeScript import statements directly from file content
func parseTypeScriptImports(content string, ctx DiagnosticsContext) []string {
	var importedElements []string

	// Parse static imports: import 'path' or import from 'path'
	staticImportRegex := regexp.MustCompile(`import\s+(?:.*\s+from\s+)?['"]([^'"]+)['"]`)
	importMatches := staticImportRegex.FindAllStringSubmatch(content, -1)

	for _, importMatch := range importMatches {
		if len(importMatch) > 1 {
			importPath := importMatch[1]
			elements := resolveImportPathToElements(importPath, ctx)
			importedElements = append(importedElements, elements...)
			helpers.SafeDebugLog("[DIAGNOSTICS] TypeScript import '%s' resolved to elements: %v", importPath, elements)
		}
	}

	// Parse dynamic imports: import('path')
	dynamicImportRegex := regexp.MustCompile(`import\s*\(\s*['"]([^'"]+)['"]\s*\)`)
	dynamicMatches := dynamicImportRegex.FindAllStringSubmatch(content, -1)

	for _, dynamicMatch := range dynamicMatches {
		if len(dynamicMatch) > 1 {
			importPath := dynamicMatch[1]
			elements := resolveImportPathToElements(importPath, ctx)
			importedElements = append(importedElements, elements...)
			helpers.SafeDebugLog("[DIAGNOSTICS] TypeScript dynamic import '%s' resolved to elements: %v", importPath, elements)
		}
	}

	return importedElements
}

// parseModuleScriptImports parses <script type="module"> tags for imports using tree-sitter
func parseModuleScriptImports(content string, ctx DiagnosticsContext, doc types.Document) []string {
	var importedElements []string

	// Use tree-sitter parsed script tags for better accuracy
	if doc != nil {
		scriptTags := doc.GetScriptTags()
		for _, scriptTag := range scriptTags {
			if scriptTag.IsModule {
				// Use the parsed imports from tree-sitter
				for _, importStmt := range scriptTag.Imports {
					elements := resolveImportPathToElements(importStmt.ImportPath, ctx)
					importedElements = append(importedElements, elements...)
				}
			}
		}

		// If we found tree-sitter data, return it
		if len(importedElements) > 0 {
			return importedElements
		}
	}

	// Fallback to regex parsing if tree-sitter data is not available
	// Regex to find script type="module" tags and their content (with multiline support)
	scriptRegex := regexp.MustCompile(`(?is)<script[^>]*type\s*=\s*["']module["'][^>]*>(.*?)</script>`)
	matches := scriptRegex.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		if len(match) > 1 {
			scriptContent := match[1]

			// Parse static imports: import 'path' or import from 'path'
			staticImportRegex := regexp.MustCompile(`import\s+(?:.*\s+from\s+)?['"]([^'"]+)['"]`)
			importMatches := staticImportRegex.FindAllStringSubmatch(scriptContent, -1)

			for _, importMatch := range importMatches {
				if len(importMatch) > 1 {
					importPath := importMatch[1]
					elements := resolveImportPathToElements(importPath, ctx)
					importedElements = append(importedElements, elements...)
				}
			}

			// Parse dynamic imports: import('path')
			dynamicImportRegex := regexp.MustCompile(`import\s*\(\s*['"]([^'"]+)['"]\s*\)`)
			dynamicMatches := dynamicImportRegex.FindAllStringSubmatch(scriptContent, -1)

			for _, dynamicMatch := range dynamicMatches {
				if len(dynamicMatch) > 1 {
					importPath := dynamicMatch[1]
					elements := resolveImportPathToElements(importPath, ctx)
					importedElements = append(importedElements, elements...)
				}
			}
		}
	}

	return importedElements
}

// parseNonModuleScriptImports parses regular <script src="..."> tags
func parseNonModuleScriptImports(content string, ctx DiagnosticsContext) []string {
	var importedElements []string

	// Find all script tags with src attribute
	scriptSrcRegex := regexp.MustCompile(`(?i)<script[^>]*src\s*=\s*["']([^"']+)["'][^>]*>`)
	matches := scriptSrcRegex.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		if len(match) > 1 {
			scriptTag := match[0]
			importPath := match[1]

			// Skip if this is a module script (contains type="module")
			if strings.Contains(strings.ToLower(scriptTag), `type="module"`) || strings.Contains(strings.ToLower(scriptTag), `type='module'`) {
				continue
			}

			elements := resolveImportPathToElements(importPath, ctx)
			importedElements = append(importedElements, elements...)
		}
	}

	return importedElements
}

// resolveImportPathToElements resolves an import path to custom element tag names
func resolveImportPathToElements(importPath string, ctx DiagnosticsContext) []string {
	var elements []string

	helpers.SafeDebugLog("[DIAGNOSTICS] Resolving import path '%s' to elements", importPath)

	// Get all available elements and their sources
	allTagNames := ctx.AllTagNames()

	for _, tagName := range allTagNames {
		if elementSource, hasSource := ctx.ElementSource(tagName); hasSource {
			// Check if the import path matches this element's source
			if pathsMatch(importPath, elementSource) {
				elements = append(elements, tagName)
				helpers.SafeDebugLog("[DIAGNOSTICS] Import '%s' provides element '%s'", importPath, tagName)
			}
		}
	}

	return elements
}

// pathsMatch checks if an import path matches an element source path
func pathsMatch(importPath, elementSource string) bool {
	// Direct match first (for exact package imports)
	if importPath == elementSource {
		return true
	}

	// Normalize paths for comparison
	normalizedImport := normalizePath(importPath)
	normalizedSource := normalizePath(elementSource)

	// Direct match on normalized paths
	if normalizedImport == normalizedSource {
		return true
	}

	// Check if import path ends with the element source (relative imports)
	if strings.HasSuffix(importPath, elementSource) {
		return true
	}

	// Check if element source ends with import path (package imports)
	if strings.HasSuffix(elementSource, importPath) {
		return true
	}

	// Extract just the filename and compare
	importFile := filepath.Base(importPath)
	elementFile := filepath.Base(elementSource)
	if importFile == elementFile {
		return true
	}

	return false
}

// normalizePath normalizes a file path for comparison
func normalizePath(path string) string {
	// Remove common prefixes/suffixes
	path = strings.TrimPrefix(path, "./")
	path = strings.TrimPrefix(path, "../")
	path = strings.TrimPrefix(path, "/")

	// Handle npm package paths like @rhds/elements/rh-card/rh-card.js
	// vs manifest paths like ./dist/rh-card.js
	if strings.Contains(path, "/") {
		// Keep the last two segments for better matching
		parts := strings.Split(path, "/")
		if len(parts) >= 2 {
			return strings.Join(parts[len(parts)-2:], "/")
		}
	}

	return path
}

// TagMatch represents a found custom element tag in the document
type TagMatch struct {
	Value string
	Range protocol.Range
}

// findCustomElementTags finds all custom element tag names in the document content
func findCustomElementTags(content string) []TagMatch {
	var matches []TagMatch
	lines := strings.Split(content, "\n")

	for lineIdx, line := range lines {
		// Look for opening tags with hyphens (custom elements)
		idx := 0
		for {
			tagStart := strings.Index(line[idx:], "<")
			if tagStart == -1 {
				break
			}
			tagStart += idx

			// Skip closing tags and comments
			if tagStart+1 < len(line) && (line[tagStart+1] == '/' || line[tagStart+1] == '!') {
				idx = tagStart + 1
				continue
			}

			// Find the end of the tag name (stop at space, /, >, or other delimiters)
			tagEnd := -1
			delimiters := " \t\n\r/>="
			spaceIdx := strings.IndexAny(line[tagStart+1:], delimiters)
			if spaceIdx != -1 {
				tagEnd = tagStart + 1 + spaceIdx
			}

			if tagEnd == -1 {
				idx = tagStart + 1
				continue
			}

			// Extract the tag name
			tagName := line[tagStart+1 : tagEnd]

			// Validate that this looks like a custom element (has a dash and is valid)
			if strings.Contains(tagName, "-") && isValidCustomElementName(tagName) {
				match := TagMatch{
					Value: tagName,
					Range: protocol.Range{
						Start: protocol.Position{
							Line:      uint32(lineIdx),
							Character: uint32(tagStart + 1),
						},
						End: protocol.Position{
							Line:      uint32(lineIdx),
							Character: uint32(tagEnd),
						},
					},
				}
				matches = append(matches, match)
			}

			idx = tagEnd
		}
	}

	return matches
}

// isValidCustomElementName checks if a tag name follows custom element naming rules
func isValidCustomElementName(tagName string) bool {
	// Basic validation - must contain hyphen and not start with invalid characters
	if !strings.Contains(tagName, "-") {
		return false
	}

	// Cannot start with certain reserved patterns
	if strings.HasPrefix(tagName, "xml") || strings.HasPrefix(tagName, "xmlns") {
		return false
	}

	// Must not contain uppercase (CE spec allows case-insensitive but lowercase is convention)
	if strings.ToLower(tagName) != tagName {
		return false
	}

	return true
}

// Helper function for minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
