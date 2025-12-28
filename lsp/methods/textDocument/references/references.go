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
package references

import (
	"io/fs"
	"path/filepath"
	"strings"

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/methods/textDocument"
	"bennypowers.dev/cem/lsp/types"
	Q "bennypowers.dev/cem/queries"
	ignore "github.com/sabhiram/go-gitignore"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// References handles textDocument/references requests
func References(ctx types.ServerContext, context *glsp.Context, params *protocol.ReferenceParams) ([]protocol.Location, error) {
	uri := params.TextDocument.URI
	helpers.SafeDebugLog("[REFERENCES] Request for %s at position %d:%d", uri, params.Position.Line, params.Position.Character)

	// Get the tracked document
	doc := ctx.Document(uri)
	if doc == nil {
		helpers.SafeDebugLog("[REFERENCES] Document not found: %s", uri)
		return []protocol.Location{}, nil
	}

	// Find what element is at the current position
	request := analyzeReferenceRequest(doc, params.Position)
	if request == nil {
		helpers.SafeDebugLog("[REFERENCES] No element found at position %d:%d", params.Position.Line, params.Position.Character)
		content, err := doc.Content()
		if err != nil {
			helpers.SafeDebugLog("[REFERENCES] Error getting document content: %v", err)
			return nil, nil
		}
		preview := content
		if len(content) > 100 {
			preview = content[:100]
		}
		helpers.SafeDebugLog("[REFERENCES] Document content preview: %s", preview)
		return []protocol.Location{}, nil
	}

	helpers.SafeDebugLog("[REFERENCES] Found element: %s", request.ElementName)

	// Search for all references across all documents
	filesystem := ctx.FileSystem()
	return findAllReferences(ctx, request, filesystem), nil
}

// ReferenceRequest contains information about what to search for
type ReferenceRequest struct {
	ElementName string
	URI         string
	Position    protocol.Position
}

// analyzeReferenceRequest determines what element is at the cursor position
func analyzeReferenceRequest(doc types.Document, position protocol.Position) *ReferenceRequest {
	// First try using the existing completion context analysis
	analysis, err := textDocument.AnalyzeCompletionContext(doc, position, "")
	if err != nil {
		helpers.SafeDebugLog("[REFERENCES] Failed to analyze completion context: %v", err)
		return nil
	}
	if analysis != nil {
		helpers.SafeDebugLog("[REFERENCES] Analysis result - Type: %d, TagName: '%s'", analysis.Type, analysis.TagName)

		if analysis.Type == types.CompletionTagName && analysis.TagName != "" {
			helpers.SafeDebugLog("[REFERENCES] Found tag name via completion context: %s", analysis.TagName)
			return &ReferenceRequest{
				ElementName: analysis.TagName,
				URI:         doc.URI(),
				Position:    position,
			}
		}
	} else {
		helpers.SafeDebugLog("[REFERENCES] AnalyzeCompletionContext returned nil")
	}

	// Fallback: use tree-sitter to find element at cursor position directly
	elementName := findElementAtCursor(doc, position)
	if elementName != "" {
		helpers.SafeDebugLog("[REFERENCES] Found tag name via tree-sitter: %s", elementName)
		return &ReferenceRequest{
			ElementName: elementName,
			URI:         doc.URI(),
			Position:    position,
		}
	}

	helpers.SafeDebugLog("[REFERENCES] No element found at position")
	return nil
}

// findElementAtCursor uses tree-sitter to find the element tag name at the cursor position
func findElementAtCursor(doc types.Document, position protocol.Position) string {
	content, err := doc.Content()
	if err != nil {
		helpers.SafeDebugLog("[REFERENCES] Error getting document content: %v", err)
		return ""
	}

	// Convert position to byte offset
	byteOffset := positionToByteOffset([]byte(content), position)
	helpers.SafeDebugLog("[REFERENCES] Looking for element at byte offset %d", byteOffset)

	// Use tree-sitter to parse and find the node at this position
	parser := Q.GetHTMLParser()
	defer Q.PutHTMLParser(parser)

	tree := parser.Parse([]byte(content), nil)
	if tree == nil {
		helpers.SafeDebugLog("[REFERENCES] Failed to parse HTML content")
		return ""
	}
	defer tree.Close()

	// Find the smallest node that contains the cursor position
	node := tree.RootNode().NamedDescendantForByteRange(byteOffset, byteOffset)
	if node == nil {
		helpers.SafeDebugLog("[REFERENCES] No node found at position")
		return ""
	}

	nodeKind := node.GrammarName()
	helpers.SafeDebugLog("[REFERENCES] Found node kind: %s", nodeKind)

	// Walk up the tree to find a tag name
	for node != nil {
		nodeKind := node.GrammarName()
		helpers.SafeDebugLog("[REFERENCES] Checking node kind: %s", nodeKind)

		// Check if this is a tag name node
		if nodeKind == "tag_name" {
			tagName := node.Utf8Text([]byte(content))
			helpers.SafeDebugLog("[REFERENCES] Found tag name: %s", tagName)

			// Check if it's a custom element (contains hyphen)
			if isCustomElement(tagName) {
				return tagName
			}
		}

		// Check if this is an element node, then look for tag_name children
		if nodeKind == "element" || nodeKind == "start_tag" || nodeKind == "end_tag" {
			for i := uint(0); i < node.ChildCount(); i++ {
				child := node.Child(i)
				if child.GrammarName() == "tag_name" {
					tagName := child.Utf8Text([]byte(content))
					helpers.SafeDebugLog("[REFERENCES] Found tag name in child: %s", tagName)

					if isCustomElement(tagName) {
						return tagName
					}
				}
			}
		}

		// Move to parent node
		node = node.Parent()
	}

	helpers.SafeDebugLog("[REFERENCES] No custom element found at cursor")
	return ""
}

// positionToByteOffset converts LSP position to byte offset
func positionToByteOffset(content []byte, position protocol.Position) uint {
	line := uint32(0)
	char := uint32(0)
	offset := uint(0)

	for i, b := range content {
		if line == position.Line && char == position.Character {
			return uint(i)
		}

		if b == '\n' {
			line++
			char = 0
		} else {
			char++
		}

		offset = uint(i + 1)
	}

	return offset
}

// isCustomElement checks if a tag name is a custom element (contains hyphen)
func isCustomElement(tagName string) bool {
	return len(tagName) > 0 && tagName != "" && tagName != "-" && len(tagName) > 1 &&
		tagName[0] != '-' && tagName[len(tagName)-1] != '-' &&
		containsHyphen(tagName)
}

// containsHyphen checks if string contains a hyphen
func containsHyphen(s string) bool {
	for _, ch := range s {
		if ch == '-' {
			return true
		}
	}
	return false
}

// findAllReferences searches for all occurrences of the element across workspace
func findAllReferences(ctx types.ServerContext, request *ReferenceRequest, filesystem platform.FileSystem) []protocol.Location {
	var locations []protocol.Location

	// First, search in all tracked (open) documents
	allDocuments := ctx.AllDocuments()
	helpers.SafeDebugLog("[REFERENCES] Searching %d open documents for element: %s", len(allDocuments), request.ElementName)

	for _, doc := range allDocuments {
		docLocations := findReferencesInDocument(ctx, doc, request.ElementName)
		locations = append(locations, docLocations...)
	}

	// Also search workspace files that aren't currently open
	workspaceRoot := ctx.WorkspaceRoot()
	if workspaceRoot != "" {
		workspaceLocations := findReferencesInWorkspaceWithFS(workspaceRoot, request.ElementName, allDocuments, filesystem)
		locations = append(locations, workspaceLocations...)
	}

	helpers.SafeDebugLog("[REFERENCES] Found %d references for %s in open documents", len(locations), request.ElementName)
	return locations
}

// findReferencesInDocument finds all references to an element in a single document
func findReferencesInDocument(ctx types.ServerContext, doc types.Document, elementName string) []protocol.Location {
	var locations []protocol.Location

	uri := doc.URI()
	helpers.SafeDebugLog("[REFERENCES] Searching document: %s", uri)

	// Get DocumentManager from context for tree-sitter queries
	dm, err := ctx.DocumentManager()
	if err != nil {
		helpers.SafeDebugLog("[REFERENCES] Failed to get DocumentManager: %v", err)
		return locations
	}

	// Use Document's built-in FindCustomElements method to find all custom elements
	elements, err := doc.FindCustomElements(dm)
	if err != nil {
		helpers.SafeDebugLog("[REFERENCES] Failed to find custom elements in %s: %v", uri, err)
		return locations
	}

	// Filter for matching element names and convert to locations
	for _, element := range elements {
		if element.TagName == elementName {
			location := protocol.Location{
				URI:   uri,
				Range: element.Range,
			}
			locations = append(locations, location)
		}
	}

	helpers.SafeDebugLog("[REFERENCES] Found %d references in %s", len(locations), uri)
	return locations
}


// findReferencesInWorkspaceWithFS searches for references using a provided filesystem
// This allows for testability with mock filesystems
func findReferencesInWorkspaceWithFS(workspaceRoot string, elementName string, openDocuments []types.Document, filesystem platform.FileSystem) []protocol.Location {
	var locations []protocol.Location

	// Normalize workspace root - remove trailing slashes for consistent path handling
	workspaceRoot = strings.TrimSuffix(workspaceRoot, "/")

	// Create a set of open document URIs to avoid duplicates
	openFiles := make(map[string]bool)
	for _, doc := range openDocuments {
		openFiles[doc.URI()] = true
	}

	// Load .gitignore if it exists
	var gitignoreMatcher *ignore.GitIgnore
	gitignorePath := filepath.Join(workspaceRoot, ".gitignore")
	if gitignoreContent, err := filesystem.ReadFile(gitignorePath); err == nil {
		// Parse gitignore content directly (works with in-memory filesystems)
		gitignoreMatcher = ignore.CompileIgnoreLines(strings.Split(string(gitignoreContent), "\n")...)
	}

	// Find all relevant files in workspace using fs.WalkDir
	walkErr := fs.WalkDir(filesystem, workspaceRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil // Skip errors and continue
		}

		// Get relative path from workspace root for gitignore matching
		relPath, err := filepath.Rel(workspaceRoot, path)
		if err != nil {
			return nil // Skip if we can't get relative path
		}

		// Skip directories that should be ignored
		if d.IsDir() {
			// Always skip .git directory
			if d.Name() == ".git" {
				return fs.SkipDir
			}

			// Check gitignore for directories
			if gitignoreMatcher != nil && gitignoreMatcher.MatchesPath(relPath+"/") {
				return fs.SkipDir
			}

			return nil
		}

		// Check gitignore for files
		if gitignoreMatcher != nil && gitignoreMatcher.MatchesPath(relPath) {
			return nil
		}

		// Only search HTML, TypeScript, and JavaScript files
		ext := filepath.Ext(path)
		if ext != ".html" && ext != ".htm" && ext != ".ts" && ext != ".js" {
			return nil
		}

		// Create file URI - ensure path is absolute
		absPath := path
		if !filepath.IsAbs(path) {
			// Relative path - join with workspace root to make absolute
			absPath = filepath.Join(workspaceRoot, path)
			// If still relative after joining (e.g., workspace root is "."), prepend / for URI
			if !filepath.IsAbs(absPath) {
				absPath = "/" + absPath
			}
		}
		fileURI := "file://" + filepath.ToSlash(absPath)

		// Skip if this file is already open (already searched)
		if openFiles[fileURI] {
			return nil
		}

		// Search for references in this file
		fileLocations := findReferencesInFileWithFS(path, fileURI, elementName, filesystem)
		locations = append(locations, fileLocations...)

		return nil
	})

	if walkErr != nil {
		helpers.SafeDebugLog("[REFERENCES] Error walking workspace: %v", walkErr)
	}

	helpers.SafeDebugLog("[REFERENCES] Found %d workspace references for %s", len(locations), elementName)
	return locations
}

