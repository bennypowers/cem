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
	"os"
	"path/filepath"
	"strings"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/methods/textDocument"
	"bennypowers.dev/cem/lsp/types"
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

// SourceParser handles parsing source files for precise definition locations
type SourceParser interface {
	// FindClassDeclaration finds the class declaration position
	FindClassDeclaration(filePath string, className string) (*protocol.Range, error)
	// FindTagNameDefinition finds where the tag name is defined (e.g., @customElement decorator)
	FindTagNameDefinition(filePath string, tagName string) (*protocol.Range, error)
	// FindAttributeDeclaration finds the attribute/property declaration
	FindAttributeDeclaration(filePath string, attributeName string) (*protocol.Range, error)
	// FindSlotDefinition finds the slot definition in template
	FindSlotDefinition(filePath string, slotName string) (*protocol.Range, error)
	// FindEventDeclaration finds the event declaration
	FindEventDeclaration(filePath string, eventName string) (*protocol.Range, error)
}

// Definition handles textDocument/definition requests
func Definition(ctx types.DefinitionContext, context *glsp.Context, params *protocol.DefinitionParams) (any, error) {
	uri := params.TextDocument.URI
	helpers.SafeDebugLog("[DEFINITION] Request for URI: %s, Position: line=%d, char=%d", uri, params.Position.Line, params.Position.Character)

	// Get the tracked document
	doc := ctx.GetDocument(uri)
	if doc == nil {
		helpers.SafeDebugLog("[DEFINITION] No document found for URI: %s", uri)
		return nil, nil
	}

	// Analyze what's at the cursor position to determine definition request
	request := analyzeDefinitionTarget(doc, params.Position, ctx.GetRawDocumentManager())
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
	definition, exists := ctx.GetElementDefinition(request.ElementName)
	if !exists {
		helpers.SafeDebugLog("[DEFINITION] No definition found for element: %s", request.ElementName)
		return nil, nil
	}

	// Resolve the source file path with TypeScript preference
	helpers.SafeDebugLog("[DEFINITION] Module path: %s, Source href: %s, Workspace root: %s",
		definition.GetModulePath(), definition.GetSourceHref(), ctx.GetWorkspaceRoot())
	sourceFile := resolveSourcePath(definition, ctx.GetWorkspaceRoot())
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

	// TODO: Add support for slot and event detection in templates
	// This would require more sophisticated parsing of the document content

	return nil
}

// findDefinitionLocation finds the precise location of the definition in the source file
func findDefinitionLocation(sourceFile string, request *DefinitionRequest) *protocol.Location {
	// For now, return basic file location
	// TODO: Implement source parsing for precise locations based on TargetType

	// Convert file:// URI back to file path for processing
	filePath := strings.TrimPrefix(sourceFile, "file://")

	// Basic existence check
	if _, err := os.Stat(filePath); err != nil {
		helpers.SafeDebugLog("[DEFINITION] Source file does not exist: %s", filePath)
		return nil
	}

	// TODO: Implement SourceParser interface and use it here
	// For different target types:
	// - DefinitionTargetTagName: Look for @customElement decorator or customElements.define call
	// - DefinitionTargetClass: Look for class declaration
	// - DefinitionTargetAttribute: Look for @property decorator or field declaration
	// - DefinitionTargetSlot: Look for <slot> tag in template
	// - DefinitionTargetEvent: Look for event declaration in JSDoc or class

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
	modulePath := definition.GetModulePath()
	sourceHref := definition.GetSourceHref()

	// Determine the base path to resolve against
	var basePath string
	if workspaceRoot != "" {
		basePath = workspaceRoot
	} else {
		// Fallback to current working directory
		cwd, err := os.Getwd()
		if err != nil {
			helpers.SafeDebugLog("[DEFINITION] Could not get current working directory: %v", err)
			return ""
		}
		basePath = cwd
	}

	// Prefer source href if available, otherwise use module path
	targetPath := sourceHref
	if targetPath == "" {
		targetPath = modulePath
	}

	if targetPath == "" {
		helpers.SafeDebugLog("[DEFINITION] No source path available")
		return ""
	}

	// Resolve the absolute path
	var resolvedPath string
	if filepath.IsAbs(targetPath) {
		resolvedPath = targetPath
	} else {
		resolvedPath = filepath.Join(basePath, targetPath)
	}

	// Check for TypeScript source preference
	tsPath := preferTypeScriptSource(resolvedPath)
	if tsPath != "" {
		helpers.SafeDebugLog("[DEFINITION] Using TypeScript source: %s", tsPath)
		return "file://" + tsPath
	}

	// Check if the original file exists
	if _, err := os.Stat(resolvedPath); err == nil {
		helpers.SafeDebugLog("[DEFINITION] Using original source: %s", resolvedPath)
		return "file://" + resolvedPath
	}

	// For testing purposes, we still return the path even if file doesn't exist
	// Real editors can handle non-existent files gracefully
	helpers.SafeDebugLog("[DEFINITION] Source file not found but returning path: %s", resolvedPath)
	return "file://" + resolvedPath
}

// preferTypeScriptSource checks if a TypeScript version of the file exists
// Returns the TypeScript path if found, empty string otherwise
func preferTypeScriptSource(jsPath string) string {
	// Only check for TypeScript version if the original is a JavaScript file
	if !strings.HasSuffix(jsPath, ".js") {
		return ""
	}

	// Try replacing .js with .ts
	tsPath := strings.TrimSuffix(jsPath, ".js") + ".ts"
	if _, err := os.Stat(tsPath); err == nil {
		return tsPath
	}

	return ""
}
