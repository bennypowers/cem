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
package completion

import (
	"encoding/json"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/methods/textDocument/hover"
	"bennypowers.dev/cem/lsp/types"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// CompletionItemData holds the data needed to resolve a completion item later
type CompletionItemData struct {
	Type          string `json:"type"`          // "tag", "attribute", "attributeValue", "event", "property", "booleanAttribute"
	TagName       string `json:"tagName"`       // Element tag name
	AttributeName string `json:"attributeName"` // Attribute name (for attribute/value completions)
}

// Resolve handles completionItem/resolve requests to lazily generate documentation
func Resolve(ctx types.ServerContext, context *glsp.Context, params *protocol.CompletionItem) (*protocol.CompletionItem, error) {
	helpers.SafeDebugLog("[RESOLVE] Request for completion item: %s", params.Label)

	// If item already has documentation, just return it
	if params.Documentation != nil {
		helpers.SafeDebugLog("[RESOLVE] Item already has documentation, returning as-is")
		return params, nil
	}

	// If no data, we can't resolve anything
	if params.Data == nil {
		helpers.SafeDebugLog("[RESOLVE] No data attached to completion item, cannot resolve")
		return params, nil
	}

	// Parse the data
	var data CompletionItemData
	dataBytes, err := json.Marshal(params.Data)
	if err != nil {
		helpers.SafeDebugLog("[RESOLVE] Failed to marshal data: %v", err)
		return params, nil
	}
	if err := json.Unmarshal(dataBytes, &data); err != nil {
		helpers.SafeDebugLog("[RESOLVE] Failed to unmarshal data: %v", err)
		return params, nil
	}

	helpers.SafeDebugLog("[RESOLVE] Resolved data: type=%s, tagName=%s, attributeName=%s",
		data.Type, data.TagName, data.AttributeName)

	// Generate documentation based on type
	switch data.Type {
	case "tag":
		// Generate element documentation
		if element, exists := ctx.Element(data.TagName); exists {
			params.Documentation = &protocol.MarkupContent{
				Kind:  protocol.MarkupKindMarkdown,
				Value: hover.CreateElementHoverContent(element),
			}
			helpers.SafeDebugLog("[RESOLVE] Generated tag documentation for %s", data.TagName)
		}

	case "attribute", "event", "property", "booleanAttribute":
		// Generate attribute documentation
		if attrs, exists := ctx.Attributes(data.TagName); exists {
			if attr, attrExists := attrs[data.AttributeName]; attrExists {
				params.Documentation = &protocol.MarkupContent{
					Kind:  protocol.MarkupKindMarkdown,
					Value: hover.CreateAttributeHoverContent(attr, data.TagName),
				}
				helpers.SafeDebugLog("[RESOLVE] Generated attribute documentation for %s.%s", data.TagName, data.AttributeName)
			}
		}

	case "attributeValue":
		// For attribute values, we could provide documentation about the specific value
		// For now, we'll just provide the attribute documentation
		if attrs, exists := ctx.Attributes(data.TagName); exists {
			if attr, attrExists := attrs[data.AttributeName]; attrExists {
				params.Documentation = &protocol.MarkupContent{
					Kind:  protocol.MarkupKindMarkdown,
					Value: hover.CreateAttributeHoverContent(attr, data.TagName),
				}
				helpers.SafeDebugLog("[RESOLVE] Generated attribute value documentation for %s.%s", data.TagName, data.AttributeName)
			}
		}

	default:
		helpers.SafeDebugLog("[RESOLVE] Unknown completion item type: %s", data.Type)
	}

	return params, nil
}

// createCompletionData creates the data object for deferred resolution
func createCompletionData(itemType, tagName, attributeName string) map[string]any {
	return map[string]any{
		"type":          itemType,
		"tagName":       tagName,
		"attributeName": attributeName,
	}
}