// findReferencesInFileWithFS searches for references using a provided filesystem
func findReferencesInFileWithFS(filePath string, fileURI string, elementName string, filesystem platform.FileSystem) []protocol.Location {
	var locations []protocol.Location

	// Read file content
	content, err := filesystem.ReadFile(filePath)
	if err != nil {
		helpers.SafeDebugLog("[REFERENCES] Failed to read file %s: %v", filePath, err)
		return locations
	}

	contentBytes := content
	ext := filepath.Ext(filePath)

	// Use tree-sitter to parse the file based on its extension
	switch ext {
	case ".html", ".htm":
		locations = findHTMLReferencesInContent(contentBytes, fileURI, elementName)
	case ".ts", ".js":
		locations = findTypeScriptReferencesInContent(contentBytes, fileURI, elementName)
	}

	helpers.SafeDebugLog("[REFERENCES] Found %d references in file %s", len(locations), filePath)
	return locations
}

// findHTMLReferencesInContent uses tree-sitter to find references in HTML content
func findHTMLReferencesInContent(content []byte, fileURI string, elementName string) []protocol.Location {
	var locations []protocol.Location

	// Get HTML parser
	parser := Q.GetHTMLParser()
	defer Q.PutHTMLParser(parser)

	tree := parser.Parse(content, nil)
	if tree == nil {
		return locations
	}
	defer tree.Close()

	// Create a query manager for HTML custom elements
	qm := getQueryManager()
	if qm == nil {
		return locations
	}

	// Use the existing findHTMLReferencesWithTreeSitter helper
	matcher, err := Q.NewQueryMatcher(qm, "html", "customElements")
	if err != nil {
		helpers.SafeDebugLog("[REFERENCES] Failed to create HTML query matcher: %v", err)
		return locations
	}
	defer matcher.Close()

	for match := range matcher.AllQueryMatches(tree.RootNode(), content) {
		for _, capture := range match.Captures {
			// Get the capture name to filter only start tags
			captureName := matcher.GetCaptureNameByIndex(capture.Index)

			// Only process start tag names, not end tag names
			if captureName != "tag.name" {
				continue
			}

			// Check if this is the element we're looking for
			tagName := string(content[capture.Node.StartByte():capture.Node.EndByte()])
			if tagName == elementName {
				// Convert byte range to position
				startByte := capture.Node.StartByte()
				endByte := capture.Node.EndByte()

				startPos := byteOffsetToPosition(content, startByte)
				endPos := byteOffsetToPosition(content, endByte)

				location := protocol.Location{
					URI: fileURI,
					Range: protocol.Range{
						Start: startPos,
						End:   endPos,
					},
				}
				locations = append(locations, location)
			}
		}
	}

	return locations
}

