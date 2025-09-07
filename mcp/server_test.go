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
package mcp

import (
	"context"
	"testing"

	W "bennypowers.dev/cem/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewMCPServer(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("./test-fixtures/basic")

	server, err := NewMCPServer(workspace)
	require.NoError(t, err)
	require.NotNil(t, server)

	assert.Equal(t, workspace, server.workspace)
	assert.NotNil(t, server.registry)
}

func TestMCPServer_Initialize(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("./test-fixtures/basic")
	server, err := NewMCPServer(workspace)
	require.NoError(t, err)

	ctx := context.Background()
	err = server.Initialize(ctx)
	require.NoError(t, err)

	// Should have loaded manifests
	assert.True(t, server.initialized)
}

func TestMCPServer_GetServerInfo(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("./test-fixtures/basic")
	server, err := NewMCPServer(workspace)
	require.NoError(t, err)

	info := server.GetServerInfo()

	assert.Equal(t, "cem", info.Name)
	assert.Equal(t, "1.0.0", info.Version)
	assert.Contains(t, info.Description, "Custom Elements Manifest")
}

func TestMCPServer_ListResources(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("./test-fixtures/basic")
	server, err := NewMCPServer(workspace)
	require.NoError(t, err)

	ctx := context.Background()
	err = server.Initialize(ctx)
	require.NoError(t, err)

	resources, err := server.ListResources(ctx)
	require.NoError(t, err)

	// Should have basic resources
	assert.Greater(t, len(resources), 0)

	// Check for expected resources
	var foundSchema, foundRegistry bool
	for _, resource := range resources {
		if resource.URI == "cem://schema" {
			foundSchema = true
		}
		if resource.URI == "cem://registry" {
			foundRegistry = true
		}
	}

	assert.True(t, foundSchema, "Should have schema resource")
	assert.True(t, foundRegistry, "Should have registry resource")
}

func TestMCPServer_GetResource(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("./test-fixtures/basic")
	server, err := NewMCPServer(workspace)
	require.NoError(t, err)

	ctx := context.Background()
	err = server.Initialize(ctx)
	require.NoError(t, err)

	// Test schema resource
	resource, err := server.GetResource(ctx, "cem://schema")
	require.NoError(t, err)
	assert.Equal(t, "cem://schema", resource.URI)
	assert.Equal(t, "application/json", resource.MimeType)
	assert.NotEmpty(t, resource.Contents)

	// Test registry resource
	resource, err = server.GetResource(ctx, "cem://registry")
	require.NoError(t, err)
	assert.Equal(t, "cem://registry", resource.URI)
	assert.Equal(t, "application/json", resource.MimeType)
	assert.NotEmpty(t, resource.Contents)
}

func TestMCPServer_ListTools(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("./test-fixtures/basic")
	server, err := NewMCPServer(workspace)
	require.NoError(t, err)

	ctx := context.Background()
	err = server.Initialize(ctx)
	require.NoError(t, err)

	tools, err := server.ListTools(ctx)
	require.NoError(t, err)

	// Should have basic tools
	assert.Greater(t, len(tools), 0)

	// Check for expected tools
	var foundValidate, foundQuery bool
	for _, tool := range tools {
		if tool.Name == "validate_element" {
			foundValidate = true
		}
		if tool.Name == "query_registry" {
			foundQuery = true
		}
	}

	assert.True(t, foundValidate, "Should have validate_element tool")
	assert.True(t, foundQuery, "Should have query_registry tool")
}

func TestMCPServer_CallTool_ValidateElement(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("./test-fixtures/basic")
	server, err := NewMCPServer(workspace)
	require.NoError(t, err)

	ctx := context.Background()
	err = server.Initialize(ctx)
	require.NoError(t, err)

	// Test with valid element
	args := map[string]interface{}{
		"tagName": "my-element",
		"html":    "<my-element></my-element>",
	}

	result, err := server.CallTool(ctx, "validate_element", args)
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestMCPServer_CallTool_QueryRegistry(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("./test-fixtures/basic")
	server, err := NewMCPServer(workspace)
	require.NoError(t, err)

	ctx := context.Background()
	err = server.Initialize(ctx)
	require.NoError(t, err)

	// Test querying all elements
	args := map[string]interface{}{}

	result, err := server.CallTool(ctx, "query_registry", args)
	require.NoError(t, err)
	assert.NotNil(t, result)
}
