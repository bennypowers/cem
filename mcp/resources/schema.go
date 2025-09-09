package mcp

import (
	"context"
	"encoding/json"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// handleSchemaResource provides the JSON schema for custom elements manifests
func (s *SimpleCEMServer) handleSchemaResource(ctx context.Context, uri string) (*mcp.Resource, error) {
	schema, err := s.registry.GetManifestSchema()
	if err != nil {
		return nil, err
	}

	contents, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		return nil, err
	}

	return &mcp.Resource{
		URI:      uri,
		Name:     "Custom Elements Manifest Schema",
		MimeType: "application/json",
		Text:     string(contents),
	}, nil
}