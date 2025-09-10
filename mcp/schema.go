package mcp

import (
	"context"
	"encoding/json"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// handleSchemaResource provides the JSON schema for custom elements manifests
func (s *Server) handleSchemaResource(ctx context.Context, uri string) (*mcp.ReadResourceResult, error) {
	schema, err := s.registry.GetManifestSchema()
	if err != nil {
		return nil, err
	}

	contents, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		return nil, err
	}

	return &mcp.ReadResourceResult{
		Contents: []*mcp.ResourceContents{{
			URI:      uri,
			MIMEType: "application/json",
			Text:     string(contents),
		}},
	}, nil
}