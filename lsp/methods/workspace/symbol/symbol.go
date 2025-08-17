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
package symbol

import (
	"fmt"
	"path/filepath"
	"strings"

	"bennypowers.dev/cem/lsp/helpers"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// SymbolContext provides the dependencies needed for workspace symbol functionality
type SymbolContext interface {
	AllTagNames() []string
	ElementSource(tagName string) (string, bool)
	ElementDescription(tagName string) (string, bool)
	WorkspaceRoot() string
}

// Symbol handles workspace/symbol requests
// Allows users to search for custom elements across the entire workspace
func Symbol(ctx SymbolContext, context *glsp.Context, params *protocol.WorkspaceSymbolParams) ([]protocol.SymbolInformation, error) {
	helpers.SafeDebugLog("[WORKSPACE_SYMBOL] Request for query: '%s'", params.Query)

	var symbols []protocol.SymbolInformation

	// Get all available custom elements
	tagNames := ctx.AllTagNames()
	helpers.SafeDebugLog("[WORKSPACE_SYMBOL] Found %d custom elements", len(tagNames))

	// Filter based on query
	query := strings.ToLower(params.Query)

	for _, tagName := range tagNames {
		// Match logic: empty query returns all, otherwise filter by substring match
		if query == "" || strings.Contains(strings.ToLower(tagName), query) {
			symbol := createSymbolInformation(ctx, tagName)
			if symbol != nil {
				symbols = append(symbols, *symbol)
			}
		}
	}

	helpers.SafeDebugLog("[WORKSPACE_SYMBOL] Returning %d symbols for query '%s'", len(symbols), params.Query)
	return symbols, nil
}

// createSymbolInformation creates a SymbolInformation for a custom element
func createSymbolInformation(ctx SymbolContext, tagName string) *protocol.SymbolInformation {
	// Get source location
	source, hasSource := ctx.ElementSource(tagName)
	if !hasSource {
		// If no source, still create symbol but without location
		return &protocol.SymbolInformation{
			Name: tagName,
			Kind: protocol.SymbolKindClass, // Custom elements are class-based
			Location: protocol.Location{
				URI: "", // No specific location
				Range: protocol.Range{
					Start: protocol.Position{Line: 0, Character: 0},
					End:   protocol.Position{Line: 0, Character: 0},
				},
			},
		}
	}

	// Get description if available
	description, _ := ctx.ElementDescription(tagName)

	// Create file URI from source path
	workspaceRoot := ctx.WorkspaceRoot()
	var uri string

	if filepath.IsAbs(source) {
		uri = "file://" + source
	} else {
		// Relative path - resolve against workspace root
		fullPath := filepath.Join(workspaceRoot, source)
		uri = "file://" + fullPath
	}

	// Create symbol name with description if available
	symbolName := tagName
	if description != "" {
		symbolName = fmt.Sprintf("%s - %s", tagName, description)
	}

	return &protocol.SymbolInformation{
		Name: symbolName,
		Kind: protocol.SymbolKindClass, // Custom elements are class-based
		Location: protocol.Location{
			URI: uri,
			Range: protocol.Range{
				// Point to start of file - could be enhanced to find exact class location
				Start: protocol.Position{Line: 0, Character: 0},
				End:   protocol.Position{Line: 0, Character: 0},
			},
		},
		ContainerName: &source, // Show the file path as container
	}
}
