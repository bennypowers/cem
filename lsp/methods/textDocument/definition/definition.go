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
package definition

import (
	"encoding/json"
	"path/filepath"
	"strings"

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/methods/textDocument"
	"bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
	Q "bennypowers.dev/cem/queries"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// DefinitionTargetType specifies what type of definition to look for
type DefinitionTargetType int

const (
	// DefinitionTargetTagName targets the custom element tag name definition (default)
	DefinitionTargetTagName DefinitionTargetType = iota
	// DefinitionTargetClass targets the class declaration
	DefinitionTargetClass
	// DefinitionTargetAttribute targets the specific attribute declaration
	DefinitionTargetAttribute
	// DefinitionTargetSlot targets the specific slot in template
	DefinitionTargetSlot
	// DefinitionTargetEvent targets the event declaration
	DefinitionTargetEvent
)

// DefinitionRequest contains all information needed to resolve a definition
type DefinitionRequest struct {
	Position      protocol.Position
	TargetType    DefinitionTargetType
	ElementName   string
	AttributeName string
	SlotName      string
	EventName     string
}

// Definition handles textDocument/definition requests
func Definition(ctx types.ServerContext, context *glsp.Context, params *protocol.DefinitionParams) (any, error) {
	uri := params.TextDocument.URI

	// Get the tracked document
	doc := ctx.Document(uri)
	if doc == nil {
		return nil, nil
	}

	// Analyze what's at the cursor position to determine definition request
	dm, err := ctx.DocumentManager()
	if err != nil {
		return nil, err
	}
	request := analyzeDefinitionTarget(doc, params.Position, dm)
	if request == nil {
		return nil, nil
	}

	// Only provide definitions for custom elements
	if !textDocument.IsCustomElementTag(request.ElementName) {
		return nil, nil
	}

	// Get the element definition with source information
	definition, exists := ctx.ElementDefinition(request.ElementName)
	if !exists {
		return nil, nil
	}

	// Resolve the source file path with TypeScript preference
	workspaceRoot := ctx.WorkspaceRoot()
	helpers.SafeDebugLog("[DEFINITION] Resolving source path for element '%s', module path '%s', workspace root '%s'", request.ElementName, definition.ModulePath(), workspaceRoot)
	sourceFile := resolveSourcePath(definition, workspaceRoot)
	if sourceFile == "" {
		helpers.SafeDebugLog("[DEFINITION] Source path resolution failed")
		return nil, nil
	}
	helpers.SafeDebugLog("[DEFINITION] Resolved source file: %s", sourceFile)

	// Find the precise location within the source file
	location := findDefinitionLocation(sourceFile, request, ctx)
	if location == nil {
		// Fallback to file start
		location = &protocol.Location{
			URI: sourceFile,
			Range: protocol.Range{
				Start: protocol.Position{Line: 0, Character: 0},
				End:   protocol.Position{Line: 0, Character: 0},
			},
		}
	}

	return *location, nil
}

// analyzeDefinitionTarget analyzes the cursor position to determine what definition is being requested
func analyzeDefinitionTarget(doc types.Document, position protocol.Position, dm any) *DefinitionRequest {
	// Find if we're on an element
	element := doc.FindElementAtPosition(position, dm)
	if element != nil {
		// Check if we're on an attribute
		attr, _ := doc.FindAttributeAtPosition(position, dm)
		if attr != nil {
			// Special handling for slot attributes
			if attr.Name == "slot" && attr.Value != "" {
				return &DefinitionRequest{
					Position:    position,
					TargetType:  DefinitionTargetSlot,
					ElementName: element.TagName,
					SlotName:    attr.Value,
				}
			}

			// Check for event bindings in Lit templates (e.g., @click, @input)
			if strings.HasPrefix(attr.Name, "@") && len(attr.Name) > 1 {
				return &DefinitionRequest{
					Position:    position,
					TargetType:  DefinitionTargetEvent,
					ElementName: element.TagName,
					EventName:   attr.Name[1:], // Remove @ prefix
				}
			}

			// Regular attribute definition
			return &DefinitionRequest{
				Position:      position,
				TargetType:    DefinitionTargetAttribute,
				ElementName:   element.TagName,
				AttributeName: attr.Name,
			}
		}

		// Default to tag name definition for elements
		return &DefinitionRequest{
			Position:    position,
			TargetType:  DefinitionTargetTagName,
			ElementName: element.TagName,
		}
	}

	// No element found at position
	return nil
}

// getQueryManagerFromContext retrieves the query manager from the context,
// falling back to creating a new one if needed (for backwards compatibility)
func getQueryManagerFromContext(ctx types.ServerContext) *Q.QueryManager {
	if qm, err := ctx.QueryManager(); err == nil && qm != nil {
		return qm
	}

	// Fallback: create a new query manager (for testing or edge cases)
	selector := Q.QuerySelector{
		HTML:       []string{"slotsAndParts"},
		TypeScript: []string{"classes", "classMemberDeclaration"},
		CSS:        []string{},
		JSDoc:      []string{},
	}
	qm, err := Q.NewQueryManager(selector)
	if err != nil {
		return nil
	}
	return qm
}

// findDefinitionLocation finds the precise location of the definition in the source file
func findDefinitionLocation(sourceFile string, request *DefinitionRequest, ctx types.ServerContext) *protocol.Location {
	// Try to get cached document content first
	doc := ctx.Document(sourceFile)
	var content []byte

	// Use cached document content if available and valid
	if doc != nil {
		// Add defensive programming to handle potential race conditions
		func() {
			defer func() {
				if r := recover(); r != nil {
					helpers.SafeDebugLog("[DEFINITION] PANIC while getting document content for %s: %v", sourceFile, r)
					doc = nil // Force fallback to disk read
				}
			}()

			docContent, err := doc.Content()
			if err != nil {
				helpers.SafeDebugLog("[DEFINITION] Error getting document content for %s: %v", sourceFile, err)
				doc = nil // Force fallback to disk read
			} else {
				content = []byte(docContent)
				helpers.SafeDebugLog("[DEFINITION] Using cached document content for %s (%d bytes)", sourceFile, len(content))
			}
		}()
	}

	if doc == nil {
		// Fallback to reading from disk if not in document manager
		filePath := strings.TrimPrefix(sourceFile, "file://")
		helpers.SafeDebugLog("[DEFINITION] Document not cached, reading from disk: %s", filePath)
		fs := platform.NewOSFileSystem()

		// Basic existence check
		if _, err := fs.Stat(filePath); err != nil {
			helpers.SafeDebugLog("[DEFINITION] File does not exist on disk: %v", err)
			return nil
		}

		// Read from disk as fallback
		var err error
		content, err = fs.ReadFile(filePath)
		if err != nil {
			helpers.SafeDebugLog("[DEFINITION] Failed to read file from disk: %v", err)
			return nil
		}
		helpers.SafeDebugLog("[DEFINITION] Successfully read %d bytes from disk", len(content))
	}

	// Get the query manager
	queryManager := getQueryManagerFromContext(ctx)
	if queryManager == nil {
		return &protocol.Location{
			URI: sourceFile,
			Range: protocol.Range{
				Start: protocol.Position{Line: 0, Character: 0},
				End:   protocol.Position{Line: 0, Character: 0},
			},
		}
	}

	// Find precise location based on target type
	var targetRange *Q.Range
	var err error

	switch request.TargetType {
	case DefinitionTargetTagName:
		// Look for @customElement decorator
		helpers.SafeDebugLog("[DEFINITION] Looking for tag name definition for '%s' in %d bytes of content", request.ElementName, len(content))
		targetRange, err = Q.FindTagNameDefinitionInSource(content, request.ElementName, queryManager)
		if err != nil {
			helpers.SafeDebugLog("[DEFINITION] Tag name definition search error: %v", err)
		}
		if targetRange == nil {
			helpers.SafeDebugLog("[DEFINITION] Tag name definition not found, trying class declaration")
			// Fallback to class declaration
			targetRange, _ = Q.FindClassDeclarationInSource(content, request.ElementName, queryManager)
			if targetRange == nil {
				helpers.SafeDebugLog("[DEFINITION] Class declaration also not found")
			} else {
				helpers.SafeDebugLog("[DEFINITION] Found class declaration at Line:%d Character:%d", targetRange.Start.Line, targetRange.Start.Character)
			}
		} else {
			helpers.SafeDebugLog("[DEFINITION] Found tag name definition at Line:%d Character:%d", targetRange.Start.Line, targetRange.Start.Character)
		}

	case DefinitionTargetClass:
		// Look for class declaration
		targetRange, _ = Q.FindClassDeclarationInSource(content, request.ElementName, queryManager)

	case DefinitionTargetAttribute:
		// Look for @property decorator or field declaration
		targetRange, _ = Q.FindAttributeDeclarationInSource(content, request.AttributeName, queryManager)

	case DefinitionTargetSlot:
		// Look for <slot name="..."> in template
		targetRange, _ = Q.FindSlotDefinitionInSource(content, request.SlotName, queryManager)

	case DefinitionTargetEvent:
		// Look for event declaration (placeholder for now)
		targetRange = nil

	default:
		// Unknown target type
		targetRange = nil
	}

	// If we found a precise location, use it
	if targetRange != nil {
		helpers.SafeDebugLog("[DEFINITION] Using precise location: Line:%d Character:%d", targetRange.Start.Line, targetRange.Start.Character)
		return &protocol.Location{
			URI: sourceFile,
			Range: protocol.Range{
				Start: protocol.Position{
					Line:      targetRange.Start.Line,
					Character: targetRange.Start.Character,
				},
				End: protocol.Position{
					Line:      targetRange.End.Line,
					Character: targetRange.End.Character,
				},
			},
		}
	}

	// Fallback to file start if precise location not found
	helpers.SafeDebugLog("[DEFINITION] No precise location found, falling back to top of file")
	return &protocol.Location{
		URI: sourceFile,
		Range: protocol.Range{
			Start: protocol.Position{Line: 0, Character: 0},
			End:   protocol.Position{Line: 0, Character: 0},
		},
	}
}

// resolveSourcePath resolves the source file path, preferring TypeScript files over JavaScript
func resolveSourcePath(definition types.ElementDefinition, workspaceRoot string) string {
	// Get the module path from the definition
	modulePath := definition.ModulePath()
	helpers.SafeDebugLog("[DEFINITION] Starting path resolution - modulePath: '%s', workspaceRoot: '%s'", modulePath, workspaceRoot)
	if modulePath == "" {
		return ""
	}

	// Convert to file:// URI if not already
	if !strings.HasPrefix(modulePath, "file://") {
		if strings.HasPrefix(modulePath, "/") {
			// Absolute path
			modulePath = "file://" + modulePath
		} else {
			// Relative path - resolve using package exports if available
			cleanPath := strings.TrimPrefix(modulePath, "./")

			// Check if workspace root is available
			if workspaceRoot == "" {
				return ""
			}

			// Try to resolve using package exports first
			resolvedPath := resolveModulePathWithExports(cleanPath, workspaceRoot)
			if resolvedPath == "" {
				// Fallback to simple concatenation
				resolvedPath = filepath.Join(workspaceRoot, cleanPath)
			}

			modulePath = "file://" + resolvedPath
		}
	}

	// Convert file:// URI to local path
	localPath := strings.TrimPrefix(modulePath, "file://")
	helpers.SafeDebugLog("[DEFINITION] Converted to local path: '%s'", localPath)

	// Try TypeScript source first (.ts), then declaration (.d.ts), then JavaScript (.js)
	fs := platform.NewOSFileSystem()

	if strings.HasSuffix(localPath, ".js") {
		// Try .ts first
		tsPath := strings.TrimSuffix(localPath, ".js") + ".ts"
		helpers.SafeDebugLog("[DEFINITION] Trying TypeScript file: %s", tsPath)
		if _, err := fs.Stat(tsPath); err == nil {
			helpers.SafeDebugLog("[DEFINITION] Found TypeScript file: %s", tsPath)
			return "file://" + tsPath
		} else {
			helpers.SafeDebugLog("[DEFINITION] TypeScript file not found: %v", err)
		}

		// Try .d.ts
		dtsPath := strings.TrimSuffix(localPath, ".js") + ".d.ts"
		helpers.SafeDebugLog("[DEFINITION] Trying declaration file: %s", dtsPath)
		if _, err := fs.Stat(dtsPath); err == nil {
			helpers.SafeDebugLog("[DEFINITION] Found declaration file: %s", dtsPath)
			return "file://" + dtsPath
		} else {
			helpers.SafeDebugLog("[DEFINITION] Declaration file not found: %v", err)
		}
	}

	// Return original path if no alternatives found
	helpers.SafeDebugLog("[DEFINITION] Final resolved path: '%s'", modulePath)
	return modulePath
}

// ResolveSourcePathForTesting is a test-friendly version of resolveSourcePath
// This function is exported for testing purposes only
func ResolveSourcePathForTesting(definition types.ElementDefinition, workspaceRoot string) string {
	return resolveSourcePath(definition, workspaceRoot)
}

// resolveModulePathWithExports resolves a module path using package.json exports field
func resolveModulePathWithExports(modulePath, workspaceRoot string) string {
	// Read package.json from workspace root
	packageJSONPath := filepath.Join(workspaceRoot, "package.json")
	fs := platform.NewOSFileSystem()

	data, err := fs.ReadFile(packageJSONPath)
	if err != nil {
		return ""
	}

	var packageJSON M.PackageJSON
	if err := json.Unmarshal(data, &packageJSON); err != nil {
		return ""
	}

	// Check if package has exports field
	if packageJSON.Exports == nil {
		return ""
	}

	// Type assert the exports field to map[string]any
	exportsMap, ok := packageJSON.Exports.(map[string]any)
	if !ok {
		return ""
	}

	// Try to resolve the module path using exports
	return resolveExportsPath(modulePath, exportsMap, workspaceRoot)
}

// resolveExportsPath resolves a path using the package.json exports field
func resolveExportsPath(modulePath string, exports map[string]any, workspaceRoot string) string {
	// Try exact match first
	if exactValue, exists := exports["./"+modulePath]; exists {
		if exactValueStr, ok := exactValue.(string); ok {
			return filepath.Join(workspaceRoot, strings.TrimPrefix(exactValueStr, "./"))
		}
	}

	// Handle pattern matching
	for exportKey, exportValue := range exports {
		if exportValueStr, ok := exportValue.(string); ok {
			if matchesExportPattern(modulePath, exportKey) {
				// Resolve the pattern
				resolved := resolveExportPattern(modulePath, exportKey, exportValueStr, workspaceRoot)
				if resolved != "" {
					return resolved
				}
			}
		}
	}

	return ""
}

// matchesExportPattern checks if a module path matches an export pattern
func matchesExportPattern(modulePath, exportKey string) bool {
	// Handle wildcard patterns like "./*"
	if strings.HasSuffix(exportKey, "/*") {
		prefix := strings.TrimSuffix(exportKey, "/*")
		// Handle "./*" -> "." -> "" (empty prefix means match everything)
		if prefix == "." {
			prefix = ""
		} else {
			prefix = strings.TrimPrefix(prefix, "./")
		}
		if prefix == "" {
			return true // "./*" matches everything
		}
		return strings.HasPrefix(modulePath, prefix+"/")
	}

	// Handle exact matches
	cleanKey := strings.TrimPrefix(exportKey, "./")
	return cleanKey == modulePath
}

// resolveExportPattern resolves a module path using an export pattern
func resolveExportPattern(modulePath, exportKey, exportValue, workspaceRoot string) string {
	// Handle wildcard patterns like "./*" -> "./elements/*"
	if strings.HasSuffix(exportKey, "/*") && strings.HasSuffix(exportValue, "/*") {
		keyPrefix := strings.TrimSuffix(exportKey, "/*")
		// Handle "./*" -> "." -> "" (empty prefix means match everything)
		if keyPrefix == "." {
			keyPrefix = ""
		} else {
			keyPrefix = strings.TrimPrefix(keyPrefix, "./")
		}

		valuePrefix := strings.TrimSuffix(exportValue, "/*")
		valuePrefix = strings.TrimPrefix(valuePrefix, "./")

		// Remove the key prefix from module path to get the wildcard part
		var wildcardPart string
		if keyPrefix == "" {
			// "./*" pattern - use the full module path as wildcard
			wildcardPart = modulePath
		} else {
			// "./lib/*" pattern - remove the lib/ prefix
			if after, ok := strings.CutPrefix(modulePath, keyPrefix+"/"); ok {
				wildcardPart = after
			} else {
				return ""
			}
		}

		// Construct the resolved path
		var resolvedPath string
		if valuePrefix == "" {
			resolvedPath = wildcardPart
		} else {
			resolvedPath = valuePrefix + "/" + wildcardPart
		}

		return filepath.Join(workspaceRoot, resolvedPath)
	}

	// Handle exact patterns
	if strings.TrimPrefix(exportKey, "./") == modulePath {
		return filepath.Join(workspaceRoot, strings.TrimPrefix(exportValue, "./"))
	}

	return ""
}
