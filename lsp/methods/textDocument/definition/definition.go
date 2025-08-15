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
	"strings"

	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/methods/textDocument"
	"bennypowers.dev/cem/lsp/types"
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
func Definition(ctx types.DefinitionContext, context *glsp.Context, params *protocol.DefinitionParams) (any, error) {
	uri := params.TextDocument.URI
	helpers.SafeDebugLog("[DEFINITION] Request for URI: %s, Position: line=%d, char=%d", uri, params.Position.Line, params.Position.Character)

	// Get the tracked document
	doc := ctx.Document(uri)
	if doc == nil {
		helpers.SafeDebugLog("[DEFINITION] No document found for URI: %s", uri)
		return nil, nil
	}

	// Analyze what's at the cursor position to determine definition request
	request := analyzeDefinitionTarget(doc, params.Position, ctx.RawDocumentManager())
	if request == nil {
		helpers.SafeDebugLog("[DEFINITION] No definition target found at position")
		return nil, nil
	}

	helpers.SafeDebugLog("[DEFINITION] Definition request: type=%d, element=%s, attr=%s",
		request.TargetType, request.ElementName, request.AttributeName)

	// Only provide definitions for custom elements
	if !textDocument.IsCustomElementTag(request.ElementName) {
		helpers.SafeDebugLog("[DEFINITION] Element %s is not a custom element", request.ElementName)
		return nil, nil
	}

	// Get the element definition with source information
	definition, exists := ctx.ElementDefinition(request.ElementName)
	if !exists {
		helpers.SafeDebugLog("[DEFINITION] No definition found for element: %s", request.ElementName)
		return nil, nil
	}

	// Resolve the source file path with TypeScript preference
	helpers.SafeDebugLog("[DEFINITION] Module path: %s, Source href: %s, Workspace root: %s",
		definition.ModulePath(), definition.SourceHref(), ctx.WorkspaceRoot())
	sourceFile := resolveSourcePath(definition, ctx.WorkspaceRoot())
	if sourceFile == "" {
		helpers.SafeDebugLog("[DEFINITION] Could not resolve source path for element: %s", request.ElementName)
		return nil, nil
	}

	helpers.SafeDebugLog("[DEFINITION] Resolved source file: %s", sourceFile)

	// Find the precise location within the source file
	location := findDefinitionLocation(sourceFile, request)
	if location == nil {
		helpers.SafeDebugLog("[DEFINITION] Could not find precise location in source file")
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
func analyzeDefinitionTarget(doc types.Document, position protocol.Position, dm interface{}) *DefinitionRequest {
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

// Global query manager instance - lazily initialized and reused
var globalQueryManager *Q.QueryManager

// getQueryManager returns the global query manager, creating it if necessary
func getQueryManager() *Q.QueryManager {
	if globalQueryManager == nil {
		// Create a query manager with the queries we need for source parsing
		selector := Q.QuerySelector{
			HTML:       []string{"slotsAndParts"},
			TypeScript: []string{"classes", "classMemberDeclaration"},
			CSS:        []string{},
			JSDoc:      []string{},
		}
		qm, err := Q.NewQueryManager(selector)
		if err != nil {
			helpers.SafeDebugLog("[DEFINITION] Failed to create query manager: %v", err)
			return nil
		}
		globalQueryManager = qm
	}
	return globalQueryManager
}

// findDefinitionLocation finds the precise location of the definition in the source file
func findDefinitionLocation(sourceFile string, request *DefinitionRequest) *protocol.Location {
	// Convert file:// URI back to file path for processing
	filePath := strings.TrimPrefix(sourceFile, "file://")

	// Basic existence check using platform abstraction
	fs := platform.NewOSFileSystem()
	if _, err := fs.Stat(filePath); err != nil {
		helpers.SafeDebugLog("[DEFINITION] Source file does not exist: %s", filePath)
		return nil
	}

	// Get the query manager
	queryManager := getQueryManager()
	if queryManager == nil {
		helpers.SafeDebugLog("[DEFINITION] Query manager not available, falling back to file start")
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

	// Read the source file content using platform abstraction
	content, err := fs.ReadFile(filePath)
	if err != nil {
		helpers.SafeDebugLog("[DEFINITION] Failed to read source file: %v", err)
		return nil
	}

	switch request.TargetType {
	case DefinitionTargetTagName:
		// Look for @customElement decorator
		targetRange, err = Q.FindTagNameDefinitionInSource(content, request.ElementName, queryManager)
		if err != nil || targetRange == nil {
			helpers.SafeDebugLog("[DEFINITION] Tag name definition not found, trying class declaration")
			// Fallback to class declaration
			targetRange, err = Q.FindClassDeclarationInSource(content, request.ElementName, queryManager)
		}

	case DefinitionTargetClass:
		// Look for class declaration
		targetRange, err = Q.FindClassDeclarationInSource(content, request.ElementName, queryManager)

	case DefinitionTargetAttribute:
		// Look for @property decorator or field declaration
		targetRange, err = Q.FindAttributeDeclarationInSource(content, request.AttributeName, queryManager)

	case DefinitionTargetSlot:
		// Look for <slot name="..."> in template
		targetRange, err = Q.FindSlotDefinitionInSource(content, request.SlotName, queryManager)

	case DefinitionTargetEvent:
		// Look for event declaration (placeholder for now)
		helpers.SafeDebugLog("[DEFINITION] Event declaration finding not yet implemented")
		targetRange = nil

	default:
		helpers.SafeDebugLog("[DEFINITION] Unknown target type: %d", request.TargetType)
	}

	if err != nil {
		helpers.SafeDebugLog("[DEFINITION] Error finding precise location: %v", err)
	}

	// If we found a precise location, use it
	if targetRange != nil {
		helpers.SafeDebugLog("[DEFINITION] Found precise location at line %d, char %d",
			targetRange.Start.Line, targetRange.Start.Character)
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
	helpers.SafeDebugLog("[DEFINITION] Precise location not found, falling back to file start")
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
	if modulePath == "" {
		return ""
	}

	// Convert to file:// URI if not already
	if !strings.HasPrefix(modulePath, "file://") {
		if strings.HasPrefix(modulePath, "/") {
			modulePath = "file://" + modulePath
		} else {
			// Relative path - resolve against workspace root
			modulePath = "file://" + workspaceRoot + "/" + modulePath
		}
	}

	// Convert file:// URI to local path
	localPath := strings.TrimPrefix(modulePath, "file://")

	// Try TypeScript source first (.ts), then declaration (.d.ts), then JavaScript (.js)
	fs := platform.NewOSFileSystem()

	if strings.HasSuffix(localPath, ".js") {
		// Try .ts first
		tsPath := strings.TrimSuffix(localPath, ".js") + ".ts"
		if _, err := fs.Stat(tsPath); err == nil {
			return "file://" + tsPath
		}

		// Try .d.ts
		dtsPath := strings.TrimSuffix(localPath, ".js") + ".d.ts"
		if _, err := fs.Stat(dtsPath); err == nil {
			return "file://" + dtsPath
		}
	}

	// Return original path if no alternatives found
	return modulePath
}
