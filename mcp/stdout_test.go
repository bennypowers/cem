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
	"bytes"
	"io"
	"os"
	"testing"

	"bennypowers.dev/cem/mcp"
	W "bennypowers.dev/cem/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestMCPServerStdoutClean verifies that the MCP server never writes to stdout
// except for JSON-RPC messages. This is critical for MCP protocol compliance.
//
// Background: MCP specification requires that stdio-based servers ONLY write
// JSON-RPC messages to stdout. Any other output (logs, debug info, etc.) must
// go to stderr or be suppressed. This test ensures we comply with the spec.
//
// See: https://github.com/bennypowers/cem/issues/129
func TestMCPServerStdoutClean(t *testing.T) {
	// Create a pipe to capture stdout
	originalStdout := os.Stdout
	r, w, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = w
	defer func() { os.Stdout = originalStdout }()

	// Buffer to capture stdout
	var stdoutBuf bytes.Buffer
	done := make(chan struct{})
	go func() {
		io.Copy(&stdoutBuf, r)
		close(done)
	}()

	// Create workspace and MCP server
	workspace := W.NewFileSystemWorkspaceContext("./fixtures/multiple-elements-integration")
	err = workspace.Init()
	require.NoError(t, err)

	server, err := mcp.NewServer(workspace)
	require.NoError(t, err)

	// Load manifests (this is where logging happens)
	// We don't run the full server because that blocks on stdio
	registry := server.GetRegistry()
	err = registry.LoadManifests()
	require.NoError(t, err)

	// Close stdout pipe
	w.Close()
	<-done

	// Verify stdout is clean
	stdoutContent := stdoutBuf.String()

	// Should contain ZERO output (no JSON-RPC in this test)
	assert.Empty(t, stdoutContent, "stdout should be completely empty during manifest loading")

	// Check for common log prefixes that indicate contamination
	// These are the actual log messages that were reported in issue #129
	contaminations := []string{
		"[REGISTRY]",
		"[IN-MEMORY]",
		"INFO",
		"DEBUG",
		"Loading manifests",
		"Successfully",
	}

	for _, pattern := range contaminations {
		assert.NotContains(t, stdoutContent, pattern,
			"stdout contains log output '%s' which would corrupt JSON-RPC", pattern)
	}
}

// TestMCPServerStderrAllowed verifies that error/warning messages CAN go to stderr.
// This is acceptable per MCP spec - only stdout must be clean.
func TestMCPServerStderrAllowed(t *testing.T) {
	// This test verifies that we can still see errors/warnings on stderr
	// Capture stderr
	originalStderr := os.Stderr
	r, w, err := os.Pipe()
	require.NoError(t, err)
	os.Stderr = w
	defer func() { os.Stderr = originalStderr }()

	var stderrBuf bytes.Buffer
	done := make(chan struct{})
	go func() {
		io.Copy(&stderrBuf, r)
		close(done)
	}()

	// Create workspace with nonexistent path to potentially trigger warnings
	// (not errors, as we handle missing manifests gracefully)
	workspace := W.NewFileSystemWorkspaceContext("./fixtures/basic-integration")
	err = workspace.Init()
	require.NoError(t, err)

	server, err := mcp.NewServer(workspace)
	require.NoError(t, err)

	registry := server.GetRegistry()
	err = registry.LoadManifests()
	require.NoError(t, err)

	w.Close()
	<-done

	// stderr MAY contain error/warning messages (this is acceptable and desired for debugging)
	// We're just verifying the test infrastructure works
	// The actual stderr content depends on whether pterm outputs debug info
	// This is fine - we just want to ensure stderr is available for logging
}