// findTypeScriptReferencesInContent uses tree-sitter to find references in TypeScript/JS content
func findTypeScriptReferencesInContent(content []byte, fileURI string, elementName string) []protocol.Location {
	var locations []protocol.Location

	// Get TypeScript parser
	parser := Q.RetrieveTypeScriptParser()
	defer Q.PutTypeScriptParser(parser)

	tree := parser.Parse(content, nil)
	if tree == nil {
		return locations
	}
	defer tree.Close()

	// Create a query manager for TypeScript HTML templates
	qm := getQueryManager()
	if qm == nil {
		return locations
	}

	matcher, err := Q.NewQueryMatcher(qm, "typescript", "htmlTemplates")
	if err != nil {
		helpers.SafeDebugLog("[REFERENCES] Failed to create TypeScript query matcher: %v", err)
		return locations
	}
	defer matcher.Close()

	// Find template literals, then search within them for HTML elements
	// Track seen templates to avoid duplicates
	seenTemplates := make(map[string]bool)

	for match := range matcher.AllQueryMatches(tree.RootNode(), content) {
		for _, capture := range match.Captures {
			captureName := matcher.GetCaptureNameByIndex(capture.Index)

			// Only process template literal captures (skip tag/function captures)
			if !strings.HasSuffix(captureName, ".literal") && captureName != "innerHTML.template" {
				continue
			}

			// Get template content
			templateContent := content[capture.Node.StartByte():capture.Node.EndByte()]
			templateOffset := capture.Node.StartByte()

			// Skip if we've already processed this template (same offset)
			templateKey := fileURI + ":" + string(rune(templateOffset))
			if seenTemplates[templateKey] {
				continue
			}
			seenTemplates[templateKey] = true

			// Parse template content as HTML
			htmlLocations := findHTMLReferencesInTemplate(templateContent, templateOffset, fileURI, elementName)
			locations = append(locations, htmlLocations...)
		}
	}

	return locations
}

