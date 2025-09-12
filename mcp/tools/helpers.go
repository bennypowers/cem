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
	"encoding/json"
	"fmt"
	"strings"

	"bennypowers.dev/cem/mcp/types"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// ParseToolArgs parses JSON arguments from an MCP request into the specified type
func ParseToolArgs[T any](req *mcp.CallToolRequest) (T, error) {
	var args T
	if req.Params.Arguments != nil {
		if argsData, err := json.Marshal(req.Params.Arguments); err != nil {
			return args, fmt.Errorf("failed to marshal args: %w", err)
		} else if err := json.Unmarshal(argsData, &args); err != nil {
			return args, fmt.Errorf("failed to unmarshal args: %w", err)
		}
	}
	return args, nil
}

// LookupElement looks up an element in the registry and returns appropriate responses
// Returns: (element, errorResponse, error)
// - If element found: (element, nil, nil)
// - If element not found: (nil, errorResponse, nil)
// - If lookup failed: (nil, nil, error)
func LookupElement(registry types.MCPContext, tagName string) (types.ElementInfo, *mcp.CallToolResult, error) {
	element, err := registry.ElementInfo(tagName)
	if err != nil {
		// Element not found - return user-friendly response, not an error
		return nil, &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{
					Text: fmt.Sprintf("Element '%s' not found in workspace", tagName),
				},
			},
		}, nil
	}
	return element, nil, nil
}

// BuildErrorResponse creates a standard error response for MCP tools
func BuildErrorResponse(message string) *mcp.CallToolResult {
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: message,
			},
		},
	}
}

// BuildSuccessResponse creates a standard success response for MCP tools
func BuildSuccessResponse(content string) *mcp.CallToolResult {
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: content,
			},
		},
	}
}

// ResponseBuilder provides a fluent interface for building tool responses
type ResponseBuilder struct {
	builder strings.Builder
}

// NewResponseBuilder creates a new response builder
func NewResponseBuilder() *ResponseBuilder {
	return &ResponseBuilder{}
}

// AddHeader adds a markdown header at the specified level
func (rb *ResponseBuilder) AddHeader(level int, text string) *ResponseBuilder {
	headerMarker := strings.Repeat("#", level)
	rb.builder.WriteString(fmt.Sprintf("%s %s\n\n", headerMarker, text))
	return rb
}

// AddDescription adds element description if available
func (rb *ResponseBuilder) AddDescription(element types.ElementInfo) *ResponseBuilder {
	if element.Description() != "" {
		rb.builder.WriteString(fmt.Sprintf("**Element Description:** %s\n\n", element.Description()))
	}
	return rb
}

// AddSection adds a section with content
func (rb *ResponseBuilder) AddSection(content string) *ResponseBuilder {
	rb.builder.WriteString(content)
	if !strings.HasSuffix(content, "\n") {
		rb.builder.WriteString("\n")
	}
	return rb
}

// AddTemplateSection renders a template and adds it as a section
func (rb *ResponseBuilder) AddTemplateSection(templateName string, data interface{}) *ResponseBuilder {
	content, err := RenderTemplate(templateName, data)
	if err != nil {
		rb.builder.WriteString(fmt.Sprintf("Error rendering template %s: %v\n", templateName, err))
	} else {
		rb.AddSection(content)
	}
	return rb
}

// Build creates the final MCP response
func (rb *ResponseBuilder) Build() *mcp.CallToolResult {
	return BuildSuccessResponse(rb.builder.String())
}

// String returns the current content as a string
func (rb *ResponseBuilder) String() string {
	return rb.builder.String()
}
