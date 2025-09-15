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
package tools

import (
	"context"
	"fmt"

	"bennypowers.dev/cem/mcp/types"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// ElementAttributesArgs represents the arguments for the element_attributes tool
type ElementAttributesArgs struct {
	TagName string `json:"tagName"`
	Context string `json:"context,omitempty"`
}

// handleElementAttributes provides detailed attribute information for a custom element
func handleElementAttributes(
	ctx context.Context,
	req *mcp.CallToolRequest,
	registry types.MCPContext,
) (*mcp.CallToolResult, error) {
	// Parse args from request
	args, err := ParseToolArgs[ElementAttributesArgs](req)
	if err != nil {
		return nil, err
	}

	// Get element information
	element, errorResponse, err := LookupElement(registry, args.TagName)
	if err != nil {
		return nil, err
	}
	if errorResponse != nil {
		return errorResponse, nil
	}

	// Get schema definitions for rich context
	schemaDefinitions, err := getSchemaDefinitions(registry)
	if err != nil {
		return nil, fmt.Errorf("failed to get schema definitions: %w", err)
	}

	// Prepare template data with schema context
	templateData := NewBaseTemplateDataWithSchema(element, args.Context, map[string]string{}, schemaDefinitions)

	// Render the complete response using template
	response, err := RenderTemplate("element_attributes", templateData)
	if err != nil {
		return nil, fmt.Errorf("failed to render element attributes template: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: response,
			},
		},
	}, nil
}