// findHTMLReferencesInTemplate finds references within a template literal
func findHTMLReferencesInTemplate(templateContent []byte, templateOffset uint, fileURI string, elementName string) []protocol.Location {
	var locations []protocol.Location

	parser := Q.GetHTMLParser()
	defer Q.PutHTMLParser(parser)

	tree := parser.Parse(templateContent, nil)
	if tree == nil {
		return locations
	}
	defer tree.Close()

	qm := getQueryManager()
	if qm == nil {
		return locations
	}

	matcher, err := Q.NewQueryMatcher(qm, "html", "customElements")
	if err != nil {
		return locations
	}
	defer matcher.Close()

	for match := range matcher.AllQueryMatches(tree.RootNode(), templateContent) {
		for _, capture := range match.Captures {
			// Get the capture name to filter only start tags
			captureName := matcher.GetCaptureNameByIndex(capture.Index)

			// Only process start tag names, not end tag names
			if captureName != "tag.name" {
				continue
			}

			tagName := string(templateContent[capture.Node.StartByte():capture.Node.EndByte()])
			if tagName == elementName {
				// Convert template-relative byte range to positions
				templateStartByte := capture.Node.StartByte()
				templateEndByte := capture.Node.EndByte()

				// For simplicity, use positions within the template
				// TODO: Properly adjust for template offset in the main document
				startPos := byteOffsetToPosition(templateContent, templateStartByte+templateOffset)
				endPos := byteOffsetToPosition(templateContent, templateEndByte+templateOffset)

				location := protocol.Location{
					URI: fileURI,
					Range: protocol.Range{
						Start: startPos,
						End:   endPos,
					},
				}
				locations = append(locations, location)
			}
		}
	}

	return locations
}

// getQueryManager returns a query manager for references
var refQueryManager *Q.QueryManager

func getQueryManager() *Q.QueryManager {
	if refQueryManager == nil {
		selector := Q.QuerySelector{
			HTML:       []string{"customElements"},
			TypeScript: []string{"htmlTemplates"},
		}
		var err error
		refQueryManager, err = Q.NewQueryManager(selector)
		if err != nil {
			helpers.SafeDebugLog("[REFERENCES] Failed to create query manager: %v", err)
			return nil
		}
	}
	return refQueryManager
}

// byteOffsetToPosition converts byte offset to line/character position
func byteOffsetToPosition(content []byte, offset uint) protocol.Position {
	line := uint32(0)
	char := uint32(0)

	for i := uint(0); i < offset && i < uint(len(content)); i++ {
		if content[i] == '\n' {
			line++
			char = 0
		} else {
			char++
		}
	}

	return protocol.Position{Line: line, Character: char}
}
