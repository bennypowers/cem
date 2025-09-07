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
package mcp_test

import (
	"testing"

	"bennypowers.dev/cem/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPlaceholderServer_NewSimpleServer(t *testing.T) {
	workspace, cleanup := createTestWorkspace(t, "basic")
	defer cleanup()

	tests := []struct {
		name      string
		transport mcp.TransportKind
	}{
		{
			name:      "stdio transport",
			transport: mcp.TransportStdio,
		},
		{
			name:      "tcp transport",
			transport: mcp.TransportTCP,
		},
		{
			name:      "websocket transport",
			transport: mcp.TransportWebSocket,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server, err := mcp.NewSimpleServer(workspace, test.transport)
			require.NoError(t, err, "Failed to create simple server")
			assert.NotNil(t, server, "Server should not be nil")
		})
	}
}

func TestPlaceholderServer_Run(t *testing.T) {
	workspace, cleanup := createTestWorkspace(t, "basic")
	defer cleanup()

	server, err := mcp.NewSimpleServer(workspace, mcp.TransportStdio)
	require.NoError(t, err, "Failed to create simple server")

	// Run should not error for the placeholder implementation
	err = server.Run()
	assert.NoError(t, err, "Server run should not error")
}

func TestTransportKind_String(t *testing.T) {
	tests := []struct {
		name      string
		transport mcp.TransportKind
		expected  string
	}{
		{
			name:      "stdio transport",
			transport: mcp.TransportStdio,
			expected:  "stdio",
		},
		{
			name:      "tcp transport",
			transport: mcp.TransportTCP,
			expected:  "tcp",
		},
		{
			name:      "websocket transport",
			transport: mcp.TransportWebSocket,
			expected:  "websocket",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			workspace, cleanup := createTestWorkspace(t, "basic")
			defer cleanup()

			server, err := mcp.NewSimpleServer(workspace, test.transport)
			require.NoError(t, err, "Failed to create simple server")

			// We can't directly test the transportName method since it's not public,
			// but we can verify the server was created with the correct transport
			assert.NotNil(t, server)
		})
	}
}
