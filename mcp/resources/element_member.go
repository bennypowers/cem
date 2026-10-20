/*
Copyright © 2026 Benny Powers <web@bennypowers.com>

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
package resources

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/mcp/types"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// makeElementMethodsMemberHandler serves
// cem://element/{tagName}/methods/{memberName}, returning a single class
// method or non-attribute field as JSON. The SDK routes URIs with a fourth
// path segment here because the parent resource's template
// (cem://element/{tagName}/methods) matches only three segments; the generic
// sub-resource filtering in declarative_framework.go cannot serve these URIs
// for the same reason. Methods are searched before fields so a name shared
// by both resolves deterministically.
func makeElementMethodsMemberHandler(registry types.MCPContext) mcp.ResourceHandler {
	return func(_ context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
		uri := req.Params.URI
		parts := strings.Split(strings.TrimPrefix(uri, "cem://element/"), "/")
		if len(parts) != 3 || parts[1] != "methods" || parts[0] == "" || parts[2] == "" {
			return nil, fmt.Errorf("invalid member URI %q", uri)
		}
		tagName, memberName := parts[0], parts[2]

		element, err := registry.ElementInfo(tagName)
		if err != nil {
			return nil, err
		}

		for _, method := range element.Methods() {
			if method.Name == memberName {
				return memberResourceResult(uri, convertMethods([]*manifest.ClassMethod{method})[0])
			}
		}
		for _, field := range element.Fields() {
			if field.Name == memberName {
				return memberResourceResult(uri, convertFields([]*manifest.ClassField{field})[0])
			}
		}
		return nil, fmt.Errorf("element %q has no method or field named %q", tagName, memberName)
	}
}

func memberResourceResult(uri string, member map[string]any) (*mcp.ReadResourceResult, error) {
	data, err := json.MarshalIndent(member, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal member: %w", err)
	}
	return &mcp.ReadResourceResult{
		Contents: []*mcp.ResourceContents{{
			URI:      uri,
			MIMEType: "application/json",
			Text:     string(data),
		}},
	}, nil
}
