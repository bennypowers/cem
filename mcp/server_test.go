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
	"testing"

	W "bennypowers.dev/cem/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewServer(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("./test-fixtures/basic")

	server, err := NewServer(workspace)
	require.NoError(t, err)
	require.NotNil(t, server)

	assert.Equal(t, workspace, server.workspace)
	assert.NotNil(t, server.registry)
	assert.NotNil(t, server.server)
}

func TestServer_GetInfo(t *testing.T) {
	workspace := W.NewFileSystemWorkspaceContext("./test-fixtures/basic")
	server, err := NewServer(workspace)
	require.NoError(t, err)

	info := server.GetInfo()

	assert.Equal(t, "cem", info.Name)
	assert.Equal(t, "1.0.0", info.Version)
	assert.Contains(t, info.Description, "Custom Elements Manifest")
}
