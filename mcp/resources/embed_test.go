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
package resources_test

import (
	"testing"

	"bennypowers.dev/cem/mcp"
	"bennypowers.dev/cem/mcp/resources"
	W "bennypowers.dev/cem/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestResourceDefinitions_Embedded is a smoke test to verify resource definitions
// are properly embedded in the binary and can be loaded without filesystem access.
// This is a regression test for the issue where runtime.Caller(0) was used to
// find .md files on disk, which failed for distributed binaries.
// See: https://github.com/anthropics/cem/issues/XXX
func TestResourceDefinitions_Embedded(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("../testdata/fixtures/multiple-elements-integration")
	err := workspace.Init()
	require.NoError(t, err)

	registry, err := mcp.NewMCPContext(workspace)
	require.NoError(t, err)

	err = registry.LoadManifests()
	require.NoError(t, err)

	adapter := mcp.NewMCPContextAdapter(registry)

	resourceDefs, err := resources.Resources(adapter)
	require.NoError(t, err, "Resources() should succeed with embedded definitions")

	// Verify expected number of resources (14 total)
	assert.Len(t, resourceDefs, 14, "Should have exactly 14 embedded resource definitions")

	// Verify expected resource names are present
	resourceNames := make(map[string]bool)
	for _, def := range resourceDefs {
		resourceNames[def.Name] = true
	}

	expectedResources := []string{
		"schema",
		"packages",
		"package",
		"elements",
		"element",
		"element-attributes",
		"element-slots",
		"element-events",
		"element-css-custom-properties",
		"element-css-parts",
		"element-css-states",
		"guidelines",
		"element-guidelines",
		"accessibility",
	}

	for _, expected := range expectedResources {
		assert.True(t, resourceNames[expected], "Should have resource: %s", expected)
	}

	// Verify each resource has required fields populated from embedded .md files
	for _, def := range resourceDefs {
		assert.NotEmpty(t, def.Name, "Resource should have name")
		assert.NotEmpty(t, def.URI, "Resource %s should have URI", def.Name)
		assert.NotEmpty(t, def.MimeType, "Resource %s should have MIME type", def.Name)
		assert.NotEmpty(t, def.Description, "Resource %s should have description", def.Name)
		assert.NotNil(t, def.Handler, "Resource %s should have handler", def.Name)
	}
}
