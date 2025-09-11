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
package resources

import (
	"context"
	"fmt"
	"strings"

	"bennypowers.dev/cem/mcp/types"
	V "bennypowers.dev/cem/validate"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// handleSchemaResource provides the JSON schema for custom elements manifests
func handleSchemaResource(ctx context.Context, req *mcp.ReadResourceRequest, registry types.Registry) (*mcp.ReadResourceResult, error) {
	// Detect schema version from loaded manifests
	schemaVersion := detectSchemaVersion(registry)

	// Get schema using the same method as the validate command
	schemaData, err := V.GetSchema(schemaVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to load schema: %w", err)
	}

	return &mcp.ReadResourceResult{
		Contents: []*mcp.ResourceContents{{
			URI:      req.Params.URI,
			MIMEType: "application/json",
			Text:     string(schemaData),
		}},
	}, nil
}

// detectSchemaVersion extracts schema version from workspace manifests
func detectSchemaVersion(registry types.Registry) string {
	versions := registry.GetManifestSchemaVersions()

	// If no manifests found, use latest stable version as fallback
	if len(versions) == 0 {
		return "2.1.1-speculative"
	}

	// If single version, use it
	if len(versions) == 1 {
		return versions[0]
	}

	// Multiple versions: prefer the highest version, with speculative versions favored
	return selectBestSchemaVersion(versions)
}

// selectBestSchemaVersion chooses the best schema version from multiple options
func selectBestSchemaVersion(versions []string) string {
	// Fallback if empty
	if len(versions) == 0 {
		return "2.1.1-speculative"
	}

	// Simple heuristic: prefer speculative versions, then highest semantic version
	var best string
	var hasSpeculative bool

	for _, version := range versions {
		if version == "" {
			continue
		}

		// Always prefer speculative versions as they are most complete
		if strings.Contains(version, "speculative") {
			if !hasSpeculative || version > best {
				best = version
				hasSpeculative = true
			}
		} else if !hasSpeculative {
			// Among non-speculative versions, take the highest
			if best == "" || version > best {
				best = version
			}
		}
	}

	// If no valid version found, use fallback
	if best == "" {
		return "2.1.1-speculative"
	}

	return best
}
