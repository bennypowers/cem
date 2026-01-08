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
package tools_test

import (
	"testing"

	"bennypowers.dev/cem/mcp"
	"bennypowers.dev/cem/mcp/tools"
	W "bennypowers.dev/cem/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestToolDefinitions_Embedded is a smoke test to verify tool definitions
// are properly embedded in the binary and can be loaded without filesystem access.
// This is a regression test for the issue where runtime.Caller(0) was used to
// find .md files on disk, which failed for distributed binaries.
// See: https://github.com/anthropics/cem/issues/XXX
func TestToolDefinitions_Embedded(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("../testdata/fixtures/multiple-elements-integration")
	err := workspace.Init()
	require.NoError(t, err)

	registry, err := mcp.NewMCPContext(workspace)
	require.NoError(t, err)

	err = registry.LoadManifests()
	require.NoError(t, err)

	adapter := mcp.NewMCPContextAdapter(registry)

	toolDefs, err := tools.Tools(adapter)
	require.NoError(t, err, "Tools() should succeed with embedded definitions")

	// Verify expected number of tools
	assert.Len(t, toolDefs, 2, "Should have exactly 2 embedded tool definitions")

	// Verify expected tool names are present
	toolNames := make(map[string]bool)
	for _, def := range toolDefs {
		toolNames[def.Name] = true
	}

	assert.True(t, toolNames["generate_html"], "Should have generate_html tool")
	assert.True(t, toolNames["validate_html"], "Should have validate_html tool")

	// Verify each tool has required fields populated from embedded .md files
	for _, def := range toolDefs {
		assert.NotEmpty(t, def.Name, "Tool should have name")
		assert.NotEmpty(t, def.Description, "Tool %s should have description", def.Name)
		assert.NotNil(t, def.Handler, "Tool %s should have handler", def.Name)
		assert.NotNil(t, def.InputSchema, "Tool %s should have input schema", def.Name)
	}
}
